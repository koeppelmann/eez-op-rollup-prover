// op-verifier: independent re-deriver. Holds the proof-signing key.
//
// Receives signBatch(rollupId, prevRoot, claimedRoot, callData,
// l1ParentHash, l1Timestamp, l1NextNumber, blobCount, publicHash) over
// JSON-RPC. Decodes callData, replays the canonical derivation against
// its OWN L2 EL, compares the resulting state root against the
// settler's claimed root.
//   - match  → signs publicHash with its key; returns the signature
//   - mismatch → returns a JSON-RPC error and the candidate batch
//                cannot land on L1 (Rollups.tmpECDSAVerifier rejects
//                an unsigned-by-this-key proof).
//
// Production deployment runs N of these with different keys behind a
// multisig verifier on L1 — see README. For the demo we run one,
// and we still split the key from the settler so the demo
// substantively shows the model.
package main

import (
	"context"
	"crypto/ecdsa"
	"encoding/json"
	"flag"
	"fmt"
	"io"
	"log"
	"math/big"
	"net/http"
	"os"
	"strings"
	"sync"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"

	"github.com/koeppelmann/op-rollups-prover/bindings/rollupmgr"
	"github.com/koeppelmann/op-rollups-prover/bindings/rollups"
	"github.com/koeppelmann/op-rollups-prover/internal/config"
	"github.com/koeppelmann/op-rollups-prover/internal/derivation"
	"github.com/koeppelmann/op-rollups-prover/internal/l1"
	"github.com/koeppelmann/op-rollups-prover/internal/l2"
	"github.com/koeppelmann/op-rollups-prover/internal/proof"
)


func main() {
	cfg := config.LoadFromFlags()
	listenAddr := flag.String("listen", env("VERIFIER_LISTEN", ":9100"), "verifier RPC listen addr")
	flag.Parse()

	log.SetFlags(log.LstdFlags | log.Lmicroseconds)
	log.Printf("op-verifier starting | l1=%s l2=%s rollupId=%d listen=%s",
		cfg.L1RPC, cfg.L2RPC, cfg.RollupID, *listenAddr)

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	l1c, err := l1.New(ctx, cfg)
	if err != nil {
		log.Fatalf("l1: %v", err)
	}
	defer l1c.Close()
	l2c, err := l2.New(ctx, cfg)
	if err != nil {
		log.Fatalf("l2: %v", err)
	}
	defer l2c.Close()
	if err := l2c.SetAutoMine(ctx, false); err != nil {
		log.Printf("warn: SetAutoMine off: %v", err)
	}
	if err := l2c.SetIntervalMining(ctx, 0); err != nil {
		log.Printf("warn: SetIntervalMining 0: %v", err)
	}

	if cfg.SignerKeyHex == "" {
		log.Fatalf("no signing key — set $PROVER_KEY (or pass -key 0x...). " +
			"This binary refuses to fall back to a default key by design.")
	}
	pk, err := crypto.HexToECDSA(stripHex(cfg.SignerKeyHex))
	if err != nil {
		log.Fatalf("invalid signer key: %v", err)
	}
	signerAddr := crypto.PubkeyToAddress(pk.PublicKey)
	log.Printf("verifier signer: %s", signerAddr.Hex())

	rb, err := rollups.NewRollups(cfg.RollupsAddr, l1c)
	if err != nil {
		log.Fatalf("bind rollups: %v", err)
	}
	rcfg, err := rb.Rollups(&bind.CallOpts{Context: ctx}, new(big.Int).SetUint64(cfg.RollupID))
	if err != nil {
		log.Fatalf("read rollup config: %v", err)
	}
	mgr, err := rollupmgr.NewRollupmgr(cfg.RollupMgrAddr, l1c)
	if err != nil {
		log.Fatalf("bind rollup manager: %v", err)
	}
	vk, err := mgr.VerificationKey(&bind.CallOpts{Context: ctx}, cfg.ProofSystemAddr)
	if err != nil {
		log.Fatalf("read vk(manager=%s, ps=%s): %v",
			cfg.RollupMgrAddr.Hex(), cfg.ProofSystemAddr.Hex(), err)
	}
	if (vk == [32]byte{}) {
		log.Fatalf("manager %s reports zero vkey for PS %s — PS not registered",
			cfg.RollupMgrAddr.Hex(), cfg.ProofSystemAddr.Hex())
	}
	log.Printf("rollup %d on L1: rollupContract=%s stateRoot=%s | PS=%s vk=%s",
		cfg.RollupID, rcfg.RollupContract.Hex(), common.Hash(rcfg.StateRoot).Hex(),
		cfg.ProofSystemAddr.Hex(), common.Hash(vk).Hex())

	// L2 verifier head must agree with L1 as we sign batches forward.
	hr, lh, err := l2c.HeadStateRoot(ctx)
	if err != nil {
		log.Fatalf("l2 head: %v", err)
	}
	log.Printf("l2 verifier head: block=%d root=%s", lh, hr.Hex())

	v := &Verifier{
		cfg: cfg, l1: l1c, l2: l2c, vk: vk,
		signerKey: pk, signerAddr: signerAddr,
		verifiedBatches: make(map[string]bool),
	}

	// Catch-up: walk L1 for every historical postBatch involving our
	// rollupId, replay the canonical derivation against our local anvil
	// until L2 head root matches the on-chain commit. This is the
	// "any node can sync from L1 alone" property — no settler trust.
	if err := v.catchUp(ctx, common.Hash(rcfg.StateRoot)); err != nil {
		log.Fatalf("catch-up failed: %v", err)
	}

	mux := http.NewServeMux()
	mux.HandleFunc("/", v.handle)
	mux.HandleFunc("/status", v.statusHandler)
	srv := &http.Server{Addr: *listenAddr, Handler: mux}
	log.Printf("listening on %s", *listenAddr)
	if err := srv.ListenAndServe(); err != nil {
		log.Fatalf("listen: %v", err)
	}
}

