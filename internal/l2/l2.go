// Package l2 wraps the L2 EL connection. We use anvil for the demo; the
// only EL-specific calls are anvil_setBalance (used to materialize
// deposits as L2 ETH) and evm_mine / evm_setNextBlockTimestamp for
// deterministic 2-second cadence.
package l2

import (
	"context"
	"encoding/hex"
	"fmt"
	"math/big"
	"strings"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/ethereum/go-ethereum/rpc"

	"github.com/koeppelmann/op-rollups-prover/internal/config"
)

type Client struct {
	*ethclient.Client
	rpc *rpc.Client
	cfg *config.Config
}

func New(ctx context.Context, cfg *config.Config) (*Client, error) {
	rc, err := rpc.DialContext(ctx, cfg.L2RPC)
	if err != nil {
		return nil, fmt.Errorf("dial l2: %w", err)
	}
	c := ethclient.NewClient(rc)
	id, err := c.ChainID(ctx)
	if err != nil {
		return nil, fmt.Errorf("l2 chain id: %w", err)
	}
	if id.Uint64() != cfg.L2ChainID {
		return nil, fmt.Errorf("l2 chain id mismatch: got %d want %d", id.Uint64(), cfg.L2ChainID)
	}
	return &Client{Client: c, rpc: rc, cfg: cfg}, nil
}

// Snapshot takes an anvil/hardhat-style EVM snapshot. Returns an
// opaque snapshot id that can be passed to Revert. Used by the
// settler to roll back a speculatively-mined L2 block when the
// matching postBatch fails to land on L1.
func (c *Client) Snapshot(ctx context.Context) (string, error) {
	var id string
	if err := c.rpc.CallContext(ctx, &id, "evm_snapshot"); err != nil {
		return "", fmt.Errorf("evm_snapshot: %w", err)
	}
	return id, nil
}

// Revert rewinds the L2 EL to the state captured by Snapshot. Anvil
// invalidates all snapshots taken AFTER the reverted-to one.
func (c *Client) Revert(ctx context.Context, id string) error {
	var ok bool
	if err := c.rpc.CallContext(ctx, &ok, "evm_revert", id); err != nil {
		return fmt.Errorf("evm_revert(%s): %w", id, err)
	}
	if !ok {
		return fmt.Errorf("evm_revert(%s) returned false", id)
	}
	return nil
}

// Mine forces a new block with the given timestamp. We disable
// auto-mining at startup so blocks happen exactly when the settler
// decides — preserving the deterministic cadence the user asked for.
func (c *Client) Mine(ctx context.Context, timestamp uint64) (*types.Header, error) {
	if err := c.rpc.CallContext(ctx, nil, "evm_setNextBlockTimestamp", timestamp); err != nil {
		// Already-set timestamps are not fatal — anvil errors but mining
		// still succeeds with current wall time, which we accept.
		if !strings.Contains(err.Error(), "Timestamp") && !strings.Contains(err.Error(), "TIMESTAMP") {
			return nil, fmt.Errorf("evm_setNextBlockTimestamp(%d): %w", timestamp, err)
		}
	}
	if err := c.rpc.CallContext(ctx, nil, "evm_mine"); err != nil {
		return nil, fmt.Errorf("evm_mine: %w", err)
	}
	return c.HeaderByNumber(ctx, nil)
}

// SetAutoMine toggles anvil's automatic mining. We turn it off so the
// settler is the sole producer of blocks.
func (c *Client) SetAutoMine(ctx context.Context, on bool) error {
	return c.rpc.CallContext(ctx, nil, "anvil_setAutomine", on)
}

// SetIntervalMining sets anvil's interval mode. 0 disables interval
// mining; we use this together with SetAutoMine(false) to ensure only
// our explicit Mine calls produce blocks.
func (c *Client) SetIntervalMining(ctx context.Context, sec uint64) error {
	return c.rpc.CallContext(ctx, nil, "evm_setIntervalMining", sec)
}

