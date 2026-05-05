// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package proofsys

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

// ProofsysMetaData contains all meta data concerning the Proofsys contract.
var ProofsysMetaData = &bind.MetaData{
	ABI: "[{\"type\":\"constructor\",\"inputs\":[{\"name\":\"initialOwner\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"initialSigner\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"owner\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"renounceOwnership\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setSigner\",\"inputs\":[{\"name\":\"newSigner\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"signer\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"transferOwnership\",\"inputs\":[{\"name\":\"newOwner\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"verify\",\"inputs\":[{\"name\":\"proof\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"publicInputsHash\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"event\",\"name\":\"OwnershipTransferred\",\"inputs\":[{\"name\":\"previousOwner\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"newOwner\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"SignerUpdated\",\"inputs\":[{\"name\":\"previousSigner\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"newSigner\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"error\",\"name\":\"ECDSAInvalidSignature\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"ECDSAInvalidSignatureLength\",\"inputs\":[{\"name\":\"length\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"type\":\"error\",\"name\":\"ECDSAInvalidSignatureS\",\"inputs\":[{\"name\":\"s\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}]},{\"type\":\"error\",\"name\":\"OwnableInvalidOwner\",\"inputs\":[{\"name\":\"owner\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"OwnableUnauthorizedAccount\",\"inputs\":[{\"name\":\"account\",\"type\":\"address\",\"internalType\":\"address\"}]}]",
	Bin: "0x6080346100ea57601f61057e38819003918201601f19168301916001600160401b038311848410176100ee5780849260409485528339810103126100ea57610052602061004b83610102565b9201610102565b6001600160a01b039091169081156100d7575f80546001600160a01b031981168417825560405193916001600160a01b03909116907f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e09080a3600180546001600160a01b0319166001600160a01b039290921691909117905561046790816101178239f35b631e4fbdf760e01b5f525f60045260245ffd5b5f80fd5b634e487b7160e01b5f52604160045260245ffd5b51906001600160a01b03821682036100ea5756fe6080806040526004361015610012575f80fd5b5f3560e01c908163238ac933146102ad575080636b406341146101d95780636c19e78314610169578063715018a6146101125780638da5cb5b146100eb5763f2fde38b1461005e575f80fd5b346100e75760203660031901126100e7576004356001600160a01b038116908190036100e75761008c6102d0565b80156100d4575f80546001600160a01b03198116831782556001600160a01b0316907f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e09080a3005b631e4fbdf760e01b5f525f60045260245ffd5b5f80fd5b346100e7575f3660031901126100e7575f546040516001600160a01b039091168152602090f35b346100e7575f3660031901126100e75761012a6102d0565b5f80546001600160a01b0319811682556001600160a01b03167f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e08280a3005b346100e75760203660031901126100e7576004356001600160a01b038116908190036100e7576101976102d0565b600154816001600160a01b0382167f2d025324f0a785e8c12d0a0d91a9caa49df4ef20ff87e0df7213a1d4f3157beb5f80a36001600160a01b03191617600155005b346100e75760403660031901126100e75760043567ffffffffffffffff81116100e757366023820112156100e757806004013567ffffffffffffffff81116100e75736602482840101116100e757604051601f8201601f19908116603f01168101919067ffffffffffffffff831181841017610299575f602083610276946024839861027f986040528287520183860137830101526024356102f6565b90929192610330565b6001546040516001600160a01b0392831691909216148152f35b634e487b7160e01b5f52604160045260245ffd5b346100e7575f3660031901126100e7576001546001600160a01b03168152602090f35b5f546001600160a01b031633036102e357565b63118cdaa760e01b5f523360045260245ffd5b81519190604183036103265761031f9250602082015190606060408401519301515f1a906103a4565b9192909190565b50505f9160029190565b60048110156103905780610342575050565b600181036103595763f645eedf60e01b5f5260045ffd5b60028103610374575063fce698f760e01b5f5260045260245ffd5b60031461037e5750565b6335e2f38360e21b5f5260045260245ffd5b634e487b7160e01b5f52602160045260245ffd5b91907f7fffffffffffffffffffffffffffffff5d576e7357a4501ddfe92f46681b20a08411610426579160209360809260ff5f9560405194855216868401526040830152606082015282805260015afa1561041b575f516001600160a01b0381161561041157905f905f90565b505f906001905f90565b6040513d5f823e3d90fd5b5050505f916003919056fea2646970667358221220100771474e0bde65b36c7a249a0fcf3cb8402fd24abca127bf0308afc5114e4164736f6c63430008220033",
}

