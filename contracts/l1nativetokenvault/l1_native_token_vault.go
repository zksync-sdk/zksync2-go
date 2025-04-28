// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package l1nativetokenvault

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

// IL1NativeTokenVaultMetaData contains all meta data concerning the IL1NativeTokenVault contract.
var IL1NativeTokenVaultMetaData = &bind.MetaData{
	ABI: "[{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"bridgedTokenBeacon\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"bridgedTokenProxyBytecodeHash\",\"type\":\"bytes32\"}],\"name\":\"BridgedTokenBeaconUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"l2TokenBeacon\",\"type\":\"address\"}],\"name\":\"TokenBeaconUpdated\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"ASSET_ROUTER\",\"outputs\":[{\"internalType\":\"contractIAssetRouterBase\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"L1_NULLIFIER\",\"outputs\":[{\"internalType\":\"contractIL1Nullifier\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"WETH_TOKEN\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"}],\"name\":\"assetId\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_chainId\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"_tokenAddress\",\"type\":\"address\"}],\"name\":\"calculateAssetId\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_originChainId\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"_originToken\",\"type\":\"address\"}],\"name\":\"calculateCreate2TokenAddress\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_chainId\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"_assetId\",\"type\":\"bytes32\"}],\"name\":\"chainBalance\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_originChainId\",\"type\":\"uint256\"}],\"name\":\"getERC20Getters\",\"outputs\":[{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"assetId\",\"type\":\"bytes32\"}],\"name\":\"originChainId\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"registerEthToken\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_l1Token\",\"type\":\"address\"}],\"name\":\"registerToken\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"assetId\",\"type\":\"bytes32\"}],\"name\":\"tokenAddress\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
}

// IL1NativeTokenVaultABI is the input ABI used to generate the binding from.
// Deprecated: Use IL1NativeTokenVaultMetaData.ABI instead.
var IL1NativeTokenVaultABI = IL1NativeTokenVaultMetaData.ABI

// IL1NativeTokenVault is an auto generated Go binding around an Ethereum contract.
type IL1NativeTokenVault struct {
	IL1NativeTokenVaultCaller     // Read-only binding to the contract
	IL1NativeTokenVaultTransactor // Write-only binding to the contract
	IL1NativeTokenVaultFilterer   // Log filterer for contract events
}

// IL1NativeTokenVaultCaller is an auto generated read-only Go binding around an Ethereum contract.
type IL1NativeTokenVaultCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IL1NativeTokenVaultTransactor is an auto generated write-only Go binding around an Ethereum contract.
type IL1NativeTokenVaultTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IL1NativeTokenVaultFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type IL1NativeTokenVaultFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IL1NativeTokenVaultSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type IL1NativeTokenVaultSession struct {
	Contract     *IL1NativeTokenVault // Generic contract binding to set the session for
	CallOpts     bind.CallOpts        // Call options to use throughout this session
	TransactOpts bind.TransactOpts    // Transaction auth options to use throughout this session
}

// IL1NativeTokenVaultCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type IL1NativeTokenVaultCallerSession struct {
	Contract *IL1NativeTokenVaultCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts              // Call options to use throughout this session
}

// IL1NativeTokenVaultTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type IL1NativeTokenVaultTransactorSession struct {
	Contract     *IL1NativeTokenVaultTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts              // Transaction auth options to use throughout this session
}

// IL1NativeTokenVaultRaw is an auto generated low-level Go binding around an Ethereum contract.
type IL1NativeTokenVaultRaw struct {
	Contract *IL1NativeTokenVault // Generic contract binding to access the raw methods on
}

// IL1NativeTokenVaultCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type IL1NativeTokenVaultCallerRaw struct {
	Contract *IL1NativeTokenVaultCaller // Generic read-only contract binding to access the raw methods on
}

// IL1NativeTokenVaultTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type IL1NativeTokenVaultTransactorRaw struct {
	Contract *IL1NativeTokenVaultTransactor // Generic write-only contract binding to access the raw methods on
}

