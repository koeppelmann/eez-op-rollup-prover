// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package withdrawal

import (
	"errors"
	"math/big"
	"strings"

	ethereum "github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/event"
)

// Reference imports to suppress errors if they are not otherwise used.
var (
	_ = errors.New
	_ = big.NewInt
	_ = strings.NewReader
	_ = ethereum.NotFound
	_ = bind.Bind
	_ = common.Big1
	_ = types.BloomLookup
	_ = event.NewSubscription
	_ = abi.ConvertType
)

// WithdrawalGatewayOutputRootProof is an auto generated low-level Go binding around an user-defined struct.
type WithdrawalGatewayOutputRootProof struct {
	Version                  [32]byte
	StateRoot                [32]byte
	MessagePasserStorageRoot [32]byte
	LatestBlockhash          [32]byte
}

// WithdrawalGatewayWithdrawalTransaction is an auto generated low-level Go binding around an user-defined struct.
type WithdrawalGatewayWithdrawalTransaction struct {
	Nonce    *big.Int
	Sender   common.Address
	Target   common.Address
	Value    *big.Int
	GasLimit *big.Int
	Data     []byte
}

// WithdrawalGatewayMetaData contains all meta data concerning the WithdrawalGateway contract.
var WithdrawalGatewayMetaData = &bind.MetaData{
	ABI: "[{\"type\":\"constructor\",\"inputs\":[{\"name\":\"rollups\",\"type\":\"address\",\"internalType\":\"contractIRollups\"},{\"name\":\"rollupId\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"mptVerifier\",\"type\":\"address\",\"internalType\":\"contractIMPTVerifier\"},{\"name\":\"depositGateway\",\"type\":\"address\",\"internalType\":\"contractDepositGateway\"},{\"name\":\"finalizationPeriodSeconds\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"receive\",\"stateMutability\":\"payable\"},{\"type\":\"function\",\"name\":\"DEPOSIT_GATEWAY\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"contractDepositGateway\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"FINALIZATION_PERIOD_SECONDS\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"L2_TO_L1_MESSAGE_PASSER\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"MPT_VERIFIER\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"contractIMPTVerifier\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"ROLLUPS\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"contractIRollups\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"ROLLUP_ID\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"SENT_MESSAGES_SLOT\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"finalizeWithdrawalTransaction\",\"inputs\":[{\"name\":\"_tx\",\"type\":\"tuple\",\"internalType\":\"structWithdrawalGateway.WithdrawalTransaction\",\"components\":[{\"name\":\"nonce\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"sender\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"target\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"value\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"gasLimit\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"data\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"finalized\",\"inputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"hashWithdrawal\",\"inputs\":[{\"name\":\"_tx\",\"type\":\"tuple\",\"internalType\":\"structWithdrawalGateway.WithdrawalTransaction\",\"components\":[{\"name\":\"nonce\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"sender\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"target\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"value\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"gasLimit\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"data\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]}],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"pure\"},{\"type\":\"function\",\"name\":\"proveWithdrawalTransaction\",\"inputs\":[{\"name\":\"_tx\",\"type\":\"tuple\",\"internalType\":\"structWithdrawalGateway.WithdrawalTransaction\",\"components\":[{\"name\":\"nonce\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"sender\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"target\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"value\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"gasLimit\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"data\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]},{\"name\":\"_outputRootProof\",\"type\":\"tuple\",\"internalType\":\"structWithdrawalGateway.OutputRootProof\",\"components\":[{\"name\":\"version\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"stateRoot\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"messagePasserStorageRoot\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"latestBlockhash\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}]},{\"name\":\"_accountProof\",\"type\":\"bytes[]\",\"internalType\":\"bytes[]\"},{\"name\":\"_withdrawalProof\",\"type\":\"bytes[]\",\"internalType\":\"bytes[]\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"provenAt\",\"inputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"event\",\"name\":\"WithdrawalFinalized\",\"inputs\":[{\"name\":\"withdrawalHash\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"},{\"name\":\"success\",\"type\":\"bool\",\"indexed\":false,\"internalType\":\"bool\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"WithdrawalProven\",\"inputs\":[{\"name\":\"withdrawalHash\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"},{\"name\":\"from\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"to\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"error\",\"name\":\"FinalizationPeriodNotElapsed\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidAccountProof\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidWithdrawalProof\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"StateRootMismatch\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"WithdrawalAlreadyFinalized\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"WithdrawalNotProven\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"ZeroStateRoot\",\"inputs\":[]}]",
	Bin: "0x610120604052348015610010575f80fd5b506040516110b53803806110b583398101604081905261002f9161006f565b6001600160a01b0394851660a05260c09390935290831660e052909116610100526080526100cc565b6001600160a01b038116811461006c575f80fd5b50565b5f805f805f60a08688031215610083575f80fd5b855161008e81610058565b6020870151604088015191965094506100a681610058565b60608701519093506100b781610058565b80925050608086015190509295509295909350565b60805160a05160c05160e05161010051610f7c6101395f395f818160b8015281816102ee015261048a01525f81816101370152818161071801526108fc01525f818161024401526105cb01525f81816101c501526105fe01525f818161032101526104010152610f7c5ff3fe6080604052600436106100a8575f3560e01c8063b84c278611610062578063b84c278614610266578063c080abef14610291578063c91ac182146102b0578063d5f3204b146102c3578063ea56249c146102dd578063f4daa29114610310575f80fd5b8063053d3037146101265780630abea268146101765780632ef01a75146101b45780637d4395ac146101e75780638c3152e914610214578063b5ed555914610233575f80fd5b3661012257336001600160a01b037f000000000000000000000000000000000000000000000000000000000000000016146101205760405162461bcd60e51b81526020600482015260146024820152736f6e6c79206465706f736974206761746577617960601b604482015260640160405180910390fd5b005b5f80fd5b348015610131575f80fd5b506101597f000000000000000000000000000000000000000000000000000000000000000081565b6040516001600160a01b0390911681526020015b60405180910390f35b348015610181575f80fd5b506101a46101903660046109e5565b60016020525f908152604090205460ff1681565b604051901515815260200161016d565b3480156101bf575f80fd5b506101597f000000000000000000000000000000000000000000000000000000000000000081565b3480156101f2575f80fd5b50610206610201366004610b60565b610343565b60405190815260200161016d565b34801561021f575f80fd5b5061012061022e366004610bb0565b61038f565b34801561023e575f80fd5b506102067f000000000000000000000000000000000000000000000000000000000000000081565b348015610271575f80fd5b506102066102803660046109e5565b5f6020819052908152604090205481565b34801561029c575f80fd5b506101206102ab366004610c2a565b6105af565b3480156102bb575f80fd5b506102065f81565b3480156102ce575f80fd5b506101596016602160991b0181565b3480156102e8575f80fd5b506101597f000000000000000000000000000000000000000000000000000000000000000081565b34801561031b575f80fd5b506102067f000000000000000000000000000000000000000000000000000000000000000081565b80516020808301516040808501516060860151608087015160a088015193515f97610372979096959101610d19565b604051602081830303815290604052805190602001209050919050565b5f61039c61020183610d58565b5f818152602081905260408120549192508190036103cd576040516367ba7fe160e11b815260040160405180910390fd5b5f8281526001602052604090205460ff16156103fc57604051632ba2651560e21b815260040160405180910390fd5b6104267f000000000000000000000000000000000000000000000000000000000000000082610d69565b42101561044657604051633a629fd560e21b815260040160405180910390fd5b5f828152600160208190526040909120805460ff191690911790556060830135156104ea57604051638091ed5360e01b8152306004820152606084013560248201527f00000000000000000000000000000000000000000000000000000000000000006001600160a01b031690638091ed53906044015f604051808303815f87803b1580156104d3575f80fd5b505af11580156104e5573d5f803e3d5ffd5b505050505b5f6104fb6060850160408601610d88565b6001600160a01b03166080850135606086013561051b60a0880188610daa565b604051610529929190610ded565b5f60405180830381858888f193505050503d805f8114610564576040519150601f19603f3d011682016040523d82523d5f602084013e610569565b606091505b50509050827fdb5c7652857aa163daadd670e116628fb42e869d8ac4251ef8971d9e5727df1b826040516105a1911515815260200190565b60405180910390a250505050565b5f6105bc61020188610d58565b60405163b794e5a360e01b81527f000000000000000000000000000000000000000000000000000000000000000060048201529091505f906001600160a01b037f0000000000000000000000000000000000000000000000000000000000000000169063b794e5a390602401608060405180830381865afa158015610643573d5f803e3d5ffd5b505050506040513d601f19601f820116820180604052508101906106679190610dfc565b50925082915061068c905057604051637eebc38360e01b815260040160405180910390fd5b808760200135146106b057604051631b2075cd60e01b815260040160405180910390fd5b6106cc87602001356016602160991b0189604001358989610871565b6106e957604051631052dd2560e21b815260040160405180910390fd5b60408051602081018490525f9181018290526060016040516020818303038152906040528051906020012090507f00000000000000000000000000000000000000000000000000000000000000006001600160a01b031663c20a15fb8260405160200161075891815260200190565b60408051601f1981840301815291905261077260016109ba565b88888d604001356040518663ffffffff1660e01b8152600401610799959493929190610e60565b602060405180830381865afa1580156107b4573d5f803e3d5ffd5b505050506040513d601f19601f820116820180604052508101906107d89190610f27565b6107f557604051632e1aaf2760e21b815260040160405180910390fd5b5f83815260208190526040908190204290556108179060608b01908b01610d88565b6001600160a01b031661083060408b0160208c01610d88565b6001600160a01b0316847f67a6208cfcc0801d50f6cbe764733f4fddf66ac0b04442061a8a8c0cb6b63f6260405160405180910390a4505050505050505050565b60408051606086901b6bffffffffffffffffffffffff19166020808301919091528251601481840301815260348301845280519101205f6054830181905260748301819052609483018790527fc5d2460186f7233c927e7db2dcc703c0e500b653ca82273b7bfad8045d85a47060b4808501919091528451808503909101815260d4909301909352907f00000000000000000000000000000000000000000000000000000000000000006001600160a01b031663c20a15fb8360405160200161093c91815260200190565b6040516020818303038152906040528388888d6040518663ffffffff1660e01b815260040161096f959493929190610e60565b602060405180830381865afa15801561098a573d5f803e3d5ffd5b505050506040513d601f19601f820116820180604052508101906109ae9190610f27565b98975050505050505050565b6060816040516020016109cf91815260200190565b6040516020818303038152906040529050919050565b5f602082840312156109f5575f80fd5b5035919050565b634e487b7160e01b5f52604160045260245ffd5b60405160c0810167ffffffffffffffff81118282101715610a3357610a336109fc565b60405290565b604051601f8201601f1916810167ffffffffffffffff81118282101715610a6257610a626109fc565b604052919050565b6001600160a01b0381168114610a7e575f80fd5b50565b5f60c08284031215610a91575f80fd5b610a99610a10565b905081358152602080830135610aae81610a6a565b828201526040830135610ac081610a6a565b80604084015250606083013560608301526080830135608083015260a083013567ffffffffffffffff80821115610af5575f80fd5b818501915085601f830112610b08575f80fd5b813581811115610b1a57610b1a6109fc565b610b2c601f8201601f19168501610a39565b91508082528684828501011115610b41575f80fd5b80848401858401375f848284010152508060a085015250505092915050565b5f60208284031215610b70575f80fd5b813567ffffffffffffffff811115610b86575f80fd5b610b9284828501610a81565b949350505050565b5f60c08284031215610baa575f80fd5b50919050565b5f60208284031215610bc0575f80fd5b813567ffffffffffffffff811115610bd6575f80fd5b610b9284828501610b9a565b5f8083601f840112610bf2575f80fd5b50813567ffffffffffffffff811115610c09575f80fd5b6020830191508360208260051b8501011115610c23575f80fd5b9250929050565b5f805f805f8086880360e0811215610c40575f80fd5b873567ffffffffffffffff80821115610c57575f80fd5b610c638b838c01610b9a565b98506080601f1984011215610c76575f80fd5b60208a01975060a08a0135925080831115610c8f575f80fd5b610c9b8b848c01610be2565b909750955060c08a0135925086915080831115610cb6575f80fd5b5050610cc489828a01610be2565b979a9699509497509295939492505050565b5f81518084525f5b81811015610cfa57602081850181015186830182015201610cde565b505f602082860101526020601f19601f83011685010191505092915050565b8681526001600160a01b03868116602083015285166040820152606081018490526080810183905260c060a082018190525f906109ae90830184610cd6565b5f610d633683610a81565b92915050565b80820180821115610d6357634e487b7160e01b5f52601160045260245ffd5b5f60208284031215610d98575f80fd5b8135610da381610a6a565b9392505050565b5f808335601e19843603018112610dbf575f80fd5b83018035915067ffffffffffffffff821115610dd9575f80fd5b602001915036819003821315610c23575f80fd5b818382375f9101908152919050565b5f805f8060808587031215610e0f575f80fd5b8451610e1a81610a6a565b60208601516040870151606090970151919890975090945092505050565b81835281816020850137505f828201602090810191909152601f909101601f19169091010190565b608081525f610e726080830188610cd6565b602083820381850152610e858289610cd6565b84810360408601528681529150808201600587901b83018201885f5b89811015610f0e57858303601f190184528135368c9003601e19018112610ec6575f80fd5b8b01858101903567ffffffffffffffff811115610ee1575f80fd5b803603821315610eef575f80fd5b610efa858284610e38565b958701959450505090840190600101610ea1565b5050809450505050508260608301529695505050505050565b5f60208284031215610f37575f80fd5b81518015158114610da3575f80fdfea26469706673582212207a5079b07a94875a7c76029a21acb1c8907847a05c698677a0017152238dcfc364736f6c63430008180033",
}

