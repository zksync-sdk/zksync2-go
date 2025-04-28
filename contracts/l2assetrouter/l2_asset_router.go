// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package l2assetrouter

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

// IL2AssetRouterMetaData contains all meta data concerning the IL2AssetRouter contract.
var IL2AssetRouterMetaData = &bind.MetaData{
	ABI: "[{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"assetId\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"_assetAddress\",\"type\":\"address\"}],\"name\":\"AssetHandlerRegistered\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"assetId\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"assetHandlerAddress\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"additionalData\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"assetDeploymentTracker\",\"type\":\"address\"}],\"name\":\"AssetHandlerRegisteredInitial\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"chainId\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"assetId\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"BridgehubDepositBaseTokenInitiated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"chainId\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"txDataHash\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"assetId\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"bridgeMintCalldata\",\"type\":\"bytes\"}],\"name\":\"BridgehubDepositInitiated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"chainId\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"assetId\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"assetDataHash\",\"type\":\"bytes32\"}],\"name\":\"BridgehubWithdrawalInitiated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"chainId\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"assetId\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"assetData\",\"type\":\"bytes\"}],\"name\":\"DepositFinalizedAssetRouter\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"chainId\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"l2Sender\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"assetId\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"assetData\",\"type\":\"bytes\"}],\"name\":\"WithdrawalInitiatedAssetRouter\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"BRIDGE_HUB\",\"outputs\":[{\"internalType\":\"contractIBridgehub\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_assetId\",\"type\":\"bytes32\"}],\"name\":\"assetHandlerAddress\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_chainId\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"_assetId\",\"type\":\"bytes32\"},{\"internalType\":\"bytes\",\"name\":\"_transferData\",\"type\":\"bytes\"}],\"name\":\"finalizeDeposit\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"l1AssetRouter\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_originChainId\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"_assetId\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"_assetAddress\",\"type\":\"address\"}],\"name\":\"setAssetHandlerAddress\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_assetRegistrationData\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"_assetHandlerAddress\",\"type\":\"address\"}],\"name\":\"setAssetHandlerAddressThisChain\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_assetId\",\"type\":\"bytes32\"},{\"internalType\":\"bytes\",\"name\":\"_transferData\",\"type\":\"bytes\"}],\"name\":\"withdraw\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_l1Receiver\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_l2Token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"_sender\",\"type\":\"address\"}],\"name\":\"withdrawLegacyBridge\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
}

// IL2AssetRouterABI is the input ABI used to generate the binding from.
// Deprecated: Use IL2AssetRouterMetaData.ABI instead.
var IL2AssetRouterABI = IL2AssetRouterMetaData.ABI

// IL2AssetRouter is an auto generated Go binding around an Ethereum contract.
type IL2AssetRouter struct {
	IL2AssetRouterCaller     // Read-only binding to the contract
	IL2AssetRouterTransactor // Write-only binding to the contract
	IL2AssetRouterFilterer   // Log filterer for contract events
}

// IL2AssetRouterCaller is an auto generated read-only Go binding around an Ethereum contract.
type IL2AssetRouterCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IL2AssetRouterTransactor is an auto generated write-only Go binding around an Ethereum contract.
type IL2AssetRouterTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IL2AssetRouterFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type IL2AssetRouterFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IL2AssetRouterSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type IL2AssetRouterSession struct {
	Contract     *IL2AssetRouter   // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// IL2AssetRouterCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type IL2AssetRouterCallerSession struct {
	Contract *IL2AssetRouterCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts         // Call options to use throughout this session
}

// IL2AssetRouterTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type IL2AssetRouterTransactorSession struct {
	Contract     *IL2AssetRouterTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts         // Transaction auth options to use throughout this session
}

// IL2AssetRouterRaw is an auto generated low-level Go binding around an Ethereum contract.
type IL2AssetRouterRaw struct {
	Contract *IL2AssetRouter // Generic contract binding to access the raw methods on
}

// IL2AssetRouterCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type IL2AssetRouterCallerRaw struct {
	Contract *IL2AssetRouterCaller // Generic read-only contract binding to access the raw methods on
}

// IL2AssetRouterTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type IL2AssetRouterTransactorRaw struct {
	Contract *IL2AssetRouterTransactor // Generic write-only contract binding to access the raw methods on
}

