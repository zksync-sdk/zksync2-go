// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package IEthToken

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

// IEthTokenMetaData contains all meta data concerning the IEthToken contract.
var IEthTokenMetaData = &bind.MetaData{
	ABI: "[{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"Mint\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Transfer\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"_l1Receiver\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"Withdrawal\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"decimals\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_account\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"mint\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"name\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"symbol\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"totalSupply\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"transferFromTo\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_l1Receiver\",\"type\":\"address\"}],\"name\":\"withdraw\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"}]",
}

// IEthTokenABI is the input ABI used to generate the binding from.
// Deprecated: Use IEthTokenMetaData.ABI instead.
var IEthTokenABI = IEthTokenMetaData.ABI

// IEthToken is an auto generated Go binding around an Ethereum contract.
type IEthToken struct {
	IEthTokenCaller     // Read-only binding to the contract
	IEthTokenTransactor // Write-only binding to the contract
	IEthTokenFilterer   // Log filterer for contract events
}

// IEthTokenCaller is an auto generated read-only Go binding around an Ethereum contract.
type IEthTokenCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IEthTokenTransactor is an auto generated write-only Go binding around an Ethereum contract.
type IEthTokenTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IEthTokenFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type IEthTokenFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IEthTokenSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type IEthTokenSession struct {
	Contract     *IEthToken        // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// IEthTokenCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type IEthTokenCallerSession struct {
	Contract *IEthTokenCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts    // Call options to use throughout this session
}

// IEthTokenTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type IEthTokenTransactorSession struct {
	Contract     *IEthTokenTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts    // Transaction auth options to use throughout this session
}

// IEthTokenRaw is an auto generated low-level Go binding around an Ethereum contract.
type IEthTokenRaw struct {
	Contract *IEthToken // Generic contract binding to access the raw methods on
}

// IEthTokenCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type IEthTokenCallerRaw struct {
	Contract *IEthTokenCaller // Generic read-only contract binding to access the raw methods on
}

// IEthTokenTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type IEthTokenTransactorRaw struct {
	Contract *IEthTokenTransactor // Generic write-only contract binding to access the raw methods on
}

