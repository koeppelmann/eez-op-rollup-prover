// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package rollupmgr

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

// RollupmgrMetaData contains all meta data concerning the Rollupmgr contract.
var RollupmgrMetaData = &bind.MetaData{
	ABI: "[{\"type\":\"constructor\",\"inputs\":[{\"name\":\"rollupsRegistry\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"_owner\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"_threshold\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"proofSystems\",\"type\":\"address[]\",\"internalType\":\"address[]\"},{\"name\":\"vkeys\",\"type\":\"bytes32[]\",\"internalType\":\"bytes32[]\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"ROLLUPS\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"addProofSystem\",\"inputs\":[{\"name\":\"proofSystem\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"vkey\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"getVkeysFromProofSystems\",\"inputs\":[{\"name\":\"proofSystems\",\"type\":\"address[]\",\"internalType\":\"address[]\"}],\"outputs\":[{\"name\":\"vkeys\",\"type\":\"bytes32[]\",\"internalType\":\"bytes32[]\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"owner\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"removeProofSystem\",\"inputs\":[{\"name\":\"proofSystem\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"rollupContractRegistered\",\"inputs\":[{\"name\":\"_rollupId\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"rollupId\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"setStateRoot\",\"inputs\":[{\"name\":\"newStateRoot\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setThreshold\",\"inputs\":[{\"name\":\"newThreshold\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setVerificationKey\",\"inputs\":[{\"name\":\"proofSystem\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"newVkey\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"threshold\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"transferOwnership\",\"inputs\":[{\"name\":\"newOwner\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"verificationKey\",\"inputs\":[{\"name\":\"proofSystem\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"vkey\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"view\"},{\"type\":\"event\",\"name\":\"OwnershipTransferred\",\"inputs\":[{\"name\":\"previousOwner\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"newOwner\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"ProofSystemAdded\",\"inputs\":[{\"name\":\"proofSystem\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"verificationKey\",\"type\":\"bytes32\",\"indexed\":false,\"internalType\":\"bytes32\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"ProofSystemRemoved\",\"inputs\":[{\"name\":\"proofSystem\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"StateRootEscape\",\"inputs\":[{\"name\":\"newStateRoot\",\"type\":\"bytes32\",\"indexed\":false,\"internalType\":\"bytes32\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"ThresholdChanged\",\"inputs\":[{\"name\":\"newThreshold\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"VerificationKeyUpdated\",\"inputs\":[{\"name\":\"proofSystem\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"newVerificationKey\",\"type\":\"bytes32\",\"indexed\":false,\"internalType\":\"bytes32\"}],\"anonymous\":false},{\"type\":\"error\",\"name\":\"InvalidConfig\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"NotOwner\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"NotRollupsRegistry\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"ProofSystemAlreadyAllowed\",\"inputs\":[{\"name\":\"proofSystem\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"ProofSystemNotAllowed\",\"inputs\":[{\"name\":\"proofSystem\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"ThresholdNotMet\",\"inputs\":[{\"name\":\"submitted\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"required\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]}]",
	Bin: "0x60a0604052346101ed57610a948038038061001981610209565b92833981019060a0818303126101ed5761003281610242565b61003e60208301610242565b60408301516060840151909291906001600160401b0381116101ed5784019385601f860112156101ed5784519461007c61007787610256565b610209565b9560208088838152019160051b830101918883116101ed57602001905b8282106101f1575050506080810151906001600160401b0382116101ed57019480601f870112156101ed5785516100d261007782610256565b9660208089848152019260051b8201019283116101ed57602001905b8282106101dd575050506001600160a01b038216156101a8576001600160a01b03169081156101a85783518551036101a85760805260018060a01b03195f5416175f556002555f5b81518110156101b7576001600160a01b03610151828461026d565b511661015d828561026d565b51156101a857805f52600360205260405f20546101965790600191610182828661026d565b51905f52600360205260405f205501610136565b635878623f60e11b5f5260045260245ffd5b6306b7c75960e31b5f5260045ffd5b6040516107fe908161029682396080518181816103c80152818161066c01526106e30152f35b81518152602091820191016100ee565b5f80fd5b602080916101fe84610242565b815201910190610099565b6040519190601f01601f191682016001600160401b0381118382101761022e57604052565b634e487b7160e01b5f52604160045260245ffd5b51906001600160a01b03821682036101ed57565b6001600160401b03811161022e5760051b60200190565b80518210156102815760209160051b010190565b634e487b7160e01b5f52603260045260245ffdfe6080806040526004361015610012575f80fd5b5f905f3560e01c9081631987f89714610712575080632ef01a75146106ce57806342cde4e8146106b157806374a8ffc4146106595780638da5cb5b14610632578063960bfe04146105db578063ad5d53511461053a578063c99e08d51461048b578063d307b7431461039c578063d97e4a8f14610300578063eea5eec7146102c7578063f2fde38b146102345763fb4e73ce146100ad575f80fd5b346102315760203660031901126102315760043567ffffffffffffffff811161022f573660238201121561022f57806004013567ffffffffffffffff811161022b576024820191602436918360051b01011161022b57600254808210610213575061011781610778565b906101256040519283610742565b80825261013181610778565b602083019390601f1901368537845b82811061018c5750505090604051928392602084019060208552518091526040840192915b818110610173575050500390f35b8251845285945060209384019390920191600101610165565b6001600160a01b036101a76101a2838686610790565b6107b4565b1686526003602052604086205480156101ea5784518210156101d657600582901b850160200152600101610140565b634e487b7160e01b87526032600452602487fd5b6024876101fb6101a2858888610790565b632ce975db60e21b82526001600160a01b0316600452fd5b6305bc216760e51b8452600491909152602452604482fd5b8280fd5b505b80fd5b50346102315760203660031901126102315761024e61072c565b8154906001600160a01b03821690338290036102b8576001600160a01b03169182156102a9576001600160a01b031916821783557f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e08380a380f35b6306b7c75960e31b8452600484fd5b6330cd747160e01b8452600484fd5b5034610231576020366003190112610231576020906040906001600160a01b036102ef61072c565b168152600383522054604051908152f35b50346102315760203660031901126102315761031a61072c565b81546001600160a01b0316330361038d576001600160a01b0316808252600360205260408220541561037b5780825260036020528160408120557fc37d448cc261a57e962ef8f52ee14d7d425f2c49312bef3907b61f5cedc3eda18280a280f35b632ce975db60e21b8252600452602490fd5b6330cd747160e01b8252600482fd5b5034610478576020366003190112610478575f54600435906001600160a01b0316330361047c576001547f00000000000000000000000000000000000000000000000000000000000000006001600160a01b031690813b15610478575f916044839260405194859384926315ee8f8b60e11b845260048401528760248401525af1801561046d57610457575b5060207f332cb801cc233238da18836ad2c6338134e30f41fbcb9290bafa9f2f14cd8d7991604051908152a180f35b6104649192505f90610742565b5f906020610428565b6040513d5f823e3d90fd5b5f80fd5b6330cd747160e01b5f5260045ffd5b34610478576040366003190112610478576104a461072c565b5f5460243591906001600160a01b0316330361047c57811561052b576001600160a01b03165f818152600360205260409020549091906105185760207fe536f233c69520d2c365bffd1be5f2dafb937d58460c322821cb8d04b8794b6191835f52600382528060405f2055604051908152a2005b50635878623f60e11b5f5260045260245ffd5b6306b7c75960e31b5f5260045ffd5b346104785760403660031901126104785761055361072c565b5f5460243591906001600160a01b0316330361047c57811561052b576001600160a01b03165f81815260036020526040902054909190156105c85760207fbbd5cf29bfda0a9ab13b746170b7f79b4faf6c0a12c9af86089d216912e6313391835f52600382528060405f2055604051908152a2005b50632ce975db60e21b5f5260045260245ffd5b34610478576020366003190112610478575f54600435906001600160a01b0316330361047c576020817f6c4ce60fd690e1216286a10b875c5662555f10774484e58142cedd7a90781baa92600255604051908152a1005b34610478575f366003190112610478575f546040516001600160a01b039091168152602090f35b34610478576020366003190112610478577f00000000000000000000000000000000000000000000000000000000000000006001600160a01b031633036106a257600435600155005b63e6b9616d60e01b5f5260045ffd5b34610478575f366003190112610478576020600254604051908152f35b34610478575f366003190112610478576040517f00000000000000000000000000000000000000000000000000000000000000006001600160a01b03168152602090f35b34610478575f366003190112610478576020906001548152f35b600435906001600160a01b038216820361047857565b90601f8019910116810190811067ffffffffffffffff82111761076457604052565b634e487b7160e01b5f52604160045260245ffd5b67ffffffffffffffff81116107645760051b60200190565b91908110156107a05760051b0190565b634e487b7160e01b5f52603260045260245ffd5b356001600160a01b0381168103610478579056fea264697066735822122034e6140147cae5c338095677c2d98b6b4fade58378ef2e33520681ba192dcdf064736f6c63430008220033",
}

