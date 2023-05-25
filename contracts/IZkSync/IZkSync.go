// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package IZkSync

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

// IExecutorCommitBlockInfo is an auto generated low-level Go binding around an user-defined struct.
type IExecutorCommitBlockInfo struct {
	BlockNumber                 uint64
	Timestamp                   uint64
	IndexRepeatedStorageChanges uint64
	NewStateRoot                [32]byte
	NumberOfLayer1Txs           *big.Int
	L2LogsTreeRoot              [32]byte
	PriorityOperationsHash      [32]byte
	InitialStorageChanges       []byte
	RepeatedStorageChanges      []byte
	L2Logs                      []byte
	L2ArbitraryLengthMessages   [][]byte
	FactoryDeps                 [][]byte
}

// IExecutorProofInput is an auto generated low-level Go binding around an user-defined struct.
type IExecutorProofInput struct {
	RecursiveAggregationInput []*big.Int
	SerializedProof           []*big.Int
}

// IExecutorStoredBlockInfo is an auto generated low-level Go binding around an user-defined struct.
type IExecutorStoredBlockInfo struct {
	BlockNumber                 uint64
	BlockHash                   [32]byte
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
	TxNumberInBlock uint16
	Sender          common.Address
	Key             [32]byte
	Value           [32]byte
}

