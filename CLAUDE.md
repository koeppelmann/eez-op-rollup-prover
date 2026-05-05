# Claude project notes — op-rollups-prover

This is the **standalone prover** for the multi-prover Rollups demo on
Gnosis Chiado. The full architecture lives in two other places:

- `sync-rollups-v2` — the canonical Solidity protocol (PR-22 multi-prover
  branch). The L1 contracts on Chiado are built from there.
- `op-rollups` — the settler + verifier monorepo on the operator's
  machine. The verifier here is a stripped-down copy of the one in
  `op-rollups/settler/cmd/verifier/`, with the same protocol semantics.

**Read `README.md` first** — it has the full operator-handoff flow.
This file is just for context if you (Claude) need to debug or extend.

## Mental model

```
   ┌─────────────────────────┐         ┌─────────────────────────┐
   │  settler operator       │ HTTPS   │   prover (this repo)    │
   │  (other machine)        │ ───────▶│                         │
   │                         │ signBatch│  - own L2 anvil :8547   │
   │  - sequencer L2 :8546   │         │  - own L1 RPC connection│
   │  - settler              │         │  - own signing key      │
   │  - own L1 RPC           │         │  - own ECDSAProofSystem │
   │  - own ECDSAProofSystem │         │    on Chiado            │
   │    on Chiado            │         │                         │
   └─────────────────────────┘         └─────────────────────────┘
                ▲                                    ▲
                │ postBatch with TWO sigs            │
                │ (threshold=2)                      │
                ▼                                    │
   ┌─────────────────────────────────────────────────┘
   │              L1 (Chiado)                        │
   │  Rollups + Rollup manager + N×ECDSAProofSystem  │
   └──────────────────────────────────────────────────
```

The settler builds a `ProofSystemBatch` listing both proof systems
(addresses sorted ascending), gets a sig from each verifier independently,
and posts both. On-chain `Rollups._verifyProofSystemBatch` requires every
sig to verify against its PS — a single mismatch reverts the whole
postBatch.

## Verification flow inside `signBatch`

The full state machine is in `cmd/verifier/main.go:255` (`signBatch`).
Key invariants:

1. **L2 head root must equal `prev`** — protects against being asked to
   re-attest a divergent history. If your anvil is drifting the only
   correct response is to refuse and surface `verifier is out of sync`.
2. **Snapshot before every state mutation** — `evm_snapshot` is taken
   before fetching deposits / mining, and reverted on every error path
   (mine error, deposits error, root mismatch, etc.). Without this the
   anvil drifts forward and every subsequent `signBatch` fails.
3. **publicInputsHash is recomputed locally** — never trust
   `req.publicHash`. Hash from the same inputs the settler hashed from
   and compare. This is what prevents a malicious settler from asking
   for a sig over a different digest with our key.
4. **Already-verified batches cache** (`verifiedBatches`) — a batch
   keyed by `(prev, claimed, keccak(callData))` that we already
   re-derived once skips the re-derivation on resign. This is how we
   handle the settler's retry-on-bad-prediction loop on real Chiado
   without re-mining the same L2 block.

## State machine for the L2 anvil

- Block 0: empty trie root (`0x56e81f17…`). Settler must initialize the
  on-chain rollup with this same root.
- Each successful `signBatch` advances exactly one L2 block.
- The verifier holds **one** snapshot id (`preMineSnapshot`) — the most
  recent successful signBatch's pre-mine state — so the settler can
  call `verifier_rollbackBatch` if its postBatch fails to land on L1.
  This keeps both ELs in lockstep with what L1 actually committed.

## Things that look like bugs but aren't

- **Verifier mines L2 blocks "out of order"** with respect to the
  sequencer. That's fine — it independently produces its own L2 chain
  and the only check is "does the resulting state root match what the
  settler claims". The block hashes will differ between sequencer and
  verifier (different miner, different timestamps if rounding differs)
  but the state root is what the protocol cares about.
- **The proof's vkey is a non-zero placeholder** (`bytes32(uint256(1))`).
  ECDSAProofSystem ignores vkey contents — it just verifies a 65-byte
  secp256k1 sig recovers to the on-chain `signer`. The vkey field exists
  for future ZK proof systems where it really is the proving key.
- **`finalized` lags `safe` by ~3 minutes on Chiado.** That's correct —
  the settler queries Chiado's `finalized` block tag (real beacon-chain
  finality, ~12-16 epochs). Don't "fix" this by lowering the lag.

## What's NOT in this repo

- The settler. It runs on the operator's machine and is the only thing
  that talks to L1 with the postBatch tx. We just provide one of N sigs.
- The arb-demo + dashboard. Settler operator runs those.
- The L1 contract source for `Rollups`/`Rollup`/etc. Those live in
  `sync-rollups-v2`. We only need their bindings and our own
  `ECDSAProofSystem.sol`.
