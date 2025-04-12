// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package exampleauthserverpaymaster

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

// ExampleAuthServerPaymasterMetaData contains all meta data concerning the ExampleAuthServerPaymaster contract.
var ExampleAuthServerPaymasterMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"aaFactoryAddress\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"sessionKeyValidatorAddress\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"accountRecoveryValidatorAddress\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"webAuthValidatorAddress\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"AA_FACTORY_CONTRACT_ADDRESS\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"ACCOUNT_RECOVERY_VALIDATOR_CONTRACT_ADDRESS\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"SESSION_KEY_VALIDATOR_CONTRACT_ADDRESS\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"WEB_AUTH_VALIDATOR_CONTRACT_ADDRESS\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"_context\",\"type\":\"bytes\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"txType\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"from\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"to\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"gasLimit\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"gasPerPubdataByteLimit\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"maxFeePerGas\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"maxPriorityFeePerGas\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"paymaster\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"nonce\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"},{\"internalType\":\"uint256[4]\",\"name\":\"reserved\",\"type\":\"uint256[4]\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"signature\",\"type\":\"bytes\"},{\"internalType\":\"bytes32[]\",\"name\":\"factoryDeps\",\"type\":\"bytes32[]\"},{\"internalType\":\"bytes\",\"name\":\"paymasterInput\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"reservedDynamic\",\"type\":\"bytes\"}],\"internalType\":\"structTransaction\",\"name\":\"_transaction\",\"type\":\"tuple\"},{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"},{\"internalType\":\"enumExecutionResult\",\"name\":\"_txResult\",\"type\":\"uint8\"},{\"internalType\":\"uint256\",\"name\":\"_maxRefundedGas\",\"type\":\"uint256\"}],\"name\":\"postTransaction\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"txType\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"from\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"to\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"gasLimit\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"gasPerPubdataByteLimit\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"maxFeePerGas\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"maxPriorityFeePerGas\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"paymaster\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"nonce\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"},{\"internalType\":\"uint256[4]\",\"name\":\"reserved\",\"type\":\"uint256[4]\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"signature\",\"type\":\"bytes\"},{\"internalType\":\"bytes32[]\",\"name\":\"factoryDeps\",\"type\":\"bytes32[]\"},{\"internalType\":\"bytes\",\"name\":\"paymasterInput\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"reservedDynamic\",\"type\":\"bytes\"}],\"internalType\":\"structTransaction\",\"name\":\"_transaction\",\"type\":\"tuple\"}],\"name\":\"validateAndPayForPaymasterTransaction\",\"outputs\":[{\"internalType\":\"bytes4\",\"name\":\"magic\",\"type\":\"bytes4\"},{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"addresspayable\",\"name\":\"_to\",\"type\":\"address\"}],\"name\":\"withdraw\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"stateMutability\":\"payable\",\"type\":\"receive\"}]",
}

// ExampleAuthServerPaymasterABI is the input ABI used to generate the binding from.
// Deprecated: Use ExampleAuthServerPaymasterMetaData.ABI instead.
var ExampleAuthServerPaymasterABI = ExampleAuthServerPaymasterMetaData.ABI

// ExampleAuthServerPaymaster is an auto generated Go binding around an Ethereum contract.
type ExampleAuthServerPaymaster struct {
	ExampleAuthServerPaymasterCaller     // Read-only binding to the contract
	ExampleAuthServerPaymasterTransactor // Write-only binding to the contract
	ExampleAuthServerPaymasterFilterer   // Log filterer for contract events
}

// ExampleAuthServerPaymasterCaller is an auto generated read-only Go binding around an Ethereum contract.
type ExampleAuthServerPaymasterCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ExampleAuthServerPaymasterTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ExampleAuthServerPaymasterTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ExampleAuthServerPaymasterFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ExampleAuthServerPaymasterFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ExampleAuthServerPaymasterSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ExampleAuthServerPaymasterSession struct {
	Contract     *ExampleAuthServerPaymaster // Generic contract binding to set the session for
	CallOpts     bind.CallOpts               // Call options to use throughout this session
	TransactOpts bind.TransactOpts           // Transaction auth options to use throughout this session
}

// ExampleAuthServerPaymasterCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ExampleAuthServerPaymasterCallerSession struct {
	Contract *ExampleAuthServerPaymasterCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts                     // Call options to use throughout this session
}

// ExampleAuthServerPaymasterTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ExampleAuthServerPaymasterTransactorSession struct {
	Contract     *ExampleAuthServerPaymasterTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts                     // Transaction auth options to use throughout this session
}

// ExampleAuthServerPaymasterRaw is an auto generated low-level Go binding around an Ethereum contract.
type ExampleAuthServerPaymasterRaw struct {
	Contract *ExampleAuthServerPaymaster // Generic contract binding to access the raw methods on
}

// ExampleAuthServerPaymasterCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ExampleAuthServerPaymasterCallerRaw struct {
	Contract *ExampleAuthServerPaymasterCaller // Generic read-only contract binding to access the raw methods on
}

// ExampleAuthServerPaymasterTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ExampleAuthServerPaymasterTransactorRaw struct {
	Contract *ExampleAuthServerPaymasterTransactor // Generic write-only contract binding to access the raw methods on
}

// NewExampleAuthServerPaymaster creates a new instance of ExampleAuthServerPaymaster, bound to a specific deployed contract.
func NewExampleAuthServerPaymaster(address common.Address, backend bind.ContractBackend) (*ExampleAuthServerPaymaster, error) {
	contract, err := bindExampleAuthServerPaymaster(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &ExampleAuthServerPaymaster{ExampleAuthServerPaymasterCaller: ExampleAuthServerPaymasterCaller{contract: contract}, ExampleAuthServerPaymasterTransactor: ExampleAuthServerPaymasterTransactor{contract: contract}, ExampleAuthServerPaymasterFilterer: ExampleAuthServerPaymasterFilterer{contract: contract}}, nil
}

// NewExampleAuthServerPaymasterCaller creates a new read-only instance of ExampleAuthServerPaymaster, bound to a specific deployed contract.
func NewExampleAuthServerPaymasterCaller(address common.Address, caller bind.ContractCaller) (*ExampleAuthServerPaymasterCaller, error) {
	contract, err := bindExampleAuthServerPaymaster(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ExampleAuthServerPaymasterCaller{contract: contract}, nil
}

// NewExampleAuthServerPaymasterTransactor creates a new write-only instance of ExampleAuthServerPaymaster, bound to a specific deployed contract.
func NewExampleAuthServerPaymasterTransactor(address common.Address, transactor bind.ContractTransactor) (*ExampleAuthServerPaymasterTransactor, error) {
	contract, err := bindExampleAuthServerPaymaster(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ExampleAuthServerPaymasterTransactor{contract: contract}, nil
}

// NewExampleAuthServerPaymasterFilterer creates a new log filterer instance of ExampleAuthServerPaymaster, bound to a specific deployed contract.
func NewExampleAuthServerPaymasterFilterer(address common.Address, filterer bind.ContractFilterer) (*ExampleAuthServerPaymasterFilterer, error) {
	contract, err := bindExampleAuthServerPaymaster(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ExampleAuthServerPaymasterFilterer{contract: contract}, nil
}

// bindExampleAuthServerPaymaster binds a generic wrapper to an already deployed contract.
func bindExampleAuthServerPaymaster(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := ExampleAuthServerPaymasterMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ExampleAuthServerPaymaster *ExampleAuthServerPaymasterRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ExampleAuthServerPaymaster.Contract.ExampleAuthServerPaymasterCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ExampleAuthServerPaymaster *ExampleAuthServerPaymasterRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ExampleAuthServerPaymaster.Contract.ExampleAuthServerPaymasterTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ExampleAuthServerPaymaster *ExampleAuthServerPaymasterRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ExampleAuthServerPaymaster.Contract.ExampleAuthServerPaymasterTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ExampleAuthServerPaymaster *ExampleAuthServerPaymasterCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ExampleAuthServerPaymaster.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ExampleAuthServerPaymaster *ExampleAuthServerPaymasterTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ExampleAuthServerPaymaster.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ExampleAuthServerPaymaster *ExampleAuthServerPaymasterTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ExampleAuthServerPaymaster.Contract.contract.Transact(opts, method, params...)
}

// AAFACTORYCONTRACTADDRESS is a free data retrieval call binding the contract method 0x6ab2d73f.
//
// Solidity: function AA_FACTORY_CONTRACT_ADDRESS() view returns(address)
func (_ExampleAuthServerPaymaster *ExampleAuthServerPaymasterCaller) AAFACTORYCONTRACTADDRESS(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _ExampleAuthServerPaymaster.contract.Call(opts, &out, "AA_FACTORY_CONTRACT_ADDRESS")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// AAFACTORYCONTRACTADDRESS is a free data retrieval call binding the contract method 0x6ab2d73f.
//
// Solidity: function AA_FACTORY_CONTRACT_ADDRESS() view returns(address)
func (_ExampleAuthServerPaymaster *ExampleAuthServerPaymasterSession) AAFACTORYCONTRACTADDRESS() (common.Address, error) {
	return _ExampleAuthServerPaymaster.Contract.AAFACTORYCONTRACTADDRESS(&_ExampleAuthServerPaymaster.CallOpts)
}

// AAFACTORYCONTRACTADDRESS is a free data retrieval call binding the contract method 0x6ab2d73f.
//
// Solidity: function AA_FACTORY_CONTRACT_ADDRESS() view returns(address)
func (_ExampleAuthServerPaymaster *ExampleAuthServerPaymasterCallerSession) AAFACTORYCONTRACTADDRESS() (common.Address, error) {
	return _ExampleAuthServerPaymaster.Contract.AAFACTORYCONTRACTADDRESS(&_ExampleAuthServerPaymaster.CallOpts)
}

// ACCOUNTRECOVERYVALIDATORCONTRACTADDRESS is a free data retrieval call binding the contract method 0xb45c72e4.
//
// Solidity: function ACCOUNT_RECOVERY_VALIDATOR_CONTRACT_ADDRESS() view returns(address)
func (_ExampleAuthServerPaymaster *ExampleAuthServerPaymasterCaller) ACCOUNTRECOVERYVALIDATORCONTRACTADDRESS(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _ExampleAuthServerPaymaster.contract.Call(opts, &out, "ACCOUNT_RECOVERY_VALIDATOR_CONTRACT_ADDRESS")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// ACCOUNTRECOVERYVALIDATORCONTRACTADDRESS is a free data retrieval call binding the contract method 0xb45c72e4.
//
// Solidity: function ACCOUNT_RECOVERY_VALIDATOR_CONTRACT_ADDRESS() view returns(address)
func (_ExampleAuthServerPaymaster *ExampleAuthServerPaymasterSession) ACCOUNTRECOVERYVALIDATORCONTRACTADDRESS() (common.Address, error) {
	return _ExampleAuthServerPaymaster.Contract.ACCOUNTRECOVERYVALIDATORCONTRACTADDRESS(&_ExampleAuthServerPaymaster.CallOpts)
}

// ACCOUNTRECOVERYVALIDATORCONTRACTADDRESS is a free data retrieval call binding the contract method 0xb45c72e4.
//
// Solidity: function ACCOUNT_RECOVERY_VALIDATOR_CONTRACT_ADDRESS() view returns(address)
func (_ExampleAuthServerPaymaster *ExampleAuthServerPaymasterCallerSession) ACCOUNTRECOVERYVALIDATORCONTRACTADDRESS() (common.Address, error) {
	return _ExampleAuthServerPaymaster.Contract.ACCOUNTRECOVERYVALIDATORCONTRACTADDRESS(&_ExampleAuthServerPaymaster.CallOpts)
}

// SESSIONKEYVALIDATORCONTRACTADDRESS is a free data retrieval call binding the contract method 0x10677c76.
//
// Solidity: function SESSION_KEY_VALIDATOR_CONTRACT_ADDRESS() view returns(address)
func (_ExampleAuthServerPaymaster *ExampleAuthServerPaymasterCaller) SESSIONKEYVALIDATORCONTRACTADDRESS(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _ExampleAuthServerPaymaster.contract.Call(opts, &out, "SESSION_KEY_VALIDATOR_CONTRACT_ADDRESS")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// SESSIONKEYVALIDATORCONTRACTADDRESS is a free data retrieval call binding the contract method 0x10677c76.
//
// Solidity: function SESSION_KEY_VALIDATOR_CONTRACT_ADDRESS() view returns(address)
func (_ExampleAuthServerPaymaster *ExampleAuthServerPaymasterSession) SESSIONKEYVALIDATORCONTRACTADDRESS() (common.Address, error) {
	return _ExampleAuthServerPaymaster.Contract.SESSIONKEYVALIDATORCONTRACTADDRESS(&_ExampleAuthServerPaymaster.CallOpts)
}

// SESSIONKEYVALIDATORCONTRACTADDRESS is a free data retrieval call binding the contract method 0x10677c76.
//
// Solidity: function SESSION_KEY_VALIDATOR_CONTRACT_ADDRESS() view returns(address)
func (_ExampleAuthServerPaymaster *ExampleAuthServerPaymasterCallerSession) SESSIONKEYVALIDATORCONTRACTADDRESS() (common.Address, error) {
	return _ExampleAuthServerPaymaster.Contract.SESSIONKEYVALIDATORCONTRACTADDRESS(&_ExampleAuthServerPaymaster.CallOpts)
}

// WEBAUTHVALIDATORCONTRACTADDRESS is a free data retrieval call binding the contract method 0x8875a5da.
//
// Solidity: function WEB_AUTH_VALIDATOR_CONTRACT_ADDRESS() view returns(address)
func (_ExampleAuthServerPaymaster *ExampleAuthServerPaymasterCaller) WEBAUTHVALIDATORCONTRACTADDRESS(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _ExampleAuthServerPaymaster.contract.Call(opts, &out, "WEB_AUTH_VALIDATOR_CONTRACT_ADDRESS")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// WEBAUTHVALIDATORCONTRACTADDRESS is a free data retrieval call binding the contract method 0x8875a5da.
//
// Solidity: function WEB_AUTH_VALIDATOR_CONTRACT_ADDRESS() view returns(address)
func (_ExampleAuthServerPaymaster *ExampleAuthServerPaymasterSession) WEBAUTHVALIDATORCONTRACTADDRESS() (common.Address, error) {
	return _ExampleAuthServerPaymaster.Contract.WEBAUTHVALIDATORCONTRACTADDRESS(&_ExampleAuthServerPaymaster.CallOpts)
}

// WEBAUTHVALIDATORCONTRACTADDRESS is a free data retrieval call binding the contract method 0x8875a5da.
//
// Solidity: function WEB_AUTH_VALIDATOR_CONTRACT_ADDRESS() view returns(address)
func (_ExampleAuthServerPaymaster *ExampleAuthServerPaymasterCallerSession) WEBAUTHVALIDATORCONTRACTADDRESS() (common.Address, error) {
	return _ExampleAuthServerPaymaster.Contract.WEBAUTHVALIDATORCONTRACTADDRESS(&_ExampleAuthServerPaymaster.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_ExampleAuthServerPaymaster *ExampleAuthServerPaymasterCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _ExampleAuthServerPaymaster.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_ExampleAuthServerPaymaster *ExampleAuthServerPaymasterSession) Owner() (common.Address, error) {
	return _ExampleAuthServerPaymaster.Contract.Owner(&_ExampleAuthServerPaymaster.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_ExampleAuthServerPaymaster *ExampleAuthServerPaymasterCallerSession) Owner() (common.Address, error) {
	return _ExampleAuthServerPaymaster.Contract.Owner(&_ExampleAuthServerPaymaster.CallOpts)
}

// PostTransaction is a paid mutator transaction binding the contract method 0x817b17f0.
//
// Solidity: function postTransaction(bytes _context, (uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256[4],bytes,bytes,bytes32[],bytes,bytes) _transaction, bytes32 , bytes32 , uint8 _txResult, uint256 _maxRefundedGas) payable returns()
func (_ExampleAuthServerPaymaster *ExampleAuthServerPaymasterTransactor) PostTransaction(opts *bind.TransactOpts, _context []byte, _transaction Transaction, arg2 [32]byte, arg3 [32]byte, _txResult uint8, _maxRefundedGas *big.Int) (*types.Transaction, error) {
	return _ExampleAuthServerPaymaster.contract.Transact(opts, "postTransaction", _context, _transaction, arg2, arg3, _txResult, _maxRefundedGas)
}

// PostTransaction is a paid mutator transaction binding the contract method 0x817b17f0.
//
// Solidity: function postTransaction(bytes _context, (uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256[4],bytes,bytes,bytes32[],bytes,bytes) _transaction, bytes32 , bytes32 , uint8 _txResult, uint256 _maxRefundedGas) payable returns()
func (_ExampleAuthServerPaymaster *ExampleAuthServerPaymasterSession) PostTransaction(_context []byte, _transaction Transaction, arg2 [32]byte, arg3 [32]byte, _txResult uint8, _maxRefundedGas *big.Int) (*types.Transaction, error) {
	return _ExampleAuthServerPaymaster.Contract.PostTransaction(&_ExampleAuthServerPaymaster.TransactOpts, _context, _transaction, arg2, arg3, _txResult, _maxRefundedGas)
}

// PostTransaction is a paid mutator transaction binding the contract method 0x817b17f0.
//
// Solidity: function postTransaction(bytes _context, (uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256[4],bytes,bytes,bytes32[],bytes,bytes) _transaction, bytes32 , bytes32 , uint8 _txResult, uint256 _maxRefundedGas) payable returns()
func (_ExampleAuthServerPaymaster *ExampleAuthServerPaymasterTransactorSession) PostTransaction(_context []byte, _transaction Transaction, arg2 [32]byte, arg3 [32]byte, _txResult uint8, _maxRefundedGas *big.Int) (*types.Transaction, error) {
	return _ExampleAuthServerPaymaster.Contract.PostTransaction(&_ExampleAuthServerPaymaster.TransactOpts, _context, _transaction, arg2, arg3, _txResult, _maxRefundedGas)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_ExampleAuthServerPaymaster *ExampleAuthServerPaymasterTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ExampleAuthServerPaymaster.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_ExampleAuthServerPaymaster *ExampleAuthServerPaymasterSession) RenounceOwnership() (*types.Transaction, error) {
	return _ExampleAuthServerPaymaster.Contract.RenounceOwnership(&_ExampleAuthServerPaymaster.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_ExampleAuthServerPaymaster *ExampleAuthServerPaymasterTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _ExampleAuthServerPaymaster.Contract.RenounceOwnership(&_ExampleAuthServerPaymaster.TransactOpts)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_ExampleAuthServerPaymaster *ExampleAuthServerPaymasterTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _ExampleAuthServerPaymaster.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_ExampleAuthServerPaymaster *ExampleAuthServerPaymasterSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _ExampleAuthServerPaymaster.Contract.TransferOwnership(&_ExampleAuthServerPaymaster.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_ExampleAuthServerPaymaster *ExampleAuthServerPaymasterTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _ExampleAuthServerPaymaster.Contract.TransferOwnership(&_ExampleAuthServerPaymaster.TransactOpts, newOwner)
}

// ValidateAndPayForPaymasterTransaction is a paid mutator transaction binding the contract method 0x038a24bc.
//
// Solidity: function validateAndPayForPaymasterTransaction(bytes32 , bytes32 , (uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256[4],bytes,bytes,bytes32[],bytes,bytes) _transaction) payable returns(bytes4 magic, bytes)
func (_ExampleAuthServerPaymaster *ExampleAuthServerPaymasterTransactor) ValidateAndPayForPaymasterTransaction(opts *bind.TransactOpts, arg0 [32]byte, arg1 [32]byte, _transaction Transaction) (*types.Transaction, error) {
	return _ExampleAuthServerPaymaster.contract.Transact(opts, "validateAndPayForPaymasterTransaction", arg0, arg1, _transaction)
}

// ValidateAndPayForPaymasterTransaction is a paid mutator transaction binding the contract method 0x038a24bc.
//
// Solidity: function validateAndPayForPaymasterTransaction(bytes32 , bytes32 , (uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256[4],bytes,bytes,bytes32[],bytes,bytes) _transaction) payable returns(bytes4 magic, bytes)
func (_ExampleAuthServerPaymaster *ExampleAuthServerPaymasterSession) ValidateAndPayForPaymasterTransaction(arg0 [32]byte, arg1 [32]byte, _transaction Transaction) (*types.Transaction, error) {
	return _ExampleAuthServerPaymaster.Contract.ValidateAndPayForPaymasterTransaction(&_ExampleAuthServerPaymaster.TransactOpts, arg0, arg1, _transaction)
}

// ValidateAndPayForPaymasterTransaction is a paid mutator transaction binding the contract method 0x038a24bc.
//
// Solidity: function validateAndPayForPaymasterTransaction(bytes32 , bytes32 , (uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256[4],bytes,bytes,bytes32[],bytes,bytes) _transaction) payable returns(bytes4 magic, bytes)
func (_ExampleAuthServerPaymaster *ExampleAuthServerPaymasterTransactorSession) ValidateAndPayForPaymasterTransaction(arg0 [32]byte, arg1 [32]byte, _transaction Transaction) (*types.Transaction, error) {
	return _ExampleAuthServerPaymaster.Contract.ValidateAndPayForPaymasterTransaction(&_ExampleAuthServerPaymaster.TransactOpts, arg0, arg1, _transaction)
}

// Withdraw is a paid mutator transaction binding the contract method 0x51cff8d9.
//
// Solidity: function withdraw(address _to) returns()
func (_ExampleAuthServerPaymaster *ExampleAuthServerPaymasterTransactor) Withdraw(opts *bind.TransactOpts, _to common.Address) (*types.Transaction, error) {
	return _ExampleAuthServerPaymaster.contract.Transact(opts, "withdraw", _to)
}

// Withdraw is a paid mutator transaction binding the contract method 0x51cff8d9.
//
// Solidity: function withdraw(address _to) returns()
func (_ExampleAuthServerPaymaster *ExampleAuthServerPaymasterSession) Withdraw(_to common.Address) (*types.Transaction, error) {
	return _ExampleAuthServerPaymaster.Contract.Withdraw(&_ExampleAuthServerPaymaster.TransactOpts, _to)
}

// Withdraw is a paid mutator transaction binding the contract method 0x51cff8d9.
//
// Solidity: function withdraw(address _to) returns()
func (_ExampleAuthServerPaymaster *ExampleAuthServerPaymasterTransactorSession) Withdraw(_to common.Address) (*types.Transaction, error) {
	return _ExampleAuthServerPaymaster.Contract.Withdraw(&_ExampleAuthServerPaymaster.TransactOpts, _to)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_ExampleAuthServerPaymaster *ExampleAuthServerPaymasterTransactor) Receive(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ExampleAuthServerPaymaster.contract.RawTransact(opts, nil) // calldata is disallowed for receive function
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_ExampleAuthServerPaymaster *ExampleAuthServerPaymasterSession) Receive() (*types.Transaction, error) {
	return _ExampleAuthServerPaymaster.Contract.Receive(&_ExampleAuthServerPaymaster.TransactOpts)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_ExampleAuthServerPaymaster *ExampleAuthServerPaymasterTransactorSession) Receive() (*types.Transaction, error) {
	return _ExampleAuthServerPaymaster.Contract.Receive(&_ExampleAuthServerPaymaster.TransactOpts)
}

// ExampleAuthServerPaymasterOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the ExampleAuthServerPaymaster contract.
type ExampleAuthServerPaymasterOwnershipTransferredIterator struct {
	Event *ExampleAuthServerPaymasterOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *ExampleAuthServerPaymasterOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ExampleAuthServerPaymasterOwnershipTransferred)
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
		it.Event = new(ExampleAuthServerPaymasterOwnershipTransferred)
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
func (it *ExampleAuthServerPaymasterOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ExampleAuthServerPaymasterOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ExampleAuthServerPaymasterOwnershipTransferred represents a OwnershipTransferred event raised by the ExampleAuthServerPaymaster contract.
type ExampleAuthServerPaymasterOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_ExampleAuthServerPaymaster *ExampleAuthServerPaymasterFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*ExampleAuthServerPaymasterOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _ExampleAuthServerPaymaster.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &ExampleAuthServerPaymasterOwnershipTransferredIterator{contract: _ExampleAuthServerPaymaster.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_ExampleAuthServerPaymaster *ExampleAuthServerPaymasterFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *ExampleAuthServerPaymasterOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _ExampleAuthServerPaymaster.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ExampleAuthServerPaymasterOwnershipTransferred)
				if err := _ExampleAuthServerPaymaster.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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

// ParseOwnershipTransferred is a log parse operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_ExampleAuthServerPaymaster *ExampleAuthServerPaymasterFilterer) ParseOwnershipTransferred(log types.Log) (*ExampleAuthServerPaymasterOwnershipTransferred, error) {
	event := new(ExampleAuthServerPaymasterOwnershipTransferred)
	if err := _ExampleAuthServerPaymaster.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
