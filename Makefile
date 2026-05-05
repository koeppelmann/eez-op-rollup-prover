.PHONY: build setup-contracts deploy-ps start-anvil start-verifier status stop clean

BIN := bin/op-prover-verifier

build:
	@mkdir -p bin
	go build -o $(BIN) ./cmd/verifier
	@echo "✓ built $(BIN)"

setup-contracts:
	bash scripts/setup_contracts.sh

deploy-ps:
	bash scripts/deploy_proof_system.sh

start-anvil:
	bash scripts/start_l2_anvil.sh

start-verifier: build
	bash scripts/start_verifier.sh

status:
	@echo "=== anvil ===" && curl -s http://127.0.0.1:$${L2_PORT:-8547} -X POST -H "Content-Type: application/json" --data '{"jsonrpc":"2.0","method":"eth_blockNumber","params":[],"id":1}' | jq . || echo "anvil down"
	@echo "=== verifier ===" && curl -s http://127.0.0.1:9100/status | jq . || echo "verifier down"

check-registration:
	bash scripts/check_registration.sh

stop:
	@pkill -f "$(BIN)" 2>/dev/null || true
	@pkill -f "anvil --port $${L2_PORT:-8547}" 2>/dev/null || true
	@echo "✓ stopped"

clean:
	rm -rf bin/ contracts/out/ contracts/cache/
