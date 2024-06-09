// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package contractdeployer

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

// IContractDeployerAccountInfo is an auto generated low-level Go binding around an user-defined struct.
type IContractDeployerAccountInfo struct {
	SupportedAAVersion uint8
	NonceOrdering      uint8
}

// IContractDeployerMetaData contains all meta data concerning the IContractDeployer contract.
var IContractDeployerMetaData = &bind.MetaData{
	ABI: "[{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"accountAddress\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"enumIContractDeployer.AccountNonceOrdering\",\"name\":\"nonceOrdering\",\"type\":\"uint8\"}],\"name\":\"AccountNonceOrderingUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"accountAddress\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"enumIContractDeployer.AccountAbstractionVersion\",\"name\":\"aaVersion\",\"type\":\"uint8\"}],\"name\":\"AccountVersionUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"deployerAddress\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"bytecodeHash\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"contractAddress\",\"type\":\"address\"}],\"name\":\"ContractDeployed\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_salt\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"_bytecodeHash\",\"type\":\"bytes32\"},{\"internalType\":\"bytes\",\"name\":\"_input\",\"type\":\"bytes\"}],\"name\":\"create\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"newAddress\",\"type\":\"address\"}],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_salt\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"_bytecodeHash\",\"type\":\"bytes32\"},{\"internalType\":\"bytes\",\"name\":\"_input\",\"type\":\"bytes\"}],\"name\":\"create2\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"newAddress\",\"type\":\"address\"}],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_salt\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"_bytecodeHash\",\"type\":\"bytes32\"},{\"internalType\":\"bytes\",\"name\":\"_input\",\"type\":\"bytes\"},{\"internalType\":\"enumIContractDeployer.AccountAbstractionVersion\",\"name\":\"_aaVersion\",\"type\":\"uint8\"}],\"name\":\"create2Account\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"newAddress\",\"type\":\"address\"}],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_salt\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"_bytecodeHash\",\"type\":\"bytes32\"},{\"internalType\":\"bytes\",\"name\":\"_input\",\"type\":\"bytes\"},{\"internalType\":\"enumIContractDeployer.AccountAbstractionVersion\",\"name\":\"_aaVersion\",\"type\":\"uint8\"}],\"name\":\"createAccount\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"newAddress\",\"type\":\"address\"}],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_address\",\"type\":\"address\"}],\"name\":\"getAccountInfo\",\"outputs\":[{\"components\":[{\"internalType\":\"enumIContractDeployer.AccountAbstractionVersion\",\"name\":\"supportedAAVersion\",\"type\":\"uint8\"},{\"internalType\":\"enumIContractDeployer.AccountNonceOrdering\",\"name\":\"nonceOrdering\",\"type\":\"uint8\"}],\"internalType\":\"structIContractDeployer.AccountInfo\",\"name\":\"info\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_sender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_senderNonce\",\"type\":\"uint256\"}],\"name\":\"getNewAddressCreate\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"newAddress\",\"type\":\"address\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_sender\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"_bytecodeHash\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"_salt\",\"type\":\"bytes32\"},{\"internalType\":\"bytes\",\"name\":\"_input\",\"type\":\"bytes\"}],\"name\":\"getNewAddressCreate2\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"newAddress\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"enumIContractDeployer.AccountAbstractionVersion\",\"name\":\"_version\",\"type\":\"uint8\"}],\"name\":\"updateAccountVersion\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"enumIContractDeployer.AccountNonceOrdering\",\"name\":\"_nonceOrdering\",\"type\":\"uint8\"}],\"name\":\"updateNonceOrdering\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
}

// IContractDeployerABI is the input ABI used to generate the binding from.
// Deprecated: Use IContractDeployerMetaData.ABI instead.
var IContractDeployerABI = IContractDeployerMetaData.ABI

// IContractDeployer is an auto generated Go binding around an Ethereum contract.
type IContractDeployer struct {
	IContractDeployerCaller     // Read-only binding to the contract
	IContractDeployerTransactor // Write-only binding to the contract
	IContractDeployerFilterer   // Log filterer for contract events
}