// Mint adds ETH directly to an L2 account. Used when an L1 deposit lands
// — the deposit's `mint` field becomes credited L2 balance.
func (c *Client) Mint(ctx context.Context, to common.Address, amount *big.Int) error {
	bal, err := c.BalanceAt(ctx, to, nil)
	if err != nil {
		return err
	}
	newBal := new(big.Int).Add(bal, amount)
	hexBal := "0x" + newBal.Text(16)
	return c.rpc.CallContext(ctx, nil, "anvil_setBalance", to, hexBal)
}

// Burn removes ETH from an L2 account; used during withdrawal so the
// vault on L1 stays in sync with L2 supply.
func (c *Client) Burn(ctx context.Context, from common.Address, amount *big.Int) error {
	bal, err := c.BalanceAt(ctx, from, nil)
	if err != nil {
		return err
	}
	if bal.Cmp(amount) < 0 {
		return fmt.Errorf("insufficient L2 balance: have %s want %s", bal, amount)
	}
	newBal := new(big.Int).Sub(bal, amount)
	hexBal := "0x" + newBal.Text(16)
	return c.rpc.CallContext(ctx, nil, "anvil_setBalance", from, hexBal)
}

// HeadStateRoot returns the post-state root of the latest L2 block.
func (c *Client) HeadStateRoot(ctx context.Context) (common.Hash, uint64, error) {
	h, err := c.HeaderByNumber(ctx, nil)
	if err != nil {
		return common.Hash{}, 0, err
	}
	return h.Root, h.Number.Uint64(), nil
}

// SendDepositCalldata sends a synthetic L2 deposit transaction (raw call
// from `from` to `to` with `value` and `data`). Anvil's
// anvil_impersonateAccount lets us send unsigned txs from arbitrary
// senders, which is how OP-style deposit txs land on L2 in our demo.
func (c *Client) SendDepositCalldata(ctx context.Context, from, to common.Address, value *big.Int, gas uint64, data []byte, isCreation bool) (common.Hash, error) {
	if err := c.rpc.CallContext(ctx, nil, "anvil_impersonateAccount", from); err != nil {
		return common.Hash{}, fmt.Errorf("impersonate: %w", err)
	}
	defer func() {
		_ = c.rpc.CallContext(context.Background(), nil, "anvil_stopImpersonatingAccount", from)
	}()
	args := map[string]any{
		"from":  from,
		"value": "0x" + value.Text(16),
		"gas":   fmt.Sprintf("0x%x", gas),
	}
	if !isCreation {
		args["to"] = to
	}
	if len(data) > 0 {
		args["data"] = "0x" + hex.EncodeToString(data)
	}
	var h common.Hash
	if err := c.rpc.CallContext(ctx, &h, "eth_sendTransaction", args); err != nil {
		return common.Hash{}, fmt.Errorf("send: %w", err)
	}
	return h, nil
}

// RPCClient exposes the underlying rpc client for the RPC frontend.
func (c *Client) RPCClient() *rpc.Client { return c.rpc }

// PoolPending returns the number of pending txs in anvil's mempool.
// Used by the sequencer to decide whether the next 2s tick has any
// real work (saves a mine when idle, keeping verifier in sync).
func (c *Client) PoolPending(ctx context.Context) (int, error) {
	var status struct {
		Pending string `json:"pending"`
		Queued  string `json:"queued"`
	}
	if err := c.rpc.CallContext(ctx, &status, "txpool_status"); err != nil {
		return 0, err
	}
	parse := func(h string) int {
		if h == "" {
			return 0
		}
		var n int
		_, _ = fmt.Sscanf(h, "0x%x", &n)
		return n
	}
	return parse(status.Pending) + parse(status.Queued), nil
}

// SendRawTx submits a pre-signed raw transaction. Used by the verifier
// when replaying L2 batches and (potentially) by the settler too.
func (c *Client) SendRawTx(ctx context.Context, raw []byte) (common.Hash, error) {
	var h common.Hash
	if err := c.rpc.CallContext(ctx, &h, "eth_sendRawTransaction", "0x"+hexEncode(raw)); err != nil {
		return common.Hash{}, err
	}
	return h, nil
}

func hexEncode(b []byte) string {
	const hexChars = "0123456789abcdef"
	out := make([]byte, len(b)*2)
	for i, v := range b {
		out[2*i] = hexChars[v>>4]
		out[2*i+1] = hexChars[v&0x0f]
	}
	return string(out)
}
