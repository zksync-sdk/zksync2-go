// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package L1EthBridge

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

// L1EthBridgeMetaData contains all meta data concerning the L1EthBridge contract.
var L1EthBridgeMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"contractIMailbox\",\"name\":\"_mailbox\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_depositSender\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_l1Token\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"_l2TxHash\",\"type\":\"bytes32\"},{\"internalType\":\"uint32\",\"name\":\"_l2BlockNumber\",\"type\":\"uint32\"},{\"internalType\":\"uint256\",\"name\":\"_l2MessageIndex\",\"type\":\"uint256\"},{\"internalType\":\"bytes32[]\",\"name\":\"_merkleProof\",\"type\":\"bytes32[]\"}],\"name\":\"claimFailedDeposit\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_l2Receiver\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_l1Token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"},{\"internalType\":\"enumQueueType\",\"name\":\"_queueType\",\"type\":\"uint8\"}],\"name\":\"deposit\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"txHash\",\"type\":\"bytes32\"}],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint32\",\"name\":\"_l2BlockNumber\",\"type\":\"uint32\"},{\"internalType\":\"uint256\",\"name\":\"_l2MessageIndex\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"_message\",\"type\":\"bytes\"},{\"internalType\":\"bytes32[]\",\"name\":\"_merkleProof\",\"type\":\"bytes32[]\"}],\"name\":\"finalizeWithdrawal\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"_l2BridgeBytecode\",\"type\":\"bytes\"}],\"name\":\"initialize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"l2Bridge\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_l1Token\",\"type\":\"address\"}],\"name\":\"l2TokenAddress\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"l2Token\",\"type\":\"address\"}],\"stateMutability\":\"pure\",\"type\":\"function\"}]",
}

// L1EthBridgeABI is the input ABI used to generate the binding from.
// Deprecated: Use L1EthBridgeMetaData.ABI instead.
var L1EthBridgeABI = L1EthBridgeMetaData.ABI

// L1EthBridge is an auto generated Go binding around an Ethereum contract.
type L1EthBridge struct {
	L1EthBridgeCaller     // Read-only binding to the contract
	L1EthBridgeTransactor // Write-only binding to the contract
	L1EthBridgeFilterer   // Log filterer for contract events
}

// L1EthBridgeCaller is an auto generated read-only Go binding around an Ethereum contract.
type L1EthBridgeCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// L1EthBridgeTransactor is an auto generated write-only Go binding around an Ethereum contract.
type L1EthBridgeTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// L1EthBridgeFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type L1EthBridgeFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// L1EthBridgeSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type L1EthBridgeSession struct {
	Contract     *L1EthBridge      // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// L1EthBridgeCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type L1EthBridgeCallerSession struct {
	Contract *L1EthBridgeCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts      // Call options to use throughout this session
}

// L1EthBridgeTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type L1EthBridgeTransactorSession struct {
	Contract     *L1EthBridgeTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts      // Transaction auth options to use throughout this session
}

// L1EthBridgeRaw is an auto generated low-level Go binding around an Ethereum contract.
type L1EthBridgeRaw struct {
	Contract *L1EthBridge // Generic contract binding to access the raw methods on
}

// L1EthBridgeCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type L1EthBridgeCallerRaw struct {
	Contract *L1EthBridgeCaller // Generic read-only contract binding to access the raw methods on
}

// L1EthBridgeTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type L1EthBridgeTransactorRaw struct {
	Contract *L1EthBridgeTransactor // Generic write-only contract binding to access the raw methods on
}

