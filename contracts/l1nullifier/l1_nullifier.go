// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package l1nullifier

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

// FinalizeL1DepositParams is an auto generated low-level Go binding around an user-defined struct.
type FinalizeL1DepositParams struct {
	ChainId           *big.Int
	L2BatchNumber     *big.Int
	L2MessageIndex    *big.Int
	L2Sender          common.Address
	L2TxNumberInBatch uint16
	Message           []byte
	MerkleProof       [][32]byte
}

// IL1NullifierMetaData contains all meta data concerning the IL1Nullifier contract.
var IL1NullifierMetaData = &bind.MetaData{
	ABI: "[{\"type\":\"constructor\",\"inputs\":[{\"name\":\"_bridgehub\",\"type\":\"address\",\"internalType\":\"contractIBridgehub\"},{\"name\":\"_interopCenter\",\"type\":\"address\",\"internalType\":\"contractIInteropCenter\"},{\"name\":\"_eraChainId\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"_eraDiamondProxy\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"BRIDGE_HUB\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"contractIBridgehub\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"__DEPRECATED_admin\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"__DEPRECATED_chainBalance\",\"inputs\":[{\"name\":\"chainId\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"l1Token\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"balance\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"__DEPRECATED_l2BridgeAddress\",\"inputs\":[{\"name\":\"chainId\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"l2Bridge\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"__DEPRECATED_pendingAdmin\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"acceptOwnership\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"bridgeRecoverFailedTransfer\",\"inputs\":[{\"name\":\"_chainId\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"_depositSender\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"_assetId\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"_assetData\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"_l2TxHash\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"_l2BatchNumber\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"_l2MessageIndex\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"_l2TxNumberInBatch\",\"type\":\"uint16\",\"internalType\":\"uint16\"},{\"name\":\"_merkleProof\",\"type\":\"bytes32[]\",\"internalType\":\"bytes32[]\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"bridgehubConfirmL2TransactionForwarded\",\"inputs\":[{\"name\":\"_chainId\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"_txDataHash\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"_txHash\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"chainBalance\",\"inputs\":[{\"name\":\"_chainId\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"_token\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"claimFailedDeposit\",\"inputs\":[{\"name\":\"_chainId\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"_depositSender\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"_l1Token\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"_amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"_l2TxHash\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"_l2BatchNumber\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"_l2MessageIndex\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"_l2TxNumberInBatch\",\"type\":\"uint16\",\"internalType\":\"uint16\"},{\"name\":\"_merkleProof\",\"type\":\"bytes32[]\",\"internalType\":\"bytes32[]\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"claimFailedDepositLegacyErc20Bridge\",\"inputs\":[{\"name\":\"_depositSender\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"_l1Token\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"_amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"_l2TxHash\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"_l2BatchNumber\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"_l2MessageIndex\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"_l2TxNumberInBatch\",\"type\":\"uint16\",\"internalType\":\"uint16\"},{\"name\":\"_merkleProof\",\"type\":\"bytes32[]\",\"internalType\":\"bytes32[]\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"depositHappened\",\"inputs\":[{\"name\":\"chainId\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"l2DepositTxHash\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"outputs\":[{\"name\":\"depositDataHash\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"encodeTxDataHash\",\"inputs\":[{\"name\":\"_encodingVersion\",\"type\":\"bytes1\",\"internalType\":\"bytes1\"},{\"name\":\"_originalCaller\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"_assetId\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"_transferData\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[{\"name\":\"txDataHash\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"finalizeDeposit\",\"inputs\":[{\"name\":\"_finalizeWithdrawalParams\",\"type\":\"tuple\",\"internalType\":\"structFinalizeL1DepositParams\",\"components\":[{\"name\":\"chainId\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"l2BatchNumber\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"l2MessageIndex\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"l2Sender\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"l2TxNumberInBatch\",\"type\":\"uint16\",\"internalType\":\"uint16\"},{\"name\":\"message\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"merkleProof\",\"type\":\"bytes32[]\",\"internalType\":\"bytes32[]\"}]}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"finalizeWithdrawal\",\"inputs\":[{\"name\":\"_chainId\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"_l2BatchNumber\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"_l2MessageIndex\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"_l2TxNumberInBatch\",\"type\":\"uint16\",\"internalType\":\"uint16\"},{\"name\":\"_message\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"_merkleProof\",\"type\":\"bytes32[]\",\"internalType\":\"bytes32[]\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"initialize\",\"inputs\":[{\"name\":\"_owner\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"_eraPostDiamondUpgradeFirstBatch\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"_eraPostLegacyBridgeUpgradeFirstBatch\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"_eraLegacyBridgeLastDepositBatch\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"_eraLegacyBridgeLastDepositTxNumber\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"isWithdrawalFinalized\",\"inputs\":[{\"name\":\"chainId\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"l2BatchNumber\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"l2ToL1MessageNumber\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"isFinalized\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"l1AssetRouter\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"contractIL1AssetRouter\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"l1NativeTokenVault\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"contractIL1NativeTokenVault\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"l2BridgeAddress\",\"inputs\":[{\"name\":\"_chainId\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"legacyBridge\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"contractIL1ERC20Bridge\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"nullifyChainBalanceByNTV\",\"inputs\":[{\"name\":\"_chainId\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"_token\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"owner\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"pause\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"paused\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"pendingOwner\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"renounceOwnership\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setL1AssetRouter\",\"inputs\":[{\"name\":\"_l1AssetRouter\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setL1Erc20Bridge\",\"inputs\":[{\"name\":\"_legacyBridge\",\"type\":\"address\",\"internalType\":\"contractIL1ERC20Bridge\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setL1NativeTokenVault\",\"inputs\":[{\"name\":\"_l1NativeTokenVault\",\"type\":\"address\",\"internalType\":\"contractIL1NativeTokenVault\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"transferOwnership\",\"inputs\":[{\"name\":\"newOwner\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"transferTokenToNTV\",\"inputs\":[{\"name\":\"_token\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"unpause\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"event\",\"name\":\"BridgehubDepositFinalized\",\"inputs\":[{\"name\":\"chainId\",\"type\":\"uint256\",\"indexed\":true,\"internalType\":\"uint256\"},{\"name\":\"txDataHash\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"},{\"name\":\"l2DepositTxHash\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"Initialized\",\"inputs\":[{\"name\":\"version\",\"type\":\"uint8\",\"indexed\":false,\"internalType\":\"uint8\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"OwnershipTransferStarted\",\"inputs\":[{\"name\":\"previousOwner\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"newOwner\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"OwnershipTransferred\",\"inputs\":[{\"name\":\"previousOwner\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"newOwner\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"Paused\",\"inputs\":[{\"name\":\"account\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"Unpaused\",\"inputs\":[{\"name\":\"account\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"error\",\"name\":\"AddressAlreadySet\",\"inputs\":[{\"name\":\"addr\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"DepositDoesNotExist\",\"inputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}]},{\"type\":\"error\",\"name\":\"DepositExists\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"EthTransferFailed\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"IncorrectTokenAddressFromNTV\",\"inputs\":[{\"name\":\"assetId\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"tokenAddress\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"InvalidNTVBurnData\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidProof\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidSelector\",\"inputs\":[{\"name\":\"func\",\"type\":\"bytes4\",\"internalType\":\"bytes4\"}]},{\"type\":\"error\",\"name\":\"L2WithdrawalMessageWrongLength\",\"inputs\":[{\"name\":\"messageLen\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"type\":\"error\",\"name\":\"LegacyBridgeNotSet\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"LegacyMethodForNonL1Token\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"NativeTokenVaultAlreadySet\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"NotInitializedReentrancyGuard\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"Reentrancy\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"SharedBridgeValueNotSet\",\"inputs\":[{\"name\":\"\",\"type\":\"uint8\",\"internalType\":\"enumSharedBridgeKey\"}]},{\"type\":\"error\",\"name\":\"SlotOccupied\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"TokenNotLegacy\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"Unauthorized\",\"inputs\":[{\"name\":\"caller\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"UnsupportedEncodingVersion\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"WithdrawalAlreadyFinalized\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"WrongL2Sender\",\"inputs\":[{\"name\":\"providedL2Sender\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"WrongMsgLength\",\"inputs\":[{\"name\":\"expected\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"length\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"type\":\"error\",\"name\":\"ZeroAddress\",\"inputs\":[]}]",
}

// IL1NullifierABI is the input ABI used to generate the binding from.
// Deprecated: Use IL1NullifierMetaData.ABI instead.
var IL1NullifierABI = IL1NullifierMetaData.ABI

// IL1Nullifier is an auto generated Go binding around an Ethereum contract.
type IL1Nullifier struct {
	IL1NullifierCaller     // Read-only binding to the contract
	IL1NullifierTransactor // Write-only binding to the contract
	IL1NullifierFilterer   // Log filterer for contract events
}

// IL1NullifierCaller is an auto generated read-only Go binding around an Ethereum contract.
type IL1NullifierCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IL1NullifierTransactor is an auto generated write-only Go binding around an Ethereum contract.
type IL1NullifierTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IL1NullifierFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type IL1NullifierFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IL1NullifierSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type IL1NullifierSession struct {
	Contract     *IL1Nullifier     // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// IL1NullifierCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type IL1NullifierCallerSession struct {
	Contract *IL1NullifierCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts       // Call options to use throughout this session
}

// IL1NullifierTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type IL1NullifierTransactorSession struct {
	Contract     *IL1NullifierTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts       // Transaction auth options to use throughout this session
}

