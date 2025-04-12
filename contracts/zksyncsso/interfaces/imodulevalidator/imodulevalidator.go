// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package imodulevalidator

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

// IModuleValidatorMetaData contains all meta data concerning the IModuleValidator contract.
var IModuleValidatorMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"onInstall\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"onUninstall\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes4\",\"name\":\"interfaceId\",\"type\":\"bytes4\"}],\"name\":\"supportsInterface\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"signedHash\",\"type\":\"bytes32\"},{\"internalType\":\"bytes\",\"name\":\"signature\",\"type\":\"bytes\"}],\"name\":\"validateSignature\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"signedHash\",\"type\":\"bytes32\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"txType\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"from\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"to\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"gasLimit\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"gasPerPubdataByteLimit\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"maxFeePerGas\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"maxPriorityFeePerGas\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"paymaster\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"nonce\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"},{\"internalType\":\"uint256[4]\",\"name\":\"reserved\",\"type\":\"uint256[4]\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"signature\",\"type\":\"bytes\"},{\"internalType\":\"bytes32[]\",\"name\":\"factoryDeps\",\"type\":\"bytes32[]\"},{\"internalType\":\"bytes\",\"name\":\"paymasterInput\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"reservedDynamic\",\"type\":\"bytes\"}],\"internalType\":\"structTransaction\",\"name\":\"transaction\",\"type\":\"tuple\"}],\"name\":\"validateTransaction\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
}

// IModuleValidatorABI is the input ABI used to generate the binding from.
// Deprecated: Use IModuleValidatorMetaData.ABI instead.
var IModuleValidatorABI = IModuleValidatorMetaData.ABI

// IModuleValidator is an auto generated Go binding around an Ethereum contract.
type IModuleValidator struct {
	IModuleValidatorCaller     // Read-only binding to the contract
	IModuleValidatorTransactor // Write-only binding to the contract
	IModuleValidatorFilterer   // Log filterer for contract events
}

// IModuleValidatorCaller is an auto generated read-only Go binding around an Ethereum contract.
type IModuleValidatorCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IModuleValidatorTransactor is an auto generated write-only Go binding around an Ethereum contract.
type IModuleValidatorTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IModuleValidatorFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type IModuleValidatorFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IModuleValidatorSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type IModuleValidatorSession struct {
	Contract     *IModuleValidator // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// IModuleValidatorCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type IModuleValidatorCallerSession struct {
	Contract *IModuleValidatorCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts           // Call options to use throughout this session
}

// IModuleValidatorTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type IModuleValidatorTransactorSession struct {
	Contract     *IModuleValidatorTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts           // Transaction auth options to use throughout this session
}

// IModuleValidatorRaw is an auto generated low-level Go binding around an Ethereum contract.
type IModuleValidatorRaw struct {
	Contract *IModuleValidator // Generic contract binding to access the raw methods on
}

// IModuleValidatorCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type IModuleValidatorCallerRaw struct {
	Contract *IModuleValidatorCaller // Generic read-only contract binding to access the raw methods on
}

// IModuleValidatorTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type IModuleValidatorTransactorRaw struct {
	Contract *IModuleValidatorTransactor // Generic write-only contract binding to access the raw methods on
}

