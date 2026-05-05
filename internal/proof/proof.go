// Package proof exports the publicInputsHash computation and signing
// helpers used by both the sequencer (settler) and the verifier.
//
// Multi-prover Rollups (sync-rollups-v2 / PR-22) layout:
//
//	sharedPublicInput = keccak256(abi.encodePacked(
//	    blockhash(block.number - 1),
//	    block.timestamp,
//	    abi.encode(rollupIds),
//	    abi.encode(entryHashes),
//	    abi.encode(lookupCallHashes),
//	    abi.encode(blobHashes),
//	    keccak256(callData),
//	    crossProofSystemInteractions
//	))
//	publicInputsHash[k] = keccak256(abi.encodePacked(
//	    sharedPublicInput,
//	    abi.encode(rollupVks)
//	))
//
// Per-entry hash is just keccak256(abi.encode(ExecutionEntry)) — the full
// struct is hashed, no per-rollup vkey lookup. The vkey enters the proof
// indirectly via the per-PS publicInputsHash.
package proof

import (
	"crypto/ecdsa"
	"fmt"
	"math/big"

	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"

	"github.com/koeppelmann/op-rollups-prover/bindings/rollups"
)

// EntryHash mirrors Rollups.sol's per-entry hash: keccak256(abi.encode(entry)).
func EntryHash(e rollups.ExecutionEntry) [32]byte {
	enc, err := encodeExecutionEntry(e)
	if err != nil {
		// Match the Solidity behavior: a malformed entry would have already
		// failed _validateStructure long before we get here. Returning a
		// zero hash on encode failure is safe — a wrong hash just produces
		// an InvalidProof revert downstream.
		return [32]byte{}
	}
	return crypto.Keccak256Hash(enc)
}

// LookupCallHash mirrors keccak256(abi.encode(lookupCall)).
func LookupCallHash(l rollups.LookupCall) [32]byte {
	enc, err := encodeLookupCall(l)
	if err != nil {
		return [32]byte{}
	}
	return crypto.Keccak256Hash(enc)
}

// SharedPublicInput computes the cross-PS portion of the publicInputsHash.
// `pending` is the L1 header THIS postBatch will land in — its ParentHash
// plays the role of `blockhash(block.number - 1)` and `Time` plays the
// role of `block.timestamp`.
func SharedPublicInput(
	rollupIds []*big.Int,
	entries []rollups.ExecutionEntry,
	lookupCalls []rollups.LookupCall,
	pending *types.Header,
	callData []byte,
	blobCount uint64,
	crossProofSystemInteractions [32]byte,
) [32]byte {
	entryHashes := make([][32]byte, len(entries))
	for i, e := range entries {
		entryHashes[i] = EntryHash(e)
	}
	lookupHashes := make([][32]byte, len(lookupCalls))
	for i, l := range lookupCalls {
		lookupHashes[i] = LookupCallHash(l)
	}
	blobHashes := make([][32]byte, blobCount)
	// blob hashes left zero — we don't post blobs.

	cdHash := crypto.Keccak256Hash(callData)

	encRollupIds := encodeUint256Array(rollupIds)
	encEntryHashes := encodeBytes32Array(entryHashes)
	encLookupHashes := encodeBytes32Array(lookupHashes)
	encBlobHashes := encodeBytes32Array(blobHashes)

	prevHash := pending.ParentHash
	timestamp := uint256BE(pending.Time)

	var concat []byte
	concat = append(concat, prevHash.Bytes()...)
	concat = append(concat, timestamp[:]...)
	concat = append(concat, encRollupIds...)
	concat = append(concat, encEntryHashes...)
	concat = append(concat, encLookupHashes...)
	concat = append(concat, encBlobHashes...)
	concat = append(concat, cdHash.Bytes()...)
	concat = append(concat, crossProofSystemInteractions[:]...)
	return crypto.Keccak256Hash(concat)
}

// PublicInputsHashForPS combines the cross-PS shared input with the
// per-PS rollupVks (one vkey per rollup, in rollupIds order).
func PublicInputsHashForPS(shared [32]byte, rollupVks [][32]byte) [32]byte {
	enc := encodeBytes32Array(rollupVks)
	var concat []byte
	concat = append(concat, shared[:]...)
	concat = append(concat, enc...)
	return crypto.Keccak256Hash(concat)
}

