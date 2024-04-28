// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package l1sharedbridge

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

// IL1SharedBridgeMetaData contains all meta data concerning the IL1SharedBridge contract.
var IL1SharedBridgeMetaData = &bind.MetaData{
	ABI: "[{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"chainId\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"l1Token\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"BridgehubDepositBaseTokenInitiated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"chainId\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"txDataHash\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"l2DepositTxHash\",\"type\":\"bytes32\"}],\"name\":\"BridgehubDepositFinalized\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"chainId\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"txDataHash\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"l1Token\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"BridgehubDepositInitiated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"chainId\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"l1Token\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"ClaimedFailedDepositSharedBridge\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"chainId\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"l2DepositTxHash\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"l1Token\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"LegacyDepositInitiated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"chainId\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"l1Token\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"WithdrawalFinalizedSharedBridge\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"bridgehub\",\"outputs\":[{\"internalType\":\"contractIBridgehub\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_chainId\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"_txDataHash\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"_txHash\",\"type\":\"bytes32\"}],\"name\":\"bridgehubConfirmL2Transaction\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_chainId\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"_prevMsgSender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_l2Value\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"_data\",\"type\":\"bytes\"}],\"name\":\"bridgehubDeposit\",\"outputs\":[{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"magicValue\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"l2Contract\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"l2Calldata\",\"type\":\"bytes\"},{\"internalType\":\"bytes[]\",\"name\":\"factoryDeps\",\"type\":\"bytes[]\"},{\"internalType\":\"bytes32\",\"name\":\"txDataHash\",\"type\":\"bytes32\"}],\"internalType\":\"structL2TransactionRequestTwoBridgesInner\",\"name\":\"request\",\"type\":\"tuple\"}],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_chainId\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"_prevMsgSender\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_l1Token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"bridgehubDepositBaseToken\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_chainId\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"_depositSender\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_l1Token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"_l2TxHash\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"_l2BatchNumber\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_l2MessageIndex\",\"type\":\"uint256\"},{\"internalType\":\"uint16\",\"name\":\"_l2TxNumberInBatch\",\"type\":\"uint16\"},{\"internalType\":\"bytes32[]\",\"name\":\"_merkleProof\",\"type\":\"bytes32[]\"}],\"name\":\"claimFailedDeposit\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_depositSender\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_l1Token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"_l2TxHash\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"_l2BatchNumber\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_l2MessageIndex\",\"type\":\"uint256\"},{\"internalType\":\"uint16\",\"name\":\"_l2TxNumberInBatch\",\"type\":\"uint16\"},{\"internalType\":\"bytes32[]\",\"name\":\"_merkleProof\",\"type\":\"bytes32[]\"}],\"name\":\"claimFailedDepositLegacyErc20Bridge\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_chainId\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"_l2TxHash\",\"type\":\"bytes32\"}],\"name\":\"depositHappened\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_msgSender\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_l2Receiver\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_l1Token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_l2TxGasLimit\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_l2TxGasPerPubdataByte\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"_refundRecipient\",\"type\":\"address\"}],\"name\":\"depositLegacyErc20Bridge\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"txHash\",\"type\":\"bytes32\"}],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_chainId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_l2BatchNumber\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_l2MessageIndex\",\"type\":\"uint256\"},{\"internalType\":\"uint16\",\"name\":\"_l2TxNumberInBatch\",\"type\":\"uint16\"},{\"internalType\":\"bytes\",\"name\":\"_message\",\"type\":\"bytes\"},{\"internalType\":\"bytes32[]\",\"name\":\"_merkleProof\",\"type\":\"bytes32[]\"}],\"name\":\"finalizeWithdrawal\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_l2BatchNumber\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_l2MessageIndex\",\"type\":\"uint256\"},{\"internalType\":\"uint16\",\"name\":\"_l2TxNumberInBatch\",\"type\":\"uint16\"},{\"internalType\":\"bytes\",\"name\":\"_message\",\"type\":\"bytes\"},{\"internalType\":\"bytes32[]\",\"name\":\"_merkleProof\",\"type\":\"bytes32[]\"}],\"name\":\"finalizeWithdrawalLegacyErc20Bridge\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"l1Receiver\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"l1Token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_chainId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_l2BatchNumber\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_l2MessageIndex\",\"type\":\"uint256\"}],\"name\":\"isWithdrawalFinalized\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"l1WethAddress\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_chainId\",\"type\":\"uint256\"}],\"name\":\"l2BridgeAddress\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"legacyBridge\",\"outputs\":[{\"internalType\":\"contractIL1ERC20Bridge\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_chainId\",\"type\":\"uint256\"}],\"name\":\"receiveEth\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_eraFirstPostUpgradeBatch\",\"type\":\"uint256\"}],\"name\":\"setEraFirstPostUpgradeBatch\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
}

// IL1SharedBridgeABI is the input ABI used to generate the binding from.
// Deprecated: Use IL1SharedBridgeMetaData.ABI instead.
var IL1SharedBridgeABI = IL1SharedBridgeMetaData.ABI

// IL1SharedBridge is an auto generated Go binding around an Ethereum contract.
type IL1SharedBridge struct {
	IL1SharedBridgeCaller     // Read-only binding to the contract
	IL1SharedBridgeTransactor // Write-only binding to the contract
	IL1SharedBridgeFilterer   // Log filterer for contract events
}

// IL1SharedBridgeCaller is an auto generated read-only Go binding around an Ethereum contract.
type IL1SharedBridgeCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IL1SharedBridgeTransactor is an auto generated write-only Go binding around an Ethereum contract.
type IL1SharedBridgeTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IL1SharedBridgeFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type IL1SharedBridgeFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IL1SharedBridgeSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type IL1SharedBridgeSession struct {
	Contract     *IL1SharedBridge  // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// IL1SharedBridgeCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type IL1SharedBridgeCallerSession struct {
	Contract *IL1SharedBridgeCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts          // Call options to use throughout this session
}

// IL1SharedBridgeTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type IL1SharedBridgeTransactorSession struct {
	Contract     *IL1SharedBridgeTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts          // Transaction auth options to use throughout this session
}

// IL1SharedBridgeRaw is an auto generated low-level Go binding around an Ethereum contract.
type IL1SharedBridgeRaw struct {
	Contract *IL1SharedBridge // Generic contract binding to access the raw methods on
}

// IL1SharedBridgeCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type IL1SharedBridgeCallerRaw struct {
	Contract *IL1SharedBridgeCaller // Generic read-only contract binding to access the raw methods on
}

// IL1SharedBridgeTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type IL1SharedBridgeTransactorRaw struct {
	Contract *IL1SharedBridgeTransactor // Generic write-only contract binding to access the raw methods on
}

