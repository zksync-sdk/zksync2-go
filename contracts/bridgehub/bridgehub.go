// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package bridgehub

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

// L2Log is an auto generated low-level Go binding around an user-defined struct.
type L2Log struct {
	L2ShardId       uint8
	IsService       bool
	TxNumberInBatch uint16
	Sender          common.Address
	Key             [32]byte
	Value           [32]byte
}

// L2Message is an auto generated low-level Go binding around an user-defined struct.
type L2Message struct {
	TxNumberInBatch uint16
	Sender          common.Address
	Data            []byte
}

// L2TransactionRequestDirect is an auto generated low-level Go binding around an user-defined struct.
type L2TransactionRequestDirect struct {
	ChainId                  *big.Int
	MintValue                *big.Int
	L2Contract               common.Address
	L2Value                  *big.Int
	L2Calldata               []byte
	L2GasLimit               *big.Int
	L2GasPerPubdataByteLimit *big.Int
	FactoryDeps              [][]byte
	RefundRecipient          common.Address
}

// L2TransactionRequestTwoBridgesOuter is an auto generated low-level Go binding around an user-defined struct.
type L2TransactionRequestTwoBridgesOuter struct {
	ChainId                  *big.Int
	MintValue                *big.Int
	L2Value                  *big.Int
	L2GasLimit               *big.Int
	L2GasPerPubdataByteLimit *big.Int
	RefundRecipient          common.Address
	SecondBridgeAddress      common.Address
	SecondBridgeValue        *big.Int
	SecondBridgeCalldata     []byte
}

// IBridgehubMetaData contains all meta data concerning the IBridgehub contract.
var IBridgehubMetaData = &bind.MetaData{
	ABI: "[{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"oldAdmin\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newAdmin\",\"type\":\"address\"}],\"name\":\"NewAdmin\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"chainId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"stateTransitionManager\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"chainGovernance\",\"type\":\"address\"}],\"name\":\"NewChain\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"oldPendingAdmin\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newPendingAdmin\",\"type\":\"address\"}],\"name\":\"NewPendingAdmin\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"acceptAdmin\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_stateTransitionManager\",\"type\":\"address\"}],\"name\":\"addStateTransitionManager\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_token\",\"type\":\"address\"}],\"name\":\"addToken\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_chainId\",\"type\":\"uint256\"}],\"name\":\"baseToken\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_chainId\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"_stateTransitionManager\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_baseToken\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_salt\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"_admin\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"_initData\",\"type\":\"bytes\"}],\"name\":\"createNewChain\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"chainId\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_chainId\",\"type\":\"uint256\"}],\"name\":\"getHyperchain\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_chainId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_gasPrice\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_l2GasLimit\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_l2GasPerPubdataByteLimit\",\"type\":\"uint256\"}],\"name\":\"l2TransactionBaseCost\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_chainId\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"_l2TxHash\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"_l2BatchNumber\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_l2MessageIndex\",\"type\":\"uint256\"},{\"internalType\":\"uint16\",\"name\":\"_l2TxNumberInBatch\",\"type\":\"uint16\"},{\"internalType\":\"bytes32[]\",\"name\":\"_merkleProof\",\"type\":\"bytes32[]\"},{\"internalType\":\"enumTxStatus\",\"name\":\"_status\",\"type\":\"uint8\"}],\"name\":\"proveL1ToL2TransactionStatus\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_chainId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_batchNumber\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_index\",\"type\":\"uint256\"},{\"components\":[{\"internalType\":\"uint8\",\"name\":\"l2ShardId\",\"type\":\"uint8\"},{\"internalType\":\"bool\",\"name\":\"isService\",\"type\":\"bool\"},{\"internalType\":\"uint16\",\"name\":\"txNumberInBatch\",\"type\":\"uint16\"},{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"key\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"value\",\"type\":\"bytes32\"}],\"internalType\":\"structL2Log\",\"name\":\"_log\",\"type\":\"tuple\"},{\"internalType\":\"bytes32[]\",\"name\":\"_proof\",\"type\":\"bytes32[]\"}],\"name\":\"proveL2LogInclusion\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_chainId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_batchNumber\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_index\",\"type\":\"uint256\"},{\"components\":[{\"internalType\":\"uint16\",\"name\":\"txNumberInBatch\",\"type\":\"uint16\"},{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"internalType\":\"structL2Message\",\"name\":\"_message\",\"type\":\"tuple\"},{\"internalType\":\"bytes32[]\",\"name\":\"_proof\",\"type\":\"bytes32[]\"}],\"name\":\"proveL2MessageInclusion\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_stateTransitionManager\",\"type\":\"address\"}],\"name\":\"removeStateTransitionManager\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"chainId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"mintValue\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"l2Contract\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"l2Value\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"l2Calldata\",\"type\":\"bytes\"},{\"internalType\":\"uint256\",\"name\":\"l2GasLimit\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"l2GasPerPubdataByteLimit\",\"type\":\"uint256\"},{\"internalType\":\"bytes[]\",\"name\":\"factoryDeps\",\"type\":\"bytes[]\"},{\"internalType\":\"address\",\"name\":\"refundRecipient\",\"type\":\"address\"}],\"internalType\":\"structL2TransactionRequestDirect\",\"name\":\"_request\",\"type\":\"tuple\"}],\"name\":\"requestL2TransactionDirect\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"canonicalTxHash\",\"type\":\"bytes32\"}],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"chainId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"mintValue\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"l2Value\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"l2GasLimit\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"l2GasPerPubdataByteLimit\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"refundRecipient\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"secondBridgeAddress\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"secondBridgeValue\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"secondBridgeCalldata\",\"type\":\"bytes\"}],\"internalType\":\"structL2TransactionRequestTwoBridgesOuter\",\"name\":\"_request\",\"type\":\"tuple\"}],\"name\":\"requestL2TransactionTwoBridges\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"canonicalTxHash\",\"type\":\"bytes32\"}],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_newPendingAdmin\",\"type\":\"address\"}],\"name\":\"setPendingAdmin\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_sharedBridge\",\"type\":\"address\"}],\"name\":\"setSharedBridge\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"sharedBridge\",\"outputs\":[{\"internalType\":\"contractIL1SharedBridge\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_chainId\",\"type\":\"uint256\"}],\"name\":\"stateTransitionManager\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_stateTransitionManager\",\"type\":\"address\"}],\"name\":\"stateTransitionManagerIsRegistered\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_baseToken\",\"type\":\"address\"}],\"name\":\"tokenIsRegistered\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
}