// SignRaw signs a raw 32-byte digest WITHOUT the EIP-191 prefix. The
// ECDSAProofSystem expects the signer to have signed publicInputsHash
// directly.
func SignRaw(key *ecdsa.PrivateKey, digest [32]byte) ([]byte, error) {
	sig, err := crypto.Sign(digest[:], key)
	if err != nil {
		return nil, err
	}
	return sig, nil
}

// NormalizeV converts geth's 0/1 recovery byte to OZ ECDSA-recover's
// expected 27/28. Mutates in place.
func NormalizeV(sig []byte) error {
	if len(sig) != 65 {
		return fmt.Errorf("expected 65-byte sig, got %d", len(sig))
	}
	if sig[64] < 27 {
		sig[64] += 27
	}
	return nil
}

// uint256BE returns 32 big-endian bytes for v.
func uint256BE(v uint64) [32]byte {
	var b [32]byte
	bi := new(big.Int).SetUint64(v)
	bi.FillBytes(b[:])
	return b
}

// encodeBytes32Array does abi.encode(bytes32[]).
func encodeBytes32Array(arr [][32]byte) []byte {
	t, _ := abi.NewType("bytes32[]", "", nil)
	args := abi.Arguments{{Type: t}}
	out, err := args.Pack(arr)
	if err != nil {
		return nil
	}
	return out
}

// encodeUint256Array does abi.encode(uint256[]).
func encodeUint256Array(arr []*big.Int) []byte {
	t, _ := abi.NewType("uint256[]", "", nil)
	args := abi.Arguments{{Type: t}}
	out, err := args.Pack(arr)
	if err != nil {
		return nil
	}
	return out
}

// executionEntryArgs is the full ABI shape of the ExecutionEntry struct.
// keccak256(abi.encode(entry)) needs every field encoded exactly the way
// Solidity sees it, including dynamic-tuple offsets — abi.Arguments handles
// that automatically.
func executionEntryArgs() (abi.Arguments, error) {
	t, err := abi.NewType("tuple", "", []abi.ArgumentMarshaling{
		{Name: "stateDeltas", Type: "tuple[]", Components: []abi.ArgumentMarshaling{
			{Name: "rollupId", Type: "uint256"},
			{Name: "currentState", Type: "bytes32"},
			{Name: "newState", Type: "bytes32"},
			{Name: "etherDelta", Type: "int256"},
		}},
		{Name: "crossChainCallHash", Type: "bytes32"},
		{Name: "destinationRollupId", Type: "uint256"},
		{Name: "calls", Type: "tuple[]", Components: []abi.ArgumentMarshaling{
			{Name: "targetAddress", Type: "address"},
			{Name: "value", Type: "uint256"},
			{Name: "data", Type: "bytes"},
			{Name: "sourceAddress", Type: "address"},
			{Name: "sourceRollupId", Type: "uint256"},
			{Name: "revertSpan", Type: "uint256"},
		}},
		{Name: "nestedActions", Type: "tuple[]", Components: []abi.ArgumentMarshaling{
			{Name: "crossChainCallHash", Type: "bytes32"},
			{Name: "callCount", Type: "uint256"},
			{Name: "returnData", Type: "bytes"},
		}},
		{Name: "callCount", Type: "uint256"},
		{Name: "returnData", Type: "bytes"},
		{Name: "rollingHash", Type: "bytes32"},
	})
	if err != nil {
		return nil, err
	}
	return abi.Arguments{{Type: t}}, nil
}

// abiStateDelta / abiCall / abiNestedAction match the field layouts the geth
// ABI encoder expects. The generated `rollups.StateDelta` etc. work too,
// but go-ethereum's abi reflection is stricter about exported field names
// (must match the json tag), so a parallel struct keeps the encoding
// surface independent of binding regeneration.
type abiStateDelta struct {
	RollupId     *big.Int
	CurrentState [32]byte
	NewState     [32]byte
	EtherDelta   *big.Int
}

type abiCall struct {
	TargetAddress  common.Address
	Value          *big.Int
	Data           []byte
	SourceAddress  common.Address
	SourceRollupId *big.Int
	RevertSpan     *big.Int
}

type abiNestedAction struct {
	CrossChainCallHash [32]byte
	CallCount          *big.Int
	ReturnData         []byte
}

