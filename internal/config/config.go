package config

import (
	"flag"
	"os"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
)

type Config struct {
	L1RPC               string
	L2RPC               string
	RPCAddr             string
	RollupsAddr         common.Address
	DepositGatewayAddr  common.Address
	WithdrawalAddr      common.Address
	// Multi-prover layout (sync-rollups-v2 / PR-22): the registry no
	// longer pins a single ECDSA verifier; each rollup's manager picks
	// proof systems. We carry both the per-rollup manager (so we can
	// look up the vkey for our PS) and the proof system address (so
	// our ProofSystemBatch points at the right one).
	RollupMgrAddr      common.Address
	ProofSystemAddr    common.Address
	RollupID            uint64
	BlockTimeSec        int
	L1BlockTimeSec      uint64 // expected L1 block time in seconds; used to predict block.timestamp at postBatch landing
	SignerKeyHex        string
	L1ChainID           uint64
	L2ChainID           uint64
	StartL1Block        uint64
	FinalizationDepth   uint64 // L1 blocks before considered "finalized"

	// Verifier wiring. If VerifierRPC is non-empty, the settler does not
	// hold the signing key — it sends candidate batches to the verifier
	// over JSON-RPC and the verifier returns the signature only if its
	// independent re-derivation matches.
	VerifierRPC string
}

func LoadFromFlags() *Config {
	c := &Config{}
	flag.StringVar(&c.L1RPC, "l1", env("L1_RPC", "http://localhost:8545"), "L1 RPC")
	flag.StringVar(&c.L2RPC, "l2", env("L2_RPC", "http://localhost:8546"), "L2 RPC (anvil)")
	flag.StringVar(&c.RPCAddr, "rpc", env("RPC_ADDR", ":9000"), "L2 RPC frontend listen addr")

	// No default for ROLLUPS_ADDR — the prover gets it from the settler
	// operator and writes it into .env. Hard-coding the previous demo's
	// Rollups address here would silently mis-target a different rollup
	// if/when the contracts get redeployed.
	flagStringVarHex(&c.RollupsAddr, "rollups", env("ROLLUPS_ADDR", ""), "Rollups registry on L1")

	flagStringVarHex(&c.DepositGatewayAddr, "deposit", env("DEPOSIT_ADDR", ""), "DepositGateway addr on L1")
	flagStringVarHex(&c.WithdrawalAddr, "withdrawal", env("WITHDRAWAL_ADDR", ""), "WithdrawalGateway addr on L1")
	flagStringVarHex(&c.RollupMgrAddr, "rollup-mgr", env("ROLLUP_MGR_ADDR", ""), "Rollup manager (IRollup) for our rollupId")
	flagStringVarHex(&c.ProofSystemAddr, "proof-system", env("ECDSA_PS_ADDR", ""), "ECDSAProofSystem addr (registered in the manager)")

	flag.Uint64Var(&c.RollupID, "rollup-id", envUint("ROLLUP_ID", 0), "rollupId we own in Rollups")
	flag.IntVar(&c.BlockTimeSec, "block-time", int(envUint("BLOCK_TIME_SEC", 2)), "L2 block time in seconds")
	flag.Uint64Var(&c.L1BlockTimeSec, "l1-block-time", envUint("L1_BLOCK_TIME_SEC", 5), "expected L1 block time in seconds (5 for Chiado)")

	// No default key — the prover MUST supply its own signing key via
	// -key, $SIGNER_KEY, or $PROVER_KEY. Refusing to start is the right
	// failure mode here; we don't want to ever fall back to a baked-in
	// dev key (which would silently produce signatures that recover to
	// some shared address, defeating the operational-separation point
	// of running an independent prover).
	keyDefault := env("PROVER_KEY", env("SIGNER_KEY", ""))
	flag.StringVar(&c.SignerKeyHex, "key", keyDefault, "verifier signing key (hex; 0x-prefixed). Required.")
	flag.Uint64Var(&c.L1ChainID, "l1-chain-id", envUint("L1_CHAIN_ID", 10200), "L1 chain id")
	flag.Uint64Var(&c.L2ChainID, "l2-chain-id", envUint("L2_CHAIN_ID", 4242), "L2 chain id")
	flag.Uint64Var(&c.StartL1Block, "start-l1", envUint("START_L1_BLOCK", 0), "L1 block to begin scanning from (0 = current head)")
	flag.Uint64Var(&c.FinalizationDepth, "finalization-depth", envUint("FINALIZATION_DEPTH", 12), "L1 blocks before postBatch is considered finalized")
	flag.StringVar(&c.VerifierRPC, "verifier-rpc", env("VERIFIER_RPC", ""), "JSON-RPC endpoint of op-verifier (empty = legacy local-signing mode)")
	return c
}

func (c *Config) SignerAddress() common.Address {
	pk, err := crypto.HexToECDSA(stripHex(c.SignerKeyHex))
	if err != nil {
		return common.Address{}
	}
	return crypto.PubkeyToAddress(pk.PublicKey)
}

func env(k, def string) string {
	if v, ok := os.LookupEnv(k); ok {
		return v
	}
	return def
}

func envUint(k string, def uint64) uint64 {
	if v, ok := os.LookupEnv(k); ok {
		var u uint64
		_, err := fmtSscanf(v, &u)
		if err == nil {
			return u
		}
	}
	return def
}

func fmtSscanf(s string, u *uint64) (int, error) {
	var n int
	for i := 0; i < len(s); i++ {
		c := s[i]
		if c < '0' || c > '9' {
			break
		}
		*u = *u*10 + uint64(c-'0')
		n++
	}
	if n == 0 {
		return 0, errParse
	}
	return n, nil
}

func stripHex(s string) string {
	if len(s) >= 2 && (s[:2] == "0x" || s[:2] == "0X") {
		return s[2:]
	}
	return s
}

func flagStringVarHex(p *common.Address, name, def, usage string) {
	wrapper := &addrFlag{p: p}
	if def != "" {
		*p = common.HexToAddress(def)
	}
	flag.Var(wrapper, name, usage)
}

type addrFlag struct{ p *common.Address }

func (a *addrFlag) String() string {
	if a.p == nil {
		return ""
	}
	return a.p.Hex()
}

func (a *addrFlag) Set(s string) error {
	*a.p = common.HexToAddress(s)
	return nil
}

type errString string

func (e errString) Error() string { return string(e) }

const errParse = errString("parse error")
