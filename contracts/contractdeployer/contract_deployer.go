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

// ContractDeployerForceDeployment is an auto generated low-level Go binding around an user-defined struct.
type ContractDeployerForceDeployment struct {
	BytecodeHash    [32]byte
	NewAddress      common.Address
	CallConstructor bool
	Value           *big.Int
	Input           []byte
}

// IContractDeployerAccountInfo is an auto generated low-level Go binding around an user-defined struct.
type IContractDeployerAccountInfo struct {
	SupportedAAVersion uint8
	NonceOrdering      uint8
}

// ContractDeployerMetaData contains all meta data concerning the ContractDeployer contract.
var ContractDeployerMetaData = &bind.MetaData{
	ABI: "[{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"accountAddress\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"enumIContractDeployer.AccountNonceOrdering\",\"name\":\"nonceOrdering\",\"type\":\"uint8\"}],\"name\":\"AccountNonceOrderingUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"accountAddress\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"enumIContractDeployer.AccountAbstractionVersion\",\"name\":\"aaVersion\",\"type\":\"uint8\"}],\"name\":\"AccountVersionUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"deployerAddress\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"bytecodeHash\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"contractAddress\",\"type\":\"address\"}],\"name\":\"ContractDeployed\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_salt\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"_bytecodeHash\",\"type\":\"bytes32\"},{\"internalType\":\"bytes\",\"name\":\"_input\",\"type\":\"bytes\"}],\"name\":\"create\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_salt\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"_bytecodeHash\",\"type\":\"bytes32\"},{\"internalType\":\"bytes\",\"name\":\"_input\",\"type\":\"bytes\"}],\"name\":\"create2\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_salt\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"_bytecodeHash\",\"type\":\"bytes32\"},{\"internalType\":\"bytes\",\"name\":\"_input\",\"type\":\"bytes\"},{\"internalType\":\"enumIContractDeployer.AccountAbstractionVersion\",\"name\":\"_aaVersion\",\"type\":\"uint8\"}],\"name\":\"create2Account\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"_bytecodeHash\",\"type\":\"bytes32\"},{\"internalType\":\"bytes\",\"name\":\"_input\",\"type\":\"bytes\"},{\"internalType\":\"enumIContractDeployer.AccountAbstractionVersion\",\"name\":\"_aaVersion\",\"type\":\"uint8\"}],\"name\":\"createAccount\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_address\",\"type\":\"address\"}],\"name\":\"extendedAccountVersion\",\"outputs\":[{\"internalType\":\"enumIContractDeployer.AccountAbstractionVersion\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"bytecodeHash\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"newAddress\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"callConstructor\",\"type\":\"bool\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"input\",\"type\":\"bytes\"}],\"internalType\":\"structContractDeployer.ForceDeployment\",\"name\":\"_deployment\",\"type\":\"tuple\"},{\"internalType\":\"address\",\"name\":\"_sender\",\"type\":\"address\"}],\"name\":\"forceDeployOnAddress\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"bytecodeHash\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"newAddress\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"callConstructor\",\"type\":\"bool\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"input\",\"type\":\"bytes\"}],\"internalType\":\"structContractDeployer.ForceDeployment[]\",\"name\":\"_deployments\",\"type\":\"tuple[]\"}],\"name\":\"forceDeployOnAddresses\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_address\",\"type\":\"address\"}],\"name\":\"getAccountInfo\",\"outputs\":[{\"components\":[{\"internalType\":\"enumIContractDeployer.AccountAbstractionVersion\",\"name\":\"supportedAAVersion\",\"type\":\"uint8\"},{\"internalType\":\"enumIContractDeployer.AccountNonceOrdering\",\"name\":\"nonceOrdering\",\"type\":\"uint8\"}],\"internalType\":\"structIContractDeployer.AccountInfo\",\"name\":\"info\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_sender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_senderNonce\",\"type\":\"uint256\"}],\"name\":\"getNewAddressCreate\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"newAddress\",\"type\":\"address\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_sender\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"_bytecodeHash\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"_salt\",\"type\":\"bytes32\"},{\"internalType\":\"bytes\",\"name\":\"_input\",\"type\":\"bytes\"}],\"name\":\"getNewAddressCreate2\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"newAddress\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"enumIContractDeployer.AccountAbstractionVersion\",\"name\":\"_version\",\"type\":\"uint8\"}],\"name\":\"updateAccountVersion\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"enumIContractDeployer.AccountNonceOrdering\",\"name\":\"_nonceOrdering\",\"type\":\"uint8\"}],\"name\":\"updateNonceOrdering\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
}