type abiEntry struct {
	StateDeltas         []abiStateDelta
	CrossChainCallHash  [32]byte
	DestinationRollupId *big.Int
	Calls               []abiCall
	NestedActions       []abiNestedAction
	CallCount           *big.Int
	ReturnData          []byte
	RollingHash         [32]byte
}

func encodeExecutionEntry(e rollups.ExecutionEntry) ([]byte, error) {
	args, err := executionEntryArgs()
	if err != nil {
		return nil, err
	}
	deltas := make([]abiStateDelta, len(e.StateDeltas))
	for i, d := range e.StateDeltas {
		deltas[i] = abiStateDelta{
			RollupId:     d.RollupId,
			CurrentState: d.CurrentState,
			NewState:     d.NewState,
			EtherDelta:   d.EtherDelta,
		}
	}
	calls := make([]abiCall, len(e.Calls))
	for i, c := range e.Calls {
		calls[i] = abiCall{
			TargetAddress:  c.TargetAddress,
			Value:          c.Value,
			Data:           c.Data,
			SourceAddress:  c.SourceAddress,
			SourceRollupId: c.SourceRollupId,
			RevertSpan:     c.RevertSpan,
		}
	}
	nested := make([]abiNestedAction, len(e.NestedActions))
	for i, n := range e.NestedActions {
		nested[i] = abiNestedAction{
			CrossChainCallHash: n.CrossChainCallHash,
			CallCount:          n.CallCount,
			ReturnData:         n.ReturnData,
		}
	}
	v := abiEntry{
		StateDeltas:         deltas,
		CrossChainCallHash:  e.CrossChainCallHash,
		DestinationRollupId: e.DestinationRollupId,
		Calls:               calls,
		NestedActions:       nested,
		CallCount:           e.CallCount,
		ReturnData:          e.ReturnData,
		RollingHash:         e.RollingHash,
	}
	return args.Pack(v)
}

// lookupCallArgs is the full ABI shape of the LookupCall struct.
func lookupCallArgs() (abi.Arguments, error) {
	t, err := abi.NewType("tuple", "", []abi.ArgumentMarshaling{
		{Name: "crossChainCallHash", Type: "bytes32"},
		{Name: "destinationRollupId", Type: "uint256"},
		{Name: "returnData", Type: "bytes"},
		{Name: "failed", Type: "bool"},
		{Name: "callNumber", Type: "uint64"},
		{Name: "lastNestedActionConsumed", Type: "uint64"},
		{Name: "calls", Type: "tuple[]", Components: []abi.ArgumentMarshaling{
			{Name: "targetAddress", Type: "address"},
			{Name: "value", Type: "uint256"},
			{Name: "data", Type: "bytes"},
			{Name: "sourceAddress", Type: "address"},
			{Name: "sourceRollupId", Type: "uint256"},
			{Name: "revertSpan", Type: "uint256"},
		}},
		{Name: "rollingHash", Type: "bytes32"},
	})
	if err != nil {
		return nil, err
	}
	return abi.Arguments{{Type: t}}, nil
}

type abiLookupCall struct {
	CrossChainCallHash       [32]byte
	DestinationRollupId      *big.Int
	ReturnData               []byte
	Failed                   bool
	CallNumber               uint64
	LastNestedActionConsumed uint64
	Calls                    []abiCall
	RollingHash              [32]byte
}

func encodeLookupCall(l rollups.LookupCall) ([]byte, error) {
	args, err := lookupCallArgs()
	if err != nil {
		return nil, err
	}
	calls := make([]abiCall, len(l.Calls))
	for i, c := range l.Calls {
		calls[i] = abiCall{
			TargetAddress:  c.TargetAddress,
			Value:          c.Value,
			Data:           c.Data,
			SourceAddress:  c.SourceAddress,
			SourceRollupId: c.SourceRollupId,
			RevertSpan:     c.RevertSpan,
		}
	}
	v := abiLookupCall{
		CrossChainCallHash:       l.CrossChainCallHash,
		DestinationRollupId:      l.DestinationRollupId,
		ReturnData:               l.ReturnData,
		Failed:                   l.Failed,
		CallNumber:               l.CallNumber,
		LastNestedActionConsumed: l.LastNestedActionConsumed,
		Calls:                    calls,
		RollingHash:              l.RollingHash,
	}
	return args.Pack(v)
}
