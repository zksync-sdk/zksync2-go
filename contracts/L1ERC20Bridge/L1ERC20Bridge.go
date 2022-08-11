// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package L1ERC20Bridge

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
)

// L1ERC20BridgeMetaData contains all meta data concerning the L1ERC20Bridge contract.
var L1ERC20BridgeMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"contractIMailbox\",\"name\":\"_mailbox\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_depositSender\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_l1Token\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"_l2TxHash\",\"type\":\"bytes32\"},{\"internalType\":\"uint32\",\"name\":\"_l2BlockNumber\",\"type\":\"uint32\"},{\"internalType\":\"uint256\",\"name\":\"_l2MessageIndex\",\"type\":\"uint256\"},{\"internalType\":\"bytes32[]\",\"name\":\"_merkleProof\",\"type\":\"bytes32[]\"}],\"name\":\"claimFailedDeposit\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_l2Receiver\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_l1Token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"},{\"internalType\":\"enumQueueType\",\"name\":\"_queueType\",\"type\":\"uint8\"}],\"name\":\"deposit\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"txHash\",\"type\":\"bytes32\"}],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint32\",\"name\":\"_l2BlockNumber\",\"type\":\"uint32\"},{\"internalType\":\"uint256\",\"name\":\"_l2MessageIndex\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"_message\",\"type\":\"bytes\"},{\"internalType\":\"bytes32[]\",\"name\":\"_merkleProof\",\"type\":\"bytes32[]\"}],\"name\":\"finalizeWithdrawal\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"_l2BridgeBytecode\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"_l2StandardERC20Bytecode\",\"type\":\"bytes\"}],\"name\":\"initialize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"l2Bridge\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"l2StandardERC20BytecodeHash\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_l1Token\",\"type\":\"address\"}],\"name\":\"l2TokenAddress\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
}

// L1ERC20BridgeABI is the input ABI used to generate the binding from.
// Deprecated: Use L1ERC20BridgeMetaData.ABI instead.
var L1ERC20BridgeABI = L1ERC20BridgeMetaData.ABI

// L1ERC20Bridge is an auto generated Go binding around an Ethereum contract.
type L1ERC20Bridge struct {
	L1ERC20BridgeCaller     // Read-only binding to the contract
	L1ERC20BridgeTransactor // Write-only binding to the contract
	L1ERC20BridgeFilterer   // Log filterer for contract events
}

// L1ERC20BridgeCaller is an auto generated read-only Go binding around an Ethereum contract.
type L1ERC20BridgeCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// L1ERC20BridgeTransactor is an auto generated write-only Go binding around an Ethereum contract.
type L1ERC20BridgeTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// L1ERC20BridgeFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type L1ERC20BridgeFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// L1ERC20BridgeSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type L1ERC20BridgeSession struct {
	Contract     *L1ERC20Bridge    // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// L1ERC20BridgeCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type L1ERC20BridgeCallerSession struct {
	Contract *L1ERC20BridgeCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts        // Call options to use throughout this session
}

// L1ERC20BridgeTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type L1ERC20BridgeTransactorSession struct {
	Contract     *L1ERC20BridgeTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts        // Transaction auth options to use throughout this session
}

// L1ERC20BridgeRaw is an auto generated low-level Go binding around an Ethereum contract.
type L1ERC20BridgeRaw struct {
	Contract *L1ERC20Bridge // Generic contract binding to access the raw methods on
}

// L1ERC20BridgeCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type L1ERC20BridgeCallerRaw struct {
	Contract *L1ERC20BridgeCaller // Generic read-only contract binding to access the raw methods on
}

// L1ERC20BridgeTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type L1ERC20BridgeTransactorRaw struct {
	Contract *L1ERC20BridgeTransactor // Generic write-only contract binding to access the raw methods on
}

