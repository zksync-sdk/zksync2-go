// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package l1erc20bridge

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

// IL1ERC20BridgeMetaData contains all meta data concerning the IL1ERC20Bridge contract.
var IL1ERC20BridgeMetaData = &bind.MetaData{
	ABI: "[{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"l1Token\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"ClaimedFailedDeposit\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"l2DepositTxHash\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"l1Token\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"DepositInitiated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"l1Token\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"WithdrawalFinalized\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_depositSender\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_l1Token\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"_l2TxHash\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"_l2BatchNumber\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_l2MessageIndex\",\"type\":\"uint256\"},{\"internalType\":\"uint16\",\"name\":\"_l2TxNumberInBatch\",\"type\":\"uint16\"},{\"internalType\":\"bytes32[]\",\"name\":\"_merkleProof\",\"type\":\"bytes32[]\"}],\"name\":\"claimFailedDeposit\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_l2Receiver\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_l1Token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_l2TxGasLimit\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_l2TxGasPerPubdataByte\",\"type\":\"uint256\"}],\"name\":\"deposit\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"txHash\",\"type\":\"bytes32\"}],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_l2Receiver\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_l1Token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_l2TxGasLimit\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_l2TxGasPerPubdataByte\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"_refundRecipient\",\"type\":\"address\"}],\"name\":\"deposit\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"txHash\",\"type\":\"bytes32\"}],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_account\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_l1Token\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"_depositL2TxHash\",\"type\":\"bytes32\"}],\"name\":\"depositAmount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_l2BatchNumber\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_l2MessageIndex\",\"type\":\"uint256\"},{\"internalType\":\"uint16\",\"name\":\"_l2TxNumberInBatch\",\"type\":\"uint16\"},{\"internalType\":\"bytes\",\"name\":\"_message\",\"type\":\"bytes\"},{\"internalType\":\"bytes32[]\",\"name\":\"_merkleProof\",\"type\":\"bytes32[]\"}],\"name\":\"finalizeWithdrawal\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_l2BatchNumber\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_l2MessageIndex\",\"type\":\"uint256\"}],\"name\":\"isWithdrawalFinalized\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"l2Bridge\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_l1Token\",\"type\":\"address\"}],\"name\":\"l2TokenAddress\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"l2TokenBeacon\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"sharedBridge\",\"outputs\":[{\"internalType\":\"contractIL1SharedBridge\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"transferTokenToSharedBridge\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
}

// IL1ERC20BridgeABI is the input ABI used to generate the binding from.
// Deprecated: Use IL1ERC20BridgeMetaData.ABI instead.
var IL1ERC20BridgeABI = IL1ERC20BridgeMetaData.ABI

// IL1ERC20Bridge is an auto generated Go binding around an Ethereum contract.
type IL1ERC20Bridge struct {
	IL1ERC20BridgeCaller     // Read-only binding to the contract
	IL1ERC20BridgeTransactor // Write-only binding to the contract
	IL1ERC20BridgeFilterer   // Log filterer for contract events
}

// IL1ERC20BridgeCaller is an auto generated read-only Go binding around an Ethereum contract.
type IL1ERC20BridgeCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IL1ERC20BridgeTransactor is an auto generated write-only Go binding around an Ethereum contract.
type IL1ERC20BridgeTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IL1ERC20BridgeFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type IL1ERC20BridgeFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IL1ERC20BridgeSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type IL1ERC20BridgeSession struct {
	Contract     *IL1ERC20Bridge   // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// IL1ERC20BridgeCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type IL1ERC20BridgeCallerSession struct {
	Contract *IL1ERC20BridgeCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts         // Call options to use throughout this session
}

// IL1ERC20BridgeTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type IL1ERC20BridgeTransactorSession struct {
	Contract     *IL1ERC20BridgeTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts         // Transaction auth options to use throughout this session
}

// IL1ERC20BridgeRaw is an auto generated low-level Go binding around an Ethereum contract.
type IL1ERC20BridgeRaw struct {
	Contract *IL1ERC20Bridge // Generic contract binding to access the raw methods on
}

// IL1ERC20BridgeCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type IL1ERC20BridgeCallerRaw struct {
	Contract *IL1ERC20BridgeCaller // Generic read-only contract binding to access the raw methods on
}

// IL1ERC20BridgeTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type IL1ERC20BridgeTransactorRaw struct {
	Contract *IL1ERC20BridgeTransactor // Generic write-only contract binding to access the raw methods on
}