// NewIEthToken creates a new instance of IEthToken, bound to a specific deployed contract.
func NewIEthToken(address common.Address, backend bind.ContractBackend) (*IEthToken, error) {
	contract, err := bindIEthToken(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &IEthToken{IEthTokenCaller: IEthTokenCaller{contract: contract}, IEthTokenTransactor: IEthTokenTransactor{contract: contract}, IEthTokenFilterer: IEthTokenFilterer{contract: contract}}, nil
}

// NewIEthTokenCaller creates a new read-only instance of IEthToken, bound to a specific deployed contract.
func NewIEthTokenCaller(address common.Address, caller bind.ContractCaller) (*IEthTokenCaller, error) {
	contract, err := bindIEthToken(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &IEthTokenCaller{contract: contract}, nil
}

// NewIEthTokenTransactor creates a new write-only instance of IEthToken, bound to a specific deployed contract.
func NewIEthTokenTransactor(address common.Address, transactor bind.ContractTransactor) (*IEthTokenTransactor, error) {
	contract, err := bindIEthToken(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &IEthTokenTransactor{contract: contract}, nil
}

// NewIEthTokenFilterer creates a new log filterer instance of IEthToken, bound to a specific deployed contract.
func NewIEthTokenFilterer(address common.Address, filterer bind.ContractFilterer) (*IEthTokenFilterer, error) {
	contract, err := bindIEthToken(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &IEthTokenFilterer{contract: contract}, nil
}

// bindIEthToken binds a generic wrapper to an already deployed contract.
func bindIEthToken(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(IEthTokenABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IEthToken *IEthTokenRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IEthToken.Contract.IEthTokenCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IEthToken *IEthTokenRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IEthToken.Contract.IEthTokenTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IEthToken *IEthTokenRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IEthToken.Contract.IEthTokenTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IEthToken *IEthTokenCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IEthToken.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IEthToken *IEthTokenTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IEthToken.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IEthToken *IEthTokenTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IEthToken.Contract.contract.Transact(opts, method, params...)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address ) view returns(uint256)
func (_IEthToken *IEthTokenCaller) BalanceOf(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _IEthToken.contract.Call(opts, &out, "balanceOf", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address ) view returns(uint256)
func (_IEthToken *IEthTokenSession) BalanceOf(arg0 common.Address) (*big.Int, error) {
	return _IEthToken.Contract.BalanceOf(&_IEthToken.CallOpts, arg0)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address ) view returns(uint256)
func (_IEthToken *IEthTokenCallerSession) BalanceOf(arg0 common.Address) (*big.Int, error) {
	return _IEthToken.Contract.BalanceOf(&_IEthToken.CallOpts, arg0)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() pure returns(uint8)
func (_IEthToken *IEthTokenCaller) Decimals(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _IEthToken.contract.Call(opts, &out, "decimals")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() pure returns(uint8)
func (_IEthToken *IEthTokenSession) Decimals() (uint8, error) {
	return _IEthToken.Contract.Decimals(&_IEthToken.CallOpts)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() pure returns(uint8)
func (_IEthToken *IEthTokenCallerSession) Decimals() (uint8, error) {
	return _IEthToken.Contract.Decimals(&_IEthToken.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() pure returns(string)
func (_IEthToken *IEthTokenCaller) Name(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _IEthToken.contract.Call(opts, &out, "name")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() pure returns(string)
func (_IEthToken *IEthTokenSession) Name() (string, error) {
	return _IEthToken.Contract.Name(&_IEthToken.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() pure returns(string)
func (_IEthToken *IEthTokenCallerSession) Name() (string, error) {
	return _IEthToken.Contract.Name(&_IEthToken.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() pure returns(string)
func (_IEthToken *IEthTokenCaller) Symbol(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _IEthToken.contract.Call(opts, &out, "symbol")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() pure returns(string)
func (_IEthToken *IEthTokenSession) Symbol() (string, error) {
	return _IEthToken.Contract.Symbol(&_IEthToken.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() pure returns(string)
func (_IEthToken *IEthTokenCallerSession) Symbol() (string, error) {
	return _IEthToken.Contract.Symbol(&_IEthToken.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_IEthToken *IEthTokenCaller) TotalSupply(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _IEthToken.contract.Call(opts, &out, "totalSupply")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_IEthToken *IEthTokenSession) TotalSupply() (*big.Int, error) {
	return _IEthToken.Contract.TotalSupply(&_IEthToken.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_IEthToken *IEthTokenCallerSession) TotalSupply() (*big.Int, error) {
	return _IEthToken.Contract.TotalSupply(&_IEthToken.CallOpts)
}

// Mint is a paid mutator transaction binding the contract method 0x40c10f19.
//
// Solidity: function mint(address _account, uint256 _amount) returns()
func (_IEthToken *IEthTokenTransactor) Mint(opts *bind.TransactOpts, _account common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _IEthToken.contract.Transact(opts, "mint", _account, _amount)
}

// Mint is a paid mutator transaction binding the contract method 0x40c10f19.
//
// Solidity: function mint(address _account, uint256 _amount) returns()
func (_IEthToken *IEthTokenSession) Mint(_account common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _IEthToken.Contract.Mint(&_IEthToken.TransactOpts, _account, _amount)
}

// Mint is a paid mutator transaction binding the contract method 0x40c10f19.
//
// Solidity: function mint(address _account, uint256 _amount) returns()
func (_IEthToken *IEthTokenTransactorSession) Mint(_account common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _IEthToken.Contract.Mint(&_IEthToken.TransactOpts, _account, _amount)
}

// TransferFromTo is a paid mutator transaction binding the contract method 0x579952fc.
//
// Solidity: function transferFromTo(address _from, address _to, uint256 _amount) returns()
func (_IEthToken *IEthTokenTransactor) TransferFromTo(opts *bind.TransactOpts, _from common.Address, _to common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _IEthToken.contract.Transact(opts, "transferFromTo", _from, _to, _amount)
}

// TransferFromTo is a paid mutator transaction binding the contract method 0x579952fc.
//
// Solidity: function transferFromTo(address _from, address _to, uint256 _amount) returns()
func (_IEthToken *IEthTokenSession) TransferFromTo(_from common.Address, _to common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _IEthToken.Contract.TransferFromTo(&_IEthToken.TransactOpts, _from, _to, _amount)
}

// TransferFromTo is a paid mutator transaction binding the contract method 0x579952fc.
//
// Solidity: function transferFromTo(address _from, address _to, uint256 _amount) returns()
func (_IEthToken *IEthTokenTransactorSession) TransferFromTo(_from common.Address, _to common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _IEthToken.Contract.TransferFromTo(&_IEthToken.TransactOpts, _from, _to, _amount)
}

// Withdraw is a paid mutator transaction binding the contract method 0x51cff8d9.
//
// Solidity: function withdraw(address _l1Receiver) payable returns()
func (_IEthToken *IEthTokenTransactor) Withdraw(opts *bind.TransactOpts, _l1Receiver common.Address) (*types.Transaction, error) {
	return _IEthToken.contract.Transact(opts, "withdraw", _l1Receiver)
}

// Withdraw is a paid mutator transaction binding the contract method 0x51cff8d9.
//
// Solidity: function withdraw(address _l1Receiver) payable returns()
func (_IEthToken *IEthTokenSession) Withdraw(_l1Receiver common.Address) (*types.Transaction, error) {
	return _IEthToken.Contract.Withdraw(&_IEthToken.TransactOpts, _l1Receiver)
}

// Withdraw is a paid mutator transaction binding the contract method 0x51cff8d9.
//
// Solidity: function withdraw(address _l1Receiver) payable returns()
func (_IEthToken *IEthTokenTransactorSession) Withdraw(_l1Receiver common.Address) (*types.Transaction, error) {
	return _IEthToken.Contract.Withdraw(&_IEthToken.TransactOpts, _l1Receiver)
}

// IEthTokenMintIterator is returned from FilterMint and is used to iterate over the raw logs and unpacked data for Mint events raised by the IEthToken contract.
type IEthTokenMintIterator struct {
	Event *IEthTokenMint // Event containing the contract specifics and raw log

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
func (it *IEthTokenMintIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IEthTokenMint)
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
		it.Event = new(IEthTokenMint)
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
func (it *IEthTokenMintIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IEthTokenMintIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IEthTokenMint represents a Mint event raised by the IEthToken contract.
type IEthTokenMint struct {
	Account common.Address
	Amount  *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterMint is a free log retrieval operation binding the contract event 0x0f6798a560793a54c3bcfe86a93cde1e73087d944c0ea20544137d4121396885.
//
// Solidity: event Mint(address indexed account, uint256 amount)
func (_IEthToken *IEthTokenFilterer) FilterMint(opts *bind.FilterOpts, account []common.Address) (*IEthTokenMintIterator, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}

	logs, sub, err := _IEthToken.contract.FilterLogs(opts, "Mint", accountRule)
	if err != nil {
		return nil, err
	}
	return &IEthTokenMintIterator{contract: _IEthToken.contract, event: "Mint", logs: logs, sub: sub}, nil
}

// WatchMint is a free log subscription operation binding the contract event 0x0f6798a560793a54c3bcfe86a93cde1e73087d944c0ea20544137d4121396885.
//
// Solidity: event Mint(address indexed account, uint256 amount)
func (_IEthToken *IEthTokenFilterer) WatchMint(opts *bind.WatchOpts, sink chan<- *IEthTokenMint, account []common.Address) (event.Subscription, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}

	logs, sub, err := _IEthToken.contract.WatchLogs(opts, "Mint", accountRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IEthTokenMint)
				if err := _IEthToken.contract.UnpackLog(event, "Mint", log); err != nil {
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

// ParseMint is a log parse operation binding the contract event 0x0f6798a560793a54c3bcfe86a93cde1e73087d944c0ea20544137d4121396885.
//
// Solidity: event Mint(address indexed account, uint256 amount)
func (_IEthToken *IEthTokenFilterer) ParseMint(log types.Log) (*IEthTokenMint, error) {
	event := new(IEthTokenMint)
	if err := _IEthToken.contract.UnpackLog(event, "Mint", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// IEthTokenTransferIterator is returned from FilterTransfer and is used to iterate over the raw logs and unpacked data for Transfer events raised by the IEthToken contract.
type IEthTokenTransferIterator struct {
	Event *IEthTokenTransfer // Event containing the contract specifics and raw log

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
func (it *IEthTokenTransferIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IEthTokenTransfer)
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
		it.Event = new(IEthTokenTransfer)
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
func (it *IEthTokenTransferIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IEthTokenTransferIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IEthTokenTransfer represents a Transfer event raised by the IEthToken contract.
type IEthTokenTransfer struct {
	From  common.Address
	To    common.Address
	Value *big.Int
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterTransfer is a free log retrieval operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_IEthToken *IEthTokenFilterer) FilterTransfer(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*IEthTokenTransferIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _IEthToken.contract.FilterLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &IEthTokenTransferIterator{contract: _IEthToken.contract, event: "Transfer", logs: logs, sub: sub}, nil
}

// WatchTransfer is a free log subscription operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_IEthToken *IEthTokenFilterer) WatchTransfer(opts *bind.WatchOpts, sink chan<- *IEthTokenTransfer, from []common.Address, to []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _IEthToken.contract.WatchLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IEthTokenTransfer)
				if err := _IEthToken.contract.UnpackLog(event, "Transfer", log); err != nil {
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

// ParseTransfer is a log parse operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_IEthToken *IEthTokenFilterer) ParseTransfer(log types.Log) (*IEthTokenTransfer, error) {
	event := new(IEthTokenTransfer)
	if err := _IEthToken.contract.UnpackLog(event, "Transfer", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// IEthTokenWithdrawalIterator is returned from FilterWithdrawal and is used to iterate over the raw logs and unpacked data for Withdrawal events raised by the IEthToken contract.
type IEthTokenWithdrawalIterator struct {
	Event *IEthTokenWithdrawal // Event containing the contract specifics and raw log

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
func (it *IEthTokenWithdrawalIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IEthTokenWithdrawal)
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
		it.Event = new(IEthTokenWithdrawal)
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
func (it *IEthTokenWithdrawalIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IEthTokenWithdrawalIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IEthTokenWithdrawal represents a Withdrawal event raised by the IEthToken contract.
type IEthTokenWithdrawal struct {
	L1Receiver common.Address
	Amount     *big.Int
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterWithdrawal is a free log retrieval operation binding the contract event 0x7fcf532c15f0a6db0bd6d0e038bea71d30d808c7d98cb3bf7268a95bf5081b65.
//
// Solidity: event Withdrawal(address indexed _l1Receiver, uint256 _amount)
func (_IEthToken *IEthTokenFilterer) FilterWithdrawal(opts *bind.FilterOpts, _l1Receiver []common.Address) (*IEthTokenWithdrawalIterator, error) {

	var _l1ReceiverRule []interface{}
	for _, _l1ReceiverItem := range _l1Receiver {
		_l1ReceiverRule = append(_l1ReceiverRule, _l1ReceiverItem)
	}

	logs, sub, err := _IEthToken.contract.FilterLogs(opts, "Withdrawal", _l1ReceiverRule)
	if err != nil {
		return nil, err
	}
	return &IEthTokenWithdrawalIterator{contract: _IEthToken.contract, event: "Withdrawal", logs: logs, sub: sub}, nil
}

// WatchWithdrawal is a free log subscription operation binding the contract event 0x7fcf532c15f0a6db0bd6d0e038bea71d30d808c7d98cb3bf7268a95bf5081b65.
//
// Solidity: event Withdrawal(address indexed _l1Receiver, uint256 _amount)
func (_IEthToken *IEthTokenFilterer) WatchWithdrawal(opts *bind.WatchOpts, sink chan<- *IEthTokenWithdrawal, _l1Receiver []common.Address) (event.Subscription, error) {

	var _l1ReceiverRule []interface{}
	for _, _l1ReceiverItem := range _l1Receiver {
		_l1ReceiverRule = append(_l1ReceiverRule, _l1ReceiverItem)
	}

	logs, sub, err := _IEthToken.contract.WatchLogs(opts, "Withdrawal", _l1ReceiverRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IEthTokenWithdrawal)
				if err := _IEthToken.contract.UnpackLog(event, "Withdrawal", log); err != nil {
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

// ParseWithdrawal is a log parse operation binding the contract event 0x7fcf532c15f0a6db0bd6d0e038bea71d30d808c7d98cb3bf7268a95bf5081b65.
//
// Solidity: event Withdrawal(address indexed _l1Receiver, uint256 _amount)
func (_IEthToken *IEthTokenFilterer) ParseWithdrawal(log types.Log) (*IEthTokenWithdrawal, error) {
	event := new(IEthTokenWithdrawal)
	if err := _IEthToken.contract.UnpackLog(event, "Withdrawal", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
