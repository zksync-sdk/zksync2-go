// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package IL1Bridge

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

// IL1BridgeMetaData contains all meta data concerning the IL1Bridge contract.
var IL1BridgeMetaData = &bind.MetaData{
	ABI: "[{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"l1Token\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"ClaimedFailedDeposit\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"l2DepositTxHash\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"l1Token\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"DepositInitiated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"l1Token\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"WithdrawalFinalized\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_depositSender\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_l1Token\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"_l2TxHash\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"_l2BlockNumber\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_l2MessageIndex\",\"type\":\"uint256\"},{\"internalType\":\"uint16\",\"name\":\"_l2TxNumberInBlock\",\"type\":\"uint16\"},{\"internalType\":\"bytes32[]\",\"name\":\"_merkleProof\",\"type\":\"bytes32[]\"}],\"name\":\"claimFailedDeposit\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_l2Receiver\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_l1Token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_l2TxGasLimit\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_l2TxGasPerPubdataByte\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"_refundRecipient\",\"type\":\"address\"}],\"name\":\"deposit\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"txHash\",\"type\":\"bytes32\"}],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_l2BlockNumber\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_l2MessageIndex\",\"type\":\"uint256\"},{\"internalType\":\"uint16\",\"name\":\"_l2TxNumberInBlock\",\"type\":\"uint16\"},{\"internalType\":\"bytes\",\"name\":\"_message\",\"type\":\"bytes\"},{\"internalType\":\"bytes32[]\",\"name\":\"_merkleProof\",\"type\":\"bytes32[]\"}],\"name\":\"finalizeWithdrawal\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_l2BlockNumber\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_l2MessageIndex\",\"type\":\"uint256\"}],\"name\":\"isWithdrawalFinalized\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"l2Bridge\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_l1Token\",\"type\":\"address\"}],\"name\":\"l2TokenAddress\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
}

// IL1BridgeABI is the input ABI used to generate the binding from.
// Deprecated: Use IL1BridgeMetaData.ABI instead.
var IL1BridgeABI = IL1BridgeMetaData.ABI

// IL1Bridge is an auto generated Go binding around an Ethereum contract.
type IL1Bridge struct {
	IL1BridgeCaller     // Read-only binding to the contract
	IL1BridgeTransactor // Write-only binding to the contract
	IL1BridgeFilterer   // Log filterer for contract events
}

// IL1BridgeCaller is an auto generated read-only Go binding around an Ethereum contract.
type IL1BridgeCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IL1BridgeTransactor is an auto generated write-only Go binding around an Ethereum contract.
type IL1BridgeTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IL1BridgeFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type IL1BridgeFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IL1BridgeSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type IL1BridgeSession struct {
	Contract     *IL1Bridge        // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// IL1BridgeCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type IL1BridgeCallerSession struct {
	Contract *IL1BridgeCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts    // Call options to use throughout this session
}

// IL1BridgeTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type IL1BridgeTransactorSession struct {
	Contract     *IL1BridgeTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts    // Transaction auth options to use throughout this session
}

// IL1BridgeRaw is an auto generated low-level Go binding around an Ethereum contract.
type IL1BridgeRaw struct {
	Contract *IL1Bridge // Generic contract binding to access the raw methods on
}

// IL1BridgeCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type IL1BridgeCallerRaw struct {
	Contract *IL1BridgeCaller // Generic read-only contract binding to access the raw methods on
}

// IL1BridgeTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type IL1BridgeTransactorRaw struct {
	Contract *IL1BridgeTransactor // Generic write-only contract binding to access the raw methods on
}