// RollupmgrABI is the input ABI used to generate the binding from.
// Deprecated: Use RollupmgrMetaData.ABI instead.
var RollupmgrABI = RollupmgrMetaData.ABI

// RollupmgrBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use RollupmgrMetaData.Bin instead.
var RollupmgrBin = RollupmgrMetaData.Bin

// DeployRollupmgr deploys a new Ethereum contract, binding an instance of Rollupmgr to it.
func DeployRollupmgr(auth *bind.TransactOpts, backend bind.ContractBackend, rollupsRegistry common.Address, _owner common.Address, _threshold *big.Int, proofSystems []common.Address, vkeys [][32]byte) (common.Address, *types.Transaction, *Rollupmgr, error) {
	parsed, err := RollupmgrMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(RollupmgrBin), backend, rollupsRegistry, _owner, _threshold, proofSystems, vkeys)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Rollupmgr{RollupmgrCaller: RollupmgrCaller{contract: contract}, RollupmgrTransactor: RollupmgrTransactor{contract: contract}, RollupmgrFilterer: RollupmgrFilterer{contract: contract}}, nil
}

// Rollupmgr is an auto generated Go binding around an Ethereum contract.
type Rollupmgr struct {
	RollupmgrCaller     // Read-only binding to the contract
	RollupmgrTransactor // Write-only binding to the contract
	RollupmgrFilterer   // Log filterer for contract events
}

// RollupmgrCaller is an auto generated read-only Go binding around an Ethereum contract.
type RollupmgrCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// RollupmgrTransactor is an auto generated write-only Go binding around an Ethereum contract.
type RollupmgrTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// RollupmgrFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type RollupmgrFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// RollupmgrSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type RollupmgrSession struct {
	Contract     *Rollupmgr        // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// RollupmgrCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type RollupmgrCallerSession struct {
	Contract *RollupmgrCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts    // Call options to use throughout this session
}

// RollupmgrTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type RollupmgrTransactorSession struct {
	Contract     *RollupmgrTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts    // Transaction auth options to use throughout this session
}

// RollupmgrRaw is an auto generated low-level Go binding around an Ethereum contract.
type RollupmgrRaw struct {
	Contract *Rollupmgr // Generic contract binding to access the raw methods on
}

// RollupmgrCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type RollupmgrCallerRaw struct {
	Contract *RollupmgrCaller // Generic read-only contract binding to access the raw methods on
}

// RollupmgrTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type RollupmgrTransactorRaw struct {
	Contract *RollupmgrTransactor // Generic write-only contract binding to access the raw methods on
}

// NewRollupmgr creates a new instance of Rollupmgr, bound to a specific deployed contract.
func NewRollupmgr(address common.Address, backend bind.ContractBackend) (*Rollupmgr, error) {
	contract, err := bindRollupmgr(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Rollupmgr{RollupmgrCaller: RollupmgrCaller{contract: contract}, RollupmgrTransactor: RollupmgrTransactor{contract: contract}, RollupmgrFilterer: RollupmgrFilterer{contract: contract}}, nil
}

// NewRollupmgrCaller creates a new read-only instance of Rollupmgr, bound to a specific deployed contract.
func NewRollupmgrCaller(address common.Address, caller bind.ContractCaller) (*RollupmgrCaller, error) {
	contract, err := bindRollupmgr(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &RollupmgrCaller{contract: contract}, nil
}

// NewRollupmgrTransactor creates a new write-only instance of Rollupmgr, bound to a specific deployed contract.
func NewRollupmgrTransactor(address common.Address, transactor bind.ContractTransactor) (*RollupmgrTransactor, error) {
	contract, err := bindRollupmgr(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &RollupmgrTransactor{contract: contract}, nil
}

// NewRollupmgrFilterer creates a new log filterer instance of Rollupmgr, bound to a specific deployed contract.
func NewRollupmgrFilterer(address common.Address, filterer bind.ContractFilterer) (*RollupmgrFilterer, error) {
	contract, err := bindRollupmgr(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &RollupmgrFilterer{contract: contract}, nil
}

// bindRollupmgr binds a generic wrapper to an already deployed contract.
func bindRollupmgr(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := RollupmgrMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Rollupmgr *RollupmgrRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Rollupmgr.Contract.RollupmgrCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Rollupmgr *RollupmgrRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Rollupmgr.Contract.RollupmgrTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Rollupmgr *RollupmgrRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Rollupmgr.Contract.RollupmgrTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Rollupmgr *RollupmgrCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Rollupmgr.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Rollupmgr *RollupmgrTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Rollupmgr.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Rollupmgr *RollupmgrTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Rollupmgr.Contract.contract.Transact(opts, method, params...)
}

// ROLLUPS is a free data retrieval call binding the contract method 0x2ef01a75.
//
// Solidity: function ROLLUPS() view returns(address)
func (_Rollupmgr *RollupmgrCaller) ROLLUPS(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Rollupmgr.contract.Call(opts, &out, "ROLLUPS")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// ROLLUPS is a free data retrieval call binding the contract method 0x2ef01a75.
//
// Solidity: function ROLLUPS() view returns(address)
func (_Rollupmgr *RollupmgrSession) ROLLUPS() (common.Address, error) {
	return _Rollupmgr.Contract.ROLLUPS(&_Rollupmgr.CallOpts)
}

// ROLLUPS is a free data retrieval call binding the contract method 0x2ef01a75.
//
// Solidity: function ROLLUPS() view returns(address)
func (_Rollupmgr *RollupmgrCallerSession) ROLLUPS() (common.Address, error) {
	return _Rollupmgr.Contract.ROLLUPS(&_Rollupmgr.CallOpts)
}

// GetVkeysFromProofSystems is a free data retrieval call binding the contract method 0xfb4e73ce.
//
// Solidity: function getVkeysFromProofSystems(address[] proofSystems) view returns(bytes32[] vkeys)
func (_Rollupmgr *RollupmgrCaller) GetVkeysFromProofSystems(opts *bind.CallOpts, proofSystems []common.Address) ([][32]byte, error) {
	var out []interface{}
	err := _Rollupmgr.contract.Call(opts, &out, "getVkeysFromProofSystems", proofSystems)

	if err != nil {
		return *new([][32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([][32]byte)).(*[][32]byte)

	return out0, err

}

// GetVkeysFromProofSystems is a free data retrieval call binding the contract method 0xfb4e73ce.
//
// Solidity: function getVkeysFromProofSystems(address[] proofSystems) view returns(bytes32[] vkeys)
func (_Rollupmgr *RollupmgrSession) GetVkeysFromProofSystems(proofSystems []common.Address) ([][32]byte, error) {
	return _Rollupmgr.Contract.GetVkeysFromProofSystems(&_Rollupmgr.CallOpts, proofSystems)
}

// GetVkeysFromProofSystems is a free data retrieval call binding the contract method 0xfb4e73ce.
//
// Solidity: function getVkeysFromProofSystems(address[] proofSystems) view returns(bytes32[] vkeys)
func (_Rollupmgr *RollupmgrCallerSession) GetVkeysFromProofSystems(proofSystems []common.Address) ([][32]byte, error) {
	return _Rollupmgr.Contract.GetVkeysFromProofSystems(&_Rollupmgr.CallOpts, proofSystems)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Rollupmgr *RollupmgrCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Rollupmgr.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Rollupmgr *RollupmgrSession) Owner() (common.Address, error) {
	return _Rollupmgr.Contract.Owner(&_Rollupmgr.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Rollupmgr *RollupmgrCallerSession) Owner() (common.Address, error) {
	return _Rollupmgr.Contract.Owner(&_Rollupmgr.CallOpts)
}

// RollupId is a free data retrieval call binding the contract method 0x1987f897.
//
// Solidity: function rollupId() view returns(uint256)
func (_Rollupmgr *RollupmgrCaller) RollupId(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Rollupmgr.contract.Call(opts, &out, "rollupId")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// RollupId is a free data retrieval call binding the contract method 0x1987f897.
//
// Solidity: function rollupId() view returns(uint256)
func (_Rollupmgr *RollupmgrSession) RollupId() (*big.Int, error) {
	return _Rollupmgr.Contract.RollupId(&_Rollupmgr.CallOpts)
}

// RollupId is a free data retrieval call binding the contract method 0x1987f897.
//
// Solidity: function rollupId() view returns(uint256)
func (_Rollupmgr *RollupmgrCallerSession) RollupId() (*big.Int, error) {
	return _Rollupmgr.Contract.RollupId(&_Rollupmgr.CallOpts)
}

// Threshold is a free data retrieval call binding the contract method 0x42cde4e8.
//
// Solidity: function threshold() view returns(uint256)
func (_Rollupmgr *RollupmgrCaller) Threshold(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Rollupmgr.contract.Call(opts, &out, "threshold")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Threshold is a free data retrieval call binding the contract method 0x42cde4e8.
//
// Solidity: function threshold() view returns(uint256)
func (_Rollupmgr *RollupmgrSession) Threshold() (*big.Int, error) {
	return _Rollupmgr.Contract.Threshold(&_Rollupmgr.CallOpts)
}

// Threshold is a free data retrieval call binding the contract method 0x42cde4e8.
//
// Solidity: function threshold() view returns(uint256)
func (_Rollupmgr *RollupmgrCallerSession) Threshold() (*big.Int, error) {
	return _Rollupmgr.Contract.Threshold(&_Rollupmgr.CallOpts)
}

// VerificationKey is a free data retrieval call binding the contract method 0xeea5eec7.
//
// Solidity: function verificationKey(address proofSystem) view returns(bytes32 vkey)
func (_Rollupmgr *RollupmgrCaller) VerificationKey(opts *bind.CallOpts, proofSystem common.Address) ([32]byte, error) {
	var out []interface{}
	err := _Rollupmgr.contract.Call(opts, &out, "verificationKey", proofSystem)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// VerificationKey is a free data retrieval call binding the contract method 0xeea5eec7.
//
// Solidity: function verificationKey(address proofSystem) view returns(bytes32 vkey)
func (_Rollupmgr *RollupmgrSession) VerificationKey(proofSystem common.Address) ([32]byte, error) {
	return _Rollupmgr.Contract.VerificationKey(&_Rollupmgr.CallOpts, proofSystem)
}

// VerificationKey is a free data retrieval call binding the contract method 0xeea5eec7.
//
// Solidity: function verificationKey(address proofSystem) view returns(bytes32 vkey)
func (_Rollupmgr *RollupmgrCallerSession) VerificationKey(proofSystem common.Address) ([32]byte, error) {
	return _Rollupmgr.Contract.VerificationKey(&_Rollupmgr.CallOpts, proofSystem)
}

// AddProofSystem is a paid mutator transaction binding the contract method 0xc99e08d5.
//
// Solidity: function addProofSystem(address proofSystem, bytes32 vkey) returns()
func (_Rollupmgr *RollupmgrTransactor) AddProofSystem(opts *bind.TransactOpts, proofSystem common.Address, vkey [32]byte) (*types.Transaction, error) {
	return _Rollupmgr.contract.Transact(opts, "addProofSystem", proofSystem, vkey)
}

// AddProofSystem is a paid mutator transaction binding the contract method 0xc99e08d5.
//
// Solidity: function addProofSystem(address proofSystem, bytes32 vkey) returns()
func (_Rollupmgr *RollupmgrSession) AddProofSystem(proofSystem common.Address, vkey [32]byte) (*types.Transaction, error) {
	return _Rollupmgr.Contract.AddProofSystem(&_Rollupmgr.TransactOpts, proofSystem, vkey)
}

// AddProofSystem is a paid mutator transaction binding the contract method 0xc99e08d5.
//
// Solidity: function addProofSystem(address proofSystem, bytes32 vkey) returns()
func (_Rollupmgr *RollupmgrTransactorSession) AddProofSystem(proofSystem common.Address, vkey [32]byte) (*types.Transaction, error) {
	return _Rollupmgr.Contract.AddProofSystem(&_Rollupmgr.TransactOpts, proofSystem, vkey)
}

// RemoveProofSystem is a paid mutator transaction binding the contract method 0xd97e4a8f.
//
// Solidity: function removeProofSystem(address proofSystem) returns()
func (_Rollupmgr *RollupmgrTransactor) RemoveProofSystem(opts *bind.TransactOpts, proofSystem common.Address) (*types.Transaction, error) {
	return _Rollupmgr.contract.Transact(opts, "removeProofSystem", proofSystem)
}

// RemoveProofSystem is a paid mutator transaction binding the contract method 0xd97e4a8f.
//
// Solidity: function removeProofSystem(address proofSystem) returns()
func (_Rollupmgr *RollupmgrSession) RemoveProofSystem(proofSystem common.Address) (*types.Transaction, error) {
	return _Rollupmgr.Contract.RemoveProofSystem(&_Rollupmgr.TransactOpts, proofSystem)
}

// RemoveProofSystem is a paid mutator transaction binding the contract method 0xd97e4a8f.
//
// Solidity: function removeProofSystem(address proofSystem) returns()
func (_Rollupmgr *RollupmgrTransactorSession) RemoveProofSystem(proofSystem common.Address) (*types.Transaction, error) {
	return _Rollupmgr.Contract.RemoveProofSystem(&_Rollupmgr.TransactOpts, proofSystem)
}

// RollupContractRegistered is a paid mutator transaction binding the contract method 0x74a8ffc4.
//
// Solidity: function rollupContractRegistered(uint256 _rollupId) returns()
func (_Rollupmgr *RollupmgrTransactor) RollupContractRegistered(opts *bind.TransactOpts, _rollupId *big.Int) (*types.Transaction, error) {
	return _Rollupmgr.contract.Transact(opts, "rollupContractRegistered", _rollupId)
}

// RollupContractRegistered is a paid mutator transaction binding the contract method 0x74a8ffc4.
//
// Solidity: function rollupContractRegistered(uint256 _rollupId) returns()
func (_Rollupmgr *RollupmgrSession) RollupContractRegistered(_rollupId *big.Int) (*types.Transaction, error) {
	return _Rollupmgr.Contract.RollupContractRegistered(&_Rollupmgr.TransactOpts, _rollupId)
}

// RollupContractRegistered is a paid mutator transaction binding the contract method 0x74a8ffc4.
//
// Solidity: function rollupContractRegistered(uint256 _rollupId) returns()
func (_Rollupmgr *RollupmgrTransactorSession) RollupContractRegistered(_rollupId *big.Int) (*types.Transaction, error) {
	return _Rollupmgr.Contract.RollupContractRegistered(&_Rollupmgr.TransactOpts, _rollupId)
}

// SetStateRoot is a paid mutator transaction binding the contract method 0xd307b743.
//
// Solidity: function setStateRoot(bytes32 newStateRoot) returns()
func (_Rollupmgr *RollupmgrTransactor) SetStateRoot(opts *bind.TransactOpts, newStateRoot [32]byte) (*types.Transaction, error) {
	return _Rollupmgr.contract.Transact(opts, "setStateRoot", newStateRoot)
}

// SetStateRoot is a paid mutator transaction binding the contract method 0xd307b743.
//
// Solidity: function setStateRoot(bytes32 newStateRoot) returns()
func (_Rollupmgr *RollupmgrSession) SetStateRoot(newStateRoot [32]byte) (*types.Transaction, error) {
	return _Rollupmgr.Contract.SetStateRoot(&_Rollupmgr.TransactOpts, newStateRoot)
}

// SetStateRoot is a paid mutator transaction binding the contract method 0xd307b743.
//
// Solidity: function setStateRoot(bytes32 newStateRoot) returns()
func (_Rollupmgr *RollupmgrTransactorSession) SetStateRoot(newStateRoot [32]byte) (*types.Transaction, error) {
	return _Rollupmgr.Contract.SetStateRoot(&_Rollupmgr.TransactOpts, newStateRoot)
}

// SetThreshold is a paid mutator transaction binding the contract method 0x960bfe04.
//
// Solidity: function setThreshold(uint256 newThreshold) returns()
func (_Rollupmgr *RollupmgrTransactor) SetThreshold(opts *bind.TransactOpts, newThreshold *big.Int) (*types.Transaction, error) {
	return _Rollupmgr.contract.Transact(opts, "setThreshold", newThreshold)
}

// SetThreshold is a paid mutator transaction binding the contract method 0x960bfe04.
//
// Solidity: function setThreshold(uint256 newThreshold) returns()
func (_Rollupmgr *RollupmgrSession) SetThreshold(newThreshold *big.Int) (*types.Transaction, error) {
	return _Rollupmgr.Contract.SetThreshold(&_Rollupmgr.TransactOpts, newThreshold)
}

// SetThreshold is a paid mutator transaction binding the contract method 0x960bfe04.
//
// Solidity: function setThreshold(uint256 newThreshold) returns()
func (_Rollupmgr *RollupmgrTransactorSession) SetThreshold(newThreshold *big.Int) (*types.Transaction, error) {
	return _Rollupmgr.Contract.SetThreshold(&_Rollupmgr.TransactOpts, newThreshold)
}

// SetVerificationKey is a paid mutator transaction binding the contract method 0xad5d5351.
//
// Solidity: function setVerificationKey(address proofSystem, bytes32 newVkey) returns()
func (_Rollupmgr *RollupmgrTransactor) SetVerificationKey(opts *bind.TransactOpts, proofSystem common.Address, newVkey [32]byte) (*types.Transaction, error) {
	return _Rollupmgr.contract.Transact(opts, "setVerificationKey", proofSystem, newVkey)
}

// SetVerificationKey is a paid mutator transaction binding the contract method 0xad5d5351.
//
// Solidity: function setVerificationKey(address proofSystem, bytes32 newVkey) returns()
func (_Rollupmgr *RollupmgrSession) SetVerificationKey(proofSystem common.Address, newVkey [32]byte) (*types.Transaction, error) {
	return _Rollupmgr.Contract.SetVerificationKey(&_Rollupmgr.TransactOpts, proofSystem, newVkey)
}

// SetVerificationKey is a paid mutator transaction binding the contract method 0xad5d5351.
//
// Solidity: function setVerificationKey(address proofSystem, bytes32 newVkey) returns()
func (_Rollupmgr *RollupmgrTransactorSession) SetVerificationKey(proofSystem common.Address, newVkey [32]byte) (*types.Transaction, error) {
	return _Rollupmgr.Contract.SetVerificationKey(&_Rollupmgr.TransactOpts, proofSystem, newVkey)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Rollupmgr *RollupmgrTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _Rollupmgr.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Rollupmgr *RollupmgrSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Rollupmgr.Contract.TransferOwnership(&_Rollupmgr.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Rollupmgr *RollupmgrTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Rollupmgr.Contract.TransferOwnership(&_Rollupmgr.TransactOpts, newOwner)
}

// RollupmgrOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the Rollupmgr contract.
type RollupmgrOwnershipTransferredIterator struct {
	Event *RollupmgrOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *RollupmgrOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(RollupmgrOwnershipTransferred)
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
		it.Event = new(RollupmgrOwnershipTransferred)
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
func (it *RollupmgrOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *RollupmgrOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// RollupmgrOwnershipTransferred represents a OwnershipTransferred event raised by the Rollupmgr contract.
type RollupmgrOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Rollupmgr *RollupmgrFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*RollupmgrOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Rollupmgr.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &RollupmgrOwnershipTransferredIterator{contract: _Rollupmgr.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Rollupmgr *RollupmgrFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *RollupmgrOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Rollupmgr.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(RollupmgrOwnershipTransferred)
				if err := _Rollupmgr.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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
func (_Rollupmgr *RollupmgrFilterer) ParseOwnershipTransferred(log types.Log) (*RollupmgrOwnershipTransferred, error) {
	event := new(RollupmgrOwnershipTransferred)
	if err := _Rollupmgr.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// RollupmgrProofSystemAddedIterator is returned from FilterProofSystemAdded and is used to iterate over the raw logs and unpacked data for ProofSystemAdded events raised by the Rollupmgr contract.
type RollupmgrProofSystemAddedIterator struct {
	Event *RollupmgrProofSystemAdded // Event containing the contract specifics and raw log

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
func (it *RollupmgrProofSystemAddedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(RollupmgrProofSystemAdded)
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
		it.Event = new(RollupmgrProofSystemAdded)
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
func (it *RollupmgrProofSystemAddedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *RollupmgrProofSystemAddedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// RollupmgrProofSystemAdded represents a ProofSystemAdded event raised by the Rollupmgr contract.
type RollupmgrProofSystemAdded struct {
	ProofSystem     common.Address
	VerificationKey [32]byte
	Raw             types.Log // Blockchain specific contextual infos
}

// FilterProofSystemAdded is a free log retrieval operation binding the contract event 0xe536f233c69520d2c365bffd1be5f2dafb937d58460c322821cb8d04b8794b61.
//
// Solidity: event ProofSystemAdded(address indexed proofSystem, bytes32 verificationKey)
func (_Rollupmgr *RollupmgrFilterer) FilterProofSystemAdded(opts *bind.FilterOpts, proofSystem []common.Address) (*RollupmgrProofSystemAddedIterator, error) {

	var proofSystemRule []interface{}
	for _, proofSystemItem := range proofSystem {
		proofSystemRule = append(proofSystemRule, proofSystemItem)
	}

	logs, sub, err := _Rollupmgr.contract.FilterLogs(opts, "ProofSystemAdded", proofSystemRule)
	if err != nil {
		return nil, err
	}
	return &RollupmgrProofSystemAddedIterator{contract: _Rollupmgr.contract, event: "ProofSystemAdded", logs: logs, sub: sub}, nil
}

// WatchProofSystemAdded is a free log subscription operation binding the contract event 0xe536f233c69520d2c365bffd1be5f2dafb937d58460c322821cb8d04b8794b61.
//
// Solidity: event ProofSystemAdded(address indexed proofSystem, bytes32 verificationKey)
func (_Rollupmgr *RollupmgrFilterer) WatchProofSystemAdded(opts *bind.WatchOpts, sink chan<- *RollupmgrProofSystemAdded, proofSystem []common.Address) (event.Subscription, error) {

	var proofSystemRule []interface{}
	for _, proofSystemItem := range proofSystem {
		proofSystemRule = append(proofSystemRule, proofSystemItem)
	}

	logs, sub, err := _Rollupmgr.contract.WatchLogs(opts, "ProofSystemAdded", proofSystemRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(RollupmgrProofSystemAdded)
				if err := _Rollupmgr.contract.UnpackLog(event, "ProofSystemAdded", log); err != nil {
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

// ParseProofSystemAdded is a log parse operation binding the contract event 0xe536f233c69520d2c365bffd1be5f2dafb937d58460c322821cb8d04b8794b61.
//
// Solidity: event ProofSystemAdded(address indexed proofSystem, bytes32 verificationKey)
func (_Rollupmgr *RollupmgrFilterer) ParseProofSystemAdded(log types.Log) (*RollupmgrProofSystemAdded, error) {
	event := new(RollupmgrProofSystemAdded)
	if err := _Rollupmgr.contract.UnpackLog(event, "ProofSystemAdded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// RollupmgrProofSystemRemovedIterator is returned from FilterProofSystemRemoved and is used to iterate over the raw logs and unpacked data for ProofSystemRemoved events raised by the Rollupmgr contract.
type RollupmgrProofSystemRemovedIterator struct {
	Event *RollupmgrProofSystemRemoved // Event containing the contract specifics and raw log

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
func (it *RollupmgrProofSystemRemovedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(RollupmgrProofSystemRemoved)
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
		it.Event = new(RollupmgrProofSystemRemoved)
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
func (it *RollupmgrProofSystemRemovedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *RollupmgrProofSystemRemovedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// RollupmgrProofSystemRemoved represents a ProofSystemRemoved event raised by the Rollupmgr contract.
type RollupmgrProofSystemRemoved struct {
	ProofSystem common.Address
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterProofSystemRemoved is a free log retrieval operation binding the contract event 0xc37d448cc261a57e962ef8f52ee14d7d425f2c49312bef3907b61f5cedc3eda1.
//
// Solidity: event ProofSystemRemoved(address indexed proofSystem)
func (_Rollupmgr *RollupmgrFilterer) FilterProofSystemRemoved(opts *bind.FilterOpts, proofSystem []common.Address) (*RollupmgrProofSystemRemovedIterator, error) {

	var proofSystemRule []interface{}
	for _, proofSystemItem := range proofSystem {
		proofSystemRule = append(proofSystemRule, proofSystemItem)
	}

	logs, sub, err := _Rollupmgr.contract.FilterLogs(opts, "ProofSystemRemoved", proofSystemRule)
	if err != nil {
		return nil, err
	}
	return &RollupmgrProofSystemRemovedIterator{contract: _Rollupmgr.contract, event: "ProofSystemRemoved", logs: logs, sub: sub}, nil
}

// WatchProofSystemRemoved is a free log subscription operation binding the contract event 0xc37d448cc261a57e962ef8f52ee14d7d425f2c49312bef3907b61f5cedc3eda1.
//
// Solidity: event ProofSystemRemoved(address indexed proofSystem)
func (_Rollupmgr *RollupmgrFilterer) WatchProofSystemRemoved(opts *bind.WatchOpts, sink chan<- *RollupmgrProofSystemRemoved, proofSystem []common.Address) (event.Subscription, error) {

	var proofSystemRule []interface{}
	for _, proofSystemItem := range proofSystem {
		proofSystemRule = append(proofSystemRule, proofSystemItem)
	}

	logs, sub, err := _Rollupmgr.contract.WatchLogs(opts, "ProofSystemRemoved", proofSystemRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(RollupmgrProofSystemRemoved)
				if err := _Rollupmgr.contract.UnpackLog(event, "ProofSystemRemoved", log); err != nil {
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

// ParseProofSystemRemoved is a log parse operation binding the contract event 0xc37d448cc261a57e962ef8f52ee14d7d425f2c49312bef3907b61f5cedc3eda1.
//
// Solidity: event ProofSystemRemoved(address indexed proofSystem)
func (_Rollupmgr *RollupmgrFilterer) ParseProofSystemRemoved(log types.Log) (*RollupmgrProofSystemRemoved, error) {
	event := new(RollupmgrProofSystemRemoved)
	if err := _Rollupmgr.contract.UnpackLog(event, "ProofSystemRemoved", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// RollupmgrStateRootEscapeIterator is returned from FilterStateRootEscape and is used to iterate over the raw logs and unpacked data for StateRootEscape events raised by the Rollupmgr contract.
type RollupmgrStateRootEscapeIterator struct {
	Event *RollupmgrStateRootEscape // Event containing the contract specifics and raw log

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
func (it *RollupmgrStateRootEscapeIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(RollupmgrStateRootEscape)
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
		it.Event = new(RollupmgrStateRootEscape)
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
func (it *RollupmgrStateRootEscapeIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *RollupmgrStateRootEscapeIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// RollupmgrStateRootEscape represents a StateRootEscape event raised by the Rollupmgr contract.
type RollupmgrStateRootEscape struct {
	NewStateRoot [32]byte
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterStateRootEscape is a free log retrieval operation binding the contract event 0x332cb801cc233238da18836ad2c6338134e30f41fbcb9290bafa9f2f14cd8d79.
//
// Solidity: event StateRootEscape(bytes32 newStateRoot)
func (_Rollupmgr *RollupmgrFilterer) FilterStateRootEscape(opts *bind.FilterOpts) (*RollupmgrStateRootEscapeIterator, error) {

	logs, sub, err := _Rollupmgr.contract.FilterLogs(opts, "StateRootEscape")
	if err != nil {
		return nil, err
	}
	return &RollupmgrStateRootEscapeIterator{contract: _Rollupmgr.contract, event: "StateRootEscape", logs: logs, sub: sub}, nil
}

// WatchStateRootEscape is a free log subscription operation binding the contract event 0x332cb801cc233238da18836ad2c6338134e30f41fbcb9290bafa9f2f14cd8d79.
//
// Solidity: event StateRootEscape(bytes32 newStateRoot)
func (_Rollupmgr *RollupmgrFilterer) WatchStateRootEscape(opts *bind.WatchOpts, sink chan<- *RollupmgrStateRootEscape) (event.Subscription, error) {

	logs, sub, err := _Rollupmgr.contract.WatchLogs(opts, "StateRootEscape")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(RollupmgrStateRootEscape)
				if err := _Rollupmgr.contract.UnpackLog(event, "StateRootEscape", log); err != nil {
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

// ParseStateRootEscape is a log parse operation binding the contract event 0x332cb801cc233238da18836ad2c6338134e30f41fbcb9290bafa9f2f14cd8d79.
//
// Solidity: event StateRootEscape(bytes32 newStateRoot)
func (_Rollupmgr *RollupmgrFilterer) ParseStateRootEscape(log types.Log) (*RollupmgrStateRootEscape, error) {
	event := new(RollupmgrStateRootEscape)
	if err := _Rollupmgr.contract.UnpackLog(event, "StateRootEscape", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// RollupmgrThresholdChangedIterator is returned from FilterThresholdChanged and is used to iterate over the raw logs and unpacked data for ThresholdChanged events raised by the Rollupmgr contract.
type RollupmgrThresholdChangedIterator struct {
	Event *RollupmgrThresholdChanged // Event containing the contract specifics and raw log

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
func (it *RollupmgrThresholdChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(RollupmgrThresholdChanged)
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
		it.Event = new(RollupmgrThresholdChanged)
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
func (it *RollupmgrThresholdChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *RollupmgrThresholdChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// RollupmgrThresholdChanged represents a ThresholdChanged event raised by the Rollupmgr contract.
type RollupmgrThresholdChanged struct {
	NewThreshold *big.Int
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterThresholdChanged is a free log retrieval operation binding the contract event 0x6c4ce60fd690e1216286a10b875c5662555f10774484e58142cedd7a90781baa.
//
// Solidity: event ThresholdChanged(uint256 newThreshold)
func (_Rollupmgr *RollupmgrFilterer) FilterThresholdChanged(opts *bind.FilterOpts) (*RollupmgrThresholdChangedIterator, error) {

	logs, sub, err := _Rollupmgr.contract.FilterLogs(opts, "ThresholdChanged")
	if err != nil {
		return nil, err
	}
	return &RollupmgrThresholdChangedIterator{contract: _Rollupmgr.contract, event: "ThresholdChanged", logs: logs, sub: sub}, nil
}

// WatchThresholdChanged is a free log subscription operation binding the contract event 0x6c4ce60fd690e1216286a10b875c5662555f10774484e58142cedd7a90781baa.
//
// Solidity: event ThresholdChanged(uint256 newThreshold)
func (_Rollupmgr *RollupmgrFilterer) WatchThresholdChanged(opts *bind.WatchOpts, sink chan<- *RollupmgrThresholdChanged) (event.Subscription, error) {

	logs, sub, err := _Rollupmgr.contract.WatchLogs(opts, "ThresholdChanged")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(RollupmgrThresholdChanged)
				if err := _Rollupmgr.contract.UnpackLog(event, "ThresholdChanged", log); err != nil {
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

// ParseThresholdChanged is a log parse operation binding the contract event 0x6c4ce60fd690e1216286a10b875c5662555f10774484e58142cedd7a90781baa.
//
// Solidity: event ThresholdChanged(uint256 newThreshold)
func (_Rollupmgr *RollupmgrFilterer) ParseThresholdChanged(log types.Log) (*RollupmgrThresholdChanged, error) {
	event := new(RollupmgrThresholdChanged)
	if err := _Rollupmgr.contract.UnpackLog(event, "ThresholdChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// RollupmgrVerificationKeyUpdatedIterator is returned from FilterVerificationKeyUpdated and is used to iterate over the raw logs and unpacked data for VerificationKeyUpdated events raised by the Rollupmgr contract.
type RollupmgrVerificationKeyUpdatedIterator struct {
	Event *RollupmgrVerificationKeyUpdated // Event containing the contract specifics and raw log

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
func (it *RollupmgrVerificationKeyUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(RollupmgrVerificationKeyUpdated)
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
		it.Event = new(RollupmgrVerificationKeyUpdated)
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
func (it *RollupmgrVerificationKeyUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *RollupmgrVerificationKeyUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// RollupmgrVerificationKeyUpdated represents a VerificationKeyUpdated event raised by the Rollupmgr contract.
type RollupmgrVerificationKeyUpdated struct {
	ProofSystem        common.Address
	NewVerificationKey [32]byte
	Raw                types.Log // Blockchain specific contextual infos
}

// FilterVerificationKeyUpdated is a free log retrieval operation binding the contract event 0xbbd5cf29bfda0a9ab13b746170b7f79b4faf6c0a12c9af86089d216912e63133.
//
// Solidity: event VerificationKeyUpdated(address indexed proofSystem, bytes32 newVerificationKey)
func (_Rollupmgr *RollupmgrFilterer) FilterVerificationKeyUpdated(opts *bind.FilterOpts, proofSystem []common.Address) (*RollupmgrVerificationKeyUpdatedIterator, error) {

	var proofSystemRule []interface{}
	for _, proofSystemItem := range proofSystem {
		proofSystemRule = append(proofSystemRule, proofSystemItem)
	}

	logs, sub, err := _Rollupmgr.contract.FilterLogs(opts, "VerificationKeyUpdated", proofSystemRule)
	if err != nil {
		return nil, err
	}
	return &RollupmgrVerificationKeyUpdatedIterator{contract: _Rollupmgr.contract, event: "VerificationKeyUpdated", logs: logs, sub: sub}, nil
}

// WatchVerificationKeyUpdated is a free log subscription operation binding the contract event 0xbbd5cf29bfda0a9ab13b746170b7f79b4faf6c0a12c9af86089d216912e63133.
//
// Solidity: event VerificationKeyUpdated(address indexed proofSystem, bytes32 newVerificationKey)
func (_Rollupmgr *RollupmgrFilterer) WatchVerificationKeyUpdated(opts *bind.WatchOpts, sink chan<- *RollupmgrVerificationKeyUpdated, proofSystem []common.Address) (event.Subscription, error) {

	var proofSystemRule []interface{}
	for _, proofSystemItem := range proofSystem {
		proofSystemRule = append(proofSystemRule, proofSystemItem)
	}

	logs, sub, err := _Rollupmgr.contract.WatchLogs(opts, "VerificationKeyUpdated", proofSystemRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(RollupmgrVerificationKeyUpdated)
				if err := _Rollupmgr.contract.UnpackLog(event, "VerificationKeyUpdated", log); err != nil {
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

// ParseVerificationKeyUpdated is a log parse operation binding the contract event 0xbbd5cf29bfda0a9ab13b746170b7f79b4faf6c0a12c9af86089d216912e63133.
//
// Solidity: event VerificationKeyUpdated(address indexed proofSystem, bytes32 newVerificationKey)
func (_Rollupmgr *RollupmgrFilterer) ParseVerificationKeyUpdated(log types.Log) (*RollupmgrVerificationKeyUpdated, error) {
	event := new(RollupmgrVerificationKeyUpdated)
	if err := _Rollupmgr.contract.UnpackLog(event, "VerificationKeyUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