type Verifier struct {
	cfg *config.Config
	l1  *l1.Client
	l2  *l2.Client
	vk  [32]byte

	signerKey  *ecdsa.PrivateKey
	signerAddr common.Address

	mu       sync.Mutex
	signed   uint64
	rejected uint64
	lastErr  string

	// Batches we have already independently re-derived and verified. Key
	// is hex(prev|claimed|keccak256(callData)). When a settler has to
	// resubmit a batch on real Chiado (because publicInputsHash binds to
	// L1 block context that was wrong on the first attempt), it asks us
	// to sign a fresh digest covering the same (prev, claimed, callData)
	// but a new (parentHash, timestamp). Without this cache we'd refuse
	// the resign because our anvil already moved past `prev` — re-doing
	// the derivation from a stale head produces nonsense.
	verifiedBatches map[string]bool

	// L2 anvil snapshot id taken just BEFORE mining the most recent
	// batch. If the settler tells us via verifier_rollbackBatch that
	// the post-derive state never landed on L1, we revert to this
	// snapshot so our anvil stays in lockstep with what L1 actually
	// committed. Empty string = no snapshot held (e.g. after a
	// successful commit).
	preMineSnapshot string
	preMineKey      string

	// depositCache is populated once at the start of catchUp() with every
	// L1 deposit emitted by our gateway in the relevant range, indexed by
	// L1 block number. Per-batch lookups become a map probe instead of
	// an L1 filter-logs roundtrip (which is critical against rate-limited
	// public RPCs). Nil outside catch-up.
	depositCache map[uint64][]l1.Deposit
}

type rollbackReq struct {
	PrevRoot     string `json:"prevRoot"`
	ClaimedRoot  string `json:"claimedRoot"`
	CallDataHash string `json:"callDataHash"`
}

