// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package selfauth

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

// SelfAuthMetaData contains all meta data concerning the SelfAuth contract.
var SelfAuthMetaData = &bind.MetaData{
	ABI: "[]",
}

// SelfAuthABI is the input ABI used to generate the binding from.
// Deprecated: Use SelfAuthMetaData.ABI instead.
var SelfAuthABI = SelfAuthMetaData.ABI

// SelfAuth is an auto generated Go binding around an Ethereum contract.
type SelfAuth struct {
	SelfAuthCaller     // Read-only binding to the contract
	SelfAuthTransactor // Write-only binding to the contract
	SelfAuthFilterer   // Log filterer for contract events
}

// SelfAuthCaller is an auto generated read-only Go binding around an Ethereum contract.
type SelfAuthCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SelfAuthTransactor is an auto generated write-only Go binding around an Ethereum contract.
type SelfAuthTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SelfAuthFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type SelfAuthFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SelfAuthSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type SelfAuthSession struct {
	Contract     *SelfAuth         // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// SelfAuthCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type SelfAuthCallerSession struct {
	Contract *SelfAuthCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts   // Call options to use throughout this session
}

// SelfAuthTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type SelfAuthTransactorSession struct {
	Contract     *SelfAuthTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts   // Transaction auth options to use throughout this session
}

// SelfAuthRaw is an auto generated low-level Go binding around an Ethereum contract.
type SelfAuthRaw struct {
	Contract *SelfAuth // Generic contract binding to access the raw methods on
}

// SelfAuthCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type SelfAuthCallerRaw struct {
	Contract *SelfAuthCaller // Generic read-only contract binding to access the raw methods on
}

// SelfAuthTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type SelfAuthTransactorRaw struct {
	Contract *SelfAuthTransactor // Generic write-only contract binding to access the raw methods on
}

// NewSelfAuth creates a new instance of SelfAuth, bound to a specific deployed contract.
func NewSelfAuth(address common.Address, backend bind.ContractBackend) (*SelfAuth, error) {
	contract, err := bindSelfAuth(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &SelfAuth{SelfAuthCaller: SelfAuthCaller{contract: contract}, SelfAuthTransactor: SelfAuthTransactor{contract: contract}, SelfAuthFilterer: SelfAuthFilterer{contract: contract}}, nil
}

// NewSelfAuthCaller creates a new read-only instance of SelfAuth, bound to a specific deployed contract.
func NewSelfAuthCaller(address common.Address, caller bind.ContractCaller) (*SelfAuthCaller, error) {
	contract, err := bindSelfAuth(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &SelfAuthCaller{contract: contract}, nil
}

// NewSelfAuthTransactor creates a new write-only instance of SelfAuth, bound to a specific deployed contract.
func NewSelfAuthTransactor(address common.Address, transactor bind.ContractTransactor) (*SelfAuthTransactor, error) {
	contract, err := bindSelfAuth(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &SelfAuthTransactor{contract: contract}, nil
}

// NewSelfAuthFilterer creates a new log filterer instance of SelfAuth, bound to a specific deployed contract.
func NewSelfAuthFilterer(address common.Address, filterer bind.ContractFilterer) (*SelfAuthFilterer, error) {
	contract, err := bindSelfAuth(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &SelfAuthFilterer{contract: contract}, nil
}

// bindSelfAuth binds a generic wrapper to an already deployed contract.
func bindSelfAuth(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := SelfAuthMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_SelfAuth *SelfAuthRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _SelfAuth.Contract.SelfAuthCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_SelfAuth *SelfAuthRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _SelfAuth.Contract.SelfAuthTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_SelfAuth *SelfAuthRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _SelfAuth.Contract.SelfAuthTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_SelfAuth *SelfAuthCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _SelfAuth.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_SelfAuth *SelfAuthTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _SelfAuth.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_SelfAuth *SelfAuthTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _SelfAuth.Contract.contract.Transact(opts, method, params...)
}