// NewIL1SharedBridge creates a new instance of IL1SharedBridge, bound to a specific deployed contract.
func NewIL1SharedBridge(address common.Address, backend bind.ContractBackend) (*IL1SharedBridge, error) {
	contract, err := bindIL1SharedBridge(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &IL1SharedBridge{IL1SharedBridgeCaller: IL1SharedBridgeCaller{contract: contract}, IL1SharedBridgeTransactor: IL1SharedBridgeTransactor{contract: contract}, IL1SharedBridgeFilterer: IL1SharedBridgeFilterer{contract: contract}}, nil
}

// NewIL1SharedBridgeCaller creates a new read-only instance of IL1SharedBridge, bound to a specific deployed contract.
func NewIL1SharedBridgeCaller(address common.Address, caller bind.ContractCaller) (*IL1SharedBridgeCaller, error) {
	contract, err := bindIL1SharedBridge(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &IL1SharedBridgeCaller{contract: contract}, nil
}

// NewIL1SharedBridgeTransactor creates a new write-only instance of IL1SharedBridge, bound to a specific deployed contract.
func NewIL1SharedBridgeTransactor(address common.Address, transactor bind.ContractTransactor) (*IL1SharedBridgeTransactor, error) {
	contract, err := bindIL1SharedBridge(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &IL1SharedBridgeTransactor{contract: contract}, nil
}

// NewIL1SharedBridgeFilterer creates a new log filterer instance of IL1SharedBridge, bound to a specific deployed contract.
func NewIL1SharedBridgeFilterer(address common.Address, filterer bind.ContractFilterer) (*IL1SharedBridgeFilterer, error) {
	contract, err := bindIL1SharedBridge(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &IL1SharedBridgeFilterer{contract: contract}, nil
}

// bindIL1SharedBridge binds a generic wrapper to an already deployed contract.
func bindIL1SharedBridge(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := IL1SharedBridgeMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IL1SharedBridge *IL1SharedBridgeRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IL1SharedBridge.Contract.IL1SharedBridgeCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IL1SharedBridge *IL1SharedBridgeRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IL1SharedBridge.Contract.IL1SharedBridgeTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IL1SharedBridge *IL1SharedBridgeRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IL1SharedBridge.Contract.IL1SharedBridgeTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IL1SharedBridge *IL1SharedBridgeCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IL1SharedBridge.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IL1SharedBridge *IL1SharedBridgeTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IL1SharedBridge.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IL1SharedBridge *IL1SharedBridgeTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IL1SharedBridge.Contract.contract.Transact(opts, method, params...)
}

// Bridgehub is a free data retrieval call binding the contract method 0x5d38962d.
//
// Solidity: function bridgehub() view returns(address)
func (_IL1SharedBridge *IL1SharedBridgeCaller) Bridgehub(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _IL1SharedBridge.contract.Call(opts, &out, "bridgehub")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Bridgehub is a free data retrieval call binding the contract method 0x5d38962d.
//
// Solidity: function bridgehub() view returns(address)
func (_IL1SharedBridge *IL1SharedBridgeSession) Bridgehub() (common.Address, error) {
	return _IL1SharedBridge.Contract.Bridgehub(&_IL1SharedBridge.CallOpts)
}

// Bridgehub is a free data retrieval call binding the contract method 0x5d38962d.
//
// Solidity: function bridgehub() view returns(address)
func (_IL1SharedBridge *IL1SharedBridgeCallerSession) Bridgehub() (common.Address, error) {
	return _IL1SharedBridge.Contract.Bridgehub(&_IL1SharedBridge.CallOpts)
}

// DepositHappened is a free data retrieval call binding the contract method 0x9fa8826b.
//
// Solidity: function depositHappened(uint256 _chainId, bytes32 _l2TxHash) view returns(bytes32)
func (_IL1SharedBridge *IL1SharedBridgeCaller) DepositHappened(opts *bind.CallOpts, _chainId *big.Int, _l2TxHash [32]byte) ([32]byte, error) {
	var out []interface{}
	err := _IL1SharedBridge.contract.Call(opts, &out, "depositHappened", _chainId, _l2TxHash)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// DepositHappened is a free data retrieval call binding the contract method 0x9fa8826b.
//
// Solidity: function depositHappened(uint256 _chainId, bytes32 _l2TxHash) view returns(bytes32)
func (_IL1SharedBridge *IL1SharedBridgeSession) DepositHappened(_chainId *big.Int, _l2TxHash [32]byte) ([32]byte, error) {
	return _IL1SharedBridge.Contract.DepositHappened(&_IL1SharedBridge.CallOpts, _chainId, _l2TxHash)
}

// DepositHappened is a free data retrieval call binding the contract method 0x9fa8826b.
//
// Solidity: function depositHappened(uint256 _chainId, bytes32 _l2TxHash) view returns(bytes32)
func (_IL1SharedBridge *IL1SharedBridgeCallerSession) DepositHappened(_chainId *big.Int, _l2TxHash [32]byte) ([32]byte, error) {
	return _IL1SharedBridge.Contract.DepositHappened(&_IL1SharedBridge.CallOpts, _chainId, _l2TxHash)
}

// IsWithdrawalFinalized is a free data retrieval call binding the contract method 0x8f31f052.
//
// Solidity: function isWithdrawalFinalized(uint256 _chainId, uint256 _l2BatchNumber, uint256 _l2MessageIndex) view returns(bool)
func (_IL1SharedBridge *IL1SharedBridgeCaller) IsWithdrawalFinalized(opts *bind.CallOpts, _chainId *big.Int, _l2BatchNumber *big.Int, _l2MessageIndex *big.Int) (bool, error) {
	var out []interface{}
	err := _IL1SharedBridge.contract.Call(opts, &out, "isWithdrawalFinalized", _chainId, _l2BatchNumber, _l2MessageIndex)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsWithdrawalFinalized is a free data retrieval call binding the contract method 0x8f31f052.
//
// Solidity: function isWithdrawalFinalized(uint256 _chainId, uint256 _l2BatchNumber, uint256 _l2MessageIndex) view returns(bool)
func (_IL1SharedBridge *IL1SharedBridgeSession) IsWithdrawalFinalized(_chainId *big.Int, _l2BatchNumber *big.Int, _l2MessageIndex *big.Int) (bool, error) {
	return _IL1SharedBridge.Contract.IsWithdrawalFinalized(&_IL1SharedBridge.CallOpts, _chainId, _l2BatchNumber, _l2MessageIndex)
}

// IsWithdrawalFinalized is a free data retrieval call binding the contract method 0x8f31f052.
//
// Solidity: function isWithdrawalFinalized(uint256 _chainId, uint256 _l2BatchNumber, uint256 _l2MessageIndex) view returns(bool)
func (_IL1SharedBridge *IL1SharedBridgeCallerSession) IsWithdrawalFinalized(_chainId *big.Int, _l2BatchNumber *big.Int, _l2MessageIndex *big.Int) (bool, error) {
	return _IL1SharedBridge.Contract.IsWithdrawalFinalized(&_IL1SharedBridge.CallOpts, _chainId, _l2BatchNumber, _l2MessageIndex)
}

// L1WethAddress is a free data retrieval call binding the contract method 0x6ace8bbb.
//
// Solidity: function l1WethAddress() view returns(address)
func (_IL1SharedBridge *IL1SharedBridgeCaller) L1WethAddress(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _IL1SharedBridge.contract.Call(opts, &out, "l1WethAddress")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// L1WethAddress is a free data retrieval call binding the contract method 0x6ace8bbb.
//
// Solidity: function l1WethAddress() view returns(address)
func (_IL1SharedBridge *IL1SharedBridgeSession) L1WethAddress() (common.Address, error) {
	return _IL1SharedBridge.Contract.L1WethAddress(&_IL1SharedBridge.CallOpts)
}

// L1WethAddress is a free data retrieval call binding the contract method 0x6ace8bbb.
//
// Solidity: function l1WethAddress() view returns(address)
func (_IL1SharedBridge *IL1SharedBridgeCallerSession) L1WethAddress() (common.Address, error) {
	return _IL1SharedBridge.Contract.L1WethAddress(&_IL1SharedBridge.CallOpts)
}

// L2BridgeAddress is a free data retrieval call binding the contract method 0x07ee9355.
//
// Solidity: function l2BridgeAddress(uint256 _chainId) view returns(address)
func (_IL1SharedBridge *IL1SharedBridgeCaller) L2BridgeAddress(opts *bind.CallOpts, _chainId *big.Int) (common.Address, error) {
	var out []interface{}
	err := _IL1SharedBridge.contract.Call(opts, &out, "l2BridgeAddress", _chainId)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// L2BridgeAddress is a free data retrieval call binding the contract method 0x07ee9355.
//
// Solidity: function l2BridgeAddress(uint256 _chainId) view returns(address)
func (_IL1SharedBridge *IL1SharedBridgeSession) L2BridgeAddress(_chainId *big.Int) (common.Address, error) {
	return _IL1SharedBridge.Contract.L2BridgeAddress(&_IL1SharedBridge.CallOpts, _chainId)
}

// L2BridgeAddress is a free data retrieval call binding the contract method 0x07ee9355.
//
// Solidity: function l2BridgeAddress(uint256 _chainId) view returns(address)
func (_IL1SharedBridge *IL1SharedBridgeCallerSession) L2BridgeAddress(_chainId *big.Int) (common.Address, error) {
	return _IL1SharedBridge.Contract.L2BridgeAddress(&_IL1SharedBridge.CallOpts, _chainId)
}

// LegacyBridge is a free data retrieval call binding the contract method 0x6e9d7899.
//
// Solidity: function legacyBridge() view returns(address)
func (_IL1SharedBridge *IL1SharedBridgeCaller) LegacyBridge(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _IL1SharedBridge.contract.Call(opts, &out, "legacyBridge")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// LegacyBridge is a free data retrieval call binding the contract method 0x6e9d7899.
//
// Solidity: function legacyBridge() view returns(address)
func (_IL1SharedBridge *IL1SharedBridgeSession) LegacyBridge() (common.Address, error) {
	return _IL1SharedBridge.Contract.LegacyBridge(&_IL1SharedBridge.CallOpts)
}

// LegacyBridge is a free data retrieval call binding the contract method 0x6e9d7899.
//
// Solidity: function legacyBridge() view returns(address)
func (_IL1SharedBridge *IL1SharedBridgeCallerSession) LegacyBridge() (common.Address, error) {
	return _IL1SharedBridge.Contract.LegacyBridge(&_IL1SharedBridge.CallOpts)
}

// BridgehubConfirmL2Transaction is a paid mutator transaction binding the contract method 0x8eb7db57.
//
// Solidity: function bridgehubConfirmL2Transaction(uint256 _chainId, bytes32 _txDataHash, bytes32 _txHash) returns()
func (_IL1SharedBridge *IL1SharedBridgeTransactor) BridgehubConfirmL2Transaction(opts *bind.TransactOpts, _chainId *big.Int, _txDataHash [32]byte, _txHash [32]byte) (*types.Transaction, error) {
	return _IL1SharedBridge.contract.Transact(opts, "bridgehubConfirmL2Transaction", _chainId, _txDataHash, _txHash)
}

// BridgehubConfirmL2Transaction is a paid mutator transaction binding the contract method 0x8eb7db57.
//
// Solidity: function bridgehubConfirmL2Transaction(uint256 _chainId, bytes32 _txDataHash, bytes32 _txHash) returns()
func (_IL1SharedBridge *IL1SharedBridgeSession) BridgehubConfirmL2Transaction(_chainId *big.Int, _txDataHash [32]byte, _txHash [32]byte) (*types.Transaction, error) {
	return _IL1SharedBridge.Contract.BridgehubConfirmL2Transaction(&_IL1SharedBridge.TransactOpts, _chainId, _txDataHash, _txHash)
}

// BridgehubConfirmL2Transaction is a paid mutator transaction binding the contract method 0x8eb7db57.
//
// Solidity: function bridgehubConfirmL2Transaction(uint256 _chainId, bytes32 _txDataHash, bytes32 _txHash) returns()
func (_IL1SharedBridge *IL1SharedBridgeTransactorSession) BridgehubConfirmL2Transaction(_chainId *big.Int, _txDataHash [32]byte, _txHash [32]byte) (*types.Transaction, error) {
	return _IL1SharedBridge.Contract.BridgehubConfirmL2Transaction(&_IL1SharedBridge.TransactOpts, _chainId, _txDataHash, _txHash)
}

// BridgehubDeposit is a paid mutator transaction binding the contract method 0xca408c23.
//
// Solidity: function bridgehubDeposit(uint256 _chainId, address _prevMsgSender, uint256 _l2Value, bytes _data) payable returns((bytes32,address,bytes,bytes[],bytes32) request)
func (_IL1SharedBridge *IL1SharedBridgeTransactor) BridgehubDeposit(opts *bind.TransactOpts, _chainId *big.Int, _prevMsgSender common.Address, _l2Value *big.Int, _data []byte) (*types.Transaction, error) {
	return _IL1SharedBridge.contract.Transact(opts, "bridgehubDeposit", _chainId, _prevMsgSender, _l2Value, _data)
}

// BridgehubDeposit is a paid mutator transaction binding the contract method 0xca408c23.
//
// Solidity: function bridgehubDeposit(uint256 _chainId, address _prevMsgSender, uint256 _l2Value, bytes _data) payable returns((bytes32,address,bytes,bytes[],bytes32) request)
func (_IL1SharedBridge *IL1SharedBridgeSession) BridgehubDeposit(_chainId *big.Int, _prevMsgSender common.Address, _l2Value *big.Int, _data []byte) (*types.Transaction, error) {
	return _IL1SharedBridge.Contract.BridgehubDeposit(&_IL1SharedBridge.TransactOpts, _chainId, _prevMsgSender, _l2Value, _data)
}

// BridgehubDeposit is a paid mutator transaction binding the contract method 0xca408c23.
//
// Solidity: function bridgehubDeposit(uint256 _chainId, address _prevMsgSender, uint256 _l2Value, bytes _data) payable returns((bytes32,address,bytes,bytes[],bytes32) request)
func (_IL1SharedBridge *IL1SharedBridgeTransactorSession) BridgehubDeposit(_chainId *big.Int, _prevMsgSender common.Address, _l2Value *big.Int, _data []byte) (*types.Transaction, error) {
	return _IL1SharedBridge.Contract.BridgehubDeposit(&_IL1SharedBridge.TransactOpts, _chainId, _prevMsgSender, _l2Value, _data)
}

// BridgehubDepositBaseToken is a paid mutator transaction binding the contract method 0x2c4f2a58.
//
// Solidity: function bridgehubDepositBaseToken(uint256 _chainId, address _prevMsgSender, address _l1Token, uint256 _amount) payable returns()
func (_IL1SharedBridge *IL1SharedBridgeTransactor) BridgehubDepositBaseToken(opts *bind.TransactOpts, _chainId *big.Int, _prevMsgSender common.Address, _l1Token common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _IL1SharedBridge.contract.Transact(opts, "bridgehubDepositBaseToken", _chainId, _prevMsgSender, _l1Token, _amount)
}

// BridgehubDepositBaseToken is a paid mutator transaction binding the contract method 0x2c4f2a58.
//
// Solidity: function bridgehubDepositBaseToken(uint256 _chainId, address _prevMsgSender, address _l1Token, uint256 _amount) payable returns()
func (_IL1SharedBridge *IL1SharedBridgeSession) BridgehubDepositBaseToken(_chainId *big.Int, _prevMsgSender common.Address, _l1Token common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _IL1SharedBridge.Contract.BridgehubDepositBaseToken(&_IL1SharedBridge.TransactOpts, _chainId, _prevMsgSender, _l1Token, _amount)
}

// BridgehubDepositBaseToken is a paid mutator transaction binding the contract method 0x2c4f2a58.
//
// Solidity: function bridgehubDepositBaseToken(uint256 _chainId, address _prevMsgSender, address _l1Token, uint256 _amount) payable returns()
func (_IL1SharedBridge *IL1SharedBridgeTransactorSession) BridgehubDepositBaseToken(_chainId *big.Int, _prevMsgSender common.Address, _l1Token common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _IL1SharedBridge.Contract.BridgehubDepositBaseToken(&_IL1SharedBridge.TransactOpts, _chainId, _prevMsgSender, _l1Token, _amount)
}

// ClaimFailedDeposit is a paid mutator transaction binding the contract method 0xc0991525.
//
// Solidity: function claimFailedDeposit(uint256 _chainId, address _depositSender, address _l1Token, uint256 _amount, bytes32 _l2TxHash, uint256 _l2BatchNumber, uint256 _l2MessageIndex, uint16 _l2TxNumberInBatch, bytes32[] _merkleProof) returns()
func (_IL1SharedBridge *IL1SharedBridgeTransactor) ClaimFailedDeposit(opts *bind.TransactOpts, _chainId *big.Int, _depositSender common.Address, _l1Token common.Address, _amount *big.Int, _l2TxHash [32]byte, _l2BatchNumber *big.Int, _l2MessageIndex *big.Int, _l2TxNumberInBatch uint16, _merkleProof [][32]byte) (*types.Transaction, error) {
	return _IL1SharedBridge.contract.Transact(opts, "claimFailedDeposit", _chainId, _depositSender, _l1Token, _amount, _l2TxHash, _l2BatchNumber, _l2MessageIndex, _l2TxNumberInBatch, _merkleProof)
}

// ClaimFailedDeposit is a paid mutator transaction binding the contract method 0xc0991525.
//
// Solidity: function claimFailedDeposit(uint256 _chainId, address _depositSender, address _l1Token, uint256 _amount, bytes32 _l2TxHash, uint256 _l2BatchNumber, uint256 _l2MessageIndex, uint16 _l2TxNumberInBatch, bytes32[] _merkleProof) returns()
func (_IL1SharedBridge *IL1SharedBridgeSession) ClaimFailedDeposit(_chainId *big.Int, _depositSender common.Address, _l1Token common.Address, _amount *big.Int, _l2TxHash [32]byte, _l2BatchNumber *big.Int, _l2MessageIndex *big.Int, _l2TxNumberInBatch uint16, _merkleProof [][32]byte) (*types.Transaction, error) {
	return _IL1SharedBridge.Contract.ClaimFailedDeposit(&_IL1SharedBridge.TransactOpts, _chainId, _depositSender, _l1Token, _amount, _l2TxHash, _l2BatchNumber, _l2MessageIndex, _l2TxNumberInBatch, _merkleProof)
}

// ClaimFailedDeposit is a paid mutator transaction binding the contract method 0xc0991525.
//
// Solidity: function claimFailedDeposit(uint256 _chainId, address _depositSender, address _l1Token, uint256 _amount, bytes32 _l2TxHash, uint256 _l2BatchNumber, uint256 _l2MessageIndex, uint16 _l2TxNumberInBatch, bytes32[] _merkleProof) returns()
func (_IL1SharedBridge *IL1SharedBridgeTransactorSession) ClaimFailedDeposit(_chainId *big.Int, _depositSender common.Address, _l1Token common.Address, _amount *big.Int, _l2TxHash [32]byte, _l2BatchNumber *big.Int, _l2MessageIndex *big.Int, _l2TxNumberInBatch uint16, _merkleProof [][32]byte) (*types.Transaction, error) {
	return _IL1SharedBridge.Contract.ClaimFailedDeposit(&_IL1SharedBridge.TransactOpts, _chainId, _depositSender, _l1Token, _amount, _l2TxHash, _l2BatchNumber, _l2MessageIndex, _l2TxNumberInBatch, _merkleProof)
}

// ClaimFailedDepositLegacyErc20Bridge is a paid mutator transaction binding the contract method 0x8fbb3711.
//
// Solidity: function claimFailedDepositLegacyErc20Bridge(address _depositSender, address _l1Token, uint256 _amount, bytes32 _l2TxHash, uint256 _l2BatchNumber, uint256 _l2MessageIndex, uint16 _l2TxNumberInBatch, bytes32[] _merkleProof) returns()
func (_IL1SharedBridge *IL1SharedBridgeTransactor) ClaimFailedDepositLegacyErc20Bridge(opts *bind.TransactOpts, _depositSender common.Address, _l1Token common.Address, _amount *big.Int, _l2TxHash [32]byte, _l2BatchNumber *big.Int, _l2MessageIndex *big.Int, _l2TxNumberInBatch uint16, _merkleProof [][32]byte) (*types.Transaction, error) {
	return _IL1SharedBridge.contract.Transact(opts, "claimFailedDepositLegacyErc20Bridge", _depositSender, _l1Token, _amount, _l2TxHash, _l2BatchNumber, _l2MessageIndex, _l2TxNumberInBatch, _merkleProof)
}

// ClaimFailedDepositLegacyErc20Bridge is a paid mutator transaction binding the contract method 0x8fbb3711.
//
// Solidity: function claimFailedDepositLegacyErc20Bridge(address _depositSender, address _l1Token, uint256 _amount, bytes32 _l2TxHash, uint256 _l2BatchNumber, uint256 _l2MessageIndex, uint16 _l2TxNumberInBatch, bytes32[] _merkleProof) returns()
func (_IL1SharedBridge *IL1SharedBridgeSession) ClaimFailedDepositLegacyErc20Bridge(_depositSender common.Address, _l1Token common.Address, _amount *big.Int, _l2TxHash [32]byte, _l2BatchNumber *big.Int, _l2MessageIndex *big.Int, _l2TxNumberInBatch uint16, _merkleProof [][32]byte) (*types.Transaction, error) {
	return _IL1SharedBridge.Contract.ClaimFailedDepositLegacyErc20Bridge(&_IL1SharedBridge.TransactOpts, _depositSender, _l1Token, _amount, _l2TxHash, _l2BatchNumber, _l2MessageIndex, _l2TxNumberInBatch, _merkleProof)
}

// ClaimFailedDepositLegacyErc20Bridge is a paid mutator transaction binding the contract method 0x8fbb3711.
//
// Solidity: function claimFailedDepositLegacyErc20Bridge(address _depositSender, address _l1Token, uint256 _amount, bytes32 _l2TxHash, uint256 _l2BatchNumber, uint256 _l2MessageIndex, uint16 _l2TxNumberInBatch, bytes32[] _merkleProof) returns()
func (_IL1SharedBridge *IL1SharedBridgeTransactorSession) ClaimFailedDepositLegacyErc20Bridge(_depositSender common.Address, _l1Token common.Address, _amount *big.Int, _l2TxHash [32]byte, _l2BatchNumber *big.Int, _l2MessageIndex *big.Int, _l2TxNumberInBatch uint16, _merkleProof [][32]byte) (*types.Transaction, error) {
	return _IL1SharedBridge.Contract.ClaimFailedDepositLegacyErc20Bridge(&_IL1SharedBridge.TransactOpts, _depositSender, _l1Token, _amount, _l2TxHash, _l2BatchNumber, _l2MessageIndex, _l2TxNumberInBatch, _merkleProof)
}

// DepositLegacyErc20Bridge is a paid mutator transaction binding the contract method 0x9e6ea417.
//
// Solidity: function depositLegacyErc20Bridge(address _msgSender, address _l2Receiver, address _l1Token, uint256 _amount, uint256 _l2TxGasLimit, uint256 _l2TxGasPerPubdataByte, address _refundRecipient) payable returns(bytes32 txHash)
func (_IL1SharedBridge *IL1SharedBridgeTransactor) DepositLegacyErc20Bridge(opts *bind.TransactOpts, _msgSender common.Address, _l2Receiver common.Address, _l1Token common.Address, _amount *big.Int, _l2TxGasLimit *big.Int, _l2TxGasPerPubdataByte *big.Int, _refundRecipient common.Address) (*types.Transaction, error) {
	return _IL1SharedBridge.contract.Transact(opts, "depositLegacyErc20Bridge", _msgSender, _l2Receiver, _l1Token, _amount, _l2TxGasLimit, _l2TxGasPerPubdataByte, _refundRecipient)
}

// DepositLegacyErc20Bridge is a paid mutator transaction binding the contract method 0x9e6ea417.
//
// Solidity: function depositLegacyErc20Bridge(address _msgSender, address _l2Receiver, address _l1Token, uint256 _amount, uint256 _l2TxGasLimit, uint256 _l2TxGasPerPubdataByte, address _refundRecipient) payable returns(bytes32 txHash)
func (_IL1SharedBridge *IL1SharedBridgeSession) DepositLegacyErc20Bridge(_msgSender common.Address, _l2Receiver common.Address, _l1Token common.Address, _amount *big.Int, _l2TxGasLimit *big.Int, _l2TxGasPerPubdataByte *big.Int, _refundRecipient common.Address) (*types.Transaction, error) {
	return _IL1SharedBridge.Contract.DepositLegacyErc20Bridge(&_IL1SharedBridge.TransactOpts, _msgSender, _l2Receiver, _l1Token, _amount, _l2TxGasLimit, _l2TxGasPerPubdataByte, _refundRecipient)
}

// DepositLegacyErc20Bridge is a paid mutator transaction binding the contract method 0x9e6ea417.
//
// Solidity: function depositLegacyErc20Bridge(address _msgSender, address _l2Receiver, address _l1Token, uint256 _amount, uint256 _l2TxGasLimit, uint256 _l2TxGasPerPubdataByte, address _refundRecipient) payable returns(bytes32 txHash)
func (_IL1SharedBridge *IL1SharedBridgeTransactorSession) DepositLegacyErc20Bridge(_msgSender common.Address, _l2Receiver common.Address, _l1Token common.Address, _amount *big.Int, _l2TxGasLimit *big.Int, _l2TxGasPerPubdataByte *big.Int, _refundRecipient common.Address) (*types.Transaction, error) {
	return _IL1SharedBridge.Contract.DepositLegacyErc20Bridge(&_IL1SharedBridge.TransactOpts, _msgSender, _l2Receiver, _l1Token, _amount, _l2TxGasLimit, _l2TxGasPerPubdataByte, _refundRecipient)
}

// FinalizeWithdrawal is a paid mutator transaction binding the contract method 0xc87325f1.
//
// Solidity: function finalizeWithdrawal(uint256 _chainId, uint256 _l2BatchNumber, uint256 _l2MessageIndex, uint16 _l2TxNumberInBatch, bytes _message, bytes32[] _merkleProof) returns()
func (_IL1SharedBridge *IL1SharedBridgeTransactor) FinalizeWithdrawal(opts *bind.TransactOpts, _chainId *big.Int, _l2BatchNumber *big.Int, _l2MessageIndex *big.Int, _l2TxNumberInBatch uint16, _message []byte, _merkleProof [][32]byte) (*types.Transaction, error) {
	return _IL1SharedBridge.contract.Transact(opts, "finalizeWithdrawal", _chainId, _l2BatchNumber, _l2MessageIndex, _l2TxNumberInBatch, _message, _merkleProof)
}

// FinalizeWithdrawal is a paid mutator transaction binding the contract method 0xc87325f1.
//
// Solidity: function finalizeWithdrawal(uint256 _chainId, uint256 _l2BatchNumber, uint256 _l2MessageIndex, uint16 _l2TxNumberInBatch, bytes _message, bytes32[] _merkleProof) returns()
func (_IL1SharedBridge *IL1SharedBridgeSession) FinalizeWithdrawal(_chainId *big.Int, _l2BatchNumber *big.Int, _l2MessageIndex *big.Int, _l2TxNumberInBatch uint16, _message []byte, _merkleProof [][32]byte) (*types.Transaction, error) {
	return _IL1SharedBridge.Contract.FinalizeWithdrawal(&_IL1SharedBridge.TransactOpts, _chainId, _l2BatchNumber, _l2MessageIndex, _l2TxNumberInBatch, _message, _merkleProof)
}

// FinalizeWithdrawal is a paid mutator transaction binding the contract method 0xc87325f1.
//
// Solidity: function finalizeWithdrawal(uint256 _chainId, uint256 _l2BatchNumber, uint256 _l2MessageIndex, uint16 _l2TxNumberInBatch, bytes _message, bytes32[] _merkleProof) returns()
func (_IL1SharedBridge *IL1SharedBridgeTransactorSession) FinalizeWithdrawal(_chainId *big.Int, _l2BatchNumber *big.Int, _l2MessageIndex *big.Int, _l2TxNumberInBatch uint16, _message []byte, _merkleProof [][32]byte) (*types.Transaction, error) {
	return _IL1SharedBridge.Contract.FinalizeWithdrawal(&_IL1SharedBridge.TransactOpts, _chainId, _l2BatchNumber, _l2MessageIndex, _l2TxNumberInBatch, _message, _merkleProof)
}

// FinalizeWithdrawalLegacyErc20Bridge is a paid mutator transaction binding the contract method 0x7ab08472.
//
// Solidity: function finalizeWithdrawalLegacyErc20Bridge(uint256 _l2BatchNumber, uint256 _l2MessageIndex, uint16 _l2TxNumberInBatch, bytes _message, bytes32[] _merkleProof) returns(address l1Receiver, address l1Token, uint256 amount)
func (_IL1SharedBridge *IL1SharedBridgeTransactor) FinalizeWithdrawalLegacyErc20Bridge(opts *bind.TransactOpts, _l2BatchNumber *big.Int, _l2MessageIndex *big.Int, _l2TxNumberInBatch uint16, _message []byte, _merkleProof [][32]byte) (*types.Transaction, error) {
	return _IL1SharedBridge.contract.Transact(opts, "finalizeWithdrawalLegacyErc20Bridge", _l2BatchNumber, _l2MessageIndex, _l2TxNumberInBatch, _message, _merkleProof)
}

// FinalizeWithdrawalLegacyErc20Bridge is a paid mutator transaction binding the contract method 0x7ab08472.
//
// Solidity: function finalizeWithdrawalLegacyErc20Bridge(uint256 _l2BatchNumber, uint256 _l2MessageIndex, uint16 _l2TxNumberInBatch, bytes _message, bytes32[] _merkleProof) returns(address l1Receiver, address l1Token, uint256 amount)
func (_IL1SharedBridge *IL1SharedBridgeSession) FinalizeWithdrawalLegacyErc20Bridge(_l2BatchNumber *big.Int, _l2MessageIndex *big.Int, _l2TxNumberInBatch uint16, _message []byte, _merkleProof [][32]byte) (*types.Transaction, error) {
	return _IL1SharedBridge.Contract.FinalizeWithdrawalLegacyErc20Bridge(&_IL1SharedBridge.TransactOpts, _l2BatchNumber, _l2MessageIndex, _l2TxNumberInBatch, _message, _merkleProof)
}

// FinalizeWithdrawalLegacyErc20Bridge is a paid mutator transaction binding the contract method 0x7ab08472.
//
// Solidity: function finalizeWithdrawalLegacyErc20Bridge(uint256 _l2BatchNumber, uint256 _l2MessageIndex, uint16 _l2TxNumberInBatch, bytes _message, bytes32[] _merkleProof) returns(address l1Receiver, address l1Token, uint256 amount)
func (_IL1SharedBridge *IL1SharedBridgeTransactorSession) FinalizeWithdrawalLegacyErc20Bridge(_l2BatchNumber *big.Int, _l2MessageIndex *big.Int, _l2TxNumberInBatch uint16, _message []byte, _merkleProof [][32]byte) (*types.Transaction, error) {
	return _IL1SharedBridge.Contract.FinalizeWithdrawalLegacyErc20Bridge(&_IL1SharedBridge.TransactOpts, _l2BatchNumber, _l2MessageIndex, _l2TxNumberInBatch, _message, _merkleProof)
}

// ReceiveEth is a paid mutator transaction binding the contract method 0xc2aaf9c4.
//
// Solidity: function receiveEth(uint256 _chainId) payable returns()
func (_IL1SharedBridge *IL1SharedBridgeTransactor) ReceiveEth(opts *bind.TransactOpts, _chainId *big.Int) (*types.Transaction, error) {
	return _IL1SharedBridge.contract.Transact(opts, "receiveEth", _chainId)
}

// ReceiveEth is a paid mutator transaction binding the contract method 0xc2aaf9c4.
//
// Solidity: function receiveEth(uint256 _chainId) payable returns()
func (_IL1SharedBridge *IL1SharedBridgeSession) ReceiveEth(_chainId *big.Int) (*types.Transaction, error) {
	return _IL1SharedBridge.Contract.ReceiveEth(&_IL1SharedBridge.TransactOpts, _chainId)
}

// ReceiveEth is a paid mutator transaction binding the contract method 0xc2aaf9c4.
//
// Solidity: function receiveEth(uint256 _chainId) payable returns()
func (_IL1SharedBridge *IL1SharedBridgeTransactorSession) ReceiveEth(_chainId *big.Int) (*types.Transaction, error) {
	return _IL1SharedBridge.Contract.ReceiveEth(&_IL1SharedBridge.TransactOpts, _chainId)
}

// SetEraFirstPostUpgradeBatch is a paid mutator transaction binding the contract method 0x22d1c19b.
//
// Solidity: function setEraFirstPostUpgradeBatch(uint256 _eraFirstPostUpgradeBatch) returns()
func (_IL1SharedBridge *IL1SharedBridgeTransactor) SetEraFirstPostUpgradeBatch(opts *bind.TransactOpts, _eraFirstPostUpgradeBatch *big.Int) (*types.Transaction, error) {
	return _IL1SharedBridge.contract.Transact(opts, "setEraFirstPostUpgradeBatch", _eraFirstPostUpgradeBatch)
}

// SetEraFirstPostUpgradeBatch is a paid mutator transaction binding the contract method 0x22d1c19b.
//
// Solidity: function setEraFirstPostUpgradeBatch(uint256 _eraFirstPostUpgradeBatch) returns()
func (_IL1SharedBridge *IL1SharedBridgeSession) SetEraFirstPostUpgradeBatch(_eraFirstPostUpgradeBatch *big.Int) (*types.Transaction, error) {
	return _IL1SharedBridge.Contract.SetEraFirstPostUpgradeBatch(&_IL1SharedBridge.TransactOpts, _eraFirstPostUpgradeBatch)
}

// SetEraFirstPostUpgradeBatch is a paid mutator transaction binding the contract method 0x22d1c19b.
//
// Solidity: function setEraFirstPostUpgradeBatch(uint256 _eraFirstPostUpgradeBatch) returns()
func (_IL1SharedBridge *IL1SharedBridgeTransactorSession) SetEraFirstPostUpgradeBatch(_eraFirstPostUpgradeBatch *big.Int) (*types.Transaction, error) {
	return _IL1SharedBridge.Contract.SetEraFirstPostUpgradeBatch(&_IL1SharedBridge.TransactOpts, _eraFirstPostUpgradeBatch)
}

// IL1SharedBridgeBridgehubDepositBaseTokenInitiatedIterator is returned from FilterBridgehubDepositBaseTokenInitiated and is used to iterate over the raw logs and unpacked data for BridgehubDepositBaseTokenInitiated events raised by the IL1SharedBridge contract.
type IL1SharedBridgeBridgehubDepositBaseTokenInitiatedIterator struct {
	Event *IL1SharedBridgeBridgehubDepositBaseTokenInitiated // Event containing the contract specifics and raw log

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
func (it *IL1SharedBridgeBridgehubDepositBaseTokenInitiatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IL1SharedBridgeBridgehubDepositBaseTokenInitiated)
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
		it.Event = new(IL1SharedBridgeBridgehubDepositBaseTokenInitiated)
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
func (it *IL1SharedBridgeBridgehubDepositBaseTokenInitiatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IL1SharedBridgeBridgehubDepositBaseTokenInitiatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IL1SharedBridgeBridgehubDepositBaseTokenInitiated represents a BridgehubDepositBaseTokenInitiated event raised by the IL1SharedBridge contract.
type IL1SharedBridgeBridgehubDepositBaseTokenInitiated struct {
	ChainId *big.Int
	From    common.Address
	L1Token common.Address
	Amount  *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterBridgehubDepositBaseTokenInitiated is a free log retrieval operation binding the contract event 0x249bc8a55d0c4a0034b9aaa6be739bec2d4466e5d859bec9566a8553c405c838.
//
// Solidity: event BridgehubDepositBaseTokenInitiated(uint256 indexed chainId, address indexed from, address l1Token, uint256 amount)
func (_IL1SharedBridge *IL1SharedBridgeFilterer) FilterBridgehubDepositBaseTokenInitiated(opts *bind.FilterOpts, chainId []*big.Int, from []common.Address) (*IL1SharedBridgeBridgehubDepositBaseTokenInitiatedIterator, error) {

	var chainIdRule []interface{}
	for _, chainIdItem := range chainId {
		chainIdRule = append(chainIdRule, chainIdItem)
	}
	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}

	logs, sub, err := _IL1SharedBridge.contract.FilterLogs(opts, "BridgehubDepositBaseTokenInitiated", chainIdRule, fromRule)
	if err != nil {
		return nil, err
	}
	return &IL1SharedBridgeBridgehubDepositBaseTokenInitiatedIterator{contract: _IL1SharedBridge.contract, event: "BridgehubDepositBaseTokenInitiated", logs: logs, sub: sub}, nil
}

// WatchBridgehubDepositBaseTokenInitiated is a free log subscription operation binding the contract event 0x249bc8a55d0c4a0034b9aaa6be739bec2d4466e5d859bec9566a8553c405c838.
//
// Solidity: event BridgehubDepositBaseTokenInitiated(uint256 indexed chainId, address indexed from, address l1Token, uint256 amount)
func (_IL1SharedBridge *IL1SharedBridgeFilterer) WatchBridgehubDepositBaseTokenInitiated(opts *bind.WatchOpts, sink chan<- *IL1SharedBridgeBridgehubDepositBaseTokenInitiated, chainId []*big.Int, from []common.Address) (event.Subscription, error) {

	var chainIdRule []interface{}
	for _, chainIdItem := range chainId {
		chainIdRule = append(chainIdRule, chainIdItem)
	}
	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}

	logs, sub, err := _IL1SharedBridge.contract.WatchLogs(opts, "BridgehubDepositBaseTokenInitiated", chainIdRule, fromRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IL1SharedBridgeBridgehubDepositBaseTokenInitiated)
				if err := _IL1SharedBridge.contract.UnpackLog(event, "BridgehubDepositBaseTokenInitiated", log); err != nil {
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

// ParseBridgehubDepositBaseTokenInitiated is a log parse operation binding the contract event 0x249bc8a55d0c4a0034b9aaa6be739bec2d4466e5d859bec9566a8553c405c838.
//
// Solidity: event BridgehubDepositBaseTokenInitiated(uint256 indexed chainId, address indexed from, address l1Token, uint256 amount)
func (_IL1SharedBridge *IL1SharedBridgeFilterer) ParseBridgehubDepositBaseTokenInitiated(log types.Log) (*IL1SharedBridgeBridgehubDepositBaseTokenInitiated, error) {
	event := new(IL1SharedBridgeBridgehubDepositBaseTokenInitiated)
	if err := _IL1SharedBridge.contract.UnpackLog(event, "BridgehubDepositBaseTokenInitiated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// IL1SharedBridgeBridgehubDepositFinalizedIterator is returned from FilterBridgehubDepositFinalized and is used to iterate over the raw logs and unpacked data for BridgehubDepositFinalized events raised by the IL1SharedBridge contract.
type IL1SharedBridgeBridgehubDepositFinalizedIterator struct {
	Event *IL1SharedBridgeBridgehubDepositFinalized // Event containing the contract specifics and raw log

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
func (it *IL1SharedBridgeBridgehubDepositFinalizedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IL1SharedBridgeBridgehubDepositFinalized)
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
		it.Event = new(IL1SharedBridgeBridgehubDepositFinalized)
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
func (it *IL1SharedBridgeBridgehubDepositFinalizedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IL1SharedBridgeBridgehubDepositFinalizedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IL1SharedBridgeBridgehubDepositFinalized represents a BridgehubDepositFinalized event raised by the IL1SharedBridge contract.
type IL1SharedBridgeBridgehubDepositFinalized struct {
	ChainId         *big.Int
	TxDataHash      [32]byte
	L2DepositTxHash [32]byte
	Raw             types.Log // Blockchain specific contextual infos
}

// FilterBridgehubDepositFinalized is a free log retrieval operation binding the contract event 0xe4def01b981193a97a9e81230d7b9f31812ceaf23f864a828a82c687911cb2df.
//
// Solidity: event BridgehubDepositFinalized(uint256 indexed chainId, bytes32 indexed txDataHash, bytes32 indexed l2DepositTxHash)
func (_IL1SharedBridge *IL1SharedBridgeFilterer) FilterBridgehubDepositFinalized(opts *bind.FilterOpts, chainId []*big.Int, txDataHash [][32]byte, l2DepositTxHash [][32]byte) (*IL1SharedBridgeBridgehubDepositFinalizedIterator, error) {

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

	logs, sub, err := _IL1SharedBridge.contract.FilterLogs(opts, "BridgehubDepositFinalized", chainIdRule, txDataHashRule, l2DepositTxHashRule)
	if err != nil {
		return nil, err
	}
	return &IL1SharedBridgeBridgehubDepositFinalizedIterator{contract: _IL1SharedBridge.contract, event: "BridgehubDepositFinalized", logs: logs, sub: sub}, nil
}

// WatchBridgehubDepositFinalized is a free log subscription operation binding the contract event 0xe4def01b981193a97a9e81230d7b9f31812ceaf23f864a828a82c687911cb2df.
//
// Solidity: event BridgehubDepositFinalized(uint256 indexed chainId, bytes32 indexed txDataHash, bytes32 indexed l2DepositTxHash)
func (_IL1SharedBridge *IL1SharedBridgeFilterer) WatchBridgehubDepositFinalized(opts *bind.WatchOpts, sink chan<- *IL1SharedBridgeBridgehubDepositFinalized, chainId []*big.Int, txDataHash [][32]byte, l2DepositTxHash [][32]byte) (event.Subscription, error) {

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

	logs, sub, err := _IL1SharedBridge.contract.WatchLogs(opts, "BridgehubDepositFinalized", chainIdRule, txDataHashRule, l2DepositTxHashRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IL1SharedBridgeBridgehubDepositFinalized)
				if err := _IL1SharedBridge.contract.UnpackLog(event, "BridgehubDepositFinalized", log); err != nil {
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
func (_IL1SharedBridge *IL1SharedBridgeFilterer) ParseBridgehubDepositFinalized(log types.Log) (*IL1SharedBridgeBridgehubDepositFinalized, error) {
	event := new(IL1SharedBridgeBridgehubDepositFinalized)
	if err := _IL1SharedBridge.contract.UnpackLog(event, "BridgehubDepositFinalized", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// IL1SharedBridgeBridgehubDepositInitiatedIterator is returned from FilterBridgehubDepositInitiated and is used to iterate over the raw logs and unpacked data for BridgehubDepositInitiated events raised by the IL1SharedBridge contract.
type IL1SharedBridgeBridgehubDepositInitiatedIterator struct {
	Event *IL1SharedBridgeBridgehubDepositInitiated // Event containing the contract specifics and raw log

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
func (it *IL1SharedBridgeBridgehubDepositInitiatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IL1SharedBridgeBridgehubDepositInitiated)
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
		it.Event = new(IL1SharedBridgeBridgehubDepositInitiated)
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
func (it *IL1SharedBridgeBridgehubDepositInitiatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IL1SharedBridgeBridgehubDepositInitiatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IL1SharedBridgeBridgehubDepositInitiated represents a BridgehubDepositInitiated event raised by the IL1SharedBridge contract.
type IL1SharedBridgeBridgehubDepositInitiated struct {
	ChainId    *big.Int
	TxDataHash [32]byte
	From       common.Address
	To         common.Address
	L1Token    common.Address
	Amount     *big.Int
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterBridgehubDepositInitiated is a free log retrieval operation binding the contract event 0x8768405a01370685449c74c293804d6c9cc216d170acdbdba50b33ed4144447f.
//
// Solidity: event BridgehubDepositInitiated(uint256 indexed chainId, bytes32 indexed txDataHash, address indexed from, address to, address l1Token, uint256 amount)
func (_IL1SharedBridge *IL1SharedBridgeFilterer) FilterBridgehubDepositInitiated(opts *bind.FilterOpts, chainId []*big.Int, txDataHash [][32]byte, from []common.Address) (*IL1SharedBridgeBridgehubDepositInitiatedIterator, error) {

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

	logs, sub, err := _IL1SharedBridge.contract.FilterLogs(opts, "BridgehubDepositInitiated", chainIdRule, txDataHashRule, fromRule)
	if err != nil {
		return nil, err
	}
	return &IL1SharedBridgeBridgehubDepositInitiatedIterator{contract: _IL1SharedBridge.contract, event: "BridgehubDepositInitiated", logs: logs, sub: sub}, nil
}

// WatchBridgehubDepositInitiated is a free log subscription operation binding the contract event 0x8768405a01370685449c74c293804d6c9cc216d170acdbdba50b33ed4144447f.
//
// Solidity: event BridgehubDepositInitiated(uint256 indexed chainId, bytes32 indexed txDataHash, address indexed from, address to, address l1Token, uint256 amount)
func (_IL1SharedBridge *IL1SharedBridgeFilterer) WatchBridgehubDepositInitiated(opts *bind.WatchOpts, sink chan<- *IL1SharedBridgeBridgehubDepositInitiated, chainId []*big.Int, txDataHash [][32]byte, from []common.Address) (event.Subscription, error) {

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

	logs, sub, err := _IL1SharedBridge.contract.WatchLogs(opts, "BridgehubDepositInitiated", chainIdRule, txDataHashRule, fromRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IL1SharedBridgeBridgehubDepositInitiated)
				if err := _IL1SharedBridge.contract.UnpackLog(event, "BridgehubDepositInitiated", log); err != nil {
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

// ParseBridgehubDepositInitiated is a log parse operation binding the contract event 0x8768405a01370685449c74c293804d6c9cc216d170acdbdba50b33ed4144447f.
//
// Solidity: event BridgehubDepositInitiated(uint256 indexed chainId, bytes32 indexed txDataHash, address indexed from, address to, address l1Token, uint256 amount)
func (_IL1SharedBridge *IL1SharedBridgeFilterer) ParseBridgehubDepositInitiated(log types.Log) (*IL1SharedBridgeBridgehubDepositInitiated, error) {
	event := new(IL1SharedBridgeBridgehubDepositInitiated)
	if err := _IL1SharedBridge.contract.UnpackLog(event, "BridgehubDepositInitiated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// IL1SharedBridgeClaimedFailedDepositSharedBridgeIterator is returned from FilterClaimedFailedDepositSharedBridge and is used to iterate over the raw logs and unpacked data for ClaimedFailedDepositSharedBridge events raised by the IL1SharedBridge contract.
type IL1SharedBridgeClaimedFailedDepositSharedBridgeIterator struct {
	Event *IL1SharedBridgeClaimedFailedDepositSharedBridge // Event containing the contract specifics and raw log

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
func (it *IL1SharedBridgeClaimedFailedDepositSharedBridgeIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IL1SharedBridgeClaimedFailedDepositSharedBridge)
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
		it.Event = new(IL1SharedBridgeClaimedFailedDepositSharedBridge)
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
func (it *IL1SharedBridgeClaimedFailedDepositSharedBridgeIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IL1SharedBridgeClaimedFailedDepositSharedBridgeIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IL1SharedBridgeClaimedFailedDepositSharedBridge represents a ClaimedFailedDepositSharedBridge event raised by the IL1SharedBridge contract.
type IL1SharedBridgeClaimedFailedDepositSharedBridge struct {
	ChainId *big.Int
	To      common.Address
	L1Token common.Address
	Amount  *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterClaimedFailedDepositSharedBridge is a free log retrieval operation binding the contract event 0x3bd55dc610580f1af96f4071b65e94fe7fedb1ccd1c57e2befd807fb806dd047.
//
// Solidity: event ClaimedFailedDepositSharedBridge(uint256 indexed chainId, address indexed to, address indexed l1Token, uint256 amount)
func (_IL1SharedBridge *IL1SharedBridgeFilterer) FilterClaimedFailedDepositSharedBridge(opts *bind.FilterOpts, chainId []*big.Int, to []common.Address, l1Token []common.Address) (*IL1SharedBridgeClaimedFailedDepositSharedBridgeIterator, error) {

	var chainIdRule []interface{}
	for _, chainIdItem := range chainId {
		chainIdRule = append(chainIdRule, chainIdItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}
	var l1TokenRule []interface{}
	for _, l1TokenItem := range l1Token {
		l1TokenRule = append(l1TokenRule, l1TokenItem)
	}

	logs, sub, err := _IL1SharedBridge.contract.FilterLogs(opts, "ClaimedFailedDepositSharedBridge", chainIdRule, toRule, l1TokenRule)
	if err != nil {
		return nil, err
	}
	return &IL1SharedBridgeClaimedFailedDepositSharedBridgeIterator{contract: _IL1SharedBridge.contract, event: "ClaimedFailedDepositSharedBridge", logs: logs, sub: sub}, nil
}

// WatchClaimedFailedDepositSharedBridge is a free log subscription operation binding the contract event 0x3bd55dc610580f1af96f4071b65e94fe7fedb1ccd1c57e2befd807fb806dd047.
//
// Solidity: event ClaimedFailedDepositSharedBridge(uint256 indexed chainId, address indexed to, address indexed l1Token, uint256 amount)
func (_IL1SharedBridge *IL1SharedBridgeFilterer) WatchClaimedFailedDepositSharedBridge(opts *bind.WatchOpts, sink chan<- *IL1SharedBridgeClaimedFailedDepositSharedBridge, chainId []*big.Int, to []common.Address, l1Token []common.Address) (event.Subscription, error) {

	var chainIdRule []interface{}
	for _, chainIdItem := range chainId {
		chainIdRule = append(chainIdRule, chainIdItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}
	var l1TokenRule []interface{}
	for _, l1TokenItem := range l1Token {
		l1TokenRule = append(l1TokenRule, l1TokenItem)
	}

	logs, sub, err := _IL1SharedBridge.contract.WatchLogs(opts, "ClaimedFailedDepositSharedBridge", chainIdRule, toRule, l1TokenRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IL1SharedBridgeClaimedFailedDepositSharedBridge)
				if err := _IL1SharedBridge.contract.UnpackLog(event, "ClaimedFailedDepositSharedBridge", log); err != nil {
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

// ParseClaimedFailedDepositSharedBridge is a log parse operation binding the contract event 0x3bd55dc610580f1af96f4071b65e94fe7fedb1ccd1c57e2befd807fb806dd047.
//
// Solidity: event ClaimedFailedDepositSharedBridge(uint256 indexed chainId, address indexed to, address indexed l1Token, uint256 amount)
func (_IL1SharedBridge *IL1SharedBridgeFilterer) ParseClaimedFailedDepositSharedBridge(log types.Log) (*IL1SharedBridgeClaimedFailedDepositSharedBridge, error) {
	event := new(IL1SharedBridgeClaimedFailedDepositSharedBridge)
	if err := _IL1SharedBridge.contract.UnpackLog(event, "ClaimedFailedDepositSharedBridge", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// IL1SharedBridgeLegacyDepositInitiatedIterator is returned from FilterLegacyDepositInitiated and is used to iterate over the raw logs and unpacked data for LegacyDepositInitiated events raised by the IL1SharedBridge contract.
type IL1SharedBridgeLegacyDepositInitiatedIterator struct {
	Event *IL1SharedBridgeLegacyDepositInitiated // Event containing the contract specifics and raw log

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
func (it *IL1SharedBridgeLegacyDepositInitiatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IL1SharedBridgeLegacyDepositInitiated)
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
		it.Event = new(IL1SharedBridgeLegacyDepositInitiated)
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
func (it *IL1SharedBridgeLegacyDepositInitiatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IL1SharedBridgeLegacyDepositInitiatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IL1SharedBridgeLegacyDepositInitiated represents a LegacyDepositInitiated event raised by the IL1SharedBridge contract.
type IL1SharedBridgeLegacyDepositInitiated struct {
	ChainId         *big.Int
	L2DepositTxHash [32]byte
	From            common.Address
	To              common.Address
	L1Token         common.Address
	Amount          *big.Int
	Raw             types.Log // Blockchain specific contextual infos
}

// FilterLegacyDepositInitiated is a free log retrieval operation binding the contract event 0xa1846a4248529db592da99da276f761d9f37a84d0f3d4e83819b869759000700.
//
// Solidity: event LegacyDepositInitiated(uint256 indexed chainId, bytes32 indexed l2DepositTxHash, address indexed from, address to, address l1Token, uint256 amount)
func (_IL1SharedBridge *IL1SharedBridgeFilterer) FilterLegacyDepositInitiated(opts *bind.FilterOpts, chainId []*big.Int, l2DepositTxHash [][32]byte, from []common.Address) (*IL1SharedBridgeLegacyDepositInitiatedIterator, error) {

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

	logs, sub, err := _IL1SharedBridge.contract.FilterLogs(opts, "LegacyDepositInitiated", chainIdRule, l2DepositTxHashRule, fromRule)
	if err != nil {
		return nil, err
	}
	return &IL1SharedBridgeLegacyDepositInitiatedIterator{contract: _IL1SharedBridge.contract, event: "LegacyDepositInitiated", logs: logs, sub: sub}, nil
}

// WatchLegacyDepositInitiated is a free log subscription operation binding the contract event 0xa1846a4248529db592da99da276f761d9f37a84d0f3d4e83819b869759000700.
//
// Solidity: event LegacyDepositInitiated(uint256 indexed chainId, bytes32 indexed l2DepositTxHash, address indexed from, address to, address l1Token, uint256 amount)
func (_IL1SharedBridge *IL1SharedBridgeFilterer) WatchLegacyDepositInitiated(opts *bind.WatchOpts, sink chan<- *IL1SharedBridgeLegacyDepositInitiated, chainId []*big.Int, l2DepositTxHash [][32]byte, from []common.Address) (event.Subscription, error) {

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

	logs, sub, err := _IL1SharedBridge.contract.WatchLogs(opts, "LegacyDepositInitiated", chainIdRule, l2DepositTxHashRule, fromRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IL1SharedBridgeLegacyDepositInitiated)
				if err := _IL1SharedBridge.contract.UnpackLog(event, "LegacyDepositInitiated", log); err != nil {
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
// Solidity: event LegacyDepositInitiated(uint256 indexed chainId, bytes32 indexed l2DepositTxHash, address indexed from, address to, address l1Token, uint256 amount)
func (_IL1SharedBridge *IL1SharedBridgeFilterer) ParseLegacyDepositInitiated(log types.Log) (*IL1SharedBridgeLegacyDepositInitiated, error) {
	event := new(IL1SharedBridgeLegacyDepositInitiated)
	if err := _IL1SharedBridge.contract.UnpackLog(event, "LegacyDepositInitiated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// IL1SharedBridgeWithdrawalFinalizedSharedBridgeIterator is returned from FilterWithdrawalFinalizedSharedBridge and is used to iterate over the raw logs and unpacked data for WithdrawalFinalizedSharedBridge events raised by the IL1SharedBridge contract.
type IL1SharedBridgeWithdrawalFinalizedSharedBridgeIterator struct {
	Event *IL1SharedBridgeWithdrawalFinalizedSharedBridge // Event containing the contract specifics and raw log

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
func (it *IL1SharedBridgeWithdrawalFinalizedSharedBridgeIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IL1SharedBridgeWithdrawalFinalizedSharedBridge)
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
		it.Event = new(IL1SharedBridgeWithdrawalFinalizedSharedBridge)
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
func (it *IL1SharedBridgeWithdrawalFinalizedSharedBridgeIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IL1SharedBridgeWithdrawalFinalizedSharedBridgeIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IL1SharedBridgeWithdrawalFinalizedSharedBridge represents a WithdrawalFinalizedSharedBridge event raised by the IL1SharedBridge contract.
type IL1SharedBridgeWithdrawalFinalizedSharedBridge struct {
	ChainId *big.Int
	To      common.Address
	L1Token common.Address
	Amount  *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterWithdrawalFinalizedSharedBridge is a free log retrieval operation binding the contract event 0x05518b128f0a9b11ddddebd5211a7fc2f4a689dab3a3e258d93eb13049983c3e.
//
// Solidity: event WithdrawalFinalizedSharedBridge(uint256 indexed chainId, address indexed to, address indexed l1Token, uint256 amount)
func (_IL1SharedBridge *IL1SharedBridgeFilterer) FilterWithdrawalFinalizedSharedBridge(opts *bind.FilterOpts, chainId []*big.Int, to []common.Address, l1Token []common.Address) (*IL1SharedBridgeWithdrawalFinalizedSharedBridgeIterator, error) {

	var chainIdRule []interface{}
	for _, chainIdItem := range chainId {
		chainIdRule = append(chainIdRule, chainIdItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}
	var l1TokenRule []interface{}
	for _, l1TokenItem := range l1Token {
		l1TokenRule = append(l1TokenRule, l1TokenItem)
	}

	logs, sub, err := _IL1SharedBridge.contract.FilterLogs(opts, "WithdrawalFinalizedSharedBridge", chainIdRule, toRule, l1TokenRule)
	if err != nil {
		return nil, err
	}
	return &IL1SharedBridgeWithdrawalFinalizedSharedBridgeIterator{contract: _IL1SharedBridge.contract, event: "WithdrawalFinalizedSharedBridge", logs: logs, sub: sub}, nil
}

// WatchWithdrawalFinalizedSharedBridge is a free log subscription operation binding the contract event 0x05518b128f0a9b11ddddebd5211a7fc2f4a689dab3a3e258d93eb13049983c3e.
//
// Solidity: event WithdrawalFinalizedSharedBridge(uint256 indexed chainId, address indexed to, address indexed l1Token, uint256 amount)
func (_IL1SharedBridge *IL1SharedBridgeFilterer) WatchWithdrawalFinalizedSharedBridge(opts *bind.WatchOpts, sink chan<- *IL1SharedBridgeWithdrawalFinalizedSharedBridge, chainId []*big.Int, to []common.Address, l1Token []common.Address) (event.Subscription, error) {

	var chainIdRule []interface{}
	for _, chainIdItem := range chainId {
		chainIdRule = append(chainIdRule, chainIdItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}
	var l1TokenRule []interface{}
	for _, l1TokenItem := range l1Token {
		l1TokenRule = append(l1TokenRule, l1TokenItem)
	}

	logs, sub, err := _IL1SharedBridge.contract.WatchLogs(opts, "WithdrawalFinalizedSharedBridge", chainIdRule, toRule, l1TokenRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IL1SharedBridgeWithdrawalFinalizedSharedBridge)
				if err := _IL1SharedBridge.contract.UnpackLog(event, "WithdrawalFinalizedSharedBridge", log); err != nil {
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

// ParseWithdrawalFinalizedSharedBridge is a log parse operation binding the contract event 0x05518b128f0a9b11ddddebd5211a7fc2f4a689dab3a3e258d93eb13049983c3e.
//
// Solidity: event WithdrawalFinalizedSharedBridge(uint256 indexed chainId, address indexed to, address indexed l1Token, uint256 amount)
func (_IL1SharedBridge *IL1SharedBridgeFilterer) ParseWithdrawalFinalizedSharedBridge(log types.Log) (*IL1SharedBridgeWithdrawalFinalizedSharedBridge, error) {
	event := new(IL1SharedBridgeWithdrawalFinalizedSharedBridge)
	if err := _IL1SharedBridge.contract.UnpackLog(event, "WithdrawalFinalizedSharedBridge", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
