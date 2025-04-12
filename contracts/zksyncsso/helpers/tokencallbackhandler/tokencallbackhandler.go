// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package tokencallbackhandler

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

// TokenCallbackHandlerMetaData contains all meta data concerning the TokenCallbackHandler contract.
var TokenCallbackHandlerMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"uint256[]\",\"name\":\"\",\"type\":\"uint256[]\"},{\"internalType\":\"uint256[]\",\"name\":\"\",\"type\":\"uint256[]\"},{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"name\":\"onERC1155BatchReceived\",\"outputs\":[{\"internalType\":\"bytes4\",\"name\":\"\",\"type\":\"bytes4\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"name\":\"onERC1155Received\",\"outputs\":[{\"internalType\":\"bytes4\",\"name\":\"\",\"type\":\"bytes4\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"name\":\"onERC721Received\",\"outputs\":[{\"internalType\":\"bytes4\",\"name\":\"\",\"type\":\"bytes4\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes4\",\"name\":\"interfaceId\",\"type\":\"bytes4\"}],\"name\":\"supportsInterface\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
}

// TokenCallbackHandlerABI is the input ABI used to generate the binding from.
// Deprecated: Use TokenCallbackHandlerMetaData.ABI instead.
var TokenCallbackHandlerABI = TokenCallbackHandlerMetaData.ABI

// TokenCallbackHandler is an auto generated Go binding around an Ethereum contract.
type TokenCallbackHandler struct {
	TokenCallbackHandlerCaller     // Read-only binding to the contract
	TokenCallbackHandlerTransactor // Write-only binding to the contract
	TokenCallbackHandlerFilterer   // Log filterer for contract events
}

// TokenCallbackHandlerCaller is an auto generated read-only Go binding around an Ethereum contract.
type TokenCallbackHandlerCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TokenCallbackHandlerTransactor is an auto generated write-only Go binding around an Ethereum contract.
type TokenCallbackHandlerTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TokenCallbackHandlerFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type TokenCallbackHandlerFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TokenCallbackHandlerSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type TokenCallbackHandlerSession struct {
	Contract     *TokenCallbackHandler // Generic contract binding to set the session for
	CallOpts     bind.CallOpts         // Call options to use throughout this session
	TransactOpts bind.TransactOpts     // Transaction auth options to use throughout this session
}

// TokenCallbackHandlerCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type TokenCallbackHandlerCallerSession struct {
	Contract *TokenCallbackHandlerCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts               // Call options to use throughout this session
}

// TokenCallbackHandlerTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type TokenCallbackHandlerTransactorSession struct {
	Contract     *TokenCallbackHandlerTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts               // Transaction auth options to use throughout this session
}

// TokenCallbackHandlerRaw is an auto generated low-level Go binding around an Ethereum contract.
type TokenCallbackHandlerRaw struct {
	Contract *TokenCallbackHandler // Generic contract binding to access the raw methods on
}

// TokenCallbackHandlerCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type TokenCallbackHandlerCallerRaw struct {
	Contract *TokenCallbackHandlerCaller // Generic read-only contract binding to access the raw methods on
}

// TokenCallbackHandlerTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type TokenCallbackHandlerTransactorRaw struct {
	Contract *TokenCallbackHandlerTransactor // Generic write-only contract binding to access the raw methods on
}

