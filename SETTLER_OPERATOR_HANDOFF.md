# Settler-side handoff

When the prover operator finishes `make deploy-ps` on their machine,
they'll send you two values:

```
PROOF_SYSTEM = 0x...
SIGNER       = 0x...
```

Three things to do before they can start their verifier:

## 1. Register their proof system in the manager

```bash
ROLLUP_MGR=0xEE43bb490a511d2Ed343553333DCD77c2F0056dE   # current rollupId=2 manager
NEW_PS=<their PROOF_SYSTEM>
VKEY=0x0000000000000000000000000000000000000000000000000000000000000001

cast send "$ROLLUP_MGR" \
  "addProofSystem(address,bytes32)" \
  "$NEW_PS" "$VKEY" \
  --private-key "$WALLET_KEY" --rpc-url https://rpc.chiadochain.net \
  --gas-price 2000000000 --legacy
```

## 2. Bump the threshold to 2

```bash
cast send "$ROLLUP_MGR" \
  "setThreshold(uint256)" 2 \
  --private-key "$WALLET_KEY" --rpc-url https://rpc.chiadochain.net \
  --gas-price 2000000000 --legacy
```

## 3. Update the settler to query both verifiers

The current settler binary only takes a single `-verifier-rpc`. To use
two verifiers properly the settler needs to:

- Query each verifier independently for `verifier_signBatch`
- Build the `ProofSystemBatch` with `proofSystems = [psA, psB]`
  (sorted ascending by address)
- Build `proof = [sigA, sigB]` in the same order

This requires a settler code change. Sketch:

```go
// in settler/internal/settler/settler.go
type verifierEndpoint struct {
    proofSystem common.Address
    rpc         *gethrpc.Client
}
// Sort by proofSystem address (ascending) — Rollups.sol enforces this.
sort.Slice(s.verifiers, func(i, j int) bool {
    return bytes.Compare(s.verifiers[i].proofSystem.Bytes(),
                         s.verifiers[j].proofSystem.Bytes()) < 0
})

// signProof now collects sigs from all of them in order
sigs := make([][]byte, len(s.verifiers))
psAddrs := make([]common.Address, len(s.verifiers))
for i, v := range s.verifiers {
    sigs[i] = ... // RPC call
    psAddrs[i] = v.proofSystem
}

// And the postBatch construction uses these arrays
batch := rollups.ProofSystemBatch{
    ProofSystems: psAddrs,
    ...
    Proof: sigs,
    ...
}
```

Until that lands, the practical workaround is to keep `threshold=1`,
have the prover operator run their verifier alongside the local one,
and verify it returns the SAME signature for the SAME digest (a trivial
co-attestation check, not yet enforced on-chain).

## 4. Tell the prover operator to start their verifier

Once `addProofSystem` and `setThreshold(2)` are mined, the prover's
`make check-registration` will flip to `✅`, and they can run
`make start-verifier`. Send them back the verifier RPC URL of their
verifier so we can hit it from the settler.