// IContractDeployerCaller is an auto generated read-only Go binding around an Ethereum contract.
type IContractDeployerCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IContractDeployerTransactor is an auto generated write-only Go binding around an Ethereum contract.
type IContractDeployerTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IContractDeployerFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type IContractDeployerFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IContractDeployerSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type IContractDeployerSession struct {
	Contract     *IContractDeployer // Generic contract binding to set the session for
	CallOpts     bind.CallOpts      // Call options to use throughout this session
	TransactOpts bind.TransactOpts  // Transaction auth options to use throughout this session
}

// IContractDeployerCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type IContractDeployerCallerSession struct {
	Contract *IContractDeployerCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts            // Call options to use throughout this session
}

// IContractDeployerTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type IContractDeployerTransactorSession struct {
	Contract     *IContractDeployerTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts            // Transaction auth options to use throughout this session
}

// IContractDeployerRaw is an auto generated low-level Go binding around an Ethereum contract.
type IContractDeployerRaw struct {
	Contract *IContractDeployer // Generic contract binding to access the raw methods on
}

// IContractDeployerCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type IContractDeployerCallerRaw struct {
	Contract *IContractDeployerCaller // Generic read-only contract binding to access the raw methods on
}

// IContractDeployerTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type IContractDeployerTransactorRaw struct {
	Contract *IContractDeployerTransactor // Generic write-only contract binding to access the raw methods on
}