// NewIModuleValidator creates a new instance of IModuleValidator, bound to a specific deployed contract.
func NewIModuleValidator(address common.Address, backend bind.ContractBackend) (*IModuleValidator, error) {
	contract, err := bindIModuleValidator(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &IModuleValidator{IModuleValidatorCaller: IModuleValidatorCaller{contract: contract}, IModuleValidatorTransactor: IModuleValidatorTransactor{contract: contract}, IModuleValidatorFilterer: IModuleValidatorFilterer{contract: contract}}, nil
}

// NewIModuleValidatorCaller creates a new read-only instance of IModuleValidator, bound to a specific deployed contract.
func NewIModuleValidatorCaller(address common.Address, caller bind.ContractCaller) (*IModuleValidatorCaller, error) {
	contract, err := bindIModuleValidator(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &IModuleValidatorCaller{contract: contract}, nil
}

// NewIModuleValidatorTransactor creates a new write-only instance of IModuleValidator, bound to a specific deployed contract.
func NewIModuleValidatorTransactor(address common.Address, transactor bind.ContractTransactor) (*IModuleValidatorTransactor, error) {
	contract, err := bindIModuleValidator(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &IModuleValidatorTransactor{contract: contract}, nil
}

// NewIModuleValidatorFilterer creates a new log filterer instance of IModuleValidator, bound to a specific deployed contract.
func NewIModuleValidatorFilterer(address common.Address, filterer bind.ContractFilterer) (*IModuleValidatorFilterer, error) {
	contract, err := bindIModuleValidator(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &IModuleValidatorFilterer{contract: contract}, nil
}

// bindIModuleValidator binds a generic wrapper to an already deployed contract.
func bindIModuleValidator(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := IModuleValidatorMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IModuleValidator *IModuleValidatorRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IModuleValidator.Contract.IModuleValidatorCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IModuleValidator *IModuleValidatorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IModuleValidator.Contract.IModuleValidatorTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IModuleValidator *IModuleValidatorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IModuleValidator.Contract.IModuleValidatorTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IModuleValidator *IModuleValidatorCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IModuleValidator.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IModuleValidator *IModuleValidatorTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IModuleValidator.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IModuleValidator *IModuleValidatorTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IModuleValidator.Contract.contract.Transact(opts, method, params...)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_IModuleValidator *IModuleValidatorCaller) SupportsInterface(opts *bind.CallOpts, interfaceId [4]byte) (bool, error) {
	var out []interface{}
	err := _IModuleValidator.contract.Call(opts, &out, "supportsInterface", interfaceId)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_IModuleValidator *IModuleValidatorSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _IModuleValidator.Contract.SupportsInterface(&_IModuleValidator.CallOpts, interfaceId)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_IModuleValidator *IModuleValidatorCallerSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _IModuleValidator.Contract.SupportsInterface(&_IModuleValidator.CallOpts, interfaceId)
}

// ValidateSignature is a free data retrieval call binding the contract method 0x333daf92.
//
// Solidity: function validateSignature(bytes32 signedHash, bytes signature) view returns(bool)
func (_IModuleValidator *IModuleValidatorCaller) ValidateSignature(opts *bind.CallOpts, signedHash [32]byte, signature []byte) (bool, error) {
	var out []interface{}
	err := _IModuleValidator.contract.Call(opts, &out, "validateSignature", signedHash, signature)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// ValidateSignature is a free data retrieval call binding the contract method 0x333daf92.
//
// Solidity: function validateSignature(bytes32 signedHash, bytes signature) view returns(bool)
func (_IModuleValidator *IModuleValidatorSession) ValidateSignature(signedHash [32]byte, signature []byte) (bool, error) {
	return _IModuleValidator.Contract.ValidateSignature(&_IModuleValidator.CallOpts, signedHash, signature)
}

// ValidateSignature is a free data retrieval call binding the contract method 0x333daf92.
//
// Solidity: function validateSignature(bytes32 signedHash, bytes signature) view returns(bool)
func (_IModuleValidator *IModuleValidatorCallerSession) ValidateSignature(signedHash [32]byte, signature []byte) (bool, error) {
	return _IModuleValidator.Contract.ValidateSignature(&_IModuleValidator.CallOpts, signedHash, signature)
}

// OnInstall is a paid mutator transaction binding the contract method 0x6d61fe70.
//
// Solidity: function onInstall(bytes data) returns()
func (_IModuleValidator *IModuleValidatorTransactor) OnInstall(opts *bind.TransactOpts, data []byte) (*types.Transaction, error) {
	return _IModuleValidator.contract.Transact(opts, "onInstall", data)
}

// OnInstall is a paid mutator transaction binding the contract method 0x6d61fe70.
//
// Solidity: function onInstall(bytes data) returns()
func (_IModuleValidator *IModuleValidatorSession) OnInstall(data []byte) (*types.Transaction, error) {
	return _IModuleValidator.Contract.OnInstall(&_IModuleValidator.TransactOpts, data)
}

// OnInstall is a paid mutator transaction binding the contract method 0x6d61fe70.
//
// Solidity: function onInstall(bytes data) returns()
func (_IModuleValidator *IModuleValidatorTransactorSession) OnInstall(data []byte) (*types.Transaction, error) {
	return _IModuleValidator.Contract.OnInstall(&_IModuleValidator.TransactOpts, data)
}

// OnUninstall is a paid mutator transaction binding the contract method 0x8a91b0e3.
//
// Solidity: function onUninstall(bytes data) returns()
func (_IModuleValidator *IModuleValidatorTransactor) OnUninstall(opts *bind.TransactOpts, data []byte) (*types.Transaction, error) {
	return _IModuleValidator.contract.Transact(opts, "onUninstall", data)
}

// OnUninstall is a paid mutator transaction binding the contract method 0x8a91b0e3.
//
// Solidity: function onUninstall(bytes data) returns()
func (_IModuleValidator *IModuleValidatorSession) OnUninstall(data []byte) (*types.Transaction, error) {
	return _IModuleValidator.Contract.OnUninstall(&_IModuleValidator.TransactOpts, data)
}

// OnUninstall is a paid mutator transaction binding the contract method 0x8a91b0e3.
//
// Solidity: function onUninstall(bytes data) returns()
func (_IModuleValidator *IModuleValidatorTransactorSession) OnUninstall(data []byte) (*types.Transaction, error) {
	return _IModuleValidator.Contract.OnUninstall(&_IModuleValidator.TransactOpts, data)
}

// ValidateTransaction is a paid mutator transaction binding the contract method 0xbf1733f9.
//
// Solidity: function validateTransaction(bytes32 signedHash, (uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256[4],bytes,bytes,bytes32[],bytes,bytes) transaction) returns(bool)
func (_IModuleValidator *IModuleValidatorTransactor) ValidateTransaction(opts *bind.TransactOpts, signedHash [32]byte, transaction Transaction) (*types.Transaction, error) {
	return _IModuleValidator.contract.Transact(opts, "validateTransaction", signedHash, transaction)
}

// ValidateTransaction is a paid mutator transaction binding the contract method 0xbf1733f9.
//
// Solidity: function validateTransaction(bytes32 signedHash, (uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256[4],bytes,bytes,bytes32[],bytes,bytes) transaction) returns(bool)
func (_IModuleValidator *IModuleValidatorSession) ValidateTransaction(signedHash [32]byte, transaction Transaction) (*types.Transaction, error) {
	return _IModuleValidator.Contract.ValidateTransaction(&_IModuleValidator.TransactOpts, signedHash, transaction)
}

// ValidateTransaction is a paid mutator transaction binding the contract method 0xbf1733f9.
//
// Solidity: function validateTransaction(bytes32 signedHash, (uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256[4],bytes,bytes,bytes32[],bytes,bytes) transaction) returns(bool)
func (_IModuleValidator *IModuleValidatorTransactorSession) ValidateTransaction(signedHash [32]byte, transaction Transaction) (*types.Transaction, error) {
	return _IModuleValidator.Contract.ValidateTransaction(&_IModuleValidator.TransactOpts, signedHash, transaction)
}
