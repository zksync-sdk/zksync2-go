// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package testpaymaster

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

// TestPaymasterMetaData contains all meta data concerning the TestPaymaster contract.
var TestPaymasterMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"_context\",\"type\":\"bytes\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"txType\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"from\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"to\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"gasLimit\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"gasPerPubdataByteLimit\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"maxFeePerGas\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"maxPriorityFeePerGas\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"paymaster\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"nonce\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"},{\"internalType\":\"uint256[4]\",\"name\":\"reserved\",\"type\":\"uint256[4]\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"signature\",\"type\":\"bytes\"},{\"internalType\":\"bytes32[]\",\"name\":\"factoryDeps\",\"type\":\"bytes32[]\"},{\"internalType\":\"bytes\",\"name\":\"paymasterInput\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"reservedDynamic\",\"type\":\"bytes\"}],\"internalType\":\"structTransaction\",\"name\":\"transaction\",\"type\":\"tuple\"},{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"},{\"internalType\":\"enumExecutionResult\",\"name\":\"_txResult\",\"type\":\"uint8\"},{\"internalType\":\"uint256\",\"name\":\"_maxRefundedGas\",\"type\":\"uint256\"}],\"name\":\"postTransaction\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"txType\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"from\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"to\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"gasLimit\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"gasPerPubdataByteLimit\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"maxFeePerGas\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"maxPriorityFeePerGas\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"paymaster\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"nonce\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"},{\"internalType\":\"uint256[4]\",\"name\":\"reserved\",\"type\":\"uint256[4]\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"signature\",\"type\":\"bytes\"},{\"internalType\":\"bytes32[]\",\"name\":\"factoryDeps\",\"type\":\"bytes32[]\"},{\"internalType\":\"bytes\",\"name\":\"paymasterInput\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"reservedDynamic\",\"type\":\"bytes\"}],\"internalType\":\"structTransaction\",\"name\":\"transaction\",\"type\":\"tuple\"}],\"name\":\"validateAndPayForPaymasterTransaction\",\"outputs\":[{\"internalType\":\"bytes4\",\"name\":\"magic\",\"type\":\"bytes4\"},{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"stateMutability\":\"payable\",\"type\":\"receive\"}]",
}

// TestPaymasterABI is the input ABI used to generate the binding from.
// Deprecated: Use TestPaymasterMetaData.ABI instead.
var TestPaymasterABI = TestPaymasterMetaData.ABI

// TestPaymaster is an auto generated Go binding around an Ethereum contract.
type TestPaymaster struct {
	TestPaymasterCaller     // Read-only binding to the contract
	TestPaymasterTransactor // Write-only binding to the contract
	TestPaymasterFilterer   // Log filterer for contract events
}

// TestPaymasterCaller is an auto generated read-only Go binding around an Ethereum contract.
type TestPaymasterCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TestPaymasterTransactor is an auto generated write-only Go binding around an Ethereum contract.
type TestPaymasterTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TestPaymasterFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type TestPaymasterFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TestPaymasterSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type TestPaymasterSession struct {
	Contract     *TestPaymaster    // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// TestPaymasterCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type TestPaymasterCallerSession struct {
	Contract *TestPaymasterCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts        // Call options to use throughout this session
}

// TestPaymasterTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type TestPaymasterTransactorSession struct {
	Contract     *TestPaymasterTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts        // Transaction auth options to use throughout this session
}

// TestPaymasterRaw is an auto generated low-level Go binding around an Ethereum contract.
type TestPaymasterRaw struct {
	Contract *TestPaymaster // Generic contract binding to access the raw methods on
}

// TestPaymasterCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type TestPaymasterCallerRaw struct {
	Contract *TestPaymasterCaller // Generic read-only contract binding to access the raw methods on
}

// TestPaymasterTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type TestPaymasterTransactorRaw struct {
	Contract *TestPaymasterTransactor // Generic write-only contract binding to access the raw methods on
}