// NewIContractDeployer creates a new instance of IContractDeployer, bound to a specific deployed contract.
func NewIContractDeployer(address common.Address, backend bind.ContractBackend) (*IContractDeployer, error) {
	contract, err := bindIContractDeployer(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &IContractDeployer{IContractDeployerCaller: IContractDeployerCaller{contract: contract}, IContractDeployerTransactor: IContractDeployerTransactor{contract: contract}, IContractDeployerFilterer: IContractDeployerFilterer{contract: contract}}, nil
}

// NewIContractDeployerCaller creates a new read-only instance of IContractDeployer, bound to a specific deployed contract.
func NewIContractDeployerCaller(address common.Address, caller bind.ContractCaller) (*IContractDeployerCaller, error) {
	contract, err := bindIContractDeployer(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &IContractDeployerCaller{contract: contract}, nil
}

// NewIContractDeployerTransactor creates a new write-only instance of IContractDeployer, bound to a specific deployed contract.
func NewIContractDeployerTransactor(address common.Address, transactor bind.ContractTransactor) (*IContractDeployerTransactor, error) {
	contract, err := bindIContractDeployer(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &IContractDeployerTransactor{contract: contract}, nil
}

// NewIContractDeployerFilterer creates a new log filterer instance of IContractDeployer, bound to a specific deployed contract.
func NewIContractDeployerFilterer(address common.Address, filterer bind.ContractFilterer) (*IContractDeployerFilterer, error) {
	contract, err := bindIContractDeployer(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &IContractDeployerFilterer{contract: contract}, nil
}

// bindIContractDeployer binds a generic wrapper to an already deployed contract.
func bindIContractDeployer(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := IContractDeployerMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IContractDeployer *IContractDeployerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IContractDeployer.Contract.IContractDeployerCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IContractDeployer *IContractDeployerRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IContractDeployer.Contract.IContractDeployerTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IContractDeployer *IContractDeployerRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IContractDeployer.Contract.IContractDeployerTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IContractDeployer *IContractDeployerCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IContractDeployer.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IContractDeployer *IContractDeployerTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IContractDeployer.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IContractDeployer *IContractDeployerTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IContractDeployer.Contract.contract.Transact(opts, method, params...)
}

// GetAccountInfo is a free data retrieval call binding the contract method 0x7b510fe8.
//
// Solidity: function getAccountInfo(address _address) view returns((uint8,uint8) info)
func (_IContractDeployer *IContractDeployerCaller) GetAccountInfo(opts *bind.CallOpts, _address common.Address) (IContractDeployerAccountInfo, error) {
	var out []interface{}
	err := _IContractDeployer.contract.Call(opts, &out, "getAccountInfo", _address)

	if err != nil {
		return *new(IContractDeployerAccountInfo), err
	}

	out0 := *abi.ConvertType(out[0], new(IContractDeployerAccountInfo)).(*IContractDeployerAccountInfo)

	return out0, err

}

// GetAccountInfo is a free data retrieval call binding the contract method 0x7b510fe8.
//
// Solidity: function getAccountInfo(address _address) view returns((uint8,uint8) info)
func (_IContractDeployer *IContractDeployerSession) GetAccountInfo(_address common.Address) (IContractDeployerAccountInfo, error) {
	return _IContractDeployer.Contract.GetAccountInfo(&_IContractDeployer.CallOpts, _address)
}

// GetAccountInfo is a free data retrieval call binding the contract method 0x7b510fe8.
//
// Solidity: function getAccountInfo(address _address) view returns((uint8,uint8) info)
func (_IContractDeployer *IContractDeployerCallerSession) GetAccountInfo(_address common.Address) (IContractDeployerAccountInfo, error) {
	return _IContractDeployer.Contract.GetAccountInfo(&_IContractDeployer.CallOpts, _address)
}

// GetNewAddressCreate is a free data retrieval call binding the contract method 0x187598a5.
//
// Solidity: function getNewAddressCreate(address _sender, uint256 _senderNonce) pure returns(address newAddress)
func (_IContractDeployer *IContractDeployerCaller) GetNewAddressCreate(opts *bind.CallOpts, _sender common.Address, _senderNonce *big.Int) (common.Address, error) {
	var out []interface{}
	err := _IContractDeployer.contract.Call(opts, &out, "getNewAddressCreate", _sender, _senderNonce)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetNewAddressCreate is a free data retrieval call binding the contract method 0x187598a5.
//
// Solidity: function getNewAddressCreate(address _sender, uint256 _senderNonce) pure returns(address newAddress)
func (_IContractDeployer *IContractDeployerSession) GetNewAddressCreate(_sender common.Address, _senderNonce *big.Int) (common.Address, error) {
	return _IContractDeployer.Contract.GetNewAddressCreate(&_IContractDeployer.CallOpts, _sender, _senderNonce)
}

// GetNewAddressCreate is a free data retrieval call binding the contract method 0x187598a5.
//
// Solidity: function getNewAddressCreate(address _sender, uint256 _senderNonce) pure returns(address newAddress)
func (_IContractDeployer *IContractDeployerCallerSession) GetNewAddressCreate(_sender common.Address, _senderNonce *big.Int) (common.Address, error) {
	return _IContractDeployer.Contract.GetNewAddressCreate(&_IContractDeployer.CallOpts, _sender, _senderNonce)
}

// GetNewAddressCreate2 is a free data retrieval call binding the contract method 0x84da1fb4.
//
// Solidity: function getNewAddressCreate2(address _sender, bytes32 _bytecodeHash, bytes32 _salt, bytes _input) view returns(address newAddress)
func (_IContractDeployer *IContractDeployerCaller) GetNewAddressCreate2(opts *bind.CallOpts, _sender common.Address, _bytecodeHash [32]byte, _salt [32]byte, _input []byte) (common.Address, error) {
	var out []interface{}
	err := _IContractDeployer.contract.Call(opts, &out, "getNewAddressCreate2", _sender, _bytecodeHash, _salt, _input)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetNewAddressCreate2 is a free data retrieval call binding the contract method 0x84da1fb4.
//
// Solidity: function getNewAddressCreate2(address _sender, bytes32 _bytecodeHash, bytes32 _salt, bytes _input) view returns(address newAddress)
func (_IContractDeployer *IContractDeployerSession) GetNewAddressCreate2(_sender common.Address, _bytecodeHash [32]byte, _salt [32]byte, _input []byte) (common.Address, error) {
	return _IContractDeployer.Contract.GetNewAddressCreate2(&_IContractDeployer.CallOpts, _sender, _bytecodeHash, _salt, _input)
}

// GetNewAddressCreate2 is a free data retrieval call binding the contract method 0x84da1fb4.
//
// Solidity: function getNewAddressCreate2(address _sender, bytes32 _bytecodeHash, bytes32 _salt, bytes _input) view returns(address newAddress)
func (_IContractDeployer *IContractDeployerCallerSession) GetNewAddressCreate2(_sender common.Address, _bytecodeHash [32]byte, _salt [32]byte, _input []byte) (common.Address, error) {
	return _IContractDeployer.Contract.GetNewAddressCreate2(&_IContractDeployer.CallOpts, _sender, _bytecodeHash, _salt, _input)
}

// Create is a paid mutator transaction binding the contract method 0x9c4d535b.
//
// Solidity: function create(bytes32 _salt, bytes32 _bytecodeHash, bytes _input) payable returns(address newAddress)
func (_IContractDeployer *IContractDeployerTransactor) Create(opts *bind.TransactOpts, _salt [32]byte, _bytecodeHash [32]byte, _input []byte) (*types.Transaction, error) {
	return _IContractDeployer.contract.Transact(opts, "create", _salt, _bytecodeHash, _input)
}

// Create is a paid mutator transaction binding the contract method 0x9c4d535b.
//
// Solidity: function create(bytes32 _salt, bytes32 _bytecodeHash, bytes _input) payable returns(address newAddress)
func (_IContractDeployer *IContractDeployerSession) Create(_salt [32]byte, _bytecodeHash [32]byte, _input []byte) (*types.Transaction, error) {
	return _IContractDeployer.Contract.Create(&_IContractDeployer.TransactOpts, _salt, _bytecodeHash, _input)
}

// Create is a paid mutator transaction binding the contract method 0x9c4d535b.
//
// Solidity: function create(bytes32 _salt, bytes32 _bytecodeHash, bytes _input) payable returns(address newAddress)
func (_IContractDeployer *IContractDeployerTransactorSession) Create(_salt [32]byte, _bytecodeHash [32]byte, _input []byte) (*types.Transaction, error) {
	return _IContractDeployer.Contract.Create(&_IContractDeployer.TransactOpts, _salt, _bytecodeHash, _input)
}

// Create2 is a paid mutator transaction binding the contract method 0x3cda3351.
//
// Solidity: function create2(bytes32 _salt, bytes32 _bytecodeHash, bytes _input) payable returns(address newAddress)
func (_IContractDeployer *IContractDeployerTransactor) Create2(opts *bind.TransactOpts, _salt [32]byte, _bytecodeHash [32]byte, _input []byte) (*types.Transaction, error) {
	return _IContractDeployer.contract.Transact(opts, "create2", _salt, _bytecodeHash, _input)
}

// Create2 is a paid mutator transaction binding the contract method 0x3cda3351.
//
// Solidity: function create2(bytes32 _salt, bytes32 _bytecodeHash, bytes _input) payable returns(address newAddress)
func (_IContractDeployer *IContractDeployerSession) Create2(_salt [32]byte, _bytecodeHash [32]byte, _input []byte) (*types.Transaction, error) {
	return _IContractDeployer.Contract.Create2(&_IContractDeployer.TransactOpts, _salt, _bytecodeHash, _input)
}

// Create2 is a paid mutator transaction binding the contract method 0x3cda3351.
//
// Solidity: function create2(bytes32 _salt, bytes32 _bytecodeHash, bytes _input) payable returns(address newAddress)
func (_IContractDeployer *IContractDeployerTransactorSession) Create2(_salt [32]byte, _bytecodeHash [32]byte, _input []byte) (*types.Transaction, error) {
	return _IContractDeployer.Contract.Create2(&_IContractDeployer.TransactOpts, _salt, _bytecodeHash, _input)
}

// Create2Account is a paid mutator transaction binding the contract method 0x5d382700.
//
// Solidity: function create2Account(bytes32 _salt, bytes32 _bytecodeHash, bytes _input, uint8 _aaVersion) payable returns(address newAddress)
func (_IContractDeployer *IContractDeployerTransactor) Create2Account(opts *bind.TransactOpts, _salt [32]byte, _bytecodeHash [32]byte, _input []byte, _aaVersion uint8) (*types.Transaction, error) {
	return _IContractDeployer.contract.Transact(opts, "create2Account", _salt, _bytecodeHash, _input, _aaVersion)
}

// Create2Account is a paid mutator transaction binding the contract method 0x5d382700.
//
// Solidity: function create2Account(bytes32 _salt, bytes32 _bytecodeHash, bytes _input, uint8 _aaVersion) payable returns(address newAddress)
func (_IContractDeployer *IContractDeployerSession) Create2Account(_salt [32]byte, _bytecodeHash [32]byte, _input []byte, _aaVersion uint8) (*types.Transaction, error) {
	return _IContractDeployer.Contract.Create2Account(&_IContractDeployer.TransactOpts, _salt, _bytecodeHash, _input, _aaVersion)
}

// Create2Account is a paid mutator transaction binding the contract method 0x5d382700.
//
// Solidity: function create2Account(bytes32 _salt, bytes32 _bytecodeHash, bytes _input, uint8 _aaVersion) payable returns(address newAddress)
func (_IContractDeployer *IContractDeployerTransactorSession) Create2Account(_salt [32]byte, _bytecodeHash [32]byte, _input []byte, _aaVersion uint8) (*types.Transaction, error) {
	return _IContractDeployer.Contract.Create2Account(&_IContractDeployer.TransactOpts, _salt, _bytecodeHash, _input, _aaVersion)
}

// CreateAccount is a paid mutator transaction binding the contract method 0xecf95b8a.
//
// Solidity: function createAccount(bytes32 _salt, bytes32 _bytecodeHash, bytes _input, uint8 _aaVersion) payable returns(address newAddress)
func (_IContractDeployer *IContractDeployerTransactor) CreateAccount(opts *bind.TransactOpts, _salt [32]byte, _bytecodeHash [32]byte, _input []byte, _aaVersion uint8) (*types.Transaction, error) {
	return _IContractDeployer.contract.Transact(opts, "createAccount", _salt, _bytecodeHash, _input, _aaVersion)
}

// CreateAccount is a paid mutator transaction binding the contract method 0xecf95b8a.
//
// Solidity: function createAccount(bytes32 _salt, bytes32 _bytecodeHash, bytes _input, uint8 _aaVersion) payable returns(address newAddress)
func (_IContractDeployer *IContractDeployerSession) CreateAccount(_salt [32]byte, _bytecodeHash [32]byte, _input []byte, _aaVersion uint8) (*types.Transaction, error) {
	return _IContractDeployer.Contract.CreateAccount(&_IContractDeployer.TransactOpts, _salt, _bytecodeHash, _input, _aaVersion)
}

// CreateAccount is a paid mutator transaction binding the contract method 0xecf95b8a.
//
// Solidity: function createAccount(bytes32 _salt, bytes32 _bytecodeHash, bytes _input, uint8 _aaVersion) payable returns(address newAddress)
func (_IContractDeployer *IContractDeployerTransactorSession) CreateAccount(_salt [32]byte, _bytecodeHash [32]byte, _input []byte, _aaVersion uint8) (*types.Transaction, error) {
	return _IContractDeployer.Contract.CreateAccount(&_IContractDeployer.TransactOpts, _salt, _bytecodeHash, _input, _aaVersion)
}

// UpdateAccountVersion is a paid mutator transaction binding the contract method 0x57180981.
//
// Solidity: function updateAccountVersion(uint8 _version) returns()
func (_IContractDeployer *IContractDeployerTransactor) UpdateAccountVersion(opts *bind.TransactOpts, _version uint8) (*types.Transaction, error) {
	return _IContractDeployer.contract.Transact(opts, "updateAccountVersion", _version)
}

// UpdateAccountVersion is a paid mutator transaction binding the contract method 0x57180981.
//
// Solidity: function updateAccountVersion(uint8 _version) returns()
func (_IContractDeployer *IContractDeployerSession) UpdateAccountVersion(_version uint8) (*types.Transaction, error) {
	return _IContractDeployer.Contract.UpdateAccountVersion(&_IContractDeployer.TransactOpts, _version)
}

// UpdateAccountVersion is a paid mutator transaction binding the contract method 0x57180981.
//
// Solidity: function updateAccountVersion(uint8 _version) returns()
func (_IContractDeployer *IContractDeployerTransactorSession) UpdateAccountVersion(_version uint8) (*types.Transaction, error) {
	return _IContractDeployer.Contract.UpdateAccountVersion(&_IContractDeployer.TransactOpts, _version)
}

// UpdateNonceOrdering is a paid mutator transaction binding the contract method 0xec8067c7.
//
// Solidity: function updateNonceOrdering(uint8 _nonceOrdering) returns()
func (_IContractDeployer *IContractDeployerTransactor) UpdateNonceOrdering(opts *bind.TransactOpts, _nonceOrdering uint8) (*types.Transaction, error) {
	return _IContractDeployer.contract.Transact(opts, "updateNonceOrdering", _nonceOrdering)
}

// UpdateNonceOrdering is a paid mutator transaction binding the contract method 0xec8067c7.
//
// Solidity: function updateNonceOrdering(uint8 _nonceOrdering) returns()
func (_IContractDeployer *IContractDeployerSession) UpdateNonceOrdering(_nonceOrdering uint8) (*types.Transaction, error) {
	return _IContractDeployer.Contract.UpdateNonceOrdering(&_IContractDeployer.TransactOpts, _nonceOrdering)
}

// UpdateNonceOrdering is a paid mutator transaction binding the contract method 0xec8067c7.
//
// Solidity: function updateNonceOrdering(uint8 _nonceOrdering) returns()
func (_IContractDeployer *IContractDeployerTransactorSession) UpdateNonceOrdering(_nonceOrdering uint8) (*types.Transaction, error) {
	return _IContractDeployer.Contract.UpdateNonceOrdering(&_IContractDeployer.TransactOpts, _nonceOrdering)
}

// IContractDeployerAccountNonceOrderingUpdatedIterator is returned from FilterAccountNonceOrderingUpdated and is used to iterate over the raw logs and unpacked data for AccountNonceOrderingUpdated events raised by the IContractDeployer contract.
type IContractDeployerAccountNonceOrderingUpdatedIterator struct {
	Event *IContractDeployerAccountNonceOrderingUpdated // Event containing the contract specifics and raw log

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
func (it *IContractDeployerAccountNonceOrderingUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IContractDeployerAccountNonceOrderingUpdated)
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
		it.Event = new(IContractDeployerAccountNonceOrderingUpdated)
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
func (it *IContractDeployerAccountNonceOrderingUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IContractDeployerAccountNonceOrderingUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IContractDeployerAccountNonceOrderingUpdated represents a AccountNonceOrderingUpdated event raised by the IContractDeployer contract.
type IContractDeployerAccountNonceOrderingUpdated struct {
	AccountAddress common.Address
	NonceOrdering  uint8
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterAccountNonceOrderingUpdated is a free log retrieval operation binding the contract event 0xc7544194dab38b1652f35439b9b4806d8b71e113f2cf5c1351cb2ecf7c83959a.
//
// Solidity: event AccountNonceOrderingUpdated(address indexed accountAddress, uint8 nonceOrdering)
func (_IContractDeployer *IContractDeployerFilterer) FilterAccountNonceOrderingUpdated(opts *bind.FilterOpts, accountAddress []common.Address) (*IContractDeployerAccountNonceOrderingUpdatedIterator, error) {

	var accountAddressRule []interface{}
	for _, accountAddressItem := range accountAddress {
		accountAddressRule = append(accountAddressRule, accountAddressItem)
	}

	logs, sub, err := _IContractDeployer.contract.FilterLogs(opts, "AccountNonceOrderingUpdated", accountAddressRule)
	if err != nil {
		return nil, err
	}
	return &IContractDeployerAccountNonceOrderingUpdatedIterator{contract: _IContractDeployer.contract, event: "AccountNonceOrderingUpdated", logs: logs, sub: sub}, nil
}

// WatchAccountNonceOrderingUpdated is a free log subscription operation binding the contract event 0xc7544194dab38b1652f35439b9b4806d8b71e113f2cf5c1351cb2ecf7c83959a.
//
// Solidity: event AccountNonceOrderingUpdated(address indexed accountAddress, uint8 nonceOrdering)
func (_IContractDeployer *IContractDeployerFilterer) WatchAccountNonceOrderingUpdated(opts *bind.WatchOpts, sink chan<- *IContractDeployerAccountNonceOrderingUpdated, accountAddress []common.Address) (event.Subscription, error) {

	var accountAddressRule []interface{}
	for _, accountAddressItem := range accountAddress {
		accountAddressRule = append(accountAddressRule, accountAddressItem)
	}

	logs, sub, err := _IContractDeployer.contract.WatchLogs(opts, "AccountNonceOrderingUpdated", accountAddressRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IContractDeployerAccountNonceOrderingUpdated)
				if err := _IContractDeployer.contract.UnpackLog(event, "AccountNonceOrderingUpdated", log); err != nil {
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

// ParseAccountNonceOrderingUpdated is a log parse operation binding the contract event 0xc7544194dab38b1652f35439b9b4806d8b71e113f2cf5c1351cb2ecf7c83959a.
//
// Solidity: event AccountNonceOrderingUpdated(address indexed accountAddress, uint8 nonceOrdering)
func (_IContractDeployer *IContractDeployerFilterer) ParseAccountNonceOrderingUpdated(log types.Log) (*IContractDeployerAccountNonceOrderingUpdated, error) {
	event := new(IContractDeployerAccountNonceOrderingUpdated)
	if err := _IContractDeployer.contract.UnpackLog(event, "AccountNonceOrderingUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// IContractDeployerAccountVersionUpdatedIterator is returned from FilterAccountVersionUpdated and is used to iterate over the raw logs and unpacked data for AccountVersionUpdated events raised by the IContractDeployer contract.
type IContractDeployerAccountVersionUpdatedIterator struct {
	Event *IContractDeployerAccountVersionUpdated // Event containing the contract specifics and raw log

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
func (it *IContractDeployerAccountVersionUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IContractDeployerAccountVersionUpdated)
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
		it.Event = new(IContractDeployerAccountVersionUpdated)
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
func (it *IContractDeployerAccountVersionUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IContractDeployerAccountVersionUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IContractDeployerAccountVersionUpdated represents a AccountVersionUpdated event raised by the IContractDeployer contract.
type IContractDeployerAccountVersionUpdated struct {
	AccountAddress common.Address
	AaVersion      uint8
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterAccountVersionUpdated is a free log retrieval operation binding the contract event 0x3fb6f4f15ddd4a75588ca934894ad2cdcab25a5012e2515e1783433d0128611a.
//
// Solidity: event AccountVersionUpdated(address indexed accountAddress, uint8 aaVersion)
func (_IContractDeployer *IContractDeployerFilterer) FilterAccountVersionUpdated(opts *bind.FilterOpts, accountAddress []common.Address) (*IContractDeployerAccountVersionUpdatedIterator, error) {

	var accountAddressRule []interface{}
	for _, accountAddressItem := range accountAddress {
		accountAddressRule = append(accountAddressRule, accountAddressItem)
	}

	logs, sub, err := _IContractDeployer.contract.FilterLogs(opts, "AccountVersionUpdated", accountAddressRule)
	if err != nil {
		return nil, err
	}
	return &IContractDeployerAccountVersionUpdatedIterator{contract: _IContractDeployer.contract, event: "AccountVersionUpdated", logs: logs, sub: sub}, nil
}

// WatchAccountVersionUpdated is a free log subscription operation binding the contract event 0x3fb6f4f15ddd4a75588ca934894ad2cdcab25a5012e2515e1783433d0128611a.
//
// Solidity: event AccountVersionUpdated(address indexed accountAddress, uint8 aaVersion)
func (_IContractDeployer *IContractDeployerFilterer) WatchAccountVersionUpdated(opts *bind.WatchOpts, sink chan<- *IContractDeployerAccountVersionUpdated, accountAddress []common.Address) (event.Subscription, error) {

	var accountAddressRule []interface{}
	for _, accountAddressItem := range accountAddress {
		accountAddressRule = append(accountAddressRule, accountAddressItem)
	}

	logs, sub, err := _IContractDeployer.contract.WatchLogs(opts, "AccountVersionUpdated", accountAddressRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IContractDeployerAccountVersionUpdated)
				if err := _IContractDeployer.contract.UnpackLog(event, "AccountVersionUpdated", log); err != nil {
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

// ParseAccountVersionUpdated is a log parse operation binding the contract event 0x3fb6f4f15ddd4a75588ca934894ad2cdcab25a5012e2515e1783433d0128611a.
//
// Solidity: event AccountVersionUpdated(address indexed accountAddress, uint8 aaVersion)
func (_IContractDeployer *IContractDeployerFilterer) ParseAccountVersionUpdated(log types.Log) (*IContractDeployerAccountVersionUpdated, error) {
	event := new(IContractDeployerAccountVersionUpdated)
	if err := _IContractDeployer.contract.UnpackLog(event, "AccountVersionUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// IContractDeployerContractDeployedIterator is returned from FilterContractDeployed and is used to iterate over the raw logs and unpacked data for ContractDeployed events raised by the IContractDeployer contract.
type IContractDeployerContractDeployedIterator struct {
	Event *IContractDeployerContractDeployed // Event containing the contract specifics and raw log

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
func (it *IContractDeployerContractDeployedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IContractDeployerContractDeployed)
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
		it.Event = new(IContractDeployerContractDeployed)
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
func (it *IContractDeployerContractDeployedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IContractDeployerContractDeployedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IContractDeployerContractDeployed represents a ContractDeployed event raised by the IContractDeployer contract.
type IContractDeployerContractDeployed struct {
	DeployerAddress common.Address
	BytecodeHash    [32]byte
	ContractAddress common.Address
	Raw             types.Log // Blockchain specific contextual infos
}

// FilterContractDeployed is a free log retrieval operation binding the contract event 0x290afdae231a3fc0bbae8b1af63698b0a1d79b21ad17df0342dfb952fe74f8e5.
//
// Solidity: event ContractDeployed(address indexed deployerAddress, bytes32 indexed bytecodeHash, address indexed contractAddress)
func (_IContractDeployer *IContractDeployerFilterer) FilterContractDeployed(opts *bind.FilterOpts, deployerAddress []common.Address, bytecodeHash [][32]byte, contractAddress []common.Address) (*IContractDeployerContractDeployedIterator, error) {

	var deployerAddressRule []interface{}
	for _, deployerAddressItem := range deployerAddress {
		deployerAddressRule = append(deployerAddressRule, deployerAddressItem)
	}
	var bytecodeHashRule []interface{}
	for _, bytecodeHashItem := range bytecodeHash {
		bytecodeHashRule = append(bytecodeHashRule, bytecodeHashItem)
	}
	var contractAddressRule []interface{}
	for _, contractAddressItem := range contractAddress {
		contractAddressRule = append(contractAddressRule, contractAddressItem)
	}

	logs, sub, err := _IContractDeployer.contract.FilterLogs(opts, "ContractDeployed", deployerAddressRule, bytecodeHashRule, contractAddressRule)
	if err != nil {
		return nil, err
	}
	return &IContractDeployerContractDeployedIterator{contract: _IContractDeployer.contract, event: "ContractDeployed", logs: logs, sub: sub}, nil
}

// WatchContractDeployed is a free log subscription operation binding the contract event 0x290afdae231a3fc0bbae8b1af63698b0a1d79b21ad17df0342dfb952fe74f8e5.
//
// Solidity: event ContractDeployed(address indexed deployerAddress, bytes32 indexed bytecodeHash, address indexed contractAddress)
func (_IContractDeployer *IContractDeployerFilterer) WatchContractDeployed(opts *bind.WatchOpts, sink chan<- *IContractDeployerContractDeployed, deployerAddress []common.Address, bytecodeHash [][32]byte, contractAddress []common.Address) (event.Subscription, error) {

	var deployerAddressRule []interface{}
	for _, deployerAddressItem := range deployerAddress {
		deployerAddressRule = append(deployerAddressRule, deployerAddressItem)
	}
	var bytecodeHashRule []interface{}
	for _, bytecodeHashItem := range bytecodeHash {
		bytecodeHashRule = append(bytecodeHashRule, bytecodeHashItem)
	}
	var contractAddressRule []interface{}
	for _, contractAddressItem := range contractAddress {
		contractAddressRule = append(contractAddressRule, contractAddressItem)
	}

	logs, sub, err := _IContractDeployer.contract.WatchLogs(opts, "ContractDeployed", deployerAddressRule, bytecodeHashRule, contractAddressRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IContractDeployerContractDeployed)
				if err := _IContractDeployer.contract.UnpackLog(event, "ContractDeployed", log); err != nil {
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

// ParseContractDeployed is a log parse operation binding the contract event 0x290afdae231a3fc0bbae8b1af63698b0a1d79b21ad17df0342dfb952fe74f8e5.
//
// Solidity: event ContractDeployed(address indexed deployerAddress, bytes32 indexed bytecodeHash, address indexed contractAddress)
func (_IContractDeployer *IContractDeployerFilterer) ParseContractDeployed(log types.Log) (*IContractDeployerContractDeployed, error) {
	event := new(IContractDeployerContractDeployed)
	if err := _IContractDeployer.contract.UnpackLog(event, "ContractDeployed", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
