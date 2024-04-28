// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package zksyncstatetransition

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

// BridgehubL2TransactionRequest is an auto generated low-level Go binding around an user-defined struct.
type BridgehubL2TransactionRequest struct {
	Sender                   common.Address
	ContractL2               common.Address
	MintValue                *big.Int
	L2Value                  *big.Int
	L2Calldata               []byte
	L2GasLimit               *big.Int
	L2GasPerPubdataByteLimit *big.Int
	FactoryDeps              [][]byte
	RefundRecipient          common.Address
}

// DiamondDiamondCutData is an auto generated low-level Go binding around an user-defined struct.
type DiamondDiamondCutData struct {
	FacetCuts    []DiamondFacetCut
	InitAddress  common.Address
	InitCalldata []byte
}

// DiamondFacetCut is an auto generated low-level Go binding around an user-defined struct.
type DiamondFacetCut struct {
	Facet       common.Address
	Action      uint8
	IsFreezable bool
	Selectors   [][4]byte
}

// FeeParams is an auto generated low-level Go binding around an user-defined struct.
type FeeParams struct {
	PubdataPricingMode   uint8
	BatchOverheadL1Gas   uint32
	MaxPubdataPerBatch   uint32
	MaxL2GasPerBatch     uint32
	PriorityTxMaxPubdata uint32
	MinimalL2GasPrice    uint64
}

// IExecutorCommitBatchInfo is an auto generated low-level Go binding around an user-defined struct.
type IExecutorCommitBatchInfo struct {
	BatchNumber                       uint64
	Timestamp                         uint64
	IndexRepeatedStorageChanges       uint64
	NewStateRoot                      [32]byte
	NumberOfLayer1Txs                 *big.Int
	PriorityOperationsHash            [32]byte
	BootloaderHeapInitialContentsHash [32]byte
	EventsQueueStateHash              [32]byte
	SystemLogs                        []byte
	PubdataCommitments                []byte
}

// IExecutorProofInput is an auto generated low-level Go binding around an user-defined struct.
type IExecutorProofInput struct {
	RecursiveAggregationInput []*big.Int
	SerializedProof           []*big.Int
}

// IExecutorStoredBatchInfo is an auto generated low-level Go binding around an user-defined struct.
type IExecutorStoredBatchInfo struct {
	BatchNumber                 uint64
	BatchHash                   [32]byte
	IndexRepeatedStorageChanges uint64
	NumberOfLayer1Txs           *big.Int
	PriorityOperationsHash      [32]byte
	L2LogsTreeRoot              [32]byte
	Timestamp                   *big.Int
	Commitment                  [32]byte
}

// IGettersFacet is an auto generated low-level Go binding around an user-defined struct.
type IGettersFacet struct {
	Addr      common.Address
	Selectors [][4]byte
}

// L2CanonicalTransaction is an auto generated low-level Go binding around an user-defined struct.
type L2CanonicalTransaction struct {
	TxType                 *big.Int
	From                   *big.Int
	To                     *big.Int
	GasLimit               *big.Int
	GasPerPubdataByteLimit *big.Int
	MaxFeePerGas           *big.Int
	MaxPriorityFeePerGas   *big.Int
	Paymaster              *big.Int
	Nonce                  *big.Int
	Value                  *big.Int
	Reserved               [4]*big.Int
	Data                   []byte
	Signature              []byte
	FactoryDeps            []*big.Int
	PaymasterInput         []byte
	ReservedDynamic        []byte
}

// L2Log is an auto generated low-level Go binding around an user-defined struct.
type L2Log struct {
	L2ShardId       uint8
	IsService       bool
	TxNumberInBatch uint16
	Sender          common.Address
	Key             [32]byte
	Value           [32]byte
}

// L2Message is an auto generated low-level Go binding around an user-defined struct.
type L2Message struct {
	TxNumberInBatch uint16
	Sender          common.Address
	Data            []byte
}

// PriorityOperation is an auto generated low-level Go binding around an user-defined struct.
type PriorityOperation struct {
	CanonicalTxHash     [32]byte
	ExpirationTimestamp uint64
	Layer2Tip           *big.Int
}

// VerifierParams is an auto generated low-level Go binding around an user-defined struct.
type VerifierParams struct {
	RecursionNodeLevelVkHash    [32]byte
	RecursionLeafLevelVkHash    [32]byte
	RecursionCircuitsSetVksHash [32]byte
}

// IZkSyncStateTransitionMetaData contains all meta data concerning the IZkSyncStateTransition contract.
var IZkSyncStateTransitionMetaData = &bind.MetaData{
	ABI: "[{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"batchNumber\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"batchHash\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"commitment\",\"type\":\"bytes32\"}],\"name\":\"BlockCommit\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"batchNumber\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"batchHash\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"commitment\",\"type\":\"bytes32\"}],\"name\":\"BlockExecution\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"totalBatchesCommitted\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"totalBatchesVerified\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"totalBatchesExecuted\",\"type\":\"uint256\"}],\"name\":\"BlocksRevert\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"previousLastVerifiedBatch\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"currentLastVerifiedBatch\",\"type\":\"uint256\"}],\"name\":\"BlocksVerification\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"EthWithdrawalFinalized\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"components\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"facet\",\"type\":\"address\"},{\"internalType\":\"enumDiamond.Action\",\"name\":\"action\",\"type\":\"uint8\"},{\"internalType\":\"bool\",\"name\":\"isFreezable\",\"type\":\"bool\"},{\"internalType\":\"bytes4[]\",\"name\":\"selectors\",\"type\":\"bytes4[]\"}],\"internalType\":\"structDiamond.FacetCut[]\",\"name\":\"facetCuts\",\"type\":\"tuple[]\"},{\"internalType\":\"address\",\"name\":\"initAddress\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"initCalldata\",\"type\":\"bytes\"}],\"indexed\":false,\"internalType\":\"structDiamond.DiamondCutData\",\"name\":\"diamondCut\",\"type\":\"tuple\"}],\"name\":\"ExecuteUpgrade\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[],\"name\":\"Freeze\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"isPorterAvailable\",\"type\":\"bool\"}],\"name\":\"IsPorterAvailableStatusUpdate\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"oldAdmin\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newAdmin\",\"type\":\"address\"}],\"name\":\"NewAdmin\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint128\",\"name\":\"oldNominator\",\"type\":\"uint128\"},{\"indexed\":false,\"internalType\":\"uint128\",\"name\":\"oldDenominator\",\"type\":\"uint128\"},{\"indexed\":false,\"internalType\":\"uint128\",\"name\":\"newNominator\",\"type\":\"uint128\"},{\"indexed\":false,\"internalType\":\"uint128\",\"name\":\"newDenominator\",\"type\":\"uint128\"}],\"name\":\"NewBaseTokenMultiplier\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"components\":[{\"internalType\":\"enumPubdataPricingMode\",\"name\":\"pubdataPricingMode\",\"type\":\"uint8\"},{\"internalType\":\"uint32\",\"name\":\"batchOverheadL1Gas\",\"type\":\"uint32\"},{\"internalType\":\"uint32\",\"name\":\"maxPubdataPerBatch\",\"type\":\"uint32\"},{\"internalType\":\"uint32\",\"name\":\"maxL2GasPerBatch\",\"type\":\"uint32\"},{\"internalType\":\"uint32\",\"name\":\"priorityTxMaxPubdata\",\"type\":\"uint32\"},{\"internalType\":\"uint64\",\"name\":\"minimalL2GasPrice\",\"type\":\"uint64\"}],\"indexed\":false,\"internalType\":\"structFeeParams\",\"name\":\"oldFeeParams\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"enumPubdataPricingMode\",\"name\":\"pubdataPricingMode\",\"type\":\"uint8\"},{\"internalType\":\"uint32\",\"name\":\"batchOverheadL1Gas\",\"type\":\"uint32\"},{\"internalType\":\"uint32\",\"name\":\"maxPubdataPerBatch\",\"type\":\"uint32\"},{\"internalType\":\"uint32\",\"name\":\"maxL2GasPerBatch\",\"type\":\"uint32\"},{\"internalType\":\"uint32\",\"name\":\"priorityTxMaxPubdata\",\"type\":\"uint32\"},{\"internalType\":\"uint64\",\"name\":\"minimalL2GasPrice\",\"type\":\"uint64\"}],\"indexed\":false,\"internalType\":\"structFeeParams\",\"name\":\"newFeeParams\",\"type\":\"tuple\"}],\"name\":\"NewFeeParams\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"oldPendingAdmin\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newPendingAdmin\",\"type\":\"address\"}],\"name\":\"NewPendingAdmin\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"txId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"txHash\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"expirationTimestamp\",\"type\":\"uint64\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"txType\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"from\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"to\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"gasLimit\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"gasPerPubdataByteLimit\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"maxFeePerGas\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"maxPriorityFeePerGas\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"paymaster\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"nonce\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"},{\"internalType\":\"uint256[4]\",\"name\":\"reserved\",\"type\":\"uint256[4]\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"signature\",\"type\":\"bytes\"},{\"internalType\":\"uint256[]\",\"name\":\"factoryDeps\",\"type\":\"uint256[]\"},{\"internalType\":\"bytes\",\"name\":\"paymasterInput\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"reservedDynamic\",\"type\":\"bytes\"}],\"indexed\":false,\"internalType\":\"structL2CanonicalTransaction\",\"name\":\"transaction\",\"type\":\"tuple\"},{\"indexed\":false,\"internalType\":\"bytes[]\",\"name\":\"factoryDeps\",\"type\":\"bytes[]\"}],\"name\":\"NewPriorityRequest\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"oldPriorityTxMaxGasLimit\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"newPriorityTxMaxGasLimit\",\"type\":\"uint256\"}],\"name\":\"NewPriorityTxMaxGasLimit\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"oldTransactionFilterer\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"newTransactionFilterer\",\"type\":\"address\"}],\"name\":\"NewTransactionFilterer\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"components\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"facet\",\"type\":\"address\"},{\"internalType\":\"enumDiamond.Action\",\"name\":\"action\",\"type\":\"uint8\"},{\"internalType\":\"bool\",\"name\":\"isFreezable\",\"type\":\"bool\"},{\"internalType\":\"bytes4[]\",\"name\":\"selectors\",\"type\":\"bytes4[]\"}],\"internalType\":\"structDiamond.FacetCut[]\",\"name\":\"facetCuts\",\"type\":\"tuple[]\"},{\"internalType\":\"address\",\"name\":\"initAddress\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"initCalldata\",\"type\":\"bytes\"}],\"indexed\":false,\"internalType\":\"structDiamond.DiamondCutData\",\"name\":\"diamondCut\",\"type\":\"tuple\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"proposalId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"proposalSalt\",\"type\":\"bytes32\"}],\"name\":\"ProposeTransparentUpgrade\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[],\"name\":\"Unfreeze\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"validatorAddress\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"isActive\",\"type\":\"bool\"}],\"name\":\"ValidatorStatusUpdate\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"enumPubdataPricingMode\",\"name\":\"validiumMode\",\"type\":\"uint8\"}],\"name\":\"ValidiumModeStatusUpdate\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"acceptAdmin\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"baseTokenGasPriceMultiplierDenominator\",\"outputs\":[{\"internalType\":\"uint128\",\"name\":\"\",\"type\":\"uint128\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"baseTokenGasPriceMultiplierNominator\",\"outputs\":[{\"internalType\":\"uint128\",\"name\":\"\",\"type\":\"uint128\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"contractL2\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"mintValue\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"l2Value\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"l2Calldata\",\"type\":\"bytes\"},{\"internalType\":\"uint256\",\"name\":\"l2GasLimit\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"l2GasPerPubdataByteLimit\",\"type\":\"uint256\"},{\"internalType\":\"bytes[]\",\"name\":\"factoryDeps\",\"type\":\"bytes[]\"},{\"internalType\":\"address\",\"name\":\"refundRecipient\",\"type\":\"address\"}],\"internalType\":\"structBridgehubL2TransactionRequest\",\"name\":\"_request\",\"type\":\"tuple\"}],\"name\":\"bridgehubRequestL2Transaction\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"canonicalTxHash\",\"type\":\"bytes32\"}],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"enumPubdataPricingMode\",\"name\":\"pubdataPricingMode\",\"type\":\"uint8\"},{\"internalType\":\"uint32\",\"name\":\"batchOverheadL1Gas\",\"type\":\"uint32\"},{\"internalType\":\"uint32\",\"name\":\"maxPubdataPerBatch\",\"type\":\"uint32\"},{\"internalType\":\"uint32\",\"name\":\"maxL2GasPerBatch\",\"type\":\"uint32\"},{\"internalType\":\"uint32\",\"name\":\"priorityTxMaxPubdata\",\"type\":\"uint32\"},{\"internalType\":\"uint64\",\"name\":\"minimalL2GasPrice\",\"type\":\"uint64\"}],\"internalType\":\"structFeeParams\",\"name\":\"_newFeeParams\",\"type\":\"tuple\"}],\"name\":\"changeFeeParams\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint64\",\"name\":\"batchNumber\",\"type\":\"uint64\"},{\"internalType\":\"bytes32\",\"name\":\"batchHash\",\"type\":\"bytes32\"},{\"internalType\":\"uint64\",\"name\":\"indexRepeatedStorageChanges\",\"type\":\"uint64\"},{\"internalType\":\"uint256\",\"name\":\"numberOfLayer1Txs\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"priorityOperationsHash\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"l2LogsTreeRoot\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"timestamp\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"commitment\",\"type\":\"bytes32\"}],\"internalType\":\"structIExecutor.StoredBatchInfo\",\"name\":\"_lastCommittedBatchData\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"uint64\",\"name\":\"batchNumber\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"timestamp\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"indexRepeatedStorageChanges\",\"type\":\"uint64\"},{\"internalType\":\"bytes32\",\"name\":\"newStateRoot\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"numberOfLayer1Txs\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"priorityOperationsHash\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"bootloaderHeapInitialContentsHash\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"eventsQueueStateHash\",\"type\":\"bytes32\"},{\"internalType\":\"bytes\",\"name\":\"systemLogs\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"pubdataCommitments\",\"type\":\"bytes\"}],\"internalType\":\"structIExecutor.CommitBatchInfo[]\",\"name\":\"_newBatchesData\",\"type\":\"tuple[]\"}],\"name\":\"commitBatches\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_chainId\",\"type\":\"uint256\"},{\"components\":[{\"internalType\":\"uint64\",\"name\":\"batchNumber\",\"type\":\"uint64\"},{\"internalType\":\"bytes32\",\"name\":\"batchHash\",\"type\":\"bytes32\"},{\"internalType\":\"uint64\",\"name\":\"indexRepeatedStorageChanges\",\"type\":\"uint64\"},{\"internalType\":\"uint256\",\"name\":\"numberOfLayer1Txs\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"priorityOperationsHash\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"l2LogsTreeRoot\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"timestamp\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"commitment\",\"type\":\"bytes32\"}],\"internalType\":\"structIExecutor.StoredBatchInfo\",\"name\":\"_lastCommittedBatchData\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"uint64\",\"name\":\"batchNumber\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"timestamp\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"indexRepeatedStorageChanges\",\"type\":\"uint64\"},{\"internalType\":\"bytes32\",\"name\":\"newStateRoot\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"numberOfLayer1Txs\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"priorityOperationsHash\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"bootloaderHeapInitialContentsHash\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"eventsQueueStateHash\",\"type\":\"bytes32\"},{\"internalType\":\"bytes\",\"name\":\"systemLogs\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"pubdataCommitments\",\"type\":\"bytes\"}],\"internalType\":\"structIExecutor.CommitBatchInfo[]\",\"name\":\"_newBatchesData\",\"type\":\"tuple[]\"}],\"name\":\"commitBatchesSharedBridge\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint64\",\"name\":\"batchNumber\",\"type\":\"uint64\"},{\"internalType\":\"bytes32\",\"name\":\"batchHash\",\"type\":\"bytes32\"},{\"internalType\":\"uint64\",\"name\":\"indexRepeatedStorageChanges\",\"type\":\"uint64\"},{\"internalType\":\"uint256\",\"name\":\"numberOfLayer1Txs\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"priorityOperationsHash\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"l2LogsTreeRoot\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"timestamp\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"commitment\",\"type\":\"bytes32\"}],\"internalType\":\"structIExecutor.StoredBatchInfo[]\",\"name\":\"_batchesData\",\"type\":\"tuple[]\"}],\"name\":\"executeBatches\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_chainId\",\"type\":\"uint256\"},{\"components\":[{\"internalType\":\"uint64\",\"name\":\"batchNumber\",\"type\":\"uint64\"},{\"internalType\":\"bytes32\",\"name\":\"batchHash\",\"type\":\"bytes32\"},{\"internalType\":\"uint64\",\"name\":\"indexRepeatedStorageChanges\",\"type\":\"uint64\"},{\"internalType\":\"uint256\",\"name\":\"numberOfLayer1Txs\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"priorityOperationsHash\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"l2LogsTreeRoot\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"timestamp\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"commitment\",\"type\":\"bytes32\"}],\"internalType\":\"structIExecutor.StoredBatchInfo[]\",\"name\":\"_batchesData\",\"type\":\"tuple[]\"}],\"name\":\"executeBatchesSharedBridge\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"facet\",\"type\":\"address\"},{\"internalType\":\"enumDiamond.Action\",\"name\":\"action\",\"type\":\"uint8\"},{\"internalType\":\"bool\",\"name\":\"isFreezable\",\"type\":\"bool\"},{\"internalType\":\"bytes4[]\",\"name\":\"selectors\",\"type\":\"bytes4[]\"}],\"internalType\":\"structDiamond.FacetCut[]\",\"name\":\"facetCuts\",\"type\":\"tuple[]\"},{\"internalType\":\"address\",\"name\":\"initAddress\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"initCalldata\",\"type\":\"bytes\"}],\"internalType\":\"structDiamond.DiamondCutData\",\"name\":\"_diamondCut\",\"type\":\"tuple\"}],\"name\":\"executeUpgrade\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes4\",\"name\":\"_selector\",\"type\":\"bytes4\"}],\"name\":\"facetAddress\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"facet\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"facetAddresses\",\"outputs\":[{\"internalType\":\"address[]\",\"name\":\"facets\",\"type\":\"address[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_facet\",\"type\":\"address\"}],\"name\":\"facetFunctionSelectors\",\"outputs\":[{\"internalType\":\"bytes4[]\",\"name\":\"\",\"type\":\"bytes4[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"facets\",\"outputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"addr\",\"type\":\"address\"},{\"internalType\":\"bytes4[]\",\"name\":\"selectors\",\"type\":\"bytes4[]\"}],\"internalType\":\"structIGetters.Facet[]\",\"name\":\"\",\"type\":\"tuple[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_l2BatchNumber\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_l2MessageIndex\",\"type\":\"uint256\"},{\"internalType\":\"uint16\",\"name\":\"_l2TxNumberInBatch\",\"type\":\"uint16\"},{\"internalType\":\"bytes\",\"name\":\"_message\",\"type\":\"bytes\"},{\"internalType\":\"bytes32[]\",\"name\":\"_merkleProof\",\"type\":\"bytes32[]\"}],\"name\":\"finalizeEthWithdrawal\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"freezeDiamond\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getAdmin\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getBaseToken\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getBaseTokenBridge\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getBridgehub\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getFirstUnprocessedPriorityTx\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getL2BootloaderBytecodeHash\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getL2DefaultAccountBytecodeHash\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getL2SystemContractsUpgradeBatchNumber\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getL2SystemContractsUpgradeTxHash\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getName\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getPendingAdmin\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getPriorityQueueSize\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getPriorityTxMaxGasLimit\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getProtocolVersion\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getPubdataPricingMode\",\"outputs\":[{\"internalType\":\"enumPubdataPricingMode\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getStateTransitionManager\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getTotalBatchesCommitted\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getTotalBatchesExecuted\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getTotalBatchesVerified\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getTotalPriorityTxs\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getVerifier\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getVerifierParams\",\"outputs\":[{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"recursionNodeLevelVkHash\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"recursionLeafLevelVkHash\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"recursionCircuitsSetVksHash\",\"type\":\"bytes32\"}],\"internalType\":\"structVerifierParams\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"isDiamondStorageFrozen\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_l2BatchNumber\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_l2MessageIndex\",\"type\":\"uint256\"}],\"name\":\"isEthWithdrawalFinalized\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_facet\",\"type\":\"address\"}],\"name\":\"isFacetFreezable\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"isFreezable\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes4\",\"name\":\"_selector\",\"type\":\"bytes4\"}],\"name\":\"isFunctionFreezable\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_address\",\"type\":\"address\"}],\"name\":\"isValidator\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_batchNumber\",\"type\":\"uint256\"}],\"name\":\"l2LogsRootHash\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"merkleRoot\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_gasPrice\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_l2GasLimit\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_l2GasPerPubdataByteLimit\",\"type\":\"uint256\"}],\"name\":\"l2TransactionBaseCost\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"priorityQueueFrontOperation\",\"outputs\":[{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"canonicalTxHash\",\"type\":\"bytes32\"},{\"internalType\":\"uint64\",\"name\":\"expirationTimestamp\",\"type\":\"uint64\"},{\"internalType\":\"uint192\",\"name\":\"layer2Tip\",\"type\":\"uint192\"}],\"internalType\":\"structPriorityOperation\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint64\",\"name\":\"batchNumber\",\"type\":\"uint64\"},{\"internalType\":\"bytes32\",\"name\":\"batchHash\",\"type\":\"bytes32\"},{\"internalType\":\"uint64\",\"name\":\"indexRepeatedStorageChanges\",\"type\":\"uint64\"},{\"internalType\":\"uint256\",\"name\":\"numberOfLayer1Txs\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"priorityOperationsHash\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"l2LogsTreeRoot\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"timestamp\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"commitment\",\"type\":\"bytes32\"}],\"internalType\":\"structIExecutor.StoredBatchInfo\",\"name\":\"_prevBatch\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"uint64\",\"name\":\"batchNumber\",\"type\":\"uint64\"},{\"internalType\":\"bytes32\",\"name\":\"batchHash\",\"type\":\"bytes32\"},{\"internalType\":\"uint64\",\"name\":\"indexRepeatedStorageChanges\",\"type\":\"uint64\"},{\"internalType\":\"uint256\",\"name\":\"numberOfLayer1Txs\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"priorityOperationsHash\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"l2LogsTreeRoot\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"timestamp\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"commitment\",\"type\":\"bytes32\"}],\"internalType\":\"structIExecutor.StoredBatchInfo[]\",\"name\":\"_committedBatches\",\"type\":\"tuple[]\"},{\"components\":[{\"internalType\":\"uint256[]\",\"name\":\"recursiveAggregationInput\",\"type\":\"uint256[]\"},{\"internalType\":\"uint256[]\",\"name\":\"serializedProof\",\"type\":\"uint256[]\"}],\"internalType\":\"structIExecutor.ProofInput\",\"name\":\"_proof\",\"type\":\"tuple\"}],\"name\":\"proveBatches\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_chainId\",\"type\":\"uint256\"},{\"components\":[{\"internalType\":\"uint64\",\"name\":\"batchNumber\",\"type\":\"uint64\"},{\"internalType\":\"bytes32\",\"name\":\"batchHash\",\"type\":\"bytes32\"},{\"internalType\":\"uint64\",\"name\":\"indexRepeatedStorageChanges\",\"type\":\"uint64\"},{\"internalType\":\"uint256\",\"name\":\"numberOfLayer1Txs\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"priorityOperationsHash\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"l2LogsTreeRoot\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"timestamp\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"commitment\",\"type\":\"bytes32\"}],\"internalType\":\"structIExecutor.StoredBatchInfo\",\"name\":\"_prevBatch\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"uint64\",\"name\":\"batchNumber\",\"type\":\"uint64\"},{\"internalType\":\"bytes32\",\"name\":\"batchHash\",\"type\":\"bytes32\"},{\"internalType\":\"uint64\",\"name\":\"indexRepeatedStorageChanges\",\"type\":\"uint64\"},{\"internalType\":\"uint256\",\"name\":\"numberOfLayer1Txs\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"priorityOperationsHash\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"l2LogsTreeRoot\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"timestamp\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"commitment\",\"type\":\"bytes32\"}],\"internalType\":\"structIExecutor.StoredBatchInfo[]\",\"name\":\"_committedBatches\",\"type\":\"tuple[]\"},{\"components\":[{\"internalType\":\"uint256[]\",\"name\":\"recursiveAggregationInput\",\"type\":\"uint256[]\"},{\"internalType\":\"uint256[]\",\"name\":\"serializedProof\",\"type\":\"uint256[]\"}],\"internalType\":\"structIExecutor.ProofInput\",\"name\":\"_proof\",\"type\":\"tuple\"}],\"name\":\"proveBatchesSharedBridge\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_l2TxHash\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"_l2BatchNumber\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_l2MessageIndex\",\"type\":\"uint256\"},{\"internalType\":\"uint16\",\"name\":\"_l2TxNumberInBatch\",\"type\":\"uint16\"},{\"internalType\":\"bytes32[]\",\"name\":\"_merkleProof\",\"type\":\"bytes32[]\"},{\"internalType\":\"enumTxStatus\",\"name\":\"_status\",\"type\":\"uint8\"}],\"name\":\"proveL1ToL2TransactionStatus\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_batchNumber\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_index\",\"type\":\"uint256\"},{\"components\":[{\"internalType\":\"uint8\",\"name\":\"l2ShardId\",\"type\":\"uint8\"},{\"internalType\":\"bool\",\"name\":\"isService\",\"type\":\"bool\"},{\"internalType\":\"uint16\",\"name\":\"txNumberInBatch\",\"type\":\"uint16\"},{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"key\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"value\",\"type\":\"bytes32\"}],\"internalType\":\"structL2Log\",\"name\":\"_log\",\"type\":\"tuple\"},{\"internalType\":\"bytes32[]\",\"name\":\"_proof\",\"type\":\"bytes32[]\"}],\"name\":\"proveL2LogInclusion\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_batchNumber\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_index\",\"type\":\"uint256\"},{\"components\":[{\"internalType\":\"uint16\",\"name\":\"txNumberInBatch\",\"type\":\"uint16\"},{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"internalType\":\"structL2Message\",\"name\":\"_message\",\"type\":\"tuple\"},{\"internalType\":\"bytes32[]\",\"name\":\"_proof\",\"type\":\"bytes32[]\"}],\"name\":\"proveL2MessageInclusion\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_contractL2\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_l2Value\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"_calldata\",\"type\":\"bytes\"},{\"internalType\":\"uint256\",\"name\":\"_l2GasLimit\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_l2GasPerPubdataByteLimit\",\"type\":\"uint256\"},{\"internalType\":\"bytes[]\",\"name\":\"_factoryDeps\",\"type\":\"bytes[]\"},{\"internalType\":\"address\",\"name\":\"_refundRecipient\",\"type\":\"address\"}],\"name\":\"requestL2Transaction\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"canonicalTxHash\",\"type\":\"bytes32\"}],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_newLastBatch\",\"type\":\"uint256\"}],\"name\":\"revertBatches\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_chainId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_newLastBatch\",\"type\":\"uint256\"}],\"name\":\"revertBatchesSharedBridge\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_newPendingAdmin\",\"type\":\"address\"}],\"name\":\"setPendingAdmin\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bool\",\"name\":\"_zkPorterIsAvailable\",\"type\":\"bool\"}],\"name\":\"setPorterAvailability\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_newPriorityTxMaxGasLimit\",\"type\":\"uint256\"}],\"name\":\"setPriorityTxMaxGasLimit\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint128\",\"name\":\"_nominator\",\"type\":\"uint128\"},{\"internalType\":\"uint128\",\"name\":\"_denominator\",\"type\":\"uint128\"}],\"name\":\"setTokenMultiplier\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_transactionFilterer\",\"type\":\"address\"}],\"name\":\"setTransactionFilterer\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_validator\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"_active\",\"type\":\"bool\"}],\"name\":\"setValidator\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"enumPubdataPricingMode\",\"name\":\"_validiumMode\",\"type\":\"uint8\"}],\"name\":\"setValidiumMode\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_batchNumber\",\"type\":\"uint256\"}],\"name\":\"storedBatchHash\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"transferEthToSharedBridge\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"unfreezeDiamond\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_protocolVersion\",\"type\":\"uint256\"},{\"components\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"facet\",\"type\":\"address\"},{\"internalType\":\"enumDiamond.Action\",\"name\":\"action\",\"type\":\"uint8\"},{\"internalType\":\"bool\",\"name\":\"isFreezable\",\"type\":\"bool\"},{\"internalType\":\"bytes4[]\",\"name\":\"selectors\",\"type\":\"bytes4[]\"}],\"internalType\":\"structDiamond.FacetCut[]\",\"name\":\"facetCuts\",\"type\":\"tuple[]\"},{\"internalType\":\"address\",\"name\":\"initAddress\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"initCalldata\",\"type\":\"bytes\"}],\"internalType\":\"structDiamond.DiamondCutData\",\"name\":\"_cutData\",\"type\":\"tuple\"}],\"name\":\"upgradeChainFromVersion\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
}

