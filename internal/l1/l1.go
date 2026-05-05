// Package l1 wraps the L1 (Chiado / Chiado-fork) connection for the
// settler. It exposes:
//   - polling for new DepositGateway events
//   - tx submission
//   - "finalized" head detection (via finality depth from latest)
package l1

import (
	"context"
	"fmt"
	"math/big"
	"sync"
	"time"

	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
	gethrpc "github.com/ethereum/go-ethereum/rpc"

	"github.com/koeppelmann/op-rollups-prover/internal/config"
)

// TransactionDeposited(address indexed from, address indexed to,
//                      uint256 indexed version, bytes opaqueData)
var TransactionDepositedTopic = common.HexToHash("0xb3813568d9991fc951961fcb4c784893574240a28925604d09fc577c55bb7c32")

// computed at init
func init() {
	// keccak256("TransactionDeposited(address,address,uint256,bytes)") =
	// 0xb3813568... — verified at deploy time. We hard-code rather than
	// recomputing to keep the import surface tiny.
}

type Client struct {
	*ethclient.Client
	rpc     *gethrpc.Client
	cfg     *config.Config
	mu      sync.Mutex
	scanned uint64
}

// SetNextBlockTimestamp pins the timestamp of the next L1 block via
// anvil's cheat method. We use this so the settler's pre-computed
// publicInputsHash matches the timestamp the EVM sees when postBatch
// is mined.
func (c *Client) SetNextBlockTimestamp(ctx context.Context, ts uint64) error {
	return c.rpc.CallContext(ctx, nil, "evm_setNextBlockTimestamp", ts)
}

func New(ctx context.Context, cfg *config.Config) (*Client, error) {
	rc, err := gethrpc.DialContext(ctx, cfg.L1RPC)
	if err != nil {
		return nil, fmt.Errorf("dial l1: %w", err)
	}
	c := ethclient.NewClient(rc)
	id, err := c.ChainID(ctx)
	if err != nil {
		return nil, fmt.Errorf("l1 chain id: %w", err)
	}
	if id.Uint64() != cfg.L1ChainID {
		return nil, fmt.Errorf("l1 chain id mismatch: got %d want %d", id.Uint64(), cfg.L1ChainID)
	}
	cl := &Client{Client: c, rpc: rc, cfg: cfg}
	if cfg.StartL1Block != 0 {
		cl.scanned = cfg.StartL1Block - 1
	} else {
		head, err := c.BlockNumber(ctx)
		if err != nil {
			return nil, err
		}
		cl.scanned = head
	}
	return cl, nil
}

type Deposit struct {
	From       common.Address
	To         common.Address
	Mint       *big.Int
	Value      *big.Int
	GasLimit   uint64
	IsCreation bool
	Data       []byte
	L1Block    uint64
	L1TxHash   common.Hash
	LogIdx     uint
}

// PollDeposits scans for new TransactionDeposited events between the last
// scanned head and current head. Cap range to avoid hammering RPC after
// long downtime.
func (c *Client) PollDeposits(ctx context.Context) ([]Deposit, error) {
	c.mu.Lock()
	defer c.mu.Unlock()

	head, err := c.BlockNumber(ctx)
	if err != nil {
		return nil, err
	}
	if head <= c.scanned {
		return nil, nil
	}
	from := c.scanned + 1
	to := head
	const maxRange = 1000
	if to-from > maxRange {
		to = from + maxRange
	}

	q := ethereum.FilterQuery{
		FromBlock: big.NewInt(int64(from)),
		ToBlock:   big.NewInt(int64(to)),
		Addresses: []common.Address{c.cfg.DepositGatewayAddr},
		Topics:    [][]common.Hash{{TransactionDepositedTopic}},
	}
	logs, err := c.FilterLogs(ctx, q)
	if err != nil {
		return nil, fmt.Errorf("filter logs: %w", err)
	}

	out := make([]Deposit, 0, len(logs))
	for _, l := range logs {
		d, err := decodeDeposit(l)
		if err != nil {
			return nil, fmt.Errorf("decode deposit at %s: %w", l.TxHash.Hex(), err)
		}
		out = append(out, d)
	}
	c.scanned = to
	return out, nil
}

