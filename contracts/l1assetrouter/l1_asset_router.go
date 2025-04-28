// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package l1assetrouter

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

// L2TransactionRequestTwoBridgesInner is an auto generated low-level Go binding around an user-defined struct.
type L2TransactionRequestTwoBridgesInner struct {
	MagicValue  [32]byte
	L2Contract  common.Address
	L2Calldata  []byte
	FactoryDeps [][]byte
	TxDataHash  [32]byte
}

// IL1AssetRouterMetaData contains all meta data concerning the IL1AssetRouter contract.
var IL1AssetRouterMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_l1WethAddress\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_bridgehub\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_l1Nullifier\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_eraChainId\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"_eraDiamondProxy\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"addr\",\"type\":\"address\"}],\"name\":\"AddressAlreadyUsed\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"assetId\",\"type\":\"bytes32\"}],\"name\":\"AssetHandlerDoesNotExist\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"assetId\",\"type\":\"bytes32\"}],\"name\":\"AssetIdNotSupported\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NotInitializedReentrancyGuard\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"Reentrancy\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"SlotOccupied\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"}],\"name\":\"TokenNotSupported\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"caller\",\"type\":\"address\"}],\"name\":\"Unauthorized\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"UnsupportedEncodingVersion\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ZeroAddress\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"assetId\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"assetDeploymentTracker\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"additionalData\",\"type\":\"bytes32\"}],\"name\":\"AssetDeploymentTrackerSet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"assetId\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"_assetAddress\",\"type\":\"address\"}],\"name\":\"AssetHandlerRegistered\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"assetId\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"assetHandlerAddress\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"additionalData\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"assetDeploymentTracker\",\"type\":\"address\"}],\"name\":\"AssetHandlerRegisteredInitial\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"chainId\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"assetId\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"BridgehubDepositBaseTokenInitiated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"chainId\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"txDataHash\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"l2DepositTxHash\",\"type\":\"bytes32\"}],\"name\":\"BridgehubDepositFinalized\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"chainId\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"txDataHash\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"assetId\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"bridgeMintCalldata\",\"type\":\"bytes\"}],\"name\":\"BridgehubDepositInitiated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"bridgeMintData\",\"type\":\"bytes\"}],\"name\":\"BridgehubMintData\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"chainId\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"assetId\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"assetDataHash\",\"type\":\"bytes32\"}],\"name\":\"BridgehubWithdrawalInitiated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"chainId\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"assetId\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"assetData\",\"type\":\"bytes\"}],\"name\":\"ClaimedFailedDepositAssetRouter\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"chainId\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"assetId\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"assetData\",\"type\":\"bytes\"}],\"name\":\"DepositFinalizedAssetRouter\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint8\",\"name\":\"version\",\"type\":\"uint8\"}],\"name\":\"Initialized\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"chainId\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"l2DepositTxHash\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"l1Asset\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"LegacyDepositInitiated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferStarted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"Paused\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"Unpaused\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"BRIDGE_HUB\",\"outputs\":[{\"internalType\":\"contractIBridgehub\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"ERA_CHAIN_ID\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"L1_CHAIN_ID\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"L1_NULLIFIER\",\"outputs\":[{\"internalType\":\"contractIL1Nullifier\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"L1_WETH_TOKEN\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"acceptOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"assetId\",\"type\":\"bytes32\"}],\"name\":\"assetDeploymentTracker\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"assetDeploymentTracker\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"assetId\",\"type\":\"bytes32\"}],\"name\":\"assetHandlerAddress\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"assetHandlerAddress\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_chainId\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"_depositSender\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"_assetId\",\"type\":\"bytes32\"},{\"internalType\":\"bytes\",\"name\":\"_assetData\",\"type\":\"bytes\"}],\"name\":\"bridgeRecoverFailedTransfer\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_chainId\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"_depositSender\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"_assetId\",\"type\":\"bytes32\"},{\"internalType\":\"bytes\",\"name\":\"_assetData\",\"type\":\"bytes\"},{\"internalType\":\"bytes32\",\"name\":\"_l2TxHash\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"_l2BatchNumber\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_l2MessageIndex\",\"type\":\"uint256\"},{\"internalType\":\"uint16\",\"name\":\"_l2TxNumberInBatch\",\"type\":\"uint16\"},{\"internalType\":\"bytes32[]\",\"name\":\"_merkleProof\",\"type\":\"bytes32[]\"}],\"name\":\"bridgeRecoverFailedTransfer\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_chainId\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"_txDataHash\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"_txHash\",\"type\":\"bytes32\"}],\"name\":\"bridgehubConfirmL2Transaction\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_chainId\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"_originalCaller\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_value\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"_data\",\"type\":\"bytes\"}],\"name\":\"bridgehubDeposit\",\"outputs\":[{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"magicValue\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"l2Contract\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"l2Calldata\",\"type\":\"bytes\"},{\"internalType\":\"bytes[]\",\"name\":\"factoryDeps\",\"type\":\"bytes[]\"},{\"internalType\":\"bytes32\",\"name\":\"txDataHash\",\"type\":\"bytes32\"}],\"internalType\":\"structL2TransactionRequestTwoBridgesInner\",\"name\":\"request\",\"type\":\"tuple\"}],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_chainId\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"_assetId\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"_originalCaller\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"bridgehubDepositBaseToken\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_chainId\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"_depositSender\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_l1Token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"_l2TxHash\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"_l2BatchNumber\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_l2MessageIndex\",\"type\":\"uint256\"},{\"internalType\":\"uint16\",\"name\":\"_l2TxNumberInBatch\",\"type\":\"uint16\"},{\"internalType\":\"bytes32[]\",\"name\":\"_merkleProof\",\"type\":\"bytes32[]\"}],\"name\":\"claimFailedDeposit\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_originalCaller\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_l2Receiver\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_l1Token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_l2TxGasLimit\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_l2TxGasPerPubdataByte\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"_refundRecipient\",\"type\":\"address\"}],\"name\":\"depositLegacyErc20Bridge\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"txHash\",\"type\":\"bytes32\"}],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_chainId\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"_assetId\",\"type\":\"bytes32\"},{\"internalType\":\"bytes\",\"name\":\"_transferData\",\"type\":\"bytes\"}],\"name\":\"finalizeDeposit\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_chainId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_l2BatchNumber\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_l2MessageIndex\",\"type\":\"uint256\"},{\"internalType\":\"uint16\",\"name\":\"_l2TxNumberInBatch\",\"type\":\"uint16\"},{\"internalType\":\"bytes\",\"name\":\"_message\",\"type\":\"bytes\"},{\"internalType\":\"bytes32[]\",\"name\":\"_merkleProof\",\"type\":\"bytes32[]\"}],\"name\":\"finalizeWithdrawal\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_sender\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"_assetId\",\"type\":\"bytes32\"},{\"internalType\":\"bytes\",\"name\":\"_assetData\",\"type\":\"bytes\"}],\"name\":\"getDepositCalldata\",\"outputs\":[{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_owner\",\"type\":\"address\"}],\"name\":\"initialize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_chainId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_l2BatchNumber\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_l2MessageIndex\",\"type\":\"uint256\"}],\"name\":\"isWithdrawalFinalized\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"legacyBridge\",\"outputs\":[{\"internalType\":\"contractIL1ERC20Bridge\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"nativeTokenVault\",\"outputs\":[{\"internalType\":\"contractINativeTokenVault\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"pause\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"paused\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"pendingOwner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_assetRegistrationData\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"_assetDeploymentTracker\",\"type\":\"address\"}],\"name\":\"setAssetDeploymentTracker\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_assetRegistrationData\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"_assetHandlerAddress\",\"type\":\"address\"}],\"name\":\"setAssetHandlerAddressThisChain\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractIL1ERC20Bridge\",\"name\":\"_legacyBridge\",\"type\":\"address\"}],\"name\":\"setL1Erc20Bridge\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractINativeTokenVault\",\"name\":\"_nativeTokenVault\",\"type\":\"address\"}],\"name\":\"setNativeTokenVault\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_assetId\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"_originalCaller\",\"type\":\"address\"}],\"name\":\"transferFundsToNTV\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"unpause\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
}

// IL1AssetRouterABI is the input ABI used to generate the binding from.
// Deprecated: Use IL1AssetRouterMetaData.ABI instead.
var IL1AssetRouterABI = IL1AssetRouterMetaData.ABI

// IL1AssetRouter is an auto generated Go binding around an Ethereum contract.
type IL1AssetRouter struct {
	IL1AssetRouterCaller     // Read-only binding to the contract
	IL1AssetRouterTransactor // Write-only binding to the contract
	IL1AssetRouterFilterer   // Log filterer for contract events
}

// IL1AssetRouterCaller is an auto generated read-only Go binding around an Ethereum contract.
type IL1AssetRouterCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IL1AssetRouterTransactor is an auto generated write-only Go binding around an Ethereum contract.
type IL1AssetRouterTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IL1AssetRouterFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type IL1AssetRouterFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IL1AssetRouterSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type IL1AssetRouterSession struct {
	Contract     *IL1AssetRouter   // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// IL1AssetRouterCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type IL1AssetRouterCallerSession struct {
	Contract *IL1AssetRouterCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts         // Call options to use throughout this session
}

// IL1AssetRouterTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type IL1AssetRouterTransactorSession struct {
	Contract     *IL1AssetRouterTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts         // Transaction auth options to use throughout this session
}

// IL1AssetRouterRaw is an auto generated low-level Go binding around an Ethereum contract.
type IL1AssetRouterRaw struct {
	Contract *IL1AssetRouter // Generic contract binding to access the raw methods on
}

// IL1AssetRouterCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type IL1AssetRouterCallerRaw struct {
	Contract *IL1AssetRouterCaller // Generic read-only contract binding to access the raw methods on
}

// IL1AssetRouterTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type IL1AssetRouterTransactorRaw struct {
	Contract *IL1AssetRouterTransactor // Generic write-only contract binding to access the raw methods on
}

