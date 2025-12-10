// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package ivalidationhook

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

// IValidationHookMetaData contains all meta data concerning the IValidationHook contract.
var IValidationHookMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"onInstall\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"onUninstall\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes4\",\"name\":\"interfaceId\",\"type\":\"bytes4\"}],\"name\":\"supportsInterface\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"signedHash\",\"type\":\"bytes32\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"txType\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"from\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"to\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"gasLimit\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"gasPerPubdataByteLimit\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"maxFeePerGas\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"maxPriorityFeePerGas\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"paymaster\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"nonce\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"},{\"internalType\":\"uint256[4]\",\"name\":\"reserved\",\"type\":\"uint256[4]\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"signature\",\"type\":\"bytes\"},{\"internalType\":\"bytes32[]\",\"name\":\"factoryDeps\",\"type\":\"bytes32[]\"},{\"internalType\":\"bytes\",\"name\":\"paymasterInput\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"reservedDynamic\",\"type\":\"bytes\"}],\"internalType\":\"structTransaction\",\"name\":\"transaction\",\"type\":\"tuple\"}],\"name\":\"validationHook\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
}

// IValidationHookABI is the input ABI used to generate the binding from.
// Deprecated: Use IValidationHookMetaData.ABI instead.
var IValidationHookABI = IValidationHookMetaData.ABI

// IValidationHook is an auto generated Go binding around an Ethereum contract.
type IValidationHook struct {
	IValidationHookCaller     // Read-only binding to the contract
	IValidationHookTransactor // Write-only binding to the contract
	IValidationHookFilterer   // Log filterer for contract events
}

// IValidationHookCaller is an auto generated read-only Go binding around an Ethereum contract.
type IValidationHookCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IValidationHookTransactor is an auto generated write-only Go binding around an Ethereum contract.
type IValidationHookTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IValidationHookFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type IValidationHookFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IValidationHookSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type IValidationHookSession struct {
	Contract     *IValidationHook  // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// IValidationHookCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type IValidationHookCallerSession struct {
	Contract *IValidationHookCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts          // Call options to use throughout this session
}

// IValidationHookTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type IValidationHookTransactorSession struct {
	Contract     *IValidationHookTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts          // Transaction auth options to use throughout this session
}

// IValidationHookRaw is an auto generated low-level Go binding around an Ethereum contract.
type IValidationHookRaw struct {
	Contract *IValidationHook // Generic contract binding to access the raw methods on
}

// IValidationHookCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type IValidationHookCallerRaw struct {
	Contract *IValidationHookCaller // Generic read-only contract binding to access the raw methods on
}

// IValidationHookTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type IValidationHookTransactorRaw struct {
	Contract *IValidationHookTransactor // Generic write-only contract binding to access the raw methods on
}

