// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package zksync

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
	TotalL2ToL1Pubdata                []byte
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

// IMailboxL2CanonicalTransaction is an auto generated low-level Go binding around an user-defined struct.
type IMailboxL2CanonicalTransaction struct {
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

// IZkSyncMetaData contains all meta data concerning the IZkSync contract.
var IZkSyncMetaData = &bind.MetaData{
	ABI: "[{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"batchNumber\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"batchHash\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"commitment\",\"type\":\"bytes32\"}],\"name\":\"BlockCommit\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"batchNumber\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"batchHash\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"commitment\",\"type\":\"bytes32\"}],\"name\":\"BlockExecution\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"totalBatchesCommitted\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"totalBatchesVerified\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"totalBatchesExecuted\",\"type\":\"uint256\"}],\"name\":\"BlocksRevert\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"previousLastVerifiedBatch\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"currentLastVerifiedBatch\",\"type\":\"uint256\"}],\"name\":\"BlocksVerification\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"EthWithdrawalFinalized\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"components\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"facet\",\"type\":\"address\"},{\"internalType\":\"enumDiamond.Action\",\"name\":\"action\",\"type\":\"uint8\"},{\"internalType\":\"bool\",\"name\":\"isFreezable\",\"type\":\"bool\"},{\"internalType\":\"bytes4[]\",\"name\":\"selectors\",\"type\":\"bytes4[]\"}],\"internalType\":\"structDiamond.FacetCut[]\",\"name\":\"facetCuts\",\"type\":\"tuple[]\"},{\"internalType\":\"address\",\"name\":\"initAddress\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"initCalldata\",\"type\":\"bytes\"}],\"indexed\":false,\"internalType\":\"structDiamond.DiamondCutData\",\"name\":\"diamondCut\",\"type\":\"tuple\"}],\"name\":\"ExecuteUpgrade\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[],\"name\":\"Freeze\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"isPorterAvailable\",\"type\":\"bool\"}],\"name\":\"IsPorterAvailableStatusUpdate\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"oldAdmin\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newAdmin\",\"type\":\"address\"}],\"name\":\"NewAdmin\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"oldGovernor\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newGovernor\",\"type\":\"address\"}],\"name\":\"NewGovernor\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"oldPendingAdmin\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newPendingAdmin\",\"type\":\"address\"}],\"name\":\"NewPendingAdmin\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"oldPendingGovernor\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newPendingGovernor\",\"type\":\"address\"}],\"name\":\"NewPendingGovernor\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"txId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"txHash\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"expirationTimestamp\",\"type\":\"uint64\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"txType\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"from\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"to\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"gasLimit\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"gasPerPubdataByteLimit\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"maxFeePerGas\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"maxPriorityFeePerGas\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"paymaster\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"nonce\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"},{\"internalType\":\"uint256[4]\",\"name\":\"reserved\",\"type\":\"uint256[4]\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"signature\",\"type\":\"bytes\"},{\"internalType\":\"uint256[]\",\"name\":\"factoryDeps\",\"type\":\"uint256[]\"},{\"internalType\":\"bytes\",\"name\":\"paymasterInput\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"reservedDynamic\",\"type\":\"bytes\"}],\"indexed\":false,\"internalType\":\"structIMailbox.L2CanonicalTransaction\",\"name\":\"transaction\",\"type\":\"tuple\"},{\"indexed\":false,\"internalType\":\"bytes[]\",\"name\":\"factoryDeps\",\"type\":\"bytes[]\"}],\"name\":\"NewPriorityRequest\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"oldPriorityTxMaxGasLimit\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"newPriorityTxMaxGasLimit\",\"type\":\"uint256\"}],\"name\":\"NewPriorityTxMaxGasLimit\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[],\"name\":\"Unfreeze\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"validatorAddress\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"isActive\",\"type\":\"bool\"}],\"name\":\"ValidatorStatusUpdate\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"acceptAdmin\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"acceptGovernor\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint64\",\"name\":\"batchNumber\",\"type\":\"uint64\"},{\"internalType\":\"bytes32\",\"name\":\"batchHash\",\"type\":\"bytes32\"},{\"internalType\":\"uint64\",\"name\":\"indexRepeatedStorageChanges\",\"type\":\"uint64\"},{\"internalType\":\"uint256\",\"name\":\"numberOfLayer1Txs\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"priorityOperationsHash\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"l2LogsTreeRoot\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"timestamp\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"commitment\",\"type\":\"bytes32\"}],\"internalType\":\"structIExecutor.StoredBatchInfo\",\"name\":\"_lastCommittedBatchData\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"uint64\",\"name\":\"batchNumber\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"timestamp\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"indexRepeatedStorageChanges\",\"type\":\"uint64\"},{\"internalType\":\"bytes32\",\"name\":\"newStateRoot\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"numberOfLayer1Txs\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"priorityOperationsHash\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"bootloaderHeapInitialContentsHash\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"eventsQueueStateHash\",\"type\":\"bytes32\"},{\"internalType\":\"bytes\",\"name\":\"systemLogs\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"totalL2ToL1Pubdata\",\"type\":\"bytes\"}],\"internalType\":\"structIExecutor.CommitBatchInfo[]\",\"name\":\"_newBatchesData\",\"type\":\"tuple[]\"}],\"name\":\"commitBatches\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint64\",\"name\":\"batchNumber\",\"type\":\"uint64\"},{\"internalType\":\"bytes32\",\"name\":\"batchHash\",\"type\":\"bytes32\"},{\"internalType\":\"uint64\",\"name\":\"indexRepeatedStorageChanges\",\"type\":\"uint64\"},{\"internalType\":\"uint256\",\"name\":\"numberOfLayer1Txs\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"priorityOperationsHash\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"l2LogsTreeRoot\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"timestamp\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"commitment\",\"type\":\"bytes32\"}],\"internalType\":\"structIExecutor.StoredBatchInfo[]\",\"name\":\"_batchesData\",\"type\":\"tuple[]\"}],\"name\":\"executeBatches\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"facet\",\"type\":\"address\"},{\"internalType\":\"enumDiamond.Action\",\"name\":\"action\",\"type\":\"uint8\"},{\"internalType\":\"bool\",\"name\":\"isFreezable\",\"type\":\"bool\"},{\"internalType\":\"bytes4[]\",\"name\":\"selectors\",\"type\":\"bytes4[]\"}],\"internalType\":\"structDiamond.FacetCut[]\",\"name\":\"facetCuts\",\"type\":\"tuple[]\"},{\"internalType\":\"address\",\"name\":\"initAddress\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"initCalldata\",\"type\":\"bytes\"}],\"internalType\":\"structDiamond.DiamondCutData\",\"name\":\"_diamondCut\",\"type\":\"tuple\"}],\"name\":\"executeUpgrade\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes4\",\"name\":\"_selector\",\"type\":\"bytes4\"}],\"name\":\"facetAddress\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"facet\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"facetAddresses\",\"outputs\":[{\"internalType\":\"address[]\",\"name\":\"facets\",\"type\":\"address[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_facet\",\"type\":\"address\"}],\"name\":\"facetFunctionSelectors\",\"outputs\":[{\"internalType\":\"bytes4[]\",\"name\":\"\",\"type\":\"bytes4[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"facets\",\"outputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"addr\",\"type\":\"address\"},{\"internalType\":\"bytes4[]\",\"name\":\"selectors\",\"type\":\"bytes4[]\"}],\"internalType\":\"structIGetters.Facet[]\",\"name\":\"\",\"type\":\"tuple[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_l2BatchNumber\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_l2MessageIndex\",\"type\":\"uint256\"},{\"internalType\":\"uint16\",\"name\":\"_l2TxNumberInBatch\",\"type\":\"uint16\"},{\"internalType\":\"bytes\",\"name\":\"_message\",\"type\":\"bytes\"},{\"internalType\":\"bytes32[]\",\"name\":\"_merkleProof\",\"type\":\"bytes32[]\"}],\"name\":\"finalizeEthWithdrawal\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"freezeDiamond\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getFirstUnprocessedPriorityTx\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getGovernor\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getL2BootloaderBytecodeHash\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getL2DefaultAccountBytecodeHash\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getL2SystemContractsUpgradeBatchNumber\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getL2SystemContractsUpgradeTxHash\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getName\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getPendingGovernor\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getPriorityQueueSize\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getPriorityTxMaxGasLimit\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getProtocolVersion\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getTotalBatchesCommitted\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getTotalBatchesExecuted\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getTotalBatchesVerified\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getTotalPriorityTxs\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getVerifier\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getVerifierParams\",\"outputs\":[{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"recursionNodeLevelVkHash\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"recursionLeafLevelVkHash\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"recursionCircuitsSetVksHash\",\"type\":\"bytes32\"}],\"internalType\":\"structVerifierParams\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"isDiamondStorageFrozen\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_l2BatchNumber\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_l2MessageIndex\",\"type\":\"uint256\"}],\"name\":\"isEthWithdrawalFinalized\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_facet\",\"type\":\"address\"}],\"name\":\"isFacetFreezable\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"isFreezable\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes4\",\"name\":\"_selector\",\"type\":\"bytes4\"}],\"name\":\"isFunctionFreezable\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_address\",\"type\":\"address\"}],\"name\":\"isValidator\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_batchNumber\",\"type\":\"uint256\"}],\"name\":\"l2LogsRootHash\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"hash\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_gasPrice\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_l2GasLimit\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_l2GasPerPubdataByteLimit\",\"type\":\"uint256\"}],\"name\":\"l2TransactionBaseCost\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"priorityQueueFrontOperation\",\"outputs\":[{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"canonicalTxHash\",\"type\":\"bytes32\"},{\"internalType\":\"uint64\",\"name\":\"expirationTimestamp\",\"type\":\"uint64\"},{\"internalType\":\"uint192\",\"name\":\"layer2Tip\",\"type\":\"uint192\"}],\"internalType\":\"structPriorityOperation\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint64\",\"name\":\"batchNumber\",\"type\":\"uint64\"},{\"internalType\":\"bytes32\",\"name\":\"batchHash\",\"type\":\"bytes32\"},{\"internalType\":\"uint64\",\"name\":\"indexRepeatedStorageChanges\",\"type\":\"uint64\"},{\"internalType\":\"uint256\",\"name\":\"numberOfLayer1Txs\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"priorityOperationsHash\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"l2LogsTreeRoot\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"timestamp\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"commitment\",\"type\":\"bytes32\"}],\"internalType\":\"structIExecutor.StoredBatchInfo\",\"name\":\"_prevBatch\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"uint64\",\"name\":\"batchNumber\",\"type\":\"uint64\"},{\"internalType\":\"bytes32\",\"name\":\"batchHash\",\"type\":\"bytes32\"},{\"internalType\":\"uint64\",\"name\":\"indexRepeatedStorageChanges\",\"type\":\"uint64\"},{\"internalType\":\"uint256\",\"name\":\"numberOfLayer1Txs\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"priorityOperationsHash\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"l2LogsTreeRoot\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"timestamp\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"commitment\",\"type\":\"bytes32\"}],\"internalType\":\"structIExecutor.StoredBatchInfo[]\",\"name\":\"_committedBatches\",\"type\":\"tuple[]\"},{\"components\":[{\"internalType\":\"uint256[]\",\"name\":\"recursiveAggregationInput\",\"type\":\"uint256[]\"},{\"internalType\":\"uint256[]\",\"name\":\"serializedProof\",\"type\":\"uint256[]\"}],\"internalType\":\"structIExecutor.ProofInput\",\"name\":\"_proof\",\"type\":\"tuple\"}],\"name\":\"proveBatches\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_l2TxHash\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"_l2BatchNumber\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_l2MessageIndex\",\"type\":\"uint256\"},{\"internalType\":\"uint16\",\"name\":\"_l2TxNumberInBatch\",\"type\":\"uint16\"},{\"internalType\":\"bytes32[]\",\"name\":\"_merkleProof\",\"type\":\"bytes32[]\"},{\"internalType\":\"enumTxStatus\",\"name\":\"_status\",\"type\":\"uint8\"}],\"name\":\"proveL1ToL2TransactionStatus\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_l2BatchNumber\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_index\",\"type\":\"uint256\"},{\"components\":[{\"internalType\":\"uint8\",\"name\":\"l2ShardId\",\"type\":\"uint8\"},{\"internalType\":\"bool\",\"name\":\"isService\",\"type\":\"bool\"},{\"internalType\":\"uint16\",\"name\":\"txNumberInBatch\",\"type\":\"uint16\"},{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"key\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"value\",\"type\":\"bytes32\"}],\"internalType\":\"structL2Log\",\"name\":\"_log\",\"type\":\"tuple\"},{\"internalType\":\"bytes32[]\",\"name\":\"_proof\",\"type\":\"bytes32[]\"}],\"name\":\"proveL2LogInclusion\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_l2BatchNumber\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_index\",\"type\":\"uint256\"},{\"components\":[{\"internalType\":\"uint16\",\"name\":\"txNumberInBatch\",\"type\":\"uint16\"},{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"internalType\":\"structL2Message\",\"name\":\"_message\",\"type\":\"tuple\"},{\"internalType\":\"bytes32[]\",\"name\":\"_proof\",\"type\":\"bytes32[]\"}],\"name\":\"proveL2MessageInclusion\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_contractL2\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_l2Value\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"_calldata\",\"type\":\"bytes\"},{\"internalType\":\"uint256\",\"name\":\"_l2GasLimit\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_l2GasPerPubdataByteLimit\",\"type\":\"uint256\"},{\"internalType\":\"bytes[]\",\"name\":\"_factoryDeps\",\"type\":\"bytes[]\"},{\"internalType\":\"address\",\"name\":\"_refundRecipient\",\"type\":\"address\"}],\"name\":\"requestL2Transaction\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"canonicalTxHash\",\"type\":\"bytes32\"}],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_newLastBatch\",\"type\":\"uint256\"}],\"name\":\"revertBatches\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_newPendingAdmin\",\"type\":\"address\"}],\"name\":\"setPendingAdmin\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_newPendingGovernor\",\"type\":\"address\"}],\"name\":\"setPendingGovernor\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bool\",\"name\":\"_zkPorterIsAvailable\",\"type\":\"bool\"}],\"name\":\"setPorterAvailability\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_newPriorityTxMaxGasLimit\",\"type\":\"uint256\"}],\"name\":\"setPriorityTxMaxGasLimit\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_validator\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"_active\",\"type\":\"bool\"}],\"name\":\"setValidator\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_batchNumber\",\"type\":\"uint256\"}],\"name\":\"storedBatchHash\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"unfreezeDiamond\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
}

// IZkSyncABI is the input ABI used to generate the binding from.
// Deprecated: Use IZkSyncMetaData.ABI instead.
var IZkSyncABI = IZkSyncMetaData.ABI

// IZkSync is an auto generated Go binding around an Ethereum contract.
type IZkSync struct {
	IZkSyncCaller     // Read-only binding to the contract
	IZkSyncTransactor // Write-only binding to the contract
	IZkSyncFilterer   // Log filterer for contract events
}

// IZkSyncCaller is an auto generated read-only Go binding around an Ethereum contract.
type IZkSyncCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IZkSyncTransactor is an auto generated write-only Go binding around an Ethereum contract.
type IZkSyncTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IZkSyncFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type IZkSyncFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IZkSyncSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type IZkSyncSession struct {
	Contract     *IZkSync          // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// IZkSyncCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type IZkSyncCallerSession struct {
	Contract *IZkSyncCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts  // Call options to use throughout this session
}

// IZkSyncTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type IZkSyncTransactorSession struct {
	Contract     *IZkSyncTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts  // Transaction auth options to use throughout this session
}

// IZkSyncRaw is an auto generated low-level Go binding around an Ethereum contract.
type IZkSyncRaw struct {
	Contract *IZkSync // Generic contract binding to access the raw methods on
}

// IZkSyncCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type IZkSyncCallerRaw struct {
	Contract *IZkSyncCaller // Generic read-only contract binding to access the raw methods on
}

// IZkSyncTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type IZkSyncTransactorRaw struct {
	Contract *IZkSyncTransactor // Generic write-only contract binding to access the raw methods on
}

