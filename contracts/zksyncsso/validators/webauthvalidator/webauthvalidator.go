// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package webauthvalidator

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

// WebAuthValidatorMetaData contains all meta data concerning the WebAuthValidator contract.
var WebAuthValidatorMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"name\":\"ACCOUNT_EXISTS\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"BAD_CREDENTIAL_ID_LENGTH\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"BAD_DOMAIN_LENGTH\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"EMPTY_KEY\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"KEY_EXISTS\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"NOT_KEY_OWNER\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"keyOwner\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"originDomain\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"credentialId\",\"type\":\"bytes\"}],\"name\":\"PasskeyCreated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"keyOwner\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"originDomain\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"credentialId\",\"type\":\"bytes\"}],\"name\":\"PasskeyRemoved\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"credentialId\",\"type\":\"bytes\"},{\"internalType\":\"bytes32[2]\",\"name\":\"rawPublicKey\",\"type\":\"bytes32[2]\"},{\"internalType\":\"string\",\"name\":\"originDomain\",\"type\":\"string\"}],\"name\":\"addValidationKey\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"originDomain\",\"type\":\"string\"},{\"internalType\":\"bytes\",\"name\":\"credentialId\",\"type\":\"bytes\"},{\"internalType\":\"address\",\"name\":\"accountAddress\",\"type\":\"address\"}],\"name\":\"getAccountKey\",\"outputs\":[{\"internalType\":\"bytes32[2]\",\"name\":\"\",\"type\":\"bytes32[2]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"onInstall\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"onUninstall\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"originDomain\",\"type\":\"string\"},{\"internalType\":\"bytes\",\"name\":\"credentialId\",\"type\":\"bytes\"}],\"name\":\"registeredAddress\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"accountAddress\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"credentialId\",\"type\":\"bytes\"},{\"internalType\":\"string\",\"name\":\"domain\",\"type\":\"string\"}],\"name\":\"removeValidationKey\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes4\",\"name\":\"interfaceId\",\"type\":\"bytes4\"}],\"name\":\"supportsInterface\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"signedHash\",\"type\":\"bytes32\"},{\"internalType\":\"bytes\",\"name\":\"signature\",\"type\":\"bytes\"}],\"name\":\"validateSignature\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"signedHash\",\"type\":\"bytes32\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"txType\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"from\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"to\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"gasLimit\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"gasPerPubdataByteLimit\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"maxFeePerGas\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"maxPriorityFeePerGas\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"paymaster\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"nonce\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"},{\"internalType\":\"uint256[4]\",\"name\":\"reserved\",\"type\":\"uint256[4]\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"signature\",\"type\":\"bytes\"},{\"internalType\":\"bytes32[]\",\"name\":\"factoryDeps\",\"type\":\"bytes32[]\"},{\"internalType\":\"bytes\",\"name\":\"paymasterInput\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"reservedDynamic\",\"type\":\"bytes\"}],\"internalType\":\"structTransaction\",\"name\":\"transaction\",\"type\":\"tuple\"}],\"name\":\"validateTransaction\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
}

// WebAuthValidatorABI is the input ABI used to generate the binding from.
// Deprecated: Use WebAuthValidatorMetaData.ABI instead.
var WebAuthValidatorABI = WebAuthValidatorMetaData.ABI

// WebAuthValidator is an auto generated Go binding around an Ethereum contract.
type WebAuthValidator struct {
	WebAuthValidatorCaller     // Read-only binding to the contract
	WebAuthValidatorTransactor // Write-only binding to the contract
	WebAuthValidatorFilterer   // Log filterer for contract events
}

// WebAuthValidatorCaller is an auto generated read-only Go binding around an Ethereum contract.
type WebAuthValidatorCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// WebAuthValidatorTransactor is an auto generated write-only Go binding around an Ethereum contract.
type WebAuthValidatorTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// WebAuthValidatorFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type WebAuthValidatorFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// WebAuthValidatorSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type WebAuthValidatorSession struct {
	Contract     *WebAuthValidator // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// WebAuthValidatorCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type WebAuthValidatorCallerSession struct {
	Contract *WebAuthValidatorCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts           // Call options to use throughout this session
}

// WebAuthValidatorTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type WebAuthValidatorTransactorSession struct {
	Contract     *WebAuthValidatorTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts           // Transaction auth options to use throughout this session
}

// WebAuthValidatorRaw is an auto generated low-level Go binding around an Ethereum contract.
type WebAuthValidatorRaw struct {
	Contract *WebAuthValidator // Generic contract binding to access the raw methods on
}

// WebAuthValidatorCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type WebAuthValidatorCallerRaw struct {
	Contract *WebAuthValidatorCaller // Generic read-only contract binding to access the raw methods on
}

// WebAuthValidatorTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type WebAuthValidatorTransactorRaw struct {
	Contract *WebAuthValidatorTransactor // Generic write-only contract binding to access the raw methods on
}

// NewWebAuthValidator creates a new instance of WebAuthValidator, bound to a specific deployed contract.
func NewWebAuthValidator(address common.Address, backend bind.ContractBackend) (*WebAuthValidator, error) {
	contract, err := bindWebAuthValidator(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &WebAuthValidator{WebAuthValidatorCaller: WebAuthValidatorCaller{contract: contract}, WebAuthValidatorTransactor: WebAuthValidatorTransactor{contract: contract}, WebAuthValidatorFilterer: WebAuthValidatorFilterer{contract: contract}}, nil
}

// NewWebAuthValidatorCaller creates a new read-only instance of WebAuthValidator, bound to a specific deployed contract.
func NewWebAuthValidatorCaller(address common.Address, caller bind.ContractCaller) (*WebAuthValidatorCaller, error) {
	contract, err := bindWebAuthValidator(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &WebAuthValidatorCaller{contract: contract}, nil
}

// NewWebAuthValidatorTransactor creates a new write-only instance of WebAuthValidator, bound to a specific deployed contract.
func NewWebAuthValidatorTransactor(address common.Address, transactor bind.ContractTransactor) (*WebAuthValidatorTransactor, error) {
	contract, err := bindWebAuthValidator(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &WebAuthValidatorTransactor{contract: contract}, nil
}

// NewWebAuthValidatorFilterer creates a new log filterer instance of WebAuthValidator, bound to a specific deployed contract.
func NewWebAuthValidatorFilterer(address common.Address, filterer bind.ContractFilterer) (*WebAuthValidatorFilterer, error) {
	contract, err := bindWebAuthValidator(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &WebAuthValidatorFilterer{contract: contract}, nil
}

// bindWebAuthValidator binds a generic wrapper to an already deployed contract.
func bindWebAuthValidator(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := WebAuthValidatorMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_WebAuthValidator *WebAuthValidatorRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _WebAuthValidator.Contract.WebAuthValidatorCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_WebAuthValidator *WebAuthValidatorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _WebAuthValidator.Contract.WebAuthValidatorTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_WebAuthValidator *WebAuthValidatorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _WebAuthValidator.Contract.WebAuthValidatorTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_WebAuthValidator *WebAuthValidatorCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _WebAuthValidator.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_WebAuthValidator *WebAuthValidatorTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _WebAuthValidator.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_WebAuthValidator *WebAuthValidatorTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _WebAuthValidator.Contract.contract.Transact(opts, method, params...)
}

// GetAccountKey is a free data retrieval call binding the contract method 0x80c904d6.
//
// Solidity: function getAccountKey(string originDomain, bytes credentialId, address accountAddress) view returns(bytes32[2])
func (_WebAuthValidator *WebAuthValidatorCaller) GetAccountKey(opts *bind.CallOpts, originDomain string, credentialId []byte, accountAddress common.Address) ([2][32]byte, error) {
	var out []interface{}
	err := _WebAuthValidator.contract.Call(opts, &out, "getAccountKey", originDomain, credentialId, accountAddress)

	if err != nil {
		return *new([2][32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([2][32]byte)).(*[2][32]byte)

	return out0, err

}

// GetAccountKey is a free data retrieval call binding the contract method 0x80c904d6.
//
// Solidity: function getAccountKey(string originDomain, bytes credentialId, address accountAddress) view returns(bytes32[2])
func (_WebAuthValidator *WebAuthValidatorSession) GetAccountKey(originDomain string, credentialId []byte, accountAddress common.Address) ([2][32]byte, error) {
	return _WebAuthValidator.Contract.GetAccountKey(&_WebAuthValidator.CallOpts, originDomain, credentialId, accountAddress)
}

// GetAccountKey is a free data retrieval call binding the contract method 0x80c904d6.
//
// Solidity: function getAccountKey(string originDomain, bytes credentialId, address accountAddress) view returns(bytes32[2])
func (_WebAuthValidator *WebAuthValidatorCallerSession) GetAccountKey(originDomain string, credentialId []byte, accountAddress common.Address) ([2][32]byte, error) {
	return _WebAuthValidator.Contract.GetAccountKey(&_WebAuthValidator.CallOpts, originDomain, credentialId, accountAddress)
}

// RegisteredAddress is a free data retrieval call binding the contract method 0x90c0af30.
//
// Solidity: function registeredAddress(string originDomain, bytes credentialId) view returns(address accountAddress)
func (_WebAuthValidator *WebAuthValidatorCaller) RegisteredAddress(opts *bind.CallOpts, originDomain string, credentialId []byte) (common.Address, error) {
	var out []interface{}
	err := _WebAuthValidator.contract.Call(opts, &out, "registeredAddress", originDomain, credentialId)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// RegisteredAddress is a free data retrieval call binding the contract method 0x90c0af30.
//
// Solidity: function registeredAddress(string originDomain, bytes credentialId) view returns(address accountAddress)
func (_WebAuthValidator *WebAuthValidatorSession) RegisteredAddress(originDomain string, credentialId []byte) (common.Address, error) {
	return _WebAuthValidator.Contract.RegisteredAddress(&_WebAuthValidator.CallOpts, originDomain, credentialId)
}

// RegisteredAddress is a free data retrieval call binding the contract method 0x90c0af30.
//
// Solidity: function registeredAddress(string originDomain, bytes credentialId) view returns(address accountAddress)
func (_WebAuthValidator *WebAuthValidatorCallerSession) RegisteredAddress(originDomain string, credentialId []byte) (common.Address, error) {
	return _WebAuthValidator.Contract.RegisteredAddress(&_WebAuthValidator.CallOpts, originDomain, credentialId)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) pure returns(bool)
func (_WebAuthValidator *WebAuthValidatorCaller) SupportsInterface(opts *bind.CallOpts, interfaceId [4]byte) (bool, error) {
	var out []interface{}
	err := _WebAuthValidator.contract.Call(opts, &out, "supportsInterface", interfaceId)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) pure returns(bool)
func (_WebAuthValidator *WebAuthValidatorSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _WebAuthValidator.Contract.SupportsInterface(&_WebAuthValidator.CallOpts, interfaceId)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) pure returns(bool)
func (_WebAuthValidator *WebAuthValidatorCallerSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _WebAuthValidator.Contract.SupportsInterface(&_WebAuthValidator.CallOpts, interfaceId)
}

// ValidateSignature is a free data retrieval call binding the contract method 0x333daf92.
//
// Solidity: function validateSignature(bytes32 signedHash, bytes signature) view returns(bool)
func (_WebAuthValidator *WebAuthValidatorCaller) ValidateSignature(opts *bind.CallOpts, signedHash [32]byte, signature []byte) (bool, error) {
	var out []interface{}
	err := _WebAuthValidator.contract.Call(opts, &out, "validateSignature", signedHash, signature)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// ValidateSignature is a free data retrieval call binding the contract method 0x333daf92.
//
// Solidity: function validateSignature(bytes32 signedHash, bytes signature) view returns(bool)
func (_WebAuthValidator *WebAuthValidatorSession) ValidateSignature(signedHash [32]byte, signature []byte) (bool, error) {
	return _WebAuthValidator.Contract.ValidateSignature(&_WebAuthValidator.CallOpts, signedHash, signature)
}

// ValidateSignature is a free data retrieval call binding the contract method 0x333daf92.
//
// Solidity: function validateSignature(bytes32 signedHash, bytes signature) view returns(bool)
func (_WebAuthValidator *WebAuthValidatorCallerSession) ValidateSignature(signedHash [32]byte, signature []byte) (bool, error) {
	return _WebAuthValidator.Contract.ValidateSignature(&_WebAuthValidator.CallOpts, signedHash, signature)
}

// ValidateTransaction is a free data retrieval call binding the contract method 0xbf1733f9.
//
// Solidity: function validateTransaction(bytes32 signedHash, (uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256[4],bytes,bytes,bytes32[],bytes,bytes) transaction) view returns(bool)
func (_WebAuthValidator *WebAuthValidatorCaller) ValidateTransaction(opts *bind.CallOpts, signedHash [32]byte, transaction Transaction) (bool, error) {
	var out []interface{}
	err := _WebAuthValidator.contract.Call(opts, &out, "validateTransaction", signedHash, transaction)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// ValidateTransaction is a free data retrieval call binding the contract method 0xbf1733f9.
//
// Solidity: function validateTransaction(bytes32 signedHash, (uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256[4],bytes,bytes,bytes32[],bytes,bytes) transaction) view returns(bool)
func (_WebAuthValidator *WebAuthValidatorSession) ValidateTransaction(signedHash [32]byte, transaction Transaction) (bool, error) {
	return _WebAuthValidator.Contract.ValidateTransaction(&_WebAuthValidator.CallOpts, signedHash, transaction)
}

// ValidateTransaction is a free data retrieval call binding the contract method 0xbf1733f9.
//
// Solidity: function validateTransaction(bytes32 signedHash, (uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256[4],bytes,bytes,bytes32[],bytes,bytes) transaction) view returns(bool)
func (_WebAuthValidator *WebAuthValidatorCallerSession) ValidateTransaction(signedHash [32]byte, transaction Transaction) (bool, error) {
	return _WebAuthValidator.Contract.ValidateTransaction(&_WebAuthValidator.CallOpts, signedHash, transaction)
}

// AddValidationKey is a paid mutator transaction binding the contract method 0x8f5bc2af.
//
// Solidity: function addValidationKey(bytes credentialId, bytes32[2] rawPublicKey, string originDomain) returns()
func (_WebAuthValidator *WebAuthValidatorTransactor) AddValidationKey(opts *bind.TransactOpts, credentialId []byte, rawPublicKey [2][32]byte, originDomain string) (*types.Transaction, error) {
	return _WebAuthValidator.contract.Transact(opts, "addValidationKey", credentialId, rawPublicKey, originDomain)
}

// AddValidationKey is a paid mutator transaction binding the contract method 0x8f5bc2af.
//
// Solidity: function addValidationKey(bytes credentialId, bytes32[2] rawPublicKey, string originDomain) returns()
func (_WebAuthValidator *WebAuthValidatorSession) AddValidationKey(credentialId []byte, rawPublicKey [2][32]byte, originDomain string) (*types.Transaction, error) {
	return _WebAuthValidator.Contract.AddValidationKey(&_WebAuthValidator.TransactOpts, credentialId, rawPublicKey, originDomain)
}

// AddValidationKey is a paid mutator transaction binding the contract method 0x8f5bc2af.
//
// Solidity: function addValidationKey(bytes credentialId, bytes32[2] rawPublicKey, string originDomain) returns()
func (_WebAuthValidator *WebAuthValidatorTransactorSession) AddValidationKey(credentialId []byte, rawPublicKey [2][32]byte, originDomain string) (*types.Transaction, error) {
	return _WebAuthValidator.Contract.AddValidationKey(&_WebAuthValidator.TransactOpts, credentialId, rawPublicKey, originDomain)
}

// OnInstall is a paid mutator transaction binding the contract method 0x6d61fe70.
//
// Solidity: function onInstall(bytes data) returns()
func (_WebAuthValidator *WebAuthValidatorTransactor) OnInstall(opts *bind.TransactOpts, data []byte) (*types.Transaction, error) {
	return _WebAuthValidator.contract.Transact(opts, "onInstall", data)
}

// OnInstall is a paid mutator transaction binding the contract method 0x6d61fe70.
//
// Solidity: function onInstall(bytes data) returns()
func (_WebAuthValidator *WebAuthValidatorSession) OnInstall(data []byte) (*types.Transaction, error) {
	return _WebAuthValidator.Contract.OnInstall(&_WebAuthValidator.TransactOpts, data)
}

// OnInstall is a paid mutator transaction binding the contract method 0x6d61fe70.
//
// Solidity: function onInstall(bytes data) returns()
func (_WebAuthValidator *WebAuthValidatorTransactorSession) OnInstall(data []byte) (*types.Transaction, error) {
	return _WebAuthValidator.Contract.OnInstall(&_WebAuthValidator.TransactOpts, data)
}

// OnUninstall is a paid mutator transaction binding the contract method 0x8a91b0e3.
//
// Solidity: function onUninstall(bytes data) returns()
func (_WebAuthValidator *WebAuthValidatorTransactor) OnUninstall(opts *bind.TransactOpts, data []byte) (*types.Transaction, error) {
	return _WebAuthValidator.contract.Transact(opts, "onUninstall", data)
}

// OnUninstall is a paid mutator transaction binding the contract method 0x8a91b0e3.
//
// Solidity: function onUninstall(bytes data) returns()
func (_WebAuthValidator *WebAuthValidatorSession) OnUninstall(data []byte) (*types.Transaction, error) {
	return _WebAuthValidator.Contract.OnUninstall(&_WebAuthValidator.TransactOpts, data)
}

// OnUninstall is a paid mutator transaction binding the contract method 0x8a91b0e3.
//
// Solidity: function onUninstall(bytes data) returns()
func (_WebAuthValidator *WebAuthValidatorTransactorSession) OnUninstall(data []byte) (*types.Transaction, error) {
	return _WebAuthValidator.Contract.OnUninstall(&_WebAuthValidator.TransactOpts, data)
}

// RemoveValidationKey is a paid mutator transaction binding the contract method 0x59be375f.
//
// Solidity: function removeValidationKey(bytes credentialId, string domain) returns()
func (_WebAuthValidator *WebAuthValidatorTransactor) RemoveValidationKey(opts *bind.TransactOpts, credentialId []byte, domain string) (*types.Transaction, error) {
	return _WebAuthValidator.contract.Transact(opts, "removeValidationKey", credentialId, domain)
}

// RemoveValidationKey is a paid mutator transaction binding the contract method 0x59be375f.
//
// Solidity: function removeValidationKey(bytes credentialId, string domain) returns()
func (_WebAuthValidator *WebAuthValidatorSession) RemoveValidationKey(credentialId []byte, domain string) (*types.Transaction, error) {
	return _WebAuthValidator.Contract.RemoveValidationKey(&_WebAuthValidator.TransactOpts, credentialId, domain)
}

// RemoveValidationKey is a paid mutator transaction binding the contract method 0x59be375f.
//
// Solidity: function removeValidationKey(bytes credentialId, string domain) returns()
func (_WebAuthValidator *WebAuthValidatorTransactorSession) RemoveValidationKey(credentialId []byte, domain string) (*types.Transaction, error) {
	return _WebAuthValidator.Contract.RemoveValidationKey(&_WebAuthValidator.TransactOpts, credentialId, domain)
}

// WebAuthValidatorPasskeyCreatedIterator is returned from FilterPasskeyCreated and is used to iterate over the raw logs and unpacked data for PasskeyCreated events raised by the WebAuthValidator contract.
type WebAuthValidatorPasskeyCreatedIterator struct {
	Event *WebAuthValidatorPasskeyCreated // Event containing the contract specifics and raw log

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
func (it *WebAuthValidatorPasskeyCreatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(WebAuthValidatorPasskeyCreated)
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
		it.Event = new(WebAuthValidatorPasskeyCreated)
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
func (it *WebAuthValidatorPasskeyCreatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *WebAuthValidatorPasskeyCreatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// WebAuthValidatorPasskeyCreated represents a PasskeyCreated event raised by the WebAuthValidator contract.
type WebAuthValidatorPasskeyCreated struct {
	KeyOwner     common.Address
	OriginDomain string
	CredentialId []byte
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterPasskeyCreated is a free log retrieval operation binding the contract event 0x3c6e30f99b4a1e8aae9cdcfbc2d6683bf8347bb160ae280eb95513567c07b473.
//
// Solidity: event PasskeyCreated(address indexed keyOwner, string originDomain, bytes credentialId)
func (_WebAuthValidator *WebAuthValidatorFilterer) FilterPasskeyCreated(opts *bind.FilterOpts, keyOwner []common.Address) (*WebAuthValidatorPasskeyCreatedIterator, error) {

	var keyOwnerRule []interface{}
	for _, keyOwnerItem := range keyOwner {
		keyOwnerRule = append(keyOwnerRule, keyOwnerItem)
	}

	logs, sub, err := _WebAuthValidator.contract.FilterLogs(opts, "PasskeyCreated", keyOwnerRule)
	if err != nil {
		return nil, err
	}
	return &WebAuthValidatorPasskeyCreatedIterator{contract: _WebAuthValidator.contract, event: "PasskeyCreated", logs: logs, sub: sub}, nil
}

// WatchPasskeyCreated is a free log subscription operation binding the contract event 0x3c6e30f99b4a1e8aae9cdcfbc2d6683bf8347bb160ae280eb95513567c07b473.
//
// Solidity: event PasskeyCreated(address indexed keyOwner, string originDomain, bytes credentialId)
func (_WebAuthValidator *WebAuthValidatorFilterer) WatchPasskeyCreated(opts *bind.WatchOpts, sink chan<- *WebAuthValidatorPasskeyCreated, keyOwner []common.Address) (event.Subscription, error) {

	var keyOwnerRule []interface{}
	for _, keyOwnerItem := range keyOwner {
		keyOwnerRule = append(keyOwnerRule, keyOwnerItem)
	}

	logs, sub, err := _WebAuthValidator.contract.WatchLogs(opts, "PasskeyCreated", keyOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(WebAuthValidatorPasskeyCreated)
				if err := _WebAuthValidator.contract.UnpackLog(event, "PasskeyCreated", log); err != nil {
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

// ParsePasskeyCreated is a log parse operation binding the contract event 0x3c6e30f99b4a1e8aae9cdcfbc2d6683bf8347bb160ae280eb95513567c07b473.
//
// Solidity: event PasskeyCreated(address indexed keyOwner, string originDomain, bytes credentialId)
func (_WebAuthValidator *WebAuthValidatorFilterer) ParsePasskeyCreated(log types.Log) (*WebAuthValidatorPasskeyCreated, error) {
	event := new(WebAuthValidatorPasskeyCreated)
	if err := _WebAuthValidator.contract.UnpackLog(event, "PasskeyCreated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// WebAuthValidatorPasskeyRemovedIterator is returned from FilterPasskeyRemoved and is used to iterate over the raw logs and unpacked data for PasskeyRemoved events raised by the WebAuthValidator contract.
type WebAuthValidatorPasskeyRemovedIterator struct {
	Event *WebAuthValidatorPasskeyRemoved // Event containing the contract specifics and raw log

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
func (it *WebAuthValidatorPasskeyRemovedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(WebAuthValidatorPasskeyRemoved)
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
		it.Event = new(WebAuthValidatorPasskeyRemoved)
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
func (it *WebAuthValidatorPasskeyRemovedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *WebAuthValidatorPasskeyRemovedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// WebAuthValidatorPasskeyRemoved represents a PasskeyRemoved event raised by the WebAuthValidator contract.
type WebAuthValidatorPasskeyRemoved struct {
	KeyOwner     common.Address
	OriginDomain string
	CredentialId []byte
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterPasskeyRemoved is a free log retrieval operation binding the contract event 0x60a29a47798fec901f6df6fe2a86177130b82aecf219f6c7960182b8a0627636.
//
// Solidity: event PasskeyRemoved(address indexed keyOwner, string originDomain, bytes credentialId)
func (_WebAuthValidator *WebAuthValidatorFilterer) FilterPasskeyRemoved(opts *bind.FilterOpts, keyOwner []common.Address) (*WebAuthValidatorPasskeyRemovedIterator, error) {

	var keyOwnerRule []interface{}
	for _, keyOwnerItem := range keyOwner {
		keyOwnerRule = append(keyOwnerRule, keyOwnerItem)
	}

	logs, sub, err := _WebAuthValidator.contract.FilterLogs(opts, "PasskeyRemoved", keyOwnerRule)
	if err != nil {
		return nil, err
	}
	return &WebAuthValidatorPasskeyRemovedIterator{contract: _WebAuthValidator.contract, event: "PasskeyRemoved", logs: logs, sub: sub}, nil
}

// WatchPasskeyRemoved is a free log subscription operation binding the contract event 0x60a29a47798fec901f6df6fe2a86177130b82aecf219f6c7960182b8a0627636.
//
// Solidity: event PasskeyRemoved(address indexed keyOwner, string originDomain, bytes credentialId)
func (_WebAuthValidator *WebAuthValidatorFilterer) WatchPasskeyRemoved(opts *bind.WatchOpts, sink chan<- *WebAuthValidatorPasskeyRemoved, keyOwner []common.Address) (event.Subscription, error) {

	var keyOwnerRule []interface{}
	for _, keyOwnerItem := range keyOwner {
		keyOwnerRule = append(keyOwnerRule, keyOwnerItem)
	}

	logs, sub, err := _WebAuthValidator.contract.WatchLogs(opts, "PasskeyRemoved", keyOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(WebAuthValidatorPasskeyRemoved)
				if err := _WebAuthValidator.contract.UnpackLog(event, "PasskeyRemoved", log); err != nil {
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

// ParsePasskeyRemoved is a log parse operation binding the contract event 0x60a29a47798fec901f6df6fe2a86177130b82aecf219f6c7960182b8a0627636.
//
// Solidity: event PasskeyRemoved(address indexed keyOwner, string originDomain, bytes credentialId)
func (_WebAuthValidator *WebAuthValidatorFilterer) ParsePasskeyRemoved(log types.Log) (*WebAuthValidatorPasskeyRemoved, error) {
	event := new(WebAuthValidatorPasskeyRemoved)
	if err := _WebAuthValidator.contract.UnpackLog(event, "PasskeyRemoved", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
