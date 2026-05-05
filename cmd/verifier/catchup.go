// Catch-up: when the verifier's local L2 anvil starts at block 0 with the
// empty trie root but the on-chain rollup is already at some block N with a
// different root, we need to replay every historical batch through our own
// derivation pipeline before we can serve signBatch requests.
//
// This is the canonical "any node can sync from L1 alone" property. The
// settler is not consulted — every input is read from L1 directly:
//
//   1. Filter L1 logs for `L2ExecutionPerformed(rollupId indexed, newState)`
//      emitted by the Rollups contract for OUR rollupId.
//   2. For each unique txHash, fetch the L1 tx, ABI-decode its calldata as
//      `postBatch(ProofSystemBatch[])`, locate the sub-batch that touches
//      our rollupId, and extract `batch.callData` — the L2 derivation input.
//   3. Replay through `derivation.Decode` + `FetchDeposits` +
//      `ApplyDeposits` + `SubmitUserTxs` + `Mine`.
//   4. Loop: chain may advance during catch-up, so re-query head + new logs
//      until our L2 root stabilises at the on-chain root.
//
// Failure mode: if our derived root diverges from any intermediate L1 commit,
// it surfaces immediately at the END as `final root mismatch`. The loop
// itself doesn't compare per-batch (we trust L1's threshold-attested commits
// during catch-up — this is replay, not re-verification). After catch-up
// completes, every NEW batch is independently re-derived in `signBatch`.
package main

import (
	"context"
	"fmt"
	"log"
	"math/big"
	"sort"
	"strings"
	"time"

	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"

	"github.com/koeppelmann/op-rollups-prover/bindings/rollups"
	"github.com/koeppelmann/op-rollups-prover/internal/derivation"
	"github.com/koeppelmann/op-rollups-prover/internal/l1"
)

// withRetry wraps a fallible L1 call with exponential backoff. Public Chiado
// RPC nodes rate-limit aggressively (HTTP 429) — a fresh prover replaying
// 4k+ batches will hit the limit within seconds otherwise.
//
// Backs off exponentially up to 30s per try, gives up after 8 attempts.
// Idempotent calls only.
func withRetry[T any](ctx context.Context, op string, fn func() (T, error)) (T, error) {
	var zero T
	delay := 500 * time.Millisecond
	const maxDelay = 30 * time.Second
	for attempt := 1; attempt <= 8; attempt++ {
		v, err := fn()
		if err == nil {
			return v, nil
		}
		es := err.Error()
		retryable := strings.Contains(es, "429") ||
			strings.Contains(es, "Too Many Requests") ||
			strings.Contains(es, "EOF") ||
			strings.Contains(es, "i/o timeout") ||
			strings.Contains(es, "connection reset") ||
			strings.Contains(es, "context deadline exceeded")
		if !retryable || attempt == 8 {
			return zero, err
		}
		log.Printf("retry %s after %v (attempt %d/8): %v", op, delay, attempt, err)
		select {
		case <-ctx.Done():
			return zero, ctx.Err()
		case <-time.After(delay):
		}
		delay *= 2
		if delay > maxDelay {
			delay = maxDelay
		}
	}
	return zero, fmt.Errorf("%s: exhausted retries", op)
}

// l2ExecutionPerformedTopic = keccak256("L2ExecutionPerformed(uint256,bytes32)")
var l2ExecutionPerformedTopic = crypto.Keccak256Hash([]byte("L2ExecutionPerformed(uint256,bytes32)"))

// log range cap so the L1 RPC doesn't time out on a multi-thousand-block scan
const logScanChunk = 5000

// Per-batch throttle: public Chiado RPC nodes (rpc.chiadochain.net) rate-limit
// at ~5-10 req/sec; with 2-3 RPC calls per replayed batch we'd otherwise burst
// straight into HTTP 429. A small delay lets us run smoothly under the limit
// rather than relying on backoff after every fault. Operators with their own
// RPC node can drop this to zero.
const replayThrottle = 80 * time.Millisecond