// NewIZkSync creates a new instance of IZkSync, bound to a specific deployed contract.
func NewIZkSync(address common.Address, backend bind.ContractBackend) (*IZkSync, error) {
	contract, err := bindIZkSync(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &IZkSync{IZkSyncCaller: IZkSyncCaller{contract: contract}, IZkSyncTransactor: IZkSyncTransactor{contract: contract}, IZkSyncFilterer: IZkSyncFilterer{contract: contract}}, nil
}

// NewIZkSyncCaller creates a new read-only instance of IZkSync, bound to a specific deployed contract.
func NewIZkSyncCaller(address common.Address, caller bind.ContractCaller) (*IZkSyncCaller, error) {
	contract, err := bindIZkSync(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &IZkSyncCaller{contract: contract}, nil
}

// NewIZkSyncTransactor creates a new write-only instance of IZkSync, bound to a specific deployed contract.
func NewIZkSyncTransactor(address common.Address, transactor bind.ContractTransactor) (*IZkSyncTransactor, error) {
	contract, err := bindIZkSync(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &IZkSyncTransactor{contract: contract}, nil
}

// NewIZkSyncFilterer creates a new log filterer instance of IZkSync, bound to a specific deployed contract.
func NewIZkSyncFilterer(address common.Address, filterer bind.ContractFilterer) (*IZkSyncFilterer, error) {
	contract, err := bindIZkSync(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &IZkSyncFilterer{contract: contract}, nil
}

// bindIZkSync binds a generic wrapper to an already deployed contract.
func bindIZkSync(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := IZkSyncMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IZkSync *IZkSyncRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IZkSync.Contract.IZkSyncCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IZkSync *IZkSyncRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IZkSync.Contract.IZkSyncTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IZkSync *IZkSyncRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IZkSync.Contract.IZkSyncTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IZkSync *IZkSyncCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IZkSync.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IZkSync *IZkSyncTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IZkSync.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IZkSync *IZkSyncTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IZkSync.Contract.contract.Transact(opts, method, params...)
}

// FacetAddress is a free data retrieval call binding the contract method 0xcdffacc6.
//
// Solidity: function facetAddress(bytes4 _selector) view returns(address facet)
func (_IZkSync *IZkSyncCaller) FacetAddress(opts *bind.CallOpts, _selector [4]byte) (common.Address, error) {
	var out []interface{}
	err := _IZkSync.contract.Call(opts, &out, "facetAddress", _selector)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// FacetAddress is a free data retrieval call binding the contract method 0xcdffacc6.
//
// Solidity: function facetAddress(bytes4 _selector) view returns(address facet)
func (_IZkSync *IZkSyncSession) FacetAddress(_selector [4]byte) (common.Address, error) {
	return _IZkSync.Contract.FacetAddress(&_IZkSync.CallOpts, _selector)
}

// FacetAddress is a free data retrieval call binding the contract method 0xcdffacc6.
//
// Solidity: function facetAddress(bytes4 _selector) view returns(address facet)
func (_IZkSync *IZkSyncCallerSession) FacetAddress(_selector [4]byte) (common.Address, error) {
	return _IZkSync.Contract.FacetAddress(&_IZkSync.CallOpts, _selector)
}

// FacetAddresses is a free data retrieval call binding the contract method 0x52ef6b2c.
//
// Solidity: function facetAddresses() view returns(address[] facets)
func (_IZkSync *IZkSyncCaller) FacetAddresses(opts *bind.CallOpts) ([]common.Address, error) {
	var out []interface{}
	err := _IZkSync.contract.Call(opts, &out, "facetAddresses")

	if err != nil {
		return *new([]common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)

	return out0, err

}

// FacetAddresses is a free data retrieval call binding the contract method 0x52ef6b2c.
//
// Solidity: function facetAddresses() view returns(address[] facets)
func (_IZkSync *IZkSyncSession) FacetAddresses() ([]common.Address, error) {
	return _IZkSync.Contract.FacetAddresses(&_IZkSync.CallOpts)
}

// FacetAddresses is a free data retrieval call binding the contract method 0x52ef6b2c.
//
// Solidity: function facetAddresses() view returns(address[] facets)
func (_IZkSync *IZkSyncCallerSession) FacetAddresses() ([]common.Address, error) {
	return _IZkSync.Contract.FacetAddresses(&_IZkSync.CallOpts)
}

// FacetFunctionSelectors is a free data retrieval call binding the contract method 0xadfca15e.
//
// Solidity: function facetFunctionSelectors(address _facet) view returns(bytes4[])
func (_IZkSync *IZkSyncCaller) FacetFunctionSelectors(opts *bind.CallOpts, _facet common.Address) ([][4]byte, error) {
	var out []interface{}
	err := _IZkSync.contract.Call(opts, &out, "facetFunctionSelectors", _facet)

	if err != nil {
		return *new([][4]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([][4]byte)).(*[][4]byte)

	return out0, err

}

// FacetFunctionSelectors is a free data retrieval call binding the contract method 0xadfca15e.
//
// Solidity: function facetFunctionSelectors(address _facet) view returns(bytes4[])
func (_IZkSync *IZkSyncSession) FacetFunctionSelectors(_facet common.Address) ([][4]byte, error) {
	return _IZkSync.Contract.FacetFunctionSelectors(&_IZkSync.CallOpts, _facet)
}

// FacetFunctionSelectors is a free data retrieval call binding the contract method 0xadfca15e.
//
// Solidity: function facetFunctionSelectors(address _facet) view returns(bytes4[])
func (_IZkSync *IZkSyncCallerSession) FacetFunctionSelectors(_facet common.Address) ([][4]byte, error) {
	return _IZkSync.Contract.FacetFunctionSelectors(&_IZkSync.CallOpts, _facet)
}

// Facets is a free data retrieval call binding the contract method 0x7a0ed627.
//
// Solidity: function facets() view returns((address,bytes4[])[])
func (_IZkSync *IZkSyncCaller) Facets(opts *bind.CallOpts) ([]IGettersFacet, error) {
	var out []interface{}
	err := _IZkSync.contract.Call(opts, &out, "facets")

	if err != nil {
		return *new([]IGettersFacet), err
	}

	out0 := *abi.ConvertType(out[0], new([]IGettersFacet)).(*[]IGettersFacet)

	return out0, err

}

// Facets is a free data retrieval call binding the contract method 0x7a0ed627.
//
// Solidity: function facets() view returns((address,bytes4[])[])
func (_IZkSync *IZkSyncSession) Facets() ([]IGettersFacet, error) {
	return _IZkSync.Contract.Facets(&_IZkSync.CallOpts)
}

// Facets is a free data retrieval call binding the contract method 0x7a0ed627.
//
// Solidity: function facets() view returns((address,bytes4[])[])
func (_IZkSync *IZkSyncCallerSession) Facets() ([]IGettersFacet, error) {
	return _IZkSync.Contract.Facets(&_IZkSync.CallOpts)
}

// GetFirstUnprocessedPriorityTx is a free data retrieval call binding the contract method 0x79823c9a.
//
// Solidity: function getFirstUnprocessedPriorityTx() view returns(uint256)
func (_IZkSync *IZkSyncCaller) GetFirstUnprocessedPriorityTx(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _IZkSync.contract.Call(opts, &out, "getFirstUnprocessedPriorityTx")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetFirstUnprocessedPriorityTx is a free data retrieval call binding the contract method 0x79823c9a.
//
// Solidity: function getFirstUnprocessedPriorityTx() view returns(uint256)
func (_IZkSync *IZkSyncSession) GetFirstUnprocessedPriorityTx() (*big.Int, error) {
	return _IZkSync.Contract.GetFirstUnprocessedPriorityTx(&_IZkSync.CallOpts)
}

// GetFirstUnprocessedPriorityTx is a free data retrieval call binding the contract method 0x79823c9a.
//
// Solidity: function getFirstUnprocessedPriorityTx() view returns(uint256)
func (_IZkSync *IZkSyncCallerSession) GetFirstUnprocessedPriorityTx() (*big.Int, error) {
	return _IZkSync.Contract.GetFirstUnprocessedPriorityTx(&_IZkSync.CallOpts)
}

// GetGovernor is a free data retrieval call binding the contract method 0x4fc07d75.
//
// Solidity: function getGovernor() view returns(address)
func (_IZkSync *IZkSyncCaller) GetGovernor(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _IZkSync.contract.Call(opts, &out, "getGovernor")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetGovernor is a free data retrieval call binding the contract method 0x4fc07d75.
//
// Solidity: function getGovernor() view returns(address)
func (_IZkSync *IZkSyncSession) GetGovernor() (common.Address, error) {
	return _IZkSync.Contract.GetGovernor(&_IZkSync.CallOpts)
}

// GetGovernor is a free data retrieval call binding the contract method 0x4fc07d75.
//
// Solidity: function getGovernor() view returns(address)
func (_IZkSync *IZkSyncCallerSession) GetGovernor() (common.Address, error) {
	return _IZkSync.Contract.GetGovernor(&_IZkSync.CallOpts)
}

// GetL2BootloaderBytecodeHash is a free data retrieval call binding the contract method 0xd86970d8.
//
// Solidity: function getL2BootloaderBytecodeHash() view returns(bytes32)
func (_IZkSync *IZkSyncCaller) GetL2BootloaderBytecodeHash(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _IZkSync.contract.Call(opts, &out, "getL2BootloaderBytecodeHash")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// GetL2BootloaderBytecodeHash is a free data retrieval call binding the contract method 0xd86970d8.
//
// Solidity: function getL2BootloaderBytecodeHash() view returns(bytes32)
func (_IZkSync *IZkSyncSession) GetL2BootloaderBytecodeHash() ([32]byte, error) {
	return _IZkSync.Contract.GetL2BootloaderBytecodeHash(&_IZkSync.CallOpts)
}

// GetL2BootloaderBytecodeHash is a free data retrieval call binding the contract method 0xd86970d8.
//
// Solidity: function getL2BootloaderBytecodeHash() view returns(bytes32)
func (_IZkSync *IZkSyncCallerSession) GetL2BootloaderBytecodeHash() ([32]byte, error) {
	return _IZkSync.Contract.GetL2BootloaderBytecodeHash(&_IZkSync.CallOpts)
}

// GetL2DefaultAccountBytecodeHash is a free data retrieval call binding the contract method 0xfd791f3c.
//
// Solidity: function getL2DefaultAccountBytecodeHash() view returns(bytes32)
func (_IZkSync *IZkSyncCaller) GetL2DefaultAccountBytecodeHash(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _IZkSync.contract.Call(opts, &out, "getL2DefaultAccountBytecodeHash")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// GetL2DefaultAccountBytecodeHash is a free data retrieval call binding the contract method 0xfd791f3c.
//
// Solidity: function getL2DefaultAccountBytecodeHash() view returns(bytes32)
func (_IZkSync *IZkSyncSession) GetL2DefaultAccountBytecodeHash() ([32]byte, error) {
	return _IZkSync.Contract.GetL2DefaultAccountBytecodeHash(&_IZkSync.CallOpts)
}

// GetL2DefaultAccountBytecodeHash is a free data retrieval call binding the contract method 0xfd791f3c.
//
// Solidity: function getL2DefaultAccountBytecodeHash() view returns(bytes32)
func (_IZkSync *IZkSyncCallerSession) GetL2DefaultAccountBytecodeHash() ([32]byte, error) {
	return _IZkSync.Contract.GetL2DefaultAccountBytecodeHash(&_IZkSync.CallOpts)
}

// GetL2SystemContractsUpgradeBatchNumber is a free data retrieval call binding the contract method 0xe5355c75.
//
// Solidity: function getL2SystemContractsUpgradeBatchNumber() view returns(uint256)
func (_IZkSync *IZkSyncCaller) GetL2SystemContractsUpgradeBatchNumber(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _IZkSync.contract.Call(opts, &out, "getL2SystemContractsUpgradeBatchNumber")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetL2SystemContractsUpgradeBatchNumber is a free data retrieval call binding the contract method 0xe5355c75.
//
// Solidity: function getL2SystemContractsUpgradeBatchNumber() view returns(uint256)
func (_IZkSync *IZkSyncSession) GetL2SystemContractsUpgradeBatchNumber() (*big.Int, error) {
	return _IZkSync.Contract.GetL2SystemContractsUpgradeBatchNumber(&_IZkSync.CallOpts)
}

// GetL2SystemContractsUpgradeBatchNumber is a free data retrieval call binding the contract method 0xe5355c75.
//
// Solidity: function getL2SystemContractsUpgradeBatchNumber() view returns(uint256)
func (_IZkSync *IZkSyncCallerSession) GetL2SystemContractsUpgradeBatchNumber() (*big.Int, error) {
	return _IZkSync.Contract.GetL2SystemContractsUpgradeBatchNumber(&_IZkSync.CallOpts)
}

// GetL2SystemContractsUpgradeTxHash is a free data retrieval call binding the contract method 0x7b30c8da.
//
// Solidity: function getL2SystemContractsUpgradeTxHash() view returns(bytes32)
func (_IZkSync *IZkSyncCaller) GetL2SystemContractsUpgradeTxHash(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _IZkSync.contract.Call(opts, &out, "getL2SystemContractsUpgradeTxHash")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// GetL2SystemContractsUpgradeTxHash is a free data retrieval call binding the contract method 0x7b30c8da.
//
// Solidity: function getL2SystemContractsUpgradeTxHash() view returns(bytes32)
func (_IZkSync *IZkSyncSession) GetL2SystemContractsUpgradeTxHash() ([32]byte, error) {
	return _IZkSync.Contract.GetL2SystemContractsUpgradeTxHash(&_IZkSync.CallOpts)
}

// GetL2SystemContractsUpgradeTxHash is a free data retrieval call binding the contract method 0x7b30c8da.
//
// Solidity: function getL2SystemContractsUpgradeTxHash() view returns(bytes32)
func (_IZkSync *IZkSyncCallerSession) GetL2SystemContractsUpgradeTxHash() ([32]byte, error) {
	return _IZkSync.Contract.GetL2SystemContractsUpgradeTxHash(&_IZkSync.CallOpts)
}

// GetName is a free data retrieval call binding the contract method 0x17d7de7c.
//
// Solidity: function getName() view returns(string)
func (_IZkSync *IZkSyncCaller) GetName(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _IZkSync.contract.Call(opts, &out, "getName")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// GetName is a free data retrieval call binding the contract method 0x17d7de7c.
//
// Solidity: function getName() view returns(string)
func (_IZkSync *IZkSyncSession) GetName() (string, error) {
	return _IZkSync.Contract.GetName(&_IZkSync.CallOpts)
}

// GetName is a free data retrieval call binding the contract method 0x17d7de7c.
//
// Solidity: function getName() view returns(string)
func (_IZkSync *IZkSyncCallerSession) GetName() (string, error) {
	return _IZkSync.Contract.GetName(&_IZkSync.CallOpts)
}

// GetPendingGovernor is a free data retrieval call binding the contract method 0x8665b150.
//
// Solidity: function getPendingGovernor() view returns(address)
func (_IZkSync *IZkSyncCaller) GetPendingGovernor(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _IZkSync.contract.Call(opts, &out, "getPendingGovernor")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetPendingGovernor is a free data retrieval call binding the contract method 0x8665b150.
//
// Solidity: function getPendingGovernor() view returns(address)
func (_IZkSync *IZkSyncSession) GetPendingGovernor() (common.Address, error) {
	return _IZkSync.Contract.GetPendingGovernor(&_IZkSync.CallOpts)
}

// GetPendingGovernor is a free data retrieval call binding the contract method 0x8665b150.
//
// Solidity: function getPendingGovernor() view returns(address)
func (_IZkSync *IZkSyncCallerSession) GetPendingGovernor() (common.Address, error) {
	return _IZkSync.Contract.GetPendingGovernor(&_IZkSync.CallOpts)
}

// GetPriorityQueueSize is a free data retrieval call binding the contract method 0x631f4bac.
//
// Solidity: function getPriorityQueueSize() view returns(uint256)
func (_IZkSync *IZkSyncCaller) GetPriorityQueueSize(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _IZkSync.contract.Call(opts, &out, "getPriorityQueueSize")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetPriorityQueueSize is a free data retrieval call binding the contract method 0x631f4bac.
//
// Solidity: function getPriorityQueueSize() view returns(uint256)
func (_IZkSync *IZkSyncSession) GetPriorityQueueSize() (*big.Int, error) {
	return _IZkSync.Contract.GetPriorityQueueSize(&_IZkSync.CallOpts)
}

// GetPriorityQueueSize is a free data retrieval call binding the contract method 0x631f4bac.
//
// Solidity: function getPriorityQueueSize() view returns(uint256)
func (_IZkSync *IZkSyncCallerSession) GetPriorityQueueSize() (*big.Int, error) {
	return _IZkSync.Contract.GetPriorityQueueSize(&_IZkSync.CallOpts)
}

// GetPriorityTxMaxGasLimit is a free data retrieval call binding the contract method 0x0ec6b0b7.
//
// Solidity: function getPriorityTxMaxGasLimit() view returns(uint256)
func (_IZkSync *IZkSyncCaller) GetPriorityTxMaxGasLimit(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _IZkSync.contract.Call(opts, &out, "getPriorityTxMaxGasLimit")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetPriorityTxMaxGasLimit is a free data retrieval call binding the contract method 0x0ec6b0b7.
//
// Solidity: function getPriorityTxMaxGasLimit() view returns(uint256)
func (_IZkSync *IZkSyncSession) GetPriorityTxMaxGasLimit() (*big.Int, error) {
	return _IZkSync.Contract.GetPriorityTxMaxGasLimit(&_IZkSync.CallOpts)
}

// GetPriorityTxMaxGasLimit is a free data retrieval call binding the contract method 0x0ec6b0b7.
//
// Solidity: function getPriorityTxMaxGasLimit() view returns(uint256)
func (_IZkSync *IZkSyncCallerSession) GetPriorityTxMaxGasLimit() (*big.Int, error) {
	return _IZkSync.Contract.GetPriorityTxMaxGasLimit(&_IZkSync.CallOpts)
}

// GetProtocolVersion is a free data retrieval call binding the contract method 0x33ce93fe.
//
// Solidity: function getProtocolVersion() view returns(uint256)
func (_IZkSync *IZkSyncCaller) GetProtocolVersion(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _IZkSync.contract.Call(opts, &out, "getProtocolVersion")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetProtocolVersion is a free data retrieval call binding the contract method 0x33ce93fe.
//
// Solidity: function getProtocolVersion() view returns(uint256)
func (_IZkSync *IZkSyncSession) GetProtocolVersion() (*big.Int, error) {
	return _IZkSync.Contract.GetProtocolVersion(&_IZkSync.CallOpts)
}

// GetProtocolVersion is a free data retrieval call binding the contract method 0x33ce93fe.
//
// Solidity: function getProtocolVersion() view returns(uint256)
func (_IZkSync *IZkSyncCallerSession) GetProtocolVersion() (*big.Int, error) {
	return _IZkSync.Contract.GetProtocolVersion(&_IZkSync.CallOpts)
}

// GetTotalBatchesCommitted is a free data retrieval call binding the contract method 0xdb1f0bf9.
//
// Solidity: function getTotalBatchesCommitted() view returns(uint256)
func (_IZkSync *IZkSyncCaller) GetTotalBatchesCommitted(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _IZkSync.contract.Call(opts, &out, "getTotalBatchesCommitted")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetTotalBatchesCommitted is a free data retrieval call binding the contract method 0xdb1f0bf9.
//
// Solidity: function getTotalBatchesCommitted() view returns(uint256)
func (_IZkSync *IZkSyncSession) GetTotalBatchesCommitted() (*big.Int, error) {
	return _IZkSync.Contract.GetTotalBatchesCommitted(&_IZkSync.CallOpts)
}

// GetTotalBatchesCommitted is a free data retrieval call binding the contract method 0xdb1f0bf9.
//
// Solidity: function getTotalBatchesCommitted() view returns(uint256)
func (_IZkSync *IZkSyncCallerSession) GetTotalBatchesCommitted() (*big.Int, error) {
	return _IZkSync.Contract.GetTotalBatchesCommitted(&_IZkSync.CallOpts)
}

// GetTotalBatchesExecuted is a free data retrieval call binding the contract method 0xb8c2f66f.
//
// Solidity: function getTotalBatchesExecuted() view returns(uint256)
func (_IZkSync *IZkSyncCaller) GetTotalBatchesExecuted(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _IZkSync.contract.Call(opts, &out, "getTotalBatchesExecuted")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetTotalBatchesExecuted is a free data retrieval call binding the contract method 0xb8c2f66f.
//
// Solidity: function getTotalBatchesExecuted() view returns(uint256)
func (_IZkSync *IZkSyncSession) GetTotalBatchesExecuted() (*big.Int, error) {
	return _IZkSync.Contract.GetTotalBatchesExecuted(&_IZkSync.CallOpts)
}

// GetTotalBatchesExecuted is a free data retrieval call binding the contract method 0xb8c2f66f.
//
// Solidity: function getTotalBatchesExecuted() view returns(uint256)
func (_IZkSync *IZkSyncCallerSession) GetTotalBatchesExecuted() (*big.Int, error) {
	return _IZkSync.Contract.GetTotalBatchesExecuted(&_IZkSync.CallOpts)
}

// GetTotalBatchesVerified is a free data retrieval call binding the contract method 0xef3f0bae.
//
// Solidity: function getTotalBatchesVerified() view returns(uint256)
func (_IZkSync *IZkSyncCaller) GetTotalBatchesVerified(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _IZkSync.contract.Call(opts, &out, "getTotalBatchesVerified")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetTotalBatchesVerified is a free data retrieval call binding the contract method 0xef3f0bae.
//
// Solidity: function getTotalBatchesVerified() view returns(uint256)
func (_IZkSync *IZkSyncSession) GetTotalBatchesVerified() (*big.Int, error) {
	return _IZkSync.Contract.GetTotalBatchesVerified(&_IZkSync.CallOpts)
}

// GetTotalBatchesVerified is a free data retrieval call binding the contract method 0xef3f0bae.
//
// Solidity: function getTotalBatchesVerified() view returns(uint256)
func (_IZkSync *IZkSyncCallerSession) GetTotalBatchesVerified() (*big.Int, error) {
	return _IZkSync.Contract.GetTotalBatchesVerified(&_IZkSync.CallOpts)
}

// GetTotalPriorityTxs is a free data retrieval call binding the contract method 0xa1954fc5.
//
// Solidity: function getTotalPriorityTxs() view returns(uint256)
func (_IZkSync *IZkSyncCaller) GetTotalPriorityTxs(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _IZkSync.contract.Call(opts, &out, "getTotalPriorityTxs")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetTotalPriorityTxs is a free data retrieval call binding the contract method 0xa1954fc5.
//
// Solidity: function getTotalPriorityTxs() view returns(uint256)
func (_IZkSync *IZkSyncSession) GetTotalPriorityTxs() (*big.Int, error) {
	return _IZkSync.Contract.GetTotalPriorityTxs(&_IZkSync.CallOpts)
}

// GetTotalPriorityTxs is a free data retrieval call binding the contract method 0xa1954fc5.
//
// Solidity: function getTotalPriorityTxs() view returns(uint256)
func (_IZkSync *IZkSyncCallerSession) GetTotalPriorityTxs() (*big.Int, error) {
	return _IZkSync.Contract.GetTotalPriorityTxs(&_IZkSync.CallOpts)
}

// GetVerifier is a free data retrieval call binding the contract method 0x46657fe9.
//
// Solidity: function getVerifier() view returns(address)
func (_IZkSync *IZkSyncCaller) GetVerifier(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _IZkSync.contract.Call(opts, &out, "getVerifier")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetVerifier is a free data retrieval call binding the contract method 0x46657fe9.
//
// Solidity: function getVerifier() view returns(address)
func (_IZkSync *IZkSyncSession) GetVerifier() (common.Address, error) {
	return _IZkSync.Contract.GetVerifier(&_IZkSync.CallOpts)
}

// GetVerifier is a free data retrieval call binding the contract method 0x46657fe9.
//
// Solidity: function getVerifier() view returns(address)
func (_IZkSync *IZkSyncCallerSession) GetVerifier() (common.Address, error) {
	return _IZkSync.Contract.GetVerifier(&_IZkSync.CallOpts)
}

// GetVerifierParams is a free data retrieval call binding the contract method 0x18e3a941.
//
// Solidity: function getVerifierParams() view returns((bytes32,bytes32,bytes32))
func (_IZkSync *IZkSyncCaller) GetVerifierParams(opts *bind.CallOpts) (VerifierParams, error) {
	var out []interface{}
	err := _IZkSync.contract.Call(opts, &out, "getVerifierParams")

	if err != nil {
		return *new(VerifierParams), err
	}

	out0 := *abi.ConvertType(out[0], new(VerifierParams)).(*VerifierParams)

	return out0, err

}

// GetVerifierParams is a free data retrieval call binding the contract method 0x18e3a941.
//
// Solidity: function getVerifierParams() view returns((bytes32,bytes32,bytes32))
func (_IZkSync *IZkSyncSession) GetVerifierParams() (VerifierParams, error) {
	return _IZkSync.Contract.GetVerifierParams(&_IZkSync.CallOpts)
}

// GetVerifierParams is a free data retrieval call binding the contract method 0x18e3a941.
//
// Solidity: function getVerifierParams() view returns((bytes32,bytes32,bytes32))
func (_IZkSync *IZkSyncCallerSession) GetVerifierParams() (VerifierParams, error) {
	return _IZkSync.Contract.GetVerifierParams(&_IZkSync.CallOpts)
}

// IsDiamondStorageFrozen is a free data retrieval call binding the contract method 0x29b98c67.
//
// Solidity: function isDiamondStorageFrozen() view returns(bool)
func (_IZkSync *IZkSyncCaller) IsDiamondStorageFrozen(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _IZkSync.contract.Call(opts, &out, "isDiamondStorageFrozen")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsDiamondStorageFrozen is a free data retrieval call binding the contract method 0x29b98c67.
//
// Solidity: function isDiamondStorageFrozen() view returns(bool)
func (_IZkSync *IZkSyncSession) IsDiamondStorageFrozen() (bool, error) {
	return _IZkSync.Contract.IsDiamondStorageFrozen(&_IZkSync.CallOpts)
}

// IsDiamondStorageFrozen is a free data retrieval call binding the contract method 0x29b98c67.
//
// Solidity: function isDiamondStorageFrozen() view returns(bool)
func (_IZkSync *IZkSyncCallerSession) IsDiamondStorageFrozen() (bool, error) {
	return _IZkSync.Contract.IsDiamondStorageFrozen(&_IZkSync.CallOpts)
}

// IsEthWithdrawalFinalized is a free data retrieval call binding the contract method 0xbd7c5412.
//
// Solidity: function isEthWithdrawalFinalized(uint256 _l2BatchNumber, uint256 _l2MessageIndex) view returns(bool)
func (_IZkSync *IZkSyncCaller) IsEthWithdrawalFinalized(opts *bind.CallOpts, _l2BatchNumber *big.Int, _l2MessageIndex *big.Int) (bool, error) {
	var out []interface{}
	err := _IZkSync.contract.Call(opts, &out, "isEthWithdrawalFinalized", _l2BatchNumber, _l2MessageIndex)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsEthWithdrawalFinalized is a free data retrieval call binding the contract method 0xbd7c5412.
//
// Solidity: function isEthWithdrawalFinalized(uint256 _l2BatchNumber, uint256 _l2MessageIndex) view returns(bool)
func (_IZkSync *IZkSyncSession) IsEthWithdrawalFinalized(_l2BatchNumber *big.Int, _l2MessageIndex *big.Int) (bool, error) {
	return _IZkSync.Contract.IsEthWithdrawalFinalized(&_IZkSync.CallOpts, _l2BatchNumber, _l2MessageIndex)
}

// IsEthWithdrawalFinalized is a free data retrieval call binding the contract method 0xbd7c5412.
//
// Solidity: function isEthWithdrawalFinalized(uint256 _l2BatchNumber, uint256 _l2MessageIndex) view returns(bool)
func (_IZkSync *IZkSyncCallerSession) IsEthWithdrawalFinalized(_l2BatchNumber *big.Int, _l2MessageIndex *big.Int) (bool, error) {
	return _IZkSync.Contract.IsEthWithdrawalFinalized(&_IZkSync.CallOpts, _l2BatchNumber, _l2MessageIndex)
}

// IsFacetFreezable is a free data retrieval call binding the contract method 0xc3bbd2d7.
//
// Solidity: function isFacetFreezable(address _facet) view returns(bool isFreezable)
func (_IZkSync *IZkSyncCaller) IsFacetFreezable(opts *bind.CallOpts, _facet common.Address) (bool, error) {
	var out []interface{}
	err := _IZkSync.contract.Call(opts, &out, "isFacetFreezable", _facet)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsFacetFreezable is a free data retrieval call binding the contract method 0xc3bbd2d7.
//
// Solidity: function isFacetFreezable(address _facet) view returns(bool isFreezable)
func (_IZkSync *IZkSyncSession) IsFacetFreezable(_facet common.Address) (bool, error) {
	return _IZkSync.Contract.IsFacetFreezable(&_IZkSync.CallOpts, _facet)
}

// IsFacetFreezable is a free data retrieval call binding the contract method 0xc3bbd2d7.
//
// Solidity: function isFacetFreezable(address _facet) view returns(bool isFreezable)
func (_IZkSync *IZkSyncCallerSession) IsFacetFreezable(_facet common.Address) (bool, error) {
	return _IZkSync.Contract.IsFacetFreezable(&_IZkSync.CallOpts, _facet)
}

// IsFunctionFreezable is a free data retrieval call binding the contract method 0xe81e0ba1.
//
// Solidity: function isFunctionFreezable(bytes4 _selector) view returns(bool)
func (_IZkSync *IZkSyncCaller) IsFunctionFreezable(opts *bind.CallOpts, _selector [4]byte) (bool, error) {
	var out []interface{}
	err := _IZkSync.contract.Call(opts, &out, "isFunctionFreezable", _selector)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsFunctionFreezable is a free data retrieval call binding the contract method 0xe81e0ba1.
//
// Solidity: function isFunctionFreezable(bytes4 _selector) view returns(bool)
func (_IZkSync *IZkSyncSession) IsFunctionFreezable(_selector [4]byte) (bool, error) {
	return _IZkSync.Contract.IsFunctionFreezable(&_IZkSync.CallOpts, _selector)
}

// IsFunctionFreezable is a free data retrieval call binding the contract method 0xe81e0ba1.
//
// Solidity: function isFunctionFreezable(bytes4 _selector) view returns(bool)
func (_IZkSync *IZkSyncCallerSession) IsFunctionFreezable(_selector [4]byte) (bool, error) {
	return _IZkSync.Contract.IsFunctionFreezable(&_IZkSync.CallOpts, _selector)
}

// IsValidator is a free data retrieval call binding the contract method 0xfacd743b.
//
// Solidity: function isValidator(address _address) view returns(bool)
func (_IZkSync *IZkSyncCaller) IsValidator(opts *bind.CallOpts, _address common.Address) (bool, error) {
	var out []interface{}
	err := _IZkSync.contract.Call(opts, &out, "isValidator", _address)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsValidator is a free data retrieval call binding the contract method 0xfacd743b.
//
// Solidity: function isValidator(address _address) view returns(bool)
func (_IZkSync *IZkSyncSession) IsValidator(_address common.Address) (bool, error) {
	return _IZkSync.Contract.IsValidator(&_IZkSync.CallOpts, _address)
}

// IsValidator is a free data retrieval call binding the contract method 0xfacd743b.
//
// Solidity: function isValidator(address _address) view returns(bool)
func (_IZkSync *IZkSyncCallerSession) IsValidator(_address common.Address) (bool, error) {
	return _IZkSync.Contract.IsValidator(&_IZkSync.CallOpts, _address)
}

// L2LogsRootHash is a free data retrieval call binding the contract method 0x9cd939e4.
//
// Solidity: function l2LogsRootHash(uint256 _batchNumber) view returns(bytes32 hash)
func (_IZkSync *IZkSyncCaller) L2LogsRootHash(opts *bind.CallOpts, _batchNumber *big.Int) ([32]byte, error) {
	var out []interface{}
	err := _IZkSync.contract.Call(opts, &out, "l2LogsRootHash", _batchNumber)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// L2LogsRootHash is a free data retrieval call binding the contract method 0x9cd939e4.
//
// Solidity: function l2LogsRootHash(uint256 _batchNumber) view returns(bytes32 hash)
func (_IZkSync *IZkSyncSession) L2LogsRootHash(_batchNumber *big.Int) ([32]byte, error) {
	return _IZkSync.Contract.L2LogsRootHash(&_IZkSync.CallOpts, _batchNumber)
}

// L2LogsRootHash is a free data retrieval call binding the contract method 0x9cd939e4.
//
// Solidity: function l2LogsRootHash(uint256 _batchNumber) view returns(bytes32 hash)
func (_IZkSync *IZkSyncCallerSession) L2LogsRootHash(_batchNumber *big.Int) ([32]byte, error) {
	return _IZkSync.Contract.L2LogsRootHash(&_IZkSync.CallOpts, _batchNumber)
}

// L2TransactionBaseCost is a free data retrieval call binding the contract method 0xb473318e.
//
// Solidity: function l2TransactionBaseCost(uint256 _gasPrice, uint256 _l2GasLimit, uint256 _l2GasPerPubdataByteLimit) view returns(uint256)
func (_IZkSync *IZkSyncCaller) L2TransactionBaseCost(opts *bind.CallOpts, _gasPrice *big.Int, _l2GasLimit *big.Int, _l2GasPerPubdataByteLimit *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _IZkSync.contract.Call(opts, &out, "l2TransactionBaseCost", _gasPrice, _l2GasLimit, _l2GasPerPubdataByteLimit)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// L2TransactionBaseCost is a free data retrieval call binding the contract method 0xb473318e.
//
// Solidity: function l2TransactionBaseCost(uint256 _gasPrice, uint256 _l2GasLimit, uint256 _l2GasPerPubdataByteLimit) view returns(uint256)
func (_IZkSync *IZkSyncSession) L2TransactionBaseCost(_gasPrice *big.Int, _l2GasLimit *big.Int, _l2GasPerPubdataByteLimit *big.Int) (*big.Int, error) {
	return _IZkSync.Contract.L2TransactionBaseCost(&_IZkSync.CallOpts, _gasPrice, _l2GasLimit, _l2GasPerPubdataByteLimit)
}

// L2TransactionBaseCost is a free data retrieval call binding the contract method 0xb473318e.
//
// Solidity: function l2TransactionBaseCost(uint256 _gasPrice, uint256 _l2GasLimit, uint256 _l2GasPerPubdataByteLimit) view returns(uint256)
func (_IZkSync *IZkSyncCallerSession) L2TransactionBaseCost(_gasPrice *big.Int, _l2GasLimit *big.Int, _l2GasPerPubdataByteLimit *big.Int) (*big.Int, error) {
	return _IZkSync.Contract.L2TransactionBaseCost(&_IZkSync.CallOpts, _gasPrice, _l2GasLimit, _l2GasPerPubdataByteLimit)
}

// PriorityQueueFrontOperation is a free data retrieval call binding the contract method 0x56142d7a.
//
// Solidity: function priorityQueueFrontOperation() view returns((bytes32,uint64,uint192))
func (_IZkSync *IZkSyncCaller) PriorityQueueFrontOperation(opts *bind.CallOpts) (PriorityOperation, error) {
	var out []interface{}
	err := _IZkSync.contract.Call(opts, &out, "priorityQueueFrontOperation")

	if err != nil {
		return *new(PriorityOperation), err
	}

	out0 := *abi.ConvertType(out[0], new(PriorityOperation)).(*PriorityOperation)

	return out0, err

}

// PriorityQueueFrontOperation is a free data retrieval call binding the contract method 0x56142d7a.
//
// Solidity: function priorityQueueFrontOperation() view returns((bytes32,uint64,uint192))
func (_IZkSync *IZkSyncSession) PriorityQueueFrontOperation() (PriorityOperation, error) {
	return _IZkSync.Contract.PriorityQueueFrontOperation(&_IZkSync.CallOpts)
}

// PriorityQueueFrontOperation is a free data retrieval call binding the contract method 0x56142d7a.
//
// Solidity: function priorityQueueFrontOperation() view returns((bytes32,uint64,uint192))
func (_IZkSync *IZkSyncCallerSession) PriorityQueueFrontOperation() (PriorityOperation, error) {
	return _IZkSync.Contract.PriorityQueueFrontOperation(&_IZkSync.CallOpts)
}

// ProveL1ToL2TransactionStatus is a free data retrieval call binding the contract method 0x042901c7.
//
// Solidity: function proveL1ToL2TransactionStatus(bytes32 _l2TxHash, uint256 _l2BatchNumber, uint256 _l2MessageIndex, uint16 _l2TxNumberInBatch, bytes32[] _merkleProof, uint8 _status) view returns(bool)
func (_IZkSync *IZkSyncCaller) ProveL1ToL2TransactionStatus(opts *bind.CallOpts, _l2TxHash [32]byte, _l2BatchNumber *big.Int, _l2MessageIndex *big.Int, _l2TxNumberInBatch uint16, _merkleProof [][32]byte, _status uint8) (bool, error) {
	var out []interface{}
	err := _IZkSync.contract.Call(opts, &out, "proveL1ToL2TransactionStatus", _l2TxHash, _l2BatchNumber, _l2MessageIndex, _l2TxNumberInBatch, _merkleProof, _status)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// ProveL1ToL2TransactionStatus is a free data retrieval call binding the contract method 0x042901c7.
//
// Solidity: function proveL1ToL2TransactionStatus(bytes32 _l2TxHash, uint256 _l2BatchNumber, uint256 _l2MessageIndex, uint16 _l2TxNumberInBatch, bytes32[] _merkleProof, uint8 _status) view returns(bool)
func (_IZkSync *IZkSyncSession) ProveL1ToL2TransactionStatus(_l2TxHash [32]byte, _l2BatchNumber *big.Int, _l2MessageIndex *big.Int, _l2TxNumberInBatch uint16, _merkleProof [][32]byte, _status uint8) (bool, error) {
	return _IZkSync.Contract.ProveL1ToL2TransactionStatus(&_IZkSync.CallOpts, _l2TxHash, _l2BatchNumber, _l2MessageIndex, _l2TxNumberInBatch, _merkleProof, _status)
}

// ProveL1ToL2TransactionStatus is a free data retrieval call binding the contract method 0x042901c7.
//
// Solidity: function proveL1ToL2TransactionStatus(bytes32 _l2TxHash, uint256 _l2BatchNumber, uint256 _l2MessageIndex, uint16 _l2TxNumberInBatch, bytes32[] _merkleProof, uint8 _status) view returns(bool)
func (_IZkSync *IZkSyncCallerSession) ProveL1ToL2TransactionStatus(_l2TxHash [32]byte, _l2BatchNumber *big.Int, _l2MessageIndex *big.Int, _l2TxNumberInBatch uint16, _merkleProof [][32]byte, _status uint8) (bool, error) {
	return _IZkSync.Contract.ProveL1ToL2TransactionStatus(&_IZkSync.CallOpts, _l2TxHash, _l2BatchNumber, _l2MessageIndex, _l2TxNumberInBatch, _merkleProof, _status)
}

// ProveL2LogInclusion is a free data retrieval call binding the contract method 0x263b7f8e.
//
// Solidity: function proveL2LogInclusion(uint256 _l2BatchNumber, uint256 _index, (uint8,bool,uint16,address,bytes32,bytes32) _log, bytes32[] _proof) view returns(bool)
func (_IZkSync *IZkSyncCaller) ProveL2LogInclusion(opts *bind.CallOpts, _l2BatchNumber *big.Int, _index *big.Int, _log L2Log, _proof [][32]byte) (bool, error) {
	var out []interface{}
	err := _IZkSync.contract.Call(opts, &out, "proveL2LogInclusion", _l2BatchNumber, _index, _log, _proof)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// ProveL2LogInclusion is a free data retrieval call binding the contract method 0x263b7f8e.
//
// Solidity: function proveL2LogInclusion(uint256 _l2BatchNumber, uint256 _index, (uint8,bool,uint16,address,bytes32,bytes32) _log, bytes32[] _proof) view returns(bool)
func (_IZkSync *IZkSyncSession) ProveL2LogInclusion(_l2BatchNumber *big.Int, _index *big.Int, _log L2Log, _proof [][32]byte) (bool, error) {
	return _IZkSync.Contract.ProveL2LogInclusion(&_IZkSync.CallOpts, _l2BatchNumber, _index, _log, _proof)
}

// ProveL2LogInclusion is a free data retrieval call binding the contract method 0x263b7f8e.
//
// Solidity: function proveL2LogInclusion(uint256 _l2BatchNumber, uint256 _index, (uint8,bool,uint16,address,bytes32,bytes32) _log, bytes32[] _proof) view returns(bool)
func (_IZkSync *IZkSyncCallerSession) ProveL2LogInclusion(_l2BatchNumber *big.Int, _index *big.Int, _log L2Log, _proof [][32]byte) (bool, error) {
	return _IZkSync.Contract.ProveL2LogInclusion(&_IZkSync.CallOpts, _l2BatchNumber, _index, _log, _proof)
}

// ProveL2MessageInclusion is a free data retrieval call binding the contract method 0xe4948f43.
//
// Solidity: function proveL2MessageInclusion(uint256 _l2BatchNumber, uint256 _index, (uint16,address,bytes) _message, bytes32[] _proof) view returns(bool)
func (_IZkSync *IZkSyncCaller) ProveL2MessageInclusion(opts *bind.CallOpts, _l2BatchNumber *big.Int, _index *big.Int, _message L2Message, _proof [][32]byte) (bool, error) {
	var out []interface{}
	err := _IZkSync.contract.Call(opts, &out, "proveL2MessageInclusion", _l2BatchNumber, _index, _message, _proof)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// ProveL2MessageInclusion is a free data retrieval call binding the contract method 0xe4948f43.
//
// Solidity: function proveL2MessageInclusion(uint256 _l2BatchNumber, uint256 _index, (uint16,address,bytes) _message, bytes32[] _proof) view returns(bool)
func (_IZkSync *IZkSyncSession) ProveL2MessageInclusion(_l2BatchNumber *big.Int, _index *big.Int, _message L2Message, _proof [][32]byte) (bool, error) {
	return _IZkSync.Contract.ProveL2MessageInclusion(&_IZkSync.CallOpts, _l2BatchNumber, _index, _message, _proof)
}

// ProveL2MessageInclusion is a free data retrieval call binding the contract method 0xe4948f43.
//
// Solidity: function proveL2MessageInclusion(uint256 _l2BatchNumber, uint256 _index, (uint16,address,bytes) _message, bytes32[] _proof) view returns(bool)
func (_IZkSync *IZkSyncCallerSession) ProveL2MessageInclusion(_l2BatchNumber *big.Int, _index *big.Int, _message L2Message, _proof [][32]byte) (bool, error) {
	return _IZkSync.Contract.ProveL2MessageInclusion(&_IZkSync.CallOpts, _l2BatchNumber, _index, _message, _proof)
}

// StoredBatchHash is a free data retrieval call binding the contract method 0xb22dd78e.
//
// Solidity: function storedBatchHash(uint256 _batchNumber) view returns(bytes32)
func (_IZkSync *IZkSyncCaller) StoredBatchHash(opts *bind.CallOpts, _batchNumber *big.Int) ([32]byte, error) {
	var out []interface{}
	err := _IZkSync.contract.Call(opts, &out, "storedBatchHash", _batchNumber)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// StoredBatchHash is a free data retrieval call binding the contract method 0xb22dd78e.
//
// Solidity: function storedBatchHash(uint256 _batchNumber) view returns(bytes32)
func (_IZkSync *IZkSyncSession) StoredBatchHash(_batchNumber *big.Int) ([32]byte, error) {
	return _IZkSync.Contract.StoredBatchHash(&_IZkSync.CallOpts, _batchNumber)
}

// StoredBatchHash is a free data retrieval call binding the contract method 0xb22dd78e.
//
// Solidity: function storedBatchHash(uint256 _batchNumber) view returns(bytes32)
func (_IZkSync *IZkSyncCallerSession) StoredBatchHash(_batchNumber *big.Int) ([32]byte, error) {
	return _IZkSync.Contract.StoredBatchHash(&_IZkSync.CallOpts, _batchNumber)
}

// AcceptAdmin is a paid mutator transaction binding the contract method 0x0e18b681.
//
// Solidity: function acceptAdmin() returns()
func (_IZkSync *IZkSyncTransactor) AcceptAdmin(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IZkSync.contract.Transact(opts, "acceptAdmin")
}

// AcceptAdmin is a paid mutator transaction binding the contract method 0x0e18b681.
//
// Solidity: function acceptAdmin() returns()
func (_IZkSync *IZkSyncSession) AcceptAdmin() (*types.Transaction, error) {
	return _IZkSync.Contract.AcceptAdmin(&_IZkSync.TransactOpts)
}

// AcceptAdmin is a paid mutator transaction binding the contract method 0x0e18b681.
//
// Solidity: function acceptAdmin() returns()
func (_IZkSync *IZkSyncTransactorSession) AcceptAdmin() (*types.Transaction, error) {
	return _IZkSync.Contract.AcceptAdmin(&_IZkSync.TransactOpts)
}

// AcceptGovernor is a paid mutator transaction binding the contract method 0xe58bb639.
//
// Solidity: function acceptGovernor() returns()
func (_IZkSync *IZkSyncTransactor) AcceptGovernor(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IZkSync.contract.Transact(opts, "acceptGovernor")
}

// AcceptGovernor is a paid mutator transaction binding the contract method 0xe58bb639.
//
// Solidity: function acceptGovernor() returns()
func (_IZkSync *IZkSyncSession) AcceptGovernor() (*types.Transaction, error) {
	return _IZkSync.Contract.AcceptGovernor(&_IZkSync.TransactOpts)
}

// AcceptGovernor is a paid mutator transaction binding the contract method 0xe58bb639.
//
// Solidity: function acceptGovernor() returns()
func (_IZkSync *IZkSyncTransactorSession) AcceptGovernor() (*types.Transaction, error) {
	return _IZkSync.Contract.AcceptGovernor(&_IZkSync.TransactOpts)
}

// CommitBatches is a paid mutator transaction binding the contract method 0x701f58c5.
//
// Solidity: function commitBatches((uint64,bytes32,uint64,uint256,bytes32,bytes32,uint256,bytes32) _lastCommittedBatchData, (uint64,uint64,uint64,bytes32,uint256,bytes32,bytes32,bytes32,bytes,bytes)[] _newBatchesData) returns()
func (_IZkSync *IZkSyncTransactor) CommitBatches(opts *bind.TransactOpts, _lastCommittedBatchData IExecutorStoredBatchInfo, _newBatchesData []IExecutorCommitBatchInfo) (*types.Transaction, error) {
	return _IZkSync.contract.Transact(opts, "commitBatches", _lastCommittedBatchData, _newBatchesData)
}

// CommitBatches is a paid mutator transaction binding the contract method 0x701f58c5.
//
// Solidity: function commitBatches((uint64,bytes32,uint64,uint256,bytes32,bytes32,uint256,bytes32) _lastCommittedBatchData, (uint64,uint64,uint64,bytes32,uint256,bytes32,bytes32,bytes32,bytes,bytes)[] _newBatchesData) returns()
func (_IZkSync *IZkSyncSession) CommitBatches(_lastCommittedBatchData IExecutorStoredBatchInfo, _newBatchesData []IExecutorCommitBatchInfo) (*types.Transaction, error) {
	return _IZkSync.Contract.CommitBatches(&_IZkSync.TransactOpts, _lastCommittedBatchData, _newBatchesData)
}

// CommitBatches is a paid mutator transaction binding the contract method 0x701f58c5.
//
// Solidity: function commitBatches((uint64,bytes32,uint64,uint256,bytes32,bytes32,uint256,bytes32) _lastCommittedBatchData, (uint64,uint64,uint64,bytes32,uint256,bytes32,bytes32,bytes32,bytes,bytes)[] _newBatchesData) returns()
func (_IZkSync *IZkSyncTransactorSession) CommitBatches(_lastCommittedBatchData IExecutorStoredBatchInfo, _newBatchesData []IExecutorCommitBatchInfo) (*types.Transaction, error) {
	return _IZkSync.Contract.CommitBatches(&_IZkSync.TransactOpts, _lastCommittedBatchData, _newBatchesData)
}

// ExecuteBatches is a paid mutator transaction binding the contract method 0xc3d93e7c.
//
// Solidity: function executeBatches((uint64,bytes32,uint64,uint256,bytes32,bytes32,uint256,bytes32)[] _batchesData) returns()
func (_IZkSync *IZkSyncTransactor) ExecuteBatches(opts *bind.TransactOpts, _batchesData []IExecutorStoredBatchInfo) (*types.Transaction, error) {
	return _IZkSync.contract.Transact(opts, "executeBatches", _batchesData)
}

// ExecuteBatches is a paid mutator transaction binding the contract method 0xc3d93e7c.
//
// Solidity: function executeBatches((uint64,bytes32,uint64,uint256,bytes32,bytes32,uint256,bytes32)[] _batchesData) returns()
func (_IZkSync *IZkSyncSession) ExecuteBatches(_batchesData []IExecutorStoredBatchInfo) (*types.Transaction, error) {
	return _IZkSync.Contract.ExecuteBatches(&_IZkSync.TransactOpts, _batchesData)
}

// ExecuteBatches is a paid mutator transaction binding the contract method 0xc3d93e7c.
//
// Solidity: function executeBatches((uint64,bytes32,uint64,uint256,bytes32,bytes32,uint256,bytes32)[] _batchesData) returns()
func (_IZkSync *IZkSyncTransactorSession) ExecuteBatches(_batchesData []IExecutorStoredBatchInfo) (*types.Transaction, error) {
	return _IZkSync.Contract.ExecuteBatches(&_IZkSync.TransactOpts, _batchesData)
}

// ExecuteUpgrade is a paid mutator transaction binding the contract method 0xa9f6d941.
//
// Solidity: function executeUpgrade(((address,uint8,bool,bytes4[])[],address,bytes) _diamondCut) returns()
func (_IZkSync *IZkSyncTransactor) ExecuteUpgrade(opts *bind.TransactOpts, _diamondCut DiamondDiamondCutData) (*types.Transaction, error) {
	return _IZkSync.contract.Transact(opts, "executeUpgrade", _diamondCut)
}

// ExecuteUpgrade is a paid mutator transaction binding the contract method 0xa9f6d941.
//
// Solidity: function executeUpgrade(((address,uint8,bool,bytes4[])[],address,bytes) _diamondCut) returns()
func (_IZkSync *IZkSyncSession) ExecuteUpgrade(_diamondCut DiamondDiamondCutData) (*types.Transaction, error) {
	return _IZkSync.Contract.ExecuteUpgrade(&_IZkSync.TransactOpts, _diamondCut)
}

// ExecuteUpgrade is a paid mutator transaction binding the contract method 0xa9f6d941.
//
// Solidity: function executeUpgrade(((address,uint8,bool,bytes4[])[],address,bytes) _diamondCut) returns()
func (_IZkSync *IZkSyncTransactorSession) ExecuteUpgrade(_diamondCut DiamondDiamondCutData) (*types.Transaction, error) {
	return _IZkSync.Contract.ExecuteUpgrade(&_IZkSync.TransactOpts, _diamondCut)
}

// FinalizeEthWithdrawal is a paid mutator transaction binding the contract method 0x6c0960f9.
//
// Solidity: function finalizeEthWithdrawal(uint256 _l2BatchNumber, uint256 _l2MessageIndex, uint16 _l2TxNumberInBatch, bytes _message, bytes32[] _merkleProof) returns()
func (_IZkSync *IZkSyncTransactor) FinalizeEthWithdrawal(opts *bind.TransactOpts, _l2BatchNumber *big.Int, _l2MessageIndex *big.Int, _l2TxNumberInBatch uint16, _message []byte, _merkleProof [][32]byte) (*types.Transaction, error) {
	return _IZkSync.contract.Transact(opts, "finalizeEthWithdrawal", _l2BatchNumber, _l2MessageIndex, _l2TxNumberInBatch, _message, _merkleProof)
}

// FinalizeEthWithdrawal is a paid mutator transaction binding the contract method 0x6c0960f9.
//
// Solidity: function finalizeEthWithdrawal(uint256 _l2BatchNumber, uint256 _l2MessageIndex, uint16 _l2TxNumberInBatch, bytes _message, bytes32[] _merkleProof) returns()
func (_IZkSync *IZkSyncSession) FinalizeEthWithdrawal(_l2BatchNumber *big.Int, _l2MessageIndex *big.Int, _l2TxNumberInBatch uint16, _message []byte, _merkleProof [][32]byte) (*types.Transaction, error) {
	return _IZkSync.Contract.FinalizeEthWithdrawal(&_IZkSync.TransactOpts, _l2BatchNumber, _l2MessageIndex, _l2TxNumberInBatch, _message, _merkleProof)
}

// FinalizeEthWithdrawal is a paid mutator transaction binding the contract method 0x6c0960f9.
//
// Solidity: function finalizeEthWithdrawal(uint256 _l2BatchNumber, uint256 _l2MessageIndex, uint16 _l2TxNumberInBatch, bytes _message, bytes32[] _merkleProof) returns()
func (_IZkSync *IZkSyncTransactorSession) FinalizeEthWithdrawal(_l2BatchNumber *big.Int, _l2MessageIndex *big.Int, _l2TxNumberInBatch uint16, _message []byte, _merkleProof [][32]byte) (*types.Transaction, error) {
	return _IZkSync.Contract.FinalizeEthWithdrawal(&_IZkSync.TransactOpts, _l2BatchNumber, _l2MessageIndex, _l2TxNumberInBatch, _message, _merkleProof)
}

// FreezeDiamond is a paid mutator transaction binding the contract method 0x27ae4c16.
//
// Solidity: function freezeDiamond() returns()
func (_IZkSync *IZkSyncTransactor) FreezeDiamond(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IZkSync.contract.Transact(opts, "freezeDiamond")
}

// FreezeDiamond is a paid mutator transaction binding the contract method 0x27ae4c16.
//
// Solidity: function freezeDiamond() returns()
func (_IZkSync *IZkSyncSession) FreezeDiamond() (*types.Transaction, error) {
	return _IZkSync.Contract.FreezeDiamond(&_IZkSync.TransactOpts)
}

// FreezeDiamond is a paid mutator transaction binding the contract method 0x27ae4c16.
//
// Solidity: function freezeDiamond() returns()
func (_IZkSync *IZkSyncTransactorSession) FreezeDiamond() (*types.Transaction, error) {
	return _IZkSync.Contract.FreezeDiamond(&_IZkSync.TransactOpts)
}

// ProveBatches is a paid mutator transaction binding the contract method 0x7f61885c.
//
// Solidity: function proveBatches((uint64,bytes32,uint64,uint256,bytes32,bytes32,uint256,bytes32) _prevBatch, (uint64,bytes32,uint64,uint256,bytes32,bytes32,uint256,bytes32)[] _committedBatches, (uint256[],uint256[]) _proof) returns()
func (_IZkSync *IZkSyncTransactor) ProveBatches(opts *bind.TransactOpts, _prevBatch IExecutorStoredBatchInfo, _committedBatches []IExecutorStoredBatchInfo, _proof IExecutorProofInput) (*types.Transaction, error) {
	return _IZkSync.contract.Transact(opts, "proveBatches", _prevBatch, _committedBatches, _proof)
}

// ProveBatches is a paid mutator transaction binding the contract method 0x7f61885c.
//
// Solidity: function proveBatches((uint64,bytes32,uint64,uint256,bytes32,bytes32,uint256,bytes32) _prevBatch, (uint64,bytes32,uint64,uint256,bytes32,bytes32,uint256,bytes32)[] _committedBatches, (uint256[],uint256[]) _proof) returns()
func (_IZkSync *IZkSyncSession) ProveBatches(_prevBatch IExecutorStoredBatchInfo, _committedBatches []IExecutorStoredBatchInfo, _proof IExecutorProofInput) (*types.Transaction, error) {
	return _IZkSync.Contract.ProveBatches(&_IZkSync.TransactOpts, _prevBatch, _committedBatches, _proof)
}

// ProveBatches is a paid mutator transaction binding the contract method 0x7f61885c.
//
// Solidity: function proveBatches((uint64,bytes32,uint64,uint256,bytes32,bytes32,uint256,bytes32) _prevBatch, (uint64,bytes32,uint64,uint256,bytes32,bytes32,uint256,bytes32)[] _committedBatches, (uint256[],uint256[]) _proof) returns()
func (_IZkSync *IZkSyncTransactorSession) ProveBatches(_prevBatch IExecutorStoredBatchInfo, _committedBatches []IExecutorStoredBatchInfo, _proof IExecutorProofInput) (*types.Transaction, error) {
	return _IZkSync.Contract.ProveBatches(&_IZkSync.TransactOpts, _prevBatch, _committedBatches, _proof)
}

// RequestL2Transaction is a paid mutator transaction binding the contract method 0xeb672419.
//
// Solidity: function requestL2Transaction(address _contractL2, uint256 _l2Value, bytes _calldata, uint256 _l2GasLimit, uint256 _l2GasPerPubdataByteLimit, bytes[] _factoryDeps, address _refundRecipient) payable returns(bytes32 canonicalTxHash)
func (_IZkSync *IZkSyncTransactor) RequestL2Transaction(opts *bind.TransactOpts, _contractL2 common.Address, _l2Value *big.Int, _calldata []byte, _l2GasLimit *big.Int, _l2GasPerPubdataByteLimit *big.Int, _factoryDeps [][]byte, _refundRecipient common.Address) (*types.Transaction, error) {
	return _IZkSync.contract.Transact(opts, "requestL2Transaction", _contractL2, _l2Value, _calldata, _l2GasLimit, _l2GasPerPubdataByteLimit, _factoryDeps, _refundRecipient)
}

// RequestL2Transaction is a paid mutator transaction binding the contract method 0xeb672419.
//
// Solidity: function requestL2Transaction(address _contractL2, uint256 _l2Value, bytes _calldata, uint256 _l2GasLimit, uint256 _l2GasPerPubdataByteLimit, bytes[] _factoryDeps, address _refundRecipient) payable returns(bytes32 canonicalTxHash)
func (_IZkSync *IZkSyncSession) RequestL2Transaction(_contractL2 common.Address, _l2Value *big.Int, _calldata []byte, _l2GasLimit *big.Int, _l2GasPerPubdataByteLimit *big.Int, _factoryDeps [][]byte, _refundRecipient common.Address) (*types.Transaction, error) {
	return _IZkSync.Contract.RequestL2Transaction(&_IZkSync.TransactOpts, _contractL2, _l2Value, _calldata, _l2GasLimit, _l2GasPerPubdataByteLimit, _factoryDeps, _refundRecipient)
}

// RequestL2Transaction is a paid mutator transaction binding the contract method 0xeb672419.
//
// Solidity: function requestL2Transaction(address _contractL2, uint256 _l2Value, bytes _calldata, uint256 _l2GasLimit, uint256 _l2GasPerPubdataByteLimit, bytes[] _factoryDeps, address _refundRecipient) payable returns(bytes32 canonicalTxHash)
func (_IZkSync *IZkSyncTransactorSession) RequestL2Transaction(_contractL2 common.Address, _l2Value *big.Int, _calldata []byte, _l2GasLimit *big.Int, _l2GasPerPubdataByteLimit *big.Int, _factoryDeps [][]byte, _refundRecipient common.Address) (*types.Transaction, error) {
	return _IZkSync.Contract.RequestL2Transaction(&_IZkSync.TransactOpts, _contractL2, _l2Value, _calldata, _l2GasLimit, _l2GasPerPubdataByteLimit, _factoryDeps, _refundRecipient)
}

// RevertBatches is a paid mutator transaction binding the contract method 0x97c09d34.
//
// Solidity: function revertBatches(uint256 _newLastBatch) returns()
func (_IZkSync *IZkSyncTransactor) RevertBatches(opts *bind.TransactOpts, _newLastBatch *big.Int) (*types.Transaction, error) {
	return _IZkSync.contract.Transact(opts, "revertBatches", _newLastBatch)
}

// RevertBatches is a paid mutator transaction binding the contract method 0x97c09d34.
//
// Solidity: function revertBatches(uint256 _newLastBatch) returns()
func (_IZkSync *IZkSyncSession) RevertBatches(_newLastBatch *big.Int) (*types.Transaction, error) {
	return _IZkSync.Contract.RevertBatches(&_IZkSync.TransactOpts, _newLastBatch)
}

// RevertBatches is a paid mutator transaction binding the contract method 0x97c09d34.
//
// Solidity: function revertBatches(uint256 _newLastBatch) returns()
func (_IZkSync *IZkSyncTransactorSession) RevertBatches(_newLastBatch *big.Int) (*types.Transaction, error) {
	return _IZkSync.Contract.RevertBatches(&_IZkSync.TransactOpts, _newLastBatch)
}

// SetPendingAdmin is a paid mutator transaction binding the contract method 0x4dd18bf5.
//
// Solidity: function setPendingAdmin(address _newPendingAdmin) returns()
func (_IZkSync *IZkSyncTransactor) SetPendingAdmin(opts *bind.TransactOpts, _newPendingAdmin common.Address) (*types.Transaction, error) {
	return _IZkSync.contract.Transact(opts, "setPendingAdmin", _newPendingAdmin)
}

// SetPendingAdmin is a paid mutator transaction binding the contract method 0x4dd18bf5.
//
// Solidity: function setPendingAdmin(address _newPendingAdmin) returns()
func (_IZkSync *IZkSyncSession) SetPendingAdmin(_newPendingAdmin common.Address) (*types.Transaction, error) {
	return _IZkSync.Contract.SetPendingAdmin(&_IZkSync.TransactOpts, _newPendingAdmin)
}

// SetPendingAdmin is a paid mutator transaction binding the contract method 0x4dd18bf5.
//
// Solidity: function setPendingAdmin(address _newPendingAdmin) returns()
func (_IZkSync *IZkSyncTransactorSession) SetPendingAdmin(_newPendingAdmin common.Address) (*types.Transaction, error) {
	return _IZkSync.Contract.SetPendingAdmin(&_IZkSync.TransactOpts, _newPendingAdmin)
}

// SetPendingGovernor is a paid mutator transaction binding the contract method 0xf235757f.
//
// Solidity: function setPendingGovernor(address _newPendingGovernor) returns()
func (_IZkSync *IZkSyncTransactor) SetPendingGovernor(opts *bind.TransactOpts, _newPendingGovernor common.Address) (*types.Transaction, error) {
	return _IZkSync.contract.Transact(opts, "setPendingGovernor", _newPendingGovernor)
}

// SetPendingGovernor is a paid mutator transaction binding the contract method 0xf235757f.
//
// Solidity: function setPendingGovernor(address _newPendingGovernor) returns()
func (_IZkSync *IZkSyncSession) SetPendingGovernor(_newPendingGovernor common.Address) (*types.Transaction, error) {
	return _IZkSync.Contract.SetPendingGovernor(&_IZkSync.TransactOpts, _newPendingGovernor)
}

// SetPendingGovernor is a paid mutator transaction binding the contract method 0xf235757f.
//
// Solidity: function setPendingGovernor(address _newPendingGovernor) returns()
func (_IZkSync *IZkSyncTransactorSession) SetPendingGovernor(_newPendingGovernor common.Address) (*types.Transaction, error) {
	return _IZkSync.Contract.SetPendingGovernor(&_IZkSync.TransactOpts, _newPendingGovernor)
}

// SetPorterAvailability is a paid mutator transaction binding the contract method 0x1cc5d103.
//
// Solidity: function setPorterAvailability(bool _zkPorterIsAvailable) returns()
func (_IZkSync *IZkSyncTransactor) SetPorterAvailability(opts *bind.TransactOpts, _zkPorterIsAvailable bool) (*types.Transaction, error) {
	return _IZkSync.contract.Transact(opts, "setPorterAvailability", _zkPorterIsAvailable)
}

// SetPorterAvailability is a paid mutator transaction binding the contract method 0x1cc5d103.
//
// Solidity: function setPorterAvailability(bool _zkPorterIsAvailable) returns()
func (_IZkSync *IZkSyncSession) SetPorterAvailability(_zkPorterIsAvailable bool) (*types.Transaction, error) {
	return _IZkSync.Contract.SetPorterAvailability(&_IZkSync.TransactOpts, _zkPorterIsAvailable)
}

// SetPorterAvailability is a paid mutator transaction binding the contract method 0x1cc5d103.
//
// Solidity: function setPorterAvailability(bool _zkPorterIsAvailable) returns()
func (_IZkSync *IZkSyncTransactorSession) SetPorterAvailability(_zkPorterIsAvailable bool) (*types.Transaction, error) {
	return _IZkSync.Contract.SetPorterAvailability(&_IZkSync.TransactOpts, _zkPorterIsAvailable)
}

// SetPriorityTxMaxGasLimit is a paid mutator transaction binding the contract method 0xbe6f11cf.
//
// Solidity: function setPriorityTxMaxGasLimit(uint256 _newPriorityTxMaxGasLimit) returns()
func (_IZkSync *IZkSyncTransactor) SetPriorityTxMaxGasLimit(opts *bind.TransactOpts, _newPriorityTxMaxGasLimit *big.Int) (*types.Transaction, error) {
	return _IZkSync.contract.Transact(opts, "setPriorityTxMaxGasLimit", _newPriorityTxMaxGasLimit)
}

// SetPriorityTxMaxGasLimit is a paid mutator transaction binding the contract method 0xbe6f11cf.
//
// Solidity: function setPriorityTxMaxGasLimit(uint256 _newPriorityTxMaxGasLimit) returns()
func (_IZkSync *IZkSyncSession) SetPriorityTxMaxGasLimit(_newPriorityTxMaxGasLimit *big.Int) (*types.Transaction, error) {
	return _IZkSync.Contract.SetPriorityTxMaxGasLimit(&_IZkSync.TransactOpts, _newPriorityTxMaxGasLimit)
}

// SetPriorityTxMaxGasLimit is a paid mutator transaction binding the contract method 0xbe6f11cf.
//
// Solidity: function setPriorityTxMaxGasLimit(uint256 _newPriorityTxMaxGasLimit) returns()
func (_IZkSync *IZkSyncTransactorSession) SetPriorityTxMaxGasLimit(_newPriorityTxMaxGasLimit *big.Int) (*types.Transaction, error) {
	return _IZkSync.Contract.SetPriorityTxMaxGasLimit(&_IZkSync.TransactOpts, _newPriorityTxMaxGasLimit)
}

// SetValidator is a paid mutator transaction binding the contract method 0x4623c91d.
//
// Solidity: function setValidator(address _validator, bool _active) returns()
func (_IZkSync *IZkSyncTransactor) SetValidator(opts *bind.TransactOpts, _validator common.Address, _active bool) (*types.Transaction, error) {
	return _IZkSync.contract.Transact(opts, "setValidator", _validator, _active)
}

// SetValidator is a paid mutator transaction binding the contract method 0x4623c91d.
//
// Solidity: function setValidator(address _validator, bool _active) returns()
func (_IZkSync *IZkSyncSession) SetValidator(_validator common.Address, _active bool) (*types.Transaction, error) {
	return _IZkSync.Contract.SetValidator(&_IZkSync.TransactOpts, _validator, _active)
}

// SetValidator is a paid mutator transaction binding the contract method 0x4623c91d.
//
// Solidity: function setValidator(address _validator, bool _active) returns()
func (_IZkSync *IZkSyncTransactorSession) SetValidator(_validator common.Address, _active bool) (*types.Transaction, error) {
	return _IZkSync.Contract.SetValidator(&_IZkSync.TransactOpts, _validator, _active)
}

// UnfreezeDiamond is a paid mutator transaction binding the contract method 0x17338945.
//
// Solidity: function unfreezeDiamond() returns()
func (_IZkSync *IZkSyncTransactor) UnfreezeDiamond(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IZkSync.contract.Transact(opts, "unfreezeDiamond")
}

// UnfreezeDiamond is a paid mutator transaction binding the contract method 0x17338945.
//
// Solidity: function unfreezeDiamond() returns()
func (_IZkSync *IZkSyncSession) UnfreezeDiamond() (*types.Transaction, error) {
	return _IZkSync.Contract.UnfreezeDiamond(&_IZkSync.TransactOpts)
}

// UnfreezeDiamond is a paid mutator transaction binding the contract method 0x17338945.
//
// Solidity: function unfreezeDiamond() returns()
func (_IZkSync *IZkSyncTransactorSession) UnfreezeDiamond() (*types.Transaction, error) {
	return _IZkSync.Contract.UnfreezeDiamond(&_IZkSync.TransactOpts)
}

// IZkSyncBlockCommitIterator is returned from FilterBlockCommit and is used to iterate over the raw logs and unpacked data for BlockCommit events raised by the IZkSync contract.
type IZkSyncBlockCommitIterator struct {
	Event *IZkSyncBlockCommit // Event containing the contract specifics and raw log

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
func (it *IZkSyncBlockCommitIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IZkSyncBlockCommit)
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
		it.Event = new(IZkSyncBlockCommit)
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
func (it *IZkSyncBlockCommitIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IZkSyncBlockCommitIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IZkSyncBlockCommit represents a BlockCommit event raised by the IZkSync contract.
type IZkSyncBlockCommit struct {
	BatchNumber *big.Int
	BatchHash   [32]byte
	Commitment  [32]byte
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterBlockCommit is a free log retrieval operation binding the contract event 0x8f2916b2f2d78cc5890ead36c06c0f6d5d112c7e103589947e8e2f0d6eddb763.
//
// Solidity: event BlockCommit(uint256 indexed batchNumber, bytes32 indexed batchHash, bytes32 indexed commitment)
func (_IZkSync *IZkSyncFilterer) FilterBlockCommit(opts *bind.FilterOpts, batchNumber []*big.Int, batchHash [][32]byte, commitment [][32]byte) (*IZkSyncBlockCommitIterator, error) {

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

	logs, sub, err := _IZkSync.contract.FilterLogs(opts, "BlockCommit", batchNumberRule, batchHashRule, commitmentRule)
	if err != nil {
		return nil, err
	}
	return &IZkSyncBlockCommitIterator{contract: _IZkSync.contract, event: "BlockCommit", logs: logs, sub: sub}, nil
}

// WatchBlockCommit is a free log subscription operation binding the contract event 0x8f2916b2f2d78cc5890ead36c06c0f6d5d112c7e103589947e8e2f0d6eddb763.
//
// Solidity: event BlockCommit(uint256 indexed batchNumber, bytes32 indexed batchHash, bytes32 indexed commitment)
func (_IZkSync *IZkSyncFilterer) WatchBlockCommit(opts *bind.WatchOpts, sink chan<- *IZkSyncBlockCommit, batchNumber []*big.Int, batchHash [][32]byte, commitment [][32]byte) (event.Subscription, error) {

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

	logs, sub, err := _IZkSync.contract.WatchLogs(opts, "BlockCommit", batchNumberRule, batchHashRule, commitmentRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IZkSyncBlockCommit)
				if err := _IZkSync.contract.UnpackLog(event, "BlockCommit", log); err != nil {
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
func (_IZkSync *IZkSyncFilterer) ParseBlockCommit(log types.Log) (*IZkSyncBlockCommit, error) {
	event := new(IZkSyncBlockCommit)
	if err := _IZkSync.contract.UnpackLog(event, "BlockCommit", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// IZkSyncBlockExecutionIterator is returned from FilterBlockExecution and is used to iterate over the raw logs and unpacked data for BlockExecution events raised by the IZkSync contract.
type IZkSyncBlockExecutionIterator struct {
	Event *IZkSyncBlockExecution // Event containing the contract specifics and raw log

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
func (it *IZkSyncBlockExecutionIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IZkSyncBlockExecution)
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
		it.Event = new(IZkSyncBlockExecution)
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
func (it *IZkSyncBlockExecutionIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IZkSyncBlockExecutionIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IZkSyncBlockExecution represents a BlockExecution event raised by the IZkSync contract.
type IZkSyncBlockExecution struct {
	BatchNumber *big.Int
	BatchHash   [32]byte
	Commitment  [32]byte
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterBlockExecution is a free log retrieval operation binding the contract event 0x2402307311a4d6604e4e7b4c8a15a7e1213edb39c16a31efa70afb06030d3165.
//
// Solidity: event BlockExecution(uint256 indexed batchNumber, bytes32 indexed batchHash, bytes32 indexed commitment)
func (_IZkSync *IZkSyncFilterer) FilterBlockExecution(opts *bind.FilterOpts, batchNumber []*big.Int, batchHash [][32]byte, commitment [][32]byte) (*IZkSyncBlockExecutionIterator, error) {

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

	logs, sub, err := _IZkSync.contract.FilterLogs(opts, "BlockExecution", batchNumberRule, batchHashRule, commitmentRule)
	if err != nil {
		return nil, err
	}
	return &IZkSyncBlockExecutionIterator{contract: _IZkSync.contract, event: "BlockExecution", logs: logs, sub: sub}, nil
}

// WatchBlockExecution is a free log subscription operation binding the contract event 0x2402307311a4d6604e4e7b4c8a15a7e1213edb39c16a31efa70afb06030d3165.
//
// Solidity: event BlockExecution(uint256 indexed batchNumber, bytes32 indexed batchHash, bytes32 indexed commitment)
func (_IZkSync *IZkSyncFilterer) WatchBlockExecution(opts *bind.WatchOpts, sink chan<- *IZkSyncBlockExecution, batchNumber []*big.Int, batchHash [][32]byte, commitment [][32]byte) (event.Subscription, error) {

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

	logs, sub, err := _IZkSync.contract.WatchLogs(opts, "BlockExecution", batchNumberRule, batchHashRule, commitmentRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IZkSyncBlockExecution)
				if err := _IZkSync.contract.UnpackLog(event, "BlockExecution", log); err != nil {
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
func (_IZkSync *IZkSyncFilterer) ParseBlockExecution(log types.Log) (*IZkSyncBlockExecution, error) {
	event := new(IZkSyncBlockExecution)
	if err := _IZkSync.contract.UnpackLog(event, "BlockExecution", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// IZkSyncBlocksRevertIterator is returned from FilterBlocksRevert and is used to iterate over the raw logs and unpacked data for BlocksRevert events raised by the IZkSync contract.
type IZkSyncBlocksRevertIterator struct {
	Event *IZkSyncBlocksRevert // Event containing the contract specifics and raw log

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
func (it *IZkSyncBlocksRevertIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IZkSyncBlocksRevert)
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
		it.Event = new(IZkSyncBlocksRevert)
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
func (it *IZkSyncBlocksRevertIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IZkSyncBlocksRevertIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IZkSyncBlocksRevert represents a BlocksRevert event raised by the IZkSync contract.
type IZkSyncBlocksRevert struct {
	TotalBatchesCommitted *big.Int
	TotalBatchesVerified  *big.Int
	TotalBatchesExecuted  *big.Int
	Raw                   types.Log // Blockchain specific contextual infos
}

// FilterBlocksRevert is a free log retrieval operation binding the contract event 0x8bd4b15ea7d1bc41ea9abc3fc487ccb89cd678a00786584714faa9d751c84ee5.
//
// Solidity: event BlocksRevert(uint256 totalBatchesCommitted, uint256 totalBatchesVerified, uint256 totalBatchesExecuted)
func (_IZkSync *IZkSyncFilterer) FilterBlocksRevert(opts *bind.FilterOpts) (*IZkSyncBlocksRevertIterator, error) {

	logs, sub, err := _IZkSync.contract.FilterLogs(opts, "BlocksRevert")
	if err != nil {
		return nil, err
	}
	return &IZkSyncBlocksRevertIterator{contract: _IZkSync.contract, event: "BlocksRevert", logs: logs, sub: sub}, nil
}

// WatchBlocksRevert is a free log subscription operation binding the contract event 0x8bd4b15ea7d1bc41ea9abc3fc487ccb89cd678a00786584714faa9d751c84ee5.
//
// Solidity: event BlocksRevert(uint256 totalBatchesCommitted, uint256 totalBatchesVerified, uint256 totalBatchesExecuted)
func (_IZkSync *IZkSyncFilterer) WatchBlocksRevert(opts *bind.WatchOpts, sink chan<- *IZkSyncBlocksRevert) (event.Subscription, error) {

	logs, sub, err := _IZkSync.contract.WatchLogs(opts, "BlocksRevert")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IZkSyncBlocksRevert)
				if err := _IZkSync.contract.UnpackLog(event, "BlocksRevert", log); err != nil {
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
func (_IZkSync *IZkSyncFilterer) ParseBlocksRevert(log types.Log) (*IZkSyncBlocksRevert, error) {
	event := new(IZkSyncBlocksRevert)
	if err := _IZkSync.contract.UnpackLog(event, "BlocksRevert", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// IZkSyncBlocksVerificationIterator is returned from FilterBlocksVerification and is used to iterate over the raw logs and unpacked data for BlocksVerification events raised by the IZkSync contract.
type IZkSyncBlocksVerificationIterator struct {
	Event *IZkSyncBlocksVerification // Event containing the contract specifics and raw log

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
func (it *IZkSyncBlocksVerificationIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IZkSyncBlocksVerification)
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
		it.Event = new(IZkSyncBlocksVerification)
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
func (it *IZkSyncBlocksVerificationIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IZkSyncBlocksVerificationIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IZkSyncBlocksVerification represents a BlocksVerification event raised by the IZkSync contract.
type IZkSyncBlocksVerification struct {
	PreviousLastVerifiedBatch *big.Int
	CurrentLastVerifiedBatch  *big.Int
	Raw                       types.Log // Blockchain specific contextual infos
}

// FilterBlocksVerification is a free log retrieval operation binding the contract event 0x22c9005dd88c18b552a1cd7e8b3b937fcde9ca69213c1f658f54d572e4877a81.
//
// Solidity: event BlocksVerification(uint256 indexed previousLastVerifiedBatch, uint256 indexed currentLastVerifiedBatch)
func (_IZkSync *IZkSyncFilterer) FilterBlocksVerification(opts *bind.FilterOpts, previousLastVerifiedBatch []*big.Int, currentLastVerifiedBatch []*big.Int) (*IZkSyncBlocksVerificationIterator, error) {

	var previousLastVerifiedBatchRule []interface{}
	for _, previousLastVerifiedBatchItem := range previousLastVerifiedBatch {
		previousLastVerifiedBatchRule = append(previousLastVerifiedBatchRule, previousLastVerifiedBatchItem)
	}
	var currentLastVerifiedBatchRule []interface{}
	for _, currentLastVerifiedBatchItem := range currentLastVerifiedBatch {
		currentLastVerifiedBatchRule = append(currentLastVerifiedBatchRule, currentLastVerifiedBatchItem)
	}

	logs, sub, err := _IZkSync.contract.FilterLogs(opts, "BlocksVerification", previousLastVerifiedBatchRule, currentLastVerifiedBatchRule)
	if err != nil {
		return nil, err
	}
	return &IZkSyncBlocksVerificationIterator{contract: _IZkSync.contract, event: "BlocksVerification", logs: logs, sub: sub}, nil
}

// WatchBlocksVerification is a free log subscription operation binding the contract event 0x22c9005dd88c18b552a1cd7e8b3b937fcde9ca69213c1f658f54d572e4877a81.
//
// Solidity: event BlocksVerification(uint256 indexed previousLastVerifiedBatch, uint256 indexed currentLastVerifiedBatch)
func (_IZkSync *IZkSyncFilterer) WatchBlocksVerification(opts *bind.WatchOpts, sink chan<- *IZkSyncBlocksVerification, previousLastVerifiedBatch []*big.Int, currentLastVerifiedBatch []*big.Int) (event.Subscription, error) {

	var previousLastVerifiedBatchRule []interface{}
	for _, previousLastVerifiedBatchItem := range previousLastVerifiedBatch {
		previousLastVerifiedBatchRule = append(previousLastVerifiedBatchRule, previousLastVerifiedBatchItem)
	}
	var currentLastVerifiedBatchRule []interface{}
	for _, currentLastVerifiedBatchItem := range currentLastVerifiedBatch {
		currentLastVerifiedBatchRule = append(currentLastVerifiedBatchRule, currentLastVerifiedBatchItem)
	}

	logs, sub, err := _IZkSync.contract.WatchLogs(opts, "BlocksVerification", previousLastVerifiedBatchRule, currentLastVerifiedBatchRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IZkSyncBlocksVerification)
				if err := _IZkSync.contract.UnpackLog(event, "BlocksVerification", log); err != nil {
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
func (_IZkSync *IZkSyncFilterer) ParseBlocksVerification(log types.Log) (*IZkSyncBlocksVerification, error) {
	event := new(IZkSyncBlocksVerification)
	if err := _IZkSync.contract.UnpackLog(event, "BlocksVerification", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// IZkSyncEthWithdrawalFinalizedIterator is returned from FilterEthWithdrawalFinalized and is used to iterate over the raw logs and unpacked data for EthWithdrawalFinalized events raised by the IZkSync contract.
type IZkSyncEthWithdrawalFinalizedIterator struct {
	Event *IZkSyncEthWithdrawalFinalized // Event containing the contract specifics and raw log

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
func (it *IZkSyncEthWithdrawalFinalizedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IZkSyncEthWithdrawalFinalized)
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
		it.Event = new(IZkSyncEthWithdrawalFinalized)
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
func (it *IZkSyncEthWithdrawalFinalizedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IZkSyncEthWithdrawalFinalizedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IZkSyncEthWithdrawalFinalized represents a EthWithdrawalFinalized event raised by the IZkSync contract.
type IZkSyncEthWithdrawalFinalized struct {
	To     common.Address
	Amount *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterEthWithdrawalFinalized is a free log retrieval operation binding the contract event 0x26464d64ddb13f6d187de632d165bd1065382ec0b66c25c648957116e7bc25c8.
//
// Solidity: event EthWithdrawalFinalized(address indexed to, uint256 amount)
func (_IZkSync *IZkSyncFilterer) FilterEthWithdrawalFinalized(opts *bind.FilterOpts, to []common.Address) (*IZkSyncEthWithdrawalFinalizedIterator, error) {

	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _IZkSync.contract.FilterLogs(opts, "EthWithdrawalFinalized", toRule)
	if err != nil {
		return nil, err
	}
	return &IZkSyncEthWithdrawalFinalizedIterator{contract: _IZkSync.contract, event: "EthWithdrawalFinalized", logs: logs, sub: sub}, nil
}

// WatchEthWithdrawalFinalized is a free log subscription operation binding the contract event 0x26464d64ddb13f6d187de632d165bd1065382ec0b66c25c648957116e7bc25c8.
//
// Solidity: event EthWithdrawalFinalized(address indexed to, uint256 amount)
func (_IZkSync *IZkSyncFilterer) WatchEthWithdrawalFinalized(opts *bind.WatchOpts, sink chan<- *IZkSyncEthWithdrawalFinalized, to []common.Address) (event.Subscription, error) {

	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _IZkSync.contract.WatchLogs(opts, "EthWithdrawalFinalized", toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IZkSyncEthWithdrawalFinalized)
				if err := _IZkSync.contract.UnpackLog(event, "EthWithdrawalFinalized", log); err != nil {
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
func (_IZkSync *IZkSyncFilterer) ParseEthWithdrawalFinalized(log types.Log) (*IZkSyncEthWithdrawalFinalized, error) {
	event := new(IZkSyncEthWithdrawalFinalized)
	if err := _IZkSync.contract.UnpackLog(event, "EthWithdrawalFinalized", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// IZkSyncExecuteUpgradeIterator is returned from FilterExecuteUpgrade and is used to iterate over the raw logs and unpacked data for ExecuteUpgrade events raised by the IZkSync contract.
type IZkSyncExecuteUpgradeIterator struct {
	Event *IZkSyncExecuteUpgrade // Event containing the contract specifics and raw log

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
func (it *IZkSyncExecuteUpgradeIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IZkSyncExecuteUpgrade)
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
		it.Event = new(IZkSyncExecuteUpgrade)
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
func (it *IZkSyncExecuteUpgradeIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IZkSyncExecuteUpgradeIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IZkSyncExecuteUpgrade represents a ExecuteUpgrade event raised by the IZkSync contract.
type IZkSyncExecuteUpgrade struct {
	DiamondCut DiamondDiamondCutData
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterExecuteUpgrade is a free log retrieval operation binding the contract event 0x1dabfc3f4f6a4e74e19be10cc9d1d8e23db03e415d5d3547a1a7d5eb766513ba.
//
// Solidity: event ExecuteUpgrade(((address,uint8,bool,bytes4[])[],address,bytes) diamondCut)
func (_IZkSync *IZkSyncFilterer) FilterExecuteUpgrade(opts *bind.FilterOpts) (*IZkSyncExecuteUpgradeIterator, error) {

	logs, sub, err := _IZkSync.contract.FilterLogs(opts, "ExecuteUpgrade")
	if err != nil {
		return nil, err
	}
	return &IZkSyncExecuteUpgradeIterator{contract: _IZkSync.contract, event: "ExecuteUpgrade", logs: logs, sub: sub}, nil
}

// WatchExecuteUpgrade is a free log subscription operation binding the contract event 0x1dabfc3f4f6a4e74e19be10cc9d1d8e23db03e415d5d3547a1a7d5eb766513ba.
//
// Solidity: event ExecuteUpgrade(((address,uint8,bool,bytes4[])[],address,bytes) diamondCut)
func (_IZkSync *IZkSyncFilterer) WatchExecuteUpgrade(opts *bind.WatchOpts, sink chan<- *IZkSyncExecuteUpgrade) (event.Subscription, error) {

	logs, sub, err := _IZkSync.contract.WatchLogs(opts, "ExecuteUpgrade")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IZkSyncExecuteUpgrade)
				if err := _IZkSync.contract.UnpackLog(event, "ExecuteUpgrade", log); err != nil {
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
func (_IZkSync *IZkSyncFilterer) ParseExecuteUpgrade(log types.Log) (*IZkSyncExecuteUpgrade, error) {
	event := new(IZkSyncExecuteUpgrade)
	if err := _IZkSync.contract.UnpackLog(event, "ExecuteUpgrade", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// IZkSyncFreezeIterator is returned from FilterFreeze and is used to iterate over the raw logs and unpacked data for Freeze events raised by the IZkSync contract.
type IZkSyncFreezeIterator struct {
	Event *IZkSyncFreeze // Event containing the contract specifics and raw log

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
func (it *IZkSyncFreezeIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IZkSyncFreeze)
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
		it.Event = new(IZkSyncFreeze)
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
func (it *IZkSyncFreezeIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IZkSyncFreezeIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IZkSyncFreeze represents a Freeze event raised by the IZkSync contract.
type IZkSyncFreeze struct {
	Raw types.Log // Blockchain specific contextual infos
}

// FilterFreeze is a free log retrieval operation binding the contract event 0x615acbaede366d76a8b8cb2a9ada6a71495f0786513d71aa97aaf0c3910b78de.
//
// Solidity: event Freeze()
func (_IZkSync *IZkSyncFilterer) FilterFreeze(opts *bind.FilterOpts) (*IZkSyncFreezeIterator, error) {

	logs, sub, err := _IZkSync.contract.FilterLogs(opts, "Freeze")
	if err != nil {
		return nil, err
	}
	return &IZkSyncFreezeIterator{contract: _IZkSync.contract, event: "Freeze", logs: logs, sub: sub}, nil
}

// WatchFreeze is a free log subscription operation binding the contract event 0x615acbaede366d76a8b8cb2a9ada6a71495f0786513d71aa97aaf0c3910b78de.
//
// Solidity: event Freeze()
func (_IZkSync *IZkSyncFilterer) WatchFreeze(opts *bind.WatchOpts, sink chan<- *IZkSyncFreeze) (event.Subscription, error) {

	logs, sub, err := _IZkSync.contract.WatchLogs(opts, "Freeze")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IZkSyncFreeze)
				if err := _IZkSync.contract.UnpackLog(event, "Freeze", log); err != nil {
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
func (_IZkSync *IZkSyncFilterer) ParseFreeze(log types.Log) (*IZkSyncFreeze, error) {
	event := new(IZkSyncFreeze)
	if err := _IZkSync.contract.UnpackLog(event, "Freeze", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// IZkSyncIsPorterAvailableStatusUpdateIterator is returned from FilterIsPorterAvailableStatusUpdate and is used to iterate over the raw logs and unpacked data for IsPorterAvailableStatusUpdate events raised by the IZkSync contract.
type IZkSyncIsPorterAvailableStatusUpdateIterator struct {
	Event *IZkSyncIsPorterAvailableStatusUpdate // Event containing the contract specifics and raw log

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
func (it *IZkSyncIsPorterAvailableStatusUpdateIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IZkSyncIsPorterAvailableStatusUpdate)
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
		it.Event = new(IZkSyncIsPorterAvailableStatusUpdate)
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
func (it *IZkSyncIsPorterAvailableStatusUpdateIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IZkSyncIsPorterAvailableStatusUpdateIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IZkSyncIsPorterAvailableStatusUpdate represents a IsPorterAvailableStatusUpdate event raised by the IZkSync contract.
type IZkSyncIsPorterAvailableStatusUpdate struct {
	IsPorterAvailable bool
	Raw               types.Log // Blockchain specific contextual infos
}

// FilterIsPorterAvailableStatusUpdate is a free log retrieval operation binding the contract event 0x036b81a8a07344698cb5aa4142c5669a9317c9ce905264a08f0b9f9331883936.
//
// Solidity: event IsPorterAvailableStatusUpdate(bool isPorterAvailable)
func (_IZkSync *IZkSyncFilterer) FilterIsPorterAvailableStatusUpdate(opts *bind.FilterOpts) (*IZkSyncIsPorterAvailableStatusUpdateIterator, error) {

	logs, sub, err := _IZkSync.contract.FilterLogs(opts, "IsPorterAvailableStatusUpdate")
	if err != nil {
		return nil, err
	}
	return &IZkSyncIsPorterAvailableStatusUpdateIterator{contract: _IZkSync.contract, event: "IsPorterAvailableStatusUpdate", logs: logs, sub: sub}, nil
}

// WatchIsPorterAvailableStatusUpdate is a free log subscription operation binding the contract event 0x036b81a8a07344698cb5aa4142c5669a9317c9ce905264a08f0b9f9331883936.
//
// Solidity: event IsPorterAvailableStatusUpdate(bool isPorterAvailable)
func (_IZkSync *IZkSyncFilterer) WatchIsPorterAvailableStatusUpdate(opts *bind.WatchOpts, sink chan<- *IZkSyncIsPorterAvailableStatusUpdate) (event.Subscription, error) {

	logs, sub, err := _IZkSync.contract.WatchLogs(opts, "IsPorterAvailableStatusUpdate")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IZkSyncIsPorterAvailableStatusUpdate)
				if err := _IZkSync.contract.UnpackLog(event, "IsPorterAvailableStatusUpdate", log); err != nil {
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
func (_IZkSync *IZkSyncFilterer) ParseIsPorterAvailableStatusUpdate(log types.Log) (*IZkSyncIsPorterAvailableStatusUpdate, error) {
	event := new(IZkSyncIsPorterAvailableStatusUpdate)
	if err := _IZkSync.contract.UnpackLog(event, "IsPorterAvailableStatusUpdate", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// IZkSyncNewAdminIterator is returned from FilterNewAdmin and is used to iterate over the raw logs and unpacked data for NewAdmin events raised by the IZkSync contract.
type IZkSyncNewAdminIterator struct {
	Event *IZkSyncNewAdmin // Event containing the contract specifics and raw log

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
func (it *IZkSyncNewAdminIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IZkSyncNewAdmin)
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
		it.Event = new(IZkSyncNewAdmin)
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
func (it *IZkSyncNewAdminIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IZkSyncNewAdminIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IZkSyncNewAdmin represents a NewAdmin event raised by the IZkSync contract.
type IZkSyncNewAdmin struct {
	OldAdmin common.Address
	NewAdmin common.Address
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterNewAdmin is a free log retrieval operation binding the contract event 0xf9ffabca9c8276e99321725bcb43fb076a6c66a54b7f21c4e8146d8519b417dc.
//
// Solidity: event NewAdmin(address indexed oldAdmin, address indexed newAdmin)
func (_IZkSync *IZkSyncFilterer) FilterNewAdmin(opts *bind.FilterOpts, oldAdmin []common.Address, newAdmin []common.Address) (*IZkSyncNewAdminIterator, error) {

	var oldAdminRule []interface{}
	for _, oldAdminItem := range oldAdmin {
		oldAdminRule = append(oldAdminRule, oldAdminItem)
	}
	var newAdminRule []interface{}
	for _, newAdminItem := range newAdmin {
		newAdminRule = append(newAdminRule, newAdminItem)
	}

	logs, sub, err := _IZkSync.contract.FilterLogs(opts, "NewAdmin", oldAdminRule, newAdminRule)
	if err != nil {
		return nil, err
	}
	return &IZkSyncNewAdminIterator{contract: _IZkSync.contract, event: "NewAdmin", logs: logs, sub: sub}, nil
}

// WatchNewAdmin is a free log subscription operation binding the contract event 0xf9ffabca9c8276e99321725bcb43fb076a6c66a54b7f21c4e8146d8519b417dc.
//
// Solidity: event NewAdmin(address indexed oldAdmin, address indexed newAdmin)
func (_IZkSync *IZkSyncFilterer) WatchNewAdmin(opts *bind.WatchOpts, sink chan<- *IZkSyncNewAdmin, oldAdmin []common.Address, newAdmin []common.Address) (event.Subscription, error) {

	var oldAdminRule []interface{}
	for _, oldAdminItem := range oldAdmin {
		oldAdminRule = append(oldAdminRule, oldAdminItem)
	}
	var newAdminRule []interface{}
	for _, newAdminItem := range newAdmin {
		newAdminRule = append(newAdminRule, newAdminItem)
	}

	logs, sub, err := _IZkSync.contract.WatchLogs(opts, "NewAdmin", oldAdminRule, newAdminRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IZkSyncNewAdmin)
				if err := _IZkSync.contract.UnpackLog(event, "NewAdmin", log); err != nil {
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
func (_IZkSync *IZkSyncFilterer) ParseNewAdmin(log types.Log) (*IZkSyncNewAdmin, error) {
	event := new(IZkSyncNewAdmin)
	if err := _IZkSync.contract.UnpackLog(event, "NewAdmin", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// IZkSyncNewGovernorIterator is returned from FilterNewGovernor and is used to iterate over the raw logs and unpacked data for NewGovernor events raised by the IZkSync contract.
type IZkSyncNewGovernorIterator struct {
	Event *IZkSyncNewGovernor // Event containing the contract specifics and raw log

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
func (it *IZkSyncNewGovernorIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IZkSyncNewGovernor)
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
		it.Event = new(IZkSyncNewGovernor)
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
func (it *IZkSyncNewGovernorIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IZkSyncNewGovernorIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IZkSyncNewGovernor represents a NewGovernor event raised by the IZkSync contract.
type IZkSyncNewGovernor struct {
	OldGovernor common.Address
	NewGovernor common.Address
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterNewGovernor is a free log retrieval operation binding the contract event 0x1ba669d4a78521f2ad26e8e0fcbcdd626a63f34d68f326bc232a3abe2a5d042a.
//
// Solidity: event NewGovernor(address indexed oldGovernor, address indexed newGovernor)
func (_IZkSync *IZkSyncFilterer) FilterNewGovernor(opts *bind.FilterOpts, oldGovernor []common.Address, newGovernor []common.Address) (*IZkSyncNewGovernorIterator, error) {

	var oldGovernorRule []interface{}
	for _, oldGovernorItem := range oldGovernor {
		oldGovernorRule = append(oldGovernorRule, oldGovernorItem)
	}
	var newGovernorRule []interface{}
	for _, newGovernorItem := range newGovernor {
		newGovernorRule = append(newGovernorRule, newGovernorItem)
	}

	logs, sub, err := _IZkSync.contract.FilterLogs(opts, "NewGovernor", oldGovernorRule, newGovernorRule)
	if err != nil {
		return nil, err
	}
	return &IZkSyncNewGovernorIterator{contract: _IZkSync.contract, event: "NewGovernor", logs: logs, sub: sub}, nil
}

// WatchNewGovernor is a free log subscription operation binding the contract event 0x1ba669d4a78521f2ad26e8e0fcbcdd626a63f34d68f326bc232a3abe2a5d042a.
//
// Solidity: event NewGovernor(address indexed oldGovernor, address indexed newGovernor)
func (_IZkSync *IZkSyncFilterer) WatchNewGovernor(opts *bind.WatchOpts, sink chan<- *IZkSyncNewGovernor, oldGovernor []common.Address, newGovernor []common.Address) (event.Subscription, error) {

	var oldGovernorRule []interface{}
	for _, oldGovernorItem := range oldGovernor {
		oldGovernorRule = append(oldGovernorRule, oldGovernorItem)
	}
	var newGovernorRule []interface{}
	for _, newGovernorItem := range newGovernor {
		newGovernorRule = append(newGovernorRule, newGovernorItem)
	}

	logs, sub, err := _IZkSync.contract.WatchLogs(opts, "NewGovernor", oldGovernorRule, newGovernorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IZkSyncNewGovernor)
				if err := _IZkSync.contract.UnpackLog(event, "NewGovernor", log); err != nil {
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

// ParseNewGovernor is a log parse operation binding the contract event 0x1ba669d4a78521f2ad26e8e0fcbcdd626a63f34d68f326bc232a3abe2a5d042a.
//
// Solidity: event NewGovernor(address indexed oldGovernor, address indexed newGovernor)
func (_IZkSync *IZkSyncFilterer) ParseNewGovernor(log types.Log) (*IZkSyncNewGovernor, error) {
	event := new(IZkSyncNewGovernor)
	if err := _IZkSync.contract.UnpackLog(event, "NewGovernor", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// IZkSyncNewPendingAdminIterator is returned from FilterNewPendingAdmin and is used to iterate over the raw logs and unpacked data for NewPendingAdmin events raised by the IZkSync contract.
type IZkSyncNewPendingAdminIterator struct {
	Event *IZkSyncNewPendingAdmin // Event containing the contract specifics and raw log

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
func (it *IZkSyncNewPendingAdminIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IZkSyncNewPendingAdmin)
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
		it.Event = new(IZkSyncNewPendingAdmin)
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
func (it *IZkSyncNewPendingAdminIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IZkSyncNewPendingAdminIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IZkSyncNewPendingAdmin represents a NewPendingAdmin event raised by the IZkSync contract.
type IZkSyncNewPendingAdmin struct {
	OldPendingAdmin common.Address
	NewPendingAdmin common.Address
	Raw             types.Log // Blockchain specific contextual infos
}

// FilterNewPendingAdmin is a free log retrieval operation binding the contract event 0xca4f2f25d0898edd99413412fb94012f9e54ec8142f9b093e7720646a95b16a9.
//
// Solidity: event NewPendingAdmin(address indexed oldPendingAdmin, address indexed newPendingAdmin)
func (_IZkSync *IZkSyncFilterer) FilterNewPendingAdmin(opts *bind.FilterOpts, oldPendingAdmin []common.Address, newPendingAdmin []common.Address) (*IZkSyncNewPendingAdminIterator, error) {

	var oldPendingAdminRule []interface{}
	for _, oldPendingAdminItem := range oldPendingAdmin {
		oldPendingAdminRule = append(oldPendingAdminRule, oldPendingAdminItem)
	}
	var newPendingAdminRule []interface{}
	for _, newPendingAdminItem := range newPendingAdmin {
		newPendingAdminRule = append(newPendingAdminRule, newPendingAdminItem)
	}

	logs, sub, err := _IZkSync.contract.FilterLogs(opts, "NewPendingAdmin", oldPendingAdminRule, newPendingAdminRule)
	if err != nil {
		return nil, err
	}
	return &IZkSyncNewPendingAdminIterator{contract: _IZkSync.contract, event: "NewPendingAdmin", logs: logs, sub: sub}, nil
}

// WatchNewPendingAdmin is a free log subscription operation binding the contract event 0xca4f2f25d0898edd99413412fb94012f9e54ec8142f9b093e7720646a95b16a9.
//
// Solidity: event NewPendingAdmin(address indexed oldPendingAdmin, address indexed newPendingAdmin)
func (_IZkSync *IZkSyncFilterer) WatchNewPendingAdmin(opts *bind.WatchOpts, sink chan<- *IZkSyncNewPendingAdmin, oldPendingAdmin []common.Address, newPendingAdmin []common.Address) (event.Subscription, error) {

	var oldPendingAdminRule []interface{}
	for _, oldPendingAdminItem := range oldPendingAdmin {
		oldPendingAdminRule = append(oldPendingAdminRule, oldPendingAdminItem)
	}
	var newPendingAdminRule []interface{}
	for _, newPendingAdminItem := range newPendingAdmin {
		newPendingAdminRule = append(newPendingAdminRule, newPendingAdminItem)
	}

	logs, sub, err := _IZkSync.contract.WatchLogs(opts, "NewPendingAdmin", oldPendingAdminRule, newPendingAdminRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IZkSyncNewPendingAdmin)
				if err := _IZkSync.contract.UnpackLog(event, "NewPendingAdmin", log); err != nil {
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
func (_IZkSync *IZkSyncFilterer) ParseNewPendingAdmin(log types.Log) (*IZkSyncNewPendingAdmin, error) {
	event := new(IZkSyncNewPendingAdmin)
	if err := _IZkSync.contract.UnpackLog(event, "NewPendingAdmin", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// IZkSyncNewPendingGovernorIterator is returned from FilterNewPendingGovernor and is used to iterate over the raw logs and unpacked data for NewPendingGovernor events raised by the IZkSync contract.
type IZkSyncNewPendingGovernorIterator struct {
	Event *IZkSyncNewPendingGovernor // Event containing the contract specifics and raw log

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
func (it *IZkSyncNewPendingGovernorIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IZkSyncNewPendingGovernor)
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
		it.Event = new(IZkSyncNewPendingGovernor)
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
func (it *IZkSyncNewPendingGovernorIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IZkSyncNewPendingGovernorIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IZkSyncNewPendingGovernor represents a NewPendingGovernor event raised by the IZkSync contract.
type IZkSyncNewPendingGovernor struct {
	OldPendingGovernor common.Address
	NewPendingGovernor common.Address
	Raw                types.Log // Blockchain specific contextual infos
}

// FilterNewPendingGovernor is a free log retrieval operation binding the contract event 0x7d767be5a57784412a13945bd5114db84487d2b007bfcdb2f449fc9ea35437f7.
//
// Solidity: event NewPendingGovernor(address indexed oldPendingGovernor, address indexed newPendingGovernor)
func (_IZkSync *IZkSyncFilterer) FilterNewPendingGovernor(opts *bind.FilterOpts, oldPendingGovernor []common.Address, newPendingGovernor []common.Address) (*IZkSyncNewPendingGovernorIterator, error) {

	var oldPendingGovernorRule []interface{}
	for _, oldPendingGovernorItem := range oldPendingGovernor {
		oldPendingGovernorRule = append(oldPendingGovernorRule, oldPendingGovernorItem)
	}
	var newPendingGovernorRule []interface{}
	for _, newPendingGovernorItem := range newPendingGovernor {
		newPendingGovernorRule = append(newPendingGovernorRule, newPendingGovernorItem)
	}

	logs, sub, err := _IZkSync.contract.FilterLogs(opts, "NewPendingGovernor", oldPendingGovernorRule, newPendingGovernorRule)
	if err != nil {
		return nil, err
	}
	return &IZkSyncNewPendingGovernorIterator{contract: _IZkSync.contract, event: "NewPendingGovernor", logs: logs, sub: sub}, nil
}

// WatchNewPendingGovernor is a free log subscription operation binding the contract event 0x7d767be5a57784412a13945bd5114db84487d2b007bfcdb2f449fc9ea35437f7.
//
// Solidity: event NewPendingGovernor(address indexed oldPendingGovernor, address indexed newPendingGovernor)
func (_IZkSync *IZkSyncFilterer) WatchNewPendingGovernor(opts *bind.WatchOpts, sink chan<- *IZkSyncNewPendingGovernor, oldPendingGovernor []common.Address, newPendingGovernor []common.Address) (event.Subscription, error) {

	var oldPendingGovernorRule []interface{}
	for _, oldPendingGovernorItem := range oldPendingGovernor {
		oldPendingGovernorRule = append(oldPendingGovernorRule, oldPendingGovernorItem)
	}
	var newPendingGovernorRule []interface{}
	for _, newPendingGovernorItem := range newPendingGovernor {
		newPendingGovernorRule = append(newPendingGovernorRule, newPendingGovernorItem)
	}

	logs, sub, err := _IZkSync.contract.WatchLogs(opts, "NewPendingGovernor", oldPendingGovernorRule, newPendingGovernorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IZkSyncNewPendingGovernor)
				if err := _IZkSync.contract.UnpackLog(event, "NewPendingGovernor", log); err != nil {
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

// ParseNewPendingGovernor is a log parse operation binding the contract event 0x7d767be5a57784412a13945bd5114db84487d2b007bfcdb2f449fc9ea35437f7.
//
// Solidity: event NewPendingGovernor(address indexed oldPendingGovernor, address indexed newPendingGovernor)
func (_IZkSync *IZkSyncFilterer) ParseNewPendingGovernor(log types.Log) (*IZkSyncNewPendingGovernor, error) {
	event := new(IZkSyncNewPendingGovernor)
	if err := _IZkSync.contract.UnpackLog(event, "NewPendingGovernor", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// IZkSyncNewPriorityRequestIterator is returned from FilterNewPriorityRequest and is used to iterate over the raw logs and unpacked data for NewPriorityRequest events raised by the IZkSync contract.
type IZkSyncNewPriorityRequestIterator struct {
	Event *IZkSyncNewPriorityRequest // Event containing the contract specifics and raw log

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
func (it *IZkSyncNewPriorityRequestIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IZkSyncNewPriorityRequest)
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
		it.Event = new(IZkSyncNewPriorityRequest)
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
func (it *IZkSyncNewPriorityRequestIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IZkSyncNewPriorityRequestIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IZkSyncNewPriorityRequest represents a NewPriorityRequest event raised by the IZkSync contract.
type IZkSyncNewPriorityRequest struct {
	TxId                *big.Int
	TxHash              [32]byte
	ExpirationTimestamp uint64
	Transaction         IMailboxL2CanonicalTransaction
	FactoryDeps         [][]byte
	Raw                 types.Log // Blockchain specific contextual infos
}

// FilterNewPriorityRequest is a free log retrieval operation binding the contract event 0x4531cd5795773d7101c17bdeb9f5ab7f47d7056017506f937083be5d6e77a382.
//
// Solidity: event NewPriorityRequest(uint256 txId, bytes32 txHash, uint64 expirationTimestamp, (uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256[4],bytes,bytes,uint256[],bytes,bytes) transaction, bytes[] factoryDeps)
func (_IZkSync *IZkSyncFilterer) FilterNewPriorityRequest(opts *bind.FilterOpts) (*IZkSyncNewPriorityRequestIterator, error) {

	logs, sub, err := _IZkSync.contract.FilterLogs(opts, "NewPriorityRequest")
	if err != nil {
		return nil, err
	}
	return &IZkSyncNewPriorityRequestIterator{contract: _IZkSync.contract, event: "NewPriorityRequest", logs: logs, sub: sub}, nil
}

// WatchNewPriorityRequest is a free log subscription operation binding the contract event 0x4531cd5795773d7101c17bdeb9f5ab7f47d7056017506f937083be5d6e77a382.
//
// Solidity: event NewPriorityRequest(uint256 txId, bytes32 txHash, uint64 expirationTimestamp, (uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256[4],bytes,bytes,uint256[],bytes,bytes) transaction, bytes[] factoryDeps)
func (_IZkSync *IZkSyncFilterer) WatchNewPriorityRequest(opts *bind.WatchOpts, sink chan<- *IZkSyncNewPriorityRequest) (event.Subscription, error) {

	logs, sub, err := _IZkSync.contract.WatchLogs(opts, "NewPriorityRequest")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IZkSyncNewPriorityRequest)
				if err := _IZkSync.contract.UnpackLog(event, "NewPriorityRequest", log); err != nil {
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
func (_IZkSync *IZkSyncFilterer) ParseNewPriorityRequest(log types.Log) (*IZkSyncNewPriorityRequest, error) {
	event := new(IZkSyncNewPriorityRequest)
	if err := _IZkSync.contract.UnpackLog(event, "NewPriorityRequest", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// IZkSyncNewPriorityTxMaxGasLimitIterator is returned from FilterNewPriorityTxMaxGasLimit and is used to iterate over the raw logs and unpacked data for NewPriorityTxMaxGasLimit events raised by the IZkSync contract.
type IZkSyncNewPriorityTxMaxGasLimitIterator struct {
	Event *IZkSyncNewPriorityTxMaxGasLimit // Event containing the contract specifics and raw log

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
func (it *IZkSyncNewPriorityTxMaxGasLimitIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IZkSyncNewPriorityTxMaxGasLimit)
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
		it.Event = new(IZkSyncNewPriorityTxMaxGasLimit)
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
func (it *IZkSyncNewPriorityTxMaxGasLimitIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IZkSyncNewPriorityTxMaxGasLimitIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IZkSyncNewPriorityTxMaxGasLimit represents a NewPriorityTxMaxGasLimit event raised by the IZkSync contract.
type IZkSyncNewPriorityTxMaxGasLimit struct {
	OldPriorityTxMaxGasLimit *big.Int
	NewPriorityTxMaxGasLimit *big.Int
	Raw                      types.Log // Blockchain specific contextual infos
}

// FilterNewPriorityTxMaxGasLimit is a free log retrieval operation binding the contract event 0x83dd728f7e76a849126c55ffabdc6e299eb8c85dccf12498701376d9f5c954d1.
//
// Solidity: event NewPriorityTxMaxGasLimit(uint256 oldPriorityTxMaxGasLimit, uint256 newPriorityTxMaxGasLimit)
func (_IZkSync *IZkSyncFilterer) FilterNewPriorityTxMaxGasLimit(opts *bind.FilterOpts) (*IZkSyncNewPriorityTxMaxGasLimitIterator, error) {

	logs, sub, err := _IZkSync.contract.FilterLogs(opts, "NewPriorityTxMaxGasLimit")
	if err != nil {
		return nil, err
	}
	return &IZkSyncNewPriorityTxMaxGasLimitIterator{contract: _IZkSync.contract, event: "NewPriorityTxMaxGasLimit", logs: logs, sub: sub}, nil
}

// WatchNewPriorityTxMaxGasLimit is a free log subscription operation binding the contract event 0x83dd728f7e76a849126c55ffabdc6e299eb8c85dccf12498701376d9f5c954d1.
//
// Solidity: event NewPriorityTxMaxGasLimit(uint256 oldPriorityTxMaxGasLimit, uint256 newPriorityTxMaxGasLimit)
func (_IZkSync *IZkSyncFilterer) WatchNewPriorityTxMaxGasLimit(opts *bind.WatchOpts, sink chan<- *IZkSyncNewPriorityTxMaxGasLimit) (event.Subscription, error) {

	logs, sub, err := _IZkSync.contract.WatchLogs(opts, "NewPriorityTxMaxGasLimit")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IZkSyncNewPriorityTxMaxGasLimit)
				if err := _IZkSync.contract.UnpackLog(event, "NewPriorityTxMaxGasLimit", log); err != nil {
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
func (_IZkSync *IZkSyncFilterer) ParseNewPriorityTxMaxGasLimit(log types.Log) (*IZkSyncNewPriorityTxMaxGasLimit, error) {
	event := new(IZkSyncNewPriorityTxMaxGasLimit)
	if err := _IZkSync.contract.UnpackLog(event, "NewPriorityTxMaxGasLimit", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// IZkSyncUnfreezeIterator is returned from FilterUnfreeze and is used to iterate over the raw logs and unpacked data for Unfreeze events raised by the IZkSync contract.
type IZkSyncUnfreezeIterator struct {
	Event *IZkSyncUnfreeze // Event containing the contract specifics and raw log

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
func (it *IZkSyncUnfreezeIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IZkSyncUnfreeze)
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
		it.Event = new(IZkSyncUnfreeze)
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
func (it *IZkSyncUnfreezeIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IZkSyncUnfreezeIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IZkSyncUnfreeze represents a Unfreeze event raised by the IZkSync contract.
type IZkSyncUnfreeze struct {
	Raw types.Log // Blockchain specific contextual infos
}

// FilterUnfreeze is a free log retrieval operation binding the contract event 0x2f05ba71d0df11bf5fa562a6569d70c4f80da84284badbe015ce1456063d0ded.
//
// Solidity: event Unfreeze()
func (_IZkSync *IZkSyncFilterer) FilterUnfreeze(opts *bind.FilterOpts) (*IZkSyncUnfreezeIterator, error) {

	logs, sub, err := _IZkSync.contract.FilterLogs(opts, "Unfreeze")
	if err != nil {
		return nil, err
	}
	return &IZkSyncUnfreezeIterator{contract: _IZkSync.contract, event: "Unfreeze", logs: logs, sub: sub}, nil
}

// WatchUnfreeze is a free log subscription operation binding the contract event 0x2f05ba71d0df11bf5fa562a6569d70c4f80da84284badbe015ce1456063d0ded.
//
// Solidity: event Unfreeze()
func (_IZkSync *IZkSyncFilterer) WatchUnfreeze(opts *bind.WatchOpts, sink chan<- *IZkSyncUnfreeze) (event.Subscription, error) {

	logs, sub, err := _IZkSync.contract.WatchLogs(opts, "Unfreeze")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IZkSyncUnfreeze)
				if err := _IZkSync.contract.UnpackLog(event, "Unfreeze", log); err != nil {
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
func (_IZkSync *IZkSyncFilterer) ParseUnfreeze(log types.Log) (*IZkSyncUnfreeze, error) {
	event := new(IZkSyncUnfreeze)
	if err := _IZkSync.contract.UnpackLog(event, "Unfreeze", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// IZkSyncValidatorStatusUpdateIterator is returned from FilterValidatorStatusUpdate and is used to iterate over the raw logs and unpacked data for ValidatorStatusUpdate events raised by the IZkSync contract.
type IZkSyncValidatorStatusUpdateIterator struct {
	Event *IZkSyncValidatorStatusUpdate // Event containing the contract specifics and raw log

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
func (it *IZkSyncValidatorStatusUpdateIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IZkSyncValidatorStatusUpdate)
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
		it.Event = new(IZkSyncValidatorStatusUpdate)
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
func (it *IZkSyncValidatorStatusUpdateIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IZkSyncValidatorStatusUpdateIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IZkSyncValidatorStatusUpdate represents a ValidatorStatusUpdate event raised by the IZkSync contract.
type IZkSyncValidatorStatusUpdate struct {
	ValidatorAddress common.Address
	IsActive         bool
	Raw              types.Log // Blockchain specific contextual infos
}

// FilterValidatorStatusUpdate is a free log retrieval operation binding the contract event 0x065b77b53864e46fda3d8986acb51696223d6dde7ced42441eb150bae6d48136.
//
// Solidity: event ValidatorStatusUpdate(address indexed validatorAddress, bool isActive)
func (_IZkSync *IZkSyncFilterer) FilterValidatorStatusUpdate(opts *bind.FilterOpts, validatorAddress []common.Address) (*IZkSyncValidatorStatusUpdateIterator, error) {

	var validatorAddressRule []interface{}
	for _, validatorAddressItem := range validatorAddress {
		validatorAddressRule = append(validatorAddressRule, validatorAddressItem)
	}

	logs, sub, err := _IZkSync.contract.FilterLogs(opts, "ValidatorStatusUpdate", validatorAddressRule)
	if err != nil {
		return nil, err
	}
	return &IZkSyncValidatorStatusUpdateIterator{contract: _IZkSync.contract, event: "ValidatorStatusUpdate", logs: logs, sub: sub}, nil
}

// WatchValidatorStatusUpdate is a free log subscription operation binding the contract event 0x065b77b53864e46fda3d8986acb51696223d6dde7ced42441eb150bae6d48136.
//
// Solidity: event ValidatorStatusUpdate(address indexed validatorAddress, bool isActive)
func (_IZkSync *IZkSyncFilterer) WatchValidatorStatusUpdate(opts *bind.WatchOpts, sink chan<- *IZkSyncValidatorStatusUpdate, validatorAddress []common.Address) (event.Subscription, error) {

	var validatorAddressRule []interface{}
	for _, validatorAddressItem := range validatorAddress {
		validatorAddressRule = append(validatorAddressRule, validatorAddressItem)
	}

	logs, sub, err := _IZkSync.contract.WatchLogs(opts, "ValidatorStatusUpdate", validatorAddressRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IZkSyncValidatorStatusUpdate)
				if err := _IZkSync.contract.UnpackLog(event, "ValidatorStatusUpdate", log); err != nil {
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
func (_IZkSync *IZkSyncFilterer) ParseValidatorStatusUpdate(log types.Log) (*IZkSyncValidatorStatusUpdate, error) {
	event := new(IZkSyncValidatorStatusUpdate)
	if err := _IZkSync.contract.UnpackLog(event, "ValidatorStatusUpdate", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
