// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package deposit

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

// DepositGatewayMetaData contains all meta data concerning the DepositGateway contract.
var DepositGatewayMetaData = &bind.MetaData{
	ABI: "[{\"type\":\"constructor\",\"inputs\":[{\"name\":\"rollupId\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"withdrawalGateway\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"receive\",\"stateMutability\":\"payable\"},{\"type\":\"function\",\"name\":\"ROLLUP_ID\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"WITHDRAWAL_GATEWAY\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"depositTransaction\",\"inputs\":[{\"name\":\"_to\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"_value\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"_gasLimit\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"_isCreation\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"_data\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[],\"stateMutability\":\"payable\"},{\"type\":\"function\",\"name\":\"payOutWithdrawal\",\"inputs\":[{\"name\":\"to\",\"type\":\"address\",\"internalType\":\"addresspayable\"},{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"event\",\"name\":\"TransactionDeposited\",\"inputs\":[{\"name\":\"from\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"to\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"version\",\"type\":\"uint256\",\"indexed\":true,\"internalType\":\"uint256\"},{\"name\":\"opaqueData\",\"type\":\"bytes\",\"indexed\":false,\"internalType\":\"bytes\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"WithdrawalSettled\",\"inputs\":[{\"name\":\"to\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"error\",\"name\":\"CreateWithNonZeroTarget\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"CreateWithNonZeroValue\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"GasLimitTooHigh\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"OnlyWithdrawalGateway\",\"inputs\":[]}]",
	Bin: "0x60c060405234801561000f575f80fd5b506040516106bc3803806106bc83398101604081905261002e91610045565b6080919091526001600160a01b031660a05261007f565b5f8060408385031215610056575f80fd5b825160208401519092506001600160a01b0381168114610074575f80fd5b809150509250929050565b60805160a0516106156100a75f395f818161015f01526101a401525f61010601526106155ff3fe608060405260043610610041575f3560e01c80638091ed53146100d4578063b5ed5559146100f5578063e9e05c421461013b578063f60714881461014e575f80fd5b366100d0575f3434620186a05f60405180602001604052805f815250604051602001610071959493929190610412565b60405160208183030381529060405290505f336001600160a01b0316336001600160a01b03167fb3813568d9991fc951961fcb4c784893574240a28925604d09fc577c55bb7c32846040516100c69190610461565b60405180910390a4005b5f80fd5b3480156100df575f80fd5b506100f36100ee3660046104aa565b610199565b005b348015610100575f80fd5b506101287f000000000000000000000000000000000000000000000000000000000000000081565b6040519081526020015b60405180910390f35b6100f36101493660046104e8565b6102c3565b348015610159575f80fd5b506101817f000000000000000000000000000000000000000000000000000000000000000081565b6040516001600160a01b039091168152602001610132565b336001600160a01b037f000000000000000000000000000000000000000000000000000000000000000016146101e257604051630399193560e01b815260040160405180910390fd5b816001600160a01b03167f979ca02696ec818ded66a00c02750e26c18bb5595d8f57e972f91b3ac6aa54df8260405161021d91815260200190565b60405180910390a25f826001600160a01b0316826040515f6040518083038185875af1925050503d805f811461026e576040519150601f19603f3d011682016040523d82523d5f602084013e610273565b606091505b50509050806102be5760405162461bcd60e51b8152602060048201526013602482015272115512081d1c985b9cd9995c8819985a5b1959606a1b604482015260640160405180910390fd5b505050565b8180156102d857506001600160a01b03851615155b156102f6576040516331594c8d60e11b815260040160405180910390fd5b81801561030257508315155b156103205760405163890c3fd360e01b815260040160405180910390fd5b6301c9c38067ffffffffffffffff8416111561034f57604051637d2e4eb560e01b815260040160405180910390fd5b33328114610370575033731111000000000000000000000000000000001111015b5f348686868660405160200161038a959493929190610412565b60405160208183030381529060405290505f876001600160a01b0316836001600160a01b03167fb3813568d9991fc951961fcb4c784893574240a28925604d09fc577c55bb7c32846040516103df9190610461565b60405180910390a450505050505050565b5f5b8381101561040a5781810151838201526020016103f2565b50505f910152565b85815284602082015267ffffffffffffffff60c01b8460c01b16604082015282151560f81b60488201525f82516104508160498501602087016103f0565b919091016049019695505050505050565b602081525f825180602084015261047f8160408501602087016103f0565b601f01601f19169190910160400192915050565b6001600160a01b03811681146104a7575f80fd5b50565b5f80604083850312156104bb575f80fd5b82356104c681610493565b946020939093013593505050565b634e487b7160e01b5f52604160045260245ffd5b5f805f805f60a086880312156104fc575f80fd5b853561050781610493565b945060208601359350604086013567ffffffffffffffff808216821461052b575f80fd5b9093506060870135908115158214610541575f80fd5b90925060808701359080821115610556575f80fd5b818801915088601f830112610569575f80fd5b81358181111561057b5761057b6104d4565b604051601f8201601f19908116603f011681019083821181831017156105a3576105a36104d4565b816040528281528b60208487010111156105bb575f80fd5b826020860160208301375f602084830101528095505050505050929550929590935056fea2646970667358221220ab8b4d7533fb8d8060c128f3357a6c332fc177d3be6baf6e45400bb50d8d9d3964736f6c63430008180033",
}