// NewTokenCallbackHandler creates a new instance of TokenCallbackHandler, bound to a specific deployed contract.
func NewTokenCallbackHandler(address common.Address, backend bind.ContractBackend) (*TokenCallbackHandler, error) {
	contract, err := bindTokenCallbackHandler(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &TokenCallbackHandler{TokenCallbackHandlerCaller: TokenCallbackHandlerCaller{contract: contract}, TokenCallbackHandlerTransactor: TokenCallbackHandlerTransactor{contract: contract}, TokenCallbackHandlerFilterer: TokenCallbackHandlerFilterer{contract: contract}}, nil
}

// NewTokenCallbackHandlerCaller creates a new read-only instance of TokenCallbackHandler, bound to a specific deployed contract.
func NewTokenCallbackHandlerCaller(address common.Address, caller bind.ContractCaller) (*TokenCallbackHandlerCaller, error) {
	contract, err := bindTokenCallbackHandler(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &TokenCallbackHandlerCaller{contract: contract}, nil
}

// NewTokenCallbackHandlerTransactor creates a new write-only instance of TokenCallbackHandler, bound to a specific deployed contract.
func NewTokenCallbackHandlerTransactor(address common.Address, transactor bind.ContractTransactor) (*TokenCallbackHandlerTransactor, error) {
	contract, err := bindTokenCallbackHandler(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &TokenCallbackHandlerTransactor{contract: contract}, nil
}

// NewTokenCallbackHandlerFilterer creates a new log filterer instance of TokenCallbackHandler, bound to a specific deployed contract.
func NewTokenCallbackHandlerFilterer(address common.Address, filterer bind.ContractFilterer) (*TokenCallbackHandlerFilterer, error) {
	contract, err := bindTokenCallbackHandler(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &TokenCallbackHandlerFilterer{contract: contract}, nil
}

// bindTokenCallbackHandler binds a generic wrapper to an already deployed contract.
func bindTokenCallbackHandler(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := TokenCallbackHandlerMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_TokenCallbackHandler *TokenCallbackHandlerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _TokenCallbackHandler.Contract.TokenCallbackHandlerCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_TokenCallbackHandler *TokenCallbackHandlerRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _TokenCallbackHandler.Contract.TokenCallbackHandlerTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_TokenCallbackHandler *TokenCallbackHandlerRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _TokenCallbackHandler.Contract.TokenCallbackHandlerTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_TokenCallbackHandler *TokenCallbackHandlerCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _TokenCallbackHandler.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_TokenCallbackHandler *TokenCallbackHandlerTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _TokenCallbackHandler.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_TokenCallbackHandler *TokenCallbackHandlerTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _TokenCallbackHandler.Contract.contract.Transact(opts, method, params...)
}

// OnERC1155BatchReceived is a free data retrieval call binding the contract method 0xbc197c81.
//
// Solidity: function onERC1155BatchReceived(address , address , uint256[] , uint256[] , bytes ) pure returns(bytes4)
func (_TokenCallbackHandler *TokenCallbackHandlerCaller) OnERC1155BatchReceived(opts *bind.CallOpts, arg0 common.Address, arg1 common.Address, arg2 []*big.Int, arg3 []*big.Int, arg4 []byte) ([4]byte, error) {
	var out []interface{}
	err := _TokenCallbackHandler.contract.Call(opts, &out, "onERC1155BatchReceived", arg0, arg1, arg2, arg3, arg4)

	if err != nil {
		return *new([4]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([4]byte)).(*[4]byte)

	return out0, err

}

// OnERC1155BatchReceived is a free data retrieval call binding the contract method 0xbc197c81.
//
// Solidity: function onERC1155BatchReceived(address , address , uint256[] , uint256[] , bytes ) pure returns(bytes4)
func (_TokenCallbackHandler *TokenCallbackHandlerSession) OnERC1155BatchReceived(arg0 common.Address, arg1 common.Address, arg2 []*big.Int, arg3 []*big.Int, arg4 []byte) ([4]byte, error) {
	return _TokenCallbackHandler.Contract.OnERC1155BatchReceived(&_TokenCallbackHandler.CallOpts, arg0, arg1, arg2, arg3, arg4)
}

// OnERC1155BatchReceived is a free data retrieval call binding the contract method 0xbc197c81.
//
// Solidity: function onERC1155BatchReceived(address , address , uint256[] , uint256[] , bytes ) pure returns(bytes4)
func (_TokenCallbackHandler *TokenCallbackHandlerCallerSession) OnERC1155BatchReceived(arg0 common.Address, arg1 common.Address, arg2 []*big.Int, arg3 []*big.Int, arg4 []byte) ([4]byte, error) {
	return _TokenCallbackHandler.Contract.OnERC1155BatchReceived(&_TokenCallbackHandler.CallOpts, arg0, arg1, arg2, arg3, arg4)
}

// OnERC1155Received is a free data retrieval call binding the contract method 0xf23a6e61.
//
// Solidity: function onERC1155Received(address , address , uint256 , uint256 , bytes ) pure returns(bytes4)
func (_TokenCallbackHandler *TokenCallbackHandlerCaller) OnERC1155Received(opts *bind.CallOpts, arg0 common.Address, arg1 common.Address, arg2 *big.Int, arg3 *big.Int, arg4 []byte) ([4]byte, error) {
	var out []interface{}
	err := _TokenCallbackHandler.contract.Call(opts, &out, "onERC1155Received", arg0, arg1, arg2, arg3, arg4)

	if err != nil {
		return *new([4]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([4]byte)).(*[4]byte)

	return out0, err

}

// OnERC1155Received is a free data retrieval call binding the contract method 0xf23a6e61.
//
// Solidity: function onERC1155Received(address , address , uint256 , uint256 , bytes ) pure returns(bytes4)
func (_TokenCallbackHandler *TokenCallbackHandlerSession) OnERC1155Received(arg0 common.Address, arg1 common.Address, arg2 *big.Int, arg3 *big.Int, arg4 []byte) ([4]byte, error) {
	return _TokenCallbackHandler.Contract.OnERC1155Received(&_TokenCallbackHandler.CallOpts, arg0, arg1, arg2, arg3, arg4)
}

// OnERC1155Received is a free data retrieval call binding the contract method 0xf23a6e61.
//
// Solidity: function onERC1155Received(address , address , uint256 , uint256 , bytes ) pure returns(bytes4)
func (_TokenCallbackHandler *TokenCallbackHandlerCallerSession) OnERC1155Received(arg0 common.Address, arg1 common.Address, arg2 *big.Int, arg3 *big.Int, arg4 []byte) ([4]byte, error) {
	return _TokenCallbackHandler.Contract.OnERC1155Received(&_TokenCallbackHandler.CallOpts, arg0, arg1, arg2, arg3, arg4)
}

// OnERC721Received is a free data retrieval call binding the contract method 0x150b7a02.
//
// Solidity: function onERC721Received(address , address , uint256 , bytes ) pure returns(bytes4)
func (_TokenCallbackHandler *TokenCallbackHandlerCaller) OnERC721Received(opts *bind.CallOpts, arg0 common.Address, arg1 common.Address, arg2 *big.Int, arg3 []byte) ([4]byte, error) {
	var out []interface{}
	err := _TokenCallbackHandler.contract.Call(opts, &out, "onERC721Received", arg0, arg1, arg2, arg3)

	if err != nil {
		return *new([4]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([4]byte)).(*[4]byte)

	return out0, err

}

// OnERC721Received is a free data retrieval call binding the contract method 0x150b7a02.
//
// Solidity: function onERC721Received(address , address , uint256 , bytes ) pure returns(bytes4)
func (_TokenCallbackHandler *TokenCallbackHandlerSession) OnERC721Received(arg0 common.Address, arg1 common.Address, arg2 *big.Int, arg3 []byte) ([4]byte, error) {
	return _TokenCallbackHandler.Contract.OnERC721Received(&_TokenCallbackHandler.CallOpts, arg0, arg1, arg2, arg3)
}

// OnERC721Received is a free data retrieval call binding the contract method 0x150b7a02.
//
// Solidity: function onERC721Received(address , address , uint256 , bytes ) pure returns(bytes4)
func (_TokenCallbackHandler *TokenCallbackHandlerCallerSession) OnERC721Received(arg0 common.Address, arg1 common.Address, arg2 *big.Int, arg3 []byte) ([4]byte, error) {
	return _TokenCallbackHandler.Contract.OnERC721Received(&_TokenCallbackHandler.CallOpts, arg0, arg1, arg2, arg3)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_TokenCallbackHandler *TokenCallbackHandlerCaller) SupportsInterface(opts *bind.CallOpts, interfaceId [4]byte) (bool, error) {
	var out []interface{}
	err := _TokenCallbackHandler.contract.Call(opts, &out, "supportsInterface", interfaceId)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_TokenCallbackHandler *TokenCallbackHandlerSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _TokenCallbackHandler.Contract.SupportsInterface(&_TokenCallbackHandler.CallOpts, interfaceId)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_TokenCallbackHandler *TokenCallbackHandlerCallerSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _TokenCallbackHandler.Contract.SupportsInterface(&_TokenCallbackHandler.CallOpts, interfaceId)
}
