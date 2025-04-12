// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package jsonparserlibtest

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

// JSONParserLibItem is an auto generated low-level Go binding around an user-defined struct.
type JSONParserLibItem struct {
	Data *big.Int
}

// JSONParserLibTestMetaData contains all meta data concerning the JSONParserLibTest contract.
var JSONParserLibTestMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"string\",\"name\":\"json\",\"type\":\"string\"}],\"name\":\"parse\",\"outputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"_data\",\"type\":\"uint256\"}],\"internalType\":\"structJSONParserLib.Item\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"pure\",\"type\":\"function\"}]",
}

// JSONParserLibTestABI is the input ABI used to generate the binding from.
// Deprecated: Use JSONParserLibTestMetaData.ABI instead.
var JSONParserLibTestABI = JSONParserLibTestMetaData.ABI

// JSONParserLibTest is an auto generated Go binding around an Ethereum contract.
type JSONParserLibTest struct {
	JSONParserLibTestCaller     // Read-only binding to the contract
	JSONParserLibTestTransactor // Write-only binding to the contract
	JSONParserLibTestFilterer   // Log filterer for contract events
}

// JSONParserLibTestCaller is an auto generated read-only Go binding around an Ethereum contract.
type JSONParserLibTestCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// JSONParserLibTestTransactor is an auto generated write-only Go binding around an Ethereum contract.
type JSONParserLibTestTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// JSONParserLibTestFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type JSONParserLibTestFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// JSONParserLibTestSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type JSONParserLibTestSession struct {
	Contract     *JSONParserLibTest // Generic contract binding to set the session for
	CallOpts     bind.CallOpts      // Call options to use throughout this session
	TransactOpts bind.TransactOpts  // Transaction auth options to use throughout this session
}

// JSONParserLibTestCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type JSONParserLibTestCallerSession struct {
	Contract *JSONParserLibTestCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts            // Call options to use throughout this session
}

// JSONParserLibTestTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type JSONParserLibTestTransactorSession struct {
	Contract     *JSONParserLibTestTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts            // Transaction auth options to use throughout this session
}

// JSONParserLibTestRaw is an auto generated low-level Go binding around an Ethereum contract.
type JSONParserLibTestRaw struct {
	Contract *JSONParserLibTest // Generic contract binding to access the raw methods on
}

// JSONParserLibTestCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type JSONParserLibTestCallerRaw struct {
	Contract *JSONParserLibTestCaller // Generic read-only contract binding to access the raw methods on
}

// JSONParserLibTestTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type JSONParserLibTestTransactorRaw struct {
	Contract *JSONParserLibTestTransactor // Generic write-only contract binding to access the raw methods on
}

// NewJSONParserLibTest creates a new instance of JSONParserLibTest, bound to a specific deployed contract.
func NewJSONParserLibTest(address common.Address, backend bind.ContractBackend) (*JSONParserLibTest, error) {
	contract, err := bindJSONParserLibTest(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &JSONParserLibTest{JSONParserLibTestCaller: JSONParserLibTestCaller{contract: contract}, JSONParserLibTestTransactor: JSONParserLibTestTransactor{contract: contract}, JSONParserLibTestFilterer: JSONParserLibTestFilterer{contract: contract}}, nil
}

// NewJSONParserLibTestCaller creates a new read-only instance of JSONParserLibTest, bound to a specific deployed contract.
func NewJSONParserLibTestCaller(address common.Address, caller bind.ContractCaller) (*JSONParserLibTestCaller, error) {
	contract, err := bindJSONParserLibTest(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &JSONParserLibTestCaller{contract: contract}, nil
}

// NewJSONParserLibTestTransactor creates a new write-only instance of JSONParserLibTest, bound to a specific deployed contract.
func NewJSONParserLibTestTransactor(address common.Address, transactor bind.ContractTransactor) (*JSONParserLibTestTransactor, error) {
	contract, err := bindJSONParserLibTest(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &JSONParserLibTestTransactor{contract: contract}, nil
}

// NewJSONParserLibTestFilterer creates a new log filterer instance of JSONParserLibTest, bound to a specific deployed contract.
func NewJSONParserLibTestFilterer(address common.Address, filterer bind.ContractFilterer) (*JSONParserLibTestFilterer, error) {
	contract, err := bindJSONParserLibTest(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &JSONParserLibTestFilterer{contract: contract}, nil
}

// bindJSONParserLibTest binds a generic wrapper to an already deployed contract.
func bindJSONParserLibTest(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := JSONParserLibTestMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_JSONParserLibTest *JSONParserLibTestRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _JSONParserLibTest.Contract.JSONParserLibTestCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_JSONParserLibTest *JSONParserLibTestRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _JSONParserLibTest.Contract.JSONParserLibTestTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_JSONParserLibTest *JSONParserLibTestRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _JSONParserLibTest.Contract.JSONParserLibTestTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_JSONParserLibTest *JSONParserLibTestCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _JSONParserLibTest.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_JSONParserLibTest *JSONParserLibTestTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _JSONParserLibTest.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_JSONParserLibTest *JSONParserLibTestTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _JSONParserLibTest.Contract.contract.Transact(opts, method, params...)
}

// Parse is a free data retrieval call binding the contract method 0xbc62d8d8.
//
// Solidity: function parse(string json) pure returns((uint256))
func (_JSONParserLibTest *JSONParserLibTestCaller) Parse(opts *bind.CallOpts, json string) (JSONParserLibItem, error) {
	var out []interface{}
	err := _JSONParserLibTest.contract.Call(opts, &out, "parse", json)

	if err != nil {
		return *new(JSONParserLibItem), err
	}

	out0 := *abi.ConvertType(out[0], new(JSONParserLibItem)).(*JSONParserLibItem)

	return out0, err

}

// Parse is a free data retrieval call binding the contract method 0xbc62d8d8.
//
// Solidity: function parse(string json) pure returns((uint256))
func (_JSONParserLibTest *JSONParserLibTestSession) Parse(json string) (JSONParserLibItem, error) {
	return _JSONParserLibTest.Contract.Parse(&_JSONParserLibTest.CallOpts, json)
}

// Parse is a free data retrieval call binding the contract method 0xbc62d8d8.
//
// Solidity: function parse(string json) pure returns((uint256))
func (_JSONParserLibTest *JSONParserLibTestCallerSession) Parse(json string) (JSONParserLibItem, error) {
	return _JSONParserLibTest.Contract.Parse(&_JSONParserLibTest.CallOpts, json)
}
