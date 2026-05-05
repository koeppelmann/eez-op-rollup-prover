#!/usr/bin/env bash
# Deploy a fresh ECDSAProofSystem on Chiado, owned by your wallet, with
# your wallet as the signer. Print the deployed address — send it back to
# the settler operator so they can register it on the Rollup manager.
#
# Reads (from env or .env file in repo root):
#   PROVER_KEY     - hex private key (with or without 0x prefix)
#   L1_RPC         - default https://rpc.chiadochain.net
set -euo pipefail
ROOT="$(cd "$(dirname "$0")/.." && pwd)"
[[ -f "$ROOT/.env" ]] && source "$ROOT/.env"

L1_RPC="${L1_RPC:-https://rpc.chiadochain.net}"
PROVER_KEY="${PROVER_KEY:?PROVER_KEY required (export it or put in .env)}"

# derive address from the key for display + sanity check
SIGNER=$(cast wallet address --private-key "$PROVER_KEY")
echo "deploying ECDSAProofSystem on Chiado"
echo "  L1_RPC: $L1_RPC"
echo "  signer: $SIGNER"
echo "  owner:  $SIGNER (same as signer)"
echo

BAL=$(cast balance "$SIGNER" --rpc-url "$L1_RPC")
echo "wallet balance: $BAL wei"
if [[ "$BAL" == "0" ]]; then
  echo "ERROR: wallet has 0 xDAI. Fund it from a Chiado faucet:"
  echo "  https://gnosisfaucet.com/  (pick Chiado)"
  echo "  or https://faucet.chiadochain.net/"
  exit 1
fi

# Match the original sync-rollups-v2 build (solc 0.8.34, optimizer=200,
# via_ir=true, evm=prague) so the bytecode is bit-identical and Blockscout
# verification succeeds without metadata-hash drift.
cd "$ROOT/contracts"

# Constructor: (address initialOwner, address initialSigner)
ARGS=$(cast abi-encode "constructor(address,address)" "$SIGNER" "$SIGNER")

# Flat 2 gwei legacy. Chiado SuggestGasPrice can spike during congestion;
# 2 gwei rides through and 1 gwei sometimes hits "replacement underpriced".
forge create ECDSAProofSystem.sol:ECDSAProofSystem --broadcast \
  --rpc-url "$L1_RPC" --private-key "$PROVER_KEY" \
  --gas-price 2000000000 --legacy \
  --constructor-args "$SIGNER" "$SIGNER" \
  | tee /tmp/eps_deploy.txt

PS_ADDR=$(grep -i "Deployed to:" /tmp/eps_deploy.txt | awk '{print $NF}')
echo
echo "===================================================================="
echo "  ECDSAProofSystem: $PS_ADDR"
echo "  signer:           $SIGNER"
echo "===================================================================="
echo
echo "Next steps:"
echo "  1. Send these two values to the settler operator:"
echo "       PROOF_SYSTEM = $PS_ADDR"
echo "       SIGNER       = $SIGNER"
echo "  2. They will call:"
echo "       Rollup(\$ROLLUP_MGR).addProofSystem($PS_ADDR, <vk>)"
echo "       Rollup(\$ROLLUP_MGR).setThreshold(2)"
echo "  3. Once registered, run scripts/start_verifier.sh"
echo
echo "Save your settings to .env so the verifier picks them up:"
echo "  PROOF_SYSTEM_ADDR=$PS_ADDR"
echo
# Append to .env if not already there
if ! grep -q "^PROOF_SYSTEM_ADDR=" "$ROOT/.env" 2>/dev/null; then
  echo "PROOF_SYSTEM_ADDR=$PS_ADDR" >> "$ROOT/.env"
  echo "(appended PROOF_SYSTEM_ADDR to $ROOT/.env)"
fi