// DepositGatewayABI is the input ABI used to generate the binding from.
// Deprecated: Use DepositGatewayMetaData.ABI instead.
var DepositGatewayABI = DepositGatewayMetaData.ABI

// DepositGatewayBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use DepositGatewayMetaData.Bin instead.
var DepositGatewayBin = DepositGatewayMetaData.Bin

// DeployDepositGateway deploys a new Ethereum contract, binding an instance of DepositGateway to it.
func DeployDepositGateway(auth *bind.TransactOpts, backend bind.ContractBackend, rollupId *big.Int, withdrawalGateway common.Address) (common.Address, *types.Transaction, *DepositGateway, error) {
	parsed, err := DepositGatewayMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(DepositGatewayBin), backend, rollupId, withdrawalGateway)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &DepositGateway{DepositGatewayCaller: DepositGatewayCaller{contract: contract}, DepositGatewayTransactor: DepositGatewayTransactor{contract: contract}, DepositGatewayFilterer: DepositGatewayFilterer{contract: contract}}, nil
}

// DepositGateway is an auto generated Go binding around an Ethereum contract.
type DepositGateway struct {
	DepositGatewayCaller     // Read-only binding to the contract
	DepositGatewayTransactor // Write-only binding to the contract
	DepositGatewayFilterer   // Log filterer for contract events
}

// DepositGatewayCaller is an auto generated read-only Go binding around an Ethereum contract.
type DepositGatewayCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// DepositGatewayTransactor is an auto generated write-only Go binding around an Ethereum contract.
type DepositGatewayTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// DepositGatewayFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type DepositGatewayFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// DepositGatewaySession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type DepositGatewaySession struct {
	Contract     *DepositGateway   // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// DepositGatewayCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type DepositGatewayCallerSession struct {
	Contract *DepositGatewayCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts         // Call options to use throughout this session
}

// DepositGatewayTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type DepositGatewayTransactorSession struct {
	Contract     *DepositGatewayTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts         // Transaction auth options to use throughout this session
}

// DepositGatewayRaw is an auto generated low-level Go binding around an Ethereum contract.
type DepositGatewayRaw struct {
	Contract *DepositGateway // Generic contract binding to access the raw methods on
}

// DepositGatewayCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type DepositGatewayCallerRaw struct {
	Contract *DepositGatewayCaller // Generic read-only contract binding to access the raw methods on
}

// DepositGatewayTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type DepositGatewayTransactorRaw struct {
	Contract *DepositGatewayTransactor // Generic write-only contract binding to access the raw methods on
}

