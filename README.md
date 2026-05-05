# op-rollups-prover

Standalone prover for the **PR-22 multi-prover Rollups** (sync-rollups-v2)
deployed on Gnosis Chiado. Runs the verifier service that the settler
operator queries to co-attest each batch.

## What this does

Each `postBatch` on the central `Rollups` registry must carry one
signature **per registered proof system** for the rollup. This repo
contributes one such signature path:

1. You deploy your own `ECDSAProofSystem` on Chiado (the contract is
   under `contracts/`, owned by your wallet, signing key = your wallet).
2. The settler operator registers your proof system in the rollup's
   manager (`Rollup.addProofSystem(yourPS, vk)`) and bumps `threshold`
   so both PS sigs are required.
3. You run the verifier process here. It exposes a JSON-RPC at
   `:9100` that the settler queries with `verifier_signBatch`.
4. On each request the verifier **independently re-derives** the L2
   state from the same `callData` against its own L2 EL. It only signs
   if its derived `stateRoot` matches what the settler claims, and only
   if its locally-computed `publicInputsHash` matches the requested digest.
5. The settler bundles your sig with its own and posts to the
   `Rollups` contract. With threshold=2, neither side can land a batch
   alone — meaningful operational separation.

The architecture works regardless of how many provers there are; this
repo is the boilerplate to add **one** independent prover.

## What you need before starting

- **Linux/macOS shell** with: `git`, `curl`, `jq`
- **Foundry** (`forge`, `cast`):
  `curl -L https://foundry.paradigm.xyz | bash && foundryup`
- **Go ≥ 1.24**: https://go.dev/dl/
- **Anvil** (ships with Foundry — used as the local L2 EL)
- **Network reachability** between this machine and the settler
  operator's machine on TCP `:9100` (or whichever port you pick for
  `VERIFIER_LISTEN`).
- **A funded Chiado wallet.** A few tenths of a xDAI is enough to
  deploy the proof system and pay for occasional admin txs. Faucets:
  - https://gnosisfaucet.com/  (pick "Chiado")
  - https://faucet.chiadochain.net/

## Setup — step by step

### 1. Clone and configure

```bash
git clone <this-repo> op-rollups-prover
cd op-rollups-prover
cp .env.example .env
```

Generate a fresh signing key (don't reuse one you already have on
mainnet — this key signs batches and the settler operator will see its
address):

```bash
cast wallet new
# pick a mnemonic-free key, paste the 0x-prefixed private key into .env:
#   PROVER_KEY=0x...
```

Fund the resulting address on Chiado from a faucet.

### 2. Get the L1 deployment details from the settler operator

Ask the settler operator (the other Claude instance / human running
the settler) for these values and paste them into `.env`:

```
L1_RPC            # default https://rpc.chiadochain.net is fine
ROLLUPS_ADDR      # central registry, shared across rollups
ROLLUP_MGR_ADDR   # per-rollup manager that gates proof systems + vkeys
ROLLUP_ID         # uint256 id assigned by createRollup
DEPOSIT_ADDR      # the rollup's DepositGateway on L1
WITHDRAWAL_ADDR   # the rollup's WithdrawalGateway on L1
PREFUND_WALLETS   # CSV of L2 addresses to pre-fund on your anvil
                  # (must match what the settler funds; otherwise tx
                  #  replays will fail with insufficient balance and
                  #  produce a STATE-ROOT MISMATCH on the first batch)
```

### 3. Build everything

```bash
make setup-contracts   # forge install OZ + forge build
make build             # go build the verifier binary into ./bin/
```

### 4. Deploy your ECDSAProofSystem on Chiado

```bash
make deploy-ps
```

This prints the deployed address and appends it as
`PROOF_SYSTEM_ADDR=...` to your `.env`. It will also print the next
step you need to send back to the settler operator:

```
Send the settler operator your proof-system address:
  PROOF_SYSTEM = 0x...
  SIGNER       = 0x...
```

### 5. Wait for the settler operator to register your PS

Until they call `Rollup.addProofSystem(yourPS, vk)` and bump
`threshold` to 2, your verifier cannot start (it reads its vkey from
the manager at boot). Check status whenever you want:

```bash
make check-registration
```

When it prints `✅ registered with vkey 0x...`, you're cleared.

### 6. Start the L2 anvil + verifier

```bash
make start-anvil       # starts anvil --port 8547 --no-mining
make start-verifier    # starts the Go verifier on :9100
```

Both run as background processes; logs at `/tmp/l2_anvil.log` and
`/tmp/op_prover_verifier.log`.

Sanity check:

```bash
make status
```

You should see:

```json
{"signer":"0x...","signed":0,"rejected":0,"lastErr":""}
```

### 7. Tell the settler operator your verifier RPC URL

They'll wire it into the settler so it polls **both** verifiers for
each batch sig:

```
http://<your-public-ip>:9100/
```

(Or whatever you set `VERIFIER_LISTEN` to. Make sure the port is
reachable — open the firewall, NAT-forward, or use a tunnel if
necessary.)

## How verification actually works

When the settler asks for a sig (`verifier_signBatch`), the verifier:

1. Refuses if its own L2 anvil head root ≠ `prevRoot` — protects
   against the settler skipping or fabricating prior batches
   (`l2 head root … ≠ batch prev …`).
2. Decodes the `callData` to extract the L1 epoch range
   `(l1FromBlock, l1ToBlock]` + user-tx list.