// NewIL1AssetRouter creates a new instance of IL1AssetRouter, bound to a specific deployed contract.
func NewIL1AssetRouter(address common.Address, backend bind.ContractBackend) (*IL1AssetRouter, error) {
	contract, err := bindIL1AssetRouter(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &IL1AssetRouter{IL1AssetRouterCaller: IL1AssetRouterCaller{contract: contract}, IL1AssetRouterTransactor: IL1AssetRouterTransactor{contract: contract}, IL1AssetRouterFilterer: IL1AssetRouterFilterer{contract: contract}}, nil
}

// NewIL1AssetRouterCaller creates a new read-only instance of IL1AssetRouter, bound to a specific deployed contract.
func NewIL1AssetRouterCaller(address common.Address, caller bind.ContractCaller) (*IL1AssetRouterCaller, error) {
	contract, err := bindIL1AssetRouter(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &IL1AssetRouterCaller{contract: contract}, nil
}

// NewIL1AssetRouterTransactor creates a new write-only instance of IL1AssetRouter, bound to a specific deployed contract.
func NewIL1AssetRouterTransactor(address common.Address, transactor bind.ContractTransactor) (*IL1AssetRouterTransactor, error) {
	contract, err := bindIL1AssetRouter(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &IL1AssetRouterTransactor{contract: contract}, nil
}

// NewIL1AssetRouterFilterer creates a new log filterer instance of IL1AssetRouter, bound to a specific deployed contract.
func NewIL1AssetRouterFilterer(address common.Address, filterer bind.ContractFilterer) (*IL1AssetRouterFilterer, error) {
	contract, err := bindIL1AssetRouter(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &IL1AssetRouterFilterer{contract: contract}, nil
}

// bindIL1AssetRouter binds a generic wrapper to an already deployed contract.
func bindIL1AssetRouter(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := IL1AssetRouterMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IL1AssetRouter *IL1AssetRouterRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IL1AssetRouter.Contract.IL1AssetRouterCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IL1AssetRouter *IL1AssetRouterRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IL1AssetRouter.Contract.IL1AssetRouterTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IL1AssetRouter *IL1AssetRouterRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IL1AssetRouter.Contract.IL1AssetRouterTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IL1AssetRouter *IL1AssetRouterCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IL1AssetRouter.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IL1AssetRouter *IL1AssetRouterTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IL1AssetRouter.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IL1AssetRouter *IL1AssetRouterTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IL1AssetRouter.Contract.contract.Transact(opts, method, params...)
}

// BRIDGEHUB is a free data retrieval call binding the contract method 0x5d4edca7.
//
// Solidity: function BRIDGE_HUB() view returns(address)
func (_IL1AssetRouter *IL1AssetRouterCaller) BRIDGEHUB(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _IL1AssetRouter.contract.Call(opts, &out, "BRIDGE_HUB")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// BRIDGEHUB is a free data retrieval call binding the contract method 0x5d4edca7.
//
// Solidity: function BRIDGE_HUB() view returns(address)
func (_IL1AssetRouter *IL1AssetRouterSession) BRIDGEHUB() (common.Address, error) {
	return _IL1AssetRouter.Contract.BRIDGEHUB(&_IL1AssetRouter.CallOpts)
}

// BRIDGEHUB is a free data retrieval call binding the contract method 0x5d4edca7.
//
// Solidity: function BRIDGE_HUB() view returns(address)
func (_IL1AssetRouter *IL1AssetRouterCallerSession) BRIDGEHUB() (common.Address, error) {
	return _IL1AssetRouter.Contract.BRIDGEHUB(&_IL1AssetRouter.CallOpts)
}

// ERACHAINID is a free data retrieval call binding the contract method 0xef011dff.
//
// Solidity: function ERA_CHAIN_ID() view returns(uint256)
func (_IL1AssetRouter *IL1AssetRouterCaller) ERACHAINID(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _IL1AssetRouter.contract.Call(opts, &out, "ERA_CHAIN_ID")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// ERACHAINID is a free data retrieval call binding the contract method 0xef011dff.
//
// Solidity: function ERA_CHAIN_ID() view returns(uint256)
func (_IL1AssetRouter *IL1AssetRouterSession) ERACHAINID() (*big.Int, error) {
	return _IL1AssetRouter.Contract.ERACHAINID(&_IL1AssetRouter.CallOpts)
}

// ERACHAINID is a free data retrieval call binding the contract method 0xef011dff.
//
// Solidity: function ERA_CHAIN_ID() view returns(uint256)
func (_IL1AssetRouter *IL1AssetRouterCallerSession) ERACHAINID() (*big.Int, error) {
	return _IL1AssetRouter.Contract.ERACHAINID(&_IL1AssetRouter.CallOpts)
}

// L1CHAINID is a free data retrieval call binding the contract method 0x2f90b184.
//
// Solidity: function L1_CHAIN_ID() view returns(uint256)
func (_IL1AssetRouter *IL1AssetRouterCaller) L1CHAINID(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _IL1AssetRouter.contract.Call(opts, &out, "L1_CHAIN_ID")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// L1CHAINID is a free data retrieval call binding the contract method 0x2f90b184.
//
// Solidity: function L1_CHAIN_ID() view returns(uint256)
func (_IL1AssetRouter *IL1AssetRouterSession) L1CHAINID() (*big.Int, error) {
	return _IL1AssetRouter.Contract.L1CHAINID(&_IL1AssetRouter.CallOpts)
}

// L1CHAINID is a free data retrieval call binding the contract method 0x2f90b184.
//
// Solidity: function L1_CHAIN_ID() view returns(uint256)
func (_IL1AssetRouter *IL1AssetRouterCallerSession) L1CHAINID() (*big.Int, error) {
	return _IL1AssetRouter.Contract.L1CHAINID(&_IL1AssetRouter.CallOpts)
}

// L1NULLIFIER is a free data retrieval call binding the contract method 0xe60ccaba.
//
// Solidity: function L1_NULLIFIER() view returns(address)
func (_IL1AssetRouter *IL1AssetRouterCaller) L1NULLIFIER(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _IL1AssetRouter.contract.Call(opts, &out, "L1_NULLIFIER")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// L1NULLIFIER is a free data retrieval call binding the contract method 0xe60ccaba.
//
// Solidity: function L1_NULLIFIER() view returns(address)
func (_IL1AssetRouter *IL1AssetRouterSession) L1NULLIFIER() (common.Address, error) {
	return _IL1AssetRouter.Contract.L1NULLIFIER(&_IL1AssetRouter.CallOpts)
}

// L1NULLIFIER is a free data retrieval call binding the contract method 0xe60ccaba.
//
// Solidity: function L1_NULLIFIER() view returns(address)
func (_IL1AssetRouter *IL1AssetRouterCallerSession) L1NULLIFIER() (common.Address, error) {
	return _IL1AssetRouter.Contract.L1NULLIFIER(&_IL1AssetRouter.CallOpts)
}

// L1WETHTOKEN is a free data retrieval call binding the contract method 0x41c841c3.
//
// Solidity: function L1_WETH_TOKEN() view returns(address)
func (_IL1AssetRouter *IL1AssetRouterCaller) L1WETHTOKEN(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _IL1AssetRouter.contract.Call(opts, &out, "L1_WETH_TOKEN")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// L1WETHTOKEN is a free data retrieval call binding the contract method 0x41c841c3.
//
// Solidity: function L1_WETH_TOKEN() view returns(address)
func (_IL1AssetRouter *IL1AssetRouterSession) L1WETHTOKEN() (common.Address, error) {
	return _IL1AssetRouter.Contract.L1WETHTOKEN(&_IL1AssetRouter.CallOpts)
}

// L1WETHTOKEN is a free data retrieval call binding the contract method 0x41c841c3.
//
// Solidity: function L1_WETH_TOKEN() view returns(address)
func (_IL1AssetRouter *IL1AssetRouterCallerSession) L1WETHTOKEN() (common.Address, error) {
	return _IL1AssetRouter.Contract.L1WETHTOKEN(&_IL1AssetRouter.CallOpts)
}

// AssetDeploymentTracker is a free data retrieval call binding the contract method 0x85e4e16a.
//
// Solidity: function assetDeploymentTracker(bytes32 assetId) view returns(address assetDeploymentTracker)
func (_IL1AssetRouter *IL1AssetRouterCaller) AssetDeploymentTracker(opts *bind.CallOpts, assetId [32]byte) (common.Address, error) {
	var out []interface{}
	err := _IL1AssetRouter.contract.Call(opts, &out, "assetDeploymentTracker", assetId)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// AssetDeploymentTracker is a free data retrieval call binding the contract method 0x85e4e16a.
//
// Solidity: function assetDeploymentTracker(bytes32 assetId) view returns(address assetDeploymentTracker)
func (_IL1AssetRouter *IL1AssetRouterSession) AssetDeploymentTracker(assetId [32]byte) (common.Address, error) {
	return _IL1AssetRouter.Contract.AssetDeploymentTracker(&_IL1AssetRouter.CallOpts, assetId)
}

// AssetDeploymentTracker is a free data retrieval call binding the contract method 0x85e4e16a.
//
// Solidity: function assetDeploymentTracker(bytes32 assetId) view returns(address assetDeploymentTracker)
func (_IL1AssetRouter *IL1AssetRouterCallerSession) AssetDeploymentTracker(assetId [32]byte) (common.Address, error) {
	return _IL1AssetRouter.Contract.AssetDeploymentTracker(&_IL1AssetRouter.CallOpts, assetId)
}

// AssetHandlerAddress is a free data retrieval call binding the contract method 0x53b9e632.
//
// Solidity: function assetHandlerAddress(bytes32 assetId) view returns(address assetHandlerAddress)
func (_IL1AssetRouter *IL1AssetRouterCaller) AssetHandlerAddress(opts *bind.CallOpts, assetId [32]byte) (common.Address, error) {
	var out []interface{}
	err := _IL1AssetRouter.contract.Call(opts, &out, "assetHandlerAddress", assetId)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// AssetHandlerAddress is a free data retrieval call binding the contract method 0x53b9e632.
//
// Solidity: function assetHandlerAddress(bytes32 assetId) view returns(address assetHandlerAddress)
func (_IL1AssetRouter *IL1AssetRouterSession) AssetHandlerAddress(assetId [32]byte) (common.Address, error) {
	return _IL1AssetRouter.Contract.AssetHandlerAddress(&_IL1AssetRouter.CallOpts, assetId)
}

// AssetHandlerAddress is a free data retrieval call binding the contract method 0x53b9e632.
//
// Solidity: function assetHandlerAddress(bytes32 assetId) view returns(address assetHandlerAddress)
func (_IL1AssetRouter *IL1AssetRouterCallerSession) AssetHandlerAddress(assetId [32]byte) (common.Address, error) {
	return _IL1AssetRouter.Contract.AssetHandlerAddress(&_IL1AssetRouter.CallOpts, assetId)
}

// GetDepositCalldata is a free data retrieval call binding the contract method 0x2ff0b2ea.
//
// Solidity: function getDepositCalldata(address _sender, bytes32 _assetId, bytes _assetData) view returns(bytes)
func (_IL1AssetRouter *IL1AssetRouterCaller) GetDepositCalldata(opts *bind.CallOpts, _sender common.Address, _assetId [32]byte, _assetData []byte) ([]byte, error) {
	var out []interface{}
	err := _IL1AssetRouter.contract.Call(opts, &out, "getDepositCalldata", _sender, _assetId, _assetData)

	if err != nil {
		return *new([]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([]byte)).(*[]byte)

	return out0, err

}

// GetDepositCalldata is a free data retrieval call binding the contract method 0x2ff0b2ea.
//
// Solidity: function getDepositCalldata(address _sender, bytes32 _assetId, bytes _assetData) view returns(bytes)
func (_IL1AssetRouter *IL1AssetRouterSession) GetDepositCalldata(_sender common.Address, _assetId [32]byte, _assetData []byte) ([]byte, error) {
	return _IL1AssetRouter.Contract.GetDepositCalldata(&_IL1AssetRouter.CallOpts, _sender, _assetId, _assetData)
}

// GetDepositCalldata is a free data retrieval call binding the contract method 0x2ff0b2ea.
//
// Solidity: function getDepositCalldata(address _sender, bytes32 _assetId, bytes _assetData) view returns(bytes)
func (_IL1AssetRouter *IL1AssetRouterCallerSession) GetDepositCalldata(_sender common.Address, _assetId [32]byte, _assetData []byte) ([]byte, error) {
	return _IL1AssetRouter.Contract.GetDepositCalldata(&_IL1AssetRouter.CallOpts, _sender, _assetId, _assetData)
}

// IsWithdrawalFinalized is a free data retrieval call binding the contract method 0x8f31f052.
//
// Solidity: function isWithdrawalFinalized(uint256 _chainId, uint256 _l2BatchNumber, uint256 _l2MessageIndex) view returns(bool)
func (_IL1AssetRouter *IL1AssetRouterCaller) IsWithdrawalFinalized(opts *bind.CallOpts, _chainId *big.Int, _l2BatchNumber *big.Int, _l2MessageIndex *big.Int) (bool, error) {
	var out []interface{}
	err := _IL1AssetRouter.contract.Call(opts, &out, "isWithdrawalFinalized", _chainId, _l2BatchNumber, _l2MessageIndex)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsWithdrawalFinalized is a free data retrieval call binding the contract method 0x8f31f052.
//
// Solidity: function isWithdrawalFinalized(uint256 _chainId, uint256 _l2BatchNumber, uint256 _l2MessageIndex) view returns(bool)
func (_IL1AssetRouter *IL1AssetRouterSession) IsWithdrawalFinalized(_chainId *big.Int, _l2BatchNumber *big.Int, _l2MessageIndex *big.Int) (bool, error) {
	return _IL1AssetRouter.Contract.IsWithdrawalFinalized(&_IL1AssetRouter.CallOpts, _chainId, _l2BatchNumber, _l2MessageIndex)
}

// IsWithdrawalFinalized is a free data retrieval call binding the contract method 0x8f31f052.
//
// Solidity: function isWithdrawalFinalized(uint256 _chainId, uint256 _l2BatchNumber, uint256 _l2MessageIndex) view returns(bool)
func (_IL1AssetRouter *IL1AssetRouterCallerSession) IsWithdrawalFinalized(_chainId *big.Int, _l2BatchNumber *big.Int, _l2MessageIndex *big.Int) (bool, error) {
	return _IL1AssetRouter.Contract.IsWithdrawalFinalized(&_IL1AssetRouter.CallOpts, _chainId, _l2BatchNumber, _l2MessageIndex)
}

// LegacyBridge is a free data retrieval call binding the contract method 0x6e9d7899.
//
// Solidity: function legacyBridge() view returns(address)
func (_IL1AssetRouter *IL1AssetRouterCaller) LegacyBridge(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _IL1AssetRouter.contract.Call(opts, &out, "legacyBridge")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// LegacyBridge is a free data retrieval call binding the contract method 0x6e9d7899.
//
// Solidity: function legacyBridge() view returns(address)
func (_IL1AssetRouter *IL1AssetRouterSession) LegacyBridge() (common.Address, error) {
	return _IL1AssetRouter.Contract.LegacyBridge(&_IL1AssetRouter.CallOpts)
}

// LegacyBridge is a free data retrieval call binding the contract method 0x6e9d7899.
//
// Solidity: function legacyBridge() view returns(address)
func (_IL1AssetRouter *IL1AssetRouterCallerSession) LegacyBridge() (common.Address, error) {
	return _IL1AssetRouter.Contract.LegacyBridge(&_IL1AssetRouter.CallOpts)
}

// NativeTokenVault is a free data retrieval call binding the contract method 0x64e130cf.
//
// Solidity: function nativeTokenVault() view returns(address)
func (_IL1AssetRouter *IL1AssetRouterCaller) NativeTokenVault(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _IL1AssetRouter.contract.Call(opts, &out, "nativeTokenVault")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// NativeTokenVault is a free data retrieval call binding the contract method 0x64e130cf.
//
// Solidity: function nativeTokenVault() view returns(address)
func (_IL1AssetRouter *IL1AssetRouterSession) NativeTokenVault() (common.Address, error) {
	return _IL1AssetRouter.Contract.NativeTokenVault(&_IL1AssetRouter.CallOpts)
}

// NativeTokenVault is a free data retrieval call binding the contract method 0x64e130cf.
//
// Solidity: function nativeTokenVault() view returns(address)
func (_IL1AssetRouter *IL1AssetRouterCallerSession) NativeTokenVault() (common.Address, error) {
	return _IL1AssetRouter.Contract.NativeTokenVault(&_IL1AssetRouter.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_IL1AssetRouter *IL1AssetRouterCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _IL1AssetRouter.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_IL1AssetRouter *IL1AssetRouterSession) Owner() (common.Address, error) {
	return _IL1AssetRouter.Contract.Owner(&_IL1AssetRouter.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_IL1AssetRouter *IL1AssetRouterCallerSession) Owner() (common.Address, error) {
	return _IL1AssetRouter.Contract.Owner(&_IL1AssetRouter.CallOpts)
}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_IL1AssetRouter *IL1AssetRouterCaller) Paused(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _IL1AssetRouter.contract.Call(opts, &out, "paused")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_IL1AssetRouter *IL1AssetRouterSession) Paused() (bool, error) {
	return _IL1AssetRouter.Contract.Paused(&_IL1AssetRouter.CallOpts)
}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_IL1AssetRouter *IL1AssetRouterCallerSession) Paused() (bool, error) {
	return _IL1AssetRouter.Contract.Paused(&_IL1AssetRouter.CallOpts)
}

// PendingOwner is a free data retrieval call binding the contract method 0xe30c3978.
//
// Solidity: function pendingOwner() view returns(address)
func (_IL1AssetRouter *IL1AssetRouterCaller) PendingOwner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _IL1AssetRouter.contract.Call(opts, &out, "pendingOwner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// PendingOwner is a free data retrieval call binding the contract method 0xe30c3978.
//
// Solidity: function pendingOwner() view returns(address)
func (_IL1AssetRouter *IL1AssetRouterSession) PendingOwner() (common.Address, error) {
	return _IL1AssetRouter.Contract.PendingOwner(&_IL1AssetRouter.CallOpts)
}

// PendingOwner is a free data retrieval call binding the contract method 0xe30c3978.
//
// Solidity: function pendingOwner() view returns(address)
func (_IL1AssetRouter *IL1AssetRouterCallerSession) PendingOwner() (common.Address, error) {
	return _IL1AssetRouter.Contract.PendingOwner(&_IL1AssetRouter.CallOpts)
}

// AcceptOwnership is a paid mutator transaction binding the contract method 0x79ba5097.
//
// Solidity: function acceptOwnership() returns()
func (_IL1AssetRouter *IL1AssetRouterTransactor) AcceptOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IL1AssetRouter.contract.Transact(opts, "acceptOwnership")
}

// AcceptOwnership is a paid mutator transaction binding the contract method 0x79ba5097.
//
// Solidity: function acceptOwnership() returns()
func (_IL1AssetRouter *IL1AssetRouterSession) AcceptOwnership() (*types.Transaction, error) {
	return _IL1AssetRouter.Contract.AcceptOwnership(&_IL1AssetRouter.TransactOpts)
}

// AcceptOwnership is a paid mutator transaction binding the contract method 0x79ba5097.
//
// Solidity: function acceptOwnership() returns()
func (_IL1AssetRouter *IL1AssetRouterTransactorSession) AcceptOwnership() (*types.Transaction, error) {
	return _IL1AssetRouter.Contract.AcceptOwnership(&_IL1AssetRouter.TransactOpts)
}

// BridgeRecoverFailedTransfer is a paid mutator transaction binding the contract method 0x1346ca3b.
//
// Solidity: function bridgeRecoverFailedTransfer(uint256 _chainId, address _depositSender, bytes32 _assetId, bytes _assetData) returns()
func (_IL1AssetRouter *IL1AssetRouterTransactor) BridgeRecoverFailedTransfer(opts *bind.TransactOpts, _chainId *big.Int, _depositSender common.Address, _assetId [32]byte, _assetData []byte) (*types.Transaction, error) {
	return _IL1AssetRouter.contract.Transact(opts, "bridgeRecoverFailedTransfer", _chainId, _depositSender, _assetId, _assetData)
}

// BridgeRecoverFailedTransfer is a paid mutator transaction binding the contract method 0x1346ca3b.
//
// Solidity: function bridgeRecoverFailedTransfer(uint256 _chainId, address _depositSender, bytes32 _assetId, bytes _assetData) returns()
func (_IL1AssetRouter *IL1AssetRouterSession) BridgeRecoverFailedTransfer(_chainId *big.Int, _depositSender common.Address, _assetId [32]byte, _assetData []byte) (*types.Transaction, error) {
	return _IL1AssetRouter.Contract.BridgeRecoverFailedTransfer(&_IL1AssetRouter.TransactOpts, _chainId, _depositSender, _assetId, _assetData)
}

// BridgeRecoverFailedTransfer is a paid mutator transaction binding the contract method 0x1346ca3b.
//
// Solidity: function bridgeRecoverFailedTransfer(uint256 _chainId, address _depositSender, bytes32 _assetId, bytes _assetData) returns()
func (_IL1AssetRouter *IL1AssetRouterTransactorSession) BridgeRecoverFailedTransfer(_chainId *big.Int, _depositSender common.Address, _assetId [32]byte, _assetData []byte) (*types.Transaction, error) {
	return _IL1AssetRouter.Contract.BridgeRecoverFailedTransfer(&_IL1AssetRouter.TransactOpts, _chainId, _depositSender, _assetId, _assetData)
}

// BridgeRecoverFailedTransfer0 is a paid mutator transaction binding the contract method 0x3601e63e.
//
// Solidity: function bridgeRecoverFailedTransfer(uint256 _chainId, address _depositSender, bytes32 _assetId, bytes _assetData, bytes32 _l2TxHash, uint256 _l2BatchNumber, uint256 _l2MessageIndex, uint16 _l2TxNumberInBatch, bytes32[] _merkleProof) returns()
func (_IL1AssetRouter *IL1AssetRouterTransactor) BridgeRecoverFailedTransfer0(opts *bind.TransactOpts, _chainId *big.Int, _depositSender common.Address, _assetId [32]byte, _assetData []byte, _l2TxHash [32]byte, _l2BatchNumber *big.Int, _l2MessageIndex *big.Int, _l2TxNumberInBatch uint16, _merkleProof [][32]byte) (*types.Transaction, error) {
	return _IL1AssetRouter.contract.Transact(opts, "bridgeRecoverFailedTransfer0", _chainId, _depositSender, _assetId, _assetData, _l2TxHash, _l2BatchNumber, _l2MessageIndex, _l2TxNumberInBatch, _merkleProof)
}

// BridgeRecoverFailedTransfer0 is a paid mutator transaction binding the contract method 0x3601e63e.
//
// Solidity: function bridgeRecoverFailedTransfer(uint256 _chainId, address _depositSender, bytes32 _assetId, bytes _assetData, bytes32 _l2TxHash, uint256 _l2BatchNumber, uint256 _l2MessageIndex, uint16 _l2TxNumberInBatch, bytes32[] _merkleProof) returns()
func (_IL1AssetRouter *IL1AssetRouterSession) BridgeRecoverFailedTransfer0(_chainId *big.Int, _depositSender common.Address, _assetId [32]byte, _assetData []byte, _l2TxHash [32]byte, _l2BatchNumber *big.Int, _l2MessageIndex *big.Int, _l2TxNumberInBatch uint16, _merkleProof [][32]byte) (*types.Transaction, error) {
	return _IL1AssetRouter.Contract.BridgeRecoverFailedTransfer0(&_IL1AssetRouter.TransactOpts, _chainId, _depositSender, _assetId, _assetData, _l2TxHash, _l2BatchNumber, _l2MessageIndex, _l2TxNumberInBatch, _merkleProof)
}

// BridgeRecoverFailedTransfer0 is a paid mutator transaction binding the contract method 0x3601e63e.
//
// Solidity: function bridgeRecoverFailedTransfer(uint256 _chainId, address _depositSender, bytes32 _assetId, bytes _assetData, bytes32 _l2TxHash, uint256 _l2BatchNumber, uint256 _l2MessageIndex, uint16 _l2TxNumberInBatch, bytes32[] _merkleProof) returns()
func (_IL1AssetRouter *IL1AssetRouterTransactorSession) BridgeRecoverFailedTransfer0(_chainId *big.Int, _depositSender common.Address, _assetId [32]byte, _assetData []byte, _l2TxHash [32]byte, _l2BatchNumber *big.Int, _l2MessageIndex *big.Int, _l2TxNumberInBatch uint16, _merkleProof [][32]byte) (*types.Transaction, error) {
	return _IL1AssetRouter.Contract.BridgeRecoverFailedTransfer0(&_IL1AssetRouter.TransactOpts, _chainId, _depositSender, _assetId, _assetData, _l2TxHash, _l2BatchNumber, _l2MessageIndex, _l2TxNumberInBatch, _merkleProof)
}

// BridgehubConfirmL2Transaction is a paid mutator transaction binding the contract method 0x8eb7db57.
//
// Solidity: function bridgehubConfirmL2Transaction(uint256 _chainId, bytes32 _txDataHash, bytes32 _txHash) returns()
func (_IL1AssetRouter *IL1AssetRouterTransactor) BridgehubConfirmL2Transaction(opts *bind.TransactOpts, _chainId *big.Int, _txDataHash [32]byte, _txHash [32]byte) (*types.Transaction, error) {
	return _IL1AssetRouter.contract.Transact(opts, "bridgehubConfirmL2Transaction", _chainId, _txDataHash, _txHash)
}

// BridgehubConfirmL2Transaction is a paid mutator transaction binding the contract method 0x8eb7db57.
//
// Solidity: function bridgehubConfirmL2Transaction(uint256 _chainId, bytes32 _txDataHash, bytes32 _txHash) returns()
func (_IL1AssetRouter *IL1AssetRouterSession) BridgehubConfirmL2Transaction(_chainId *big.Int, _txDataHash [32]byte, _txHash [32]byte) (*types.Transaction, error) {
	return _IL1AssetRouter.Contract.BridgehubConfirmL2Transaction(&_IL1AssetRouter.TransactOpts, _chainId, _txDataHash, _txHash)
}

// BridgehubConfirmL2Transaction is a paid mutator transaction binding the contract method 0x8eb7db57.
//
// Solidity: function bridgehubConfirmL2Transaction(uint256 _chainId, bytes32 _txDataHash, bytes32 _txHash) returns()
func (_IL1AssetRouter *IL1AssetRouterTransactorSession) BridgehubConfirmL2Transaction(_chainId *big.Int, _txDataHash [32]byte, _txHash [32]byte) (*types.Transaction, error) {
	return _IL1AssetRouter.Contract.BridgehubConfirmL2Transaction(&_IL1AssetRouter.TransactOpts, _chainId, _txDataHash, _txHash)
}

// BridgehubDeposit is a paid mutator transaction binding the contract method 0xca408c23.
//
// Solidity: function bridgehubDeposit(uint256 _chainId, address _originalCaller, uint256 _value, bytes _data) payable returns((bytes32,address,bytes,bytes[],bytes32) request)
func (_IL1AssetRouter *IL1AssetRouterTransactor) BridgehubDeposit(opts *bind.TransactOpts, _chainId *big.Int, _originalCaller common.Address, _value *big.Int, _data []byte) (*types.Transaction, error) {
	return _IL1AssetRouter.contract.Transact(opts, "bridgehubDeposit", _chainId, _originalCaller, _value, _data)
}

// BridgehubDeposit is a paid mutator transaction binding the contract method 0xca408c23.
//
// Solidity: function bridgehubDeposit(uint256 _chainId, address _originalCaller, uint256 _value, bytes _data) payable returns((bytes32,address,bytes,bytes[],bytes32) request)
func (_IL1AssetRouter *IL1AssetRouterSession) BridgehubDeposit(_chainId *big.Int, _originalCaller common.Address, _value *big.Int, _data []byte) (*types.Transaction, error) {
	return _IL1AssetRouter.Contract.BridgehubDeposit(&_IL1AssetRouter.TransactOpts, _chainId, _originalCaller, _value, _data)
}

// BridgehubDeposit is a paid mutator transaction binding the contract method 0xca408c23.
//
// Solidity: function bridgehubDeposit(uint256 _chainId, address _originalCaller, uint256 _value, bytes _data) payable returns((bytes32,address,bytes,bytes[],bytes32) request)
func (_IL1AssetRouter *IL1AssetRouterTransactorSession) BridgehubDeposit(_chainId *big.Int, _originalCaller common.Address, _value *big.Int, _data []byte) (*types.Transaction, error) {
	return _IL1AssetRouter.Contract.BridgehubDeposit(&_IL1AssetRouter.TransactOpts, _chainId, _originalCaller, _value, _data)
}

// BridgehubDepositBaseToken is a paid mutator transaction binding the contract method 0xc4879440.
//
// Solidity: function bridgehubDepositBaseToken(uint256 _chainId, bytes32 _assetId, address _originalCaller, uint256 _amount) payable returns()
func (_IL1AssetRouter *IL1AssetRouterTransactor) BridgehubDepositBaseToken(opts *bind.TransactOpts, _chainId *big.Int, _assetId [32]byte, _originalCaller common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _IL1AssetRouter.contract.Transact(opts, "bridgehubDepositBaseToken", _chainId, _assetId, _originalCaller, _amount)
}

// BridgehubDepositBaseToken is a paid mutator transaction binding the contract method 0xc4879440.
//
// Solidity: function bridgehubDepositBaseToken(uint256 _chainId, bytes32 _assetId, address _originalCaller, uint256 _amount) payable returns()
func (_IL1AssetRouter *IL1AssetRouterSession) BridgehubDepositBaseToken(_chainId *big.Int, _assetId [32]byte, _originalCaller common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _IL1AssetRouter.Contract.BridgehubDepositBaseToken(&_IL1AssetRouter.TransactOpts, _chainId, _assetId, _originalCaller, _amount)
}

// BridgehubDepositBaseToken is a paid mutator transaction binding the contract method 0xc4879440.
//
// Solidity: function bridgehubDepositBaseToken(uint256 _chainId, bytes32 _assetId, address _originalCaller, uint256 _amount) payable returns()
func (_IL1AssetRouter *IL1AssetRouterTransactorSession) BridgehubDepositBaseToken(_chainId *big.Int, _assetId [32]byte, _originalCaller common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _IL1AssetRouter.Contract.BridgehubDepositBaseToken(&_IL1AssetRouter.TransactOpts, _chainId, _assetId, _originalCaller, _amount)
}

// ClaimFailedDeposit is a paid mutator transaction binding the contract method 0xc0991525.
//
// Solidity: function claimFailedDeposit(uint256 _chainId, address _depositSender, address _l1Token, uint256 _amount, bytes32 _l2TxHash, uint256 _l2BatchNumber, uint256 _l2MessageIndex, uint16 _l2TxNumberInBatch, bytes32[] _merkleProof) returns()
func (_IL1AssetRouter *IL1AssetRouterTransactor) ClaimFailedDeposit(opts *bind.TransactOpts, _chainId *big.Int, _depositSender common.Address, _l1Token common.Address, _amount *big.Int, _l2TxHash [32]byte, _l2BatchNumber *big.Int, _l2MessageIndex *big.Int, _l2TxNumberInBatch uint16, _merkleProof [][32]byte) (*types.Transaction, error) {
	return _IL1AssetRouter.contract.Transact(opts, "claimFailedDeposit", _chainId, _depositSender, _l1Token, _amount, _l2TxHash, _l2BatchNumber, _l2MessageIndex, _l2TxNumberInBatch, _merkleProof)
}

// ClaimFailedDeposit is a paid mutator transaction binding the contract method 0xc0991525.
//
// Solidity: function claimFailedDeposit(uint256 _chainId, address _depositSender, address _l1Token, uint256 _amount, bytes32 _l2TxHash, uint256 _l2BatchNumber, uint256 _l2MessageIndex, uint16 _l2TxNumberInBatch, bytes32[] _merkleProof) returns()
func (_IL1AssetRouter *IL1AssetRouterSession) ClaimFailedDeposit(_chainId *big.Int, _depositSender common.Address, _l1Token common.Address, _amount *big.Int, _l2TxHash [32]byte, _l2BatchNumber *big.Int, _l2MessageIndex *big.Int, _l2TxNumberInBatch uint16, _merkleProof [][32]byte) (*types.Transaction, error) {
	return _IL1AssetRouter.Contract.ClaimFailedDeposit(&_IL1AssetRouter.TransactOpts, _chainId, _depositSender, _l1Token, _amount, _l2TxHash, _l2BatchNumber, _l2MessageIndex, _l2TxNumberInBatch, _merkleProof)
}

// ClaimFailedDeposit is a paid mutator transaction binding the contract method 0xc0991525.
//
// Solidity: function claimFailedDeposit(uint256 _chainId, address _depositSender, address _l1Token, uint256 _amount, bytes32 _l2TxHash, uint256 _l2BatchNumber, uint256 _l2MessageIndex, uint16 _l2TxNumberInBatch, bytes32[] _merkleProof) returns()
func (_IL1AssetRouter *IL1AssetRouterTransactorSession) ClaimFailedDeposit(_chainId *big.Int, _depositSender common.Address, _l1Token common.Address, _amount *big.Int, _l2TxHash [32]byte, _l2BatchNumber *big.Int, _l2MessageIndex *big.Int, _l2TxNumberInBatch uint16, _merkleProof [][32]byte) (*types.Transaction, error) {
	return _IL1AssetRouter.Contract.ClaimFailedDeposit(&_IL1AssetRouter.TransactOpts, _chainId, _depositSender, _l1Token, _amount, _l2TxHash, _l2BatchNumber, _l2MessageIndex, _l2TxNumberInBatch, _merkleProof)
}

// DepositLegacyErc20Bridge is a paid mutator transaction binding the contract method 0x9e6ea417.
//
// Solidity: function depositLegacyErc20Bridge(address _originalCaller, address _l2Receiver, address _l1Token, uint256 _amount, uint256 _l2TxGasLimit, uint256 _l2TxGasPerPubdataByte, address _refundRecipient) payable returns(bytes32 txHash)
func (_IL1AssetRouter *IL1AssetRouterTransactor) DepositLegacyErc20Bridge(opts *bind.TransactOpts, _originalCaller common.Address, _l2Receiver common.Address, _l1Token common.Address, _amount *big.Int, _l2TxGasLimit *big.Int, _l2TxGasPerPubdataByte *big.Int, _refundRecipient common.Address) (*types.Transaction, error) {
	return _IL1AssetRouter.contract.Transact(opts, "depositLegacyErc20Bridge", _originalCaller, _l2Receiver, _l1Token, _amount, _l2TxGasLimit, _l2TxGasPerPubdataByte, _refundRecipient)
}

// DepositLegacyErc20Bridge is a paid mutator transaction binding the contract method 0x9e6ea417.
//
// Solidity: function depositLegacyErc20Bridge(address _originalCaller, address _l2Receiver, address _l1Token, uint256 _amount, uint256 _l2TxGasLimit, uint256 _l2TxGasPerPubdataByte, address _refundRecipient) payable returns(bytes32 txHash)
func (_IL1AssetRouter *IL1AssetRouterSession) DepositLegacyErc20Bridge(_originalCaller common.Address, _l2Receiver common.Address, _l1Token common.Address, _amount *big.Int, _l2TxGasLimit *big.Int, _l2TxGasPerPubdataByte *big.Int, _refundRecipient common.Address) (*types.Transaction, error) {
	return _IL1AssetRouter.Contract.DepositLegacyErc20Bridge(&_IL1AssetRouter.TransactOpts, _originalCaller, _l2Receiver, _l1Token, _amount, _l2TxGasLimit, _l2TxGasPerPubdataByte, _refundRecipient)
}

// DepositLegacyErc20Bridge is a paid mutator transaction binding the contract method 0x9e6ea417.
//
// Solidity: function depositLegacyErc20Bridge(address _originalCaller, address _l2Receiver, address _l1Token, uint256 _amount, uint256 _l2TxGasLimit, uint256 _l2TxGasPerPubdataByte, address _refundRecipient) payable returns(bytes32 txHash)
func (_IL1AssetRouter *IL1AssetRouterTransactorSession) DepositLegacyErc20Bridge(_originalCaller common.Address, _l2Receiver common.Address, _l1Token common.Address, _amount *big.Int, _l2TxGasLimit *big.Int, _l2TxGasPerPubdataByte *big.Int, _refundRecipient common.Address) (*types.Transaction, error) {
	return _IL1AssetRouter.Contract.DepositLegacyErc20Bridge(&_IL1AssetRouter.TransactOpts, _originalCaller, _l2Receiver, _l1Token, _amount, _l2TxGasLimit, _l2TxGasPerPubdataByte, _refundRecipient)
}

// FinalizeDeposit is a paid mutator transaction binding the contract method 0x9c884fd1.
//
// Solidity: function finalizeDeposit(uint256 _chainId, bytes32 _assetId, bytes _transferData) returns()
func (_IL1AssetRouter *IL1AssetRouterTransactor) FinalizeDeposit(opts *bind.TransactOpts, _chainId *big.Int, _assetId [32]byte, _transferData []byte) (*types.Transaction, error) {
	return _IL1AssetRouter.contract.Transact(opts, "finalizeDeposit", _chainId, _assetId, _transferData)
}

// FinalizeDeposit is a paid mutator transaction binding the contract method 0x9c884fd1.
//
// Solidity: function finalizeDeposit(uint256 _chainId, bytes32 _assetId, bytes _transferData) returns()
func (_IL1AssetRouter *IL1AssetRouterSession) FinalizeDeposit(_chainId *big.Int, _assetId [32]byte, _transferData []byte) (*types.Transaction, error) {
	return _IL1AssetRouter.Contract.FinalizeDeposit(&_IL1AssetRouter.TransactOpts, _chainId, _assetId, _transferData)
}

// FinalizeDeposit is a paid mutator transaction binding the contract method 0x9c884fd1.
//
// Solidity: function finalizeDeposit(uint256 _chainId, bytes32 _assetId, bytes _transferData) returns()
func (_IL1AssetRouter *IL1AssetRouterTransactorSession) FinalizeDeposit(_chainId *big.Int, _assetId [32]byte, _transferData []byte) (*types.Transaction, error) {
	return _IL1AssetRouter.Contract.FinalizeDeposit(&_IL1AssetRouter.TransactOpts, _chainId, _assetId, _transferData)
}

// FinalizeWithdrawal is a paid mutator transaction binding the contract method 0xc87325f1.
//
// Solidity: function finalizeWithdrawal(uint256 _chainId, uint256 _l2BatchNumber, uint256 _l2MessageIndex, uint16 _l2TxNumberInBatch, bytes _message, bytes32[] _merkleProof) returns()
func (_IL1AssetRouter *IL1AssetRouterTransactor) FinalizeWithdrawal(opts *bind.TransactOpts, _chainId *big.Int, _l2BatchNumber *big.Int, _l2MessageIndex *big.Int, _l2TxNumberInBatch uint16, _message []byte, _merkleProof [][32]byte) (*types.Transaction, error) {
	return _IL1AssetRouter.contract.Transact(opts, "finalizeWithdrawal", _chainId, _l2BatchNumber, _l2MessageIndex, _l2TxNumberInBatch, _message, _merkleProof)
}

// FinalizeWithdrawal is a paid mutator transaction binding the contract method 0xc87325f1.
//
// Solidity: function finalizeWithdrawal(uint256 _chainId, uint256 _l2BatchNumber, uint256 _l2MessageIndex, uint16 _l2TxNumberInBatch, bytes _message, bytes32[] _merkleProof) returns()
func (_IL1AssetRouter *IL1AssetRouterSession) FinalizeWithdrawal(_chainId *big.Int, _l2BatchNumber *big.Int, _l2MessageIndex *big.Int, _l2TxNumberInBatch uint16, _message []byte, _merkleProof [][32]byte) (*types.Transaction, error) {
	return _IL1AssetRouter.Contract.FinalizeWithdrawal(&_IL1AssetRouter.TransactOpts, _chainId, _l2BatchNumber, _l2MessageIndex, _l2TxNumberInBatch, _message, _merkleProof)
}

// FinalizeWithdrawal is a paid mutator transaction binding the contract method 0xc87325f1.
//
// Solidity: function finalizeWithdrawal(uint256 _chainId, uint256 _l2BatchNumber, uint256 _l2MessageIndex, uint16 _l2TxNumberInBatch, bytes _message, bytes32[] _merkleProof) returns()
func (_IL1AssetRouter *IL1AssetRouterTransactorSession) FinalizeWithdrawal(_chainId *big.Int, _l2BatchNumber *big.Int, _l2MessageIndex *big.Int, _l2TxNumberInBatch uint16, _message []byte, _merkleProof [][32]byte) (*types.Transaction, error) {
	return _IL1AssetRouter.Contract.FinalizeWithdrawal(&_IL1AssetRouter.TransactOpts, _chainId, _l2BatchNumber, _l2MessageIndex, _l2TxNumberInBatch, _message, _merkleProof)
}

// Initialize is a paid mutator transaction binding the contract method 0xc4d66de8.
//
// Solidity: function initialize(address _owner) returns()
func (_IL1AssetRouter *IL1AssetRouterTransactor) Initialize(opts *bind.TransactOpts, _owner common.Address) (*types.Transaction, error) {
	return _IL1AssetRouter.contract.Transact(opts, "initialize", _owner)
}

// Initialize is a paid mutator transaction binding the contract method 0xc4d66de8.
//
// Solidity: function initialize(address _owner) returns()
func (_IL1AssetRouter *IL1AssetRouterSession) Initialize(_owner common.Address) (*types.Transaction, error) {
	return _IL1AssetRouter.Contract.Initialize(&_IL1AssetRouter.TransactOpts, _owner)
}

// Initialize is a paid mutator transaction binding the contract method 0xc4d66de8.
//
// Solidity: function initialize(address _owner) returns()
func (_IL1AssetRouter *IL1AssetRouterTransactorSession) Initialize(_owner common.Address) (*types.Transaction, error) {
	return _IL1AssetRouter.Contract.Initialize(&_IL1AssetRouter.TransactOpts, _owner)
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns()
func (_IL1AssetRouter *IL1AssetRouterTransactor) Pause(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IL1AssetRouter.contract.Transact(opts, "pause")
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns()
func (_IL1AssetRouter *IL1AssetRouterSession) Pause() (*types.Transaction, error) {
	return _IL1AssetRouter.Contract.Pause(&_IL1AssetRouter.TransactOpts)
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns()
func (_IL1AssetRouter *IL1AssetRouterTransactorSession) Pause() (*types.Transaction, error) {
	return _IL1AssetRouter.Contract.Pause(&_IL1AssetRouter.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_IL1AssetRouter *IL1AssetRouterTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IL1AssetRouter.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_IL1AssetRouter *IL1AssetRouterSession) RenounceOwnership() (*types.Transaction, error) {
	return _IL1AssetRouter.Contract.RenounceOwnership(&_IL1AssetRouter.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_IL1AssetRouter *IL1AssetRouterTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _IL1AssetRouter.Contract.RenounceOwnership(&_IL1AssetRouter.TransactOpts)
}

// SetAssetDeploymentTracker is a paid mutator transaction binding the contract method 0xc0a16dda.
//
// Solidity: function setAssetDeploymentTracker(bytes32 _assetRegistrationData, address _assetDeploymentTracker) returns()
func (_IL1AssetRouter *IL1AssetRouterTransactor) SetAssetDeploymentTracker(opts *bind.TransactOpts, _assetRegistrationData [32]byte, _assetDeploymentTracker common.Address) (*types.Transaction, error) {
	return _IL1AssetRouter.contract.Transact(opts, "setAssetDeploymentTracker", _assetRegistrationData, _assetDeploymentTracker)
}

// SetAssetDeploymentTracker is a paid mutator transaction binding the contract method 0xc0a16dda.
//
// Solidity: function setAssetDeploymentTracker(bytes32 _assetRegistrationData, address _assetDeploymentTracker) returns()
func (_IL1AssetRouter *IL1AssetRouterSession) SetAssetDeploymentTracker(_assetRegistrationData [32]byte, _assetDeploymentTracker common.Address) (*types.Transaction, error) {
	return _IL1AssetRouter.Contract.SetAssetDeploymentTracker(&_IL1AssetRouter.TransactOpts, _assetRegistrationData, _assetDeploymentTracker)
}

// SetAssetDeploymentTracker is a paid mutator transaction binding the contract method 0xc0a16dda.
//
// Solidity: function setAssetDeploymentTracker(bytes32 _assetRegistrationData, address _assetDeploymentTracker) returns()
func (_IL1AssetRouter *IL1AssetRouterTransactorSession) SetAssetDeploymentTracker(_assetRegistrationData [32]byte, _assetDeploymentTracker common.Address) (*types.Transaction, error) {
	return _IL1AssetRouter.Contract.SetAssetDeploymentTracker(&_IL1AssetRouter.TransactOpts, _assetRegistrationData, _assetDeploymentTracker)
}

// SetAssetHandlerAddressThisChain is a paid mutator transaction binding the contract method 0x548a5a33.
//
// Solidity: function setAssetHandlerAddressThisChain(bytes32 _assetRegistrationData, address _assetHandlerAddress) returns()
func (_IL1AssetRouter *IL1AssetRouterTransactor) SetAssetHandlerAddressThisChain(opts *bind.TransactOpts, _assetRegistrationData [32]byte, _assetHandlerAddress common.Address) (*types.Transaction, error) {
	return _IL1AssetRouter.contract.Transact(opts, "setAssetHandlerAddressThisChain", _assetRegistrationData, _assetHandlerAddress)
}

// SetAssetHandlerAddressThisChain is a paid mutator transaction binding the contract method 0x548a5a33.
//
// Solidity: function setAssetHandlerAddressThisChain(bytes32 _assetRegistrationData, address _assetHandlerAddress) returns()
func (_IL1AssetRouter *IL1AssetRouterSession) SetAssetHandlerAddressThisChain(_assetRegistrationData [32]byte, _assetHandlerAddress common.Address) (*types.Transaction, error) {
	return _IL1AssetRouter.Contract.SetAssetHandlerAddressThisChain(&_IL1AssetRouter.TransactOpts, _assetRegistrationData, _assetHandlerAddress)
}

// SetAssetHandlerAddressThisChain is a paid mutator transaction binding the contract method 0x548a5a33.
//
// Solidity: function setAssetHandlerAddressThisChain(bytes32 _assetRegistrationData, address _assetHandlerAddress) returns()
func (_IL1AssetRouter *IL1AssetRouterTransactorSession) SetAssetHandlerAddressThisChain(_assetRegistrationData [32]byte, _assetHandlerAddress common.Address) (*types.Transaction, error) {
	return _IL1AssetRouter.Contract.SetAssetHandlerAddressThisChain(&_IL1AssetRouter.TransactOpts, _assetRegistrationData, _assetHandlerAddress)
}

// SetL1Erc20Bridge is a paid mutator transaction binding the contract method 0x30bda03e.
//
// Solidity: function setL1Erc20Bridge(address _legacyBridge) returns()
func (_IL1AssetRouter *IL1AssetRouterTransactor) SetL1Erc20Bridge(opts *bind.TransactOpts, _legacyBridge common.Address) (*types.Transaction, error) {
	return _IL1AssetRouter.contract.Transact(opts, "setL1Erc20Bridge", _legacyBridge)
}

// SetL1Erc20Bridge is a paid mutator transaction binding the contract method 0x30bda03e.
//
// Solidity: function setL1Erc20Bridge(address _legacyBridge) returns()
func (_IL1AssetRouter *IL1AssetRouterSession) SetL1Erc20Bridge(_legacyBridge common.Address) (*types.Transaction, error) {
	return _IL1AssetRouter.Contract.SetL1Erc20Bridge(&_IL1AssetRouter.TransactOpts, _legacyBridge)
}

// SetL1Erc20Bridge is a paid mutator transaction binding the contract method 0x30bda03e.
//
// Solidity: function setL1Erc20Bridge(address _legacyBridge) returns()
func (_IL1AssetRouter *IL1AssetRouterTransactorSession) SetL1Erc20Bridge(_legacyBridge common.Address) (*types.Transaction, error) {
	return _IL1AssetRouter.Contract.SetL1Erc20Bridge(&_IL1AssetRouter.TransactOpts, _legacyBridge)
}

// SetNativeTokenVault is a paid mutator transaction binding the contract method 0x0f3fa211.
//
// Solidity: function setNativeTokenVault(address _nativeTokenVault) returns()
func (_IL1AssetRouter *IL1AssetRouterTransactor) SetNativeTokenVault(opts *bind.TransactOpts, _nativeTokenVault common.Address) (*types.Transaction, error) {
	return _IL1AssetRouter.contract.Transact(opts, "setNativeTokenVault", _nativeTokenVault)
}

// SetNativeTokenVault is a paid mutator transaction binding the contract method 0x0f3fa211.
//
// Solidity: function setNativeTokenVault(address _nativeTokenVault) returns()
func (_IL1AssetRouter *IL1AssetRouterSession) SetNativeTokenVault(_nativeTokenVault common.Address) (*types.Transaction, error) {
	return _IL1AssetRouter.Contract.SetNativeTokenVault(&_IL1AssetRouter.TransactOpts, _nativeTokenVault)
}

// SetNativeTokenVault is a paid mutator transaction binding the contract method 0x0f3fa211.
//
// Solidity: function setNativeTokenVault(address _nativeTokenVault) returns()
func (_IL1AssetRouter *IL1AssetRouterTransactorSession) SetNativeTokenVault(_nativeTokenVault common.Address) (*types.Transaction, error) {
	return _IL1AssetRouter.Contract.SetNativeTokenVault(&_IL1AssetRouter.TransactOpts, _nativeTokenVault)
}

// TransferFundsToNTV is a paid mutator transaction binding the contract method 0x57d4ca5c.
//
// Solidity: function transferFundsToNTV(bytes32 _assetId, uint256 _amount, address _originalCaller) returns(bool)
func (_IL1AssetRouter *IL1AssetRouterTransactor) TransferFundsToNTV(opts *bind.TransactOpts, _assetId [32]byte, _amount *big.Int, _originalCaller common.Address) (*types.Transaction, error) {
	return _IL1AssetRouter.contract.Transact(opts, "transferFundsToNTV", _assetId, _amount, _originalCaller)
}

// TransferFundsToNTV is a paid mutator transaction binding the contract method 0x57d4ca5c.
//
// Solidity: function transferFundsToNTV(bytes32 _assetId, uint256 _amount, address _originalCaller) returns(bool)
func (_IL1AssetRouter *IL1AssetRouterSession) TransferFundsToNTV(_assetId [32]byte, _amount *big.Int, _originalCaller common.Address) (*types.Transaction, error) {
	return _IL1AssetRouter.Contract.TransferFundsToNTV(&_IL1AssetRouter.TransactOpts, _assetId, _amount, _originalCaller)
}

// TransferFundsToNTV is a paid mutator transaction binding the contract method 0x57d4ca5c.
//
// Solidity: function transferFundsToNTV(bytes32 _assetId, uint256 _amount, address _originalCaller) returns(bool)
func (_IL1AssetRouter *IL1AssetRouterTransactorSession) TransferFundsToNTV(_assetId [32]byte, _amount *big.Int, _originalCaller common.Address) (*types.Transaction, error) {
	return _IL1AssetRouter.Contract.TransferFundsToNTV(&_IL1AssetRouter.TransactOpts, _assetId, _amount, _originalCaller)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_IL1AssetRouter *IL1AssetRouterTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _IL1AssetRouter.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_IL1AssetRouter *IL1AssetRouterSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _IL1AssetRouter.Contract.TransferOwnership(&_IL1AssetRouter.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_IL1AssetRouter *IL1AssetRouterTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _IL1AssetRouter.Contract.TransferOwnership(&_IL1AssetRouter.TransactOpts, newOwner)
}

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns()
func (_IL1AssetRouter *IL1AssetRouterTransactor) Unpause(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IL1AssetRouter.contract.Transact(opts, "unpause")
}

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns()
func (_IL1AssetRouter *IL1AssetRouterSession) Unpause() (*types.Transaction, error) {
	return _IL1AssetRouter.Contract.Unpause(&_IL1AssetRouter.TransactOpts)
}

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns()
func (_IL1AssetRouter *IL1AssetRouterTransactorSession) Unpause() (*types.Transaction, error) {
	return _IL1AssetRouter.Contract.Unpause(&_IL1AssetRouter.TransactOpts)
}

// IL1AssetRouterAssetDeploymentTrackerSetIterator is returned from FilterAssetDeploymentTrackerSet and is used to iterate over the raw logs and unpacked data for AssetDeploymentTrackerSet events raised by the IL1AssetRouter contract.
type IL1AssetRouterAssetDeploymentTrackerSetIterator struct {
	Event *IL1AssetRouterAssetDeploymentTrackerSet // Event containing the contract specifics and raw log

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
func (it *IL1AssetRouterAssetDeploymentTrackerSetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IL1AssetRouterAssetDeploymentTrackerSet)
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
		it.Event = new(IL1AssetRouterAssetDeploymentTrackerSet)
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
func (it *IL1AssetRouterAssetDeploymentTrackerSetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IL1AssetRouterAssetDeploymentTrackerSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IL1AssetRouterAssetDeploymentTrackerSet represents a AssetDeploymentTrackerSet event raised by the IL1AssetRouter contract.
type IL1AssetRouterAssetDeploymentTrackerSet struct {
	AssetId                [32]byte
	AssetDeploymentTracker common.Address
	AdditionalData         [32]byte
	Raw                    types.Log // Blockchain specific contextual infos
}

// FilterAssetDeploymentTrackerSet is a free log retrieval operation binding the contract event 0x14c1bae9bcc3777747463b66a36584aa75e4ded1aa38089f447beecb125a2175.
//
// Solidity: event AssetDeploymentTrackerSet(bytes32 indexed assetId, address indexed assetDeploymentTracker, bytes32 indexed additionalData)
func (_IL1AssetRouter *IL1AssetRouterFilterer) FilterAssetDeploymentTrackerSet(opts *bind.FilterOpts, assetId [][32]byte, assetDeploymentTracker []common.Address, additionalData [][32]byte) (*IL1AssetRouterAssetDeploymentTrackerSetIterator, error) {

	var assetIdRule []interface{}
	for _, assetIdItem := range assetId {
		assetIdRule = append(assetIdRule, assetIdItem)
	}
	var assetDeploymentTrackerRule []interface{}
	for _, assetDeploymentTrackerItem := range assetDeploymentTracker {
		assetDeploymentTrackerRule = append(assetDeploymentTrackerRule, assetDeploymentTrackerItem)
	}
	var additionalDataRule []interface{}
	for _, additionalDataItem := range additionalData {
		additionalDataRule = append(additionalDataRule, additionalDataItem)
	}

	logs, sub, err := _IL1AssetRouter.contract.FilterLogs(opts, "AssetDeploymentTrackerSet", assetIdRule, assetDeploymentTrackerRule, additionalDataRule)
	if err != nil {
		return nil, err
	}
	return &IL1AssetRouterAssetDeploymentTrackerSetIterator{contract: _IL1AssetRouter.contract, event: "AssetDeploymentTrackerSet", logs: logs, sub: sub}, nil
}

// WatchAssetDeploymentTrackerSet is a free log subscription operation binding the contract event 0x14c1bae9bcc3777747463b66a36584aa75e4ded1aa38089f447beecb125a2175.
//
// Solidity: event AssetDeploymentTrackerSet(bytes32 indexed assetId, address indexed assetDeploymentTracker, bytes32 indexed additionalData)
func (_IL1AssetRouter *IL1AssetRouterFilterer) WatchAssetDeploymentTrackerSet(opts *bind.WatchOpts, sink chan<- *IL1AssetRouterAssetDeploymentTrackerSet, assetId [][32]byte, assetDeploymentTracker []common.Address, additionalData [][32]byte) (event.Subscription, error) {

	var assetIdRule []interface{}
	for _, assetIdItem := range assetId {
		assetIdRule = append(assetIdRule, assetIdItem)
	}
	var assetDeploymentTrackerRule []interface{}
	for _, assetDeploymentTrackerItem := range assetDeploymentTracker {
		assetDeploymentTrackerRule = append(assetDeploymentTrackerRule, assetDeploymentTrackerItem)
	}
	var additionalDataRule []interface{}
	for _, additionalDataItem := range additionalData {
		additionalDataRule = append(additionalDataRule, additionalDataItem)
	}

	logs, sub, err := _IL1AssetRouter.contract.WatchLogs(opts, "AssetDeploymentTrackerSet", assetIdRule, assetDeploymentTrackerRule, additionalDataRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IL1AssetRouterAssetDeploymentTrackerSet)
				if err := _IL1AssetRouter.contract.UnpackLog(event, "AssetDeploymentTrackerSet", log); err != nil {
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

// ParseAssetDeploymentTrackerSet is a log parse operation binding the contract event 0x14c1bae9bcc3777747463b66a36584aa75e4ded1aa38089f447beecb125a2175.
//
// Solidity: event AssetDeploymentTrackerSet(bytes32 indexed assetId, address indexed assetDeploymentTracker, bytes32 indexed additionalData)
func (_IL1AssetRouter *IL1AssetRouterFilterer) ParseAssetDeploymentTrackerSet(log types.Log) (*IL1AssetRouterAssetDeploymentTrackerSet, error) {
	event := new(IL1AssetRouterAssetDeploymentTrackerSet)
	if err := _IL1AssetRouter.contract.UnpackLog(event, "AssetDeploymentTrackerSet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// IL1AssetRouterAssetHandlerRegisteredIterator is returned from FilterAssetHandlerRegistered and is used to iterate over the raw logs and unpacked data for AssetHandlerRegistered events raised by the IL1AssetRouter contract.
type IL1AssetRouterAssetHandlerRegisteredIterator struct {
	Event *IL1AssetRouterAssetHandlerRegistered // Event containing the contract specifics and raw log

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
func (it *IL1AssetRouterAssetHandlerRegisteredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IL1AssetRouterAssetHandlerRegistered)
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
		it.Event = new(IL1AssetRouterAssetHandlerRegistered)
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
func (it *IL1AssetRouterAssetHandlerRegisteredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IL1AssetRouterAssetHandlerRegisteredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IL1AssetRouterAssetHandlerRegistered represents a AssetHandlerRegistered event raised by the IL1AssetRouter contract.
type IL1AssetRouterAssetHandlerRegistered struct {
	AssetId      [32]byte
	AssetAddress common.Address
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterAssetHandlerRegistered is a free log retrieval operation binding the contract event 0x2632cc0d58b0cb1017b99cc0b6cc66ad86440cc0dd923bfdaa294f95ba1b0201.
//
// Solidity: event AssetHandlerRegistered(bytes32 indexed assetId, address indexed _assetAddress)
func (_IL1AssetRouter *IL1AssetRouterFilterer) FilterAssetHandlerRegistered(opts *bind.FilterOpts, assetId [][32]byte, _assetAddress []common.Address) (*IL1AssetRouterAssetHandlerRegisteredIterator, error) {

	var assetIdRule []interface{}
	for _, assetIdItem := range assetId {
		assetIdRule = append(assetIdRule, assetIdItem)
	}
	var _assetAddressRule []interface{}
	for _, _assetAddressItem := range _assetAddress {
		_assetAddressRule = append(_assetAddressRule, _assetAddressItem)
	}

	logs, sub, err := _IL1AssetRouter.contract.FilterLogs(opts, "AssetHandlerRegistered", assetIdRule, _assetAddressRule)
	if err != nil {
		return nil, err
	}
	return &IL1AssetRouterAssetHandlerRegisteredIterator{contract: _IL1AssetRouter.contract, event: "AssetHandlerRegistered", logs: logs, sub: sub}, nil
}

// WatchAssetHandlerRegistered is a free log subscription operation binding the contract event 0x2632cc0d58b0cb1017b99cc0b6cc66ad86440cc0dd923bfdaa294f95ba1b0201.
//
// Solidity: event AssetHandlerRegistered(bytes32 indexed assetId, address indexed _assetAddress)
func (_IL1AssetRouter *IL1AssetRouterFilterer) WatchAssetHandlerRegistered(opts *bind.WatchOpts, sink chan<- *IL1AssetRouterAssetHandlerRegistered, assetId [][32]byte, _assetAddress []common.Address) (event.Subscription, error) {

	var assetIdRule []interface{}
	for _, assetIdItem := range assetId {
		assetIdRule = append(assetIdRule, assetIdItem)
	}
	var _assetAddressRule []interface{}
	for _, _assetAddressItem := range _assetAddress {
		_assetAddressRule = append(_assetAddressRule, _assetAddressItem)
	}

	logs, sub, err := _IL1AssetRouter.contract.WatchLogs(opts, "AssetHandlerRegistered", assetIdRule, _assetAddressRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IL1AssetRouterAssetHandlerRegistered)
				if err := _IL1AssetRouter.contract.UnpackLog(event, "AssetHandlerRegistered", log); err != nil {
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

// ParseAssetHandlerRegistered is a log parse operation binding the contract event 0x2632cc0d58b0cb1017b99cc0b6cc66ad86440cc0dd923bfdaa294f95ba1b0201.
//
// Solidity: event AssetHandlerRegistered(bytes32 indexed assetId, address indexed _assetAddress)
func (_IL1AssetRouter *IL1AssetRouterFilterer) ParseAssetHandlerRegistered(log types.Log) (*IL1AssetRouterAssetHandlerRegistered, error) {
	event := new(IL1AssetRouterAssetHandlerRegistered)
	if err := _IL1AssetRouter.contract.UnpackLog(event, "AssetHandlerRegistered", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// IL1AssetRouterAssetHandlerRegisteredInitialIterator is returned from FilterAssetHandlerRegisteredInitial and is used to iterate over the raw logs and unpacked data for AssetHandlerRegisteredInitial events raised by the IL1AssetRouter contract.
type IL1AssetRouterAssetHandlerRegisteredInitialIterator struct {
	Event *IL1AssetRouterAssetHandlerRegisteredInitial // Event containing the contract specifics and raw log

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
func (it *IL1AssetRouterAssetHandlerRegisteredInitialIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IL1AssetRouterAssetHandlerRegisteredInitial)
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
		it.Event = new(IL1AssetRouterAssetHandlerRegisteredInitial)
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
func (it *IL1AssetRouterAssetHandlerRegisteredInitialIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IL1AssetRouterAssetHandlerRegisteredInitialIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IL1AssetRouterAssetHandlerRegisteredInitial represents a AssetHandlerRegisteredInitial event raised by the IL1AssetRouter contract.
type IL1AssetRouterAssetHandlerRegisteredInitial struct {
	AssetId                [32]byte
	AssetHandlerAddress    common.Address
	AdditionalData         [32]byte
	AssetDeploymentTracker common.Address
	Raw                    types.Log // Blockchain specific contextual infos
}

// FilterAssetHandlerRegisteredInitial is a free log retrieval operation binding the contract event 0xb1e82bee3e85b2755fbceb4b7e051f5c66a7f35f0476657504e77e18ebd3a17d.
//
// Solidity: event AssetHandlerRegisteredInitial(bytes32 indexed assetId, address indexed assetHandlerAddress, bytes32 indexed additionalData, address assetDeploymentTracker)
func (_IL1AssetRouter *IL1AssetRouterFilterer) FilterAssetHandlerRegisteredInitial(opts *bind.FilterOpts, assetId [][32]byte, assetHandlerAddress []common.Address, additionalData [][32]byte) (*IL1AssetRouterAssetHandlerRegisteredInitialIterator, error) {

	var assetIdRule []interface{}
	for _, assetIdItem := range assetId {
		assetIdRule = append(assetIdRule, assetIdItem)
	}
	var assetHandlerAddressRule []interface{}
	for _, assetHandlerAddressItem := range assetHandlerAddress {
		assetHandlerAddressRule = append(assetHandlerAddressRule, assetHandlerAddressItem)
	}
	var additionalDataRule []interface{}
	for _, additionalDataItem := range additionalData {
		additionalDataRule = append(additionalDataRule, additionalDataItem)
	}

	logs, sub, err := _IL1AssetRouter.contract.FilterLogs(opts, "AssetHandlerRegisteredInitial", assetIdRule, assetHandlerAddressRule, additionalDataRule)
	if err != nil {
		return nil, err
	}
	return &IL1AssetRouterAssetHandlerRegisteredInitialIterator{contract: _IL1AssetRouter.contract, event: "AssetHandlerRegisteredInitial", logs: logs, sub: sub}, nil
}

// WatchAssetHandlerRegisteredInitial is a free log subscription operation binding the contract event 0xb1e82bee3e85b2755fbceb4b7e051f5c66a7f35f0476657504e77e18ebd3a17d.
//
// Solidity: event AssetHandlerRegisteredInitial(bytes32 indexed assetId, address indexed assetHandlerAddress, bytes32 indexed additionalData, address assetDeploymentTracker)
func (_IL1AssetRouter *IL1AssetRouterFilterer) WatchAssetHandlerRegisteredInitial(opts *bind.WatchOpts, sink chan<- *IL1AssetRouterAssetHandlerRegisteredInitial, assetId [][32]byte, assetHandlerAddress []common.Address, additionalData [][32]byte) (event.Subscription, error) {

	var assetIdRule []interface{}
	for _, assetIdItem := range assetId {
		assetIdRule = append(assetIdRule, assetIdItem)
	}
	var assetHandlerAddressRule []interface{}
	for _, assetHandlerAddressItem := range assetHandlerAddress {
		assetHandlerAddressRule = append(assetHandlerAddressRule, assetHandlerAddressItem)
	}
	var additionalDataRule []interface{}
	for _, additionalDataItem := range additionalData {
		additionalDataRule = append(additionalDataRule, additionalDataItem)
	}

	logs, sub, err := _IL1AssetRouter.contract.WatchLogs(opts, "AssetHandlerRegisteredInitial", assetIdRule, assetHandlerAddressRule, additionalDataRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IL1AssetRouterAssetHandlerRegisteredInitial)
				if err := _IL1AssetRouter.contract.UnpackLog(event, "AssetHandlerRegisteredInitial", log); err != nil {
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

// ParseAssetHandlerRegisteredInitial is a log parse operation binding the contract event 0xb1e82bee3e85b2755fbceb4b7e051f5c66a7f35f0476657504e77e18ebd3a17d.
//
// Solidity: event AssetHandlerRegisteredInitial(bytes32 indexed assetId, address indexed assetHandlerAddress, bytes32 indexed additionalData, address assetDeploymentTracker)
func (_IL1AssetRouter *IL1AssetRouterFilterer) ParseAssetHandlerRegisteredInitial(log types.Log) (*IL1AssetRouterAssetHandlerRegisteredInitial, error) {
	event := new(IL1AssetRouterAssetHandlerRegisteredInitial)
	if err := _IL1AssetRouter.contract.UnpackLog(event, "AssetHandlerRegisteredInitial", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// IL1AssetRouterBridgehubDepositBaseTokenInitiatedIterator is returned from FilterBridgehubDepositBaseTokenInitiated and is used to iterate over the raw logs and unpacked data for BridgehubDepositBaseTokenInitiated events raised by the IL1AssetRouter contract.
type IL1AssetRouterBridgehubDepositBaseTokenInitiatedIterator struct {
	Event *IL1AssetRouterBridgehubDepositBaseTokenInitiated // Event containing the contract specifics and raw log

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
func (it *IL1AssetRouterBridgehubDepositBaseTokenInitiatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IL1AssetRouterBridgehubDepositBaseTokenInitiated)
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
		it.Event = new(IL1AssetRouterBridgehubDepositBaseTokenInitiated)
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
func (it *IL1AssetRouterBridgehubDepositBaseTokenInitiatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IL1AssetRouterBridgehubDepositBaseTokenInitiatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IL1AssetRouterBridgehubDepositBaseTokenInitiated represents a BridgehubDepositBaseTokenInitiated event raised by the IL1AssetRouter contract.
type IL1AssetRouterBridgehubDepositBaseTokenInitiated struct {
	ChainId *big.Int
	From    common.Address
	AssetId [32]byte
	Amount  *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterBridgehubDepositBaseTokenInitiated is a free log retrieval operation binding the contract event 0x0f87e1ea5eb1f034a6071ef630c174063e3d48756f853efaaf4292b929298240.
//
// Solidity: event BridgehubDepositBaseTokenInitiated(uint256 indexed chainId, address indexed from, bytes32 assetId, uint256 amount)
func (_IL1AssetRouter *IL1AssetRouterFilterer) FilterBridgehubDepositBaseTokenInitiated(opts *bind.FilterOpts, chainId []*big.Int, from []common.Address) (*IL1AssetRouterBridgehubDepositBaseTokenInitiatedIterator, error) {

	var chainIdRule []interface{}
	for _, chainIdItem := range chainId {
		chainIdRule = append(chainIdRule, chainIdItem)
	}
	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}

	logs, sub, err := _IL1AssetRouter.contract.FilterLogs(opts, "BridgehubDepositBaseTokenInitiated", chainIdRule, fromRule)
	if err != nil {
		return nil, err
	}
	return &IL1AssetRouterBridgehubDepositBaseTokenInitiatedIterator{contract: _IL1AssetRouter.contract, event: "BridgehubDepositBaseTokenInitiated", logs: logs, sub: sub}, nil
}

// WatchBridgehubDepositBaseTokenInitiated is a free log subscription operation binding the contract event 0x0f87e1ea5eb1f034a6071ef630c174063e3d48756f853efaaf4292b929298240.
//
// Solidity: event BridgehubDepositBaseTokenInitiated(uint256 indexed chainId, address indexed from, bytes32 assetId, uint256 amount)
func (_IL1AssetRouter *IL1AssetRouterFilterer) WatchBridgehubDepositBaseTokenInitiated(opts *bind.WatchOpts, sink chan<- *IL1AssetRouterBridgehubDepositBaseTokenInitiated, chainId []*big.Int, from []common.Address) (event.Subscription, error) {

	var chainIdRule []interface{}
	for _, chainIdItem := range chainId {
		chainIdRule = append(chainIdRule, chainIdItem)
	}
	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}

	logs, sub, err := _IL1AssetRouter.contract.WatchLogs(opts, "BridgehubDepositBaseTokenInitiated", chainIdRule, fromRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IL1AssetRouterBridgehubDepositBaseTokenInitiated)
				if err := _IL1AssetRouter.contract.UnpackLog(event, "BridgehubDepositBaseTokenInitiated", log); err != nil {
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

// ParseBridgehubDepositBaseTokenInitiated is a log parse operation binding the contract event 0x0f87e1ea5eb1f034a6071ef630c174063e3d48756f853efaaf4292b929298240.
//
// Solidity: event BridgehubDepositBaseTokenInitiated(uint256 indexed chainId, address indexed from, bytes32 assetId, uint256 amount)
func (_IL1AssetRouter *IL1AssetRouterFilterer) ParseBridgehubDepositBaseTokenInitiated(log types.Log) (*IL1AssetRouterBridgehubDepositBaseTokenInitiated, error) {
	event := new(IL1AssetRouterBridgehubDepositBaseTokenInitiated)
	if err := _IL1AssetRouter.contract.UnpackLog(event, "BridgehubDepositBaseTokenInitiated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// IL1AssetRouterBridgehubDepositFinalizedIterator is returned from FilterBridgehubDepositFinalized and is used to iterate over the raw logs and unpacked data for BridgehubDepositFinalized events raised by the IL1AssetRouter contract.
type IL1AssetRouterBridgehubDepositFinalizedIterator struct {
	Event *IL1AssetRouterBridgehubDepositFinalized // Event containing the contract specifics and raw log

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
func (it *IL1AssetRouterBridgehubDepositFinalizedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IL1AssetRouterBridgehubDepositFinalized)
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
		it.Event = new(IL1AssetRouterBridgehubDepositFinalized)
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
func (it *IL1AssetRouterBridgehubDepositFinalizedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IL1AssetRouterBridgehubDepositFinalizedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IL1AssetRouterBridgehubDepositFinalized represents a BridgehubDepositFinalized event raised by the IL1AssetRouter contract.
type IL1AssetRouterBridgehubDepositFinalized struct {
	ChainId         *big.Int
	TxDataHash      [32]byte
	L2DepositTxHash [32]byte
	Raw             types.Log // Blockchain specific contextual infos
}

// FilterBridgehubDepositFinalized is a free log retrieval operation binding the contract event 0xe4def01b981193a97a9e81230d7b9f31812ceaf23f864a828a82c687911cb2df.
//
// Solidity: event BridgehubDepositFinalized(uint256 indexed chainId, bytes32 indexed txDataHash, bytes32 indexed l2DepositTxHash)
func (_IL1AssetRouter *IL1AssetRouterFilterer) FilterBridgehubDepositFinalized(opts *bind.FilterOpts, chainId []*big.Int, txDataHash [][32]byte, l2DepositTxHash [][32]byte) (*IL1AssetRouterBridgehubDepositFinalizedIterator, error) {

	var chainIdRule []interface{}
	for _, chainIdItem := range chainId {
		chainIdRule = append(chainIdRule, chainIdItem)
	}
	var txDataHashRule []interface{}
	for _, txDataHashItem := range txDataHash {
		txDataHashRule = append(txDataHashRule, txDataHashItem)
	}
	var l2DepositTxHashRule []interface{}
	for _, l2DepositTxHashItem := range l2DepositTxHash {
		l2DepositTxHashRule = append(l2DepositTxHashRule, l2DepositTxHashItem)
	}

	logs, sub, err := _IL1AssetRouter.contract.FilterLogs(opts, "BridgehubDepositFinalized", chainIdRule, txDataHashRule, l2DepositTxHashRule)
	if err != nil {
		return nil, err
	}
	return &IL1AssetRouterBridgehubDepositFinalizedIterator{contract: _IL1AssetRouter.contract, event: "BridgehubDepositFinalized", logs: logs, sub: sub}, nil
}

// WatchBridgehubDepositFinalized is a free log subscription operation binding the contract event 0xe4def01b981193a97a9e81230d7b9f31812ceaf23f864a828a82c687911cb2df.
//
// Solidity: event BridgehubDepositFinalized(uint256 indexed chainId, bytes32 indexed txDataHash, bytes32 indexed l2DepositTxHash)
func (_IL1AssetRouter *IL1AssetRouterFilterer) WatchBridgehubDepositFinalized(opts *bind.WatchOpts, sink chan<- *IL1AssetRouterBridgehubDepositFinalized, chainId []*big.Int, txDataHash [][32]byte, l2DepositTxHash [][32]byte) (event.Subscription, error) {

	var chainIdRule []interface{}
	for _, chainIdItem := range chainId {
		chainIdRule = append(chainIdRule, chainIdItem)
	}
	var txDataHashRule []interface{}
	for _, txDataHashItem := range txDataHash {
		txDataHashRule = append(txDataHashRule, txDataHashItem)
	}
	var l2DepositTxHashRule []interface{}
	for _, l2DepositTxHashItem := range l2DepositTxHash {
		l2DepositTxHashRule = append(l2DepositTxHashRule, l2DepositTxHashItem)
	}

	logs, sub, err := _IL1AssetRouter.contract.WatchLogs(opts, "BridgehubDepositFinalized", chainIdRule, txDataHashRule, l2DepositTxHashRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IL1AssetRouterBridgehubDepositFinalized)
				if err := _IL1AssetRouter.contract.UnpackLog(event, "BridgehubDepositFinalized", log); err != nil {
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

// ParseBridgehubDepositFinalized is a log parse operation binding the contract event 0xe4def01b981193a97a9e81230d7b9f31812ceaf23f864a828a82c687911cb2df.
//
// Solidity: event BridgehubDepositFinalized(uint256 indexed chainId, bytes32 indexed txDataHash, bytes32 indexed l2DepositTxHash)
func (_IL1AssetRouter *IL1AssetRouterFilterer) ParseBridgehubDepositFinalized(log types.Log) (*IL1AssetRouterBridgehubDepositFinalized, error) {
	event := new(IL1AssetRouterBridgehubDepositFinalized)
	if err := _IL1AssetRouter.contract.UnpackLog(event, "BridgehubDepositFinalized", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// IL1AssetRouterBridgehubDepositInitiatedIterator is returned from FilterBridgehubDepositInitiated and is used to iterate over the raw logs and unpacked data for BridgehubDepositInitiated events raised by the IL1AssetRouter contract.
type IL1AssetRouterBridgehubDepositInitiatedIterator struct {
	Event *IL1AssetRouterBridgehubDepositInitiated // Event containing the contract specifics and raw log

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
func (it *IL1AssetRouterBridgehubDepositInitiatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IL1AssetRouterBridgehubDepositInitiated)
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
		it.Event = new(IL1AssetRouterBridgehubDepositInitiated)
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
func (it *IL1AssetRouterBridgehubDepositInitiatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IL1AssetRouterBridgehubDepositInitiatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IL1AssetRouterBridgehubDepositInitiated represents a BridgehubDepositInitiated event raised by the IL1AssetRouter contract.
type IL1AssetRouterBridgehubDepositInitiated struct {
	ChainId            *big.Int
	TxDataHash         [32]byte
	From               common.Address
	AssetId            [32]byte
	BridgeMintCalldata []byte
	Raw                types.Log // Blockchain specific contextual infos
}

// FilterBridgehubDepositInitiated is a free log retrieval operation binding the contract event 0xe21913bc89c1320d9709a5d236ffe06b54cf88aecfc9509ebd68f1adba45781e.
//
// Solidity: event BridgehubDepositInitiated(uint256 indexed chainId, bytes32 indexed txDataHash, address indexed from, bytes32 assetId, bytes bridgeMintCalldata)
func (_IL1AssetRouter *IL1AssetRouterFilterer) FilterBridgehubDepositInitiated(opts *bind.FilterOpts, chainId []*big.Int, txDataHash [][32]byte, from []common.Address) (*IL1AssetRouterBridgehubDepositInitiatedIterator, error) {

	var chainIdRule []interface{}
	for _, chainIdItem := range chainId {
		chainIdRule = append(chainIdRule, chainIdItem)
	}
	var txDataHashRule []interface{}
	for _, txDataHashItem := range txDataHash {
		txDataHashRule = append(txDataHashRule, txDataHashItem)
	}
	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}

	logs, sub, err := _IL1AssetRouter.contract.FilterLogs(opts, "BridgehubDepositInitiated", chainIdRule, txDataHashRule, fromRule)
	if err != nil {
		return nil, err
	}
	return &IL1AssetRouterBridgehubDepositInitiatedIterator{contract: _IL1AssetRouter.contract, event: "BridgehubDepositInitiated", logs: logs, sub: sub}, nil
}

// WatchBridgehubDepositInitiated is a free log subscription operation binding the contract event 0xe21913bc89c1320d9709a5d236ffe06b54cf88aecfc9509ebd68f1adba45781e.
//
// Solidity: event BridgehubDepositInitiated(uint256 indexed chainId, bytes32 indexed txDataHash, address indexed from, bytes32 assetId, bytes bridgeMintCalldata)
func (_IL1AssetRouter *IL1AssetRouterFilterer) WatchBridgehubDepositInitiated(opts *bind.WatchOpts, sink chan<- *IL1AssetRouterBridgehubDepositInitiated, chainId []*big.Int, txDataHash [][32]byte, from []common.Address) (event.Subscription, error) {

	var chainIdRule []interface{}
	for _, chainIdItem := range chainId {
		chainIdRule = append(chainIdRule, chainIdItem)
	}
	var txDataHashRule []interface{}
	for _, txDataHashItem := range txDataHash {
		txDataHashRule = append(txDataHashRule, txDataHashItem)
	}
	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}

	logs, sub, err := _IL1AssetRouter.contract.WatchLogs(opts, "BridgehubDepositInitiated", chainIdRule, txDataHashRule, fromRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IL1AssetRouterBridgehubDepositInitiated)
				if err := _IL1AssetRouter.contract.UnpackLog(event, "BridgehubDepositInitiated", log); err != nil {
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

// ParseBridgehubDepositInitiated is a log parse operation binding the contract event 0xe21913bc89c1320d9709a5d236ffe06b54cf88aecfc9509ebd68f1adba45781e.
//
// Solidity: event BridgehubDepositInitiated(uint256 indexed chainId, bytes32 indexed txDataHash, address indexed from, bytes32 assetId, bytes bridgeMintCalldata)
func (_IL1AssetRouter *IL1AssetRouterFilterer) ParseBridgehubDepositInitiated(log types.Log) (*IL1AssetRouterBridgehubDepositInitiated, error) {
	event := new(IL1AssetRouterBridgehubDepositInitiated)
	if err := _IL1AssetRouter.contract.UnpackLog(event, "BridgehubDepositInitiated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// IL1AssetRouterBridgehubMintDataIterator is returned from FilterBridgehubMintData and is used to iterate over the raw logs and unpacked data for BridgehubMintData events raised by the IL1AssetRouter contract.
type IL1AssetRouterBridgehubMintDataIterator struct {
	Event *IL1AssetRouterBridgehubMintData // Event containing the contract specifics and raw log

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
func (it *IL1AssetRouterBridgehubMintDataIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IL1AssetRouterBridgehubMintData)
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
		it.Event = new(IL1AssetRouterBridgehubMintData)
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
func (it *IL1AssetRouterBridgehubMintDataIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IL1AssetRouterBridgehubMintDataIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IL1AssetRouterBridgehubMintData represents a BridgehubMintData event raised by the IL1AssetRouter contract.
type IL1AssetRouterBridgehubMintData struct {
	BridgeMintData []byte
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterBridgehubMintData is a free log retrieval operation binding the contract event 0x31a15cb4f69820f57afabeaff74feae31dc25875c07c952ba742a3acf8690f91.
//
// Solidity: event BridgehubMintData(bytes bridgeMintData)
func (_IL1AssetRouter *IL1AssetRouterFilterer) FilterBridgehubMintData(opts *bind.FilterOpts) (*IL1AssetRouterBridgehubMintDataIterator, error) {

	logs, sub, err := _IL1AssetRouter.contract.FilterLogs(opts, "BridgehubMintData")
	if err != nil {
		return nil, err
	}
	return &IL1AssetRouterBridgehubMintDataIterator{contract: _IL1AssetRouter.contract, event: "BridgehubMintData", logs: logs, sub: sub}, nil
}

// WatchBridgehubMintData is a free log subscription operation binding the contract event 0x31a15cb4f69820f57afabeaff74feae31dc25875c07c952ba742a3acf8690f91.
//
// Solidity: event BridgehubMintData(bytes bridgeMintData)
func (_IL1AssetRouter *IL1AssetRouterFilterer) WatchBridgehubMintData(opts *bind.WatchOpts, sink chan<- *IL1AssetRouterBridgehubMintData) (event.Subscription, error) {

	logs, sub, err := _IL1AssetRouter.contract.WatchLogs(opts, "BridgehubMintData")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IL1AssetRouterBridgehubMintData)
				if err := _IL1AssetRouter.contract.UnpackLog(event, "BridgehubMintData", log); err != nil {
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

// ParseBridgehubMintData is a log parse operation binding the contract event 0x31a15cb4f69820f57afabeaff74feae31dc25875c07c952ba742a3acf8690f91.
//
// Solidity: event BridgehubMintData(bytes bridgeMintData)
func (_IL1AssetRouter *IL1AssetRouterFilterer) ParseBridgehubMintData(log types.Log) (*IL1AssetRouterBridgehubMintData, error) {
	event := new(IL1AssetRouterBridgehubMintData)
	if err := _IL1AssetRouter.contract.UnpackLog(event, "BridgehubMintData", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// IL1AssetRouterBridgehubWithdrawalInitiatedIterator is returned from FilterBridgehubWithdrawalInitiated and is used to iterate over the raw logs and unpacked data for BridgehubWithdrawalInitiated events raised by the IL1AssetRouter contract.
type IL1AssetRouterBridgehubWithdrawalInitiatedIterator struct {
	Event *IL1AssetRouterBridgehubWithdrawalInitiated // Event containing the contract specifics and raw log

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
func (it *IL1AssetRouterBridgehubWithdrawalInitiatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IL1AssetRouterBridgehubWithdrawalInitiated)
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
		it.Event = new(IL1AssetRouterBridgehubWithdrawalInitiated)
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
func (it *IL1AssetRouterBridgehubWithdrawalInitiatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IL1AssetRouterBridgehubWithdrawalInitiatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IL1AssetRouterBridgehubWithdrawalInitiated represents a BridgehubWithdrawalInitiated event raised by the IL1AssetRouter contract.
type IL1AssetRouterBridgehubWithdrawalInitiated struct {
	ChainId       *big.Int
	Sender        common.Address
	AssetId       [32]byte
	AssetDataHash [32]byte
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterBridgehubWithdrawalInitiated is a free log retrieval operation binding the contract event 0x9a3d4025b7294a1754ea5b56309c1e72328d97b73718183db595c850d14a3ae0.
//
// Solidity: event BridgehubWithdrawalInitiated(uint256 chainId, address indexed sender, bytes32 indexed assetId, bytes32 assetDataHash)
func (_IL1AssetRouter *IL1AssetRouterFilterer) FilterBridgehubWithdrawalInitiated(opts *bind.FilterOpts, sender []common.Address, assetId [][32]byte) (*IL1AssetRouterBridgehubWithdrawalInitiatedIterator, error) {

	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}
	var assetIdRule []interface{}
	for _, assetIdItem := range assetId {
		assetIdRule = append(assetIdRule, assetIdItem)
	}

	logs, sub, err := _IL1AssetRouter.contract.FilterLogs(opts, "BridgehubWithdrawalInitiated", senderRule, assetIdRule)
	if err != nil {
		return nil, err
	}
	return &IL1AssetRouterBridgehubWithdrawalInitiatedIterator{contract: _IL1AssetRouter.contract, event: "BridgehubWithdrawalInitiated", logs: logs, sub: sub}, nil
}

// WatchBridgehubWithdrawalInitiated is a free log subscription operation binding the contract event 0x9a3d4025b7294a1754ea5b56309c1e72328d97b73718183db595c850d14a3ae0.
//
// Solidity: event BridgehubWithdrawalInitiated(uint256 chainId, address indexed sender, bytes32 indexed assetId, bytes32 assetDataHash)
func (_IL1AssetRouter *IL1AssetRouterFilterer) WatchBridgehubWithdrawalInitiated(opts *bind.WatchOpts, sink chan<- *IL1AssetRouterBridgehubWithdrawalInitiated, sender []common.Address, assetId [][32]byte) (event.Subscription, error) {

	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}
	var assetIdRule []interface{}
	for _, assetIdItem := range assetId {
		assetIdRule = append(assetIdRule, assetIdItem)
	}

	logs, sub, err := _IL1AssetRouter.contract.WatchLogs(opts, "BridgehubWithdrawalInitiated", senderRule, assetIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IL1AssetRouterBridgehubWithdrawalInitiated)
				if err := _IL1AssetRouter.contract.UnpackLog(event, "BridgehubWithdrawalInitiated", log); err != nil {
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

// ParseBridgehubWithdrawalInitiated is a log parse operation binding the contract event 0x9a3d4025b7294a1754ea5b56309c1e72328d97b73718183db595c850d14a3ae0.
//
// Solidity: event BridgehubWithdrawalInitiated(uint256 chainId, address indexed sender, bytes32 indexed assetId, bytes32 assetDataHash)
func (_IL1AssetRouter *IL1AssetRouterFilterer) ParseBridgehubWithdrawalInitiated(log types.Log) (*IL1AssetRouterBridgehubWithdrawalInitiated, error) {
	event := new(IL1AssetRouterBridgehubWithdrawalInitiated)
	if err := _IL1AssetRouter.contract.UnpackLog(event, "BridgehubWithdrawalInitiated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// IL1AssetRouterClaimedFailedDepositAssetRouterIterator is returned from FilterClaimedFailedDepositAssetRouter and is used to iterate over the raw logs and unpacked data for ClaimedFailedDepositAssetRouter events raised by the IL1AssetRouter contract.
type IL1AssetRouterClaimedFailedDepositAssetRouterIterator struct {
	Event *IL1AssetRouterClaimedFailedDepositAssetRouter // Event containing the contract specifics and raw log

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
func (it *IL1AssetRouterClaimedFailedDepositAssetRouterIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IL1AssetRouterClaimedFailedDepositAssetRouter)
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
		it.Event = new(IL1AssetRouterClaimedFailedDepositAssetRouter)
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
func (it *IL1AssetRouterClaimedFailedDepositAssetRouterIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IL1AssetRouterClaimedFailedDepositAssetRouterIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IL1AssetRouterClaimedFailedDepositAssetRouter represents a ClaimedFailedDepositAssetRouter event raised by the IL1AssetRouter contract.
type IL1AssetRouterClaimedFailedDepositAssetRouter struct {
	ChainId   *big.Int
	AssetId   [32]byte
	AssetData []byte
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterClaimedFailedDepositAssetRouter is a free log retrieval operation binding the contract event 0x4250817d22c13fba8067153d85ccd9706326ac2bd14d5c3898c8b1bccc440658.
//
// Solidity: event ClaimedFailedDepositAssetRouter(uint256 indexed chainId, bytes32 indexed assetId, bytes assetData)
func (_IL1AssetRouter *IL1AssetRouterFilterer) FilterClaimedFailedDepositAssetRouter(opts *bind.FilterOpts, chainId []*big.Int, assetId [][32]byte) (*IL1AssetRouterClaimedFailedDepositAssetRouterIterator, error) {

	var chainIdRule []interface{}
	for _, chainIdItem := range chainId {
		chainIdRule = append(chainIdRule, chainIdItem)
	}
	var assetIdRule []interface{}
	for _, assetIdItem := range assetId {
		assetIdRule = append(assetIdRule, assetIdItem)
	}

	logs, sub, err := _IL1AssetRouter.contract.FilterLogs(opts, "ClaimedFailedDepositAssetRouter", chainIdRule, assetIdRule)
	if err != nil {
		return nil, err
	}
	return &IL1AssetRouterClaimedFailedDepositAssetRouterIterator{contract: _IL1AssetRouter.contract, event: "ClaimedFailedDepositAssetRouter", logs: logs, sub: sub}, nil
}

// WatchClaimedFailedDepositAssetRouter is a free log subscription operation binding the contract event 0x4250817d22c13fba8067153d85ccd9706326ac2bd14d5c3898c8b1bccc440658.
//
// Solidity: event ClaimedFailedDepositAssetRouter(uint256 indexed chainId, bytes32 indexed assetId, bytes assetData)
func (_IL1AssetRouter *IL1AssetRouterFilterer) WatchClaimedFailedDepositAssetRouter(opts *bind.WatchOpts, sink chan<- *IL1AssetRouterClaimedFailedDepositAssetRouter, chainId []*big.Int, assetId [][32]byte) (event.Subscription, error) {

	var chainIdRule []interface{}
	for _, chainIdItem := range chainId {
		chainIdRule = append(chainIdRule, chainIdItem)
	}
	var assetIdRule []interface{}
	for _, assetIdItem := range assetId {
		assetIdRule = append(assetIdRule, assetIdItem)
	}

	logs, sub, err := _IL1AssetRouter.contract.WatchLogs(opts, "ClaimedFailedDepositAssetRouter", chainIdRule, assetIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IL1AssetRouterClaimedFailedDepositAssetRouter)
				if err := _IL1AssetRouter.contract.UnpackLog(event, "ClaimedFailedDepositAssetRouter", log); err != nil {
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

// ParseClaimedFailedDepositAssetRouter is a log parse operation binding the contract event 0x4250817d22c13fba8067153d85ccd9706326ac2bd14d5c3898c8b1bccc440658.
//
// Solidity: event ClaimedFailedDepositAssetRouter(uint256 indexed chainId, bytes32 indexed assetId, bytes assetData)
func (_IL1AssetRouter *IL1AssetRouterFilterer) ParseClaimedFailedDepositAssetRouter(log types.Log) (*IL1AssetRouterClaimedFailedDepositAssetRouter, error) {
	event := new(IL1AssetRouterClaimedFailedDepositAssetRouter)
	if err := _IL1AssetRouter.contract.UnpackLog(event, "ClaimedFailedDepositAssetRouter", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// IL1AssetRouterDepositFinalizedAssetRouterIterator is returned from FilterDepositFinalizedAssetRouter and is used to iterate over the raw logs and unpacked data for DepositFinalizedAssetRouter events raised by the IL1AssetRouter contract.
type IL1AssetRouterDepositFinalizedAssetRouterIterator struct {
	Event *IL1AssetRouterDepositFinalizedAssetRouter // Event containing the contract specifics and raw log

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
func (it *IL1AssetRouterDepositFinalizedAssetRouterIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IL1AssetRouterDepositFinalizedAssetRouter)
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
		it.Event = new(IL1AssetRouterDepositFinalizedAssetRouter)
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
func (it *IL1AssetRouterDepositFinalizedAssetRouterIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IL1AssetRouterDepositFinalizedAssetRouterIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IL1AssetRouterDepositFinalizedAssetRouter represents a DepositFinalizedAssetRouter event raised by the IL1AssetRouter contract.
type IL1AssetRouterDepositFinalizedAssetRouter struct {
	ChainId   *big.Int
	AssetId   [32]byte
	AssetData []byte
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterDepositFinalizedAssetRouter is a free log retrieval operation binding the contract event 0x44eb9a840094a49b3cd0a5205042598a1c08c4e87bafb5760bc2d8efa170c541.
//
// Solidity: event DepositFinalizedAssetRouter(uint256 indexed chainId, bytes32 indexed assetId, bytes assetData)
func (_IL1AssetRouter *IL1AssetRouterFilterer) FilterDepositFinalizedAssetRouter(opts *bind.FilterOpts, chainId []*big.Int, assetId [][32]byte) (*IL1AssetRouterDepositFinalizedAssetRouterIterator, error) {

	var chainIdRule []interface{}
	for _, chainIdItem := range chainId {
		chainIdRule = append(chainIdRule, chainIdItem)
	}
	var assetIdRule []interface{}
	for _, assetIdItem := range assetId {
		assetIdRule = append(assetIdRule, assetIdItem)
	}

	logs, sub, err := _IL1AssetRouter.contract.FilterLogs(opts, "DepositFinalizedAssetRouter", chainIdRule, assetIdRule)
	if err != nil {
		return nil, err
	}
	return &IL1AssetRouterDepositFinalizedAssetRouterIterator{contract: _IL1AssetRouter.contract, event: "DepositFinalizedAssetRouter", logs: logs, sub: sub}, nil
}

// WatchDepositFinalizedAssetRouter is a free log subscription operation binding the contract event 0x44eb9a840094a49b3cd0a5205042598a1c08c4e87bafb5760bc2d8efa170c541.
//
// Solidity: event DepositFinalizedAssetRouter(uint256 indexed chainId, bytes32 indexed assetId, bytes assetData)
func (_IL1AssetRouter *IL1AssetRouterFilterer) WatchDepositFinalizedAssetRouter(opts *bind.WatchOpts, sink chan<- *IL1AssetRouterDepositFinalizedAssetRouter, chainId []*big.Int, assetId [][32]byte) (event.Subscription, error) {

	var chainIdRule []interface{}
	for _, chainIdItem := range chainId {
		chainIdRule = append(chainIdRule, chainIdItem)
	}
	var assetIdRule []interface{}
	for _, assetIdItem := range assetId {
		assetIdRule = append(assetIdRule, assetIdItem)
	}

	logs, sub, err := _IL1AssetRouter.contract.WatchLogs(opts, "DepositFinalizedAssetRouter", chainIdRule, assetIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IL1AssetRouterDepositFinalizedAssetRouter)
				if err := _IL1AssetRouter.contract.UnpackLog(event, "DepositFinalizedAssetRouter", log); err != nil {
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

// ParseDepositFinalizedAssetRouter is a log parse operation binding the contract event 0x44eb9a840094a49b3cd0a5205042598a1c08c4e87bafb5760bc2d8efa170c541.
//
// Solidity: event DepositFinalizedAssetRouter(uint256 indexed chainId, bytes32 indexed assetId, bytes assetData)
func (_IL1AssetRouter *IL1AssetRouterFilterer) ParseDepositFinalizedAssetRouter(log types.Log) (*IL1AssetRouterDepositFinalizedAssetRouter, error) {
	event := new(IL1AssetRouterDepositFinalizedAssetRouter)
	if err := _IL1AssetRouter.contract.UnpackLog(event, "DepositFinalizedAssetRouter", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// IL1AssetRouterInitializedIterator is returned from FilterInitialized and is used to iterate over the raw logs and unpacked data for Initialized events raised by the IL1AssetRouter contract.
type IL1AssetRouterInitializedIterator struct {
	Event *IL1AssetRouterInitialized // Event containing the contract specifics and raw log

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
func (it *IL1AssetRouterInitializedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IL1AssetRouterInitialized)
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
		it.Event = new(IL1AssetRouterInitialized)
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
func (it *IL1AssetRouterInitializedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IL1AssetRouterInitializedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IL1AssetRouterInitialized represents a Initialized event raised by the IL1AssetRouter contract.
type IL1AssetRouterInitialized struct {
	Version uint8
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterInitialized is a free log retrieval operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_IL1AssetRouter *IL1AssetRouterFilterer) FilterInitialized(opts *bind.FilterOpts) (*IL1AssetRouterInitializedIterator, error) {

	logs, sub, err := _IL1AssetRouter.contract.FilterLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return &IL1AssetRouterInitializedIterator{contract: _IL1AssetRouter.contract, event: "Initialized", logs: logs, sub: sub}, nil
}

// WatchInitialized is a free log subscription operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_IL1AssetRouter *IL1AssetRouterFilterer) WatchInitialized(opts *bind.WatchOpts, sink chan<- *IL1AssetRouterInitialized) (event.Subscription, error) {

	logs, sub, err := _IL1AssetRouter.contract.WatchLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IL1AssetRouterInitialized)
				if err := _IL1AssetRouter.contract.UnpackLog(event, "Initialized", log); err != nil {
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

// ParseInitialized is a log parse operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_IL1AssetRouter *IL1AssetRouterFilterer) ParseInitialized(log types.Log) (*IL1AssetRouterInitialized, error) {
	event := new(IL1AssetRouterInitialized)
	if err := _IL1AssetRouter.contract.UnpackLog(event, "Initialized", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// IL1AssetRouterLegacyDepositInitiatedIterator is returned from FilterLegacyDepositInitiated and is used to iterate over the raw logs and unpacked data for LegacyDepositInitiated events raised by the IL1AssetRouter contract.
type IL1AssetRouterLegacyDepositInitiatedIterator struct {
	Event *IL1AssetRouterLegacyDepositInitiated // Event containing the contract specifics and raw log

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
func (it *IL1AssetRouterLegacyDepositInitiatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IL1AssetRouterLegacyDepositInitiated)
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
		it.Event = new(IL1AssetRouterLegacyDepositInitiated)
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
func (it *IL1AssetRouterLegacyDepositInitiatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IL1AssetRouterLegacyDepositInitiatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IL1AssetRouterLegacyDepositInitiated represents a LegacyDepositInitiated event raised by the IL1AssetRouter contract.
type IL1AssetRouterLegacyDepositInitiated struct {
	ChainId         *big.Int
	L2DepositTxHash [32]byte
	From            common.Address
	To              common.Address
	L1Asset         common.Address
	Amount          *big.Int
	Raw             types.Log // Blockchain specific contextual infos
}

// FilterLegacyDepositInitiated is a free log retrieval operation binding the contract event 0xa1846a4248529db592da99da276f761d9f37a84d0f3d4e83819b869759000700.
//
// Solidity: event LegacyDepositInitiated(uint256 indexed chainId, bytes32 indexed l2DepositTxHash, address indexed from, address to, address l1Asset, uint256 amount)
func (_IL1AssetRouter *IL1AssetRouterFilterer) FilterLegacyDepositInitiated(opts *bind.FilterOpts, chainId []*big.Int, l2DepositTxHash [][32]byte, from []common.Address) (*IL1AssetRouterLegacyDepositInitiatedIterator, error) {

	var chainIdRule []interface{}
	for _, chainIdItem := range chainId {
		chainIdRule = append(chainIdRule, chainIdItem)
	}
	var l2DepositTxHashRule []interface{}
	for _, l2DepositTxHashItem := range l2DepositTxHash {
		l2DepositTxHashRule = append(l2DepositTxHashRule, l2DepositTxHashItem)
	}
	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}

	logs, sub, err := _IL1AssetRouter.contract.FilterLogs(opts, "LegacyDepositInitiated", chainIdRule, l2DepositTxHashRule, fromRule)
	if err != nil {
		return nil, err
	}
	return &IL1AssetRouterLegacyDepositInitiatedIterator{contract: _IL1AssetRouter.contract, event: "LegacyDepositInitiated", logs: logs, sub: sub}, nil
}

// WatchLegacyDepositInitiated is a free log subscription operation binding the contract event 0xa1846a4248529db592da99da276f761d9f37a84d0f3d4e83819b869759000700.
//
// Solidity: event LegacyDepositInitiated(uint256 indexed chainId, bytes32 indexed l2DepositTxHash, address indexed from, address to, address l1Asset, uint256 amount)
func (_IL1AssetRouter *IL1AssetRouterFilterer) WatchLegacyDepositInitiated(opts *bind.WatchOpts, sink chan<- *IL1AssetRouterLegacyDepositInitiated, chainId []*big.Int, l2DepositTxHash [][32]byte, from []common.Address) (event.Subscription, error) {

	var chainIdRule []interface{}
	for _, chainIdItem := range chainId {
		chainIdRule = append(chainIdRule, chainIdItem)
	}
	var l2DepositTxHashRule []interface{}
	for _, l2DepositTxHashItem := range l2DepositTxHash {
		l2DepositTxHashRule = append(l2DepositTxHashRule, l2DepositTxHashItem)
	}
	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}

	logs, sub, err := _IL1AssetRouter.contract.WatchLogs(opts, "LegacyDepositInitiated", chainIdRule, l2DepositTxHashRule, fromRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IL1AssetRouterLegacyDepositInitiated)
				if err := _IL1AssetRouter.contract.UnpackLog(event, "LegacyDepositInitiated", log); err != nil {
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

// ParseLegacyDepositInitiated is a log parse operation binding the contract event 0xa1846a4248529db592da99da276f761d9f37a84d0f3d4e83819b869759000700.
//
// Solidity: event LegacyDepositInitiated(uint256 indexed chainId, bytes32 indexed l2DepositTxHash, address indexed from, address to, address l1Asset, uint256 amount)
func (_IL1AssetRouter *IL1AssetRouterFilterer) ParseLegacyDepositInitiated(log types.Log) (*IL1AssetRouterLegacyDepositInitiated, error) {
	event := new(IL1AssetRouterLegacyDepositInitiated)
	if err := _IL1AssetRouter.contract.UnpackLog(event, "LegacyDepositInitiated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// IL1AssetRouterOwnershipTransferStartedIterator is returned from FilterOwnershipTransferStarted and is used to iterate over the raw logs and unpacked data for OwnershipTransferStarted events raised by the IL1AssetRouter contract.
type IL1AssetRouterOwnershipTransferStartedIterator struct {
	Event *IL1AssetRouterOwnershipTransferStarted // Event containing the contract specifics and raw log

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
func (it *IL1AssetRouterOwnershipTransferStartedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IL1AssetRouterOwnershipTransferStarted)
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
		it.Event = new(IL1AssetRouterOwnershipTransferStarted)
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
func (it *IL1AssetRouterOwnershipTransferStartedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IL1AssetRouterOwnershipTransferStartedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IL1AssetRouterOwnershipTransferStarted represents a OwnershipTransferStarted event raised by the IL1AssetRouter contract.
type IL1AssetRouterOwnershipTransferStarted struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferStarted is a free log retrieval operation binding the contract event 0x38d16b8cac22d99fc7c124b9cd0de2d3fa1faef420bfe791d8c362d765e22700.
//
// Solidity: event OwnershipTransferStarted(address indexed previousOwner, address indexed newOwner)
func (_IL1AssetRouter *IL1AssetRouterFilterer) FilterOwnershipTransferStarted(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*IL1AssetRouterOwnershipTransferStartedIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _IL1AssetRouter.contract.FilterLogs(opts, "OwnershipTransferStarted", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &IL1AssetRouterOwnershipTransferStartedIterator{contract: _IL1AssetRouter.contract, event: "OwnershipTransferStarted", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferStarted is a free log subscription operation binding the contract event 0x38d16b8cac22d99fc7c124b9cd0de2d3fa1faef420bfe791d8c362d765e22700.
//
// Solidity: event OwnershipTransferStarted(address indexed previousOwner, address indexed newOwner)
func (_IL1AssetRouter *IL1AssetRouterFilterer) WatchOwnershipTransferStarted(opts *bind.WatchOpts, sink chan<- *IL1AssetRouterOwnershipTransferStarted, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _IL1AssetRouter.contract.WatchLogs(opts, "OwnershipTransferStarted", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IL1AssetRouterOwnershipTransferStarted)
				if err := _IL1AssetRouter.contract.UnpackLog(event, "OwnershipTransferStarted", log); err != nil {
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

// ParseOwnershipTransferStarted is a log parse operation binding the contract event 0x38d16b8cac22d99fc7c124b9cd0de2d3fa1faef420bfe791d8c362d765e22700.
//
// Solidity: event OwnershipTransferStarted(address indexed previousOwner, address indexed newOwner)
func (_IL1AssetRouter *IL1AssetRouterFilterer) ParseOwnershipTransferStarted(log types.Log) (*IL1AssetRouterOwnershipTransferStarted, error) {
	event := new(IL1AssetRouterOwnershipTransferStarted)
	if err := _IL1AssetRouter.contract.UnpackLog(event, "OwnershipTransferStarted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// IL1AssetRouterOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the IL1AssetRouter contract.
type IL1AssetRouterOwnershipTransferredIterator struct {
	Event *IL1AssetRouterOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *IL1AssetRouterOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IL1AssetRouterOwnershipTransferred)
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
		it.Event = new(IL1AssetRouterOwnershipTransferred)
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
func (it *IL1AssetRouterOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IL1AssetRouterOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IL1AssetRouterOwnershipTransferred represents a OwnershipTransferred event raised by the IL1AssetRouter contract.
type IL1AssetRouterOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_IL1AssetRouter *IL1AssetRouterFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*IL1AssetRouterOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _IL1AssetRouter.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &IL1AssetRouterOwnershipTransferredIterator{contract: _IL1AssetRouter.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_IL1AssetRouter *IL1AssetRouterFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *IL1AssetRouterOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _IL1AssetRouter.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IL1AssetRouterOwnershipTransferred)
				if err := _IL1AssetRouter.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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
func (_IL1AssetRouter *IL1AssetRouterFilterer) ParseOwnershipTransferred(log types.Log) (*IL1AssetRouterOwnershipTransferred, error) {
	event := new(IL1AssetRouterOwnershipTransferred)
	if err := _IL1AssetRouter.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// IL1AssetRouterPausedIterator is returned from FilterPaused and is used to iterate over the raw logs and unpacked data for Paused events raised by the IL1AssetRouter contract.
type IL1AssetRouterPausedIterator struct {
	Event *IL1AssetRouterPaused // Event containing the contract specifics and raw log

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
func (it *IL1AssetRouterPausedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IL1AssetRouterPaused)
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
		it.Event = new(IL1AssetRouterPaused)
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
func (it *IL1AssetRouterPausedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IL1AssetRouterPausedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IL1AssetRouterPaused represents a Paused event raised by the IL1AssetRouter contract.
type IL1AssetRouterPaused struct {
	Account common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterPaused is a free log retrieval operation binding the contract event 0x62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a258.
//
// Solidity: event Paused(address account)
func (_IL1AssetRouter *IL1AssetRouterFilterer) FilterPaused(opts *bind.FilterOpts) (*IL1AssetRouterPausedIterator, error) {

	logs, sub, err := _IL1AssetRouter.contract.FilterLogs(opts, "Paused")
	if err != nil {
		return nil, err
	}
	return &IL1AssetRouterPausedIterator{contract: _IL1AssetRouter.contract, event: "Paused", logs: logs, sub: sub}, nil
}

// WatchPaused is a free log subscription operation binding the contract event 0x62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a258.
//
// Solidity: event Paused(address account)
func (_IL1AssetRouter *IL1AssetRouterFilterer) WatchPaused(opts *bind.WatchOpts, sink chan<- *IL1AssetRouterPaused) (event.Subscription, error) {

	logs, sub, err := _IL1AssetRouter.contract.WatchLogs(opts, "Paused")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IL1AssetRouterPaused)
				if err := _IL1AssetRouter.contract.UnpackLog(event, "Paused", log); err != nil {
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

// ParsePaused is a log parse operation binding the contract event 0x62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a258.
//
// Solidity: event Paused(address account)
func (_IL1AssetRouter *IL1AssetRouterFilterer) ParsePaused(log types.Log) (*IL1AssetRouterPaused, error) {
	event := new(IL1AssetRouterPaused)
	if err := _IL1AssetRouter.contract.UnpackLog(event, "Paused", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// IL1AssetRouterUnpausedIterator is returned from FilterUnpaused and is used to iterate over the raw logs and unpacked data for Unpaused events raised by the IL1AssetRouter contract.
type IL1AssetRouterUnpausedIterator struct {
	Event *IL1AssetRouterUnpaused // Event containing the contract specifics and raw log

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
func (it *IL1AssetRouterUnpausedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IL1AssetRouterUnpaused)
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
		it.Event = new(IL1AssetRouterUnpaused)
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
func (it *IL1AssetRouterUnpausedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IL1AssetRouterUnpausedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IL1AssetRouterUnpaused represents a Unpaused event raised by the IL1AssetRouter contract.
type IL1AssetRouterUnpaused struct {
	Account common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterUnpaused is a free log retrieval operation binding the contract event 0x5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa.
//
// Solidity: event Unpaused(address account)
func (_IL1AssetRouter *IL1AssetRouterFilterer) FilterUnpaused(opts *bind.FilterOpts) (*IL1AssetRouterUnpausedIterator, error) {

	logs, sub, err := _IL1AssetRouter.contract.FilterLogs(opts, "Unpaused")
	if err != nil {
		return nil, err
	}
	return &IL1AssetRouterUnpausedIterator{contract: _IL1AssetRouter.contract, event: "Unpaused", logs: logs, sub: sub}, nil
}

// WatchUnpaused is a free log subscription operation binding the contract event 0x5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa.
//
// Solidity: event Unpaused(address account)
func (_IL1AssetRouter *IL1AssetRouterFilterer) WatchUnpaused(opts *bind.WatchOpts, sink chan<- *IL1AssetRouterUnpaused) (event.Subscription, error) {

	logs, sub, err := _IL1AssetRouter.contract.WatchLogs(opts, "Unpaused")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IL1AssetRouterUnpaused)
				if err := _IL1AssetRouter.contract.UnpackLog(event, "Unpaused", log); err != nil {
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

// ParseUnpaused is a log parse operation binding the contract event 0x5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa.
//
// Solidity: event Unpaused(address account)
func (_IL1AssetRouter *IL1AssetRouterFilterer) ParseUnpaused(log types.Log) (*IL1AssetRouterUnpaused, error) {
	event := new(IL1AssetRouterUnpaused)
	if err := _IL1AssetRouter.contract.UnpackLog(event, "Unpaused", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