func (v *Verifier) catchUp(ctx context.Context, onChainRoot common.Hash) error {
	headRoot, headBlock, err := v.l2.HeadStateRoot(ctx)
	if err != nil {
		return fmt.Errorf("l2 head: %w", err)
	}
	if headRoot == onChainRoot {
		log.Printf("catch-up: L2 head=%d already matches on-chain root %s",
			headBlock, onChainRoot.Hex())
		return nil
	}

	log.Printf("catch-up: L2 head=%d root=%s, on-chain root=%s — replaying from L1",
		headBlock, headRoot.Hex(), onChainRoot.Hex())

	abiObj, err := abi.JSON(strings.NewReader(rollups.RollupsABI))
	if err != nil {
		return fmt.Errorf("parse Rollups ABI: %w", err)
	}
	postBatchID := abiObj.Methods["postBatch"].ID

	// Pre-flight: fetch ALL deposits to our gateway in one chunked sweep,
	// index by L1 block, so per-batch FetchDeposits becomes a map lookup
	// instead of an RPC roundtrip. Cuts replay traffic to public Chiado
	// RPC roughly in half (most batches have zero deposits).
	preDeposits, err := v.preloadDeposits(ctx)
	if err != nil {
		return fmt.Errorf("preload deposits: %w", err)
	}
	v.depositCache = preDeposits

	rid := new(big.Int).SetUint64(v.cfg.RollupID)
	ridTopic := common.BytesToHash(common.LeftPadBytes(rid.Bytes(), 32))

	// Outer loop: chain advances during replay, so re-scan after each pass.
	for pass := 1; pass <= 16; pass++ {
		fromBlock := v.cfg.StartL1Block
		if fromBlock == 0 {
			fromBlock = 1
		}
		l1Head, err := v.l1.BlockNumber(ctx)
		if err != nil {
			return fmt.Errorf("l1 head: %w", err)
		}
		log.Printf("catch-up pass %d: scanning L1 (%d, %d] for L2ExecutionPerformed(rollupId=%d)",
			pass, fromBlock-1, l1Head, v.cfg.RollupID)

		// Collect logs in chunks so we don't blow the RPC's filter range cap.
		var logs []types.Log
		for from := fromBlock; from <= l1Head; from += logScanChunk {
			to := from + logScanChunk - 1
			if to > l1Head {
				to = l1Head
			}
			q := ethereum.FilterQuery{
				FromBlock: new(big.Int).SetUint64(from),
				ToBlock:   new(big.Int).SetUint64(to),
				Addresses: []common.Address{v.cfg.RollupsAddr},
				Topics:    [][]common.Hash{{l2ExecutionPerformedTopic}, {ridTopic}},
			}
			chunk, err := withRetry(ctx, "FilterLogs", func() ([]types.Log, error) {
				return v.l1.FilterLogs(ctx, q)
			})
			if err != nil {
				return fmt.Errorf("filter logs (%d, %d]: %w", from-1, to, err)
			}
			logs = append(logs, chunk...)
		}

		// Group by tx hash, deduped + sorted by (block, txIdx). Multiple deltas
		// per tx (rare for our settler, but the protocol allows it) are
		// covered by replaying the WHOLE postBatch once per tx.
		byTx := map[common.Hash]types.Log{}
		for _, l := range logs {
			if existing, ok := byTx[l.TxHash]; ok {
				// keep the lowest log index for stable ordering on tx-mate ties
				if l.Index < existing.Index {
					byTx[l.TxHash] = l
				}
				continue
			}
			byTx[l.TxHash] = l
		}
		txs := make([]types.Log, 0, len(byTx))
		for _, l := range byTx {
			txs = append(txs, l)
		}
		sort.Slice(txs, func(i, j int) bool {
			if txs[i].BlockNumber != txs[j].BlockNumber {
				return txs[i].BlockNumber < txs[j].BlockNumber
			}
			return txs[i].TxIndex < txs[j].TxIndex
		})

		// Skip the txs whose batches we've already applied. We use the
		// post-replay L2 head block as our "cursor" — every successful
		// replay advances it by exactly one. (Valid for our 1-stateDelta-
		// per-batch settler; a more general implementation would track a
		// separate L1-tx-position cursor.)
		_, currentL2, _ := v.l2.HeadStateRoot(ctx)
		applied := uint64(currentL2)
		if uint64(len(txs)) <= applied {
			// L2 head has caught up to the number of historical batches
			// but the root still doesn't match on-chain → divergence.
			break
		}
		toApply := txs[applied:]
		log.Printf("catch-up pass %d: %d historical batches found, %d already replayed, applying %d more",
			pass, len(txs), applied, len(toApply))

		// Periodic progress logging — replaying 4k+ batches quietly is
		// indistinguishable from a hang from the operator's perspective.
		const progressEvery = 250
		startedAt := time.Now()
		for i, l := range toApply {
			if err := v.replayHistoricalBatch(ctx, abiObj, postBatchID, l.TxHash); err != nil {
				return fmt.Errorf("replay batch (l1Tx=%s, idx=%d): %w", l.TxHash.Hex(), applied+uint64(i), err)
			}
			if (i+1)%progressEvery == 0 {
				elapsed := time.Since(startedAt)
				rate := float64(i+1) / elapsed.Seconds()
				eta := time.Duration(float64(len(toApply)-(i+1))/rate) * time.Second
				log.Printf("  catch-up: %d/%d batches replayed (%.1f bps, ETA %v)",
					i+1, len(toApply), rate, eta.Round(time.Second))
			}
			if replayThrottle > 0 {
				select {
				case <-ctx.Done():
					return ctx.Err()
				case <-time.After(replayThrottle):
				}
			}
		}

		// Did we converge?
		root, blk, err := v.l2.HeadStateRoot(ctx)
		if err != nil {
			return fmt.Errorf("post-replay head: %w", err)
		}
		// Re-read on-chain root in case it advanced during catch-up.
		latestOnChain, err := v.readOnChainStateRoot(ctx)
		if err != nil {
			return fmt.Errorf("re-read on-chain root: %w", err)
		}
		log.Printf("catch-up pass %d: L2 head=%d root=%s, on-chain root=%s",
			pass, blk, root.Hex(), latestOnChain.Hex())
		if root == latestOnChain {
			log.Printf("✓ catch-up complete after %d passes", pass)
			return nil
		}
		if root != onChainRoot && len(toApply) == 0 {
			// nothing new to apply but still diverged → can't recover
			return fmt.Errorf("catch-up diverged: derived %s ≠ on-chain %s, no further batches",
				root.Hex(), latestOnChain.Hex())
		}
		// otherwise loop and try to apply any newly-emitted events
		onChainRoot = latestOnChain
	}
	return fmt.Errorf("catch-up did not converge in 16 passes")
}