3. Snapshots the anvil so it can roll back on any error.
4. Calls **its own L1 RPC** with `derivation.FetchDeposits` —
   reads `TransactionDeposited` events from the DepositGateway. Doesn't
   trust the settler for the deposit list.
5. Applies those deposits + submits user txs to its own L2 anvil.
6. Mines a block at the requested timestamp.
7. Compares its derived state root to `claimed`. Mismatch → revert
   snapshot, return `STATE-ROOT MISMATCH`, increment `rejected`.
8. Independently re-computes `publicInputsHash` and compares to the
   settler's claimed value (defense against being asked to sign some
   unrelated digest).
9. **Only then** signs `publicInputsHash` with `PROVER_KEY`.

The `/status` endpoint exposes counters:

```bash
curl -s http://127.0.0.1:9100/status | jq
# { "signer": "0x...", "signed": N, "rejected": M, "lastErr": "..." }
```

Any rejection has its full reason in `lastErr` and in the verifier
log.

## Layout

```
.
├── cmd/verifier/        # Go verifier binary entrypoint
├── internal/
│   ├── config/          # CLI flags + env
│   ├── derivation/      # canonical L2 derivation recipe
│   ├── l1/              # L1 RPC wrapper (deposits, finalized tag)
│   ├── l2/              # L2 anvil RPC wrapper
│   └── proof/           # publicInputsHash + entry hashing
├── bindings/            # abigen output for the relevant contracts
│   ├── rollups/         # central Rollups registry
│   ├── rollupmgr/       # per-rollup manager
│   ├── proofsys/        # ECDSAProofSystem
│   ├── deposit/         # DepositGateway
│   └── withdrawal/      # WithdrawalGateway
├── contracts/
│   ├── ECDSAProofSystem.sol  # what you deploy on L1
│   ├── IProofSystem.sol      # interface it implements
│   └── foundry.toml          # solc 0.8.34, optimizer=200, via_ir=true
├── scripts/
│   ├── setup_contracts.sh    # forge install + build
│   ├── deploy_proof_system.sh
│   ├── start_l2_anvil.sh
│   ├── start_verifier.sh
│   └── check_registration.sh
└── Makefile
```

## Common failure modes

| Symptom | Diagnosis |
|---|---|
| `manager %s reports zero vkey for PS …` at startup | Settler operator hasn't run `addProofSystem` yet. Run `make check-registration`. |
| `l2 head root X ≠ batch prev Y — verifier is out of sync` | Your anvil drifted from the canonical chain (e.g. reset settler-side without resetting your anvil). Solution: kill `anvil`, `make start-anvil` again, ensure the rollup is at L2 block 0 with empty trie root on L1. |
| `STATE-ROOT MISMATCH: settler claims … verifier derived …` | Either the settler is malicious (good — refusing is the point) or your `PREFUND_WALLETS` doesn't match what the settler funded. Diff the wallet lists. |
| `nonce too low / replacement underpriced` from `forge create` | Chiado RPC mempool is lagging. Wait 30s and retry, or bump `--gas-price`. |
| Settler can't reach your `:9100` | Firewall / NAT. Either open the port, set up a reverse tunnel (e.g. `ssh -R 9100:localhost:9100 settler-host`), or expose via Cloudflare Tunnel / Tailscale. |

## Catch-up: how a fresh prover syncs from L1

When the verifier starts and finds its L2 anvil head root ≠ the on-chain
rollup root, it kicks off catch-up before serving any signBatch RPC.
This is the canonical "any node can sync from L1 alone" property —
zero settler trust:

1. Filter L1 logs for `L2ExecutionPerformed(rollupId, newState)` from
   the Rollups contract — one event per applied state delta. Public
   Chiado RPC supports the standard log filter range, chunked at 5,000
   blocks.
2. Pre-flight all `TransactionDeposited` events from the deposit
   gateway in one chunked sweep; index by L1 block. Per-batch deposit
   lookups become O(1) map probes.
3. For each unique L1 tx, fetch + ABI-decode the `postBatch(...)`
   calldata, find our sub-batch's `callData`, and replay through the
   derivation pipeline (decode → apply deposits → submit user txs →
   mine).
4. Loop: chain advances during replay, so re-poll on each pass until
   our L2 root stabilises at the on-chain root.

Public Chiado RPC rate-limits hard (~5-10 req/sec). Two defensive
mechanisms are built in:

- `withRetry` exponential backoff (500ms → 30s, 8 attempts) on the
  recoverable HTTP errors (429, EOF, timeout, connection reset).
- 80ms throttle between batches to stay under the limit proactively.

Observed throughput on free-tier rate-limited RPC: ~8.5 batches/sec
(~7 minutes for ~3,800 batches). Operators with their own RPC node
will be faster.

Progress is logged every 250 batches with rate + ETA.

## Caveats / known limitations
- **Chain id must match.** `L2_CHAIN_ID` here must equal what the
  settler runs (`4242` in the reference deployment). Otherwise the
  signed user txs decode to different recovered senders and you get
  silent state-root drift.
- **`anvil_setBalance` is the funding mechanism.** That's a cheat
  method: in production this would be the genesis state. Pre-fund
  exactly the wallets the settler pre-funds — the canonical
  derivation recipe assumes both ELs start identically.
- **Single-machine demo.** This repo only adds **one** independent
  prover. The trust model improves with each one — the goal is several
  unrelated operators running this stack with different keys.
