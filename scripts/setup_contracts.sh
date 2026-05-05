#!/usr/bin/env bash
# One-shot foundry setup. Installs OpenZeppelin contracts (used by
# ECDSAProofSystem) and builds the proof-system artifact so it's ready
# for `deploy_proof_system.sh`.
#
# Run once after cloning the repo. Idempotent.
set -euo pipefail
cd "$(dirname "$0")/../contracts"

if [[ ! -d lib/openzeppelin-contracts ]]; then
  forge install OpenZeppelin/openzeppelin-contracts --no-commit --no-git
fi

# remappings so `import {ECDSA} from "@openzeppelin/..."` resolves
cat > remappings.txt <<EOF
@openzeppelin/contracts/=lib/openzeppelin-contracts/contracts/
EOF

forge build
echo "✓ contracts ready"
