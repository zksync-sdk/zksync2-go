// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package ContractDeployer

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

// ContractDeployerMetaData contains all meta data concerning the ContractDeployer contract.
var ContractDeployerMetaData = &bind.MetaData{
	ABI: "[{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"deployerAddress\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"bytecodeHash\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"contractAddress\",\"type\":\"address\"}],\"name\":\"ContractDeployed\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_salt\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"_bytecodeHash\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"_value\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"_input\",\"type\":\"bytes\"}],\"name\":\"create\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_salt\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"_bytecodeHash\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"_value\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"_input\",\"type\":\"bytes\"}],\"name\":\"create2\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_salt\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"_bytecodeHash\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"_value\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"_input\",\"type\":\"bytes\"}],\"name\":\"create2AA\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_salt\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"_bytecodeHash\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"_value\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"_input\",\"type\":\"bytes\"}],\"name\":\"createAA\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
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
	parsed, err := abi.JSON(strings.NewReader(ContractDeployerABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
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

// Create is a paid mutator transaction binding the contract method 0xe2e9718a.
//
// Solidity: function create(bytes32 _salt, bytes32 _bytecodeHash, uint256 _value, bytes _input) returns(address)
func (_ContractDeployer *ContractDeployerTransactor) Create(opts *bind.TransactOpts, _salt [32]byte, _bytecodeHash [32]byte, _value *big.Int, _input []byte) (*types.Transaction, error) {
	return _ContractDeployer.contract.Transact(opts, "create", _salt, _bytecodeHash, _value, _input)
}

// Create is a paid mutator transaction binding the contract method 0xe2e9718a.
//
// Solidity: function create(bytes32 _salt, bytes32 _bytecodeHash, uint256 _value, bytes _input) returns(address)
func (_ContractDeployer *ContractDeployerSession) Create(_salt [32]byte, _bytecodeHash [32]byte, _value *big.Int, _input []byte) (*types.Transaction, error) {
	return _ContractDeployer.Contract.Create(&_ContractDeployer.TransactOpts, _salt, _bytecodeHash, _value, _input)
}

// Create is a paid mutator transaction binding the contract method 0xe2e9718a.
//
// Solidity: function create(bytes32 _salt, bytes32 _bytecodeHash, uint256 _value, bytes _input) returns(address)
func (_ContractDeployer *ContractDeployerTransactorSession) Create(_salt [32]byte, _bytecodeHash [32]byte, _value *big.Int, _input []byte) (*types.Transaction, error) {
	return _ContractDeployer.Contract.Create(&_ContractDeployer.TransactOpts, _salt, _bytecodeHash, _value, _input)
}

// Create2 is a paid mutator transaction binding the contract method 0x1415dae2.
//
// Solidity: function create2(bytes32 _salt, bytes32 _bytecodeHash, uint256 _value, bytes _input) returns(address)
func (_ContractDeployer *ContractDeployerTransactor) Create2(opts *bind.TransactOpts, _salt [32]byte, _bytecodeHash [32]byte, _value *big.Int, _input []byte) (*types.Transaction, error) {
	return _ContractDeployer.contract.Transact(opts, "create2", _salt, _bytecodeHash, _value, _input)
}

// Create2 is a paid mutator transaction binding the contract method 0x1415dae2.
//
// Solidity: function create2(bytes32 _salt, bytes32 _bytecodeHash, uint256 _value, bytes _input) returns(address)
func (_ContractDeployer *ContractDeployerSession) Create2(_salt [32]byte, _bytecodeHash [32]byte, _value *big.Int, _input []byte) (*types.Transaction, error) {
	return _ContractDeployer.Contract.Create2(&_ContractDeployer.TransactOpts, _salt, _bytecodeHash, _value, _input)
}

// Create2 is a paid mutator transaction binding the contract method 0x1415dae2.
//
// Solidity: function create2(bytes32 _salt, bytes32 _bytecodeHash, uint256 _value, bytes _input) returns(address)
func (_ContractDeployer *ContractDeployerTransactorSession) Create2(_salt [32]byte, _bytecodeHash [32]byte, _value *big.Int, _input []byte) (*types.Transaction, error) {
	return _ContractDeployer.Contract.Create2(&_ContractDeployer.TransactOpts, _salt, _bytecodeHash, _value, _input)
}

// Create2AA is a paid mutator transaction binding the contract method 0x350adbbb.
//
// Solidity: function create2AA(bytes32 _salt, bytes32 _bytecodeHash, uint256 _value, bytes _input) returns(address)
func (_ContractDeployer *ContractDeployerTransactor) Create2AA(opts *bind.TransactOpts, _salt [32]byte, _bytecodeHash [32]byte, _value *big.Int, _input []byte) (*types.Transaction, error) {
	return _ContractDeployer.contract.Transact(opts, "create2AA", _salt, _bytecodeHash, _value, _input)
}

// Create2AA is a paid mutator transaction binding the contract method 0x350adbbb.
//
// Solidity: function create2AA(bytes32 _salt, bytes32 _bytecodeHash, uint256 _value, bytes _input) returns(address)
func (_ContractDeployer *ContractDeployerSession) Create2AA(_salt [32]byte, _bytecodeHash [32]byte, _value *big.Int, _input []byte) (*types.Transaction, error) {
	return _ContractDeployer.Contract.Create2AA(&_ContractDeployer.TransactOpts, _salt, _bytecodeHash, _value, _input)
}

// Create2AA is a paid mutator transaction binding the contract method 0x350adbbb.
//
// Solidity: function create2AA(bytes32 _salt, bytes32 _bytecodeHash, uint256 _value, bytes _input) returns(address)
func (_ContractDeployer *ContractDeployerTransactorSession) Create2AA(_salt [32]byte, _bytecodeHash [32]byte, _value *big.Int, _input []byte) (*types.Transaction, error) {
	return _ContractDeployer.Contract.Create2AA(&_ContractDeployer.TransactOpts, _salt, _bytecodeHash, _value, _input)
}

// CreateAA is a paid mutator transaction binding the contract method 0x26bf08ce.
//
// Solidity: function createAA(bytes32 _salt, bytes32 _bytecodeHash, uint256 _value, bytes _input) returns(address)
func (_ContractDeployer *ContractDeployerTransactor) CreateAA(opts *bind.TransactOpts, _salt [32]byte, _bytecodeHash [32]byte, _value *big.Int, _input []byte) (*types.Transaction, error) {
	return _ContractDeployer.contract.Transact(opts, "createAA", _salt, _bytecodeHash, _value, _input)
}

// CreateAA is a paid mutator transaction binding the contract method 0x26bf08ce.
//
// Solidity: function createAA(bytes32 _salt, bytes32 _bytecodeHash, uint256 _value, bytes _input) returns(address)
func (_ContractDeployer *ContractDeployerSession) CreateAA(_salt [32]byte, _bytecodeHash [32]byte, _value *big.Int, _input []byte) (*types.Transaction, error) {
	return _ContractDeployer.Contract.CreateAA(&_ContractDeployer.TransactOpts, _salt, _bytecodeHash, _value, _input)
}

// CreateAA is a paid mutator transaction binding the contract method 0x26bf08ce.
//
// Solidity: function createAA(bytes32 _salt, bytes32 _bytecodeHash, uint256 _value, bytes _input) returns(address)
func (_ContractDeployer *ContractDeployerTransactorSession) CreateAA(_salt [32]byte, _bytecodeHash [32]byte, _value *big.Int, _input []byte) (*types.Transaction, error) {
	return _ContractDeployer.Contract.CreateAA(&_ContractDeployer.TransactOpts, _salt, _bytecodeHash, _value, _input)
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