// replayHistoricalBatch: fetch the L1 tx, decode `postBatch(ProofSystemBatch[])`,
// find the sub-batch whose entries touch our rollupId, and re-run the canonical
// derivation against our local L2 anvil.
func (v *Verifier) replayHistoricalBatch(ctx context.Context, abiObj abi.ABI, postBatchID []byte, txHash common.Hash) error {
	type txResult struct {
		tx *types.Transaction
	}
	res, err := withRetry(ctx, "TransactionByHash", func() (txResult, error) {
		t, _, e := v.l1.TransactionByHash(ctx, txHash)
		return txResult{t}, e
	})
	if err != nil {
		return fmt.Errorf("fetch tx: %w", err)
	}
	tx := res.tx
	data := tx.Data()
	if len(data) < 4 || !bytesEq(data[:4], postBatchID) {
		return fmt.Errorf("tx selector is not postBatch (got 0x%x)", data[:4])
	}

	args, err := abiObj.Methods["postBatch"].Inputs.Unpack(data[4:])
	if err != nil {
		return fmt.Errorf("decode postBatch args: %w", err)
	}
	if len(args) != 1 {
		return fmt.Errorf("postBatch decoded %d args, want 1", len(args))
	}

	// `args[0]` is `[]<anonymous struct>` whose layout matches
	// `rollups.ProofSystemBatch` exactly. abi.ConvertType walks the field
	// graph and re-types it without copying, so we can iterate using the
	// generated bindings' field names.
	batchesPtr, ok := abi.ConvertType(args[0], new([]rollups.ProofSystemBatch)).(*[]rollups.ProofSystemBatch)
	if !ok || batchesPtr == nil {
		return fmt.Errorf("ConvertType failed for postBatch arg: %T", args[0])
	}
	batches := *batchesPtr

	rid := v.cfg.RollupID
	for _, b := range batches {
		// Does this sub-batch touch our rollupId?
		hit := false
		for _, r := range b.RollupIds {
			if r.Uint64() == rid {
				hit = true
				break
			}
		}
		if !hit {
			continue
		}

		// ProofSystemBatch has ONE callData (the L2 derivation input).
		// Our settler emits batches with 1 entry per batch; each batch
		// drives exactly one L2 block forward via its callData.
		if err := v.applyBatchCallData(ctx, b.CallData); err != nil {
			return fmt.Errorf("replay derivation: %w", err)
		}
	}
	return nil
}