// IZkSyncStateTransitionABI is the input ABI used to generate the binding from.
// Deprecated: Use IZkSyncStateTransitionMetaData.ABI instead.
var IZkSyncStateTransitionABI = IZkSyncStateTransitionMetaData.ABI

// IZkSyncStateTransition is an auto generated Go binding around an Ethereum contract.
type IZkSyncStateTransition struct {
	IZkSyncStateTransitionCaller     // Read-only binding to the contract
	IZkSyncStateTransitionTransactor // Write-only binding to the contract
	IZkSyncStateTransitionFilterer   // Log filterer for contract events
}

// IZkSyncStateTransitionCaller is an auto generated read-only Go binding around an Ethereum contract.
type IZkSyncStateTransitionCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IZkSyncStateTransitionTransactor is an auto generated write-only Go binding around an Ethereum contract.
type IZkSyncStateTransitionTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IZkSyncStateTransitionFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type IZkSyncStateTransitionFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IZkSyncStateTransitionSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type IZkSyncStateTransitionSession struct {
	Contract     *IZkSyncStateTransition // Generic contract binding to set the session for
	CallOpts     bind.CallOpts           // Call options to use throughout this session
	TransactOpts bind.TransactOpts       // Transaction auth options to use throughout this session
}

// IZkSyncStateTransitionCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type IZkSyncStateTransitionCallerSession struct {
	Contract *IZkSyncStateTransitionCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts                 // Call options to use throughout this session
}

// IZkSyncStateTransitionTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type IZkSyncStateTransitionTransactorSession struct {
	Contract     *IZkSyncStateTransitionTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts                 // Transaction auth options to use throughout this session
}

// IZkSyncStateTransitionRaw is an auto generated low-level Go binding around an Ethereum contract.
type IZkSyncStateTransitionRaw struct {
	Contract *IZkSyncStateTransition // Generic contract binding to access the raw methods on
}

// IZkSyncStateTransitionCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type IZkSyncStateTransitionCallerRaw struct {
	Contract *IZkSyncStateTransitionCaller // Generic read-only contract binding to access the raw methods on
}

// IZkSyncStateTransitionTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type IZkSyncStateTransitionTransactorRaw struct {
	Contract *IZkSyncStateTransitionTransactor // Generic write-only contract binding to access the raw methods on
}

