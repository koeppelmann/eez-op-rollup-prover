#!/usr/bin/env bash
# Start the verifier process. Reads .env for all addresses + keys.
#
# Prerequisites (in order):
#   1. setup_contracts.sh (foundry deps + build)
#   2. deploy_proof_system.sh (your ECDSAProofSystem on Chiado)
#   3. settler operator runs Rollup.addProofSystem(yourPS, vk) on the manager
#      and bumps threshold (otherwise verifier startup fails with
#      "manager reports zero vkey for PS …")
#   4. start_l2_anvil.sh
#   5. this script
set -euo pipefail
ROOT="$(cd "$(dirname "$0")/.." && pwd)"
[[ -f "$ROOT/.env" ]] && source "$ROOT/.env"

# required
: "${L1_RPC:?L1_RPC required (Chiado RPC URL)}"
: "${ROLLUPS_ADDR:?ROLLUPS_ADDR required (sent by settler operator)}"
: "${ROLLUP_MGR_ADDR:?ROLLUP_MGR_ADDR required (sent by settler operator)}"
: "${ROLLUP_ID:?ROLLUP_ID required (sent by settler operator)}"
: "${DEPOSIT_ADDR:?DEPOSIT_ADDR required (sent by settler operator)}"
: "${WITHDRAWAL_ADDR:?WITHDRAWAL_ADDR required (sent by settler operator)}"
: "${PROOF_SYSTEM_ADDR:?PROOF_SYSTEM_ADDR required (run deploy_proof_system.sh first)}"
: "${PROVER_KEY:?PROVER_KEY required (your signing key)}"

# optional
L2_PORT="${L2_PORT:-8547}"
LISTEN="${VERIFIER_LISTEN:-:9100}"
L1_CHAIN_ID="${L1_CHAIN_ID:-10200}"
L2_CHAIN_ID="${L2_CHAIN_ID:-4242}"
START_L1_BLOCK="${START_L1_BLOCK:-0}"
FINALIZATION_DEPTH="${FINALIZATION_DEPTH:-12}"

VERIFIER_BIN="${VERIFIER_BIN:-$ROOT/bin/op-prover-verifier}"
if [[ ! -x "$VERIFIER_BIN" ]]; then
  echo "binary missing: $VERIFIER_BIN — run 'make build' first" >&2
  exit 1
fi

if pgrep -f "$VERIFIER_BIN" >/dev/null; then
  echo "verifier already running on $LISTEN"
  exit 0
fi

echo "starting verifier:"
echo "  L1_RPC:           $L1_RPC"
echo "  L2:               http://127.0.0.1:$L2_PORT"
echo "  ROLLUPS:          $ROLLUPS_ADDR"
echo "  ROLLUP_MGR:       $ROLLUP_MGR_ADDR"
echo "  PROOF_SYSTEM:     $PROOF_SYSTEM_ADDR"
echo "  ROLLUP_ID:        $ROLLUP_ID"
echo "  DEPOSIT:          $DEPOSIT_ADDR"
echo "  WITHDRAWAL:       $WITHDRAWAL_ADDR"
echo "  LISTEN:           $LISTEN"

nohup "$VERIFIER_BIN" \
  -l1 "$L1_RPC" \
  -l2 "http://127.0.0.1:$L2_PORT" \
  -rollups "$ROLLUPS_ADDR" \
  -rollup-mgr "$ROLLUP_MGR_ADDR" \
  -proof-system "$PROOF_SYSTEM_ADDR" \
  -deposit "$DEPOSIT_ADDR" \
  -withdrawal "$WITHDRAWAL_ADDR" \
  -rollup-id "$ROLLUP_ID" \
  -key "$PROVER_KEY" \
  -l1-chain-id "$L1_CHAIN_ID" \
  -l2-chain-id "$L2_CHAIN_ID" \
  -start-l1 "$START_L1_BLOCK" \
  -finalization-depth "$FINALIZATION_DEPTH" \
  -listen "$LISTEN" \
  > /tmp/op_prover_verifier.log 2>&1 &
disown

until curl -s "http://127.0.0.1${LISTEN#:}/status" >/dev/null 2>&1 \
   || curl -s "http://127.0.0.1$LISTEN/status" >/dev/null 2>&1; do
  sleep 1
done
echo "✓ verifier listening on $LISTEN"
echo "  log: /tmp/op_prover_verifier.log"
echo
echo "Send the settler operator your verifier RPC URL so they can wire it in:"
echo "  http://<this-machine-ip>:${LISTEN#:}/"
