// Package derivation defines the canonical recipe for turning an L1
// epoch + a list of L2 user txs + an L2 timestamp into a new L2 state.
//
// This is THE protocol spec, made code. Both the sequencer (settler)
// and the independent verifier MUST go through this same function;
// any divergence is a bug in one of them, not a difference of opinion.
//
// Recipe (per L2 block):
//   1. For each L1 TransactionDeposited event in the epoch
//      (l1FromBlock, l1ToBlock], in (blockNumber, logIndex) order:
//        a. mint   = opaqueData[0:32]
//        b. value  = opaqueData[32:64]
//        c. gasLim = opaqueData[64:72]
//        d. isCreation = opaqueData[72]
//        e. data   = opaqueData[73:]
//        f. from = topics[1] (already aliased on L1 if a contract called)
//        g. to   = topics[2]
//        Apply:
//           - anvil_setBalance(from, current + mint)
//           - if !isCreation && to != 0 && (data.length>0 || value>0):
//               impersonate from, send tx to=to, value, data, gas=gasLim
//           - (creation deposits left as a TODO — not used by demo)
//   2. For each user-tx in callData, in callData order:
//        eth_sendRawTransaction(rawTx)
//   3. evm_setNextBlockTimestamp(l2Timestamp); evm_mine
//
// CallData layout (binary, no padding):
//
//   |  l1FromBlock (8 BE)  |  l1ToBlock (8 BE)  |  l2Timestamp (8 BE)  |
//   |  numUserTxs (4 BE)   |
//   |  txLen_1 (4 BE)  |  rawTx_1  |
//   |  txLen_2 (4 BE)  |  rawTx_2  |
//   |  ...                                                            |
//
// The state delta on Rollups commits to (prevRoot, newRoot). The proof
// signs publicInputsHash which covers keccak256(callData), so any
// tampering with the byte stream is caught by the existing
// tmpECDSAVerifier check.
package derivation

import (
	"context"
	"encoding/binary"
	"encoding/hex"
	"fmt"
	"math/big"
	"sort"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/rlp"

	"github.com/koeppelmann/op-rollups-prover/internal/l1"
	"github.com/koeppelmann/op-rollups-prover/internal/l2"
)

const HeaderSize = 8 + 8 + 8 + 4

type Header struct {
	L1FromBlock uint64
	L1ToBlock   uint64
	L2Timestamp uint64
	NumUserTxs  uint32
}

// Encode serializes (header + raw user txs) into the canonical callData
// byte format. rawTxs are RLP-encoded signed transactions exactly as
// they would be sent via eth_sendRawTransaction.
func Encode(h Header, rawTxs [][]byte) []byte {
	out := make([]byte, 0, HeaderSize+len(rawTxs)*200)
	var hdr [HeaderSize]byte
	binary.BigEndian.PutUint64(hdr[0:8], h.L1FromBlock)
	binary.BigEndian.PutUint64(hdr[8:16], h.L1ToBlock)
	binary.BigEndian.PutUint64(hdr[16:24], h.L2Timestamp)
	binary.BigEndian.PutUint32(hdr[24:28], uint32(len(rawTxs)))
	out = append(out, hdr[:]...)
	for _, tx := range rawTxs {
		var l [4]byte
		binary.BigEndian.PutUint32(l[:], uint32(len(tx)))
		out = append(out, l[:]...)
		out = append(out, tx...)
	}
	return out
}

func Decode(b []byte) (Header, [][]byte, error) {
	if len(b) < HeaderSize {
		return Header{}, nil, fmt.Errorf("callData too short: %d < %d", len(b), HeaderSize)
	}
	h := Header{
		L1FromBlock: binary.BigEndian.Uint64(b[0:8]),
		L1ToBlock:   binary.BigEndian.Uint64(b[8:16]),
		L2Timestamp: binary.BigEndian.Uint64(b[16:24]),
		NumUserTxs:  binary.BigEndian.Uint32(b[24:28]),
	}
	off := HeaderSize
	txs := make([][]byte, 0, h.NumUserTxs)
	for i := uint32(0); i < h.NumUserTxs; i++ {
		if off+4 > len(b) {
			return h, nil, fmt.Errorf("truncated tx-len at idx %d", i)
		}
		ln := binary.BigEndian.Uint32(b[off : off+4])
		off += 4
		if off+int(ln) > len(b) {
			return h, nil, fmt.Errorf("truncated tx body at idx %d (need %d, have %d)", i, ln, len(b)-off)
		}
		tx := make([]byte, ln)
		copy(tx, b[off:off+int(ln)])
		off += int(ln)
		txs = append(txs, tx)
	}
	if off != len(b) {
		return h, nil, fmt.Errorf("trailing bytes: %d", len(b)-off)
	}
	return h, txs, nil
}

// FetchDeposits scans L1 in (l1From, l1To] for TransactionDeposited
// events from the configured DepositGateway, decodes them, and returns
// them in the canonical (blockNumber, logIndex) order.
//
// Both settler and verifier call this with the same params. Anyone with
// L1 RPC access gets the same list — that's what makes the recipe
// "deterministic from L1 alone."
func FetchDeposits(ctx context.Context, l1c *l1.Client, depositGw common.Address, l1From, l1To uint64) ([]l1.Deposit, error) {
	logs, err := l1c.RangeDepositEvents(ctx, depositGw, l1From, l1To)
	if err != nil {
		return nil, err
	}
	sort.SliceStable(logs, func(i, j int) bool {
		if logs[i].L1Block != logs[j].L1Block {
			return logs[i].L1Block < logs[j].L1Block
		}
		return logs[i].LogIdx < logs[j].LogIdx
	})
	return logs, nil
}

