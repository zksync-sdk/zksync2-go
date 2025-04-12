// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package iexecutionhook

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

// Transaction is an auto generated low-level Go binding around an user-defined struct.
type Transaction struct {
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
	FactoryDeps            [][32]byte
	PaymasterInput         []byte
	ReservedDynamic        []byte
}

// IExecutionHookMetaData contains all meta data concerning the IExecutionHook contract.
var IExecutionHookMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"onInstall\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"onUninstall\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"context\",\"type\":\"bytes\"}],\"name\":\"postExecutionHook\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"txType\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"from\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"to\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"gasLimit\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"gasPerPubdataByteLimit\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"maxFeePerGas\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"maxPriorityFeePerGas\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"paymaster\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"nonce\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"},{\"internalType\":\"uint256[4]\",\"name\":\"reserved\",\"type\":\"uint256[4]\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"signature\",\"type\":\"bytes\"},{\"internalType\":\"bytes32[]\",\"name\":\"factoryDeps\",\"type\":\"bytes32[]\"},{\"internalType\":\"bytes\",\"name\":\"paymasterInput\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"reservedDynamic\",\"type\":\"bytes\"}],\"internalType\":\"structTransaction\",\"name\":\"transaction\",\"type\":\"tuple\"}],\"name\":\"preExecutionHook\",\"outputs\":[{\"internalType\":\"bytes\",\"name\":\"context\",\"type\":\"bytes\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes4\",\"name\":\"interfaceId\",\"type\":\"bytes4\"}],\"name\":\"supportsInterface\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
}

// IExecutionHookABI is the input ABI used to generate the binding from.
// Deprecated: Use IExecutionHookMetaData.ABI instead.
var IExecutionHookABI = IExecutionHookMetaData.ABI

// IExecutionHook is an auto generated Go binding around an Ethereum contract.
type IExecutionHook struct {
	IExecutionHookCaller     // Read-only binding to the contract
	IExecutionHookTransactor // Write-only binding to the contract
	IExecutionHookFilterer   // Log filterer for contract events
}

// IExecutionHookCaller is an auto generated read-only Go binding around an Ethereum contract.
type IExecutionHookCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IExecutionHookTransactor is an auto generated write-only Go binding around an Ethereum contract.
type IExecutionHookTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IExecutionHookFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type IExecutionHookFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IExecutionHookSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type IExecutionHookSession struct {
	Contract     *IExecutionHook   // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// IExecutionHookCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type IExecutionHookCallerSession struct {
	Contract *IExecutionHookCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts         // Call options to use throughout this session
}

// IExecutionHookTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type IExecutionHookTransactorSession struct {
	Contract     *IExecutionHookTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts         // Transaction auth options to use throughout this session
}

// IExecutionHookRaw is an auto generated low-level Go binding around an Ethereum contract.
type IExecutionHookRaw struct {
	Contract *IExecutionHook // Generic contract binding to access the raw methods on
}

// IExecutionHookCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type IExecutionHookCallerRaw struct {
	Contract *IExecutionHookCaller // Generic read-only contract binding to access the raw methods on
}

// IExecutionHookTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type IExecutionHookTransactorRaw struct {
	Contract *IExecutionHookTransactor // Generic write-only contract binding to access the raw methods on
}

