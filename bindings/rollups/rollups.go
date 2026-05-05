// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package rollups

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

// CrossChainCall is an auto generated low-level Go binding around an user-defined struct.
type CrossChainCall struct {
	TargetAddress  common.Address
	Value          *big.Int
	Data           []byte
	SourceAddress  common.Address
	SourceRollupId *big.Int
	RevertSpan     *big.Int
}

// ExecutionEntry is an auto generated low-level Go binding around an user-defined struct.
type ExecutionEntry struct {
	StateDeltas         []StateDelta
	CrossChainCallHash  [32]byte
	DestinationRollupId *big.Int
	Calls               []CrossChainCall
	NestedActions       []NestedAction
	CallCount           *big.Int
	ReturnData          []byte
	RollingHash         [32]byte
}

// LookupCall is an auto generated low-level Go binding around an user-defined struct.
type LookupCall struct {
	CrossChainCallHash       [32]byte
	DestinationRollupId      *big.Int
	ReturnData               []byte
	Failed                   bool
	CallNumber               uint64
	LastNestedActionConsumed uint64
	Calls                    []CrossChainCall
	RollingHash              [32]byte
}

// NestedAction is an auto generated low-level Go binding around an user-defined struct.
type NestedAction struct {
	CrossChainCallHash [32]byte
	CallCount          *big.Int
	ReturnData         []byte
}

// ProofSystemBatch is an auto generated low-level Go binding around an user-defined struct.
type ProofSystemBatch struct {
	ProofSystems                 []common.Address
	RollupIds                    []*big.Int
	Entries                      []ExecutionEntry
	LookupCalls                  []LookupCall
	TransientCount               *big.Int
	TransientLookupCallCount     *big.Int
	BlobIndices                  []*big.Int
	CallData                     []byte
	Proof                        [][]byte
	CrossProofSystemInteractions [32]byte
}

// StateDelta is an auto generated low-level Go binding around an user-defined struct.
type StateDelta struct {
	RollupId     *big.Int
	CurrentState [32]byte
	NewState     [32]byte
	EtherDelta   *big.Int
}