// WithdrawalGatewayABI is the input ABI used to generate the binding from.
// Deprecated: Use WithdrawalGatewayMetaData.ABI instead.
var WithdrawalGatewayABI = WithdrawalGatewayMetaData.ABI

// WithdrawalGatewayBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use WithdrawalGatewayMetaData.Bin instead.
var WithdrawalGatewayBin = WithdrawalGatewayMetaData.Bin

// DeployWithdrawalGateway deploys a new Ethereum contract, binding an instance of WithdrawalGateway to it.
func DeployWithdrawalGateway(auth *bind.TransactOpts, backend bind.ContractBackend, rollups common.Address, rollupId *big.Int, mptVerifier common.Address, depositGateway common.Address, finalizationPeriodSeconds *big.Int) (common.Address, *types.Transaction, *WithdrawalGateway, error) {
	parsed, err := WithdrawalGatewayMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(WithdrawalGatewayBin), backend, rollups, rollupId, mptVerifier, depositGateway, finalizationPeriodSeconds)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &WithdrawalGateway{WithdrawalGatewayCaller: WithdrawalGatewayCaller{contract: contract}, WithdrawalGatewayTransactor: WithdrawalGatewayTransactor{contract: contract}, WithdrawalGatewayFilterer: WithdrawalGatewayFilterer{contract: contract}}, nil
}