// NewDepositGateway creates a new instance of DepositGateway, bound to a specific deployed contract.
func NewDepositGateway(address common.Address, backend bind.ContractBackend) (*DepositGateway, error) {
	contract, err := bindDepositGateway(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &DepositGateway{DepositGatewayCaller: DepositGatewayCaller{contract: contract}, DepositGatewayTransactor: DepositGatewayTransactor{contract: contract}, DepositGatewayFilterer: DepositGatewayFilterer{contract: contract}}, nil
}

// NewDepositGatewayCaller creates a new read-only instance of DepositGateway, bound to a specific deployed contract.
func NewDepositGatewayCaller(address common.Address, caller bind.ContractCaller) (*DepositGatewayCaller, error) {
	contract, err := bindDepositGateway(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &DepositGatewayCaller{contract: contract}, nil
}

// NewDepositGatewayTransactor creates a new write-only instance of DepositGateway, bound to a specific deployed contract.
func NewDepositGatewayTransactor(address common.Address, transactor bind.ContractTransactor) (*DepositGatewayTransactor, error) {
	contract, err := bindDepositGateway(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &DepositGatewayTransactor{contract: contract}, nil
}

// NewDepositGatewayFilterer creates a new log filterer instance of DepositGateway, bound to a specific deployed contract.
func NewDepositGatewayFilterer(address common.Address, filterer bind.ContractFilterer) (*DepositGatewayFilterer, error) {
	contract, err := bindDepositGateway(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &DepositGatewayFilterer{contract: contract}, nil
}

// bindDepositGateway binds a generic wrapper to an already deployed contract.
func bindDepositGateway(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := DepositGatewayMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_DepositGateway *DepositGatewayRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _DepositGateway.Contract.DepositGatewayCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_DepositGateway *DepositGatewayRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _DepositGateway.Contract.DepositGatewayTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_DepositGateway *DepositGatewayRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _DepositGateway.Contract.DepositGatewayTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_DepositGateway *DepositGatewayCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _DepositGateway.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_DepositGateway *DepositGatewayTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _DepositGateway.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_DepositGateway *DepositGatewayTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _DepositGateway.Contract.contract.Transact(opts, method, params...)
}

// ROLLUPID is a free data retrieval call binding the contract method 0xb5ed5559.
//
// Solidity: function ROLLUP_ID() view returns(uint256)
func (_DepositGateway *DepositGatewayCaller) ROLLUPID(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _DepositGateway.contract.Call(opts, &out, "ROLLUP_ID")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// ROLLUPID is a free data retrieval call binding the contract method 0xb5ed5559.
//
// Solidity: function ROLLUP_ID() view returns(uint256)
func (_DepositGateway *DepositGatewaySession) ROLLUPID() (*big.Int, error) {
	return _DepositGateway.Contract.ROLLUPID(&_DepositGateway.CallOpts)
}

// ROLLUPID is a free data retrieval call binding the contract method 0xb5ed5559.
//
// Solidity: function ROLLUP_ID() view returns(uint256)
func (_DepositGateway *DepositGatewayCallerSession) ROLLUPID() (*big.Int, error) {
	return _DepositGateway.Contract.ROLLUPID(&_DepositGateway.CallOpts)
}

// WITHDRAWALGATEWAY is a free data retrieval call binding the contract method 0xf6071488.
//
// Solidity: function WITHDRAWAL_GATEWAY() view returns(address)
func (_DepositGateway *DepositGatewayCaller) WITHDRAWALGATEWAY(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _DepositGateway.contract.Call(opts, &out, "WITHDRAWAL_GATEWAY")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// WITHDRAWALGATEWAY is a free data retrieval call binding the contract method 0xf6071488.
//
// Solidity: function WITHDRAWAL_GATEWAY() view returns(address)
func (_DepositGateway *DepositGatewaySession) WITHDRAWALGATEWAY() (common.Address, error) {
	return _DepositGateway.Contract.WITHDRAWALGATEWAY(&_DepositGateway.CallOpts)
}

// WITHDRAWALGATEWAY is a free data retrieval call binding the contract method 0xf6071488.
//
// Solidity: function WITHDRAWAL_GATEWAY() view returns(address)
func (_DepositGateway *DepositGatewayCallerSession) WITHDRAWALGATEWAY() (common.Address, error) {
	return _DepositGateway.Contract.WITHDRAWALGATEWAY(&_DepositGateway.CallOpts)
}

// DepositTransaction is a paid mutator transaction binding the contract method 0xe9e05c42.
//
// Solidity: function depositTransaction(address _to, uint256 _value, uint64 _gasLimit, bool _isCreation, bytes _data) payable returns()
func (_DepositGateway *DepositGatewayTransactor) DepositTransaction(opts *bind.TransactOpts, _to common.Address, _value *big.Int, _gasLimit uint64, _isCreation bool, _data []byte) (*types.Transaction, error) {
	return _DepositGateway.contract.Transact(opts, "depositTransaction", _to, _value, _gasLimit, _isCreation, _data)
}

// DepositTransaction is a paid mutator transaction binding the contract method 0xe9e05c42.
//
// Solidity: function depositTransaction(address _to, uint256 _value, uint64 _gasLimit, bool _isCreation, bytes _data) payable returns()
func (_DepositGateway *DepositGatewaySession) DepositTransaction(_to common.Address, _value *big.Int, _gasLimit uint64, _isCreation bool, _data []byte) (*types.Transaction, error) {
	return _DepositGateway.Contract.DepositTransaction(&_DepositGateway.TransactOpts, _to, _value, _gasLimit, _isCreation, _data)
}

// DepositTransaction is a paid mutator transaction binding the contract method 0xe9e05c42.
//
// Solidity: function depositTransaction(address _to, uint256 _value, uint64 _gasLimit, bool _isCreation, bytes _data) payable returns()
func (_DepositGateway *DepositGatewayTransactorSession) DepositTransaction(_to common.Address, _value *big.Int, _gasLimit uint64, _isCreation bool, _data []byte) (*types.Transaction, error) {
	return _DepositGateway.Contract.DepositTransaction(&_DepositGateway.TransactOpts, _to, _value, _gasLimit, _isCreation, _data)
}

// PayOutWithdrawal is a paid mutator transaction binding the contract method 0x8091ed53.
//
// Solidity: function payOutWithdrawal(address to, uint256 amount) returns()
func (_DepositGateway *DepositGatewayTransactor) PayOutWithdrawal(opts *bind.TransactOpts, to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _DepositGateway.contract.Transact(opts, "payOutWithdrawal", to, amount)
}

// PayOutWithdrawal is a paid mutator transaction binding the contract method 0x8091ed53.
//
// Solidity: function payOutWithdrawal(address to, uint256 amount) returns()
func (_DepositGateway *DepositGatewaySession) PayOutWithdrawal(to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _DepositGateway.Contract.PayOutWithdrawal(&_DepositGateway.TransactOpts, to, amount)
}

// PayOutWithdrawal is a paid mutator transaction binding the contract method 0x8091ed53.
//
// Solidity: function payOutWithdrawal(address to, uint256 amount) returns()
func (_DepositGateway *DepositGatewayTransactorSession) PayOutWithdrawal(to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _DepositGateway.Contract.PayOutWithdrawal(&_DepositGateway.TransactOpts, to, amount)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_DepositGateway *DepositGatewayTransactor) Receive(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _DepositGateway.contract.RawTransact(opts, nil) // calldata is disallowed for receive function
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_DepositGateway *DepositGatewaySession) Receive() (*types.Transaction, error) {
	return _DepositGateway.Contract.Receive(&_DepositGateway.TransactOpts)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_DepositGateway *DepositGatewayTransactorSession) Receive() (*types.Transaction, error) {
	return _DepositGateway.Contract.Receive(&_DepositGateway.TransactOpts)
}

// DepositGatewayTransactionDepositedIterator is returned from FilterTransactionDeposited and is used to iterate over the raw logs and unpacked data for TransactionDeposited events raised by the DepositGateway contract.
type DepositGatewayTransactionDepositedIterator struct {
	Event *DepositGatewayTransactionDeposited // Event containing the contract specifics and raw log

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
func (it *DepositGatewayTransactionDepositedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DepositGatewayTransactionDeposited)
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
		it.Event = new(DepositGatewayTransactionDeposited)
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
func (it *DepositGatewayTransactionDepositedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DepositGatewayTransactionDepositedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DepositGatewayTransactionDeposited represents a TransactionDeposited event raised by the DepositGateway contract.
type DepositGatewayTransactionDeposited struct {
	From       common.Address
	To         common.Address
	Version    *big.Int
	OpaqueData []byte
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterTransactionDeposited is a free log retrieval operation binding the contract event 0xb3813568d9991fc951961fcb4c784893574240a28925604d09fc577c55bb7c32.
//
// Solidity: event TransactionDeposited(address indexed from, address indexed to, uint256 indexed version, bytes opaqueData)
func (_DepositGateway *DepositGatewayFilterer) FilterTransactionDeposited(opts *bind.FilterOpts, from []common.Address, to []common.Address, version []*big.Int) (*DepositGatewayTransactionDepositedIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}
	var versionRule []interface{}
	for _, versionItem := range version {
		versionRule = append(versionRule, versionItem)
	}

	logs, sub, err := _DepositGateway.contract.FilterLogs(opts, "TransactionDeposited", fromRule, toRule, versionRule)
	if err != nil {
		return nil, err
	}
	return &DepositGatewayTransactionDepositedIterator{contract: _DepositGateway.contract, event: "TransactionDeposited", logs: logs, sub: sub}, nil
}

// WatchTransactionDeposited is a free log subscription operation binding the contract event 0xb3813568d9991fc951961fcb4c784893574240a28925604d09fc577c55bb7c32.
//
// Solidity: event TransactionDeposited(address indexed from, address indexed to, uint256 indexed version, bytes opaqueData)
func (_DepositGateway *DepositGatewayFilterer) WatchTransactionDeposited(opts *bind.WatchOpts, sink chan<- *DepositGatewayTransactionDeposited, from []common.Address, to []common.Address, version []*big.Int) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}
	var versionRule []interface{}
	for _, versionItem := range version {
		versionRule = append(versionRule, versionItem)
	}

	logs, sub, err := _DepositGateway.contract.WatchLogs(opts, "TransactionDeposited", fromRule, toRule, versionRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DepositGatewayTransactionDeposited)
				if err := _DepositGateway.contract.UnpackLog(event, "TransactionDeposited", log); err != nil {
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

// ParseTransactionDeposited is a log parse operation binding the contract event 0xb3813568d9991fc951961fcb4c784893574240a28925604d09fc577c55bb7c32.
//
// Solidity: event TransactionDeposited(address indexed from, address indexed to, uint256 indexed version, bytes opaqueData)
func (_DepositGateway *DepositGatewayFilterer) ParseTransactionDeposited(log types.Log) (*DepositGatewayTransactionDeposited, error) {
	event := new(DepositGatewayTransactionDeposited)
	if err := _DepositGateway.contract.UnpackLog(event, "TransactionDeposited", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// DepositGatewayWithdrawalSettledIterator is returned from FilterWithdrawalSettled and is used to iterate over the raw logs and unpacked data for WithdrawalSettled events raised by the DepositGateway contract.
type DepositGatewayWithdrawalSettledIterator struct {
	Event *DepositGatewayWithdrawalSettled // Event containing the contract specifics and raw log

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
func (it *DepositGatewayWithdrawalSettledIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DepositGatewayWithdrawalSettled)
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
		it.Event = new(DepositGatewayWithdrawalSettled)
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
func (it *DepositGatewayWithdrawalSettledIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DepositGatewayWithdrawalSettledIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DepositGatewayWithdrawalSettled represents a WithdrawalSettled event raised by the DepositGateway contract.
type DepositGatewayWithdrawalSettled struct {
	To     common.Address
	Amount *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterWithdrawalSettled is a free log retrieval operation binding the contract event 0x979ca02696ec818ded66a00c02750e26c18bb5595d8f57e972f91b3ac6aa54df.
//
// Solidity: event WithdrawalSettled(address indexed to, uint256 amount)
func (_DepositGateway *DepositGatewayFilterer) FilterWithdrawalSettled(opts *bind.FilterOpts, to []common.Address) (*DepositGatewayWithdrawalSettledIterator, error) {

	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _DepositGateway.contract.FilterLogs(opts, "WithdrawalSettled", toRule)
	if err != nil {
		return nil, err
	}
	return &DepositGatewayWithdrawalSettledIterator{contract: _DepositGateway.contract, event: "WithdrawalSettled", logs: logs, sub: sub}, nil
}

// WatchWithdrawalSettled is a free log subscription operation binding the contract event 0x979ca02696ec818ded66a00c02750e26c18bb5595d8f57e972f91b3ac6aa54df.
//
// Solidity: event WithdrawalSettled(address indexed to, uint256 amount)
func (_DepositGateway *DepositGatewayFilterer) WatchWithdrawalSettled(opts *bind.WatchOpts, sink chan<- *DepositGatewayWithdrawalSettled, to []common.Address) (event.Subscription, error) {

	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _DepositGateway.contract.WatchLogs(opts, "WithdrawalSettled", toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DepositGatewayWithdrawalSettled)
				if err := _DepositGateway.contract.UnpackLog(event, "WithdrawalSettled", log); err != nil {
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

// ParseWithdrawalSettled is a log parse operation binding the contract event 0x979ca02696ec818ded66a00c02750e26c18bb5595d8f57e972f91b3ac6aa54df.
//
// Solidity: event WithdrawalSettled(address indexed to, uint256 amount)
func (_DepositGateway *DepositGatewayFilterer) ParseWithdrawalSettled(log types.Log) (*DepositGatewayWithdrawalSettled, error) {
	event := new(DepositGatewayWithdrawalSettled)
	if err := _DepositGateway.contract.UnpackLog(event, "WithdrawalSettled", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