// ContractDeployerABI is the input ABI used to generate the binding from.
// Deprecated: Use ContractDeployerMetaData.ABI instead.
var ContractDeployerABI = ContractDeployerMetaData.ABI

// ContractDeployer is an auto generated Go binding around an Ethereum contract.
type ContractDeployer struct {
	ContractDeployerCaller     // Read-only binding to the contract
	ContractDeployerTransactor // Write-only binding to the contract
	ContractDeployerFilterer   // Log filterer for contract events
}

// ContractDeployerCaller is an auto generated read-only Go binding around an Ethereum contract.
type ContractDeployerCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ContractDeployerTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ContractDeployerTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ContractDeployerFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ContractDeployerFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ContractDeployerSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ContractDeployerSession struct {
	Contract     *ContractDeployer // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ContractDeployerCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ContractDeployerCallerSession struct {
	Contract *ContractDeployerCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts           // Call options to use throughout this session
}

// ContractDeployerTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ContractDeployerTransactorSession struct {
	Contract     *ContractDeployerTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts           // Transaction auth options to use throughout this session
}

// ContractDeployerRaw is an auto generated low-level Go binding around an Ethereum contract.
type ContractDeployerRaw struct {
	Contract *ContractDeployer // Generic contract binding to access the raw methods on
}

// ContractDeployerCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ContractDeployerCallerRaw struct {
	Contract *ContractDeployerCaller // Generic read-only contract binding to access the raw methods on
}

// ContractDeployerTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ContractDeployerTransactorRaw struct {
	Contract *ContractDeployerTransactor // Generic write-only contract binding to access the raw methods on
}

