// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package assetrouterbase

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

// IAssetRouterBaseMetaData contains all meta data concerning the IAssetRouterBase contract.
var IAssetRouterBaseMetaData = &bind.MetaData{
	ABI: "[{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"assetId\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"_assetAddress\",\"type\":\"address\"}],\"name\":\"AssetHandlerRegistered\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"assetId\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"assetHandlerAddress\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"additionalData\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"assetDeploymentTracker\",\"type\":\"address\"}],\"name\":\"AssetHandlerRegisteredInitial\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"chainId\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"assetId\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"BridgehubDepositBaseTokenInitiated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"chainId\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"txDataHash\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"assetId\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"bridgeMintCalldata\",\"type\":\"bytes\"}],\"name\":\"BridgehubDepositInitiated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"chainId\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"assetId\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"assetDataHash\",\"type\":\"bytes32\"}],\"name\":\"BridgehubWithdrawalInitiated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"chainId\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"assetId\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"assetData\",\"type\":\"bytes\"}],\"name\":\"DepositFinalizedAssetRouter\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"BRIDGE_HUB\",\"outputs\":[{\"internalType\":\"contractIBridgehub\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_assetId\",\"type\":\"bytes32\"}],\"name\":\"assetHandlerAddress\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_chainId\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"_assetId\",\"type\":\"bytes32\"},{\"internalType\":\"bytes\",\"name\":\"_transferData\",\"type\":\"bytes\"}],\"name\":\"finalizeDeposit\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_assetRegistrationData\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"_assetHandlerAddress\",\"type\":\"address\"}],\"name\":\"setAssetHandlerAddressThisChain\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
}

// IAssetRouterBaseABI is the input ABI used to generate the binding from.
// Deprecated: Use IAssetRouterBaseMetaData.ABI instead.
var IAssetRouterBaseABI = IAssetRouterBaseMetaData.ABI

// IAssetRouterBase is an auto generated Go binding around an Ethereum contract.
type IAssetRouterBase struct {
	IAssetRouterBaseCaller     // Read-only binding to the contract
	IAssetRouterBaseTransactor // Write-only binding to the contract
	IAssetRouterBaseFilterer   // Log filterer for contract events
}

// IAssetRouterBaseCaller is an auto generated read-only Go binding around an Ethereum contract.
type IAssetRouterBaseCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IAssetRouterBaseTransactor is an auto generated write-only Go binding around an Ethereum contract.
type IAssetRouterBaseTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IAssetRouterBaseFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type IAssetRouterBaseFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IAssetRouterBaseSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type IAssetRouterBaseSession struct {
	Contract     *IAssetRouterBase // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// IAssetRouterBaseCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type IAssetRouterBaseCallerSession struct {
	Contract *IAssetRouterBaseCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts           // Call options to use throughout this session
}

// IAssetRouterBaseTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type IAssetRouterBaseTransactorSession struct {
	Contract     *IAssetRouterBaseTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts           // Transaction auth options to use throughout this session
}

// IAssetRouterBaseRaw is an auto generated low-level Go binding around an Ethereum contract.
type IAssetRouterBaseRaw struct {
	Contract *IAssetRouterBase // Generic contract binding to access the raw methods on
}

// IAssetRouterBaseCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type IAssetRouterBaseCallerRaw struct {
	Contract *IAssetRouterBaseCaller // Generic read-only contract binding to access the raw methods on
}

// IAssetRouterBaseTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type IAssetRouterBaseTransactorRaw struct {
	Contract *IAssetRouterBaseTransactor // Generic write-only contract binding to access the raw methods on
}