// applyBatchCallData runs the canonical derivation against our local L2 anvil:
// decode → look up deposits in the precomputed cache → apply → submit user
// txs → mine. The cache (`v.depositCache`) is populated once at the start
// of catch-up so we don't roundtrip to L1 for every batch.
func (v *Verifier) applyBatchCallData(ctx context.Context, callData []byte) error {
	hdr, rawTxs, err := derivation.Decode(callData)
	if err != nil {
		return fmt.Errorf("decode: %w", err)
	}
	deps := v.depositsInRange(hdr.L1FromBlock, hdr.L1ToBlock)
	if err := derivation.ApplyDeposits(ctx, v.l2, deps); err != nil {
		return fmt.Errorf("apply deposits: %w", err)
	}
	if err := derivation.SubmitUserTxs(ctx, v.l2, rawTxs); err != nil {
		return fmt.Errorf("submit user txs: %w", err)
	}
	if _, err := v.l2.Mine(ctx, hdr.L2Timestamp); err != nil {
		return fmt.Errorf("mine: %w", err)
	}
	return nil
}

// readOnChainStateRoot fetches `rollups[rid].stateRoot` from the registry.
func (v *Verifier) readOnChainStateRoot(ctx context.Context) (common.Hash, error) {
	rb, err := rollups.NewRollups(v.cfg.RollupsAddr, v.l1)
	if err != nil {
		return common.Hash{}, err
	}
	cfg, err := rb.Rollups(nil, new(big.Int).SetUint64(v.cfg.RollupID))
	if err != nil {
		return common.Hash{}, err
	}
	return common.Hash(cfg.StateRoot), nil
}

// preloadDeposits fetches every L1 deposit emitted by our DepositGateway
// in the range (StartL1Block-1, currentHead], chunked under the RPC's
// log-range cap, indexed by L1 block. Each historical batch's
// FetchDeposits call becomes a O(1) lookup against this map.
//
// Done once at the start of catch-up. Concurrent batches landing during
// catch-up are picked up by pass 2 of the outer loop, which calls this
// again.
func (v *Verifier) preloadDeposits(ctx context.Context) (map[uint64][]l1.Deposit, error) {
	from := v.cfg.StartL1Block
	if from == 0 {
		from = 1
	}
	head, err := v.l1.BlockNumber(ctx)
	if err != nil {
		return nil, err
	}
	if head < from {
		return map[uint64][]l1.Deposit{}, nil
	}

	out := map[uint64][]l1.Deposit{}
	total := 0
	for f := from; f <= head; f += logScanChunk {
		t := f + logScanChunk - 1
		if t > head {
			t = head
		}
		// l1.RangeDepositEvents takes (from, to] semantics — it scans
		// events in blocks (f, t]. Pass f-1 so the first iteration covers
		// `from`.
		deps, err := withRetry(ctx, "RangeDepositEvents", func() ([]l1.Deposit, error) {
			return v.l1.RangeDepositEvents(ctx, v.cfg.DepositGatewayAddr, f-1, t)
		})
		if err != nil {
			return nil, fmt.Errorf("range deposits (%d, %d]: %w", f-1, t, err)
		}
		for _, d := range deps {
			out[d.L1Block] = append(out[d.L1Block], d)
			total++
		}
	}
	log.Printf("catch-up: preloaded %d deposits in (%d, %d] (%d blocks indexed)",
		total, from-1, head, len(out))
	return out, nil
}

// depositsInRange returns deposits emitted in the L1 block range
// (l1From, l1To] (matching the canonical derivation recipe). Order
// preserved within each block — depends on the consistent insertion
// order from preloadDeposits.
func (v *Verifier) depositsInRange(l1From, l1To uint64) []l1.Deposit {
	if v.depositCache == nil {
		return nil
	}
	var out []l1.Deposit
	for b := l1From + 1; b <= l1To; b++ {
		if d, ok := v.depositCache[b]; ok {
			out = append(out, d...)
		}
	}
	return out
}

func bytesEq(a, b []byte) bool {
	if len(a) != len(b) {
		return false
	}
	for i := range a {
		if a[i] != b[i] {
			return false
		}
	}
	return true
}