// NewIZkSyncStateTransition creates a new instance of IZkSyncStateTransition, bound to a specific deployed contract.
func NewIZkSyncStateTransition(address common.Address, backend bind.ContractBackend) (*IZkSyncStateTransition, error) {
	contract, err := bindIZkSyncStateTransition(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &IZkSyncStateTransition{IZkSyncStateTransitionCaller: IZkSyncStateTransitionCaller{contract: contract}, IZkSyncStateTransitionTransactor: IZkSyncStateTransitionTransactor{contract: contract}, IZkSyncStateTransitionFilterer: IZkSyncStateTransitionFilterer{contract: contract}}, nil
}

// NewIZkSyncStateTransitionCaller creates a new read-only instance of IZkSyncStateTransition, bound to a specific deployed contract.
func NewIZkSyncStateTransitionCaller(address common.Address, caller bind.ContractCaller) (*IZkSyncStateTransitionCaller, error) {
	contract, err := bindIZkSyncStateTransition(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &IZkSyncStateTransitionCaller{contract: contract}, nil
}

// NewIZkSyncStateTransitionTransactor creates a new write-only instance of IZkSyncStateTransition, bound to a specific deployed contract.
func NewIZkSyncStateTransitionTransactor(address common.Address, transactor bind.ContractTransactor) (*IZkSyncStateTransitionTransactor, error) {
	contract, err := bindIZkSyncStateTransition(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &IZkSyncStateTransitionTransactor{contract: contract}, nil
}

// NewIZkSyncStateTransitionFilterer creates a new log filterer instance of IZkSyncStateTransition, bound to a specific deployed contract.
func NewIZkSyncStateTransitionFilterer(address common.Address, filterer bind.ContractFilterer) (*IZkSyncStateTransitionFilterer, error) {
	contract, err := bindIZkSyncStateTransition(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &IZkSyncStateTransitionFilterer{contract: contract}, nil
}

// bindIZkSyncStateTransition binds a generic wrapper to an already deployed contract.
func bindIZkSyncStateTransition(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := IZkSyncStateTransitionMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IZkSyncStateTransition *IZkSyncStateTransitionRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IZkSyncStateTransition.Contract.IZkSyncStateTransitionCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IZkSyncStateTransition *IZkSyncStateTransitionRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IZkSyncStateTransition.Contract.IZkSyncStateTransitionTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IZkSyncStateTransition *IZkSyncStateTransitionRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IZkSyncStateTransition.Contract.IZkSyncStateTransitionTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IZkSyncStateTransition *IZkSyncStateTransitionCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IZkSyncStateTransition.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IZkSyncStateTransition *IZkSyncStateTransitionTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IZkSyncStateTransition.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IZkSyncStateTransition *IZkSyncStateTransitionTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IZkSyncStateTransition.Contract.contract.Transact(opts, method, params...)
}

// BaseTokenGasPriceMultiplierDenominator is a free data retrieval call binding the contract method 0x1de72e34.
//
// Solidity: function baseTokenGasPriceMultiplierDenominator() view returns(uint128)
func (_IZkSyncStateTransition *IZkSyncStateTransitionCaller) BaseTokenGasPriceMultiplierDenominator(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _IZkSyncStateTransition.contract.Call(opts, &out, "baseTokenGasPriceMultiplierDenominator")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BaseTokenGasPriceMultiplierDenominator is a free data retrieval call binding the contract method 0x1de72e34.
//
// Solidity: function baseTokenGasPriceMultiplierDenominator() view returns(uint128)
func (_IZkSyncStateTransition *IZkSyncStateTransitionSession) BaseTokenGasPriceMultiplierDenominator() (*big.Int, error) {
	return _IZkSyncStateTransition.Contract.BaseTokenGasPriceMultiplierDenominator(&_IZkSyncStateTransition.CallOpts)
}

// BaseTokenGasPriceMultiplierDenominator is a free data retrieval call binding the contract method 0x1de72e34.
//
// Solidity: function baseTokenGasPriceMultiplierDenominator() view returns(uint128)
func (_IZkSyncStateTransition *IZkSyncStateTransitionCallerSession) BaseTokenGasPriceMultiplierDenominator() (*big.Int, error) {
	return _IZkSyncStateTransition.Contract.BaseTokenGasPriceMultiplierDenominator(&_IZkSyncStateTransition.CallOpts)
}

// BaseTokenGasPriceMultiplierNominator is a free data retrieval call binding the contract method 0xea6c029c.
//
// Solidity: function baseTokenGasPriceMultiplierNominator() view returns(uint128)
func (_IZkSyncStateTransition *IZkSyncStateTransitionCaller) BaseTokenGasPriceMultiplierNominator(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _IZkSyncStateTransition.contract.Call(opts, &out, "baseTokenGasPriceMultiplierNominator")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BaseTokenGasPriceMultiplierNominator is a free data retrieval call binding the contract method 0xea6c029c.
//
// Solidity: function baseTokenGasPriceMultiplierNominator() view returns(uint128)
func (_IZkSyncStateTransition *IZkSyncStateTransitionSession) BaseTokenGasPriceMultiplierNominator() (*big.Int, error) {
	return _IZkSyncStateTransition.Contract.BaseTokenGasPriceMultiplierNominator(&_IZkSyncStateTransition.CallOpts)
}

// BaseTokenGasPriceMultiplierNominator is a free data retrieval call binding the contract method 0xea6c029c.
//
// Solidity: function baseTokenGasPriceMultiplierNominator() view returns(uint128)
func (_IZkSyncStateTransition *IZkSyncStateTransitionCallerSession) BaseTokenGasPriceMultiplierNominator() (*big.Int, error) {
	return _IZkSyncStateTransition.Contract.BaseTokenGasPriceMultiplierNominator(&_IZkSyncStateTransition.CallOpts)
}

// FacetAddress is a free data retrieval call binding the contract method 0xcdffacc6.
//
// Solidity: function facetAddress(bytes4 _selector) view returns(address facet)
func (_IZkSyncStateTransition *IZkSyncStateTransitionCaller) FacetAddress(opts *bind.CallOpts, _selector [4]byte) (common.Address, error) {
	var out []interface{}
	err := _IZkSyncStateTransition.contract.Call(opts, &out, "facetAddress", _selector)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// FacetAddress is a free data retrieval call binding the contract method 0xcdffacc6.
//
// Solidity: function facetAddress(bytes4 _selector) view returns(address facet)
func (_IZkSyncStateTransition *IZkSyncStateTransitionSession) FacetAddress(_selector [4]byte) (common.Address, error) {
	return _IZkSyncStateTransition.Contract.FacetAddress(&_IZkSyncStateTransition.CallOpts, _selector)
}

// FacetAddress is a free data retrieval call binding the contract method 0xcdffacc6.
//
// Solidity: function facetAddress(bytes4 _selector) view returns(address facet)
func (_IZkSyncStateTransition *IZkSyncStateTransitionCallerSession) FacetAddress(_selector [4]byte) (common.Address, error) {
	return _IZkSyncStateTransition.Contract.FacetAddress(&_IZkSyncStateTransition.CallOpts, _selector)
}

// FacetAddresses is a free data retrieval call binding the contract method 0x52ef6b2c.
//
// Solidity: function facetAddresses() view returns(address[] facets)
func (_IZkSyncStateTransition *IZkSyncStateTransitionCaller) FacetAddresses(opts *bind.CallOpts) ([]common.Address, error) {
	var out []interface{}
	err := _IZkSyncStateTransition.contract.Call(opts, &out, "facetAddresses")

	if err != nil {
		return *new([]common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)

	return out0, err

}

// FacetAddresses is a free data retrieval call binding the contract method 0x52ef6b2c.
//
// Solidity: function facetAddresses() view returns(address[] facets)
func (_IZkSyncStateTransition *IZkSyncStateTransitionSession) FacetAddresses() ([]common.Address, error) {
	return _IZkSyncStateTransition.Contract.FacetAddresses(&_IZkSyncStateTransition.CallOpts)
}

// FacetAddresses is a free data retrieval call binding the contract method 0x52ef6b2c.
//
// Solidity: function facetAddresses() view returns(address[] facets)
func (_IZkSyncStateTransition *IZkSyncStateTransitionCallerSession) FacetAddresses() ([]common.Address, error) {
	return _IZkSyncStateTransition.Contract.FacetAddresses(&_IZkSyncStateTransition.CallOpts)
}

// FacetFunctionSelectors is a free data retrieval call binding the contract method 0xadfca15e.
//
// Solidity: function facetFunctionSelectors(address _facet) view returns(bytes4[])
func (_IZkSyncStateTransition *IZkSyncStateTransitionCaller) FacetFunctionSelectors(opts *bind.CallOpts, _facet common.Address) ([][4]byte, error) {
	var out []interface{}
	err := _IZkSyncStateTransition.contract.Call(opts, &out, "facetFunctionSelectors", _facet)

	if err != nil {
		return *new([][4]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([][4]byte)).(*[][4]byte)

	return out0, err

}

// FacetFunctionSelectors is a free data retrieval call binding the contract method 0xadfca15e.
//
// Solidity: function facetFunctionSelectors(address _facet) view returns(bytes4[])
func (_IZkSyncStateTransition *IZkSyncStateTransitionSession) FacetFunctionSelectors(_facet common.Address) ([][4]byte, error) {
	return _IZkSyncStateTransition.Contract.FacetFunctionSelectors(&_IZkSyncStateTransition.CallOpts, _facet)
}

// FacetFunctionSelectors is a free data retrieval call binding the contract method 0xadfca15e.
//
// Solidity: function facetFunctionSelectors(address _facet) view returns(bytes4[])
func (_IZkSyncStateTransition *IZkSyncStateTransitionCallerSession) FacetFunctionSelectors(_facet common.Address) ([][4]byte, error) {
	return _IZkSyncStateTransition.Contract.FacetFunctionSelectors(&_IZkSyncStateTransition.CallOpts, _facet)
}

// Facets is a free data retrieval call binding the contract method 0x7a0ed627.
//
// Solidity: function facets() view returns((address,bytes4[])[])
func (_IZkSyncStateTransition *IZkSyncStateTransitionCaller) Facets(opts *bind.CallOpts) ([]IGettersFacet, error) {
	var out []interface{}
	err := _IZkSyncStateTransition.contract.Call(opts, &out, "facets")

	if err != nil {
		return *new([]IGettersFacet), err
	}

	out0 := *abi.ConvertType(out[0], new([]IGettersFacet)).(*[]IGettersFacet)

	return out0, err

}

// Facets is a free data retrieval call binding the contract method 0x7a0ed627.
//
// Solidity: function facets() view returns((address,bytes4[])[])
func (_IZkSyncStateTransition *IZkSyncStateTransitionSession) Facets() ([]IGettersFacet, error) {
	return _IZkSyncStateTransition.Contract.Facets(&_IZkSyncStateTransition.CallOpts)
}

// Facets is a free data retrieval call binding the contract method 0x7a0ed627.
//
// Solidity: function facets() view returns((address,bytes4[])[])
func (_IZkSyncStateTransition *IZkSyncStateTransitionCallerSession) Facets() ([]IGettersFacet, error) {
	return _IZkSyncStateTransition.Contract.Facets(&_IZkSyncStateTransition.CallOpts)
}

// GetAdmin is a free data retrieval call binding the contract method 0x6e9960c3.
//
// Solidity: function getAdmin() view returns(address)
func (_IZkSyncStateTransition *IZkSyncStateTransitionCaller) GetAdmin(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _IZkSyncStateTransition.contract.Call(opts, &out, "getAdmin")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetAdmin is a free data retrieval call binding the contract method 0x6e9960c3.
//
// Solidity: function getAdmin() view returns(address)
func (_IZkSyncStateTransition *IZkSyncStateTransitionSession) GetAdmin() (common.Address, error) {
	return _IZkSyncStateTransition.Contract.GetAdmin(&_IZkSyncStateTransition.CallOpts)
}

// GetAdmin is a free data retrieval call binding the contract method 0x6e9960c3.
//
// Solidity: function getAdmin() view returns(address)
func (_IZkSyncStateTransition *IZkSyncStateTransitionCallerSession) GetAdmin() (common.Address, error) {
	return _IZkSyncStateTransition.Contract.GetAdmin(&_IZkSyncStateTransition.CallOpts)
}

// GetBaseToken is a free data retrieval call binding the contract method 0x98acd7a6.
//
// Solidity: function getBaseToken() view returns(address)
func (_IZkSyncStateTransition *IZkSyncStateTransitionCaller) GetBaseToken(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _IZkSyncStateTransition.contract.Call(opts, &out, "getBaseToken")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetBaseToken is a free data retrieval call binding the contract method 0x98acd7a6.
//
// Solidity: function getBaseToken() view returns(address)
func (_IZkSyncStateTransition *IZkSyncStateTransitionSession) GetBaseToken() (common.Address, error) {
	return _IZkSyncStateTransition.Contract.GetBaseToken(&_IZkSyncStateTransition.CallOpts)
}

// GetBaseToken is a free data retrieval call binding the contract method 0x98acd7a6.
//
// Solidity: function getBaseToken() view returns(address)
func (_IZkSyncStateTransition *IZkSyncStateTransitionCallerSession) GetBaseToken() (common.Address, error) {
	return _IZkSyncStateTransition.Contract.GetBaseToken(&_IZkSyncStateTransition.CallOpts)
}

// GetBaseTokenBridge is a free data retrieval call binding the contract method 0x086a56f8.
//
// Solidity: function getBaseTokenBridge() view returns(address)
func (_IZkSyncStateTransition *IZkSyncStateTransitionCaller) GetBaseTokenBridge(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _IZkSyncStateTransition.contract.Call(opts, &out, "getBaseTokenBridge")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetBaseTokenBridge is a free data retrieval call binding the contract method 0x086a56f8.
//
// Solidity: function getBaseTokenBridge() view returns(address)
func (_IZkSyncStateTransition *IZkSyncStateTransitionSession) GetBaseTokenBridge() (common.Address, error) {
	return _IZkSyncStateTransition.Contract.GetBaseTokenBridge(&_IZkSyncStateTransition.CallOpts)
}

// GetBaseTokenBridge is a free data retrieval call binding the contract method 0x086a56f8.
//
// Solidity: function getBaseTokenBridge() view returns(address)
func (_IZkSyncStateTransition *IZkSyncStateTransitionCallerSession) GetBaseTokenBridge() (common.Address, error) {
	return _IZkSyncStateTransition.Contract.GetBaseTokenBridge(&_IZkSyncStateTransition.CallOpts)
}

// GetBridgehub is a free data retrieval call binding the contract method 0x3591c1a0.
//
// Solidity: function getBridgehub() view returns(address)
func (_IZkSyncStateTransition *IZkSyncStateTransitionCaller) GetBridgehub(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _IZkSyncStateTransition.contract.Call(opts, &out, "getBridgehub")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetBridgehub is a free data retrieval call binding the contract method 0x3591c1a0.
//
// Solidity: function getBridgehub() view returns(address)
func (_IZkSyncStateTransition *IZkSyncStateTransitionSession) GetBridgehub() (common.Address, error) {
	return _IZkSyncStateTransition.Contract.GetBridgehub(&_IZkSyncStateTransition.CallOpts)
}

// GetBridgehub is a free data retrieval call binding the contract method 0x3591c1a0.
//
// Solidity: function getBridgehub() view returns(address)
func (_IZkSyncStateTransition *IZkSyncStateTransitionCallerSession) GetBridgehub() (common.Address, error) {
	return _IZkSyncStateTransition.Contract.GetBridgehub(&_IZkSyncStateTransition.CallOpts)
}

// GetFirstUnprocessedPriorityTx is a free data retrieval call binding the contract method 0x79823c9a.
//
// Solidity: function getFirstUnprocessedPriorityTx() view returns(uint256)
func (_IZkSyncStateTransition *IZkSyncStateTransitionCaller) GetFirstUnprocessedPriorityTx(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _IZkSyncStateTransition.contract.Call(opts, &out, "getFirstUnprocessedPriorityTx")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetFirstUnprocessedPriorityTx is a free data retrieval call binding the contract method 0x79823c9a.
//
// Solidity: function getFirstUnprocessedPriorityTx() view returns(uint256)
func (_IZkSyncStateTransition *IZkSyncStateTransitionSession) GetFirstUnprocessedPriorityTx() (*big.Int, error) {
	return _IZkSyncStateTransition.Contract.GetFirstUnprocessedPriorityTx(&_IZkSyncStateTransition.CallOpts)
}

// GetFirstUnprocessedPriorityTx is a free data retrieval call binding the contract method 0x79823c9a.
//
// Solidity: function getFirstUnprocessedPriorityTx() view returns(uint256)
func (_IZkSyncStateTransition *IZkSyncStateTransitionCallerSession) GetFirstUnprocessedPriorityTx() (*big.Int, error) {
	return _IZkSyncStateTransition.Contract.GetFirstUnprocessedPriorityTx(&_IZkSyncStateTransition.CallOpts)
}

// GetL2BootloaderBytecodeHash is a free data retrieval call binding the contract method 0xd86970d8.
//
// Solidity: function getL2BootloaderBytecodeHash() view returns(bytes32)
func (_IZkSyncStateTransition *IZkSyncStateTransitionCaller) GetL2BootloaderBytecodeHash(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _IZkSyncStateTransition.contract.Call(opts, &out, "getL2BootloaderBytecodeHash")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// GetL2BootloaderBytecodeHash is a free data retrieval call binding the contract method 0xd86970d8.
//
// Solidity: function getL2BootloaderBytecodeHash() view returns(bytes32)
func (_IZkSyncStateTransition *IZkSyncStateTransitionSession) GetL2BootloaderBytecodeHash() ([32]byte, error) {
	return _IZkSyncStateTransition.Contract.GetL2BootloaderBytecodeHash(&_IZkSyncStateTransition.CallOpts)
}

// GetL2BootloaderBytecodeHash is a free data retrieval call binding the contract method 0xd86970d8.
//
// Solidity: function getL2BootloaderBytecodeHash() view returns(bytes32)
func (_IZkSyncStateTransition *IZkSyncStateTransitionCallerSession) GetL2BootloaderBytecodeHash() ([32]byte, error) {
	return _IZkSyncStateTransition.Contract.GetL2BootloaderBytecodeHash(&_IZkSyncStateTransition.CallOpts)
}

// GetL2DefaultAccountBytecodeHash is a free data retrieval call binding the contract method 0xfd791f3c.
//
// Solidity: function getL2DefaultAccountBytecodeHash() view returns(bytes32)
func (_IZkSyncStateTransition *IZkSyncStateTransitionCaller) GetL2DefaultAccountBytecodeHash(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _IZkSyncStateTransition.contract.Call(opts, &out, "getL2DefaultAccountBytecodeHash")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// GetL2DefaultAccountBytecodeHash is a free data retrieval call binding the contract method 0xfd791f3c.
//
// Solidity: function getL2DefaultAccountBytecodeHash() view returns(bytes32)
func (_IZkSyncStateTransition *IZkSyncStateTransitionSession) GetL2DefaultAccountBytecodeHash() ([32]byte, error) {
	return _IZkSyncStateTransition.Contract.GetL2DefaultAccountBytecodeHash(&_IZkSyncStateTransition.CallOpts)
}

// GetL2DefaultAccountBytecodeHash is a free data retrieval call binding the contract method 0xfd791f3c.
//
// Solidity: function getL2DefaultAccountBytecodeHash() view returns(bytes32)
func (_IZkSyncStateTransition *IZkSyncStateTransitionCallerSession) GetL2DefaultAccountBytecodeHash() ([32]byte, error) {
	return _IZkSyncStateTransition.Contract.GetL2DefaultAccountBytecodeHash(&_IZkSyncStateTransition.CallOpts)
}

// GetL2SystemContractsUpgradeBatchNumber is a free data retrieval call binding the contract method 0xe5355c75.
//
// Solidity: function getL2SystemContractsUpgradeBatchNumber() view returns(uint256)
func (_IZkSyncStateTransition *IZkSyncStateTransitionCaller) GetL2SystemContractsUpgradeBatchNumber(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _IZkSyncStateTransition.contract.Call(opts, &out, "getL2SystemContractsUpgradeBatchNumber")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetL2SystemContractsUpgradeBatchNumber is a free data retrieval call binding the contract method 0xe5355c75.
//
// Solidity: function getL2SystemContractsUpgradeBatchNumber() view returns(uint256)
func (_IZkSyncStateTransition *IZkSyncStateTransitionSession) GetL2SystemContractsUpgradeBatchNumber() (*big.Int, error) {
	return _IZkSyncStateTransition.Contract.GetL2SystemContractsUpgradeBatchNumber(&_IZkSyncStateTransition.CallOpts)
}

// GetL2SystemContractsUpgradeBatchNumber is a free data retrieval call binding the contract method 0xe5355c75.
//
// Solidity: function getL2SystemContractsUpgradeBatchNumber() view returns(uint256)
func (_IZkSyncStateTransition *IZkSyncStateTransitionCallerSession) GetL2SystemContractsUpgradeBatchNumber() (*big.Int, error) {
	return _IZkSyncStateTransition.Contract.GetL2SystemContractsUpgradeBatchNumber(&_IZkSyncStateTransition.CallOpts)
}

// GetL2SystemContractsUpgradeTxHash is a free data retrieval call binding the contract method 0x7b30c8da.
//
// Solidity: function getL2SystemContractsUpgradeTxHash() view returns(bytes32)
func (_IZkSyncStateTransition *IZkSyncStateTransitionCaller) GetL2SystemContractsUpgradeTxHash(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _IZkSyncStateTransition.contract.Call(opts, &out, "getL2SystemContractsUpgradeTxHash")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// GetL2SystemContractsUpgradeTxHash is a free data retrieval call binding the contract method 0x7b30c8da.
//
// Solidity: function getL2SystemContractsUpgradeTxHash() view returns(bytes32)
func (_IZkSyncStateTransition *IZkSyncStateTransitionSession) GetL2SystemContractsUpgradeTxHash() ([32]byte, error) {
	return _IZkSyncStateTransition.Contract.GetL2SystemContractsUpgradeTxHash(&_IZkSyncStateTransition.CallOpts)
}

// GetL2SystemContractsUpgradeTxHash is a free data retrieval call binding the contract method 0x7b30c8da.
//
// Solidity: function getL2SystemContractsUpgradeTxHash() view returns(bytes32)
func (_IZkSyncStateTransition *IZkSyncStateTransitionCallerSession) GetL2SystemContractsUpgradeTxHash() ([32]byte, error) {
	return _IZkSyncStateTransition.Contract.GetL2SystemContractsUpgradeTxHash(&_IZkSyncStateTransition.CallOpts)
}

// GetName is a free data retrieval call binding the contract method 0x17d7de7c.
//
// Solidity: function getName() view returns(string)
func (_IZkSyncStateTransition *IZkSyncStateTransitionCaller) GetName(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _IZkSyncStateTransition.contract.Call(opts, &out, "getName")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// GetName is a free data retrieval call binding the contract method 0x17d7de7c.
//
// Solidity: function getName() view returns(string)
func (_IZkSyncStateTransition *IZkSyncStateTransitionSession) GetName() (string, error) {
	return _IZkSyncStateTransition.Contract.GetName(&_IZkSyncStateTransition.CallOpts)
}

// GetName is a free data retrieval call binding the contract method 0x17d7de7c.
//
// Solidity: function getName() view returns(string)
func (_IZkSyncStateTransition *IZkSyncStateTransitionCallerSession) GetName() (string, error) {
	return _IZkSyncStateTransition.Contract.GetName(&_IZkSyncStateTransition.CallOpts)
}

// GetPendingAdmin is a free data retrieval call binding the contract method 0xd0468156.
//
// Solidity: function getPendingAdmin() view returns(address)
func (_IZkSyncStateTransition *IZkSyncStateTransitionCaller) GetPendingAdmin(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _IZkSyncStateTransition.contract.Call(opts, &out, "getPendingAdmin")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetPendingAdmin is a free data retrieval call binding the contract method 0xd0468156.
//
// Solidity: function getPendingAdmin() view returns(address)
func (_IZkSyncStateTransition *IZkSyncStateTransitionSession) GetPendingAdmin() (common.Address, error) {
	return _IZkSyncStateTransition.Contract.GetPendingAdmin(&_IZkSyncStateTransition.CallOpts)
}

// GetPendingAdmin is a free data retrieval call binding the contract method 0xd0468156.
//
// Solidity: function getPendingAdmin() view returns(address)
func (_IZkSyncStateTransition *IZkSyncStateTransitionCallerSession) GetPendingAdmin() (common.Address, error) {
	return _IZkSyncStateTransition.Contract.GetPendingAdmin(&_IZkSyncStateTransition.CallOpts)
}

// GetPriorityQueueSize is a free data retrieval call binding the contract method 0x631f4bac.
//
// Solidity: function getPriorityQueueSize() view returns(uint256)
func (_IZkSyncStateTransition *IZkSyncStateTransitionCaller) GetPriorityQueueSize(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _IZkSyncStateTransition.contract.Call(opts, &out, "getPriorityQueueSize")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetPriorityQueueSize is a free data retrieval call binding the contract method 0x631f4bac.
//
// Solidity: function getPriorityQueueSize() view returns(uint256)
func (_IZkSyncStateTransition *IZkSyncStateTransitionSession) GetPriorityQueueSize() (*big.Int, error) {
	return _IZkSyncStateTransition.Contract.GetPriorityQueueSize(&_IZkSyncStateTransition.CallOpts)
}

// GetPriorityQueueSize is a free data retrieval call binding the contract method 0x631f4bac.
//
// Solidity: function getPriorityQueueSize() view returns(uint256)
func (_IZkSyncStateTransition *IZkSyncStateTransitionCallerSession) GetPriorityQueueSize() (*big.Int, error) {
	return _IZkSyncStateTransition.Contract.GetPriorityQueueSize(&_IZkSyncStateTransition.CallOpts)
}

// GetPriorityTxMaxGasLimit is a free data retrieval call binding the contract method 0x0ec6b0b7.
//
// Solidity: function getPriorityTxMaxGasLimit() view returns(uint256)
func (_IZkSyncStateTransition *IZkSyncStateTransitionCaller) GetPriorityTxMaxGasLimit(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _IZkSyncStateTransition.contract.Call(opts, &out, "getPriorityTxMaxGasLimit")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetPriorityTxMaxGasLimit is a free data retrieval call binding the contract method 0x0ec6b0b7.
//
// Solidity: function getPriorityTxMaxGasLimit() view returns(uint256)
func (_IZkSyncStateTransition *IZkSyncStateTransitionSession) GetPriorityTxMaxGasLimit() (*big.Int, error) {
	return _IZkSyncStateTransition.Contract.GetPriorityTxMaxGasLimit(&_IZkSyncStateTransition.CallOpts)
}

// GetPriorityTxMaxGasLimit is a free data retrieval call binding the contract method 0x0ec6b0b7.
//
// Solidity: function getPriorityTxMaxGasLimit() view returns(uint256)
func (_IZkSyncStateTransition *IZkSyncStateTransitionCallerSession) GetPriorityTxMaxGasLimit() (*big.Int, error) {
	return _IZkSyncStateTransition.Contract.GetPriorityTxMaxGasLimit(&_IZkSyncStateTransition.CallOpts)
}

// GetProtocolVersion is a free data retrieval call binding the contract method 0x33ce93fe.
//
// Solidity: function getProtocolVersion() view returns(uint256)
func (_IZkSyncStateTransition *IZkSyncStateTransitionCaller) GetProtocolVersion(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _IZkSyncStateTransition.contract.Call(opts, &out, "getProtocolVersion")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetProtocolVersion is a free data retrieval call binding the contract method 0x33ce93fe.
//
// Solidity: function getProtocolVersion() view returns(uint256)
func (_IZkSyncStateTransition *IZkSyncStateTransitionSession) GetProtocolVersion() (*big.Int, error) {
	return _IZkSyncStateTransition.Contract.GetProtocolVersion(&_IZkSyncStateTransition.CallOpts)
}

// GetProtocolVersion is a free data retrieval call binding the contract method 0x33ce93fe.
//
// Solidity: function getProtocolVersion() view returns(uint256)
func (_IZkSyncStateTransition *IZkSyncStateTransitionCallerSession) GetProtocolVersion() (*big.Int, error) {
	return _IZkSyncStateTransition.Contract.GetProtocolVersion(&_IZkSyncStateTransition.CallOpts)
}

// GetPubdataPricingMode is a free data retrieval call binding the contract method 0x06d49e5b.
//
// Solidity: function getPubdataPricingMode() view returns(uint8)
func (_IZkSyncStateTransition *IZkSyncStateTransitionCaller) GetPubdataPricingMode(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _IZkSyncStateTransition.contract.Call(opts, &out, "getPubdataPricingMode")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// GetPubdataPricingMode is a free data retrieval call binding the contract method 0x06d49e5b.
//
// Solidity: function getPubdataPricingMode() view returns(uint8)
func (_IZkSyncStateTransition *IZkSyncStateTransitionSession) GetPubdataPricingMode() (uint8, error) {
	return _IZkSyncStateTransition.Contract.GetPubdataPricingMode(&_IZkSyncStateTransition.CallOpts)
}

// GetPubdataPricingMode is a free data retrieval call binding the contract method 0x06d49e5b.
//
// Solidity: function getPubdataPricingMode() view returns(uint8)
func (_IZkSyncStateTransition *IZkSyncStateTransitionCallerSession) GetPubdataPricingMode() (uint8, error) {
	return _IZkSyncStateTransition.Contract.GetPubdataPricingMode(&_IZkSyncStateTransition.CallOpts)
}

// GetStateTransitionManager is a free data retrieval call binding the contract method 0x5518c73b.
//
// Solidity: function getStateTransitionManager() view returns(address)
func (_IZkSyncStateTransition *IZkSyncStateTransitionCaller) GetStateTransitionManager(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _IZkSyncStateTransition.contract.Call(opts, &out, "getStateTransitionManager")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetStateTransitionManager is a free data retrieval call binding the contract method 0x5518c73b.
//
// Solidity: function getStateTransitionManager() view returns(address)
func (_IZkSyncStateTransition *IZkSyncStateTransitionSession) GetStateTransitionManager() (common.Address, error) {
	return _IZkSyncStateTransition.Contract.GetStateTransitionManager(&_IZkSyncStateTransition.CallOpts)
}

// GetStateTransitionManager is a free data retrieval call binding the contract method 0x5518c73b.
//
// Solidity: function getStateTransitionManager() view returns(address)
func (_IZkSyncStateTransition *IZkSyncStateTransitionCallerSession) GetStateTransitionManager() (common.Address, error) {
	return _IZkSyncStateTransition.Contract.GetStateTransitionManager(&_IZkSyncStateTransition.CallOpts)
}

// GetTotalBatchesCommitted is a free data retrieval call binding the contract method 0xdb1f0bf9.
//
// Solidity: function getTotalBatchesCommitted() view returns(uint256)
func (_IZkSyncStateTransition *IZkSyncStateTransitionCaller) GetTotalBatchesCommitted(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _IZkSyncStateTransition.contract.Call(opts, &out, "getTotalBatchesCommitted")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetTotalBatchesCommitted is a free data retrieval call binding the contract method 0xdb1f0bf9.
//
// Solidity: function getTotalBatchesCommitted() view returns(uint256)
func (_IZkSyncStateTransition *IZkSyncStateTransitionSession) GetTotalBatchesCommitted() (*big.Int, error) {
	return _IZkSyncStateTransition.Contract.GetTotalBatchesCommitted(&_IZkSyncStateTransition.CallOpts)
}

// GetTotalBatchesCommitted is a free data retrieval call binding the contract method 0xdb1f0bf9.
//
// Solidity: function getTotalBatchesCommitted() view returns(uint256)
func (_IZkSyncStateTransition *IZkSyncStateTransitionCallerSession) GetTotalBatchesCommitted() (*big.Int, error) {
	return _IZkSyncStateTransition.Contract.GetTotalBatchesCommitted(&_IZkSyncStateTransition.CallOpts)
}

// GetTotalBatchesExecuted is a free data retrieval call binding the contract method 0xb8c2f66f.
//
// Solidity: function getTotalBatchesExecuted() view returns(uint256)
func (_IZkSyncStateTransition *IZkSyncStateTransitionCaller) GetTotalBatchesExecuted(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _IZkSyncStateTransition.contract.Call(opts, &out, "getTotalBatchesExecuted")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetTotalBatchesExecuted is a free data retrieval call binding the contract method 0xb8c2f66f.
//
// Solidity: function getTotalBatchesExecuted() view returns(uint256)
func (_IZkSyncStateTransition *IZkSyncStateTransitionSession) GetTotalBatchesExecuted() (*big.Int, error) {
	return _IZkSyncStateTransition.Contract.GetTotalBatchesExecuted(&_IZkSyncStateTransition.CallOpts)
}

// GetTotalBatchesExecuted is a free data retrieval call binding the contract method 0xb8c2f66f.
//
// Solidity: function getTotalBatchesExecuted() view returns(uint256)
func (_IZkSyncStateTransition *IZkSyncStateTransitionCallerSession) GetTotalBatchesExecuted() (*big.Int, error) {
	return _IZkSyncStateTransition.Contract.GetTotalBatchesExecuted(&_IZkSyncStateTransition.CallOpts)
}

// GetTotalBatchesVerified is a free data retrieval call binding the contract method 0xef3f0bae.
//
// Solidity: function getTotalBatchesVerified() view returns(uint256)
func (_IZkSyncStateTransition *IZkSyncStateTransitionCaller) GetTotalBatchesVerified(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _IZkSyncStateTransition.contract.Call(opts, &out, "getTotalBatchesVerified")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetTotalBatchesVerified is a free data retrieval call binding the contract method 0xef3f0bae.
//
// Solidity: function getTotalBatchesVerified() view returns(uint256)
func (_IZkSyncStateTransition *IZkSyncStateTransitionSession) GetTotalBatchesVerified() (*big.Int, error) {
	return _IZkSyncStateTransition.Contract.GetTotalBatchesVerified(&_IZkSyncStateTransition.CallOpts)
}

// GetTotalBatchesVerified is a free data retrieval call binding the contract method 0xef3f0bae.
//
// Solidity: function getTotalBatchesVerified() view returns(uint256)
func (_IZkSyncStateTransition *IZkSyncStateTransitionCallerSession) GetTotalBatchesVerified() (*big.Int, error) {
	return _IZkSyncStateTransition.Contract.GetTotalBatchesVerified(&_IZkSyncStateTransition.CallOpts)
}

// GetTotalPriorityTxs is a free data retrieval call binding the contract method 0xa1954fc5.
//
// Solidity: function getTotalPriorityTxs() view returns(uint256)
func (_IZkSyncStateTransition *IZkSyncStateTransitionCaller) GetTotalPriorityTxs(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _IZkSyncStateTransition.contract.Call(opts, &out, "getTotalPriorityTxs")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetTotalPriorityTxs is a free data retrieval call binding the contract method 0xa1954fc5.
//
// Solidity: function getTotalPriorityTxs() view returns(uint256)
func (_IZkSyncStateTransition *IZkSyncStateTransitionSession) GetTotalPriorityTxs() (*big.Int, error) {
	return _IZkSyncStateTransition.Contract.GetTotalPriorityTxs(&_IZkSyncStateTransition.CallOpts)
}

// GetTotalPriorityTxs is a free data retrieval call binding the contract method 0xa1954fc5.
//
// Solidity: function getTotalPriorityTxs() view returns(uint256)
func (_IZkSyncStateTransition *IZkSyncStateTransitionCallerSession) GetTotalPriorityTxs() (*big.Int, error) {
	return _IZkSyncStateTransition.Contract.GetTotalPriorityTxs(&_IZkSyncStateTransition.CallOpts)
}

// GetVerifier is a free data retrieval call binding the contract method 0x46657fe9.
//
// Solidity: function getVerifier() view returns(address)
func (_IZkSyncStateTransition *IZkSyncStateTransitionCaller) GetVerifier(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _IZkSyncStateTransition.contract.Call(opts, &out, "getVerifier")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetVerifier is a free data retrieval call binding the contract method 0x46657fe9.
//
// Solidity: function getVerifier() view returns(address)
func (_IZkSyncStateTransition *IZkSyncStateTransitionSession) GetVerifier() (common.Address, error) {
	return _IZkSyncStateTransition.Contract.GetVerifier(&_IZkSyncStateTransition.CallOpts)
}

// GetVerifier is a free data retrieval call binding the contract method 0x46657fe9.
//
// Solidity: function getVerifier() view returns(address)
func (_IZkSyncStateTransition *IZkSyncStateTransitionCallerSession) GetVerifier() (common.Address, error) {
	return _IZkSyncStateTransition.Contract.GetVerifier(&_IZkSyncStateTransition.CallOpts)
}

// GetVerifierParams is a free data retrieval call binding the contract method 0x18e3a941.
//
// Solidity: function getVerifierParams() view returns((bytes32,bytes32,bytes32))
func (_IZkSyncStateTransition *IZkSyncStateTransitionCaller) GetVerifierParams(opts *bind.CallOpts) (VerifierParams, error) {
	var out []interface{}
	err := _IZkSyncStateTransition.contract.Call(opts, &out, "getVerifierParams")

	if err != nil {
		return *new(VerifierParams), err
	}

	out0 := *abi.ConvertType(out[0], new(VerifierParams)).(*VerifierParams)

	return out0, err

}

// GetVerifierParams is a free data retrieval call binding the contract method 0x18e3a941.
//
// Solidity: function getVerifierParams() view returns((bytes32,bytes32,bytes32))
func (_IZkSyncStateTransition *IZkSyncStateTransitionSession) GetVerifierParams() (VerifierParams, error) {
	return _IZkSyncStateTransition.Contract.GetVerifierParams(&_IZkSyncStateTransition.CallOpts)
}

// GetVerifierParams is a free data retrieval call binding the contract method 0x18e3a941.
//
// Solidity: function getVerifierParams() view returns((bytes32,bytes32,bytes32))
func (_IZkSyncStateTransition *IZkSyncStateTransitionCallerSession) GetVerifierParams() (VerifierParams, error) {
	return _IZkSyncStateTransition.Contract.GetVerifierParams(&_IZkSyncStateTransition.CallOpts)
}

// IsDiamondStorageFrozen is a free data retrieval call binding the contract method 0x29b98c67.
//
// Solidity: function isDiamondStorageFrozen() view returns(bool)
func (_IZkSyncStateTransition *IZkSyncStateTransitionCaller) IsDiamondStorageFrozen(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _IZkSyncStateTransition.contract.Call(opts, &out, "isDiamondStorageFrozen")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsDiamondStorageFrozen is a free data retrieval call binding the contract method 0x29b98c67.
//
// Solidity: function isDiamondStorageFrozen() view returns(bool)
func (_IZkSyncStateTransition *IZkSyncStateTransitionSession) IsDiamondStorageFrozen() (bool, error) {
	return _IZkSyncStateTransition.Contract.IsDiamondStorageFrozen(&_IZkSyncStateTransition.CallOpts)
}

// IsDiamondStorageFrozen is a free data retrieval call binding the contract method 0x29b98c67.
//
// Solidity: function isDiamondStorageFrozen() view returns(bool)
func (_IZkSyncStateTransition *IZkSyncStateTransitionCallerSession) IsDiamondStorageFrozen() (bool, error) {
	return _IZkSyncStateTransition.Contract.IsDiamondStorageFrozen(&_IZkSyncStateTransition.CallOpts)
}

// IsEthWithdrawalFinalized is a free data retrieval call binding the contract method 0xbd7c5412.
//
// Solidity: function isEthWithdrawalFinalized(uint256 _l2BatchNumber, uint256 _l2MessageIndex) view returns(bool)
func (_IZkSyncStateTransition *IZkSyncStateTransitionCaller) IsEthWithdrawalFinalized(opts *bind.CallOpts, _l2BatchNumber *big.Int, _l2MessageIndex *big.Int) (bool, error) {
	var out []interface{}
	err := _IZkSyncStateTransition.contract.Call(opts, &out, "isEthWithdrawalFinalized", _l2BatchNumber, _l2MessageIndex)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsEthWithdrawalFinalized is a free data retrieval call binding the contract method 0xbd7c5412.
//
// Solidity: function isEthWithdrawalFinalized(uint256 _l2BatchNumber, uint256 _l2MessageIndex) view returns(bool)
func (_IZkSyncStateTransition *IZkSyncStateTransitionSession) IsEthWithdrawalFinalized(_l2BatchNumber *big.Int, _l2MessageIndex *big.Int) (bool, error) {
	return _IZkSyncStateTransition.Contract.IsEthWithdrawalFinalized(&_IZkSyncStateTransition.CallOpts, _l2BatchNumber, _l2MessageIndex)
}

// IsEthWithdrawalFinalized is a free data retrieval call binding the contract method 0xbd7c5412.
//
// Solidity: function isEthWithdrawalFinalized(uint256 _l2BatchNumber, uint256 _l2MessageIndex) view returns(bool)
func (_IZkSyncStateTransition *IZkSyncStateTransitionCallerSession) IsEthWithdrawalFinalized(_l2BatchNumber *big.Int, _l2MessageIndex *big.Int) (bool, error) {
	return _IZkSyncStateTransition.Contract.IsEthWithdrawalFinalized(&_IZkSyncStateTransition.CallOpts, _l2BatchNumber, _l2MessageIndex)
}

// IsFacetFreezable is a free data retrieval call binding the contract method 0xc3bbd2d7.
//
// Solidity: function isFacetFreezable(address _facet) view returns(bool isFreezable)
func (_IZkSyncStateTransition *IZkSyncStateTransitionCaller) IsFacetFreezable(opts *bind.CallOpts, _facet common.Address) (bool, error) {
	var out []interface{}
	err := _IZkSyncStateTransition.contract.Call(opts, &out, "isFacetFreezable", _facet)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsFacetFreezable is a free data retrieval call binding the contract method 0xc3bbd2d7.
//
// Solidity: function isFacetFreezable(address _facet) view returns(bool isFreezable)
func (_IZkSyncStateTransition *IZkSyncStateTransitionSession) IsFacetFreezable(_facet common.Address) (bool, error) {
	return _IZkSyncStateTransition.Contract.IsFacetFreezable(&_IZkSyncStateTransition.CallOpts, _facet)
}

// IsFacetFreezable is a free data retrieval call binding the contract method 0xc3bbd2d7.
//
// Solidity: function isFacetFreezable(address _facet) view returns(bool isFreezable)
func (_IZkSyncStateTransition *IZkSyncStateTransitionCallerSession) IsFacetFreezable(_facet common.Address) (bool, error) {
	return _IZkSyncStateTransition.Contract.IsFacetFreezable(&_IZkSyncStateTransition.CallOpts, _facet)
}

// IsFunctionFreezable is a free data retrieval call binding the contract method 0xe81e0ba1.
//
// Solidity: function isFunctionFreezable(bytes4 _selector) view returns(bool)
func (_IZkSyncStateTransition *IZkSyncStateTransitionCaller) IsFunctionFreezable(opts *bind.CallOpts, _selector [4]byte) (bool, error) {
	var out []interface{}
	err := _IZkSyncStateTransition.contract.Call(opts, &out, "isFunctionFreezable", _selector)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsFunctionFreezable is a free data retrieval call binding the contract method 0xe81e0ba1.
//
// Solidity: function isFunctionFreezable(bytes4 _selector) view returns(bool)
func (_IZkSyncStateTransition *IZkSyncStateTransitionSession) IsFunctionFreezable(_selector [4]byte) (bool, error) {
	return _IZkSyncStateTransition.Contract.IsFunctionFreezable(&_IZkSyncStateTransition.CallOpts, _selector)
}

// IsFunctionFreezable is a free data retrieval call binding the contract method 0xe81e0ba1.
//
// Solidity: function isFunctionFreezable(bytes4 _selector) view returns(bool)
func (_IZkSyncStateTransition *IZkSyncStateTransitionCallerSession) IsFunctionFreezable(_selector [4]byte) (bool, error) {
	return _IZkSyncStateTransition.Contract.IsFunctionFreezable(&_IZkSyncStateTransition.CallOpts, _selector)
}

// IsValidator is a free data retrieval call binding the contract method 0xfacd743b.
//
// Solidity: function isValidator(address _address) view returns(bool)
func (_IZkSyncStateTransition *IZkSyncStateTransitionCaller) IsValidator(opts *bind.CallOpts, _address common.Address) (bool, error) {
	var out []interface{}
	err := _IZkSyncStateTransition.contract.Call(opts, &out, "isValidator", _address)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsValidator is a free data retrieval call binding the contract method 0xfacd743b.
//
// Solidity: function isValidator(address _address) view returns(bool)
func (_IZkSyncStateTransition *IZkSyncStateTransitionSession) IsValidator(_address common.Address) (bool, error) {
	return _IZkSyncStateTransition.Contract.IsValidator(&_IZkSyncStateTransition.CallOpts, _address)
}

// IsValidator is a free data retrieval call binding the contract method 0xfacd743b.
//
// Solidity: function isValidator(address _address) view returns(bool)
func (_IZkSyncStateTransition *IZkSyncStateTransitionCallerSession) IsValidator(_address common.Address) (bool, error) {
	return _IZkSyncStateTransition.Contract.IsValidator(&_IZkSyncStateTransition.CallOpts, _address)
}

// L2LogsRootHash is a free data retrieval call binding the contract method 0x9cd939e4.
//
// Solidity: function l2LogsRootHash(uint256 _batchNumber) view returns(bytes32 merkleRoot)
func (_IZkSyncStateTransition *IZkSyncStateTransitionCaller) L2LogsRootHash(opts *bind.CallOpts, _batchNumber *big.Int) ([32]byte, error) {
	var out []interface{}
	err := _IZkSyncStateTransition.contract.Call(opts, &out, "l2LogsRootHash", _batchNumber)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// L2LogsRootHash is a free data retrieval call binding the contract method 0x9cd939e4.
//
// Solidity: function l2LogsRootHash(uint256 _batchNumber) view returns(bytes32 merkleRoot)
func (_IZkSyncStateTransition *IZkSyncStateTransitionSession) L2LogsRootHash(_batchNumber *big.Int) ([32]byte, error) {
	return _IZkSyncStateTransition.Contract.L2LogsRootHash(&_IZkSyncStateTransition.CallOpts, _batchNumber)
}

// L2LogsRootHash is a free data retrieval call binding the contract method 0x9cd939e4.
//
// Solidity: function l2LogsRootHash(uint256 _batchNumber) view returns(bytes32 merkleRoot)
func (_IZkSyncStateTransition *IZkSyncStateTransitionCallerSession) L2LogsRootHash(_batchNumber *big.Int) ([32]byte, error) {
	return _IZkSyncStateTransition.Contract.L2LogsRootHash(&_IZkSyncStateTransition.CallOpts, _batchNumber)
}

// L2TransactionBaseCost is a free data retrieval call binding the contract method 0xb473318e.
//
// Solidity: function l2TransactionBaseCost(uint256 _gasPrice, uint256 _l2GasLimit, uint256 _l2GasPerPubdataByteLimit) view returns(uint256)
func (_IZkSyncStateTransition *IZkSyncStateTransitionCaller) L2TransactionBaseCost(opts *bind.CallOpts, _gasPrice *big.Int, _l2GasLimit *big.Int, _l2GasPerPubdataByteLimit *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _IZkSyncStateTransition.contract.Call(opts, &out, "l2TransactionBaseCost", _gasPrice, _l2GasLimit, _l2GasPerPubdataByteLimit)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// L2TransactionBaseCost is a free data retrieval call binding the contract method 0xb473318e.
//
// Solidity: function l2TransactionBaseCost(uint256 _gasPrice, uint256 _l2GasLimit, uint256 _l2GasPerPubdataByteLimit) view returns(uint256)
func (_IZkSyncStateTransition *IZkSyncStateTransitionSession) L2TransactionBaseCost(_gasPrice *big.Int, _l2GasLimit *big.Int, _l2GasPerPubdataByteLimit *big.Int) (*big.Int, error) {
	return _IZkSyncStateTransition.Contract.L2TransactionBaseCost(&_IZkSyncStateTransition.CallOpts, _gasPrice, _l2GasLimit, _l2GasPerPubdataByteLimit)
}

// L2TransactionBaseCost is a free data retrieval call binding the contract method 0xb473318e.
//
// Solidity: function l2TransactionBaseCost(uint256 _gasPrice, uint256 _l2GasLimit, uint256 _l2GasPerPubdataByteLimit) view returns(uint256)
func (_IZkSyncStateTransition *IZkSyncStateTransitionCallerSession) L2TransactionBaseCost(_gasPrice *big.Int, _l2GasLimit *big.Int, _l2GasPerPubdataByteLimit *big.Int) (*big.Int, error) {
	return _IZkSyncStateTransition.Contract.L2TransactionBaseCost(&_IZkSyncStateTransition.CallOpts, _gasPrice, _l2GasLimit, _l2GasPerPubdataByteLimit)
}

// PriorityQueueFrontOperation is a free data retrieval call binding the contract method 0x56142d7a.
//
// Solidity: function priorityQueueFrontOperation() view returns((bytes32,uint64,uint192))
func (_IZkSyncStateTransition *IZkSyncStateTransitionCaller) PriorityQueueFrontOperation(opts *bind.CallOpts) (PriorityOperation, error) {
	var out []interface{}
	err := _IZkSyncStateTransition.contract.Call(opts, &out, "priorityQueueFrontOperation")

	if err != nil {
		return *new(PriorityOperation), err
	}

	out0 := *abi.ConvertType(out[0], new(PriorityOperation)).(*PriorityOperation)

	return out0, err

}

// PriorityQueueFrontOperation is a free data retrieval call binding the contract method 0x56142d7a.
//
// Solidity: function priorityQueueFrontOperation() view returns((bytes32,uint64,uint192))
func (_IZkSyncStateTransition *IZkSyncStateTransitionSession) PriorityQueueFrontOperation() (PriorityOperation, error) {
	return _IZkSyncStateTransition.Contract.PriorityQueueFrontOperation(&_IZkSyncStateTransition.CallOpts)
}

// PriorityQueueFrontOperation is a free data retrieval call binding the contract method 0x56142d7a.
//
// Solidity: function priorityQueueFrontOperation() view returns((bytes32,uint64,uint192))
func (_IZkSyncStateTransition *IZkSyncStateTransitionCallerSession) PriorityQueueFrontOperation() (PriorityOperation, error) {
	return _IZkSyncStateTransition.Contract.PriorityQueueFrontOperation(&_IZkSyncStateTransition.CallOpts)
}

// ProveL1ToL2TransactionStatus is a free data retrieval call binding the contract method 0x042901c7.
//
// Solidity: function proveL1ToL2TransactionStatus(bytes32 _l2TxHash, uint256 _l2BatchNumber, uint256 _l2MessageIndex, uint16 _l2TxNumberInBatch, bytes32[] _merkleProof, uint8 _status) view returns(bool)
func (_IZkSyncStateTransition *IZkSyncStateTransitionCaller) ProveL1ToL2TransactionStatus(opts *bind.CallOpts, _l2TxHash [32]byte, _l2BatchNumber *big.Int, _l2MessageIndex *big.Int, _l2TxNumberInBatch uint16, _merkleProof [][32]byte, _status uint8) (bool, error) {
	var out []interface{}
	err := _IZkSyncStateTransition.contract.Call(opts, &out, "proveL1ToL2TransactionStatus", _l2TxHash, _l2BatchNumber, _l2MessageIndex, _l2TxNumberInBatch, _merkleProof, _status)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// ProveL1ToL2TransactionStatus is a free data retrieval call binding the contract method 0x042901c7.
//
// Solidity: function proveL1ToL2TransactionStatus(bytes32 _l2TxHash, uint256 _l2BatchNumber, uint256 _l2MessageIndex, uint16 _l2TxNumberInBatch, bytes32[] _merkleProof, uint8 _status) view returns(bool)
func (_IZkSyncStateTransition *IZkSyncStateTransitionSession) ProveL1ToL2TransactionStatus(_l2TxHash [32]byte, _l2BatchNumber *big.Int, _l2MessageIndex *big.Int, _l2TxNumberInBatch uint16, _merkleProof [][32]byte, _status uint8) (bool, error) {
	return _IZkSyncStateTransition.Contract.ProveL1ToL2TransactionStatus(&_IZkSyncStateTransition.CallOpts, _l2TxHash, _l2BatchNumber, _l2MessageIndex, _l2TxNumberInBatch, _merkleProof, _status)
}

// ProveL1ToL2TransactionStatus is a free data retrieval call binding the contract method 0x042901c7.
//
// Solidity: function proveL1ToL2TransactionStatus(bytes32 _l2TxHash, uint256 _l2BatchNumber, uint256 _l2MessageIndex, uint16 _l2TxNumberInBatch, bytes32[] _merkleProof, uint8 _status) view returns(bool)
func (_IZkSyncStateTransition *IZkSyncStateTransitionCallerSession) ProveL1ToL2TransactionStatus(_l2TxHash [32]byte, _l2BatchNumber *big.Int, _l2MessageIndex *big.Int, _l2TxNumberInBatch uint16, _merkleProof [][32]byte, _status uint8) (bool, error) {
	return _IZkSyncStateTransition.Contract.ProveL1ToL2TransactionStatus(&_IZkSyncStateTransition.CallOpts, _l2TxHash, _l2BatchNumber, _l2MessageIndex, _l2TxNumberInBatch, _merkleProof, _status)
}

// ProveL2LogInclusion is a free data retrieval call binding the contract method 0x263b7f8e.
//
// Solidity: function proveL2LogInclusion(uint256 _batchNumber, uint256 _index, (uint8,bool,uint16,address,bytes32,bytes32) _log, bytes32[] _proof) view returns(bool)
func (_IZkSyncStateTransition *IZkSyncStateTransitionCaller) ProveL2LogInclusion(opts *bind.CallOpts, _batchNumber *big.Int, _index *big.Int, _log L2Log, _proof [][32]byte) (bool, error) {
	var out []interface{}
	err := _IZkSyncStateTransition.contract.Call(opts, &out, "proveL2LogInclusion", _batchNumber, _index, _log, _proof)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// ProveL2LogInclusion is a free data retrieval call binding the contract method 0x263b7f8e.
//
// Solidity: function proveL2LogInclusion(uint256 _batchNumber, uint256 _index, (uint8,bool,uint16,address,bytes32,bytes32) _log, bytes32[] _proof) view returns(bool)
func (_IZkSyncStateTransition *IZkSyncStateTransitionSession) ProveL2LogInclusion(_batchNumber *big.Int, _index *big.Int, _log L2Log, _proof [][32]byte) (bool, error) {
	return _IZkSyncStateTransition.Contract.ProveL2LogInclusion(&_IZkSyncStateTransition.CallOpts, _batchNumber, _index, _log, _proof)
}

// ProveL2LogInclusion is a free data retrieval call binding the contract method 0x263b7f8e.
//
// Solidity: function proveL2LogInclusion(uint256 _batchNumber, uint256 _index, (uint8,bool,uint16,address,bytes32,bytes32) _log, bytes32[] _proof) view returns(bool)
func (_IZkSyncStateTransition *IZkSyncStateTransitionCallerSession) ProveL2LogInclusion(_batchNumber *big.Int, _index *big.Int, _log L2Log, _proof [][32]byte) (bool, error) {
	return _IZkSyncStateTransition.Contract.ProveL2LogInclusion(&_IZkSyncStateTransition.CallOpts, _batchNumber, _index, _log, _proof)
}

// ProveL2MessageInclusion is a free data retrieval call binding the contract method 0xe4948f43.
//
// Solidity: function proveL2MessageInclusion(uint256 _batchNumber, uint256 _index, (uint16,address,bytes) _message, bytes32[] _proof) view returns(bool)
func (_IZkSyncStateTransition *IZkSyncStateTransitionCaller) ProveL2MessageInclusion(opts *bind.CallOpts, _batchNumber *big.Int, _index *big.Int, _message L2Message, _proof [][32]byte) (bool, error) {
	var out []interface{}
	err := _IZkSyncStateTransition.contract.Call(opts, &out, "proveL2MessageInclusion", _batchNumber, _index, _message, _proof)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// ProveL2MessageInclusion is a free data retrieval call binding the contract method 0xe4948f43.
//
// Solidity: function proveL2MessageInclusion(uint256 _batchNumber, uint256 _index, (uint16,address,bytes) _message, bytes32[] _proof) view returns(bool)
func (_IZkSyncStateTransition *IZkSyncStateTransitionSession) ProveL2MessageInclusion(_batchNumber *big.Int, _index *big.Int, _message L2Message, _proof [][32]byte) (bool, error) {
	return _IZkSyncStateTransition.Contract.ProveL2MessageInclusion(&_IZkSyncStateTransition.CallOpts, _batchNumber, _index, _message, _proof)
}

// ProveL2MessageInclusion is a free data retrieval call binding the contract method 0xe4948f43.
//
// Solidity: function proveL2MessageInclusion(uint256 _batchNumber, uint256 _index, (uint16,address,bytes) _message, bytes32[] _proof) view returns(bool)
func (_IZkSyncStateTransition *IZkSyncStateTransitionCallerSession) ProveL2MessageInclusion(_batchNumber *big.Int, _index *big.Int, _message L2Message, _proof [][32]byte) (bool, error) {
	return _IZkSyncStateTransition.Contract.ProveL2MessageInclusion(&_IZkSyncStateTransition.CallOpts, _batchNumber, _index, _message, _proof)
}

// StoredBatchHash is a free data retrieval call binding the contract method 0xb22dd78e.
//
// Solidity: function storedBatchHash(uint256 _batchNumber) view returns(bytes32)
func (_IZkSyncStateTransition *IZkSyncStateTransitionCaller) StoredBatchHash(opts *bind.CallOpts, _batchNumber *big.Int) ([32]byte, error) {
	var out []interface{}
	err := _IZkSyncStateTransition.contract.Call(opts, &out, "storedBatchHash", _batchNumber)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// StoredBatchHash is a free data retrieval call binding the contract method 0xb22dd78e.
//
// Solidity: function storedBatchHash(uint256 _batchNumber) view returns(bytes32)
func (_IZkSyncStateTransition *IZkSyncStateTransitionSession) StoredBatchHash(_batchNumber *big.Int) ([32]byte, error) {
	return _IZkSyncStateTransition.Contract.StoredBatchHash(&_IZkSyncStateTransition.CallOpts, _batchNumber)
}

// StoredBatchHash is a free data retrieval call binding the contract method 0xb22dd78e.
//
// Solidity: function storedBatchHash(uint256 _batchNumber) view returns(bytes32)
func (_IZkSyncStateTransition *IZkSyncStateTransitionCallerSession) StoredBatchHash(_batchNumber *big.Int) ([32]byte, error) {
	return _IZkSyncStateTransition.Contract.StoredBatchHash(&_IZkSyncStateTransition.CallOpts, _batchNumber)
}

// AcceptAdmin is a paid mutator transaction binding the contract method 0x0e18b681.
//
// Solidity: function acceptAdmin() returns()
func (_IZkSyncStateTransition *IZkSyncStateTransitionTransactor) AcceptAdmin(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IZkSyncStateTransition.contract.Transact(opts, "acceptAdmin")
}

// AcceptAdmin is a paid mutator transaction binding the contract method 0x0e18b681.
//
// Solidity: function acceptAdmin() returns()
func (_IZkSyncStateTransition *IZkSyncStateTransitionSession) AcceptAdmin() (*types.Transaction, error) {
	return _IZkSyncStateTransition.Contract.AcceptAdmin(&_IZkSyncStateTransition.TransactOpts)
}

// AcceptAdmin is a paid mutator transaction binding the contract method 0x0e18b681.
//
// Solidity: function acceptAdmin() returns()
func (_IZkSyncStateTransition *IZkSyncStateTransitionTransactorSession) AcceptAdmin() (*types.Transaction, error) {
	return _IZkSyncStateTransition.Contract.AcceptAdmin(&_IZkSyncStateTransition.TransactOpts)
}

// BridgehubRequestL2Transaction is a paid mutator transaction binding the contract method 0x12f43dab.
//
// Solidity: function bridgehubRequestL2Transaction((address,address,uint256,uint256,bytes,uint256,uint256,bytes[],address) _request) payable returns(bytes32 canonicalTxHash)
func (_IZkSyncStateTransition *IZkSyncStateTransitionTransactor) BridgehubRequestL2Transaction(opts *bind.TransactOpts, _request BridgehubL2TransactionRequest) (*types.Transaction, error) {
	return _IZkSyncStateTransition.contract.Transact(opts, "bridgehubRequestL2Transaction", _request)
}

// BridgehubRequestL2Transaction is a paid mutator transaction binding the contract method 0x12f43dab.
//
// Solidity: function bridgehubRequestL2Transaction((address,address,uint256,uint256,bytes,uint256,uint256,bytes[],address) _request) payable returns(bytes32 canonicalTxHash)
func (_IZkSyncStateTransition *IZkSyncStateTransitionSession) BridgehubRequestL2Transaction(_request BridgehubL2TransactionRequest) (*types.Transaction, error) {
	return _IZkSyncStateTransition.Contract.BridgehubRequestL2Transaction(&_IZkSyncStateTransition.TransactOpts, _request)
}

// BridgehubRequestL2Transaction is a paid mutator transaction binding the contract method 0x12f43dab.
//
// Solidity: function bridgehubRequestL2Transaction((address,address,uint256,uint256,bytes,uint256,uint256,bytes[],address) _request) payable returns(bytes32 canonicalTxHash)
func (_IZkSyncStateTransition *IZkSyncStateTransitionTransactorSession) BridgehubRequestL2Transaction(_request BridgehubL2TransactionRequest) (*types.Transaction, error) {
	return _IZkSyncStateTransition.Contract.BridgehubRequestL2Transaction(&_IZkSyncStateTransition.TransactOpts, _request)
}

// ChangeFeeParams is a paid mutator transaction binding the contract method 0x64bf8d66.
//
// Solidity: function changeFeeParams((uint8,uint32,uint32,uint32,uint32,uint64) _newFeeParams) returns()
func (_IZkSyncStateTransition *IZkSyncStateTransitionTransactor) ChangeFeeParams(opts *bind.TransactOpts, _newFeeParams FeeParams) (*types.Transaction, error) {
	return _IZkSyncStateTransition.contract.Transact(opts, "changeFeeParams", _newFeeParams)
}

// ChangeFeeParams is a paid mutator transaction binding the contract method 0x64bf8d66.
//
// Solidity: function changeFeeParams((uint8,uint32,uint32,uint32,uint32,uint64) _newFeeParams) returns()
func (_IZkSyncStateTransition *IZkSyncStateTransitionSession) ChangeFeeParams(_newFeeParams FeeParams) (*types.Transaction, error) {
	return _IZkSyncStateTransition.Contract.ChangeFeeParams(&_IZkSyncStateTransition.TransactOpts, _newFeeParams)
}

// ChangeFeeParams is a paid mutator transaction binding the contract method 0x64bf8d66.
//
// Solidity: function changeFeeParams((uint8,uint32,uint32,uint32,uint32,uint64) _newFeeParams) returns()
func (_IZkSyncStateTransition *IZkSyncStateTransitionTransactorSession) ChangeFeeParams(_newFeeParams FeeParams) (*types.Transaction, error) {
	return _IZkSyncStateTransition.Contract.ChangeFeeParams(&_IZkSyncStateTransition.TransactOpts, _newFeeParams)
}

// CommitBatches is a paid mutator transaction binding the contract method 0x701f58c5.
//
// Solidity: function commitBatches((uint64,bytes32,uint64,uint256,bytes32,bytes32,uint256,bytes32) _lastCommittedBatchData, (uint64,uint64,uint64,bytes32,uint256,bytes32,bytes32,bytes32,bytes,bytes)[] _newBatchesData) returns()
func (_IZkSyncStateTransition *IZkSyncStateTransitionTransactor) CommitBatches(opts *bind.TransactOpts, _lastCommittedBatchData IExecutorStoredBatchInfo, _newBatchesData []IExecutorCommitBatchInfo) (*types.Transaction, error) {
	return _IZkSyncStateTransition.contract.Transact(opts, "commitBatches", _lastCommittedBatchData, _newBatchesData)
}

// CommitBatches is a paid mutator transaction binding the contract method 0x701f58c5.
//
// Solidity: function commitBatches((uint64,bytes32,uint64,uint256,bytes32,bytes32,uint256,bytes32) _lastCommittedBatchData, (uint64,uint64,uint64,bytes32,uint256,bytes32,bytes32,bytes32,bytes,bytes)[] _newBatchesData) returns()
func (_IZkSyncStateTransition *IZkSyncStateTransitionSession) CommitBatches(_lastCommittedBatchData IExecutorStoredBatchInfo, _newBatchesData []IExecutorCommitBatchInfo) (*types.Transaction, error) {
	return _IZkSyncStateTransition.Contract.CommitBatches(&_IZkSyncStateTransition.TransactOpts, _lastCommittedBatchData, _newBatchesData)
}

// CommitBatches is a paid mutator transaction binding the contract method 0x701f58c5.
//
// Solidity: function commitBatches((uint64,bytes32,uint64,uint256,bytes32,bytes32,uint256,bytes32) _lastCommittedBatchData, (uint64,uint64,uint64,bytes32,uint256,bytes32,bytes32,bytes32,bytes,bytes)[] _newBatchesData) returns()
func (_IZkSyncStateTransition *IZkSyncStateTransitionTransactorSession) CommitBatches(_lastCommittedBatchData IExecutorStoredBatchInfo, _newBatchesData []IExecutorCommitBatchInfo) (*types.Transaction, error) {
	return _IZkSyncStateTransition.Contract.CommitBatches(&_IZkSyncStateTransition.TransactOpts, _lastCommittedBatchData, _newBatchesData)
}

// CommitBatchesSharedBridge is a paid mutator transaction binding the contract method 0x6edd4f12.
//
// Solidity: function commitBatchesSharedBridge(uint256 _chainId, (uint64,bytes32,uint64,uint256,bytes32,bytes32,uint256,bytes32) _lastCommittedBatchData, (uint64,uint64,uint64,bytes32,uint256,bytes32,bytes32,bytes32,bytes,bytes)[] _newBatchesData) returns()
func (_IZkSyncStateTransition *IZkSyncStateTransitionTransactor) CommitBatchesSharedBridge(opts *bind.TransactOpts, _chainId *big.Int, _lastCommittedBatchData IExecutorStoredBatchInfo, _newBatchesData []IExecutorCommitBatchInfo) (*types.Transaction, error) {
	return _IZkSyncStateTransition.contract.Transact(opts, "commitBatchesSharedBridge", _chainId, _lastCommittedBatchData, _newBatchesData)
}

// CommitBatchesSharedBridge is a paid mutator transaction binding the contract method 0x6edd4f12.
//
// Solidity: function commitBatchesSharedBridge(uint256 _chainId, (uint64,bytes32,uint64,uint256,bytes32,bytes32,uint256,bytes32) _lastCommittedBatchData, (uint64,uint64,uint64,bytes32,uint256,bytes32,bytes32,bytes32,bytes,bytes)[] _newBatchesData) returns()
func (_IZkSyncStateTransition *IZkSyncStateTransitionSession) CommitBatchesSharedBridge(_chainId *big.Int, _lastCommittedBatchData IExecutorStoredBatchInfo, _newBatchesData []IExecutorCommitBatchInfo) (*types.Transaction, error) {
	return _IZkSyncStateTransition.Contract.CommitBatchesSharedBridge(&_IZkSyncStateTransition.TransactOpts, _chainId, _lastCommittedBatchData, _newBatchesData)
}

// CommitBatchesSharedBridge is a paid mutator transaction binding the contract method 0x6edd4f12.
//
// Solidity: function commitBatchesSharedBridge(uint256 _chainId, (uint64,bytes32,uint64,uint256,bytes32,bytes32,uint256,bytes32) _lastCommittedBatchData, (uint64,uint64,uint64,bytes32,uint256,bytes32,bytes32,bytes32,bytes,bytes)[] _newBatchesData) returns()
func (_IZkSyncStateTransition *IZkSyncStateTransitionTransactorSession) CommitBatchesSharedBridge(_chainId *big.Int, _lastCommittedBatchData IExecutorStoredBatchInfo, _newBatchesData []IExecutorCommitBatchInfo) (*types.Transaction, error) {
	return _IZkSyncStateTransition.Contract.CommitBatchesSharedBridge(&_IZkSyncStateTransition.TransactOpts, _chainId, _lastCommittedBatchData, _newBatchesData)
}

// ExecuteBatches is a paid mutator transaction binding the contract method 0xc3d93e7c.
//
// Solidity: function executeBatches((uint64,bytes32,uint64,uint256,bytes32,bytes32,uint256,bytes32)[] _batchesData) returns()
func (_IZkSyncStateTransition *IZkSyncStateTransitionTransactor) ExecuteBatches(opts *bind.TransactOpts, _batchesData []IExecutorStoredBatchInfo) (*types.Transaction, error) {
	return _IZkSyncStateTransition.contract.Transact(opts, "executeBatches", _batchesData)
}

// ExecuteBatches is a paid mutator transaction binding the contract method 0xc3d93e7c.
//
// Solidity: function executeBatches((uint64,bytes32,uint64,uint256,bytes32,bytes32,uint256,bytes32)[] _batchesData) returns()
func (_IZkSyncStateTransition *IZkSyncStateTransitionSession) ExecuteBatches(_batchesData []IExecutorStoredBatchInfo) (*types.Transaction, error) {
	return _IZkSyncStateTransition.Contract.ExecuteBatches(&_IZkSyncStateTransition.TransactOpts, _batchesData)
}

// ExecuteBatches is a paid mutator transaction binding the contract method 0xc3d93e7c.
//
// Solidity: function executeBatches((uint64,bytes32,uint64,uint256,bytes32,bytes32,uint256,bytes32)[] _batchesData) returns()
func (_IZkSyncStateTransition *IZkSyncStateTransitionTransactorSession) ExecuteBatches(_batchesData []IExecutorStoredBatchInfo) (*types.Transaction, error) {
	return _IZkSyncStateTransition.Contract.ExecuteBatches(&_IZkSyncStateTransition.TransactOpts, _batchesData)
}

// ExecuteBatchesSharedBridge is a paid mutator transaction binding the contract method 0x6f497ac6.
//
// Solidity: function executeBatchesSharedBridge(uint256 _chainId, (uint64,bytes32,uint64,uint256,bytes32,bytes32,uint256,bytes32)[] _batchesData) returns()
func (_IZkSyncStateTransition *IZkSyncStateTransitionTransactor) ExecuteBatchesSharedBridge(opts *bind.TransactOpts, _chainId *big.Int, _batchesData []IExecutorStoredBatchInfo) (*types.Transaction, error) {
	return _IZkSyncStateTransition.contract.Transact(opts, "executeBatchesSharedBridge", _chainId, _batchesData)
}

// ExecuteBatchesSharedBridge is a paid mutator transaction binding the contract method 0x6f497ac6.
//
// Solidity: function executeBatchesSharedBridge(uint256 _chainId, (uint64,bytes32,uint64,uint256,bytes32,bytes32,uint256,bytes32)[] _batchesData) returns()
func (_IZkSyncStateTransition *IZkSyncStateTransitionSession) ExecuteBatchesSharedBridge(_chainId *big.Int, _batchesData []IExecutorStoredBatchInfo) (*types.Transaction, error) {
	return _IZkSyncStateTransition.Contract.ExecuteBatchesSharedBridge(&_IZkSyncStateTransition.TransactOpts, _chainId, _batchesData)
}

// ExecuteBatchesSharedBridge is a paid mutator transaction binding the contract method 0x6f497ac6.
//
// Solidity: function executeBatchesSharedBridge(uint256 _chainId, (uint64,bytes32,uint64,uint256,bytes32,bytes32,uint256,bytes32)[] _batchesData) returns()
func (_IZkSyncStateTransition *IZkSyncStateTransitionTransactorSession) ExecuteBatchesSharedBridge(_chainId *big.Int, _batchesData []IExecutorStoredBatchInfo) (*types.Transaction, error) {
	return _IZkSyncStateTransition.Contract.ExecuteBatchesSharedBridge(&_IZkSyncStateTransition.TransactOpts, _chainId, _batchesData)
}

// ExecuteUpgrade is a paid mutator transaction binding the contract method 0xa9f6d941.
//
// Solidity: function executeUpgrade(((address,uint8,bool,bytes4[])[],address,bytes) _diamondCut) returns()
func (_IZkSyncStateTransition *IZkSyncStateTransitionTransactor) ExecuteUpgrade(opts *bind.TransactOpts, _diamondCut DiamondDiamondCutData) (*types.Transaction, error) {
	return _IZkSyncStateTransition.contract.Transact(opts, "executeUpgrade", _diamondCut)
}

// ExecuteUpgrade is a paid mutator transaction binding the contract method 0xa9f6d941.
//
// Solidity: function executeUpgrade(((address,uint8,bool,bytes4[])[],address,bytes) _diamondCut) returns()
func (_IZkSyncStateTransition *IZkSyncStateTransitionSession) ExecuteUpgrade(_diamondCut DiamondDiamondCutData) (*types.Transaction, error) {
	return _IZkSyncStateTransition.Contract.ExecuteUpgrade(&_IZkSyncStateTransition.TransactOpts, _diamondCut)
}

// ExecuteUpgrade is a paid mutator transaction binding the contract method 0xa9f6d941.
//
// Solidity: function executeUpgrade(((address,uint8,bool,bytes4[])[],address,bytes) _diamondCut) returns()
func (_IZkSyncStateTransition *IZkSyncStateTransitionTransactorSession) ExecuteUpgrade(_diamondCut DiamondDiamondCutData) (*types.Transaction, error) {
	return _IZkSyncStateTransition.Contract.ExecuteUpgrade(&_IZkSyncStateTransition.TransactOpts, _diamondCut)
}

// FinalizeEthWithdrawal is a paid mutator transaction binding the contract method 0x6c0960f9.
//
// Solidity: function finalizeEthWithdrawal(uint256 _l2BatchNumber, uint256 _l2MessageIndex, uint16 _l2TxNumberInBatch, bytes _message, bytes32[] _merkleProof) returns()
func (_IZkSyncStateTransition *IZkSyncStateTransitionTransactor) FinalizeEthWithdrawal(opts *bind.TransactOpts, _l2BatchNumber *big.Int, _l2MessageIndex *big.Int, _l2TxNumberInBatch uint16, _message []byte, _merkleProof [][32]byte) (*types.Transaction, error) {
	return _IZkSyncStateTransition.contract.Transact(opts, "finalizeEthWithdrawal", _l2BatchNumber, _l2MessageIndex, _l2TxNumberInBatch, _message, _merkleProof)
}

// FinalizeEthWithdrawal is a paid mutator transaction binding the contract method 0x6c0960f9.
//
// Solidity: function finalizeEthWithdrawal(uint256 _l2BatchNumber, uint256 _l2MessageIndex, uint16 _l2TxNumberInBatch, bytes _message, bytes32[] _merkleProof) returns()
func (_IZkSyncStateTransition *IZkSyncStateTransitionSession) FinalizeEthWithdrawal(_l2BatchNumber *big.Int, _l2MessageIndex *big.Int, _l2TxNumberInBatch uint16, _message []byte, _merkleProof [][32]byte) (*types.Transaction, error) {
	return _IZkSyncStateTransition.Contract.FinalizeEthWithdrawal(&_IZkSyncStateTransition.TransactOpts, _l2BatchNumber, _l2MessageIndex, _l2TxNumberInBatch, _message, _merkleProof)
}

// FinalizeEthWithdrawal is a paid mutator transaction binding the contract method 0x6c0960f9.
//
// Solidity: function finalizeEthWithdrawal(uint256 _l2BatchNumber, uint256 _l2MessageIndex, uint16 _l2TxNumberInBatch, bytes _message, bytes32[] _merkleProof) returns()
func (_IZkSyncStateTransition *IZkSyncStateTransitionTransactorSession) FinalizeEthWithdrawal(_l2BatchNumber *big.Int, _l2MessageIndex *big.Int, _l2TxNumberInBatch uint16, _message []byte, _merkleProof [][32]byte) (*types.Transaction, error) {
	return _IZkSyncStateTransition.Contract.FinalizeEthWithdrawal(&_IZkSyncStateTransition.TransactOpts, _l2BatchNumber, _l2MessageIndex, _l2TxNumberInBatch, _message, _merkleProof)
}

// FreezeDiamond is a paid mutator transaction binding the contract method 0x27ae4c16.
//
// Solidity: function freezeDiamond() returns()
func (_IZkSyncStateTransition *IZkSyncStateTransitionTransactor) FreezeDiamond(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IZkSyncStateTransition.contract.Transact(opts, "freezeDiamond")
}

// FreezeDiamond is a paid mutator transaction binding the contract method 0x27ae4c16.
//
// Solidity: function freezeDiamond() returns()
func (_IZkSyncStateTransition *IZkSyncStateTransitionSession) FreezeDiamond() (*types.Transaction, error) {
	return _IZkSyncStateTransition.Contract.FreezeDiamond(&_IZkSyncStateTransition.TransactOpts)
}

// FreezeDiamond is a paid mutator transaction binding the contract method 0x27ae4c16.
//
// Solidity: function freezeDiamond() returns()
func (_IZkSyncStateTransition *IZkSyncStateTransitionTransactorSession) FreezeDiamond() (*types.Transaction, error) {
	return _IZkSyncStateTransition.Contract.FreezeDiamond(&_IZkSyncStateTransition.TransactOpts)
}

// ProveBatches is a paid mutator transaction binding the contract method 0x7f61885c.
//
// Solidity: function proveBatches((uint64,bytes32,uint64,uint256,bytes32,bytes32,uint256,bytes32) _prevBatch, (uint64,bytes32,uint64,uint256,bytes32,bytes32,uint256,bytes32)[] _committedBatches, (uint256[],uint256[]) _proof) returns()
func (_IZkSyncStateTransition *IZkSyncStateTransitionTransactor) ProveBatches(opts *bind.TransactOpts, _prevBatch IExecutorStoredBatchInfo, _committedBatches []IExecutorStoredBatchInfo, _proof IExecutorProofInput) (*types.Transaction, error) {
	return _IZkSyncStateTransition.contract.Transact(opts, "proveBatches", _prevBatch, _committedBatches, _proof)
}

// ProveBatches is a paid mutator transaction binding the contract method 0x7f61885c.
//
// Solidity: function proveBatches((uint64,bytes32,uint64,uint256,bytes32,bytes32,uint256,bytes32) _prevBatch, (uint64,bytes32,uint64,uint256,bytes32,bytes32,uint256,bytes32)[] _committedBatches, (uint256[],uint256[]) _proof) returns()
func (_IZkSyncStateTransition *IZkSyncStateTransitionSession) ProveBatches(_prevBatch IExecutorStoredBatchInfo, _committedBatches []IExecutorStoredBatchInfo, _proof IExecutorProofInput) (*types.Transaction, error) {
	return _IZkSyncStateTransition.Contract.ProveBatches(&_IZkSyncStateTransition.TransactOpts, _prevBatch, _committedBatches, _proof)
}

// ProveBatches is a paid mutator transaction binding the contract method 0x7f61885c.
//
// Solidity: function proveBatches((uint64,bytes32,uint64,uint256,bytes32,bytes32,uint256,bytes32) _prevBatch, (uint64,bytes32,uint64,uint256,bytes32,bytes32,uint256,bytes32)[] _committedBatches, (uint256[],uint256[]) _proof) returns()
func (_IZkSyncStateTransition *IZkSyncStateTransitionTransactorSession) ProveBatches(_prevBatch IExecutorStoredBatchInfo, _committedBatches []IExecutorStoredBatchInfo, _proof IExecutorProofInput) (*types.Transaction, error) {
	return _IZkSyncStateTransition.Contract.ProveBatches(&_IZkSyncStateTransition.TransactOpts, _prevBatch, _committedBatches, _proof)
}

// ProveBatchesSharedBridge is a paid mutator transaction binding the contract method 0xc37533bb.
//
// Solidity: function proveBatchesSharedBridge(uint256 _chainId, (uint64,bytes32,uint64,uint256,bytes32,bytes32,uint256,bytes32) _prevBatch, (uint64,bytes32,uint64,uint256,bytes32,bytes32,uint256,bytes32)[] _committedBatches, (uint256[],uint256[]) _proof) returns()
func (_IZkSyncStateTransition *IZkSyncStateTransitionTransactor) ProveBatchesSharedBridge(opts *bind.TransactOpts, _chainId *big.Int, _prevBatch IExecutorStoredBatchInfo, _committedBatches []IExecutorStoredBatchInfo, _proof IExecutorProofInput) (*types.Transaction, error) {
	return _IZkSyncStateTransition.contract.Transact(opts, "proveBatchesSharedBridge", _chainId, _prevBatch, _committedBatches, _proof)
}

// ProveBatchesSharedBridge is a paid mutator transaction binding the contract method 0xc37533bb.
//
// Solidity: function proveBatchesSharedBridge(uint256 _chainId, (uint64,bytes32,uint64,uint256,bytes32,bytes32,uint256,bytes32) _prevBatch, (uint64,bytes32,uint64,uint256,bytes32,bytes32,uint256,bytes32)[] _committedBatches, (uint256[],uint256[]) _proof) returns()
func (_IZkSyncStateTransition *IZkSyncStateTransitionSession) ProveBatchesSharedBridge(_chainId *big.Int, _prevBatch IExecutorStoredBatchInfo, _committedBatches []IExecutorStoredBatchInfo, _proof IExecutorProofInput) (*types.Transaction, error) {
	return _IZkSyncStateTransition.Contract.ProveBatchesSharedBridge(&_IZkSyncStateTransition.TransactOpts, _chainId, _prevBatch, _committedBatches, _proof)
}

// ProveBatchesSharedBridge is a paid mutator transaction binding the contract method 0xc37533bb.
//
// Solidity: function proveBatchesSharedBridge(uint256 _chainId, (uint64,bytes32,uint64,uint256,bytes32,bytes32,uint256,bytes32) _prevBatch, (uint64,bytes32,uint64,uint256,bytes32,bytes32,uint256,bytes32)[] _committedBatches, (uint256[],uint256[]) _proof) returns()
func (_IZkSyncStateTransition *IZkSyncStateTransitionTransactorSession) ProveBatchesSharedBridge(_chainId *big.Int, _prevBatch IExecutorStoredBatchInfo, _committedBatches []IExecutorStoredBatchInfo, _proof IExecutorProofInput) (*types.Transaction, error) {
	return _IZkSyncStateTransition.Contract.ProveBatchesSharedBridge(&_IZkSyncStateTransition.TransactOpts, _chainId, _prevBatch, _committedBatches, _proof)
}

// RequestL2Transaction is a paid mutator transaction binding the contract method 0xeb672419.
//
// Solidity: function requestL2Transaction(address _contractL2, uint256 _l2Value, bytes _calldata, uint256 _l2GasLimit, uint256 _l2GasPerPubdataByteLimit, bytes[] _factoryDeps, address _refundRecipient) payable returns(bytes32 canonicalTxHash)
func (_IZkSyncStateTransition *IZkSyncStateTransitionTransactor) RequestL2Transaction(opts *bind.TransactOpts, _contractL2 common.Address, _l2Value *big.Int, _calldata []byte, _l2GasLimit *big.Int, _l2GasPerPubdataByteLimit *big.Int, _factoryDeps [][]byte, _refundRecipient common.Address) (*types.Transaction, error) {
	return _IZkSyncStateTransition.contract.Transact(opts, "requestL2Transaction", _contractL2, _l2Value, _calldata, _l2GasLimit, _l2GasPerPubdataByteLimit, _factoryDeps, _refundRecipient)
}

// RequestL2Transaction is a paid mutator transaction binding the contract method 0xeb672419.
//
// Solidity: function requestL2Transaction(address _contractL2, uint256 _l2Value, bytes _calldata, uint256 _l2GasLimit, uint256 _l2GasPerPubdataByteLimit, bytes[] _factoryDeps, address _refundRecipient) payable returns(bytes32 canonicalTxHash)
func (_IZkSyncStateTransition *IZkSyncStateTransitionSession) RequestL2Transaction(_contractL2 common.Address, _l2Value *big.Int, _calldata []byte, _l2GasLimit *big.Int, _l2GasPerPubdataByteLimit *big.Int, _factoryDeps [][]byte, _refundRecipient common.Address) (*types.Transaction, error) {
	return _IZkSyncStateTransition.Contract.RequestL2Transaction(&_IZkSyncStateTransition.TransactOpts, _contractL2, _l2Value, _calldata, _l2GasLimit, _l2GasPerPubdataByteLimit, _factoryDeps, _refundRecipient)
}

// RequestL2Transaction is a paid mutator transaction binding the contract method 0xeb672419.
//
// Solidity: function requestL2Transaction(address _contractL2, uint256 _l2Value, bytes _calldata, uint256 _l2GasLimit, uint256 _l2GasPerPubdataByteLimit, bytes[] _factoryDeps, address _refundRecipient) payable returns(bytes32 canonicalTxHash)
func (_IZkSyncStateTransition *IZkSyncStateTransitionTransactorSession) RequestL2Transaction(_contractL2 common.Address, _l2Value *big.Int, _calldata []byte, _l2GasLimit *big.Int, _l2GasPerPubdataByteLimit *big.Int, _factoryDeps [][]byte, _refundRecipient common.Address) (*types.Transaction, error) {
	return _IZkSyncStateTransition.Contract.RequestL2Transaction(&_IZkSyncStateTransition.TransactOpts, _contractL2, _l2Value, _calldata, _l2GasLimit, _l2GasPerPubdataByteLimit, _factoryDeps, _refundRecipient)
}

// RevertBatches is a paid mutator transaction binding the contract method 0x97c09d34.
//
// Solidity: function revertBatches(uint256 _newLastBatch) returns()
func (_IZkSyncStateTransition *IZkSyncStateTransitionTransactor) RevertBatches(opts *bind.TransactOpts, _newLastBatch *big.Int) (*types.Transaction, error) {
	return _IZkSyncStateTransition.contract.Transact(opts, "revertBatches", _newLastBatch)
}

// RevertBatches is a paid mutator transaction binding the contract method 0x97c09d34.
//
// Solidity: function revertBatches(uint256 _newLastBatch) returns()
func (_IZkSyncStateTransition *IZkSyncStateTransitionSession) RevertBatches(_newLastBatch *big.Int) (*types.Transaction, error) {
	return _IZkSyncStateTransition.Contract.RevertBatches(&_IZkSyncStateTransition.TransactOpts, _newLastBatch)
}

// RevertBatches is a paid mutator transaction binding the contract method 0x97c09d34.
//
// Solidity: function revertBatches(uint256 _newLastBatch) returns()
func (_IZkSyncStateTransition *IZkSyncStateTransitionTransactorSession) RevertBatches(_newLastBatch *big.Int) (*types.Transaction, error) {
	return _IZkSyncStateTransition.Contract.RevertBatches(&_IZkSyncStateTransition.TransactOpts, _newLastBatch)
}

// RevertBatchesSharedBridge is a paid mutator transaction binding the contract method 0x0f23da43.
//
// Solidity: function revertBatchesSharedBridge(uint256 _chainId, uint256 _newLastBatch) returns()
func (_IZkSyncStateTransition *IZkSyncStateTransitionTransactor) RevertBatchesSharedBridge(opts *bind.TransactOpts, _chainId *big.Int, _newLastBatch *big.Int) (*types.Transaction, error) {
	return _IZkSyncStateTransition.contract.Transact(opts, "revertBatchesSharedBridge", _chainId, _newLastBatch)
}

// RevertBatchesSharedBridge is a paid mutator transaction binding the contract method 0x0f23da43.
//
// Solidity: function revertBatchesSharedBridge(uint256 _chainId, uint256 _newLastBatch) returns()
func (_IZkSyncStateTransition *IZkSyncStateTransitionSession) RevertBatchesSharedBridge(_chainId *big.Int, _newLastBatch *big.Int) (*types.Transaction, error) {
	return _IZkSyncStateTransition.Contract.RevertBatchesSharedBridge(&_IZkSyncStateTransition.TransactOpts, _chainId, _newLastBatch)
}

// RevertBatchesSharedBridge is a paid mutator transaction binding the contract method 0x0f23da43.
//
// Solidity: function revertBatchesSharedBridge(uint256 _chainId, uint256 _newLastBatch) returns()
func (_IZkSyncStateTransition *IZkSyncStateTransitionTransactorSession) RevertBatchesSharedBridge(_chainId *big.Int, _newLastBatch *big.Int) (*types.Transaction, error) {
	return _IZkSyncStateTransition.Contract.RevertBatchesSharedBridge(&_IZkSyncStateTransition.TransactOpts, _chainId, _newLastBatch)
}

// SetPendingAdmin is a paid mutator transaction binding the contract method 0x4dd18bf5.
//
// Solidity: function setPendingAdmin(address _newPendingAdmin) returns()
func (_IZkSyncStateTransition *IZkSyncStateTransitionTransactor) SetPendingAdmin(opts *bind.TransactOpts, _newPendingAdmin common.Address) (*types.Transaction, error) {
	return _IZkSyncStateTransition.contract.Transact(opts, "setPendingAdmin", _newPendingAdmin)
}

// SetPendingAdmin is a paid mutator transaction binding the contract method 0x4dd18bf5.
//
// Solidity: function setPendingAdmin(address _newPendingAdmin) returns()
func (_IZkSyncStateTransition *IZkSyncStateTransitionSession) SetPendingAdmin(_newPendingAdmin common.Address) (*types.Transaction, error) {
	return _IZkSyncStateTransition.Contract.SetPendingAdmin(&_IZkSyncStateTransition.TransactOpts, _newPendingAdmin)
}

// SetPendingAdmin is a paid mutator transaction binding the contract method 0x4dd18bf5.
//
// Solidity: function setPendingAdmin(address _newPendingAdmin) returns()
func (_IZkSyncStateTransition *IZkSyncStateTransitionTransactorSession) SetPendingAdmin(_newPendingAdmin common.Address) (*types.Transaction, error) {
	return _IZkSyncStateTransition.Contract.SetPendingAdmin(&_IZkSyncStateTransition.TransactOpts, _newPendingAdmin)
}

// SetPorterAvailability is a paid mutator transaction binding the contract method 0x1cc5d103.
//
// Solidity: function setPorterAvailability(bool _zkPorterIsAvailable) returns()
func (_IZkSyncStateTransition *IZkSyncStateTransitionTransactor) SetPorterAvailability(opts *bind.TransactOpts, _zkPorterIsAvailable bool) (*types.Transaction, error) {
	return _IZkSyncStateTransition.contract.Transact(opts, "setPorterAvailability", _zkPorterIsAvailable)
}

// SetPorterAvailability is a paid mutator transaction binding the contract method 0x1cc5d103.
//
// Solidity: function setPorterAvailability(bool _zkPorterIsAvailable) returns()
func (_IZkSyncStateTransition *IZkSyncStateTransitionSession) SetPorterAvailability(_zkPorterIsAvailable bool) (*types.Transaction, error) {
	return _IZkSyncStateTransition.Contract.SetPorterAvailability(&_IZkSyncStateTransition.TransactOpts, _zkPorterIsAvailable)
}

// SetPorterAvailability is a paid mutator transaction binding the contract method 0x1cc5d103.
//
// Solidity: function setPorterAvailability(bool _zkPorterIsAvailable) returns()
func (_IZkSyncStateTransition *IZkSyncStateTransitionTransactorSession) SetPorterAvailability(_zkPorterIsAvailable bool) (*types.Transaction, error) {
	return _IZkSyncStateTransition.Contract.SetPorterAvailability(&_IZkSyncStateTransition.TransactOpts, _zkPorterIsAvailable)
}

// SetPriorityTxMaxGasLimit is a paid mutator transaction binding the contract method 0xbe6f11cf.
//
// Solidity: function setPriorityTxMaxGasLimit(uint256 _newPriorityTxMaxGasLimit) returns()
func (_IZkSyncStateTransition *IZkSyncStateTransitionTransactor) SetPriorityTxMaxGasLimit(opts *bind.TransactOpts, _newPriorityTxMaxGasLimit *big.Int) (*types.Transaction, error) {
	return _IZkSyncStateTransition.contract.Transact(opts, "setPriorityTxMaxGasLimit", _newPriorityTxMaxGasLimit)
}

// SetPriorityTxMaxGasLimit is a paid mutator transaction binding the contract method 0xbe6f11cf.
//
// Solidity: function setPriorityTxMaxGasLimit(uint256 _newPriorityTxMaxGasLimit) returns()
func (_IZkSyncStateTransition *IZkSyncStateTransitionSession) SetPriorityTxMaxGasLimit(_newPriorityTxMaxGasLimit *big.Int) (*types.Transaction, error) {
	return _IZkSyncStateTransition.Contract.SetPriorityTxMaxGasLimit(&_IZkSyncStateTransition.TransactOpts, _newPriorityTxMaxGasLimit)
}

// SetPriorityTxMaxGasLimit is a paid mutator transaction binding the contract method 0xbe6f11cf.
//
// Solidity: function setPriorityTxMaxGasLimit(uint256 _newPriorityTxMaxGasLimit) returns()
func (_IZkSyncStateTransition *IZkSyncStateTransitionTransactorSession) SetPriorityTxMaxGasLimit(_newPriorityTxMaxGasLimit *big.Int) (*types.Transaction, error) {
	return _IZkSyncStateTransition.Contract.SetPriorityTxMaxGasLimit(&_IZkSyncStateTransition.TransactOpts, _newPriorityTxMaxGasLimit)
}

// SetTokenMultiplier is a paid mutator transaction binding the contract method 0x235d9eb5.
//
// Solidity: function setTokenMultiplier(uint128 _nominator, uint128 _denominator) returns()
func (_IZkSyncStateTransition *IZkSyncStateTransitionTransactor) SetTokenMultiplier(opts *bind.TransactOpts, _nominator *big.Int, _denominator *big.Int) (*types.Transaction, error) {
	return _IZkSyncStateTransition.contract.Transact(opts, "setTokenMultiplier", _nominator, _denominator)
}

// SetTokenMultiplier is a paid mutator transaction binding the contract method 0x235d9eb5.
//
// Solidity: function setTokenMultiplier(uint128 _nominator, uint128 _denominator) returns()
func (_IZkSyncStateTransition *IZkSyncStateTransitionSession) SetTokenMultiplier(_nominator *big.Int, _denominator *big.Int) (*types.Transaction, error) {
	return _IZkSyncStateTransition.Contract.SetTokenMultiplier(&_IZkSyncStateTransition.TransactOpts, _nominator, _denominator)
}

// SetTokenMultiplier is a paid mutator transaction binding the contract method 0x235d9eb5.
//
// Solidity: function setTokenMultiplier(uint128 _nominator, uint128 _denominator) returns()
func (_IZkSyncStateTransition *IZkSyncStateTransitionTransactorSession) SetTokenMultiplier(_nominator *big.Int, _denominator *big.Int) (*types.Transaction, error) {
	return _IZkSyncStateTransition.Contract.SetTokenMultiplier(&_IZkSyncStateTransition.TransactOpts, _nominator, _denominator)
}

// SetTransactionFilterer is a paid mutator transaction binding the contract method 0x21f603d7.
//
// Solidity: function setTransactionFilterer(address _transactionFilterer) returns()
func (_IZkSyncStateTransition *IZkSyncStateTransitionTransactor) SetTransactionFilterer(opts *bind.TransactOpts, _transactionFilterer common.Address) (*types.Transaction, error) {
	return _IZkSyncStateTransition.contract.Transact(opts, "setTransactionFilterer", _transactionFilterer)
}

// SetTransactionFilterer is a paid mutator transaction binding the contract method 0x21f603d7.
//
// Solidity: function setTransactionFilterer(address _transactionFilterer) returns()
func (_IZkSyncStateTransition *IZkSyncStateTransitionSession) SetTransactionFilterer(_transactionFilterer common.Address) (*types.Transaction, error) {
	return _IZkSyncStateTransition.Contract.SetTransactionFilterer(&_IZkSyncStateTransition.TransactOpts, _transactionFilterer)
}

// SetTransactionFilterer is a paid mutator transaction binding the contract method 0x21f603d7.
//
// Solidity: function setTransactionFilterer(address _transactionFilterer) returns()
func (_IZkSyncStateTransition *IZkSyncStateTransitionTransactorSession) SetTransactionFilterer(_transactionFilterer common.Address) (*types.Transaction, error) {
	return _IZkSyncStateTransition.Contract.SetTransactionFilterer(&_IZkSyncStateTransition.TransactOpts, _transactionFilterer)
}

// SetValidator is a paid mutator transaction binding the contract method 0x4623c91d.
//
// Solidity: function setValidator(address _validator, bool _active) returns()
func (_IZkSyncStateTransition *IZkSyncStateTransitionTransactor) SetValidator(opts *bind.TransactOpts, _validator common.Address, _active bool) (*types.Transaction, error) {
	return _IZkSyncStateTransition.contract.Transact(opts, "setValidator", _validator, _active)
}

// SetValidator is a paid mutator transaction binding the contract method 0x4623c91d.
//
// Solidity: function setValidator(address _validator, bool _active) returns()
func (_IZkSyncStateTransition *IZkSyncStateTransitionSession) SetValidator(_validator common.Address, _active bool) (*types.Transaction, error) {
	return _IZkSyncStateTransition.Contract.SetValidator(&_IZkSyncStateTransition.TransactOpts, _validator, _active)
}

// SetValidator is a paid mutator transaction binding the contract method 0x4623c91d.
//
// Solidity: function setValidator(address _validator, bool _active) returns()
func (_IZkSyncStateTransition *IZkSyncStateTransitionTransactorSession) SetValidator(_validator common.Address, _active bool) (*types.Transaction, error) {
	return _IZkSyncStateTransition.Contract.SetValidator(&_IZkSyncStateTransition.TransactOpts, _validator, _active)
}

// SetValidiumMode is a paid mutator transaction binding the contract method 0x87c905b7.
//
// Solidity: function setValidiumMode(uint8 _validiumMode) returns()
func (_IZkSyncStateTransition *IZkSyncStateTransitionTransactor) SetValidiumMode(opts *bind.TransactOpts, _validiumMode uint8) (*types.Transaction, error) {
	return _IZkSyncStateTransition.contract.Transact(opts, "setValidiumMode", _validiumMode)
}

// SetValidiumMode is a paid mutator transaction binding the contract method 0x87c905b7.
//
// Solidity: function setValidiumMode(uint8 _validiumMode) returns()
func (_IZkSyncStateTransition *IZkSyncStateTransitionSession) SetValidiumMode(_validiumMode uint8) (*types.Transaction, error) {
	return _IZkSyncStateTransition.Contract.SetValidiumMode(&_IZkSyncStateTransition.TransactOpts, _validiumMode)
}

// SetValidiumMode is a paid mutator transaction binding the contract method 0x87c905b7.
//
// Solidity: function setValidiumMode(uint8 _validiumMode) returns()
func (_IZkSyncStateTransition *IZkSyncStateTransitionTransactorSession) SetValidiumMode(_validiumMode uint8) (*types.Transaction, error) {
	return _IZkSyncStateTransition.Contract.SetValidiumMode(&_IZkSyncStateTransition.TransactOpts, _validiumMode)
}

// TransferEthToSharedBridge is a paid mutator transaction binding the contract method 0xc924de35.
//
// Solidity: function transferEthToSharedBridge() returns()
func (_IZkSyncStateTransition *IZkSyncStateTransitionTransactor) TransferEthToSharedBridge(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IZkSyncStateTransition.contract.Transact(opts, "transferEthToSharedBridge")
}

// TransferEthToSharedBridge is a paid mutator transaction binding the contract method 0xc924de35.
//
// Solidity: function transferEthToSharedBridge() returns()
func (_IZkSyncStateTransition *IZkSyncStateTransitionSession) TransferEthToSharedBridge() (*types.Transaction, error) {
	return _IZkSyncStateTransition.Contract.TransferEthToSharedBridge(&_IZkSyncStateTransition.TransactOpts)
}

// TransferEthToSharedBridge is a paid mutator transaction binding the contract method 0xc924de35.
//
// Solidity: function transferEthToSharedBridge() returns()
func (_IZkSyncStateTransition *IZkSyncStateTransitionTransactorSession) TransferEthToSharedBridge() (*types.Transaction, error) {
	return _IZkSyncStateTransition.Contract.TransferEthToSharedBridge(&_IZkSyncStateTransition.TransactOpts)
}

// UnfreezeDiamond is a paid mutator transaction binding the contract method 0x17338945.
//
// Solidity: function unfreezeDiamond() returns()
func (_IZkSyncStateTransition *IZkSyncStateTransitionTransactor) UnfreezeDiamond(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IZkSyncStateTransition.contract.Transact(opts, "unfreezeDiamond")
}

// UnfreezeDiamond is a paid mutator transaction binding the contract method 0x17338945.
//
// Solidity: function unfreezeDiamond() returns()
func (_IZkSyncStateTransition *IZkSyncStateTransitionSession) UnfreezeDiamond() (*types.Transaction, error) {
	return _IZkSyncStateTransition.Contract.UnfreezeDiamond(&_IZkSyncStateTransition.TransactOpts)
}

// UnfreezeDiamond is a paid mutator transaction binding the contract method 0x17338945.
//
// Solidity: function unfreezeDiamond() returns()
func (_IZkSyncStateTransition *IZkSyncStateTransitionTransactorSession) UnfreezeDiamond() (*types.Transaction, error) {
	return _IZkSyncStateTransition.Contract.UnfreezeDiamond(&_IZkSyncStateTransition.TransactOpts)
}

// UpgradeChainFromVersion is a paid mutator transaction binding the contract method 0xfc57565f.
//
// Solidity: function upgradeChainFromVersion(uint256 _protocolVersion, ((address,uint8,bool,bytes4[])[],address,bytes) _cutData) returns()
func (_IZkSyncStateTransition *IZkSyncStateTransitionTransactor) UpgradeChainFromVersion(opts *bind.TransactOpts, _protocolVersion *big.Int, _cutData DiamondDiamondCutData) (*types.Transaction, error) {
	return _IZkSyncStateTransition.contract.Transact(opts, "upgradeChainFromVersion", _protocolVersion, _cutData)
}

// UpgradeChainFromVersion is a paid mutator transaction binding the contract method 0xfc57565f.
//
// Solidity: function upgradeChainFromVersion(uint256 _protocolVersion, ((address,uint8,bool,bytes4[])[],address,bytes) _cutData) returns()
func (_IZkSyncStateTransition *IZkSyncStateTransitionSession) UpgradeChainFromVersion(_protocolVersion *big.Int, _cutData DiamondDiamondCutData) (*types.Transaction, error) {
	return _IZkSyncStateTransition.Contract.UpgradeChainFromVersion(&_IZkSyncStateTransition.TransactOpts, _protocolVersion, _cutData)
}

// UpgradeChainFromVersion is a paid mutator transaction binding the contract method 0xfc57565f.
//
// Solidity: function upgradeChainFromVersion(uint256 _protocolVersion, ((address,uint8,bool,bytes4[])[],address,bytes) _cutData) returns()
func (_IZkSyncStateTransition *IZkSyncStateTransitionTransactorSession) UpgradeChainFromVersion(_protocolVersion *big.Int, _cutData DiamondDiamondCutData) (*types.Transaction, error) {
	return _IZkSyncStateTransition.Contract.UpgradeChainFromVersion(&_IZkSyncStateTransition.TransactOpts, _protocolVersion, _cutData)
}

// IZkSyncStateTransitionBlockCommitIterator is returned from FilterBlockCommit and is used to iterate over the raw logs and unpacked data for BlockCommit events raised by the IZkSyncStateTransition contract.
type IZkSyncStateTransitionBlockCommitIterator struct {
	Event *IZkSyncStateTransitionBlockCommit // Event containing the contract specifics and raw log

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
func (it *IZkSyncStateTransitionBlockCommitIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IZkSyncStateTransitionBlockCommit)
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
		it.Event = new(IZkSyncStateTransitionBlockCommit)
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
func (it *IZkSyncStateTransitionBlockCommitIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IZkSyncStateTransitionBlockCommitIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IZkSyncStateTransitionBlockCommit represents a BlockCommit event raised by the IZkSyncStateTransition contract.
type IZkSyncStateTransitionBlockCommit struct {
	BatchNumber *big.Int
	BatchHash   [32]byte
	Commitment  [32]byte
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterBlockCommit is a free log retrieval operation binding the contract event 0x8f2916b2f2d78cc5890ead36c06c0f6d5d112c7e103589947e8e2f0d6eddb763.
//
// Solidity: event BlockCommit(uint256 indexed batchNumber, bytes32 indexed batchHash, bytes32 indexed commitment)
func (_IZkSyncStateTransition *IZkSyncStateTransitionFilterer) FilterBlockCommit(opts *bind.FilterOpts, batchNumber []*big.Int, batchHash [][32]byte, commitment [][32]byte) (*IZkSyncStateTransitionBlockCommitIterator, error) {

	var batchNumberRule []interface{}
	for _, batchNumberItem := range batchNumber {
		batchNumberRule = append(batchNumberRule, batchNumberItem)
	}
	var batchHashRule []interface{}
	for _, batchHashItem := range batchHash {
		batchHashRule = append(batchHashRule, batchHashItem)
	}
	var commitmentRule []interface{}
	for _, commitmentItem := range commitment {
		commitmentRule = append(commitmentRule, commitmentItem)
	}

	logs, sub, err := _IZkSyncStateTransition.contract.FilterLogs(opts, "BlockCommit", batchNumberRule, batchHashRule, commitmentRule)
	if err != nil {
		return nil, err
	}
	return &IZkSyncStateTransitionBlockCommitIterator{contract: _IZkSyncStateTransition.contract, event: "BlockCommit", logs: logs, sub: sub}, nil
}

// WatchBlockCommit is a free log subscription operation binding the contract event 0x8f2916b2f2d78cc5890ead36c06c0f6d5d112c7e103589947e8e2f0d6eddb763.
//
// Solidity: event BlockCommit(uint256 indexed batchNumber, bytes32 indexed batchHash, bytes32 indexed commitment)
func (_IZkSyncStateTransition *IZkSyncStateTransitionFilterer) WatchBlockCommit(opts *bind.WatchOpts, sink chan<- *IZkSyncStateTransitionBlockCommit, batchNumber []*big.Int, batchHash [][32]byte, commitment [][32]byte) (event.Subscription, error) {

	var batchNumberRule []interface{}
	for _, batchNumberItem := range batchNumber {
		batchNumberRule = append(batchNumberRule, batchNumberItem)
	}
	var batchHashRule []interface{}
	for _, batchHashItem := range batchHash {
		batchHashRule = append(batchHashRule, batchHashItem)
	}
	var commitmentRule []interface{}
	for _, commitmentItem := range commitment {
		commitmentRule = append(commitmentRule, commitmentItem)
	}

	logs, sub, err := _IZkSyncStateTransition.contract.WatchLogs(opts, "BlockCommit", batchNumberRule, batchHashRule, commitmentRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IZkSyncStateTransitionBlockCommit)
				if err := _IZkSyncStateTransition.contract.UnpackLog(event, "BlockCommit", log); err != nil {
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

// ParseBlockCommit is a log parse operation binding the contract event 0x8f2916b2f2d78cc5890ead36c06c0f6d5d112c7e103589947e8e2f0d6eddb763.
//
// Solidity: event BlockCommit(uint256 indexed batchNumber, bytes32 indexed batchHash, bytes32 indexed commitment)
func (_IZkSyncStateTransition *IZkSyncStateTransitionFilterer) ParseBlockCommit(log types.Log) (*IZkSyncStateTransitionBlockCommit, error) {
	event := new(IZkSyncStateTransitionBlockCommit)
	if err := _IZkSyncStateTransition.contract.UnpackLog(event, "BlockCommit", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// IZkSyncStateTransitionBlockExecutionIterator is returned from FilterBlockExecution and is used to iterate over the raw logs and unpacked data for BlockExecution events raised by the IZkSyncStateTransition contract.
type IZkSyncStateTransitionBlockExecutionIterator struct {
	Event *IZkSyncStateTransitionBlockExecution // Event containing the contract specifics and raw log

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
func (it *IZkSyncStateTransitionBlockExecutionIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IZkSyncStateTransitionBlockExecution)
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
		it.Event = new(IZkSyncStateTransitionBlockExecution)
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
func (it *IZkSyncStateTransitionBlockExecutionIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IZkSyncStateTransitionBlockExecutionIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IZkSyncStateTransitionBlockExecution represents a BlockExecution event raised by the IZkSyncStateTransition contract.
type IZkSyncStateTransitionBlockExecution struct {
	BatchNumber *big.Int
	BatchHash   [32]byte
	Commitment  [32]byte
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterBlockExecution is a free log retrieval operation binding the contract event 0x2402307311a4d6604e4e7b4c8a15a7e1213edb39c16a31efa70afb06030d3165.
//
// Solidity: event BlockExecution(uint256 indexed batchNumber, bytes32 indexed batchHash, bytes32 indexed commitment)
func (_IZkSyncStateTransition *IZkSyncStateTransitionFilterer) FilterBlockExecution(opts *bind.FilterOpts, batchNumber []*big.Int, batchHash [][32]byte, commitment [][32]byte) (*IZkSyncStateTransitionBlockExecutionIterator, error) {

	var batchNumberRule []interface{}
	for _, batchNumberItem := range batchNumber {
		batchNumberRule = append(batchNumberRule, batchNumberItem)
	}
	var batchHashRule []interface{}
	for _, batchHashItem := range batchHash {
		batchHashRule = append(batchHashRule, batchHashItem)
	}
	var commitmentRule []interface{}
	for _, commitmentItem := range commitment {
		commitmentRule = append(commitmentRule, commitmentItem)
	}

	logs, sub, err := _IZkSyncStateTransition.contract.FilterLogs(opts, "BlockExecution", batchNumberRule, batchHashRule, commitmentRule)
	if err != nil {
		return nil, err
	}
	return &IZkSyncStateTransitionBlockExecutionIterator{contract: _IZkSyncStateTransition.contract, event: "BlockExecution", logs: logs, sub: sub}, nil
}

// WatchBlockExecution is a free log subscription operation binding the contract event 0x2402307311a4d6604e4e7b4c8a15a7e1213edb39c16a31efa70afb06030d3165.
//
// Solidity: event BlockExecution(uint256 indexed batchNumber, bytes32 indexed batchHash, bytes32 indexed commitment)
func (_IZkSyncStateTransition *IZkSyncStateTransitionFilterer) WatchBlockExecution(opts *bind.WatchOpts, sink chan<- *IZkSyncStateTransitionBlockExecution, batchNumber []*big.Int, batchHash [][32]byte, commitment [][32]byte) (event.Subscription, error) {

	var batchNumberRule []interface{}
	for _, batchNumberItem := range batchNumber {
		batchNumberRule = append(batchNumberRule, batchNumberItem)
	}
	var batchHashRule []interface{}
	for _, batchHashItem := range batchHash {
		batchHashRule = append(batchHashRule, batchHashItem)
	}
	var commitmentRule []interface{}
	for _, commitmentItem := range commitment {
		commitmentRule = append(commitmentRule, commitmentItem)
	}

	logs, sub, err := _IZkSyncStateTransition.contract.WatchLogs(opts, "BlockExecution", batchNumberRule, batchHashRule, commitmentRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IZkSyncStateTransitionBlockExecution)
				if err := _IZkSyncStateTransition.contract.UnpackLog(event, "BlockExecution", log); err != nil {
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

// ParseBlockExecution is a log parse operation binding the contract event 0x2402307311a4d6604e4e7b4c8a15a7e1213edb39c16a31efa70afb06030d3165.
//
// Solidity: event BlockExecution(uint256 indexed batchNumber, bytes32 indexed batchHash, bytes32 indexed commitment)
func (_IZkSyncStateTransition *IZkSyncStateTransitionFilterer) ParseBlockExecution(log types.Log) (*IZkSyncStateTransitionBlockExecution, error) {
	event := new(IZkSyncStateTransitionBlockExecution)
	if err := _IZkSyncStateTransition.contract.UnpackLog(event, "BlockExecution", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// IZkSyncStateTransitionBlocksRevertIterator is returned from FilterBlocksRevert and is used to iterate over the raw logs and unpacked data for BlocksRevert events raised by the IZkSyncStateTransition contract.
type IZkSyncStateTransitionBlocksRevertIterator struct {
	Event *IZkSyncStateTransitionBlocksRevert // Event containing the contract specifics and raw log

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
func (it *IZkSyncStateTransitionBlocksRevertIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IZkSyncStateTransitionBlocksRevert)
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
		it.Event = new(IZkSyncStateTransitionBlocksRevert)
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
func (it *IZkSyncStateTransitionBlocksRevertIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IZkSyncStateTransitionBlocksRevertIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IZkSyncStateTransitionBlocksRevert represents a BlocksRevert event raised by the IZkSyncStateTransition contract.
type IZkSyncStateTransitionBlocksRevert struct {
	TotalBatchesCommitted *big.Int
	TotalBatchesVerified  *big.Int
	TotalBatchesExecuted  *big.Int
	Raw                   types.Log // Blockchain specific contextual infos
}

// FilterBlocksRevert is a free log retrieval operation binding the contract event 0x8bd4b15ea7d1bc41ea9abc3fc487ccb89cd678a00786584714faa9d751c84ee5.
//
// Solidity: event BlocksRevert(uint256 totalBatchesCommitted, uint256 totalBatchesVerified, uint256 totalBatchesExecuted)
func (_IZkSyncStateTransition *IZkSyncStateTransitionFilterer) FilterBlocksRevert(opts *bind.FilterOpts) (*IZkSyncStateTransitionBlocksRevertIterator, error) {

	logs, sub, err := _IZkSyncStateTransition.contract.FilterLogs(opts, "BlocksRevert")
	if err != nil {
		return nil, err
	}
	return &IZkSyncStateTransitionBlocksRevertIterator{contract: _IZkSyncStateTransition.contract, event: "BlocksRevert", logs: logs, sub: sub}, nil
}

// WatchBlocksRevert is a free log subscription operation binding the contract event 0x8bd4b15ea7d1bc41ea9abc3fc487ccb89cd678a00786584714faa9d751c84ee5.
//
// Solidity: event BlocksRevert(uint256 totalBatchesCommitted, uint256 totalBatchesVerified, uint256 totalBatchesExecuted)
func (_IZkSyncStateTransition *IZkSyncStateTransitionFilterer) WatchBlocksRevert(opts *bind.WatchOpts, sink chan<- *IZkSyncStateTransitionBlocksRevert) (event.Subscription, error) {

	logs, sub, err := _IZkSyncStateTransition.contract.WatchLogs(opts, "BlocksRevert")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IZkSyncStateTransitionBlocksRevert)
				if err := _IZkSyncStateTransition.contract.UnpackLog(event, "BlocksRevert", log); err != nil {
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

// ParseBlocksRevert is a log parse operation binding the contract event 0x8bd4b15ea7d1bc41ea9abc3fc487ccb89cd678a00786584714faa9d751c84ee5.
//
// Solidity: event BlocksRevert(uint256 totalBatchesCommitted, uint256 totalBatchesVerified, uint256 totalBatchesExecuted)
func (_IZkSyncStateTransition *IZkSyncStateTransitionFilterer) ParseBlocksRevert(log types.Log) (*IZkSyncStateTransitionBlocksRevert, error) {
	event := new(IZkSyncStateTransitionBlocksRevert)
	if err := _IZkSyncStateTransition.contract.UnpackLog(event, "BlocksRevert", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// IZkSyncStateTransitionBlocksVerificationIterator is returned from FilterBlocksVerification and is used to iterate over the raw logs and unpacked data for BlocksVerification events raised by the IZkSyncStateTransition contract.
type IZkSyncStateTransitionBlocksVerificationIterator struct {
	Event *IZkSyncStateTransitionBlocksVerification // Event containing the contract specifics and raw log

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
func (it *IZkSyncStateTransitionBlocksVerificationIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IZkSyncStateTransitionBlocksVerification)
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
		it.Event = new(IZkSyncStateTransitionBlocksVerification)
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
func (it *IZkSyncStateTransitionBlocksVerificationIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IZkSyncStateTransitionBlocksVerificationIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IZkSyncStateTransitionBlocksVerification represents a BlocksVerification event raised by the IZkSyncStateTransition contract.
type IZkSyncStateTransitionBlocksVerification struct {
	PreviousLastVerifiedBatch *big.Int
	CurrentLastVerifiedBatch  *big.Int
	Raw                       types.Log // Blockchain specific contextual infos
}

// FilterBlocksVerification is a free log retrieval operation binding the contract event 0x22c9005dd88c18b552a1cd7e8b3b937fcde9ca69213c1f658f54d572e4877a81.
//
// Solidity: event BlocksVerification(uint256 indexed previousLastVerifiedBatch, uint256 indexed currentLastVerifiedBatch)
func (_IZkSyncStateTransition *IZkSyncStateTransitionFilterer) FilterBlocksVerification(opts *bind.FilterOpts, previousLastVerifiedBatch []*big.Int, currentLastVerifiedBatch []*big.Int) (*IZkSyncStateTransitionBlocksVerificationIterator, error) {

	var previousLastVerifiedBatchRule []interface{}
	for _, previousLastVerifiedBatchItem := range previousLastVerifiedBatch {
		previousLastVerifiedBatchRule = append(previousLastVerifiedBatchRule, previousLastVerifiedBatchItem)
	}
	var currentLastVerifiedBatchRule []interface{}
	for _, currentLastVerifiedBatchItem := range currentLastVerifiedBatch {
		currentLastVerifiedBatchRule = append(currentLastVerifiedBatchRule, currentLastVerifiedBatchItem)
	}

	logs, sub, err := _IZkSyncStateTransition.contract.FilterLogs(opts, "BlocksVerification", previousLastVerifiedBatchRule, currentLastVerifiedBatchRule)
	if err != nil {
		return nil, err
	}
	return &IZkSyncStateTransitionBlocksVerificationIterator{contract: _IZkSyncStateTransition.contract, event: "BlocksVerification", logs: logs, sub: sub}, nil
}

// WatchBlocksVerification is a free log subscription operation binding the contract event 0x22c9005dd88c18b552a1cd7e8b3b937fcde9ca69213c1f658f54d572e4877a81.
//
// Solidity: event BlocksVerification(uint256 indexed previousLastVerifiedBatch, uint256 indexed currentLastVerifiedBatch)
func (_IZkSyncStateTransition *IZkSyncStateTransitionFilterer) WatchBlocksVerification(opts *bind.WatchOpts, sink chan<- *IZkSyncStateTransitionBlocksVerification, previousLastVerifiedBatch []*big.Int, currentLastVerifiedBatch []*big.Int) (event.Subscription, error) {

	var previousLastVerifiedBatchRule []interface{}
	for _, previousLastVerifiedBatchItem := range previousLastVerifiedBatch {
		previousLastVerifiedBatchRule = append(previousLastVerifiedBatchRule, previousLastVerifiedBatchItem)
	}
	var currentLastVerifiedBatchRule []interface{}
	for _, currentLastVerifiedBatchItem := range currentLastVerifiedBatch {
		currentLastVerifiedBatchRule = append(currentLastVerifiedBatchRule, currentLastVerifiedBatchItem)
	}

	logs, sub, err := _IZkSyncStateTransition.contract.WatchLogs(opts, "BlocksVerification", previousLastVerifiedBatchRule, currentLastVerifiedBatchRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IZkSyncStateTransitionBlocksVerification)
				if err := _IZkSyncStateTransition.contract.UnpackLog(event, "BlocksVerification", log); err != nil {
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

// ParseBlocksVerification is a log parse operation binding the contract event 0x22c9005dd88c18b552a1cd7e8b3b937fcde9ca69213c1f658f54d572e4877a81.
//
// Solidity: event BlocksVerification(uint256 indexed previousLastVerifiedBatch, uint256 indexed currentLastVerifiedBatch)
func (_IZkSyncStateTransition *IZkSyncStateTransitionFilterer) ParseBlocksVerification(log types.Log) (*IZkSyncStateTransitionBlocksVerification, error) {
	event := new(IZkSyncStateTransitionBlocksVerification)
	if err := _IZkSyncStateTransition.contract.UnpackLog(event, "BlocksVerification", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// IZkSyncStateTransitionEthWithdrawalFinalizedIterator is returned from FilterEthWithdrawalFinalized and is used to iterate over the raw logs and unpacked data for EthWithdrawalFinalized events raised by the IZkSyncStateTransition contract.
type IZkSyncStateTransitionEthWithdrawalFinalizedIterator struct {
	Event *IZkSyncStateTransitionEthWithdrawalFinalized // Event containing the contract specifics and raw log

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
func (it *IZkSyncStateTransitionEthWithdrawalFinalizedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IZkSyncStateTransitionEthWithdrawalFinalized)
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
		it.Event = new(IZkSyncStateTransitionEthWithdrawalFinalized)
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
func (it *IZkSyncStateTransitionEthWithdrawalFinalizedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IZkSyncStateTransitionEthWithdrawalFinalizedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IZkSyncStateTransitionEthWithdrawalFinalized represents a EthWithdrawalFinalized event raised by the IZkSyncStateTransition contract.
type IZkSyncStateTransitionEthWithdrawalFinalized struct {
	To     common.Address
	Amount *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterEthWithdrawalFinalized is a free log retrieval operation binding the contract event 0x26464d64ddb13f6d187de632d165bd1065382ec0b66c25c648957116e7bc25c8.
//
// Solidity: event EthWithdrawalFinalized(address indexed to, uint256 amount)
func (_IZkSyncStateTransition *IZkSyncStateTransitionFilterer) FilterEthWithdrawalFinalized(opts *bind.FilterOpts, to []common.Address) (*IZkSyncStateTransitionEthWithdrawalFinalizedIterator, error) {

	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _IZkSyncStateTransition.contract.FilterLogs(opts, "EthWithdrawalFinalized", toRule)
	if err != nil {
		return nil, err
	}
	return &IZkSyncStateTransitionEthWithdrawalFinalizedIterator{contract: _IZkSyncStateTransition.contract, event: "EthWithdrawalFinalized", logs: logs, sub: sub}, nil
}

// WatchEthWithdrawalFinalized is a free log subscription operation binding the contract event 0x26464d64ddb13f6d187de632d165bd1065382ec0b66c25c648957116e7bc25c8.
//
// Solidity: event EthWithdrawalFinalized(address indexed to, uint256 amount)
func (_IZkSyncStateTransition *IZkSyncStateTransitionFilterer) WatchEthWithdrawalFinalized(opts *bind.WatchOpts, sink chan<- *IZkSyncStateTransitionEthWithdrawalFinalized, to []common.Address) (event.Subscription, error) {

	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _IZkSyncStateTransition.contract.WatchLogs(opts, "EthWithdrawalFinalized", toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IZkSyncStateTransitionEthWithdrawalFinalized)
				if err := _IZkSyncStateTransition.contract.UnpackLog(event, "EthWithdrawalFinalized", log); err != nil {
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

// ParseEthWithdrawalFinalized is a log parse operation binding the contract event 0x26464d64ddb13f6d187de632d165bd1065382ec0b66c25c648957116e7bc25c8.
//
// Solidity: event EthWithdrawalFinalized(address indexed to, uint256 amount)
func (_IZkSyncStateTransition *IZkSyncStateTransitionFilterer) ParseEthWithdrawalFinalized(log types.Log) (*IZkSyncStateTransitionEthWithdrawalFinalized, error) {
	event := new(IZkSyncStateTransitionEthWithdrawalFinalized)
	if err := _IZkSyncStateTransition.contract.UnpackLog(event, "EthWithdrawalFinalized", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// IZkSyncStateTransitionExecuteUpgradeIterator is returned from FilterExecuteUpgrade and is used to iterate over the raw logs and unpacked data for ExecuteUpgrade events raised by the IZkSyncStateTransition contract.
type IZkSyncStateTransitionExecuteUpgradeIterator struct {
	Event *IZkSyncStateTransitionExecuteUpgrade // Event containing the contract specifics and raw log

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
func (it *IZkSyncStateTransitionExecuteUpgradeIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IZkSyncStateTransitionExecuteUpgrade)
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
		it.Event = new(IZkSyncStateTransitionExecuteUpgrade)
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
func (it *IZkSyncStateTransitionExecuteUpgradeIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IZkSyncStateTransitionExecuteUpgradeIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IZkSyncStateTransitionExecuteUpgrade represents a ExecuteUpgrade event raised by the IZkSyncStateTransition contract.
type IZkSyncStateTransitionExecuteUpgrade struct {
	DiamondCut DiamondDiamondCutData
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterExecuteUpgrade is a free log retrieval operation binding the contract event 0x1dabfc3f4f6a4e74e19be10cc9d1d8e23db03e415d5d3547a1a7d5eb766513ba.
//
// Solidity: event ExecuteUpgrade(((address,uint8,bool,bytes4[])[],address,bytes) diamondCut)
func (_IZkSyncStateTransition *IZkSyncStateTransitionFilterer) FilterExecuteUpgrade(opts *bind.FilterOpts) (*IZkSyncStateTransitionExecuteUpgradeIterator, error) {

	logs, sub, err := _IZkSyncStateTransition.contract.FilterLogs(opts, "ExecuteUpgrade")
	if err != nil {
		return nil, err
	}
	return &IZkSyncStateTransitionExecuteUpgradeIterator{contract: _IZkSyncStateTransition.contract, event: "ExecuteUpgrade", logs: logs, sub: sub}, nil
}

// WatchExecuteUpgrade is a free log subscription operation binding the contract event 0x1dabfc3f4f6a4e74e19be10cc9d1d8e23db03e415d5d3547a1a7d5eb766513ba.
//
// Solidity: event ExecuteUpgrade(((address,uint8,bool,bytes4[])[],address,bytes) diamondCut)
func (_IZkSyncStateTransition *IZkSyncStateTransitionFilterer) WatchExecuteUpgrade(opts *bind.WatchOpts, sink chan<- *IZkSyncStateTransitionExecuteUpgrade) (event.Subscription, error) {

	logs, sub, err := _IZkSyncStateTransition.contract.WatchLogs(opts, "ExecuteUpgrade")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IZkSyncStateTransitionExecuteUpgrade)
				if err := _IZkSyncStateTransition.contract.UnpackLog(event, "ExecuteUpgrade", log); err != nil {
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

// ParseExecuteUpgrade is a log parse operation binding the contract event 0x1dabfc3f4f6a4e74e19be10cc9d1d8e23db03e415d5d3547a1a7d5eb766513ba.
//
// Solidity: event ExecuteUpgrade(((address,uint8,bool,bytes4[])[],address,bytes) diamondCut)
func (_IZkSyncStateTransition *IZkSyncStateTransitionFilterer) ParseExecuteUpgrade(log types.Log) (*IZkSyncStateTransitionExecuteUpgrade, error) {
	event := new(IZkSyncStateTransitionExecuteUpgrade)
	if err := _IZkSyncStateTransition.contract.UnpackLog(event, "ExecuteUpgrade", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// IZkSyncStateTransitionFreezeIterator is returned from FilterFreeze and is used to iterate over the raw logs and unpacked data for Freeze events raised by the IZkSyncStateTransition contract.
type IZkSyncStateTransitionFreezeIterator struct {
	Event *IZkSyncStateTransitionFreeze // Event containing the contract specifics and raw log

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
func (it *IZkSyncStateTransitionFreezeIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IZkSyncStateTransitionFreeze)
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
		it.Event = new(IZkSyncStateTransitionFreeze)
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
func (it *IZkSyncStateTransitionFreezeIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IZkSyncStateTransitionFreezeIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IZkSyncStateTransitionFreeze represents a Freeze event raised by the IZkSyncStateTransition contract.
type IZkSyncStateTransitionFreeze struct {
	Raw types.Log // Blockchain specific contextual infos
}

// FilterFreeze is a free log retrieval operation binding the contract event 0x615acbaede366d76a8b8cb2a9ada6a71495f0786513d71aa97aaf0c3910b78de.
//
// Solidity: event Freeze()
func (_IZkSyncStateTransition *IZkSyncStateTransitionFilterer) FilterFreeze(opts *bind.FilterOpts) (*IZkSyncStateTransitionFreezeIterator, error) {

	logs, sub, err := _IZkSyncStateTransition.contract.FilterLogs(opts, "Freeze")
	if err != nil {
		return nil, err
	}
	return &IZkSyncStateTransitionFreezeIterator{contract: _IZkSyncStateTransition.contract, event: "Freeze", logs: logs, sub: sub}, nil
}

// WatchFreeze is a free log subscription operation binding the contract event 0x615acbaede366d76a8b8cb2a9ada6a71495f0786513d71aa97aaf0c3910b78de.
//
// Solidity: event Freeze()
func (_IZkSyncStateTransition *IZkSyncStateTransitionFilterer) WatchFreeze(opts *bind.WatchOpts, sink chan<- *IZkSyncStateTransitionFreeze) (event.Subscription, error) {

	logs, sub, err := _IZkSyncStateTransition.contract.WatchLogs(opts, "Freeze")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IZkSyncStateTransitionFreeze)
				if err := _IZkSyncStateTransition.contract.UnpackLog(event, "Freeze", log); err != nil {
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

// ParseFreeze is a log parse operation binding the contract event 0x615acbaede366d76a8b8cb2a9ada6a71495f0786513d71aa97aaf0c3910b78de.
//
// Solidity: event Freeze()
func (_IZkSyncStateTransition *IZkSyncStateTransitionFilterer) ParseFreeze(log types.Log) (*IZkSyncStateTransitionFreeze, error) {
	event := new(IZkSyncStateTransitionFreeze)
	if err := _IZkSyncStateTransition.contract.UnpackLog(event, "Freeze", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// IZkSyncStateTransitionIsPorterAvailableStatusUpdateIterator is returned from FilterIsPorterAvailableStatusUpdate and is used to iterate over the raw logs and unpacked data for IsPorterAvailableStatusUpdate events raised by the IZkSyncStateTransition contract.
type IZkSyncStateTransitionIsPorterAvailableStatusUpdateIterator struct {
	Event *IZkSyncStateTransitionIsPorterAvailableStatusUpdate // Event containing the contract specifics and raw log

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
func (it *IZkSyncStateTransitionIsPorterAvailableStatusUpdateIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IZkSyncStateTransitionIsPorterAvailableStatusUpdate)
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
		it.Event = new(IZkSyncStateTransitionIsPorterAvailableStatusUpdate)
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
func (it *IZkSyncStateTransitionIsPorterAvailableStatusUpdateIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IZkSyncStateTransitionIsPorterAvailableStatusUpdateIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IZkSyncStateTransitionIsPorterAvailableStatusUpdate represents a IsPorterAvailableStatusUpdate event raised by the IZkSyncStateTransition contract.
type IZkSyncStateTransitionIsPorterAvailableStatusUpdate struct {
	IsPorterAvailable bool
	Raw               types.Log // Blockchain specific contextual infos
}

// FilterIsPorterAvailableStatusUpdate is a free log retrieval operation binding the contract event 0x036b81a8a07344698cb5aa4142c5669a9317c9ce905264a08f0b9f9331883936.
//
// Solidity: event IsPorterAvailableStatusUpdate(bool isPorterAvailable)
func (_IZkSyncStateTransition *IZkSyncStateTransitionFilterer) FilterIsPorterAvailableStatusUpdate(opts *bind.FilterOpts) (*IZkSyncStateTransitionIsPorterAvailableStatusUpdateIterator, error) {

	logs, sub, err := _IZkSyncStateTransition.contract.FilterLogs(opts, "IsPorterAvailableStatusUpdate")
	if err != nil {
		return nil, err
	}
	return &IZkSyncStateTransitionIsPorterAvailableStatusUpdateIterator{contract: _IZkSyncStateTransition.contract, event: "IsPorterAvailableStatusUpdate", logs: logs, sub: sub}, nil
}

// WatchIsPorterAvailableStatusUpdate is a free log subscription operation binding the contract event 0x036b81a8a07344698cb5aa4142c5669a9317c9ce905264a08f0b9f9331883936.
//
// Solidity: event IsPorterAvailableStatusUpdate(bool isPorterAvailable)
func (_IZkSyncStateTransition *IZkSyncStateTransitionFilterer) WatchIsPorterAvailableStatusUpdate(opts *bind.WatchOpts, sink chan<- *IZkSyncStateTransitionIsPorterAvailableStatusUpdate) (event.Subscription, error) {

	logs, sub, err := _IZkSyncStateTransition.contract.WatchLogs(opts, "IsPorterAvailableStatusUpdate")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IZkSyncStateTransitionIsPorterAvailableStatusUpdate)
				if err := _IZkSyncStateTransition.contract.UnpackLog(event, "IsPorterAvailableStatusUpdate", log); err != nil {
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

// ParseIsPorterAvailableStatusUpdate is a log parse operation binding the contract event 0x036b81a8a07344698cb5aa4142c5669a9317c9ce905264a08f0b9f9331883936.
//
// Solidity: event IsPorterAvailableStatusUpdate(bool isPorterAvailable)
func (_IZkSyncStateTransition *IZkSyncStateTransitionFilterer) ParseIsPorterAvailableStatusUpdate(log types.Log) (*IZkSyncStateTransitionIsPorterAvailableStatusUpdate, error) {
	event := new(IZkSyncStateTransitionIsPorterAvailableStatusUpdate)
	if err := _IZkSyncStateTransition.contract.UnpackLog(event, "IsPorterAvailableStatusUpdate", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// IZkSyncStateTransitionNewAdminIterator is returned from FilterNewAdmin and is used to iterate over the raw logs and unpacked data for NewAdmin events raised by the IZkSyncStateTransition contract.
type IZkSyncStateTransitionNewAdminIterator struct {
	Event *IZkSyncStateTransitionNewAdmin // Event containing the contract specifics and raw log

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
func (it *IZkSyncStateTransitionNewAdminIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IZkSyncStateTransitionNewAdmin)
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
		it.Event = new(IZkSyncStateTransitionNewAdmin)
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
func (it *IZkSyncStateTransitionNewAdminIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IZkSyncStateTransitionNewAdminIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IZkSyncStateTransitionNewAdmin represents a NewAdmin event raised by the IZkSyncStateTransition contract.
type IZkSyncStateTransitionNewAdmin struct {
	OldAdmin common.Address
	NewAdmin common.Address
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterNewAdmin is a free log retrieval operation binding the contract event 0xf9ffabca9c8276e99321725bcb43fb076a6c66a54b7f21c4e8146d8519b417dc.
//
// Solidity: event NewAdmin(address indexed oldAdmin, address indexed newAdmin)
func (_IZkSyncStateTransition *IZkSyncStateTransitionFilterer) FilterNewAdmin(opts *bind.FilterOpts, oldAdmin []common.Address, newAdmin []common.Address) (*IZkSyncStateTransitionNewAdminIterator, error) {

	var oldAdminRule []interface{}
	for _, oldAdminItem := range oldAdmin {
		oldAdminRule = append(oldAdminRule, oldAdminItem)
	}
	var newAdminRule []interface{}
	for _, newAdminItem := range newAdmin {
		newAdminRule = append(newAdminRule, newAdminItem)
	}

	logs, sub, err := _IZkSyncStateTransition.contract.FilterLogs(opts, "NewAdmin", oldAdminRule, newAdminRule)
	if err != nil {
		return nil, err
	}
	return &IZkSyncStateTransitionNewAdminIterator{contract: _IZkSyncStateTransition.contract, event: "NewAdmin", logs: logs, sub: sub}, nil
}

// WatchNewAdmin is a free log subscription operation binding the contract event 0xf9ffabca9c8276e99321725bcb43fb076a6c66a54b7f21c4e8146d8519b417dc.
//
// Solidity: event NewAdmin(address indexed oldAdmin, address indexed newAdmin)
func (_IZkSyncStateTransition *IZkSyncStateTransitionFilterer) WatchNewAdmin(opts *bind.WatchOpts, sink chan<- *IZkSyncStateTransitionNewAdmin, oldAdmin []common.Address, newAdmin []common.Address) (event.Subscription, error) {

	var oldAdminRule []interface{}
	for _, oldAdminItem := range oldAdmin {
		oldAdminRule = append(oldAdminRule, oldAdminItem)
	}
	var newAdminRule []interface{}
	for _, newAdminItem := range newAdmin {
		newAdminRule = append(newAdminRule, newAdminItem)
	}

	logs, sub, err := _IZkSyncStateTransition.contract.WatchLogs(opts, "NewAdmin", oldAdminRule, newAdminRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IZkSyncStateTransitionNewAdmin)
				if err := _IZkSyncStateTransition.contract.UnpackLog(event, "NewAdmin", log); err != nil {
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

// ParseNewAdmin is a log parse operation binding the contract event 0xf9ffabca9c8276e99321725bcb43fb076a6c66a54b7f21c4e8146d8519b417dc.
//
// Solidity: event NewAdmin(address indexed oldAdmin, address indexed newAdmin)
func (_IZkSyncStateTransition *IZkSyncStateTransitionFilterer) ParseNewAdmin(log types.Log) (*IZkSyncStateTransitionNewAdmin, error) {
	event := new(IZkSyncStateTransitionNewAdmin)
	if err := _IZkSyncStateTransition.contract.UnpackLog(event, "NewAdmin", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// IZkSyncStateTransitionNewBaseTokenMultiplierIterator is returned from FilterNewBaseTokenMultiplier and is used to iterate over the raw logs and unpacked data for NewBaseTokenMultiplier events raised by the IZkSyncStateTransition contract.
type IZkSyncStateTransitionNewBaseTokenMultiplierIterator struct {
	Event *IZkSyncStateTransitionNewBaseTokenMultiplier // Event containing the contract specifics and raw log

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
func (it *IZkSyncStateTransitionNewBaseTokenMultiplierIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IZkSyncStateTransitionNewBaseTokenMultiplier)
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
		it.Event = new(IZkSyncStateTransitionNewBaseTokenMultiplier)
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
func (it *IZkSyncStateTransitionNewBaseTokenMultiplierIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IZkSyncStateTransitionNewBaseTokenMultiplierIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IZkSyncStateTransitionNewBaseTokenMultiplier represents a NewBaseTokenMultiplier event raised by the IZkSyncStateTransition contract.
type IZkSyncStateTransitionNewBaseTokenMultiplier struct {
	OldNominator   *big.Int
	OldDenominator *big.Int
	NewNominator   *big.Int
	NewDenominator *big.Int
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterNewBaseTokenMultiplier is a free log retrieval operation binding the contract event 0xc9cbdb110bd58a133e82c62b1488e57677be1f326df10a4d8096b5f170c370c8.
//
// Solidity: event NewBaseTokenMultiplier(uint128 oldNominator, uint128 oldDenominator, uint128 newNominator, uint128 newDenominator)
func (_IZkSyncStateTransition *IZkSyncStateTransitionFilterer) FilterNewBaseTokenMultiplier(opts *bind.FilterOpts) (*IZkSyncStateTransitionNewBaseTokenMultiplierIterator, error) {

	logs, sub, err := _IZkSyncStateTransition.contract.FilterLogs(opts, "NewBaseTokenMultiplier")
	if err != nil {
		return nil, err
	}
	return &IZkSyncStateTransitionNewBaseTokenMultiplierIterator{contract: _IZkSyncStateTransition.contract, event: "NewBaseTokenMultiplier", logs: logs, sub: sub}, nil
}

// WatchNewBaseTokenMultiplier is a free log subscription operation binding the contract event 0xc9cbdb110bd58a133e82c62b1488e57677be1f326df10a4d8096b5f170c370c8.
//
// Solidity: event NewBaseTokenMultiplier(uint128 oldNominator, uint128 oldDenominator, uint128 newNominator, uint128 newDenominator)
func (_IZkSyncStateTransition *IZkSyncStateTransitionFilterer) WatchNewBaseTokenMultiplier(opts *bind.WatchOpts, sink chan<- *IZkSyncStateTransitionNewBaseTokenMultiplier) (event.Subscription, error) {

	logs, sub, err := _IZkSyncStateTransition.contract.WatchLogs(opts, "NewBaseTokenMultiplier")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IZkSyncStateTransitionNewBaseTokenMultiplier)
				if err := _IZkSyncStateTransition.contract.UnpackLog(event, "NewBaseTokenMultiplier", log); err != nil {
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

// ParseNewBaseTokenMultiplier is a log parse operation binding the contract event 0xc9cbdb110bd58a133e82c62b1488e57677be1f326df10a4d8096b5f170c370c8.
//
// Solidity: event NewBaseTokenMultiplier(uint128 oldNominator, uint128 oldDenominator, uint128 newNominator, uint128 newDenominator)
func (_IZkSyncStateTransition *IZkSyncStateTransitionFilterer) ParseNewBaseTokenMultiplier(log types.Log) (*IZkSyncStateTransitionNewBaseTokenMultiplier, error) {
	event := new(IZkSyncStateTransitionNewBaseTokenMultiplier)
	if err := _IZkSyncStateTransition.contract.UnpackLog(event, "NewBaseTokenMultiplier", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// IZkSyncStateTransitionNewFeeParamsIterator is returned from FilterNewFeeParams and is used to iterate over the raw logs and unpacked data for NewFeeParams events raised by the IZkSyncStateTransition contract.
type IZkSyncStateTransitionNewFeeParamsIterator struct {
	Event *IZkSyncStateTransitionNewFeeParams // Event containing the contract specifics and raw log

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
func (it *IZkSyncStateTransitionNewFeeParamsIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IZkSyncStateTransitionNewFeeParams)
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
		it.Event = new(IZkSyncStateTransitionNewFeeParams)
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
func (it *IZkSyncStateTransitionNewFeeParamsIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IZkSyncStateTransitionNewFeeParamsIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IZkSyncStateTransitionNewFeeParams represents a NewFeeParams event raised by the IZkSyncStateTransition contract.
type IZkSyncStateTransitionNewFeeParams struct {
	OldFeeParams FeeParams
	NewFeeParams FeeParams
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterNewFeeParams is a free log retrieval operation binding the contract event 0xc8b245ac8b138b17b6b1dbbbb8860adc66b373afa000d99f3cdc775d8ae0bbed.
//
// Solidity: event NewFeeParams((uint8,uint32,uint32,uint32,uint32,uint64) oldFeeParams, (uint8,uint32,uint32,uint32,uint32,uint64) newFeeParams)
func (_IZkSyncStateTransition *IZkSyncStateTransitionFilterer) FilterNewFeeParams(opts *bind.FilterOpts) (*IZkSyncStateTransitionNewFeeParamsIterator, error) {

	logs, sub, err := _IZkSyncStateTransition.contract.FilterLogs(opts, "NewFeeParams")
	if err != nil {
		return nil, err
	}
	return &IZkSyncStateTransitionNewFeeParamsIterator{contract: _IZkSyncStateTransition.contract, event: "NewFeeParams", logs: logs, sub: sub}, nil
}

// WatchNewFeeParams is a free log subscription operation binding the contract event 0xc8b245ac8b138b17b6b1dbbbb8860adc66b373afa000d99f3cdc775d8ae0bbed.
//
// Solidity: event NewFeeParams((uint8,uint32,uint32,uint32,uint32,uint64) oldFeeParams, (uint8,uint32,uint32,uint32,uint32,uint64) newFeeParams)
func (_IZkSyncStateTransition *IZkSyncStateTransitionFilterer) WatchNewFeeParams(opts *bind.WatchOpts, sink chan<- *IZkSyncStateTransitionNewFeeParams) (event.Subscription, error) {

	logs, sub, err := _IZkSyncStateTransition.contract.WatchLogs(opts, "NewFeeParams")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IZkSyncStateTransitionNewFeeParams)
				if err := _IZkSyncStateTransition.contract.UnpackLog(event, "NewFeeParams", log); err != nil {
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

// ParseNewFeeParams is a log parse operation binding the contract event 0xc8b245ac8b138b17b6b1dbbbb8860adc66b373afa000d99f3cdc775d8ae0bbed.
//
// Solidity: event NewFeeParams((uint8,uint32,uint32,uint32,uint32,uint64) oldFeeParams, (uint8,uint32,uint32,uint32,uint32,uint64) newFeeParams)
func (_IZkSyncStateTransition *IZkSyncStateTransitionFilterer) ParseNewFeeParams(log types.Log) (*IZkSyncStateTransitionNewFeeParams, error) {
	event := new(IZkSyncStateTransitionNewFeeParams)
	if err := _IZkSyncStateTransition.contract.UnpackLog(event, "NewFeeParams", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// IZkSyncStateTransitionNewPendingAdminIterator is returned from FilterNewPendingAdmin and is used to iterate over the raw logs and unpacked data for NewPendingAdmin events raised by the IZkSyncStateTransition contract.
type IZkSyncStateTransitionNewPendingAdminIterator struct {
	Event *IZkSyncStateTransitionNewPendingAdmin // Event containing the contract specifics and raw log

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
func (it *IZkSyncStateTransitionNewPendingAdminIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IZkSyncStateTransitionNewPendingAdmin)
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
		it.Event = new(IZkSyncStateTransitionNewPendingAdmin)
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
func (it *IZkSyncStateTransitionNewPendingAdminIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IZkSyncStateTransitionNewPendingAdminIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IZkSyncStateTransitionNewPendingAdmin represents a NewPendingAdmin event raised by the IZkSyncStateTransition contract.
type IZkSyncStateTransitionNewPendingAdmin struct {
	OldPendingAdmin common.Address
	NewPendingAdmin common.Address
	Raw             types.Log // Blockchain specific contextual infos
}

// FilterNewPendingAdmin is a free log retrieval operation binding the contract event 0xca4f2f25d0898edd99413412fb94012f9e54ec8142f9b093e7720646a95b16a9.
//
// Solidity: event NewPendingAdmin(address indexed oldPendingAdmin, address indexed newPendingAdmin)
func (_IZkSyncStateTransition *IZkSyncStateTransitionFilterer) FilterNewPendingAdmin(opts *bind.FilterOpts, oldPendingAdmin []common.Address, newPendingAdmin []common.Address) (*IZkSyncStateTransitionNewPendingAdminIterator, error) {

	var oldPendingAdminRule []interface{}
	for _, oldPendingAdminItem := range oldPendingAdmin {
		oldPendingAdminRule = append(oldPendingAdminRule, oldPendingAdminItem)
	}
	var newPendingAdminRule []interface{}
	for _, newPendingAdminItem := range newPendingAdmin {
		newPendingAdminRule = append(newPendingAdminRule, newPendingAdminItem)
	}

	logs, sub, err := _IZkSyncStateTransition.contract.FilterLogs(opts, "NewPendingAdmin", oldPendingAdminRule, newPendingAdminRule)
	if err != nil {
		return nil, err
	}
	return &IZkSyncStateTransitionNewPendingAdminIterator{contract: _IZkSyncStateTransition.contract, event: "NewPendingAdmin", logs: logs, sub: sub}, nil
}

// WatchNewPendingAdmin is a free log subscription operation binding the contract event 0xca4f2f25d0898edd99413412fb94012f9e54ec8142f9b093e7720646a95b16a9.
//
// Solidity: event NewPendingAdmin(address indexed oldPendingAdmin, address indexed newPendingAdmin)
func (_IZkSyncStateTransition *IZkSyncStateTransitionFilterer) WatchNewPendingAdmin(opts *bind.WatchOpts, sink chan<- *IZkSyncStateTransitionNewPendingAdmin, oldPendingAdmin []common.Address, newPendingAdmin []common.Address) (event.Subscription, error) {

	var oldPendingAdminRule []interface{}
	for _, oldPendingAdminItem := range oldPendingAdmin {
		oldPendingAdminRule = append(oldPendingAdminRule, oldPendingAdminItem)
	}
	var newPendingAdminRule []interface{}
	for _, newPendingAdminItem := range newPendingAdmin {
		newPendingAdminRule = append(newPendingAdminRule, newPendingAdminItem)
	}

	logs, sub, err := _IZkSyncStateTransition.contract.WatchLogs(opts, "NewPendingAdmin", oldPendingAdminRule, newPendingAdminRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IZkSyncStateTransitionNewPendingAdmin)
				if err := _IZkSyncStateTransition.contract.UnpackLog(event, "NewPendingAdmin", log); err != nil {
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

// ParseNewPendingAdmin is a log parse operation binding the contract event 0xca4f2f25d0898edd99413412fb94012f9e54ec8142f9b093e7720646a95b16a9.
//
// Solidity: event NewPendingAdmin(address indexed oldPendingAdmin, address indexed newPendingAdmin)
func (_IZkSyncStateTransition *IZkSyncStateTransitionFilterer) ParseNewPendingAdmin(log types.Log) (*IZkSyncStateTransitionNewPendingAdmin, error) {
	event := new(IZkSyncStateTransitionNewPendingAdmin)
	if err := _IZkSyncStateTransition.contract.UnpackLog(event, "NewPendingAdmin", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// IZkSyncStateTransitionNewPriorityRequestIterator is returned from FilterNewPriorityRequest and is used to iterate over the raw logs and unpacked data for NewPriorityRequest events raised by the IZkSyncStateTransition contract.
type IZkSyncStateTransitionNewPriorityRequestIterator struct {
	Event *IZkSyncStateTransitionNewPriorityRequest // Event containing the contract specifics and raw log

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
func (it *IZkSyncStateTransitionNewPriorityRequestIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IZkSyncStateTransitionNewPriorityRequest)
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
		it.Event = new(IZkSyncStateTransitionNewPriorityRequest)
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
func (it *IZkSyncStateTransitionNewPriorityRequestIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IZkSyncStateTransitionNewPriorityRequestIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IZkSyncStateTransitionNewPriorityRequest represents a NewPriorityRequest event raised by the IZkSyncStateTransition contract.
type IZkSyncStateTransitionNewPriorityRequest struct {
	TxId                *big.Int
	TxHash              [32]byte
	ExpirationTimestamp uint64
	Transaction         L2CanonicalTransaction
	FactoryDeps         [][]byte
	Raw                 types.Log // Blockchain specific contextual infos
}

// FilterNewPriorityRequest is a free log retrieval operation binding the contract event 0x4531cd5795773d7101c17bdeb9f5ab7f47d7056017506f937083be5d6e77a382.
//
// Solidity: event NewPriorityRequest(uint256 txId, bytes32 txHash, uint64 expirationTimestamp, (uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256[4],bytes,bytes,uint256[],bytes,bytes) transaction, bytes[] factoryDeps)
func (_IZkSyncStateTransition *IZkSyncStateTransitionFilterer) FilterNewPriorityRequest(opts *bind.FilterOpts) (*IZkSyncStateTransitionNewPriorityRequestIterator, error) {

	logs, sub, err := _IZkSyncStateTransition.contract.FilterLogs(opts, "NewPriorityRequest")
	if err != nil {
		return nil, err
	}
	return &IZkSyncStateTransitionNewPriorityRequestIterator{contract: _IZkSyncStateTransition.contract, event: "NewPriorityRequest", logs: logs, sub: sub}, nil
}

// WatchNewPriorityRequest is a free log subscription operation binding the contract event 0x4531cd5795773d7101c17bdeb9f5ab7f47d7056017506f937083be5d6e77a382.
//
// Solidity: event NewPriorityRequest(uint256 txId, bytes32 txHash, uint64 expirationTimestamp, (uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256[4],bytes,bytes,uint256[],bytes,bytes) transaction, bytes[] factoryDeps)
func (_IZkSyncStateTransition *IZkSyncStateTransitionFilterer) WatchNewPriorityRequest(opts *bind.WatchOpts, sink chan<- *IZkSyncStateTransitionNewPriorityRequest) (event.Subscription, error) {

	logs, sub, err := _IZkSyncStateTransition.contract.WatchLogs(opts, "NewPriorityRequest")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IZkSyncStateTransitionNewPriorityRequest)
				if err := _IZkSyncStateTransition.contract.UnpackLog(event, "NewPriorityRequest", log); err != nil {
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

// ParseNewPriorityRequest is a log parse operation binding the contract event 0x4531cd5795773d7101c17bdeb9f5ab7f47d7056017506f937083be5d6e77a382.
//
// Solidity: event NewPriorityRequest(uint256 txId, bytes32 txHash, uint64 expirationTimestamp, (uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256[4],bytes,bytes,uint256[],bytes,bytes) transaction, bytes[] factoryDeps)
func (_IZkSyncStateTransition *IZkSyncStateTransitionFilterer) ParseNewPriorityRequest(log types.Log) (*IZkSyncStateTransitionNewPriorityRequest, error) {
	event := new(IZkSyncStateTransitionNewPriorityRequest)
	if err := _IZkSyncStateTransition.contract.UnpackLog(event, "NewPriorityRequest", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// IZkSyncStateTransitionNewPriorityTxMaxGasLimitIterator is returned from FilterNewPriorityTxMaxGasLimit and is used to iterate over the raw logs and unpacked data for NewPriorityTxMaxGasLimit events raised by the IZkSyncStateTransition contract.
type IZkSyncStateTransitionNewPriorityTxMaxGasLimitIterator struct {
	Event *IZkSyncStateTransitionNewPriorityTxMaxGasLimit // Event containing the contract specifics and raw log

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
func (it *IZkSyncStateTransitionNewPriorityTxMaxGasLimitIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IZkSyncStateTransitionNewPriorityTxMaxGasLimit)
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
		it.Event = new(IZkSyncStateTransitionNewPriorityTxMaxGasLimit)
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
func (it *IZkSyncStateTransitionNewPriorityTxMaxGasLimitIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IZkSyncStateTransitionNewPriorityTxMaxGasLimitIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IZkSyncStateTransitionNewPriorityTxMaxGasLimit represents a NewPriorityTxMaxGasLimit event raised by the IZkSyncStateTransition contract.
type IZkSyncStateTransitionNewPriorityTxMaxGasLimit struct {
	OldPriorityTxMaxGasLimit *big.Int
	NewPriorityTxMaxGasLimit *big.Int
	Raw                      types.Log // Blockchain specific contextual infos
}

// FilterNewPriorityTxMaxGasLimit is a free log retrieval operation binding the contract event 0x83dd728f7e76a849126c55ffabdc6e299eb8c85dccf12498701376d9f5c954d1.
//
// Solidity: event NewPriorityTxMaxGasLimit(uint256 oldPriorityTxMaxGasLimit, uint256 newPriorityTxMaxGasLimit)
func (_IZkSyncStateTransition *IZkSyncStateTransitionFilterer) FilterNewPriorityTxMaxGasLimit(opts *bind.FilterOpts) (*IZkSyncStateTransitionNewPriorityTxMaxGasLimitIterator, error) {

	logs, sub, err := _IZkSyncStateTransition.contract.FilterLogs(opts, "NewPriorityTxMaxGasLimit")
	if err != nil {
		return nil, err
	}
	return &IZkSyncStateTransitionNewPriorityTxMaxGasLimitIterator{contract: _IZkSyncStateTransition.contract, event: "NewPriorityTxMaxGasLimit", logs: logs, sub: sub}, nil
}

// WatchNewPriorityTxMaxGasLimit is a free log subscription operation binding the contract event 0x83dd728f7e76a849126c55ffabdc6e299eb8c85dccf12498701376d9f5c954d1.
//
// Solidity: event NewPriorityTxMaxGasLimit(uint256 oldPriorityTxMaxGasLimit, uint256 newPriorityTxMaxGasLimit)
func (_IZkSyncStateTransition *IZkSyncStateTransitionFilterer) WatchNewPriorityTxMaxGasLimit(opts *bind.WatchOpts, sink chan<- *IZkSyncStateTransitionNewPriorityTxMaxGasLimit) (event.Subscription, error) {

	logs, sub, err := _IZkSyncStateTransition.contract.WatchLogs(opts, "NewPriorityTxMaxGasLimit")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IZkSyncStateTransitionNewPriorityTxMaxGasLimit)
				if err := _IZkSyncStateTransition.contract.UnpackLog(event, "NewPriorityTxMaxGasLimit", log); err != nil {
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

// ParseNewPriorityTxMaxGasLimit is a log parse operation binding the contract event 0x83dd728f7e76a849126c55ffabdc6e299eb8c85dccf12498701376d9f5c954d1.
//
// Solidity: event NewPriorityTxMaxGasLimit(uint256 oldPriorityTxMaxGasLimit, uint256 newPriorityTxMaxGasLimit)
func (_IZkSyncStateTransition *IZkSyncStateTransitionFilterer) ParseNewPriorityTxMaxGasLimit(log types.Log) (*IZkSyncStateTransitionNewPriorityTxMaxGasLimit, error) {
	event := new(IZkSyncStateTransitionNewPriorityTxMaxGasLimit)
	if err := _IZkSyncStateTransition.contract.UnpackLog(event, "NewPriorityTxMaxGasLimit", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// IZkSyncStateTransitionNewTransactionFiltererIterator is returned from FilterNewTransactionFilterer and is used to iterate over the raw logs and unpacked data for NewTransactionFilterer events raised by the IZkSyncStateTransition contract.
type IZkSyncStateTransitionNewTransactionFiltererIterator struct {
	Event *IZkSyncStateTransitionNewTransactionFilterer // Event containing the contract specifics and raw log

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
func (it *IZkSyncStateTransitionNewTransactionFiltererIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IZkSyncStateTransitionNewTransactionFilterer)
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
		it.Event = new(IZkSyncStateTransitionNewTransactionFilterer)
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
func (it *IZkSyncStateTransitionNewTransactionFiltererIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IZkSyncStateTransitionNewTransactionFiltererIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IZkSyncStateTransitionNewTransactionFilterer represents a NewTransactionFilterer event raised by the IZkSyncStateTransition contract.
type IZkSyncStateTransitionNewTransactionFilterer struct {
	OldTransactionFilterer common.Address
	NewTransactionFilterer common.Address
	Raw                    types.Log // Blockchain specific contextual infos
}

// FilterNewTransactionFilterer is a free log retrieval operation binding the contract event 0xa9b43a66c5d1c607986f49e932a0c08058d843210db03dd6e8a621b96b4770f4.
//
// Solidity: event NewTransactionFilterer(address oldTransactionFilterer, address newTransactionFilterer)
func (_IZkSyncStateTransition *IZkSyncStateTransitionFilterer) FilterNewTransactionFilterer(opts *bind.FilterOpts) (*IZkSyncStateTransitionNewTransactionFiltererIterator, error) {

	logs, sub, err := _IZkSyncStateTransition.contract.FilterLogs(opts, "NewTransactionFilterer")
	if err != nil {
		return nil, err
	}
	return &IZkSyncStateTransitionNewTransactionFiltererIterator{contract: _IZkSyncStateTransition.contract, event: "NewTransactionFilterer", logs: logs, sub: sub}, nil
}

// WatchNewTransactionFilterer is a free log subscription operation binding the contract event 0xa9b43a66c5d1c607986f49e932a0c08058d843210db03dd6e8a621b96b4770f4.
//
// Solidity: event NewTransactionFilterer(address oldTransactionFilterer, address newTransactionFilterer)
func (_IZkSyncStateTransition *IZkSyncStateTransitionFilterer) WatchNewTransactionFilterer(opts *bind.WatchOpts, sink chan<- *IZkSyncStateTransitionNewTransactionFilterer) (event.Subscription, error) {

	logs, sub, err := _IZkSyncStateTransition.contract.WatchLogs(opts, "NewTransactionFilterer")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IZkSyncStateTransitionNewTransactionFilterer)
				if err := _IZkSyncStateTransition.contract.UnpackLog(event, "NewTransactionFilterer", log); err != nil {
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

// ParseNewTransactionFilterer is a log parse operation binding the contract event 0xa9b43a66c5d1c607986f49e932a0c08058d843210db03dd6e8a621b96b4770f4.
//
// Solidity: event NewTransactionFilterer(address oldTransactionFilterer, address newTransactionFilterer)
func (_IZkSyncStateTransition *IZkSyncStateTransitionFilterer) ParseNewTransactionFilterer(log types.Log) (*IZkSyncStateTransitionNewTransactionFilterer, error) {
	event := new(IZkSyncStateTransitionNewTransactionFilterer)
	if err := _IZkSyncStateTransition.contract.UnpackLog(event, "NewTransactionFilterer", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// IZkSyncStateTransitionProposeTransparentUpgradeIterator is returned from FilterProposeTransparentUpgrade and is used to iterate over the raw logs and unpacked data for ProposeTransparentUpgrade events raised by the IZkSyncStateTransition contract.
type IZkSyncStateTransitionProposeTransparentUpgradeIterator struct {
	Event *IZkSyncStateTransitionProposeTransparentUpgrade // Event containing the contract specifics and raw log

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
func (it *IZkSyncStateTransitionProposeTransparentUpgradeIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IZkSyncStateTransitionProposeTransparentUpgrade)
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
		it.Event = new(IZkSyncStateTransitionProposeTransparentUpgrade)
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
func (it *IZkSyncStateTransitionProposeTransparentUpgradeIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IZkSyncStateTransitionProposeTransparentUpgradeIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IZkSyncStateTransitionProposeTransparentUpgrade represents a ProposeTransparentUpgrade event raised by the IZkSyncStateTransition contract.
type IZkSyncStateTransitionProposeTransparentUpgrade struct {
	DiamondCut   DiamondDiamondCutData
	ProposalId   *big.Int
	ProposalSalt [32]byte
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterProposeTransparentUpgrade is a free log retrieval operation binding the contract event 0x69115b49afe7a6101a2e7af17d421eda1dc153bd26d699f013c4fff0404646a6.
//
// Solidity: event ProposeTransparentUpgrade(((address,uint8,bool,bytes4[])[],address,bytes) diamondCut, uint256 indexed proposalId, bytes32 proposalSalt)
func (_IZkSyncStateTransition *IZkSyncStateTransitionFilterer) FilterProposeTransparentUpgrade(opts *bind.FilterOpts, proposalId []*big.Int) (*IZkSyncStateTransitionProposeTransparentUpgradeIterator, error) {

	var proposalIdRule []interface{}
	for _, proposalIdItem := range proposalId {
		proposalIdRule = append(proposalIdRule, proposalIdItem)
	}

	logs, sub, err := _IZkSyncStateTransition.contract.FilterLogs(opts, "ProposeTransparentUpgrade", proposalIdRule)
	if err != nil {
		return nil, err
	}
	return &IZkSyncStateTransitionProposeTransparentUpgradeIterator{contract: _IZkSyncStateTransition.contract, event: "ProposeTransparentUpgrade", logs: logs, sub: sub}, nil
}

// WatchProposeTransparentUpgrade is a free log subscription operation binding the contract event 0x69115b49afe7a6101a2e7af17d421eda1dc153bd26d699f013c4fff0404646a6.
//
// Solidity: event ProposeTransparentUpgrade(((address,uint8,bool,bytes4[])[],address,bytes) diamondCut, uint256 indexed proposalId, bytes32 proposalSalt)
func (_IZkSyncStateTransition *IZkSyncStateTransitionFilterer) WatchProposeTransparentUpgrade(opts *bind.WatchOpts, sink chan<- *IZkSyncStateTransitionProposeTransparentUpgrade, proposalId []*big.Int) (event.Subscription, error) {

	var proposalIdRule []interface{}
	for _, proposalIdItem := range proposalId {
		proposalIdRule = append(proposalIdRule, proposalIdItem)
	}

	logs, sub, err := _IZkSyncStateTransition.contract.WatchLogs(opts, "ProposeTransparentUpgrade", proposalIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IZkSyncStateTransitionProposeTransparentUpgrade)
				if err := _IZkSyncStateTransition.contract.UnpackLog(event, "ProposeTransparentUpgrade", log); err != nil {
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

// ParseProposeTransparentUpgrade is a log parse operation binding the contract event 0x69115b49afe7a6101a2e7af17d421eda1dc153bd26d699f013c4fff0404646a6.
//
// Solidity: event ProposeTransparentUpgrade(((address,uint8,bool,bytes4[])[],address,bytes) diamondCut, uint256 indexed proposalId, bytes32 proposalSalt)
func (_IZkSyncStateTransition *IZkSyncStateTransitionFilterer) ParseProposeTransparentUpgrade(log types.Log) (*IZkSyncStateTransitionProposeTransparentUpgrade, error) {
	event := new(IZkSyncStateTransitionProposeTransparentUpgrade)
	if err := _IZkSyncStateTransition.contract.UnpackLog(event, "ProposeTransparentUpgrade", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// IZkSyncStateTransitionUnfreezeIterator is returned from FilterUnfreeze and is used to iterate over the raw logs and unpacked data for Unfreeze events raised by the IZkSyncStateTransition contract.
type IZkSyncStateTransitionUnfreezeIterator struct {
	Event *IZkSyncStateTransitionUnfreeze // Event containing the contract specifics and raw log

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
func (it *IZkSyncStateTransitionUnfreezeIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IZkSyncStateTransitionUnfreeze)
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
		it.Event = new(IZkSyncStateTransitionUnfreeze)
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
func (it *IZkSyncStateTransitionUnfreezeIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IZkSyncStateTransitionUnfreezeIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IZkSyncStateTransitionUnfreeze represents a Unfreeze event raised by the IZkSyncStateTransition contract.
type IZkSyncStateTransitionUnfreeze struct {
	Raw types.Log // Blockchain specific contextual infos
}

// FilterUnfreeze is a free log retrieval operation binding the contract event 0x2f05ba71d0df11bf5fa562a6569d70c4f80da84284badbe015ce1456063d0ded.
//
// Solidity: event Unfreeze()
func (_IZkSyncStateTransition *IZkSyncStateTransitionFilterer) FilterUnfreeze(opts *bind.FilterOpts) (*IZkSyncStateTransitionUnfreezeIterator, error) {

	logs, sub, err := _IZkSyncStateTransition.contract.FilterLogs(opts, "Unfreeze")
	if err != nil {
		return nil, err
	}
	return &IZkSyncStateTransitionUnfreezeIterator{contract: _IZkSyncStateTransition.contract, event: "Unfreeze", logs: logs, sub: sub}, nil
}

// WatchUnfreeze is a free log subscription operation binding the contract event 0x2f05ba71d0df11bf5fa562a6569d70c4f80da84284badbe015ce1456063d0ded.
//
// Solidity: event Unfreeze()
func (_IZkSyncStateTransition *IZkSyncStateTransitionFilterer) WatchUnfreeze(opts *bind.WatchOpts, sink chan<- *IZkSyncStateTransitionUnfreeze) (event.Subscription, error) {

	logs, sub, err := _IZkSyncStateTransition.contract.WatchLogs(opts, "Unfreeze")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IZkSyncStateTransitionUnfreeze)
				if err := _IZkSyncStateTransition.contract.UnpackLog(event, "Unfreeze", log); err != nil {
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

// ParseUnfreeze is a log parse operation binding the contract event 0x2f05ba71d0df11bf5fa562a6569d70c4f80da84284badbe015ce1456063d0ded.
//
// Solidity: event Unfreeze()
func (_IZkSyncStateTransition *IZkSyncStateTransitionFilterer) ParseUnfreeze(log types.Log) (*IZkSyncStateTransitionUnfreeze, error) {
	event := new(IZkSyncStateTransitionUnfreeze)
	if err := _IZkSyncStateTransition.contract.UnpackLog(event, "Unfreeze", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// IZkSyncStateTransitionValidatorStatusUpdateIterator is returned from FilterValidatorStatusUpdate and is used to iterate over the raw logs and unpacked data for ValidatorStatusUpdate events raised by the IZkSyncStateTransition contract.
type IZkSyncStateTransitionValidatorStatusUpdateIterator struct {
	Event *IZkSyncStateTransitionValidatorStatusUpdate // Event containing the contract specifics and raw log

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
func (it *IZkSyncStateTransitionValidatorStatusUpdateIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IZkSyncStateTransitionValidatorStatusUpdate)
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
		it.Event = new(IZkSyncStateTransitionValidatorStatusUpdate)
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
func (it *IZkSyncStateTransitionValidatorStatusUpdateIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IZkSyncStateTransitionValidatorStatusUpdateIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IZkSyncStateTransitionValidatorStatusUpdate represents a ValidatorStatusUpdate event raised by the IZkSyncStateTransition contract.
type IZkSyncStateTransitionValidatorStatusUpdate struct {
	ValidatorAddress common.Address
	IsActive         bool
	Raw              types.Log // Blockchain specific contextual infos
}

// FilterValidatorStatusUpdate is a free log retrieval operation binding the contract event 0x065b77b53864e46fda3d8986acb51696223d6dde7ced42441eb150bae6d48136.
//
// Solidity: event ValidatorStatusUpdate(address indexed validatorAddress, bool isActive)
func (_IZkSyncStateTransition *IZkSyncStateTransitionFilterer) FilterValidatorStatusUpdate(opts *bind.FilterOpts, validatorAddress []common.Address) (*IZkSyncStateTransitionValidatorStatusUpdateIterator, error) {

	var validatorAddressRule []interface{}
	for _, validatorAddressItem := range validatorAddress {
		validatorAddressRule = append(validatorAddressRule, validatorAddressItem)
	}

	logs, sub, err := _IZkSyncStateTransition.contract.FilterLogs(opts, "ValidatorStatusUpdate", validatorAddressRule)
	if err != nil {
		return nil, err
	}
	return &IZkSyncStateTransitionValidatorStatusUpdateIterator{contract: _IZkSyncStateTransition.contract, event: "ValidatorStatusUpdate", logs: logs, sub: sub}, nil
}

// WatchValidatorStatusUpdate is a free log subscription operation binding the contract event 0x065b77b53864e46fda3d8986acb51696223d6dde7ced42441eb150bae6d48136.
//
// Solidity: event ValidatorStatusUpdate(address indexed validatorAddress, bool isActive)
func (_IZkSyncStateTransition *IZkSyncStateTransitionFilterer) WatchValidatorStatusUpdate(opts *bind.WatchOpts, sink chan<- *IZkSyncStateTransitionValidatorStatusUpdate, validatorAddress []common.Address) (event.Subscription, error) {

	var validatorAddressRule []interface{}
	for _, validatorAddressItem := range validatorAddress {
		validatorAddressRule = append(validatorAddressRule, validatorAddressItem)
	}

	logs, sub, err := _IZkSyncStateTransition.contract.WatchLogs(opts, "ValidatorStatusUpdate", validatorAddressRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IZkSyncStateTransitionValidatorStatusUpdate)
				if err := _IZkSyncStateTransition.contract.UnpackLog(event, "ValidatorStatusUpdate", log); err != nil {
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

// ParseValidatorStatusUpdate is a log parse operation binding the contract event 0x065b77b53864e46fda3d8986acb51696223d6dde7ced42441eb150bae6d48136.
//
// Solidity: event ValidatorStatusUpdate(address indexed validatorAddress, bool isActive)
func (_IZkSyncStateTransition *IZkSyncStateTransitionFilterer) ParseValidatorStatusUpdate(log types.Log) (*IZkSyncStateTransitionValidatorStatusUpdate, error) {
	event := new(IZkSyncStateTransitionValidatorStatusUpdate)
	if err := _IZkSyncStateTransition.contract.UnpackLog(event, "ValidatorStatusUpdate", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// IZkSyncStateTransitionValidiumModeStatusUpdateIterator is returned from FilterValidiumModeStatusUpdate and is used to iterate over the raw logs and unpacked data for ValidiumModeStatusUpdate events raised by the IZkSyncStateTransition contract.
type IZkSyncStateTransitionValidiumModeStatusUpdateIterator struct {
	Event *IZkSyncStateTransitionValidiumModeStatusUpdate // Event containing the contract specifics and raw log

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
func (it *IZkSyncStateTransitionValidiumModeStatusUpdateIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IZkSyncStateTransitionValidiumModeStatusUpdate)
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
		it.Event = new(IZkSyncStateTransitionValidiumModeStatusUpdate)
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
func (it *IZkSyncStateTransitionValidiumModeStatusUpdateIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IZkSyncStateTransitionValidiumModeStatusUpdateIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IZkSyncStateTransitionValidiumModeStatusUpdate represents a ValidiumModeStatusUpdate event raised by the IZkSyncStateTransition contract.
type IZkSyncStateTransitionValidiumModeStatusUpdate struct {
	ValidiumMode uint8
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterValidiumModeStatusUpdate is a free log retrieval operation binding the contract event 0xaa01146df0a628c6b214e79a281f7524b792de4a52ad6f07c78e6e035d58c896.
//
// Solidity: event ValidiumModeStatusUpdate(uint8 validiumMode)
func (_IZkSyncStateTransition *IZkSyncStateTransitionFilterer) FilterValidiumModeStatusUpdate(opts *bind.FilterOpts) (*IZkSyncStateTransitionValidiumModeStatusUpdateIterator, error) {

	logs, sub, err := _IZkSyncStateTransition.contract.FilterLogs(opts, "ValidiumModeStatusUpdate")
	if err != nil {
		return nil, err
	}
	return &IZkSyncStateTransitionValidiumModeStatusUpdateIterator{contract: _IZkSyncStateTransition.contract, event: "ValidiumModeStatusUpdate", logs: logs, sub: sub}, nil
}

// WatchValidiumModeStatusUpdate is a free log subscription operation binding the contract event 0xaa01146df0a628c6b214e79a281f7524b792de4a52ad6f07c78e6e035d58c896.
//
// Solidity: event ValidiumModeStatusUpdate(uint8 validiumMode)
func (_IZkSyncStateTransition *IZkSyncStateTransitionFilterer) WatchValidiumModeStatusUpdate(opts *bind.WatchOpts, sink chan<- *IZkSyncStateTransitionValidiumModeStatusUpdate) (event.Subscription, error) {

	logs, sub, err := _IZkSyncStateTransition.contract.WatchLogs(opts, "ValidiumModeStatusUpdate")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IZkSyncStateTransitionValidiumModeStatusUpdate)
				if err := _IZkSyncStateTransition.contract.UnpackLog(event, "ValidiumModeStatusUpdate", log); err != nil {
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

// ParseValidiumModeStatusUpdate is a log parse operation binding the contract event 0xaa01146df0a628c6b214e79a281f7524b792de4a52ad6f07c78e6e035d58c896.
//
// Solidity: event ValidiumModeStatusUpdate(uint8 validiumMode)
func (_IZkSyncStateTransition *IZkSyncStateTransitionFilterer) ParseValidiumModeStatusUpdate(log types.Log) (*IZkSyncStateTransitionValidiumModeStatusUpdate, error) {
	event := new(IZkSyncStateTransitionValidiumModeStatusUpdate)
	if err := _IZkSyncStateTransition.contract.UnpackLog(event, "ValidiumModeStatusUpdate", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
