// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package testneterc20token

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

// ITestnetERC20TokenMetaData contains all meta data concerning the ITestnetERC20Token contract.
var ITestnetERC20TokenMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"name\":\"decimals\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"mint\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
}

// ITestnetERC20TokenABI is the input ABI used to generate the binding from.
// Deprecated: Use ITestnetERC20TokenMetaData.ABI instead.
var ITestnetERC20TokenABI = ITestnetERC20TokenMetaData.ABI

// ITestnetERC20Token is an auto generated Go binding around an Ethereum contract.
type ITestnetERC20Token struct {
	ITestnetERC20TokenCaller     // Read-only binding to the contract
	ITestnetERC20TokenTransactor // Write-only binding to the contract
	ITestnetERC20TokenFilterer   // Log filterer for contract events
}

// ITestnetERC20TokenCaller is an auto generated read-only Go binding around an Ethereum contract.
type ITestnetERC20TokenCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ITestnetERC20TokenTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ITestnetERC20TokenTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ITestnetERC20TokenFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ITestnetERC20TokenFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ITestnetERC20TokenSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ITestnetERC20TokenSession struct {
	Contract     *ITestnetERC20Token // Generic contract binding to set the session for
	CallOpts     bind.CallOpts       // Call options to use throughout this session
	TransactOpts bind.TransactOpts   // Transaction auth options to use throughout this session
}

// ITestnetERC20TokenCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ITestnetERC20TokenCallerSession struct {
	Contract *ITestnetERC20TokenCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts             // Call options to use throughout this session
}

// ITestnetERC20TokenTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ITestnetERC20TokenTransactorSession struct {
	Contract     *ITestnetERC20TokenTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts             // Transaction auth options to use throughout this session
}

// ITestnetERC20TokenRaw is an auto generated low-level Go binding around an Ethereum contract.
type ITestnetERC20TokenRaw struct {
	Contract *ITestnetERC20Token // Generic contract binding to access the raw methods on
}

// ITestnetERC20TokenCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ITestnetERC20TokenCallerRaw struct {
	Contract *ITestnetERC20TokenCaller // Generic read-only contract binding to access the raw methods on
}

// ITestnetERC20TokenTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ITestnetERC20TokenTransactorRaw struct {
	Contract *ITestnetERC20TokenTransactor // Generic write-only contract binding to access the raw methods on
}

// NewITestnetERC20Token creates a new instance of ITestnetERC20Token, bound to a specific deployed contract.
func NewITestnetERC20Token(address common.Address, backend bind.ContractBackend) (*ITestnetERC20Token, error) {
	contract, err := bindITestnetERC20Token(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &ITestnetERC20Token{ITestnetERC20TokenCaller: ITestnetERC20TokenCaller{contract: contract}, ITestnetERC20TokenTransactor: ITestnetERC20TokenTransactor{contract: contract}, ITestnetERC20TokenFilterer: ITestnetERC20TokenFilterer{contract: contract}}, nil
}

// NewITestnetERC20TokenCaller creates a new read-only instance of ITestnetERC20Token, bound to a specific deployed contract.
func NewITestnetERC20TokenCaller(address common.Address, caller bind.ContractCaller) (*ITestnetERC20TokenCaller, error) {
	contract, err := bindITestnetERC20Token(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ITestnetERC20TokenCaller{contract: contract}, nil
}

// NewITestnetERC20TokenTransactor creates a new write-only instance of ITestnetERC20Token, bound to a specific deployed contract.
func NewITestnetERC20TokenTransactor(address common.Address, transactor bind.ContractTransactor) (*ITestnetERC20TokenTransactor, error) {
	contract, err := bindITestnetERC20Token(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ITestnetERC20TokenTransactor{contract: contract}, nil
}

// NewITestnetERC20TokenFilterer creates a new log filterer instance of ITestnetERC20Token, bound to a specific deployed contract.
func NewITestnetERC20TokenFilterer(address common.Address, filterer bind.ContractFilterer) (*ITestnetERC20TokenFilterer, error) {
	contract, err := bindITestnetERC20Token(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ITestnetERC20TokenFilterer{contract: contract}, nil
}

// bindITestnetERC20Token binds a generic wrapper to an already deployed contract.
func bindITestnetERC20Token(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := ITestnetERC20TokenMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ITestnetERC20Token *ITestnetERC20TokenRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ITestnetERC20Token.Contract.ITestnetERC20TokenCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ITestnetERC20Token *ITestnetERC20TokenRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ITestnetERC20Token.Contract.ITestnetERC20TokenTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ITestnetERC20Token *ITestnetERC20TokenRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ITestnetERC20Token.Contract.ITestnetERC20TokenTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ITestnetERC20Token *ITestnetERC20TokenCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ITestnetERC20Token.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ITestnetERC20Token *ITestnetERC20TokenTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ITestnetERC20Token.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ITestnetERC20Token *ITestnetERC20TokenTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ITestnetERC20Token.Contract.contract.Transact(opts, method, params...)
}

// Decimals is a paid mutator transaction binding the contract method 0x313ce567.
//
// Solidity: function decimals() returns(uint8)
func (_ITestnetERC20Token *ITestnetERC20TokenTransactor) Decimals(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ITestnetERC20Token.contract.Transact(opts, "decimals")
}

// Decimals is a paid mutator transaction binding the contract method 0x313ce567.
//
// Solidity: function decimals() returns(uint8)
func (_ITestnetERC20Token *ITestnetERC20TokenSession) Decimals() (*types.Transaction, error) {
	return _ITestnetERC20Token.Contract.Decimals(&_ITestnetERC20Token.TransactOpts)
}

// Decimals is a paid mutator transaction binding the contract method 0x313ce567.
//
// Solidity: function decimals() returns(uint8)
func (_ITestnetERC20Token *ITestnetERC20TokenTransactorSession) Decimals() (*types.Transaction, error) {
	return _ITestnetERC20Token.Contract.Decimals(&_ITestnetERC20Token.TransactOpts)
}

// Mint is a paid mutator transaction binding the contract method 0x40c10f19.
//
// Solidity: function mint(address _to, uint256 _amount) returns(bool)
func (_ITestnetERC20Token *ITestnetERC20TokenTransactor) Mint(opts *bind.TransactOpts, _to common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _ITestnetERC20Token.contract.Transact(opts, "mint", _to, _amount)
}

// Mint is a paid mutator transaction binding the contract method 0x40c10f19.
//
// Solidity: function mint(address _to, uint256 _amount) returns(bool)
func (_ITestnetERC20Token *ITestnetERC20TokenSession) Mint(_to common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _ITestnetERC20Token.Contract.Mint(&_ITestnetERC20Token.TransactOpts, _to, _amount)
}

// Mint is a paid mutator transaction binding the contract method 0x40c10f19.
//
// Solidity: function mint(address _to, uint256 _amount) returns(bool)
func (_ITestnetERC20Token *ITestnetERC20TokenTransactorSession) Mint(_to common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _ITestnetERC20Token.Contract.Mint(&_ITestnetERC20Token.TransactOpts, _to, _amount)
}