// NewIValidationHook creates a new instance of IValidationHook, bound to a specific deployed contract.
func NewIValidationHook(address common.Address, backend bind.ContractBackend) (*IValidationHook, error) {
	contract, err := bindIValidationHook(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &IValidationHook{IValidationHookCaller: IValidationHookCaller{contract: contract}, IValidationHookTransactor: IValidationHookTransactor{contract: contract}, IValidationHookFilterer: IValidationHookFilterer{contract: contract}}, nil
}

// NewIValidationHookCaller creates a new read-only instance of IValidationHook, bound to a specific deployed contract.
func NewIValidationHookCaller(address common.Address, caller bind.ContractCaller) (*IValidationHookCaller, error) {
	contract, err := bindIValidationHook(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &IValidationHookCaller{contract: contract}, nil
}

// NewIValidationHookTransactor creates a new write-only instance of IValidationHook, bound to a specific deployed contract.
func NewIValidationHookTransactor(address common.Address, transactor bind.ContractTransactor) (*IValidationHookTransactor, error) {
	contract, err := bindIValidationHook(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &IValidationHookTransactor{contract: contract}, nil
}

// NewIValidationHookFilterer creates a new log filterer instance of IValidationHook, bound to a specific deployed contract.
func NewIValidationHookFilterer(address common.Address, filterer bind.ContractFilterer) (*IValidationHookFilterer, error) {
	contract, err := bindIValidationHook(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &IValidationHookFilterer{contract: contract}, nil
}

// bindIValidationHook binds a generic wrapper to an already deployed contract.
func bindIValidationHook(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := IValidationHookMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IValidationHook *IValidationHookRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IValidationHook.Contract.IValidationHookCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IValidationHook *IValidationHookRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IValidationHook.Contract.IValidationHookTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IValidationHook *IValidationHookRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IValidationHook.Contract.IValidationHookTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IValidationHook *IValidationHookCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IValidationHook.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IValidationHook *IValidationHookTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IValidationHook.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IValidationHook *IValidationHookTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IValidationHook.Contract.contract.Transact(opts, method, params...)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_IValidationHook *IValidationHookCaller) SupportsInterface(opts *bind.CallOpts, interfaceId [4]byte) (bool, error) {
	var out []interface{}
	err := _IValidationHook.contract.Call(opts, &out, "supportsInterface", interfaceId)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_IValidationHook *IValidationHookSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _IValidationHook.Contract.SupportsInterface(&_IValidationHook.CallOpts, interfaceId)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_IValidationHook *IValidationHookCallerSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _IValidationHook.Contract.SupportsInterface(&_IValidationHook.CallOpts, interfaceId)
}

// OnInstall is a paid mutator transaction binding the contract method 0x6d61fe70.
//
// Solidity: function onInstall(bytes data) returns()
func (_IValidationHook *IValidationHookTransactor) OnInstall(opts *bind.TransactOpts, data []byte) (*types.Transaction, error) {
	return _IValidationHook.contract.Transact(opts, "onInstall", data)
}

// OnInstall is a paid mutator transaction binding the contract method 0x6d61fe70.
//
// Solidity: function onInstall(bytes data) returns()
func (_IValidationHook *IValidationHookSession) OnInstall(data []byte) (*types.Transaction, error) {
	return _IValidationHook.Contract.OnInstall(&_IValidationHook.TransactOpts, data)
}

// OnInstall is a paid mutator transaction binding the contract method 0x6d61fe70.
//
// Solidity: function onInstall(bytes data) returns()
func (_IValidationHook *IValidationHookTransactorSession) OnInstall(data []byte) (*types.Transaction, error) {
	return _IValidationHook.Contract.OnInstall(&_IValidationHook.TransactOpts, data)
}

// OnUninstall is a paid mutator transaction binding the contract method 0x8a91b0e3.
//
// Solidity: function onUninstall(bytes data) returns()
func (_IValidationHook *IValidationHookTransactor) OnUninstall(opts *bind.TransactOpts, data []byte) (*types.Transaction, error) {
	return _IValidationHook.contract.Transact(opts, "onUninstall", data)
}

// OnUninstall is a paid mutator transaction binding the contract method 0x8a91b0e3.
//
// Solidity: function onUninstall(bytes data) returns()
func (_IValidationHook *IValidationHookSession) OnUninstall(data []byte) (*types.Transaction, error) {
	return _IValidationHook.Contract.OnUninstall(&_IValidationHook.TransactOpts, data)
}

// OnUninstall is a paid mutator transaction binding the contract method 0x8a91b0e3.
//
// Solidity: function onUninstall(bytes data) returns()
func (_IValidationHook *IValidationHookTransactorSession) OnUninstall(data []byte) (*types.Transaction, error) {
	return _IValidationHook.Contract.OnUninstall(&_IValidationHook.TransactOpts, data)
}

// ValidationHook is a paid mutator transaction binding the contract method 0x37d5f03a.
//
// Solidity: function validationHook(bytes32 signedHash, (uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256[4],bytes,bytes,bytes32[],bytes,bytes) transaction) returns()
func (_IValidationHook *IValidationHookTransactor) ValidationHook(opts *bind.TransactOpts, signedHash [32]byte, transaction Transaction) (*types.Transaction, error) {
	return _IValidationHook.contract.Transact(opts, "validationHook", signedHash, transaction)
}

// ValidationHook is a paid mutator transaction binding the contract method 0x37d5f03a.
//
// Solidity: function validationHook(bytes32 signedHash, (uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256[4],bytes,bytes,bytes32[],bytes,bytes) transaction) returns()
func (_IValidationHook *IValidationHookSession) ValidationHook(signedHash [32]byte, transaction Transaction) (*types.Transaction, error) {
	return _IValidationHook.Contract.ValidationHook(&_IValidationHook.TransactOpts, signedHash, transaction)
}

// ValidationHook is a paid mutator transaction binding the contract method 0x37d5f03a.
//
// Solidity: function validationHook(bytes32 signedHash, (uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256[4],bytes,bytes,bytes32[],bytes,bytes) transaction) returns()
func (_IValidationHook *IValidationHookTransactorSession) ValidationHook(signedHash [32]byte, transaction Transaction) (*types.Transaction, error) {
	return _IValidationHook.Contract.ValidationHook(&_IValidationHook.TransactOpts, signedHash, transaction)
}