// RangeDepositEvents reads TransactionDeposited events from `gw` in
// the L1 block range (from, to]. Used by the canonical derivation
// recipe — anyone with L1 RPC access gets the same list.
func (c *Client) RangeDepositEvents(ctx context.Context, gw common.Address, from, to uint64) ([]Deposit, error) {
	// Range is (from, to] — i.e. blocks from+1..to inclusive. If
	// from >= to the range is empty (settler tick faster than L1
	// produces blocks); skip the RPC call entirely. Real Chiado's RPC
	// returns "invalid params" for fromBlock > toBlock instead of an
	// empty result.
	if from >= to {
		return nil, nil
	}
	q := ethereum.FilterQuery{
		FromBlock: big.NewInt(int64(from + 1)),
		ToBlock:   big.NewInt(int64(to)),
		Addresses: []common.Address{gw},
		Topics:    [][]common.Hash{{TransactionDepositedTopic}},
	}
	logs, err := c.FilterLogs(ctx, q)
	if err != nil {
		return nil, fmt.Errorf("filter logs (%d, %d]: %w", from, to, err)
	}
	out := make([]Deposit, 0, len(logs))
	for _, l := range logs {
		d, err := decodeDeposit(l)
		if err != nil {
			return nil, fmt.Errorf("decode deposit at %s: %w", l.TxHash.Hex(), err)
		}
		out = append(out, d)
	}
	return out, nil
}

func (c *Client) ScannedHead() uint64 {
	c.mu.Lock()
	defer c.mu.Unlock()
	return c.scanned
}

// FinalizedHead returns the L1 block number considered finalized.
// On real Chiado we ask the EL for its `finalized` tag, which the node
// derives from the beacon-chain finalized checkpoint (~12-16 epochs of
// lag, i.e. several minutes on Chiado). On a local anvil fork the
// `finalized` tag isn't tracked, so we fall back to the depth-from-latest
// approximation configured by FinalizationDepth.
func (c *Client) FinalizedHead(ctx context.Context) (uint64, error) {
	var hdr struct {
		Number string `json:"number"`
	}
	err := c.rpc.CallContext(ctx, &hdr, "eth_getBlockByNumber", "finalized", false)
	if err == nil && hdr.Number != "" {
		n, ok := new(big.Int).SetString(hdr.Number[2:], 16)
		if ok {
			return n.Uint64(), nil
		}
	}
	// Fallback: depth-from-latest (anvil / non-finality-aware nodes).
	head, err := c.BlockNumber(ctx)
	if err != nil {
		return 0, err
	}
	if head < c.cfg.FinalizationDepth {
		return 0, nil
	}
	return head - c.cfg.FinalizationDepth, nil
}

// WaitForReceipt returns the receipt for txHash, polling every poll
// interval up to timeout. nil receipt + nil error means timeout.
func (c *Client) WaitForReceipt(ctx context.Context, txHash common.Hash, poll, timeout time.Duration) (*types.Receipt, error) {
	deadline := time.Now().Add(timeout)
	for {
		r, err := c.TransactionReceipt(ctx, txHash)
		if err == nil {
			return r, nil
		}
		if time.Now().After(deadline) {
			return nil, nil
		}
		select {
		case <-ctx.Done():
			return nil, ctx.Err()
		case <-time.After(poll):
		}
	}
}

func decodeDeposit(l types.Log) (Deposit, error) {
	if len(l.Topics) < 4 {
		return Deposit{}, fmt.Errorf("expected 4 topics, got %d", len(l.Topics))
	}
	d := Deposit{
		From:     common.HexToAddress(l.Topics[1].Hex()),
		To:       common.HexToAddress(l.Topics[2].Hex()),
		L1Block:  l.BlockNumber,
		L1TxHash: l.TxHash,
		LogIdx:   l.Index,
	}
	// opaqueData = abi.encodePacked(mint, value, gasLimit, isCreation, data)
	// mint:       uint256 (32 bytes)
	// value:      uint256 (32 bytes)
	// gasLimit:   uint64  (8 bytes)
	// isCreation: bool    (1 byte)
	// data:       bytes   (rest)
	od, err := decodeBytesABI(l.Data)
	if err != nil {
		return d, err
	}
	if len(od) < 32+32+8+1 {
		return d, fmt.Errorf("opaqueData too short: %d", len(od))
	}
	d.Mint = new(big.Int).SetBytes(od[0:32])
	d.Value = new(big.Int).SetBytes(od[32:64])
	d.GasLimit = bigEndianUint64(od[64:72])
	d.IsCreation = od[72] != 0
	d.Data = od[73:]
	return d, nil
}

// decodeBytesABI extracts a single `bytes` value from ABI-encoded
// non-indexed event data: 32 bytes offset, 32 bytes length, raw bytes.
func decodeBytesABI(data []byte) ([]byte, error) {
	if len(data) < 64 {
		return nil, fmt.Errorf("event data too short: %d", len(data))
	}
	length := new(big.Int).SetBytes(data[32:64]).Uint64()
	if uint64(len(data)) < 64+length {
		return nil, fmt.Errorf("event data truncated: have %d, want %d", len(data), 64+length)
	}
	return data[64 : 64+length], nil
}

func bigEndianUint64(b []byte) uint64 {
	if len(b) != 8 {
		return 0
	}
	var v uint64
	for i := 0; i < 8; i++ {
		v = (v << 8) | uint64(b[i])
	}
	return v
}