// ProofsysABI is the input ABI used to generate the binding from.
// Deprecated: Use ProofsysMetaData.ABI instead.
var ProofsysABI = ProofsysMetaData.ABI

// ProofsysBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use ProofsysMetaData.Bin instead.
var ProofsysBin = ProofsysMetaData.Bin

// DeployProofsys deploys a new Ethereum contract, binding an instance of Proofsys to it.
func DeployProofsys(auth *bind.TransactOpts, backend bind.ContractBackend, initialOwner common.Address, initialSigner common.Address) (common.Address, *types.Transaction, *Proofsys, error) {
	parsed, err := ProofsysMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(ProofsysBin), backend, initialOwner, initialSigner)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Proofsys{ProofsysCaller: ProofsysCaller{contract: contract}, ProofsysTransactor: ProofsysTransactor{contract: contract}, ProofsysFilterer: ProofsysFilterer{contract: contract}}, nil
}

// Proofsys is an auto generated Go binding around an Ethereum contract.
type Proofsys struct {
	ProofsysCaller     // Read-only binding to the contract
	ProofsysTransactor // Write-only binding to the contract
	ProofsysFilterer   // Log filterer for contract events
}

// ProofsysCaller is an auto generated read-only Go binding around an Ethereum contract.
type ProofsysCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ProofsysTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ProofsysTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ProofsysFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ProofsysFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ProofsysSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ProofsysSession struct {
	Contract     *Proofsys         // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ProofsysCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ProofsysCallerSession struct {
	Contract *ProofsysCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts   // Call options to use throughout this session
}

// ProofsysTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ProofsysTransactorSession struct {
	Contract     *ProofsysTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts   // Transaction auth options to use throughout this session
}

// ProofsysRaw is an auto generated low-level Go binding around an Ethereum contract.
type ProofsysRaw struct {
	Contract *Proofsys // Generic contract binding to access the raw methods on
}

// ProofsysCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ProofsysCallerRaw struct {
	Contract *ProofsysCaller // Generic read-only contract binding to access the raw methods on
}

// ProofsysTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ProofsysTransactorRaw struct {
	Contract *ProofsysTransactor // Generic write-only contract binding to access the raw methods on
}