type signReq struct {
	RollupID     uint64 `json:"rollupId"`
	PrevRoot     string `json:"prevRoot"`
	ClaimedRoot  string `json:"claimedRoot"`
	CallData     string `json:"callData"`
	L1ParentHash string `json:"l1ParentHash"`
	L1Timestamp  uint64 `json:"l1Timestamp"`
	L1NextNumber uint64 `json:"l1NextNumber"`
	BlobCount    uint64 `json:"blobCount"`
	PublicHash   string `json:"publicHash"`
}

type rpcReq struct {
	JSONRPC string          `json:"jsonrpc"`
	Method  string          `json:"method"`
	Params  json.RawMessage `json:"params"`
	ID      json.RawMessage `json:"id"`
}
type rpcResp struct {
	JSONRPC string          `json:"jsonrpc"`
	Result  any             `json:"result,omitempty"`
	Error   *rpcErr         `json:"error,omitempty"`
	ID      json.RawMessage `json:"id"`
}
type rpcErr struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

func (v *Verifier) handle(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "POST only", http.StatusMethodNotAllowed)
		return
	}
	body, err := io.ReadAll(r.Body)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	var req rpcReq
	if err := json.Unmarshal(body, &req); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	enc := json.NewEncoder(w)

	if req.Method == "verifier_rollbackBatch" {
		// Settler tells us a batch we signed never landed on L1 — roll
		// our anvil back to the pre-mine snapshot we took for it. Without
		// this our anvil drifts ahead of L1 and the next signBatch
		// refuses with "verifier out of sync".
		var arr []rollbackReq
		if err := json.Unmarshal(req.Params, &arr); err != nil || len(arr) != 1 {
			_ = enc.Encode(rpcResp{JSONRPC: "2.0", ID: req.ID,
				Error: &rpcErr{Code: -32602, Message: "expected [rollbackReq]"}})
			return
		}
		if err := v.rollbackBatch(r.Context(), &arr[0]); err != nil {
			log.Printf("ROLLBACK refused: %v", err)
			_ = enc.Encode(rpcResp{JSONRPC: "2.0", ID: req.ID,
				Error: &rpcErr{Code: -32000, Message: err.Error()}})
			return
		}
		_ = enc.Encode(rpcResp{JSONRPC: "2.0", ID: req.ID, Result: "ok"})
		return
	}

	if req.Method != "verifier_signBatch" {
		_ = enc.Encode(rpcResp{JSONRPC: "2.0", ID: req.ID,
			Error: &rpcErr{Code: -32601, Message: "unknown method " + req.Method}})
		return
	}

	// Single-arg call: params is `[req]`.
	var arr []signReq
	if err := json.Unmarshal(req.Params, &arr); err != nil || len(arr) != 1 {
		_ = enc.Encode(rpcResp{JSONRPC: "2.0", ID: req.ID,
			Error: &rpcErr{Code: -32602, Message: "expected [signReq]"}})
		return
	}
	sig, err := v.signBatch(r.Context(), &arr[0])
	if err != nil {
		v.mu.Lock()
		v.rejected++
		v.lastErr = err.Error()
		v.mu.Unlock()
		log.Printf("REFUSE: %v", err)
		_ = enc.Encode(rpcResp{JSONRPC: "2.0", ID: req.ID,
			Error: &rpcErr{Code: -32000, Message: err.Error()}})
		return
	}
	v.mu.Lock()
	v.signed++
	v.mu.Unlock()
	_ = enc.Encode(rpcResp{JSONRPC: "2.0", ID: req.ID,
		Result: map[string]string{"signature": "0x" + hexEnc(sig)}})
}

