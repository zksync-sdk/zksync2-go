// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package aafactory

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

// AAFactoryMetaData contains all meta data concerning the AAFactory contract.
var AAFactoryMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_beaconProxyBytecodeHash\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"_beacon\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_passKeyModule\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_sessionKeyModule\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"ACCOUNT_ALREADY_EXISTS\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"EMPTY_BEACON_ADDRESS\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"EMPTY_BEACON_BYTECODE_HASH\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"EMPTY_PASSKEY_ADDRESS\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"EMPTY_SESSIONKEY_ADDRESS\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"INVALID_ACCOUNT_KEYS\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"accountAddress\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"uniqueAccountId\",\"type\":\"bytes32\"}],\"name\":\"AccountCreated\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"accountId\",\"type\":\"bytes32\"}],\"name\":\"accountMappings\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"deployedAccount\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"beacon\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"beaconProxyBytecodeHash\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"uniqueId\",\"type\":\"bytes32\"},{\"internalType\":\"bytes\",\"name\":\"passKey\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"sessionKey\",\"type\":\"bytes\"},{\"internalType\":\"address[]\",\"name\":\"ownerKeys\",\"type\":\"address[]\"}],\"name\":\"deployModularAccount\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"accountAddress\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"uniqueId\",\"type\":\"bytes32\"},{\"internalType\":\"bytes[]\",\"name\":\"initialValidators\",\"type\":\"bytes[]\"},{\"internalType\":\"address[]\",\"name\":\"initialK1Owners\",\"type\":\"address[]\"}],\"name\":\"deployProxySsoAccount\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"accountAddress\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getEncodedBeacon\",\"outputs\":[{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"passKeyModule\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"sessionKeyModule\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
}

// AAFactoryABI is the input ABI used to generate the binding from.
// Deprecated: Use AAFactoryMetaData.ABI instead.
var AAFactoryABI = AAFactoryMetaData.ABI

// AAFactory is an auto generated Go binding around an Ethereum contract.
type AAFactory struct {
	AAFactoryCaller     // Read-only binding to the contract
	AAFactoryTransactor // Write-only binding to the contract
	AAFactoryFilterer   // Log filterer for contract events
}

// AAFactoryCaller is an auto generated read-only Go binding around an Ethereum contract.
type AAFactoryCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AAFactoryTransactor is an auto generated write-only Go binding around an Ethereum contract.
type AAFactoryTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AAFactoryFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type AAFactoryFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AAFactorySession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type AAFactorySession struct {
	Contract     *AAFactory        // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// AAFactoryCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type AAFactoryCallerSession struct {
	Contract *AAFactoryCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts    // Call options to use throughout this session
}

// AAFactoryTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type AAFactoryTransactorSession struct {
	Contract     *AAFactoryTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts    // Transaction auth options to use throughout this session
}

// AAFactoryRaw is an auto generated low-level Go binding around an Ethereum contract.
type AAFactoryRaw struct {
	Contract *AAFactory // Generic contract binding to access the raw methods on
}

// AAFactoryCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type AAFactoryCallerRaw struct {
	Contract *AAFactoryCaller // Generic read-only contract binding to access the raw methods on
}

// AAFactoryTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type AAFactoryTransactorRaw struct {
	Contract *AAFactoryTransactor // Generic write-only contract binding to access the raw methods on
}