// NewIL1NativeTokenVault creates a new instance of IL1NativeTokenVault, bound to a specific deployed contract.
func NewIL1NativeTokenVault(address common.Address, backend bind.ContractBackend) (*IL1NativeTokenVault, error) {
	contract, err := bindIL1NativeTokenVault(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &IL1NativeTokenVault{IL1NativeTokenVaultCaller: IL1NativeTokenVaultCaller{contract: contract}, IL1NativeTokenVaultTransactor: IL1NativeTokenVaultTransactor{contract: contract}, IL1NativeTokenVaultFilterer: IL1NativeTokenVaultFilterer{contract: contract}}, nil
}

// NewIL1NativeTokenVaultCaller creates a new read-only instance of IL1NativeTokenVault, bound to a specific deployed contract.
func NewIL1NativeTokenVaultCaller(address common.Address, caller bind.ContractCaller) (*IL1NativeTokenVaultCaller, error) {
	contract, err := bindIL1NativeTokenVault(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &IL1NativeTokenVaultCaller{contract: contract}, nil
}

// NewIL1NativeTokenVaultTransactor creates a new write-only instance of IL1NativeTokenVault, bound to a specific deployed contract.
func NewIL1NativeTokenVaultTransactor(address common.Address, transactor bind.ContractTransactor) (*IL1NativeTokenVaultTransactor, error) {
	contract, err := bindIL1NativeTokenVault(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &IL1NativeTokenVaultTransactor{contract: contract}, nil
}

// NewIL1NativeTokenVaultFilterer creates a new log filterer instance of IL1NativeTokenVault, bound to a specific deployed contract.
func NewIL1NativeTokenVaultFilterer(address common.Address, filterer bind.ContractFilterer) (*IL1NativeTokenVaultFilterer, error) {
	contract, err := bindIL1NativeTokenVault(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &IL1NativeTokenVaultFilterer{contract: contract}, nil
}

// bindIL1NativeTokenVault binds a generic wrapper to an already deployed contract.
func bindIL1NativeTokenVault(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := IL1NativeTokenVaultMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IL1NativeTokenVault *IL1NativeTokenVaultRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IL1NativeTokenVault.Contract.IL1NativeTokenVaultCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IL1NativeTokenVault *IL1NativeTokenVaultRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IL1NativeTokenVault.Contract.IL1NativeTokenVaultTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IL1NativeTokenVault *IL1NativeTokenVaultRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IL1NativeTokenVault.Contract.IL1NativeTokenVaultTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IL1NativeTokenVault *IL1NativeTokenVaultCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IL1NativeTokenVault.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IL1NativeTokenVault *IL1NativeTokenVaultTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IL1NativeTokenVault.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IL1NativeTokenVault *IL1NativeTokenVaultTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IL1NativeTokenVault.Contract.contract.Transact(opts, method, params...)
}

// ASSETROUTER is a free data retrieval call binding the contract method 0xc6a70bbb.
//
// Solidity: function ASSET_ROUTER() view returns(address)
func (_IL1NativeTokenVault *IL1NativeTokenVaultCaller) ASSETROUTER(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _IL1NativeTokenVault.contract.Call(opts, &out, "ASSET_ROUTER")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// ASSETROUTER is a free data retrieval call binding the contract method 0xc6a70bbb.
//
// Solidity: function ASSET_ROUTER() view returns(address)
func (_IL1NativeTokenVault *IL1NativeTokenVaultSession) ASSETROUTER() (common.Address, error) {
	return _IL1NativeTokenVault.Contract.ASSETROUTER(&_IL1NativeTokenVault.CallOpts)
}

// ASSETROUTER is a free data retrieval call binding the contract method 0xc6a70bbb.
//
// Solidity: function ASSET_ROUTER() view returns(address)
func (_IL1NativeTokenVault *IL1NativeTokenVaultCallerSession) ASSETROUTER() (common.Address, error) {
	return _IL1NativeTokenVault.Contract.ASSETROUTER(&_IL1NativeTokenVault.CallOpts)
}

// L1NULLIFIER is a free data retrieval call binding the contract method 0xe60ccaba.
//
// Solidity: function L1_NULLIFIER() view returns(address)
func (_IL1NativeTokenVault *IL1NativeTokenVaultCaller) L1NULLIFIER(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _IL1NativeTokenVault.contract.Call(opts, &out, "L1_NULLIFIER")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// L1NULLIFIER is a free data retrieval call binding the contract method 0xe60ccaba.
//
// Solidity: function L1_NULLIFIER() view returns(address)
func (_IL1NativeTokenVault *IL1NativeTokenVaultSession) L1NULLIFIER() (common.Address, error) {
	return _IL1NativeTokenVault.Contract.L1NULLIFIER(&_IL1NativeTokenVault.CallOpts)
}

// L1NULLIFIER is a free data retrieval call binding the contract method 0xe60ccaba.
//
// Solidity: function L1_NULLIFIER() view returns(address)
func (_IL1NativeTokenVault *IL1NativeTokenVaultCallerSession) L1NULLIFIER() (common.Address, error) {
	return _IL1NativeTokenVault.Contract.L1NULLIFIER(&_IL1NativeTokenVault.CallOpts)
}

// WETHTOKEN is a free data retrieval call binding the contract method 0x37d277d4.
//
// Solidity: function WETH_TOKEN() view returns(address)
func (_IL1NativeTokenVault *IL1NativeTokenVaultCaller) WETHTOKEN(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _IL1NativeTokenVault.contract.Call(opts, &out, "WETH_TOKEN")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// WETHTOKEN is a free data retrieval call binding the contract method 0x37d277d4.
//
// Solidity: function WETH_TOKEN() view returns(address)
func (_IL1NativeTokenVault *IL1NativeTokenVaultSession) WETHTOKEN() (common.Address, error) {
	return _IL1NativeTokenVault.Contract.WETHTOKEN(&_IL1NativeTokenVault.CallOpts)
}

// WETHTOKEN is a free data retrieval call binding the contract method 0x37d277d4.
//
// Solidity: function WETH_TOKEN() view returns(address)
func (_IL1NativeTokenVault *IL1NativeTokenVaultCallerSession) WETHTOKEN() (common.Address, error) {
	return _IL1NativeTokenVault.Contract.WETHTOKEN(&_IL1NativeTokenVault.CallOpts)
}

// AssetId is a free data retrieval call binding the contract method 0xfd3f60df.
//
// Solidity: function assetId(address token) view returns(bytes32)
func (_IL1NativeTokenVault *IL1NativeTokenVaultCaller) AssetId(opts *bind.CallOpts, token common.Address) ([32]byte, error) {
	var out []interface{}
	err := _IL1NativeTokenVault.contract.Call(opts, &out, "assetId", token)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// AssetId is a free data retrieval call binding the contract method 0xfd3f60df.
//
// Solidity: function assetId(address token) view returns(bytes32)
func (_IL1NativeTokenVault *IL1NativeTokenVaultSession) AssetId(token common.Address) ([32]byte, error) {
	return _IL1NativeTokenVault.Contract.AssetId(&_IL1NativeTokenVault.CallOpts, token)
}

// AssetId is a free data retrieval call binding the contract method 0xfd3f60df.
//
// Solidity: function assetId(address token) view returns(bytes32)
func (_IL1NativeTokenVault *IL1NativeTokenVaultCallerSession) AssetId(token common.Address) ([32]byte, error) {
	return _IL1NativeTokenVault.Contract.AssetId(&_IL1NativeTokenVault.CallOpts, token)
}

// CalculateAssetId is a free data retrieval call binding the contract method 0xa42c88a2.
//
// Solidity: function calculateAssetId(uint256 _chainId, address _tokenAddress) view returns(bytes32)
func (_IL1NativeTokenVault *IL1NativeTokenVaultCaller) CalculateAssetId(opts *bind.CallOpts, _chainId *big.Int, _tokenAddress common.Address) ([32]byte, error) {
	var out []interface{}
	err := _IL1NativeTokenVault.contract.Call(opts, &out, "calculateAssetId", _chainId, _tokenAddress)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// CalculateAssetId is a free data retrieval call binding the contract method 0xa42c88a2.
//
// Solidity: function calculateAssetId(uint256 _chainId, address _tokenAddress) view returns(bytes32)
func (_IL1NativeTokenVault *IL1NativeTokenVaultSession) CalculateAssetId(_chainId *big.Int, _tokenAddress common.Address) ([32]byte, error) {
	return _IL1NativeTokenVault.Contract.CalculateAssetId(&_IL1NativeTokenVault.CallOpts, _chainId, _tokenAddress)
}

// CalculateAssetId is a free data retrieval call binding the contract method 0xa42c88a2.
//
// Solidity: function calculateAssetId(uint256 _chainId, address _tokenAddress) view returns(bytes32)
func (_IL1NativeTokenVault *IL1NativeTokenVaultCallerSession) CalculateAssetId(_chainId *big.Int, _tokenAddress common.Address) ([32]byte, error) {
	return _IL1NativeTokenVault.Contract.CalculateAssetId(&_IL1NativeTokenVault.CallOpts, _chainId, _tokenAddress)
}

// CalculateCreate2TokenAddress is a free data retrieval call binding the contract method 0xc487412c.
//
// Solidity: function calculateCreate2TokenAddress(uint256 _originChainId, address _originToken) view returns(address)
func (_IL1NativeTokenVault *IL1NativeTokenVaultCaller) CalculateCreate2TokenAddress(opts *bind.CallOpts, _originChainId *big.Int, _originToken common.Address) (common.Address, error) {
	var out []interface{}
	err := _IL1NativeTokenVault.contract.Call(opts, &out, "calculateCreate2TokenAddress", _originChainId, _originToken)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// CalculateCreate2TokenAddress is a free data retrieval call binding the contract method 0xc487412c.
//
// Solidity: function calculateCreate2TokenAddress(uint256 _originChainId, address _originToken) view returns(address)
func (_IL1NativeTokenVault *IL1NativeTokenVaultSession) CalculateCreate2TokenAddress(_originChainId *big.Int, _originToken common.Address) (common.Address, error) {
	return _IL1NativeTokenVault.Contract.CalculateCreate2TokenAddress(&_IL1NativeTokenVault.CallOpts, _originChainId, _originToken)
}

// CalculateCreate2TokenAddress is a free data retrieval call binding the contract method 0xc487412c.
//
// Solidity: function calculateCreate2TokenAddress(uint256 _originChainId, address _originToken) view returns(address)
func (_IL1NativeTokenVault *IL1NativeTokenVaultCallerSession) CalculateCreate2TokenAddress(_originChainId *big.Int, _originToken common.Address) (common.Address, error) {
	return _IL1NativeTokenVault.Contract.CalculateCreate2TokenAddress(&_IL1NativeTokenVault.CallOpts, _originChainId, _originToken)
}

// ChainBalance is a free data retrieval call binding the contract method 0x3345359b.
//
// Solidity: function chainBalance(uint256 _chainId, bytes32 _assetId) view returns(uint256)
func (_IL1NativeTokenVault *IL1NativeTokenVaultCaller) ChainBalance(opts *bind.CallOpts, _chainId *big.Int, _assetId [32]byte) (*big.Int, error) {
	var out []interface{}
	err := _IL1NativeTokenVault.contract.Call(opts, &out, "chainBalance", _chainId, _assetId)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// ChainBalance is a free data retrieval call binding the contract method 0x3345359b.
//
// Solidity: function chainBalance(uint256 _chainId, bytes32 _assetId) view returns(uint256)
func (_IL1NativeTokenVault *IL1NativeTokenVaultSession) ChainBalance(_chainId *big.Int, _assetId [32]byte) (*big.Int, error) {
	return _IL1NativeTokenVault.Contract.ChainBalance(&_IL1NativeTokenVault.CallOpts, _chainId, _assetId)
}

// ChainBalance is a free data retrieval call binding the contract method 0x3345359b.
//
// Solidity: function chainBalance(uint256 _chainId, bytes32 _assetId) view returns(uint256)
func (_IL1NativeTokenVault *IL1NativeTokenVaultCallerSession) ChainBalance(_chainId *big.Int, _assetId [32]byte) (*big.Int, error) {
	return _IL1NativeTokenVault.Contract.ChainBalance(&_IL1NativeTokenVault.CallOpts, _chainId, _assetId)
}

// GetERC20Getters is a free data retrieval call binding the contract method 0xa7236d16.
//
// Solidity: function getERC20Getters(address _token, uint256 _originChainId) view returns(bytes)
func (_IL1NativeTokenVault *IL1NativeTokenVaultCaller) GetERC20Getters(opts *bind.CallOpts, _token common.Address, _originChainId *big.Int) ([]byte, error) {
	var out []interface{}
	err := _IL1NativeTokenVault.contract.Call(opts, &out, "getERC20Getters", _token, _originChainId)

	if err != nil {
		return *new([]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([]byte)).(*[]byte)

	return out0, err

}

// GetERC20Getters is a free data retrieval call binding the contract method 0xa7236d16.
//
// Solidity: function getERC20Getters(address _token, uint256 _originChainId) view returns(bytes)
func (_IL1NativeTokenVault *IL1NativeTokenVaultSession) GetERC20Getters(_token common.Address, _originChainId *big.Int) ([]byte, error) {
	return _IL1NativeTokenVault.Contract.GetERC20Getters(&_IL1NativeTokenVault.CallOpts, _token, _originChainId)
}

// GetERC20Getters is a free data retrieval call binding the contract method 0xa7236d16.
//
// Solidity: function getERC20Getters(address _token, uint256 _originChainId) view returns(bytes)
func (_IL1NativeTokenVault *IL1NativeTokenVaultCallerSession) GetERC20Getters(_token common.Address, _originChainId *big.Int) ([]byte, error) {
	return _IL1NativeTokenVault.Contract.GetERC20Getters(&_IL1NativeTokenVault.CallOpts, _token, _originChainId)
}

// OriginChainId is a free data retrieval call binding the contract method 0x5f3455b5.
//
// Solidity: function originChainId(bytes32 assetId) view returns(uint256)
func (_IL1NativeTokenVault *IL1NativeTokenVaultCaller) OriginChainId(opts *bind.CallOpts, assetId [32]byte) (*big.Int, error) {
	var out []interface{}
	err := _IL1NativeTokenVault.contract.Call(opts, &out, "originChainId", assetId)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// OriginChainId is a free data retrieval call binding the contract method 0x5f3455b5.
//
// Solidity: function originChainId(bytes32 assetId) view returns(uint256)
func (_IL1NativeTokenVault *IL1NativeTokenVaultSession) OriginChainId(assetId [32]byte) (*big.Int, error) {
	return _IL1NativeTokenVault.Contract.OriginChainId(&_IL1NativeTokenVault.CallOpts, assetId)
}

// OriginChainId is a free data retrieval call binding the contract method 0x5f3455b5.
//
// Solidity: function originChainId(bytes32 assetId) view returns(uint256)
func (_IL1NativeTokenVault *IL1NativeTokenVaultCallerSession) OriginChainId(assetId [32]byte) (*big.Int, error) {
	return _IL1NativeTokenVault.Contract.OriginChainId(&_IL1NativeTokenVault.CallOpts, assetId)
}

// TokenAddress is a free data retrieval call binding the contract method 0x97bb3ce9.
//
// Solidity: function tokenAddress(bytes32 assetId) view returns(address)
func (_IL1NativeTokenVault *IL1NativeTokenVaultCaller) TokenAddress(opts *bind.CallOpts, assetId [32]byte) (common.Address, error) {
	var out []interface{}
	err := _IL1NativeTokenVault.contract.Call(opts, &out, "tokenAddress", assetId)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// TokenAddress is a free data retrieval call binding the contract method 0x97bb3ce9.
//
// Solidity: function tokenAddress(bytes32 assetId) view returns(address)
func (_IL1NativeTokenVault *IL1NativeTokenVaultSession) TokenAddress(assetId [32]byte) (common.Address, error) {
	return _IL1NativeTokenVault.Contract.TokenAddress(&_IL1NativeTokenVault.CallOpts, assetId)
}

// TokenAddress is a free data retrieval call binding the contract method 0x97bb3ce9.
//
// Solidity: function tokenAddress(bytes32 assetId) view returns(address)
func (_IL1NativeTokenVault *IL1NativeTokenVaultCallerSession) TokenAddress(assetId [32]byte) (common.Address, error) {
	return _IL1NativeTokenVault.Contract.TokenAddress(&_IL1NativeTokenVault.CallOpts, assetId)
}

// RegisterEthToken is a paid mutator transaction binding the contract method 0xcb6da609.
//
// Solidity: function registerEthToken() returns()
func (_IL1NativeTokenVault *IL1NativeTokenVaultTransactor) RegisterEthToken(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IL1NativeTokenVault.contract.Transact(opts, "registerEthToken")
}

// RegisterEthToken is a paid mutator transaction binding the contract method 0xcb6da609.
//
// Solidity: function registerEthToken() returns()
func (_IL1NativeTokenVault *IL1NativeTokenVaultSession) RegisterEthToken() (*types.Transaction, error) {
	return _IL1NativeTokenVault.Contract.RegisterEthToken(&_IL1NativeTokenVault.TransactOpts)
}

// RegisterEthToken is a paid mutator transaction binding the contract method 0xcb6da609.
//
// Solidity: function registerEthToken() returns()
func (_IL1NativeTokenVault *IL1NativeTokenVaultTransactorSession) RegisterEthToken() (*types.Transaction, error) {
	return _IL1NativeTokenVault.Contract.RegisterEthToken(&_IL1NativeTokenVault.TransactOpts)
}

// RegisterToken is a paid mutator transaction binding the contract method 0x09824a80.
//
// Solidity: function registerToken(address _l1Token) returns()
func (_IL1NativeTokenVault *IL1NativeTokenVaultTransactor) RegisterToken(opts *bind.TransactOpts, _l1Token common.Address) (*types.Transaction, error) {
	return _IL1NativeTokenVault.contract.Transact(opts, "registerToken", _l1Token)
}

// RegisterToken is a paid mutator transaction binding the contract method 0x09824a80.
//
// Solidity: function registerToken(address _l1Token) returns()
func (_IL1NativeTokenVault *IL1NativeTokenVaultSession) RegisterToken(_l1Token common.Address) (*types.Transaction, error) {
	return _IL1NativeTokenVault.Contract.RegisterToken(&_IL1NativeTokenVault.TransactOpts, _l1Token)
}

// RegisterToken is a paid mutator transaction binding the contract method 0x09824a80.
//
// Solidity: function registerToken(address _l1Token) returns()
func (_IL1NativeTokenVault *IL1NativeTokenVaultTransactorSession) RegisterToken(_l1Token common.Address) (*types.Transaction, error) {
	return _IL1NativeTokenVault.Contract.RegisterToken(&_IL1NativeTokenVault.TransactOpts, _l1Token)
}

// IL1NativeTokenVaultBridgedTokenBeaconUpdatedIterator is returned from FilterBridgedTokenBeaconUpdated and is used to iterate over the raw logs and unpacked data for BridgedTokenBeaconUpdated events raised by the IL1NativeTokenVault contract.
type IL1NativeTokenVaultBridgedTokenBeaconUpdatedIterator struct {
	Event *IL1NativeTokenVaultBridgedTokenBeaconUpdated // Event containing the contract specifics and raw log

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
func (it *IL1NativeTokenVaultBridgedTokenBeaconUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IL1NativeTokenVaultBridgedTokenBeaconUpdated)
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
		it.Event = new(IL1NativeTokenVaultBridgedTokenBeaconUpdated)
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
func (it *IL1NativeTokenVaultBridgedTokenBeaconUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IL1NativeTokenVaultBridgedTokenBeaconUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IL1NativeTokenVaultBridgedTokenBeaconUpdated represents a BridgedTokenBeaconUpdated event raised by the IL1NativeTokenVault contract.
type IL1NativeTokenVaultBridgedTokenBeaconUpdated struct {
	BridgedTokenBeacon            common.Address
	BridgedTokenProxyBytecodeHash [32]byte
	Raw                           types.Log // Blockchain specific contextual infos
}

// FilterBridgedTokenBeaconUpdated is a free log retrieval operation binding the contract event 0xc3f14dba68f86c42f518e5c0e8a5cbc9514da6f388e2f52c5b1a6263d8588bfb.
//
// Solidity: event BridgedTokenBeaconUpdated(address bridgedTokenBeacon, bytes32 bridgedTokenProxyBytecodeHash)
func (_IL1NativeTokenVault *IL1NativeTokenVaultFilterer) FilterBridgedTokenBeaconUpdated(opts *bind.FilterOpts) (*IL1NativeTokenVaultBridgedTokenBeaconUpdatedIterator, error) {

	logs, sub, err := _IL1NativeTokenVault.contract.FilterLogs(opts, "BridgedTokenBeaconUpdated")
	if err != nil {
		return nil, err
	}
	return &IL1NativeTokenVaultBridgedTokenBeaconUpdatedIterator{contract: _IL1NativeTokenVault.contract, event: "BridgedTokenBeaconUpdated", logs: logs, sub: sub}, nil
}

// WatchBridgedTokenBeaconUpdated is a free log subscription operation binding the contract event 0xc3f14dba68f86c42f518e5c0e8a5cbc9514da6f388e2f52c5b1a6263d8588bfb.
//
// Solidity: event BridgedTokenBeaconUpdated(address bridgedTokenBeacon, bytes32 bridgedTokenProxyBytecodeHash)
func (_IL1NativeTokenVault *IL1NativeTokenVaultFilterer) WatchBridgedTokenBeaconUpdated(opts *bind.WatchOpts, sink chan<- *IL1NativeTokenVaultBridgedTokenBeaconUpdated) (event.Subscription, error) {

	logs, sub, err := _IL1NativeTokenVault.contract.WatchLogs(opts, "BridgedTokenBeaconUpdated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IL1NativeTokenVaultBridgedTokenBeaconUpdated)
				if err := _IL1NativeTokenVault.contract.UnpackLog(event, "BridgedTokenBeaconUpdated", log); err != nil {
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

// ParseBridgedTokenBeaconUpdated is a log parse operation binding the contract event 0xc3f14dba68f86c42f518e5c0e8a5cbc9514da6f388e2f52c5b1a6263d8588bfb.
//
// Solidity: event BridgedTokenBeaconUpdated(address bridgedTokenBeacon, bytes32 bridgedTokenProxyBytecodeHash)
func (_IL1NativeTokenVault *IL1NativeTokenVaultFilterer) ParseBridgedTokenBeaconUpdated(log types.Log) (*IL1NativeTokenVaultBridgedTokenBeaconUpdated, error) {
	event := new(IL1NativeTokenVaultBridgedTokenBeaconUpdated)
	if err := _IL1NativeTokenVault.contract.UnpackLog(event, "BridgedTokenBeaconUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// IL1NativeTokenVaultTokenBeaconUpdatedIterator is returned from FilterTokenBeaconUpdated and is used to iterate over the raw logs and unpacked data for TokenBeaconUpdated events raised by the IL1NativeTokenVault contract.
type IL1NativeTokenVaultTokenBeaconUpdatedIterator struct {
	Event *IL1NativeTokenVaultTokenBeaconUpdated // Event containing the contract specifics and raw log

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
func (it *IL1NativeTokenVaultTokenBeaconUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IL1NativeTokenVaultTokenBeaconUpdated)
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
		it.Event = new(IL1NativeTokenVaultTokenBeaconUpdated)
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
func (it *IL1NativeTokenVaultTokenBeaconUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IL1NativeTokenVaultTokenBeaconUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IL1NativeTokenVaultTokenBeaconUpdated represents a TokenBeaconUpdated event raised by the IL1NativeTokenVault contract.
type IL1NativeTokenVaultTokenBeaconUpdated struct {
	L2TokenBeacon common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterTokenBeaconUpdated is a free log retrieval operation binding the contract event 0x5ed5e4f58bf9a324a38beaa1177fb96fcb7bf3a5f4c4585ebb78c4a8c0249d0f.
//
// Solidity: event TokenBeaconUpdated(address indexed l2TokenBeacon)
func (_IL1NativeTokenVault *IL1NativeTokenVaultFilterer) FilterTokenBeaconUpdated(opts *bind.FilterOpts, l2TokenBeacon []common.Address) (*IL1NativeTokenVaultTokenBeaconUpdatedIterator, error) {

	var l2TokenBeaconRule []interface{}
	for _, l2TokenBeaconItem := range l2TokenBeacon {
		l2TokenBeaconRule = append(l2TokenBeaconRule, l2TokenBeaconItem)
	}

	logs, sub, err := _IL1NativeTokenVault.contract.FilterLogs(opts, "TokenBeaconUpdated", l2TokenBeaconRule)
	if err != nil {
		return nil, err
	}
	return &IL1NativeTokenVaultTokenBeaconUpdatedIterator{contract: _IL1NativeTokenVault.contract, event: "TokenBeaconUpdated", logs: logs, sub: sub}, nil
}

// WatchTokenBeaconUpdated is a free log subscription operation binding the contract event 0x5ed5e4f58bf9a324a38beaa1177fb96fcb7bf3a5f4c4585ebb78c4a8c0249d0f.
//
// Solidity: event TokenBeaconUpdated(address indexed l2TokenBeacon)
func (_IL1NativeTokenVault *IL1NativeTokenVaultFilterer) WatchTokenBeaconUpdated(opts *bind.WatchOpts, sink chan<- *IL1NativeTokenVaultTokenBeaconUpdated, l2TokenBeacon []common.Address) (event.Subscription, error) {

	var l2TokenBeaconRule []interface{}
	for _, l2TokenBeaconItem := range l2TokenBeacon {
		l2TokenBeaconRule = append(l2TokenBeaconRule, l2TokenBeaconItem)
	}

	logs, sub, err := _IL1NativeTokenVault.contract.WatchLogs(opts, "TokenBeaconUpdated", l2TokenBeaconRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IL1NativeTokenVaultTokenBeaconUpdated)
				if err := _IL1NativeTokenVault.contract.UnpackLog(event, "TokenBeaconUpdated", log); err != nil {
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

// ParseTokenBeaconUpdated is a log parse operation binding the contract event 0x5ed5e4f58bf9a324a38beaa1177fb96fcb7bf3a5f4c4585ebb78c4a8c0249d0f.
//
// Solidity: event TokenBeaconUpdated(address indexed l2TokenBeacon)
func (_IL1NativeTokenVault *IL1NativeTokenVaultFilterer) ParseTokenBeaconUpdated(log types.Log) (*IL1NativeTokenVaultTokenBeaconUpdated, error) {
	event := new(IL1NativeTokenVaultTokenBeaconUpdated)
	if err := _IL1NativeTokenVault.contract.UnpackLog(event, "TokenBeaconUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