// NewIAssetRouterBase creates a new instance of IAssetRouterBase, bound to a specific deployed contract.
func NewIAssetRouterBase(address common.Address, backend bind.ContractBackend) (*IAssetRouterBase, error) {
	contract, err := bindIAssetRouterBase(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &IAssetRouterBase{IAssetRouterBaseCaller: IAssetRouterBaseCaller{contract: contract}, IAssetRouterBaseTransactor: IAssetRouterBaseTransactor{contract: contract}, IAssetRouterBaseFilterer: IAssetRouterBaseFilterer{contract: contract}}, nil
}

// NewIAssetRouterBaseCaller creates a new read-only instance of IAssetRouterBase, bound to a specific deployed contract.
func NewIAssetRouterBaseCaller(address common.Address, caller bind.ContractCaller) (*IAssetRouterBaseCaller, error) {
	contract, err := bindIAssetRouterBase(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &IAssetRouterBaseCaller{contract: contract}, nil
}

// NewIAssetRouterBaseTransactor creates a new write-only instance of IAssetRouterBase, bound to a specific deployed contract.
func NewIAssetRouterBaseTransactor(address common.Address, transactor bind.ContractTransactor) (*IAssetRouterBaseTransactor, error) {
	contract, err := bindIAssetRouterBase(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &IAssetRouterBaseTransactor{contract: contract}, nil
}

// NewIAssetRouterBaseFilterer creates a new log filterer instance of IAssetRouterBase, bound to a specific deployed contract.
func NewIAssetRouterBaseFilterer(address common.Address, filterer bind.ContractFilterer) (*IAssetRouterBaseFilterer, error) {
	contract, err := bindIAssetRouterBase(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &IAssetRouterBaseFilterer{contract: contract}, nil
}

// bindIAssetRouterBase binds a generic wrapper to an already deployed contract.
func bindIAssetRouterBase(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := IAssetRouterBaseMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IAssetRouterBase *IAssetRouterBaseRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IAssetRouterBase.Contract.IAssetRouterBaseCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IAssetRouterBase *IAssetRouterBaseRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IAssetRouterBase.Contract.IAssetRouterBaseTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IAssetRouterBase *IAssetRouterBaseRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IAssetRouterBase.Contract.IAssetRouterBaseTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IAssetRouterBase *IAssetRouterBaseCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IAssetRouterBase.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IAssetRouterBase *IAssetRouterBaseTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IAssetRouterBase.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IAssetRouterBase *IAssetRouterBaseTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IAssetRouterBase.Contract.contract.Transact(opts, method, params...)
}

// BRIDGEHUB is a free data retrieval call binding the contract method 0x5d4edca7.
//
// Solidity: function BRIDGE_HUB() view returns(address)
func (_IAssetRouterBase *IAssetRouterBaseCaller) BRIDGEHUB(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _IAssetRouterBase.contract.Call(opts, &out, "BRIDGE_HUB")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// BRIDGEHUB is a free data retrieval call binding the contract method 0x5d4edca7.
//
// Solidity: function BRIDGE_HUB() view returns(address)
func (_IAssetRouterBase *IAssetRouterBaseSession) BRIDGEHUB() (common.Address, error) {
	return _IAssetRouterBase.Contract.BRIDGEHUB(&_IAssetRouterBase.CallOpts)
}

// BRIDGEHUB is a free data retrieval call binding the contract method 0x5d4edca7.
//
// Solidity: function BRIDGE_HUB() view returns(address)
func (_IAssetRouterBase *IAssetRouterBaseCallerSession) BRIDGEHUB() (common.Address, error) {
	return _IAssetRouterBase.Contract.BRIDGEHUB(&_IAssetRouterBase.CallOpts)
}

// AssetHandlerAddress is a free data retrieval call binding the contract method 0x53b9e632.
//
// Solidity: function assetHandlerAddress(bytes32 _assetId) view returns(address)
func (_IAssetRouterBase *IAssetRouterBaseCaller) AssetHandlerAddress(opts *bind.CallOpts, _assetId [32]byte) (common.Address, error) {
	var out []interface{}
	err := _IAssetRouterBase.contract.Call(opts, &out, "assetHandlerAddress", _assetId)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// AssetHandlerAddress is a free data retrieval call binding the contract method 0x53b9e632.
//
// Solidity: function assetHandlerAddress(bytes32 _assetId) view returns(address)
func (_IAssetRouterBase *IAssetRouterBaseSession) AssetHandlerAddress(_assetId [32]byte) (common.Address, error) {
	return _IAssetRouterBase.Contract.AssetHandlerAddress(&_IAssetRouterBase.CallOpts, _assetId)
}

// AssetHandlerAddress is a free data retrieval call binding the contract method 0x53b9e632.
//
// Solidity: function assetHandlerAddress(bytes32 _assetId) view returns(address)
func (_IAssetRouterBase *IAssetRouterBaseCallerSession) AssetHandlerAddress(_assetId [32]byte) (common.Address, error) {
	return _IAssetRouterBase.Contract.AssetHandlerAddress(&_IAssetRouterBase.CallOpts, _assetId)
}

// FinalizeDeposit is a paid mutator transaction binding the contract method 0x9c884fd1.
//
// Solidity: function finalizeDeposit(uint256 _chainId, bytes32 _assetId, bytes _transferData) returns()
func (_IAssetRouterBase *IAssetRouterBaseTransactor) FinalizeDeposit(opts *bind.TransactOpts, _chainId *big.Int, _assetId [32]byte, _transferData []byte) (*types.Transaction, error) {
	return _IAssetRouterBase.contract.Transact(opts, "finalizeDeposit", _chainId, _assetId, _transferData)
}

// FinalizeDeposit is a paid mutator transaction binding the contract method 0x9c884fd1.
//
// Solidity: function finalizeDeposit(uint256 _chainId, bytes32 _assetId, bytes _transferData) returns()
func (_IAssetRouterBase *IAssetRouterBaseSession) FinalizeDeposit(_chainId *big.Int, _assetId [32]byte, _transferData []byte) (*types.Transaction, error) {
	return _IAssetRouterBase.Contract.FinalizeDeposit(&_IAssetRouterBase.TransactOpts, _chainId, _assetId, _transferData)
}

// FinalizeDeposit is a paid mutator transaction binding the contract method 0x9c884fd1.
//
// Solidity: function finalizeDeposit(uint256 _chainId, bytes32 _assetId, bytes _transferData) returns()
func (_IAssetRouterBase *IAssetRouterBaseTransactorSession) FinalizeDeposit(_chainId *big.Int, _assetId [32]byte, _transferData []byte) (*types.Transaction, error) {
	return _IAssetRouterBase.Contract.FinalizeDeposit(&_IAssetRouterBase.TransactOpts, _chainId, _assetId, _transferData)
}

// SetAssetHandlerAddressThisChain is a paid mutator transaction binding the contract method 0x548a5a33.
//
// Solidity: function setAssetHandlerAddressThisChain(bytes32 _assetRegistrationData, address _assetHandlerAddress) returns()
func (_IAssetRouterBase *IAssetRouterBaseTransactor) SetAssetHandlerAddressThisChain(opts *bind.TransactOpts, _assetRegistrationData [32]byte, _assetHandlerAddress common.Address) (*types.Transaction, error) {
	return _IAssetRouterBase.contract.Transact(opts, "setAssetHandlerAddressThisChain", _assetRegistrationData, _assetHandlerAddress)
}

// SetAssetHandlerAddressThisChain is a paid mutator transaction binding the contract method 0x548a5a33.
//
// Solidity: function setAssetHandlerAddressThisChain(bytes32 _assetRegistrationData, address _assetHandlerAddress) returns()
func (_IAssetRouterBase *IAssetRouterBaseSession) SetAssetHandlerAddressThisChain(_assetRegistrationData [32]byte, _assetHandlerAddress common.Address) (*types.Transaction, error) {
	return _IAssetRouterBase.Contract.SetAssetHandlerAddressThisChain(&_IAssetRouterBase.TransactOpts, _assetRegistrationData, _assetHandlerAddress)
}

// SetAssetHandlerAddressThisChain is a paid mutator transaction binding the contract method 0x548a5a33.
//
// Solidity: function setAssetHandlerAddressThisChain(bytes32 _assetRegistrationData, address _assetHandlerAddress) returns()
func (_IAssetRouterBase *IAssetRouterBaseTransactorSession) SetAssetHandlerAddressThisChain(_assetRegistrationData [32]byte, _assetHandlerAddress common.Address) (*types.Transaction, error) {
	return _IAssetRouterBase.Contract.SetAssetHandlerAddressThisChain(&_IAssetRouterBase.TransactOpts, _assetRegistrationData, _assetHandlerAddress)
}

// IAssetRouterBaseAssetHandlerRegisteredIterator is returned from FilterAssetHandlerRegistered and is used to iterate over the raw logs and unpacked data for AssetHandlerRegistered events raised by the IAssetRouterBase contract.
type IAssetRouterBaseAssetHandlerRegisteredIterator struct {
	Event *IAssetRouterBaseAssetHandlerRegistered // Event containing the contract specifics and raw log

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
func (it *IAssetRouterBaseAssetHandlerRegisteredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IAssetRouterBaseAssetHandlerRegistered)
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
		it.Event = new(IAssetRouterBaseAssetHandlerRegistered)
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
func (it *IAssetRouterBaseAssetHandlerRegisteredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IAssetRouterBaseAssetHandlerRegisteredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IAssetRouterBaseAssetHandlerRegistered represents a AssetHandlerRegistered event raised by the IAssetRouterBase contract.
type IAssetRouterBaseAssetHandlerRegistered struct {
	AssetId      [32]byte
	AssetAddress common.Address
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterAssetHandlerRegistered is a free log retrieval operation binding the contract event 0x2632cc0d58b0cb1017b99cc0b6cc66ad86440cc0dd923bfdaa294f95ba1b0201.
//
// Solidity: event AssetHandlerRegistered(bytes32 indexed assetId, address indexed _assetAddress)
func (_IAssetRouterBase *IAssetRouterBaseFilterer) FilterAssetHandlerRegistered(opts *bind.FilterOpts, assetId [][32]byte, _assetAddress []common.Address) (*IAssetRouterBaseAssetHandlerRegisteredIterator, error) {

	var assetIdRule []interface{}
	for _, assetIdItem := range assetId {
		assetIdRule = append(assetIdRule, assetIdItem)
	}
	var _assetAddressRule []interface{}
	for _, _assetAddressItem := range _assetAddress {
		_assetAddressRule = append(_assetAddressRule, _assetAddressItem)
	}

	logs, sub, err := _IAssetRouterBase.contract.FilterLogs(opts, "AssetHandlerRegistered", assetIdRule, _assetAddressRule)
	if err != nil {
		return nil, err
	}
	return &IAssetRouterBaseAssetHandlerRegisteredIterator{contract: _IAssetRouterBase.contract, event: "AssetHandlerRegistered", logs: logs, sub: sub}, nil
}

// WatchAssetHandlerRegistered is a free log subscription operation binding the contract event 0x2632cc0d58b0cb1017b99cc0b6cc66ad86440cc0dd923bfdaa294f95ba1b0201.
//
// Solidity: event AssetHandlerRegistered(bytes32 indexed assetId, address indexed _assetAddress)
func (_IAssetRouterBase *IAssetRouterBaseFilterer) WatchAssetHandlerRegistered(opts *bind.WatchOpts, sink chan<- *IAssetRouterBaseAssetHandlerRegistered, assetId [][32]byte, _assetAddress []common.Address) (event.Subscription, error) {

	var assetIdRule []interface{}
	for _, assetIdItem := range assetId {
		assetIdRule = append(assetIdRule, assetIdItem)
	}
	var _assetAddressRule []interface{}
	for _, _assetAddressItem := range _assetAddress {
		_assetAddressRule = append(_assetAddressRule, _assetAddressItem)
	}

	logs, sub, err := _IAssetRouterBase.contract.WatchLogs(opts, "AssetHandlerRegistered", assetIdRule, _assetAddressRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IAssetRouterBaseAssetHandlerRegistered)
				if err := _IAssetRouterBase.contract.UnpackLog(event, "AssetHandlerRegistered", log); err != nil {
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
func (_IAssetRouterBase *IAssetRouterBaseFilterer) ParseAssetHandlerRegistered(log types.Log) (*IAssetRouterBaseAssetHandlerRegistered, error) {
	event := new(IAssetRouterBaseAssetHandlerRegistered)
	if err := _IAssetRouterBase.contract.UnpackLog(event, "AssetHandlerRegistered", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// IAssetRouterBaseAssetHandlerRegisteredInitialIterator is returned from FilterAssetHandlerRegisteredInitial and is used to iterate over the raw logs and unpacked data for AssetHandlerRegisteredInitial events raised by the IAssetRouterBase contract.
type IAssetRouterBaseAssetHandlerRegisteredInitialIterator struct {
	Event *IAssetRouterBaseAssetHandlerRegisteredInitial // Event containing the contract specifics and raw log

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
func (it *IAssetRouterBaseAssetHandlerRegisteredInitialIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IAssetRouterBaseAssetHandlerRegisteredInitial)
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
		it.Event = new(IAssetRouterBaseAssetHandlerRegisteredInitial)
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
func (it *IAssetRouterBaseAssetHandlerRegisteredInitialIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IAssetRouterBaseAssetHandlerRegisteredInitialIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IAssetRouterBaseAssetHandlerRegisteredInitial represents a AssetHandlerRegisteredInitial event raised by the IAssetRouterBase contract.
type IAssetRouterBaseAssetHandlerRegisteredInitial struct {
	AssetId                [32]byte
	AssetHandlerAddress    common.Address
	AdditionalData         [32]byte
	AssetDeploymentTracker common.Address
	Raw                    types.Log // Blockchain specific contextual infos
}

// FilterAssetHandlerRegisteredInitial is a free log retrieval operation binding the contract event 0xb1e82bee3e85b2755fbceb4b7e051f5c66a7f35f0476657504e77e18ebd3a17d.
//
// Solidity: event AssetHandlerRegisteredInitial(bytes32 indexed assetId, address indexed assetHandlerAddress, bytes32 indexed additionalData, address assetDeploymentTracker)
func (_IAssetRouterBase *IAssetRouterBaseFilterer) FilterAssetHandlerRegisteredInitial(opts *bind.FilterOpts, assetId [][32]byte, assetHandlerAddress []common.Address, additionalData [][32]byte) (*IAssetRouterBaseAssetHandlerRegisteredInitialIterator, error) {

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

	logs, sub, err := _IAssetRouterBase.contract.FilterLogs(opts, "AssetHandlerRegisteredInitial", assetIdRule, assetHandlerAddressRule, additionalDataRule)
	if err != nil {
		return nil, err
	}
	return &IAssetRouterBaseAssetHandlerRegisteredInitialIterator{contract: _IAssetRouterBase.contract, event: "AssetHandlerRegisteredInitial", logs: logs, sub: sub}, nil
}

// WatchAssetHandlerRegisteredInitial is a free log subscription operation binding the contract event 0xb1e82bee3e85b2755fbceb4b7e051f5c66a7f35f0476657504e77e18ebd3a17d.
//
// Solidity: event AssetHandlerRegisteredInitial(bytes32 indexed assetId, address indexed assetHandlerAddress, bytes32 indexed additionalData, address assetDeploymentTracker)
func (_IAssetRouterBase *IAssetRouterBaseFilterer) WatchAssetHandlerRegisteredInitial(opts *bind.WatchOpts, sink chan<- *IAssetRouterBaseAssetHandlerRegisteredInitial, assetId [][32]byte, assetHandlerAddress []common.Address, additionalData [][32]byte) (event.Subscription, error) {

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

	logs, sub, err := _IAssetRouterBase.contract.WatchLogs(opts, "AssetHandlerRegisteredInitial", assetIdRule, assetHandlerAddressRule, additionalDataRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IAssetRouterBaseAssetHandlerRegisteredInitial)
				if err := _IAssetRouterBase.contract.UnpackLog(event, "AssetHandlerRegisteredInitial", log); err != nil {
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
func (_IAssetRouterBase *IAssetRouterBaseFilterer) ParseAssetHandlerRegisteredInitial(log types.Log) (*IAssetRouterBaseAssetHandlerRegisteredInitial, error) {
	event := new(IAssetRouterBaseAssetHandlerRegisteredInitial)
	if err := _IAssetRouterBase.contract.UnpackLog(event, "AssetHandlerRegisteredInitial", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// IAssetRouterBaseBridgehubDepositBaseTokenInitiatedIterator is returned from FilterBridgehubDepositBaseTokenInitiated and is used to iterate over the raw logs and unpacked data for BridgehubDepositBaseTokenInitiated events raised by the IAssetRouterBase contract.
type IAssetRouterBaseBridgehubDepositBaseTokenInitiatedIterator struct {
	Event *IAssetRouterBaseBridgehubDepositBaseTokenInitiated // Event containing the contract specifics and raw log

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
func (it *IAssetRouterBaseBridgehubDepositBaseTokenInitiatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IAssetRouterBaseBridgehubDepositBaseTokenInitiated)
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
		it.Event = new(IAssetRouterBaseBridgehubDepositBaseTokenInitiated)
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
func (it *IAssetRouterBaseBridgehubDepositBaseTokenInitiatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IAssetRouterBaseBridgehubDepositBaseTokenInitiatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IAssetRouterBaseBridgehubDepositBaseTokenInitiated represents a BridgehubDepositBaseTokenInitiated event raised by the IAssetRouterBase contract.
type IAssetRouterBaseBridgehubDepositBaseTokenInitiated struct {
	ChainId *big.Int
	From    common.Address
	AssetId [32]byte
	Amount  *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterBridgehubDepositBaseTokenInitiated is a free log retrieval operation binding the contract event 0x0f87e1ea5eb1f034a6071ef630c174063e3d48756f853efaaf4292b929298240.
//
// Solidity: event BridgehubDepositBaseTokenInitiated(uint256 indexed chainId, address indexed from, bytes32 assetId, uint256 amount)
func (_IAssetRouterBase *IAssetRouterBaseFilterer) FilterBridgehubDepositBaseTokenInitiated(opts *bind.FilterOpts, chainId []*big.Int, from []common.Address) (*IAssetRouterBaseBridgehubDepositBaseTokenInitiatedIterator, error) {

	var chainIdRule []interface{}
	for _, chainIdItem := range chainId {
		chainIdRule = append(chainIdRule, chainIdItem)
	}
	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}

	logs, sub, err := _IAssetRouterBase.contract.FilterLogs(opts, "BridgehubDepositBaseTokenInitiated", chainIdRule, fromRule)
	if err != nil {
		return nil, err
	}
	return &IAssetRouterBaseBridgehubDepositBaseTokenInitiatedIterator{contract: _IAssetRouterBase.contract, event: "BridgehubDepositBaseTokenInitiated", logs: logs, sub: sub}, nil
}

// WatchBridgehubDepositBaseTokenInitiated is a free log subscription operation binding the contract event 0x0f87e1ea5eb1f034a6071ef630c174063e3d48756f853efaaf4292b929298240.
//
// Solidity: event BridgehubDepositBaseTokenInitiated(uint256 indexed chainId, address indexed from, bytes32 assetId, uint256 amount)
func (_IAssetRouterBase *IAssetRouterBaseFilterer) WatchBridgehubDepositBaseTokenInitiated(opts *bind.WatchOpts, sink chan<- *IAssetRouterBaseBridgehubDepositBaseTokenInitiated, chainId []*big.Int, from []common.Address) (event.Subscription, error) {

	var chainIdRule []interface{}
	for _, chainIdItem := range chainId {
		chainIdRule = append(chainIdRule, chainIdItem)
	}
	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}

	logs, sub, err := _IAssetRouterBase.contract.WatchLogs(opts, "BridgehubDepositBaseTokenInitiated", chainIdRule, fromRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IAssetRouterBaseBridgehubDepositBaseTokenInitiated)
				if err := _IAssetRouterBase.contract.UnpackLog(event, "BridgehubDepositBaseTokenInitiated", log); err != nil {
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
func (_IAssetRouterBase *IAssetRouterBaseFilterer) ParseBridgehubDepositBaseTokenInitiated(log types.Log) (*IAssetRouterBaseBridgehubDepositBaseTokenInitiated, error) {
	event := new(IAssetRouterBaseBridgehubDepositBaseTokenInitiated)
	if err := _IAssetRouterBase.contract.UnpackLog(event, "BridgehubDepositBaseTokenInitiated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// IAssetRouterBaseBridgehubDepositInitiatedIterator is returned from FilterBridgehubDepositInitiated and is used to iterate over the raw logs and unpacked data for BridgehubDepositInitiated events raised by the IAssetRouterBase contract.
type IAssetRouterBaseBridgehubDepositInitiatedIterator struct {
	Event *IAssetRouterBaseBridgehubDepositInitiated // Event containing the contract specifics and raw log

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
func (it *IAssetRouterBaseBridgehubDepositInitiatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IAssetRouterBaseBridgehubDepositInitiated)
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
		it.Event = new(IAssetRouterBaseBridgehubDepositInitiated)
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
func (it *IAssetRouterBaseBridgehubDepositInitiatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IAssetRouterBaseBridgehubDepositInitiatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IAssetRouterBaseBridgehubDepositInitiated represents a BridgehubDepositInitiated event raised by the IAssetRouterBase contract.
type IAssetRouterBaseBridgehubDepositInitiated struct {
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
func (_IAssetRouterBase *IAssetRouterBaseFilterer) FilterBridgehubDepositInitiated(opts *bind.FilterOpts, chainId []*big.Int, txDataHash [][32]byte, from []common.Address) (*IAssetRouterBaseBridgehubDepositInitiatedIterator, error) {

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

	logs, sub, err := _IAssetRouterBase.contract.FilterLogs(opts, "BridgehubDepositInitiated", chainIdRule, txDataHashRule, fromRule)
	if err != nil {
		return nil, err
	}
	return &IAssetRouterBaseBridgehubDepositInitiatedIterator{contract: _IAssetRouterBase.contract, event: "BridgehubDepositInitiated", logs: logs, sub: sub}, nil
}

// WatchBridgehubDepositInitiated is a free log subscription operation binding the contract event 0xe21913bc89c1320d9709a5d236ffe06b54cf88aecfc9509ebd68f1adba45781e.
//
// Solidity: event BridgehubDepositInitiated(uint256 indexed chainId, bytes32 indexed txDataHash, address indexed from, bytes32 assetId, bytes bridgeMintCalldata)
func (_IAssetRouterBase *IAssetRouterBaseFilterer) WatchBridgehubDepositInitiated(opts *bind.WatchOpts, sink chan<- *IAssetRouterBaseBridgehubDepositInitiated, chainId []*big.Int, txDataHash [][32]byte, from []common.Address) (event.Subscription, error) {

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

	logs, sub, err := _IAssetRouterBase.contract.WatchLogs(opts, "BridgehubDepositInitiated", chainIdRule, txDataHashRule, fromRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IAssetRouterBaseBridgehubDepositInitiated)
				if err := _IAssetRouterBase.contract.UnpackLog(event, "BridgehubDepositInitiated", log); err != nil {
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
func (_IAssetRouterBase *IAssetRouterBaseFilterer) ParseBridgehubDepositInitiated(log types.Log) (*IAssetRouterBaseBridgehubDepositInitiated, error) {
	event := new(IAssetRouterBaseBridgehubDepositInitiated)
	if err := _IAssetRouterBase.contract.UnpackLog(event, "BridgehubDepositInitiated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// IAssetRouterBaseBridgehubWithdrawalInitiatedIterator is returned from FilterBridgehubWithdrawalInitiated and is used to iterate over the raw logs and unpacked data for BridgehubWithdrawalInitiated events raised by the IAssetRouterBase contract.
type IAssetRouterBaseBridgehubWithdrawalInitiatedIterator struct {
	Event *IAssetRouterBaseBridgehubWithdrawalInitiated // Event containing the contract specifics and raw log

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
func (it *IAssetRouterBaseBridgehubWithdrawalInitiatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IAssetRouterBaseBridgehubWithdrawalInitiated)
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
		it.Event = new(IAssetRouterBaseBridgehubWithdrawalInitiated)
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
func (it *IAssetRouterBaseBridgehubWithdrawalInitiatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IAssetRouterBaseBridgehubWithdrawalInitiatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IAssetRouterBaseBridgehubWithdrawalInitiated represents a BridgehubWithdrawalInitiated event raised by the IAssetRouterBase contract.
type IAssetRouterBaseBridgehubWithdrawalInitiated struct {
	ChainId       *big.Int
	Sender        common.Address
	AssetId       [32]byte
	AssetDataHash [32]byte
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterBridgehubWithdrawalInitiated is a free log retrieval operation binding the contract event 0x9a3d4025b7294a1754ea5b56309c1e72328d97b73718183db595c850d14a3ae0.
//
// Solidity: event BridgehubWithdrawalInitiated(uint256 chainId, address indexed sender, bytes32 indexed assetId, bytes32 assetDataHash)
func (_IAssetRouterBase *IAssetRouterBaseFilterer) FilterBridgehubWithdrawalInitiated(opts *bind.FilterOpts, sender []common.Address, assetId [][32]byte) (*IAssetRouterBaseBridgehubWithdrawalInitiatedIterator, error) {

	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}
	var assetIdRule []interface{}
	for _, assetIdItem := range assetId {
		assetIdRule = append(assetIdRule, assetIdItem)
	}

	logs, sub, err := _IAssetRouterBase.contract.FilterLogs(opts, "BridgehubWithdrawalInitiated", senderRule, assetIdRule)
	if err != nil {
		return nil, err
	}
	return &IAssetRouterBaseBridgehubWithdrawalInitiatedIterator{contract: _IAssetRouterBase.contract, event: "BridgehubWithdrawalInitiated", logs: logs, sub: sub}, nil
}

// WatchBridgehubWithdrawalInitiated is a free log subscription operation binding the contract event 0x9a3d4025b7294a1754ea5b56309c1e72328d97b73718183db595c850d14a3ae0.
//
// Solidity: event BridgehubWithdrawalInitiated(uint256 chainId, address indexed sender, bytes32 indexed assetId, bytes32 assetDataHash)
func (_IAssetRouterBase *IAssetRouterBaseFilterer) WatchBridgehubWithdrawalInitiated(opts *bind.WatchOpts, sink chan<- *IAssetRouterBaseBridgehubWithdrawalInitiated, sender []common.Address, assetId [][32]byte) (event.Subscription, error) {

	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}
	var assetIdRule []interface{}
	for _, assetIdItem := range assetId {
		assetIdRule = append(assetIdRule, assetIdItem)
	}

	logs, sub, err := _IAssetRouterBase.contract.WatchLogs(opts, "BridgehubWithdrawalInitiated", senderRule, assetIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IAssetRouterBaseBridgehubWithdrawalInitiated)
				if err := _IAssetRouterBase.contract.UnpackLog(event, "BridgehubWithdrawalInitiated", log); err != nil {
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
func (_IAssetRouterBase *IAssetRouterBaseFilterer) ParseBridgehubWithdrawalInitiated(log types.Log) (*IAssetRouterBaseBridgehubWithdrawalInitiated, error) {
	event := new(IAssetRouterBaseBridgehubWithdrawalInitiated)
	if err := _IAssetRouterBase.contract.UnpackLog(event, "BridgehubWithdrawalInitiated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// IAssetRouterBaseDepositFinalizedAssetRouterIterator is returned from FilterDepositFinalizedAssetRouter and is used to iterate over the raw logs and unpacked data for DepositFinalizedAssetRouter events raised by the IAssetRouterBase contract.
type IAssetRouterBaseDepositFinalizedAssetRouterIterator struct {
	Event *IAssetRouterBaseDepositFinalizedAssetRouter // Event containing the contract specifics and raw log

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
func (it *IAssetRouterBaseDepositFinalizedAssetRouterIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IAssetRouterBaseDepositFinalizedAssetRouter)
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
		it.Event = new(IAssetRouterBaseDepositFinalizedAssetRouter)
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
func (it *IAssetRouterBaseDepositFinalizedAssetRouterIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IAssetRouterBaseDepositFinalizedAssetRouterIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IAssetRouterBaseDepositFinalizedAssetRouter represents a DepositFinalizedAssetRouter event raised by the IAssetRouterBase contract.
type IAssetRouterBaseDepositFinalizedAssetRouter struct {
	ChainId   *big.Int
	AssetId   [32]byte
	AssetData []byte
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterDepositFinalizedAssetRouter is a free log retrieval operation binding the contract event 0x44eb9a840094a49b3cd0a5205042598a1c08c4e87bafb5760bc2d8efa170c541.
//
// Solidity: event DepositFinalizedAssetRouter(uint256 indexed chainId, bytes32 indexed assetId, bytes assetData)
func (_IAssetRouterBase *IAssetRouterBaseFilterer) FilterDepositFinalizedAssetRouter(opts *bind.FilterOpts, chainId []*big.Int, assetId [][32]byte) (*IAssetRouterBaseDepositFinalizedAssetRouterIterator, error) {

	var chainIdRule []interface{}
	for _, chainIdItem := range chainId {
		chainIdRule = append(chainIdRule, chainIdItem)
	}
	var assetIdRule []interface{}
	for _, assetIdItem := range assetId {
		assetIdRule = append(assetIdRule, assetIdItem)
	}

	logs, sub, err := _IAssetRouterBase.contract.FilterLogs(opts, "DepositFinalizedAssetRouter", chainIdRule, assetIdRule)
	if err != nil {
		return nil, err
	}
	return &IAssetRouterBaseDepositFinalizedAssetRouterIterator{contract: _IAssetRouterBase.contract, event: "DepositFinalizedAssetRouter", logs: logs, sub: sub}, nil
}

// WatchDepositFinalizedAssetRouter is a free log subscription operation binding the contract event 0x44eb9a840094a49b3cd0a5205042598a1c08c4e87bafb5760bc2d8efa170c541.
//
// Solidity: event DepositFinalizedAssetRouter(uint256 indexed chainId, bytes32 indexed assetId, bytes assetData)
func (_IAssetRouterBase *IAssetRouterBaseFilterer) WatchDepositFinalizedAssetRouter(opts *bind.WatchOpts, sink chan<- *IAssetRouterBaseDepositFinalizedAssetRouter, chainId []*big.Int, assetId [][32]byte) (event.Subscription, error) {

	var chainIdRule []interface{}
	for _, chainIdItem := range chainId {
		chainIdRule = append(chainIdRule, chainIdItem)
	}
	var assetIdRule []interface{}
	for _, assetIdItem := range assetId {
		assetIdRule = append(assetIdRule, assetIdItem)
	}

	logs, sub, err := _IAssetRouterBase.contract.WatchLogs(opts, "DepositFinalizedAssetRouter", chainIdRule, assetIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IAssetRouterBaseDepositFinalizedAssetRouter)
				if err := _IAssetRouterBase.contract.UnpackLog(event, "DepositFinalizedAssetRouter", log); err != nil {
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
func (_IAssetRouterBase *IAssetRouterBaseFilterer) ParseDepositFinalizedAssetRouter(log types.Log) (*IAssetRouterBaseDepositFinalizedAssetRouter, error) {
	event := new(IAssetRouterBaseDepositFinalizedAssetRouter)
	if err := _IAssetRouterBase.contract.UnpackLog(event, "DepositFinalizedAssetRouter", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