// IBridgehubABI is the input ABI used to generate the binding from.
// Deprecated: Use IBridgehubMetaData.ABI instead.
var IBridgehubABI = IBridgehubMetaData.ABI

// IBridgehub is an auto generated Go binding around an Ethereum contract.
type IBridgehub struct {
	IBridgehubCaller     // Read-only binding to the contract
	IBridgehubTransactor // Write-only binding to the contract
	IBridgehubFilterer   // Log filterer for contract events
}

// IBridgehubCaller is an auto generated read-only Go binding around an Ethereum contract.
type IBridgehubCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IBridgehubTransactor is an auto generated write-only Go binding around an Ethereum contract.
type IBridgehubTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IBridgehubFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type IBridgehubFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IBridgehubSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type IBridgehubSession struct {
	Contract     *IBridgehub       // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// IBridgehubCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type IBridgehubCallerSession struct {
	Contract *IBridgehubCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts     // Call options to use throughout this session
}

// IBridgehubTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type IBridgehubTransactorSession struct {
	Contract     *IBridgehubTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts     // Transaction auth options to use throughout this session
}

// IBridgehubRaw is an auto generated low-level Go binding around an Ethereum contract.
type IBridgehubRaw struct {
	Contract *IBridgehub // Generic contract binding to access the raw methods on
}

// IBridgehubCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type IBridgehubCallerRaw struct {
	Contract *IBridgehubCaller // Generic read-only contract binding to access the raw methods on
}

// IBridgehubTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type IBridgehubTransactorRaw struct {
	Contract *IBridgehubTransactor // Generic write-only contract binding to access the raw methods on
}