// NewIL2AssetRouter creates a new instance of IL2AssetRouter, bound to a specific deployed contract.
func NewIL2AssetRouter(address common.Address, backend bind.ContractBackend) (*IL2AssetRouter, error) {
	contract, err := bindIL2AssetRouter(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &IL2AssetRouter{IL2AssetRouterCaller: IL2AssetRouterCaller{contract: contract}, IL2AssetRouterTransactor: IL2AssetRouterTransactor{contract: contract}, IL2AssetRouterFilterer: IL2AssetRouterFilterer{contract: contract}}, nil
}

// NewIL2AssetRouterCaller creates a new read-only instance of IL2AssetRouter, bound to a specific deployed contract.
func NewIL2AssetRouterCaller(address common.Address, caller bind.ContractCaller) (*IL2AssetRouterCaller, error) {
	contract, err := bindIL2AssetRouter(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &IL2AssetRouterCaller{contract: contract}, nil
}

// NewIL2AssetRouterTransactor creates a new write-only instance of IL2AssetRouter, bound to a specific deployed contract.
func NewIL2AssetRouterTransactor(address common.Address, transactor bind.ContractTransactor) (*IL2AssetRouterTransactor, error) {
	contract, err := bindIL2AssetRouter(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &IL2AssetRouterTransactor{contract: contract}, nil
}

// NewIL2AssetRouterFilterer creates a new log filterer instance of IL2AssetRouter, bound to a specific deployed contract.
func NewIL2AssetRouterFilterer(address common.Address, filterer bind.ContractFilterer) (*IL2AssetRouterFilterer, error) {
	contract, err := bindIL2AssetRouter(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &IL2AssetRouterFilterer{contract: contract}, nil
}

// bindIL2AssetRouter binds a generic wrapper to an already deployed contract.
func bindIL2AssetRouter(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := IL2AssetRouterMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IL2AssetRouter *IL2AssetRouterRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IL2AssetRouter.Contract.IL2AssetRouterCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IL2AssetRouter *IL2AssetRouterRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IL2AssetRouter.Contract.IL2AssetRouterTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IL2AssetRouter *IL2AssetRouterRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IL2AssetRouter.Contract.IL2AssetRouterTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IL2AssetRouter *IL2AssetRouterCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IL2AssetRouter.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IL2AssetRouter *IL2AssetRouterTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IL2AssetRouter.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IL2AssetRouter *IL2AssetRouterTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IL2AssetRouter.Contract.contract.Transact(opts, method, params...)
}

// BRIDGEHUB is a free data retrieval call binding the contract method 0x5d4edca7.
//
// Solidity: function BRIDGE_HUB() view returns(address)
func (_IL2AssetRouter *IL2AssetRouterCaller) BRIDGEHUB(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _IL2AssetRouter.contract.Call(opts, &out, "BRIDGE_HUB")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// BRIDGEHUB is a free data retrieval call binding the contract method 0x5d4edca7.
//
// Solidity: function BRIDGE_HUB() view returns(address)
func (_IL2AssetRouter *IL2AssetRouterSession) BRIDGEHUB() (common.Address, error) {
	return _IL2AssetRouter.Contract.BRIDGEHUB(&_IL2AssetRouter.CallOpts)
}

// BRIDGEHUB is a free data retrieval call binding the contract method 0x5d4edca7.
//
// Solidity: function BRIDGE_HUB() view returns(address)
func (_IL2AssetRouter *IL2AssetRouterCallerSession) BRIDGEHUB() (common.Address, error) {
	return _IL2AssetRouter.Contract.BRIDGEHUB(&_IL2AssetRouter.CallOpts)
}

// AssetHandlerAddress is a free data retrieval call binding the contract method 0x53b9e632.
//
// Solidity: function assetHandlerAddress(bytes32 _assetId) view returns(address)
func (_IL2AssetRouter *IL2AssetRouterCaller) AssetHandlerAddress(opts *bind.CallOpts, _assetId [32]byte) (common.Address, error) {
	var out []interface{}
	err := _IL2AssetRouter.contract.Call(opts, &out, "assetHandlerAddress", _assetId)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// AssetHandlerAddress is a free data retrieval call binding the contract method 0x53b9e632.
//
// Solidity: function assetHandlerAddress(bytes32 _assetId) view returns(address)
func (_IL2AssetRouter *IL2AssetRouterSession) AssetHandlerAddress(_assetId [32]byte) (common.Address, error) {
	return _IL2AssetRouter.Contract.AssetHandlerAddress(&_IL2AssetRouter.CallOpts, _assetId)
}

// AssetHandlerAddress is a free data retrieval call binding the contract method 0x53b9e632.
//
// Solidity: function assetHandlerAddress(bytes32 _assetId) view returns(address)
func (_IL2AssetRouter *IL2AssetRouterCallerSession) AssetHandlerAddress(_assetId [32]byte) (common.Address, error) {
	return _IL2AssetRouter.Contract.AssetHandlerAddress(&_IL2AssetRouter.CallOpts, _assetId)
}

// L1AssetRouter is a free data retrieval call binding the contract method 0x6d9860e1.
//
// Solidity: function l1AssetRouter() view returns(address)
func (_IL2AssetRouter *IL2AssetRouterCaller) L1AssetRouter(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _IL2AssetRouter.contract.Call(opts, &out, "l1AssetRouter")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// L1AssetRouter is a free data retrieval call binding the contract method 0x6d9860e1.
//
// Solidity: function l1AssetRouter() view returns(address)
func (_IL2AssetRouter *IL2AssetRouterSession) L1AssetRouter() (common.Address, error) {
	return _IL2AssetRouter.Contract.L1AssetRouter(&_IL2AssetRouter.CallOpts)
}

// L1AssetRouter is a free data retrieval call binding the contract method 0x6d9860e1.
//
// Solidity: function l1AssetRouter() view returns(address)
func (_IL2AssetRouter *IL2AssetRouterCallerSession) L1AssetRouter() (common.Address, error) {
	return _IL2AssetRouter.Contract.L1AssetRouter(&_IL2AssetRouter.CallOpts)
}

// FinalizeDeposit is a paid mutator transaction binding the contract method 0x9c884fd1.
//
// Solidity: function finalizeDeposit(uint256 _chainId, bytes32 _assetId, bytes _transferData) returns()
func (_IL2AssetRouter *IL2AssetRouterTransactor) FinalizeDeposit(opts *bind.TransactOpts, _chainId *big.Int, _assetId [32]byte, _transferData []byte) (*types.Transaction, error) {
	return _IL2AssetRouter.contract.Transact(opts, "finalizeDeposit", _chainId, _assetId, _transferData)
}

// FinalizeDeposit is a paid mutator transaction binding the contract method 0x9c884fd1.
//
// Solidity: function finalizeDeposit(uint256 _chainId, bytes32 _assetId, bytes _transferData) returns()
func (_IL2AssetRouter *IL2AssetRouterSession) FinalizeDeposit(_chainId *big.Int, _assetId [32]byte, _transferData []byte) (*types.Transaction, error) {
	return _IL2AssetRouter.Contract.FinalizeDeposit(&_IL2AssetRouter.TransactOpts, _chainId, _assetId, _transferData)
}

// FinalizeDeposit is a paid mutator transaction binding the contract method 0x9c884fd1.
//
// Solidity: function finalizeDeposit(uint256 _chainId, bytes32 _assetId, bytes _transferData) returns()
func (_IL2AssetRouter *IL2AssetRouterTransactorSession) FinalizeDeposit(_chainId *big.Int, _assetId [32]byte, _transferData []byte) (*types.Transaction, error) {
	return _IL2AssetRouter.Contract.FinalizeDeposit(&_IL2AssetRouter.TransactOpts, _chainId, _assetId, _transferData)
}

// SetAssetHandlerAddress is a paid mutator transaction binding the contract method 0xda556bdc.
//
// Solidity: function setAssetHandlerAddress(uint256 _originChainId, bytes32 _assetId, address _assetAddress) returns()
func (_IL2AssetRouter *IL2AssetRouterTransactor) SetAssetHandlerAddress(opts *bind.TransactOpts, _originChainId *big.Int, _assetId [32]byte, _assetAddress common.Address) (*types.Transaction, error) {
	return _IL2AssetRouter.contract.Transact(opts, "setAssetHandlerAddress", _originChainId, _assetId, _assetAddress)
}

// SetAssetHandlerAddress is a paid mutator transaction binding the contract method 0xda556bdc.
//
// Solidity: function setAssetHandlerAddress(uint256 _originChainId, bytes32 _assetId, address _assetAddress) returns()
func (_IL2AssetRouter *IL2AssetRouterSession) SetAssetHandlerAddress(_originChainId *big.Int, _assetId [32]byte, _assetAddress common.Address) (*types.Transaction, error) {
	return _IL2AssetRouter.Contract.SetAssetHandlerAddress(&_IL2AssetRouter.TransactOpts, _originChainId, _assetId, _assetAddress)
}

// SetAssetHandlerAddress is a paid mutator transaction binding the contract method 0xda556bdc.
//
// Solidity: function setAssetHandlerAddress(uint256 _originChainId, bytes32 _assetId, address _assetAddress) returns()
func (_IL2AssetRouter *IL2AssetRouterTransactorSession) SetAssetHandlerAddress(_originChainId *big.Int, _assetId [32]byte, _assetAddress common.Address) (*types.Transaction, error) {
	return _IL2AssetRouter.Contract.SetAssetHandlerAddress(&_IL2AssetRouter.TransactOpts, _originChainId, _assetId, _assetAddress)
}

// SetAssetHandlerAddressThisChain is a paid mutator transaction binding the contract method 0x548a5a33.
//
// Solidity: function setAssetHandlerAddressThisChain(bytes32 _assetRegistrationData, address _assetHandlerAddress) returns()
func (_IL2AssetRouter *IL2AssetRouterTransactor) SetAssetHandlerAddressThisChain(opts *bind.TransactOpts, _assetRegistrationData [32]byte, _assetHandlerAddress common.Address) (*types.Transaction, error) {
	return _IL2AssetRouter.contract.Transact(opts, "setAssetHandlerAddressThisChain", _assetRegistrationData, _assetHandlerAddress)
}

// SetAssetHandlerAddressThisChain is a paid mutator transaction binding the contract method 0x548a5a33.
//
// Solidity: function setAssetHandlerAddressThisChain(bytes32 _assetRegistrationData, address _assetHandlerAddress) returns()
func (_IL2AssetRouter *IL2AssetRouterSession) SetAssetHandlerAddressThisChain(_assetRegistrationData [32]byte, _assetHandlerAddress common.Address) (*types.Transaction, error) {
	return _IL2AssetRouter.Contract.SetAssetHandlerAddressThisChain(&_IL2AssetRouter.TransactOpts, _assetRegistrationData, _assetHandlerAddress)
}

// SetAssetHandlerAddressThisChain is a paid mutator transaction binding the contract method 0x548a5a33.
//
// Solidity: function setAssetHandlerAddressThisChain(bytes32 _assetRegistrationData, address _assetHandlerAddress) returns()
func (_IL2AssetRouter *IL2AssetRouterTransactorSession) SetAssetHandlerAddressThisChain(_assetRegistrationData [32]byte, _assetHandlerAddress common.Address) (*types.Transaction, error) {
	return _IL2AssetRouter.Contract.SetAssetHandlerAddressThisChain(&_IL2AssetRouter.TransactOpts, _assetRegistrationData, _assetHandlerAddress)
}

// Withdraw is a paid mutator transaction binding the contract method 0x4a2e35ba.
//
// Solidity: function withdraw(bytes32 _assetId, bytes _transferData) returns()
func (_IL2AssetRouter *IL2AssetRouterTransactor) Withdraw(opts *bind.TransactOpts, _assetId [32]byte, _transferData []byte) (*types.Transaction, error) {
	return _IL2AssetRouter.contract.Transact(opts, "withdraw", _assetId, _transferData)
}

// Withdraw is a paid mutator transaction binding the contract method 0x4a2e35ba.
//
// Solidity: function withdraw(bytes32 _assetId, bytes _transferData) returns()
func (_IL2AssetRouter *IL2AssetRouterSession) Withdraw(_assetId [32]byte, _transferData []byte) (*types.Transaction, error) {
	return _IL2AssetRouter.Contract.Withdraw(&_IL2AssetRouter.TransactOpts, _assetId, _transferData)
}

// Withdraw is a paid mutator transaction binding the contract method 0x4a2e35ba.
//
// Solidity: function withdraw(bytes32 _assetId, bytes _transferData) returns()
func (_IL2AssetRouter *IL2AssetRouterTransactorSession) Withdraw(_assetId [32]byte, _transferData []byte) (*types.Transaction, error) {
	return _IL2AssetRouter.Contract.Withdraw(&_IL2AssetRouter.TransactOpts, _assetId, _transferData)
}

// WithdrawLegacyBridge is a paid mutator transaction binding the contract method 0x7ac3a553.
//
// Solidity: function withdrawLegacyBridge(address _l1Receiver, address _l2Token, uint256 _amount, address _sender) returns()
func (_IL2AssetRouter *IL2AssetRouterTransactor) WithdrawLegacyBridge(opts *bind.TransactOpts, _l1Receiver common.Address, _l2Token common.Address, _amount *big.Int, _sender common.Address) (*types.Transaction, error) {
	return _IL2AssetRouter.contract.Transact(opts, "withdrawLegacyBridge", _l1Receiver, _l2Token, _amount, _sender)
}

// WithdrawLegacyBridge is a paid mutator transaction binding the contract method 0x7ac3a553.
//
// Solidity: function withdrawLegacyBridge(address _l1Receiver, address _l2Token, uint256 _amount, address _sender) returns()
func (_IL2AssetRouter *IL2AssetRouterSession) WithdrawLegacyBridge(_l1Receiver common.Address, _l2Token common.Address, _amount *big.Int, _sender common.Address) (*types.Transaction, error) {
	return _IL2AssetRouter.Contract.WithdrawLegacyBridge(&_IL2AssetRouter.TransactOpts, _l1Receiver, _l2Token, _amount, _sender)
}

// WithdrawLegacyBridge is a paid mutator transaction binding the contract method 0x7ac3a553.
//
// Solidity: function withdrawLegacyBridge(address _l1Receiver, address _l2Token, uint256 _amount, address _sender) returns()
func (_IL2AssetRouter *IL2AssetRouterTransactorSession) WithdrawLegacyBridge(_l1Receiver common.Address, _l2Token common.Address, _amount *big.Int, _sender common.Address) (*types.Transaction, error) {
	return _IL2AssetRouter.Contract.WithdrawLegacyBridge(&_IL2AssetRouter.TransactOpts, _l1Receiver, _l2Token, _amount, _sender)
}

// IL2AssetRouterAssetHandlerRegisteredIterator is returned from FilterAssetHandlerRegistered and is used to iterate over the raw logs and unpacked data for AssetHandlerRegistered events raised by the IL2AssetRouter contract.
type IL2AssetRouterAssetHandlerRegisteredIterator struct {
	Event *IL2AssetRouterAssetHandlerRegistered // Event containing the contract specifics and raw log

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
func (it *IL2AssetRouterAssetHandlerRegisteredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IL2AssetRouterAssetHandlerRegistered)
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
		it.Event = new(IL2AssetRouterAssetHandlerRegistered)
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
func (it *IL2AssetRouterAssetHandlerRegisteredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IL2AssetRouterAssetHandlerRegisteredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IL2AssetRouterAssetHandlerRegistered represents a AssetHandlerRegistered event raised by the IL2AssetRouter contract.
type IL2AssetRouterAssetHandlerRegistered struct {
	AssetId      [32]byte
	AssetAddress common.Address
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterAssetHandlerRegistered is a free log retrieval operation binding the contract event 0x2632cc0d58b0cb1017b99cc0b6cc66ad86440cc0dd923bfdaa294f95ba1b0201.
//
// Solidity: event AssetHandlerRegistered(bytes32 indexed assetId, address indexed _assetAddress)
func (_IL2AssetRouter *IL2AssetRouterFilterer) FilterAssetHandlerRegistered(opts *bind.FilterOpts, assetId [][32]byte, _assetAddress []common.Address) (*IL2AssetRouterAssetHandlerRegisteredIterator, error) {

	var assetIdRule []interface{}
	for _, assetIdItem := range assetId {
		assetIdRule = append(assetIdRule, assetIdItem)
	}
	var _assetAddressRule []interface{}
	for _, _assetAddressItem := range _assetAddress {
		_assetAddressRule = append(_assetAddressRule, _assetAddressItem)
	}

	logs, sub, err := _IL2AssetRouter.contract.FilterLogs(opts, "AssetHandlerRegistered", assetIdRule, _assetAddressRule)
	if err != nil {
		return nil, err
	}
	return &IL2AssetRouterAssetHandlerRegisteredIterator{contract: _IL2AssetRouter.contract, event: "AssetHandlerRegistered", logs: logs, sub: sub}, nil
}

// WatchAssetHandlerRegistered is a free log subscription operation binding the contract event 0x2632cc0d58b0cb1017b99cc0b6cc66ad86440cc0dd923bfdaa294f95ba1b0201.
//
// Solidity: event AssetHandlerRegistered(bytes32 indexed assetId, address indexed _assetAddress)
func (_IL2AssetRouter *IL2AssetRouterFilterer) WatchAssetHandlerRegistered(opts *bind.WatchOpts, sink chan<- *IL2AssetRouterAssetHandlerRegistered, assetId [][32]byte, _assetAddress []common.Address) (event.Subscription, error) {

	var assetIdRule []interface{}
	for _, assetIdItem := range assetId {
		assetIdRule = append(assetIdRule, assetIdItem)
	}
	var _assetAddressRule []interface{}
	for _, _assetAddressItem := range _assetAddress {
		_assetAddressRule = append(_assetAddressRule, _assetAddressItem)
	}

	logs, sub, err := _IL2AssetRouter.contract.WatchLogs(opts, "AssetHandlerRegistered", assetIdRule, _assetAddressRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IL2AssetRouterAssetHandlerRegistered)
				if err := _IL2AssetRouter.contract.UnpackLog(event, "AssetHandlerRegistered", log); err != nil {
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

// ParseAssetHandlerRegistered is a log parse operation binding the contract event 0x2632cc0d58b0cb1017b99cc0b6cc66ad86440cc0dd923bfdaa294f95ba1b0201.
//
// Solidity: event AssetHandlerRegistered(bytes32 indexed assetId, address indexed _assetAddress)
func (_IL2AssetRouter *IL2AssetRouterFilterer) ParseAssetHandlerRegistered(log types.Log) (*IL2AssetRouterAssetHandlerRegistered, error) {
	event := new(IL2AssetRouterAssetHandlerRegistered)
	if err := _IL2AssetRouter.contract.UnpackLog(event, "AssetHandlerRegistered", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// IL2AssetRouterAssetHandlerRegisteredInitialIterator is returned from FilterAssetHandlerRegisteredInitial and is used to iterate over the raw logs and unpacked data for AssetHandlerRegisteredInitial events raised by the IL2AssetRouter contract.
type IL2AssetRouterAssetHandlerRegisteredInitialIterator struct {
	Event *IL2AssetRouterAssetHandlerRegisteredInitial // Event containing the contract specifics and raw log

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
func (it *IL2AssetRouterAssetHandlerRegisteredInitialIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IL2AssetRouterAssetHandlerRegisteredInitial)
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
		it.Event = new(IL2AssetRouterAssetHandlerRegisteredInitial)
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
func (it *IL2AssetRouterAssetHandlerRegisteredInitialIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IL2AssetRouterAssetHandlerRegisteredInitialIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IL2AssetRouterAssetHandlerRegisteredInitial represents a AssetHandlerRegisteredInitial event raised by the IL2AssetRouter contract.
type IL2AssetRouterAssetHandlerRegisteredInitial struct {
	AssetId                [32]byte
	AssetHandlerAddress    common.Address
	AdditionalData         [32]byte
	AssetDeploymentTracker common.Address
	Raw                    types.Log // Blockchain specific contextual infos
}

// FilterAssetHandlerRegisteredInitial is a free log retrieval operation binding the contract event 0xb1e82bee3e85b2755fbceb4b7e051f5c66a7f35f0476657504e77e18ebd3a17d.
//
// Solidity: event AssetHandlerRegisteredInitial(bytes32 indexed assetId, address indexed assetHandlerAddress, bytes32 indexed additionalData, address assetDeploymentTracker)
func (_IL2AssetRouter *IL2AssetRouterFilterer) FilterAssetHandlerRegisteredInitial(opts *bind.FilterOpts, assetId [][32]byte, assetHandlerAddress []common.Address, additionalData [][32]byte) (*IL2AssetRouterAssetHandlerRegisteredInitialIterator, error) {

	var assetIdRule []interface{}
	for _, assetIdItem := range assetId {
		assetIdRule = append(assetIdRule, assetIdItem)
	}
	var assetHandlerAddressRule []interface{}
	for _, assetHandlerAddressItem := range assetHandlerAddress {
		assetHandlerAddressRule = append(assetHandlerAddressRule, assetHandlerAddressItem)
	}
	var additionalDataRule []interface{}
	for _, additionalDataItem := range additionalData {
		additionalDataRule = append(additionalDataRule, additionalDataItem)
	}

	logs, sub, err := _IL2AssetRouter.contract.FilterLogs(opts, "AssetHandlerRegisteredInitial", assetIdRule, assetHandlerAddressRule, additionalDataRule)
	if err != nil {
		return nil, err
	}
	return &IL2AssetRouterAssetHandlerRegisteredInitialIterator{contract: _IL2AssetRouter.contract, event: "AssetHandlerRegisteredInitial", logs: logs, sub: sub}, nil
}

// WatchAssetHandlerRegisteredInitial is a free log subscription operation binding the contract event 0xb1e82bee3e85b2755fbceb4b7e051f5c66a7f35f0476657504e77e18ebd3a17d.
//
// Solidity: event AssetHandlerRegisteredInitial(bytes32 indexed assetId, address indexed assetHandlerAddress, bytes32 indexed additionalData, address assetDeploymentTracker)
func (_IL2AssetRouter *IL2AssetRouterFilterer) WatchAssetHandlerRegisteredInitial(opts *bind.WatchOpts, sink chan<- *IL2AssetRouterAssetHandlerRegisteredInitial, assetId [][32]byte, assetHandlerAddress []common.Address, additionalData [][32]byte) (event.Subscription, error) {

	var assetIdRule []interface{}
	for _, assetIdItem := range assetId {
		assetIdRule = append(assetIdRule, assetIdItem)
	}
	var assetHandlerAddressRule []interface{}
	for _, assetHandlerAddressItem := range assetHandlerAddress {
		assetHandlerAddressRule = append(assetHandlerAddressRule, assetHandlerAddressItem)
	}
	var additionalDataRule []interface{}
	for _, additionalDataItem := range additionalData {
		additionalDataRule = append(additionalDataRule, additionalDataItem)
	}

	logs, sub, err := _IL2AssetRouter.contract.WatchLogs(opts, "AssetHandlerRegisteredInitial", assetIdRule, assetHandlerAddressRule, additionalDataRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IL2AssetRouterAssetHandlerRegisteredInitial)
				if err := _IL2AssetRouter.contract.UnpackLog(event, "AssetHandlerRegisteredInitial", log); err != nil {
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

// ParseAssetHandlerRegisteredInitial is a log parse operation binding the contract event 0xb1e82bee3e85b2755fbceb4b7e051f5c66a7f35f0476657504e77e18ebd3a17d.
//
// Solidity: event AssetHandlerRegisteredInitial(bytes32 indexed assetId, address indexed assetHandlerAddress, bytes32 indexed additionalData, address assetDeploymentTracker)
func (_IL2AssetRouter *IL2AssetRouterFilterer) ParseAssetHandlerRegisteredInitial(log types.Log) (*IL2AssetRouterAssetHandlerRegisteredInitial, error) {
	event := new(IL2AssetRouterAssetHandlerRegisteredInitial)
	if err := _IL2AssetRouter.contract.UnpackLog(event, "AssetHandlerRegisteredInitial", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// IL2AssetRouterBridgehubDepositBaseTokenInitiatedIterator is returned from FilterBridgehubDepositBaseTokenInitiated and is used to iterate over the raw logs and unpacked data for BridgehubDepositBaseTokenInitiated events raised by the IL2AssetRouter contract.
type IL2AssetRouterBridgehubDepositBaseTokenInitiatedIterator struct {
	Event *IL2AssetRouterBridgehubDepositBaseTokenInitiated // Event containing the contract specifics and raw log

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
func (it *IL2AssetRouterBridgehubDepositBaseTokenInitiatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IL2AssetRouterBridgehubDepositBaseTokenInitiated)
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
		it.Event = new(IL2AssetRouterBridgehubDepositBaseTokenInitiated)
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
func (it *IL2AssetRouterBridgehubDepositBaseTokenInitiatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IL2AssetRouterBridgehubDepositBaseTokenInitiatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IL2AssetRouterBridgehubDepositBaseTokenInitiated represents a BridgehubDepositBaseTokenInitiated event raised by the IL2AssetRouter contract.
type IL2AssetRouterBridgehubDepositBaseTokenInitiated struct {
	ChainId *big.Int
	From    common.Address
	AssetId [32]byte
	Amount  *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterBridgehubDepositBaseTokenInitiated is a free log retrieval operation binding the contract event 0x0f87e1ea5eb1f034a6071ef630c174063e3d48756f853efaaf4292b929298240.
//
// Solidity: event BridgehubDepositBaseTokenInitiated(uint256 indexed chainId, address indexed from, bytes32 assetId, uint256 amount)
func (_IL2AssetRouter *IL2AssetRouterFilterer) FilterBridgehubDepositBaseTokenInitiated(opts *bind.FilterOpts, chainId []*big.Int, from []common.Address) (*IL2AssetRouterBridgehubDepositBaseTokenInitiatedIterator, error) {

	var chainIdRule []interface{}
	for _, chainIdItem := range chainId {
		chainIdRule = append(chainIdRule, chainIdItem)
	}
	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}

	logs, sub, err := _IL2AssetRouter.contract.FilterLogs(opts, "BridgehubDepositBaseTokenInitiated", chainIdRule, fromRule)
	if err != nil {
		return nil, err
	}
	return &IL2AssetRouterBridgehubDepositBaseTokenInitiatedIterator{contract: _IL2AssetRouter.contract, event: "BridgehubDepositBaseTokenInitiated", logs: logs, sub: sub}, nil
}

// WatchBridgehubDepositBaseTokenInitiated is a free log subscription operation binding the contract event 0x0f87e1ea5eb1f034a6071ef630c174063e3d48756f853efaaf4292b929298240.
//
// Solidity: event BridgehubDepositBaseTokenInitiated(uint256 indexed chainId, address indexed from, bytes32 assetId, uint256 amount)
func (_IL2AssetRouter *IL2AssetRouterFilterer) WatchBridgehubDepositBaseTokenInitiated(opts *bind.WatchOpts, sink chan<- *IL2AssetRouterBridgehubDepositBaseTokenInitiated, chainId []*big.Int, from []common.Address) (event.Subscription, error) {

	var chainIdRule []interface{}
	for _, chainIdItem := range chainId {
		chainIdRule = append(chainIdRule, chainIdItem)
	}
	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}

	logs, sub, err := _IL2AssetRouter.contract.WatchLogs(opts, "BridgehubDepositBaseTokenInitiated", chainIdRule, fromRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IL2AssetRouterBridgehubDepositBaseTokenInitiated)
				if err := _IL2AssetRouter.contract.UnpackLog(event, "BridgehubDepositBaseTokenInitiated", log); err != nil {
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

// ParseBridgehubDepositBaseTokenInitiated is a log parse operation binding the contract event 0x0f87e1ea5eb1f034a6071ef630c174063e3d48756f853efaaf4292b929298240.
//
// Solidity: event BridgehubDepositBaseTokenInitiated(uint256 indexed chainId, address indexed from, bytes32 assetId, uint256 amount)
func (_IL2AssetRouter *IL2AssetRouterFilterer) ParseBridgehubDepositBaseTokenInitiated(log types.Log) (*IL2AssetRouterBridgehubDepositBaseTokenInitiated, error) {
	event := new(IL2AssetRouterBridgehubDepositBaseTokenInitiated)
	if err := _IL2AssetRouter.contract.UnpackLog(event, "BridgehubDepositBaseTokenInitiated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// IL2AssetRouterBridgehubDepositInitiatedIterator is returned from FilterBridgehubDepositInitiated and is used to iterate over the raw logs and unpacked data for BridgehubDepositInitiated events raised by the IL2AssetRouter contract.
type IL2AssetRouterBridgehubDepositInitiatedIterator struct {
	Event *IL2AssetRouterBridgehubDepositInitiated // Event containing the contract specifics and raw log

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
func (it *IL2AssetRouterBridgehubDepositInitiatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IL2AssetRouterBridgehubDepositInitiated)
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
		it.Event = new(IL2AssetRouterBridgehubDepositInitiated)
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
func (it *IL2AssetRouterBridgehubDepositInitiatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IL2AssetRouterBridgehubDepositInitiatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IL2AssetRouterBridgehubDepositInitiated represents a BridgehubDepositInitiated event raised by the IL2AssetRouter contract.
type IL2AssetRouterBridgehubDepositInitiated struct {
	ChainId            *big.Int
	TxDataHash         [32]byte
	From               common.Address
	AssetId            [32]byte
	BridgeMintCalldata []byte
	Raw                types.Log // Blockchain specific contextual infos
}

// FilterBridgehubDepositInitiated is a free log retrieval operation binding the contract event 0xe21913bc89c1320d9709a5d236ffe06b54cf88aecfc9509ebd68f1adba45781e.
//
// Solidity: event BridgehubDepositInitiated(uint256 indexed chainId, bytes32 indexed txDataHash, address indexed from, bytes32 assetId, bytes bridgeMintCalldata)
func (_IL2AssetRouter *IL2AssetRouterFilterer) FilterBridgehubDepositInitiated(opts *bind.FilterOpts, chainId []*big.Int, txDataHash [][32]byte, from []common.Address) (*IL2AssetRouterBridgehubDepositInitiatedIterator, error) {

	var chainIdRule []interface{}
	for _, chainIdItem := range chainId {
		chainIdRule = append(chainIdRule, chainIdItem)
	}
	var txDataHashRule []interface{}
	for _, txDataHashItem := range txDataHash {
		txDataHashRule = append(txDataHashRule, txDataHashItem)
	}
	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}

	logs, sub, err := _IL2AssetRouter.contract.FilterLogs(opts, "BridgehubDepositInitiated", chainIdRule, txDataHashRule, fromRule)
	if err != nil {
		return nil, err
	}
	return &IL2AssetRouterBridgehubDepositInitiatedIterator{contract: _IL2AssetRouter.contract, event: "BridgehubDepositInitiated", logs: logs, sub: sub}, nil
}

// WatchBridgehubDepositInitiated is a free log subscription operation binding the contract event 0xe21913bc89c1320d9709a5d236ffe06b54cf88aecfc9509ebd68f1adba45781e.
//
// Solidity: event BridgehubDepositInitiated(uint256 indexed chainId, bytes32 indexed txDataHash, address indexed from, bytes32 assetId, bytes bridgeMintCalldata)
func (_IL2AssetRouter *IL2AssetRouterFilterer) WatchBridgehubDepositInitiated(opts *bind.WatchOpts, sink chan<- *IL2AssetRouterBridgehubDepositInitiated, chainId []*big.Int, txDataHash [][32]byte, from []common.Address) (event.Subscription, error) {

	var chainIdRule []interface{}
	for _, chainIdItem := range chainId {
		chainIdRule = append(chainIdRule, chainIdItem)
	}
	var txDataHashRule []interface{}
	for _, txDataHashItem := range txDataHash {
		txDataHashRule = append(txDataHashRule, txDataHashItem)
	}
	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}

	logs, sub, err := _IL2AssetRouter.contract.WatchLogs(opts, "BridgehubDepositInitiated", chainIdRule, txDataHashRule, fromRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IL2AssetRouterBridgehubDepositInitiated)
				if err := _IL2AssetRouter.contract.UnpackLog(event, "BridgehubDepositInitiated", log); err != nil {
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

// ParseBridgehubDepositInitiated is a log parse operation binding the contract event 0xe21913bc89c1320d9709a5d236ffe06b54cf88aecfc9509ebd68f1adba45781e.
//
// Solidity: event BridgehubDepositInitiated(uint256 indexed chainId, bytes32 indexed txDataHash, address indexed from, bytes32 assetId, bytes bridgeMintCalldata)
func (_IL2AssetRouter *IL2AssetRouterFilterer) ParseBridgehubDepositInitiated(log types.Log) (*IL2AssetRouterBridgehubDepositInitiated, error) {
	event := new(IL2AssetRouterBridgehubDepositInitiated)
	if err := _IL2AssetRouter.contract.UnpackLog(event, "BridgehubDepositInitiated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// IL2AssetRouterBridgehubWithdrawalInitiatedIterator is returned from FilterBridgehubWithdrawalInitiated and is used to iterate over the raw logs and unpacked data for BridgehubWithdrawalInitiated events raised by the IL2AssetRouter contract.
type IL2AssetRouterBridgehubWithdrawalInitiatedIterator struct {
	Event *IL2AssetRouterBridgehubWithdrawalInitiated // Event containing the contract specifics and raw log

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
func (it *IL2AssetRouterBridgehubWithdrawalInitiatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IL2AssetRouterBridgehubWithdrawalInitiated)
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
		it.Event = new(IL2AssetRouterBridgehubWithdrawalInitiated)
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
func (it *IL2AssetRouterBridgehubWithdrawalInitiatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IL2AssetRouterBridgehubWithdrawalInitiatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IL2AssetRouterBridgehubWithdrawalInitiated represents a BridgehubWithdrawalInitiated event raised by the IL2AssetRouter contract.
type IL2AssetRouterBridgehubWithdrawalInitiated struct {
	ChainId       *big.Int
	Sender        common.Address
	AssetId       [32]byte
	AssetDataHash [32]byte
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterBridgehubWithdrawalInitiated is a free log retrieval operation binding the contract event 0x9a3d4025b7294a1754ea5b56309c1e72328d97b73718183db595c850d14a3ae0.
//
// Solidity: event BridgehubWithdrawalInitiated(uint256 chainId, address indexed sender, bytes32 indexed assetId, bytes32 assetDataHash)
func (_IL2AssetRouter *IL2AssetRouterFilterer) FilterBridgehubWithdrawalInitiated(opts *bind.FilterOpts, sender []common.Address, assetId [][32]byte) (*IL2AssetRouterBridgehubWithdrawalInitiatedIterator, error) {

	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}
	var assetIdRule []interface{}
	for _, assetIdItem := range assetId {
		assetIdRule = append(assetIdRule, assetIdItem)
	}

	logs, sub, err := _IL2AssetRouter.contract.FilterLogs(opts, "BridgehubWithdrawalInitiated", senderRule, assetIdRule)
	if err != nil {
		return nil, err
	}
	return &IL2AssetRouterBridgehubWithdrawalInitiatedIterator{contract: _IL2AssetRouter.contract, event: "BridgehubWithdrawalInitiated", logs: logs, sub: sub}, nil
}

// WatchBridgehubWithdrawalInitiated is a free log subscription operation binding the contract event 0x9a3d4025b7294a1754ea5b56309c1e72328d97b73718183db595c850d14a3ae0.
//
// Solidity: event BridgehubWithdrawalInitiated(uint256 chainId, address indexed sender, bytes32 indexed assetId, bytes32 assetDataHash)
func (_IL2AssetRouter *IL2AssetRouterFilterer) WatchBridgehubWithdrawalInitiated(opts *bind.WatchOpts, sink chan<- *IL2AssetRouterBridgehubWithdrawalInitiated, sender []common.Address, assetId [][32]byte) (event.Subscription, error) {

	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}
	var assetIdRule []interface{}
	for _, assetIdItem := range assetId {
		assetIdRule = append(assetIdRule, assetIdItem)
	}

	logs, sub, err := _IL2AssetRouter.contract.WatchLogs(opts, "BridgehubWithdrawalInitiated", senderRule, assetIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IL2AssetRouterBridgehubWithdrawalInitiated)
				if err := _IL2AssetRouter.contract.UnpackLog(event, "BridgehubWithdrawalInitiated", log); err != nil {
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

// ParseBridgehubWithdrawalInitiated is a log parse operation binding the contract event 0x9a3d4025b7294a1754ea5b56309c1e72328d97b73718183db595c850d14a3ae0.
//
// Solidity: event BridgehubWithdrawalInitiated(uint256 chainId, address indexed sender, bytes32 indexed assetId, bytes32 assetDataHash)
func (_IL2AssetRouter *IL2AssetRouterFilterer) ParseBridgehubWithdrawalInitiated(log types.Log) (*IL2AssetRouterBridgehubWithdrawalInitiated, error) {
	event := new(IL2AssetRouterBridgehubWithdrawalInitiated)
	if err := _IL2AssetRouter.contract.UnpackLog(event, "BridgehubWithdrawalInitiated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// IL2AssetRouterDepositFinalizedAssetRouterIterator is returned from FilterDepositFinalizedAssetRouter and is used to iterate over the raw logs and unpacked data for DepositFinalizedAssetRouter events raised by the IL2AssetRouter contract.
type IL2AssetRouterDepositFinalizedAssetRouterIterator struct {
	Event *IL2AssetRouterDepositFinalizedAssetRouter // Event containing the contract specifics and raw log

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
func (it *IL2AssetRouterDepositFinalizedAssetRouterIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IL2AssetRouterDepositFinalizedAssetRouter)
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
		it.Event = new(IL2AssetRouterDepositFinalizedAssetRouter)
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
func (it *IL2AssetRouterDepositFinalizedAssetRouterIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IL2AssetRouterDepositFinalizedAssetRouterIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IL2AssetRouterDepositFinalizedAssetRouter represents a DepositFinalizedAssetRouter event raised by the IL2AssetRouter contract.
type IL2AssetRouterDepositFinalizedAssetRouter struct {
	ChainId   *big.Int
	AssetId   [32]byte
	AssetData []byte
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterDepositFinalizedAssetRouter is a free log retrieval operation binding the contract event 0x44eb9a840094a49b3cd0a5205042598a1c08c4e87bafb5760bc2d8efa170c541.
//
// Solidity: event DepositFinalizedAssetRouter(uint256 indexed chainId, bytes32 indexed assetId, bytes assetData)
func (_IL2AssetRouter *IL2AssetRouterFilterer) FilterDepositFinalizedAssetRouter(opts *bind.FilterOpts, chainId []*big.Int, assetId [][32]byte) (*IL2AssetRouterDepositFinalizedAssetRouterIterator, error) {

	var chainIdRule []interface{}
	for _, chainIdItem := range chainId {
		chainIdRule = append(chainIdRule, chainIdItem)
	}
	var assetIdRule []interface{}
	for _, assetIdItem := range assetId {
		assetIdRule = append(assetIdRule, assetIdItem)
	}

	logs, sub, err := _IL2AssetRouter.contract.FilterLogs(opts, "DepositFinalizedAssetRouter", chainIdRule, assetIdRule)
	if err != nil {
		return nil, err
	}
	return &IL2AssetRouterDepositFinalizedAssetRouterIterator{contract: _IL2AssetRouter.contract, event: "DepositFinalizedAssetRouter", logs: logs, sub: sub}, nil
}

// WatchDepositFinalizedAssetRouter is a free log subscription operation binding the contract event 0x44eb9a840094a49b3cd0a5205042598a1c08c4e87bafb5760bc2d8efa170c541.
//
// Solidity: event DepositFinalizedAssetRouter(uint256 indexed chainId, bytes32 indexed assetId, bytes assetData)
func (_IL2AssetRouter *IL2AssetRouterFilterer) WatchDepositFinalizedAssetRouter(opts *bind.WatchOpts, sink chan<- *IL2AssetRouterDepositFinalizedAssetRouter, chainId []*big.Int, assetId [][32]byte) (event.Subscription, error) {

	var chainIdRule []interface{}
	for _, chainIdItem := range chainId {
		chainIdRule = append(chainIdRule, chainIdItem)
	}
	var assetIdRule []interface{}
	for _, assetIdItem := range assetId {
		assetIdRule = append(assetIdRule, assetIdItem)
	}

	logs, sub, err := _IL2AssetRouter.contract.WatchLogs(opts, "DepositFinalizedAssetRouter", chainIdRule, assetIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IL2AssetRouterDepositFinalizedAssetRouter)
				if err := _IL2AssetRouter.contract.UnpackLog(event, "DepositFinalizedAssetRouter", log); err != nil {
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

// ParseDepositFinalizedAssetRouter is a log parse operation binding the contract event 0x44eb9a840094a49b3cd0a5205042598a1c08c4e87bafb5760bc2d8efa170c541.
//
// Solidity: event DepositFinalizedAssetRouter(uint256 indexed chainId, bytes32 indexed assetId, bytes assetData)
func (_IL2AssetRouter *IL2AssetRouterFilterer) ParseDepositFinalizedAssetRouter(log types.Log) (*IL2AssetRouterDepositFinalizedAssetRouter, error) {
	event := new(IL2AssetRouterDepositFinalizedAssetRouter)
	if err := _IL2AssetRouter.contract.UnpackLog(event, "DepositFinalizedAssetRouter", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// IL2AssetRouterWithdrawalInitiatedAssetRouterIterator is returned from FilterWithdrawalInitiatedAssetRouter and is used to iterate over the raw logs and unpacked data for WithdrawalInitiatedAssetRouter events raised by the IL2AssetRouter contract.
type IL2AssetRouterWithdrawalInitiatedAssetRouterIterator struct {
	Event *IL2AssetRouterWithdrawalInitiatedAssetRouter // Event containing the contract specifics and raw log

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
func (it *IL2AssetRouterWithdrawalInitiatedAssetRouterIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IL2AssetRouterWithdrawalInitiatedAssetRouter)
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
		it.Event = new(IL2AssetRouterWithdrawalInitiatedAssetRouter)
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
func (it *IL2AssetRouterWithdrawalInitiatedAssetRouterIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IL2AssetRouterWithdrawalInitiatedAssetRouterIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IL2AssetRouterWithdrawalInitiatedAssetRouter represents a WithdrawalInitiatedAssetRouter event raised by the IL2AssetRouter contract.
type IL2AssetRouterWithdrawalInitiatedAssetRouter struct {
	ChainId   *big.Int
	L2Sender  common.Address
	AssetId   [32]byte
	AssetData []byte
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterWithdrawalInitiatedAssetRouter is a free log retrieval operation binding the contract event 0x55362fc62473cb1255e770af5d5e02ba6ee5bc7ed6969c30eb11ca31b92384dc.
//
// Solidity: event WithdrawalInitiatedAssetRouter(uint256 chainId, address indexed l2Sender, bytes32 indexed assetId, bytes assetData)
func (_IL2AssetRouter *IL2AssetRouterFilterer) FilterWithdrawalInitiatedAssetRouter(opts *bind.FilterOpts, l2Sender []common.Address, assetId [][32]byte) (*IL2AssetRouterWithdrawalInitiatedAssetRouterIterator, error) {

	var l2SenderRule []interface{}
	for _, l2SenderItem := range l2Sender {
		l2SenderRule = append(l2SenderRule, l2SenderItem)
	}
	var assetIdRule []interface{}
	for _, assetIdItem := range assetId {
		assetIdRule = append(assetIdRule, assetIdItem)
	}

	logs, sub, err := _IL2AssetRouter.contract.FilterLogs(opts, "WithdrawalInitiatedAssetRouter", l2SenderRule, assetIdRule)
	if err != nil {
		return nil, err
	}
	return &IL2AssetRouterWithdrawalInitiatedAssetRouterIterator{contract: _IL2AssetRouter.contract, event: "WithdrawalInitiatedAssetRouter", logs: logs, sub: sub}, nil
}

// WatchWithdrawalInitiatedAssetRouter is a free log subscription operation binding the contract event 0x55362fc62473cb1255e770af5d5e02ba6ee5bc7ed6969c30eb11ca31b92384dc.
//
// Solidity: event WithdrawalInitiatedAssetRouter(uint256 chainId, address indexed l2Sender, bytes32 indexed assetId, bytes assetData)
func (_IL2AssetRouter *IL2AssetRouterFilterer) WatchWithdrawalInitiatedAssetRouter(opts *bind.WatchOpts, sink chan<- *IL2AssetRouterWithdrawalInitiatedAssetRouter, l2Sender []common.Address, assetId [][32]byte) (event.Subscription, error) {

	var l2SenderRule []interface{}
	for _, l2SenderItem := range l2Sender {
		l2SenderRule = append(l2SenderRule, l2SenderItem)
	}
	var assetIdRule []interface{}
	for _, assetIdItem := range assetId {
		assetIdRule = append(assetIdRule, assetIdItem)
	}

	logs, sub, err := _IL2AssetRouter.contract.WatchLogs(opts, "WithdrawalInitiatedAssetRouter", l2SenderRule, assetIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IL2AssetRouterWithdrawalInitiatedAssetRouter)
				if err := _IL2AssetRouter.contract.UnpackLog(event, "WithdrawalInitiatedAssetRouter", log); err != nil {
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

// ParseWithdrawalInitiatedAssetRouter is a log parse operation binding the contract event 0x55362fc62473cb1255e770af5d5e02ba6ee5bc7ed6969c30eb11ca31b92384dc.
//
// Solidity: event WithdrawalInitiatedAssetRouter(uint256 chainId, address indexed l2Sender, bytes32 indexed assetId, bytes assetData)
func (_IL2AssetRouter *IL2AssetRouterFilterer) ParseWithdrawalInitiatedAssetRouter(log types.Log) (*IL2AssetRouterWithdrawalInitiatedAssetRouter, error) {
	event := new(IL2AssetRouterWithdrawalInitiatedAssetRouter)
	if err := _IL2AssetRouter.contract.UnpackLog(event, "WithdrawalInitiatedAssetRouter", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