// NewProofsys creates a new instance of Proofsys, bound to a specific deployed contract.
func NewProofsys(address common.Address, backend bind.ContractBackend) (*Proofsys, error) {
	contract, err := bindProofsys(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Proofsys{ProofsysCaller: ProofsysCaller{contract: contract}, ProofsysTransactor: ProofsysTransactor{contract: contract}, ProofsysFilterer: ProofsysFilterer{contract: contract}}, nil
}

// NewProofsysCaller creates a new read-only instance of Proofsys, bound to a specific deployed contract.
func NewProofsysCaller(address common.Address, caller bind.ContractCaller) (*ProofsysCaller, error) {
	contract, err := bindProofsys(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ProofsysCaller{contract: contract}, nil
}

// NewProofsysTransactor creates a new write-only instance of Proofsys, bound to a specific deployed contract.
func NewProofsysTransactor(address common.Address, transactor bind.ContractTransactor) (*ProofsysTransactor, error) {
	contract, err := bindProofsys(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ProofsysTransactor{contract: contract}, nil
}

// NewProofsysFilterer creates a new log filterer instance of Proofsys, bound to a specific deployed contract.
func NewProofsysFilterer(address common.Address, filterer bind.ContractFilterer) (*ProofsysFilterer, error) {
	contract, err := bindProofsys(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ProofsysFilterer{contract: contract}, nil
}

// bindProofsys binds a generic wrapper to an already deployed contract.
func bindProofsys(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := ProofsysMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Proofsys *ProofsysRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Proofsys.Contract.ProofsysCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Proofsys *ProofsysRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Proofsys.Contract.ProofsysTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Proofsys *ProofsysRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Proofsys.Contract.ProofsysTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Proofsys *ProofsysCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Proofsys.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Proofsys *ProofsysTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Proofsys.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Proofsys *ProofsysTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Proofsys.Contract.contract.Transact(opts, method, params...)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Proofsys *ProofsysCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Proofsys.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Proofsys *ProofsysSession) Owner() (common.Address, error) {
	return _Proofsys.Contract.Owner(&_Proofsys.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Proofsys *ProofsysCallerSession) Owner() (common.Address, error) {
	return _Proofsys.Contract.Owner(&_Proofsys.CallOpts)
}

// Signer is a free data retrieval call binding the contract method 0x238ac933.
//
// Solidity: function signer() view returns(address)
func (_Proofsys *ProofsysCaller) Signer(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Proofsys.contract.Call(opts, &out, "signer")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Signer is a free data retrieval call binding the contract method 0x238ac933.
//
// Solidity: function signer() view returns(address)
func (_Proofsys *ProofsysSession) Signer() (common.Address, error) {
	return _Proofsys.Contract.Signer(&_Proofsys.CallOpts)
}

// Signer is a free data retrieval call binding the contract method 0x238ac933.
//
// Solidity: function signer() view returns(address)
func (_Proofsys *ProofsysCallerSession) Signer() (common.Address, error) {
	return _Proofsys.Contract.Signer(&_Proofsys.CallOpts)
}

// Verify is a free data retrieval call binding the contract method 0x6b406341.
//
// Solidity: function verify(bytes proof, bytes32 publicInputsHash) view returns(bool)
func (_Proofsys *ProofsysCaller) Verify(opts *bind.CallOpts, proof []byte, publicInputsHash [32]byte) (bool, error) {
	var out []interface{}
	err := _Proofsys.contract.Call(opts, &out, "verify", proof, publicInputsHash)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// Verify is a free data retrieval call binding the contract method 0x6b406341.
//
// Solidity: function verify(bytes proof, bytes32 publicInputsHash) view returns(bool)
func (_Proofsys *ProofsysSession) Verify(proof []byte, publicInputsHash [32]byte) (bool, error) {
	return _Proofsys.Contract.Verify(&_Proofsys.CallOpts, proof, publicInputsHash)
}

// Verify is a free data retrieval call binding the contract method 0x6b406341.
//
// Solidity: function verify(bytes proof, bytes32 publicInputsHash) view returns(bool)
func (_Proofsys *ProofsysCallerSession) Verify(proof []byte, publicInputsHash [32]byte) (bool, error) {
	return _Proofsys.Contract.Verify(&_Proofsys.CallOpts, proof, publicInputsHash)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Proofsys *ProofsysTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Proofsys.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Proofsys *ProofsysSession) RenounceOwnership() (*types.Transaction, error) {
	return _Proofsys.Contract.RenounceOwnership(&_Proofsys.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Proofsys *ProofsysTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _Proofsys.Contract.RenounceOwnership(&_Proofsys.TransactOpts)
}

// SetSigner is a paid mutator transaction binding the contract method 0x6c19e783.
//
// Solidity: function setSigner(address newSigner) returns()
func (_Proofsys *ProofsysTransactor) SetSigner(opts *bind.TransactOpts, newSigner common.Address) (*types.Transaction, error) {
	return _Proofsys.contract.Transact(opts, "setSigner", newSigner)
}

// SetSigner is a paid mutator transaction binding the contract method 0x6c19e783.
//
// Solidity: function setSigner(address newSigner) returns()
func (_Proofsys *ProofsysSession) SetSigner(newSigner common.Address) (*types.Transaction, error) {
	return _Proofsys.Contract.SetSigner(&_Proofsys.TransactOpts, newSigner)
}

// SetSigner is a paid mutator transaction binding the contract method 0x6c19e783.
//
// Solidity: function setSigner(address newSigner) returns()
func (_Proofsys *ProofsysTransactorSession) SetSigner(newSigner common.Address) (*types.Transaction, error) {
	return _Proofsys.Contract.SetSigner(&_Proofsys.TransactOpts, newSigner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Proofsys *ProofsysTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _Proofsys.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Proofsys *ProofsysSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Proofsys.Contract.TransferOwnership(&_Proofsys.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Proofsys *ProofsysTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Proofsys.Contract.TransferOwnership(&_Proofsys.TransactOpts, newOwner)
}

// ProofsysOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the Proofsys contract.
type ProofsysOwnershipTransferredIterator struct {
	Event *ProofsysOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *ProofsysOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ProofsysOwnershipTransferred)
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
		it.Event = new(ProofsysOwnershipTransferred)
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
func (it *ProofsysOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ProofsysOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ProofsysOwnershipTransferred represents a OwnershipTransferred event raised by the Proofsys contract.
type ProofsysOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Proofsys *ProofsysFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*ProofsysOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Proofsys.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &ProofsysOwnershipTransferredIterator{contract: _Proofsys.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Proofsys *ProofsysFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *ProofsysOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Proofsys.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ProofsysOwnershipTransferred)
				if err := _Proofsys.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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

// ParseOwnershipTransferred is a log parse operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Proofsys *ProofsysFilterer) ParseOwnershipTransferred(log types.Log) (*ProofsysOwnershipTransferred, error) {
	event := new(ProofsysOwnershipTransferred)
	if err := _Proofsys.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ProofsysSignerUpdatedIterator is returned from FilterSignerUpdated and is used to iterate over the raw logs and unpacked data for SignerUpdated events raised by the Proofsys contract.
type ProofsysSignerUpdatedIterator struct {
	Event *ProofsysSignerUpdated // Event containing the contract specifics and raw log

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
func (it *ProofsysSignerUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ProofsysSignerUpdated)
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
		it.Event = new(ProofsysSignerUpdated)
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
func (it *ProofsysSignerUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ProofsysSignerUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ProofsysSignerUpdated represents a SignerUpdated event raised by the Proofsys contract.
type ProofsysSignerUpdated struct {
	PreviousSigner common.Address
	NewSigner      common.Address
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterSignerUpdated is a free log retrieval operation binding the contract event 0x2d025324f0a785e8c12d0a0d91a9caa49df4ef20ff87e0df7213a1d4f3157beb.
//
// Solidity: event SignerUpdated(address indexed previousSigner, address indexed newSigner)
func (_Proofsys *ProofsysFilterer) FilterSignerUpdated(opts *bind.FilterOpts, previousSigner []common.Address, newSigner []common.Address) (*ProofsysSignerUpdatedIterator, error) {

	var previousSignerRule []interface{}
	for _, previousSignerItem := range previousSigner {
		previousSignerRule = append(previousSignerRule, previousSignerItem)
	}
	var newSignerRule []interface{}
	for _, newSignerItem := range newSigner {
		newSignerRule = append(newSignerRule, newSignerItem)
	}

	logs, sub, err := _Proofsys.contract.FilterLogs(opts, "SignerUpdated", previousSignerRule, newSignerRule)
	if err != nil {
		return nil, err
	}
	return &ProofsysSignerUpdatedIterator{contract: _Proofsys.contract, event: "SignerUpdated", logs: logs, sub: sub}, nil
}

// WatchSignerUpdated is a free log subscription operation binding the contract event 0x2d025324f0a785e8c12d0a0d91a9caa49df4ef20ff87e0df7213a1d4f3157beb.
//
// Solidity: event SignerUpdated(address indexed previousSigner, address indexed newSigner)
func (_Proofsys *ProofsysFilterer) WatchSignerUpdated(opts *bind.WatchOpts, sink chan<- *ProofsysSignerUpdated, previousSigner []common.Address, newSigner []common.Address) (event.Subscription, error) {

	var previousSignerRule []interface{}
	for _, previousSignerItem := range previousSigner {
		previousSignerRule = append(previousSignerRule, previousSignerItem)
	}
	var newSignerRule []interface{}
	for _, newSignerItem := range newSigner {
		newSignerRule = append(newSignerRule, newSignerItem)
	}

	logs, sub, err := _Proofsys.contract.WatchLogs(opts, "SignerUpdated", previousSignerRule, newSignerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ProofsysSignerUpdated)
				if err := _Proofsys.contract.UnpackLog(event, "SignerUpdated", log); err != nil {
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

// ParseSignerUpdated is a log parse operation binding the contract event 0x2d025324f0a785e8c12d0a0d91a9caa49df4ef20ff87e0df7213a1d4f3157beb.
//
// Solidity: event SignerUpdated(address indexed previousSigner, address indexed newSigner)
func (_Proofsys *ProofsysFilterer) ParseSignerUpdated(log types.Log) (*ProofsysSignerUpdated, error) {
	event := new(ProofsysSignerUpdated)
	if err := _Proofsys.contract.UnpackLog(event, "SignerUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