// NewContractDeployer creates a new instance of ContractDeployer, bound to a specific deployed contract.
func NewContractDeployer(address common.Address, backend bind.ContractBackend) (*ContractDeployer, error) {
	contract, err := bindContractDeployer(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &ContractDeployer{ContractDeployerCaller: ContractDeployerCaller{contract: contract}, ContractDeployerTransactor: ContractDeployerTransactor{contract: contract}, ContractDeployerFilterer: ContractDeployerFilterer{contract: contract}}, nil
}

// NewContractDeployerCaller creates a new read-only instance of ContractDeployer, bound to a specific deployed contract.
func NewContractDeployerCaller(address common.Address, caller bind.ContractCaller) (*ContractDeployerCaller, error) {
	contract, err := bindContractDeployer(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ContractDeployerCaller{contract: contract}, nil
}

// NewContractDeployerTransactor creates a new write-only instance of ContractDeployer, bound to a specific deployed contract.
func NewContractDeployerTransactor(address common.Address, transactor bind.ContractTransactor) (*ContractDeployerTransactor, error) {
	contract, err := bindContractDeployer(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ContractDeployerTransactor{contract: contract}, nil
}

// NewContractDeployerFilterer creates a new log filterer instance of ContractDeployer, bound to a specific deployed contract.
func NewContractDeployerFilterer(address common.Address, filterer bind.ContractFilterer) (*ContractDeployerFilterer, error) {
	contract, err := bindContractDeployer(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ContractDeployerFilterer{contract: contract}, nil
}

// bindContractDeployer binds a generic wrapper to an already deployed contract.
func bindContractDeployer(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := ContractDeployerMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ContractDeployer *ContractDeployerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ContractDeployer.Contract.ContractDeployerCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ContractDeployer *ContractDeployerRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ContractDeployer.Contract.ContractDeployerTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ContractDeployer *ContractDeployerRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ContractDeployer.Contract.ContractDeployerTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ContractDeployer *ContractDeployerCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ContractDeployer.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ContractDeployer *ContractDeployerTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ContractDeployer.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ContractDeployer *ContractDeployerTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ContractDeployer.Contract.contract.Transact(opts, method, params...)
}

// ExtendedAccountVersion is a free data retrieval call binding the contract method 0xbb0fd610.
//
// Solidity: function extendedAccountVersion(address _address) view returns(uint8)
func (_ContractDeployer *ContractDeployerCaller) ExtendedAccountVersion(opts *bind.CallOpts, _address common.Address) (uint8, error) {
	var out []interface{}
	err := _ContractDeployer.contract.Call(opts, &out, "extendedAccountVersion", _address)

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// ExtendedAccountVersion is a free data retrieval call binding the contract method 0xbb0fd610.
//
// Solidity: function extendedAccountVersion(address _address) view returns(uint8)
func (_ContractDeployer *ContractDeployerSession) ExtendedAccountVersion(_address common.Address) (uint8, error) {
	return _ContractDeployer.Contract.ExtendedAccountVersion(&_ContractDeployer.CallOpts, _address)
}

// ExtendedAccountVersion is a free data retrieval call binding the contract method 0xbb0fd610.
//
// Solidity: function extendedAccountVersion(address _address) view returns(uint8)
func (_ContractDeployer *ContractDeployerCallerSession) ExtendedAccountVersion(_address common.Address) (uint8, error) {
	return _ContractDeployer.Contract.ExtendedAccountVersion(&_ContractDeployer.CallOpts, _address)
}

// GetAccountInfo is a free data retrieval call binding the contract method 0x7b510fe8.
//
// Solidity: function getAccountInfo(address _address) view returns((uint8,uint8) info)
func (_ContractDeployer *ContractDeployerCaller) GetAccountInfo(opts *bind.CallOpts, _address common.Address) (IContractDeployerAccountInfo, error) {
	var out []interface{}
	err := _ContractDeployer.contract.Call(opts, &out, "getAccountInfo", _address)

	if err != nil {
		return *new(IContractDeployerAccountInfo), err
	}

	out0 := *abi.ConvertType(out[0], new(IContractDeployerAccountInfo)).(*IContractDeployerAccountInfo)

	return out0, err

}

// GetAccountInfo is a free data retrieval call binding the contract method 0x7b510fe8.
//
// Solidity: function getAccountInfo(address _address) view returns((uint8,uint8) info)
func (_ContractDeployer *ContractDeployerSession) GetAccountInfo(_address common.Address) (IContractDeployerAccountInfo, error) {
	return _ContractDeployer.Contract.GetAccountInfo(&_ContractDeployer.CallOpts, _address)
}

// GetAccountInfo is a free data retrieval call binding the contract method 0x7b510fe8.
//
// Solidity: function getAccountInfo(address _address) view returns((uint8,uint8) info)
func (_ContractDeployer *ContractDeployerCallerSession) GetAccountInfo(_address common.Address) (IContractDeployerAccountInfo, error) {
	return _ContractDeployer.Contract.GetAccountInfo(&_ContractDeployer.CallOpts, _address)
}

// GetNewAddressCreate is a free data retrieval call binding the contract method 0x187598a5.
//
// Solidity: function getNewAddressCreate(address _sender, uint256 _senderNonce) pure returns(address newAddress)
func (_ContractDeployer *ContractDeployerCaller) GetNewAddressCreate(opts *bind.CallOpts, _sender common.Address, _senderNonce *big.Int) (common.Address, error) {
	var out []interface{}
	err := _ContractDeployer.contract.Call(opts, &out, "getNewAddressCreate", _sender, _senderNonce)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetNewAddressCreate is a free data retrieval call binding the contract method 0x187598a5.
//
// Solidity: function getNewAddressCreate(address _sender, uint256 _senderNonce) pure returns(address newAddress)
func (_ContractDeployer *ContractDeployerSession) GetNewAddressCreate(_sender common.Address, _senderNonce *big.Int) (common.Address, error) {
	return _ContractDeployer.Contract.GetNewAddressCreate(&_ContractDeployer.CallOpts, _sender, _senderNonce)
}

// GetNewAddressCreate is a free data retrieval call binding the contract method 0x187598a5.
//
// Solidity: function getNewAddressCreate(address _sender, uint256 _senderNonce) pure returns(address newAddress)
func (_ContractDeployer *ContractDeployerCallerSession) GetNewAddressCreate(_sender common.Address, _senderNonce *big.Int) (common.Address, error) {
	return _ContractDeployer.Contract.GetNewAddressCreate(&_ContractDeployer.CallOpts, _sender, _senderNonce)
}

// GetNewAddressCreate2 is a free data retrieval call binding the contract method 0x84da1fb4.
//
// Solidity: function getNewAddressCreate2(address _sender, bytes32 _bytecodeHash, bytes32 _salt, bytes _input) view returns(address newAddress)
func (_ContractDeployer *ContractDeployerCaller) GetNewAddressCreate2(opts *bind.CallOpts, _sender common.Address, _bytecodeHash [32]byte, _salt [32]byte, _input []byte) (common.Address, error) {
	var out []interface{}
	err := _ContractDeployer.contract.Call(opts, &out, "getNewAddressCreate2", _sender, _bytecodeHash, _salt, _input)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetNewAddressCreate2 is a free data retrieval call binding the contract method 0x84da1fb4.
//
// Solidity: function getNewAddressCreate2(address _sender, bytes32 _bytecodeHash, bytes32 _salt, bytes _input) view returns(address newAddress)
func (_ContractDeployer *ContractDeployerSession) GetNewAddressCreate2(_sender common.Address, _bytecodeHash [32]byte, _salt [32]byte, _input []byte) (common.Address, error) {
	return _ContractDeployer.Contract.GetNewAddressCreate2(&_ContractDeployer.CallOpts, _sender, _bytecodeHash, _salt, _input)
}

// GetNewAddressCreate2 is a free data retrieval call binding the contract method 0x84da1fb4.
//
// Solidity: function getNewAddressCreate2(address _sender, bytes32 _bytecodeHash, bytes32 _salt, bytes _input) view returns(address newAddress)
func (_ContractDeployer *ContractDeployerCallerSession) GetNewAddressCreate2(_sender common.Address, _bytecodeHash [32]byte, _salt [32]byte, _input []byte) (common.Address, error) {
	return _ContractDeployer.Contract.GetNewAddressCreate2(&_ContractDeployer.CallOpts, _sender, _bytecodeHash, _salt, _input)
}

// Create is a paid mutator transaction binding the contract method 0x9c4d535b.
//
// Solidity: function create(bytes32 _salt, bytes32 _bytecodeHash, bytes _input) payable returns(address)
func (_ContractDeployer *ContractDeployerTransactor) Create(opts *bind.TransactOpts, _salt [32]byte, _bytecodeHash [32]byte, _input []byte) (*types.Transaction, error) {
	return _ContractDeployer.contract.Transact(opts, "create", _salt, _bytecodeHash, _input)
}

// Create is a paid mutator transaction binding the contract method 0x9c4d535b.
//
// Solidity: function create(bytes32 _salt, bytes32 _bytecodeHash, bytes _input) payable returns(address)
func (_ContractDeployer *ContractDeployerSession) Create(_salt [32]byte, _bytecodeHash [32]byte, _input []byte) (*types.Transaction, error) {
	return _ContractDeployer.Contract.Create(&_ContractDeployer.TransactOpts, _salt, _bytecodeHash, _input)
}

// Create is a paid mutator transaction binding the contract method 0x9c4d535b.
//
// Solidity: function create(bytes32 _salt, bytes32 _bytecodeHash, bytes _input) payable returns(address)
func (_ContractDeployer *ContractDeployerTransactorSession) Create(_salt [32]byte, _bytecodeHash [32]byte, _input []byte) (*types.Transaction, error) {
	return _ContractDeployer.Contract.Create(&_ContractDeployer.TransactOpts, _salt, _bytecodeHash, _input)
}

// Create2 is a paid mutator transaction binding the contract method 0x3cda3351.
//
// Solidity: function create2(bytes32 _salt, bytes32 _bytecodeHash, bytes _input) payable returns(address)
func (_ContractDeployer *ContractDeployerTransactor) Create2(opts *bind.TransactOpts, _salt [32]byte, _bytecodeHash [32]byte, _input []byte) (*types.Transaction, error) {
	return _ContractDeployer.contract.Transact(opts, "create2", _salt, _bytecodeHash, _input)
}

// Create2 is a paid mutator transaction binding the contract method 0x3cda3351.
//
// Solidity: function create2(bytes32 _salt, bytes32 _bytecodeHash, bytes _input) payable returns(address)
func (_ContractDeployer *ContractDeployerSession) Create2(_salt [32]byte, _bytecodeHash [32]byte, _input []byte) (*types.Transaction, error) {
	return _ContractDeployer.Contract.Create2(&_ContractDeployer.TransactOpts, _salt, _bytecodeHash, _input)
}

// Create2 is a paid mutator transaction binding the contract method 0x3cda3351.
//
// Solidity: function create2(bytes32 _salt, bytes32 _bytecodeHash, bytes _input) payable returns(address)
func (_ContractDeployer *ContractDeployerTransactorSession) Create2(_salt [32]byte, _bytecodeHash [32]byte, _input []byte) (*types.Transaction, error) {
	return _ContractDeployer.Contract.Create2(&_ContractDeployer.TransactOpts, _salt, _bytecodeHash, _input)
}

// Create2Account is a paid mutator transaction binding the contract method 0x5d382700.
//
// Solidity: function create2Account(bytes32 _salt, bytes32 _bytecodeHash, bytes _input, uint8 _aaVersion) payable returns(address)
func (_ContractDeployer *ContractDeployerTransactor) Create2Account(opts *bind.TransactOpts, _salt [32]byte, _bytecodeHash [32]byte, _input []byte, _aaVersion uint8) (*types.Transaction, error) {
	return _ContractDeployer.contract.Transact(opts, "create2Account", _salt, _bytecodeHash, _input, _aaVersion)
}

// Create2Account is a paid mutator transaction binding the contract method 0x5d382700.
//
// Solidity: function create2Account(bytes32 _salt, bytes32 _bytecodeHash, bytes _input, uint8 _aaVersion) payable returns(address)
func (_ContractDeployer *ContractDeployerSession) Create2Account(_salt [32]byte, _bytecodeHash [32]byte, _input []byte, _aaVersion uint8) (*types.Transaction, error) {
	return _ContractDeployer.Contract.Create2Account(&_ContractDeployer.TransactOpts, _salt, _bytecodeHash, _input, _aaVersion)
}

// Create2Account is a paid mutator transaction binding the contract method 0x5d382700.
//
// Solidity: function create2Account(bytes32 _salt, bytes32 _bytecodeHash, bytes _input, uint8 _aaVersion) payable returns(address)
func (_ContractDeployer *ContractDeployerTransactorSession) Create2Account(_salt [32]byte, _bytecodeHash [32]byte, _input []byte, _aaVersion uint8) (*types.Transaction, error) {
	return _ContractDeployer.Contract.Create2Account(&_ContractDeployer.TransactOpts, _salt, _bytecodeHash, _input, _aaVersion)
}

// CreateAccount is a paid mutator transaction binding the contract method 0xecf95b8a.
//
// Solidity: function createAccount(bytes32 , bytes32 _bytecodeHash, bytes _input, uint8 _aaVersion) payable returns(address)
func (_ContractDeployer *ContractDeployerTransactor) CreateAccount(opts *bind.TransactOpts, arg0 [32]byte, _bytecodeHash [32]byte, _input []byte, _aaVersion uint8) (*types.Transaction, error) {
	return _ContractDeployer.contract.Transact(opts, "createAccount", arg0, _bytecodeHash, _input, _aaVersion)
}

// CreateAccount is a paid mutator transaction binding the contract method 0xecf95b8a.
//
// Solidity: function createAccount(bytes32 , bytes32 _bytecodeHash, bytes _input, uint8 _aaVersion) payable returns(address)
func (_ContractDeployer *ContractDeployerSession) CreateAccount(arg0 [32]byte, _bytecodeHash [32]byte, _input []byte, _aaVersion uint8) (*types.Transaction, error) {
	return _ContractDeployer.Contract.CreateAccount(&_ContractDeployer.TransactOpts, arg0, _bytecodeHash, _input, _aaVersion)
}

// CreateAccount is a paid mutator transaction binding the contract method 0xecf95b8a.
//
// Solidity: function createAccount(bytes32 , bytes32 _bytecodeHash, bytes _input, uint8 _aaVersion) payable returns(address)
func (_ContractDeployer *ContractDeployerTransactorSession) CreateAccount(arg0 [32]byte, _bytecodeHash [32]byte, _input []byte, _aaVersion uint8) (*types.Transaction, error) {
	return _ContractDeployer.Contract.CreateAccount(&_ContractDeployer.TransactOpts, arg0, _bytecodeHash, _input, _aaVersion)
}

// ForceDeployOnAddress is a paid mutator transaction binding the contract method 0xf3385fb6.
//
// Solidity: function forceDeployOnAddress((bytes32,address,bool,uint256,bytes) _deployment, address _sender) payable returns()
func (_ContractDeployer *ContractDeployerTransactor) ForceDeployOnAddress(opts *bind.TransactOpts, _deployment ContractDeployerForceDeployment, _sender common.Address) (*types.Transaction, error) {
	return _ContractDeployer.contract.Transact(opts, "forceDeployOnAddress", _deployment, _sender)
}

// ForceDeployOnAddress is a paid mutator transaction binding the contract method 0xf3385fb6.
//
// Solidity: function forceDeployOnAddress((bytes32,address,bool,uint256,bytes) _deployment, address _sender) payable returns()
func (_ContractDeployer *ContractDeployerSession) ForceDeployOnAddress(_deployment ContractDeployerForceDeployment, _sender common.Address) (*types.Transaction, error) {
	return _ContractDeployer.Contract.ForceDeployOnAddress(&_ContractDeployer.TransactOpts, _deployment, _sender)
}

// ForceDeployOnAddress is a paid mutator transaction binding the contract method 0xf3385fb6.
//
// Solidity: function forceDeployOnAddress((bytes32,address,bool,uint256,bytes) _deployment, address _sender) payable returns()
func (_ContractDeployer *ContractDeployerTransactorSession) ForceDeployOnAddress(_deployment ContractDeployerForceDeployment, _sender common.Address) (*types.Transaction, error) {
	return _ContractDeployer.Contract.ForceDeployOnAddress(&_ContractDeployer.TransactOpts, _deployment, _sender)
}

// ForceDeployOnAddresses is a paid mutator transaction binding the contract method 0xe9f18c17.
//
// Solidity: function forceDeployOnAddresses((bytes32,address,bool,uint256,bytes)[] _deployments) payable returns()
func (_ContractDeployer *ContractDeployerTransactor) ForceDeployOnAddresses(opts *bind.TransactOpts, _deployments []ContractDeployerForceDeployment) (*types.Transaction, error) {
	return _ContractDeployer.contract.Transact(opts, "forceDeployOnAddresses", _deployments)
}

// ForceDeployOnAddresses is a paid mutator transaction binding the contract method 0xe9f18c17.
//
// Solidity: function forceDeployOnAddresses((bytes32,address,bool,uint256,bytes)[] _deployments) payable returns()
func (_ContractDeployer *ContractDeployerSession) ForceDeployOnAddresses(_deployments []ContractDeployerForceDeployment) (*types.Transaction, error) {
	return _ContractDeployer.Contract.ForceDeployOnAddresses(&_ContractDeployer.TransactOpts, _deployments)
}

// ForceDeployOnAddresses is a paid mutator transaction binding the contract method 0xe9f18c17.
//
// Solidity: function forceDeployOnAddresses((bytes32,address,bool,uint256,bytes)[] _deployments) payable returns()
func (_ContractDeployer *ContractDeployerTransactorSession) ForceDeployOnAddresses(_deployments []ContractDeployerForceDeployment) (*types.Transaction, error) {
	return _ContractDeployer.Contract.ForceDeployOnAddresses(&_ContractDeployer.TransactOpts, _deployments)
}

// UpdateAccountVersion is a paid mutator transaction binding the contract method 0x57180981.
//
// Solidity: function updateAccountVersion(uint8 _version) returns()
func (_ContractDeployer *ContractDeployerTransactor) UpdateAccountVersion(opts *bind.TransactOpts, _version uint8) (*types.Transaction, error) {
	return _ContractDeployer.contract.Transact(opts, "updateAccountVersion", _version)
}

// UpdateAccountVersion is a paid mutator transaction binding the contract method 0x57180981.
//
// Solidity: function updateAccountVersion(uint8 _version) returns()
func (_ContractDeployer *ContractDeployerSession) UpdateAccountVersion(_version uint8) (*types.Transaction, error) {
	return _ContractDeployer.Contract.UpdateAccountVersion(&_ContractDeployer.TransactOpts, _version)
}

// UpdateAccountVersion is a paid mutator transaction binding the contract method 0x57180981.
//
// Solidity: function updateAccountVersion(uint8 _version) returns()
func (_ContractDeployer *ContractDeployerTransactorSession) UpdateAccountVersion(_version uint8) (*types.Transaction, error) {
	return _ContractDeployer.Contract.UpdateAccountVersion(&_ContractDeployer.TransactOpts, _version)
}

// UpdateNonceOrdering is a paid mutator transaction binding the contract method 0xec8067c7.
//
// Solidity: function updateNonceOrdering(uint8 _nonceOrdering) returns()
func (_ContractDeployer *ContractDeployerTransactor) UpdateNonceOrdering(opts *bind.TransactOpts, _nonceOrdering uint8) (*types.Transaction, error) {
	return _ContractDeployer.contract.Transact(opts, "updateNonceOrdering", _nonceOrdering)
}

// UpdateNonceOrdering is a paid mutator transaction binding the contract method 0xec8067c7.
//
// Solidity: function updateNonceOrdering(uint8 _nonceOrdering) returns()
func (_ContractDeployer *ContractDeployerSession) UpdateNonceOrdering(_nonceOrdering uint8) (*types.Transaction, error) {
	return _ContractDeployer.Contract.UpdateNonceOrdering(&_ContractDeployer.TransactOpts, _nonceOrdering)
}

// UpdateNonceOrdering is a paid mutator transaction binding the contract method 0xec8067c7.
//
// Solidity: function updateNonceOrdering(uint8 _nonceOrdering) returns()
func (_ContractDeployer *ContractDeployerTransactorSession) UpdateNonceOrdering(_nonceOrdering uint8) (*types.Transaction, error) {
	return _ContractDeployer.Contract.UpdateNonceOrdering(&_ContractDeployer.TransactOpts, _nonceOrdering)
}

// ContractDeployerAccountNonceOrderingUpdatedIterator is returned from FilterAccountNonceOrderingUpdated and is used to iterate over the raw logs and unpacked data for AccountNonceOrderingUpdated events raised by the ContractDeployer contract.
type ContractDeployerAccountNonceOrderingUpdatedIterator struct {
	Event *ContractDeployerAccountNonceOrderingUpdated // Event containing the contract specifics and raw log

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
func (it *ContractDeployerAccountNonceOrderingUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractDeployerAccountNonceOrderingUpdated)
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
		it.Event = new(ContractDeployerAccountNonceOrderingUpdated)
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
func (it *ContractDeployerAccountNonceOrderingUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractDeployerAccountNonceOrderingUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractDeployerAccountNonceOrderingUpdated represents a AccountNonceOrderingUpdated event raised by the ContractDeployer contract.
type ContractDeployerAccountNonceOrderingUpdated struct {
	AccountAddress common.Address
	NonceOrdering  uint8
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterAccountNonceOrderingUpdated is a free log retrieval operation binding the contract event 0xc7544194dab38b1652f35439b9b4806d8b71e113f2cf5c1351cb2ecf7c83959a.
//
// Solidity: event AccountNonceOrderingUpdated(address indexed accountAddress, uint8 nonceOrdering)
func (_ContractDeployer *ContractDeployerFilterer) FilterAccountNonceOrderingUpdated(opts *bind.FilterOpts, accountAddress []common.Address) (*ContractDeployerAccountNonceOrderingUpdatedIterator, error) {

	var accountAddressRule []interface{}
	for _, accountAddressItem := range accountAddress {
		accountAddressRule = append(accountAddressRule, accountAddressItem)
	}

	logs, sub, err := _ContractDeployer.contract.FilterLogs(opts, "AccountNonceOrderingUpdated", accountAddressRule)
	if err != nil {
		return nil, err
	}
	return &ContractDeployerAccountNonceOrderingUpdatedIterator{contract: _ContractDeployer.contract, event: "AccountNonceOrderingUpdated", logs: logs, sub: sub}, nil
}

// WatchAccountNonceOrderingUpdated is a free log subscription operation binding the contract event 0xc7544194dab38b1652f35439b9b4806d8b71e113f2cf5c1351cb2ecf7c83959a.
//
// Solidity: event AccountNonceOrderingUpdated(address indexed accountAddress, uint8 nonceOrdering)
func (_ContractDeployer *ContractDeployerFilterer) WatchAccountNonceOrderingUpdated(opts *bind.WatchOpts, sink chan<- *ContractDeployerAccountNonceOrderingUpdated, accountAddress []common.Address) (event.Subscription, error) {

	var accountAddressRule []interface{}
	for _, accountAddressItem := range accountAddress {
		accountAddressRule = append(accountAddressRule, accountAddressItem)
	}

	logs, sub, err := _ContractDeployer.contract.WatchLogs(opts, "AccountNonceOrderingUpdated", accountAddressRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractDeployerAccountNonceOrderingUpdated)
				if err := _ContractDeployer.contract.UnpackLog(event, "AccountNonceOrderingUpdated", log); err != nil {
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
func (_ContractDeployer *ContractDeployerFilterer) ParseAccountNonceOrderingUpdated(log types.Log) (*ContractDeployerAccountNonceOrderingUpdated, error) {
	event := new(ContractDeployerAccountNonceOrderingUpdated)
	if err := _ContractDeployer.contract.UnpackLog(event, "AccountNonceOrderingUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ContractDeployerAccountVersionUpdatedIterator is returned from FilterAccountVersionUpdated and is used to iterate over the raw logs and unpacked data for AccountVersionUpdated events raised by the ContractDeployer contract.
type ContractDeployerAccountVersionUpdatedIterator struct {
	Event *ContractDeployerAccountVersionUpdated // Event containing the contract specifics and raw log

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
func (it *ContractDeployerAccountVersionUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractDeployerAccountVersionUpdated)
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
		it.Event = new(ContractDeployerAccountVersionUpdated)
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
func (it *ContractDeployerAccountVersionUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractDeployerAccountVersionUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractDeployerAccountVersionUpdated represents a AccountVersionUpdated event raised by the ContractDeployer contract.
type ContractDeployerAccountVersionUpdated struct {
	AccountAddress common.Address
	AaVersion      uint8
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterAccountVersionUpdated is a free log retrieval operation binding the contract event 0x3fb6f4f15ddd4a75588ca934894ad2cdcab25a5012e2515e1783433d0128611a.
//
// Solidity: event AccountVersionUpdated(address indexed accountAddress, uint8 aaVersion)
func (_ContractDeployer *ContractDeployerFilterer) FilterAccountVersionUpdated(opts *bind.FilterOpts, accountAddress []common.Address) (*ContractDeployerAccountVersionUpdatedIterator, error) {

	var accountAddressRule []interface{}
	for _, accountAddressItem := range accountAddress {
		accountAddressRule = append(accountAddressRule, accountAddressItem)
	}

	logs, sub, err := _ContractDeployer.contract.FilterLogs(opts, "AccountVersionUpdated", accountAddressRule)
	if err != nil {
		return nil, err
	}
	return &ContractDeployerAccountVersionUpdatedIterator{contract: _ContractDeployer.contract, event: "AccountVersionUpdated", logs: logs, sub: sub}, nil
}

// WatchAccountVersionUpdated is a free log subscription operation binding the contract event 0x3fb6f4f15ddd4a75588ca934894ad2cdcab25a5012e2515e1783433d0128611a.
//
// Solidity: event AccountVersionUpdated(address indexed accountAddress, uint8 aaVersion)
func (_ContractDeployer *ContractDeployerFilterer) WatchAccountVersionUpdated(opts *bind.WatchOpts, sink chan<- *ContractDeployerAccountVersionUpdated, accountAddress []common.Address) (event.Subscription, error) {

	var accountAddressRule []interface{}
	for _, accountAddressItem := range accountAddress {
		accountAddressRule = append(accountAddressRule, accountAddressItem)
	}

	logs, sub, err := _ContractDeployer.contract.WatchLogs(opts, "AccountVersionUpdated", accountAddressRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractDeployerAccountVersionUpdated)
				if err := _ContractDeployer.contract.UnpackLog(event, "AccountVersionUpdated", log); err != nil {
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
func (_ContractDeployer *ContractDeployerFilterer) ParseAccountVersionUpdated(log types.Log) (*ContractDeployerAccountVersionUpdated, error) {
	event := new(ContractDeployerAccountVersionUpdated)
	if err := _ContractDeployer.contract.UnpackLog(event, "AccountVersionUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ContractDeployerContractDeployedIterator is returned from FilterContractDeployed and is used to iterate over the raw logs and unpacked data for ContractDeployed events raised by the ContractDeployer contract.
type ContractDeployerContractDeployedIterator struct {
	Event *ContractDeployerContractDeployed // Event containing the contract specifics and raw log

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
func (it *ContractDeployerContractDeployedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractDeployerContractDeployed)
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
		it.Event = new(ContractDeployerContractDeployed)
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
func (it *ContractDeployerContractDeployedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractDeployerContractDeployedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractDeployerContractDeployed represents a ContractDeployed event raised by the ContractDeployer contract.
type ContractDeployerContractDeployed struct {
	DeployerAddress common.Address
	BytecodeHash    [32]byte
	ContractAddress common.Address
	Raw             types.Log // Blockchain specific contextual infos
}

// FilterContractDeployed is a free log retrieval operation binding the contract event 0x290afdae231a3fc0bbae8b1af63698b0a1d79b21ad17df0342dfb952fe74f8e5.
//
// Solidity: event ContractDeployed(address indexed deployerAddress, bytes32 indexed bytecodeHash, address indexed contractAddress)
func (_ContractDeployer *ContractDeployerFilterer) FilterContractDeployed(opts *bind.FilterOpts, deployerAddress []common.Address, bytecodeHash [][32]byte, contractAddress []common.Address) (*ContractDeployerContractDeployedIterator, error) {

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

	logs, sub, err := _ContractDeployer.contract.FilterLogs(opts, "ContractDeployed", deployerAddressRule, bytecodeHashRule, contractAddressRule)
	if err != nil {
		return nil, err
	}
	return &ContractDeployerContractDeployedIterator{contract: _ContractDeployer.contract, event: "ContractDeployed", logs: logs, sub: sub}, nil
}

// WatchContractDeployed is a free log subscription operation binding the contract event 0x290afdae231a3fc0bbae8b1af63698b0a1d79b21ad17df0342dfb952fe74f8e5.
//
// Solidity: event ContractDeployed(address indexed deployerAddress, bytes32 indexed bytecodeHash, address indexed contractAddress)
func (_ContractDeployer *ContractDeployerFilterer) WatchContractDeployed(opts *bind.WatchOpts, sink chan<- *ContractDeployerContractDeployed, deployerAddress []common.Address, bytecodeHash [][32]byte, contractAddress []common.Address) (event.Subscription, error) {

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

	logs, sub, err := _ContractDeployer.contract.WatchLogs(opts, "ContractDeployed", deployerAddressRule, bytecodeHashRule, contractAddressRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractDeployerContractDeployed)
				if err := _ContractDeployer.contract.UnpackLog(event, "ContractDeployed", log); err != nil {
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
func (_ContractDeployer *ContractDeployerFilterer) ParseContractDeployed(log types.Log) (*ContractDeployerContractDeployed, error) {
	event := new(ContractDeployerContractDeployed)
	if err := _ContractDeployer.contract.UnpackLog(event, "ContractDeployed", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