// NewTestPaymaster creates a new instance of TestPaymaster, bound to a specific deployed contract.
func NewTestPaymaster(address common.Address, backend bind.ContractBackend) (*TestPaymaster, error) {
	contract, err := bindTestPaymaster(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &TestPaymaster{TestPaymasterCaller: TestPaymasterCaller{contract: contract}, TestPaymasterTransactor: TestPaymasterTransactor{contract: contract}, TestPaymasterFilterer: TestPaymasterFilterer{contract: contract}}, nil
}

// NewTestPaymasterCaller creates a new read-only instance of TestPaymaster, bound to a specific deployed contract.
func NewTestPaymasterCaller(address common.Address, caller bind.ContractCaller) (*TestPaymasterCaller, error) {
	contract, err := bindTestPaymaster(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &TestPaymasterCaller{contract: contract}, nil
}

// NewTestPaymasterTransactor creates a new write-only instance of TestPaymaster, bound to a specific deployed contract.
func NewTestPaymasterTransactor(address common.Address, transactor bind.ContractTransactor) (*TestPaymasterTransactor, error) {
	contract, err := bindTestPaymaster(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &TestPaymasterTransactor{contract: contract}, nil
}

// NewTestPaymasterFilterer creates a new log filterer instance of TestPaymaster, bound to a specific deployed contract.
func NewTestPaymasterFilterer(address common.Address, filterer bind.ContractFilterer) (*TestPaymasterFilterer, error) {
	contract, err := bindTestPaymaster(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &TestPaymasterFilterer{contract: contract}, nil
}

// bindTestPaymaster binds a generic wrapper to an already deployed contract.
func bindTestPaymaster(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := TestPaymasterMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_TestPaymaster *TestPaymasterRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _TestPaymaster.Contract.TestPaymasterCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_TestPaymaster *TestPaymasterRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _TestPaymaster.Contract.TestPaymasterTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_TestPaymaster *TestPaymasterRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _TestPaymaster.Contract.TestPaymasterTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_TestPaymaster *TestPaymasterCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _TestPaymaster.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_TestPaymaster *TestPaymasterTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _TestPaymaster.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_TestPaymaster *TestPaymasterTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _TestPaymaster.Contract.contract.Transact(opts, method, params...)
}

// PostTransaction is a paid mutator transaction binding the contract method 0x817b17f0.
//
// Solidity: function postTransaction(bytes _context, (uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256[4],bytes,bytes,bytes32[],bytes,bytes) transaction, bytes32 , bytes32 , uint8 _txResult, uint256 _maxRefundedGas) payable returns()
func (_TestPaymaster *TestPaymasterTransactor) PostTransaction(opts *bind.TransactOpts, _context []byte, transaction Transaction, arg2 [32]byte, arg3 [32]byte, _txResult uint8, _maxRefundedGas *big.Int) (*types.Transaction, error) {
	return _TestPaymaster.contract.Transact(opts, "postTransaction", _context, transaction, arg2, arg3, _txResult, _maxRefundedGas)
}

// PostTransaction is a paid mutator transaction binding the contract method 0x817b17f0.
//
// Solidity: function postTransaction(bytes _context, (uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256[4],bytes,bytes,bytes32[],bytes,bytes) transaction, bytes32 , bytes32 , uint8 _txResult, uint256 _maxRefundedGas) payable returns()
func (_TestPaymaster *TestPaymasterSession) PostTransaction(_context []byte, transaction Transaction, arg2 [32]byte, arg3 [32]byte, _txResult uint8, _maxRefundedGas *big.Int) (*types.Transaction, error) {
	return _TestPaymaster.Contract.PostTransaction(&_TestPaymaster.TransactOpts, _context, transaction, arg2, arg3, _txResult, _maxRefundedGas)
}

// PostTransaction is a paid mutator transaction binding the contract method 0x817b17f0.
//
// Solidity: function postTransaction(bytes _context, (uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256[4],bytes,bytes,bytes32[],bytes,bytes) transaction, bytes32 , bytes32 , uint8 _txResult, uint256 _maxRefundedGas) payable returns()
func (_TestPaymaster *TestPaymasterTransactorSession) PostTransaction(_context []byte, transaction Transaction, arg2 [32]byte, arg3 [32]byte, _txResult uint8, _maxRefundedGas *big.Int) (*types.Transaction, error) {
	return _TestPaymaster.Contract.PostTransaction(&_TestPaymaster.TransactOpts, _context, transaction, arg2, arg3, _txResult, _maxRefundedGas)
}

// ValidateAndPayForPaymasterTransaction is a paid mutator transaction binding the contract method 0x038a24bc.
//
// Solidity: function validateAndPayForPaymasterTransaction(bytes32 , bytes32 , (uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256[4],bytes,bytes,bytes32[],bytes,bytes) transaction) payable returns(bytes4 magic, bytes)
func (_TestPaymaster *TestPaymasterTransactor) ValidateAndPayForPaymasterTransaction(opts *bind.TransactOpts, arg0 [32]byte, arg1 [32]byte, transaction Transaction) (*types.Transaction, error) {
	return _TestPaymaster.contract.Transact(opts, "validateAndPayForPaymasterTransaction", arg0, arg1, transaction)
}

// ValidateAndPayForPaymasterTransaction is a paid mutator transaction binding the contract method 0x038a24bc.
//
// Solidity: function validateAndPayForPaymasterTransaction(bytes32 , bytes32 , (uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256[4],bytes,bytes,bytes32[],bytes,bytes) transaction) payable returns(bytes4 magic, bytes)
func (_TestPaymaster *TestPaymasterSession) ValidateAndPayForPaymasterTransaction(arg0 [32]byte, arg1 [32]byte, transaction Transaction) (*types.Transaction, error) {
	return _TestPaymaster.Contract.ValidateAndPayForPaymasterTransaction(&_TestPaymaster.TransactOpts, arg0, arg1, transaction)
}

// ValidateAndPayForPaymasterTransaction is a paid mutator transaction binding the contract method 0x038a24bc.
//
// Solidity: function validateAndPayForPaymasterTransaction(bytes32 , bytes32 , (uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256[4],bytes,bytes,bytes32[],bytes,bytes) transaction) payable returns(bytes4 magic, bytes)
func (_TestPaymaster *TestPaymasterTransactorSession) ValidateAndPayForPaymasterTransaction(arg0 [32]byte, arg1 [32]byte, transaction Transaction) (*types.Transaction, error) {
	return _TestPaymaster.Contract.ValidateAndPayForPaymasterTransaction(&_TestPaymaster.TransactOpts, arg0, arg1, transaction)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_TestPaymaster *TestPaymasterTransactor) Receive(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _TestPaymaster.contract.RawTransact(opts, nil) // calldata is disallowed for receive function
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_TestPaymaster *TestPaymasterSession) Receive() (*types.Transaction, error) {
	return _TestPaymaster.Contract.Receive(&_TestPaymaster.TransactOpts)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_TestPaymaster *TestPaymasterTransactorSession) Receive() (*types.Transaction, error) {
	return _TestPaymaster.Contract.Receive(&_TestPaymaster.TransactOpts)
}