// NewIL1ERC20Bridge creates a new instance of IL1ERC20Bridge, bound to a specific deployed contract.
func NewIL1ERC20Bridge(address common.Address, backend bind.ContractBackend) (*IL1ERC20Bridge, error) {
	contract, err := bindIL1ERC20Bridge(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &IL1ERC20Bridge{IL1ERC20BridgeCaller: IL1ERC20BridgeCaller{contract: contract}, IL1ERC20BridgeTransactor: IL1ERC20BridgeTransactor{contract: contract}, IL1ERC20BridgeFilterer: IL1ERC20BridgeFilterer{contract: contract}}, nil
}

// NewIL1ERC20BridgeCaller creates a new read-only instance of IL1ERC20Bridge, bound to a specific deployed contract.
func NewIL1ERC20BridgeCaller(address common.Address, caller bind.ContractCaller) (*IL1ERC20BridgeCaller, error) {
	contract, err := bindIL1ERC20Bridge(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &IL1ERC20BridgeCaller{contract: contract}, nil
}

// NewIL1ERC20BridgeTransactor creates a new write-only instance of IL1ERC20Bridge, bound to a specific deployed contract.
func NewIL1ERC20BridgeTransactor(address common.Address, transactor bind.ContractTransactor) (*IL1ERC20BridgeTransactor, error) {
	contract, err := bindIL1ERC20Bridge(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &IL1ERC20BridgeTransactor{contract: contract}, nil
}

// NewIL1ERC20BridgeFilterer creates a new log filterer instance of IL1ERC20Bridge, bound to a specific deployed contract.
func NewIL1ERC20BridgeFilterer(address common.Address, filterer bind.ContractFilterer) (*IL1ERC20BridgeFilterer, error) {
	contract, err := bindIL1ERC20Bridge(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &IL1ERC20BridgeFilterer{contract: contract}, nil
}

// bindIL1ERC20Bridge binds a generic wrapper to an already deployed contract.
func bindIL1ERC20Bridge(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := IL1ERC20BridgeMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IL1ERC20Bridge *IL1ERC20BridgeRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IL1ERC20Bridge.Contract.IL1ERC20BridgeCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IL1ERC20Bridge *IL1ERC20BridgeRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IL1ERC20Bridge.Contract.IL1ERC20BridgeTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IL1ERC20Bridge *IL1ERC20BridgeRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IL1ERC20Bridge.Contract.IL1ERC20BridgeTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IL1ERC20Bridge *IL1ERC20BridgeCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IL1ERC20Bridge.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IL1ERC20Bridge *IL1ERC20BridgeTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IL1ERC20Bridge.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IL1ERC20Bridge *IL1ERC20BridgeTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IL1ERC20Bridge.Contract.contract.Transact(opts, method, params...)
}

// IsWithdrawalFinalized is a free data retrieval call binding the contract method 0x4bed8212.
//
// Solidity: function isWithdrawalFinalized(uint256 _l2BatchNumber, uint256 _l2MessageIndex) view returns(bool)
func (_IL1ERC20Bridge *IL1ERC20BridgeCaller) IsWithdrawalFinalized(opts *bind.CallOpts, _l2BatchNumber *big.Int, _l2MessageIndex *big.Int) (bool, error) {
	var out []interface{}
	err := _IL1ERC20Bridge.contract.Call(opts, &out, "isWithdrawalFinalized", _l2BatchNumber, _l2MessageIndex)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsWithdrawalFinalized is a free data retrieval call binding the contract method 0x4bed8212.
//
// Solidity: function isWithdrawalFinalized(uint256 _l2BatchNumber, uint256 _l2MessageIndex) view returns(bool)
func (_IL1ERC20Bridge *IL1ERC20BridgeSession) IsWithdrawalFinalized(_l2BatchNumber *big.Int, _l2MessageIndex *big.Int) (bool, error) {
	return _IL1ERC20Bridge.Contract.IsWithdrawalFinalized(&_IL1ERC20Bridge.CallOpts, _l2BatchNumber, _l2MessageIndex)
}

// IsWithdrawalFinalized is a free data retrieval call binding the contract method 0x4bed8212.
//
// Solidity: function isWithdrawalFinalized(uint256 _l2BatchNumber, uint256 _l2MessageIndex) view returns(bool)
func (_IL1ERC20Bridge *IL1ERC20BridgeCallerSession) IsWithdrawalFinalized(_l2BatchNumber *big.Int, _l2MessageIndex *big.Int) (bool, error) {
	return _IL1ERC20Bridge.Contract.IsWithdrawalFinalized(&_IL1ERC20Bridge.CallOpts, _l2BatchNumber, _l2MessageIndex)
}

// L2Bridge is a free data retrieval call binding the contract method 0xae1f6aaf.
//
// Solidity: function l2Bridge() view returns(address)
func (_IL1ERC20Bridge *IL1ERC20BridgeCaller) L2Bridge(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _IL1ERC20Bridge.contract.Call(opts, &out, "l2Bridge")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// L2Bridge is a free data retrieval call binding the contract method 0xae1f6aaf.
//
// Solidity: function l2Bridge() view returns(address)
func (_IL1ERC20Bridge *IL1ERC20BridgeSession) L2Bridge() (common.Address, error) {
	return _IL1ERC20Bridge.Contract.L2Bridge(&_IL1ERC20Bridge.CallOpts)
}

// L2Bridge is a free data retrieval call binding the contract method 0xae1f6aaf.
//
// Solidity: function l2Bridge() view returns(address)
func (_IL1ERC20Bridge *IL1ERC20BridgeCallerSession) L2Bridge() (common.Address, error) {
	return _IL1ERC20Bridge.Contract.L2Bridge(&_IL1ERC20Bridge.CallOpts)
}

// L2TokenAddress is a free data retrieval call binding the contract method 0xf5f15168.
//
// Solidity: function l2TokenAddress(address _l1Token) view returns(address)
func (_IL1ERC20Bridge *IL1ERC20BridgeCaller) L2TokenAddress(opts *bind.CallOpts, _l1Token common.Address) (common.Address, error) {
	var out []interface{}
	err := _IL1ERC20Bridge.contract.Call(opts, &out, "l2TokenAddress", _l1Token)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// L2TokenAddress is a free data retrieval call binding the contract method 0xf5f15168.
//
// Solidity: function l2TokenAddress(address _l1Token) view returns(address)
func (_IL1ERC20Bridge *IL1ERC20BridgeSession) L2TokenAddress(_l1Token common.Address) (common.Address, error) {
	return _IL1ERC20Bridge.Contract.L2TokenAddress(&_IL1ERC20Bridge.CallOpts, _l1Token)
}

// L2TokenAddress is a free data retrieval call binding the contract method 0xf5f15168.
//
// Solidity: function l2TokenAddress(address _l1Token) view returns(address)
func (_IL1ERC20Bridge *IL1ERC20BridgeCallerSession) L2TokenAddress(_l1Token common.Address) (common.Address, error) {
	return _IL1ERC20Bridge.Contract.L2TokenAddress(&_IL1ERC20Bridge.CallOpts, _l1Token)
}

// L2TokenBeacon is a free data retrieval call binding the contract method 0x6dde7209.
//
// Solidity: function l2TokenBeacon() view returns(address)
func (_IL1ERC20Bridge *IL1ERC20BridgeCaller) L2TokenBeacon(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _IL1ERC20Bridge.contract.Call(opts, &out, "l2TokenBeacon")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// L2TokenBeacon is a free data retrieval call binding the contract method 0x6dde7209.
//
// Solidity: function l2TokenBeacon() view returns(address)
func (_IL1ERC20Bridge *IL1ERC20BridgeSession) L2TokenBeacon() (common.Address, error) {
	return _IL1ERC20Bridge.Contract.L2TokenBeacon(&_IL1ERC20Bridge.CallOpts)
}

// L2TokenBeacon is a free data retrieval call binding the contract method 0x6dde7209.
//
// Solidity: function l2TokenBeacon() view returns(address)
func (_IL1ERC20Bridge *IL1ERC20BridgeCallerSession) L2TokenBeacon() (common.Address, error) {
	return _IL1ERC20Bridge.Contract.L2TokenBeacon(&_IL1ERC20Bridge.CallOpts)
}

// SharedBridge is a free data retrieval call binding the contract method 0x38720778.
//
// Solidity: function sharedBridge() view returns(address)
func (_IL1ERC20Bridge *IL1ERC20BridgeCaller) SharedBridge(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _IL1ERC20Bridge.contract.Call(opts, &out, "sharedBridge")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// SharedBridge is a free data retrieval call binding the contract method 0x38720778.
//
// Solidity: function sharedBridge() view returns(address)
func (_IL1ERC20Bridge *IL1ERC20BridgeSession) SharedBridge() (common.Address, error) {
	return _IL1ERC20Bridge.Contract.SharedBridge(&_IL1ERC20Bridge.CallOpts)
}

// SharedBridge is a free data retrieval call binding the contract method 0x38720778.
//
// Solidity: function sharedBridge() view returns(address)
func (_IL1ERC20Bridge *IL1ERC20BridgeCallerSession) SharedBridge() (common.Address, error) {
	return _IL1ERC20Bridge.Contract.SharedBridge(&_IL1ERC20Bridge.CallOpts)
}

// ClaimFailedDeposit is a paid mutator transaction binding the contract method 0x19fa7f62.
//
// Solidity: function claimFailedDeposit(address _depositSender, address _l1Token, bytes32 _l2TxHash, uint256 _l2BatchNumber, uint256 _l2MessageIndex, uint16 _l2TxNumberInBatch, bytes32[] _merkleProof) returns()
func (_IL1ERC20Bridge *IL1ERC20BridgeTransactor) ClaimFailedDeposit(opts *bind.TransactOpts, _depositSender common.Address, _l1Token common.Address, _l2TxHash [32]byte, _l2BatchNumber *big.Int, _l2MessageIndex *big.Int, _l2TxNumberInBatch uint16, _merkleProof [][32]byte) (*types.Transaction, error) {
	return _IL1ERC20Bridge.contract.Transact(opts, "claimFailedDeposit", _depositSender, _l1Token, _l2TxHash, _l2BatchNumber, _l2MessageIndex, _l2TxNumberInBatch, _merkleProof)
}

// ClaimFailedDeposit is a paid mutator transaction binding the contract method 0x19fa7f62.
//
// Solidity: function claimFailedDeposit(address _depositSender, address _l1Token, bytes32 _l2TxHash, uint256 _l2BatchNumber, uint256 _l2MessageIndex, uint16 _l2TxNumberInBatch, bytes32[] _merkleProof) returns()
func (_IL1ERC20Bridge *IL1ERC20BridgeSession) ClaimFailedDeposit(_depositSender common.Address, _l1Token common.Address, _l2TxHash [32]byte, _l2BatchNumber *big.Int, _l2MessageIndex *big.Int, _l2TxNumberInBatch uint16, _merkleProof [][32]byte) (*types.Transaction, error) {
	return _IL1ERC20Bridge.Contract.ClaimFailedDeposit(&_IL1ERC20Bridge.TransactOpts, _depositSender, _l1Token, _l2TxHash, _l2BatchNumber, _l2MessageIndex, _l2TxNumberInBatch, _merkleProof)
}

// ClaimFailedDeposit is a paid mutator transaction binding the contract method 0x19fa7f62.
//
// Solidity: function claimFailedDeposit(address _depositSender, address _l1Token, bytes32 _l2TxHash, uint256 _l2BatchNumber, uint256 _l2MessageIndex, uint16 _l2TxNumberInBatch, bytes32[] _merkleProof) returns()
func (_IL1ERC20Bridge *IL1ERC20BridgeTransactorSession) ClaimFailedDeposit(_depositSender common.Address, _l1Token common.Address, _l2TxHash [32]byte, _l2BatchNumber *big.Int, _l2MessageIndex *big.Int, _l2TxNumberInBatch uint16, _merkleProof [][32]byte) (*types.Transaction, error) {
	return _IL1ERC20Bridge.Contract.ClaimFailedDeposit(&_IL1ERC20Bridge.TransactOpts, _depositSender, _l1Token, _l2TxHash, _l2BatchNumber, _l2MessageIndex, _l2TxNumberInBatch, _merkleProof)
}

// Deposit is a paid mutator transaction binding the contract method 0x933999fb.
//
// Solidity: function deposit(address _l2Receiver, address _l1Token, uint256 _amount, uint256 _l2TxGasLimit, uint256 _l2TxGasPerPubdataByte) payable returns(bytes32 txHash)
func (_IL1ERC20Bridge *IL1ERC20BridgeTransactor) Deposit(opts *bind.TransactOpts, _l2Receiver common.Address, _l1Token common.Address, _amount *big.Int, _l2TxGasLimit *big.Int, _l2TxGasPerPubdataByte *big.Int) (*types.Transaction, error) {
	return _IL1ERC20Bridge.contract.Transact(opts, "deposit", _l2Receiver, _l1Token, _amount, _l2TxGasLimit, _l2TxGasPerPubdataByte)
}

// Deposit is a paid mutator transaction binding the contract method 0x933999fb.
//
// Solidity: function deposit(address _l2Receiver, address _l1Token, uint256 _amount, uint256 _l2TxGasLimit, uint256 _l2TxGasPerPubdataByte) payable returns(bytes32 txHash)
func (_IL1ERC20Bridge *IL1ERC20BridgeSession) Deposit(_l2Receiver common.Address, _l1Token common.Address, _amount *big.Int, _l2TxGasLimit *big.Int, _l2TxGasPerPubdataByte *big.Int) (*types.Transaction, error) {
	return _IL1ERC20Bridge.Contract.Deposit(&_IL1ERC20Bridge.TransactOpts, _l2Receiver, _l1Token, _amount, _l2TxGasLimit, _l2TxGasPerPubdataByte)
}

// Deposit is a paid mutator transaction binding the contract method 0x933999fb.
//
// Solidity: function deposit(address _l2Receiver, address _l1Token, uint256 _amount, uint256 _l2TxGasLimit, uint256 _l2TxGasPerPubdataByte) payable returns(bytes32 txHash)
func (_IL1ERC20Bridge *IL1ERC20BridgeTransactorSession) Deposit(_l2Receiver common.Address, _l1Token common.Address, _amount *big.Int, _l2TxGasLimit *big.Int, _l2TxGasPerPubdataByte *big.Int) (*types.Transaction, error) {
	return _IL1ERC20Bridge.Contract.Deposit(&_IL1ERC20Bridge.TransactOpts, _l2Receiver, _l1Token, _amount, _l2TxGasLimit, _l2TxGasPerPubdataByte)
}

// Deposit0 is a paid mutator transaction binding the contract method 0xe8b99b1b.
//
// Solidity: function deposit(address _l2Receiver, address _l1Token, uint256 _amount, uint256 _l2TxGasLimit, uint256 _l2TxGasPerPubdataByte, address _refundRecipient) payable returns(bytes32 txHash)
func (_IL1ERC20Bridge *IL1ERC20BridgeTransactor) Deposit0(opts *bind.TransactOpts, _l2Receiver common.Address, _l1Token common.Address, _amount *big.Int, _l2TxGasLimit *big.Int, _l2TxGasPerPubdataByte *big.Int, _refundRecipient common.Address) (*types.Transaction, error) {
	return _IL1ERC20Bridge.contract.Transact(opts, "deposit0", _l2Receiver, _l1Token, _amount, _l2TxGasLimit, _l2TxGasPerPubdataByte, _refundRecipient)
}

// Deposit0 is a paid mutator transaction binding the contract method 0xe8b99b1b.
//
// Solidity: function deposit(address _l2Receiver, address _l1Token, uint256 _amount, uint256 _l2TxGasLimit, uint256 _l2TxGasPerPubdataByte, address _refundRecipient) payable returns(bytes32 txHash)
func (_IL1ERC20Bridge *IL1ERC20BridgeSession) Deposit0(_l2Receiver common.Address, _l1Token common.Address, _amount *big.Int, _l2TxGasLimit *big.Int, _l2TxGasPerPubdataByte *big.Int, _refundRecipient common.Address) (*types.Transaction, error) {
	return _IL1ERC20Bridge.Contract.Deposit0(&_IL1ERC20Bridge.TransactOpts, _l2Receiver, _l1Token, _amount, _l2TxGasLimit, _l2TxGasPerPubdataByte, _refundRecipient)
}

// Deposit0 is a paid mutator transaction binding the contract method 0xe8b99b1b.
//
// Solidity: function deposit(address _l2Receiver, address _l1Token, uint256 _amount, uint256 _l2TxGasLimit, uint256 _l2TxGasPerPubdataByte, address _refundRecipient) payable returns(bytes32 txHash)
func (_IL1ERC20Bridge *IL1ERC20BridgeTransactorSession) Deposit0(_l2Receiver common.Address, _l1Token common.Address, _amount *big.Int, _l2TxGasLimit *big.Int, _l2TxGasPerPubdataByte *big.Int, _refundRecipient common.Address) (*types.Transaction, error) {
	return _IL1ERC20Bridge.Contract.Deposit0(&_IL1ERC20Bridge.TransactOpts, _l2Receiver, _l1Token, _amount, _l2TxGasLimit, _l2TxGasPerPubdataByte, _refundRecipient)
}

// DepositAmount is a paid mutator transaction binding the contract method 0x01eae183.
//
// Solidity: function depositAmount(address _account, address _l1Token, bytes32 _depositL2TxHash) returns(uint256 amount)
func (_IL1ERC20Bridge *IL1ERC20BridgeTransactor) DepositAmount(opts *bind.TransactOpts, _account common.Address, _l1Token common.Address, _depositL2TxHash [32]byte) (*types.Transaction, error) {
	return _IL1ERC20Bridge.contract.Transact(opts, "depositAmount", _account, _l1Token, _depositL2TxHash)
}

// DepositAmount is a paid mutator transaction binding the contract method 0x01eae183.
//
// Solidity: function depositAmount(address _account, address _l1Token, bytes32 _depositL2TxHash) returns(uint256 amount)
func (_IL1ERC20Bridge *IL1ERC20BridgeSession) DepositAmount(_account common.Address, _l1Token common.Address, _depositL2TxHash [32]byte) (*types.Transaction, error) {
	return _IL1ERC20Bridge.Contract.DepositAmount(&_IL1ERC20Bridge.TransactOpts, _account, _l1Token, _depositL2TxHash)
}

// DepositAmount is a paid mutator transaction binding the contract method 0x01eae183.
//
// Solidity: function depositAmount(address _account, address _l1Token, bytes32 _depositL2TxHash) returns(uint256 amount)
func (_IL1ERC20Bridge *IL1ERC20BridgeTransactorSession) DepositAmount(_account common.Address, _l1Token common.Address, _depositL2TxHash [32]byte) (*types.Transaction, error) {
	return _IL1ERC20Bridge.Contract.DepositAmount(&_IL1ERC20Bridge.TransactOpts, _account, _l1Token, _depositL2TxHash)
}

// FinalizeWithdrawal is a paid mutator transaction binding the contract method 0x11a2ccc1.
//
// Solidity: function finalizeWithdrawal(uint256 _l2BatchNumber, uint256 _l2MessageIndex, uint16 _l2TxNumberInBatch, bytes _message, bytes32[] _merkleProof) returns()
func (_IL1ERC20Bridge *IL1ERC20BridgeTransactor) FinalizeWithdrawal(opts *bind.TransactOpts, _l2BatchNumber *big.Int, _l2MessageIndex *big.Int, _l2TxNumberInBatch uint16, _message []byte, _merkleProof [][32]byte) (*types.Transaction, error) {
	return _IL1ERC20Bridge.contract.Transact(opts, "finalizeWithdrawal", _l2BatchNumber, _l2MessageIndex, _l2TxNumberInBatch, _message, _merkleProof)
}

// FinalizeWithdrawal is a paid mutator transaction binding the contract method 0x11a2ccc1.
//
// Solidity: function finalizeWithdrawal(uint256 _l2BatchNumber, uint256 _l2MessageIndex, uint16 _l2TxNumberInBatch, bytes _message, bytes32[] _merkleProof) returns()
func (_IL1ERC20Bridge *IL1ERC20BridgeSession) FinalizeWithdrawal(_l2BatchNumber *big.Int, _l2MessageIndex *big.Int, _l2TxNumberInBatch uint16, _message []byte, _merkleProof [][32]byte) (*types.Transaction, error) {
	return _IL1ERC20Bridge.Contract.FinalizeWithdrawal(&_IL1ERC20Bridge.TransactOpts, _l2BatchNumber, _l2MessageIndex, _l2TxNumberInBatch, _message, _merkleProof)
}

// FinalizeWithdrawal is a paid mutator transaction binding the contract method 0x11a2ccc1.
//
// Solidity: function finalizeWithdrawal(uint256 _l2BatchNumber, uint256 _l2MessageIndex, uint16 _l2TxNumberInBatch, bytes _message, bytes32[] _merkleProof) returns()
func (_IL1ERC20Bridge *IL1ERC20BridgeTransactorSession) FinalizeWithdrawal(_l2BatchNumber *big.Int, _l2MessageIndex *big.Int, _l2TxNumberInBatch uint16, _message []byte, _merkleProof [][32]byte) (*types.Transaction, error) {
	return _IL1ERC20Bridge.Contract.FinalizeWithdrawal(&_IL1ERC20Bridge.TransactOpts, _l2BatchNumber, _l2MessageIndex, _l2TxNumberInBatch, _message, _merkleProof)
}

// TransferTokenToSharedBridge is a paid mutator transaction binding the contract method 0x03a33104.
//
// Solidity: function transferTokenToSharedBridge(address _token, uint256 _amount) returns()
func (_IL1ERC20Bridge *IL1ERC20BridgeTransactor) TransferTokenToSharedBridge(opts *bind.TransactOpts, _token common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _IL1ERC20Bridge.contract.Transact(opts, "transferTokenToSharedBridge", _token, _amount)
}

// TransferTokenToSharedBridge is a paid mutator transaction binding the contract method 0x03a33104.
//
// Solidity: function transferTokenToSharedBridge(address _token, uint256 _amount) returns()
func (_IL1ERC20Bridge *IL1ERC20BridgeSession) TransferTokenToSharedBridge(_token common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _IL1ERC20Bridge.Contract.TransferTokenToSharedBridge(&_IL1ERC20Bridge.TransactOpts, _token, _amount)
}

// TransferTokenToSharedBridge is a paid mutator transaction binding the contract method 0x03a33104.
//
// Solidity: function transferTokenToSharedBridge(address _token, uint256 _amount) returns()
func (_IL1ERC20Bridge *IL1ERC20BridgeTransactorSession) TransferTokenToSharedBridge(_token common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _IL1ERC20Bridge.Contract.TransferTokenToSharedBridge(&_IL1ERC20Bridge.TransactOpts, _token, _amount)
}

// IL1ERC20BridgeClaimedFailedDepositIterator is returned from FilterClaimedFailedDeposit and is used to iterate over the raw logs and unpacked data for ClaimedFailedDeposit events raised by the IL1ERC20Bridge contract.
type IL1ERC20BridgeClaimedFailedDepositIterator struct {
	Event *IL1ERC20BridgeClaimedFailedDeposit // Event containing the contract specifics and raw log

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
func (it *IL1ERC20BridgeClaimedFailedDepositIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IL1ERC20BridgeClaimedFailedDeposit)
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
		it.Event = new(IL1ERC20BridgeClaimedFailedDeposit)
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
func (it *IL1ERC20BridgeClaimedFailedDepositIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IL1ERC20BridgeClaimedFailedDepositIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IL1ERC20BridgeClaimedFailedDeposit represents a ClaimedFailedDeposit event raised by the IL1ERC20Bridge contract.
type IL1ERC20BridgeClaimedFailedDeposit struct {
	To      common.Address
	L1Token common.Address
	Amount  *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterClaimedFailedDeposit is a free log retrieval operation binding the contract event 0xbe066dc591f4a444f75176d387c3e6c775e5706d9ea9a91d11eb49030c66cf60.
//
// Solidity: event ClaimedFailedDeposit(address indexed to, address indexed l1Token, uint256 amount)
func (_IL1ERC20Bridge *IL1ERC20BridgeFilterer) FilterClaimedFailedDeposit(opts *bind.FilterOpts, to []common.Address, l1Token []common.Address) (*IL1ERC20BridgeClaimedFailedDepositIterator, error) {

	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}
	var l1TokenRule []interface{}
	for _, l1TokenItem := range l1Token {
		l1TokenRule = append(l1TokenRule, l1TokenItem)
	}

	logs, sub, err := _IL1ERC20Bridge.contract.FilterLogs(opts, "ClaimedFailedDeposit", toRule, l1TokenRule)
	if err != nil {
		return nil, err
	}
	return &IL1ERC20BridgeClaimedFailedDepositIterator{contract: _IL1ERC20Bridge.contract, event: "ClaimedFailedDeposit", logs: logs, sub: sub}, nil
}

// WatchClaimedFailedDeposit is a free log subscription operation binding the contract event 0xbe066dc591f4a444f75176d387c3e6c775e5706d9ea9a91d11eb49030c66cf60.
//
// Solidity: event ClaimedFailedDeposit(address indexed to, address indexed l1Token, uint256 amount)
func (_IL1ERC20Bridge *IL1ERC20BridgeFilterer) WatchClaimedFailedDeposit(opts *bind.WatchOpts, sink chan<- *IL1ERC20BridgeClaimedFailedDeposit, to []common.Address, l1Token []common.Address) (event.Subscription, error) {

	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}
	var l1TokenRule []interface{}
	for _, l1TokenItem := range l1Token {
		l1TokenRule = append(l1TokenRule, l1TokenItem)
	}

	logs, sub, err := _IL1ERC20Bridge.contract.WatchLogs(opts, "ClaimedFailedDeposit", toRule, l1TokenRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IL1ERC20BridgeClaimedFailedDeposit)
				if err := _IL1ERC20Bridge.contract.UnpackLog(event, "ClaimedFailedDeposit", log); err != nil {
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

// ParseClaimedFailedDeposit is a log parse operation binding the contract event 0xbe066dc591f4a444f75176d387c3e6c775e5706d9ea9a91d11eb49030c66cf60.
//
// Solidity: event ClaimedFailedDeposit(address indexed to, address indexed l1Token, uint256 amount)
func (_IL1ERC20Bridge *IL1ERC20BridgeFilterer) ParseClaimedFailedDeposit(log types.Log) (*IL1ERC20BridgeClaimedFailedDeposit, error) {
	event := new(IL1ERC20BridgeClaimedFailedDeposit)
	if err := _IL1ERC20Bridge.contract.UnpackLog(event, "ClaimedFailedDeposit", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// IL1ERC20BridgeDepositInitiatedIterator is returned from FilterDepositInitiated and is used to iterate over the raw logs and unpacked data for DepositInitiated events raised by the IL1ERC20Bridge contract.
type IL1ERC20BridgeDepositInitiatedIterator struct {
	Event *IL1ERC20BridgeDepositInitiated // Event containing the contract specifics and raw log

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
func (it *IL1ERC20BridgeDepositInitiatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IL1ERC20BridgeDepositInitiated)
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
		it.Event = new(IL1ERC20BridgeDepositInitiated)
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
func (it *IL1ERC20BridgeDepositInitiatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IL1ERC20BridgeDepositInitiatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IL1ERC20BridgeDepositInitiated represents a DepositInitiated event raised by the IL1ERC20Bridge contract.
type IL1ERC20BridgeDepositInitiated struct {
	L2DepositTxHash [32]byte
	From            common.Address
	To              common.Address
	L1Token         common.Address
	Amount          *big.Int
	Raw             types.Log // Blockchain specific contextual infos
}

// FilterDepositInitiated is a free log retrieval operation binding the contract event 0xdd341179f4edc78148d894d0213a96d212af2cbaf223d19ef6d483bdd47ab81d.
//
// Solidity: event DepositInitiated(bytes32 indexed l2DepositTxHash, address indexed from, address indexed to, address l1Token, uint256 amount)
func (_IL1ERC20Bridge *IL1ERC20BridgeFilterer) FilterDepositInitiated(opts *bind.FilterOpts, l2DepositTxHash [][32]byte, from []common.Address, to []common.Address) (*IL1ERC20BridgeDepositInitiatedIterator, error) {

	var l2DepositTxHashRule []interface{}
	for _, l2DepositTxHashItem := range l2DepositTxHash {
		l2DepositTxHashRule = append(l2DepositTxHashRule, l2DepositTxHashItem)
	}
	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _IL1ERC20Bridge.contract.FilterLogs(opts, "DepositInitiated", l2DepositTxHashRule, fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &IL1ERC20BridgeDepositInitiatedIterator{contract: _IL1ERC20Bridge.contract, event: "DepositInitiated", logs: logs, sub: sub}, nil
}

// WatchDepositInitiated is a free log subscription operation binding the contract event 0xdd341179f4edc78148d894d0213a96d212af2cbaf223d19ef6d483bdd47ab81d.
//
// Solidity: event DepositInitiated(bytes32 indexed l2DepositTxHash, address indexed from, address indexed to, address l1Token, uint256 amount)
func (_IL1ERC20Bridge *IL1ERC20BridgeFilterer) WatchDepositInitiated(opts *bind.WatchOpts, sink chan<- *IL1ERC20BridgeDepositInitiated, l2DepositTxHash [][32]byte, from []common.Address, to []common.Address) (event.Subscription, error) {

	var l2DepositTxHashRule []interface{}
	for _, l2DepositTxHashItem := range l2DepositTxHash {
		l2DepositTxHashRule = append(l2DepositTxHashRule, l2DepositTxHashItem)
	}
	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _IL1ERC20Bridge.contract.WatchLogs(opts, "DepositInitiated", l2DepositTxHashRule, fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IL1ERC20BridgeDepositInitiated)
				if err := _IL1ERC20Bridge.contract.UnpackLog(event, "DepositInitiated", log); err != nil {
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

// ParseDepositInitiated is a log parse operation binding the contract event 0xdd341179f4edc78148d894d0213a96d212af2cbaf223d19ef6d483bdd47ab81d.
//
// Solidity: event DepositInitiated(bytes32 indexed l2DepositTxHash, address indexed from, address indexed to, address l1Token, uint256 amount)
func (_IL1ERC20Bridge *IL1ERC20BridgeFilterer) ParseDepositInitiated(log types.Log) (*IL1ERC20BridgeDepositInitiated, error) {
	event := new(IL1ERC20BridgeDepositInitiated)
	if err := _IL1ERC20Bridge.contract.UnpackLog(event, "DepositInitiated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// IL1ERC20BridgeWithdrawalFinalizedIterator is returned from FilterWithdrawalFinalized and is used to iterate over the raw logs and unpacked data for WithdrawalFinalized events raised by the IL1ERC20Bridge contract.
type IL1ERC20BridgeWithdrawalFinalizedIterator struct {
	Event *IL1ERC20BridgeWithdrawalFinalized // Event containing the contract specifics and raw log

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
func (it *IL1ERC20BridgeWithdrawalFinalizedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IL1ERC20BridgeWithdrawalFinalized)
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
		it.Event = new(IL1ERC20BridgeWithdrawalFinalized)
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
func (it *IL1ERC20BridgeWithdrawalFinalizedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IL1ERC20BridgeWithdrawalFinalizedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IL1ERC20BridgeWithdrawalFinalized represents a WithdrawalFinalized event raised by the IL1ERC20Bridge contract.
type IL1ERC20BridgeWithdrawalFinalized struct {
	To      common.Address
	L1Token common.Address
	Amount  *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterWithdrawalFinalized is a free log retrieval operation binding the contract event 0xac1b18083978656d557d6e91c88203585cfda1031bdb14538327121ef140d383.
//
// Solidity: event WithdrawalFinalized(address indexed to, address indexed l1Token, uint256 amount)
func (_IL1ERC20Bridge *IL1ERC20BridgeFilterer) FilterWithdrawalFinalized(opts *bind.FilterOpts, to []common.Address, l1Token []common.Address) (*IL1ERC20BridgeWithdrawalFinalizedIterator, error) {

	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}
	var l1TokenRule []interface{}
	for _, l1TokenItem := range l1Token {
		l1TokenRule = append(l1TokenRule, l1TokenItem)
	}

	logs, sub, err := _IL1ERC20Bridge.contract.FilterLogs(opts, "WithdrawalFinalized", toRule, l1TokenRule)
	if err != nil {
		return nil, err
	}
	return &IL1ERC20BridgeWithdrawalFinalizedIterator{contract: _IL1ERC20Bridge.contract, event: "WithdrawalFinalized", logs: logs, sub: sub}, nil
}

// WatchWithdrawalFinalized is a free log subscription operation binding the contract event 0xac1b18083978656d557d6e91c88203585cfda1031bdb14538327121ef140d383.
//
// Solidity: event WithdrawalFinalized(address indexed to, address indexed l1Token, uint256 amount)
func (_IL1ERC20Bridge *IL1ERC20BridgeFilterer) WatchWithdrawalFinalized(opts *bind.WatchOpts, sink chan<- *IL1ERC20BridgeWithdrawalFinalized, to []common.Address, l1Token []common.Address) (event.Subscription, error) {

	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}
	var l1TokenRule []interface{}
	for _, l1TokenItem := range l1Token {
		l1TokenRule = append(l1TokenRule, l1TokenItem)
	}

	logs, sub, err := _IL1ERC20Bridge.contract.WatchLogs(opts, "WithdrawalFinalized", toRule, l1TokenRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IL1ERC20BridgeWithdrawalFinalized)
				if err := _IL1ERC20Bridge.contract.UnpackLog(event, "WithdrawalFinalized", log); err != nil {
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

// ParseWithdrawalFinalized is a log parse operation binding the contract event 0xac1b18083978656d557d6e91c88203585cfda1031bdb14538327121ef140d383.
//
// Solidity: event WithdrawalFinalized(address indexed to, address indexed l1Token, uint256 amount)
func (_IL1ERC20Bridge *IL1ERC20BridgeFilterer) ParseWithdrawalFinalized(log types.Log) (*IL1ERC20BridgeWithdrawalFinalized, error) {
	event := new(IL1ERC20BridgeWithdrawalFinalized)
	if err := _IL1ERC20Bridge.contract.UnpackLog(event, "WithdrawalFinalized", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