// IL1NullifierRaw is an auto generated low-level Go binding around an Ethereum contract.
type IL1NullifierRaw struct {
	Contract *IL1Nullifier // Generic contract binding to access the raw methods on
}

// IL1NullifierCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type IL1NullifierCallerRaw struct {
	Contract *IL1NullifierCaller // Generic read-only contract binding to access the raw methods on
}

// IL1NullifierTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type IL1NullifierTransactorRaw struct {
	Contract *IL1NullifierTransactor // Generic write-only contract binding to access the raw methods on
}

// NewIL1Nullifier creates a new instance of IL1Nullifier, bound to a specific deployed contract.
func NewIL1Nullifier(address common.Address, backend bind.ContractBackend) (*IL1Nullifier, error) {
	contract, err := bindIL1Nullifier(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &IL1Nullifier{IL1NullifierCaller: IL1NullifierCaller{contract: contract}, IL1NullifierTransactor: IL1NullifierTransactor{contract: contract}, IL1NullifierFilterer: IL1NullifierFilterer{contract: contract}}, nil
}

// NewIL1NullifierCaller creates a new read-only instance of IL1Nullifier, bound to a specific deployed contract.
func NewIL1NullifierCaller(address common.Address, caller bind.ContractCaller) (*IL1NullifierCaller, error) {
	contract, err := bindIL1Nullifier(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &IL1NullifierCaller{contract: contract}, nil
}

// NewIL1NullifierTransactor creates a new write-only instance of IL1Nullifier, bound to a specific deployed contract.
func NewIL1NullifierTransactor(address common.Address, transactor bind.ContractTransactor) (*IL1NullifierTransactor, error) {
	contract, err := bindIL1Nullifier(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &IL1NullifierTransactor{contract: contract}, nil
}

// NewIL1NullifierFilterer creates a new log filterer instance of IL1Nullifier, bound to a specific deployed contract.
func NewIL1NullifierFilterer(address common.Address, filterer bind.ContractFilterer) (*IL1NullifierFilterer, error) {
	contract, err := bindIL1Nullifier(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &IL1NullifierFilterer{contract: contract}, nil
}

// bindIL1Nullifier binds a generic wrapper to an already deployed contract.
func bindIL1Nullifier(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := IL1NullifierMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IL1Nullifier *IL1NullifierRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IL1Nullifier.Contract.IL1NullifierCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IL1Nullifier *IL1NullifierRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IL1Nullifier.Contract.IL1NullifierTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IL1Nullifier *IL1NullifierRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IL1Nullifier.Contract.IL1NullifierTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IL1Nullifier *IL1NullifierCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IL1Nullifier.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IL1Nullifier *IL1NullifierTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IL1Nullifier.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IL1Nullifier *IL1NullifierTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IL1Nullifier.Contract.contract.Transact(opts, method, params...)
}

// BRIDGEHUB is a free data retrieval call binding the contract method 0x5d4edca7.
//
// Solidity: function BRIDGE_HUB() view returns(address)
func (_IL1Nullifier *IL1NullifierCaller) BRIDGEHUB(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _IL1Nullifier.contract.Call(opts, &out, "BRIDGE_HUB")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// BRIDGEHUB is a free data retrieval call binding the contract method 0x5d4edca7.
//
// Solidity: function BRIDGE_HUB() view returns(address)
func (_IL1Nullifier *IL1NullifierSession) BRIDGEHUB() (common.Address, error) {
	return _IL1Nullifier.Contract.BRIDGEHUB(&_IL1Nullifier.CallOpts)
}

// BRIDGEHUB is a free data retrieval call binding the contract method 0x5d4edca7.
//
// Solidity: function BRIDGE_HUB() view returns(address)
func (_IL1Nullifier *IL1NullifierCallerSession) BRIDGEHUB() (common.Address, error) {
	return _IL1Nullifier.Contract.BRIDGEHUB(&_IL1Nullifier.CallOpts)
}

// DEPRECATEDAdmin is a free data retrieval call binding the contract method 0xf7a5cec0.
//
// Solidity: function __DEPRECATED_admin() view returns(address)
func (_IL1Nullifier *IL1NullifierCaller) DEPRECATEDAdmin(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _IL1Nullifier.contract.Call(opts, &out, "__DEPRECATED_admin")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// DEPRECATEDAdmin is a free data retrieval call binding the contract method 0xf7a5cec0.
//
// Solidity: function __DEPRECATED_admin() view returns(address)
func (_IL1Nullifier *IL1NullifierSession) DEPRECATEDAdmin() (common.Address, error) {
	return _IL1Nullifier.Contract.DEPRECATEDAdmin(&_IL1Nullifier.CallOpts)
}

// DEPRECATEDAdmin is a free data retrieval call binding the contract method 0xf7a5cec0.
//
// Solidity: function __DEPRECATED_admin() view returns(address)
func (_IL1Nullifier *IL1NullifierCallerSession) DEPRECATEDAdmin() (common.Address, error) {
	return _IL1Nullifier.Contract.DEPRECATEDAdmin(&_IL1Nullifier.CallOpts)
}

// DEPRECATEDChainBalance is a free data retrieval call binding the contract method 0x6182877b.
//
// Solidity: function __DEPRECATED_chainBalance(uint256 chainId, address l1Token) view returns(uint256 balance)
func (_IL1Nullifier *IL1NullifierCaller) DEPRECATEDChainBalance(opts *bind.CallOpts, chainId *big.Int, l1Token common.Address) (*big.Int, error) {
	var out []interface{}
	err := _IL1Nullifier.contract.Call(opts, &out, "__DEPRECATED_chainBalance", chainId, l1Token)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// DEPRECATEDChainBalance is a free data retrieval call binding the contract method 0x6182877b.
//
// Solidity: function __DEPRECATED_chainBalance(uint256 chainId, address l1Token) view returns(uint256 balance)
func (_IL1Nullifier *IL1NullifierSession) DEPRECATEDChainBalance(chainId *big.Int, l1Token common.Address) (*big.Int, error) {
	return _IL1Nullifier.Contract.DEPRECATEDChainBalance(&_IL1Nullifier.CallOpts, chainId, l1Token)
}

// DEPRECATEDChainBalance is a free data retrieval call binding the contract method 0x6182877b.
//
// Solidity: function __DEPRECATED_chainBalance(uint256 chainId, address l1Token) view returns(uint256 balance)
func (_IL1Nullifier *IL1NullifierCallerSession) DEPRECATEDChainBalance(chainId *big.Int, l1Token common.Address) (*big.Int, error) {
	return _IL1Nullifier.Contract.DEPRECATEDChainBalance(&_IL1Nullifier.CallOpts, chainId, l1Token)
}

// DEPRECATEDL2BridgeAddress is a free data retrieval call binding the contract method 0xfdbb0301.
//
// Solidity: function __DEPRECATED_l2BridgeAddress(uint256 chainId) view returns(address l2Bridge)
func (_IL1Nullifier *IL1NullifierCaller) DEPRECATEDL2BridgeAddress(opts *bind.CallOpts, chainId *big.Int) (common.Address, error) {
	var out []interface{}
	err := _IL1Nullifier.contract.Call(opts, &out, "__DEPRECATED_l2BridgeAddress", chainId)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// DEPRECATEDL2BridgeAddress is a free data retrieval call binding the contract method 0xfdbb0301.
//
// Solidity: function __DEPRECATED_l2BridgeAddress(uint256 chainId) view returns(address l2Bridge)
func (_IL1Nullifier *IL1NullifierSession) DEPRECATEDL2BridgeAddress(chainId *big.Int) (common.Address, error) {
	return _IL1Nullifier.Contract.DEPRECATEDL2BridgeAddress(&_IL1Nullifier.CallOpts, chainId)
}

// DEPRECATEDL2BridgeAddress is a free data retrieval call binding the contract method 0xfdbb0301.
//
// Solidity: function __DEPRECATED_l2BridgeAddress(uint256 chainId) view returns(address l2Bridge)
func (_IL1Nullifier *IL1NullifierCallerSession) DEPRECATEDL2BridgeAddress(chainId *big.Int) (common.Address, error) {
	return _IL1Nullifier.Contract.DEPRECATEDL2BridgeAddress(&_IL1Nullifier.CallOpts, chainId)
}

// DEPRECATEDPendingAdmin is a free data retrieval call binding the contract method 0x6cdecb2b.
//
// Solidity: function __DEPRECATED_pendingAdmin() view returns(address)
func (_IL1Nullifier *IL1NullifierCaller) DEPRECATEDPendingAdmin(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _IL1Nullifier.contract.Call(opts, &out, "__DEPRECATED_pendingAdmin")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// DEPRECATEDPendingAdmin is a free data retrieval call binding the contract method 0x6cdecb2b.
//
// Solidity: function __DEPRECATED_pendingAdmin() view returns(address)
func (_IL1Nullifier *IL1NullifierSession) DEPRECATEDPendingAdmin() (common.Address, error) {
	return _IL1Nullifier.Contract.DEPRECATEDPendingAdmin(&_IL1Nullifier.CallOpts)
}

// DEPRECATEDPendingAdmin is a free data retrieval call binding the contract method 0x6cdecb2b.
//
// Solidity: function __DEPRECATED_pendingAdmin() view returns(address)
func (_IL1Nullifier *IL1NullifierCallerSession) DEPRECATEDPendingAdmin() (common.Address, error) {
	return _IL1Nullifier.Contract.DEPRECATEDPendingAdmin(&_IL1Nullifier.CallOpts)
}

// ChainBalance is a free data retrieval call binding the contract method 0x9cd45184.
//
// Solidity: function chainBalance(uint256 _chainId, address _token) view returns(uint256)
func (_IL1Nullifier *IL1NullifierCaller) ChainBalance(opts *bind.CallOpts, _chainId *big.Int, _token common.Address) (*big.Int, error) {
	var out []interface{}
	err := _IL1Nullifier.contract.Call(opts, &out, "chainBalance", _chainId, _token)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// ChainBalance is a free data retrieval call binding the contract method 0x9cd45184.
//
// Solidity: function chainBalance(uint256 _chainId, address _token) view returns(uint256)
func (_IL1Nullifier *IL1NullifierSession) ChainBalance(_chainId *big.Int, _token common.Address) (*big.Int, error) {
	return _IL1Nullifier.Contract.ChainBalance(&_IL1Nullifier.CallOpts, _chainId, _token)
}

// ChainBalance is a free data retrieval call binding the contract method 0x9cd45184.
//
// Solidity: function chainBalance(uint256 _chainId, address _token) view returns(uint256)
func (_IL1Nullifier *IL1NullifierCallerSession) ChainBalance(_chainId *big.Int, _token common.Address) (*big.Int, error) {
	return _IL1Nullifier.Contract.ChainBalance(&_IL1Nullifier.CallOpts, _chainId, _token)
}

// DepositHappened is a free data retrieval call binding the contract method 0x9fa8826b.
//
// Solidity: function depositHappened(uint256 chainId, bytes32 l2DepositTxHash) view returns(bytes32 depositDataHash)
func (_IL1Nullifier *IL1NullifierCaller) DepositHappened(opts *bind.CallOpts, chainId *big.Int, l2DepositTxHash [32]byte) ([32]byte, error) {
	var out []interface{}
	err := _IL1Nullifier.contract.Call(opts, &out, "depositHappened", chainId, l2DepositTxHash)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// DepositHappened is a free data retrieval call binding the contract method 0x9fa8826b.
//
// Solidity: function depositHappened(uint256 chainId, bytes32 l2DepositTxHash) view returns(bytes32 depositDataHash)
func (_IL1Nullifier *IL1NullifierSession) DepositHappened(chainId *big.Int, l2DepositTxHash [32]byte) ([32]byte, error) {
	return _IL1Nullifier.Contract.DepositHappened(&_IL1Nullifier.CallOpts, chainId, l2DepositTxHash)
}

// DepositHappened is a free data retrieval call binding the contract method 0x9fa8826b.
//
// Solidity: function depositHappened(uint256 chainId, bytes32 l2DepositTxHash) view returns(bytes32 depositDataHash)
func (_IL1Nullifier *IL1NullifierCallerSession) DepositHappened(chainId *big.Int, l2DepositTxHash [32]byte) ([32]byte, error) {
	return _IL1Nullifier.Contract.DepositHappened(&_IL1Nullifier.CallOpts, chainId, l2DepositTxHash)
}

// EncodeTxDataHash is a free data retrieval call binding the contract method 0xf120e6c4.
//
// Solidity: function encodeTxDataHash(bytes1 _encodingVersion, address _originalCaller, bytes32 _assetId, bytes _transferData) view returns(bytes32 txDataHash)
func (_IL1Nullifier *IL1NullifierCaller) EncodeTxDataHash(opts *bind.CallOpts, _encodingVersion [1]byte, _originalCaller common.Address, _assetId [32]byte, _transferData []byte) ([32]byte, error) {
	var out []interface{}
	err := _IL1Nullifier.contract.Call(opts, &out, "encodeTxDataHash", _encodingVersion, _originalCaller, _assetId, _transferData)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// EncodeTxDataHash is a free data retrieval call binding the contract method 0xf120e6c4.
//
// Solidity: function encodeTxDataHash(bytes1 _encodingVersion, address _originalCaller, bytes32 _assetId, bytes _transferData) view returns(bytes32 txDataHash)
func (_IL1Nullifier *IL1NullifierSession) EncodeTxDataHash(_encodingVersion [1]byte, _originalCaller common.Address, _assetId [32]byte, _transferData []byte) ([32]byte, error) {
	return _IL1Nullifier.Contract.EncodeTxDataHash(&_IL1Nullifier.CallOpts, _encodingVersion, _originalCaller, _assetId, _transferData)
}

// EncodeTxDataHash is a free data retrieval call binding the contract method 0xf120e6c4.
//
// Solidity: function encodeTxDataHash(bytes1 _encodingVersion, address _originalCaller, bytes32 _assetId, bytes _transferData) view returns(bytes32 txDataHash)
func (_IL1Nullifier *IL1NullifierCallerSession) EncodeTxDataHash(_encodingVersion [1]byte, _originalCaller common.Address, _assetId [32]byte, _transferData []byte) ([32]byte, error) {
	return _IL1Nullifier.Contract.EncodeTxDataHash(&_IL1Nullifier.CallOpts, _encodingVersion, _originalCaller, _assetId, _transferData)
}

// IsWithdrawalFinalized is a free data retrieval call binding the contract method 0x8f31f052.
//
// Solidity: function isWithdrawalFinalized(uint256 chainId, uint256 l2BatchNumber, uint256 l2ToL1MessageNumber) view returns(bool isFinalized)
func (_IL1Nullifier *IL1NullifierCaller) IsWithdrawalFinalized(opts *bind.CallOpts, chainId *big.Int, l2BatchNumber *big.Int, l2ToL1MessageNumber *big.Int) (bool, error) {
	var out []interface{}
	err := _IL1Nullifier.contract.Call(opts, &out, "isWithdrawalFinalized", chainId, l2BatchNumber, l2ToL1MessageNumber)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsWithdrawalFinalized is a free data retrieval call binding the contract method 0x8f31f052.
//
// Solidity: function isWithdrawalFinalized(uint256 chainId, uint256 l2BatchNumber, uint256 l2ToL1MessageNumber) view returns(bool isFinalized)
func (_IL1Nullifier *IL1NullifierSession) IsWithdrawalFinalized(chainId *big.Int, l2BatchNumber *big.Int, l2ToL1MessageNumber *big.Int) (bool, error) {
	return _IL1Nullifier.Contract.IsWithdrawalFinalized(&_IL1Nullifier.CallOpts, chainId, l2BatchNumber, l2ToL1MessageNumber)
}

// IsWithdrawalFinalized is a free data retrieval call binding the contract method 0x8f31f052.
//
// Solidity: function isWithdrawalFinalized(uint256 chainId, uint256 l2BatchNumber, uint256 l2ToL1MessageNumber) view returns(bool isFinalized)
func (_IL1Nullifier *IL1NullifierCallerSession) IsWithdrawalFinalized(chainId *big.Int, l2BatchNumber *big.Int, l2ToL1MessageNumber *big.Int) (bool, error) {
	return _IL1Nullifier.Contract.IsWithdrawalFinalized(&_IL1Nullifier.CallOpts, chainId, l2BatchNumber, l2ToL1MessageNumber)
}

// L1AssetRouter is a free data retrieval call binding the contract method 0x6d9860e1.
//
// Solidity: function l1AssetRouter() view returns(address)
func (_IL1Nullifier *IL1NullifierCaller) L1AssetRouter(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _IL1Nullifier.contract.Call(opts, &out, "l1AssetRouter")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// L1AssetRouter is a free data retrieval call binding the contract method 0x6d9860e1.
//
// Solidity: function l1AssetRouter() view returns(address)
func (_IL1Nullifier *IL1NullifierSession) L1AssetRouter() (common.Address, error) {
	return _IL1Nullifier.Contract.L1AssetRouter(&_IL1Nullifier.CallOpts)
}

// L1AssetRouter is a free data retrieval call binding the contract method 0x6d9860e1.
//
// Solidity: function l1AssetRouter() view returns(address)
func (_IL1Nullifier *IL1NullifierCallerSession) L1AssetRouter() (common.Address, error) {
	return _IL1Nullifier.Contract.L1AssetRouter(&_IL1Nullifier.CallOpts)
}

// L1NativeTokenVault is a free data retrieval call binding the contract method 0x6f513211.
//
// Solidity: function l1NativeTokenVault() view returns(address)
func (_IL1Nullifier *IL1NullifierCaller) L1NativeTokenVault(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _IL1Nullifier.contract.Call(opts, &out, "l1NativeTokenVault")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// L1NativeTokenVault is a free data retrieval call binding the contract method 0x6f513211.
//
// Solidity: function l1NativeTokenVault() view returns(address)
func (_IL1Nullifier *IL1NullifierSession) L1NativeTokenVault() (common.Address, error) {
	return _IL1Nullifier.Contract.L1NativeTokenVault(&_IL1Nullifier.CallOpts)
}

// L1NativeTokenVault is a free data retrieval call binding the contract method 0x6f513211.
//
// Solidity: function l1NativeTokenVault() view returns(address)
func (_IL1Nullifier *IL1NullifierCallerSession) L1NativeTokenVault() (common.Address, error) {
	return _IL1Nullifier.Contract.L1NativeTokenVault(&_IL1Nullifier.CallOpts)
}

// L2BridgeAddress is a free data retrieval call binding the contract method 0x07ee9355.
//
// Solidity: function l2BridgeAddress(uint256 _chainId) view returns(address)
func (_IL1Nullifier *IL1NullifierCaller) L2BridgeAddress(opts *bind.CallOpts, _chainId *big.Int) (common.Address, error) {
	var out []interface{}
	err := _IL1Nullifier.contract.Call(opts, &out, "l2BridgeAddress", _chainId)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// L2BridgeAddress is a free data retrieval call binding the contract method 0x07ee9355.
//
// Solidity: function l2BridgeAddress(uint256 _chainId) view returns(address)
func (_IL1Nullifier *IL1NullifierSession) L2BridgeAddress(_chainId *big.Int) (common.Address, error) {
	return _IL1Nullifier.Contract.L2BridgeAddress(&_IL1Nullifier.CallOpts, _chainId)
}

// L2BridgeAddress is a free data retrieval call binding the contract method 0x07ee9355.
//
// Solidity: function l2BridgeAddress(uint256 _chainId) view returns(address)
func (_IL1Nullifier *IL1NullifierCallerSession) L2BridgeAddress(_chainId *big.Int) (common.Address, error) {
	return _IL1Nullifier.Contract.L2BridgeAddress(&_IL1Nullifier.CallOpts, _chainId)
}

// LegacyBridge is a free data retrieval call binding the contract method 0x6e9d7899.
//
// Solidity: function legacyBridge() view returns(address)
func (_IL1Nullifier *IL1NullifierCaller) LegacyBridge(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _IL1Nullifier.contract.Call(opts, &out, "legacyBridge")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// LegacyBridge is a free data retrieval call binding the contract method 0x6e9d7899.
//
// Solidity: function legacyBridge() view returns(address)
func (_IL1Nullifier *IL1NullifierSession) LegacyBridge() (common.Address, error) {
	return _IL1Nullifier.Contract.LegacyBridge(&_IL1Nullifier.CallOpts)
}

// LegacyBridge is a free data retrieval call binding the contract method 0x6e9d7899.
//
// Solidity: function legacyBridge() view returns(address)
func (_IL1Nullifier *IL1NullifierCallerSession) LegacyBridge() (common.Address, error) {
	return _IL1Nullifier.Contract.LegacyBridge(&_IL1Nullifier.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_IL1Nullifier *IL1NullifierCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _IL1Nullifier.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_IL1Nullifier *IL1NullifierSession) Owner() (common.Address, error) {
	return _IL1Nullifier.Contract.Owner(&_IL1Nullifier.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_IL1Nullifier *IL1NullifierCallerSession) Owner() (common.Address, error) {
	return _IL1Nullifier.Contract.Owner(&_IL1Nullifier.CallOpts)
}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_IL1Nullifier *IL1NullifierCaller) Paused(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _IL1Nullifier.contract.Call(opts, &out, "paused")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_IL1Nullifier *IL1NullifierSession) Paused() (bool, error) {
	return _IL1Nullifier.Contract.Paused(&_IL1Nullifier.CallOpts)
}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_IL1Nullifier *IL1NullifierCallerSession) Paused() (bool, error) {
	return _IL1Nullifier.Contract.Paused(&_IL1Nullifier.CallOpts)
}

// PendingOwner is a free data retrieval call binding the contract method 0xe30c3978.
//
// Solidity: function pendingOwner() view returns(address)
func (_IL1Nullifier *IL1NullifierCaller) PendingOwner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _IL1Nullifier.contract.Call(opts, &out, "pendingOwner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// PendingOwner is a free data retrieval call binding the contract method 0xe30c3978.
//
// Solidity: function pendingOwner() view returns(address)
func (_IL1Nullifier *IL1NullifierSession) PendingOwner() (common.Address, error) {
	return _IL1Nullifier.Contract.PendingOwner(&_IL1Nullifier.CallOpts)
}

// PendingOwner is a free data retrieval call binding the contract method 0xe30c3978.
//
// Solidity: function pendingOwner() view returns(address)
func (_IL1Nullifier *IL1NullifierCallerSession) PendingOwner() (common.Address, error) {
	return _IL1Nullifier.Contract.PendingOwner(&_IL1Nullifier.CallOpts)
}

// AcceptOwnership is a paid mutator transaction binding the contract method 0x79ba5097.
//
// Solidity: function acceptOwnership() returns()
func (_IL1Nullifier *IL1NullifierTransactor) AcceptOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IL1Nullifier.contract.Transact(opts, "acceptOwnership")
}

// AcceptOwnership is a paid mutator transaction binding the contract method 0x79ba5097.
//
// Solidity: function acceptOwnership() returns()
func (_IL1Nullifier *IL1NullifierSession) AcceptOwnership() (*types.Transaction, error) {
	return _IL1Nullifier.Contract.AcceptOwnership(&_IL1Nullifier.TransactOpts)
}

// AcceptOwnership is a paid mutator transaction binding the contract method 0x79ba5097.
//
// Solidity: function acceptOwnership() returns()
func (_IL1Nullifier *IL1NullifierTransactorSession) AcceptOwnership() (*types.Transaction, error) {
	return _IL1Nullifier.Contract.AcceptOwnership(&_IL1Nullifier.TransactOpts)
}

// BridgeRecoverFailedTransfer is a paid mutator transaction binding the contract method 0x3601e63e.
//
// Solidity: function bridgeRecoverFailedTransfer(uint256 _chainId, address _depositSender, bytes32 _assetId, bytes _assetData, bytes32 _l2TxHash, uint256 _l2BatchNumber, uint256 _l2MessageIndex, uint16 _l2TxNumberInBatch, bytes32[] _merkleProof) returns()
func (_IL1Nullifier *IL1NullifierTransactor) BridgeRecoverFailedTransfer(opts *bind.TransactOpts, _chainId *big.Int, _depositSender common.Address, _assetId [32]byte, _assetData []byte, _l2TxHash [32]byte, _l2BatchNumber *big.Int, _l2MessageIndex *big.Int, _l2TxNumberInBatch uint16, _merkleProof [][32]byte) (*types.Transaction, error) {
	return _IL1Nullifier.contract.Transact(opts, "bridgeRecoverFailedTransfer", _chainId, _depositSender, _assetId, _assetData, _l2TxHash, _l2BatchNumber, _l2MessageIndex, _l2TxNumberInBatch, _merkleProof)
}

// BridgeRecoverFailedTransfer is a paid mutator transaction binding the contract method 0x3601e63e.
//
// Solidity: function bridgeRecoverFailedTransfer(uint256 _chainId, address _depositSender, bytes32 _assetId, bytes _assetData, bytes32 _l2TxHash, uint256 _l2BatchNumber, uint256 _l2MessageIndex, uint16 _l2TxNumberInBatch, bytes32[] _merkleProof) returns()
func (_IL1Nullifier *IL1NullifierSession) BridgeRecoverFailedTransfer(_chainId *big.Int, _depositSender common.Address, _assetId [32]byte, _assetData []byte, _l2TxHash [32]byte, _l2BatchNumber *big.Int, _l2MessageIndex *big.Int, _l2TxNumberInBatch uint16, _merkleProof [][32]byte) (*types.Transaction, error) {
	return _IL1Nullifier.Contract.BridgeRecoverFailedTransfer(&_IL1Nullifier.TransactOpts, _chainId, _depositSender, _assetId, _assetData, _l2TxHash, _l2BatchNumber, _l2MessageIndex, _l2TxNumberInBatch, _merkleProof)
}

// BridgeRecoverFailedTransfer is a paid mutator transaction binding the contract method 0x3601e63e.
//
// Solidity: function bridgeRecoverFailedTransfer(uint256 _chainId, address _depositSender, bytes32 _assetId, bytes _assetData, bytes32 _l2TxHash, uint256 _l2BatchNumber, uint256 _l2MessageIndex, uint16 _l2TxNumberInBatch, bytes32[] _merkleProof) returns()
func (_IL1Nullifier *IL1NullifierTransactorSession) BridgeRecoverFailedTransfer(_chainId *big.Int, _depositSender common.Address, _assetId [32]byte, _assetData []byte, _l2TxHash [32]byte, _l2BatchNumber *big.Int, _l2MessageIndex *big.Int, _l2TxNumberInBatch uint16, _merkleProof [][32]byte) (*types.Transaction, error) {
	return _IL1Nullifier.Contract.BridgeRecoverFailedTransfer(&_IL1Nullifier.TransactOpts, _chainId, _depositSender, _assetId, _assetData, _l2TxHash, _l2BatchNumber, _l2MessageIndex, _l2TxNumberInBatch, _merkleProof)
}

// BridgehubConfirmL2TransactionForwarded is a paid mutator transaction binding the contract method 0x4bc2c8c0.
//
// Solidity: function bridgehubConfirmL2TransactionForwarded(uint256 _chainId, bytes32 _txDataHash, bytes32 _txHash) returns()
func (_IL1Nullifier *IL1NullifierTransactor) BridgehubConfirmL2TransactionForwarded(opts *bind.TransactOpts, _chainId *big.Int, _txDataHash [32]byte, _txHash [32]byte) (*types.Transaction, error) {
	return _IL1Nullifier.contract.Transact(opts, "bridgehubConfirmL2TransactionForwarded", _chainId, _txDataHash, _txHash)
}

// BridgehubConfirmL2TransactionForwarded is a paid mutator transaction binding the contract method 0x4bc2c8c0.
//
// Solidity: function bridgehubConfirmL2TransactionForwarded(uint256 _chainId, bytes32 _txDataHash, bytes32 _txHash) returns()
func (_IL1Nullifier *IL1NullifierSession) BridgehubConfirmL2TransactionForwarded(_chainId *big.Int, _txDataHash [32]byte, _txHash [32]byte) (*types.Transaction, error) {
	return _IL1Nullifier.Contract.BridgehubConfirmL2TransactionForwarded(&_IL1Nullifier.TransactOpts, _chainId, _txDataHash, _txHash)
}

// BridgehubConfirmL2TransactionForwarded is a paid mutator transaction binding the contract method 0x4bc2c8c0.
//
// Solidity: function bridgehubConfirmL2TransactionForwarded(uint256 _chainId, bytes32 _txDataHash, bytes32 _txHash) returns()
func (_IL1Nullifier *IL1NullifierTransactorSession) BridgehubConfirmL2TransactionForwarded(_chainId *big.Int, _txDataHash [32]byte, _txHash [32]byte) (*types.Transaction, error) {
	return _IL1Nullifier.Contract.BridgehubConfirmL2TransactionForwarded(&_IL1Nullifier.TransactOpts, _chainId, _txDataHash, _txHash)
}

// ClaimFailedDeposit is a paid mutator transaction binding the contract method 0xc0991525.
//
// Solidity: function claimFailedDeposit(uint256 _chainId, address _depositSender, address _l1Token, uint256 _amount, bytes32 _l2TxHash, uint256 _l2BatchNumber, uint256 _l2MessageIndex, uint16 _l2TxNumberInBatch, bytes32[] _merkleProof) returns()
func (_IL1Nullifier *IL1NullifierTransactor) ClaimFailedDeposit(opts *bind.TransactOpts, _chainId *big.Int, _depositSender common.Address, _l1Token common.Address, _amount *big.Int, _l2TxHash [32]byte, _l2BatchNumber *big.Int, _l2MessageIndex *big.Int, _l2TxNumberInBatch uint16, _merkleProof [][32]byte) (*types.Transaction, error) {
	return _IL1Nullifier.contract.Transact(opts, "claimFailedDeposit", _chainId, _depositSender, _l1Token, _amount, _l2TxHash, _l2BatchNumber, _l2MessageIndex, _l2TxNumberInBatch, _merkleProof)
}

// ClaimFailedDeposit is a paid mutator transaction binding the contract method 0xc0991525.
//
// Solidity: function claimFailedDeposit(uint256 _chainId, address _depositSender, address _l1Token, uint256 _amount, bytes32 _l2TxHash, uint256 _l2BatchNumber, uint256 _l2MessageIndex, uint16 _l2TxNumberInBatch, bytes32[] _merkleProof) returns()
func (_IL1Nullifier *IL1NullifierSession) ClaimFailedDeposit(_chainId *big.Int, _depositSender common.Address, _l1Token common.Address, _amount *big.Int, _l2TxHash [32]byte, _l2BatchNumber *big.Int, _l2MessageIndex *big.Int, _l2TxNumberInBatch uint16, _merkleProof [][32]byte) (*types.Transaction, error) {
	return _IL1Nullifier.Contract.ClaimFailedDeposit(&_IL1Nullifier.TransactOpts, _chainId, _depositSender, _l1Token, _amount, _l2TxHash, _l2BatchNumber, _l2MessageIndex, _l2TxNumberInBatch, _merkleProof)
}

// ClaimFailedDeposit is a paid mutator transaction binding the contract method 0xc0991525.
//
// Solidity: function claimFailedDeposit(uint256 _chainId, address _depositSender, address _l1Token, uint256 _amount, bytes32 _l2TxHash, uint256 _l2BatchNumber, uint256 _l2MessageIndex, uint16 _l2TxNumberInBatch, bytes32[] _merkleProof) returns()
func (_IL1Nullifier *IL1NullifierTransactorSession) ClaimFailedDeposit(_chainId *big.Int, _depositSender common.Address, _l1Token common.Address, _amount *big.Int, _l2TxHash [32]byte, _l2BatchNumber *big.Int, _l2MessageIndex *big.Int, _l2TxNumberInBatch uint16, _merkleProof [][32]byte) (*types.Transaction, error) {
	return _IL1Nullifier.Contract.ClaimFailedDeposit(&_IL1Nullifier.TransactOpts, _chainId, _depositSender, _l1Token, _amount, _l2TxHash, _l2BatchNumber, _l2MessageIndex, _l2TxNumberInBatch, _merkleProof)
}

// ClaimFailedDepositLegacyErc20Bridge is a paid mutator transaction binding the contract method 0x8fbb3711.
//
// Solidity: function claimFailedDepositLegacyErc20Bridge(address _depositSender, address _l1Token, uint256 _amount, bytes32 _l2TxHash, uint256 _l2BatchNumber, uint256 _l2MessageIndex, uint16 _l2TxNumberInBatch, bytes32[] _merkleProof) returns()
func (_IL1Nullifier *IL1NullifierTransactor) ClaimFailedDepositLegacyErc20Bridge(opts *bind.TransactOpts, _depositSender common.Address, _l1Token common.Address, _amount *big.Int, _l2TxHash [32]byte, _l2BatchNumber *big.Int, _l2MessageIndex *big.Int, _l2TxNumberInBatch uint16, _merkleProof [][32]byte) (*types.Transaction, error) {
	return _IL1Nullifier.contract.Transact(opts, "claimFailedDepositLegacyErc20Bridge", _depositSender, _l1Token, _amount, _l2TxHash, _l2BatchNumber, _l2MessageIndex, _l2TxNumberInBatch, _merkleProof)
}

// ClaimFailedDepositLegacyErc20Bridge is a paid mutator transaction binding the contract method 0x8fbb3711.
//
// Solidity: function claimFailedDepositLegacyErc20Bridge(address _depositSender, address _l1Token, uint256 _amount, bytes32 _l2TxHash, uint256 _l2BatchNumber, uint256 _l2MessageIndex, uint16 _l2TxNumberInBatch, bytes32[] _merkleProof) returns()
func (_IL1Nullifier *IL1NullifierSession) ClaimFailedDepositLegacyErc20Bridge(_depositSender common.Address, _l1Token common.Address, _amount *big.Int, _l2TxHash [32]byte, _l2BatchNumber *big.Int, _l2MessageIndex *big.Int, _l2TxNumberInBatch uint16, _merkleProof [][32]byte) (*types.Transaction, error) {
	return _IL1Nullifier.Contract.ClaimFailedDepositLegacyErc20Bridge(&_IL1Nullifier.TransactOpts, _depositSender, _l1Token, _amount, _l2TxHash, _l2BatchNumber, _l2MessageIndex, _l2TxNumberInBatch, _merkleProof)
}

// ClaimFailedDepositLegacyErc20Bridge is a paid mutator transaction binding the contract method 0x8fbb3711.
//
// Solidity: function claimFailedDepositLegacyErc20Bridge(address _depositSender, address _l1Token, uint256 _amount, bytes32 _l2TxHash, uint256 _l2BatchNumber, uint256 _l2MessageIndex, uint16 _l2TxNumberInBatch, bytes32[] _merkleProof) returns()
func (_IL1Nullifier *IL1NullifierTransactorSession) ClaimFailedDepositLegacyErc20Bridge(_depositSender common.Address, _l1Token common.Address, _amount *big.Int, _l2TxHash [32]byte, _l2BatchNumber *big.Int, _l2MessageIndex *big.Int, _l2TxNumberInBatch uint16, _merkleProof [][32]byte) (*types.Transaction, error) {
	return _IL1Nullifier.Contract.ClaimFailedDepositLegacyErc20Bridge(&_IL1Nullifier.TransactOpts, _depositSender, _l1Token, _amount, _l2TxHash, _l2BatchNumber, _l2MessageIndex, _l2TxNumberInBatch, _merkleProof)
}

// FinalizeDeposit is a paid mutator transaction binding the contract method 0x74beea82.
//
// Solidity: function finalizeDeposit((uint256,uint256,uint256,address,uint16,bytes,bytes32[]) _finalizeWithdrawalParams) returns()
func (_IL1Nullifier *IL1NullifierTransactor) FinalizeDeposit(opts *bind.TransactOpts, _finalizeWithdrawalParams FinalizeL1DepositParams) (*types.Transaction, error) {
	return _IL1Nullifier.contract.Transact(opts, "finalizeDeposit", _finalizeWithdrawalParams)
}

// FinalizeDeposit is a paid mutator transaction binding the contract method 0x74beea82.
//
// Solidity: function finalizeDeposit((uint256,uint256,uint256,address,uint16,bytes,bytes32[]) _finalizeWithdrawalParams) returns()
func (_IL1Nullifier *IL1NullifierSession) FinalizeDeposit(_finalizeWithdrawalParams FinalizeL1DepositParams) (*types.Transaction, error) {
	return _IL1Nullifier.Contract.FinalizeDeposit(&_IL1Nullifier.TransactOpts, _finalizeWithdrawalParams)
}

// FinalizeDeposit is a paid mutator transaction binding the contract method 0x74beea82.
//
// Solidity: function finalizeDeposit((uint256,uint256,uint256,address,uint16,bytes,bytes32[]) _finalizeWithdrawalParams) returns()
func (_IL1Nullifier *IL1NullifierTransactorSession) FinalizeDeposit(_finalizeWithdrawalParams FinalizeL1DepositParams) (*types.Transaction, error) {
	return _IL1Nullifier.Contract.FinalizeDeposit(&_IL1Nullifier.TransactOpts, _finalizeWithdrawalParams)
}

// FinalizeWithdrawal is a paid mutator transaction binding the contract method 0xc87325f1.
//
// Solidity: function finalizeWithdrawal(uint256 _chainId, uint256 _l2BatchNumber, uint256 _l2MessageIndex, uint16 _l2TxNumberInBatch, bytes _message, bytes32[] _merkleProof) returns()
func (_IL1Nullifier *IL1NullifierTransactor) FinalizeWithdrawal(opts *bind.TransactOpts, _chainId *big.Int, _l2BatchNumber *big.Int, _l2MessageIndex *big.Int, _l2TxNumberInBatch uint16, _message []byte, _merkleProof [][32]byte) (*types.Transaction, error) {
	return _IL1Nullifier.contract.Transact(opts, "finalizeWithdrawal", _chainId, _l2BatchNumber, _l2MessageIndex, _l2TxNumberInBatch, _message, _merkleProof)
}

// FinalizeWithdrawal is a paid mutator transaction binding the contract method 0xc87325f1.
//
// Solidity: function finalizeWithdrawal(uint256 _chainId, uint256 _l2BatchNumber, uint256 _l2MessageIndex, uint16 _l2TxNumberInBatch, bytes _message, bytes32[] _merkleProof) returns()
func (_IL1Nullifier *IL1NullifierSession) FinalizeWithdrawal(_chainId *big.Int, _l2BatchNumber *big.Int, _l2MessageIndex *big.Int, _l2TxNumberInBatch uint16, _message []byte, _merkleProof [][32]byte) (*types.Transaction, error) {
	return _IL1Nullifier.Contract.FinalizeWithdrawal(&_IL1Nullifier.TransactOpts, _chainId, _l2BatchNumber, _l2MessageIndex, _l2TxNumberInBatch, _message, _merkleProof)
}

// FinalizeWithdrawal is a paid mutator transaction binding the contract method 0xc87325f1.
//
// Solidity: function finalizeWithdrawal(uint256 _chainId, uint256 _l2BatchNumber, uint256 _l2MessageIndex, uint16 _l2TxNumberInBatch, bytes _message, bytes32[] _merkleProof) returns()
func (_IL1Nullifier *IL1NullifierTransactorSession) FinalizeWithdrawal(_chainId *big.Int, _l2BatchNumber *big.Int, _l2MessageIndex *big.Int, _l2TxNumberInBatch uint16, _message []byte, _merkleProof [][32]byte) (*types.Transaction, error) {
	return _IL1Nullifier.Contract.FinalizeWithdrawal(&_IL1Nullifier.TransactOpts, _chainId, _l2BatchNumber, _l2MessageIndex, _l2TxNumberInBatch, _message, _merkleProof)
}

// Initialize is a paid mutator transaction binding the contract method 0xf92ad219.
//
// Solidity: function initialize(address _owner, uint256 _eraPostDiamondUpgradeFirstBatch, uint256 _eraPostLegacyBridgeUpgradeFirstBatch, uint256 _eraLegacyBridgeLastDepositBatch, uint256 _eraLegacyBridgeLastDepositTxNumber) returns()
func (_IL1Nullifier *IL1NullifierTransactor) Initialize(opts *bind.TransactOpts, _owner common.Address, _eraPostDiamondUpgradeFirstBatch *big.Int, _eraPostLegacyBridgeUpgradeFirstBatch *big.Int, _eraLegacyBridgeLastDepositBatch *big.Int, _eraLegacyBridgeLastDepositTxNumber *big.Int) (*types.Transaction, error) {
	return _IL1Nullifier.contract.Transact(opts, "initialize", _owner, _eraPostDiamondUpgradeFirstBatch, _eraPostLegacyBridgeUpgradeFirstBatch, _eraLegacyBridgeLastDepositBatch, _eraLegacyBridgeLastDepositTxNumber)
}

// Initialize is a paid mutator transaction binding the contract method 0xf92ad219.
//
// Solidity: function initialize(address _owner, uint256 _eraPostDiamondUpgradeFirstBatch, uint256 _eraPostLegacyBridgeUpgradeFirstBatch, uint256 _eraLegacyBridgeLastDepositBatch, uint256 _eraLegacyBridgeLastDepositTxNumber) returns()
func (_IL1Nullifier *IL1NullifierSession) Initialize(_owner common.Address, _eraPostDiamondUpgradeFirstBatch *big.Int, _eraPostLegacyBridgeUpgradeFirstBatch *big.Int, _eraLegacyBridgeLastDepositBatch *big.Int, _eraLegacyBridgeLastDepositTxNumber *big.Int) (*types.Transaction, error) {
	return _IL1Nullifier.Contract.Initialize(&_IL1Nullifier.TransactOpts, _owner, _eraPostDiamondUpgradeFirstBatch, _eraPostLegacyBridgeUpgradeFirstBatch, _eraLegacyBridgeLastDepositBatch, _eraLegacyBridgeLastDepositTxNumber)
}

// Initialize is a paid mutator transaction binding the contract method 0xf92ad219.
//
// Solidity: function initialize(address _owner, uint256 _eraPostDiamondUpgradeFirstBatch, uint256 _eraPostLegacyBridgeUpgradeFirstBatch, uint256 _eraLegacyBridgeLastDepositBatch, uint256 _eraLegacyBridgeLastDepositTxNumber) returns()
func (_IL1Nullifier *IL1NullifierTransactorSession) Initialize(_owner common.Address, _eraPostDiamondUpgradeFirstBatch *big.Int, _eraPostLegacyBridgeUpgradeFirstBatch *big.Int, _eraLegacyBridgeLastDepositBatch *big.Int, _eraLegacyBridgeLastDepositTxNumber *big.Int) (*types.Transaction, error) {
	return _IL1Nullifier.Contract.Initialize(&_IL1Nullifier.TransactOpts, _owner, _eraPostDiamondUpgradeFirstBatch, _eraPostLegacyBridgeUpgradeFirstBatch, _eraLegacyBridgeLastDepositBatch, _eraLegacyBridgeLastDepositTxNumber)
}

// NullifyChainBalanceByNTV is a paid mutator transaction binding the contract method 0x5de097b1.
//
// Solidity: function nullifyChainBalanceByNTV(uint256 _chainId, address _token) returns()
func (_IL1Nullifier *IL1NullifierTransactor) NullifyChainBalanceByNTV(opts *bind.TransactOpts, _chainId *big.Int, _token common.Address) (*types.Transaction, error) {
	return _IL1Nullifier.contract.Transact(opts, "nullifyChainBalanceByNTV", _chainId, _token)
}

// NullifyChainBalanceByNTV is a paid mutator transaction binding the contract method 0x5de097b1.
//
// Solidity: function nullifyChainBalanceByNTV(uint256 _chainId, address _token) returns()
func (_IL1Nullifier *IL1NullifierSession) NullifyChainBalanceByNTV(_chainId *big.Int, _token common.Address) (*types.Transaction, error) {
	return _IL1Nullifier.Contract.NullifyChainBalanceByNTV(&_IL1Nullifier.TransactOpts, _chainId, _token)
}

// NullifyChainBalanceByNTV is a paid mutator transaction binding the contract method 0x5de097b1.
//
// Solidity: function nullifyChainBalanceByNTV(uint256 _chainId, address _token) returns()
func (_IL1Nullifier *IL1NullifierTransactorSession) NullifyChainBalanceByNTV(_chainId *big.Int, _token common.Address) (*types.Transaction, error) {
	return _IL1Nullifier.Contract.NullifyChainBalanceByNTV(&_IL1Nullifier.TransactOpts, _chainId, _token)
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns()
func (_IL1Nullifier *IL1NullifierTransactor) Pause(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IL1Nullifier.contract.Transact(opts, "pause")
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns()
func (_IL1Nullifier *IL1NullifierSession) Pause() (*types.Transaction, error) {
	return _IL1Nullifier.Contract.Pause(&_IL1Nullifier.TransactOpts)
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns()
func (_IL1Nullifier *IL1NullifierTransactorSession) Pause() (*types.Transaction, error) {
	return _IL1Nullifier.Contract.Pause(&_IL1Nullifier.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_IL1Nullifier *IL1NullifierTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IL1Nullifier.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_IL1Nullifier *IL1NullifierSession) RenounceOwnership() (*types.Transaction, error) {
	return _IL1Nullifier.Contract.RenounceOwnership(&_IL1Nullifier.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_IL1Nullifier *IL1NullifierTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _IL1Nullifier.Contract.RenounceOwnership(&_IL1Nullifier.TransactOpts)
}

// SetL1AssetRouter is a paid mutator transaction binding the contract method 0x780ce114.
//
// Solidity: function setL1AssetRouter(address _l1AssetRouter) returns()
func (_IL1Nullifier *IL1NullifierTransactor) SetL1AssetRouter(opts *bind.TransactOpts, _l1AssetRouter common.Address) (*types.Transaction, error) {
	return _IL1Nullifier.contract.Transact(opts, "setL1AssetRouter", _l1AssetRouter)
}

// SetL1AssetRouter is a paid mutator transaction binding the contract method 0x780ce114.
//
// Solidity: function setL1AssetRouter(address _l1AssetRouter) returns()
func (_IL1Nullifier *IL1NullifierSession) SetL1AssetRouter(_l1AssetRouter common.Address) (*types.Transaction, error) {
	return _IL1Nullifier.Contract.SetL1AssetRouter(&_IL1Nullifier.TransactOpts, _l1AssetRouter)
}

// SetL1AssetRouter is a paid mutator transaction binding the contract method 0x780ce114.
//
// Solidity: function setL1AssetRouter(address _l1AssetRouter) returns()
func (_IL1Nullifier *IL1NullifierTransactorSession) SetL1AssetRouter(_l1AssetRouter common.Address) (*types.Transaction, error) {
	return _IL1Nullifier.Contract.SetL1AssetRouter(&_IL1Nullifier.TransactOpts, _l1AssetRouter)
}

// SetL1Erc20Bridge is a paid mutator transaction binding the contract method 0x30bda03e.
//
// Solidity: function setL1Erc20Bridge(address _legacyBridge) returns()
func (_IL1Nullifier *IL1NullifierTransactor) SetL1Erc20Bridge(opts *bind.TransactOpts, _legacyBridge common.Address) (*types.Transaction, error) {
	return _IL1Nullifier.contract.Transact(opts, "setL1Erc20Bridge", _legacyBridge)
}

// SetL1Erc20Bridge is a paid mutator transaction binding the contract method 0x30bda03e.
//
// Solidity: function setL1Erc20Bridge(address _legacyBridge) returns()
func (_IL1Nullifier *IL1NullifierSession) SetL1Erc20Bridge(_legacyBridge common.Address) (*types.Transaction, error) {
	return _IL1Nullifier.Contract.SetL1Erc20Bridge(&_IL1Nullifier.TransactOpts, _legacyBridge)
}

// SetL1Erc20Bridge is a paid mutator transaction binding the contract method 0x30bda03e.
//
// Solidity: function setL1Erc20Bridge(address _legacyBridge) returns()
func (_IL1Nullifier *IL1NullifierTransactorSession) SetL1Erc20Bridge(_legacyBridge common.Address) (*types.Transaction, error) {
	return _IL1Nullifier.Contract.SetL1Erc20Bridge(&_IL1Nullifier.TransactOpts, _legacyBridge)
}

// SetL1NativeTokenVault is a paid mutator transaction binding the contract method 0xb7cc6f46.
//
// Solidity: function setL1NativeTokenVault(address _l1NativeTokenVault) returns()
func (_IL1Nullifier *IL1NullifierTransactor) SetL1NativeTokenVault(opts *bind.TransactOpts, _l1NativeTokenVault common.Address) (*types.Transaction, error) {
	return _IL1Nullifier.contract.Transact(opts, "setL1NativeTokenVault", _l1NativeTokenVault)
}

// SetL1NativeTokenVault is a paid mutator transaction binding the contract method 0xb7cc6f46.
//
// Solidity: function setL1NativeTokenVault(address _l1NativeTokenVault) returns()
func (_IL1Nullifier *IL1NullifierSession) SetL1NativeTokenVault(_l1NativeTokenVault common.Address) (*types.Transaction, error) {
	return _IL1Nullifier.Contract.SetL1NativeTokenVault(&_IL1Nullifier.TransactOpts, _l1NativeTokenVault)
}

// SetL1NativeTokenVault is a paid mutator transaction binding the contract method 0xb7cc6f46.
//
// Solidity: function setL1NativeTokenVault(address _l1NativeTokenVault) returns()
func (_IL1Nullifier *IL1NullifierTransactorSession) SetL1NativeTokenVault(_l1NativeTokenVault common.Address) (*types.Transaction, error) {
	return _IL1Nullifier.Contract.SetL1NativeTokenVault(&_IL1Nullifier.TransactOpts, _l1NativeTokenVault)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_IL1Nullifier *IL1NullifierTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _IL1Nullifier.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_IL1Nullifier *IL1NullifierSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _IL1Nullifier.Contract.TransferOwnership(&_IL1Nullifier.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_IL1Nullifier *IL1NullifierTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _IL1Nullifier.Contract.TransferOwnership(&_IL1Nullifier.TransactOpts, newOwner)
}

// TransferTokenToNTV is a paid mutator transaction binding the contract method 0x40a434d5.
//
// Solidity: function transferTokenToNTV(address _token) returns()
func (_IL1Nullifier *IL1NullifierTransactor) TransferTokenToNTV(opts *bind.TransactOpts, _token common.Address) (*types.Transaction, error) {
	return _IL1Nullifier.contract.Transact(opts, "transferTokenToNTV", _token)
}

// TransferTokenToNTV is a paid mutator transaction binding the contract method 0x40a434d5.
//
// Solidity: function transferTokenToNTV(address _token) returns()
func (_IL1Nullifier *IL1NullifierSession) TransferTokenToNTV(_token common.Address) (*types.Transaction, error) {
	return _IL1Nullifier.Contract.TransferTokenToNTV(&_IL1Nullifier.TransactOpts, _token)
}

// TransferTokenToNTV is a paid mutator transaction binding the contract method 0x40a434d5.
//
// Solidity: function transferTokenToNTV(address _token) returns()
func (_IL1Nullifier *IL1NullifierTransactorSession) TransferTokenToNTV(_token common.Address) (*types.Transaction, error) {
	return _IL1Nullifier.Contract.TransferTokenToNTV(&_IL1Nullifier.TransactOpts, _token)
}

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns()
func (_IL1Nullifier *IL1NullifierTransactor) Unpause(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IL1Nullifier.contract.Transact(opts, "unpause")
}

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns()
func (_IL1Nullifier *IL1NullifierSession) Unpause() (*types.Transaction, error) {
	return _IL1Nullifier.Contract.Unpause(&_IL1Nullifier.TransactOpts)
}

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns()
func (_IL1Nullifier *IL1NullifierTransactorSession) Unpause() (*types.Transaction, error) {
	return _IL1Nullifier.Contract.Unpause(&_IL1Nullifier.TransactOpts)
}

// IL1NullifierBridgehubDepositFinalizedIterator is returned from FilterBridgehubDepositFinalized and is used to iterate over the raw logs and unpacked data for BridgehubDepositFinalized events raised by the IL1Nullifier contract.
type IL1NullifierBridgehubDepositFinalizedIterator struct {
	Event *IL1NullifierBridgehubDepositFinalized // Event containing the contract specifics and raw log

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
func (it *IL1NullifierBridgehubDepositFinalizedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IL1NullifierBridgehubDepositFinalized)
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
		it.Event = new(IL1NullifierBridgehubDepositFinalized)
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
func (it *IL1NullifierBridgehubDepositFinalizedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IL1NullifierBridgehubDepositFinalizedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IL1NullifierBridgehubDepositFinalized represents a BridgehubDepositFinalized event raised by the IL1Nullifier contract.
type IL1NullifierBridgehubDepositFinalized struct {
	ChainId         *big.Int
	TxDataHash      [32]byte
	L2DepositTxHash [32]byte
	Raw             types.Log // Blockchain specific contextual infos
}

// FilterBridgehubDepositFinalized is a free log retrieval operation binding the contract event 0xe4def01b981193a97a9e81230d7b9f31812ceaf23f864a828a82c687911cb2df.
//
// Solidity: event BridgehubDepositFinalized(uint256 indexed chainId, bytes32 indexed txDataHash, bytes32 indexed l2DepositTxHash)
func (_IL1Nullifier *IL1NullifierFilterer) FilterBridgehubDepositFinalized(opts *bind.FilterOpts, chainId []*big.Int, txDataHash [][32]byte, l2DepositTxHash [][32]byte) (*IL1NullifierBridgehubDepositFinalizedIterator, error) {

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

	logs, sub, err := _IL1Nullifier.contract.FilterLogs(opts, "BridgehubDepositFinalized", chainIdRule, txDataHashRule, l2DepositTxHashRule)
	if err != nil {
		return nil, err
	}
	return &IL1NullifierBridgehubDepositFinalizedIterator{contract: _IL1Nullifier.contract, event: "BridgehubDepositFinalized", logs: logs, sub: sub}, nil
}

// WatchBridgehubDepositFinalized is a free log subscription operation binding the contract event 0xe4def01b981193a97a9e81230d7b9f31812ceaf23f864a828a82c687911cb2df.
//
// Solidity: event BridgehubDepositFinalized(uint256 indexed chainId, bytes32 indexed txDataHash, bytes32 indexed l2DepositTxHash)
func (_IL1Nullifier *IL1NullifierFilterer) WatchBridgehubDepositFinalized(opts *bind.WatchOpts, sink chan<- *IL1NullifierBridgehubDepositFinalized, chainId []*big.Int, txDataHash [][32]byte, l2DepositTxHash [][32]byte) (event.Subscription, error) {

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

	logs, sub, err := _IL1Nullifier.contract.WatchLogs(opts, "BridgehubDepositFinalized", chainIdRule, txDataHashRule, l2DepositTxHashRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IL1NullifierBridgehubDepositFinalized)
				if err := _IL1Nullifier.contract.UnpackLog(event, "BridgehubDepositFinalized", log); err != nil {
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
func (_IL1Nullifier *IL1NullifierFilterer) ParseBridgehubDepositFinalized(log types.Log) (*IL1NullifierBridgehubDepositFinalized, error) {
	event := new(IL1NullifierBridgehubDepositFinalized)
	if err := _IL1Nullifier.contract.UnpackLog(event, "BridgehubDepositFinalized", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// IL1NullifierInitializedIterator is returned from FilterInitialized and is used to iterate over the raw logs and unpacked data for Initialized events raised by the IL1Nullifier contract.
type IL1NullifierInitializedIterator struct {
	Event *IL1NullifierInitialized // Event containing the contract specifics and raw log

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
func (it *IL1NullifierInitializedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IL1NullifierInitialized)
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
		it.Event = new(IL1NullifierInitialized)
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
func (it *IL1NullifierInitializedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IL1NullifierInitializedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IL1NullifierInitialized represents a Initialized event raised by the IL1Nullifier contract.
type IL1NullifierInitialized struct {
	Version uint8
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterInitialized is a free log retrieval operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_IL1Nullifier *IL1NullifierFilterer) FilterInitialized(opts *bind.FilterOpts) (*IL1NullifierInitializedIterator, error) {

	logs, sub, err := _IL1Nullifier.contract.FilterLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return &IL1NullifierInitializedIterator{contract: _IL1Nullifier.contract, event: "Initialized", logs: logs, sub: sub}, nil
}

// WatchInitialized is a free log subscription operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_IL1Nullifier *IL1NullifierFilterer) WatchInitialized(opts *bind.WatchOpts, sink chan<- *IL1NullifierInitialized) (event.Subscription, error) {

	logs, sub, err := _IL1Nullifier.contract.WatchLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IL1NullifierInitialized)
				if err := _IL1Nullifier.contract.UnpackLog(event, "Initialized", log); err != nil {
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
func (_IL1Nullifier *IL1NullifierFilterer) ParseInitialized(log types.Log) (*IL1NullifierInitialized, error) {
	event := new(IL1NullifierInitialized)
	if err := _IL1Nullifier.contract.UnpackLog(event, "Initialized", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// IL1NullifierOwnershipTransferStartedIterator is returned from FilterOwnershipTransferStarted and is used to iterate over the raw logs and unpacked data for OwnershipTransferStarted events raised by the IL1Nullifier contract.
type IL1NullifierOwnershipTransferStartedIterator struct {
	Event *IL1NullifierOwnershipTransferStarted // Event containing the contract specifics and raw log

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
func (it *IL1NullifierOwnershipTransferStartedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IL1NullifierOwnershipTransferStarted)
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
		it.Event = new(IL1NullifierOwnershipTransferStarted)
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
func (it *IL1NullifierOwnershipTransferStartedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IL1NullifierOwnershipTransferStartedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IL1NullifierOwnershipTransferStarted represents a OwnershipTransferStarted event raised by the IL1Nullifier contract.
type IL1NullifierOwnershipTransferStarted struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferStarted is a free log retrieval operation binding the contract event 0x38d16b8cac22d99fc7c124b9cd0de2d3fa1faef420bfe791d8c362d765e22700.
//
// Solidity: event OwnershipTransferStarted(address indexed previousOwner, address indexed newOwner)
func (_IL1Nullifier *IL1NullifierFilterer) FilterOwnershipTransferStarted(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*IL1NullifierOwnershipTransferStartedIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _IL1Nullifier.contract.FilterLogs(opts, "OwnershipTransferStarted", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &IL1NullifierOwnershipTransferStartedIterator{contract: _IL1Nullifier.contract, event: "OwnershipTransferStarted", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferStarted is a free log subscription operation binding the contract event 0x38d16b8cac22d99fc7c124b9cd0de2d3fa1faef420bfe791d8c362d765e22700.
//
// Solidity: event OwnershipTransferStarted(address indexed previousOwner, address indexed newOwner)
func (_IL1Nullifier *IL1NullifierFilterer) WatchOwnershipTransferStarted(opts *bind.WatchOpts, sink chan<- *IL1NullifierOwnershipTransferStarted, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _IL1Nullifier.contract.WatchLogs(opts, "OwnershipTransferStarted", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IL1NullifierOwnershipTransferStarted)
				if err := _IL1Nullifier.contract.UnpackLog(event, "OwnershipTransferStarted", log); err != nil {
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
func (_IL1Nullifier *IL1NullifierFilterer) ParseOwnershipTransferStarted(log types.Log) (*IL1NullifierOwnershipTransferStarted, error) {
	event := new(IL1NullifierOwnershipTransferStarted)
	if err := _IL1Nullifier.contract.UnpackLog(event, "OwnershipTransferStarted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// IL1NullifierOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the IL1Nullifier contract.
type IL1NullifierOwnershipTransferredIterator struct {
	Event *IL1NullifierOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *IL1NullifierOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IL1NullifierOwnershipTransferred)
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
		it.Event = new(IL1NullifierOwnershipTransferred)
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
func (it *IL1NullifierOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IL1NullifierOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IL1NullifierOwnershipTransferred represents a OwnershipTransferred event raised by the IL1Nullifier contract.
type IL1NullifierOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_IL1Nullifier *IL1NullifierFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*IL1NullifierOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _IL1Nullifier.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &IL1NullifierOwnershipTransferredIterator{contract: _IL1Nullifier.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_IL1Nullifier *IL1NullifierFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *IL1NullifierOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _IL1Nullifier.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IL1NullifierOwnershipTransferred)
				if err := _IL1Nullifier.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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
func (_IL1Nullifier *IL1NullifierFilterer) ParseOwnershipTransferred(log types.Log) (*IL1NullifierOwnershipTransferred, error) {
	event := new(IL1NullifierOwnershipTransferred)
	if err := _IL1Nullifier.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// IL1NullifierPausedIterator is returned from FilterPaused and is used to iterate over the raw logs and unpacked data for Paused events raised by the IL1Nullifier contract.
type IL1NullifierPausedIterator struct {
	Event *IL1NullifierPaused // Event containing the contract specifics and raw log

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
func (it *IL1NullifierPausedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IL1NullifierPaused)
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
		it.Event = new(IL1NullifierPaused)
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
func (it *IL1NullifierPausedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IL1NullifierPausedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IL1NullifierPaused represents a Paused event raised by the IL1Nullifier contract.
type IL1NullifierPaused struct {
	Account common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterPaused is a free log retrieval operation binding the contract event 0x62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a258.
//
// Solidity: event Paused(address account)
func (_IL1Nullifier *IL1NullifierFilterer) FilterPaused(opts *bind.FilterOpts) (*IL1NullifierPausedIterator, error) {

	logs, sub, err := _IL1Nullifier.contract.FilterLogs(opts, "Paused")
	if err != nil {
		return nil, err
	}
	return &IL1NullifierPausedIterator{contract: _IL1Nullifier.contract, event: "Paused", logs: logs, sub: sub}, nil
}

// WatchPaused is a free log subscription operation binding the contract event 0x62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a258.
//
// Solidity: event Paused(address account)
func (_IL1Nullifier *IL1NullifierFilterer) WatchPaused(opts *bind.WatchOpts, sink chan<- *IL1NullifierPaused) (event.Subscription, error) {

	logs, sub, err := _IL1Nullifier.contract.WatchLogs(opts, "Paused")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IL1NullifierPaused)
				if err := _IL1Nullifier.contract.UnpackLog(event, "Paused", log); err != nil {
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
func (_IL1Nullifier *IL1NullifierFilterer) ParsePaused(log types.Log) (*IL1NullifierPaused, error) {
	event := new(IL1NullifierPaused)
	if err := _IL1Nullifier.contract.UnpackLog(event, "Paused", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// IL1NullifierUnpausedIterator is returned from FilterUnpaused and is used to iterate over the raw logs and unpacked data for Unpaused events raised by the IL1Nullifier contract.
type IL1NullifierUnpausedIterator struct {
	Event *IL1NullifierUnpaused // Event containing the contract specifics and raw log

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
func (it *IL1NullifierUnpausedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IL1NullifierUnpaused)
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
		it.Event = new(IL1NullifierUnpaused)
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
func (it *IL1NullifierUnpausedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IL1NullifierUnpausedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IL1NullifierUnpaused represents a Unpaused event raised by the IL1Nullifier contract.
type IL1NullifierUnpaused struct {
	Account common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterUnpaused is a free log retrieval operation binding the contract event 0x5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa.
//
// Solidity: event Unpaused(address account)
func (_IL1Nullifier *IL1NullifierFilterer) FilterUnpaused(opts *bind.FilterOpts) (*IL1NullifierUnpausedIterator, error) {

	logs, sub, err := _IL1Nullifier.contract.FilterLogs(opts, "Unpaused")
	if err != nil {
		return nil, err
	}
	return &IL1NullifierUnpausedIterator{contract: _IL1Nullifier.contract, event: "Unpaused", logs: logs, sub: sub}, nil
}

// WatchUnpaused is a free log subscription operation binding the contract event 0x5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa.
//
// Solidity: event Unpaused(address account)
func (_IL1Nullifier *IL1NullifierFilterer) WatchUnpaused(opts *bind.WatchOpts, sink chan<- *IL1NullifierUnpaused) (event.Subscription, error) {

	logs, sub, err := _IL1Nullifier.contract.WatchLogs(opts, "Unpaused")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IL1NullifierUnpaused)
				if err := _IL1Nullifier.contract.UnpackLog(event, "Unpaused", log); err != nil {
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
func (_IL1Nullifier *IL1NullifierFilterer) ParseUnpaused(log types.Log) (*IL1NullifierUnpaused, error) {
	event := new(IL1NullifierUnpaused)
	if err := _IL1Nullifier.contract.UnpackLog(event, "Unpaused", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