// NewL1EthBridge creates a new instance of L1EthBridge, bound to a specific deployed contract.
func NewL1EthBridge(address common.Address, backend bind.ContractBackend) (*L1EthBridge, error) {
	contract, err := bindL1EthBridge(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &L1EthBridge{L1EthBridgeCaller: L1EthBridgeCaller{contract: contract}, L1EthBridgeTransactor: L1EthBridgeTransactor{contract: contract}, L1EthBridgeFilterer: L1EthBridgeFilterer{contract: contract}}, nil
}

// NewL1EthBridgeCaller creates a new read-only instance of L1EthBridge, bound to a specific deployed contract.
func NewL1EthBridgeCaller(address common.Address, caller bind.ContractCaller) (*L1EthBridgeCaller, error) {
	contract, err := bindL1EthBridge(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &L1EthBridgeCaller{contract: contract}, nil
}

// NewL1EthBridgeTransactor creates a new write-only instance of L1EthBridge, bound to a specific deployed contract.
func NewL1EthBridgeTransactor(address common.Address, transactor bind.ContractTransactor) (*L1EthBridgeTransactor, error) {
	contract, err := bindL1EthBridge(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &L1EthBridgeTransactor{contract: contract}, nil
}

// NewL1EthBridgeFilterer creates a new log filterer instance of L1EthBridge, bound to a specific deployed contract.
func NewL1EthBridgeFilterer(address common.Address, filterer bind.ContractFilterer) (*L1EthBridgeFilterer, error) {
	contract, err := bindL1EthBridge(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &L1EthBridgeFilterer{contract: contract}, nil
}

// bindL1EthBridge binds a generic wrapper to an already deployed contract.
func bindL1EthBridge(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(L1EthBridgeABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_L1EthBridge *L1EthBridgeRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _L1EthBridge.Contract.L1EthBridgeCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_L1EthBridge *L1EthBridgeRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _L1EthBridge.Contract.L1EthBridgeTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_L1EthBridge *L1EthBridgeRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _L1EthBridge.Contract.L1EthBridgeTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_L1EthBridge *L1EthBridgeCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _L1EthBridge.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_L1EthBridge *L1EthBridgeTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _L1EthBridge.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_L1EthBridge *L1EthBridgeTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _L1EthBridge.Contract.contract.Transact(opts, method, params...)
}

// L2Bridge is a free data retrieval call binding the contract method 0xae1f6aaf.
//
// Solidity: function l2Bridge() view returns(address)
func (_L1EthBridge *L1EthBridgeCaller) L2Bridge(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _L1EthBridge.contract.Call(opts, &out, "l2Bridge")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// L2Bridge is a free data retrieval call binding the contract method 0xae1f6aaf.
//
// Solidity: function l2Bridge() view returns(address)
func (_L1EthBridge *L1EthBridgeSession) L2Bridge() (common.Address, error) {
	return _L1EthBridge.Contract.L2Bridge(&_L1EthBridge.CallOpts)
}

// L2Bridge is a free data retrieval call binding the contract method 0xae1f6aaf.
//
// Solidity: function l2Bridge() view returns(address)
func (_L1EthBridge *L1EthBridgeCallerSession) L2Bridge() (common.Address, error) {
	return _L1EthBridge.Contract.L2Bridge(&_L1EthBridge.CallOpts)
}

// L2TokenAddress is a free data retrieval call binding the contract method 0xf5f15168.
//
// Solidity: function l2TokenAddress(address _l1Token) pure returns(address l2Token)
func (_L1EthBridge *L1EthBridgeCaller) L2TokenAddress(opts *bind.CallOpts, _l1Token common.Address) (common.Address, error) {
	var out []interface{}
	err := _L1EthBridge.contract.Call(opts, &out, "l2TokenAddress", _l1Token)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// L2TokenAddress is a free data retrieval call binding the contract method 0xf5f15168.
//
// Solidity: function l2TokenAddress(address _l1Token) pure returns(address l2Token)
func (_L1EthBridge *L1EthBridgeSession) L2TokenAddress(_l1Token common.Address) (common.Address, error) {
	return _L1EthBridge.Contract.L2TokenAddress(&_L1EthBridge.CallOpts, _l1Token)
}

// L2TokenAddress is a free data retrieval call binding the contract method 0xf5f15168.
//
// Solidity: function l2TokenAddress(address _l1Token) pure returns(address l2Token)
func (_L1EthBridge *L1EthBridgeCallerSession) L2TokenAddress(_l1Token common.Address) (common.Address, error) {
	return _L1EthBridge.Contract.L2TokenAddress(&_L1EthBridge.CallOpts, _l1Token)
}

// ClaimFailedDeposit is a paid mutator transaction binding the contract method 0x772f7128.
//
// Solidity: function claimFailedDeposit(address _depositSender, address _l1Token, bytes32 _l2TxHash, uint32 _l2BlockNumber, uint256 _l2MessageIndex, bytes32[] _merkleProof) returns()
func (_L1EthBridge *L1EthBridgeTransactor) ClaimFailedDeposit(opts *bind.TransactOpts, _depositSender common.Address, _l1Token common.Address, _l2TxHash [32]byte, _l2BlockNumber uint32, _l2MessageIndex *big.Int, _merkleProof [][32]byte) (*types.Transaction, error) {
	return _L1EthBridge.contract.Transact(opts, "claimFailedDeposit", _depositSender, _l1Token, _l2TxHash, _l2BlockNumber, _l2MessageIndex, _merkleProof)
}

// ClaimFailedDeposit is a paid mutator transaction binding the contract method 0x772f7128.
//
// Solidity: function claimFailedDeposit(address _depositSender, address _l1Token, bytes32 _l2TxHash, uint32 _l2BlockNumber, uint256 _l2MessageIndex, bytes32[] _merkleProof) returns()
func (_L1EthBridge *L1EthBridgeSession) ClaimFailedDeposit(_depositSender common.Address, _l1Token common.Address, _l2TxHash [32]byte, _l2BlockNumber uint32, _l2MessageIndex *big.Int, _merkleProof [][32]byte) (*types.Transaction, error) {
	return _L1EthBridge.Contract.ClaimFailedDeposit(&_L1EthBridge.TransactOpts, _depositSender, _l1Token, _l2TxHash, _l2BlockNumber, _l2MessageIndex, _merkleProof)
}

// ClaimFailedDeposit is a paid mutator transaction binding the contract method 0x772f7128.
//
// Solidity: function claimFailedDeposit(address _depositSender, address _l1Token, bytes32 _l2TxHash, uint32 _l2BlockNumber, uint256 _l2MessageIndex, bytes32[] _merkleProof) returns()
func (_L1EthBridge *L1EthBridgeTransactorSession) ClaimFailedDeposit(_depositSender common.Address, _l1Token common.Address, _l2TxHash [32]byte, _l2BlockNumber uint32, _l2MessageIndex *big.Int, _merkleProof [][32]byte) (*types.Transaction, error) {
	return _L1EthBridge.Contract.ClaimFailedDeposit(&_L1EthBridge.TransactOpts, _depositSender, _l1Token, _l2TxHash, _l2BlockNumber, _l2MessageIndex, _merkleProof)
}

// Deposit is a paid mutator transaction binding the contract method 0x7f63f618.
//
// Solidity: function deposit(address _l2Receiver, address _l1Token, uint256 _amount, uint8 _queueType) payable returns(bytes32 txHash)
func (_L1EthBridge *L1EthBridgeTransactor) Deposit(opts *bind.TransactOpts, _l2Receiver common.Address, _l1Token common.Address, _amount *big.Int, _queueType uint8) (*types.Transaction, error) {
	return _L1EthBridge.contract.Transact(opts, "deposit", _l2Receiver, _l1Token, _amount, _queueType)
}

// Deposit is a paid mutator transaction binding the contract method 0x7f63f618.
//
// Solidity: function deposit(address _l2Receiver, address _l1Token, uint256 _amount, uint8 _queueType) payable returns(bytes32 txHash)
func (_L1EthBridge *L1EthBridgeSession) Deposit(_l2Receiver common.Address, _l1Token common.Address, _amount *big.Int, _queueType uint8) (*types.Transaction, error) {
	return _L1EthBridge.Contract.Deposit(&_L1EthBridge.TransactOpts, _l2Receiver, _l1Token, _amount, _queueType)
}

// Deposit is a paid mutator transaction binding the contract method 0x7f63f618.
//
// Solidity: function deposit(address _l2Receiver, address _l1Token, uint256 _amount, uint8 _queueType) payable returns(bytes32 txHash)
func (_L1EthBridge *L1EthBridgeTransactorSession) Deposit(_l2Receiver common.Address, _l1Token common.Address, _amount *big.Int, _queueType uint8) (*types.Transaction, error) {
	return _L1EthBridge.Contract.Deposit(&_L1EthBridge.TransactOpts, _l2Receiver, _l1Token, _amount, _queueType)
}

// FinalizeWithdrawal is a paid mutator transaction binding the contract method 0x882f6b96.
//
// Solidity: function finalizeWithdrawal(uint32 _l2BlockNumber, uint256 _l2MessageIndex, bytes _message, bytes32[] _merkleProof) returns()
func (_L1EthBridge *L1EthBridgeTransactor) FinalizeWithdrawal(opts *bind.TransactOpts, _l2BlockNumber uint32, _l2MessageIndex *big.Int, _message []byte, _merkleProof [][32]byte) (*types.Transaction, error) {
	return _L1EthBridge.contract.Transact(opts, "finalizeWithdrawal", _l2BlockNumber, _l2MessageIndex, _message, _merkleProof)
}

// FinalizeWithdrawal is a paid mutator transaction binding the contract method 0x882f6b96.
//
// Solidity: function finalizeWithdrawal(uint32 _l2BlockNumber, uint256 _l2MessageIndex, bytes _message, bytes32[] _merkleProof) returns()
func (_L1EthBridge *L1EthBridgeSession) FinalizeWithdrawal(_l2BlockNumber uint32, _l2MessageIndex *big.Int, _message []byte, _merkleProof [][32]byte) (*types.Transaction, error) {
	return _L1EthBridge.Contract.FinalizeWithdrawal(&_L1EthBridge.TransactOpts, _l2BlockNumber, _l2MessageIndex, _message, _merkleProof)
}

// FinalizeWithdrawal is a paid mutator transaction binding the contract method 0x882f6b96.
//
// Solidity: function finalizeWithdrawal(uint32 _l2BlockNumber, uint256 _l2MessageIndex, bytes _message, bytes32[] _merkleProof) returns()
func (_L1EthBridge *L1EthBridgeTransactorSession) FinalizeWithdrawal(_l2BlockNumber uint32, _l2MessageIndex *big.Int, _message []byte, _merkleProof [][32]byte) (*types.Transaction, error) {
	return _L1EthBridge.Contract.FinalizeWithdrawal(&_L1EthBridge.TransactOpts, _l2BlockNumber, _l2MessageIndex, _message, _merkleProof)
}

// Initialize is a paid mutator transaction binding the contract method 0x439fab91.
//
// Solidity: function initialize(bytes _l2BridgeBytecode) returns()
func (_L1EthBridge *L1EthBridgeTransactor) Initialize(opts *bind.TransactOpts, _l2BridgeBytecode []byte) (*types.Transaction, error) {
	return _L1EthBridge.contract.Transact(opts, "initialize", _l2BridgeBytecode)
}

// Initialize is a paid mutator transaction binding the contract method 0x439fab91.
//
// Solidity: function initialize(bytes _l2BridgeBytecode) returns()
func (_L1EthBridge *L1EthBridgeSession) Initialize(_l2BridgeBytecode []byte) (*types.Transaction, error) {
	return _L1EthBridge.Contract.Initialize(&_L1EthBridge.TransactOpts, _l2BridgeBytecode)
}

// Initialize is a paid mutator transaction binding the contract method 0x439fab91.
//
// Solidity: function initialize(bytes _l2BridgeBytecode) returns()
func (_L1EthBridge *L1EthBridgeTransactorSession) Initialize(_l2BridgeBytecode []byte) (*types.Transaction, error) {
	return _L1EthBridge.Contract.Initialize(&_L1EthBridge.TransactOpts, _l2BridgeBytecode)
}