// NewIExecutionHook creates a new instance of IExecutionHook, bound to a specific deployed contract.
func NewIExecutionHook(address common.Address, backend bind.ContractBackend) (*IExecutionHook, error) {
	contract, err := bindIExecutionHook(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &IExecutionHook{IExecutionHookCaller: IExecutionHookCaller{contract: contract}, IExecutionHookTransactor: IExecutionHookTransactor{contract: contract}, IExecutionHookFilterer: IExecutionHookFilterer{contract: contract}}, nil
}

// NewIExecutionHookCaller creates a new read-only instance of IExecutionHook, bound to a specific deployed contract.
func NewIExecutionHookCaller(address common.Address, caller bind.ContractCaller) (*IExecutionHookCaller, error) {
	contract, err := bindIExecutionHook(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &IExecutionHookCaller{contract: contract}, nil
}

// NewIExecutionHookTransactor creates a new write-only instance of IExecutionHook, bound to a specific deployed contract.
func NewIExecutionHookTransactor(address common.Address, transactor bind.ContractTransactor) (*IExecutionHookTransactor, error) {
	contract, err := bindIExecutionHook(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &IExecutionHookTransactor{contract: contract}, nil
}

// NewIExecutionHookFilterer creates a new log filterer instance of IExecutionHook, bound to a specific deployed contract.
func NewIExecutionHookFilterer(address common.Address, filterer bind.ContractFilterer) (*IExecutionHookFilterer, error) {
	contract, err := bindIExecutionHook(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &IExecutionHookFilterer{contract: contract}, nil
}

// bindIExecutionHook binds a generic wrapper to an already deployed contract.
func bindIExecutionHook(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := IExecutionHookMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IExecutionHook *IExecutionHookRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IExecutionHook.Contract.IExecutionHookCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IExecutionHook *IExecutionHookRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IExecutionHook.Contract.IExecutionHookTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IExecutionHook *IExecutionHookRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IExecutionHook.Contract.IExecutionHookTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IExecutionHook *IExecutionHookCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IExecutionHook.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IExecutionHook *IExecutionHookTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IExecutionHook.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IExecutionHook *IExecutionHookTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IExecutionHook.Contract.contract.Transact(opts, method, params...)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_IExecutionHook *IExecutionHookCaller) SupportsInterface(opts *bind.CallOpts, interfaceId [4]byte) (bool, error) {
	var out []interface{}
	err := _IExecutionHook.contract.Call(opts, &out, "supportsInterface", interfaceId)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_IExecutionHook *IExecutionHookSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _IExecutionHook.Contract.SupportsInterface(&_IExecutionHook.CallOpts, interfaceId)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_IExecutionHook *IExecutionHookCallerSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _IExecutionHook.Contract.SupportsInterface(&_IExecutionHook.CallOpts, interfaceId)
}

// OnInstall is a paid mutator transaction binding the contract method 0x6d61fe70.
//
// Solidity: function onInstall(bytes data) returns()
func (_IExecutionHook *IExecutionHookTransactor) OnInstall(opts *bind.TransactOpts, data []byte) (*types.Transaction, error) {
	return _IExecutionHook.contract.Transact(opts, "onInstall", data)
}

// OnInstall is a paid mutator transaction binding the contract method 0x6d61fe70.
//
// Solidity: function onInstall(bytes data) returns()
func (_IExecutionHook *IExecutionHookSession) OnInstall(data []byte) (*types.Transaction, error) {
	return _IExecutionHook.Contract.OnInstall(&_IExecutionHook.TransactOpts, data)
}

// OnInstall is a paid mutator transaction binding the contract method 0x6d61fe70.
//
// Solidity: function onInstall(bytes data) returns()
func (_IExecutionHook *IExecutionHookTransactorSession) OnInstall(data []byte) (*types.Transaction, error) {
	return _IExecutionHook.Contract.OnInstall(&_IExecutionHook.TransactOpts, data)
}

// OnUninstall is a paid mutator transaction binding the contract method 0x8a91b0e3.
//
// Solidity: function onUninstall(bytes data) returns()
func (_IExecutionHook *IExecutionHookTransactor) OnUninstall(opts *bind.TransactOpts, data []byte) (*types.Transaction, error) {
	return _IExecutionHook.contract.Transact(opts, "onUninstall", data)
}

// OnUninstall is a paid mutator transaction binding the contract method 0x8a91b0e3.
//
// Solidity: function onUninstall(bytes data) returns()
func (_IExecutionHook *IExecutionHookSession) OnUninstall(data []byte) (*types.Transaction, error) {
	return _IExecutionHook.Contract.OnUninstall(&_IExecutionHook.TransactOpts, data)
}

// OnUninstall is a paid mutator transaction binding the contract method 0x8a91b0e3.
//
// Solidity: function onUninstall(bytes data) returns()
func (_IExecutionHook *IExecutionHookTransactorSession) OnUninstall(data []byte) (*types.Transaction, error) {
	return _IExecutionHook.Contract.OnUninstall(&_IExecutionHook.TransactOpts, data)
}

// PostExecutionHook is a paid mutator transaction binding the contract method 0x3920e539.
//
// Solidity: function postExecutionHook(bytes context) returns()
func (_IExecutionHook *IExecutionHookTransactor) PostExecutionHook(opts *bind.TransactOpts, context []byte) (*types.Transaction, error) {
	return _IExecutionHook.contract.Transact(opts, "postExecutionHook", context)
}

// PostExecutionHook is a paid mutator transaction binding the contract method 0x3920e539.
//
// Solidity: function postExecutionHook(bytes context) returns()
func (_IExecutionHook *IExecutionHookSession) PostExecutionHook(context []byte) (*types.Transaction, error) {
	return _IExecutionHook.Contract.PostExecutionHook(&_IExecutionHook.TransactOpts, context)
}

// PostExecutionHook is a paid mutator transaction binding the contract method 0x3920e539.
//
// Solidity: function postExecutionHook(bytes context) returns()
func (_IExecutionHook *IExecutionHookTransactorSession) PostExecutionHook(context []byte) (*types.Transaction, error) {
	return _IExecutionHook.Contract.PostExecutionHook(&_IExecutionHook.TransactOpts, context)
}

// PreExecutionHook is a paid mutator transaction binding the contract method 0x926a74a6.
//
// Solidity: function preExecutionHook((uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256[4],bytes,bytes,bytes32[],bytes,bytes) transaction) returns(bytes context)
func (_IExecutionHook *IExecutionHookTransactor) PreExecutionHook(opts *bind.TransactOpts, transaction Transaction) (*types.Transaction, error) {
	return _IExecutionHook.contract.Transact(opts, "preExecutionHook", transaction)
}

// PreExecutionHook is a paid mutator transaction binding the contract method 0x926a74a6.
//
// Solidity: function preExecutionHook((uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256[4],bytes,bytes,bytes32[],bytes,bytes) transaction) returns(bytes context)
func (_IExecutionHook *IExecutionHookSession) PreExecutionHook(transaction Transaction) (*types.Transaction, error) {
	return _IExecutionHook.Contract.PreExecutionHook(&_IExecutionHook.TransactOpts, transaction)
}

// PreExecutionHook is a paid mutator transaction binding the contract method 0x926a74a6.
//
// Solidity: function preExecutionHook((uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256[4],bytes,bytes,bytes32[],bytes,bytes) transaction) returns(bytes context)
func (_IExecutionHook *IExecutionHookTransactorSession) PreExecutionHook(transaction Transaction) (*types.Transaction, error) {
	return _IExecutionHook.Contract.PreExecutionHook(&_IExecutionHook.TransactOpts, transaction)
}