// NewL1ERC20Bridge creates a new instance of L1ERC20Bridge, bound to a specific deployed contract.
func NewL1ERC20Bridge(address common.Address, backend bind.ContractBackend) (*L1ERC20Bridge, error) {
	contract, err := bindL1ERC20Bridge(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &L1ERC20Bridge{L1ERC20BridgeCaller: L1ERC20BridgeCaller{contract: contract}, L1ERC20BridgeTransactor: L1ERC20BridgeTransactor{contract: contract}, L1ERC20BridgeFilterer: L1ERC20BridgeFilterer{contract: contract}}, nil
}

// NewL1ERC20BridgeCaller creates a new read-only instance of L1ERC20Bridge, bound to a specific deployed contract.
func NewL1ERC20BridgeCaller(address common.Address, caller bind.ContractCaller) (*L1ERC20BridgeCaller, error) {
	contract, err := bindL1ERC20Bridge(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &L1ERC20BridgeCaller{contract: contract}, nil
}

// NewL1ERC20BridgeTransactor creates a new write-only instance of L1ERC20Bridge, bound to a specific deployed contract.
func NewL1ERC20BridgeTransactor(address common.Address, transactor bind.ContractTransactor) (*L1ERC20BridgeTransactor, error) {
	contract, err := bindL1ERC20Bridge(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &L1ERC20BridgeTransactor{contract: contract}, nil
}

// NewL1ERC20BridgeFilterer creates a new log filterer instance of L1ERC20Bridge, bound to a specific deployed contract.
func NewL1ERC20BridgeFilterer(address common.Address, filterer bind.ContractFilterer) (*L1ERC20BridgeFilterer, error) {
	contract, err := bindL1ERC20Bridge(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &L1ERC20BridgeFilterer{contract: contract}, nil
}

// bindL1ERC20Bridge binds a generic wrapper to an already deployed contract.
func bindL1ERC20Bridge(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(L1ERC20BridgeABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_L1ERC20Bridge *L1ERC20BridgeRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _L1ERC20Bridge.Contract.L1ERC20BridgeCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_L1ERC20Bridge *L1ERC20BridgeRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _L1ERC20Bridge.Contract.L1ERC20BridgeTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_L1ERC20Bridge *L1ERC20BridgeRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _L1ERC20Bridge.Contract.L1ERC20BridgeTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_L1ERC20Bridge *L1ERC20BridgeCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _L1ERC20Bridge.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_L1ERC20Bridge *L1ERC20BridgeTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _L1ERC20Bridge.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_L1ERC20Bridge *L1ERC20BridgeTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _L1ERC20Bridge.Contract.contract.Transact(opts, method, params...)
}

// L2Bridge is a free data retrieval call binding the contract method 0xae1f6aaf.
//
// Solidity: function l2Bridge() view returns(address)
func (_L1ERC20Bridge *L1ERC20BridgeCaller) L2Bridge(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _L1ERC20Bridge.contract.Call(opts, &out, "l2Bridge")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// L2Bridge is a free data retrieval call binding the contract method 0xae1f6aaf.
//
// Solidity: function l2Bridge() view returns(address)
func (_L1ERC20Bridge *L1ERC20BridgeSession) L2Bridge() (common.Address, error) {
	return _L1ERC20Bridge.Contract.L2Bridge(&_L1ERC20Bridge.CallOpts)
}

// L2Bridge is a free data retrieval call binding the contract method 0xae1f6aaf.
//
// Solidity: function l2Bridge() view returns(address)
func (_L1ERC20Bridge *L1ERC20BridgeCallerSession) L2Bridge() (common.Address, error) {
	return _L1ERC20Bridge.Contract.L2Bridge(&_L1ERC20Bridge.CallOpts)
}

// L2StandardERC20BytecodeHash is a free data retrieval call binding the contract method 0xb4e5634a.
//
// Solidity: function l2StandardERC20BytecodeHash() view returns(bytes32)
func (_L1ERC20Bridge *L1ERC20BridgeCaller) L2StandardERC20BytecodeHash(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _L1ERC20Bridge.contract.Call(opts, &out, "l2StandardERC20BytecodeHash")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// L2StandardERC20BytecodeHash is a free data retrieval call binding the contract method 0xb4e5634a.
//
// Solidity: function l2StandardERC20BytecodeHash() view returns(bytes32)
func (_L1ERC20Bridge *L1ERC20BridgeSession) L2StandardERC20BytecodeHash() ([32]byte, error) {
	return _L1ERC20Bridge.Contract.L2StandardERC20BytecodeHash(&_L1ERC20Bridge.CallOpts)
}

// L2StandardERC20BytecodeHash is a free data retrieval call binding the contract method 0xb4e5634a.
//
// Solidity: function l2StandardERC20BytecodeHash() view returns(bytes32)
func (_L1ERC20Bridge *L1ERC20BridgeCallerSession) L2StandardERC20BytecodeHash() ([32]byte, error) {
	return _L1ERC20Bridge.Contract.L2StandardERC20BytecodeHash(&_L1ERC20Bridge.CallOpts)
}

// L2TokenAddress is a free data retrieval call binding the contract method 0xf5f15168.
//
// Solidity: function l2TokenAddress(address _l1Token) view returns(address)
func (_L1ERC20Bridge *L1ERC20BridgeCaller) L2TokenAddress(opts *bind.CallOpts, _l1Token common.Address) (common.Address, error) {
	var out []interface{}
	err := _L1ERC20Bridge.contract.Call(opts, &out, "l2TokenAddress", _l1Token)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// L2TokenAddress is a free data retrieval call binding the contract method 0xf5f15168.
//
// Solidity: function l2TokenAddress(address _l1Token) view returns(address)
func (_L1ERC20Bridge *L1ERC20BridgeSession) L2TokenAddress(_l1Token common.Address) (common.Address, error) {
	return _L1ERC20Bridge.Contract.L2TokenAddress(&_L1ERC20Bridge.CallOpts, _l1Token)
}

// L2TokenAddress is a free data retrieval call binding the contract method 0xf5f15168.
//
// Solidity: function l2TokenAddress(address _l1Token) view returns(address)
func (_L1ERC20Bridge *L1ERC20BridgeCallerSession) L2TokenAddress(_l1Token common.Address) (common.Address, error) {
	return _L1ERC20Bridge.Contract.L2TokenAddress(&_L1ERC20Bridge.CallOpts, _l1Token)
}

// ClaimFailedDeposit is a paid mutator transaction binding the contract method 0x772f7128.
//
// Solidity: function claimFailedDeposit(address _depositSender, address _l1Token, bytes32 _l2TxHash, uint32 _l2BlockNumber, uint256 _l2MessageIndex, bytes32[] _merkleProof) returns()
func (_L1ERC20Bridge *L1ERC20BridgeTransactor) ClaimFailedDeposit(opts *bind.TransactOpts, _depositSender common.Address, _l1Token common.Address, _l2TxHash [32]byte, _l2BlockNumber uint32, _l2MessageIndex *big.Int, _merkleProof [][32]byte) (*types.Transaction, error) {
	return _L1ERC20Bridge.contract.Transact(opts, "claimFailedDeposit", _depositSender, _l1Token, _l2TxHash, _l2BlockNumber, _l2MessageIndex, _merkleProof)
}

// ClaimFailedDeposit is a paid mutator transaction binding the contract method 0x772f7128.
//
// Solidity: function claimFailedDeposit(address _depositSender, address _l1Token, bytes32 _l2TxHash, uint32 _l2BlockNumber, uint256 _l2MessageIndex, bytes32[] _merkleProof) returns()
func (_L1ERC20Bridge *L1ERC20BridgeSession) ClaimFailedDeposit(_depositSender common.Address, _l1Token common.Address, _l2TxHash [32]byte, _l2BlockNumber uint32, _l2MessageIndex *big.Int, _merkleProof [][32]byte) (*types.Transaction, error) {
	return _L1ERC20Bridge.Contract.ClaimFailedDeposit(&_L1ERC20Bridge.TransactOpts, _depositSender, _l1Token, _l2TxHash, _l2BlockNumber, _l2MessageIndex, _merkleProof)
}

// ClaimFailedDeposit is a paid mutator transaction binding the contract method 0x772f7128.
//
// Solidity: function claimFailedDeposit(address _depositSender, address _l1Token, bytes32 _l2TxHash, uint32 _l2BlockNumber, uint256 _l2MessageIndex, bytes32[] _merkleProof) returns()
func (_L1ERC20Bridge *L1ERC20BridgeTransactorSession) ClaimFailedDeposit(_depositSender common.Address, _l1Token common.Address, _l2TxHash [32]byte, _l2BlockNumber uint32, _l2MessageIndex *big.Int, _merkleProof [][32]byte) (*types.Transaction, error) {
	return _L1ERC20Bridge.Contract.ClaimFailedDeposit(&_L1ERC20Bridge.TransactOpts, _depositSender, _l1Token, _l2TxHash, _l2BlockNumber, _l2MessageIndex, _merkleProof)
}

// Deposit is a paid mutator transaction binding the contract method 0x7f63f618.
//
// Solidity: function deposit(address _l2Receiver, address _l1Token, uint256 _amount, uint8 _queueType) payable returns(bytes32 txHash)
func (_L1ERC20Bridge *L1ERC20BridgeTransactor) Deposit(opts *bind.TransactOpts, _l2Receiver common.Address, _l1Token common.Address, _amount *big.Int, _queueType uint8) (*types.Transaction, error) {
	return _L1ERC20Bridge.contract.Transact(opts, "deposit", _l2Receiver, _l1Token, _amount, _queueType)
}

// Deposit is a paid mutator transaction binding the contract method 0x7f63f618.
//
// Solidity: function deposit(address _l2Receiver, address _l1Token, uint256 _amount, uint8 _queueType) payable returns(bytes32 txHash)
func (_L1ERC20Bridge *L1ERC20BridgeSession) Deposit(_l2Receiver common.Address, _l1Token common.Address, _amount *big.Int, _queueType uint8) (*types.Transaction, error) {
	return _L1ERC20Bridge.Contract.Deposit(&_L1ERC20Bridge.TransactOpts, _l2Receiver, _l1Token, _amount, _queueType)
}

// Deposit is a paid mutator transaction binding the contract method 0x7f63f618.
//
// Solidity: function deposit(address _l2Receiver, address _l1Token, uint256 _amount, uint8 _queueType) payable returns(bytes32 txHash)
func (_L1ERC20Bridge *L1ERC20BridgeTransactorSession) Deposit(_l2Receiver common.Address, _l1Token common.Address, _amount *big.Int, _queueType uint8) (*types.Transaction, error) {
	return _L1ERC20Bridge.Contract.Deposit(&_L1ERC20Bridge.TransactOpts, _l2Receiver, _l1Token, _amount, _queueType)
}

// FinalizeWithdrawal is a paid mutator transaction binding the contract method 0x882f6b96.
//
// Solidity: function finalizeWithdrawal(uint32 _l2BlockNumber, uint256 _l2MessageIndex, bytes _message, bytes32[] _merkleProof) returns()
func (_L1ERC20Bridge *L1ERC20BridgeTransactor) FinalizeWithdrawal(opts *bind.TransactOpts, _l2BlockNumber uint32, _l2MessageIndex *big.Int, _message []byte, _merkleProof [][32]byte) (*types.Transaction, error) {
	return _L1ERC20Bridge.contract.Transact(opts, "finalizeWithdrawal", _l2BlockNumber, _l2MessageIndex, _message, _merkleProof)
}

// FinalizeWithdrawal is a paid mutator transaction binding the contract method 0x882f6b96.
//
// Solidity: function finalizeWithdrawal(uint32 _l2BlockNumber, uint256 _l2MessageIndex, bytes _message, bytes32[] _merkleProof) returns()
func (_L1ERC20Bridge *L1ERC20BridgeSession) FinalizeWithdrawal(_l2BlockNumber uint32, _l2MessageIndex *big.Int, _message []byte, _merkleProof [][32]byte) (*types.Transaction, error) {
	return _L1ERC20Bridge.Contract.FinalizeWithdrawal(&_L1ERC20Bridge.TransactOpts, _l2BlockNumber, _l2MessageIndex, _message, _merkleProof)
}

// FinalizeWithdrawal is a paid mutator transaction binding the contract method 0x882f6b96.
//
// Solidity: function finalizeWithdrawal(uint32 _l2BlockNumber, uint256 _l2MessageIndex, bytes _message, bytes32[] _merkleProof) returns()
func (_L1ERC20Bridge *L1ERC20BridgeTransactorSession) FinalizeWithdrawal(_l2BlockNumber uint32, _l2MessageIndex *big.Int, _message []byte, _merkleProof [][32]byte) (*types.Transaction, error) {
	return _L1ERC20Bridge.Contract.FinalizeWithdrawal(&_L1ERC20Bridge.TransactOpts, _l2BlockNumber, _l2MessageIndex, _message, _merkleProof)
}

// Initialize is a paid mutator transaction binding the contract method 0x1af19f77.
//
// Solidity: function initialize(bytes _l2BridgeBytecode, bytes _l2StandardERC20Bytecode) returns()
func (_L1ERC20Bridge *L1ERC20BridgeTransactor) Initialize(opts *bind.TransactOpts, _l2BridgeBytecode []byte, _l2StandardERC20Bytecode []byte) (*types.Transaction, error) {
	return _L1ERC20Bridge.contract.Transact(opts, "initialize", _l2BridgeBytecode, _l2StandardERC20Bytecode)
}

// Initialize is a paid mutator transaction binding the contract method 0x1af19f77.
//
// Solidity: function initialize(bytes _l2BridgeBytecode, bytes _l2StandardERC20Bytecode) returns()
func (_L1ERC20Bridge *L1ERC20BridgeSession) Initialize(_l2BridgeBytecode []byte, _l2StandardERC20Bytecode []byte) (*types.Transaction, error) {
	return _L1ERC20Bridge.Contract.Initialize(&_L1ERC20Bridge.TransactOpts, _l2BridgeBytecode, _l2StandardERC20Bytecode)
}

// Initialize is a paid mutator transaction binding the contract method 0x1af19f77.
//
// Solidity: function initialize(bytes _l2BridgeBytecode, bytes _l2StandardERC20Bytecode) returns()
func (_L1ERC20Bridge *L1ERC20BridgeTransactorSession) Initialize(_l2BridgeBytecode []byte, _l2StandardERC20Bytecode []byte) (*types.Transaction, error) {
	return _L1ERC20Bridge.Contract.Initialize(&_L1ERC20Bridge.TransactOpts, _l2BridgeBytecode, _l2StandardERC20Bytecode)
}