// NewAAFactory creates a new instance of AAFactory, bound to a specific deployed contract.
func NewAAFactory(address common.Address, backend bind.ContractBackend) (*AAFactory, error) {
	contract, err := bindAAFactory(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &AAFactory{AAFactoryCaller: AAFactoryCaller{contract: contract}, AAFactoryTransactor: AAFactoryTransactor{contract: contract}, AAFactoryFilterer: AAFactoryFilterer{contract: contract}}, nil
}

// NewAAFactoryCaller creates a new read-only instance of AAFactory, bound to a specific deployed contract.
func NewAAFactoryCaller(address common.Address, caller bind.ContractCaller) (*AAFactoryCaller, error) {
	contract, err := bindAAFactory(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &AAFactoryCaller{contract: contract}, nil
}

// NewAAFactoryTransactor creates a new write-only instance of AAFactory, bound to a specific deployed contract.
func NewAAFactoryTransactor(address common.Address, transactor bind.ContractTransactor) (*AAFactoryTransactor, error) {
	contract, err := bindAAFactory(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &AAFactoryTransactor{contract: contract}, nil
}

// NewAAFactoryFilterer creates a new log filterer instance of AAFactory, bound to a specific deployed contract.
func NewAAFactoryFilterer(address common.Address, filterer bind.ContractFilterer) (*AAFactoryFilterer, error) {
	contract, err := bindAAFactory(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &AAFactoryFilterer{contract: contract}, nil
}

// bindAAFactory binds a generic wrapper to an already deployed contract.
func bindAAFactory(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := AAFactoryMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_AAFactory *AAFactoryRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _AAFactory.Contract.AAFactoryCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_AAFactory *AAFactoryRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _AAFactory.Contract.AAFactoryTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_AAFactory *AAFactoryRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _AAFactory.Contract.AAFactoryTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_AAFactory *AAFactoryCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _AAFactory.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_AAFactory *AAFactoryTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _AAFactory.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_AAFactory *AAFactoryTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _AAFactory.Contract.contract.Transact(opts, method, params...)
}

// AccountMappings is a free data retrieval call binding the contract method 0x08614c35.
//
// Solidity: function accountMappings(bytes32 accountId) view returns(address deployedAccount)
func (_AAFactory *AAFactoryCaller) AccountMappings(opts *bind.CallOpts, accountId [32]byte) (common.Address, error) {
	var out []interface{}
	err := _AAFactory.contract.Call(opts, &out, "accountMappings", accountId)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// AccountMappings is a free data retrieval call binding the contract method 0x08614c35.
//
// Solidity: function accountMappings(bytes32 accountId) view returns(address deployedAccount)
func (_AAFactory *AAFactorySession) AccountMappings(accountId [32]byte) (common.Address, error) {
	return _AAFactory.Contract.AccountMappings(&_AAFactory.CallOpts, accountId)
}

// AccountMappings is a free data retrieval call binding the contract method 0x08614c35.
//
// Solidity: function accountMappings(bytes32 accountId) view returns(address deployedAccount)
func (_AAFactory *AAFactoryCallerSession) AccountMappings(accountId [32]byte) (common.Address, error) {
	return _AAFactory.Contract.AccountMappings(&_AAFactory.CallOpts, accountId)
}

// Beacon is a free data retrieval call binding the contract method 0x59659e90.
//
// Solidity: function beacon() view returns(address)
func (_AAFactory *AAFactoryCaller) Beacon(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _AAFactory.contract.Call(opts, &out, "beacon")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Beacon is a free data retrieval call binding the contract method 0x59659e90.
//
// Solidity: function beacon() view returns(address)
func (_AAFactory *AAFactorySession) Beacon() (common.Address, error) {
	return _AAFactory.Contract.Beacon(&_AAFactory.CallOpts)
}

// Beacon is a free data retrieval call binding the contract method 0x59659e90.
//
// Solidity: function beacon() view returns(address)
func (_AAFactory *AAFactoryCallerSession) Beacon() (common.Address, error) {
	return _AAFactory.Contract.Beacon(&_AAFactory.CallOpts)
}

// BeaconProxyBytecodeHash is a free data retrieval call binding the contract method 0xc23f8a8d.
//
// Solidity: function beaconProxyBytecodeHash() view returns(bytes32)
func (_AAFactory *AAFactoryCaller) BeaconProxyBytecodeHash(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _AAFactory.contract.Call(opts, &out, "beaconProxyBytecodeHash")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// BeaconProxyBytecodeHash is a free data retrieval call binding the contract method 0xc23f8a8d.
//
// Solidity: function beaconProxyBytecodeHash() view returns(bytes32)
func (_AAFactory *AAFactorySession) BeaconProxyBytecodeHash() ([32]byte, error) {
	return _AAFactory.Contract.BeaconProxyBytecodeHash(&_AAFactory.CallOpts)
}

// BeaconProxyBytecodeHash is a free data retrieval call binding the contract method 0xc23f8a8d.
//
// Solidity: function beaconProxyBytecodeHash() view returns(bytes32)
func (_AAFactory *AAFactoryCallerSession) BeaconProxyBytecodeHash() ([32]byte, error) {
	return _AAFactory.Contract.BeaconProxyBytecodeHash(&_AAFactory.CallOpts)
}

// GetEncodedBeacon is a free data retrieval call binding the contract method 0xf9958a10.
//
// Solidity: function getEncodedBeacon() view returns(bytes)
func (_AAFactory *AAFactoryCaller) GetEncodedBeacon(opts *bind.CallOpts) ([]byte, error) {
	var out []interface{}
	err := _AAFactory.contract.Call(opts, &out, "getEncodedBeacon")

	if err != nil {
		return *new([]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([]byte)).(*[]byte)

	return out0, err

}

// GetEncodedBeacon is a free data retrieval call binding the contract method 0xf9958a10.
//
// Solidity: function getEncodedBeacon() view returns(bytes)
func (_AAFactory *AAFactorySession) GetEncodedBeacon() ([]byte, error) {
	return _AAFactory.Contract.GetEncodedBeacon(&_AAFactory.CallOpts)
}

// GetEncodedBeacon is a free data retrieval call binding the contract method 0xf9958a10.
//
// Solidity: function getEncodedBeacon() view returns(bytes)
func (_AAFactory *AAFactoryCallerSession) GetEncodedBeacon() ([]byte, error) {
	return _AAFactory.Contract.GetEncodedBeacon(&_AAFactory.CallOpts)
}

// PassKeyModule is a free data retrieval call binding the contract method 0x8c9682d8.
//
// Solidity: function passKeyModule() view returns(address)
func (_AAFactory *AAFactoryCaller) PassKeyModule(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _AAFactory.contract.Call(opts, &out, "passKeyModule")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// PassKeyModule is a free data retrieval call binding the contract method 0x8c9682d8.
//
// Solidity: function passKeyModule() view returns(address)
func (_AAFactory *AAFactorySession) PassKeyModule() (common.Address, error) {
	return _AAFactory.Contract.PassKeyModule(&_AAFactory.CallOpts)
}

// PassKeyModule is a free data retrieval call binding the contract method 0x8c9682d8.
//
// Solidity: function passKeyModule() view returns(address)
func (_AAFactory *AAFactoryCallerSession) PassKeyModule() (common.Address, error) {
	return _AAFactory.Contract.PassKeyModule(&_AAFactory.CallOpts)
}

// SessionKeyModule is a free data retrieval call binding the contract method 0xf31b26fc.
//
// Solidity: function sessionKeyModule() view returns(address)
func (_AAFactory *AAFactoryCaller) SessionKeyModule(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _AAFactory.contract.Call(opts, &out, "sessionKeyModule")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// SessionKeyModule is a free data retrieval call binding the contract method 0xf31b26fc.
//
// Solidity: function sessionKeyModule() view returns(address)
func (_AAFactory *AAFactorySession) SessionKeyModule() (common.Address, error) {
	return _AAFactory.Contract.SessionKeyModule(&_AAFactory.CallOpts)
}

// SessionKeyModule is a free data retrieval call binding the contract method 0xf31b26fc.
//
// Solidity: function sessionKeyModule() view returns(address)
func (_AAFactory *AAFactoryCallerSession) SessionKeyModule() (common.Address, error) {
	return _AAFactory.Contract.SessionKeyModule(&_AAFactory.CallOpts)
}

// DeployModularAccount is a paid mutator transaction binding the contract method 0x54385730.
//
// Solidity: function deployModularAccount(bytes32 uniqueId, bytes passKey, bytes sessionKey, address[] ownerKeys) returns(address accountAddress)
func (_AAFactory *AAFactoryTransactor) DeployModularAccount(opts *bind.TransactOpts, uniqueId [32]byte, passKey []byte, sessionKey []byte, ownerKeys []common.Address) (*types.Transaction, error) {
	return _AAFactory.contract.Transact(opts, "deployModularAccount", uniqueId, passKey, sessionKey, ownerKeys)
}

// DeployModularAccount is a paid mutator transaction binding the contract method 0x54385730.
//
// Solidity: function deployModularAccount(bytes32 uniqueId, bytes passKey, bytes sessionKey, address[] ownerKeys) returns(address accountAddress)
func (_AAFactory *AAFactorySession) DeployModularAccount(uniqueId [32]byte, passKey []byte, sessionKey []byte, ownerKeys []common.Address) (*types.Transaction, error) {
	return _AAFactory.Contract.DeployModularAccount(&_AAFactory.TransactOpts, uniqueId, passKey, sessionKey, ownerKeys)
}

// DeployModularAccount is a paid mutator transaction binding the contract method 0x54385730.
//
// Solidity: function deployModularAccount(bytes32 uniqueId, bytes passKey, bytes sessionKey, address[] ownerKeys) returns(address accountAddress)
func (_AAFactory *AAFactoryTransactorSession) DeployModularAccount(uniqueId [32]byte, passKey []byte, sessionKey []byte, ownerKeys []common.Address) (*types.Transaction, error) {
	return _AAFactory.Contract.DeployModularAccount(&_AAFactory.TransactOpts, uniqueId, passKey, sessionKey, ownerKeys)
}

// DeployProxySsoAccount is a paid mutator transaction binding the contract method 0x0b8a4ed4.
//
// Solidity: function deployProxySsoAccount(bytes32 uniqueId, bytes[] initialValidators, address[] initialK1Owners) returns(address accountAddress)
func (_AAFactory *AAFactoryTransactor) DeployProxySsoAccount(opts *bind.TransactOpts, uniqueId [32]byte, initialValidators [][]byte, initialK1Owners []common.Address) (*types.Transaction, error) {
	return _AAFactory.contract.Transact(opts, "deployProxySsoAccount", uniqueId, initialValidators, initialK1Owners)
}

// DeployProxySsoAccount is a paid mutator transaction binding the contract method 0x0b8a4ed4.
//
// Solidity: function deployProxySsoAccount(bytes32 uniqueId, bytes[] initialValidators, address[] initialK1Owners) returns(address accountAddress)
func (_AAFactory *AAFactorySession) DeployProxySsoAccount(uniqueId [32]byte, initialValidators [][]byte, initialK1Owners []common.Address) (*types.Transaction, error) {
	return _AAFactory.Contract.DeployProxySsoAccount(&_AAFactory.TransactOpts, uniqueId, initialValidators, initialK1Owners)
}

// DeployProxySsoAccount is a paid mutator transaction binding the contract method 0x0b8a4ed4.
//
// Solidity: function deployProxySsoAccount(bytes32 uniqueId, bytes[] initialValidators, address[] initialK1Owners) returns(address accountAddress)
func (_AAFactory *AAFactoryTransactorSession) DeployProxySsoAccount(uniqueId [32]byte, initialValidators [][]byte, initialK1Owners []common.Address) (*types.Transaction, error) {
	return _AAFactory.Contract.DeployProxySsoAccount(&_AAFactory.TransactOpts, uniqueId, initialValidators, initialK1Owners)
}

// AAFactoryAccountCreatedIterator is returned from FilterAccountCreated and is used to iterate over the raw logs and unpacked data for AccountCreated events raised by the AAFactory contract.
type AAFactoryAccountCreatedIterator struct {
	Event *AAFactoryAccountCreated // Event containing the contract specifics and raw log

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
func (it *AAFactoryAccountCreatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AAFactoryAccountCreated)
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
		it.Event = new(AAFactoryAccountCreated)
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
func (it *AAFactoryAccountCreatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AAFactoryAccountCreatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AAFactoryAccountCreated represents a AccountCreated event raised by the AAFactory contract.
type AAFactoryAccountCreated struct {
	AccountAddress  common.Address
	UniqueAccountId [32]byte
	Raw             types.Log // Blockchain specific contextual infos
}

// FilterAccountCreated is a free log retrieval operation binding the contract event 0x8fe66a5d954d6d3e0306797e31e226812a9916895165c96c367ef52807631951.
//
// Solidity: event AccountCreated(address indexed accountAddress, bytes32 uniqueAccountId)
func (_AAFactory *AAFactoryFilterer) FilterAccountCreated(opts *bind.FilterOpts, accountAddress []common.Address) (*AAFactoryAccountCreatedIterator, error) {

	var accountAddressRule []interface{}
	for _, accountAddressItem := range accountAddress {
		accountAddressRule = append(accountAddressRule, accountAddressItem)
	}

	logs, sub, err := _AAFactory.contract.FilterLogs(opts, "AccountCreated", accountAddressRule)
	if err != nil {
		return nil, err
	}
	return &AAFactoryAccountCreatedIterator{contract: _AAFactory.contract, event: "AccountCreated", logs: logs, sub: sub}, nil
}

// WatchAccountCreated is a free log subscription operation binding the contract event 0x8fe66a5d954d6d3e0306797e31e226812a9916895165c96c367ef52807631951.
//
// Solidity: event AccountCreated(address indexed accountAddress, bytes32 uniqueAccountId)
func (_AAFactory *AAFactoryFilterer) WatchAccountCreated(opts *bind.WatchOpts, sink chan<- *AAFactoryAccountCreated, accountAddress []common.Address) (event.Subscription, error) {

	var accountAddressRule []interface{}
	for _, accountAddressItem := range accountAddress {
		accountAddressRule = append(accountAddressRule, accountAddressItem)
	}

	logs, sub, err := _AAFactory.contract.WatchLogs(opts, "AccountCreated", accountAddressRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AAFactoryAccountCreated)
				if err := _AAFactory.contract.UnpackLog(event, "AccountCreated", log); err != nil {
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

// ParseAccountCreated is a log parse operation binding the contract event 0x8fe66a5d954d6d3e0306797e31e226812a9916895165c96c367ef52807631951.
//
// Solidity: event AccountCreated(address indexed accountAddress, bytes32 uniqueAccountId)
func (_AAFactory *AAFactoryFilterer) ParseAccountCreated(log types.Log) (*AAFactoryAccountCreated, error) {
	event := new(AAFactoryAccountCreated)
	if err := _AAFactory.contract.UnpackLog(event, "AccountCreated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