// RollupsMetaData contains all meta data concerning the Rollups contract.
var RollupsMetaData = &bind.MetaData{
	ABI: "[{\"type\":\"function\",\"name\":\"MAINNET_ROLLUP_ID\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"_transientExecutions\",\"inputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"crossChainCallHash\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"destinationRollupId\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"callCount\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"returnData\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"rollingHash\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"_transientLookupCalls\",\"inputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"crossChainCallHash\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"destinationRollupId\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"returnData\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"failed\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"callNumber\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"lastNestedActionConsumed\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"rollingHash\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"attemptApplyImmediate\",\"inputs\":[{\"name\":\"transientIdx\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"authorizedProxies\",\"inputs\":[{\"name\":\"proxy\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"originalAddress\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"originalRollupId\",\"type\":\"uint64\",\"internalType\":\"uint64\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"computeCrossChainCallHash\",\"inputs\":[{\"name\":\"targetRollupId\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"targetAddress\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"value\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"data\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"sourceAddress\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"sourceRollupId\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"pure\"},{\"type\":\"function\",\"name\":\"computeCrossChainProxyAddress\",\"inputs\":[{\"name\":\"originalAddress\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"originalRollupId\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"createCrossChainProxy\",\"inputs\":[{\"name\":\"originalAddress\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"originalRollupId\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"proxy\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"createRollup\",\"inputs\":[{\"name\":\"rollupContract\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"initialState\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"outputs\":[{\"name\":\"rollupId\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"executeCrossChainCall\",\"inputs\":[{\"name\":\"sourceAddress\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"callData\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[{\"name\":\"result\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"stateMutability\":\"payable\"},{\"type\":\"function\",\"name\":\"executeInContextAndRevert\",\"inputs\":[{\"name\":\"callCount\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"executeL2TX\",\"inputs\":[{\"name\":\"rollupId\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"result\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"lastVerifiedBlock\",\"inputs\":[{\"name\":\"_rollupId\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"postBatch\",\"inputs\":[{\"name\":\"batches\",\"type\":\"tuple[]\",\"internalType\":\"structProofSystemBatch[]\",\"components\":[{\"name\":\"proofSystems\",\"type\":\"address[]\",\"internalType\":\"address[]\"},{\"name\":\"rollupIds\",\"type\":\"uint256[]\",\"internalType\":\"uint256[]\"},{\"name\":\"entries\",\"type\":\"tuple[]\",\"internalType\":\"structExecutionEntry[]\",\"components\":[{\"name\":\"stateDeltas\",\"type\":\"tuple[]\",\"internalType\":\"structStateDelta[]\",\"components\":[{\"name\":\"rollupId\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"currentState\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"newState\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"etherDelta\",\"type\":\"int256\",\"internalType\":\"int256\"}]},{\"name\":\"crossChainCallHash\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"destinationRollupId\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"calls\",\"type\":\"tuple[]\",\"internalType\":\"structCrossChainCall[]\",\"components\":[{\"name\":\"targetAddress\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"value\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"data\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"sourceAddress\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"sourceRollupId\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"revertSpan\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"name\":\"nestedActions\",\"type\":\"tuple[]\",\"internalType\":\"structNestedAction[]\",\"components\":[{\"name\":\"crossChainCallHash\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"callCount\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"returnData\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]},{\"name\":\"callCount\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"returnData\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"rollingHash\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}]},{\"name\":\"lookupCalls\",\"type\":\"tuple[]\",\"internalType\":\"structLookupCall[]\",\"components\":[{\"name\":\"crossChainCallHash\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"destinationRollupId\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"returnData\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"failed\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"callNumber\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"lastNestedActionConsumed\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"calls\",\"type\":\"tuple[]\",\"internalType\":\"structCrossChainCall[]\",\"components\":[{\"name\":\"targetAddress\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"value\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"data\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"sourceAddress\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"sourceRollupId\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"revertSpan\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"name\":\"rollingHash\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}]},{\"name\":\"transientCount\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"transientLookupCallCount\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"blobIndices\",\"type\":\"uint256[]\",\"internalType\":\"uint256[]\"},{\"name\":\"callData\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"proof\",\"type\":\"bytes[]\",\"internalType\":\"bytes[]\"},{\"name\":\"crossProofSystemInteractions\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}]}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"queueCursor\",\"inputs\":[{\"name\":\"_rollupId\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"queueLength\",\"inputs\":[{\"name\":\"_rollupId\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"rollupCounter\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"rollups\",\"inputs\":[{\"name\":\"rollupId\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"rollupContract\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"stateRoot\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"etherBalance\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"setRollupContract\",\"inputs\":[{\"name\":\"rollupId\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"newContract\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setStateRoot\",\"inputs\":[{\"name\":\"rollupId\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"newStateRoot\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"staticCallLookup\",\"inputs\":[{\"name\":\"sourceAddress\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"callData\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"stateMutability\":\"view\"},{\"type\":\"event\",\"name\":\"BatchPosted\",\"inputs\":[{\"name\":\"subBatchCount\",\"type\":\"uint256\",\"indexed\":true,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"CallResult\",\"inputs\":[{\"name\":\"entryIndex\",\"type\":\"uint256\",\"indexed\":true,\"internalType\":\"uint256\"},{\"name\":\"callNumber\",\"type\":\"uint256\",\"indexed\":true,\"internalType\":\"uint256\"},{\"name\":\"success\",\"type\":\"bool\",\"indexed\":false,\"internalType\":\"bool\"},{\"name\":\"returnData\",\"type\":\"bytes\",\"indexed\":false,\"internalType\":\"bytes\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"CrossChainCallExecuted\",\"inputs\":[{\"name\":\"crossChainCallHash\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"},{\"name\":\"proxy\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"sourceAddress\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"callData\",\"type\":\"bytes\",\"indexed\":false,\"internalType\":\"bytes\"},{\"name\":\"value\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"CrossChainProxyCreated\",\"inputs\":[{\"name\":\"proxy\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"originalAddress\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"originalRollupId\",\"type\":\"uint256\",\"indexed\":true,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"EntryExecuted\",\"inputs\":[{\"name\":\"entryIndex\",\"type\":\"uint256\",\"indexed\":true,\"internalType\":\"uint256\"},{\"name\":\"rollingHash\",\"type\":\"bytes32\",\"indexed\":false,\"internalType\":\"bytes32\"},{\"name\":\"callsProcessed\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"nestedActionsConsumed\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"ExecutionConsumed\",\"inputs\":[{\"name\":\"crossChainCallHash\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"},{\"name\":\"rollupId\",\"type\":\"uint256\",\"indexed\":true,\"internalType\":\"uint256\"},{\"name\":\"cursor\",\"type\":\"uint256\",\"indexed\":true,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"ImmediateEntrySkipped\",\"inputs\":[{\"name\":\"transientIdx\",\"type\":\"uint256\",\"indexed\":true,\"internalType\":\"uint256\"},{\"name\":\"revertData\",\"type\":\"bytes\",\"indexed\":false,\"internalType\":\"bytes\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"L2ExecutionPerformed\",\"inputs\":[{\"name\":\"rollupId\",\"type\":\"uint256\",\"indexed\":true,\"internalType\":\"uint256\"},{\"name\":\"newState\",\"type\":\"bytes32\",\"indexed\":false,\"internalType\":\"bytes32\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"L2TXExecuted\",\"inputs\":[{\"name\":\"rollupId\",\"type\":\"uint256\",\"indexed\":true,\"internalType\":\"uint256\"},{\"name\":\"cursor\",\"type\":\"uint256\",\"indexed\":true,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"NestedActionConsumed\",\"inputs\":[{\"name\":\"entryIndex\",\"type\":\"uint256\",\"indexed\":true,\"internalType\":\"uint256\"},{\"name\":\"nestedNumber\",\"type\":\"uint256\",\"indexed\":true,\"internalType\":\"uint256\"},{\"name\":\"crossChainCallHash\",\"type\":\"bytes32\",\"indexed\":false,\"internalType\":\"bytes32\"},{\"name\":\"callCount\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"RevertSpanExecuted\",\"inputs\":[{\"name\":\"entryIndex\",\"type\":\"uint256\",\"indexed\":true,\"internalType\":\"uint256\"},{\"name\":\"startCallNumber\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"span\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"RollupContractChanged\",\"inputs\":[{\"name\":\"rollupId\",\"type\":\"uint256\",\"indexed\":true,\"internalType\":\"uint256\"},{\"name\":\"previousContract\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"newContract\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"RollupCreated\",\"inputs\":[{\"name\":\"rollupId\",\"type\":\"uint256\",\"indexed\":true,\"internalType\":\"uint256\"},{\"name\":\"rollupContract\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"initialState\",\"type\":\"bytes32\",\"indexed\":false,\"internalType\":\"bytes32\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"StateUpdated\",\"inputs\":[{\"name\":\"rollupId\",\"type\":\"uint256\",\"indexed\":true,\"internalType\":\"uint256\"},{\"name\":\"newStateRoot\",\"type\":\"bytes32\",\"indexed\":false,\"internalType\":\"bytes32\"}],\"anonymous\":false},{\"type\":\"error\",\"name\":\"ContextResult\",\"inputs\":[{\"name\":\"rollingHash\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"lastNestedActionConsumed\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"currentCallNumber\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"type\":\"error\",\"name\":\"DuplicateProofSystem\",\"inputs\":[{\"name\":\"proofSystem\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"EtherDeltaMismatch\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"ExecutionNotFound\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"ExecutionNotInCurrentBlock\",\"inputs\":[{\"name\":\"rollupId\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"type\":\"error\",\"name\":\"InsufficientRollupBalance\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidProof\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidProofSystemConfig\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidRollupContract\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"L2TXNotAllowedDuringExecution\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"NotRollupContract\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"NotSelf\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"PostBatchReentry\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"RollingHashMismatch\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"RollupAlreadyVerifiedThisBlock\",\"inputs\":[{\"name\":\"rollupId\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"type\":\"error\",\"name\":\"RollupBatchActiveThisBlock\",\"inputs\":[{\"name\":\"rollupId\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"type\":\"error\",\"name\":\"RollupNotInBatch\",\"inputs\":[{\"name\":\"rollupId\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"type\":\"error\",\"name\":\"StateRootMismatch\",\"inputs\":[{\"name\":\"rollupId\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"type\":\"error\",\"name\":\"TransientCountExceedsEntries\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"TransientLookupCallCountExceedsLookupCalls\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"UnauthorizedProxy\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"UnconsumedCalls\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"UnconsumedNestedActions\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"UnexpectedContextRevert\",\"inputs\":[{\"name\":\"revertData\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]}]",
	Bin: "0x6080806040523460155761462a908161001a8239f35b5f80fdfe60806040526004361015610011575f80fd5b5f3560e01c8063199350cf146101645780631f314db11461015f578063206bbf071461015a5780632bdd1f16146101555780632dd721201461015057806335bcb88a1461014b578063360d95b61461014657806337b99a561461014157806348cae3e21461013c5780637aa4da2e146101375780639af5325914610132578063a32718321461012d578063b761ba7e14610128578063b794e5a314610123578063cc21dd3c1461011e578063ccdcf58114610119578063d3f6036014610114578063d4ccaf871461010f578063d84220371461010a578063f41558a5146101055763ff2f742f14610100575f80fd5b610cb7565b610c5c565b610be1565b610a7c565b610a04565b6107fd565b610690565b610647565b61061c565b610600565b6105eb565b6105bd565b6104fd565b6104a7565b610453565b610426565b6103e9565b610327565b6102fd565b610276565b610177565b5f91031261017357565b5f80fd5b34610173575f3660031901126101735760206040515f8152f35b6001600160a01b0381160361017357565b634e487b7160e01b5f52604160045260245ffd5b608081019081106001600160401b038211176101d157604052565b6101a2565b90601f801991011681019081106001600160401b038211176101d157604052565b604051906102066060836101d6565b565b604051906102066040836101d6565b9061020660405192836101d6565b6001600160401b0381116101d157601f01601f191660200190565b92919261024c82610225565b9161025a60405193846101d6565b829481845281830111610173578281602093845f960137010152565b346101735760c03660031901126101735760043560243561029681610191565b60443591606435916001600160401b0383116101735736602384011215610173576102f9936102d26102e9943690602481600401359101610240565b90608435926102e084610191565b60a43594610e6d565b6040519081529081906020820190565b0390f35b34610173576020366003190112610173576004355f526002602052602060405f2054604051908152f35b34610173576040366003190112610173576004356024355f828152600160205260409020546001600160a01b031633036103da5761036d825f52600260205260405f2090565b5443146103c5576103c0817f3060356fe01e6782d366a741b8a2f65e53c64fefa69c21cf85ad79a846fd472e9260016103ae865f52600160205260405f2090565b01556040519081529081906020820190565b0390a2005b63033feec760e01b5f52600482905260245b5ffd5b633484b4e160e11b5f5260045ffd5b3461017357604036600319011261017357602061041460043561040b81610191565b60243590611b1c565b6040516001600160a01b039091168152f35b34610173576020366003190112610173576004355f5260026020526020600360405f200154604051908152f35b346101735760203660031901126101735760043561047081610191565b6001600160a01b039081165f90815260036020908152604091829020548251938116845260a01c6001600160401b03169083015290f35b34610173576020366003190112610173576004353033036104ee576104cb90611d1b565b5060035c60055c60045c9163251cb09960e21b5f5260045260245260445260645ffd5b6314e1dbf760e11b5f5260045ffd5b34610173576020366003190112610173576004355f5260026020526020600160405f200154604051908152f35b60406003198201126101735760043561054281610191565b916024356001600160401b0381116101735782602382011215610173578060040135926001600160401b0384116101735760248483010111610173576024019190565b805180835260209291819084018484015e5f828201840152601f01601f1916010190565b9060206105ba928181520190610585565b90565b34610173576102f96105d76105d13661052a565b91610ecc565b604051918291602083526020830190610585565b6102f96105d76105fa3661052a565b916110c6565b34610173575f3660031901126101735760205f54604051908152f35b3461017357604036600319011261017357602061041460043561063e81610191565b602435906111ad565b34610173576020366003190112610173576004355f526001602052606060405f2060018060a01b0381541690600260018201549101549060405192835260208301526040820152f35b34610173576040366003190112610173576004356106ad81610191565b602435906001600160a01b038116801580156107f4575b6107e5576107465f54926106d7846112d7565b5f556106f36106e46101f7565b6001600160a01b039092168252565b8460208201525f6040820152610711845f52600160205260405f2090565b815181546001600160a01b0319166001600160a01b039190911617815590600290604090602081015160018501550151910155565b803b1561017357604051631d2a3ff160e21b815260048101839052925f8460248183865af19384156107e0576107b384927f65aca4a0612183373cd8f999716ed850923221124a6c3b5c55b90e964cbe34d9926102f9976107c6575b506040519081529081906020820190565b0390a36040519081529081906020820190565b806107d45f6107da936101d6565b80610169565b5f6107a2565b6112fd565b6334387b3b60e21b5f5260045ffd5b503081146106c4565b3461017357602036600319011261017357600435805f52600260205260405f205443036108875760045c6108785761086c816102f9925f526002602052600360405f200154817f110e37308dbd81a8e4d42e316af40e7cf78577d19f52c170928664fe4a6a377f5f80a3612290565b604051918291826105a9565b63bd2663f560e01b5f5260045ffd5b632ae204bd60e01b5f5260045260245ffd5b634e487b7160e01b5f52603260045260245ffd5b6005548110156108c95760055f52600660205f20910201905f90565b610899565b80548210156108c9575f52600660205f20910201905f90565b90600182811c92168015610915575b602083101461090157565b634e487b7160e01b5f52602260045260245ffd5b91607f16916108f6565b9060405191825f825492610932846108e7565b808452936001811690811561099b5750600114610957575b50610206925003836101d6565b90505f9291925260205f20905f915b81831061097f575050906020610206928201015f61094a565b6020919350806001915483858901015201910190918492610966565b90506020925061020694915060ff191682840152151560051b8201015f61094a565b9397969591946109ec6001600160401b039460c09786948852602088015260e0604088015260e0870190610585565b98151560608601521660808401521660a08201520152565b346101735760203660031901126101735760043560055481101561017357610a2b906108ad565b508054906102f9600182015491610a446002820161091f565b9260056003830154920154916040519586956001600160401b038360481c169360ff6001600160401b038560081c16941692886109bd565b34610173576020366003190112610173576004353033036104ee57610aa081610bac565b509060015d5f60025d6005810154610ae6610ae0610ac2600785015494611350565b92610acb611327565b610ad361132d565b610adb611333565b611d1b565b91613fc7565b91610aef613d91565b9060035c03610b9d5760045c600382015403610b8e57600460055c91015403610b7f57610b1c905f612452565b03610b70576040805160035c815260045c602082015260055c9181019190915260015c907f99403777f3e390813f268a6307c0cb1b1ae83271187473a2c07e34f81e4a293690606090a2610b6e61132d565b005b63378c57b960e21b5f5260045ffd5b63efacbf5560e01b5f5260045ffd5b632888d75b60e11b5f5260045ffd5b633ce8ed9f60e21b5f5260045ffd5b6004548110156108c95760045f5260205f209060031b01905f90565b80548210156108c9575f5260205f209060031b01905f90565b346101735760203660031901126101735760043560045481101561017357610c0890610bac565b506001810154600282015491610c526005820154916007610c2b6006830161091f565b9101549260405195869586526020860152604085015260a0606085015260a0840190610585565b9060808301520390f35b34610173576020366003190112610173576004356001600160401b03811161017357366023820112156101735780600401356001600160401b038111610173573660248260051b84010111610173576024610b6e92016118d7565b3461017357604036600319011261017357602435600435610cd782610191565b610d08610cfc610cef835f52600160205260405f2090565b546001600160a01b031690565b6001600160a01b031690565b33036103da576001600160a01b0382169182158015610dc7575b6107e557610d5b90610d3c835f52600160205260405f2090565b80546001600160a01b0319166001600160a01b03909216919091179055565b813b1561017357604051631d2a3ff160e21b8152600481018290525f8160248183875af180156107e057610db3575b5033907fcb1f89865b4ce341946f7baead20a32f5c5c8694fcfc0d2e7109ad2f7754a7e35f80a4005b806107d45f610dc1936101d6565b5f610d8a565b50308314610d22565b9290610e2e91610e0a6040519485936020850197885260018060a01b031660408501525f606085015260c0608085015260e0840190610585565b6001600160a01b0390911660a08301525f60c083015203601f1981018352826101d6565b51902090565b93610e0a610e2e93946040519586946020860198895260018060a01b03166040860152606085015260c0608085015260e0840190610585565b94610ea9610e2e949592939560405196879560208701998a5260018060a01b03166040870152606086015260c0608086015260e0850190610585565b6001600160a01b0390921660a084015260c083015203601f1981018352826101d6565b335f908152600360205260409020546001600160a01b03811690811561106c57610f2093610f1991610f119060a01c6001600160401b03165b6001600160401b031690565b953691610240565b9084610dd0565b6001600160401b0360045c166001600160401b0360055c16925f600554905b818110610ffb575050610f5c6002915f52600260205260405f2090565b01805493905f5b858110610f7957630ed6bc7560e41b5f5260045ffd5b610f8381836108ce565b508581541480610fd6575b80610fb1575b610fa15750600101610f63565b9450505050506105ba9150611fa1565b50600381015460481c6001600160401b03166001600160401b03808616911614610f94565b50600381015460081c6001600160401b03166001600160401b03808716911614610f8e565b611004816108ad565b508581541480611047575b80611022575b610fa15750600101610f3f565b50600381015460481c6001600160401b03166001600160401b03808916911614611015565b50600381015460081c6001600160401b03166001600160401b0380871691161461100f565b63729ee4a560e11b5f5260045ffd5b908060209392818452848401375f828201840152601f01601f1916010190565b949392916040926110c19260018060a01b0316875260606020880152606087019161107b565b930152565b335f9081526003602052604090205490926001600160a01b0382169291831561106c576110fe9060a01c6001600160401b0316610f05565b92611111845f52600260205260405f2090565b54430361118757611152857fad4275800b3b34f59dea39a4bd58ae171a852e3555a2ed1a6f9bad83eb6fb59892611149368688610240565b90349088610e34565b948592611168604051928392339734928561109b565b0390a360045c61117d576105ba913491612377565b506105ba9061205f565b632ae204bd60e01b5f52600484905260245ffd5b805191908290602001825e015f815290565b6105ba916111f66112b4610cfc936111f661126c6040516020810190611204816111f6878b86909160349282526bffffffffffffffffffffffff199060601b1660208201520190565b03601f1981018352826101d6565b519020956112526104579461121b60208701610217565b9580875261419e6020880139604080513060208201526001600160a01b0390921690820152606081019290925281608081016111f6565b60405192839161126660208401809761119b565b9061119b565b5190206040516001600160f81b0319602082019081523060601b6bffffffffffffffffffffffff19166021830152603582019590955260558101919091529182906075820190565b5190206001600160a01b031690565b634e487b7160e01b5f52601160045260245ffd5b5f1981146112e55760010190565b6112c3565b634e487b7160e01b5f525f60045260245ffd5b6040513d5f823e3d90fd5b60045d565b60055d565b60035d565b60025d565b5f60025d565b60015d565b5f60035d565b5f60045d565b5f60055d565b6001600160401b0381116101d15760051b60200190565b90815461135c81611339565b9261136a60405194856101d6565b81845260208401905f5260205f205f915b8383106113885750505050565b6004602060019260405161139b816101b6565b855481528486015483820152600286015460408201526003860154606082015281520192019201919061137b565b91908110156108c95760051b8101359061013e1981360301821215610173570190565b903590601e198136030182121561017357018035906001600160401b03821161017357602001918160051b3603831361017357565b91908110156108c95760051b0190565b3d1561145b573d9061144282610225565b9161145060405193846101d6565b82523d5f602084013e565b606090565b90600182018092116112e557565b919082018092116112e557565b8054905f81558161148a575050565b600282901b916001600160fe1b038116036112e5575f5260205f205f5b8281106114b357505050565b805f600360049385018281558260018201558260028201550155016114a7565b90600160401b81116101d1578154918181558282106114f157505050565b6001600160fe1b03831683036112e557600282901b916001600160fe1b038116036112e5575f528060205f20019160021b03905f5b82811061153257505050565b805f60036004938501828155826001820155826002820155015501611526565b5f5b82811061156057505050565b5f82820155600101611554565b61157781546108e7565b9081611581575050565b81601f5f9311600114611592575055565b818352602083206115af91601f0160051c84190190600101611552565b808252602082209081548360011b9084198560031b1c191617905555565b8054905f8155816115dc575050565b816006029160068304036112e5575f5260205f205f5b8281106115fe57505050565b805f6005600693850182815582600182015561161c6002820161156d565b8260038201558260048201550155016115f2565b90600160401b81116101d15781549181815582821061164e57505050565b826006029260068404036112e557816006029160068304036112e5575f528060205f20019103905f5b82811061168357505050565b805f600560069385018281558260018201556116a16002820161156d565b826003820155826004820155015501611677565b8054905f8155816116c4575050565b816003029160038304036112e5575f5260205f205f5b8281106116e657505050565b80611700600260039385015f81555f60018201550161156d565b016116da565b90600160401b81116101d15781549181815582821061172457505050565b826003029260038404036112e557816003029160038304036112e5575f528060205f20019103905f5b82811061175957505050565b80611773600260039385015f81555f60018201550161156d565b0161174d565b6004545f600455806117885750565b600381901b906001600160fd1b038116036112e55760045f5260205f205f5b8281106117b357505050565b805f600760089385016117c58161147b565b8260018201558260028201556117dd600382016115cd565b6117e9600482016116b5565b8260058201556117fb6006820161156d565b0155016117a7565b6005545f600555806118125750565b806006029060068204036112e55760055f5260205f205f5b82811061183657505050565b805f600560069385018281558260018201556118546002820161156d565b826003820155611866600482016115cd565b01550161182a565b8054905f81558161187d575050565b816006029160068304036112e5575f5260205f205f5b82811061189f57505050565b805f600560069385018281558260018201556118bd6002820161156d565b8260038201556118cf600482016115cd565b015501611893565b600454611b0d578115611afe575f5b828110611ae257505f5b828110611ab357505f5b828110611a6a575061190c8282613c0f565b5f5c6004541180611a54575b156119a9575f5c303b156101735760405163d4ccaf8760e01b81526004810182905261198991905f8160248183305af19081611995575b5061199057807f62cc6fa8d0d1b1170559640dfa2d36932712a5d761c0924ea23aabf3b602cb3c61198161086c611431565b0390a2611460565b5f5d61190c565b611460565b806107d45f6119a3936101d6565b5f61194f565b815f5c6004541180611a4a575b6119fd575b6119d7916119c7611779565b6119cf611803565b5f805d613ca9565b7fd6f8d71ce42a799b91f399271f4b0e91f85eb87fac7bb2cedd4b3a52fad361825f80a2565b50333b156101735760405163455d16c760e01b8152905f8260048183335af19081156107e0576119d7928492611a36575b5091506119bb565b806107d45f611a44936101d6565b5f611a2e565b50333b15156119b6565b506001611a615f5c610bac565b50015415611918565b611a82611a788285856113c9565b60208101906113ec565b5f5b818110611a96575050506001016118fa565b80611aad611aa76001938587611421565b356132f2565b01611a84565b80611adc611acc611ac760019487876113c9565b61294e565b611ad78387876113c9565b612f7e565b016118f0565b80611af8611af360019386866113c9565b612575565b016118e6565b6328a1d92760e21b5f5260045ffd5b638b4f56d960e01b5f5260045ffd5b91906040516020810190611b53816111f6878686909160349282526bffffffffffffffffffffffff199060601b1660208201520190565b519020604051610457808201908282106001600160401b038311176101d15784878493611ba09361419e86393081526001600160a01b039091166020820152604081019190915260600190565b03905ff580156107e0576001600160a01b031692611c1b611bbf610208565b6001600160a01b0383811682526001600160401b03851660208084019182528883165f90815260039091526040902092518354915192166001600160e01b03199091161760a09190911b67ffffffffffffffff60a01b16179055565b6001600160a01b0316837f318838c0aec28d3eb1d3a48be7a8b2576548cde8922741021503e81b01a4c4475f80a4565b9060405160c081018181106001600160401b038211176101d15760405282546001600160a01b031681526001830154602082015291829060a090600590611c946002820161091f565b604085015260038101546001841b5f1901166001600160a01b03166060850152600481015460808501520154910152565b6001600160a01b0390911681526040602082018190526105ba92910190610585565b9190915f83820193841291129080158216911516176112e557565b6040906105ba9392151581528160208201520190610585565b905f91611d26613d91565b600301905f5b818110611d3857505050565b6005611d4660045c856108ce565b50015480611ec25750611d64611d5e60045c856108ce565b50611c4b565b611d77611d7260045c6112d7565b611308565b611d8260045c613e57565b6060810180515f9182916001600160a01b031690611da660808601928351906111ad565b6001600160a01b038181165f908152600360205260409020549193911615611ea4575b5050602084019384516111f6611e0b6040611dea855160018060a01b031690565b940151604051928391602083019663532f083960e01b885260248401611cc5565b51925af190611e18611431565b905180151580611e9d575b611e84575b509081611e3c611e7c94938360045c613e8c565b7f6ab31fd4e1fedcfcbb1a56d82ca9c2aa41d65466dff790e7287a415348b7576560015c9160045c93611e7460405192839283611d02565b0390a36112d7565b905b90611d2c565b611e9490611e7c94939298611ce7565b96909192611e28565b5082611e23565b90519051611eba916001600160a01b0316611b1c565b505f80611dc9565b60045c915f6005611ed385886108ce565b500155303b1561017357604051631bdccd2b60e11b815260048101839052611f8293905f8160248183305af19081611f8d575b50611f8857611f37611f2d611f32611f24611f1f611431565b613dc5565b93909193611308565b61130d565b611312565b826005611f4483896108ce565b500155604080519182526020820184905260015c917ff805bc66587a02266edd3197dc8dc57351fa8ae0b55a1e75f270f70d80159a7c9190a261146e565b90611e7e565b611f37565b806107d45f611f9b936101d6565b5f611f06565b60048101805480611fdb575b5050600381015460ff16611fc75760026105ba910161091f565b6002611fd3910161091f565b602081519101fd5b611fe481611339565b91611ff260405193846101d6565b81835260208301905f5260205f205f915b838310612028575050505061201790613ec7565b600582015403610b9d575f80611fad565b6006602060019261203885611c4b565b815201920192019190612003565b80548210156108c9575f52600360205f20910201905f90565b90612068613d91565b9060055c90612079611f2d836112d7565b60048301805483108061227c575b6121f75750506001600160401b036120a760045c6001600160401b031690565b9116915f600554905b8181106121735750506120d06002809201545f52600260205260405f2090565b01925f928454935b8481106120ee57630ed6bc7560e41b5f5260045ffd5b806120fb600192886108ce565b5060038101548460ff82169182612168575b8261214d575b82612137575b5050612127575b50016120d8565b61213090611fa1565b505f612120565b60481c6001600160401b0316149050845f612119565b600881901c6001600160401b03908116908a16149250612113565b83548814925061210d565b8061217f6001926108ad565b5060038101548760ff821691826121ec575b826121d1575b826121bb575b50506121ab575b50016120b0565b6121b490611fa1565b505f6121a4565b60481c6001600160401b0316149050875f61219d565b600881901c6001600160401b03908116908a16149250612197565b83548b149250612191565b6105ba935082919461227061221b61221460029661227695612046565b5094611460565b918260015c7fed18c2605a6158b1b555df8746ce7bd8bffd2154ceefbe38fab02b889a400612600188015480946122646040519283928360209093929193604081019481520152565b0390a3610adb83613f5d565b50613f8a565b0161091f565b50846122888483612046565b505414612087565b6004541561232c575f5c6122a3816112d7565b5f5d8060045481101561231d576122b990610bac565b506122c261131c565b600181015461231d57612300826006936105ba955f7f207e371295f3ff0efe443424ce128c96f8ecea6a25636fc38ad7c222c446699c8180a4611322565b6122765f600583015460078401549061231885611350565b61246a565b630ed6bc7560e41b5f5260045ffd5b61233e815f52600260205260405f2090565b6003810180549061234e826112d7565b9055600181920190815481101561231d5761236891610bc8565b5061237283611317565b6122c2565b60045490929015612405575f5c9161238e836112d7565b5f5d8260045481101561231d576123a490610bac565b50916123ae61131c565b8160018401540361231d576123f1846006956105ba97612276957f207e371295f3ff0efe443424ce128c96f8ecea6a25636fc38ad7c222c446699c5f80a4611322565b600583015460078401549061231885611350565b612417835f52600260205260405f2090565b9160038301805490612428826112d7565b9055600181940190815481101561231d5761244291610bc8565b509161244d85611317565b6123ae565b81810392915f1380158285131691841216176112e557565b90610ae061247d91949394610acb611327565b92612486613d91565b9060035c03610b9d5760045c600382015403610b8e57600460055c91015403610b7f576124b291612452565b03610b70576040805160035c815260045c602082015260055c9181019190915260015c907f99403777f3e390813f268a6307c0cb1b1ae83271187473a2c07e34f81e4a293690606090a261020661132d565b356105ba81610191565b91908110156108c95760051b8101359060fe1981360301821215610173570190565b903590601e198136030182121561017357018035906001600160401b03821161017357602001918160071b3603831361017357565b91908110156108c95760071b0190565b9061258082806113ec565b905015611afe5761259182806113ec565b90506125a16101008401846113ec565b91905003611afe57602082016125b781846113ec565b905015611afe575f92835b6125cc83836113ec565b905085101561261a576125e9856125e385856113ec565b90611421565b3590811115611afe5761260a610cfc610cef835f52600160205260405f2090565b15611afe576001909401936125c2565b505f93509091835b61262c83806113ec565b90508510156126865761264b612646866125e386806113ec565b612504565b906001600160a01b03908116908216111561266b57600190940193612622565b6308f396ef60e21b5f526001600160a01b031660045260245ffd5b50919092505f5b6040840161269b81866113ec565b905082101561275f57816126b26126b892876113ec565b9061250e565b6126d76126db6126c885886113ec565b9290604085013593849161410d565b1590565b61274d5750806126ea91612530565b905f5b8281106126ff5750505060010161268d565b6127216126d761270f878a6113ec565b61271a858888612565565b359161410d565b61272d576001016126ed565b6103d79261273a92612565565b63785cdf4160e11b5f5235600452602490565b63785cdf4160e11b5f5260045260245ffd5b5050915f5b6060820161277281846113ec565b90508210156127d7576127a36126d761278b87866113ec565b602061279b876126b2888b6113ec565b01359161410d565b6127b05750600101612764565b6020916126b26127c3926103d7956113ec565b63785cdf4160e11b5f520135600452602490565b505090915060808101356127ee60408301836113ec565b919050116128225761280960a08201359160608101906113ec565b9190501161281357565b633444431b60e11b5f5260045ffd5b63dd148d7560e01b5f5260045ffd5b9061283b82611339565b61284860405191826101d6565b8281528092612859601f1991611339565b01905f5b82811061286957505050565b80606060208093850101520161285d565b602081830312610173578051906001600160401b03821161017357019080601f830112156101735781516128ad81611339565b926128bb60405194856101d6565b81845260208085019260051b82010192831161017357602001905b8282106128e35750505090565b81518152602091820191016128d6565b60208082528101839052604001915f5b8181106129105750505090565b909192602080600192863561292481610191565b848060a01b031681520194019101919091612903565b80518210156108c95760209160051b010190565b6020810161296661295f82846113ec565b9050612831565b905f5b61297382856113ec565b9050811015612a43576129d25f6129a9610cfc610cfc610cef61299a876125e38a8d6113ec565b355f52600160205260405f2090565b6129b387806113ec565b604051637da739e760e11b8152948593928492839291600484016128f3565b03915afa9081156107e0575f91612a21575b506129ef828561293a565b526129fa818461293a565b50612a05818461293a565b5151612a1185806113ec565b91905003611afe57600101612969565b612a3d91503d805f833e612a3581836101d6565b81019061287a565b5f6129e4565b5050919050565b90612a5482611339565b612a6160405191826101d6565b8281528092612a72601f1991611339565b0190602036910137565b9035601e19823603018112156101735701602081359101916001600160401b038211610173578160071b3603831361017357565b916020908281520191905f5b818110612ac95750505090565b8235845260208084013590850152604080840135908501526060808401359085015260809384019390920191600101612abc565b9035601e19823603018112156101735701602081359101916001600160401b038211610173578160051b3603831361017357565b9035601e19823603018112156101735701602081359101916001600160401b03821161017357813603831361017357565b90602083828152019160208260051b8501019381935f915b848310612b8a5750505050505090565b909192939495601f1982820301835286359060be198636030182121561017357602080918760019401908135612bbf81610191565b858060a01b03168152828201358382015260a080612bf4612be36040860186612b31565b60c0604087015260c086019161107b565b936060810135612c0381610191565b8880841b031660608501526080810135608085015201359101529801930193019194939290612b7a565b90602083828152019060208160051b85010193835f915b838310612c545750505050505090565b909192939495601f198282030186528635605e1984360301811215610173576020612caa6001936060612c9a888596018035845285810135868501526040810190612b31565b919092816040820152019161107b565b980196019493019190612c44565b6020815261010060e0612d62612d3a612d1b612ce8612cd78880612a7c565b8760208a0152610120890191612ab0565b6020880135604088015260408801356060880152612d096060890189612afd565b888303601f190160808a015290612b62565b612d286080880188612afd565b878303601f190160a089015290612c2d565b60a086013560c0860152612d5160c0870187612b31565b868303601f1901858801529061107b565b93013591015290565b8015150361017357565b6001600160401b0381160361017357565b359061020682612d75565b60208152813560208201526020820135604082015261010060e0612d62612dcf612dbe6040870187612b31565b85606088015261012087019161107b565b612de96060870135612de081612d6b565b15156080870152565b612e08612df860808801612d86565b6001600160401b031660a0870152565b612e27612e1760a08801612d86565b6001600160401b031660c0870152565b612e3460c0870187612afd565b868303601f19018588015290612b62565b5f198101919082116112e557565b919082039182116112e557565b602080825281018390526001600160fb1b0383116101735760409260051b809284830137010190565b60206040818301928281528451809452019201905f5b818110612eac5750505090565b8251845260209384019390920191600101612e9f565b903590601e198136030182121561017357018035906001600160401b0382116101735760200191813603831361017357565b93611266611266946040809b9997612f1999966112669682526020820152019061119b565b91825260208201520190565b6020906105ba93928152019061119b565b908210156108c957612f4d9160051b810190612ec2565b9091565b9081602091031261017357516105ba81612d6b565b9392916020916110c19160408752604087019161107b565b919060c0830192612f99612f9285836113ec565b9050612a4a565b915f5b612fa686846113ec565b9050811015612fd35780612fc06001926125e389876113ec565b3549612fcc828761293a565b5201612f9c565b5093506040810191612fe8612f9284846113ec565b945f5b612ff585856113ec565b9050811015613038578061300f6001926126b288886113ec565b604051613024816111f6602082019485612cb8565b519020613031828a61293a565b5201612feb565b50925092606082019461304e612f9287856113ec565b935f5b61305b88866113ec565b905081101561309e57806130756001926126b28b896113ec565b60405161308a816111f6602082019485612d91565b519020613097828961293a565b5201613051565b509295509390926130ae43612e45565b409460208701956130bf87896113ec565b604051959186916130d39160208401612e60565b03601f19810186526130e590866101d6565b604051809360208201906130f891612e89565b03601f198101845261310a90846101d6565b6040518094602082019061311d91612e89565b03601f198101855261312f90856101d6565b6040518091602082019061314291612e89565b03601f198101825261315490826101d6565b61316160e08a018a612ec2565b369061316c92610240565b805190602001206101208a013591604051958695602087019842613190978b612ef4565b03601f19810182526131a290826101d6565b519020935f945b6131b382806113ec565b90508610156132ea576131c9612f9286846113ec565b935f5b6131d687856113ec565b905081101561320857806131f6896131f06001948961293a565b5161293a565b51613201828961293a565b52016131cc565b5095909360206111f661322661328b93604051928391858301612e89565b60405161323b816111f6858201948b86612f25565b519020613255610cfc610cfc612646876125e38a806113ec565b61326d856132676101008901896113ec565b90612f36565b604051636b40634160e01b8152958694859384939060048501612f66565b03915afa9081156107e0575f916132bc575b50156132ad5760010194926131a9565b6309bde33960e01b5f5260045ffd5b6132dd915060203d81116132e3575b6132d581836101d6565b810190612f51565b5f61329d565b503d6132cb565b505050915050565b805f52600260205260405f2090815443146133c35750438155600181018054613337575b5060038160025f9301805461332a57500155565b6133339061186e565b0155565b8054905f815581613349575b50613316565b600382901b916001600160fd1b038116036112e5575f5260205f205f5b8281106133735750613343565b805f600760089385016133858161147b565b82600182015582600282015561339d600382016115cd565b6133a9600482016116b5565b8260058201556133bb6006820161156d565b015501613366565b6366524c2160e01b5f5260045260245ffd5b91601f82116133e357505050565b8082116133ef57505050565b610206925f5260205f20916020601f830160051c921061341a575b601f82910160051c039101611552565b5f915061340a565b9092916001600160401b0381116101d1576134478161344184546108e7565b846133d5565b5f601f82116001146134855781906134769394955f9261347a575b50508160011b915f199060031b1c19161790565b9055565b013590505f80613462565b601f19821694613498845f5260205f2090565b915f5b8781106134d25750836001959697106134b9575b505050811b019055565b01355f19600384901b60f8161c191690555f80806134af565b9092602060018192868601358155019401910161349b565b91906134f68284611630565b5f9283526020832090805b83851061350f575050505050565b803560be198336030181121561017357820161354b813561352f81610191565b85546001600160a01b0319166001600160a01b03909116178555565b60208101356001850155600284016135666040830183612ec2565b906001600160401b0382116101d1576135898261358385546108e7565b856133d5565b5f90601f831160011461361457936020936135c2846001989560a0956006995f9261347a5750508160011b915f199060031b1c19161790565b90555b6135f76135d460608301612504565b60038a0180546001600160a01b0319166001600160a01b03909216919091179055565b608081013560048901550135600587015501930194019391613501565b601f19831691613627855f5260205f2090565b925f5b81811061367057508460a094600698946020989460019b988c9510613657575b505050811b0190556135c5565b01355f19600384901b60f8161c191690555f808061364a565b9193602060018192878701358155019501920161362a565b91906136948284611706565b5f9283526020832090805b8385106136ad575050505050565b8035605e198336030181121561017357820180358455602081013560018501556136df60028501916040810190612ec2565b906001600160401b0382116101d1576136fc8261358385546108e7565b5f90601f8311600114613741579261373283602094600197946003975f9261347a5750508160011b915f199060031b1c19161790565b90555b0193019401939161369f565b601f19831691613754855f5260205f2090565b925f5b818110613798575093600196936003969388938360209810613781575b505050811b019055613735565b01355f1983891b60f8161c191690555f8080613774565b91936020600181928787013581550195019201613757565b600454600160401b8110156101d1578060016137d192016004556004610bc8565b9190916138ae576137e28180612530565b906137ed82856114d3565b835f5260205f205f915b83831061387a575050505060e0600791602081013560018501556040810135600285015561383561382b60608301836113ec565b90600387016134ea565b61384f61384560808301836113ec565b9060048701613688565b60a0810135600585015561387361386960c0830183612ec2565b9060068701613422565b0135910155565b60046080600192803585556020810135848601556040810135600286015560608101356003860155019201920191906137f7565b6112ea565b8054600160401b8110156101d1576138d091600182018155610bc8565b9190916138ae576138e18180612530565b906138ec82856114d3565b835f5260205f205f915b83831061392a575050505060e0600791602081013560018501556040810135600285015561383561382b60608301836113ec565b60046080600192803585556020810135848601556040810135600286015560608101356003860155019201920191906138f6565b356105ba81612d6b565b356105ba81612d75565b600554600160401b8110156101d157806001613993920160055560056108ce565b9190916138ae578035825560208101356001830155600282016139b96040830183612ec2565b906001600160401b0382116101d1576139d68261358385546108e7565b5f90601f8311600114613ab557826005959360e09593613a0a935f9261347a5750508160011b915f199060031b1c19161790565b90555b613a9b60038501613a35613a236060850161395e565b829060ff801983541691151516179055565b613a68613a4460808501613968565b825468ffffffffffffffff00191660089190911b68ffffffffffffffff0016178255565b613a7460a08401613968565b67ffffffffffffffff60481b82549160481b169067ffffffffffffffff60481b1916179055565b613873613aab60c08301836113ec565b90600487016134ea565b601f19831691613ac8855f5260205f2090565b925f5b818110613b0d57509260019285926005989660e0989610613af4575b505050811b019055613a0d565b01355f19600384901b60f8161c191690555f8080613ae7565b91936020600181928787013581550195019201613acb565b8054600160401b8110156101d157613b42916001820181556108ce565b9190916138ae57803582556020810135600183015560028201613b686040830183612ec2565b906001600160401b0382116101d157613b858261358385546108e7565b5f90601f8311600114613bb957826005959360e09593613a0a935f9261347a5750508160011b915f199060031b1c19161790565b601f19831691613bcc855f5260205f2090565b925f5b818110613bf757509260019285926005989660e0989610613af457505050811b019055613a0d565b91936020600181928787013581550195019201613bcf565b905f5b818110613c1e57505050565b613c2b81838596956113c9565b905f9360408301936080840135955b86811015613c615780613c5b613c566001936126b28a8a6113ec565b6137b0565b01613c3a565b509560a084013595506060840194505f5b86811015613c995780613c93613c8e6001936126b28a8a6113ec565b613972565b01613c72565b5094509492506001915001613c12565b5f905b828210613cb857505050565b9091613cc58382846113c9565b91604083019160808401355b613cdb84866113ec565b9050811015613d295780613d236001613d0e6040613cfe83966126b28b8d6113ec565b01355f52600260205260405f2090565b01613d1d836126b2898b6113ec565b906138b3565b01613cd1565b50946060840192509060a08401355b613d4284866113ec565b9050811015613d815780613d7b6002613d666020613cfe6001966126b28b8d6113ec565b01613d75836126b2898b6113ec565b90613b25565b01613d38565b5090949360010192509050613cac565b60045415613da857613da460015c610bac565b5090565b60025c5f526002602052613da4600160405f200160015c90610bc8565b805160208201516001600160e01b03198116919060048210613e37575b50506001600160e01b031916631ae34f6760e21b01613e0e576024810151916064604483015192015190565b604051638664b1a560e01b815260206004820152908190613e33906024830190610585565b0390fd5b6001600160e01b031960049290920360031b82901b161690505f80613de2565b60035c906040519060208201928352600160f81b6040830152604182015260418152613e846061826101d6565b51902060035d565b613e84906111f660035c9460405194859360208501978852600160f91b60408601526041850152151560f81b6061840152606283019061119b565b5f9190825b8151841015613f5857613f506001915f80613ee7888761293a565b5160608101516080820151613f04916001600160a01b03166111ad565b906111f6613f3e6040613f1d845160018060a01b031690565b930151604051928391602083019563532f083960e01b875260248401611cc5565b51915afa613f4a611431565b91614173565b930192613ecc565b925050565b60035c906040519060208201928352600360f81b6040830152604182015260418152613e846061826101d6565b60035c906040519060208201928352600160fa1b6040830152604182015260418152613e846061826101d6565b600160ff1b81146112e5575f0390565b905f915f5b815181101561410957613fdf818361293a565b5190613ff482515f52600160205260405f2090565b91600183019081546020820151036140f65760408101918251905561401f6060820197885190611ce7565b9680515f81125f146140a4575061403860029151613fb7565b9401938454918183106140955760019561407661408c937f0133f662c29e67eedfc9b53c0c1f657b30ebaf9748094d09fa4659d769dd4f7895612e53565b90555b5192516040519081529081906020820190565b0390a201613fcc565b631d0a366f60e31b5f5260045ffd5b61408c915091826001965f7f0133f662c29e67eedfc9b53c0c1f657b30ebaf9748094d09fa4659d769dd4f7895136140de575b5050614079565b60026140ed910191825461146e565b90555f806140d7565b51631e32d08560e21b5f5260045260245ffd5b5050565b90919082905f5b8281106141245750505050505f90565b8281018082116112e55760011c9061413d828785611421565b35858114614167578511156141605750600181018091116112e557915b91614114565b925061415a565b50505050505050600190565b9190610e2e906111f660405193849260208401968752151560f81b6040840152604183019061119b56fe60e03461009557601f61045738819003918201601f19168301916001600160401b038311848410176100995780849260609460405283398101031261009557610047816100ad565b906040610056602083016100ad565b9101519160805260a05260c05260405161039590816100c28239608051818181610140015281816102ac015261032f015260a05181505060c051815050f35b5f80fd5b634e487b7160e01b5f52604160045260245ffd5b51906001600160a01b03821682036100955756fe60806040526004361061023e575f3560e01c8063532f08391461002b57639f149e1b0361023e57610096565b6040366003190112610092576004356001600160a01b0381168103610092576024359067ffffffffffffffff821161009257366023830112156100925781600401359167ffffffffffffffff8311610092573660248483010111610092576024019061013d565b5f80fd5b34610092575f366003190112610092573330036100b2575f805d005b61023e565b634e487b7160e01b5f52604160045260245ffd5b90601f8019910116810190811067ffffffffffffffff8211176100ed57604052565b6100b7565b67ffffffffffffffff81116100ed57601f01601f191660200190565b3d15610138573d9061011f826100f2565b9161012d60405193846100cb565b82523d5f602084013e565b606090565b337f00000000000000000000000000000000000000000000000000000000000000006001600160a01b0316036100b257825f9392849360405192839283378101848152039134905af161018e61010e565b901561019c57602081519101f35b602081519101fd5b6001600160a01b0390911681526040602082018190528101829052606091805f848401375f828201840152601f01601f1916010190565b6020818303126100925780519067ffffffffffffffff8211610092570181601f820112156100925780519061020f826100f2565b9261021d60405194856100cb565b8284526020838301011161009257815f9260208093018386015e8301015290565b5f806040516020810190639f149e1b60e01b8252600481526102616024826100cb565b519082305af161026f61010e565b50610304575f806040516020810190633d526d1760e11b82526102a88161029a3633602484016101a4565b03601f1981018352826100cb565b51907f00000000000000000000000000000000000000000000000000000000000000005afa6102d561010e565b90805b6102ea575b1561019c57602081519101f35b90806020806102fe935183010191016101db565b906102dd565b5f806040516020810190639af5325960e01b825261032a8161029a3633602484016101a4565b5190347f00000000000000000000000000000000000000000000000000000000000000005af161035861010e565b90806102d856fea26469706673582212206f9e2c868818139e4b88abf5d9da78dafbbddbcbe43a7e6fe72d2cfe13fce84864736f6c63430008220033a26469706673582212206e705f63f1ccd2a49709d5e7510c62460797e7b18d9eef524514b5f84db8962264736f6c63430008220033",
}

// RollupsABI is the input ABI used to generate the binding from.
// Deprecated: Use RollupsMetaData.ABI instead.
var RollupsABI = RollupsMetaData.ABI

// RollupsBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use RollupsMetaData.Bin instead.
var RollupsBin = RollupsMetaData.Bin

// DeployRollups deploys a new Ethereum contract, binding an instance of Rollups to it.
func DeployRollups(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *Rollups, error) {
	parsed, err := RollupsMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(RollupsBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Rollups{RollupsCaller: RollupsCaller{contract: contract}, RollupsTransactor: RollupsTransactor{contract: contract}, RollupsFilterer: RollupsFilterer{contract: contract}}, nil
}

// Rollups is an auto generated Go binding around an Ethereum contract.
type Rollups struct {
	RollupsCaller     // Read-only binding to the contract
	RollupsTransactor // Write-only binding to the contract
	RollupsFilterer   // Log filterer for contract events
}

// RollupsCaller is an auto generated read-only Go binding around an Ethereum contract.
type RollupsCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// RollupsTransactor is an auto generated write-only Go binding around an Ethereum contract.
type RollupsTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// RollupsFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type RollupsFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// RollupsSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type RollupsSession struct {
	Contract     *Rollups          // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// RollupsCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type RollupsCallerSession struct {
	Contract *RollupsCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts  // Call options to use throughout this session
}

// RollupsTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type RollupsTransactorSession struct {
	Contract     *RollupsTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts  // Transaction auth options to use throughout this session
}

// RollupsRaw is an auto generated low-level Go binding around an Ethereum contract.
type RollupsRaw struct {
	Contract *Rollups // Generic contract binding to access the raw methods on
}

// RollupsCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type RollupsCallerRaw struct {
	Contract *RollupsCaller // Generic read-only contract binding to access the raw methods on
}

// RollupsTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type RollupsTransactorRaw struct {
	Contract *RollupsTransactor // Generic write-only contract binding to access the raw methods on
}

// NewRollups creates a new instance of Rollups, bound to a specific deployed contract.
func NewRollups(address common.Address, backend bind.ContractBackend) (*Rollups, error) {
	contract, err := bindRollups(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Rollups{RollupsCaller: RollupsCaller{contract: contract}, RollupsTransactor: RollupsTransactor{contract: contract}, RollupsFilterer: RollupsFilterer{contract: contract}}, nil
}

// NewRollupsCaller creates a new read-only instance of Rollups, bound to a specific deployed contract.
func NewRollupsCaller(address common.Address, caller bind.ContractCaller) (*RollupsCaller, error) {
	contract, err := bindRollups(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &RollupsCaller{contract: contract}, nil
}

// NewRollupsTransactor creates a new write-only instance of Rollups, bound to a specific deployed contract.
func NewRollupsTransactor(address common.Address, transactor bind.ContractTransactor) (*RollupsTransactor, error) {
	contract, err := bindRollups(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &RollupsTransactor{contract: contract}, nil
}

// NewRollupsFilterer creates a new log filterer instance of Rollups, bound to a specific deployed contract.
func NewRollupsFilterer(address common.Address, filterer bind.ContractFilterer) (*RollupsFilterer, error) {
	contract, err := bindRollups(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &RollupsFilterer{contract: contract}, nil
}

// bindRollups binds a generic wrapper to an already deployed contract.
func bindRollups(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := RollupsMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Rollups *RollupsRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Rollups.Contract.RollupsCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Rollups *RollupsRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Rollups.Contract.RollupsTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Rollups *RollupsRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Rollups.Contract.RollupsTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Rollups *RollupsCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Rollups.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Rollups *RollupsTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Rollups.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Rollups *RollupsTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Rollups.Contract.contract.Transact(opts, method, params...)
}

// MAINNETROLLUPID is a free data retrieval call binding the contract method 0x199350cf.
//
// Solidity: function MAINNET_ROLLUP_ID() view returns(uint256)
func (_Rollups *RollupsCaller) MAINNETROLLUPID(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Rollups.contract.Call(opts, &out, "MAINNET_ROLLUP_ID")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MAINNETROLLUPID is a free data retrieval call binding the contract method 0x199350cf.
//
// Solidity: function MAINNET_ROLLUP_ID() view returns(uint256)
func (_Rollups *RollupsSession) MAINNETROLLUPID() (*big.Int, error) {
	return _Rollups.Contract.MAINNETROLLUPID(&_Rollups.CallOpts)
}

// MAINNETROLLUPID is a free data retrieval call binding the contract method 0x199350cf.
//
// Solidity: function MAINNET_ROLLUP_ID() view returns(uint256)
func (_Rollups *RollupsCallerSession) MAINNETROLLUPID() (*big.Int, error) {
	return _Rollups.Contract.MAINNETROLLUPID(&_Rollups.CallOpts)
}

// TransientExecutions is a free data retrieval call binding the contract method 0xd8422037.
//
// Solidity: function _transientExecutions(uint256 ) view returns(bytes32 crossChainCallHash, uint256 destinationRollupId, uint256 callCount, bytes returnData, bytes32 rollingHash)
func (_Rollups *RollupsCaller) TransientExecutions(opts *bind.CallOpts, arg0 *big.Int) (struct {
	CrossChainCallHash  [32]byte
	DestinationRollupId *big.Int
	CallCount           *big.Int
	ReturnData          []byte
	RollingHash         [32]byte
}, error) {
	var out []interface{}
	err := _Rollups.contract.Call(opts, &out, "_transientExecutions", arg0)

	outstruct := new(struct {
		CrossChainCallHash  [32]byte
		DestinationRollupId *big.Int
		CallCount           *big.Int
		ReturnData          []byte
		RollingHash         [32]byte
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.CrossChainCallHash = *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)
	outstruct.DestinationRollupId = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)
	outstruct.CallCount = *abi.ConvertType(out[2], new(*big.Int)).(**big.Int)
	outstruct.ReturnData = *abi.ConvertType(out[3], new([]byte)).(*[]byte)
	outstruct.RollingHash = *abi.ConvertType(out[4], new([32]byte)).(*[32]byte)

	return *outstruct, err

}

// TransientExecutions is a free data retrieval call binding the contract method 0xd8422037.
//
// Solidity: function _transientExecutions(uint256 ) view returns(bytes32 crossChainCallHash, uint256 destinationRollupId, uint256 callCount, bytes returnData, bytes32 rollingHash)
func (_Rollups *RollupsSession) TransientExecutions(arg0 *big.Int) (struct {
	CrossChainCallHash  [32]byte
	DestinationRollupId *big.Int
	CallCount           *big.Int
	ReturnData          []byte
	RollingHash         [32]byte
}, error) {
	return _Rollups.Contract.TransientExecutions(&_Rollups.CallOpts, arg0)
}

// TransientExecutions is a free data retrieval call binding the contract method 0xd8422037.
//
// Solidity: function _transientExecutions(uint256 ) view returns(bytes32 crossChainCallHash, uint256 destinationRollupId, uint256 callCount, bytes returnData, bytes32 rollingHash)
func (_Rollups *RollupsCallerSession) TransientExecutions(arg0 *big.Int) (struct {
	CrossChainCallHash  [32]byte
	DestinationRollupId *big.Int
	CallCount           *big.Int
	ReturnData          []byte
	RollingHash         [32]byte
}, error) {
	return _Rollups.Contract.TransientExecutions(&_Rollups.CallOpts, arg0)
}

// TransientLookupCalls is a free data retrieval call binding the contract method 0xd3f60360.
//
// Solidity: function _transientLookupCalls(uint256 ) view returns(bytes32 crossChainCallHash, uint256 destinationRollupId, bytes returnData, bool failed, uint64 callNumber, uint64 lastNestedActionConsumed, bytes32 rollingHash)
func (_Rollups *RollupsCaller) TransientLookupCalls(opts *bind.CallOpts, arg0 *big.Int) (struct {
	CrossChainCallHash       [32]byte
	DestinationRollupId      *big.Int
	ReturnData               []byte
	Failed                   bool
	CallNumber               uint64
	LastNestedActionConsumed uint64
	RollingHash              [32]byte
}, error) {
	var out []interface{}
	err := _Rollups.contract.Call(opts, &out, "_transientLookupCalls", arg0)

	outstruct := new(struct {
		CrossChainCallHash       [32]byte
		DestinationRollupId      *big.Int
		ReturnData               []byte
		Failed                   bool
		CallNumber               uint64
		LastNestedActionConsumed uint64
		RollingHash              [32]byte
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.CrossChainCallHash = *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)
	outstruct.DestinationRollupId = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)
	outstruct.ReturnData = *abi.ConvertType(out[2], new([]byte)).(*[]byte)
	outstruct.Failed = *abi.ConvertType(out[3], new(bool)).(*bool)
	outstruct.CallNumber = *abi.ConvertType(out[4], new(uint64)).(*uint64)
	outstruct.LastNestedActionConsumed = *abi.ConvertType(out[5], new(uint64)).(*uint64)
	outstruct.RollingHash = *abi.ConvertType(out[6], new([32]byte)).(*[32]byte)

	return *outstruct, err

}

// TransientLookupCalls is a free data retrieval call binding the contract method 0xd3f60360.
//
// Solidity: function _transientLookupCalls(uint256 ) view returns(bytes32 crossChainCallHash, uint256 destinationRollupId, bytes returnData, bool failed, uint64 callNumber, uint64 lastNestedActionConsumed, bytes32 rollingHash)
func (_Rollups *RollupsSession) TransientLookupCalls(arg0 *big.Int) (struct {
	CrossChainCallHash       [32]byte
	DestinationRollupId      *big.Int
	ReturnData               []byte
	Failed                   bool
	CallNumber               uint64
	LastNestedActionConsumed uint64
	RollingHash              [32]byte
}, error) {
	return _Rollups.Contract.TransientLookupCalls(&_Rollups.CallOpts, arg0)
}

// TransientLookupCalls is a free data retrieval call binding the contract method 0xd3f60360.
//
// Solidity: function _transientLookupCalls(uint256 ) view returns(bytes32 crossChainCallHash, uint256 destinationRollupId, bytes returnData, bool failed, uint64 callNumber, uint64 lastNestedActionConsumed, bytes32 rollingHash)
func (_Rollups *RollupsCallerSession) TransientLookupCalls(arg0 *big.Int) (struct {
	CrossChainCallHash       [32]byte
	DestinationRollupId      *big.Int
	ReturnData               []byte
	Failed                   bool
	CallNumber               uint64
	LastNestedActionConsumed uint64
	RollingHash              [32]byte
}, error) {
	return _Rollups.Contract.TransientLookupCalls(&_Rollups.CallOpts, arg0)
}

// AuthorizedProxies is a free data retrieval call binding the contract method 0x360d95b6.
//
// Solidity: function authorizedProxies(address proxy) view returns(address originalAddress, uint64 originalRollupId)
func (_Rollups *RollupsCaller) AuthorizedProxies(opts *bind.CallOpts, proxy common.Address) (struct {
	OriginalAddress  common.Address
	OriginalRollupId uint64
}, error) {
	var out []interface{}
	err := _Rollups.contract.Call(opts, &out, "authorizedProxies", proxy)

	outstruct := new(struct {
		OriginalAddress  common.Address
		OriginalRollupId uint64
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.OriginalAddress = *abi.ConvertType(out[0], new(common.Address)).(*common.Address)
	outstruct.OriginalRollupId = *abi.ConvertType(out[1], new(uint64)).(*uint64)

	return *outstruct, err

}

// AuthorizedProxies is a free data retrieval call binding the contract method 0x360d95b6.
//
// Solidity: function authorizedProxies(address proxy) view returns(address originalAddress, uint64 originalRollupId)
func (_Rollups *RollupsSession) AuthorizedProxies(proxy common.Address) (struct {
	OriginalAddress  common.Address
	OriginalRollupId uint64
}, error) {
	return _Rollups.Contract.AuthorizedProxies(&_Rollups.CallOpts, proxy)
}

// AuthorizedProxies is a free data retrieval call binding the contract method 0x360d95b6.
//
// Solidity: function authorizedProxies(address proxy) view returns(address originalAddress, uint64 originalRollupId)
func (_Rollups *RollupsCallerSession) AuthorizedProxies(proxy common.Address) (struct {
	OriginalAddress  common.Address
	OriginalRollupId uint64
}, error) {
	return _Rollups.Contract.AuthorizedProxies(&_Rollups.CallOpts, proxy)
}

// ComputeCrossChainCallHash is a free data retrieval call binding the contract method 0x1f314db1.
//
// Solidity: function computeCrossChainCallHash(uint256 targetRollupId, address targetAddress, uint256 value, bytes data, address sourceAddress, uint256 sourceRollupId) pure returns(bytes32)
func (_Rollups *RollupsCaller) ComputeCrossChainCallHash(opts *bind.CallOpts, targetRollupId *big.Int, targetAddress common.Address, value *big.Int, data []byte, sourceAddress common.Address, sourceRollupId *big.Int) ([32]byte, error) {
	var out []interface{}
	err := _Rollups.contract.Call(opts, &out, "computeCrossChainCallHash", targetRollupId, targetAddress, value, data, sourceAddress, sourceRollupId)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// ComputeCrossChainCallHash is a free data retrieval call binding the contract method 0x1f314db1.
//
// Solidity: function computeCrossChainCallHash(uint256 targetRollupId, address targetAddress, uint256 value, bytes data, address sourceAddress, uint256 sourceRollupId) pure returns(bytes32)
func (_Rollups *RollupsSession) ComputeCrossChainCallHash(targetRollupId *big.Int, targetAddress common.Address, value *big.Int, data []byte, sourceAddress common.Address, sourceRollupId *big.Int) ([32]byte, error) {
	return _Rollups.Contract.ComputeCrossChainCallHash(&_Rollups.CallOpts, targetRollupId, targetAddress, value, data, sourceAddress, sourceRollupId)
}

// ComputeCrossChainCallHash is a free data retrieval call binding the contract method 0x1f314db1.
//
// Solidity: function computeCrossChainCallHash(uint256 targetRollupId, address targetAddress, uint256 value, bytes data, address sourceAddress, uint256 sourceRollupId) pure returns(bytes32)
func (_Rollups *RollupsCallerSession) ComputeCrossChainCallHash(targetRollupId *big.Int, targetAddress common.Address, value *big.Int, data []byte, sourceAddress common.Address, sourceRollupId *big.Int) ([32]byte, error) {
	return _Rollups.Contract.ComputeCrossChainCallHash(&_Rollups.CallOpts, targetRollupId, targetAddress, value, data, sourceAddress, sourceRollupId)
}

// ComputeCrossChainProxyAddress is a free data retrieval call binding the contract method 0xb761ba7e.
//
// Solidity: function computeCrossChainProxyAddress(address originalAddress, uint256 originalRollupId) view returns(address)
func (_Rollups *RollupsCaller) ComputeCrossChainProxyAddress(opts *bind.CallOpts, originalAddress common.Address, originalRollupId *big.Int) (common.Address, error) {
	var out []interface{}
	err := _Rollups.contract.Call(opts, &out, "computeCrossChainProxyAddress", originalAddress, originalRollupId)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// ComputeCrossChainProxyAddress is a free data retrieval call binding the contract method 0xb761ba7e.
//
// Solidity: function computeCrossChainProxyAddress(address originalAddress, uint256 originalRollupId) view returns(address)
func (_Rollups *RollupsSession) ComputeCrossChainProxyAddress(originalAddress common.Address, originalRollupId *big.Int) (common.Address, error) {
	return _Rollups.Contract.ComputeCrossChainProxyAddress(&_Rollups.CallOpts, originalAddress, originalRollupId)
}

// ComputeCrossChainProxyAddress is a free data retrieval call binding the contract method 0xb761ba7e.
//
// Solidity: function computeCrossChainProxyAddress(address originalAddress, uint256 originalRollupId) view returns(address)
func (_Rollups *RollupsCallerSession) ComputeCrossChainProxyAddress(originalAddress common.Address, originalRollupId *big.Int) (common.Address, error) {
	return _Rollups.Contract.ComputeCrossChainProxyAddress(&_Rollups.CallOpts, originalAddress, originalRollupId)
}

// LastVerifiedBlock is a free data retrieval call binding the contract method 0x206bbf07.
//
// Solidity: function lastVerifiedBlock(uint256 _rollupId) view returns(uint256)
func (_Rollups *RollupsCaller) LastVerifiedBlock(opts *bind.CallOpts, _rollupId *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _Rollups.contract.Call(opts, &out, "lastVerifiedBlock", _rollupId)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// LastVerifiedBlock is a free data retrieval call binding the contract method 0x206bbf07.
//
// Solidity: function lastVerifiedBlock(uint256 _rollupId) view returns(uint256)
func (_Rollups *RollupsSession) LastVerifiedBlock(_rollupId *big.Int) (*big.Int, error) {
	return _Rollups.Contract.LastVerifiedBlock(&_Rollups.CallOpts, _rollupId)
}

// LastVerifiedBlock is a free data retrieval call binding the contract method 0x206bbf07.
//
// Solidity: function lastVerifiedBlock(uint256 _rollupId) view returns(uint256)
func (_Rollups *RollupsCallerSession) LastVerifiedBlock(_rollupId *big.Int) (*big.Int, error) {
	return _Rollups.Contract.LastVerifiedBlock(&_Rollups.CallOpts, _rollupId)
}

// QueueCursor is a free data retrieval call binding the contract method 0x35bcb88a.
//
// Solidity: function queueCursor(uint256 _rollupId) view returns(uint256)
func (_Rollups *RollupsCaller) QueueCursor(opts *bind.CallOpts, _rollupId *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _Rollups.contract.Call(opts, &out, "queueCursor", _rollupId)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// QueueCursor is a free data retrieval call binding the contract method 0x35bcb88a.
//
// Solidity: function queueCursor(uint256 _rollupId) view returns(uint256)
func (_Rollups *RollupsSession) QueueCursor(_rollupId *big.Int) (*big.Int, error) {
	return _Rollups.Contract.QueueCursor(&_Rollups.CallOpts, _rollupId)
}

// QueueCursor is a free data retrieval call binding the contract method 0x35bcb88a.
//
// Solidity: function queueCursor(uint256 _rollupId) view returns(uint256)
func (_Rollups *RollupsCallerSession) QueueCursor(_rollupId *big.Int) (*big.Int, error) {
	return _Rollups.Contract.QueueCursor(&_Rollups.CallOpts, _rollupId)
}

// QueueLength is a free data retrieval call binding the contract method 0x48cae3e2.
//
// Solidity: function queueLength(uint256 _rollupId) view returns(uint256)
func (_Rollups *RollupsCaller) QueueLength(opts *bind.CallOpts, _rollupId *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _Rollups.contract.Call(opts, &out, "queueLength", _rollupId)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// QueueLength is a free data retrieval call binding the contract method 0x48cae3e2.
//
// Solidity: function queueLength(uint256 _rollupId) view returns(uint256)
func (_Rollups *RollupsSession) QueueLength(_rollupId *big.Int) (*big.Int, error) {
	return _Rollups.Contract.QueueLength(&_Rollups.CallOpts, _rollupId)
}

// QueueLength is a free data retrieval call binding the contract method 0x48cae3e2.
//
// Solidity: function queueLength(uint256 _rollupId) view returns(uint256)
func (_Rollups *RollupsCallerSession) QueueLength(_rollupId *big.Int) (*big.Int, error) {
	return _Rollups.Contract.QueueLength(&_Rollups.CallOpts, _rollupId)
}

// RollupCounter is a free data retrieval call binding the contract method 0xa3271832.
//
// Solidity: function rollupCounter() view returns(uint256)
func (_Rollups *RollupsCaller) RollupCounter(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Rollups.contract.Call(opts, &out, "rollupCounter")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// RollupCounter is a free data retrieval call binding the contract method 0xa3271832.
//
// Solidity: function rollupCounter() view returns(uint256)
func (_Rollups *RollupsSession) RollupCounter() (*big.Int, error) {
	return _Rollups.Contract.RollupCounter(&_Rollups.CallOpts)
}

// RollupCounter is a free data retrieval call binding the contract method 0xa3271832.
//
// Solidity: function rollupCounter() view returns(uint256)
func (_Rollups *RollupsCallerSession) RollupCounter() (*big.Int, error) {
	return _Rollups.Contract.RollupCounter(&_Rollups.CallOpts)
}

// Rollups is a free data retrieval call binding the contract method 0xb794e5a3.
//
// Solidity: function rollups(uint256 rollupId) view returns(address rollupContract, bytes32 stateRoot, uint256 etherBalance)
func (_Rollups *RollupsCaller) Rollups(opts *bind.CallOpts, rollupId *big.Int) (struct {
	RollupContract common.Address
	StateRoot      [32]byte
	EtherBalance   *big.Int
}, error) {
	var out []interface{}
	err := _Rollups.contract.Call(opts, &out, "rollups", rollupId)

	outstruct := new(struct {
		RollupContract common.Address
		StateRoot      [32]byte
		EtherBalance   *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.RollupContract = *abi.ConvertType(out[0], new(common.Address)).(*common.Address)
	outstruct.StateRoot = *abi.ConvertType(out[1], new([32]byte)).(*[32]byte)
	outstruct.EtherBalance = *abi.ConvertType(out[2], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// Rollups is a free data retrieval call binding the contract method 0xb794e5a3.
//
// Solidity: function rollups(uint256 rollupId) view returns(address rollupContract, bytes32 stateRoot, uint256 etherBalance)
func (_Rollups *RollupsSession) Rollups(rollupId *big.Int) (struct {
	RollupContract common.Address
	StateRoot      [32]byte
	EtherBalance   *big.Int
}, error) {
	return _Rollups.Contract.Rollups(&_Rollups.CallOpts, rollupId)
}

// Rollups is a free data retrieval call binding the contract method 0xb794e5a3.
//
// Solidity: function rollups(uint256 rollupId) view returns(address rollupContract, bytes32 stateRoot, uint256 etherBalance)
func (_Rollups *RollupsCallerSession) Rollups(rollupId *big.Int) (struct {
	RollupContract common.Address
	StateRoot      [32]byte
	EtherBalance   *big.Int
}, error) {
	return _Rollups.Contract.Rollups(&_Rollups.CallOpts, rollupId)
}

// StaticCallLookup is a free data retrieval call binding the contract method 0x7aa4da2e.
//
// Solidity: function staticCallLookup(address sourceAddress, bytes callData) view returns(bytes)
func (_Rollups *RollupsCaller) StaticCallLookup(opts *bind.CallOpts, sourceAddress common.Address, callData []byte) ([]byte, error) {
	var out []interface{}
	err := _Rollups.contract.Call(opts, &out, "staticCallLookup", sourceAddress, callData)

	if err != nil {
		return *new([]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([]byte)).(*[]byte)

	return out0, err

}

// StaticCallLookup is a free data retrieval call binding the contract method 0x7aa4da2e.
//
// Solidity: function staticCallLookup(address sourceAddress, bytes callData) view returns(bytes)
func (_Rollups *RollupsSession) StaticCallLookup(sourceAddress common.Address, callData []byte) ([]byte, error) {
	return _Rollups.Contract.StaticCallLookup(&_Rollups.CallOpts, sourceAddress, callData)
}

// StaticCallLookup is a free data retrieval call binding the contract method 0x7aa4da2e.
//
// Solidity: function staticCallLookup(address sourceAddress, bytes callData) view returns(bytes)
func (_Rollups *RollupsCallerSession) StaticCallLookup(sourceAddress common.Address, callData []byte) ([]byte, error) {
	return _Rollups.Contract.StaticCallLookup(&_Rollups.CallOpts, sourceAddress, callData)
}

// AttemptApplyImmediate is a paid mutator transaction binding the contract method 0xd4ccaf87.
//
// Solidity: function attemptApplyImmediate(uint256 transientIdx) returns()
func (_Rollups *RollupsTransactor) AttemptApplyImmediate(opts *bind.TransactOpts, transientIdx *big.Int) (*types.Transaction, error) {
	return _Rollups.contract.Transact(opts, "attemptApplyImmediate", transientIdx)
}

// AttemptApplyImmediate is a paid mutator transaction binding the contract method 0xd4ccaf87.
//
// Solidity: function attemptApplyImmediate(uint256 transientIdx) returns()
func (_Rollups *RollupsSession) AttemptApplyImmediate(transientIdx *big.Int) (*types.Transaction, error) {
	return _Rollups.Contract.AttemptApplyImmediate(&_Rollups.TransactOpts, transientIdx)
}

// AttemptApplyImmediate is a paid mutator transaction binding the contract method 0xd4ccaf87.
//
// Solidity: function attemptApplyImmediate(uint256 transientIdx) returns()
func (_Rollups *RollupsTransactorSession) AttemptApplyImmediate(transientIdx *big.Int) (*types.Transaction, error) {
	return _Rollups.Contract.AttemptApplyImmediate(&_Rollups.TransactOpts, transientIdx)
}

// CreateCrossChainProxy is a paid mutator transaction binding the contract method 0x2dd72120.
//
// Solidity: function createCrossChainProxy(address originalAddress, uint256 originalRollupId) returns(address proxy)
func (_Rollups *RollupsTransactor) CreateCrossChainProxy(opts *bind.TransactOpts, originalAddress common.Address, originalRollupId *big.Int) (*types.Transaction, error) {
	return _Rollups.contract.Transact(opts, "createCrossChainProxy", originalAddress, originalRollupId)
}

// CreateCrossChainProxy is a paid mutator transaction binding the contract method 0x2dd72120.
//
// Solidity: function createCrossChainProxy(address originalAddress, uint256 originalRollupId) returns(address proxy)
func (_Rollups *RollupsSession) CreateCrossChainProxy(originalAddress common.Address, originalRollupId *big.Int) (*types.Transaction, error) {
	return _Rollups.Contract.CreateCrossChainProxy(&_Rollups.TransactOpts, originalAddress, originalRollupId)
}

// CreateCrossChainProxy is a paid mutator transaction binding the contract method 0x2dd72120.
//
// Solidity: function createCrossChainProxy(address originalAddress, uint256 originalRollupId) returns(address proxy)
func (_Rollups *RollupsTransactorSession) CreateCrossChainProxy(originalAddress common.Address, originalRollupId *big.Int) (*types.Transaction, error) {
	return _Rollups.Contract.CreateCrossChainProxy(&_Rollups.TransactOpts, originalAddress, originalRollupId)
}

// CreateRollup is a paid mutator transaction binding the contract method 0xcc21dd3c.
//
// Solidity: function createRollup(address rollupContract, bytes32 initialState) returns(uint256 rollupId)
func (_Rollups *RollupsTransactor) CreateRollup(opts *bind.TransactOpts, rollupContract common.Address, initialState [32]byte) (*types.Transaction, error) {
	return _Rollups.contract.Transact(opts, "createRollup", rollupContract, initialState)
}

// CreateRollup is a paid mutator transaction binding the contract method 0xcc21dd3c.
//
// Solidity: function createRollup(address rollupContract, bytes32 initialState) returns(uint256 rollupId)
func (_Rollups *RollupsSession) CreateRollup(rollupContract common.Address, initialState [32]byte) (*types.Transaction, error) {
	return _Rollups.Contract.CreateRollup(&_Rollups.TransactOpts, rollupContract, initialState)
}

// CreateRollup is a paid mutator transaction binding the contract method 0xcc21dd3c.
//
// Solidity: function createRollup(address rollupContract, bytes32 initialState) returns(uint256 rollupId)
func (_Rollups *RollupsTransactorSession) CreateRollup(rollupContract common.Address, initialState [32]byte) (*types.Transaction, error) {
	return _Rollups.Contract.CreateRollup(&_Rollups.TransactOpts, rollupContract, initialState)
}

// ExecuteCrossChainCall is a paid mutator transaction binding the contract method 0x9af53259.
//
// Solidity: function executeCrossChainCall(address sourceAddress, bytes callData) payable returns(bytes result)
func (_Rollups *RollupsTransactor) ExecuteCrossChainCall(opts *bind.TransactOpts, sourceAddress common.Address, callData []byte) (*types.Transaction, error) {
	return _Rollups.contract.Transact(opts, "executeCrossChainCall", sourceAddress, callData)
}

// ExecuteCrossChainCall is a paid mutator transaction binding the contract method 0x9af53259.
//
// Solidity: function executeCrossChainCall(address sourceAddress, bytes callData) payable returns(bytes result)
func (_Rollups *RollupsSession) ExecuteCrossChainCall(sourceAddress common.Address, callData []byte) (*types.Transaction, error) {
	return _Rollups.Contract.ExecuteCrossChainCall(&_Rollups.TransactOpts, sourceAddress, callData)
}

// ExecuteCrossChainCall is a paid mutator transaction binding the contract method 0x9af53259.
//
// Solidity: function executeCrossChainCall(address sourceAddress, bytes callData) payable returns(bytes result)
func (_Rollups *RollupsTransactorSession) ExecuteCrossChainCall(sourceAddress common.Address, callData []byte) (*types.Transaction, error) {
	return _Rollups.Contract.ExecuteCrossChainCall(&_Rollups.TransactOpts, sourceAddress, callData)
}

// ExecuteInContextAndRevert is a paid mutator transaction binding the contract method 0x37b99a56.
//
// Solidity: function executeInContextAndRevert(uint256 callCount) returns()
func (_Rollups *RollupsTransactor) ExecuteInContextAndRevert(opts *bind.TransactOpts, callCount *big.Int) (*types.Transaction, error) {
	return _Rollups.contract.Transact(opts, "executeInContextAndRevert", callCount)
}

// ExecuteInContextAndRevert is a paid mutator transaction binding the contract method 0x37b99a56.
//
// Solidity: function executeInContextAndRevert(uint256 callCount) returns()
func (_Rollups *RollupsSession) ExecuteInContextAndRevert(callCount *big.Int) (*types.Transaction, error) {
	return _Rollups.Contract.ExecuteInContextAndRevert(&_Rollups.TransactOpts, callCount)
}

// ExecuteInContextAndRevert is a paid mutator transaction binding the contract method 0x37b99a56.
//
// Solidity: function executeInContextAndRevert(uint256 callCount) returns()
func (_Rollups *RollupsTransactorSession) ExecuteInContextAndRevert(callCount *big.Int) (*types.Transaction, error) {
	return _Rollups.Contract.ExecuteInContextAndRevert(&_Rollups.TransactOpts, callCount)
}

// ExecuteL2TX is a paid mutator transaction binding the contract method 0xccdcf581.
//
// Solidity: function executeL2TX(uint256 rollupId) returns(bytes result)
func (_Rollups *RollupsTransactor) ExecuteL2TX(opts *bind.TransactOpts, rollupId *big.Int) (*types.Transaction, error) {
	return _Rollups.contract.Transact(opts, "executeL2TX", rollupId)
}

// ExecuteL2TX is a paid mutator transaction binding the contract method 0xccdcf581.
//
// Solidity: function executeL2TX(uint256 rollupId) returns(bytes result)
func (_Rollups *RollupsSession) ExecuteL2TX(rollupId *big.Int) (*types.Transaction, error) {
	return _Rollups.Contract.ExecuteL2TX(&_Rollups.TransactOpts, rollupId)
}

// ExecuteL2TX is a paid mutator transaction binding the contract method 0xccdcf581.
//
// Solidity: function executeL2TX(uint256 rollupId) returns(bytes result)
func (_Rollups *RollupsTransactorSession) ExecuteL2TX(rollupId *big.Int) (*types.Transaction, error) {
	return _Rollups.Contract.ExecuteL2TX(&_Rollups.TransactOpts, rollupId)
}

// PostBatch is a paid mutator transaction binding the contract method 0xf41558a5.
//
// Solidity: function postBatch((address[],uint256[],((uint256,bytes32,bytes32,int256)[],bytes32,uint256,(address,uint256,bytes,address,uint256,uint256)[],(bytes32,uint256,bytes)[],uint256,bytes,bytes32)[],(bytes32,uint256,bytes,bool,uint64,uint64,(address,uint256,bytes,address,uint256,uint256)[],bytes32)[],uint256,uint256,uint256[],bytes,bytes[],bytes32)[] batches) returns()
func (_Rollups *RollupsTransactor) PostBatch(opts *bind.TransactOpts, batches []ProofSystemBatch) (*types.Transaction, error) {
	return _Rollups.contract.Transact(opts, "postBatch", batches)
}

// PostBatch is a paid mutator transaction binding the contract method 0xf41558a5.
//
// Solidity: function postBatch((address[],uint256[],((uint256,bytes32,bytes32,int256)[],bytes32,uint256,(address,uint256,bytes,address,uint256,uint256)[],(bytes32,uint256,bytes)[],uint256,bytes,bytes32)[],(bytes32,uint256,bytes,bool,uint64,uint64,(address,uint256,bytes,address,uint256,uint256)[],bytes32)[],uint256,uint256,uint256[],bytes,bytes[],bytes32)[] batches) returns()
func (_Rollups *RollupsSession) PostBatch(batches []ProofSystemBatch) (*types.Transaction, error) {
	return _Rollups.Contract.PostBatch(&_Rollups.TransactOpts, batches)
}

// PostBatch is a paid mutator transaction binding the contract method 0xf41558a5.
//
// Solidity: function postBatch((address[],uint256[],((uint256,bytes32,bytes32,int256)[],bytes32,uint256,(address,uint256,bytes,address,uint256,uint256)[],(bytes32,uint256,bytes)[],uint256,bytes,bytes32)[],(bytes32,uint256,bytes,bool,uint64,uint64,(address,uint256,bytes,address,uint256,uint256)[],bytes32)[],uint256,uint256,uint256[],bytes,bytes[],bytes32)[] batches) returns()
func (_Rollups *RollupsTransactorSession) PostBatch(batches []ProofSystemBatch) (*types.Transaction, error) {
	return _Rollups.Contract.PostBatch(&_Rollups.TransactOpts, batches)
}

// SetRollupContract is a paid mutator transaction binding the contract method 0xff2f742f.
//
// Solidity: function setRollupContract(uint256 rollupId, address newContract) returns()
func (_Rollups *RollupsTransactor) SetRollupContract(opts *bind.TransactOpts, rollupId *big.Int, newContract common.Address) (*types.Transaction, error) {
	return _Rollups.contract.Transact(opts, "setRollupContract", rollupId, newContract)
}

// SetRollupContract is a paid mutator transaction binding the contract method 0xff2f742f.
//
// Solidity: function setRollupContract(uint256 rollupId, address newContract) returns()
func (_Rollups *RollupsSession) SetRollupContract(rollupId *big.Int, newContract common.Address) (*types.Transaction, error) {
	return _Rollups.Contract.SetRollupContract(&_Rollups.TransactOpts, rollupId, newContract)
}

// SetRollupContract is a paid mutator transaction binding the contract method 0xff2f742f.
//
// Solidity: function setRollupContract(uint256 rollupId, address newContract) returns()
func (_Rollups *RollupsTransactorSession) SetRollupContract(rollupId *big.Int, newContract common.Address) (*types.Transaction, error) {
	return _Rollups.Contract.SetRollupContract(&_Rollups.TransactOpts, rollupId, newContract)
}

// SetStateRoot is a paid mutator transaction binding the contract method 0x2bdd1f16.
//
// Solidity: function setStateRoot(uint256 rollupId, bytes32 newStateRoot) returns()
func (_Rollups *RollupsTransactor) SetStateRoot(opts *bind.TransactOpts, rollupId *big.Int, newStateRoot [32]byte) (*types.Transaction, error) {
	return _Rollups.contract.Transact(opts, "setStateRoot", rollupId, newStateRoot)
}

// SetStateRoot is a paid mutator transaction binding the contract method 0x2bdd1f16.
//
// Solidity: function setStateRoot(uint256 rollupId, bytes32 newStateRoot) returns()
func (_Rollups *RollupsSession) SetStateRoot(rollupId *big.Int, newStateRoot [32]byte) (*types.Transaction, error) {
	return _Rollups.Contract.SetStateRoot(&_Rollups.TransactOpts, rollupId, newStateRoot)
}

// SetStateRoot is a paid mutator transaction binding the contract method 0x2bdd1f16.
//
// Solidity: function setStateRoot(uint256 rollupId, bytes32 newStateRoot) returns()
func (_Rollups *RollupsTransactorSession) SetStateRoot(rollupId *big.Int, newStateRoot [32]byte) (*types.Transaction, error) {
	return _Rollups.Contract.SetStateRoot(&_Rollups.TransactOpts, rollupId, newStateRoot)
}

// RollupsBatchPostedIterator is returned from FilterBatchPosted and is used to iterate over the raw logs and unpacked data for BatchPosted events raised by the Rollups contract.
type RollupsBatchPostedIterator struct {
	Event *RollupsBatchPosted // Event containing the contract specifics and raw log

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
func (it *RollupsBatchPostedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(RollupsBatchPosted)
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
		it.Event = new(RollupsBatchPosted)
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
func (it *RollupsBatchPostedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *RollupsBatchPostedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// RollupsBatchPosted represents a BatchPosted event raised by the Rollups contract.
type RollupsBatchPosted struct {
	SubBatchCount *big.Int
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterBatchPosted is a free log retrieval operation binding the contract event 0xd6f8d71ce42a799b91f399271f4b0e91f85eb87fac7bb2cedd4b3a52fad36182.
//
// Solidity: event BatchPosted(uint256 indexed subBatchCount)
func (_Rollups *RollupsFilterer) FilterBatchPosted(opts *bind.FilterOpts, subBatchCount []*big.Int) (*RollupsBatchPostedIterator, error) {

	var subBatchCountRule []interface{}
	for _, subBatchCountItem := range subBatchCount {
		subBatchCountRule = append(subBatchCountRule, subBatchCountItem)
	}

	logs, sub, err := _Rollups.contract.FilterLogs(opts, "BatchPosted", subBatchCountRule)
	if err != nil {
		return nil, err
	}
	return &RollupsBatchPostedIterator{contract: _Rollups.contract, event: "BatchPosted", logs: logs, sub: sub}, nil
}

// WatchBatchPosted is a free log subscription operation binding the contract event 0xd6f8d71ce42a799b91f399271f4b0e91f85eb87fac7bb2cedd4b3a52fad36182.
//
// Solidity: event BatchPosted(uint256 indexed subBatchCount)
func (_Rollups *RollupsFilterer) WatchBatchPosted(opts *bind.WatchOpts, sink chan<- *RollupsBatchPosted, subBatchCount []*big.Int) (event.Subscription, error) {

	var subBatchCountRule []interface{}
	for _, subBatchCountItem := range subBatchCount {
		subBatchCountRule = append(subBatchCountRule, subBatchCountItem)
	}

	logs, sub, err := _Rollups.contract.WatchLogs(opts, "BatchPosted", subBatchCountRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(RollupsBatchPosted)
				if err := _Rollups.contract.UnpackLog(event, "BatchPosted", log); err != nil {
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

// ParseBatchPosted is a log parse operation binding the contract event 0xd6f8d71ce42a799b91f399271f4b0e91f85eb87fac7bb2cedd4b3a52fad36182.
//
// Solidity: event BatchPosted(uint256 indexed subBatchCount)
func (_Rollups *RollupsFilterer) ParseBatchPosted(log types.Log) (*RollupsBatchPosted, error) {
	event := new(RollupsBatchPosted)
	if err := _Rollups.contract.UnpackLog(event, "BatchPosted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// RollupsCallResultIterator is returned from FilterCallResult and is used to iterate over the raw logs and unpacked data for CallResult events raised by the Rollups contract.
type RollupsCallResultIterator struct {
	Event *RollupsCallResult // Event containing the contract specifics and raw log

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
func (it *RollupsCallResultIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(RollupsCallResult)
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
		it.Event = new(RollupsCallResult)
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
func (it *RollupsCallResultIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *RollupsCallResultIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// RollupsCallResult represents a CallResult event raised by the Rollups contract.
type RollupsCallResult struct {
	EntryIndex *big.Int
	CallNumber *big.Int
	Success    bool
	ReturnData []byte
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterCallResult is a free log retrieval operation binding the contract event 0x6ab31fd4e1fedcfcbb1a56d82ca9c2aa41d65466dff790e7287a415348b75765.
//
// Solidity: event CallResult(uint256 indexed entryIndex, uint256 indexed callNumber, bool success, bytes returnData)
func (_Rollups *RollupsFilterer) FilterCallResult(opts *bind.FilterOpts, entryIndex []*big.Int, callNumber []*big.Int) (*RollupsCallResultIterator, error) {

	var entryIndexRule []interface{}
	for _, entryIndexItem := range entryIndex {
		entryIndexRule = append(entryIndexRule, entryIndexItem)
	}
	var callNumberRule []interface{}
	for _, callNumberItem := range callNumber {
		callNumberRule = append(callNumberRule, callNumberItem)
	}

	logs, sub, err := _Rollups.contract.FilterLogs(opts, "CallResult", entryIndexRule, callNumberRule)
	if err != nil {
		return nil, err
	}
	return &RollupsCallResultIterator{contract: _Rollups.contract, event: "CallResult", logs: logs, sub: sub}, nil
}

// WatchCallResult is a free log subscription operation binding the contract event 0x6ab31fd4e1fedcfcbb1a56d82ca9c2aa41d65466dff790e7287a415348b75765.
//
// Solidity: event CallResult(uint256 indexed entryIndex, uint256 indexed callNumber, bool success, bytes returnData)
func (_Rollups *RollupsFilterer) WatchCallResult(opts *bind.WatchOpts, sink chan<- *RollupsCallResult, entryIndex []*big.Int, callNumber []*big.Int) (event.Subscription, error) {

	var entryIndexRule []interface{}
	for _, entryIndexItem := range entryIndex {
		entryIndexRule = append(entryIndexRule, entryIndexItem)
	}
	var callNumberRule []interface{}
	for _, callNumberItem := range callNumber {
		callNumberRule = append(callNumberRule, callNumberItem)
	}

	logs, sub, err := _Rollups.contract.WatchLogs(opts, "CallResult", entryIndexRule, callNumberRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(RollupsCallResult)
				if err := _Rollups.contract.UnpackLog(event, "CallResult", log); err != nil {
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

// ParseCallResult is a log parse operation binding the contract event 0x6ab31fd4e1fedcfcbb1a56d82ca9c2aa41d65466dff790e7287a415348b75765.
//
// Solidity: event CallResult(uint256 indexed entryIndex, uint256 indexed callNumber, bool success, bytes returnData)
func (_Rollups *RollupsFilterer) ParseCallResult(log types.Log) (*RollupsCallResult, error) {
	event := new(RollupsCallResult)
	if err := _Rollups.contract.UnpackLog(event, "CallResult", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// RollupsCrossChainCallExecutedIterator is returned from FilterCrossChainCallExecuted and is used to iterate over the raw logs and unpacked data for CrossChainCallExecuted events raised by the Rollups contract.
type RollupsCrossChainCallExecutedIterator struct {
	Event *RollupsCrossChainCallExecuted // Event containing the contract specifics and raw log

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
func (it *RollupsCrossChainCallExecutedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(RollupsCrossChainCallExecuted)
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
		it.Event = new(RollupsCrossChainCallExecuted)
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
func (it *RollupsCrossChainCallExecutedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *RollupsCrossChainCallExecutedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// RollupsCrossChainCallExecuted represents a CrossChainCallExecuted event raised by the Rollups contract.
type RollupsCrossChainCallExecuted struct {
	CrossChainCallHash [32]byte
	Proxy              common.Address
	SourceAddress      common.Address
	CallData           []byte
	Value              *big.Int
	Raw                types.Log // Blockchain specific contextual infos
}

// FilterCrossChainCallExecuted is a free log retrieval operation binding the contract event 0xad4275800b3b34f59dea39a4bd58ae171a852e3555a2ed1a6f9bad83eb6fb598.
//
// Solidity: event CrossChainCallExecuted(bytes32 indexed crossChainCallHash, address indexed proxy, address sourceAddress, bytes callData, uint256 value)
func (_Rollups *RollupsFilterer) FilterCrossChainCallExecuted(opts *bind.FilterOpts, crossChainCallHash [][32]byte, proxy []common.Address) (*RollupsCrossChainCallExecutedIterator, error) {

	var crossChainCallHashRule []interface{}
	for _, crossChainCallHashItem := range crossChainCallHash {
		crossChainCallHashRule = append(crossChainCallHashRule, crossChainCallHashItem)
	}
	var proxyRule []interface{}
	for _, proxyItem := range proxy {
		proxyRule = append(proxyRule, proxyItem)
	}

	logs, sub, err := _Rollups.contract.FilterLogs(opts, "CrossChainCallExecuted", crossChainCallHashRule, proxyRule)
	if err != nil {
		return nil, err
	}
	return &RollupsCrossChainCallExecutedIterator{contract: _Rollups.contract, event: "CrossChainCallExecuted", logs: logs, sub: sub}, nil
}

// WatchCrossChainCallExecuted is a free log subscription operation binding the contract event 0xad4275800b3b34f59dea39a4bd58ae171a852e3555a2ed1a6f9bad83eb6fb598.
//
// Solidity: event CrossChainCallExecuted(bytes32 indexed crossChainCallHash, address indexed proxy, address sourceAddress, bytes callData, uint256 value)
func (_Rollups *RollupsFilterer) WatchCrossChainCallExecuted(opts *bind.WatchOpts, sink chan<- *RollupsCrossChainCallExecuted, crossChainCallHash [][32]byte, proxy []common.Address) (event.Subscription, error) {

	var crossChainCallHashRule []interface{}
	for _, crossChainCallHashItem := range crossChainCallHash {
		crossChainCallHashRule = append(crossChainCallHashRule, crossChainCallHashItem)
	}
	var proxyRule []interface{}
	for _, proxyItem := range proxy {
		proxyRule = append(proxyRule, proxyItem)
	}

	logs, sub, err := _Rollups.contract.WatchLogs(opts, "CrossChainCallExecuted", crossChainCallHashRule, proxyRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(RollupsCrossChainCallExecuted)
				if err := _Rollups.contract.UnpackLog(event, "CrossChainCallExecuted", log); err != nil {
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

// ParseCrossChainCallExecuted is a log parse operation binding the contract event 0xad4275800b3b34f59dea39a4bd58ae171a852e3555a2ed1a6f9bad83eb6fb598.
//
// Solidity: event CrossChainCallExecuted(bytes32 indexed crossChainCallHash, address indexed proxy, address sourceAddress, bytes callData, uint256 value)
func (_Rollups *RollupsFilterer) ParseCrossChainCallExecuted(log types.Log) (*RollupsCrossChainCallExecuted, error) {
	event := new(RollupsCrossChainCallExecuted)
	if err := _Rollups.contract.UnpackLog(event, "CrossChainCallExecuted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// RollupsCrossChainProxyCreatedIterator is returned from FilterCrossChainProxyCreated and is used to iterate over the raw logs and unpacked data for CrossChainProxyCreated events raised by the Rollups contract.
type RollupsCrossChainProxyCreatedIterator struct {
	Event *RollupsCrossChainProxyCreated // Event containing the contract specifics and raw log

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
func (it *RollupsCrossChainProxyCreatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(RollupsCrossChainProxyCreated)
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
		it.Event = new(RollupsCrossChainProxyCreated)
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
func (it *RollupsCrossChainProxyCreatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *RollupsCrossChainProxyCreatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// RollupsCrossChainProxyCreated represents a CrossChainProxyCreated event raised by the Rollups contract.
type RollupsCrossChainProxyCreated struct {
	Proxy            common.Address
	OriginalAddress  common.Address
	OriginalRollupId *big.Int
	Raw              types.Log // Blockchain specific contextual infos
}

// FilterCrossChainProxyCreated is a free log retrieval operation binding the contract event 0x318838c0aec28d3eb1d3a48be7a8b2576548cde8922741021503e81b01a4c447.
//
// Solidity: event CrossChainProxyCreated(address indexed proxy, address indexed originalAddress, uint256 indexed originalRollupId)
func (_Rollups *RollupsFilterer) FilterCrossChainProxyCreated(opts *bind.FilterOpts, proxy []common.Address, originalAddress []common.Address, originalRollupId []*big.Int) (*RollupsCrossChainProxyCreatedIterator, error) {

	var proxyRule []interface{}
	for _, proxyItem := range proxy {
		proxyRule = append(proxyRule, proxyItem)
	}
	var originalAddressRule []interface{}
	for _, originalAddressItem := range originalAddress {
		originalAddressRule = append(originalAddressRule, originalAddressItem)
	}
	var originalRollupIdRule []interface{}
	for _, originalRollupIdItem := range originalRollupId {
		originalRollupIdRule = append(originalRollupIdRule, originalRollupIdItem)
	}

	logs, sub, err := _Rollups.contract.FilterLogs(opts, "CrossChainProxyCreated", proxyRule, originalAddressRule, originalRollupIdRule)
	if err != nil {
		return nil, err
	}
	return &RollupsCrossChainProxyCreatedIterator{contract: _Rollups.contract, event: "CrossChainProxyCreated", logs: logs, sub: sub}, nil
}

// WatchCrossChainProxyCreated is a free log subscription operation binding the contract event 0x318838c0aec28d3eb1d3a48be7a8b2576548cde8922741021503e81b01a4c447.
//
// Solidity: event CrossChainProxyCreated(address indexed proxy, address indexed originalAddress, uint256 indexed originalRollupId)
func (_Rollups *RollupsFilterer) WatchCrossChainProxyCreated(opts *bind.WatchOpts, sink chan<- *RollupsCrossChainProxyCreated, proxy []common.Address, originalAddress []common.Address, originalRollupId []*big.Int) (event.Subscription, error) {

	var proxyRule []interface{}
	for _, proxyItem := range proxy {
		proxyRule = append(proxyRule, proxyItem)
	}
	var originalAddressRule []interface{}
	for _, originalAddressItem := range originalAddress {
		originalAddressRule = append(originalAddressRule, originalAddressItem)
	}
	var originalRollupIdRule []interface{}
	for _, originalRollupIdItem := range originalRollupId {
		originalRollupIdRule = append(originalRollupIdRule, originalRollupIdItem)
	}

	logs, sub, err := _Rollups.contract.WatchLogs(opts, "CrossChainProxyCreated", proxyRule, originalAddressRule, originalRollupIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(RollupsCrossChainProxyCreated)
				if err := _Rollups.contract.UnpackLog(event, "CrossChainProxyCreated", log); err != nil {
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

// ParseCrossChainProxyCreated is a log parse operation binding the contract event 0x318838c0aec28d3eb1d3a48be7a8b2576548cde8922741021503e81b01a4c447.
//
// Solidity: event CrossChainProxyCreated(address indexed proxy, address indexed originalAddress, uint256 indexed originalRollupId)
func (_Rollups *RollupsFilterer) ParseCrossChainProxyCreated(log types.Log) (*RollupsCrossChainProxyCreated, error) {
	event := new(RollupsCrossChainProxyCreated)
	if err := _Rollups.contract.UnpackLog(event, "CrossChainProxyCreated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// RollupsEntryExecutedIterator is returned from FilterEntryExecuted and is used to iterate over the raw logs and unpacked data for EntryExecuted events raised by the Rollups contract.
type RollupsEntryExecutedIterator struct {
	Event *RollupsEntryExecuted // Event containing the contract specifics and raw log

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
func (it *RollupsEntryExecutedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(RollupsEntryExecuted)
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
		it.Event = new(RollupsEntryExecuted)
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
func (it *RollupsEntryExecutedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *RollupsEntryExecutedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// RollupsEntryExecuted represents a EntryExecuted event raised by the Rollups contract.
type RollupsEntryExecuted struct {
	EntryIndex            *big.Int
	RollingHash           [32]byte
	CallsProcessed        *big.Int
	NestedActionsConsumed *big.Int
	Raw                   types.Log // Blockchain specific contextual infos
}

// FilterEntryExecuted is a free log retrieval operation binding the contract event 0x99403777f3e390813f268a6307c0cb1b1ae83271187473a2c07e34f81e4a2936.
//
// Solidity: event EntryExecuted(uint256 indexed entryIndex, bytes32 rollingHash, uint256 callsProcessed, uint256 nestedActionsConsumed)
func (_Rollups *RollupsFilterer) FilterEntryExecuted(opts *bind.FilterOpts, entryIndex []*big.Int) (*RollupsEntryExecutedIterator, error) {

	var entryIndexRule []interface{}
	for _, entryIndexItem := range entryIndex {
		entryIndexRule = append(entryIndexRule, entryIndexItem)
	}

	logs, sub, err := _Rollups.contract.FilterLogs(opts, "EntryExecuted", entryIndexRule)
	if err != nil {
		return nil, err
	}
	return &RollupsEntryExecutedIterator{contract: _Rollups.contract, event: "EntryExecuted", logs: logs, sub: sub}, nil
}

// WatchEntryExecuted is a free log subscription operation binding the contract event 0x99403777f3e390813f268a6307c0cb1b1ae83271187473a2c07e34f81e4a2936.
//
// Solidity: event EntryExecuted(uint256 indexed entryIndex, bytes32 rollingHash, uint256 callsProcessed, uint256 nestedActionsConsumed)
func (_Rollups *RollupsFilterer) WatchEntryExecuted(opts *bind.WatchOpts, sink chan<- *RollupsEntryExecuted, entryIndex []*big.Int) (event.Subscription, error) {

	var entryIndexRule []interface{}
	for _, entryIndexItem := range entryIndex {
		entryIndexRule = append(entryIndexRule, entryIndexItem)
	}

	logs, sub, err := _Rollups.contract.WatchLogs(opts, "EntryExecuted", entryIndexRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(RollupsEntryExecuted)
				if err := _Rollups.contract.UnpackLog(event, "EntryExecuted", log); err != nil {
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

// ParseEntryExecuted is a log parse operation binding the contract event 0x99403777f3e390813f268a6307c0cb1b1ae83271187473a2c07e34f81e4a2936.
//
// Solidity: event EntryExecuted(uint256 indexed entryIndex, bytes32 rollingHash, uint256 callsProcessed, uint256 nestedActionsConsumed)
func (_Rollups *RollupsFilterer) ParseEntryExecuted(log types.Log) (*RollupsEntryExecuted, error) {
	event := new(RollupsEntryExecuted)
	if err := _Rollups.contract.UnpackLog(event, "EntryExecuted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// RollupsExecutionConsumedIterator is returned from FilterExecutionConsumed and is used to iterate over the raw logs and unpacked data for ExecutionConsumed events raised by the Rollups contract.
type RollupsExecutionConsumedIterator struct {
	Event *RollupsExecutionConsumed // Event containing the contract specifics and raw log

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
func (it *RollupsExecutionConsumedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(RollupsExecutionConsumed)
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
		it.Event = new(RollupsExecutionConsumed)
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
func (it *RollupsExecutionConsumedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *RollupsExecutionConsumedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// RollupsExecutionConsumed represents a ExecutionConsumed event raised by the Rollups contract.
type RollupsExecutionConsumed struct {
	CrossChainCallHash [32]byte
	RollupId           *big.Int
	Cursor             *big.Int
	Raw                types.Log // Blockchain specific contextual infos
}

// FilterExecutionConsumed is a free log retrieval operation binding the contract event 0x207e371295f3ff0efe443424ce128c96f8ecea6a25636fc38ad7c222c446699c.
//
// Solidity: event ExecutionConsumed(bytes32 indexed crossChainCallHash, uint256 indexed rollupId, uint256 indexed cursor)
func (_Rollups *RollupsFilterer) FilterExecutionConsumed(opts *bind.FilterOpts, crossChainCallHash [][32]byte, rollupId []*big.Int, cursor []*big.Int) (*RollupsExecutionConsumedIterator, error) {

	var crossChainCallHashRule []interface{}
	for _, crossChainCallHashItem := range crossChainCallHash {
		crossChainCallHashRule = append(crossChainCallHashRule, crossChainCallHashItem)
	}
	var rollupIdRule []interface{}
	for _, rollupIdItem := range rollupId {
		rollupIdRule = append(rollupIdRule, rollupIdItem)
	}
	var cursorRule []interface{}
	for _, cursorItem := range cursor {
		cursorRule = append(cursorRule, cursorItem)
	}

	logs, sub, err := _Rollups.contract.FilterLogs(opts, "ExecutionConsumed", crossChainCallHashRule, rollupIdRule, cursorRule)
	if err != nil {
		return nil, err
	}
	return &RollupsExecutionConsumedIterator{contract: _Rollups.contract, event: "ExecutionConsumed", logs: logs, sub: sub}, nil
}

// WatchExecutionConsumed is a free log subscription operation binding the contract event 0x207e371295f3ff0efe443424ce128c96f8ecea6a25636fc38ad7c222c446699c.
//
// Solidity: event ExecutionConsumed(bytes32 indexed crossChainCallHash, uint256 indexed rollupId, uint256 indexed cursor)
func (_Rollups *RollupsFilterer) WatchExecutionConsumed(opts *bind.WatchOpts, sink chan<- *RollupsExecutionConsumed, crossChainCallHash [][32]byte, rollupId []*big.Int, cursor []*big.Int) (event.Subscription, error) {

	var crossChainCallHashRule []interface{}
	for _, crossChainCallHashItem := range crossChainCallHash {
		crossChainCallHashRule = append(crossChainCallHashRule, crossChainCallHashItem)
	}
	var rollupIdRule []interface{}
	for _, rollupIdItem := range rollupId {
		rollupIdRule = append(rollupIdRule, rollupIdItem)
	}
	var cursorRule []interface{}
	for _, cursorItem := range cursor {
		cursorRule = append(cursorRule, cursorItem)
	}

	logs, sub, err := _Rollups.contract.WatchLogs(opts, "ExecutionConsumed", crossChainCallHashRule, rollupIdRule, cursorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(RollupsExecutionConsumed)
				if err := _Rollups.contract.UnpackLog(event, "ExecutionConsumed", log); err != nil {
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

// ParseExecutionConsumed is a log parse operation binding the contract event 0x207e371295f3ff0efe443424ce128c96f8ecea6a25636fc38ad7c222c446699c.
//
// Solidity: event ExecutionConsumed(bytes32 indexed crossChainCallHash, uint256 indexed rollupId, uint256 indexed cursor)
func (_Rollups *RollupsFilterer) ParseExecutionConsumed(log types.Log) (*RollupsExecutionConsumed, error) {
	event := new(RollupsExecutionConsumed)
	if err := _Rollups.contract.UnpackLog(event, "ExecutionConsumed", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// RollupsImmediateEntrySkippedIterator is returned from FilterImmediateEntrySkipped and is used to iterate over the raw logs and unpacked data for ImmediateEntrySkipped events raised by the Rollups contract.
type RollupsImmediateEntrySkippedIterator struct {
	Event *RollupsImmediateEntrySkipped // Event containing the contract specifics and raw log

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
func (it *RollupsImmediateEntrySkippedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(RollupsImmediateEntrySkipped)
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
		it.Event = new(RollupsImmediateEntrySkipped)
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
func (it *RollupsImmediateEntrySkippedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *RollupsImmediateEntrySkippedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// RollupsImmediateEntrySkipped represents a ImmediateEntrySkipped event raised by the Rollups contract.
type RollupsImmediateEntrySkipped struct {
	TransientIdx *big.Int
	RevertData   []byte
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterImmediateEntrySkipped is a free log retrieval operation binding the contract event 0x62cc6fa8d0d1b1170559640dfa2d36932712a5d761c0924ea23aabf3b602cb3c.
//
// Solidity: event ImmediateEntrySkipped(uint256 indexed transientIdx, bytes revertData)
func (_Rollups *RollupsFilterer) FilterImmediateEntrySkipped(opts *bind.FilterOpts, transientIdx []*big.Int) (*RollupsImmediateEntrySkippedIterator, error) {

	var transientIdxRule []interface{}
	for _, transientIdxItem := range transientIdx {
		transientIdxRule = append(transientIdxRule, transientIdxItem)
	}

	logs, sub, err := _Rollups.contract.FilterLogs(opts, "ImmediateEntrySkipped", transientIdxRule)
	if err != nil {
		return nil, err
	}
	return &RollupsImmediateEntrySkippedIterator{contract: _Rollups.contract, event: "ImmediateEntrySkipped", logs: logs, sub: sub}, nil
}

// WatchImmediateEntrySkipped is a free log subscription operation binding the contract event 0x62cc6fa8d0d1b1170559640dfa2d36932712a5d761c0924ea23aabf3b602cb3c.
//
// Solidity: event ImmediateEntrySkipped(uint256 indexed transientIdx, bytes revertData)
func (_Rollups *RollupsFilterer) WatchImmediateEntrySkipped(opts *bind.WatchOpts, sink chan<- *RollupsImmediateEntrySkipped, transientIdx []*big.Int) (event.Subscription, error) {

	var transientIdxRule []interface{}
	for _, transientIdxItem := range transientIdx {
		transientIdxRule = append(transientIdxRule, transientIdxItem)
	}

	logs, sub, err := _Rollups.contract.WatchLogs(opts, "ImmediateEntrySkipped", transientIdxRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(RollupsImmediateEntrySkipped)
				if err := _Rollups.contract.UnpackLog(event, "ImmediateEntrySkipped", log); err != nil {
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

// ParseImmediateEntrySkipped is a log parse operation binding the contract event 0x62cc6fa8d0d1b1170559640dfa2d36932712a5d761c0924ea23aabf3b602cb3c.
//
// Solidity: event ImmediateEntrySkipped(uint256 indexed transientIdx, bytes revertData)
func (_Rollups *RollupsFilterer) ParseImmediateEntrySkipped(log types.Log) (*RollupsImmediateEntrySkipped, error) {
	event := new(RollupsImmediateEntrySkipped)
	if err := _Rollups.contract.UnpackLog(event, "ImmediateEntrySkipped", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// RollupsL2ExecutionPerformedIterator is returned from FilterL2ExecutionPerformed and is used to iterate over the raw logs and unpacked data for L2ExecutionPerformed events raised by the Rollups contract.
type RollupsL2ExecutionPerformedIterator struct {
	Event *RollupsL2ExecutionPerformed // Event containing the contract specifics and raw log

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
func (it *RollupsL2ExecutionPerformedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(RollupsL2ExecutionPerformed)
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
		it.Event = new(RollupsL2ExecutionPerformed)
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
func (it *RollupsL2ExecutionPerformedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *RollupsL2ExecutionPerformedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// RollupsL2ExecutionPerformed represents a L2ExecutionPerformed event raised by the Rollups contract.
type RollupsL2ExecutionPerformed struct {
	RollupId *big.Int
	NewState [32]byte
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterL2ExecutionPerformed is a free log retrieval operation binding the contract event 0x0133f662c29e67eedfc9b53c0c1f657b30ebaf9748094d09fa4659d769dd4f78.
//
// Solidity: event L2ExecutionPerformed(uint256 indexed rollupId, bytes32 newState)
func (_Rollups *RollupsFilterer) FilterL2ExecutionPerformed(opts *bind.FilterOpts, rollupId []*big.Int) (*RollupsL2ExecutionPerformedIterator, error) {

	var rollupIdRule []interface{}
	for _, rollupIdItem := range rollupId {
		rollupIdRule = append(rollupIdRule, rollupIdItem)
	}

	logs, sub, err := _Rollups.contract.FilterLogs(opts, "L2ExecutionPerformed", rollupIdRule)
	if err != nil {
		return nil, err
	}
	return &RollupsL2ExecutionPerformedIterator{contract: _Rollups.contract, event: "L2ExecutionPerformed", logs: logs, sub: sub}, nil
}

// WatchL2ExecutionPerformed is a free log subscription operation binding the contract event 0x0133f662c29e67eedfc9b53c0c1f657b30ebaf9748094d09fa4659d769dd4f78.
//
// Solidity: event L2ExecutionPerformed(uint256 indexed rollupId, bytes32 newState)
func (_Rollups *RollupsFilterer) WatchL2ExecutionPerformed(opts *bind.WatchOpts, sink chan<- *RollupsL2ExecutionPerformed, rollupId []*big.Int) (event.Subscription, error) {

	var rollupIdRule []interface{}
	for _, rollupIdItem := range rollupId {
		rollupIdRule = append(rollupIdRule, rollupIdItem)
	}

	logs, sub, err := _Rollups.contract.WatchLogs(opts, "L2ExecutionPerformed", rollupIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(RollupsL2ExecutionPerformed)
				if err := _Rollups.contract.UnpackLog(event, "L2ExecutionPerformed", log); err != nil {
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

// ParseL2ExecutionPerformed is a log parse operation binding the contract event 0x0133f662c29e67eedfc9b53c0c1f657b30ebaf9748094d09fa4659d769dd4f78.
//
// Solidity: event L2ExecutionPerformed(uint256 indexed rollupId, bytes32 newState)
func (_Rollups *RollupsFilterer) ParseL2ExecutionPerformed(log types.Log) (*RollupsL2ExecutionPerformed, error) {
	event := new(RollupsL2ExecutionPerformed)
	if err := _Rollups.contract.UnpackLog(event, "L2ExecutionPerformed", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// RollupsL2TXExecutedIterator is returned from FilterL2TXExecuted and is used to iterate over the raw logs and unpacked data for L2TXExecuted events raised by the Rollups contract.
type RollupsL2TXExecutedIterator struct {
	Event *RollupsL2TXExecuted // Event containing the contract specifics and raw log

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
func (it *RollupsL2TXExecutedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(RollupsL2TXExecuted)
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
		it.Event = new(RollupsL2TXExecuted)
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
func (it *RollupsL2TXExecutedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *RollupsL2TXExecutedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// RollupsL2TXExecuted represents a L2TXExecuted event raised by the Rollups contract.
type RollupsL2TXExecuted struct {
	RollupId *big.Int
	Cursor   *big.Int
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterL2TXExecuted is a free log retrieval operation binding the contract event 0x110e37308dbd81a8e4d42e316af40e7cf78577d19f52c170928664fe4a6a377f.
//
// Solidity: event L2TXExecuted(uint256 indexed rollupId, uint256 indexed cursor)
func (_Rollups *RollupsFilterer) FilterL2TXExecuted(opts *bind.FilterOpts, rollupId []*big.Int, cursor []*big.Int) (*RollupsL2TXExecutedIterator, error) {

	var rollupIdRule []interface{}
	for _, rollupIdItem := range rollupId {
		rollupIdRule = append(rollupIdRule, rollupIdItem)
	}
	var cursorRule []interface{}
	for _, cursorItem := range cursor {
		cursorRule = append(cursorRule, cursorItem)
	}

	logs, sub, err := _Rollups.contract.FilterLogs(opts, "L2TXExecuted", rollupIdRule, cursorRule)
	if err != nil {
		return nil, err
	}
	return &RollupsL2TXExecutedIterator{contract: _Rollups.contract, event: "L2TXExecuted", logs: logs, sub: sub}, nil
}

// WatchL2TXExecuted is a free log subscription operation binding the contract event 0x110e37308dbd81a8e4d42e316af40e7cf78577d19f52c170928664fe4a6a377f.
//
// Solidity: event L2TXExecuted(uint256 indexed rollupId, uint256 indexed cursor)
func (_Rollups *RollupsFilterer) WatchL2TXExecuted(opts *bind.WatchOpts, sink chan<- *RollupsL2TXExecuted, rollupId []*big.Int, cursor []*big.Int) (event.Subscription, error) {

	var rollupIdRule []interface{}
	for _, rollupIdItem := range rollupId {
		rollupIdRule = append(rollupIdRule, rollupIdItem)
	}
	var cursorRule []interface{}
	for _, cursorItem := range cursor {
		cursorRule = append(cursorRule, cursorItem)
	}

	logs, sub, err := _Rollups.contract.WatchLogs(opts, "L2TXExecuted", rollupIdRule, cursorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(RollupsL2TXExecuted)
				if err := _Rollups.contract.UnpackLog(event, "L2TXExecuted", log); err != nil {
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

// ParseL2TXExecuted is a log parse operation binding the contract event 0x110e37308dbd81a8e4d42e316af40e7cf78577d19f52c170928664fe4a6a377f.
//
// Solidity: event L2TXExecuted(uint256 indexed rollupId, uint256 indexed cursor)
func (_Rollups *RollupsFilterer) ParseL2TXExecuted(log types.Log) (*RollupsL2TXExecuted, error) {
	event := new(RollupsL2TXExecuted)
	if err := _Rollups.contract.UnpackLog(event, "L2TXExecuted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// RollupsNestedActionConsumedIterator is returned from FilterNestedActionConsumed and is used to iterate over the raw logs and unpacked data for NestedActionConsumed events raised by the Rollups contract.
type RollupsNestedActionConsumedIterator struct {
	Event *RollupsNestedActionConsumed // Event containing the contract specifics and raw log

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
func (it *RollupsNestedActionConsumedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(RollupsNestedActionConsumed)
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
		it.Event = new(RollupsNestedActionConsumed)
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
func (it *RollupsNestedActionConsumedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *RollupsNestedActionConsumedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// RollupsNestedActionConsumed represents a NestedActionConsumed event raised by the Rollups contract.
type RollupsNestedActionConsumed struct {
	EntryIndex         *big.Int
	NestedNumber       *big.Int
	CrossChainCallHash [32]byte
	CallCount          *big.Int
	Raw                types.Log // Blockchain specific contextual infos
}

// FilterNestedActionConsumed is a free log retrieval operation binding the contract event 0xed18c2605a6158b1b555df8746ce7bd8bffd2154ceefbe38fab02b889a400612.
//
// Solidity: event NestedActionConsumed(uint256 indexed entryIndex, uint256 indexed nestedNumber, bytes32 crossChainCallHash, uint256 callCount)
func (_Rollups *RollupsFilterer) FilterNestedActionConsumed(opts *bind.FilterOpts, entryIndex []*big.Int, nestedNumber []*big.Int) (*RollupsNestedActionConsumedIterator, error) {

	var entryIndexRule []interface{}
	for _, entryIndexItem := range entryIndex {
		entryIndexRule = append(entryIndexRule, entryIndexItem)
	}
	var nestedNumberRule []interface{}
	for _, nestedNumberItem := range nestedNumber {
		nestedNumberRule = append(nestedNumberRule, nestedNumberItem)
	}

	logs, sub, err := _Rollups.contract.FilterLogs(opts, "NestedActionConsumed", entryIndexRule, nestedNumberRule)
	if err != nil {
		return nil, err
	}
	return &RollupsNestedActionConsumedIterator{contract: _Rollups.contract, event: "NestedActionConsumed", logs: logs, sub: sub}, nil
}

// WatchNestedActionConsumed is a free log subscription operation binding the contract event 0xed18c2605a6158b1b555df8746ce7bd8bffd2154ceefbe38fab02b889a400612.
//
// Solidity: event NestedActionConsumed(uint256 indexed entryIndex, uint256 indexed nestedNumber, bytes32 crossChainCallHash, uint256 callCount)
func (_Rollups *RollupsFilterer) WatchNestedActionConsumed(opts *bind.WatchOpts, sink chan<- *RollupsNestedActionConsumed, entryIndex []*big.Int, nestedNumber []*big.Int) (event.Subscription, error) {

	var entryIndexRule []interface{}
	for _, entryIndexItem := range entryIndex {
		entryIndexRule = append(entryIndexRule, entryIndexItem)
	}
	var nestedNumberRule []interface{}
	for _, nestedNumberItem := range nestedNumber {
		nestedNumberRule = append(nestedNumberRule, nestedNumberItem)
	}

	logs, sub, err := _Rollups.contract.WatchLogs(opts, "NestedActionConsumed", entryIndexRule, nestedNumberRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(RollupsNestedActionConsumed)
				if err := _Rollups.contract.UnpackLog(event, "NestedActionConsumed", log); err != nil {
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

// ParseNestedActionConsumed is a log parse operation binding the contract event 0xed18c2605a6158b1b555df8746ce7bd8bffd2154ceefbe38fab02b889a400612.
//
// Solidity: event NestedActionConsumed(uint256 indexed entryIndex, uint256 indexed nestedNumber, bytes32 crossChainCallHash, uint256 callCount)
func (_Rollups *RollupsFilterer) ParseNestedActionConsumed(log types.Log) (*RollupsNestedActionConsumed, error) {
	event := new(RollupsNestedActionConsumed)
	if err := _Rollups.contract.UnpackLog(event, "NestedActionConsumed", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// RollupsRevertSpanExecutedIterator is returned from FilterRevertSpanExecuted and is used to iterate over the raw logs and unpacked data for RevertSpanExecuted events raised by the Rollups contract.
type RollupsRevertSpanExecutedIterator struct {
	Event *RollupsRevertSpanExecuted // Event containing the contract specifics and raw log

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
func (it *RollupsRevertSpanExecutedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(RollupsRevertSpanExecuted)
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
		it.Event = new(RollupsRevertSpanExecuted)
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
func (it *RollupsRevertSpanExecutedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *RollupsRevertSpanExecutedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// RollupsRevertSpanExecuted represents a RevertSpanExecuted event raised by the Rollups contract.
type RollupsRevertSpanExecuted struct {
	EntryIndex      *big.Int
	StartCallNumber *big.Int
	Span            *big.Int
	Raw             types.Log // Blockchain specific contextual infos
}

// FilterRevertSpanExecuted is a free log retrieval operation binding the contract event 0xf805bc66587a02266edd3197dc8dc57351fa8ae0b55a1e75f270f70d80159a7c.
//
// Solidity: event RevertSpanExecuted(uint256 indexed entryIndex, uint256 startCallNumber, uint256 span)
func (_Rollups *RollupsFilterer) FilterRevertSpanExecuted(opts *bind.FilterOpts, entryIndex []*big.Int) (*RollupsRevertSpanExecutedIterator, error) {

	var entryIndexRule []interface{}
	for _, entryIndexItem := range entryIndex {
		entryIndexRule = append(entryIndexRule, entryIndexItem)
	}

	logs, sub, err := _Rollups.contract.FilterLogs(opts, "RevertSpanExecuted", entryIndexRule)
	if err != nil {
		return nil, err
	}
	return &RollupsRevertSpanExecutedIterator{contract: _Rollups.contract, event: "RevertSpanExecuted", logs: logs, sub: sub}, nil
}

// WatchRevertSpanExecuted is a free log subscription operation binding the contract event 0xf805bc66587a02266edd3197dc8dc57351fa8ae0b55a1e75f270f70d80159a7c.
//
// Solidity: event RevertSpanExecuted(uint256 indexed entryIndex, uint256 startCallNumber, uint256 span)
func (_Rollups *RollupsFilterer) WatchRevertSpanExecuted(opts *bind.WatchOpts, sink chan<- *RollupsRevertSpanExecuted, entryIndex []*big.Int) (event.Subscription, error) {

	var entryIndexRule []interface{}
	for _, entryIndexItem := range entryIndex {
		entryIndexRule = append(entryIndexRule, entryIndexItem)
	}

	logs, sub, err := _Rollups.contract.WatchLogs(opts, "RevertSpanExecuted", entryIndexRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(RollupsRevertSpanExecuted)
				if err := _Rollups.contract.UnpackLog(event, "RevertSpanExecuted", log); err != nil {
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

// ParseRevertSpanExecuted is a log parse operation binding the contract event 0xf805bc66587a02266edd3197dc8dc57351fa8ae0b55a1e75f270f70d80159a7c.
//
// Solidity: event RevertSpanExecuted(uint256 indexed entryIndex, uint256 startCallNumber, uint256 span)
func (_Rollups *RollupsFilterer) ParseRevertSpanExecuted(log types.Log) (*RollupsRevertSpanExecuted, error) {
	event := new(RollupsRevertSpanExecuted)
	if err := _Rollups.contract.UnpackLog(event, "RevertSpanExecuted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// RollupsRollupContractChangedIterator is returned from FilterRollupContractChanged and is used to iterate over the raw logs and unpacked data for RollupContractChanged events raised by the Rollups contract.
type RollupsRollupContractChangedIterator struct {
	Event *RollupsRollupContractChanged // Event containing the contract specifics and raw log

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
func (it *RollupsRollupContractChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(RollupsRollupContractChanged)
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
		it.Event = new(RollupsRollupContractChanged)
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
func (it *RollupsRollupContractChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *RollupsRollupContractChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// RollupsRollupContractChanged represents a RollupContractChanged event raised by the Rollups contract.
type RollupsRollupContractChanged struct {
	RollupId         *big.Int
	PreviousContract common.Address
	NewContract      common.Address
	Raw              types.Log // Blockchain specific contextual infos
}

// FilterRollupContractChanged is a free log retrieval operation binding the contract event 0xcb1f89865b4ce341946f7baead20a32f5c5c8694fcfc0d2e7109ad2f7754a7e3.
//
// Solidity: event RollupContractChanged(uint256 indexed rollupId, address indexed previousContract, address indexed newContract)
func (_Rollups *RollupsFilterer) FilterRollupContractChanged(opts *bind.FilterOpts, rollupId []*big.Int, previousContract []common.Address, newContract []common.Address) (*RollupsRollupContractChangedIterator, error) {

	var rollupIdRule []interface{}
	for _, rollupIdItem := range rollupId {
		rollupIdRule = append(rollupIdRule, rollupIdItem)
	}
	var previousContractRule []interface{}
	for _, previousContractItem := range previousContract {
		previousContractRule = append(previousContractRule, previousContractItem)
	}
	var newContractRule []interface{}
	for _, newContractItem := range newContract {
		newContractRule = append(newContractRule, newContractItem)
	}

	logs, sub, err := _Rollups.contract.FilterLogs(opts, "RollupContractChanged", rollupIdRule, previousContractRule, newContractRule)
	if err != nil {
		return nil, err
	}
	return &RollupsRollupContractChangedIterator{contract: _Rollups.contract, event: "RollupContractChanged", logs: logs, sub: sub}, nil
}

// WatchRollupContractChanged is a free log subscription operation binding the contract event 0xcb1f89865b4ce341946f7baead20a32f5c5c8694fcfc0d2e7109ad2f7754a7e3.
//
// Solidity: event RollupContractChanged(uint256 indexed rollupId, address indexed previousContract, address indexed newContract)
func (_Rollups *RollupsFilterer) WatchRollupContractChanged(opts *bind.WatchOpts, sink chan<- *RollupsRollupContractChanged, rollupId []*big.Int, previousContract []common.Address, newContract []common.Address) (event.Subscription, error) {

	var rollupIdRule []interface{}
	for _, rollupIdItem := range rollupId {
		rollupIdRule = append(rollupIdRule, rollupIdItem)
	}
	var previousContractRule []interface{}
	for _, previousContractItem := range previousContract {
		previousContractRule = append(previousContractRule, previousContractItem)
	}
	var newContractRule []interface{}
	for _, newContractItem := range newContract {
		newContractRule = append(newContractRule, newContractItem)
	}

	logs, sub, err := _Rollups.contract.WatchLogs(opts, "RollupContractChanged", rollupIdRule, previousContractRule, newContractRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(RollupsRollupContractChanged)
				if err := _Rollups.contract.UnpackLog(event, "RollupContractChanged", log); err != nil {
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

// ParseRollupContractChanged is a log parse operation binding the contract event 0xcb1f89865b4ce341946f7baead20a32f5c5c8694fcfc0d2e7109ad2f7754a7e3.
//
// Solidity: event RollupContractChanged(uint256 indexed rollupId, address indexed previousContract, address indexed newContract)
func (_Rollups *RollupsFilterer) ParseRollupContractChanged(log types.Log) (*RollupsRollupContractChanged, error) {
	event := new(RollupsRollupContractChanged)
	if err := _Rollups.contract.UnpackLog(event, "RollupContractChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// RollupsRollupCreatedIterator is returned from FilterRollupCreated and is used to iterate over the raw logs and unpacked data for RollupCreated events raised by the Rollups contract.
type RollupsRollupCreatedIterator struct {
	Event *RollupsRollupCreated // Event containing the contract specifics and raw log

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
func (it *RollupsRollupCreatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(RollupsRollupCreated)
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
		it.Event = new(RollupsRollupCreated)
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
func (it *RollupsRollupCreatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *RollupsRollupCreatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// RollupsRollupCreated represents a RollupCreated event raised by the Rollups contract.
type RollupsRollupCreated struct {
	RollupId       *big.Int
	RollupContract common.Address
	InitialState   [32]byte
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterRollupCreated is a free log retrieval operation binding the contract event 0x65aca4a0612183373cd8f999716ed850923221124a6c3b5c55b90e964cbe34d9.
//
// Solidity: event RollupCreated(uint256 indexed rollupId, address indexed rollupContract, bytes32 initialState)
func (_Rollups *RollupsFilterer) FilterRollupCreated(opts *bind.FilterOpts, rollupId []*big.Int, rollupContract []common.Address) (*RollupsRollupCreatedIterator, error) {

	var rollupIdRule []interface{}
	for _, rollupIdItem := range rollupId {
		rollupIdRule = append(rollupIdRule, rollupIdItem)
	}
	var rollupContractRule []interface{}
	for _, rollupContractItem := range rollupContract {
		rollupContractRule = append(rollupContractRule, rollupContractItem)
	}

	logs, sub, err := _Rollups.contract.FilterLogs(opts, "RollupCreated", rollupIdRule, rollupContractRule)
	if err != nil {
		return nil, err
	}
	return &RollupsRollupCreatedIterator{contract: _Rollups.contract, event: "RollupCreated", logs: logs, sub: sub}, nil
}

// WatchRollupCreated is a free log subscription operation binding the contract event 0x65aca4a0612183373cd8f999716ed850923221124a6c3b5c55b90e964cbe34d9.
//
// Solidity: event RollupCreated(uint256 indexed rollupId, address indexed rollupContract, bytes32 initialState)
func (_Rollups *RollupsFilterer) WatchRollupCreated(opts *bind.WatchOpts, sink chan<- *RollupsRollupCreated, rollupId []*big.Int, rollupContract []common.Address) (event.Subscription, error) {

	var rollupIdRule []interface{}
	for _, rollupIdItem := range rollupId {
		rollupIdRule = append(rollupIdRule, rollupIdItem)
	}
	var rollupContractRule []interface{}
	for _, rollupContractItem := range rollupContract {
		rollupContractRule = append(rollupContractRule, rollupContractItem)
	}

	logs, sub, err := _Rollups.contract.WatchLogs(opts, "RollupCreated", rollupIdRule, rollupContractRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(RollupsRollupCreated)
				if err := _Rollups.contract.UnpackLog(event, "RollupCreated", log); err != nil {
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

// ParseRollupCreated is a log parse operation binding the contract event 0x65aca4a0612183373cd8f999716ed850923221124a6c3b5c55b90e964cbe34d9.
//
// Solidity: event RollupCreated(uint256 indexed rollupId, address indexed rollupContract, bytes32 initialState)
func (_Rollups *RollupsFilterer) ParseRollupCreated(log types.Log) (*RollupsRollupCreated, error) {
	event := new(RollupsRollupCreated)
	if err := _Rollups.contract.UnpackLog(event, "RollupCreated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// RollupsStateUpdatedIterator is returned from FilterStateUpdated and is used to iterate over the raw logs and unpacked data for StateUpdated events raised by the Rollups contract.
type RollupsStateUpdatedIterator struct {
	Event *RollupsStateUpdated // Event containing the contract specifics and raw log

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
func (it *RollupsStateUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(RollupsStateUpdated)
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
		it.Event = new(RollupsStateUpdated)
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
func (it *RollupsStateUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *RollupsStateUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// RollupsStateUpdated represents a StateUpdated event raised by the Rollups contract.
type RollupsStateUpdated struct {
	RollupId     *big.Int
	NewStateRoot [32]byte
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterStateUpdated is a free log retrieval operation binding the contract event 0x3060356fe01e6782d366a741b8a2f65e53c64fefa69c21cf85ad79a846fd472e.
//
// Solidity: event StateUpdated(uint256 indexed rollupId, bytes32 newStateRoot)
func (_Rollups *RollupsFilterer) FilterStateUpdated(opts *bind.FilterOpts, rollupId []*big.Int) (*RollupsStateUpdatedIterator, error) {

	var rollupIdRule []interface{}
	for _, rollupIdItem := range rollupId {
		rollupIdRule = append(rollupIdRule, rollupIdItem)
	}

	logs, sub, err := _Rollups.contract.FilterLogs(opts, "StateUpdated", rollupIdRule)
	if err != nil {
		return nil, err
	}
	return &RollupsStateUpdatedIterator{contract: _Rollups.contract, event: "StateUpdated", logs: logs, sub: sub}, nil
}

// WatchStateUpdated is a free log subscription operation binding the contract event 0x3060356fe01e6782d366a741b8a2f65e53c64fefa69c21cf85ad79a846fd472e.
//
// Solidity: event StateUpdated(uint256 indexed rollupId, bytes32 newStateRoot)
func (_Rollups *RollupsFilterer) WatchStateUpdated(opts *bind.WatchOpts, sink chan<- *RollupsStateUpdated, rollupId []*big.Int) (event.Subscription, error) {

	var rollupIdRule []interface{}
	for _, rollupIdItem := range rollupId {
		rollupIdRule = append(rollupIdRule, rollupIdItem)
	}

	logs, sub, err := _Rollups.contract.WatchLogs(opts, "StateUpdated", rollupIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(RollupsStateUpdated)
				if err := _Rollups.contract.UnpackLog(event, "StateUpdated", log); err != nil {
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

// ParseStateUpdated is a log parse operation binding the contract event 0x3060356fe01e6782d366a741b8a2f65e53c64fefa69c21cf85ad79a846fd472e.
//
// Solidity: event StateUpdated(uint256 indexed rollupId, bytes32 newStateRoot)
func (_Rollups *RollupsFilterer) ParseStateUpdated(log types.Log) (*RollupsStateUpdated, error) {
	event := new(RollupsStateUpdated)
	if err := _Rollups.contract.UnpackLog(event, "StateUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
