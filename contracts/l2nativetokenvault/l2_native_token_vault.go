// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package l2nativetokenvault

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

// IL2NativeTokenVaultMetaData contains all meta data concerning the IL2NativeTokenVault contract.
var IL2NativeTokenVaultMetaData = &bind.MetaData{
	ABI: "[{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"bridgedTokenBeacon\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"bridgedTokenProxyBytecodeHash\",\"type\":\"bytes32\"}],\"name\":\"BridgedTokenBeaconUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"l1Sender\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"l2Receiver\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"l2Token\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"FinalizeDeposit\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"l2TokenBeacon\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"l2TokenProxyBytecodeHash\",\"type\":\"bytes32\"}],\"name\":\"L2TokenBeaconUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"l2Sender\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"l1Receiver\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"l2Token\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"WithdrawalInitiated\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"ASSET_ROUTER\",\"outputs\":[{\"internalType\":\"contractIAssetRouterBase\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"WETH_TOKEN\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"}],\"name\":\"assetId\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_chainId\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"_tokenAddress\",\"type\":\"address\"}],\"name\":\"calculateAssetId\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_originChainId\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"_originToken\",\"type\":\"address\"}],\"name\":\"calculateCreate2TokenAddress\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_originChainId\",\"type\":\"uint256\"}],\"name\":\"getERC20Getters\",\"outputs\":[{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_l1Token\",\"type\":\"address\"}],\"name\":\"l2TokenAddress\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"assetId\",\"type\":\"bytes32\"}],\"name\":\"originChainId\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_l1Token\",\"type\":\"address\"}],\"name\":\"registerToken\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"assetId\",\"type\":\"bytes32\"}],\"name\":\"tokenAddress\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
}

// IL2NativeTokenVaultABI is the input ABI used to generate the binding from.
// Deprecated: Use IL2NativeTokenVaultMetaData.ABI instead.
var IL2NativeTokenVaultABI = IL2NativeTokenVaultMetaData.ABI

// IL2NativeTokenVault is an auto generated Go binding around an Ethereum contract.
type IL2NativeTokenVault struct {
	IL2NativeTokenVaultCaller     // Read-only binding to the contract
	IL2NativeTokenVaultTransactor // Write-only binding to the contract
	IL2NativeTokenVaultFilterer   // Log filterer for contract events
}

// IL2NativeTokenVaultCaller is an auto generated read-only Go binding around an Ethereum contract.
type IL2NativeTokenVaultCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IL2NativeTokenVaultTransactor is an auto generated write-only Go binding around an Ethereum contract.
type IL2NativeTokenVaultTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IL2NativeTokenVaultFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type IL2NativeTokenVaultFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IL2NativeTokenVaultSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type IL2NativeTokenVaultSession struct {
	Contract     *IL2NativeTokenVault // Generic contract binding to set the session for
	CallOpts     bind.CallOpts        // Call options to use throughout this session
	TransactOpts bind.TransactOpts    // Transaction auth options to use throughout this session
}

// IL2NativeTokenVaultCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type IL2NativeTokenVaultCallerSession struct {
	Contract *IL2NativeTokenVaultCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts              // Call options to use throughout this session
}

// IL2NativeTokenVaultTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type IL2NativeTokenVaultTransactorSession struct {
	Contract     *IL2NativeTokenVaultTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts              // Transaction auth options to use throughout this session
}

// IL2NativeTokenVaultRaw is an auto generated low-level Go binding around an Ethereum contract.
type IL2NativeTokenVaultRaw struct {
	Contract *IL2NativeTokenVault // Generic contract binding to access the raw methods on
}

// IL2NativeTokenVaultCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type IL2NativeTokenVaultCallerRaw struct {
	Contract *IL2NativeTokenVaultCaller // Generic read-only contract binding to access the raw methods on
}

// IL2NativeTokenVaultTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type IL2NativeTokenVaultTransactorRaw struct {
	Contract *IL2NativeTokenVaultTransactor // Generic write-only contract binding to access the raw methods on
}