// ApplyDeposits runs deposit-side effects against an L2 EL: mint to the
// deposit's `from`, and (for non-creation deposits with payload) send
// an impersonated tx with the deposit's data/value. Used identically by
// both the sequencer and the verifier.
//
// Deposit-call txs that revert are still considered "applied" — mint
// is permanent regardless. We log nothing here; callers do.
func ApplyDeposits(ctx context.Context, l2c *l2.Client, deposits []l1.Deposit) error {
	for _, d := range deposits {
		if d.Mint != nil && d.Mint.Sign() > 0 {
			if err := l2c.Mint(ctx, d.From, d.Mint); err != nil {
				return fmt.Errorf("mint to %s: %w", d.From.Hex(), err)
			}
		}
		needsCall := !d.IsCreation && d.To != (common.Address{}) && (len(d.Data) > 0 || (d.Value != nil && d.Value.Sign() > 0))
		if needsCall {
			value := d.Value
			if value == nil {
				value = big.NewInt(0)
			}
			_, _ = l2c.SendDepositCalldata(ctx, d.From, d.To, value, d.GasLimit, d.Data, false)
		}
	}
	return nil
}

// SubmitUserTxs replays each raw RLP tx via eth_sendRawTransaction.
// Used only on the verifier's L2 EL — the sequencer's pool already has
// these txs (they came in over the user-facing RPC).
//
// Caller must have set the EL to --order=fifo so anvil mines them in
// submission order, matching the sequencer's mined order.
func SubmitUserTxs(ctx context.Context, l2c *l2.Client, rawTxs [][]byte) error {
	for i, raw := range rawTxs {
		if _, err := l2c.SendRawTx(ctx, raw); err != nil {
			// A tx that fails admission on the verifier's EL but was
			// happily mined by the sequencer's EL is a divergence —
			// surface it.
			return fmt.Errorf("verifier rejected user tx idx %d: %w", i, err)
		}
	}
	return nil
}

// CaptureBlockTxs returns the EIP-2718 binary form of every tx in the
// given L2 block — the wire format expected by
// eth_sendRawTransaction, in mined order.
//
// Note: tx.MarshalBinary() (NOT rlp.EncodeToBytes(tx)) — for typed
// (EIP-2718) txs the two differ: MarshalBinary returns the canonical
// `type ‖ rlp(payload)`, while rlp.EncodeToBytes wraps that in an
// outer RLP byte-string (the form used inside block bodies). Using
// the wrong one was the original cause of verifier divergence.
func CaptureBlockTxs(ctx context.Context, l2c *l2.Client, blockNum uint64) ([][]byte, error) {
	block, err := l2c.BlockByNumber(ctx, new(big.Int).SetUint64(blockNum))
	if err != nil {
		return nil, fmt.Errorf("get block %d: %w", blockNum, err)
	}
	out := make([][]byte, 0, len(block.Transactions()))
	for _, tx := range block.Transactions() {
		raw, err := tx.MarshalBinary()
		if err != nil {
			return nil, fmt.Errorf("marshal tx %s: %w", tx.Hash().Hex(), err)
		}
		out = append(out, raw)
	}
	return out, nil
}

// FilterImpersonatedDepositCalls splits a list of just-mined L2 txs
// into (depositCalls, userTxs) so the sequencer doesn't double-count
// deposits when packing callData. Heuristic: a tx is a deposit-call
// iff its from address matches one of the just-applied deposits and
// the same (to, data, value) tuple is present in the deposit list.
//
// This is a sequencer-side helper only; the verifier doesn't need it
// because the canonical callData layout already separates
// L1-derived-deposits (encoded as L1 range) from user-supplied txs
// (encoded explicitly). If the sequencer misclassifies, its callData
// won't reproduce on the verifier's L2 EL and the batch is rejected.
func FilterImpersonatedDepositCalls(rawTxs [][]byte, deposits []l1.Deposit) [][]byte {
	if len(deposits) == 0 {
		return rawTxs
	}
	depKey := func(from, to common.Address, value *big.Int, data []byte) string {
		v := big.NewInt(0)
		if value != nil {
			v = value
		}
		return from.Hex() + "|" + to.Hex() + "|" + v.String() + "|" + hex.EncodeToString(data)
	}
	skip := make(map[string]int)
	for _, d := range deposits {
		if d.IsCreation || d.To == (common.Address{}) {
			continue
		}
		if (d.Value == nil || d.Value.Sign() == 0) && len(d.Data) == 0 {
			continue
		}
		v := d.Value
		if v == nil {
			v = big.NewInt(0)
		}
		k := depKey(d.From, d.To, v, d.Data)
		skip[k]++
	}
	out := make([][]byte, 0, len(rawTxs))
	for _, raw := range rawTxs {
		var tx types.Transaction
		if err := rlp.DecodeBytes(raw, &tx); err != nil {
			out = append(out, raw)
			continue
		}
		// Recover sender (works for signed legacy/eip1559/eip2930 txs).
		signer := types.LatestSignerForChainID(tx.ChainId())
		from, err := types.Sender(signer, &tx)
		if err != nil {
			// Impersonated txs from anvil arrive without a real sig —
			// they show up here. Best-effort: skip them all.
			continue
		}
		to := common.Address{}
		if tx.To() != nil {
			to = *tx.To()
		}
		v := tx.Value()
		if v == nil {
			v = big.NewInt(0)
		}
		k := depKey(from, to, v, tx.Data())
		if c, ok := skip[k]; ok && c > 0 {
			skip[k] = c - 1
			continue
		}
		out = append(out, raw)
	}
	return out
}