// NewIBridgehub creates a new instance of IBridgehub, bound to a specific deployed contract.
func NewIBridgehub(address common.Address, backend bind.ContractBackend) (*IBridgehub, error) {
	contract, err := bindIBridgehub(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &IBridgehub{IBridgehubCaller: IBridgehubCaller{contract: contract}, IBridgehubTransactor: IBridgehubTransactor{contract: contract}, IBridgehubFilterer: IBridgehubFilterer{contract: contract}}, nil
}

// NewIBridgehubCaller creates a new read-only instance of IBridgehub, bound to a specific deployed contract.
func NewIBridgehubCaller(address common.Address, caller bind.ContractCaller) (*IBridgehubCaller, error) {
	contract, err := bindIBridgehub(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &IBridgehubCaller{contract: contract}, nil
}

// NewIBridgehubTransactor creates a new write-only instance of IBridgehub, bound to a specific deployed contract.
func NewIBridgehubTransactor(address common.Address, transactor bind.ContractTransactor) (*IBridgehubTransactor, error) {
	contract, err := bindIBridgehub(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &IBridgehubTransactor{contract: contract}, nil
}

// NewIBridgehubFilterer creates a new log filterer instance of IBridgehub, bound to a specific deployed contract.
func NewIBridgehubFilterer(address common.Address, filterer bind.ContractFilterer) (*IBridgehubFilterer, error) {
	contract, err := bindIBridgehub(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &IBridgehubFilterer{contract: contract}, nil
}

// bindIBridgehub binds a generic wrapper to an already deployed contract.
func bindIBridgehub(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := IBridgehubMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IBridgehub *IBridgehubRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IBridgehub.Contract.IBridgehubCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IBridgehub *IBridgehubRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IBridgehub.Contract.IBridgehubTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IBridgehub *IBridgehubRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IBridgehub.Contract.IBridgehubTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IBridgehub *IBridgehubCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IBridgehub.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IBridgehub *IBridgehubTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IBridgehub.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IBridgehub *IBridgehubTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IBridgehub.Contract.contract.Transact(opts, method, params...)
}

// BaseToken is a free data retrieval call binding the contract method 0x59ec65a2.
//
// Solidity: function baseToken(uint256 _chainId) view returns(address)
func (_IBridgehub *IBridgehubCaller) BaseToken(opts *bind.CallOpts, _chainId *big.Int) (common.Address, error) {
	var out []interface{}
	err := _IBridgehub.contract.Call(opts, &out, "baseToken", _chainId)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// BaseToken is a free data retrieval call binding the contract method 0x59ec65a2.
//
// Solidity: function baseToken(uint256 _chainId) view returns(address)
func (_IBridgehub *IBridgehubSession) BaseToken(_chainId *big.Int) (common.Address, error) {
	return _IBridgehub.Contract.BaseToken(&_IBridgehub.CallOpts, _chainId)
}

// BaseToken is a free data retrieval call binding the contract method 0x59ec65a2.
//
// Solidity: function baseToken(uint256 _chainId) view returns(address)
func (_IBridgehub *IBridgehubCallerSession) BaseToken(_chainId *big.Int) (common.Address, error) {
	return _IBridgehub.Contract.BaseToken(&_IBridgehub.CallOpts, _chainId)
}

// GetHyperchain is a free data retrieval call binding the contract method 0xdead6f7f.
//
// Solidity: function getHyperchain(uint256 _chainId) view returns(address)
func (_IBridgehub *IBridgehubCaller) GetHyperchain(opts *bind.CallOpts, _chainId *big.Int) (common.Address, error) {
	var out []interface{}
	err := _IBridgehub.contract.Call(opts, &out, "getHyperchain", _chainId)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetHyperchain is a free data retrieval call binding the contract method 0xdead6f7f.
//
// Solidity: function getHyperchain(uint256 _chainId) view returns(address)
func (_IBridgehub *IBridgehubSession) GetHyperchain(_chainId *big.Int) (common.Address, error) {
	return _IBridgehub.Contract.GetHyperchain(&_IBridgehub.CallOpts, _chainId)
}

// GetHyperchain is a free data retrieval call binding the contract method 0xdead6f7f.
//
// Solidity: function getHyperchain(uint256 _chainId) view returns(address)
func (_IBridgehub *IBridgehubCallerSession) GetHyperchain(_chainId *big.Int) (common.Address, error) {
	return _IBridgehub.Contract.GetHyperchain(&_IBridgehub.CallOpts, _chainId)
}

// L2TransactionBaseCost is a free data retrieval call binding the contract method 0x71623274.
//
// Solidity: function l2TransactionBaseCost(uint256 _chainId, uint256 _gasPrice, uint256 _l2GasLimit, uint256 _l2GasPerPubdataByteLimit) view returns(uint256)
func (_IBridgehub *IBridgehubCaller) L2TransactionBaseCost(opts *bind.CallOpts, _chainId *big.Int, _gasPrice *big.Int, _l2GasLimit *big.Int, _l2GasPerPubdataByteLimit *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _IBridgehub.contract.Call(opts, &out, "l2TransactionBaseCost", _chainId, _gasPrice, _l2GasLimit, _l2GasPerPubdataByteLimit)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// L2TransactionBaseCost is a free data retrieval call binding the contract method 0x71623274.
//
// Solidity: function l2TransactionBaseCost(uint256 _chainId, uint256 _gasPrice, uint256 _l2GasLimit, uint256 _l2GasPerPubdataByteLimit) view returns(uint256)
func (_IBridgehub *IBridgehubSession) L2TransactionBaseCost(_chainId *big.Int, _gasPrice *big.Int, _l2GasLimit *big.Int, _l2GasPerPubdataByteLimit *big.Int) (*big.Int, error) {
	return _IBridgehub.Contract.L2TransactionBaseCost(&_IBridgehub.CallOpts, _chainId, _gasPrice, _l2GasLimit, _l2GasPerPubdataByteLimit)
}

// L2TransactionBaseCost is a free data retrieval call binding the contract method 0x71623274.
//
// Solidity: function l2TransactionBaseCost(uint256 _chainId, uint256 _gasPrice, uint256 _l2GasLimit, uint256 _l2GasPerPubdataByteLimit) view returns(uint256)
func (_IBridgehub *IBridgehubCallerSession) L2TransactionBaseCost(_chainId *big.Int, _gasPrice *big.Int, _l2GasLimit *big.Int, _l2GasPerPubdataByteLimit *big.Int) (*big.Int, error) {
	return _IBridgehub.Contract.L2TransactionBaseCost(&_IBridgehub.CallOpts, _chainId, _gasPrice, _l2GasLimit, _l2GasPerPubdataByteLimit)
}

// ProveL1ToL2TransactionStatus is a free data retrieval call binding the contract method 0xb292f5f1.
//
// Solidity: function proveL1ToL2TransactionStatus(uint256 _chainId, bytes32 _l2TxHash, uint256 _l2BatchNumber, uint256 _l2MessageIndex, uint16 _l2TxNumberInBatch, bytes32[] _merkleProof, uint8 _status) view returns(bool)
func (_IBridgehub *IBridgehubCaller) ProveL1ToL2TransactionStatus(opts *bind.CallOpts, _chainId *big.Int, _l2TxHash [32]byte, _l2BatchNumber *big.Int, _l2MessageIndex *big.Int, _l2TxNumberInBatch uint16, _merkleProof [][32]byte, _status uint8) (bool, error) {
	var out []interface{}
	err := _IBridgehub.contract.Call(opts, &out, "proveL1ToL2TransactionStatus", _chainId, _l2TxHash, _l2BatchNumber, _l2MessageIndex, _l2TxNumberInBatch, _merkleProof, _status)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// ProveL1ToL2TransactionStatus is a free data retrieval call binding the contract method 0xb292f5f1.
//
// Solidity: function proveL1ToL2TransactionStatus(uint256 _chainId, bytes32 _l2TxHash, uint256 _l2BatchNumber, uint256 _l2MessageIndex, uint16 _l2TxNumberInBatch, bytes32[] _merkleProof, uint8 _status) view returns(bool)
func (_IBridgehub *IBridgehubSession) ProveL1ToL2TransactionStatus(_chainId *big.Int, _l2TxHash [32]byte, _l2BatchNumber *big.Int, _l2MessageIndex *big.Int, _l2TxNumberInBatch uint16, _merkleProof [][32]byte, _status uint8) (bool, error) {
	return _IBridgehub.Contract.ProveL1ToL2TransactionStatus(&_IBridgehub.CallOpts, _chainId, _l2TxHash, _l2BatchNumber, _l2MessageIndex, _l2TxNumberInBatch, _merkleProof, _status)
}

// ProveL1ToL2TransactionStatus is a free data retrieval call binding the contract method 0xb292f5f1.
//
// Solidity: function proveL1ToL2TransactionStatus(uint256 _chainId, bytes32 _l2TxHash, uint256 _l2BatchNumber, uint256 _l2MessageIndex, uint16 _l2TxNumberInBatch, bytes32[] _merkleProof, uint8 _status) view returns(bool)
func (_IBridgehub *IBridgehubCallerSession) ProveL1ToL2TransactionStatus(_chainId *big.Int, _l2TxHash [32]byte, _l2BatchNumber *big.Int, _l2MessageIndex *big.Int, _l2TxNumberInBatch uint16, _merkleProof [][32]byte, _status uint8) (bool, error) {
	return _IBridgehub.Contract.ProveL1ToL2TransactionStatus(&_IBridgehub.CallOpts, _chainId, _l2TxHash, _l2BatchNumber, _l2MessageIndex, _l2TxNumberInBatch, _merkleProof, _status)
}

// ProveL2LogInclusion is a free data retrieval call binding the contract method 0xe6d9923b.
//
// Solidity: function proveL2LogInclusion(uint256 _chainId, uint256 _batchNumber, uint256 _index, (uint8,bool,uint16,address,bytes32,bytes32) _log, bytes32[] _proof) view returns(bool)
func (_IBridgehub *IBridgehubCaller) ProveL2LogInclusion(opts *bind.CallOpts, _chainId *big.Int, _batchNumber *big.Int, _index *big.Int, _log L2Log, _proof [][32]byte) (bool, error) {
	var out []interface{}
	err := _IBridgehub.contract.Call(opts, &out, "proveL2LogInclusion", _chainId, _batchNumber, _index, _log, _proof)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// ProveL2LogInclusion is a free data retrieval call binding the contract method 0xe6d9923b.
//
// Solidity: function proveL2LogInclusion(uint256 _chainId, uint256 _batchNumber, uint256 _index, (uint8,bool,uint16,address,bytes32,bytes32) _log, bytes32[] _proof) view returns(bool)
func (_IBridgehub *IBridgehubSession) ProveL2LogInclusion(_chainId *big.Int, _batchNumber *big.Int, _index *big.Int, _log L2Log, _proof [][32]byte) (bool, error) {
	return _IBridgehub.Contract.ProveL2LogInclusion(&_IBridgehub.CallOpts, _chainId, _batchNumber, _index, _log, _proof)
}

// ProveL2LogInclusion is a free data retrieval call binding the contract method 0xe6d9923b.
//
// Solidity: function proveL2LogInclusion(uint256 _chainId, uint256 _batchNumber, uint256 _index, (uint8,bool,uint16,address,bytes32,bytes32) _log, bytes32[] _proof) view returns(bool)
func (_IBridgehub *IBridgehubCallerSession) ProveL2LogInclusion(_chainId *big.Int, _batchNumber *big.Int, _index *big.Int, _log L2Log, _proof [][32]byte) (bool, error) {
	return _IBridgehub.Contract.ProveL2LogInclusion(&_IBridgehub.CallOpts, _chainId, _batchNumber, _index, _log, _proof)
}

// ProveL2MessageInclusion is a free data retrieval call binding the contract method 0x99c16d1a.
//
// Solidity: function proveL2MessageInclusion(uint256 _chainId, uint256 _batchNumber, uint256 _index, (uint16,address,bytes) _message, bytes32[] _proof) view returns(bool)
func (_IBridgehub *IBridgehubCaller) ProveL2MessageInclusion(opts *bind.CallOpts, _chainId *big.Int, _batchNumber *big.Int, _index *big.Int, _message L2Message, _proof [][32]byte) (bool, error) {
	var out []interface{}
	err := _IBridgehub.contract.Call(opts, &out, "proveL2MessageInclusion", _chainId, _batchNumber, _index, _message, _proof)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// ProveL2MessageInclusion is a free data retrieval call binding the contract method 0x99c16d1a.
//
// Solidity: function proveL2MessageInclusion(uint256 _chainId, uint256 _batchNumber, uint256 _index, (uint16,address,bytes) _message, bytes32[] _proof) view returns(bool)
func (_IBridgehub *IBridgehubSession) ProveL2MessageInclusion(_chainId *big.Int, _batchNumber *big.Int, _index *big.Int, _message L2Message, _proof [][32]byte) (bool, error) {
	return _IBridgehub.Contract.ProveL2MessageInclusion(&_IBridgehub.CallOpts, _chainId, _batchNumber, _index, _message, _proof)
}

// ProveL2MessageInclusion is a free data retrieval call binding the contract method 0x99c16d1a.
//
// Solidity: function proveL2MessageInclusion(uint256 _chainId, uint256 _batchNumber, uint256 _index, (uint16,address,bytes) _message, bytes32[] _proof) view returns(bool)
func (_IBridgehub *IBridgehubCallerSession) ProveL2MessageInclusion(_chainId *big.Int, _batchNumber *big.Int, _index *big.Int, _message L2Message, _proof [][32]byte) (bool, error) {
	return _IBridgehub.Contract.ProveL2MessageInclusion(&_IBridgehub.CallOpts, _chainId, _batchNumber, _index, _message, _proof)
}

// SharedBridge is a free data retrieval call binding the contract method 0x38720778.
//
// Solidity: function sharedBridge() view returns(address)
func (_IBridgehub *IBridgehubCaller) SharedBridge(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _IBridgehub.contract.Call(opts, &out, "sharedBridge")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// SharedBridge is a free data retrieval call binding the contract method 0x38720778.
//
// Solidity: function sharedBridge() view returns(address)
func (_IBridgehub *IBridgehubSession) SharedBridge() (common.Address, error) {
	return _IBridgehub.Contract.SharedBridge(&_IBridgehub.CallOpts)
}

// SharedBridge is a free data retrieval call binding the contract method 0x38720778.
//
// Solidity: function sharedBridge() view returns(address)
func (_IBridgehub *IBridgehubCallerSession) SharedBridge() (common.Address, error) {
	return _IBridgehub.Contract.SharedBridge(&_IBridgehub.CallOpts)
}

// StateTransitionManager is a free data retrieval call binding the contract method 0x402efc91.
//
// Solidity: function stateTransitionManager(uint256 _chainId) view returns(address)
func (_IBridgehub *IBridgehubCaller) StateTransitionManager(opts *bind.CallOpts, _chainId *big.Int) (common.Address, error) {
	var out []interface{}
	err := _IBridgehub.contract.Call(opts, &out, "stateTransitionManager", _chainId)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// StateTransitionManager is a free data retrieval call binding the contract method 0x402efc91.
//
// Solidity: function stateTransitionManager(uint256 _chainId) view returns(address)
func (_IBridgehub *IBridgehubSession) StateTransitionManager(_chainId *big.Int) (common.Address, error) {
	return _IBridgehub.Contract.StateTransitionManager(&_IBridgehub.CallOpts, _chainId)
}

// StateTransitionManager is a free data retrieval call binding the contract method 0x402efc91.
//
// Solidity: function stateTransitionManager(uint256 _chainId) view returns(address)
func (_IBridgehub *IBridgehubCallerSession) StateTransitionManager(_chainId *big.Int) (common.Address, error) {
	return _IBridgehub.Contract.StateTransitionManager(&_IBridgehub.CallOpts, _chainId)
}

// StateTransitionManagerIsRegistered is a free data retrieval call binding the contract method 0xbb7044b6.
//
// Solidity: function stateTransitionManagerIsRegistered(address _stateTransitionManager) view returns(bool)
func (_IBridgehub *IBridgehubCaller) StateTransitionManagerIsRegistered(opts *bind.CallOpts, _stateTransitionManager common.Address) (bool, error) {
	var out []interface{}
	err := _IBridgehub.contract.Call(opts, &out, "stateTransitionManagerIsRegistered", _stateTransitionManager)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// StateTransitionManagerIsRegistered is a free data retrieval call binding the contract method 0xbb7044b6.
//
// Solidity: function stateTransitionManagerIsRegistered(address _stateTransitionManager) view returns(bool)
func (_IBridgehub *IBridgehubSession) StateTransitionManagerIsRegistered(_stateTransitionManager common.Address) (bool, error) {
	return _IBridgehub.Contract.StateTransitionManagerIsRegistered(&_IBridgehub.CallOpts, _stateTransitionManager)
}

// StateTransitionManagerIsRegistered is a free data retrieval call binding the contract method 0xbb7044b6.
//
// Solidity: function stateTransitionManagerIsRegistered(address _stateTransitionManager) view returns(bool)
func (_IBridgehub *IBridgehubCallerSession) StateTransitionManagerIsRegistered(_stateTransitionManager common.Address) (bool, error) {
	return _IBridgehub.Contract.StateTransitionManagerIsRegistered(&_IBridgehub.CallOpts, _stateTransitionManager)
}

// TokenIsRegistered is a free data retrieval call binding the contract method 0x805b5b74.
//
// Solidity: function tokenIsRegistered(address _baseToken) view returns(bool)
func (_IBridgehub *IBridgehubCaller) TokenIsRegistered(opts *bind.CallOpts, _baseToken common.Address) (bool, error) {
	var out []interface{}
	err := _IBridgehub.contract.Call(opts, &out, "tokenIsRegistered", _baseToken)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// TokenIsRegistered is a free data retrieval call binding the contract method 0x805b5b74.
//
// Solidity: function tokenIsRegistered(address _baseToken) view returns(bool)
func (_IBridgehub *IBridgehubSession) TokenIsRegistered(_baseToken common.Address) (bool, error) {
	return _IBridgehub.Contract.TokenIsRegistered(&_IBridgehub.CallOpts, _baseToken)
}

// TokenIsRegistered is a free data retrieval call binding the contract method 0x805b5b74.
//
// Solidity: function tokenIsRegistered(address _baseToken) view returns(bool)
func (_IBridgehub *IBridgehubCallerSession) TokenIsRegistered(_baseToken common.Address) (bool, error) {
	return _IBridgehub.Contract.TokenIsRegistered(&_IBridgehub.CallOpts, _baseToken)
}

// AcceptAdmin is a paid mutator transaction binding the contract method 0x0e18b681.
//
// Solidity: function acceptAdmin() returns()
func (_IBridgehub *IBridgehubTransactor) AcceptAdmin(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IBridgehub.contract.Transact(opts, "acceptAdmin")
}

// AcceptAdmin is a paid mutator transaction binding the contract method 0x0e18b681.
//
// Solidity: function acceptAdmin() returns()
func (_IBridgehub *IBridgehubSession) AcceptAdmin() (*types.Transaction, error) {
	return _IBridgehub.Contract.AcceptAdmin(&_IBridgehub.TransactOpts)
}

// AcceptAdmin is a paid mutator transaction binding the contract method 0x0e18b681.
//
// Solidity: function acceptAdmin() returns()
func (_IBridgehub *IBridgehubTransactorSession) AcceptAdmin() (*types.Transaction, error) {
	return _IBridgehub.Contract.AcceptAdmin(&_IBridgehub.TransactOpts)
}

// AddStateTransitionManager is a paid mutator transaction binding the contract method 0x74044673.
//
// Solidity: function addStateTransitionManager(address _stateTransitionManager) returns()
func (_IBridgehub *IBridgehubTransactor) AddStateTransitionManager(opts *bind.TransactOpts, _stateTransitionManager common.Address) (*types.Transaction, error) {
	return _IBridgehub.contract.Transact(opts, "addStateTransitionManager", _stateTransitionManager)
}

// AddStateTransitionManager is a paid mutator transaction binding the contract method 0x74044673.
//
// Solidity: function addStateTransitionManager(address _stateTransitionManager) returns()
func (_IBridgehub *IBridgehubSession) AddStateTransitionManager(_stateTransitionManager common.Address) (*types.Transaction, error) {
	return _IBridgehub.Contract.AddStateTransitionManager(&_IBridgehub.TransactOpts, _stateTransitionManager)
}

// AddStateTransitionManager is a paid mutator transaction binding the contract method 0x74044673.
//
// Solidity: function addStateTransitionManager(address _stateTransitionManager) returns()
func (_IBridgehub *IBridgehubTransactorSession) AddStateTransitionManager(_stateTransitionManager common.Address) (*types.Transaction, error) {
	return _IBridgehub.Contract.AddStateTransitionManager(&_IBridgehub.TransactOpts, _stateTransitionManager)
}

// AddToken is a paid mutator transaction binding the contract method 0xd48bfca7.
//
// Solidity: function addToken(address _token) returns()
func (_IBridgehub *IBridgehubTransactor) AddToken(opts *bind.TransactOpts, _token common.Address) (*types.Transaction, error) {
	return _IBridgehub.contract.Transact(opts, "addToken", _token)
}

// AddToken is a paid mutator transaction binding the contract method 0xd48bfca7.
//
// Solidity: function addToken(address _token) returns()
func (_IBridgehub *IBridgehubSession) AddToken(_token common.Address) (*types.Transaction, error) {
	return _IBridgehub.Contract.AddToken(&_IBridgehub.TransactOpts, _token)
}

// AddToken is a paid mutator transaction binding the contract method 0xd48bfca7.
//
// Solidity: function addToken(address _token) returns()
func (_IBridgehub *IBridgehubTransactorSession) AddToken(_token common.Address) (*types.Transaction, error) {
	return _IBridgehub.Contract.AddToken(&_IBridgehub.TransactOpts, _token)
}

// CreateNewChain is a paid mutator transaction binding the contract method 0x3f58f5b5.
//
// Solidity: function createNewChain(uint256 _chainId, address _stateTransitionManager, address _baseToken, uint256 _salt, address _admin, bytes _initData) returns(uint256 chainId)
func (_IBridgehub *IBridgehubTransactor) CreateNewChain(opts *bind.TransactOpts, _chainId *big.Int, _stateTransitionManager common.Address, _baseToken common.Address, _salt *big.Int, _admin common.Address, _initData []byte) (*types.Transaction, error) {
	return _IBridgehub.contract.Transact(opts, "createNewChain", _chainId, _stateTransitionManager, _baseToken, _salt, _admin, _initData)
}

// CreateNewChain is a paid mutator transaction binding the contract method 0x3f58f5b5.
//
// Solidity: function createNewChain(uint256 _chainId, address _stateTransitionManager, address _baseToken, uint256 _salt, address _admin, bytes _initData) returns(uint256 chainId)
func (_IBridgehub *IBridgehubSession) CreateNewChain(_chainId *big.Int, _stateTransitionManager common.Address, _baseToken common.Address, _salt *big.Int, _admin common.Address, _initData []byte) (*types.Transaction, error) {
	return _IBridgehub.Contract.CreateNewChain(&_IBridgehub.TransactOpts, _chainId, _stateTransitionManager, _baseToken, _salt, _admin, _initData)
}

// CreateNewChain is a paid mutator transaction binding the contract method 0x3f58f5b5.
//
// Solidity: function createNewChain(uint256 _chainId, address _stateTransitionManager, address _baseToken, uint256 _salt, address _admin, bytes _initData) returns(uint256 chainId)
func (_IBridgehub *IBridgehubTransactorSession) CreateNewChain(_chainId *big.Int, _stateTransitionManager common.Address, _baseToken common.Address, _salt *big.Int, _admin common.Address, _initData []byte) (*types.Transaction, error) {
	return _IBridgehub.Contract.CreateNewChain(&_IBridgehub.TransactOpts, _chainId, _stateTransitionManager, _baseToken, _salt, _admin, _initData)
}

// RemoveStateTransitionManager is a paid mutator transaction binding the contract method 0xf5ba4232.
//
// Solidity: function removeStateTransitionManager(address _stateTransitionManager) returns()
func (_IBridgehub *IBridgehubTransactor) RemoveStateTransitionManager(opts *bind.TransactOpts, _stateTransitionManager common.Address) (*types.Transaction, error) {
	return _IBridgehub.contract.Transact(opts, "removeStateTransitionManager", _stateTransitionManager)
}

// RemoveStateTransitionManager is a paid mutator transaction binding the contract method 0xf5ba4232.
//
// Solidity: function removeStateTransitionManager(address _stateTransitionManager) returns()
func (_IBridgehub *IBridgehubSession) RemoveStateTransitionManager(_stateTransitionManager common.Address) (*types.Transaction, error) {
	return _IBridgehub.Contract.RemoveStateTransitionManager(&_IBridgehub.TransactOpts, _stateTransitionManager)
}

// RemoveStateTransitionManager is a paid mutator transaction binding the contract method 0xf5ba4232.
//
// Solidity: function removeStateTransitionManager(address _stateTransitionManager) returns()
func (_IBridgehub *IBridgehubTransactorSession) RemoveStateTransitionManager(_stateTransitionManager common.Address) (*types.Transaction, error) {
	return _IBridgehub.Contract.RemoveStateTransitionManager(&_IBridgehub.TransactOpts, _stateTransitionManager)
}

// RequestL2TransactionDirect is a paid mutator transaction binding the contract method 0xd52471c1.
//
// Solidity: function requestL2TransactionDirect((uint256,uint256,address,uint256,bytes,uint256,uint256,bytes[],address) _request) payable returns(bytes32 canonicalTxHash)
func (_IBridgehub *IBridgehubTransactor) RequestL2TransactionDirect(opts *bind.TransactOpts, _request L2TransactionRequestDirect) (*types.Transaction, error) {
	return _IBridgehub.contract.Transact(opts, "requestL2TransactionDirect", _request)
}

// RequestL2TransactionDirect is a paid mutator transaction binding the contract method 0xd52471c1.
//
// Solidity: function requestL2TransactionDirect((uint256,uint256,address,uint256,bytes,uint256,uint256,bytes[],address) _request) payable returns(bytes32 canonicalTxHash)
func (_IBridgehub *IBridgehubSession) RequestL2TransactionDirect(_request L2TransactionRequestDirect) (*types.Transaction, error) {
	return _IBridgehub.Contract.RequestL2TransactionDirect(&_IBridgehub.TransactOpts, _request)
}

// RequestL2TransactionDirect is a paid mutator transaction binding the contract method 0xd52471c1.
//
// Solidity: function requestL2TransactionDirect((uint256,uint256,address,uint256,bytes,uint256,uint256,bytes[],address) _request) payable returns(bytes32 canonicalTxHash)
func (_IBridgehub *IBridgehubTransactorSession) RequestL2TransactionDirect(_request L2TransactionRequestDirect) (*types.Transaction, error) {
	return _IBridgehub.Contract.RequestL2TransactionDirect(&_IBridgehub.TransactOpts, _request)
}

// RequestL2TransactionTwoBridges is a paid mutator transaction binding the contract method 0x24fd57fb.
//
// Solidity: function requestL2TransactionTwoBridges((uint256,uint256,uint256,uint256,uint256,address,address,uint256,bytes) _request) payable returns(bytes32 canonicalTxHash)
func (_IBridgehub *IBridgehubTransactor) RequestL2TransactionTwoBridges(opts *bind.TransactOpts, _request L2TransactionRequestTwoBridgesOuter) (*types.Transaction, error) {
	return _IBridgehub.contract.Transact(opts, "requestL2TransactionTwoBridges", _request)
}

// RequestL2TransactionTwoBridges is a paid mutator transaction binding the contract method 0x24fd57fb.
//
// Solidity: function requestL2TransactionTwoBridges((uint256,uint256,uint256,uint256,uint256,address,address,uint256,bytes) _request) payable returns(bytes32 canonicalTxHash)
func (_IBridgehub *IBridgehubSession) RequestL2TransactionTwoBridges(_request L2TransactionRequestTwoBridgesOuter) (*types.Transaction, error) {
	return _IBridgehub.Contract.RequestL2TransactionTwoBridges(&_IBridgehub.TransactOpts, _request)
}

// RequestL2TransactionTwoBridges is a paid mutator transaction binding the contract method 0x24fd57fb.
//
// Solidity: function requestL2TransactionTwoBridges((uint256,uint256,uint256,uint256,uint256,address,address,uint256,bytes) _request) payable returns(bytes32 canonicalTxHash)
func (_IBridgehub *IBridgehubTransactorSession) RequestL2TransactionTwoBridges(_request L2TransactionRequestTwoBridgesOuter) (*types.Transaction, error) {
	return _IBridgehub.Contract.RequestL2TransactionTwoBridges(&_IBridgehub.TransactOpts, _request)
}

// SetPendingAdmin is a paid mutator transaction binding the contract method 0x4dd18bf5.
//
// Solidity: function setPendingAdmin(address _newPendingAdmin) returns()
func (_IBridgehub *IBridgehubTransactor) SetPendingAdmin(opts *bind.TransactOpts, _newPendingAdmin common.Address) (*types.Transaction, error) {
	return _IBridgehub.contract.Transact(opts, "setPendingAdmin", _newPendingAdmin)
}

// SetPendingAdmin is a paid mutator transaction binding the contract method 0x4dd18bf5.
//
// Solidity: function setPendingAdmin(address _newPendingAdmin) returns()
func (_IBridgehub *IBridgehubSession) SetPendingAdmin(_newPendingAdmin common.Address) (*types.Transaction, error) {
	return _IBridgehub.Contract.SetPendingAdmin(&_IBridgehub.TransactOpts, _newPendingAdmin)
}

// SetPendingAdmin is a paid mutator transaction binding the contract method 0x4dd18bf5.
//
// Solidity: function setPendingAdmin(address _newPendingAdmin) returns()
func (_IBridgehub *IBridgehubTransactorSession) SetPendingAdmin(_newPendingAdmin common.Address) (*types.Transaction, error) {
	return _IBridgehub.Contract.SetPendingAdmin(&_IBridgehub.TransactOpts, _newPendingAdmin)
}

// SetSharedBridge is a paid mutator transaction binding the contract method 0xd0bf6fd4.
//
// Solidity: function setSharedBridge(address _sharedBridge) returns()
func (_IBridgehub *IBridgehubTransactor) SetSharedBridge(opts *bind.TransactOpts, _sharedBridge common.Address) (*types.Transaction, error) {
	return _IBridgehub.contract.Transact(opts, "setSharedBridge", _sharedBridge)
}

// SetSharedBridge is a paid mutator transaction binding the contract method 0xd0bf6fd4.
//
// Solidity: function setSharedBridge(address _sharedBridge) returns()
func (_IBridgehub *IBridgehubSession) SetSharedBridge(_sharedBridge common.Address) (*types.Transaction, error) {
	return _IBridgehub.Contract.SetSharedBridge(&_IBridgehub.TransactOpts, _sharedBridge)
}

// SetSharedBridge is a paid mutator transaction binding the contract method 0xd0bf6fd4.
//
// Solidity: function setSharedBridge(address _sharedBridge) returns()
func (_IBridgehub *IBridgehubTransactorSession) SetSharedBridge(_sharedBridge common.Address) (*types.Transaction, error) {
	return _IBridgehub.Contract.SetSharedBridge(&_IBridgehub.TransactOpts, _sharedBridge)
}

// IBridgehubNewAdminIterator is returned from FilterNewAdmin and is used to iterate over the raw logs and unpacked data for NewAdmin events raised by the IBridgehub contract.
type IBridgehubNewAdminIterator struct {
	Event *IBridgehubNewAdmin // Event containing the contract specifics and raw log

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
func (it *IBridgehubNewAdminIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IBridgehubNewAdmin)
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
		it.Event = new(IBridgehubNewAdmin)
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
func (it *IBridgehubNewAdminIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IBridgehubNewAdminIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IBridgehubNewAdmin represents a NewAdmin event raised by the IBridgehub contract.
type IBridgehubNewAdmin struct {
	OldAdmin common.Address
	NewAdmin common.Address
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterNewAdmin is a free log retrieval operation binding the contract event 0xf9ffabca9c8276e99321725bcb43fb076a6c66a54b7f21c4e8146d8519b417dc.
//
// Solidity: event NewAdmin(address indexed oldAdmin, address indexed newAdmin)
func (_IBridgehub *IBridgehubFilterer) FilterNewAdmin(opts *bind.FilterOpts, oldAdmin []common.Address, newAdmin []common.Address) (*IBridgehubNewAdminIterator, error) {

	var oldAdminRule []interface{}
	for _, oldAdminItem := range oldAdmin {
		oldAdminRule = append(oldAdminRule, oldAdminItem)
	}
	var newAdminRule []interface{}
	for _, newAdminItem := range newAdmin {
		newAdminRule = append(newAdminRule, newAdminItem)
	}

	logs, sub, err := _IBridgehub.contract.FilterLogs(opts, "NewAdmin", oldAdminRule, newAdminRule)
	if err != nil {
		return nil, err
	}
	return &IBridgehubNewAdminIterator{contract: _IBridgehub.contract, event: "NewAdmin", logs: logs, sub: sub}, nil
}

// WatchNewAdmin is a free log subscription operation binding the contract event 0xf9ffabca9c8276e99321725bcb43fb076a6c66a54b7f21c4e8146d8519b417dc.
//
// Solidity: event NewAdmin(address indexed oldAdmin, address indexed newAdmin)
func (_IBridgehub *IBridgehubFilterer) WatchNewAdmin(opts *bind.WatchOpts, sink chan<- *IBridgehubNewAdmin, oldAdmin []common.Address, newAdmin []common.Address) (event.Subscription, error) {

	var oldAdminRule []interface{}
	for _, oldAdminItem := range oldAdmin {
		oldAdminRule = append(oldAdminRule, oldAdminItem)
	}
	var newAdminRule []interface{}
	for _, newAdminItem := range newAdmin {
		newAdminRule = append(newAdminRule, newAdminItem)
	}

	logs, sub, err := _IBridgehub.contract.WatchLogs(opts, "NewAdmin", oldAdminRule, newAdminRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IBridgehubNewAdmin)
				if err := _IBridgehub.contract.UnpackLog(event, "NewAdmin", log); err != nil {
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

// ParseNewAdmin is a log parse operation binding the contract event 0xf9ffabca9c8276e99321725bcb43fb076a6c66a54b7f21c4e8146d8519b417dc.
//
// Solidity: event NewAdmin(address indexed oldAdmin, address indexed newAdmin)
func (_IBridgehub *IBridgehubFilterer) ParseNewAdmin(log types.Log) (*IBridgehubNewAdmin, error) {
	event := new(IBridgehubNewAdmin)
	if err := _IBridgehub.contract.UnpackLog(event, "NewAdmin", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// IBridgehubNewChainIterator is returned from FilterNewChain and is used to iterate over the raw logs and unpacked data for NewChain events raised by the IBridgehub contract.
type IBridgehubNewChainIterator struct {
	Event *IBridgehubNewChain // Event containing the contract specifics and raw log

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
func (it *IBridgehubNewChainIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IBridgehubNewChain)
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
		it.Event = new(IBridgehubNewChain)
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
func (it *IBridgehubNewChainIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IBridgehubNewChainIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IBridgehubNewChain represents a NewChain event raised by the IBridgehub contract.
type IBridgehubNewChain struct {
	ChainId                *big.Int
	StateTransitionManager common.Address
	ChainGovernance        common.Address
	Raw                    types.Log // Blockchain specific contextual infos
}

// FilterNewChain is a free log retrieval operation binding the contract event 0x1e9125bc72db22c58abff6821d7333551967e26454b419ffa958e4cb8ef47600.
//
// Solidity: event NewChain(uint256 indexed chainId, address stateTransitionManager, address indexed chainGovernance)
func (_IBridgehub *IBridgehubFilterer) FilterNewChain(opts *bind.FilterOpts, chainId []*big.Int, chainGovernance []common.Address) (*IBridgehubNewChainIterator, error) {

	var chainIdRule []interface{}
	for _, chainIdItem := range chainId {
		chainIdRule = append(chainIdRule, chainIdItem)
	}

	var chainGovernanceRule []interface{}
	for _, chainGovernanceItem := range chainGovernance {
		chainGovernanceRule = append(chainGovernanceRule, chainGovernanceItem)
	}

	logs, sub, err := _IBridgehub.contract.FilterLogs(opts, "NewChain", chainIdRule, chainGovernanceRule)
	if err != nil {
		return nil, err
	}
	return &IBridgehubNewChainIterator{contract: _IBridgehub.contract, event: "NewChain", logs: logs, sub: sub}, nil
}

// WatchNewChain is a free log subscription operation binding the contract event 0x1e9125bc72db22c58abff6821d7333551967e26454b419ffa958e4cb8ef47600.
//
// Solidity: event NewChain(uint256 indexed chainId, address stateTransitionManager, address indexed chainGovernance)
func (_IBridgehub *IBridgehubFilterer) WatchNewChain(opts *bind.WatchOpts, sink chan<- *IBridgehubNewChain, chainId []*big.Int, chainGovernance []common.Address) (event.Subscription, error) {

	var chainIdRule []interface{}
	for _, chainIdItem := range chainId {
		chainIdRule = append(chainIdRule, chainIdItem)
	}

	var chainGovernanceRule []interface{}
	for _, chainGovernanceItem := range chainGovernance {
		chainGovernanceRule = append(chainGovernanceRule, chainGovernanceItem)
	}

	logs, sub, err := _IBridgehub.contract.WatchLogs(opts, "NewChain", chainIdRule, chainGovernanceRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IBridgehubNewChain)
				if err := _IBridgehub.contract.UnpackLog(event, "NewChain", log); err != nil {
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

// ParseNewChain is a log parse operation binding the contract event 0x1e9125bc72db22c58abff6821d7333551967e26454b419ffa958e4cb8ef47600.
//
// Solidity: event NewChain(uint256 indexed chainId, address stateTransitionManager, address indexed chainGovernance)
func (_IBridgehub *IBridgehubFilterer) ParseNewChain(log types.Log) (*IBridgehubNewChain, error) {
	event := new(IBridgehubNewChain)
	if err := _IBridgehub.contract.UnpackLog(event, "NewChain", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// IBridgehubNewPendingAdminIterator is returned from FilterNewPendingAdmin and is used to iterate over the raw logs and unpacked data for NewPendingAdmin events raised by the IBridgehub contract.
type IBridgehubNewPendingAdminIterator struct {
	Event *IBridgehubNewPendingAdmin // Event containing the contract specifics and raw log

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
func (it *IBridgehubNewPendingAdminIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IBridgehubNewPendingAdmin)
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
		it.Event = new(IBridgehubNewPendingAdmin)
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
func (it *IBridgehubNewPendingAdminIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IBridgehubNewPendingAdminIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IBridgehubNewPendingAdmin represents a NewPendingAdmin event raised by the IBridgehub contract.
type IBridgehubNewPendingAdmin struct {
	OldPendingAdmin common.Address
	NewPendingAdmin common.Address
	Raw             types.Log // Blockchain specific contextual infos
}

// FilterNewPendingAdmin is a free log retrieval operation binding the contract event 0xca4f2f25d0898edd99413412fb94012f9e54ec8142f9b093e7720646a95b16a9.
//
// Solidity: event NewPendingAdmin(address indexed oldPendingAdmin, address indexed newPendingAdmin)
func (_IBridgehub *IBridgehubFilterer) FilterNewPendingAdmin(opts *bind.FilterOpts, oldPendingAdmin []common.Address, newPendingAdmin []common.Address) (*IBridgehubNewPendingAdminIterator, error) {

	var oldPendingAdminRule []interface{}
	for _, oldPendingAdminItem := range oldPendingAdmin {
		oldPendingAdminRule = append(oldPendingAdminRule, oldPendingAdminItem)
	}
	var newPendingAdminRule []interface{}
	for _, newPendingAdminItem := range newPendingAdmin {
		newPendingAdminRule = append(newPendingAdminRule, newPendingAdminItem)
	}

	logs, sub, err := _IBridgehub.contract.FilterLogs(opts, "NewPendingAdmin", oldPendingAdminRule, newPendingAdminRule)
	if err != nil {
		return nil, err
	}
	return &IBridgehubNewPendingAdminIterator{contract: _IBridgehub.contract, event: "NewPendingAdmin", logs: logs, sub: sub}, nil
}

// WatchNewPendingAdmin is a free log subscription operation binding the contract event 0xca4f2f25d0898edd99413412fb94012f9e54ec8142f9b093e7720646a95b16a9.
//
// Solidity: event NewPendingAdmin(address indexed oldPendingAdmin, address indexed newPendingAdmin)
func (_IBridgehub *IBridgehubFilterer) WatchNewPendingAdmin(opts *bind.WatchOpts, sink chan<- *IBridgehubNewPendingAdmin, oldPendingAdmin []common.Address, newPendingAdmin []common.Address) (event.Subscription, error) {

	var oldPendingAdminRule []interface{}
	for _, oldPendingAdminItem := range oldPendingAdmin {
		oldPendingAdminRule = append(oldPendingAdminRule, oldPendingAdminItem)
	}
	var newPendingAdminRule []interface{}
	for _, newPendingAdminItem := range newPendingAdmin {
		newPendingAdminRule = append(newPendingAdminRule, newPendingAdminItem)
	}

	logs, sub, err := _IBridgehub.contract.WatchLogs(opts, "NewPendingAdmin", oldPendingAdminRule, newPendingAdminRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IBridgehubNewPendingAdmin)
				if err := _IBridgehub.contract.UnpackLog(event, "NewPendingAdmin", log); err != nil {
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

// ParseNewPendingAdmin is a log parse operation binding the contract event 0xca4f2f25d0898edd99413412fb94012f9e54ec8142f9b093e7720646a95b16a9.
//
// Solidity: event NewPendingAdmin(address indexed oldPendingAdmin, address indexed newPendingAdmin)
func (_IBridgehub *IBridgehubFilterer) ParseNewPendingAdmin(log types.Log) (*IBridgehubNewPendingAdmin, error) {
	event := new(IBridgehubNewPendingAdmin)
	if err := _IBridgehub.contract.UnpackLog(event, "NewPendingAdmin", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