// L2Message is an auto generated low-level Go binding around an user-defined struct.
type L2Message struct {
	TxNumberInBlock uint16
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
	ABI: "[{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"blockNumber\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"blockHash\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"commitment\",\"type\":\"bytes32\"}],\"name\":\"BlockCommit\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"blockNumber\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"blockHash\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"commitment\",\"type\":\"bytes32\"}],\"name\":\"BlockExecution\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"totalBlocksCommitted\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"totalBlocksVerified\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"totalBlocksExecuted\",\"type\":\"uint256\"}],\"name\":\"BlocksRevert\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"previousLastVerifiedBlock\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"currentLastVerifiedBlock\",\"type\":\"uint256\"}],\"name\":\"BlocksVerification\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"proposalId\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"proposalHash\",\"type\":\"bytes32\"}],\"name\":\"CancelUpgradeProposal\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"EthWithdrawalFinalized\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"proposalId\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"proposalHash\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"proposalSalt\",\"type\":\"bytes32\"}],\"name\":\"ExecuteUpgrade\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[],\"name\":\"Freeze\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"isPorterAvailable\",\"type\":\"bool\"}],\"name\":\"IsPorterAvailableStatusUpdate\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"oldAllowList\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newAllowList\",\"type\":\"address\"}],\"name\":\"NewAllowList\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"oldGovernor\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newGovernor\",\"type\":\"address\"}],\"name\":\"NewGovernor\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"previousBytecodeHash\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"newBytecodeHash\",\"type\":\"bytes32\"}],\"name\":\"NewL2BootloaderBytecodeHash\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"previousBytecodeHash\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"newBytecodeHash\",\"type\":\"bytes32\"}],\"name\":\"NewL2DefaultAccountBytecodeHash\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"oldPendingGovernor\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newPendingGovernor\",\"type\":\"address\"}],\"name\":\"NewPendingGovernor\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"txId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"txHash\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"expirationTimestamp\",\"type\":\"uint64\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"txType\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"from\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"to\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"gasLimit\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"gasPerPubdataByteLimit\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"maxFeePerGas\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"maxPriorityFeePerGas\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"paymaster\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"nonce\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"},{\"internalType\":\"uint256[4]\",\"name\":\"reserved\",\"type\":\"uint256[4]\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"signature\",\"type\":\"bytes\"},{\"internalType\":\"uint256[]\",\"name\":\"factoryDeps\",\"type\":\"uint256[]\"},{\"internalType\":\"bytes\",\"name\":\"paymasterInput\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"reservedDynamic\",\"type\":\"bytes\"}],\"indexed\":false,\"internalType\":\"structIMailbox.L2CanonicalTransaction\",\"name\":\"transaction\",\"type\":\"tuple\"},{\"indexed\":false,\"internalType\":\"bytes[]\",\"name\":\"factoryDeps\",\"type\":\"bytes[]\"}],\"name\":\"NewPriorityRequest\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"oldPriorityTxMaxGasLimit\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"newPriorityTxMaxGasLimit\",\"type\":\"uint256\"}],\"name\":\"NewPriorityTxMaxGasLimit\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"oldVerifier\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newVerifier\",\"type\":\"address\"}],\"name\":\"NewVerifier\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"recursionNodeLevelVkHash\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"recursionLeafLevelVkHash\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"recursionCircuitsSetVksHash\",\"type\":\"bytes32\"}],\"indexed\":false,\"internalType\":\"structVerifierParams\",\"name\":\"oldVerifierParams\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"recursionNodeLevelVkHash\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"recursionLeafLevelVkHash\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"recursionCircuitsSetVksHash\",\"type\":\"bytes32\"}],\"indexed\":false,\"internalType\":\"structVerifierParams\",\"name\":\"newVerifierParams\",\"type\":\"tuple\"}],\"name\":\"NewVerifierParams\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"proposalId\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"proposalHash\",\"type\":\"bytes32\"}],\"name\":\"ProposeShadowUpgrade\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"components\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"facet\",\"type\":\"address\"},{\"internalType\":\"enumDiamond.Action\",\"name\":\"action\",\"type\":\"uint8\"},{\"internalType\":\"bool\",\"name\":\"isFreezable\",\"type\":\"bool\"},{\"internalType\":\"bytes4[]\",\"name\":\"selectors\",\"type\":\"bytes4[]\"}],\"internalType\":\"structDiamond.FacetCut[]\",\"name\":\"facetCuts\",\"type\":\"tuple[]\"},{\"internalType\":\"address\",\"name\":\"initAddress\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"initCalldata\",\"type\":\"bytes\"}],\"indexed\":false,\"internalType\":\"structDiamond.DiamondCutData\",\"name\":\"diamondCut\",\"type\":\"tuple\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"proposalId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"proposalSalt\",\"type\":\"bytes32\"}],\"name\":\"ProposeTransparentUpgrade\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"proposalId\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"proposalHash\",\"type\":\"bytes32\"}],\"name\":\"SecurityCouncilUpgradeApprove\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[],\"name\":\"Unfreeze\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"validatorAddress\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"isActive\",\"type\":\"bool\"}],\"name\":\"ValidatorStatusUpdate\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"acceptGovernor\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_proposedUpgradeHash\",\"type\":\"bytes32\"}],\"name\":\"cancelUpgradeProposal\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint64\",\"name\":\"blockNumber\",\"type\":\"uint64\"},{\"internalType\":\"bytes32\",\"name\":\"blockHash\",\"type\":\"bytes32\"},{\"internalType\":\"uint64\",\"name\":\"indexRepeatedStorageChanges\",\"type\":\"uint64\"},{\"internalType\":\"uint256\",\"name\":\"numberOfLayer1Txs\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"priorityOperationsHash\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"l2LogsTreeRoot\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"timestamp\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"commitment\",\"type\":\"bytes32\"}],\"internalType\":\"structIExecutor.StoredBlockInfo\",\"name\":\"_lastCommittedBlockData\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"uint64\",\"name\":\"blockNumber\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"timestamp\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"indexRepeatedStorageChanges\",\"type\":\"uint64\"},{\"internalType\":\"bytes32\",\"name\":\"newStateRoot\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"numberOfLayer1Txs\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"l2LogsTreeRoot\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"priorityOperationsHash\",\"type\":\"bytes32\"},{\"internalType\":\"bytes\",\"name\":\"initialStorageChanges\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"repeatedStorageChanges\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"l2Logs\",\"type\":\"bytes\"},{\"internalType\":\"bytes[]\",\"name\":\"l2ArbitraryLengthMessages\",\"type\":\"bytes[]\"},{\"internalType\":\"bytes[]\",\"name\":\"factoryDeps\",\"type\":\"bytes[]\"}],\"internalType\":\"structIExecutor.CommitBlockInfo[]\",\"name\":\"_newBlocksData\",\"type\":\"tuple[]\"}],\"name\":\"commitBlocks\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint64\",\"name\":\"blockNumber\",\"type\":\"uint64\"},{\"internalType\":\"bytes32\",\"name\":\"blockHash\",\"type\":\"bytes32\"},{\"internalType\":\"uint64\",\"name\":\"indexRepeatedStorageChanges\",\"type\":\"uint64\"},{\"internalType\":\"uint256\",\"name\":\"numberOfLayer1Txs\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"priorityOperationsHash\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"l2LogsTreeRoot\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"timestamp\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"commitment\",\"type\":\"bytes32\"}],\"internalType\":\"structIExecutor.StoredBlockInfo[]\",\"name\":\"_blocksData\",\"type\":\"tuple[]\"}],\"name\":\"executeBlocks\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"facet\",\"type\":\"address\"},{\"internalType\":\"enumDiamond.Action\",\"name\":\"action\",\"type\":\"uint8\"},{\"internalType\":\"bool\",\"name\":\"isFreezable\",\"type\":\"bool\"},{\"internalType\":\"bytes4[]\",\"name\":\"selectors\",\"type\":\"bytes4[]\"}],\"internalType\":\"structDiamond.FacetCut[]\",\"name\":\"facetCuts\",\"type\":\"tuple[]\"},{\"internalType\":\"address\",\"name\":\"initAddress\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"initCalldata\",\"type\":\"bytes\"}],\"internalType\":\"structDiamond.DiamondCutData\",\"name\":\"_diamondCut\",\"type\":\"tuple\"},{\"internalType\":\"bytes32\",\"name\":\"_proposalSalt\",\"type\":\"bytes32\"}],\"name\":\"executeUpgrade\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes4\",\"name\":\"_selector\",\"type\":\"bytes4\"}],\"name\":\"facetAddress\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"facet\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"facetAddresses\",\"outputs\":[{\"internalType\":\"address[]\",\"name\":\"facets\",\"type\":\"address[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_facet\",\"type\":\"address\"}],\"name\":\"facetFunctionSelectors\",\"outputs\":[{\"internalType\":\"bytes4[]\",\"name\":\"\",\"type\":\"bytes4[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"facets\",\"outputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"addr\",\"type\":\"address\"},{\"internalType\":\"bytes4[]\",\"name\":\"selectors\",\"type\":\"bytes4[]\"}],\"internalType\":\"structIGetters.Facet[]\",\"name\":\"\",\"type\":\"tuple[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_l2BlockNumber\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_l2MessageIndex\",\"type\":\"uint256\"},{\"internalType\":\"uint16\",\"name\":\"_l2TxNumberInBlock\",\"type\":\"uint16\"},{\"internalType\":\"bytes\",\"name\":\"_message\",\"type\":\"bytes\"},{\"internalType\":\"bytes32[]\",\"name\":\"_merkleProof\",\"type\":\"bytes32[]\"}],\"name\":\"finalizeEthWithdrawal\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"freezeDiamond\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getAllowList\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getCurrentProposalId\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getFirstUnprocessedPriorityTx\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getGovernor\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getL2BootloaderBytecodeHash\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getL2DefaultAccountBytecodeHash\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getPendingGovernor\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getPriorityQueueSize\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getPriorityTxMaxGasLimit\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getProposedUpgradeHash\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getProposedUpgradeTimestamp\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getSecurityCouncil\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getTotalBlocksCommitted\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getTotalBlocksExecuted\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getTotalBlocksVerified\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getTotalPriorityTxs\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getUpgradeProposalState\",\"outputs\":[{\"internalType\":\"enumUpgradeState\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getVerifier\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getVerifierParams\",\"outputs\":[{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"recursionNodeLevelVkHash\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"recursionLeafLevelVkHash\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"recursionCircuitsSetVksHash\",\"type\":\"bytes32\"}],\"internalType\":\"structVerifierParams\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"isApprovedBySecurityCouncil\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"isDiamondStorageFrozen\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_l2BlockNumber\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_l2MessageIndex\",\"type\":\"uint256\"}],\"name\":\"isEthWithdrawalFinalized\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_facet\",\"type\":\"address\"}],\"name\":\"isFacetFreezable\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"isFreezable\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes4\",\"name\":\"_selector\",\"type\":\"bytes4\"}],\"name\":\"isFunctionFreezable\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_address\",\"type\":\"address\"}],\"name\":\"isValidator\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_blockNumber\",\"type\":\"uint256\"}],\"name\":\"l2LogsRootHash\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"hash\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_gasPrice\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_l2GasLimit\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_l2GasPerPubdataByteLimit\",\"type\":\"uint256\"}],\"name\":\"l2TransactionBaseCost\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"priorityQueueFrontOperation\",\"outputs\":[{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"canonicalTxHash\",\"type\":\"bytes32\"},{\"internalType\":\"uint64\",\"name\":\"expirationTimestamp\",\"type\":\"uint64\"},{\"internalType\":\"uint192\",\"name\":\"layer2Tip\",\"type\":\"uint192\"}],\"internalType\":\"structPriorityOperation\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_proposalHash\",\"type\":\"bytes32\"},{\"internalType\":\"uint40\",\"name\":\"_proposalId\",\"type\":\"uint40\"}],\"name\":\"proposeShadowUpgrade\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"facet\",\"type\":\"address\"},{\"internalType\":\"enumDiamond.Action\",\"name\":\"action\",\"type\":\"uint8\"},{\"internalType\":\"bool\",\"name\":\"isFreezable\",\"type\":\"bool\"},{\"internalType\":\"bytes4[]\",\"name\":\"selectors\",\"type\":\"bytes4[]\"}],\"internalType\":\"structDiamond.FacetCut[]\",\"name\":\"facetCuts\",\"type\":\"tuple[]\"},{\"internalType\":\"address\",\"name\":\"initAddress\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"initCalldata\",\"type\":\"bytes\"}],\"internalType\":\"structDiamond.DiamondCutData\",\"name\":\"_diamondCut\",\"type\":\"tuple\"},{\"internalType\":\"uint40\",\"name\":\"_proposalId\",\"type\":\"uint40\"}],\"name\":\"proposeTransparentUpgrade\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint64\",\"name\":\"blockNumber\",\"type\":\"uint64\"},{\"internalType\":\"bytes32\",\"name\":\"blockHash\",\"type\":\"bytes32\"},{\"internalType\":\"uint64\",\"name\":\"indexRepeatedStorageChanges\",\"type\":\"uint64\"},{\"internalType\":\"uint256\",\"name\":\"numberOfLayer1Txs\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"priorityOperationsHash\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"l2LogsTreeRoot\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"timestamp\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"commitment\",\"type\":\"bytes32\"}],\"internalType\":\"structIExecutor.StoredBlockInfo\",\"name\":\"_prevBlock\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"uint64\",\"name\":\"blockNumber\",\"type\":\"uint64\"},{\"internalType\":\"bytes32\",\"name\":\"blockHash\",\"type\":\"bytes32\"},{\"internalType\":\"uint64\",\"name\":\"indexRepeatedStorageChanges\",\"type\":\"uint64\"},{\"internalType\":\"uint256\",\"name\":\"numberOfLayer1Txs\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"priorityOperationsHash\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"l2LogsTreeRoot\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"timestamp\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"commitment\",\"type\":\"bytes32\"}],\"internalType\":\"structIExecutor.StoredBlockInfo[]\",\"name\":\"_committedBlocks\",\"type\":\"tuple[]\"},{\"components\":[{\"internalType\":\"uint256[]\",\"name\":\"recursiveAggregationInput\",\"type\":\"uint256[]\"},{\"internalType\":\"uint256[]\",\"name\":\"serializedProof\",\"type\":\"uint256[]\"}],\"internalType\":\"structIExecutor.ProofInput\",\"name\":\"_proof\",\"type\":\"tuple\"}],\"name\":\"proveBlocks\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_l2TxHash\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"_l2BlockNumber\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_l2MessageIndex\",\"type\":\"uint256\"},{\"internalType\":\"uint16\",\"name\":\"_l2TxNumberInBlock\",\"type\":\"uint16\"},{\"internalType\":\"bytes32[]\",\"name\":\"_merkleProof\",\"type\":\"bytes32[]\"},{\"internalType\":\"enumTxStatus\",\"name\":\"_status\",\"type\":\"uint8\"}],\"name\":\"proveL1ToL2TransactionStatus\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_blockNumber\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_index\",\"type\":\"uint256\"},{\"components\":[{\"internalType\":\"uint8\",\"name\":\"l2ShardId\",\"type\":\"uint8\"},{\"internalType\":\"bool\",\"name\":\"isService\",\"type\":\"bool\"},{\"internalType\":\"uint16\",\"name\":\"txNumberInBlock\",\"type\":\"uint16\"},{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"key\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"value\",\"type\":\"bytes32\"}],\"internalType\":\"structL2Log\",\"name\":\"_log\",\"type\":\"tuple\"},{\"internalType\":\"bytes32[]\",\"name\":\"_proof\",\"type\":\"bytes32[]\"}],\"name\":\"proveL2LogInclusion\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_blockNumber\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_index\",\"type\":\"uint256\"},{\"components\":[{\"internalType\":\"uint16\",\"name\":\"txNumberInBlock\",\"type\":\"uint16\"},{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"internalType\":\"structL2Message\",\"name\":\"_message\",\"type\":\"tuple\"},{\"internalType\":\"bytes32[]\",\"name\":\"_proof\",\"type\":\"bytes32[]\"}],\"name\":\"proveL2MessageInclusion\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_contractL2\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_l2Value\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"_calldata\",\"type\":\"bytes\"},{\"internalType\":\"uint256\",\"name\":\"_l2GasLimit\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_l2GasPerPubdataByteLimit\",\"type\":\"uint256\"},{\"internalType\":\"bytes[]\",\"name\":\"_factoryDeps\",\"type\":\"bytes[]\"},{\"internalType\":\"address\",\"name\":\"_refundRecipient\",\"type\":\"address\"}],\"name\":\"requestL2Transaction\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"canonicalTxHash\",\"type\":\"bytes32\"}],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_newLastBlock\",\"type\":\"uint256\"}],\"name\":\"revertBlocks\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_upgradeProposalHash\",\"type\":\"bytes32\"}],\"name\":\"securityCouncilUpgradeApprove\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractIAllowList\",\"name\":\"_newAllowList\",\"type\":\"address\"}],\"name\":\"setAllowList\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_l2BootloaderBytecodeHash\",\"type\":\"bytes32\"}],\"name\":\"setL2BootloaderBytecodeHash\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_l2DefaultAccountBytecodeHash\",\"type\":\"bytes32\"}],\"name\":\"setL2DefaultAccountBytecodeHash\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_newPendingGovernor\",\"type\":\"address\"}],\"name\":\"setPendingGovernor\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bool\",\"name\":\"_zkPorterIsAvailable\",\"type\":\"bool\"}],\"name\":\"setPorterAvailability\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_newPriorityTxMaxGasLimit\",\"type\":\"uint256\"}],\"name\":\"setPriorityTxMaxGasLimit\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_validator\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"_active\",\"type\":\"bool\"}],\"name\":\"setValidator\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractVerifier\",\"name\":\"_newVerifier\",\"type\":\"address\"}],\"name\":\"setVerifier\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"recursionNodeLevelVkHash\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"recursionLeafLevelVkHash\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"recursionCircuitsSetVksHash\",\"type\":\"bytes32\"}],\"internalType\":\"structVerifierParams\",\"name\":\"_newVerifierParams\",\"type\":\"tuple\"}],\"name\":\"setVerifierParams\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_blockNumber\",\"type\":\"uint256\"}],\"name\":\"storedBlockHash\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"unfreezeDiamond\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"facet\",\"type\":\"address\"},{\"internalType\":\"enumDiamond.Action\",\"name\":\"action\",\"type\":\"uint8\"},{\"internalType\":\"bool\",\"name\":\"isFreezable\",\"type\":\"bool\"},{\"internalType\":\"bytes4[]\",\"name\":\"selectors\",\"type\":\"bytes4[]\"}],\"internalType\":\"structDiamond.FacetCut[]\",\"name\":\"facetCuts\",\"type\":\"tuple[]\"},{\"internalType\":\"address\",\"name\":\"initAddress\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"initCalldata\",\"type\":\"bytes\"}],\"internalType\":\"structDiamond.DiamondCutData\",\"name\":\"_diamondCut\",\"type\":\"tuple\"},{\"internalType\":\"uint256\",\"name\":\"_proposalId\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"_salt\",\"type\":\"bytes32\"}],\"name\":\"upgradeProposalHash\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"pure\",\"type\":\"function\"}]",
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

// GetAllowList is a free data retrieval call binding the contract method 0xa7cd63b7.
//
// Solidity: function getAllowList() view returns(address)
func (_IZkSync *IZkSyncCaller) GetAllowList(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _IZkSync.contract.Call(opts, &out, "getAllowList")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetAllowList is a free data retrieval call binding the contract method 0xa7cd63b7.
//
// Solidity: function getAllowList() view returns(address)
func (_IZkSync *IZkSyncSession) GetAllowList() (common.Address, error) {
	return _IZkSync.Contract.GetAllowList(&_IZkSync.CallOpts)
}

// GetAllowList is a free data retrieval call binding the contract method 0xa7cd63b7.
//
// Solidity: function getAllowList() view returns(address)
func (_IZkSync *IZkSyncCallerSession) GetAllowList() (common.Address, error) {
	return _IZkSync.Contract.GetAllowList(&_IZkSync.CallOpts)
}

// GetCurrentProposalId is a free data retrieval call binding the contract method 0xfe10226d.
//
// Solidity: function getCurrentProposalId() view returns(uint256)
func (_IZkSync *IZkSyncCaller) GetCurrentProposalId(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _IZkSync.contract.Call(opts, &out, "getCurrentProposalId")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetCurrentProposalId is a free data retrieval call binding the contract method 0xfe10226d.
//
// Solidity: function getCurrentProposalId() view returns(uint256)
func (_IZkSync *IZkSyncSession) GetCurrentProposalId() (*big.Int, error) {
	return _IZkSync.Contract.GetCurrentProposalId(&_IZkSync.CallOpts)
}

// GetCurrentProposalId is a free data retrieval call binding the contract method 0xfe10226d.
//
// Solidity: function getCurrentProposalId() view returns(uint256)
func (_IZkSync *IZkSyncCallerSession) GetCurrentProposalId() (*big.Int, error) {
	return _IZkSync.Contract.GetCurrentProposalId(&_IZkSync.CallOpts)
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

// GetProposedUpgradeHash is a free data retrieval call binding the contract method 0x1b60e626.
//
// Solidity: function getProposedUpgradeHash() view returns(bytes32)
func (_IZkSync *IZkSyncCaller) GetProposedUpgradeHash(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _IZkSync.contract.Call(opts, &out, "getProposedUpgradeHash")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// GetProposedUpgradeHash is a free data retrieval call binding the contract method 0x1b60e626.
//
// Solidity: function getProposedUpgradeHash() view returns(bytes32)
func (_IZkSync *IZkSyncSession) GetProposedUpgradeHash() ([32]byte, error) {
	return _IZkSync.Contract.GetProposedUpgradeHash(&_IZkSync.CallOpts)
}

// GetProposedUpgradeHash is a free data retrieval call binding the contract method 0x1b60e626.
//
// Solidity: function getProposedUpgradeHash() view returns(bytes32)
func (_IZkSync *IZkSyncCallerSession) GetProposedUpgradeHash() ([32]byte, error) {
	return _IZkSync.Contract.GetProposedUpgradeHash(&_IZkSync.CallOpts)
}

// GetProposedUpgradeTimestamp is a free data retrieval call binding the contract method 0xe39d3bff.
//
// Solidity: function getProposedUpgradeTimestamp() view returns(uint256)
func (_IZkSync *IZkSyncCaller) GetProposedUpgradeTimestamp(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _IZkSync.contract.Call(opts, &out, "getProposedUpgradeTimestamp")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetProposedUpgradeTimestamp is a free data retrieval call binding the contract method 0xe39d3bff.
//
// Solidity: function getProposedUpgradeTimestamp() view returns(uint256)
func (_IZkSync *IZkSyncSession) GetProposedUpgradeTimestamp() (*big.Int, error) {
	return _IZkSync.Contract.GetProposedUpgradeTimestamp(&_IZkSync.CallOpts)
}

// GetProposedUpgradeTimestamp is a free data retrieval call binding the contract method 0xe39d3bff.
//
// Solidity: function getProposedUpgradeTimestamp() view returns(uint256)
func (_IZkSync *IZkSyncCallerSession) GetProposedUpgradeTimestamp() (*big.Int, error) {
	return _IZkSync.Contract.GetProposedUpgradeTimestamp(&_IZkSync.CallOpts)
}

// GetSecurityCouncil is a free data retrieval call binding the contract method 0x0ef240a0.
//
// Solidity: function getSecurityCouncil() view returns(address)
func (_IZkSync *IZkSyncCaller) GetSecurityCouncil(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _IZkSync.contract.Call(opts, &out, "getSecurityCouncil")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetSecurityCouncil is a free data retrieval call binding the contract method 0x0ef240a0.
//
// Solidity: function getSecurityCouncil() view returns(address)
func (_IZkSync *IZkSyncSession) GetSecurityCouncil() (common.Address, error) {
	return _IZkSync.Contract.GetSecurityCouncil(&_IZkSync.CallOpts)
}

// GetSecurityCouncil is a free data retrieval call binding the contract method 0x0ef240a0.
//
// Solidity: function getSecurityCouncil() view returns(address)
func (_IZkSync *IZkSyncCallerSession) GetSecurityCouncil() (common.Address, error) {
	return _IZkSync.Contract.GetSecurityCouncil(&_IZkSync.CallOpts)
}

// GetTotalBlocksCommitted is a free data retrieval call binding the contract method 0xfe26699e.
//
// Solidity: function getTotalBlocksCommitted() view returns(uint256)
func (_IZkSync *IZkSyncCaller) GetTotalBlocksCommitted(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _IZkSync.contract.Call(opts, &out, "getTotalBlocksCommitted")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetTotalBlocksCommitted is a free data retrieval call binding the contract method 0xfe26699e.
//
// Solidity: function getTotalBlocksCommitted() view returns(uint256)
func (_IZkSync *IZkSyncSession) GetTotalBlocksCommitted() (*big.Int, error) {
	return _IZkSync.Contract.GetTotalBlocksCommitted(&_IZkSync.CallOpts)
}

// GetTotalBlocksCommitted is a free data retrieval call binding the contract method 0xfe26699e.
//
// Solidity: function getTotalBlocksCommitted() view returns(uint256)
func (_IZkSync *IZkSyncCallerSession) GetTotalBlocksCommitted() (*big.Int, error) {
	return _IZkSync.Contract.GetTotalBlocksCommitted(&_IZkSync.CallOpts)
}

// GetTotalBlocksExecuted is a free data retrieval call binding the contract method 0x39607382.
//
// Solidity: function getTotalBlocksExecuted() view returns(uint256)
func (_IZkSync *IZkSyncCaller) GetTotalBlocksExecuted(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _IZkSync.contract.Call(opts, &out, "getTotalBlocksExecuted")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetTotalBlocksExecuted is a free data retrieval call binding the contract method 0x39607382.
//
// Solidity: function getTotalBlocksExecuted() view returns(uint256)
func (_IZkSync *IZkSyncSession) GetTotalBlocksExecuted() (*big.Int, error) {
	return _IZkSync.Contract.GetTotalBlocksExecuted(&_IZkSync.CallOpts)
}

// GetTotalBlocksExecuted is a free data retrieval call binding the contract method 0x39607382.
//
// Solidity: function getTotalBlocksExecuted() view returns(uint256)
func (_IZkSync *IZkSyncCallerSession) GetTotalBlocksExecuted() (*big.Int, error) {
	return _IZkSync.Contract.GetTotalBlocksExecuted(&_IZkSync.CallOpts)
}

// GetTotalBlocksVerified is a free data retrieval call binding the contract method 0xaf6a2dcd.
//
// Solidity: function getTotalBlocksVerified() view returns(uint256)
func (_IZkSync *IZkSyncCaller) GetTotalBlocksVerified(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _IZkSync.contract.Call(opts, &out, "getTotalBlocksVerified")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetTotalBlocksVerified is a free data retrieval call binding the contract method 0xaf6a2dcd.
//
// Solidity: function getTotalBlocksVerified() view returns(uint256)
func (_IZkSync *IZkSyncSession) GetTotalBlocksVerified() (*big.Int, error) {
	return _IZkSync.Contract.GetTotalBlocksVerified(&_IZkSync.CallOpts)
}

// GetTotalBlocksVerified is a free data retrieval call binding the contract method 0xaf6a2dcd.
//
// Solidity: function getTotalBlocksVerified() view returns(uint256)
func (_IZkSync *IZkSyncCallerSession) GetTotalBlocksVerified() (*big.Int, error) {
	return _IZkSync.Contract.GetTotalBlocksVerified(&_IZkSync.CallOpts)
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

// GetUpgradeProposalState is a free data retrieval call binding the contract method 0xa39980a0.
//
// Solidity: function getUpgradeProposalState() view returns(uint8)
func (_IZkSync *IZkSyncCaller) GetUpgradeProposalState(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _IZkSync.contract.Call(opts, &out, "getUpgradeProposalState")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// GetUpgradeProposalState is a free data retrieval call binding the contract method 0xa39980a0.
//
// Solidity: function getUpgradeProposalState() view returns(uint8)
func (_IZkSync *IZkSyncSession) GetUpgradeProposalState() (uint8, error) {
	return _IZkSync.Contract.GetUpgradeProposalState(&_IZkSync.CallOpts)
}

// GetUpgradeProposalState is a free data retrieval call binding the contract method 0xa39980a0.
//
// Solidity: function getUpgradeProposalState() view returns(uint8)
func (_IZkSync *IZkSyncCallerSession) GetUpgradeProposalState() (uint8, error) {
	return _IZkSync.Contract.GetUpgradeProposalState(&_IZkSync.CallOpts)
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

// IsApprovedBySecurityCouncil is a free data retrieval call binding the contract method 0x3db920ce.
//
// Solidity: function isApprovedBySecurityCouncil() view returns(bool)
func (_IZkSync *IZkSyncCaller) IsApprovedBySecurityCouncil(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _IZkSync.contract.Call(opts, &out, "isApprovedBySecurityCouncil")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsApprovedBySecurityCouncil is a free data retrieval call binding the contract method 0x3db920ce.
//
// Solidity: function isApprovedBySecurityCouncil() view returns(bool)
func (_IZkSync *IZkSyncSession) IsApprovedBySecurityCouncil() (bool, error) {
	return _IZkSync.Contract.IsApprovedBySecurityCouncil(&_IZkSync.CallOpts)
}

// IsApprovedBySecurityCouncil is a free data retrieval call binding the contract method 0x3db920ce.
//
// Solidity: function isApprovedBySecurityCouncil() view returns(bool)
func (_IZkSync *IZkSyncCallerSession) IsApprovedBySecurityCouncil() (bool, error) {
	return _IZkSync.Contract.IsApprovedBySecurityCouncil(&_IZkSync.CallOpts)
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
// Solidity: function isEthWithdrawalFinalized(uint256 _l2BlockNumber, uint256 _l2MessageIndex) view returns(bool)
func (_IZkSync *IZkSyncCaller) IsEthWithdrawalFinalized(opts *bind.CallOpts, _l2BlockNumber *big.Int, _l2MessageIndex *big.Int) (bool, error) {
	var out []interface{}
	err := _IZkSync.contract.Call(opts, &out, "isEthWithdrawalFinalized", _l2BlockNumber, _l2MessageIndex)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsEthWithdrawalFinalized is a free data retrieval call binding the contract method 0xbd7c5412.
//
// Solidity: function isEthWithdrawalFinalized(uint256 _l2BlockNumber, uint256 _l2MessageIndex) view returns(bool)
func (_IZkSync *IZkSyncSession) IsEthWithdrawalFinalized(_l2BlockNumber *big.Int, _l2MessageIndex *big.Int) (bool, error) {
	return _IZkSync.Contract.IsEthWithdrawalFinalized(&_IZkSync.CallOpts, _l2BlockNumber, _l2MessageIndex)
}

// IsEthWithdrawalFinalized is a free data retrieval call binding the contract method 0xbd7c5412.
//
// Solidity: function isEthWithdrawalFinalized(uint256 _l2BlockNumber, uint256 _l2MessageIndex) view returns(bool)
func (_IZkSync *IZkSyncCallerSession) IsEthWithdrawalFinalized(_l2BlockNumber *big.Int, _l2MessageIndex *big.Int) (bool, error) {
	return _IZkSync.Contract.IsEthWithdrawalFinalized(&_IZkSync.CallOpts, _l2BlockNumber, _l2MessageIndex)
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
// Solidity: function l2LogsRootHash(uint256 _blockNumber) view returns(bytes32 hash)
func (_IZkSync *IZkSyncCaller) L2LogsRootHash(opts *bind.CallOpts, _blockNumber *big.Int) ([32]byte, error) {
	var out []interface{}
	err := _IZkSync.contract.Call(opts, &out, "l2LogsRootHash", _blockNumber)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// L2LogsRootHash is a free data retrieval call binding the contract method 0x9cd939e4.
//
// Solidity: function l2LogsRootHash(uint256 _blockNumber) view returns(bytes32 hash)
func (_IZkSync *IZkSyncSession) L2LogsRootHash(_blockNumber *big.Int) ([32]byte, error) {
	return _IZkSync.Contract.L2LogsRootHash(&_IZkSync.CallOpts, _blockNumber)
}

// L2LogsRootHash is a free data retrieval call binding the contract method 0x9cd939e4.
//
// Solidity: function l2LogsRootHash(uint256 _blockNumber) view returns(bytes32 hash)
func (_IZkSync *IZkSyncCallerSession) L2LogsRootHash(_blockNumber *big.Int) ([32]byte, error) {
	return _IZkSync.Contract.L2LogsRootHash(&_IZkSync.CallOpts, _blockNumber)
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
// Solidity: function proveL1ToL2TransactionStatus(bytes32 _l2TxHash, uint256 _l2BlockNumber, uint256 _l2MessageIndex, uint16 _l2TxNumberInBlock, bytes32[] _merkleProof, uint8 _status) view returns(bool)
func (_IZkSync *IZkSyncCaller) ProveL1ToL2TransactionStatus(opts *bind.CallOpts, _l2TxHash [32]byte, _l2BlockNumber *big.Int, _l2MessageIndex *big.Int, _l2TxNumberInBlock uint16, _merkleProof [][32]byte, _status uint8) (bool, error) {
	var out []interface{}
	err := _IZkSync.contract.Call(opts, &out, "proveL1ToL2TransactionStatus", _l2TxHash, _l2BlockNumber, _l2MessageIndex, _l2TxNumberInBlock, _merkleProof, _status)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// ProveL1ToL2TransactionStatus is a free data retrieval call binding the contract method 0x042901c7.
//
// Solidity: function proveL1ToL2TransactionStatus(bytes32 _l2TxHash, uint256 _l2BlockNumber, uint256 _l2MessageIndex, uint16 _l2TxNumberInBlock, bytes32[] _merkleProof, uint8 _status) view returns(bool)
func (_IZkSync *IZkSyncSession) ProveL1ToL2TransactionStatus(_l2TxHash [32]byte, _l2BlockNumber *big.Int, _l2MessageIndex *big.Int, _l2TxNumberInBlock uint16, _merkleProof [][32]byte, _status uint8) (bool, error) {
	return _IZkSync.Contract.ProveL1ToL2TransactionStatus(&_IZkSync.CallOpts, _l2TxHash, _l2BlockNumber, _l2MessageIndex, _l2TxNumberInBlock, _merkleProof, _status)
}

// ProveL1ToL2TransactionStatus is a free data retrieval call binding the contract method 0x042901c7.
//
// Solidity: function proveL1ToL2TransactionStatus(bytes32 _l2TxHash, uint256 _l2BlockNumber, uint256 _l2MessageIndex, uint16 _l2TxNumberInBlock, bytes32[] _merkleProof, uint8 _status) view returns(bool)
func (_IZkSync *IZkSyncCallerSession) ProveL1ToL2TransactionStatus(_l2TxHash [32]byte, _l2BlockNumber *big.Int, _l2MessageIndex *big.Int, _l2TxNumberInBlock uint16, _merkleProof [][32]byte, _status uint8) (bool, error) {
	return _IZkSync.Contract.ProveL1ToL2TransactionStatus(&_IZkSync.CallOpts, _l2TxHash, _l2BlockNumber, _l2MessageIndex, _l2TxNumberInBlock, _merkleProof, _status)
}

// ProveL2LogInclusion is a free data retrieval call binding the contract method 0x263b7f8e.
//
// Solidity: function proveL2LogInclusion(uint256 _blockNumber, uint256 _index, (uint8,bool,uint16,address,bytes32,bytes32) _log, bytes32[] _proof) view returns(bool)
func (_IZkSync *IZkSyncCaller) ProveL2LogInclusion(opts *bind.CallOpts, _blockNumber *big.Int, _index *big.Int, _log L2Log, _proof [][32]byte) (bool, error) {
	var out []interface{}
	err := _IZkSync.contract.Call(opts, &out, "proveL2LogInclusion", _blockNumber, _index, _log, _proof)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// ProveL2LogInclusion is a free data retrieval call binding the contract method 0x263b7f8e.
//
// Solidity: function proveL2LogInclusion(uint256 _blockNumber, uint256 _index, (uint8,bool,uint16,address,bytes32,bytes32) _log, bytes32[] _proof) view returns(bool)
func (_IZkSync *IZkSyncSession) ProveL2LogInclusion(_blockNumber *big.Int, _index *big.Int, _log L2Log, _proof [][32]byte) (bool, error) {
	return _IZkSync.Contract.ProveL2LogInclusion(&_IZkSync.CallOpts, _blockNumber, _index, _log, _proof)
}

// ProveL2LogInclusion is a free data retrieval call binding the contract method 0x263b7f8e.
//
// Solidity: function proveL2LogInclusion(uint256 _blockNumber, uint256 _index, (uint8,bool,uint16,address,bytes32,bytes32) _log, bytes32[] _proof) view returns(bool)
func (_IZkSync *IZkSyncCallerSession) ProveL2LogInclusion(_blockNumber *big.Int, _index *big.Int, _log L2Log, _proof [][32]byte) (bool, error) {
	return _IZkSync.Contract.ProveL2LogInclusion(&_IZkSync.CallOpts, _blockNumber, _index, _log, _proof)
}

// ProveL2MessageInclusion is a free data retrieval call binding the contract method 0xe4948f43.
//
// Solidity: function proveL2MessageInclusion(uint256 _blockNumber, uint256 _index, (uint16,address,bytes) _message, bytes32[] _proof) view returns(bool)
func (_IZkSync *IZkSyncCaller) ProveL2MessageInclusion(opts *bind.CallOpts, _blockNumber *big.Int, _index *big.Int, _message L2Message, _proof [][32]byte) (bool, error) {
	var out []interface{}
	err := _IZkSync.contract.Call(opts, &out, "proveL2MessageInclusion", _blockNumber, _index, _message, _proof)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// ProveL2MessageInclusion is a free data retrieval call binding the contract method 0xe4948f43.
//
// Solidity: function proveL2MessageInclusion(uint256 _blockNumber, uint256 _index, (uint16,address,bytes) _message, bytes32[] _proof) view returns(bool)
func (_IZkSync *IZkSyncSession) ProveL2MessageInclusion(_blockNumber *big.Int, _index *big.Int, _message L2Message, _proof [][32]byte) (bool, error) {
	return _IZkSync.Contract.ProveL2MessageInclusion(&_IZkSync.CallOpts, _blockNumber, _index, _message, _proof)
}

// ProveL2MessageInclusion is a free data retrieval call binding the contract method 0xe4948f43.
//
// Solidity: function proveL2MessageInclusion(uint256 _blockNumber, uint256 _index, (uint16,address,bytes) _message, bytes32[] _proof) view returns(bool)
func (_IZkSync *IZkSyncCallerSession) ProveL2MessageInclusion(_blockNumber *big.Int, _index *big.Int, _message L2Message, _proof [][32]byte) (bool, error) {
	return _IZkSync.Contract.ProveL2MessageInclusion(&_IZkSync.CallOpts, _blockNumber, _index, _message, _proof)
}

// StoredBlockHash is a free data retrieval call binding the contract method 0x74f4d30d.
//
// Solidity: function storedBlockHash(uint256 _blockNumber) view returns(bytes32)
func (_IZkSync *IZkSyncCaller) StoredBlockHash(opts *bind.CallOpts, _blockNumber *big.Int) ([32]byte, error) {
	var out []interface{}
	err := _IZkSync.contract.Call(opts, &out, "storedBlockHash", _blockNumber)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// StoredBlockHash is a free data retrieval call binding the contract method 0x74f4d30d.
//
// Solidity: function storedBlockHash(uint256 _blockNumber) view returns(bytes32)
func (_IZkSync *IZkSyncSession) StoredBlockHash(_blockNumber *big.Int) ([32]byte, error) {
	return _IZkSync.Contract.StoredBlockHash(&_IZkSync.CallOpts, _blockNumber)
}

// StoredBlockHash is a free data retrieval call binding the contract method 0x74f4d30d.
//
// Solidity: function storedBlockHash(uint256 _blockNumber) view returns(bytes32)
func (_IZkSync *IZkSyncCallerSession) StoredBlockHash(_blockNumber *big.Int) ([32]byte, error) {
	return _IZkSync.Contract.StoredBlockHash(&_IZkSync.CallOpts, _blockNumber)
}

// UpgradeProposalHash is a free data retrieval call binding the contract method 0x587809c7.
//
// Solidity: function upgradeProposalHash(((address,uint8,bool,bytes4[])[],address,bytes) _diamondCut, uint256 _proposalId, bytes32 _salt) pure returns(bytes32)
func (_IZkSync *IZkSyncCaller) UpgradeProposalHash(opts *bind.CallOpts, _diamondCut DiamondDiamondCutData, _proposalId *big.Int, _salt [32]byte) ([32]byte, error) {
	var out []interface{}
	err := _IZkSync.contract.Call(opts, &out, "upgradeProposalHash", _diamondCut, _proposalId, _salt)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// UpgradeProposalHash is a free data retrieval call binding the contract method 0x587809c7.
//
// Solidity: function upgradeProposalHash(((address,uint8,bool,bytes4[])[],address,bytes) _diamondCut, uint256 _proposalId, bytes32 _salt) pure returns(bytes32)
func (_IZkSync *IZkSyncSession) UpgradeProposalHash(_diamondCut DiamondDiamondCutData, _proposalId *big.Int, _salt [32]byte) ([32]byte, error) {
	return _IZkSync.Contract.UpgradeProposalHash(&_IZkSync.CallOpts, _diamondCut, _proposalId, _salt)
}

// UpgradeProposalHash is a free data retrieval call binding the contract method 0x587809c7.
//
// Solidity: function upgradeProposalHash(((address,uint8,bool,bytes4[])[],address,bytes) _diamondCut, uint256 _proposalId, bytes32 _salt) pure returns(bytes32)
func (_IZkSync *IZkSyncCallerSession) UpgradeProposalHash(_diamondCut DiamondDiamondCutData, _proposalId *big.Int, _salt [32]byte) ([32]byte, error) {
	return _IZkSync.Contract.UpgradeProposalHash(&_IZkSync.CallOpts, _diamondCut, _proposalId, _salt)
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

// CancelUpgradeProposal is a paid mutator transaction binding the contract method 0x73fb9297.
//
// Solidity: function cancelUpgradeProposal(bytes32 _proposedUpgradeHash) returns()
func (_IZkSync *IZkSyncTransactor) CancelUpgradeProposal(opts *bind.TransactOpts, _proposedUpgradeHash [32]byte) (*types.Transaction, error) {
	return _IZkSync.contract.Transact(opts, "cancelUpgradeProposal", _proposedUpgradeHash)
}

// CancelUpgradeProposal is a paid mutator transaction binding the contract method 0x73fb9297.
//
// Solidity: function cancelUpgradeProposal(bytes32 _proposedUpgradeHash) returns()
func (_IZkSync *IZkSyncSession) CancelUpgradeProposal(_proposedUpgradeHash [32]byte) (*types.Transaction, error) {
	return _IZkSync.Contract.CancelUpgradeProposal(&_IZkSync.TransactOpts, _proposedUpgradeHash)
}

// CancelUpgradeProposal is a paid mutator transaction binding the contract method 0x73fb9297.
//
// Solidity: function cancelUpgradeProposal(bytes32 _proposedUpgradeHash) returns()
func (_IZkSync *IZkSyncTransactorSession) CancelUpgradeProposal(_proposedUpgradeHash [32]byte) (*types.Transaction, error) {
	return _IZkSync.Contract.CancelUpgradeProposal(&_IZkSync.TransactOpts, _proposedUpgradeHash)
}

// CommitBlocks is a paid mutator transaction binding the contract method 0x0c4dd810.
//
// Solidity: function commitBlocks((uint64,bytes32,uint64,uint256,bytes32,bytes32,uint256,bytes32) _lastCommittedBlockData, (uint64,uint64,uint64,bytes32,uint256,bytes32,bytes32,bytes,bytes,bytes,bytes[],bytes[])[] _newBlocksData) returns()
func (_IZkSync *IZkSyncTransactor) CommitBlocks(opts *bind.TransactOpts, _lastCommittedBlockData IExecutorStoredBlockInfo, _newBlocksData []IExecutorCommitBlockInfo) (*types.Transaction, error) {
	return _IZkSync.contract.Transact(opts, "commitBlocks", _lastCommittedBlockData, _newBlocksData)
}

// CommitBlocks is a paid mutator transaction binding the contract method 0x0c4dd810.
//
// Solidity: function commitBlocks((uint64,bytes32,uint64,uint256,bytes32,bytes32,uint256,bytes32) _lastCommittedBlockData, (uint64,uint64,uint64,bytes32,uint256,bytes32,bytes32,bytes,bytes,bytes,bytes[],bytes[])[] _newBlocksData) returns()
func (_IZkSync *IZkSyncSession) CommitBlocks(_lastCommittedBlockData IExecutorStoredBlockInfo, _newBlocksData []IExecutorCommitBlockInfo) (*types.Transaction, error) {
	return _IZkSync.Contract.CommitBlocks(&_IZkSync.TransactOpts, _lastCommittedBlockData, _newBlocksData)
}

// CommitBlocks is a paid mutator transaction binding the contract method 0x0c4dd810.
//
// Solidity: function commitBlocks((uint64,bytes32,uint64,uint256,bytes32,bytes32,uint256,bytes32) _lastCommittedBlockData, (uint64,uint64,uint64,bytes32,uint256,bytes32,bytes32,bytes,bytes,bytes,bytes[],bytes[])[] _newBlocksData) returns()
func (_IZkSync *IZkSyncTransactorSession) CommitBlocks(_lastCommittedBlockData IExecutorStoredBlockInfo, _newBlocksData []IExecutorCommitBlockInfo) (*types.Transaction, error) {
	return _IZkSync.Contract.CommitBlocks(&_IZkSync.TransactOpts, _lastCommittedBlockData, _newBlocksData)
}

// ExecuteBlocks is a paid mutator transaction binding the contract method 0xce9dcf16.
//
// Solidity: function executeBlocks((uint64,bytes32,uint64,uint256,bytes32,bytes32,uint256,bytes32)[] _blocksData) returns()
func (_IZkSync *IZkSyncTransactor) ExecuteBlocks(opts *bind.TransactOpts, _blocksData []IExecutorStoredBlockInfo) (*types.Transaction, error) {
	return _IZkSync.contract.Transact(opts, "executeBlocks", _blocksData)
}

// ExecuteBlocks is a paid mutator transaction binding the contract method 0xce9dcf16.
//
// Solidity: function executeBlocks((uint64,bytes32,uint64,uint256,bytes32,bytes32,uint256,bytes32)[] _blocksData) returns()
func (_IZkSync *IZkSyncSession) ExecuteBlocks(_blocksData []IExecutorStoredBlockInfo) (*types.Transaction, error) {
	return _IZkSync.Contract.ExecuteBlocks(&_IZkSync.TransactOpts, _blocksData)
}

// ExecuteBlocks is a paid mutator transaction binding the contract method 0xce9dcf16.
//
// Solidity: function executeBlocks((uint64,bytes32,uint64,uint256,bytes32,bytes32,uint256,bytes32)[] _blocksData) returns()
func (_IZkSync *IZkSyncTransactorSession) ExecuteBlocks(_blocksData []IExecutorStoredBlockInfo) (*types.Transaction, error) {
	return _IZkSync.Contract.ExecuteBlocks(&_IZkSync.TransactOpts, _blocksData)
}

// ExecuteUpgrade is a paid mutator transaction binding the contract method 0x36d4eb84.
//
// Solidity: function executeUpgrade(((address,uint8,bool,bytes4[])[],address,bytes) _diamondCut, bytes32 _proposalSalt) returns()
func (_IZkSync *IZkSyncTransactor) ExecuteUpgrade(opts *bind.TransactOpts, _diamondCut DiamondDiamondCutData, _proposalSalt [32]byte) (*types.Transaction, error) {
	return _IZkSync.contract.Transact(opts, "executeUpgrade", _diamondCut, _proposalSalt)
}

// ExecuteUpgrade is a paid mutator transaction binding the contract method 0x36d4eb84.
//
// Solidity: function executeUpgrade(((address,uint8,bool,bytes4[])[],address,bytes) _diamondCut, bytes32 _proposalSalt) returns()
func (_IZkSync *IZkSyncSession) ExecuteUpgrade(_diamondCut DiamondDiamondCutData, _proposalSalt [32]byte) (*types.Transaction, error) {
	return _IZkSync.Contract.ExecuteUpgrade(&_IZkSync.TransactOpts, _diamondCut, _proposalSalt)
}

// ExecuteUpgrade is a paid mutator transaction binding the contract method 0x36d4eb84.
//
// Solidity: function executeUpgrade(((address,uint8,bool,bytes4[])[],address,bytes) _diamondCut, bytes32 _proposalSalt) returns()
func (_IZkSync *IZkSyncTransactorSession) ExecuteUpgrade(_diamondCut DiamondDiamondCutData, _proposalSalt [32]byte) (*types.Transaction, error) {
	return _IZkSync.Contract.ExecuteUpgrade(&_IZkSync.TransactOpts, _diamondCut, _proposalSalt)
}

// FinalizeEthWithdrawal is a paid mutator transaction binding the contract method 0x6c0960f9.
//
// Solidity: function finalizeEthWithdrawal(uint256 _l2BlockNumber, uint256 _l2MessageIndex, uint16 _l2TxNumberInBlock, bytes _message, bytes32[] _merkleProof) returns()
func (_IZkSync *IZkSyncTransactor) FinalizeEthWithdrawal(opts *bind.TransactOpts, _l2BlockNumber *big.Int, _l2MessageIndex *big.Int, _l2TxNumberInBlock uint16, _message []byte, _merkleProof [][32]byte) (*types.Transaction, error) {
	return _IZkSync.contract.Transact(opts, "finalizeEthWithdrawal", _l2BlockNumber, _l2MessageIndex, _l2TxNumberInBlock, _message, _merkleProof)
}

// FinalizeEthWithdrawal is a paid mutator transaction binding the contract method 0x6c0960f9.
//
// Solidity: function finalizeEthWithdrawal(uint256 _l2BlockNumber, uint256 _l2MessageIndex, uint16 _l2TxNumberInBlock, bytes _message, bytes32[] _merkleProof) returns()
func (_IZkSync *IZkSyncSession) FinalizeEthWithdrawal(_l2BlockNumber *big.Int, _l2MessageIndex *big.Int, _l2TxNumberInBlock uint16, _message []byte, _merkleProof [][32]byte) (*types.Transaction, error) {
	return _IZkSync.Contract.FinalizeEthWithdrawal(&_IZkSync.TransactOpts, _l2BlockNumber, _l2MessageIndex, _l2TxNumberInBlock, _message, _merkleProof)
}

// FinalizeEthWithdrawal is a paid mutator transaction binding the contract method 0x6c0960f9.
//
// Solidity: function finalizeEthWithdrawal(uint256 _l2BlockNumber, uint256 _l2MessageIndex, uint16 _l2TxNumberInBlock, bytes _message, bytes32[] _merkleProof) returns()
func (_IZkSync *IZkSyncTransactorSession) FinalizeEthWithdrawal(_l2BlockNumber *big.Int, _l2MessageIndex *big.Int, _l2TxNumberInBlock uint16, _message []byte, _merkleProof [][32]byte) (*types.Transaction, error) {
	return _IZkSync.Contract.FinalizeEthWithdrawal(&_IZkSync.TransactOpts, _l2BlockNumber, _l2MessageIndex, _l2TxNumberInBlock, _message, _merkleProof)
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

// ProposeShadowUpgrade is a paid mutator transaction binding the contract method 0x0551448c.
//
// Solidity: function proposeShadowUpgrade(bytes32 _proposalHash, uint40 _proposalId) returns()
func (_IZkSync *IZkSyncTransactor) ProposeShadowUpgrade(opts *bind.TransactOpts, _proposalHash [32]byte, _proposalId *big.Int) (*types.Transaction, error) {
	return _IZkSync.contract.Transact(opts, "proposeShadowUpgrade", _proposalHash, _proposalId)
}

// ProposeShadowUpgrade is a paid mutator transaction binding the contract method 0x0551448c.
//
// Solidity: function proposeShadowUpgrade(bytes32 _proposalHash, uint40 _proposalId) returns()
func (_IZkSync *IZkSyncSession) ProposeShadowUpgrade(_proposalHash [32]byte, _proposalId *big.Int) (*types.Transaction, error) {
	return _IZkSync.Contract.ProposeShadowUpgrade(&_IZkSync.TransactOpts, _proposalHash, _proposalId)
}

// ProposeShadowUpgrade is a paid mutator transaction binding the contract method 0x0551448c.
//
// Solidity: function proposeShadowUpgrade(bytes32 _proposalHash, uint40 _proposalId) returns()
func (_IZkSync *IZkSyncTransactorSession) ProposeShadowUpgrade(_proposalHash [32]byte, _proposalId *big.Int) (*types.Transaction, error) {
	return _IZkSync.Contract.ProposeShadowUpgrade(&_IZkSync.TransactOpts, _proposalHash, _proposalId)
}

// ProposeTransparentUpgrade is a paid mutator transaction binding the contract method 0x8043760a.
//
// Solidity: function proposeTransparentUpgrade(((address,uint8,bool,bytes4[])[],address,bytes) _diamondCut, uint40 _proposalId) returns()
func (_IZkSync *IZkSyncTransactor) ProposeTransparentUpgrade(opts *bind.TransactOpts, _diamondCut DiamondDiamondCutData, _proposalId *big.Int) (*types.Transaction, error) {
	return _IZkSync.contract.Transact(opts, "proposeTransparentUpgrade", _diamondCut, _proposalId)
}

// ProposeTransparentUpgrade is a paid mutator transaction binding the contract method 0x8043760a.
//
// Solidity: function proposeTransparentUpgrade(((address,uint8,bool,bytes4[])[],address,bytes) _diamondCut, uint40 _proposalId) returns()
func (_IZkSync *IZkSyncSession) ProposeTransparentUpgrade(_diamondCut DiamondDiamondCutData, _proposalId *big.Int) (*types.Transaction, error) {
	return _IZkSync.Contract.ProposeTransparentUpgrade(&_IZkSync.TransactOpts, _diamondCut, _proposalId)
}

// ProposeTransparentUpgrade is a paid mutator transaction binding the contract method 0x8043760a.
//
// Solidity: function proposeTransparentUpgrade(((address,uint8,bool,bytes4[])[],address,bytes) _diamondCut, uint40 _proposalId) returns()
func (_IZkSync *IZkSyncTransactorSession) ProposeTransparentUpgrade(_diamondCut DiamondDiamondCutData, _proposalId *big.Int) (*types.Transaction, error) {
	return _IZkSync.Contract.ProposeTransparentUpgrade(&_IZkSync.TransactOpts, _diamondCut, _proposalId)
}

// ProveBlocks is a paid mutator transaction binding the contract method 0x7739cbe7.
//
// Solidity: function proveBlocks((uint64,bytes32,uint64,uint256,bytes32,bytes32,uint256,bytes32) _prevBlock, (uint64,bytes32,uint64,uint256,bytes32,bytes32,uint256,bytes32)[] _committedBlocks, (uint256[],uint256[]) _proof) returns()
func (_IZkSync *IZkSyncTransactor) ProveBlocks(opts *bind.TransactOpts, _prevBlock IExecutorStoredBlockInfo, _committedBlocks []IExecutorStoredBlockInfo, _proof IExecutorProofInput) (*types.Transaction, error) {
	return _IZkSync.contract.Transact(opts, "proveBlocks", _prevBlock, _committedBlocks, _proof)
}

// ProveBlocks is a paid mutator transaction binding the contract method 0x7739cbe7.
//
// Solidity: function proveBlocks((uint64,bytes32,uint64,uint256,bytes32,bytes32,uint256,bytes32) _prevBlock, (uint64,bytes32,uint64,uint256,bytes32,bytes32,uint256,bytes32)[] _committedBlocks, (uint256[],uint256[]) _proof) returns()
func (_IZkSync *IZkSyncSession) ProveBlocks(_prevBlock IExecutorStoredBlockInfo, _committedBlocks []IExecutorStoredBlockInfo, _proof IExecutorProofInput) (*types.Transaction, error) {
	return _IZkSync.Contract.ProveBlocks(&_IZkSync.TransactOpts, _prevBlock, _committedBlocks, _proof)
}

// ProveBlocks is a paid mutator transaction binding the contract method 0x7739cbe7.
//
// Solidity: function proveBlocks((uint64,bytes32,uint64,uint256,bytes32,bytes32,uint256,bytes32) _prevBlock, (uint64,bytes32,uint64,uint256,bytes32,bytes32,uint256,bytes32)[] _committedBlocks, (uint256[],uint256[]) _proof) returns()
func (_IZkSync *IZkSyncTransactorSession) ProveBlocks(_prevBlock IExecutorStoredBlockInfo, _committedBlocks []IExecutorStoredBlockInfo, _proof IExecutorProofInput) (*types.Transaction, error) {
	return _IZkSync.Contract.ProveBlocks(&_IZkSync.TransactOpts, _prevBlock, _committedBlocks, _proof)
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

// RevertBlocks is a paid mutator transaction binding the contract method 0xa9a2d18a.
//
// Solidity: function revertBlocks(uint256 _newLastBlock) returns()
func (_IZkSync *IZkSyncTransactor) RevertBlocks(opts *bind.TransactOpts, _newLastBlock *big.Int) (*types.Transaction, error) {
	return _IZkSync.contract.Transact(opts, "revertBlocks", _newLastBlock)
}

// RevertBlocks is a paid mutator transaction binding the contract method 0xa9a2d18a.
//
// Solidity: function revertBlocks(uint256 _newLastBlock) returns()
func (_IZkSync *IZkSyncSession) RevertBlocks(_newLastBlock *big.Int) (*types.Transaction, error) {
	return _IZkSync.Contract.RevertBlocks(&_IZkSync.TransactOpts, _newLastBlock)
}

// RevertBlocks is a paid mutator transaction binding the contract method 0xa9a2d18a.
//
// Solidity: function revertBlocks(uint256 _newLastBlock) returns()
func (_IZkSync *IZkSyncTransactorSession) RevertBlocks(_newLastBlock *big.Int) (*types.Transaction, error) {
	return _IZkSync.Contract.RevertBlocks(&_IZkSync.TransactOpts, _newLastBlock)
}

// SecurityCouncilUpgradeApprove is a paid mutator transaction binding the contract method 0xbeda4b12.
//
// Solidity: function securityCouncilUpgradeApprove(bytes32 _upgradeProposalHash) returns()
func (_IZkSync *IZkSyncTransactor) SecurityCouncilUpgradeApprove(opts *bind.TransactOpts, _upgradeProposalHash [32]byte) (*types.Transaction, error) {
	return _IZkSync.contract.Transact(opts, "securityCouncilUpgradeApprove", _upgradeProposalHash)
}

// SecurityCouncilUpgradeApprove is a paid mutator transaction binding the contract method 0xbeda4b12.
//
// Solidity: function securityCouncilUpgradeApprove(bytes32 _upgradeProposalHash) returns()
func (_IZkSync *IZkSyncSession) SecurityCouncilUpgradeApprove(_upgradeProposalHash [32]byte) (*types.Transaction, error) {
	return _IZkSync.Contract.SecurityCouncilUpgradeApprove(&_IZkSync.TransactOpts, _upgradeProposalHash)
}

// SecurityCouncilUpgradeApprove is a paid mutator transaction binding the contract method 0xbeda4b12.
//
// Solidity: function securityCouncilUpgradeApprove(bytes32 _upgradeProposalHash) returns()
func (_IZkSync *IZkSyncTransactorSession) SecurityCouncilUpgradeApprove(_upgradeProposalHash [32]byte) (*types.Transaction, error) {
	return _IZkSync.Contract.SecurityCouncilUpgradeApprove(&_IZkSync.TransactOpts, _upgradeProposalHash)
}

// SetAllowList is a paid mutator transaction binding the contract method 0xed6d06c0.
//
// Solidity: function setAllowList(address _newAllowList) returns()
func (_IZkSync *IZkSyncTransactor) SetAllowList(opts *bind.TransactOpts, _newAllowList common.Address) (*types.Transaction, error) {
	return _IZkSync.contract.Transact(opts, "setAllowList", _newAllowList)
}

// SetAllowList is a paid mutator transaction binding the contract method 0xed6d06c0.
//
// Solidity: function setAllowList(address _newAllowList) returns()
func (_IZkSync *IZkSyncSession) SetAllowList(_newAllowList common.Address) (*types.Transaction, error) {
	return _IZkSync.Contract.SetAllowList(&_IZkSync.TransactOpts, _newAllowList)
}

// SetAllowList is a paid mutator transaction binding the contract method 0xed6d06c0.
//
// Solidity: function setAllowList(address _newAllowList) returns()
func (_IZkSync *IZkSyncTransactorSession) SetAllowList(_newAllowList common.Address) (*types.Transaction, error) {
	return _IZkSync.Contract.SetAllowList(&_IZkSync.TransactOpts, _newAllowList)
}

// SetL2BootloaderBytecodeHash is a paid mutator transaction binding the contract method 0x86cb9909.
//
// Solidity: function setL2BootloaderBytecodeHash(bytes32 _l2BootloaderBytecodeHash) returns()
func (_IZkSync *IZkSyncTransactor) SetL2BootloaderBytecodeHash(opts *bind.TransactOpts, _l2BootloaderBytecodeHash [32]byte) (*types.Transaction, error) {
	return _IZkSync.contract.Transact(opts, "setL2BootloaderBytecodeHash", _l2BootloaderBytecodeHash)
}

// SetL2BootloaderBytecodeHash is a paid mutator transaction binding the contract method 0x86cb9909.
//
// Solidity: function setL2BootloaderBytecodeHash(bytes32 _l2BootloaderBytecodeHash) returns()
func (_IZkSync *IZkSyncSession) SetL2BootloaderBytecodeHash(_l2BootloaderBytecodeHash [32]byte) (*types.Transaction, error) {
	return _IZkSync.Contract.SetL2BootloaderBytecodeHash(&_IZkSync.TransactOpts, _l2BootloaderBytecodeHash)
}

// SetL2BootloaderBytecodeHash is a paid mutator transaction binding the contract method 0x86cb9909.
//
// Solidity: function setL2BootloaderBytecodeHash(bytes32 _l2BootloaderBytecodeHash) returns()
func (_IZkSync *IZkSyncTransactorSession) SetL2BootloaderBytecodeHash(_l2BootloaderBytecodeHash [32]byte) (*types.Transaction, error) {
	return _IZkSync.Contract.SetL2BootloaderBytecodeHash(&_IZkSync.TransactOpts, _l2BootloaderBytecodeHash)
}

// SetL2DefaultAccountBytecodeHash is a paid mutator transaction binding the contract method 0x0707ac09.
//
// Solidity: function setL2DefaultAccountBytecodeHash(bytes32 _l2DefaultAccountBytecodeHash) returns()
func (_IZkSync *IZkSyncTransactor) SetL2DefaultAccountBytecodeHash(opts *bind.TransactOpts, _l2DefaultAccountBytecodeHash [32]byte) (*types.Transaction, error) {
	return _IZkSync.contract.Transact(opts, "setL2DefaultAccountBytecodeHash", _l2DefaultAccountBytecodeHash)
}

// SetL2DefaultAccountBytecodeHash is a paid mutator transaction binding the contract method 0x0707ac09.
//
// Solidity: function setL2DefaultAccountBytecodeHash(bytes32 _l2DefaultAccountBytecodeHash) returns()
func (_IZkSync *IZkSyncSession) SetL2DefaultAccountBytecodeHash(_l2DefaultAccountBytecodeHash [32]byte) (*types.Transaction, error) {
	return _IZkSync.Contract.SetL2DefaultAccountBytecodeHash(&_IZkSync.TransactOpts, _l2DefaultAccountBytecodeHash)
}

// SetL2DefaultAccountBytecodeHash is a paid mutator transaction binding the contract method 0x0707ac09.
//
// Solidity: function setL2DefaultAccountBytecodeHash(bytes32 _l2DefaultAccountBytecodeHash) returns()
func (_IZkSync *IZkSyncTransactorSession) SetL2DefaultAccountBytecodeHash(_l2DefaultAccountBytecodeHash [32]byte) (*types.Transaction, error) {
	return _IZkSync.Contract.SetL2DefaultAccountBytecodeHash(&_IZkSync.TransactOpts, _l2DefaultAccountBytecodeHash)
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

// SetVerifier is a paid mutator transaction binding the contract method 0x5437988d.
//
// Solidity: function setVerifier(address _newVerifier) returns()
func (_IZkSync *IZkSyncTransactor) SetVerifier(opts *bind.TransactOpts, _newVerifier common.Address) (*types.Transaction, error) {
	return _IZkSync.contract.Transact(opts, "setVerifier", _newVerifier)
}

// SetVerifier is a paid mutator transaction binding the contract method 0x5437988d.
//
// Solidity: function setVerifier(address _newVerifier) returns()
func (_IZkSync *IZkSyncSession) SetVerifier(_newVerifier common.Address) (*types.Transaction, error) {
	return _IZkSync.Contract.SetVerifier(&_IZkSync.TransactOpts, _newVerifier)
}

// SetVerifier is a paid mutator transaction binding the contract method 0x5437988d.
//
// Solidity: function setVerifier(address _newVerifier) returns()
func (_IZkSync *IZkSyncTransactorSession) SetVerifier(_newVerifier common.Address) (*types.Transaction, error) {
	return _IZkSync.Contract.SetVerifier(&_IZkSync.TransactOpts, _newVerifier)
}

// SetVerifierParams is a paid mutator transaction binding the contract method 0x0b508883.
//
// Solidity: function setVerifierParams((bytes32,bytes32,bytes32) _newVerifierParams) returns()
func (_IZkSync *IZkSyncTransactor) SetVerifierParams(opts *bind.TransactOpts, _newVerifierParams VerifierParams) (*types.Transaction, error) {
	return _IZkSync.contract.Transact(opts, "setVerifierParams", _newVerifierParams)
}

// SetVerifierParams is a paid mutator transaction binding the contract method 0x0b508883.
//
// Solidity: function setVerifierParams((bytes32,bytes32,bytes32) _newVerifierParams) returns()
func (_IZkSync *IZkSyncSession) SetVerifierParams(_newVerifierParams VerifierParams) (*types.Transaction, error) {
	return _IZkSync.Contract.SetVerifierParams(&_IZkSync.TransactOpts, _newVerifierParams)
}

// SetVerifierParams is a paid mutator transaction binding the contract method 0x0b508883.
//
// Solidity: function setVerifierParams((bytes32,bytes32,bytes32) _newVerifierParams) returns()
func (_IZkSync *IZkSyncTransactorSession) SetVerifierParams(_newVerifierParams VerifierParams) (*types.Transaction, error) {
	return _IZkSync.Contract.SetVerifierParams(&_IZkSync.TransactOpts, _newVerifierParams)
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
	BlockNumber *big.Int
	BlockHash   [32]byte
	Commitment  [32]byte
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterBlockCommit is a free log retrieval operation binding the contract event 0x8f2916b2f2d78cc5890ead36c06c0f6d5d112c7e103589947e8e2f0d6eddb763.
//
// Solidity: event BlockCommit(uint256 indexed blockNumber, bytes32 indexed blockHash, bytes32 indexed commitment)
func (_IZkSync *IZkSyncFilterer) FilterBlockCommit(opts *bind.FilterOpts, blockNumber []*big.Int, blockHash [][32]byte, commitment [][32]byte) (*IZkSyncBlockCommitIterator, error) {

	var blockNumberRule []interface{}
	for _, blockNumberItem := range blockNumber {
		blockNumberRule = append(blockNumberRule, blockNumberItem)
	}
	var blockHashRule []interface{}
	for _, blockHashItem := range blockHash {
		blockHashRule = append(blockHashRule, blockHashItem)
	}
	var commitmentRule []interface{}
	for _, commitmentItem := range commitment {
		commitmentRule = append(commitmentRule, commitmentItem)
	}

	logs, sub, err := _IZkSync.contract.FilterLogs(opts, "BlockCommit", blockNumberRule, blockHashRule, commitmentRule)
	if err != nil {
		return nil, err
	}
	return &IZkSyncBlockCommitIterator{contract: _IZkSync.contract, event: "BlockCommit", logs: logs, sub: sub}, nil
}

// WatchBlockCommit is a free log subscription operation binding the contract event 0x8f2916b2f2d78cc5890ead36c06c0f6d5d112c7e103589947e8e2f0d6eddb763.
//
// Solidity: event BlockCommit(uint256 indexed blockNumber, bytes32 indexed blockHash, bytes32 indexed commitment)
func (_IZkSync *IZkSyncFilterer) WatchBlockCommit(opts *bind.WatchOpts, sink chan<- *IZkSyncBlockCommit, blockNumber []*big.Int, blockHash [][32]byte, commitment [][32]byte) (event.Subscription, error) {

	var blockNumberRule []interface{}
	for _, blockNumberItem := range blockNumber {
		blockNumberRule = append(blockNumberRule, blockNumberItem)
	}
	var blockHashRule []interface{}
	for _, blockHashItem := range blockHash {
		blockHashRule = append(blockHashRule, blockHashItem)
	}
	var commitmentRule []interface{}
	for _, commitmentItem := range commitment {
		commitmentRule = append(commitmentRule, commitmentItem)
	}

	logs, sub, err := _IZkSync.contract.WatchLogs(opts, "BlockCommit", blockNumberRule, blockHashRule, commitmentRule)
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
// Solidity: event BlockCommit(uint256 indexed blockNumber, bytes32 indexed blockHash, bytes32 indexed commitment)
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
	BlockNumber *big.Int
	BlockHash   [32]byte
	Commitment  [32]byte
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterBlockExecution is a free log retrieval operation binding the contract event 0x2402307311a4d6604e4e7b4c8a15a7e1213edb39c16a31efa70afb06030d3165.
//
// Solidity: event BlockExecution(uint256 indexed blockNumber, bytes32 indexed blockHash, bytes32 indexed commitment)
func (_IZkSync *IZkSyncFilterer) FilterBlockExecution(opts *bind.FilterOpts, blockNumber []*big.Int, blockHash [][32]byte, commitment [][32]byte) (*IZkSyncBlockExecutionIterator, error) {

	var blockNumberRule []interface{}
	for _, blockNumberItem := range blockNumber {
		blockNumberRule = append(blockNumberRule, blockNumberItem)
	}
	var blockHashRule []interface{}
	for _, blockHashItem := range blockHash {
		blockHashRule = append(blockHashRule, blockHashItem)
	}
	var commitmentRule []interface{}
	for _, commitmentItem := range commitment {
		commitmentRule = append(commitmentRule, commitmentItem)
	}

	logs, sub, err := _IZkSync.contract.FilterLogs(opts, "BlockExecution", blockNumberRule, blockHashRule, commitmentRule)
	if err != nil {
		return nil, err
	}
	return &IZkSyncBlockExecutionIterator{contract: _IZkSync.contract, event: "BlockExecution", logs: logs, sub: sub}, nil
}

// WatchBlockExecution is a free log subscription operation binding the contract event 0x2402307311a4d6604e4e7b4c8a15a7e1213edb39c16a31efa70afb06030d3165.
//
// Solidity: event BlockExecution(uint256 indexed blockNumber, bytes32 indexed blockHash, bytes32 indexed commitment)
func (_IZkSync *IZkSyncFilterer) WatchBlockExecution(opts *bind.WatchOpts, sink chan<- *IZkSyncBlockExecution, blockNumber []*big.Int, blockHash [][32]byte, commitment [][32]byte) (event.Subscription, error) {

	var blockNumberRule []interface{}
	for _, blockNumberItem := range blockNumber {
		blockNumberRule = append(blockNumberRule, blockNumberItem)
	}
	var blockHashRule []interface{}
	for _, blockHashItem := range blockHash {
		blockHashRule = append(blockHashRule, blockHashItem)
	}
	var commitmentRule []interface{}
	for _, commitmentItem := range commitment {
		commitmentRule = append(commitmentRule, commitmentItem)
	}

	logs, sub, err := _IZkSync.contract.WatchLogs(opts, "BlockExecution", blockNumberRule, blockHashRule, commitmentRule)
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
// Solidity: event BlockExecution(uint256 indexed blockNumber, bytes32 indexed blockHash, bytes32 indexed commitment)
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
	TotalBlocksCommitted *big.Int
	TotalBlocksVerified  *big.Int
	TotalBlocksExecuted  *big.Int
	Raw                  types.Log // Blockchain specific contextual infos
}

// FilterBlocksRevert is a free log retrieval operation binding the contract event 0x8bd4b15ea7d1bc41ea9abc3fc487ccb89cd678a00786584714faa9d751c84ee5.
//
// Solidity: event BlocksRevert(uint256 totalBlocksCommitted, uint256 totalBlocksVerified, uint256 totalBlocksExecuted)
func (_IZkSync *IZkSyncFilterer) FilterBlocksRevert(opts *bind.FilterOpts) (*IZkSyncBlocksRevertIterator, error) {

	logs, sub, err := _IZkSync.contract.FilterLogs(opts, "BlocksRevert")
	if err != nil {
		return nil, err
	}
	return &IZkSyncBlocksRevertIterator{contract: _IZkSync.contract, event: "BlocksRevert", logs: logs, sub: sub}, nil
}

// WatchBlocksRevert is a free log subscription operation binding the contract event 0x8bd4b15ea7d1bc41ea9abc3fc487ccb89cd678a00786584714faa9d751c84ee5.
//
// Solidity: event BlocksRevert(uint256 totalBlocksCommitted, uint256 totalBlocksVerified, uint256 totalBlocksExecuted)
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
// Solidity: event BlocksRevert(uint256 totalBlocksCommitted, uint256 totalBlocksVerified, uint256 totalBlocksExecuted)
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
	PreviousLastVerifiedBlock *big.Int
	CurrentLastVerifiedBlock  *big.Int
	Raw                       types.Log // Blockchain specific contextual infos
}

// FilterBlocksVerification is a free log retrieval operation binding the contract event 0x22c9005dd88c18b552a1cd7e8b3b937fcde9ca69213c1f658f54d572e4877a81.
//
// Solidity: event BlocksVerification(uint256 indexed previousLastVerifiedBlock, uint256 indexed currentLastVerifiedBlock)
func (_IZkSync *IZkSyncFilterer) FilterBlocksVerification(opts *bind.FilterOpts, previousLastVerifiedBlock []*big.Int, currentLastVerifiedBlock []*big.Int) (*IZkSyncBlocksVerificationIterator, error) {

	var previousLastVerifiedBlockRule []interface{}
	for _, previousLastVerifiedBlockItem := range previousLastVerifiedBlock {
		previousLastVerifiedBlockRule = append(previousLastVerifiedBlockRule, previousLastVerifiedBlockItem)
	}
	var currentLastVerifiedBlockRule []interface{}
	for _, currentLastVerifiedBlockItem := range currentLastVerifiedBlock {
		currentLastVerifiedBlockRule = append(currentLastVerifiedBlockRule, currentLastVerifiedBlockItem)
	}

	logs, sub, err := _IZkSync.contract.FilterLogs(opts, "BlocksVerification", previousLastVerifiedBlockRule, currentLastVerifiedBlockRule)
	if err != nil {
		return nil, err
	}
	return &IZkSyncBlocksVerificationIterator{contract: _IZkSync.contract, event: "BlocksVerification", logs: logs, sub: sub}, nil
}

// WatchBlocksVerification is a free log subscription operation binding the contract event 0x22c9005dd88c18b552a1cd7e8b3b937fcde9ca69213c1f658f54d572e4877a81.
//
// Solidity: event BlocksVerification(uint256 indexed previousLastVerifiedBlock, uint256 indexed currentLastVerifiedBlock)
func (_IZkSync *IZkSyncFilterer) WatchBlocksVerification(opts *bind.WatchOpts, sink chan<- *IZkSyncBlocksVerification, previousLastVerifiedBlock []*big.Int, currentLastVerifiedBlock []*big.Int) (event.Subscription, error) {

	var previousLastVerifiedBlockRule []interface{}
	for _, previousLastVerifiedBlockItem := range previousLastVerifiedBlock {
		previousLastVerifiedBlockRule = append(previousLastVerifiedBlockRule, previousLastVerifiedBlockItem)
	}
	var currentLastVerifiedBlockRule []interface{}
	for _, currentLastVerifiedBlockItem := range currentLastVerifiedBlock {
		currentLastVerifiedBlockRule = append(currentLastVerifiedBlockRule, currentLastVerifiedBlockItem)
	}

	logs, sub, err := _IZkSync.contract.WatchLogs(opts, "BlocksVerification", previousLastVerifiedBlockRule, currentLastVerifiedBlockRule)
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
// Solidity: event BlocksVerification(uint256 indexed previousLastVerifiedBlock, uint256 indexed currentLastVerifiedBlock)
func (_IZkSync *IZkSyncFilterer) ParseBlocksVerification(log types.Log) (*IZkSyncBlocksVerification, error) {
	event := new(IZkSyncBlocksVerification)
	if err := _IZkSync.contract.UnpackLog(event, "BlocksVerification", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// IZkSyncCancelUpgradeProposalIterator is returned from FilterCancelUpgradeProposal and is used to iterate over the raw logs and unpacked data for CancelUpgradeProposal events raised by the IZkSync contract.
type IZkSyncCancelUpgradeProposalIterator struct {
	Event *IZkSyncCancelUpgradeProposal // Event containing the contract specifics and raw log

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
func (it *IZkSyncCancelUpgradeProposalIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IZkSyncCancelUpgradeProposal)
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
		it.Event = new(IZkSyncCancelUpgradeProposal)
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
func (it *IZkSyncCancelUpgradeProposalIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IZkSyncCancelUpgradeProposalIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IZkSyncCancelUpgradeProposal represents a CancelUpgradeProposal event raised by the IZkSync contract.
type IZkSyncCancelUpgradeProposal struct {
	ProposalId   *big.Int
	ProposalHash [32]byte
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterCancelUpgradeProposal is a free log retrieval operation binding the contract event 0x502f7e4ccbd6b35356f90b3cca05afdb6afea61cea0803229edf974a981c3030.
//
// Solidity: event CancelUpgradeProposal(uint256 indexed proposalId, bytes32 indexed proposalHash)
func (_IZkSync *IZkSyncFilterer) FilterCancelUpgradeProposal(opts *bind.FilterOpts, proposalId []*big.Int, proposalHash [][32]byte) (*IZkSyncCancelUpgradeProposalIterator, error) {

	var proposalIdRule []interface{}
	for _, proposalIdItem := range proposalId {
		proposalIdRule = append(proposalIdRule, proposalIdItem)
	}
	var proposalHashRule []interface{}
	for _, proposalHashItem := range proposalHash {
		proposalHashRule = append(proposalHashRule, proposalHashItem)
	}

	logs, sub, err := _IZkSync.contract.FilterLogs(opts, "CancelUpgradeProposal", proposalIdRule, proposalHashRule)
	if err != nil {
		return nil, err
	}
	return &IZkSyncCancelUpgradeProposalIterator{contract: _IZkSync.contract, event: "CancelUpgradeProposal", logs: logs, sub: sub}, nil
}

// WatchCancelUpgradeProposal is a free log subscription operation binding the contract event 0x502f7e4ccbd6b35356f90b3cca05afdb6afea61cea0803229edf974a981c3030.
//
// Solidity: event CancelUpgradeProposal(uint256 indexed proposalId, bytes32 indexed proposalHash)
func (_IZkSync *IZkSyncFilterer) WatchCancelUpgradeProposal(opts *bind.WatchOpts, sink chan<- *IZkSyncCancelUpgradeProposal, proposalId []*big.Int, proposalHash [][32]byte) (event.Subscription, error) {

	var proposalIdRule []interface{}
	for _, proposalIdItem := range proposalId {
		proposalIdRule = append(proposalIdRule, proposalIdItem)
	}
	var proposalHashRule []interface{}
	for _, proposalHashItem := range proposalHash {
		proposalHashRule = append(proposalHashRule, proposalHashItem)
	}

	logs, sub, err := _IZkSync.contract.WatchLogs(opts, "CancelUpgradeProposal", proposalIdRule, proposalHashRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IZkSyncCancelUpgradeProposal)
				if err := _IZkSync.contract.UnpackLog(event, "CancelUpgradeProposal", log); err != nil {
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

// ParseCancelUpgradeProposal is a log parse operation binding the contract event 0x502f7e4ccbd6b35356f90b3cca05afdb6afea61cea0803229edf974a981c3030.
//
// Solidity: event CancelUpgradeProposal(uint256 indexed proposalId, bytes32 indexed proposalHash)
func (_IZkSync *IZkSyncFilterer) ParseCancelUpgradeProposal(log types.Log) (*IZkSyncCancelUpgradeProposal, error) {
	event := new(IZkSyncCancelUpgradeProposal)
	if err := _IZkSync.contract.UnpackLog(event, "CancelUpgradeProposal", log); err != nil {
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
	ProposalId   *big.Int
	ProposalHash [32]byte
	ProposalSalt [32]byte
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterExecuteUpgrade is a free log retrieval operation binding the contract event 0xf24fdb893d8f88e816b4de16cbe044c13dae94c460c753ef7fabee194aa30bbb.
//
// Solidity: event ExecuteUpgrade(uint256 indexed proposalId, bytes32 indexed proposalHash, bytes32 proposalSalt)
func (_IZkSync *IZkSyncFilterer) FilterExecuteUpgrade(opts *bind.FilterOpts, proposalId []*big.Int, proposalHash [][32]byte) (*IZkSyncExecuteUpgradeIterator, error) {

	var proposalIdRule []interface{}
	for _, proposalIdItem := range proposalId {
		proposalIdRule = append(proposalIdRule, proposalIdItem)
	}
	var proposalHashRule []interface{}
	for _, proposalHashItem := range proposalHash {
		proposalHashRule = append(proposalHashRule, proposalHashItem)
	}

	logs, sub, err := _IZkSync.contract.FilterLogs(opts, "ExecuteUpgrade", proposalIdRule, proposalHashRule)
	if err != nil {
		return nil, err
	}
	return &IZkSyncExecuteUpgradeIterator{contract: _IZkSync.contract, event: "ExecuteUpgrade", logs: logs, sub: sub}, nil
}

// WatchExecuteUpgrade is a free log subscription operation binding the contract event 0xf24fdb893d8f88e816b4de16cbe044c13dae94c460c753ef7fabee194aa30bbb.
//
// Solidity: event ExecuteUpgrade(uint256 indexed proposalId, bytes32 indexed proposalHash, bytes32 proposalSalt)
func (_IZkSync *IZkSyncFilterer) WatchExecuteUpgrade(opts *bind.WatchOpts, sink chan<- *IZkSyncExecuteUpgrade, proposalId []*big.Int, proposalHash [][32]byte) (event.Subscription, error) {

	var proposalIdRule []interface{}
	for _, proposalIdItem := range proposalId {
		proposalIdRule = append(proposalIdRule, proposalIdItem)
	}
	var proposalHashRule []interface{}
	for _, proposalHashItem := range proposalHash {
		proposalHashRule = append(proposalHashRule, proposalHashItem)
	}

	logs, sub, err := _IZkSync.contract.WatchLogs(opts, "ExecuteUpgrade", proposalIdRule, proposalHashRule)
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

// ParseExecuteUpgrade is a log parse operation binding the contract event 0xf24fdb893d8f88e816b4de16cbe044c13dae94c460c753ef7fabee194aa30bbb.
//
// Solidity: event ExecuteUpgrade(uint256 indexed proposalId, bytes32 indexed proposalHash, bytes32 proposalSalt)
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

// IZkSyncNewAllowListIterator is returned from FilterNewAllowList and is used to iterate over the raw logs and unpacked data for NewAllowList events raised by the IZkSync contract.
type IZkSyncNewAllowListIterator struct {
	Event *IZkSyncNewAllowList // Event containing the contract specifics and raw log

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
func (it *IZkSyncNewAllowListIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IZkSyncNewAllowList)
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
		it.Event = new(IZkSyncNewAllowList)
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
func (it *IZkSyncNewAllowListIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IZkSyncNewAllowListIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IZkSyncNewAllowList represents a NewAllowList event raised by the IZkSync contract.
type IZkSyncNewAllowList struct {
	OldAllowList common.Address
	NewAllowList common.Address
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterNewAllowList is a free log retrieval operation binding the contract event 0xa9dcddb809b52807b673fa313949dd73df5b0e1a42afdc45fe4e8bc13f0ebceb.
//
// Solidity: event NewAllowList(address indexed oldAllowList, address indexed newAllowList)
func (_IZkSync *IZkSyncFilterer) FilterNewAllowList(opts *bind.FilterOpts, oldAllowList []common.Address, newAllowList []common.Address) (*IZkSyncNewAllowListIterator, error) {

	var oldAllowListRule []interface{}
	for _, oldAllowListItem := range oldAllowList {
		oldAllowListRule = append(oldAllowListRule, oldAllowListItem)
	}
	var newAllowListRule []interface{}
	for _, newAllowListItem := range newAllowList {
		newAllowListRule = append(newAllowListRule, newAllowListItem)
	}

	logs, sub, err := _IZkSync.contract.FilterLogs(opts, "NewAllowList", oldAllowListRule, newAllowListRule)
	if err != nil {
		return nil, err
	}
	return &IZkSyncNewAllowListIterator{contract: _IZkSync.contract, event: "NewAllowList", logs: logs, sub: sub}, nil
}

// WatchNewAllowList is a free log subscription operation binding the contract event 0xa9dcddb809b52807b673fa313949dd73df5b0e1a42afdc45fe4e8bc13f0ebceb.
//
// Solidity: event NewAllowList(address indexed oldAllowList, address indexed newAllowList)
func (_IZkSync *IZkSyncFilterer) WatchNewAllowList(opts *bind.WatchOpts, sink chan<- *IZkSyncNewAllowList, oldAllowList []common.Address, newAllowList []common.Address) (event.Subscription, error) {

	var oldAllowListRule []interface{}
	for _, oldAllowListItem := range oldAllowList {
		oldAllowListRule = append(oldAllowListRule, oldAllowListItem)
	}
	var newAllowListRule []interface{}
	for _, newAllowListItem := range newAllowList {
		newAllowListRule = append(newAllowListRule, newAllowListItem)
	}

	logs, sub, err := _IZkSync.contract.WatchLogs(opts, "NewAllowList", oldAllowListRule, newAllowListRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IZkSyncNewAllowList)
				if err := _IZkSync.contract.UnpackLog(event, "NewAllowList", log); err != nil {
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

// ParseNewAllowList is a log parse operation binding the contract event 0xa9dcddb809b52807b673fa313949dd73df5b0e1a42afdc45fe4e8bc13f0ebceb.
//
// Solidity: event NewAllowList(address indexed oldAllowList, address indexed newAllowList)
func (_IZkSync *IZkSyncFilterer) ParseNewAllowList(log types.Log) (*IZkSyncNewAllowList, error) {
	event := new(IZkSyncNewAllowList)
	if err := _IZkSync.contract.UnpackLog(event, "NewAllowList", log); err != nil {
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

// IZkSyncNewL2BootloaderBytecodeHashIterator is returned from FilterNewL2BootloaderBytecodeHash and is used to iterate over the raw logs and unpacked data for NewL2BootloaderBytecodeHash events raised by the IZkSync contract.
type IZkSyncNewL2BootloaderBytecodeHashIterator struct {
	Event *IZkSyncNewL2BootloaderBytecodeHash // Event containing the contract specifics and raw log

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
func (it *IZkSyncNewL2BootloaderBytecodeHashIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IZkSyncNewL2BootloaderBytecodeHash)
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
		it.Event = new(IZkSyncNewL2BootloaderBytecodeHash)
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
func (it *IZkSyncNewL2BootloaderBytecodeHashIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IZkSyncNewL2BootloaderBytecodeHashIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IZkSyncNewL2BootloaderBytecodeHash represents a NewL2BootloaderBytecodeHash event raised by the IZkSync contract.
type IZkSyncNewL2BootloaderBytecodeHash struct {
	PreviousBytecodeHash [32]byte
	NewBytecodeHash      [32]byte
	Raw                  types.Log // Blockchain specific contextual infos
}

// FilterNewL2BootloaderBytecodeHash is a free log retrieval operation binding the contract event 0x271b33af94e3f065ecd8659833e6b1daf851f063700c36ddefefab35d4ce4746.
//
// Solidity: event NewL2BootloaderBytecodeHash(bytes32 indexed previousBytecodeHash, bytes32 indexed newBytecodeHash)
func (_IZkSync *IZkSyncFilterer) FilterNewL2BootloaderBytecodeHash(opts *bind.FilterOpts, previousBytecodeHash [][32]byte, newBytecodeHash [][32]byte) (*IZkSyncNewL2BootloaderBytecodeHashIterator, error) {

	var previousBytecodeHashRule []interface{}
	for _, previousBytecodeHashItem := range previousBytecodeHash {
		previousBytecodeHashRule = append(previousBytecodeHashRule, previousBytecodeHashItem)
	}
	var newBytecodeHashRule []interface{}
	for _, newBytecodeHashItem := range newBytecodeHash {
		newBytecodeHashRule = append(newBytecodeHashRule, newBytecodeHashItem)
	}

	logs, sub, err := _IZkSync.contract.FilterLogs(opts, "NewL2BootloaderBytecodeHash", previousBytecodeHashRule, newBytecodeHashRule)
	if err != nil {
		return nil, err
	}
	return &IZkSyncNewL2BootloaderBytecodeHashIterator{contract: _IZkSync.contract, event: "NewL2BootloaderBytecodeHash", logs: logs, sub: sub}, nil
}

// WatchNewL2BootloaderBytecodeHash is a free log subscription operation binding the contract event 0x271b33af94e3f065ecd8659833e6b1daf851f063700c36ddefefab35d4ce4746.
//
// Solidity: event NewL2BootloaderBytecodeHash(bytes32 indexed previousBytecodeHash, bytes32 indexed newBytecodeHash)
func (_IZkSync *IZkSyncFilterer) WatchNewL2BootloaderBytecodeHash(opts *bind.WatchOpts, sink chan<- *IZkSyncNewL2BootloaderBytecodeHash, previousBytecodeHash [][32]byte, newBytecodeHash [][32]byte) (event.Subscription, error) {

	var previousBytecodeHashRule []interface{}
	for _, previousBytecodeHashItem := range previousBytecodeHash {
		previousBytecodeHashRule = append(previousBytecodeHashRule, previousBytecodeHashItem)
	}
	var newBytecodeHashRule []interface{}
	for _, newBytecodeHashItem := range newBytecodeHash {
		newBytecodeHashRule = append(newBytecodeHashRule, newBytecodeHashItem)
	}

	logs, sub, err := _IZkSync.contract.WatchLogs(opts, "NewL2BootloaderBytecodeHash", previousBytecodeHashRule, newBytecodeHashRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IZkSyncNewL2BootloaderBytecodeHash)
				if err := _IZkSync.contract.UnpackLog(event, "NewL2BootloaderBytecodeHash", log); err != nil {
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

// ParseNewL2BootloaderBytecodeHash is a log parse operation binding the contract event 0x271b33af94e3f065ecd8659833e6b1daf851f063700c36ddefefab35d4ce4746.
//
// Solidity: event NewL2BootloaderBytecodeHash(bytes32 indexed previousBytecodeHash, bytes32 indexed newBytecodeHash)
func (_IZkSync *IZkSyncFilterer) ParseNewL2BootloaderBytecodeHash(log types.Log) (*IZkSyncNewL2BootloaderBytecodeHash, error) {
	event := new(IZkSyncNewL2BootloaderBytecodeHash)
	if err := _IZkSync.contract.UnpackLog(event, "NewL2BootloaderBytecodeHash", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// IZkSyncNewL2DefaultAccountBytecodeHashIterator is returned from FilterNewL2DefaultAccountBytecodeHash and is used to iterate over the raw logs and unpacked data for NewL2DefaultAccountBytecodeHash events raised by the IZkSync contract.
type IZkSyncNewL2DefaultAccountBytecodeHashIterator struct {
	Event *IZkSyncNewL2DefaultAccountBytecodeHash // Event containing the contract specifics and raw log

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
func (it *IZkSyncNewL2DefaultAccountBytecodeHashIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IZkSyncNewL2DefaultAccountBytecodeHash)
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
		it.Event = new(IZkSyncNewL2DefaultAccountBytecodeHash)
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
func (it *IZkSyncNewL2DefaultAccountBytecodeHashIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IZkSyncNewL2DefaultAccountBytecodeHashIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IZkSyncNewL2DefaultAccountBytecodeHash represents a NewL2DefaultAccountBytecodeHash event raised by the IZkSync contract.
type IZkSyncNewL2DefaultAccountBytecodeHash struct {
	PreviousBytecodeHash [32]byte
	NewBytecodeHash      [32]byte
	Raw                  types.Log // Blockchain specific contextual infos
}

// FilterNewL2DefaultAccountBytecodeHash is a free log retrieval operation binding the contract event 0x36df93a47cc02081d9d8208022ab736fdf98fac566e5fc6f5762bf7666e521f3.
//
// Solidity: event NewL2DefaultAccountBytecodeHash(bytes32 indexed previousBytecodeHash, bytes32 indexed newBytecodeHash)
func (_IZkSync *IZkSyncFilterer) FilterNewL2DefaultAccountBytecodeHash(opts *bind.FilterOpts, previousBytecodeHash [][32]byte, newBytecodeHash [][32]byte) (*IZkSyncNewL2DefaultAccountBytecodeHashIterator, error) {

	var previousBytecodeHashRule []interface{}
	for _, previousBytecodeHashItem := range previousBytecodeHash {
		previousBytecodeHashRule = append(previousBytecodeHashRule, previousBytecodeHashItem)
	}
	var newBytecodeHashRule []interface{}
	for _, newBytecodeHashItem := range newBytecodeHash {
		newBytecodeHashRule = append(newBytecodeHashRule, newBytecodeHashItem)
	}

	logs, sub, err := _IZkSync.contract.FilterLogs(opts, "NewL2DefaultAccountBytecodeHash", previousBytecodeHashRule, newBytecodeHashRule)
	if err != nil {
		return nil, err
	}
	return &IZkSyncNewL2DefaultAccountBytecodeHashIterator{contract: _IZkSync.contract, event: "NewL2DefaultAccountBytecodeHash", logs: logs, sub: sub}, nil
}

// WatchNewL2DefaultAccountBytecodeHash is a free log subscription operation binding the contract event 0x36df93a47cc02081d9d8208022ab736fdf98fac566e5fc6f5762bf7666e521f3.
//
// Solidity: event NewL2DefaultAccountBytecodeHash(bytes32 indexed previousBytecodeHash, bytes32 indexed newBytecodeHash)
func (_IZkSync *IZkSyncFilterer) WatchNewL2DefaultAccountBytecodeHash(opts *bind.WatchOpts, sink chan<- *IZkSyncNewL2DefaultAccountBytecodeHash, previousBytecodeHash [][32]byte, newBytecodeHash [][32]byte) (event.Subscription, error) {

	var previousBytecodeHashRule []interface{}
	for _, previousBytecodeHashItem := range previousBytecodeHash {
		previousBytecodeHashRule = append(previousBytecodeHashRule, previousBytecodeHashItem)
	}
	var newBytecodeHashRule []interface{}
	for _, newBytecodeHashItem := range newBytecodeHash {
		newBytecodeHashRule = append(newBytecodeHashRule, newBytecodeHashItem)
	}

	logs, sub, err := _IZkSync.contract.WatchLogs(opts, "NewL2DefaultAccountBytecodeHash", previousBytecodeHashRule, newBytecodeHashRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IZkSyncNewL2DefaultAccountBytecodeHash)
				if err := _IZkSync.contract.UnpackLog(event, "NewL2DefaultAccountBytecodeHash", log); err != nil {
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

// ParseNewL2DefaultAccountBytecodeHash is a log parse operation binding the contract event 0x36df93a47cc02081d9d8208022ab736fdf98fac566e5fc6f5762bf7666e521f3.
//
// Solidity: event NewL2DefaultAccountBytecodeHash(bytes32 indexed previousBytecodeHash, bytes32 indexed newBytecodeHash)
func (_IZkSync *IZkSyncFilterer) ParseNewL2DefaultAccountBytecodeHash(log types.Log) (*IZkSyncNewL2DefaultAccountBytecodeHash, error) {
	event := new(IZkSyncNewL2DefaultAccountBytecodeHash)
	if err := _IZkSync.contract.UnpackLog(event, "NewL2DefaultAccountBytecodeHash", log); err != nil {
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

// IZkSyncNewVerifierIterator is returned from FilterNewVerifier and is used to iterate over the raw logs and unpacked data for NewVerifier events raised by the IZkSync contract.
type IZkSyncNewVerifierIterator struct {
	Event *IZkSyncNewVerifier // Event containing the contract specifics and raw log

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
func (it *IZkSyncNewVerifierIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IZkSyncNewVerifier)
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
		it.Event = new(IZkSyncNewVerifier)
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
func (it *IZkSyncNewVerifierIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IZkSyncNewVerifierIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IZkSyncNewVerifier represents a NewVerifier event raised by the IZkSync contract.
type IZkSyncNewVerifier struct {
	OldVerifier common.Address
	NewVerifier common.Address
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterNewVerifier is a free log retrieval operation binding the contract event 0x2ff4895c300d6993c27f2bb507b4b59d29464dc640af727383451365631ba8b2.
//
// Solidity: event NewVerifier(address indexed oldVerifier, address indexed newVerifier)
func (_IZkSync *IZkSyncFilterer) FilterNewVerifier(opts *bind.FilterOpts, oldVerifier []common.Address, newVerifier []common.Address) (*IZkSyncNewVerifierIterator, error) {

	var oldVerifierRule []interface{}
	for _, oldVerifierItem := range oldVerifier {
		oldVerifierRule = append(oldVerifierRule, oldVerifierItem)
	}
	var newVerifierRule []interface{}
	for _, newVerifierItem := range newVerifier {
		newVerifierRule = append(newVerifierRule, newVerifierItem)
	}

	logs, sub, err := _IZkSync.contract.FilterLogs(opts, "NewVerifier", oldVerifierRule, newVerifierRule)
	if err != nil {
		return nil, err
	}
	return &IZkSyncNewVerifierIterator{contract: _IZkSync.contract, event: "NewVerifier", logs: logs, sub: sub}, nil
}

// WatchNewVerifier is a free log subscription operation binding the contract event 0x2ff4895c300d6993c27f2bb507b4b59d29464dc640af727383451365631ba8b2.
//
// Solidity: event NewVerifier(address indexed oldVerifier, address indexed newVerifier)
func (_IZkSync *IZkSyncFilterer) WatchNewVerifier(opts *bind.WatchOpts, sink chan<- *IZkSyncNewVerifier, oldVerifier []common.Address, newVerifier []common.Address) (event.Subscription, error) {

	var oldVerifierRule []interface{}
	for _, oldVerifierItem := range oldVerifier {
		oldVerifierRule = append(oldVerifierRule, oldVerifierItem)
	}
	var newVerifierRule []interface{}
	for _, newVerifierItem := range newVerifier {
		newVerifierRule = append(newVerifierRule, newVerifierItem)
	}

	logs, sub, err := _IZkSync.contract.WatchLogs(opts, "NewVerifier", oldVerifierRule, newVerifierRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IZkSyncNewVerifier)
				if err := _IZkSync.contract.UnpackLog(event, "NewVerifier", log); err != nil {
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

// ParseNewVerifier is a log parse operation binding the contract event 0x2ff4895c300d6993c27f2bb507b4b59d29464dc640af727383451365631ba8b2.
//
// Solidity: event NewVerifier(address indexed oldVerifier, address indexed newVerifier)
func (_IZkSync *IZkSyncFilterer) ParseNewVerifier(log types.Log) (*IZkSyncNewVerifier, error) {
	event := new(IZkSyncNewVerifier)
	if err := _IZkSync.contract.UnpackLog(event, "NewVerifier", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// IZkSyncNewVerifierParamsIterator is returned from FilterNewVerifierParams and is used to iterate over the raw logs and unpacked data for NewVerifierParams events raised by the IZkSync contract.
type IZkSyncNewVerifierParamsIterator struct {
	Event *IZkSyncNewVerifierParams // Event containing the contract specifics and raw log

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
func (it *IZkSyncNewVerifierParamsIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IZkSyncNewVerifierParams)
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
		it.Event = new(IZkSyncNewVerifierParams)
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
func (it *IZkSyncNewVerifierParamsIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IZkSyncNewVerifierParamsIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IZkSyncNewVerifierParams represents a NewVerifierParams event raised by the IZkSync contract.
type IZkSyncNewVerifierParams struct {
	OldVerifierParams VerifierParams
	NewVerifierParams VerifierParams
	Raw               types.Log // Blockchain specific contextual infos
}

// FilterNewVerifierParams is a free log retrieval operation binding the contract event 0x4c055dbc5f14dcb6e081c9421d9657d950dcd6372f6db0a714b9135171cbc15d.
//
// Solidity: event NewVerifierParams((bytes32,bytes32,bytes32) oldVerifierParams, (bytes32,bytes32,bytes32) newVerifierParams)
func (_IZkSync *IZkSyncFilterer) FilterNewVerifierParams(opts *bind.FilterOpts) (*IZkSyncNewVerifierParamsIterator, error) {

	logs, sub, err := _IZkSync.contract.FilterLogs(opts, "NewVerifierParams")
	if err != nil {
		return nil, err
	}
	return &IZkSyncNewVerifierParamsIterator{contract: _IZkSync.contract, event: "NewVerifierParams", logs: logs, sub: sub}, nil
}

// WatchNewVerifierParams is a free log subscription operation binding the contract event 0x4c055dbc5f14dcb6e081c9421d9657d950dcd6372f6db0a714b9135171cbc15d.
//
// Solidity: event NewVerifierParams((bytes32,bytes32,bytes32) oldVerifierParams, (bytes32,bytes32,bytes32) newVerifierParams)
func (_IZkSync *IZkSyncFilterer) WatchNewVerifierParams(opts *bind.WatchOpts, sink chan<- *IZkSyncNewVerifierParams) (event.Subscription, error) {

	logs, sub, err := _IZkSync.contract.WatchLogs(opts, "NewVerifierParams")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IZkSyncNewVerifierParams)
				if err := _IZkSync.contract.UnpackLog(event, "NewVerifierParams", log); err != nil {
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

// ParseNewVerifierParams is a log parse operation binding the contract event 0x4c055dbc5f14dcb6e081c9421d9657d950dcd6372f6db0a714b9135171cbc15d.
//
// Solidity: event NewVerifierParams((bytes32,bytes32,bytes32) oldVerifierParams, (bytes32,bytes32,bytes32) newVerifierParams)
func (_IZkSync *IZkSyncFilterer) ParseNewVerifierParams(log types.Log) (*IZkSyncNewVerifierParams, error) {
	event := new(IZkSyncNewVerifierParams)
	if err := _IZkSync.contract.UnpackLog(event, "NewVerifierParams", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// IZkSyncProposeShadowUpgradeIterator is returned from FilterProposeShadowUpgrade and is used to iterate over the raw logs and unpacked data for ProposeShadowUpgrade events raised by the IZkSync contract.
type IZkSyncProposeShadowUpgradeIterator struct {
	Event *IZkSyncProposeShadowUpgrade // Event containing the contract specifics and raw log

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
func (it *IZkSyncProposeShadowUpgradeIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IZkSyncProposeShadowUpgrade)
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
		it.Event = new(IZkSyncProposeShadowUpgrade)
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
func (it *IZkSyncProposeShadowUpgradeIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IZkSyncProposeShadowUpgradeIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IZkSyncProposeShadowUpgrade represents a ProposeShadowUpgrade event raised by the IZkSync contract.
type IZkSyncProposeShadowUpgrade struct {
	ProposalId   *big.Int
	ProposalHash [32]byte
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterProposeShadowUpgrade is a free log retrieval operation binding the contract event 0x0ba6a8e70bc3b2350bd5159eb57c31da5dfb6dc65e99421498dd536b601bbd11.
//
// Solidity: event ProposeShadowUpgrade(uint256 indexed proposalId, bytes32 indexed proposalHash)
func (_IZkSync *IZkSyncFilterer) FilterProposeShadowUpgrade(opts *bind.FilterOpts, proposalId []*big.Int, proposalHash [][32]byte) (*IZkSyncProposeShadowUpgradeIterator, error) {

	var proposalIdRule []interface{}
	for _, proposalIdItem := range proposalId {
		proposalIdRule = append(proposalIdRule, proposalIdItem)
	}
	var proposalHashRule []interface{}
	for _, proposalHashItem := range proposalHash {
		proposalHashRule = append(proposalHashRule, proposalHashItem)
	}

	logs, sub, err := _IZkSync.contract.FilterLogs(opts, "ProposeShadowUpgrade", proposalIdRule, proposalHashRule)
	if err != nil {
		return nil, err
	}
	return &IZkSyncProposeShadowUpgradeIterator{contract: _IZkSync.contract, event: "ProposeShadowUpgrade", logs: logs, sub: sub}, nil
}

// WatchProposeShadowUpgrade is a free log subscription operation binding the contract event 0x0ba6a8e70bc3b2350bd5159eb57c31da5dfb6dc65e99421498dd536b601bbd11.
//
// Solidity: event ProposeShadowUpgrade(uint256 indexed proposalId, bytes32 indexed proposalHash)
func (_IZkSync *IZkSyncFilterer) WatchProposeShadowUpgrade(opts *bind.WatchOpts, sink chan<- *IZkSyncProposeShadowUpgrade, proposalId []*big.Int, proposalHash [][32]byte) (event.Subscription, error) {

	var proposalIdRule []interface{}
	for _, proposalIdItem := range proposalId {
		proposalIdRule = append(proposalIdRule, proposalIdItem)
	}
	var proposalHashRule []interface{}
	for _, proposalHashItem := range proposalHash {
		proposalHashRule = append(proposalHashRule, proposalHashItem)
	}

	logs, sub, err := _IZkSync.contract.WatchLogs(opts, "ProposeShadowUpgrade", proposalIdRule, proposalHashRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IZkSyncProposeShadowUpgrade)
				if err := _IZkSync.contract.UnpackLog(event, "ProposeShadowUpgrade", log); err != nil {
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

// ParseProposeShadowUpgrade is a log parse operation binding the contract event 0x0ba6a8e70bc3b2350bd5159eb57c31da5dfb6dc65e99421498dd536b601bbd11.
//
// Solidity: event ProposeShadowUpgrade(uint256 indexed proposalId, bytes32 indexed proposalHash)
func (_IZkSync *IZkSyncFilterer) ParseProposeShadowUpgrade(log types.Log) (*IZkSyncProposeShadowUpgrade, error) {
	event := new(IZkSyncProposeShadowUpgrade)
	if err := _IZkSync.contract.UnpackLog(event, "ProposeShadowUpgrade", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// IZkSyncProposeTransparentUpgradeIterator is returned from FilterProposeTransparentUpgrade and is used to iterate over the raw logs and unpacked data for ProposeTransparentUpgrade events raised by the IZkSync contract.
type IZkSyncProposeTransparentUpgradeIterator struct {
	Event *IZkSyncProposeTransparentUpgrade // Event containing the contract specifics and raw log

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
func (it *IZkSyncProposeTransparentUpgradeIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IZkSyncProposeTransparentUpgrade)
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
		it.Event = new(IZkSyncProposeTransparentUpgrade)
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
func (it *IZkSyncProposeTransparentUpgradeIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IZkSyncProposeTransparentUpgradeIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IZkSyncProposeTransparentUpgrade represents a ProposeTransparentUpgrade event raised by the IZkSync contract.
type IZkSyncProposeTransparentUpgrade struct {
	DiamondCut   DiamondDiamondCutData
	ProposalId   *big.Int
	ProposalSalt [32]byte
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterProposeTransparentUpgrade is a free log retrieval operation binding the contract event 0x69115b49afe7a6101a2e7af17d421eda1dc153bd26d699f013c4fff0404646a6.
//
// Solidity: event ProposeTransparentUpgrade(((address,uint8,bool,bytes4[])[],address,bytes) diamondCut, uint256 indexed proposalId, bytes32 proposalSalt)
func (_IZkSync *IZkSyncFilterer) FilterProposeTransparentUpgrade(opts *bind.FilterOpts, proposalId []*big.Int) (*IZkSyncProposeTransparentUpgradeIterator, error) {

	var proposalIdRule []interface{}
	for _, proposalIdItem := range proposalId {
		proposalIdRule = append(proposalIdRule, proposalIdItem)
	}

	logs, sub, err := _IZkSync.contract.FilterLogs(opts, "ProposeTransparentUpgrade", proposalIdRule)
	if err != nil {
		return nil, err
	}
	return &IZkSyncProposeTransparentUpgradeIterator{contract: _IZkSync.contract, event: "ProposeTransparentUpgrade", logs: logs, sub: sub}, nil
}

// WatchProposeTransparentUpgrade is a free log subscription operation binding the contract event 0x69115b49afe7a6101a2e7af17d421eda1dc153bd26d699f013c4fff0404646a6.
//
// Solidity: event ProposeTransparentUpgrade(((address,uint8,bool,bytes4[])[],address,bytes) diamondCut, uint256 indexed proposalId, bytes32 proposalSalt)
func (_IZkSync *IZkSyncFilterer) WatchProposeTransparentUpgrade(opts *bind.WatchOpts, sink chan<- *IZkSyncProposeTransparentUpgrade, proposalId []*big.Int) (event.Subscription, error) {

	var proposalIdRule []interface{}
	for _, proposalIdItem := range proposalId {
		proposalIdRule = append(proposalIdRule, proposalIdItem)
	}

	logs, sub, err := _IZkSync.contract.WatchLogs(opts, "ProposeTransparentUpgrade", proposalIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IZkSyncProposeTransparentUpgrade)
				if err := _IZkSync.contract.UnpackLog(event, "ProposeTransparentUpgrade", log); err != nil {
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
func (_IZkSync *IZkSyncFilterer) ParseProposeTransparentUpgrade(log types.Log) (*IZkSyncProposeTransparentUpgrade, error) {
	event := new(IZkSyncProposeTransparentUpgrade)
	if err := _IZkSync.contract.UnpackLog(event, "ProposeTransparentUpgrade", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// IZkSyncSecurityCouncilUpgradeApproveIterator is returned from FilterSecurityCouncilUpgradeApprove and is used to iterate over the raw logs and unpacked data for SecurityCouncilUpgradeApprove events raised by the IZkSync contract.
type IZkSyncSecurityCouncilUpgradeApproveIterator struct {
	Event *IZkSyncSecurityCouncilUpgradeApprove // Event containing the contract specifics and raw log

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
func (it *IZkSyncSecurityCouncilUpgradeApproveIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IZkSyncSecurityCouncilUpgradeApprove)
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
		it.Event = new(IZkSyncSecurityCouncilUpgradeApprove)
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
func (it *IZkSyncSecurityCouncilUpgradeApproveIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IZkSyncSecurityCouncilUpgradeApproveIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IZkSyncSecurityCouncilUpgradeApprove represents a SecurityCouncilUpgradeApprove event raised by the IZkSync contract.
type IZkSyncSecurityCouncilUpgradeApprove struct {
	ProposalId   *big.Int
	ProposalHash [32]byte
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterSecurityCouncilUpgradeApprove is a free log retrieval operation binding the contract event 0xc1e63a10a7f348f746b83edbe62a9ad09a81c7090c0813318c782cfe0f7a5a3e.
//
// Solidity: event SecurityCouncilUpgradeApprove(uint256 indexed proposalId, bytes32 indexed proposalHash)
func (_IZkSync *IZkSyncFilterer) FilterSecurityCouncilUpgradeApprove(opts *bind.FilterOpts, proposalId []*big.Int, proposalHash [][32]byte) (*IZkSyncSecurityCouncilUpgradeApproveIterator, error) {

	var proposalIdRule []interface{}
	for _, proposalIdItem := range proposalId {
		proposalIdRule = append(proposalIdRule, proposalIdItem)
	}
	var proposalHashRule []interface{}
	for _, proposalHashItem := range proposalHash {
		proposalHashRule = append(proposalHashRule, proposalHashItem)
	}

	logs, sub, err := _IZkSync.contract.FilterLogs(opts, "SecurityCouncilUpgradeApprove", proposalIdRule, proposalHashRule)
	if err != nil {
		return nil, err
	}
	return &IZkSyncSecurityCouncilUpgradeApproveIterator{contract: _IZkSync.contract, event: "SecurityCouncilUpgradeApprove", logs: logs, sub: sub}, nil
}

// WatchSecurityCouncilUpgradeApprove is a free log subscription operation binding the contract event 0xc1e63a10a7f348f746b83edbe62a9ad09a81c7090c0813318c782cfe0f7a5a3e.
//
// Solidity: event SecurityCouncilUpgradeApprove(uint256 indexed proposalId, bytes32 indexed proposalHash)
func (_IZkSync *IZkSyncFilterer) WatchSecurityCouncilUpgradeApprove(opts *bind.WatchOpts, sink chan<- *IZkSyncSecurityCouncilUpgradeApprove, proposalId []*big.Int, proposalHash [][32]byte) (event.Subscription, error) {

	var proposalIdRule []interface{}
	for _, proposalIdItem := range proposalId {
		proposalIdRule = append(proposalIdRule, proposalIdItem)
	}
	var proposalHashRule []interface{}
	for _, proposalHashItem := range proposalHash {
		proposalHashRule = append(proposalHashRule, proposalHashItem)
	}

	logs, sub, err := _IZkSync.contract.WatchLogs(opts, "SecurityCouncilUpgradeApprove", proposalIdRule, proposalHashRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IZkSyncSecurityCouncilUpgradeApprove)
				if err := _IZkSync.contract.UnpackLog(event, "SecurityCouncilUpgradeApprove", log); err != nil {
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

// ParseSecurityCouncilUpgradeApprove is a log parse operation binding the contract event 0xc1e63a10a7f348f746b83edbe62a9ad09a81c7090c0813318c782cfe0f7a5a3e.
//
// Solidity: event SecurityCouncilUpgradeApprove(uint256 indexed proposalId, bytes32 indexed proposalHash)
func (_IZkSync *IZkSyncFilterer) ParseSecurityCouncilUpgradeApprove(log types.Log) (*IZkSyncSecurityCouncilUpgradeApprove, error) {
	event := new(IZkSyncSecurityCouncilUpgradeApprove)
	if err := _IZkSync.contract.UnpackLog(event, "SecurityCouncilUpgradeApprove", log); err != nil {
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