func (v *Verifier) signBatch(ctx context.Context, req *signReq) ([]byte, error) {
	if req.RollupID != v.cfg.RollupID {
		return nil, fmt.Errorf("rollupId mismatch: req=%d, configured=%d", req.RollupID, v.cfg.RollupID)
	}
	prev := common.HexToHash(req.PrevRoot)
	claimed := common.HexToHash(req.ClaimedRoot)
	callData := common.FromHex(req.CallData)

	cdHash := crypto.Keccak256Hash(callData)
	batchKey := prev.Hex() + claimed.Hex() + cdHash.Hex()
	v.mu.Lock()
	alreadyVerified := v.verifiedBatches[batchKey]
	v.mu.Unlock()

	// Decode callData (cheap, also a sanity check on the resign path).
	hdr, rawTxs, err := derivation.Decode(callData)
	if err != nil {
		return nil, fmt.Errorf("decode callData: %w", err)
	}

	if !alreadyVerified {
		// Fresh batch: re-derive on our L2 EL and confirm the claimed
		// root. Our head must currently be at `prev` (otherwise the
		// settler has skipped a batch or we missed one).
		headRoot, headNum, err := v.l2.HeadStateRoot(ctx)
		if err != nil {
			return nil, fmt.Errorf("l2 head: %w", err)
		}
		if headRoot != prev {
			return nil, fmt.Errorf("l2 head root %s ≠ batch prev %s — verifier is out of sync (head block %d)",
				headRoot.Hex(), prev.Hex(), headNum)
		}

		log.Printf("signBatch: prev=%s claimed=%s epoch=L1(%d→%d) ts=%d userTxs=%d",
			prev.Hex()[:10], claimed.Hex()[:10], hdr.L1FromBlock, hdr.L1ToBlock,
			hdr.L2Timestamp, len(rawTxs))

		// Snapshot BEFORE mining so the settler can roll us back if
		// this batch's postBatch never lands on L1.
		snapID, err := v.l2.Snapshot(ctx)
		if err != nil {
			return nil, fmt.Errorf("snapshot: %w", err)
		}

		// Helper: any error path after the snapshot must revert it,
		// otherwise the verifier's anvil drifts ahead of L1 (its head ≠
		// the on-chain root) and every subsequent signBatch is rejected
		// with "verifier is out of sync".
		revert := func(why error) error {
			if rerr := v.l2.Revert(ctx, snapID); rerr != nil {
				log.Printf("warn: revert snapshot %s after %v: %v", snapID, why, rerr)
			}
			return why
		}

		deps, err := derivation.FetchDeposits(ctx, v.l1, v.cfg.DepositGatewayAddr, hdr.L1FromBlock, hdr.L1ToBlock)
		if err != nil {
			return nil, revert(fmt.Errorf("fetch deposits: %w", err))
		}
		if err := derivation.ApplyDeposits(ctx, v.l2, deps); err != nil {
			return nil, revert(fmt.Errorf("apply deposits: %w", err))
		}
		if err := derivation.SubmitUserTxs(ctx, v.l2, rawTxs); err != nil {
			return nil, revert(fmt.Errorf("submit user txs: %w", err))
		}
		mineHdr, err := v.l2.Mine(ctx, hdr.L2Timestamp)
		if err != nil {
			return nil, revert(fmt.Errorf("mine: %w", err))
		}
		derived := mineHdr.Root
		if derived != claimed {
			return nil, revert(fmt.Errorf("STATE-ROOT MISMATCH: settler claims %s, verifier derived %s (l2 block %d)",
				claimed.Hex(), derived.Hex(), mineHdr.Number))
		}
		log.Printf("derived: l2 block %d root=%s deps=%d userTxs=%d",
			mineHdr.Number, derived.Hex()[:10], len(deps), len(rawTxs))
		v.mu.Lock()
		v.verifiedBatches[batchKey] = true
		v.preMineSnapshot = snapID
		v.preMineKey = batchKey
		v.mu.Unlock()
	} else {
		log.Printf("re-sign request for already-verified batch (prev=%s claimed=%s cd=%s) — skipping re-derive",
			prev.Hex()[:10], claimed.Hex()[:10], cdHash.Hex()[:10])
	}

	// 5) Re-compute publicInputsHash from the same inputs the settler
	// hashed. We don't blindly trust req.PublicHash — we recompute and
	// compare so the settler can't trick us into signing some other
	// digest.
	rid := new(big.Int).SetUint64(req.RollupID)
	deltas := []rollups.StateDelta{{
		RollupId:     rid,
		CurrentState: prev,
		NewState:     claimed,
		EtherDelta:   big.NewInt(0),
	}}
	entries := []rollups.ExecutionEntry{{
		StateDeltas:         deltas,
		CrossChainCallHash:  [32]byte{},
		DestinationRollupId: rid,
		Calls:               []rollups.CrossChainCall{},
		NestedActions:       []rollups.NestedAction{},
		CallCount:           big.NewInt(0),
		ReturnData:          []byte{},
		RollingHash:         [32]byte{},
	}}

	pendingHdr := &types.Header{
		ParentHash: common.HexToHash(req.L1ParentHash),
		Time:       req.L1Timestamp,
	}
	shared := proof.SharedPublicInput(
		[]*big.Int{rid},
		entries,
		[]rollups.LookupCall{},
		pendingHdr,
		callData,
		req.BlobCount,
		[32]byte{},
	)
	myHash := proof.PublicInputsHashForPS(shared, [][32]byte{v.vk})
	expected := common.HexToHash(req.PublicHash)
	if common.Hash(myHash) != expected {
		return nil, fmt.Errorf("publicInputsHash mismatch: settler=%s verifier=%s",
			expected.Hex(), common.Hash(myHash).Hex())
	}

	// 6) Sign.
	sig, err := crypto.Sign(myHash[:], v.signerKey)
	if err != nil {
		return nil, fmt.Errorf("sign: %w", err)
	}
	if sig[64] < 27 {
		sig[64] += 27
	}
	log.Printf("SIGN: prev=%s claimed=%s userTxs=%d (cached=%v)",
		prev.Hex()[:10], claimed.Hex()[:10], len(rawTxs), alreadyVerified)
	return sig, nil
}