// NewIL2NativeTokenVault creates a new instance of IL2NativeTokenVault, bound to a specific deployed contract.
func NewIL2NativeTokenVault(address common.Address, backend bind.ContractBackend) (*IL2NativeTokenVault, error) {
	contract, err := bindIL2NativeTokenVault(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &IL2NativeTokenVault{IL2NativeTokenVaultCaller: IL2NativeTokenVaultCaller{contract: contract}, IL2NativeTokenVaultTransactor: IL2NativeTokenVaultTransactor{contract: contract}, IL2NativeTokenVaultFilterer: IL2NativeTokenVaultFilterer{contract: contract}}, nil
}

// NewIL2NativeTokenVaultCaller creates a new read-only instance of IL2NativeTokenVault, bound to a specific deployed contract.
func NewIL2NativeTokenVaultCaller(address common.Address, caller bind.ContractCaller) (*IL2NativeTokenVaultCaller, error) {
	contract, err := bindIL2NativeTokenVault(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &IL2NativeTokenVaultCaller{contract: contract}, nil
}

// NewIL2NativeTokenVaultTransactor creates a new write-only instance of IL2NativeTokenVault, bound to a specific deployed contract.
func NewIL2NativeTokenVaultTransactor(address common.Address, transactor bind.ContractTransactor) (*IL2NativeTokenVaultTransactor, error) {
	contract, err := bindIL2NativeTokenVault(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &IL2NativeTokenVaultTransactor{contract: contract}, nil
}

// NewIL2NativeTokenVaultFilterer creates a new log filterer instance of IL2NativeTokenVault, bound to a specific deployed contract.
func NewIL2NativeTokenVaultFilterer(address common.Address, filterer bind.ContractFilterer) (*IL2NativeTokenVaultFilterer, error) {
	contract, err := bindIL2NativeTokenVault(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &IL2NativeTokenVaultFilterer{contract: contract}, nil
}

// bindIL2NativeTokenVault binds a generic wrapper to an already deployed contract.
func bindIL2NativeTokenVault(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := IL2NativeTokenVaultMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IL2NativeTokenVault *IL2NativeTokenVaultRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IL2NativeTokenVault.Contract.IL2NativeTokenVaultCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IL2NativeTokenVault *IL2NativeTokenVaultRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IL2NativeTokenVault.Contract.IL2NativeTokenVaultTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IL2NativeTokenVault *IL2NativeTokenVaultRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IL2NativeTokenVault.Contract.IL2NativeTokenVaultTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IL2NativeTokenVault *IL2NativeTokenVaultCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IL2NativeTokenVault.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IL2NativeTokenVault *IL2NativeTokenVaultTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IL2NativeTokenVault.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IL2NativeTokenVault *IL2NativeTokenVaultTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IL2NativeTokenVault.Contract.contract.Transact(opts, method, params...)
}

// ASSETROUTER is a free data retrieval call binding the contract method 0xc6a70bbb.
//
// Solidity: function ASSET_ROUTER() view returns(address)
func (_IL2NativeTokenVault *IL2NativeTokenVaultCaller) ASSETROUTER(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _IL2NativeTokenVault.contract.Call(opts, &out, "ASSET_ROUTER")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// ASSETROUTER is a free data retrieval call binding the contract method 0xc6a70bbb.
//
// Solidity: function ASSET_ROUTER() view returns(address)
func (_IL2NativeTokenVault *IL2NativeTokenVaultSession) ASSETROUTER() (common.Address, error) {
	return _IL2NativeTokenVault.Contract.ASSETROUTER(&_IL2NativeTokenVault.CallOpts)
}

// ASSETROUTER is a free data retrieval call binding the contract method 0xc6a70bbb.
//
// Solidity: function ASSET_ROUTER() view returns(address)
func (_IL2NativeTokenVault *IL2NativeTokenVaultCallerSession) ASSETROUTER() (common.Address, error) {
	return _IL2NativeTokenVault.Contract.ASSETROUTER(&_IL2NativeTokenVault.CallOpts)
}

// WETHTOKEN is a free data retrieval call binding the contract method 0x37d277d4.
//
// Solidity: function WETH_TOKEN() view returns(address)
func (_IL2NativeTokenVault *IL2NativeTokenVaultCaller) WETHTOKEN(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _IL2NativeTokenVault.contract.Call(opts, &out, "WETH_TOKEN")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// WETHTOKEN is a free data retrieval call binding the contract method 0x37d277d4.
//
// Solidity: function WETH_TOKEN() view returns(address)
func (_IL2NativeTokenVault *IL2NativeTokenVaultSession) WETHTOKEN() (common.Address, error) {
	return _IL2NativeTokenVault.Contract.WETHTOKEN(&_IL2NativeTokenVault.CallOpts)
}

// WETHTOKEN is a free data retrieval call binding the contract method 0x37d277d4.
//
// Solidity: function WETH_TOKEN() view returns(address)
func (_IL2NativeTokenVault *IL2NativeTokenVaultCallerSession) WETHTOKEN() (common.Address, error) {
	return _IL2NativeTokenVault.Contract.WETHTOKEN(&_IL2NativeTokenVault.CallOpts)
}

// AssetId is a free data retrieval call binding the contract method 0xfd3f60df.
//
// Solidity: function assetId(address token) view returns(bytes32)
func (_IL2NativeTokenVault *IL2NativeTokenVaultCaller) AssetId(opts *bind.CallOpts, token common.Address) ([32]byte, error) {
	var out []interface{}
	err := _IL2NativeTokenVault.contract.Call(opts, &out, "assetId", token)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// AssetId is a free data retrieval call binding the contract method 0xfd3f60df.
//
// Solidity: function assetId(address token) view returns(bytes32)
func (_IL2NativeTokenVault *IL2NativeTokenVaultSession) AssetId(token common.Address) ([32]byte, error) {
	return _IL2NativeTokenVault.Contract.AssetId(&_IL2NativeTokenVault.CallOpts, token)
}

// AssetId is a free data retrieval call binding the contract method 0xfd3f60df.
//
// Solidity: function assetId(address token) view returns(bytes32)
func (_IL2NativeTokenVault *IL2NativeTokenVaultCallerSession) AssetId(token common.Address) ([32]byte, error) {
	return _IL2NativeTokenVault.Contract.AssetId(&_IL2NativeTokenVault.CallOpts, token)
}

// CalculateAssetId is a free data retrieval call binding the contract method 0xa42c88a2.
//
// Solidity: function calculateAssetId(uint256 _chainId, address _tokenAddress) view returns(bytes32)
func (_IL2NativeTokenVault *IL2NativeTokenVaultCaller) CalculateAssetId(opts *bind.CallOpts, _chainId *big.Int, _tokenAddress common.Address) ([32]byte, error) {
	var out []interface{}
	err := _IL2NativeTokenVault.contract.Call(opts, &out, "calculateAssetId", _chainId, _tokenAddress)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// CalculateAssetId is a free data retrieval call binding the contract method 0xa42c88a2.
//
// Solidity: function calculateAssetId(uint256 _chainId, address _tokenAddress) view returns(bytes32)
func (_IL2NativeTokenVault *IL2NativeTokenVaultSession) CalculateAssetId(_chainId *big.Int, _tokenAddress common.Address) ([32]byte, error) {
	return _IL2NativeTokenVault.Contract.CalculateAssetId(&_IL2NativeTokenVault.CallOpts, _chainId, _tokenAddress)
}

// CalculateAssetId is a free data retrieval call binding the contract method 0xa42c88a2.
//
// Solidity: function calculateAssetId(uint256 _chainId, address _tokenAddress) view returns(bytes32)
func (_IL2NativeTokenVault *IL2NativeTokenVaultCallerSession) CalculateAssetId(_chainId *big.Int, _tokenAddress common.Address) ([32]byte, error) {
	return _IL2NativeTokenVault.Contract.CalculateAssetId(&_IL2NativeTokenVault.CallOpts, _chainId, _tokenAddress)
}

// CalculateCreate2TokenAddress is a free data retrieval call binding the contract method 0xc487412c.
//
// Solidity: function calculateCreate2TokenAddress(uint256 _originChainId, address _originToken) view returns(address)
func (_IL2NativeTokenVault *IL2NativeTokenVaultCaller) CalculateCreate2TokenAddress(opts *bind.CallOpts, _originChainId *big.Int, _originToken common.Address) (common.Address, error) {
	var out []interface{}
	err := _IL2NativeTokenVault.contract.Call(opts, &out, "calculateCreate2TokenAddress", _originChainId, _originToken)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// CalculateCreate2TokenAddress is a free data retrieval call binding the contract method 0xc487412c.
//
// Solidity: function calculateCreate2TokenAddress(uint256 _originChainId, address _originToken) view returns(address)
func (_IL2NativeTokenVault *IL2NativeTokenVaultSession) CalculateCreate2TokenAddress(_originChainId *big.Int, _originToken common.Address) (common.Address, error) {
	return _IL2NativeTokenVault.Contract.CalculateCreate2TokenAddress(&_IL2NativeTokenVault.CallOpts, _originChainId, _originToken)
}

// CalculateCreate2TokenAddress is a free data retrieval call binding the contract method 0xc487412c.
//
// Solidity: function calculateCreate2TokenAddress(uint256 _originChainId, address _originToken) view returns(address)
func (_IL2NativeTokenVault *IL2NativeTokenVaultCallerSession) CalculateCreate2TokenAddress(_originChainId *big.Int, _originToken common.Address) (common.Address, error) {
	return _IL2NativeTokenVault.Contract.CalculateCreate2TokenAddress(&_IL2NativeTokenVault.CallOpts, _originChainId, _originToken)
}

// GetERC20Getters is a free data retrieval call binding the contract method 0xa7236d16.
//
// Solidity: function getERC20Getters(address _token, uint256 _originChainId) view returns(bytes)
func (_IL2NativeTokenVault *IL2NativeTokenVaultCaller) GetERC20Getters(opts *bind.CallOpts, _token common.Address, _originChainId *big.Int) ([]byte, error) {
	var out []interface{}
	err := _IL2NativeTokenVault.contract.Call(opts, &out, "getERC20Getters", _token, _originChainId)

	if err != nil {
		return *new([]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([]byte)).(*[]byte)

	return out0, err

}

// GetERC20Getters is a free data retrieval call binding the contract method 0xa7236d16.
//
// Solidity: function getERC20Getters(address _token, uint256 _originChainId) view returns(bytes)
func (_IL2NativeTokenVault *IL2NativeTokenVaultSession) GetERC20Getters(_token common.Address, _originChainId *big.Int) ([]byte, error) {
	return _IL2NativeTokenVault.Contract.GetERC20Getters(&_IL2NativeTokenVault.CallOpts, _token, _originChainId)
}

// GetERC20Getters is a free data retrieval call binding the contract method 0xa7236d16.
//
// Solidity: function getERC20Getters(address _token, uint256 _originChainId) view returns(bytes)
func (_IL2NativeTokenVault *IL2NativeTokenVaultCallerSession) GetERC20Getters(_token common.Address, _originChainId *big.Int) ([]byte, error) {
	return _IL2NativeTokenVault.Contract.GetERC20Getters(&_IL2NativeTokenVault.CallOpts, _token, _originChainId)
}

// L2TokenAddress is a free data retrieval call binding the contract method 0xf5f15168.
//
// Solidity: function l2TokenAddress(address _l1Token) view returns(address)
func (_IL2NativeTokenVault *IL2NativeTokenVaultCaller) L2TokenAddress(opts *bind.CallOpts, _l1Token common.Address) (common.Address, error) {
	var out []interface{}
	err := _IL2NativeTokenVault.contract.Call(opts, &out, "l2TokenAddress", _l1Token)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// L2TokenAddress is a free data retrieval call binding the contract method 0xf5f15168.
//
// Solidity: function l2TokenAddress(address _l1Token) view returns(address)
func (_IL2NativeTokenVault *IL2NativeTokenVaultSession) L2TokenAddress(_l1Token common.Address) (common.Address, error) {
	return _IL2NativeTokenVault.Contract.L2TokenAddress(&_IL2NativeTokenVault.CallOpts, _l1Token)
}

// L2TokenAddress is a free data retrieval call binding the contract method 0xf5f15168.
//
// Solidity: function l2TokenAddress(address _l1Token) view returns(address)
func (_IL2NativeTokenVault *IL2NativeTokenVaultCallerSession) L2TokenAddress(_l1Token common.Address) (common.Address, error) {
	return _IL2NativeTokenVault.Contract.L2TokenAddress(&_IL2NativeTokenVault.CallOpts, _l1Token)
}

// OriginChainId is a free data retrieval call binding the contract method 0x5f3455b5.
//
// Solidity: function originChainId(bytes32 assetId) view returns(uint256)
func (_IL2NativeTokenVault *IL2NativeTokenVaultCaller) OriginChainId(opts *bind.CallOpts, assetId [32]byte) (*big.Int, error) {
	var out []interface{}
	err := _IL2NativeTokenVault.contract.Call(opts, &out, "originChainId", assetId)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// OriginChainId is a free data retrieval call binding the contract method 0x5f3455b5.
//
// Solidity: function originChainId(bytes32 assetId) view returns(uint256)
func (_IL2NativeTokenVault *IL2NativeTokenVaultSession) OriginChainId(assetId [32]byte) (*big.Int, error) {
	return _IL2NativeTokenVault.Contract.OriginChainId(&_IL2NativeTokenVault.CallOpts, assetId)
}

// OriginChainId is a free data retrieval call binding the contract method 0x5f3455b5.
//
// Solidity: function originChainId(bytes32 assetId) view returns(uint256)
func (_IL2NativeTokenVault *IL2NativeTokenVaultCallerSession) OriginChainId(assetId [32]byte) (*big.Int, error) {
	return _IL2NativeTokenVault.Contract.OriginChainId(&_IL2NativeTokenVault.CallOpts, assetId)
}

// TokenAddress is a free data retrieval call binding the contract method 0x97bb3ce9.
//
// Solidity: function tokenAddress(bytes32 assetId) view returns(address)
func (_IL2NativeTokenVault *IL2NativeTokenVaultCaller) TokenAddress(opts *bind.CallOpts, assetId [32]byte) (common.Address, error) {
	var out []interface{}
	err := _IL2NativeTokenVault.contract.Call(opts, &out, "tokenAddress", assetId)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// TokenAddress is a free data retrieval call binding the contract method 0x97bb3ce9.
//
// Solidity: function tokenAddress(bytes32 assetId) view returns(address)
func (_IL2NativeTokenVault *IL2NativeTokenVaultSession) TokenAddress(assetId [32]byte) (common.Address, error) {
	return _IL2NativeTokenVault.Contract.TokenAddress(&_IL2NativeTokenVault.CallOpts, assetId)
}

// TokenAddress is a free data retrieval call binding the contract method 0x97bb3ce9.
//
// Solidity: function tokenAddress(bytes32 assetId) view returns(address)
func (_IL2NativeTokenVault *IL2NativeTokenVaultCallerSession) TokenAddress(assetId [32]byte) (common.Address, error) {
	return _IL2NativeTokenVault.Contract.TokenAddress(&_IL2NativeTokenVault.CallOpts, assetId)
}

// RegisterToken is a paid mutator transaction binding the contract method 0x09824a80.
//
// Solidity: function registerToken(address _l1Token) returns()
func (_IL2NativeTokenVault *IL2NativeTokenVaultTransactor) RegisterToken(opts *bind.TransactOpts, _l1Token common.Address) (*types.Transaction, error) {
	return _IL2NativeTokenVault.contract.Transact(opts, "registerToken", _l1Token)
}

// RegisterToken is a paid mutator transaction binding the contract method 0x09824a80.
//
// Solidity: function registerToken(address _l1Token) returns()
func (_IL2NativeTokenVault *IL2NativeTokenVaultSession) RegisterToken(_l1Token common.Address) (*types.Transaction, error) {
	return _IL2NativeTokenVault.Contract.RegisterToken(&_IL2NativeTokenVault.TransactOpts, _l1Token)
}

// RegisterToken is a paid mutator transaction binding the contract method 0x09824a80.
//
// Solidity: function registerToken(address _l1Token) returns()
func (_IL2NativeTokenVault *IL2NativeTokenVaultTransactorSession) RegisterToken(_l1Token common.Address) (*types.Transaction, error) {
	return _IL2NativeTokenVault.Contract.RegisterToken(&_IL2NativeTokenVault.TransactOpts, _l1Token)
}

// IL2NativeTokenVaultBridgedTokenBeaconUpdatedIterator is returned from FilterBridgedTokenBeaconUpdated and is used to iterate over the raw logs and unpacked data for BridgedTokenBeaconUpdated events raised by the IL2NativeTokenVault contract.
type IL2NativeTokenVaultBridgedTokenBeaconUpdatedIterator struct {
	Event *IL2NativeTokenVaultBridgedTokenBeaconUpdated // Event containing the contract specifics and raw log

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
func (it *IL2NativeTokenVaultBridgedTokenBeaconUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IL2NativeTokenVaultBridgedTokenBeaconUpdated)
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
		it.Event = new(IL2NativeTokenVaultBridgedTokenBeaconUpdated)
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
func (it *IL2NativeTokenVaultBridgedTokenBeaconUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IL2NativeTokenVaultBridgedTokenBeaconUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IL2NativeTokenVaultBridgedTokenBeaconUpdated represents a BridgedTokenBeaconUpdated event raised by the IL2NativeTokenVault contract.
type IL2NativeTokenVaultBridgedTokenBeaconUpdated struct {
	BridgedTokenBeacon            common.Address
	BridgedTokenProxyBytecodeHash [32]byte
	Raw                           types.Log // Blockchain specific contextual infos
}

// FilterBridgedTokenBeaconUpdated is a free log retrieval operation binding the contract event 0xc3f14dba68f86c42f518e5c0e8a5cbc9514da6f388e2f52c5b1a6263d8588bfb.
//
// Solidity: event BridgedTokenBeaconUpdated(address bridgedTokenBeacon, bytes32 bridgedTokenProxyBytecodeHash)
func (_IL2NativeTokenVault *IL2NativeTokenVaultFilterer) FilterBridgedTokenBeaconUpdated(opts *bind.FilterOpts) (*IL2NativeTokenVaultBridgedTokenBeaconUpdatedIterator, error) {

	logs, sub, err := _IL2NativeTokenVault.contract.FilterLogs(opts, "BridgedTokenBeaconUpdated")
	if err != nil {
		return nil, err
	}
	return &IL2NativeTokenVaultBridgedTokenBeaconUpdatedIterator{contract: _IL2NativeTokenVault.contract, event: "BridgedTokenBeaconUpdated", logs: logs, sub: sub}, nil
}

// WatchBridgedTokenBeaconUpdated is a free log subscription operation binding the contract event 0xc3f14dba68f86c42f518e5c0e8a5cbc9514da6f388e2f52c5b1a6263d8588bfb.
//
// Solidity: event BridgedTokenBeaconUpdated(address bridgedTokenBeacon, bytes32 bridgedTokenProxyBytecodeHash)
func (_IL2NativeTokenVault *IL2NativeTokenVaultFilterer) WatchBridgedTokenBeaconUpdated(opts *bind.WatchOpts, sink chan<- *IL2NativeTokenVaultBridgedTokenBeaconUpdated) (event.Subscription, error) {

	logs, sub, err := _IL2NativeTokenVault.contract.WatchLogs(opts, "BridgedTokenBeaconUpdated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IL2NativeTokenVaultBridgedTokenBeaconUpdated)
				if err := _IL2NativeTokenVault.contract.UnpackLog(event, "BridgedTokenBeaconUpdated", log); err != nil {
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
func (_IL2NativeTokenVault *IL2NativeTokenVaultFilterer) ParseBridgedTokenBeaconUpdated(log types.Log) (*IL2NativeTokenVaultBridgedTokenBeaconUpdated, error) {
	event := new(IL2NativeTokenVaultBridgedTokenBeaconUpdated)
	if err := _IL2NativeTokenVault.contract.UnpackLog(event, "BridgedTokenBeaconUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// IL2NativeTokenVaultFinalizeDepositIterator is returned from FilterFinalizeDeposit and is used to iterate over the raw logs and unpacked data for FinalizeDeposit events raised by the IL2NativeTokenVault contract.
type IL2NativeTokenVaultFinalizeDepositIterator struct {
	Event *IL2NativeTokenVaultFinalizeDeposit // Event containing the contract specifics and raw log

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
func (it *IL2NativeTokenVaultFinalizeDepositIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IL2NativeTokenVaultFinalizeDeposit)
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
		it.Event = new(IL2NativeTokenVaultFinalizeDeposit)
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
func (it *IL2NativeTokenVaultFinalizeDepositIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IL2NativeTokenVaultFinalizeDepositIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IL2NativeTokenVaultFinalizeDeposit represents a FinalizeDeposit event raised by the IL2NativeTokenVault contract.
type IL2NativeTokenVaultFinalizeDeposit struct {
	L1Sender   common.Address
	L2Receiver common.Address
	L2Token    common.Address
	Amount     *big.Int
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterFinalizeDeposit is a free log retrieval operation binding the contract event 0xb84fba9af218da60d299dc177abd5805e7ac541d2673cbee7808c10017874f63.
//
// Solidity: event FinalizeDeposit(address indexed l1Sender, address indexed l2Receiver, address indexed l2Token, uint256 amount)
func (_IL2NativeTokenVault *IL2NativeTokenVaultFilterer) FilterFinalizeDeposit(opts *bind.FilterOpts, l1Sender []common.Address, l2Receiver []common.Address, l2Token []common.Address) (*IL2NativeTokenVaultFinalizeDepositIterator, error) {

	var l1SenderRule []interface{}
	for _, l1SenderItem := range l1Sender {
		l1SenderRule = append(l1SenderRule, l1SenderItem)
	}
	var l2ReceiverRule []interface{}
	for _, l2ReceiverItem := range l2Receiver {
		l2ReceiverRule = append(l2ReceiverRule, l2ReceiverItem)
	}
	var l2TokenRule []interface{}
	for _, l2TokenItem := range l2Token {
		l2TokenRule = append(l2TokenRule, l2TokenItem)
	}

	logs, sub, err := _IL2NativeTokenVault.contract.FilterLogs(opts, "FinalizeDeposit", l1SenderRule, l2ReceiverRule, l2TokenRule)
	if err != nil {
		return nil, err
	}
	return &IL2NativeTokenVaultFinalizeDepositIterator{contract: _IL2NativeTokenVault.contract, event: "FinalizeDeposit", logs: logs, sub: sub}, nil
}

// WatchFinalizeDeposit is a free log subscription operation binding the contract event 0xb84fba9af218da60d299dc177abd5805e7ac541d2673cbee7808c10017874f63.
//
// Solidity: event FinalizeDeposit(address indexed l1Sender, address indexed l2Receiver, address indexed l2Token, uint256 amount)
func (_IL2NativeTokenVault *IL2NativeTokenVaultFilterer) WatchFinalizeDeposit(opts *bind.WatchOpts, sink chan<- *IL2NativeTokenVaultFinalizeDeposit, l1Sender []common.Address, l2Receiver []common.Address, l2Token []common.Address) (event.Subscription, error) {

	var l1SenderRule []interface{}
	for _, l1SenderItem := range l1Sender {
		l1SenderRule = append(l1SenderRule, l1SenderItem)
	}
	var l2ReceiverRule []interface{}
	for _, l2ReceiverItem := range l2Receiver {
		l2ReceiverRule = append(l2ReceiverRule, l2ReceiverItem)
	}
	var l2TokenRule []interface{}
	for _, l2TokenItem := range l2Token {
		l2TokenRule = append(l2TokenRule, l2TokenItem)
	}

	logs, sub, err := _IL2NativeTokenVault.contract.WatchLogs(opts, "FinalizeDeposit", l1SenderRule, l2ReceiverRule, l2TokenRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IL2NativeTokenVaultFinalizeDeposit)
				if err := _IL2NativeTokenVault.contract.UnpackLog(event, "FinalizeDeposit", log); err != nil {
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

// ParseFinalizeDeposit is a log parse operation binding the contract event 0xb84fba9af218da60d299dc177abd5805e7ac541d2673cbee7808c10017874f63.
//
// Solidity: event FinalizeDeposit(address indexed l1Sender, address indexed l2Receiver, address indexed l2Token, uint256 amount)
func (_IL2NativeTokenVault *IL2NativeTokenVaultFilterer) ParseFinalizeDeposit(log types.Log) (*IL2NativeTokenVaultFinalizeDeposit, error) {
	event := new(IL2NativeTokenVaultFinalizeDeposit)
	if err := _IL2NativeTokenVault.contract.UnpackLog(event, "FinalizeDeposit", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// IL2NativeTokenVaultL2TokenBeaconUpdatedIterator is returned from FilterL2TokenBeaconUpdated and is used to iterate over the raw logs and unpacked data for L2TokenBeaconUpdated events raised by the IL2NativeTokenVault contract.
type IL2NativeTokenVaultL2TokenBeaconUpdatedIterator struct {
	Event *IL2NativeTokenVaultL2TokenBeaconUpdated // Event containing the contract specifics and raw log

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
func (it *IL2NativeTokenVaultL2TokenBeaconUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IL2NativeTokenVaultL2TokenBeaconUpdated)
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
		it.Event = new(IL2NativeTokenVaultL2TokenBeaconUpdated)
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
func (it *IL2NativeTokenVaultL2TokenBeaconUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IL2NativeTokenVaultL2TokenBeaconUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IL2NativeTokenVaultL2TokenBeaconUpdated represents a L2TokenBeaconUpdated event raised by the IL2NativeTokenVault contract.
type IL2NativeTokenVaultL2TokenBeaconUpdated struct {
	L2TokenBeacon            common.Address
	L2TokenProxyBytecodeHash [32]byte
	Raw                      types.Log // Blockchain specific contextual infos
}

// FilterL2TokenBeaconUpdated is a free log retrieval operation binding the contract event 0x01fd5911e6d04aec6b21f19752502ad7f3e9876279643c8fa7a4d30c88a29fb2.
//
// Solidity: event L2TokenBeaconUpdated(address indexed l2TokenBeacon, bytes32 indexed l2TokenProxyBytecodeHash)
func (_IL2NativeTokenVault *IL2NativeTokenVaultFilterer) FilterL2TokenBeaconUpdated(opts *bind.FilterOpts, l2TokenBeacon []common.Address, l2TokenProxyBytecodeHash [][32]byte) (*IL2NativeTokenVaultL2TokenBeaconUpdatedIterator, error) {

	var l2TokenBeaconRule []interface{}
	for _, l2TokenBeaconItem := range l2TokenBeacon {
		l2TokenBeaconRule = append(l2TokenBeaconRule, l2TokenBeaconItem)
	}
	var l2TokenProxyBytecodeHashRule []interface{}
	for _, l2TokenProxyBytecodeHashItem := range l2TokenProxyBytecodeHash {
		l2TokenProxyBytecodeHashRule = append(l2TokenProxyBytecodeHashRule, l2TokenProxyBytecodeHashItem)
	}

	logs, sub, err := _IL2NativeTokenVault.contract.FilterLogs(opts, "L2TokenBeaconUpdated", l2TokenBeaconRule, l2TokenProxyBytecodeHashRule)
	if err != nil {
		return nil, err
	}
	return &IL2NativeTokenVaultL2TokenBeaconUpdatedIterator{contract: _IL2NativeTokenVault.contract, event: "L2TokenBeaconUpdated", logs: logs, sub: sub}, nil
}

// WatchL2TokenBeaconUpdated is a free log subscription operation binding the contract event 0x01fd5911e6d04aec6b21f19752502ad7f3e9876279643c8fa7a4d30c88a29fb2.
//
// Solidity: event L2TokenBeaconUpdated(address indexed l2TokenBeacon, bytes32 indexed l2TokenProxyBytecodeHash)
func (_IL2NativeTokenVault *IL2NativeTokenVaultFilterer) WatchL2TokenBeaconUpdated(opts *bind.WatchOpts, sink chan<- *IL2NativeTokenVaultL2TokenBeaconUpdated, l2TokenBeacon []common.Address, l2TokenProxyBytecodeHash [][32]byte) (event.Subscription, error) {

	var l2TokenBeaconRule []interface{}
	for _, l2TokenBeaconItem := range l2TokenBeacon {
		l2TokenBeaconRule = append(l2TokenBeaconRule, l2TokenBeaconItem)
	}
	var l2TokenProxyBytecodeHashRule []interface{}
	for _, l2TokenProxyBytecodeHashItem := range l2TokenProxyBytecodeHash {
		l2TokenProxyBytecodeHashRule = append(l2TokenProxyBytecodeHashRule, l2TokenProxyBytecodeHashItem)
	}

	logs, sub, err := _IL2NativeTokenVault.contract.WatchLogs(opts, "L2TokenBeaconUpdated", l2TokenBeaconRule, l2TokenProxyBytecodeHashRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IL2NativeTokenVaultL2TokenBeaconUpdated)
				if err := _IL2NativeTokenVault.contract.UnpackLog(event, "L2TokenBeaconUpdated", log); err != nil {
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

// ParseL2TokenBeaconUpdated is a log parse operation binding the contract event 0x01fd5911e6d04aec6b21f19752502ad7f3e9876279643c8fa7a4d30c88a29fb2.
//
// Solidity: event L2TokenBeaconUpdated(address indexed l2TokenBeacon, bytes32 indexed l2TokenProxyBytecodeHash)
func (_IL2NativeTokenVault *IL2NativeTokenVaultFilterer) ParseL2TokenBeaconUpdated(log types.Log) (*IL2NativeTokenVaultL2TokenBeaconUpdated, error) {
	event := new(IL2NativeTokenVaultL2TokenBeaconUpdated)
	if err := _IL2NativeTokenVault.contract.UnpackLog(event, "L2TokenBeaconUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// IL2NativeTokenVaultWithdrawalInitiatedIterator is returned from FilterWithdrawalInitiated and is used to iterate over the raw logs and unpacked data for WithdrawalInitiated events raised by the IL2NativeTokenVault contract.
type IL2NativeTokenVaultWithdrawalInitiatedIterator struct {
	Event *IL2NativeTokenVaultWithdrawalInitiated // Event containing the contract specifics and raw log

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
func (it *IL2NativeTokenVaultWithdrawalInitiatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IL2NativeTokenVaultWithdrawalInitiated)
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
		it.Event = new(IL2NativeTokenVaultWithdrawalInitiated)
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
func (it *IL2NativeTokenVaultWithdrawalInitiatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IL2NativeTokenVaultWithdrawalInitiatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IL2NativeTokenVaultWithdrawalInitiated represents a WithdrawalInitiated event raised by the IL2NativeTokenVault contract.
type IL2NativeTokenVaultWithdrawalInitiated struct {
	L2Sender   common.Address
	L1Receiver common.Address
	L2Token    common.Address
	Amount     *big.Int
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterWithdrawalInitiated is a free log retrieval operation binding the contract event 0x2fc3848834aac8e883a2d2a17a7514dc4f2d3dd268089df9b9f5d918259ef3b0.
//
// Solidity: event WithdrawalInitiated(address indexed l2Sender, address indexed l1Receiver, address indexed l2Token, uint256 amount)
func (_IL2NativeTokenVault *IL2NativeTokenVaultFilterer) FilterWithdrawalInitiated(opts *bind.FilterOpts, l2Sender []common.Address, l1Receiver []common.Address, l2Token []common.Address) (*IL2NativeTokenVaultWithdrawalInitiatedIterator, error) {

	var l2SenderRule []interface{}
	for _, l2SenderItem := range l2Sender {
		l2SenderRule = append(l2SenderRule, l2SenderItem)
	}
	var l1ReceiverRule []interface{}
	for _, l1ReceiverItem := range l1Receiver {
		l1ReceiverRule = append(l1ReceiverRule, l1ReceiverItem)
	}
	var l2TokenRule []interface{}
	for _, l2TokenItem := range l2Token {
		l2TokenRule = append(l2TokenRule, l2TokenItem)
	}

	logs, sub, err := _IL2NativeTokenVault.contract.FilterLogs(opts, "WithdrawalInitiated", l2SenderRule, l1ReceiverRule, l2TokenRule)
	if err != nil {
		return nil, err
	}
	return &IL2NativeTokenVaultWithdrawalInitiatedIterator{contract: _IL2NativeTokenVault.contract, event: "WithdrawalInitiated", logs: logs, sub: sub}, nil
}

// WatchWithdrawalInitiated is a free log subscription operation binding the contract event 0x2fc3848834aac8e883a2d2a17a7514dc4f2d3dd268089df9b9f5d918259ef3b0.
//
// Solidity: event WithdrawalInitiated(address indexed l2Sender, address indexed l1Receiver, address indexed l2Token, uint256 amount)
func (_IL2NativeTokenVault *IL2NativeTokenVaultFilterer) WatchWithdrawalInitiated(opts *bind.WatchOpts, sink chan<- *IL2NativeTokenVaultWithdrawalInitiated, l2Sender []common.Address, l1Receiver []common.Address, l2Token []common.Address) (event.Subscription, error) {

	var l2SenderRule []interface{}
	for _, l2SenderItem := range l2Sender {
		l2SenderRule = append(l2SenderRule, l2SenderItem)
	}
	var l1ReceiverRule []interface{}
	for _, l1ReceiverItem := range l1Receiver {
		l1ReceiverRule = append(l1ReceiverRule, l1ReceiverItem)
	}
	var l2TokenRule []interface{}
	for _, l2TokenItem := range l2Token {
		l2TokenRule = append(l2TokenRule, l2TokenItem)
	}

	logs, sub, err := _IL2NativeTokenVault.contract.WatchLogs(opts, "WithdrawalInitiated", l2SenderRule, l1ReceiverRule, l2TokenRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IL2NativeTokenVaultWithdrawalInitiated)
				if err := _IL2NativeTokenVault.contract.UnpackLog(event, "WithdrawalInitiated", log); err != nil {
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

// ParseWithdrawalInitiated is a log parse operation binding the contract event 0x2fc3848834aac8e883a2d2a17a7514dc4f2d3dd268089df9b9f5d918259ef3b0.
//
// Solidity: event WithdrawalInitiated(address indexed l2Sender, address indexed l1Receiver, address indexed l2Token, uint256 amount)
func (_IL2NativeTokenVault *IL2NativeTokenVaultFilterer) ParseWithdrawalInitiated(log types.Log) (*IL2NativeTokenVaultWithdrawalInitiated, error) {
	event := new(IL2NativeTokenVaultWithdrawalInitiated)
	if err := _IL2NativeTokenVault.contract.UnpackLog(event, "WithdrawalInitiated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