// NewIL1Bridge creates a new instance of IL1Bridge, bound to a specific deployed contract.
func NewIL1Bridge(address common.Address, backend bind.ContractBackend) (*IL1Bridge, error) {
	contract, err := bindIL1Bridge(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &IL1Bridge{IL1BridgeCaller: IL1BridgeCaller{contract: contract}, IL1BridgeTransactor: IL1BridgeTransactor{contract: contract}, IL1BridgeFilterer: IL1BridgeFilterer{contract: contract}}, nil
}

// NewIL1BridgeCaller creates a new read-only instance of IL1Bridge, bound to a specific deployed contract.
func NewIL1BridgeCaller(address common.Address, caller bind.ContractCaller) (*IL1BridgeCaller, error) {
	contract, err := bindIL1Bridge(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &IL1BridgeCaller{contract: contract}, nil
}

// NewIL1BridgeTransactor creates a new write-only instance of IL1Bridge, bound to a specific deployed contract.
func NewIL1BridgeTransactor(address common.Address, transactor bind.ContractTransactor) (*IL1BridgeTransactor, error) {
	contract, err := bindIL1Bridge(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &IL1BridgeTransactor{contract: contract}, nil
}

// NewIL1BridgeFilterer creates a new log filterer instance of IL1Bridge, bound to a specific deployed contract.
func NewIL1BridgeFilterer(address common.Address, filterer bind.ContractFilterer) (*IL1BridgeFilterer, error) {
	contract, err := bindIL1Bridge(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &IL1BridgeFilterer{contract: contract}, nil
}

// bindIL1Bridge binds a generic wrapper to an already deployed contract.
func bindIL1Bridge(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := IL1BridgeMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IL1Bridge *IL1BridgeRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IL1Bridge.Contract.IL1BridgeCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IL1Bridge *IL1BridgeRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IL1Bridge.Contract.IL1BridgeTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IL1Bridge *IL1BridgeRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IL1Bridge.Contract.IL1BridgeTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IL1Bridge *IL1BridgeCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IL1Bridge.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IL1Bridge *IL1BridgeTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IL1Bridge.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IL1Bridge *IL1BridgeTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IL1Bridge.Contract.contract.Transact(opts, method, params...)
}

// IsWithdrawalFinalized is a free data retrieval call binding the contract method 0x4bed8212.
//
// Solidity: function isWithdrawalFinalized(uint256 _l2BlockNumber, uint256 _l2MessageIndex) view returns(bool)
func (_IL1Bridge *IL1BridgeCaller) IsWithdrawalFinalized(opts *bind.CallOpts, _l2BlockNumber *big.Int, _l2MessageIndex *big.Int) (bool, error) {
	var out []interface{}
	err := _IL1Bridge.contract.Call(opts, &out, "isWithdrawalFinalized", _l2BlockNumber, _l2MessageIndex)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsWithdrawalFinalized is a free data retrieval call binding the contract method 0x4bed8212.
//
// Solidity: function isWithdrawalFinalized(uint256 _l2BlockNumber, uint256 _l2MessageIndex) view returns(bool)
func (_IL1Bridge *IL1BridgeSession) IsWithdrawalFinalized(_l2BlockNumber *big.Int, _l2MessageIndex *big.Int) (bool, error) {
	return _IL1Bridge.Contract.IsWithdrawalFinalized(&_IL1Bridge.CallOpts, _l2BlockNumber, _l2MessageIndex)
}

// IsWithdrawalFinalized is a free data retrieval call binding the contract method 0x4bed8212.
//
// Solidity: function isWithdrawalFinalized(uint256 _l2BlockNumber, uint256 _l2MessageIndex) view returns(bool)
func (_IL1Bridge *IL1BridgeCallerSession) IsWithdrawalFinalized(_l2BlockNumber *big.Int, _l2MessageIndex *big.Int) (bool, error) {
	return _IL1Bridge.Contract.IsWithdrawalFinalized(&_IL1Bridge.CallOpts, _l2BlockNumber, _l2MessageIndex)
}

// L2Bridge is a free data retrieval call binding the contract method 0xae1f6aaf.
//
// Solidity: function l2Bridge() view returns(address)
func (_IL1Bridge *IL1BridgeCaller) L2Bridge(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _IL1Bridge.contract.Call(opts, &out, "l2Bridge")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// L2Bridge is a free data retrieval call binding the contract method 0xae1f6aaf.
//
// Solidity: function l2Bridge() view returns(address)
func (_IL1Bridge *IL1BridgeSession) L2Bridge() (common.Address, error) {
	return _IL1Bridge.Contract.L2Bridge(&_IL1Bridge.CallOpts)
}

// L2Bridge is a free data retrieval call binding the contract method 0xae1f6aaf.
//
// Solidity: function l2Bridge() view returns(address)
func (_IL1Bridge *IL1BridgeCallerSession) L2Bridge() (common.Address, error) {
	return _IL1Bridge.Contract.L2Bridge(&_IL1Bridge.CallOpts)
}

// L2TokenAddress is a free data retrieval call binding the contract method 0xf5f15168.
//
// Solidity: function l2TokenAddress(address _l1Token) view returns(address)
func (_IL1Bridge *IL1BridgeCaller) L2TokenAddress(opts *bind.CallOpts, _l1Token common.Address) (common.Address, error) {
	var out []interface{}
	err := _IL1Bridge.contract.Call(opts, &out, "l2TokenAddress", _l1Token)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// L2TokenAddress is a free data retrieval call binding the contract method 0xf5f15168.
//
// Solidity: function l2TokenAddress(address _l1Token) view returns(address)
func (_IL1Bridge *IL1BridgeSession) L2TokenAddress(_l1Token common.Address) (common.Address, error) {
	return _IL1Bridge.Contract.L2TokenAddress(&_IL1Bridge.CallOpts, _l1Token)
}

// L2TokenAddress is a free data retrieval call binding the contract method 0xf5f15168.
//
// Solidity: function l2TokenAddress(address _l1Token) view returns(address)
func (_IL1Bridge *IL1BridgeCallerSession) L2TokenAddress(_l1Token common.Address) (common.Address, error) {
	return _IL1Bridge.Contract.L2TokenAddress(&_IL1Bridge.CallOpts, _l1Token)
}

// ClaimFailedDeposit is a paid mutator transaction binding the contract method 0x19fa7f62.
//
// Solidity: function claimFailedDeposit(address _depositSender, address _l1Token, bytes32 _l2TxHash, uint256 _l2BlockNumber, uint256 _l2MessageIndex, uint16 _l2TxNumberInBlock, bytes32[] _merkleProof) returns()
func (_IL1Bridge *IL1BridgeTransactor) ClaimFailedDeposit(opts *bind.TransactOpts, _depositSender common.Address, _l1Token common.Address, _l2TxHash [32]byte, _l2BlockNumber *big.Int, _l2MessageIndex *big.Int, _l2TxNumberInBlock uint16, _merkleProof [][32]byte) (*types.Transaction, error) {
	return _IL1Bridge.contract.Transact(opts, "claimFailedDeposit", _depositSender, _l1Token, _l2TxHash, _l2BlockNumber, _l2MessageIndex, _l2TxNumberInBlock, _merkleProof)
}

// ClaimFailedDeposit is a paid mutator transaction binding the contract method 0x19fa7f62.
//
// Solidity: function claimFailedDeposit(address _depositSender, address _l1Token, bytes32 _l2TxHash, uint256 _l2BlockNumber, uint256 _l2MessageIndex, uint16 _l2TxNumberInBlock, bytes32[] _merkleProof) returns()
func (_IL1Bridge *IL1BridgeSession) ClaimFailedDeposit(_depositSender common.Address, _l1Token common.Address, _l2TxHash [32]byte, _l2BlockNumber *big.Int, _l2MessageIndex *big.Int, _l2TxNumberInBlock uint16, _merkleProof [][32]byte) (*types.Transaction, error) {
	return _IL1Bridge.Contract.ClaimFailedDeposit(&_IL1Bridge.TransactOpts, _depositSender, _l1Token, _l2TxHash, _l2BlockNumber, _l2MessageIndex, _l2TxNumberInBlock, _merkleProof)
}

// ClaimFailedDeposit is a paid mutator transaction binding the contract method 0x19fa7f62.
//
// Solidity: function claimFailedDeposit(address _depositSender, address _l1Token, bytes32 _l2TxHash, uint256 _l2BlockNumber, uint256 _l2MessageIndex, uint16 _l2TxNumberInBlock, bytes32[] _merkleProof) returns()
func (_IL1Bridge *IL1BridgeTransactorSession) ClaimFailedDeposit(_depositSender common.Address, _l1Token common.Address, _l2TxHash [32]byte, _l2BlockNumber *big.Int, _l2MessageIndex *big.Int, _l2TxNumberInBlock uint16, _merkleProof [][32]byte) (*types.Transaction, error) {
	return _IL1Bridge.Contract.ClaimFailedDeposit(&_IL1Bridge.TransactOpts, _depositSender, _l1Token, _l2TxHash, _l2BlockNumber, _l2MessageIndex, _l2TxNumberInBlock, _merkleProof)
}

// Deposit is a paid mutator transaction binding the contract method 0xe8b99b1b.
//
// Solidity: function deposit(address _l2Receiver, address _l1Token, uint256 _amount, uint256 _l2TxGasLimit, uint256 _l2TxGasPerPubdataByte, address _refundRecipient) payable returns(bytes32 txHash)
func (_IL1Bridge *IL1BridgeTransactor) Deposit(opts *bind.TransactOpts, _l2Receiver common.Address, _l1Token common.Address, _amount *big.Int, _l2TxGasLimit *big.Int, _l2TxGasPerPubdataByte *big.Int, _refundRecipient common.Address) (*types.Transaction, error) {
	return _IL1Bridge.contract.Transact(opts, "deposit", _l2Receiver, _l1Token, _amount, _l2TxGasLimit, _l2TxGasPerPubdataByte, _refundRecipient)
}

// Deposit is a paid mutator transaction binding the contract method 0xe8b99b1b.
//
// Solidity: function deposit(address _l2Receiver, address _l1Token, uint256 _amount, uint256 _l2TxGasLimit, uint256 _l2TxGasPerPubdataByte, address _refundRecipient) payable returns(bytes32 txHash)
func (_IL1Bridge *IL1BridgeSession) Deposit(_l2Receiver common.Address, _l1Token common.Address, _amount *big.Int, _l2TxGasLimit *big.Int, _l2TxGasPerPubdataByte *big.Int, _refundRecipient common.Address) (*types.Transaction, error) {
	return _IL1Bridge.Contract.Deposit(&_IL1Bridge.TransactOpts, _l2Receiver, _l1Token, _amount, _l2TxGasLimit, _l2TxGasPerPubdataByte, _refundRecipient)
}

// Deposit is a paid mutator transaction binding the contract method 0xe8b99b1b.
//
// Solidity: function deposit(address _l2Receiver, address _l1Token, uint256 _amount, uint256 _l2TxGasLimit, uint256 _l2TxGasPerPubdataByte, address _refundRecipient) payable returns(bytes32 txHash)
func (_IL1Bridge *IL1BridgeTransactorSession) Deposit(_l2Receiver common.Address, _l1Token common.Address, _amount *big.Int, _l2TxGasLimit *big.Int, _l2TxGasPerPubdataByte *big.Int, _refundRecipient common.Address) (*types.Transaction, error) {
	return _IL1Bridge.Contract.Deposit(&_IL1Bridge.TransactOpts, _l2Receiver, _l1Token, _amount, _l2TxGasLimit, _l2TxGasPerPubdataByte, _refundRecipient)
}

// FinalizeWithdrawal is a paid mutator transaction binding the contract method 0x11a2ccc1.
//
// Solidity: function finalizeWithdrawal(uint256 _l2BlockNumber, uint256 _l2MessageIndex, uint16 _l2TxNumberInBlock, bytes _message, bytes32[] _merkleProof) returns()
func (_IL1Bridge *IL1BridgeTransactor) FinalizeWithdrawal(opts *bind.TransactOpts, _l2BlockNumber *big.Int, _l2MessageIndex *big.Int, _l2TxNumberInBlock uint16, _message []byte, _merkleProof [][32]byte) (*types.Transaction, error) {
	return _IL1Bridge.contract.Transact(opts, "finalizeWithdrawal", _l2BlockNumber, _l2MessageIndex, _l2TxNumberInBlock, _message, _merkleProof)
}

// FinalizeWithdrawal is a paid mutator transaction binding the contract method 0x11a2ccc1.
//
// Solidity: function finalizeWithdrawal(uint256 _l2BlockNumber, uint256 _l2MessageIndex, uint16 _l2TxNumberInBlock, bytes _message, bytes32[] _merkleProof) returns()
func (_IL1Bridge *IL1BridgeSession) FinalizeWithdrawal(_l2BlockNumber *big.Int, _l2MessageIndex *big.Int, _l2TxNumberInBlock uint16, _message []byte, _merkleProof [][32]byte) (*types.Transaction, error) {
	return _IL1Bridge.Contract.FinalizeWithdrawal(&_IL1Bridge.TransactOpts, _l2BlockNumber, _l2MessageIndex, _l2TxNumberInBlock, _message, _merkleProof)
}

// FinalizeWithdrawal is a paid mutator transaction binding the contract method 0x11a2ccc1.
//
// Solidity: function finalizeWithdrawal(uint256 _l2BlockNumber, uint256 _l2MessageIndex, uint16 _l2TxNumberInBlock, bytes _message, bytes32[] _merkleProof) returns()
func (_IL1Bridge *IL1BridgeTransactorSession) FinalizeWithdrawal(_l2BlockNumber *big.Int, _l2MessageIndex *big.Int, _l2TxNumberInBlock uint16, _message []byte, _merkleProof [][32]byte) (*types.Transaction, error) {
	return _IL1Bridge.Contract.FinalizeWithdrawal(&_IL1Bridge.TransactOpts, _l2BlockNumber, _l2MessageIndex, _l2TxNumberInBlock, _message, _merkleProof)
}

// IL1BridgeClaimedFailedDepositIterator is returned from FilterClaimedFailedDeposit and is used to iterate over the raw logs and unpacked data for ClaimedFailedDeposit events raised by the IL1Bridge contract.
type IL1BridgeClaimedFailedDepositIterator struct {
	Event *IL1BridgeClaimedFailedDeposit // Event containing the contract specifics and raw log

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
func (it *IL1BridgeClaimedFailedDepositIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IL1BridgeClaimedFailedDeposit)
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
		it.Event = new(IL1BridgeClaimedFailedDeposit)
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
func (it *IL1BridgeClaimedFailedDepositIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IL1BridgeClaimedFailedDepositIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IL1BridgeClaimedFailedDeposit represents a ClaimedFailedDeposit event raised by the IL1Bridge contract.
type IL1BridgeClaimedFailedDeposit struct {
	To      common.Address
	L1Token common.Address
	Amount  *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterClaimedFailedDeposit is a free log retrieval operation binding the contract event 0xbe066dc591f4a444f75176d387c3e6c775e5706d9ea9a91d11eb49030c66cf60.
//
// Solidity: event ClaimedFailedDeposit(address indexed to, address indexed l1Token, uint256 amount)
func (_IL1Bridge *IL1BridgeFilterer) FilterClaimedFailedDeposit(opts *bind.FilterOpts, to []common.Address, l1Token []common.Address) (*IL1BridgeClaimedFailedDepositIterator, error) {

	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}
	var l1TokenRule []interface{}
	for _, l1TokenItem := range l1Token {
		l1TokenRule = append(l1TokenRule, l1TokenItem)
	}

	logs, sub, err := _IL1Bridge.contract.FilterLogs(opts, "ClaimedFailedDeposit", toRule, l1TokenRule)
	if err != nil {
		return nil, err
	}
	return &IL1BridgeClaimedFailedDepositIterator{contract: _IL1Bridge.contract, event: "ClaimedFailedDeposit", logs: logs, sub: sub}, nil
}

// WatchClaimedFailedDeposit is a free log subscription operation binding the contract event 0xbe066dc591f4a444f75176d387c3e6c775e5706d9ea9a91d11eb49030c66cf60.
//
// Solidity: event ClaimedFailedDeposit(address indexed to, address indexed l1Token, uint256 amount)
func (_IL1Bridge *IL1BridgeFilterer) WatchClaimedFailedDeposit(opts *bind.WatchOpts, sink chan<- *IL1BridgeClaimedFailedDeposit, to []common.Address, l1Token []common.Address) (event.Subscription, error) {

	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}
	var l1TokenRule []interface{}
	for _, l1TokenItem := range l1Token {
		l1TokenRule = append(l1TokenRule, l1TokenItem)
	}

	logs, sub, err := _IL1Bridge.contract.WatchLogs(opts, "ClaimedFailedDeposit", toRule, l1TokenRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IL1BridgeClaimedFailedDeposit)
				if err := _IL1Bridge.contract.UnpackLog(event, "ClaimedFailedDeposit", log); err != nil {
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

// ParseClaimedFailedDeposit is a log parse operation binding the contract event 0xbe066dc591f4a444f75176d387c3e6c775e5706d9ea9a91d11eb49030c66cf60.
//
// Solidity: event ClaimedFailedDeposit(address indexed to, address indexed l1Token, uint256 amount)
func (_IL1Bridge *IL1BridgeFilterer) ParseClaimedFailedDeposit(log types.Log) (*IL1BridgeClaimedFailedDeposit, error) {
	event := new(IL1BridgeClaimedFailedDeposit)
	if err := _IL1Bridge.contract.UnpackLog(event, "ClaimedFailedDeposit", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// IL1BridgeDepositInitiatedIterator is returned from FilterDepositInitiated and is used to iterate over the raw logs and unpacked data for DepositInitiated events raised by the IL1Bridge contract.
type IL1BridgeDepositInitiatedIterator struct {
	Event *IL1BridgeDepositInitiated // Event containing the contract specifics and raw log

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
func (it *IL1BridgeDepositInitiatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IL1BridgeDepositInitiated)
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
		it.Event = new(IL1BridgeDepositInitiated)
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
func (it *IL1BridgeDepositInitiatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IL1BridgeDepositInitiatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IL1BridgeDepositInitiated represents a DepositInitiated event raised by the IL1Bridge contract.
type IL1BridgeDepositInitiated struct {
	L2DepositTxHash [32]byte
	From            common.Address
	To              common.Address
	L1Token         common.Address
	Amount          *big.Int
	Raw             types.Log // Blockchain specific contextual infos
}

// FilterDepositInitiated is a free log retrieval operation binding the contract event 0xdd341179f4edc78148d894d0213a96d212af2cbaf223d19ef6d483bdd47ab81d.
//
// Solidity: event DepositInitiated(bytes32 indexed l2DepositTxHash, address indexed from, address indexed to, address l1Token, uint256 amount)
func (_IL1Bridge *IL1BridgeFilterer) FilterDepositInitiated(opts *bind.FilterOpts, l2DepositTxHash [][32]byte, from []common.Address, to []common.Address) (*IL1BridgeDepositInitiatedIterator, error) {

	var l2DepositTxHashRule []interface{}
	for _, l2DepositTxHashItem := range l2DepositTxHash {
		l2DepositTxHashRule = append(l2DepositTxHashRule, l2DepositTxHashItem)
	}
	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _IL1Bridge.contract.FilterLogs(opts, "DepositInitiated", l2DepositTxHashRule, fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &IL1BridgeDepositInitiatedIterator{contract: _IL1Bridge.contract, event: "DepositInitiated", logs: logs, sub: sub}, nil
}

// WatchDepositInitiated is a free log subscription operation binding the contract event 0xdd341179f4edc78148d894d0213a96d212af2cbaf223d19ef6d483bdd47ab81d.
//
// Solidity: event DepositInitiated(bytes32 indexed l2DepositTxHash, address indexed from, address indexed to, address l1Token, uint256 amount)
func (_IL1Bridge *IL1BridgeFilterer) WatchDepositInitiated(opts *bind.WatchOpts, sink chan<- *IL1BridgeDepositInitiated, l2DepositTxHash [][32]byte, from []common.Address, to []common.Address) (event.Subscription, error) {

	var l2DepositTxHashRule []interface{}
	for _, l2DepositTxHashItem := range l2DepositTxHash {
		l2DepositTxHashRule = append(l2DepositTxHashRule, l2DepositTxHashItem)
	}
	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _IL1Bridge.contract.WatchLogs(opts, "DepositInitiated", l2DepositTxHashRule, fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IL1BridgeDepositInitiated)
				if err := _IL1Bridge.contract.UnpackLog(event, "DepositInitiated", log); err != nil {
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

// ParseDepositInitiated is a log parse operation binding the contract event 0xdd341179f4edc78148d894d0213a96d212af2cbaf223d19ef6d483bdd47ab81d.
//
// Solidity: event DepositInitiated(bytes32 indexed l2DepositTxHash, address indexed from, address indexed to, address l1Token, uint256 amount)
func (_IL1Bridge *IL1BridgeFilterer) ParseDepositInitiated(log types.Log) (*IL1BridgeDepositInitiated, error) {
	event := new(IL1BridgeDepositInitiated)
	if err := _IL1Bridge.contract.UnpackLog(event, "DepositInitiated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// IL1BridgeWithdrawalFinalizedIterator is returned from FilterWithdrawalFinalized and is used to iterate over the raw logs and unpacked data for WithdrawalFinalized events raised by the IL1Bridge contract.
type IL1BridgeWithdrawalFinalizedIterator struct {
	Event *IL1BridgeWithdrawalFinalized // Event containing the contract specifics and raw log

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
func (it *IL1BridgeWithdrawalFinalizedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IL1BridgeWithdrawalFinalized)
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
		it.Event = new(IL1BridgeWithdrawalFinalized)
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
func (it *IL1BridgeWithdrawalFinalizedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IL1BridgeWithdrawalFinalizedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IL1BridgeWithdrawalFinalized represents a WithdrawalFinalized event raised by the IL1Bridge contract.
type IL1BridgeWithdrawalFinalized struct {
	To      common.Address
	L1Token common.Address
	Amount  *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterWithdrawalFinalized is a free log retrieval operation binding the contract event 0xac1b18083978656d557d6e91c88203585cfda1031bdb14538327121ef140d383.
//
// Solidity: event WithdrawalFinalized(address indexed to, address indexed l1Token, uint256 amount)
func (_IL1Bridge *IL1BridgeFilterer) FilterWithdrawalFinalized(opts *bind.FilterOpts, to []common.Address, l1Token []common.Address) (*IL1BridgeWithdrawalFinalizedIterator, error) {

	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}
	var l1TokenRule []interface{}
	for _, l1TokenItem := range l1Token {
		l1TokenRule = append(l1TokenRule, l1TokenItem)
	}

	logs, sub, err := _IL1Bridge.contract.FilterLogs(opts, "WithdrawalFinalized", toRule, l1TokenRule)
	if err != nil {
		return nil, err
	}
	return &IL1BridgeWithdrawalFinalizedIterator{contract: _IL1Bridge.contract, event: "WithdrawalFinalized", logs: logs, sub: sub}, nil
}

// WatchWithdrawalFinalized is a free log subscription operation binding the contract event 0xac1b18083978656d557d6e91c88203585cfda1031bdb14538327121ef140d383.
//
// Solidity: event WithdrawalFinalized(address indexed to, address indexed l1Token, uint256 amount)
func (_IL1Bridge *IL1BridgeFilterer) WatchWithdrawalFinalized(opts *bind.WatchOpts, sink chan<- *IL1BridgeWithdrawalFinalized, to []common.Address, l1Token []common.Address) (event.Subscription, error) {

	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}
	var l1TokenRule []interface{}
	for _, l1TokenItem := range l1Token {
		l1TokenRule = append(l1TokenRule, l1TokenItem)
	}

	logs, sub, err := _IL1Bridge.contract.WatchLogs(opts, "WithdrawalFinalized", toRule, l1TokenRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IL1BridgeWithdrawalFinalized)
				if err := _IL1Bridge.contract.UnpackLog(event, "WithdrawalFinalized", log); err != nil {
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

// ParseWithdrawalFinalized is a log parse operation binding the contract event 0xac1b18083978656d557d6e91c88203585cfda1031bdb14538327121ef140d383.
//
// Solidity: event WithdrawalFinalized(address indexed to, address indexed l1Token, uint256 amount)
func (_IL1Bridge *IL1BridgeFilterer) ParseWithdrawalFinalized(log types.Log) (*IL1BridgeWithdrawalFinalized, error) {
	event := new(IL1BridgeWithdrawalFinalized)
	if err := _IL1Bridge.contract.UnpackLog(event, "WithdrawalFinalized", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