func (v *Verifier) rollbackBatch(ctx context.Context, req *rollbackReq) error {
	wantKey := common.HexToHash(req.PrevRoot).Hex() +
		common.HexToHash(req.ClaimedRoot).Hex() +
		common.HexToHash(req.CallDataHash).Hex()

	v.mu.Lock()
	defer v.mu.Unlock()
	if v.preMineSnapshot == "" || v.preMineKey != wantKey {
		// Either we have no held snapshot (settler asking to roll back
		// something we already committed past), or the snapshot we hold
		// is for a different batch. No-op rather than error — the
		// settler treats rollback as best-effort.
		log.Printf("rollback no-op: held key=%s want key=%s", v.preMineKey[:20], wantKey[:20])
		return nil
	}
	if err := v.l2.Revert(ctx, v.preMineSnapshot); err != nil {
		return fmt.Errorf("revert: %w", err)
	}
	delete(v.verifiedBatches, wantKey)
	v.preMineSnapshot = ""
	v.preMineKey = ""
	log.Printf("ROLLBACK: reverted to pre-batch snapshot for prev=%s claimed=%s",
		req.PrevRoot[:10], req.ClaimedRoot[:10])
	return nil
}

func (v *Verifier) statusHandler(w http.ResponseWriter, r *http.Request) {
	v.mu.Lock()
	defer v.mu.Unlock()
	w.Header().Set("Content-Type", "application/json")
	_ = json.NewEncoder(w).Encode(map[string]any{
		"signer":   v.signerAddr.Hex(),
		"signed":   v.signed,
		"rejected": v.rejected,
		"lastErr":  v.lastErr,
	})
}

func env(k, def string) string {
	if v, ok := os.LookupEnv(k); ok {
		return v
	}
	return def
}

func stripHex(s string) string {
	if strings.HasPrefix(s, "0x") || strings.HasPrefix(s, "0X") {
		return s[2:]
	}
	return s
}

func hexEnc(b []byte) string {
	const hexChars = "0123456789abcdef"
	out := make([]byte, len(b)*2)
	for i, v := range b {
		out[2*i] = hexChars[v>>4]
		out[2*i+1] = hexChars[v&0x0f]
	}
	return string(out)
}