// WithdrawalGateway is an auto generated Go binding around an Ethereum contract.
type WithdrawalGateway struct {
	WithdrawalGatewayCaller     // Read-only binding to the contract
	WithdrawalGatewayTransactor // Write-only binding to the contract
	WithdrawalGatewayFilterer   // Log filterer for contract events
}

// WithdrawalGatewayCaller is an auto generated read-only Go binding around an Ethereum contract.
type WithdrawalGatewayCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// WithdrawalGatewayTransactor is an auto generated write-only Go binding around an Ethereum contract.
type WithdrawalGatewayTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// WithdrawalGatewayFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type WithdrawalGatewayFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// WithdrawalGatewaySession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type WithdrawalGatewaySession struct {
	Contract     *WithdrawalGateway // Generic contract binding to set the session for
	CallOpts     bind.CallOpts      // Call options to use throughout this session
	TransactOpts bind.TransactOpts  // Transaction auth options to use throughout this session
}

// WithdrawalGatewayCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type WithdrawalGatewayCallerSession struct {
	Contract *WithdrawalGatewayCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts            // Call options to use throughout this session
}

// WithdrawalGatewayTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type WithdrawalGatewayTransactorSession struct {
	Contract     *WithdrawalGatewayTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts            // Transaction auth options to use throughout this session
}

// WithdrawalGatewayRaw is an auto generated low-level Go binding around an Ethereum contract.
type WithdrawalGatewayRaw struct {
	Contract *WithdrawalGateway // Generic contract binding to access the raw methods on
}

// WithdrawalGatewayCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type WithdrawalGatewayCallerRaw struct {
	Contract *WithdrawalGatewayCaller // Generic read-only contract binding to access the raw methods on
}

// WithdrawalGatewayTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type WithdrawalGatewayTransactorRaw struct {
	Contract *WithdrawalGatewayTransactor // Generic write-only contract binding to access the raw methods on
}

// NewWithdrawalGateway creates a new instance of WithdrawalGateway, bound to a specific deployed contract.
func NewWithdrawalGateway(address common.Address, backend bind.ContractBackend) (*WithdrawalGateway, error) {
	contract, err := bindWithdrawalGateway(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &WithdrawalGateway{WithdrawalGatewayCaller: WithdrawalGatewayCaller{contract: contract}, WithdrawalGatewayTransactor: WithdrawalGatewayTransactor{contract: contract}, WithdrawalGatewayFilterer: WithdrawalGatewayFilterer{contract: contract}}, nil
}

// NewWithdrawalGatewayCaller creates a new read-only instance of WithdrawalGateway, bound to a specific deployed contract.
func NewWithdrawalGatewayCaller(address common.Address, caller bind.ContractCaller) (*WithdrawalGatewayCaller, error) {
	contract, err := bindWithdrawalGateway(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &WithdrawalGatewayCaller{contract: contract}, nil
}

// NewWithdrawalGatewayTransactor creates a new write-only instance of WithdrawalGateway, bound to a specific deployed contract.
func NewWithdrawalGatewayTransactor(address common.Address, transactor bind.ContractTransactor) (*WithdrawalGatewayTransactor, error) {
	contract, err := bindWithdrawalGateway(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &WithdrawalGatewayTransactor{contract: contract}, nil
}

// NewWithdrawalGatewayFilterer creates a new log filterer instance of WithdrawalGateway, bound to a specific deployed contract.
func NewWithdrawalGatewayFilterer(address common.Address, filterer bind.ContractFilterer) (*WithdrawalGatewayFilterer, error) {
	contract, err := bindWithdrawalGateway(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &WithdrawalGatewayFilterer{contract: contract}, nil
}

// bindWithdrawalGateway binds a generic wrapper to an already deployed contract.
func bindWithdrawalGateway(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := WithdrawalGatewayMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_WithdrawalGateway *WithdrawalGatewayRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _WithdrawalGateway.Contract.WithdrawalGatewayCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_WithdrawalGateway *WithdrawalGatewayRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _WithdrawalGateway.Contract.WithdrawalGatewayTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_WithdrawalGateway *WithdrawalGatewayRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _WithdrawalGateway.Contract.WithdrawalGatewayTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_WithdrawalGateway *WithdrawalGatewayCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _WithdrawalGateway.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_WithdrawalGateway *WithdrawalGatewayTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _WithdrawalGateway.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_WithdrawalGateway *WithdrawalGatewayTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _WithdrawalGateway.Contract.contract.Transact(opts, method, params...)
}

// DEPOSITGATEWAY is a free data retrieval call binding the contract method 0xea56249c.
//
// Solidity: function DEPOSIT_GATEWAY() view returns(address)
func (_WithdrawalGateway *WithdrawalGatewayCaller) DEPOSITGATEWAY(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _WithdrawalGateway.contract.Call(opts, &out, "DEPOSIT_GATEWAY")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// DEPOSITGATEWAY is a free data retrieval call binding the contract method 0xea56249c.
//
// Solidity: function DEPOSIT_GATEWAY() view returns(address)
func (_WithdrawalGateway *WithdrawalGatewaySession) DEPOSITGATEWAY() (common.Address, error) {
	return _WithdrawalGateway.Contract.DEPOSITGATEWAY(&_WithdrawalGateway.CallOpts)
}

// DEPOSITGATEWAY is a free data retrieval call binding the contract method 0xea56249c.
//
// Solidity: function DEPOSIT_GATEWAY() view returns(address)
func (_WithdrawalGateway *WithdrawalGatewayCallerSession) DEPOSITGATEWAY() (common.Address, error) {
	return _WithdrawalGateway.Contract.DEPOSITGATEWAY(&_WithdrawalGateway.CallOpts)
}

// FINALIZATIONPERIODSECONDS is a free data retrieval call binding the contract method 0xf4daa291.
//
// Solidity: function FINALIZATION_PERIOD_SECONDS() view returns(uint256)
func (_WithdrawalGateway *WithdrawalGatewayCaller) FINALIZATIONPERIODSECONDS(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _WithdrawalGateway.contract.Call(opts, &out, "FINALIZATION_PERIOD_SECONDS")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// FINALIZATIONPERIODSECONDS is a free data retrieval call binding the contract method 0xf4daa291.
//
// Solidity: function FINALIZATION_PERIOD_SECONDS() view returns(uint256)
func (_WithdrawalGateway *WithdrawalGatewaySession) FINALIZATIONPERIODSECONDS() (*big.Int, error) {
	return _WithdrawalGateway.Contract.FINALIZATIONPERIODSECONDS(&_WithdrawalGateway.CallOpts)
}

// FINALIZATIONPERIODSECONDS is a free data retrieval call binding the contract method 0xf4daa291.
//
// Solidity: function FINALIZATION_PERIOD_SECONDS() view returns(uint256)
func (_WithdrawalGateway *WithdrawalGatewayCallerSession) FINALIZATIONPERIODSECONDS() (*big.Int, error) {
	return _WithdrawalGateway.Contract.FINALIZATIONPERIODSECONDS(&_WithdrawalGateway.CallOpts)
}

// L2TOL1MESSAGEPASSER is a free data retrieval call binding the contract method 0xd5f3204b.
//
// Solidity: function L2_TO_L1_MESSAGE_PASSER() view returns(address)
func (_WithdrawalGateway *WithdrawalGatewayCaller) L2TOL1MESSAGEPASSER(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _WithdrawalGateway.contract.Call(opts, &out, "L2_TO_L1_MESSAGE_PASSER")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// L2TOL1MESSAGEPASSER is a free data retrieval call binding the contract method 0xd5f3204b.
//
// Solidity: function L2_TO_L1_MESSAGE_PASSER() view returns(address)
func (_WithdrawalGateway *WithdrawalGatewaySession) L2TOL1MESSAGEPASSER() (common.Address, error) {
	return _WithdrawalGateway.Contract.L2TOL1MESSAGEPASSER(&_WithdrawalGateway.CallOpts)
}

// L2TOL1MESSAGEPASSER is a free data retrieval call binding the contract method 0xd5f3204b.
//
// Solidity: function L2_TO_L1_MESSAGE_PASSER() view returns(address)
func (_WithdrawalGateway *WithdrawalGatewayCallerSession) L2TOL1MESSAGEPASSER() (common.Address, error) {
	return _WithdrawalGateway.Contract.L2TOL1MESSAGEPASSER(&_WithdrawalGateway.CallOpts)
}

// MPTVERIFIER is a free data retrieval call binding the contract method 0x053d3037.
//
// Solidity: function MPT_VERIFIER() view returns(address)
func (_WithdrawalGateway *WithdrawalGatewayCaller) MPTVERIFIER(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _WithdrawalGateway.contract.Call(opts, &out, "MPT_VERIFIER")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// MPTVERIFIER is a free data retrieval call binding the contract method 0x053d3037.
//
// Solidity: function MPT_VERIFIER() view returns(address)
func (_WithdrawalGateway *WithdrawalGatewaySession) MPTVERIFIER() (common.Address, error) {
	return _WithdrawalGateway.Contract.MPTVERIFIER(&_WithdrawalGateway.CallOpts)
}

// MPTVERIFIER is a free data retrieval call binding the contract method 0x053d3037.
//
// Solidity: function MPT_VERIFIER() view returns(address)
func (_WithdrawalGateway *WithdrawalGatewayCallerSession) MPTVERIFIER() (common.Address, error) {
	return _WithdrawalGateway.Contract.MPTVERIFIER(&_WithdrawalGateway.CallOpts)
}

// ROLLUPS is a free data retrieval call binding the contract method 0x2ef01a75.
//
// Solidity: function ROLLUPS() view returns(address)
func (_WithdrawalGateway *WithdrawalGatewayCaller) ROLLUPS(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _WithdrawalGateway.contract.Call(opts, &out, "ROLLUPS")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// ROLLUPS is a free data retrieval call binding the contract method 0x2ef01a75.
//
// Solidity: function ROLLUPS() view returns(address)
func (_WithdrawalGateway *WithdrawalGatewaySession) ROLLUPS() (common.Address, error) {
	return _WithdrawalGateway.Contract.ROLLUPS(&_WithdrawalGateway.CallOpts)
}

// ROLLUPS is a free data retrieval call binding the contract method 0x2ef01a75.
//
// Solidity: function ROLLUPS() view returns(address)
func (_WithdrawalGateway *WithdrawalGatewayCallerSession) ROLLUPS() (common.Address, error) {
	return _WithdrawalGateway.Contract.ROLLUPS(&_WithdrawalGateway.CallOpts)
}

// ROLLUPID is a free data retrieval call binding the contract method 0xb5ed5559.
//
// Solidity: function ROLLUP_ID() view returns(uint256)
func (_WithdrawalGateway *WithdrawalGatewayCaller) ROLLUPID(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _WithdrawalGateway.contract.Call(opts, &out, "ROLLUP_ID")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// ROLLUPID is a free data retrieval call binding the contract method 0xb5ed5559.
//
// Solidity: function ROLLUP_ID() view returns(uint256)
func (_WithdrawalGateway *WithdrawalGatewaySession) ROLLUPID() (*big.Int, error) {
	return _WithdrawalGateway.Contract.ROLLUPID(&_WithdrawalGateway.CallOpts)
}

// ROLLUPID is a free data retrieval call binding the contract method 0xb5ed5559.
//
// Solidity: function ROLLUP_ID() view returns(uint256)
func (_WithdrawalGateway *WithdrawalGatewayCallerSession) ROLLUPID() (*big.Int, error) {
	return _WithdrawalGateway.Contract.ROLLUPID(&_WithdrawalGateway.CallOpts)
}

// SENTMESSAGESSLOT is a free data retrieval call binding the contract method 0xc91ac182.
//
// Solidity: function SENT_MESSAGES_SLOT() view returns(uint256)
func (_WithdrawalGateway *WithdrawalGatewayCaller) SENTMESSAGESSLOT(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _WithdrawalGateway.contract.Call(opts, &out, "SENT_MESSAGES_SLOT")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// SENTMESSAGESSLOT is a free data retrieval call binding the contract method 0xc91ac182.
//
// Solidity: function SENT_MESSAGES_SLOT() view returns(uint256)
func (_WithdrawalGateway *WithdrawalGatewaySession) SENTMESSAGESSLOT() (*big.Int, error) {
	return _WithdrawalGateway.Contract.SENTMESSAGESSLOT(&_WithdrawalGateway.CallOpts)
}

// SENTMESSAGESSLOT is a free data retrieval call binding the contract method 0xc91ac182.
//
// Solidity: function SENT_MESSAGES_SLOT() view returns(uint256)
func (_WithdrawalGateway *WithdrawalGatewayCallerSession) SENTMESSAGESSLOT() (*big.Int, error) {
	return _WithdrawalGateway.Contract.SENTMESSAGESSLOT(&_WithdrawalGateway.CallOpts)
}

// Finalized is a free data retrieval call binding the contract method 0x0abea268.
//
// Solidity: function finalized(bytes32 ) view returns(bool)
func (_WithdrawalGateway *WithdrawalGatewayCaller) Finalized(opts *bind.CallOpts, arg0 [32]byte) (bool, error) {
	var out []interface{}
	err := _WithdrawalGateway.contract.Call(opts, &out, "finalized", arg0)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// Finalized is a free data retrieval call binding the contract method 0x0abea268.
//
// Solidity: function finalized(bytes32 ) view returns(bool)
func (_WithdrawalGateway *WithdrawalGatewaySession) Finalized(arg0 [32]byte) (bool, error) {
	return _WithdrawalGateway.Contract.Finalized(&_WithdrawalGateway.CallOpts, arg0)
}

// Finalized is a free data retrieval call binding the contract method 0x0abea268.
//
// Solidity: function finalized(bytes32 ) view returns(bool)
func (_WithdrawalGateway *WithdrawalGatewayCallerSession) Finalized(arg0 [32]byte) (bool, error) {
	return _WithdrawalGateway.Contract.Finalized(&_WithdrawalGateway.CallOpts, arg0)
}

// HashWithdrawal is a free data retrieval call binding the contract method 0x7d4395ac.
//
// Solidity: function hashWithdrawal((uint256,address,address,uint256,uint256,bytes) _tx) pure returns(bytes32)
func (_WithdrawalGateway *WithdrawalGatewayCaller) HashWithdrawal(opts *bind.CallOpts, _tx WithdrawalGatewayWithdrawalTransaction) ([32]byte, error) {
	var out []interface{}
	err := _WithdrawalGateway.contract.Call(opts, &out, "hashWithdrawal", _tx)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// HashWithdrawal is a free data retrieval call binding the contract method 0x7d4395ac.
//
// Solidity: function hashWithdrawal((uint256,address,address,uint256,uint256,bytes) _tx) pure returns(bytes32)
func (_WithdrawalGateway *WithdrawalGatewaySession) HashWithdrawal(_tx WithdrawalGatewayWithdrawalTransaction) ([32]byte, error) {
	return _WithdrawalGateway.Contract.HashWithdrawal(&_WithdrawalGateway.CallOpts, _tx)
}

// HashWithdrawal is a free data retrieval call binding the contract method 0x7d4395ac.
//
// Solidity: function hashWithdrawal((uint256,address,address,uint256,uint256,bytes) _tx) pure returns(bytes32)
func (_WithdrawalGateway *WithdrawalGatewayCallerSession) HashWithdrawal(_tx WithdrawalGatewayWithdrawalTransaction) ([32]byte, error) {
	return _WithdrawalGateway.Contract.HashWithdrawal(&_WithdrawalGateway.CallOpts, _tx)
}

// ProvenAt is a free data retrieval call binding the contract method 0xb84c2786.
//
// Solidity: function provenAt(bytes32 ) view returns(uint256)
func (_WithdrawalGateway *WithdrawalGatewayCaller) ProvenAt(opts *bind.CallOpts, arg0 [32]byte) (*big.Int, error) {
	var out []interface{}
	err := _WithdrawalGateway.contract.Call(opts, &out, "provenAt", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// ProvenAt is a free data retrieval call binding the contract method 0xb84c2786.
//
// Solidity: function provenAt(bytes32 ) view returns(uint256)
func (_WithdrawalGateway *WithdrawalGatewaySession) ProvenAt(arg0 [32]byte) (*big.Int, error) {
	return _WithdrawalGateway.Contract.ProvenAt(&_WithdrawalGateway.CallOpts, arg0)
}

// ProvenAt is a free data retrieval call binding the contract method 0xb84c2786.
//
// Solidity: function provenAt(bytes32 ) view returns(uint256)
func (_WithdrawalGateway *WithdrawalGatewayCallerSession) ProvenAt(arg0 [32]byte) (*big.Int, error) {
	return _WithdrawalGateway.Contract.ProvenAt(&_WithdrawalGateway.CallOpts, arg0)
}

// FinalizeWithdrawalTransaction is a paid mutator transaction binding the contract method 0x8c3152e9.
//
// Solidity: function finalizeWithdrawalTransaction((uint256,address,address,uint256,uint256,bytes) _tx) returns()
func (_WithdrawalGateway *WithdrawalGatewayTransactor) FinalizeWithdrawalTransaction(opts *bind.TransactOpts, _tx WithdrawalGatewayWithdrawalTransaction) (*types.Transaction, error) {
	return _WithdrawalGateway.contract.Transact(opts, "finalizeWithdrawalTransaction", _tx)
}

// FinalizeWithdrawalTransaction is a paid mutator transaction binding the contract method 0x8c3152e9.
//
// Solidity: function finalizeWithdrawalTransaction((uint256,address,address,uint256,uint256,bytes) _tx) returns()
func (_WithdrawalGateway *WithdrawalGatewaySession) FinalizeWithdrawalTransaction(_tx WithdrawalGatewayWithdrawalTransaction) (*types.Transaction, error) {
	return _WithdrawalGateway.Contract.FinalizeWithdrawalTransaction(&_WithdrawalGateway.TransactOpts, _tx)
}

// FinalizeWithdrawalTransaction is a paid mutator transaction binding the contract method 0x8c3152e9.
//
// Solidity: function finalizeWithdrawalTransaction((uint256,address,address,uint256,uint256,bytes) _tx) returns()
func (_WithdrawalGateway *WithdrawalGatewayTransactorSession) FinalizeWithdrawalTransaction(_tx WithdrawalGatewayWithdrawalTransaction) (*types.Transaction, error) {
	return _WithdrawalGateway.Contract.FinalizeWithdrawalTransaction(&_WithdrawalGateway.TransactOpts, _tx)
}

// ProveWithdrawalTransaction is a paid mutator transaction binding the contract method 0xc080abef.
//
// Solidity: function proveWithdrawalTransaction((uint256,address,address,uint256,uint256,bytes) _tx, (bytes32,bytes32,bytes32,bytes32) _outputRootProof, bytes[] _accountProof, bytes[] _withdrawalProof) returns()
func (_WithdrawalGateway *WithdrawalGatewayTransactor) ProveWithdrawalTransaction(opts *bind.TransactOpts, _tx WithdrawalGatewayWithdrawalTransaction, _outputRootProof WithdrawalGatewayOutputRootProof, _accountProof [][]byte, _withdrawalProof [][]byte) (*types.Transaction, error) {
	return _WithdrawalGateway.contract.Transact(opts, "proveWithdrawalTransaction", _tx, _outputRootProof, _accountProof, _withdrawalProof)
}

// ProveWithdrawalTransaction is a paid mutator transaction binding the contract method 0xc080abef.
//
// Solidity: function proveWithdrawalTransaction((uint256,address,address,uint256,uint256,bytes) _tx, (bytes32,bytes32,bytes32,bytes32) _outputRootProof, bytes[] _accountProof, bytes[] _withdrawalProof) returns()
func (_WithdrawalGateway *WithdrawalGatewaySession) ProveWithdrawalTransaction(_tx WithdrawalGatewayWithdrawalTransaction, _outputRootProof WithdrawalGatewayOutputRootProof, _accountProof [][]byte, _withdrawalProof [][]byte) (*types.Transaction, error) {
	return _WithdrawalGateway.Contract.ProveWithdrawalTransaction(&_WithdrawalGateway.TransactOpts, _tx, _outputRootProof, _accountProof, _withdrawalProof)
}

// ProveWithdrawalTransaction is a paid mutator transaction binding the contract method 0xc080abef.
//
// Solidity: function proveWithdrawalTransaction((uint256,address,address,uint256,uint256,bytes) _tx, (bytes32,bytes32,bytes32,bytes32) _outputRootProof, bytes[] _accountProof, bytes[] _withdrawalProof) returns()
func (_WithdrawalGateway *WithdrawalGatewayTransactorSession) ProveWithdrawalTransaction(_tx WithdrawalGatewayWithdrawalTransaction, _outputRootProof WithdrawalGatewayOutputRootProof, _accountProof [][]byte, _withdrawalProof [][]byte) (*types.Transaction, error) {
	return _WithdrawalGateway.Contract.ProveWithdrawalTransaction(&_WithdrawalGateway.TransactOpts, _tx, _outputRootProof, _accountProof, _withdrawalProof)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_WithdrawalGateway *WithdrawalGatewayTransactor) Receive(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _WithdrawalGateway.contract.RawTransact(opts, nil) // calldata is disallowed for receive function
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_WithdrawalGateway *WithdrawalGatewaySession) Receive() (*types.Transaction, error) {
	return _WithdrawalGateway.Contract.Receive(&_WithdrawalGateway.TransactOpts)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_WithdrawalGateway *WithdrawalGatewayTransactorSession) Receive() (*types.Transaction, error) {
	return _WithdrawalGateway.Contract.Receive(&_WithdrawalGateway.TransactOpts)
}

// WithdrawalGatewayWithdrawalFinalizedIterator is returned from FilterWithdrawalFinalized and is used to iterate over the raw logs and unpacked data for WithdrawalFinalized events raised by the WithdrawalGateway contract.
type WithdrawalGatewayWithdrawalFinalizedIterator struct {
	Event *WithdrawalGatewayWithdrawalFinalized // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *WithdrawalGatewayWithdrawalFinalizedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(WithdrawalGatewayWithdrawalFinalized)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(WithdrawalGatewayWithdrawalFinalized)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *WithdrawalGatewayWithdrawalFinalizedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *WithdrawalGatewayWithdrawalFinalizedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// WithdrawalGatewayWithdrawalFinalized represents a WithdrawalFinalized event raised by the WithdrawalGateway contract.
type WithdrawalGatewayWithdrawalFinalized struct {
	WithdrawalHash [32]byte
	Success        bool
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterWithdrawalFinalized is a free log retrieval operation binding the contract event 0xdb5c7652857aa163daadd670e116628fb42e869d8ac4251ef8971d9e5727df1b.
//
// Solidity: event WithdrawalFinalized(bytes32 indexed withdrawalHash, bool success)
func (_WithdrawalGateway *WithdrawalGatewayFilterer) FilterWithdrawalFinalized(opts *bind.FilterOpts, withdrawalHash [][32]byte) (*WithdrawalGatewayWithdrawalFinalizedIterator, error) {

	var withdrawalHashRule []interface{}
	for _, withdrawalHashItem := range withdrawalHash {
		withdrawalHashRule = append(withdrawalHashRule, withdrawalHashItem)
	}

	logs, sub, err := _WithdrawalGateway.contract.FilterLogs(opts, "WithdrawalFinalized", withdrawalHashRule)
	if err != nil {
		return nil, err
	}
	return &WithdrawalGatewayWithdrawalFinalizedIterator{contract: _WithdrawalGateway.contract, event: "WithdrawalFinalized", logs: logs, sub: sub}, nil
}

// WatchWithdrawalFinalized is a free log subscription operation binding the contract event 0xdb5c7652857aa163daadd670e116628fb42e869d8ac4251ef8971d9e5727df1b.
//
// Solidity: event WithdrawalFinalized(bytes32 indexed withdrawalHash, bool success)
func (_WithdrawalGateway *WithdrawalGatewayFilterer) WatchWithdrawalFinalized(opts *bind.WatchOpts, sink chan<- *WithdrawalGatewayWithdrawalFinalized, withdrawalHash [][32]byte) (event.Subscription, error) {

	var withdrawalHashRule []interface{}
	for _, withdrawalHashItem := range withdrawalHash {
		withdrawalHashRule = append(withdrawalHashRule, withdrawalHashItem)
	}

	logs, sub, err := _WithdrawalGateway.contract.WatchLogs(opts, "WithdrawalFinalized", withdrawalHashRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(WithdrawalGatewayWithdrawalFinalized)
				if err := _WithdrawalGateway.contract.UnpackLog(event, "WithdrawalFinalized", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseWithdrawalFinalized is a log parse operation binding the contract event 0xdb5c7652857aa163daadd670e116628fb42e869d8ac4251ef8971d9e5727df1b.
//
// Solidity: event WithdrawalFinalized(bytes32 indexed withdrawalHash, bool success)
func (_WithdrawalGateway *WithdrawalGatewayFilterer) ParseWithdrawalFinalized(log types.Log) (*WithdrawalGatewayWithdrawalFinalized, error) {
	event := new(WithdrawalGatewayWithdrawalFinalized)
	if err := _WithdrawalGateway.contract.UnpackLog(event, "WithdrawalFinalized", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// WithdrawalGatewayWithdrawalProvenIterator is returned from FilterWithdrawalProven and is used to iterate over the raw logs and unpacked data for WithdrawalProven events raised by the WithdrawalGateway contract.
type WithdrawalGatewayWithdrawalProvenIterator struct {
	Event *WithdrawalGatewayWithdrawalProven // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *WithdrawalGatewayWithdrawalProvenIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(WithdrawalGatewayWithdrawalProven)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(WithdrawalGatewayWithdrawalProven)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *WithdrawalGatewayWithdrawalProvenIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *WithdrawalGatewayWithdrawalProvenIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// WithdrawalGatewayWithdrawalProven represents a WithdrawalProven event raised by the WithdrawalGateway contract.
type WithdrawalGatewayWithdrawalProven struct {
	WithdrawalHash [32]byte
	From           common.Address
	To             common.Address
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterWithdrawalProven is a free log retrieval operation binding the contract event 0x67a6208cfcc0801d50f6cbe764733f4fddf66ac0b04442061a8a8c0cb6b63f62.
//
// Solidity: event WithdrawalProven(bytes32 indexed withdrawalHash, address indexed from, address indexed to)
func (_WithdrawalGateway *WithdrawalGatewayFilterer) FilterWithdrawalProven(opts *bind.FilterOpts, withdrawalHash [][32]byte, from []common.Address, to []common.Address) (*WithdrawalGatewayWithdrawalProvenIterator, error) {

	var withdrawalHashRule []interface{}
	for _, withdrawalHashItem := range withdrawalHash {
		withdrawalHashRule = append(withdrawalHashRule, withdrawalHashItem)
	}
	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _WithdrawalGateway.contract.FilterLogs(opts, "WithdrawalProven", withdrawalHashRule, fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &WithdrawalGatewayWithdrawalProvenIterator{contract: _WithdrawalGateway.contract, event: "WithdrawalProven", logs: logs, sub: sub}, nil
}

// WatchWithdrawalProven is a free log subscription operation binding the contract event 0x67a6208cfcc0801d50f6cbe764733f4fddf66ac0b04442061a8a8c0cb6b63f62.
//
// Solidity: event WithdrawalProven(bytes32 indexed withdrawalHash, address indexed from, address indexed to)
func (_WithdrawalGateway *WithdrawalGatewayFilterer) WatchWithdrawalProven(opts *bind.WatchOpts, sink chan<- *WithdrawalGatewayWithdrawalProven, withdrawalHash [][32]byte, from []common.Address, to []common.Address) (event.Subscription, error) {

	var withdrawalHashRule []interface{}
	for _, withdrawalHashItem := range withdrawalHash {
		withdrawalHashRule = append(withdrawalHashRule, withdrawalHashItem)
	}
	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _WithdrawalGateway.contract.WatchLogs(opts, "WithdrawalProven", withdrawalHashRule, fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(WithdrawalGatewayWithdrawalProven)
				if err := _WithdrawalGateway.contract.UnpackLog(event, "WithdrawalProven", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseWithdrawalProven is a log parse operation binding the contract event 0x67a6208cfcc0801d50f6cbe764733f4fddf66ac0b04442061a8a8c0cb6b63f62.
//
// Solidity: event WithdrawalProven(bytes32 indexed withdrawalHash, address indexed from, address indexed to)
func (_WithdrawalGateway *WithdrawalGatewayFilterer) ParseWithdrawalProven(log types.Log) (*WithdrawalGatewayWithdrawalProven, error) {
	event := new(WithdrawalGatewayWithdrawalProven)
	if err := _WithdrawalGateway.contract.UnpackLog(event, "WithdrawalProven", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
