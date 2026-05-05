#!/usr/bin/env bash
# Convenience: check whether the settler operator has registered our
# proof system in the Rollup manager. Until they do, start_verifier.sh
# will fail at boot with "manager reports zero vkey for PS …".
set -euo pipefail
ROOT="$(cd "$(dirname "$0")/.." && pwd)"
[[ -f "$ROOT/.env" ]] && source "$ROOT/.env"

L1_RPC="${L1_RPC:-https://rpc.chiadochain.net}"
: "${ROLLUP_MGR_ADDR:?ROLLUP_MGR_ADDR required}"
: "${PROOF_SYSTEM_ADDR:?PROOF_SYSTEM_ADDR required (run deploy_proof_system.sh first)}"

vk=$(cast call "$ROLLUP_MGR_ADDR" "verificationKey(address)(bytes32)" \
  "$PROOF_SYSTEM_ADDR" --rpc-url "$L1_RPC")
threshold=$(cast call "$ROLLUP_MGR_ADDR" "threshold()(uint256)" --rpc-url "$L1_RPC")

echo "manager:    $ROLLUP_MGR_ADDR"
echo "PS:         $PROOF_SYSTEM_ADDR"
echo "vk:         $vk"
echo "threshold:  $threshold"
echo

if [[ "$vk" == "0x0000000000000000000000000000000000000000000000000000000000000000" ]]; then
  echo "❌ NOT REGISTERED — ask the settler operator to call:"
  echo "   Rollup($ROLLUP_MGR_ADDR).addProofSystem($PROOF_SYSTEM_ADDR, <vk>)"
else
  echo "✅ registered with vkey $vk"
fi
