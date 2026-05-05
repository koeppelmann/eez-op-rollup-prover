#!/usr/bin/env bash
# Start the prover's local L2 EL. Anvil with --no-mining so only our
# verifier produces blocks (via signBatch).
#
# Default port: 8547 (mirrors the settler-side verifier convention).
# Default chain id: 4242 (must match what the settler uses; if the
# settler operator picked something else, override via L2_CHAIN_ID env).
set -euo pipefail
ROOT="$(cd "$(dirname "$0")/.." && pwd)"
[[ -f "$ROOT/.env" ]] && source "$ROOT/.env"

L2_PORT="${L2_PORT:-8547}"
L2_CHAIN_ID="${L2_CHAIN_ID:-4242}"

if pgrep -f "anvil --port ${L2_PORT}" >/dev/null; then
  echo "anvil already running on :${L2_PORT}"
  exit 0
fi

nohup anvil \
  --port "$L2_PORT" \
  --chain-id "$L2_CHAIN_ID" \
  --no-mining \
  --order fifo \
  --silent \
  > /tmp/l2_anvil.log 2>&1 &
disown
until cast chain-id --rpc-url "http://127.0.0.1:${L2_PORT}" >/dev/null 2>&1; do sleep 1; done
echo "✓ L2 anvil up: http://127.0.0.1:${L2_PORT}  chainId=${L2_CHAIN_ID}"
echo "  log: /tmp/l2_anvil.log"

# Pre-fund our signer + the deployer wallet on this EL so re-derived
# user txs that use those addresses don't fail with insufficient funds.
# Settler operator should send the SAME list of wallets to pre-fund —
# any divergence will surface immediately as a STATE-ROOT MISMATCH.
if [[ -n "${PROVER_KEY:-}" ]]; then
  PROVER_ADDR=$(cast wallet address --private-key "$PROVER_KEY")
  cast rpc anvil_setBalance "$PROVER_ADDR" 0x21E19E0C9BAB2400000 \
    --rpc-url "http://127.0.0.1:${L2_PORT}" >/dev/null
  echo "  funded prover signer $PROVER_ADDR"
fi

# The settler operator's wallet list is shared via PREFUND_WALLETS (CSV).
# Without this, txs from those wallets fail in re-derivation.
if [[ -n "${PREFUND_WALLETS:-}" ]]; then
  IFS=',' read -ra WALLETS <<< "$PREFUND_WALLETS"
  for w in "${WALLETS[@]}"; do
    cast rpc anvil_setBalance "$w" 0x21E19E0C9BAB2400000 \
      --rpc-url "http://127.0.0.1:${L2_PORT}" >/dev/null
    echo "  funded $w"
  done
fi
