// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package guardianrecoveryvalidator

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

// IGuardianRecoveryValidatorGuardian is an auto generated low-level Go binding around an user-defined struct.
type IGuardianRecoveryValidatorGuardian struct {
	Addr    common.Address
	IsReady bool
	AddedAt uint64
}

// IGuardianRecoveryValidatorRecoveryRequest is an auto generated low-level Go binding around an user-defined struct.
type IGuardianRecoveryValidatorRecoveryRequest struct {
	HashedCredentialId [32]byte
	RawPublicKey       [2][32]byte
	Timestamp          *big.Int
}

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

// GuardianRecoveryValidatorMetaData contains all meta data concerning the GuardianRecoveryValidator contract.
var GuardianRecoveryValidatorMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"ADDRESS_CAST_OVERFLOW\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"guardian\",\"type\":\"address\"}],\"name\":\"AccountAlreadyGuardedByGuardian\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"guardian\",\"type\":\"address\"}],\"name\":\"AccountNotGuardedByAddress\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"AccountRecoveryInProgress\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"GuardianCannotBeSelf\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"guardian\",\"type\":\"address\"}],\"name\":\"GuardianNotFound\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"guardian\",\"type\":\"address\"}],\"name\":\"GuardianNotProposed\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidAccountToGuardAddress\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidAccountToRecoverAddress\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidGuardianAddress\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidWebAuthValidatorAddress\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"chainId\",\"type\":\"uint256\"}],\"name\":\"NO_TIMESTAMP_ASSERTER\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NonFunctionCallTransaction\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"hashedOriginDomain\",\"type\":\"bytes32\"}],\"name\":\"UnknownHashedOriginDomain\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"WebAuthValidatorNotEnabled\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"hashedOriginDomain\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"guardian\",\"type\":\"address\"}],\"name\":\"GuardianAdded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"hashedOriginDomain\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"guardian\",\"type\":\"address\"}],\"name\":\"GuardianProposed\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"hashedOriginDomain\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"guardian\",\"type\":\"address\"}],\"name\":\"GuardianRemoved\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"hashedOriginDomain\",\"type\":\"bytes32\"}],\"name\":\"HashedOriginDomainDisabledForAccount\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"hashedOriginDomain\",\"type\":\"bytes32\"}],\"name\":\"HashedOriginDomainEnabledForAccount\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint8\",\"name\":\"version\",\"type\":\"uint8\"}],\"name\":\"Initialized\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"hashedOriginDomain\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"hashedCredentialId\",\"type\":\"bytes32\"}],\"name\":\"RecoveryDiscarded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"hashedOriginDomain\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"hashedCredentialId\",\"type\":\"bytes32\"}],\"name\":\"RecoveryFinished\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"hashedOriginDomain\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"hashedCredentialId\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"guardian\",\"type\":\"address\"}],\"name\":\"RecoveryInitiated\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"REQUEST_DELAY_TIME\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"REQUEST_VALIDITY_TIME\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"hashedOriginDomain\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"guardian\",\"type\":\"address\"}],\"name\":\"accountGuardianData\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"addr\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"isReady\",\"type\":\"bool\"},{\"internalType\":\"uint64\",\"name\":\"addedAt\",\"type\":\"uint64\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"hashedOriginDomain\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"accountToGuard\",\"type\":\"address\"}],\"name\":\"addGuardian\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"hashedOriginDomain\",\"type\":\"bytes32\"}],\"name\":\"discardRecovery\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"hashedOriginDomain\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"getPendingRecoveryData\",\"outputs\":[{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"hashedCredentialId\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32[2]\",\"name\":\"rawPublicKey\",\"type\":\"bytes32[2]\"},{\"internalType\":\"uint256\",\"name\":\"timestamp\",\"type\":\"uint256\"}],\"internalType\":\"structIGuardianRecoveryValidator.RecoveryRequest\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"hashedOriginDomain\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"guardian\",\"type\":\"address\"}],\"name\":\"guardianOf\",\"outputs\":[{\"internalType\":\"address[]\",\"name\":\"\",\"type\":\"address[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"hashedOriginDomain\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"addr\",\"type\":\"address\"}],\"name\":\"guardiansFor\",\"outputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"addr\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"isReady\",\"type\":\"bool\"},{\"internalType\":\"uint64\",\"name\":\"addedAt\",\"type\":\"uint64\"}],\"internalType\":\"structIGuardianRecoveryValidator.Guardian[]\",\"name\":\"\",\"type\":\"tuple[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"accountToRecover\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"hashedCredentialId\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32[2]\",\"name\":\"rawPublicKey\",\"type\":\"bytes32[2]\"},{\"internalType\":\"bytes32\",\"name\":\"hashedOriginDomain\",\"type\":\"bytes32\"}],\"name\":\"initRecovery\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractWebAuthValidator\",\"name\":\"_webAuthValidator\",\"type\":\"address\"}],\"name\":\"initialize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"name\":\"onInstall\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"name\":\"onUninstall\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"hashedOriginDomain\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"newGuardian\",\"type\":\"address\"}],\"name\":\"proposeGuardian\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"hashedOriginDomain\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"guardianToRemove\",\"type\":\"address\"}],\"name\":\"removeGuardian\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes4\",\"name\":\"interfaceId\",\"type\":\"bytes4\"}],\"name\":\"supportsInterface\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"},{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"name\":\"validateSignature\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"txType\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"from\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"to\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"gasLimit\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"gasPerPubdataByteLimit\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"maxFeePerGas\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"maxPriorityFeePerGas\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"paymaster\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"nonce\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"},{\"internalType\":\"uint256[4]\",\"name\":\"reserved\",\"type\":\"uint256[4]\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"signature\",\"type\":\"bytes\"},{\"internalType\":\"bytes32[]\",\"name\":\"factoryDeps\",\"type\":\"bytes32[]\"},{\"internalType\":\"bytes\",\"name\":\"paymasterInput\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"reservedDynamic\",\"type\":\"bytes\"}],\"internalType\":\"structTransaction\",\"name\":\"transaction\",\"type\":\"tuple\"}],\"name\":\"validateTransaction\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"webAuthValidator\",\"outputs\":[{\"internalType\":\"contractWebAuthValidator\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
}

// GuardianRecoveryValidatorABI is the input ABI used to generate the binding from.
// Deprecated: Use GuardianRecoveryValidatorMetaData.ABI instead.
var GuardianRecoveryValidatorABI = GuardianRecoveryValidatorMetaData.ABI

// GuardianRecoveryValidator is an auto generated Go binding around an Ethereum contract.
type GuardianRecoveryValidator struct {
	GuardianRecoveryValidatorCaller     // Read-only binding to the contract
	GuardianRecoveryValidatorTransactor // Write-only binding to the contract
	GuardianRecoveryValidatorFilterer   // Log filterer for contract events
}

// GuardianRecoveryValidatorCaller is an auto generated read-only Go binding around an Ethereum contract.
type GuardianRecoveryValidatorCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// GuardianRecoveryValidatorTransactor is an auto generated write-only Go binding around an Ethereum contract.
type GuardianRecoveryValidatorTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// GuardianRecoveryValidatorFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type GuardianRecoveryValidatorFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// GuardianRecoveryValidatorSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type GuardianRecoveryValidatorSession struct {
	Contract     *GuardianRecoveryValidator // Generic contract binding to set the session for
	CallOpts     bind.CallOpts              // Call options to use throughout this session
	TransactOpts bind.TransactOpts          // Transaction auth options to use throughout this session
}

// GuardianRecoveryValidatorCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type GuardianRecoveryValidatorCallerSession struct {
	Contract *GuardianRecoveryValidatorCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts                    // Call options to use throughout this session
}

// GuardianRecoveryValidatorTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type GuardianRecoveryValidatorTransactorSession struct {
	Contract     *GuardianRecoveryValidatorTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts                    // Transaction auth options to use throughout this session
}

// GuardianRecoveryValidatorRaw is an auto generated low-level Go binding around an Ethereum contract.
type GuardianRecoveryValidatorRaw struct {
	Contract *GuardianRecoveryValidator // Generic contract binding to access the raw methods on
}

// GuardianRecoveryValidatorCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type GuardianRecoveryValidatorCallerRaw struct {
	Contract *GuardianRecoveryValidatorCaller // Generic read-only contract binding to access the raw methods on
}

// GuardianRecoveryValidatorTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type GuardianRecoveryValidatorTransactorRaw struct {
	Contract *GuardianRecoveryValidatorTransactor // Generic write-only contract binding to access the raw methods on
}

// NewGuardianRecoveryValidator creates a new instance of GuardianRecoveryValidator, bound to a specific deployed contract.
func NewGuardianRecoveryValidator(address common.Address, backend bind.ContractBackend) (*GuardianRecoveryValidator, error) {
	contract, err := bindGuardianRecoveryValidator(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &GuardianRecoveryValidator{GuardianRecoveryValidatorCaller: GuardianRecoveryValidatorCaller{contract: contract}, GuardianRecoveryValidatorTransactor: GuardianRecoveryValidatorTransactor{contract: contract}, GuardianRecoveryValidatorFilterer: GuardianRecoveryValidatorFilterer{contract: contract}}, nil
}

// NewGuardianRecoveryValidatorCaller creates a new read-only instance of GuardianRecoveryValidator, bound to a specific deployed contract.
func NewGuardianRecoveryValidatorCaller(address common.Address, caller bind.ContractCaller) (*GuardianRecoveryValidatorCaller, error) {
	contract, err := bindGuardianRecoveryValidator(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &GuardianRecoveryValidatorCaller{contract: contract}, nil
}

// NewGuardianRecoveryValidatorTransactor creates a new write-only instance of GuardianRecoveryValidator, bound to a specific deployed contract.
func NewGuardianRecoveryValidatorTransactor(address common.Address, transactor bind.ContractTransactor) (*GuardianRecoveryValidatorTransactor, error) {
	contract, err := bindGuardianRecoveryValidator(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &GuardianRecoveryValidatorTransactor{contract: contract}, nil
}

// NewGuardianRecoveryValidatorFilterer creates a new log filterer instance of GuardianRecoveryValidator, bound to a specific deployed contract.
func NewGuardianRecoveryValidatorFilterer(address common.Address, filterer bind.ContractFilterer) (*GuardianRecoveryValidatorFilterer, error) {
	contract, err := bindGuardianRecoveryValidator(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &GuardianRecoveryValidatorFilterer{contract: contract}, nil
}

// bindGuardianRecoveryValidator binds a generic wrapper to an already deployed contract.
func bindGuardianRecoveryValidator(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := GuardianRecoveryValidatorMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_GuardianRecoveryValidator *GuardianRecoveryValidatorRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _GuardianRecoveryValidator.Contract.GuardianRecoveryValidatorCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_GuardianRecoveryValidator *GuardianRecoveryValidatorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _GuardianRecoveryValidator.Contract.GuardianRecoveryValidatorTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_GuardianRecoveryValidator *GuardianRecoveryValidatorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _GuardianRecoveryValidator.Contract.GuardianRecoveryValidatorTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_GuardianRecoveryValidator *GuardianRecoveryValidatorCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _GuardianRecoveryValidator.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_GuardianRecoveryValidator *GuardianRecoveryValidatorTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _GuardianRecoveryValidator.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_GuardianRecoveryValidator *GuardianRecoveryValidatorTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _GuardianRecoveryValidator.Contract.contract.Transact(opts, method, params...)
}

// REQUESTDELAYTIME is a free data retrieval call binding the contract method 0xfe8d85ab.
//
// Solidity: function REQUEST_DELAY_TIME() view returns(uint256)
func (_GuardianRecoveryValidator *GuardianRecoveryValidatorCaller) REQUESTDELAYTIME(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _GuardianRecoveryValidator.contract.Call(opts, &out, "REQUEST_DELAY_TIME")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// REQUESTDELAYTIME is a free data retrieval call binding the contract method 0xfe8d85ab.
//
// Solidity: function REQUEST_DELAY_TIME() view returns(uint256)
func (_GuardianRecoveryValidator *GuardianRecoveryValidatorSession) REQUESTDELAYTIME() (*big.Int, error) {
	return _GuardianRecoveryValidator.Contract.REQUESTDELAYTIME(&_GuardianRecoveryValidator.CallOpts)
}

// REQUESTDELAYTIME is a free data retrieval call binding the contract method 0xfe8d85ab.
//
// Solidity: function REQUEST_DELAY_TIME() view returns(uint256)
func (_GuardianRecoveryValidator *GuardianRecoveryValidatorCallerSession) REQUESTDELAYTIME() (*big.Int, error) {
	return _GuardianRecoveryValidator.Contract.REQUESTDELAYTIME(&_GuardianRecoveryValidator.CallOpts)
}

// REQUESTVALIDITYTIME is a free data retrieval call binding the contract method 0x32c768a9.
//
// Solidity: function REQUEST_VALIDITY_TIME() view returns(uint256)
func (_GuardianRecoveryValidator *GuardianRecoveryValidatorCaller) REQUESTVALIDITYTIME(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _GuardianRecoveryValidator.contract.Call(opts, &out, "REQUEST_VALIDITY_TIME")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// REQUESTVALIDITYTIME is a free data retrieval call binding the contract method 0x32c768a9.
//
// Solidity: function REQUEST_VALIDITY_TIME() view returns(uint256)
func (_GuardianRecoveryValidator *GuardianRecoveryValidatorSession) REQUESTVALIDITYTIME() (*big.Int, error) {
	return _GuardianRecoveryValidator.Contract.REQUESTVALIDITYTIME(&_GuardianRecoveryValidator.CallOpts)
}

// REQUESTVALIDITYTIME is a free data retrieval call binding the contract method 0x32c768a9.
//
// Solidity: function REQUEST_VALIDITY_TIME() view returns(uint256)
func (_GuardianRecoveryValidator *GuardianRecoveryValidatorCallerSession) REQUESTVALIDITYTIME() (*big.Int, error) {
	return _GuardianRecoveryValidator.Contract.REQUESTVALIDITYTIME(&_GuardianRecoveryValidator.CallOpts)
}

// AccountGuardianData is a free data retrieval call binding the contract method 0x9b99752a.
//
// Solidity: function accountGuardianData(bytes32 hashedOriginDomain, address account, address guardian) view returns(address addr, bool isReady, uint64 addedAt)
func (_GuardianRecoveryValidator *GuardianRecoveryValidatorCaller) AccountGuardianData(opts *bind.CallOpts, hashedOriginDomain [32]byte, account common.Address, guardian common.Address) (struct {
	Addr    common.Address
	IsReady bool
	AddedAt uint64
}, error) {
	var out []interface{}
	err := _GuardianRecoveryValidator.contract.Call(opts, &out, "accountGuardianData", hashedOriginDomain, account, guardian)

	outstruct := new(struct {
		Addr    common.Address
		IsReady bool
		AddedAt uint64
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Addr = *abi.ConvertType(out[0], new(common.Address)).(*common.Address)
	outstruct.IsReady = *abi.ConvertType(out[1], new(bool)).(*bool)
	outstruct.AddedAt = *abi.ConvertType(out[2], new(uint64)).(*uint64)

	return *outstruct, err

}

// AccountGuardianData is a free data retrieval call binding the contract method 0x9b99752a.
//
// Solidity: function accountGuardianData(bytes32 hashedOriginDomain, address account, address guardian) view returns(address addr, bool isReady, uint64 addedAt)
func (_GuardianRecoveryValidator *GuardianRecoveryValidatorSession) AccountGuardianData(hashedOriginDomain [32]byte, account common.Address, guardian common.Address) (struct {
	Addr    common.Address
	IsReady bool
	AddedAt uint64
}, error) {
	return _GuardianRecoveryValidator.Contract.AccountGuardianData(&_GuardianRecoveryValidator.CallOpts, hashedOriginDomain, account, guardian)
}

// AccountGuardianData is a free data retrieval call binding the contract method 0x9b99752a.
//
// Solidity: function accountGuardianData(bytes32 hashedOriginDomain, address account, address guardian) view returns(address addr, bool isReady, uint64 addedAt)
func (_GuardianRecoveryValidator *GuardianRecoveryValidatorCallerSession) AccountGuardianData(hashedOriginDomain [32]byte, account common.Address, guardian common.Address) (struct {
	Addr    common.Address
	IsReady bool
	AddedAt uint64
}, error) {
	return _GuardianRecoveryValidator.Contract.AccountGuardianData(&_GuardianRecoveryValidator.CallOpts, hashedOriginDomain, account, guardian)
}

// GetPendingRecoveryData is a free data retrieval call binding the contract method 0xab9d5fff.
//
// Solidity: function getPendingRecoveryData(bytes32 hashedOriginDomain, address account) view returns((bytes32,bytes32[2],uint256))
func (_GuardianRecoveryValidator *GuardianRecoveryValidatorCaller) GetPendingRecoveryData(opts *bind.CallOpts, hashedOriginDomain [32]byte, account common.Address) (IGuardianRecoveryValidatorRecoveryRequest, error) {
	var out []interface{}
	err := _GuardianRecoveryValidator.contract.Call(opts, &out, "getPendingRecoveryData", hashedOriginDomain, account)

	if err != nil {
		return *new(IGuardianRecoveryValidatorRecoveryRequest), err
	}

	out0 := *abi.ConvertType(out[0], new(IGuardianRecoveryValidatorRecoveryRequest)).(*IGuardianRecoveryValidatorRecoveryRequest)

	return out0, err

}

// GetPendingRecoveryData is a free data retrieval call binding the contract method 0xab9d5fff.
//
// Solidity: function getPendingRecoveryData(bytes32 hashedOriginDomain, address account) view returns((bytes32,bytes32[2],uint256))
func (_GuardianRecoveryValidator *GuardianRecoveryValidatorSession) GetPendingRecoveryData(hashedOriginDomain [32]byte, account common.Address) (IGuardianRecoveryValidatorRecoveryRequest, error) {
	return _GuardianRecoveryValidator.Contract.GetPendingRecoveryData(&_GuardianRecoveryValidator.CallOpts, hashedOriginDomain, account)
}

// GetPendingRecoveryData is a free data retrieval call binding the contract method 0xab9d5fff.
//
// Solidity: function getPendingRecoveryData(bytes32 hashedOriginDomain, address account) view returns((bytes32,bytes32[2],uint256))
func (_GuardianRecoveryValidator *GuardianRecoveryValidatorCallerSession) GetPendingRecoveryData(hashedOriginDomain [32]byte, account common.Address) (IGuardianRecoveryValidatorRecoveryRequest, error) {
	return _GuardianRecoveryValidator.Contract.GetPendingRecoveryData(&_GuardianRecoveryValidator.CallOpts, hashedOriginDomain, account)
}

// GuardianOf is a free data retrieval call binding the contract method 0xcb339850.
//
// Solidity: function guardianOf(bytes32 hashedOriginDomain, address guardian) view returns(address[])
func (_GuardianRecoveryValidator *GuardianRecoveryValidatorCaller) GuardianOf(opts *bind.CallOpts, hashedOriginDomain [32]byte, guardian common.Address) ([]common.Address, error) {
	var out []interface{}
	err := _GuardianRecoveryValidator.contract.Call(opts, &out, "guardianOf", hashedOriginDomain, guardian)

	if err != nil {
		return *new([]common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)

	return out0, err

}

// GuardianOf is a free data retrieval call binding the contract method 0xcb339850.
//
// Solidity: function guardianOf(bytes32 hashedOriginDomain, address guardian) view returns(address[])
func (_GuardianRecoveryValidator *GuardianRecoveryValidatorSession) GuardianOf(hashedOriginDomain [32]byte, guardian common.Address) ([]common.Address, error) {
	return _GuardianRecoveryValidator.Contract.GuardianOf(&_GuardianRecoveryValidator.CallOpts, hashedOriginDomain, guardian)
}

// GuardianOf is a free data retrieval call binding the contract method 0xcb339850.
//
// Solidity: function guardianOf(bytes32 hashedOriginDomain, address guardian) view returns(address[])
func (_GuardianRecoveryValidator *GuardianRecoveryValidatorCallerSession) GuardianOf(hashedOriginDomain [32]byte, guardian common.Address) ([]common.Address, error) {
	return _GuardianRecoveryValidator.Contract.GuardianOf(&_GuardianRecoveryValidator.CallOpts, hashedOriginDomain, guardian)
}

// GuardiansFor is a free data retrieval call binding the contract method 0x511f723d.
//
// Solidity: function guardiansFor(bytes32 hashedOriginDomain, address addr) view returns((address,bool,uint64)[])
func (_GuardianRecoveryValidator *GuardianRecoveryValidatorCaller) GuardiansFor(opts *bind.CallOpts, hashedOriginDomain [32]byte, addr common.Address) ([]IGuardianRecoveryValidatorGuardian, error) {
	var out []interface{}
	err := _GuardianRecoveryValidator.contract.Call(opts, &out, "guardiansFor", hashedOriginDomain, addr)

	if err != nil {
		return *new([]IGuardianRecoveryValidatorGuardian), err
	}

	out0 := *abi.ConvertType(out[0], new([]IGuardianRecoveryValidatorGuardian)).(*[]IGuardianRecoveryValidatorGuardian)

	return out0, err

}

// GuardiansFor is a free data retrieval call binding the contract method 0x511f723d.
//
// Solidity: function guardiansFor(bytes32 hashedOriginDomain, address addr) view returns((address,bool,uint64)[])
func (_GuardianRecoveryValidator *GuardianRecoveryValidatorSession) GuardiansFor(hashedOriginDomain [32]byte, addr common.Address) ([]IGuardianRecoveryValidatorGuardian, error) {
	return _GuardianRecoveryValidator.Contract.GuardiansFor(&_GuardianRecoveryValidator.CallOpts, hashedOriginDomain, addr)
}

// GuardiansFor is a free data retrieval call binding the contract method 0x511f723d.
//
// Solidity: function guardiansFor(bytes32 hashedOriginDomain, address addr) view returns((address,bool,uint64)[])
func (_GuardianRecoveryValidator *GuardianRecoveryValidatorCallerSession) GuardiansFor(hashedOriginDomain [32]byte, addr common.Address) ([]IGuardianRecoveryValidatorGuardian, error) {
	return _GuardianRecoveryValidator.Contract.GuardiansFor(&_GuardianRecoveryValidator.CallOpts, hashedOriginDomain, addr)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) pure returns(bool)
func (_GuardianRecoveryValidator *GuardianRecoveryValidatorCaller) SupportsInterface(opts *bind.CallOpts, interfaceId [4]byte) (bool, error) {
	var out []interface{}
	err := _GuardianRecoveryValidator.contract.Call(opts, &out, "supportsInterface", interfaceId)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) pure returns(bool)
func (_GuardianRecoveryValidator *GuardianRecoveryValidatorSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _GuardianRecoveryValidator.Contract.SupportsInterface(&_GuardianRecoveryValidator.CallOpts, interfaceId)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) pure returns(bool)
func (_GuardianRecoveryValidator *GuardianRecoveryValidatorCallerSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _GuardianRecoveryValidator.Contract.SupportsInterface(&_GuardianRecoveryValidator.CallOpts, interfaceId)
}

// ValidateSignature is a free data retrieval call binding the contract method 0x333daf92.
//
// Solidity: function validateSignature(bytes32 , bytes ) pure returns(bool)
func (_GuardianRecoveryValidator *GuardianRecoveryValidatorCaller) ValidateSignature(opts *bind.CallOpts, arg0 [32]byte, arg1 []byte) (bool, error) {
	var out []interface{}
	err := _GuardianRecoveryValidator.contract.Call(opts, &out, "validateSignature", arg0, arg1)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// ValidateSignature is a free data retrieval call binding the contract method 0x333daf92.
//
// Solidity: function validateSignature(bytes32 , bytes ) pure returns(bool)
func (_GuardianRecoveryValidator *GuardianRecoveryValidatorSession) ValidateSignature(arg0 [32]byte, arg1 []byte) (bool, error) {
	return _GuardianRecoveryValidator.Contract.ValidateSignature(&_GuardianRecoveryValidator.CallOpts, arg0, arg1)
}

// ValidateSignature is a free data retrieval call binding the contract method 0x333daf92.
//
// Solidity: function validateSignature(bytes32 , bytes ) pure returns(bool)
func (_GuardianRecoveryValidator *GuardianRecoveryValidatorCallerSession) ValidateSignature(arg0 [32]byte, arg1 []byte) (bool, error) {
	return _GuardianRecoveryValidator.Contract.ValidateSignature(&_GuardianRecoveryValidator.CallOpts, arg0, arg1)
}

// WebAuthValidator is a free data retrieval call binding the contract method 0x25290a37.
//
// Solidity: function webAuthValidator() view returns(address)
func (_GuardianRecoveryValidator *GuardianRecoveryValidatorCaller) WebAuthValidator(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _GuardianRecoveryValidator.contract.Call(opts, &out, "webAuthValidator")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// WebAuthValidator is a free data retrieval call binding the contract method 0x25290a37.
//
// Solidity: function webAuthValidator() view returns(address)
func (_GuardianRecoveryValidator *GuardianRecoveryValidatorSession) WebAuthValidator() (common.Address, error) {
	return _GuardianRecoveryValidator.Contract.WebAuthValidator(&_GuardianRecoveryValidator.CallOpts)
}

// WebAuthValidator is a free data retrieval call binding the contract method 0x25290a37.
//
// Solidity: function webAuthValidator() view returns(address)
func (_GuardianRecoveryValidator *GuardianRecoveryValidatorCallerSession) WebAuthValidator() (common.Address, error) {
	return _GuardianRecoveryValidator.Contract.WebAuthValidator(&_GuardianRecoveryValidator.CallOpts)
}

// AddGuardian is a paid mutator transaction binding the contract method 0x85431596.
//
// Solidity: function addGuardian(bytes32 hashedOriginDomain, address accountToGuard) returns(bool)
func (_GuardianRecoveryValidator *GuardianRecoveryValidatorTransactor) AddGuardian(opts *bind.TransactOpts, hashedOriginDomain [32]byte, accountToGuard common.Address) (*types.Transaction, error) {
	return _GuardianRecoveryValidator.contract.Transact(opts, "addGuardian", hashedOriginDomain, accountToGuard)
}

// AddGuardian is a paid mutator transaction binding the contract method 0x85431596.
//
// Solidity: function addGuardian(bytes32 hashedOriginDomain, address accountToGuard) returns(bool)
func (_GuardianRecoveryValidator *GuardianRecoveryValidatorSession) AddGuardian(hashedOriginDomain [32]byte, accountToGuard common.Address) (*types.Transaction, error) {
	return _GuardianRecoveryValidator.Contract.AddGuardian(&_GuardianRecoveryValidator.TransactOpts, hashedOriginDomain, accountToGuard)
}

// AddGuardian is a paid mutator transaction binding the contract method 0x85431596.
//
// Solidity: function addGuardian(bytes32 hashedOriginDomain, address accountToGuard) returns(bool)
func (_GuardianRecoveryValidator *GuardianRecoveryValidatorTransactorSession) AddGuardian(hashedOriginDomain [32]byte, accountToGuard common.Address) (*types.Transaction, error) {
	return _GuardianRecoveryValidator.Contract.AddGuardian(&_GuardianRecoveryValidator.TransactOpts, hashedOriginDomain, accountToGuard)
}

// DiscardRecovery is a paid mutator transaction binding the contract method 0x51bd1f69.
//
// Solidity: function discardRecovery(bytes32 hashedOriginDomain) returns()
func (_GuardianRecoveryValidator *GuardianRecoveryValidatorTransactor) DiscardRecovery(opts *bind.TransactOpts, hashedOriginDomain [32]byte) (*types.Transaction, error) {
	return _GuardianRecoveryValidator.contract.Transact(opts, "discardRecovery", hashedOriginDomain)
}

// DiscardRecovery is a paid mutator transaction binding the contract method 0x51bd1f69.
//
// Solidity: function discardRecovery(bytes32 hashedOriginDomain) returns()
func (_GuardianRecoveryValidator *GuardianRecoveryValidatorSession) DiscardRecovery(hashedOriginDomain [32]byte) (*types.Transaction, error) {
	return _GuardianRecoveryValidator.Contract.DiscardRecovery(&_GuardianRecoveryValidator.TransactOpts, hashedOriginDomain)
}

// DiscardRecovery is a paid mutator transaction binding the contract method 0x51bd1f69.
//
// Solidity: function discardRecovery(bytes32 hashedOriginDomain) returns()
func (_GuardianRecoveryValidator *GuardianRecoveryValidatorTransactorSession) DiscardRecovery(hashedOriginDomain [32]byte) (*types.Transaction, error) {
	return _GuardianRecoveryValidator.Contract.DiscardRecovery(&_GuardianRecoveryValidator.TransactOpts, hashedOriginDomain)
}

// InitRecovery is a paid mutator transaction binding the contract method 0x236ad679.
//
// Solidity: function initRecovery(address accountToRecover, bytes32 hashedCredentialId, bytes32[2] rawPublicKey, bytes32 hashedOriginDomain) returns()
func (_GuardianRecoveryValidator *GuardianRecoveryValidatorTransactor) InitRecovery(opts *bind.TransactOpts, accountToRecover common.Address, hashedCredentialId [32]byte, rawPublicKey [2][32]byte, hashedOriginDomain [32]byte) (*types.Transaction, error) {
	return _GuardianRecoveryValidator.contract.Transact(opts, "initRecovery", accountToRecover, hashedCredentialId, rawPublicKey, hashedOriginDomain)
}

// InitRecovery is a paid mutator transaction binding the contract method 0x236ad679.
//
// Solidity: function initRecovery(address accountToRecover, bytes32 hashedCredentialId, bytes32[2] rawPublicKey, bytes32 hashedOriginDomain) returns()
func (_GuardianRecoveryValidator *GuardianRecoveryValidatorSession) InitRecovery(accountToRecover common.Address, hashedCredentialId [32]byte, rawPublicKey [2][32]byte, hashedOriginDomain [32]byte) (*types.Transaction, error) {
	return _GuardianRecoveryValidator.Contract.InitRecovery(&_GuardianRecoveryValidator.TransactOpts, accountToRecover, hashedCredentialId, rawPublicKey, hashedOriginDomain)
}

// InitRecovery is a paid mutator transaction binding the contract method 0x236ad679.
//
// Solidity: function initRecovery(address accountToRecover, bytes32 hashedCredentialId, bytes32[2] rawPublicKey, bytes32 hashedOriginDomain) returns()
func (_GuardianRecoveryValidator *GuardianRecoveryValidatorTransactorSession) InitRecovery(accountToRecover common.Address, hashedCredentialId [32]byte, rawPublicKey [2][32]byte, hashedOriginDomain [32]byte) (*types.Transaction, error) {
	return _GuardianRecoveryValidator.Contract.InitRecovery(&_GuardianRecoveryValidator.TransactOpts, accountToRecover, hashedCredentialId, rawPublicKey, hashedOriginDomain)
}

// Initialize is a paid mutator transaction binding the contract method 0xc4d66de8.
//
// Solidity: function initialize(address _webAuthValidator) returns()
func (_GuardianRecoveryValidator *GuardianRecoveryValidatorTransactor) Initialize(opts *bind.TransactOpts, _webAuthValidator common.Address) (*types.Transaction, error) {
	return _GuardianRecoveryValidator.contract.Transact(opts, "initialize", _webAuthValidator)
}

// Initialize is a paid mutator transaction binding the contract method 0xc4d66de8.
//
// Solidity: function initialize(address _webAuthValidator) returns()
func (_GuardianRecoveryValidator *GuardianRecoveryValidatorSession) Initialize(_webAuthValidator common.Address) (*types.Transaction, error) {
	return _GuardianRecoveryValidator.Contract.Initialize(&_GuardianRecoveryValidator.TransactOpts, _webAuthValidator)
}

// Initialize is a paid mutator transaction binding the contract method 0xc4d66de8.
//
// Solidity: function initialize(address _webAuthValidator) returns()
func (_GuardianRecoveryValidator *GuardianRecoveryValidatorTransactorSession) Initialize(_webAuthValidator common.Address) (*types.Transaction, error) {
	return _GuardianRecoveryValidator.Contract.Initialize(&_GuardianRecoveryValidator.TransactOpts, _webAuthValidator)
}

// OnInstall is a paid mutator transaction binding the contract method 0x6d61fe70.
//
// Solidity: function onInstall(bytes ) returns()
func (_GuardianRecoveryValidator *GuardianRecoveryValidatorTransactor) OnInstall(opts *bind.TransactOpts, arg0 []byte) (*types.Transaction, error) {
	return _GuardianRecoveryValidator.contract.Transact(opts, "onInstall", arg0)
}

// OnInstall is a paid mutator transaction binding the contract method 0x6d61fe70.
//
// Solidity: function onInstall(bytes ) returns()
func (_GuardianRecoveryValidator *GuardianRecoveryValidatorSession) OnInstall(arg0 []byte) (*types.Transaction, error) {
	return _GuardianRecoveryValidator.Contract.OnInstall(&_GuardianRecoveryValidator.TransactOpts, arg0)
}

// OnInstall is a paid mutator transaction binding the contract method 0x6d61fe70.
//
// Solidity: function onInstall(bytes ) returns()
func (_GuardianRecoveryValidator *GuardianRecoveryValidatorTransactorSession) OnInstall(arg0 []byte) (*types.Transaction, error) {
	return _GuardianRecoveryValidator.Contract.OnInstall(&_GuardianRecoveryValidator.TransactOpts, arg0)
}

// OnUninstall is a paid mutator transaction binding the contract method 0x8a91b0e3.
//
// Solidity: function onUninstall(bytes ) returns()
func (_GuardianRecoveryValidator *GuardianRecoveryValidatorTransactor) OnUninstall(opts *bind.TransactOpts, arg0 []byte) (*types.Transaction, error) {
	return _GuardianRecoveryValidator.contract.Transact(opts, "onUninstall", arg0)
}

// OnUninstall is a paid mutator transaction binding the contract method 0x8a91b0e3.
//
// Solidity: function onUninstall(bytes ) returns()
func (_GuardianRecoveryValidator *GuardianRecoveryValidatorSession) OnUninstall(arg0 []byte) (*types.Transaction, error) {
	return _GuardianRecoveryValidator.Contract.OnUninstall(&_GuardianRecoveryValidator.TransactOpts, arg0)
}

// OnUninstall is a paid mutator transaction binding the contract method 0x8a91b0e3.
//
// Solidity: function onUninstall(bytes ) returns()
func (_GuardianRecoveryValidator *GuardianRecoveryValidatorTransactorSession) OnUninstall(arg0 []byte) (*types.Transaction, error) {
	return _GuardianRecoveryValidator.Contract.OnUninstall(&_GuardianRecoveryValidator.TransactOpts, arg0)
}

// ProposeGuardian is a paid mutator transaction binding the contract method 0x9f5d29c9.
//
// Solidity: function proposeGuardian(bytes32 hashedOriginDomain, address newGuardian) returns()
func (_GuardianRecoveryValidator *GuardianRecoveryValidatorTransactor) ProposeGuardian(opts *bind.TransactOpts, hashedOriginDomain [32]byte, newGuardian common.Address) (*types.Transaction, error) {
	return _GuardianRecoveryValidator.contract.Transact(opts, "proposeGuardian", hashedOriginDomain, newGuardian)
}

// ProposeGuardian is a paid mutator transaction binding the contract method 0x9f5d29c9.
//
// Solidity: function proposeGuardian(bytes32 hashedOriginDomain, address newGuardian) returns()
func (_GuardianRecoveryValidator *GuardianRecoveryValidatorSession) ProposeGuardian(hashedOriginDomain [32]byte, newGuardian common.Address) (*types.Transaction, error) {
	return _GuardianRecoveryValidator.Contract.ProposeGuardian(&_GuardianRecoveryValidator.TransactOpts, hashedOriginDomain, newGuardian)
}

// ProposeGuardian is a paid mutator transaction binding the contract method 0x9f5d29c9.
//
// Solidity: function proposeGuardian(bytes32 hashedOriginDomain, address newGuardian) returns()
func (_GuardianRecoveryValidator *GuardianRecoveryValidatorTransactorSession) ProposeGuardian(hashedOriginDomain [32]byte, newGuardian common.Address) (*types.Transaction, error) {
	return _GuardianRecoveryValidator.Contract.ProposeGuardian(&_GuardianRecoveryValidator.TransactOpts, hashedOriginDomain, newGuardian)
}

// RemoveGuardian is a paid mutator transaction binding the contract method 0x3cab947d.
//
// Solidity: function removeGuardian(bytes32 hashedOriginDomain, address guardianToRemove) returns()
func (_GuardianRecoveryValidator *GuardianRecoveryValidatorTransactor) RemoveGuardian(opts *bind.TransactOpts, hashedOriginDomain [32]byte, guardianToRemove common.Address) (*types.Transaction, error) {
	return _GuardianRecoveryValidator.contract.Transact(opts, "removeGuardian", hashedOriginDomain, guardianToRemove)
}

// RemoveGuardian is a paid mutator transaction binding the contract method 0x3cab947d.
//
// Solidity: function removeGuardian(bytes32 hashedOriginDomain, address guardianToRemove) returns()
func (_GuardianRecoveryValidator *GuardianRecoveryValidatorSession) RemoveGuardian(hashedOriginDomain [32]byte, guardianToRemove common.Address) (*types.Transaction, error) {
	return _GuardianRecoveryValidator.Contract.RemoveGuardian(&_GuardianRecoveryValidator.TransactOpts, hashedOriginDomain, guardianToRemove)
}

// RemoveGuardian is a paid mutator transaction binding the contract method 0x3cab947d.
//
// Solidity: function removeGuardian(bytes32 hashedOriginDomain, address guardianToRemove) returns()
func (_GuardianRecoveryValidator *GuardianRecoveryValidatorTransactorSession) RemoveGuardian(hashedOriginDomain [32]byte, guardianToRemove common.Address) (*types.Transaction, error) {
	return _GuardianRecoveryValidator.Contract.RemoveGuardian(&_GuardianRecoveryValidator.TransactOpts, hashedOriginDomain, guardianToRemove)
}

// ValidateTransaction is a paid mutator transaction binding the contract method 0xbf1733f9.
//
// Solidity: function validateTransaction(bytes32 , (uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256[4],bytes,bytes,bytes32[],bytes,bytes) transaction) returns(bool)
func (_GuardianRecoveryValidator *GuardianRecoveryValidatorTransactor) ValidateTransaction(opts *bind.TransactOpts, arg0 [32]byte, transaction Transaction) (*types.Transaction, error) {
	return _GuardianRecoveryValidator.contract.Transact(opts, "validateTransaction", arg0, transaction)
}

// ValidateTransaction is a paid mutator transaction binding the contract method 0xbf1733f9.
//
// Solidity: function validateTransaction(bytes32 , (uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256[4],bytes,bytes,bytes32[],bytes,bytes) transaction) returns(bool)
func (_GuardianRecoveryValidator *GuardianRecoveryValidatorSession) ValidateTransaction(arg0 [32]byte, transaction Transaction) (*types.Transaction, error) {
	return _GuardianRecoveryValidator.Contract.ValidateTransaction(&_GuardianRecoveryValidator.TransactOpts, arg0, transaction)
}

// ValidateTransaction is a paid mutator transaction binding the contract method 0xbf1733f9.
//
// Solidity: function validateTransaction(bytes32 , (uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256[4],bytes,bytes,bytes32[],bytes,bytes) transaction) returns(bool)
func (_GuardianRecoveryValidator *GuardianRecoveryValidatorTransactorSession) ValidateTransaction(arg0 [32]byte, transaction Transaction) (*types.Transaction, error) {
	return _GuardianRecoveryValidator.Contract.ValidateTransaction(&_GuardianRecoveryValidator.TransactOpts, arg0, transaction)
}

// GuardianRecoveryValidatorGuardianAddedIterator is returned from FilterGuardianAdded and is used to iterate over the raw logs and unpacked data for GuardianAdded events raised by the GuardianRecoveryValidator contract.
type GuardianRecoveryValidatorGuardianAddedIterator struct {
	Event *GuardianRecoveryValidatorGuardianAdded // Event containing the contract specifics and raw log

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
func (it *GuardianRecoveryValidatorGuardianAddedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(GuardianRecoveryValidatorGuardianAdded)
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
		it.Event = new(GuardianRecoveryValidatorGuardianAdded)
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
func (it *GuardianRecoveryValidatorGuardianAddedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *GuardianRecoveryValidatorGuardianAddedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// GuardianRecoveryValidatorGuardianAdded represents a GuardianAdded event raised by the GuardianRecoveryValidator contract.
type GuardianRecoveryValidatorGuardianAdded struct {
	Account            common.Address
	HashedOriginDomain [32]byte
	Guardian           common.Address
	Raw                types.Log // Blockchain specific contextual infos
}

// FilterGuardianAdded is a free log retrieval operation binding the contract event 0x6aba2dccf4edc3319e24e776ffc5d7f8fa89456e0f15421abf316d28d427ed33.
//
// Solidity: event GuardianAdded(address indexed account, bytes32 indexed hashedOriginDomain, address indexed guardian)
func (_GuardianRecoveryValidator *GuardianRecoveryValidatorFilterer) FilterGuardianAdded(opts *bind.FilterOpts, account []common.Address, hashedOriginDomain [][32]byte, guardian []common.Address) (*GuardianRecoveryValidatorGuardianAddedIterator, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}
	var hashedOriginDomainRule []interface{}
	for _, hashedOriginDomainItem := range hashedOriginDomain {
		hashedOriginDomainRule = append(hashedOriginDomainRule, hashedOriginDomainItem)
	}
	var guardianRule []interface{}
	for _, guardianItem := range guardian {
		guardianRule = append(guardianRule, guardianItem)
	}

	logs, sub, err := _GuardianRecoveryValidator.contract.FilterLogs(opts, "GuardianAdded", accountRule, hashedOriginDomainRule, guardianRule)
	if err != nil {
		return nil, err
	}
	return &GuardianRecoveryValidatorGuardianAddedIterator{contract: _GuardianRecoveryValidator.contract, event: "GuardianAdded", logs: logs, sub: sub}, nil
}

// WatchGuardianAdded is a free log subscription operation binding the contract event 0x6aba2dccf4edc3319e24e776ffc5d7f8fa89456e0f15421abf316d28d427ed33.
//
// Solidity: event GuardianAdded(address indexed account, bytes32 indexed hashedOriginDomain, address indexed guardian)
func (_GuardianRecoveryValidator *GuardianRecoveryValidatorFilterer) WatchGuardianAdded(opts *bind.WatchOpts, sink chan<- *GuardianRecoveryValidatorGuardianAdded, account []common.Address, hashedOriginDomain [][32]byte, guardian []common.Address) (event.Subscription, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}
	var hashedOriginDomainRule []interface{}
	for _, hashedOriginDomainItem := range hashedOriginDomain {
		hashedOriginDomainRule = append(hashedOriginDomainRule, hashedOriginDomainItem)
	}
	var guardianRule []interface{}
	for _, guardianItem := range guardian {
		guardianRule = append(guardianRule, guardianItem)
	}

	logs, sub, err := _GuardianRecoveryValidator.contract.WatchLogs(opts, "GuardianAdded", accountRule, hashedOriginDomainRule, guardianRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(GuardianRecoveryValidatorGuardianAdded)
				if err := _GuardianRecoveryValidator.contract.UnpackLog(event, "GuardianAdded", log); err != nil {
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

// ParseGuardianAdded is a log parse operation binding the contract event 0x6aba2dccf4edc3319e24e776ffc5d7f8fa89456e0f15421abf316d28d427ed33.
//
// Solidity: event GuardianAdded(address indexed account, bytes32 indexed hashedOriginDomain, address indexed guardian)
func (_GuardianRecoveryValidator *GuardianRecoveryValidatorFilterer) ParseGuardianAdded(log types.Log) (*GuardianRecoveryValidatorGuardianAdded, error) {
	event := new(GuardianRecoveryValidatorGuardianAdded)
	if err := _GuardianRecoveryValidator.contract.UnpackLog(event, "GuardianAdded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// GuardianRecoveryValidatorGuardianProposedIterator is returned from FilterGuardianProposed and is used to iterate over the raw logs and unpacked data for GuardianProposed events raised by the GuardianRecoveryValidator contract.
type GuardianRecoveryValidatorGuardianProposedIterator struct {
	Event *GuardianRecoveryValidatorGuardianProposed // Event containing the contract specifics and raw log

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
func (it *GuardianRecoveryValidatorGuardianProposedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(GuardianRecoveryValidatorGuardianProposed)
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
		it.Event = new(GuardianRecoveryValidatorGuardianProposed)
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
func (it *GuardianRecoveryValidatorGuardianProposedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *GuardianRecoveryValidatorGuardianProposedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// GuardianRecoveryValidatorGuardianProposed represents a GuardianProposed event raised by the GuardianRecoveryValidator contract.
type GuardianRecoveryValidatorGuardianProposed struct {
	Account            common.Address
	HashedOriginDomain [32]byte
	Guardian           common.Address
	Raw                types.Log // Blockchain specific contextual infos
}

// FilterGuardianProposed is a free log retrieval operation binding the contract event 0x3deaf0877005efc7ea2c571de4f96ff33bf60c0b32e188f1fa32bab3da8b870e.
//
// Solidity: event GuardianProposed(address indexed account, bytes32 indexed hashedOriginDomain, address indexed guardian)
func (_GuardianRecoveryValidator *GuardianRecoveryValidatorFilterer) FilterGuardianProposed(opts *bind.FilterOpts, account []common.Address, hashedOriginDomain [][32]byte, guardian []common.Address) (*GuardianRecoveryValidatorGuardianProposedIterator, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}
	var hashedOriginDomainRule []interface{}
	for _, hashedOriginDomainItem := range hashedOriginDomain {
		hashedOriginDomainRule = append(hashedOriginDomainRule, hashedOriginDomainItem)
	}
	var guardianRule []interface{}
	for _, guardianItem := range guardian {
		guardianRule = append(guardianRule, guardianItem)
	}

	logs, sub, err := _GuardianRecoveryValidator.contract.FilterLogs(opts, "GuardianProposed", accountRule, hashedOriginDomainRule, guardianRule)
	if err != nil {
		return nil, err
	}
	return &GuardianRecoveryValidatorGuardianProposedIterator{contract: _GuardianRecoveryValidator.contract, event: "GuardianProposed", logs: logs, sub: sub}, nil
}

// WatchGuardianProposed is a free log subscription operation binding the contract event 0x3deaf0877005efc7ea2c571de4f96ff33bf60c0b32e188f1fa32bab3da8b870e.
//
// Solidity: event GuardianProposed(address indexed account, bytes32 indexed hashedOriginDomain, address indexed guardian)
func (_GuardianRecoveryValidator *GuardianRecoveryValidatorFilterer) WatchGuardianProposed(opts *bind.WatchOpts, sink chan<- *GuardianRecoveryValidatorGuardianProposed, account []common.Address, hashedOriginDomain [][32]byte, guardian []common.Address) (event.Subscription, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}
	var hashedOriginDomainRule []interface{}
	for _, hashedOriginDomainItem := range hashedOriginDomain {
		hashedOriginDomainRule = append(hashedOriginDomainRule, hashedOriginDomainItem)
	}
	var guardianRule []interface{}
	for _, guardianItem := range guardian {
		guardianRule = append(guardianRule, guardianItem)
	}

	logs, sub, err := _GuardianRecoveryValidator.contract.WatchLogs(opts, "GuardianProposed", accountRule, hashedOriginDomainRule, guardianRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(GuardianRecoveryValidatorGuardianProposed)
				if err := _GuardianRecoveryValidator.contract.UnpackLog(event, "GuardianProposed", log); err != nil {
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

// ParseGuardianProposed is a log parse operation binding the contract event 0x3deaf0877005efc7ea2c571de4f96ff33bf60c0b32e188f1fa32bab3da8b870e.
//
// Solidity: event GuardianProposed(address indexed account, bytes32 indexed hashedOriginDomain, address indexed guardian)
func (_GuardianRecoveryValidator *GuardianRecoveryValidatorFilterer) ParseGuardianProposed(log types.Log) (*GuardianRecoveryValidatorGuardianProposed, error) {
	event := new(GuardianRecoveryValidatorGuardianProposed)
	if err := _GuardianRecoveryValidator.contract.UnpackLog(event, "GuardianProposed", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// GuardianRecoveryValidatorGuardianRemovedIterator is returned from FilterGuardianRemoved and is used to iterate over the raw logs and unpacked data for GuardianRemoved events raised by the GuardianRecoveryValidator contract.
type GuardianRecoveryValidatorGuardianRemovedIterator struct {
	Event *GuardianRecoveryValidatorGuardianRemoved // Event containing the contract specifics and raw log

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
func (it *GuardianRecoveryValidatorGuardianRemovedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(GuardianRecoveryValidatorGuardianRemoved)
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
		it.Event = new(GuardianRecoveryValidatorGuardianRemoved)
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
func (it *GuardianRecoveryValidatorGuardianRemovedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *GuardianRecoveryValidatorGuardianRemovedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// GuardianRecoveryValidatorGuardianRemoved represents a GuardianRemoved event raised by the GuardianRecoveryValidator contract.
type GuardianRecoveryValidatorGuardianRemoved struct {
	Account            common.Address
	HashedOriginDomain [32]byte
	Guardian           common.Address
	Raw                types.Log // Blockchain specific contextual infos
}

// FilterGuardianRemoved is a free log retrieval operation binding the contract event 0xb9c57fa4d0e004062979bfe03380ddabcf23ec66f36e471ee5b5ab67349485b2.
//
// Solidity: event GuardianRemoved(address indexed account, bytes32 indexed hashedOriginDomain, address indexed guardian)
func (_GuardianRecoveryValidator *GuardianRecoveryValidatorFilterer) FilterGuardianRemoved(opts *bind.FilterOpts, account []common.Address, hashedOriginDomain [][32]byte, guardian []common.Address) (*GuardianRecoveryValidatorGuardianRemovedIterator, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}
	var hashedOriginDomainRule []interface{}
	for _, hashedOriginDomainItem := range hashedOriginDomain {
		hashedOriginDomainRule = append(hashedOriginDomainRule, hashedOriginDomainItem)
	}
	var guardianRule []interface{}
	for _, guardianItem := range guardian {
		guardianRule = append(guardianRule, guardianItem)
	}

	logs, sub, err := _GuardianRecoveryValidator.contract.FilterLogs(opts, "GuardianRemoved", accountRule, hashedOriginDomainRule, guardianRule)
	if err != nil {
		return nil, err
	}
	return &GuardianRecoveryValidatorGuardianRemovedIterator{contract: _GuardianRecoveryValidator.contract, event: "GuardianRemoved", logs: logs, sub: sub}, nil
}

// WatchGuardianRemoved is a free log subscription operation binding the contract event 0xb9c57fa4d0e004062979bfe03380ddabcf23ec66f36e471ee5b5ab67349485b2.
//
// Solidity: event GuardianRemoved(address indexed account, bytes32 indexed hashedOriginDomain, address indexed guardian)
func (_GuardianRecoveryValidator *GuardianRecoveryValidatorFilterer) WatchGuardianRemoved(opts *bind.WatchOpts, sink chan<- *GuardianRecoveryValidatorGuardianRemoved, account []common.Address, hashedOriginDomain [][32]byte, guardian []common.Address) (event.Subscription, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}
	var hashedOriginDomainRule []interface{}
	for _, hashedOriginDomainItem := range hashedOriginDomain {
		hashedOriginDomainRule = append(hashedOriginDomainRule, hashedOriginDomainItem)
	}
	var guardianRule []interface{}
	for _, guardianItem := range guardian {
		guardianRule = append(guardianRule, guardianItem)
	}

	logs, sub, err := _GuardianRecoveryValidator.contract.WatchLogs(opts, "GuardianRemoved", accountRule, hashedOriginDomainRule, guardianRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(GuardianRecoveryValidatorGuardianRemoved)
				if err := _GuardianRecoveryValidator.contract.UnpackLog(event, "GuardianRemoved", log); err != nil {
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

// ParseGuardianRemoved is a log parse operation binding the contract event 0xb9c57fa4d0e004062979bfe03380ddabcf23ec66f36e471ee5b5ab67349485b2.
//
// Solidity: event GuardianRemoved(address indexed account, bytes32 indexed hashedOriginDomain, address indexed guardian)
func (_GuardianRecoveryValidator *GuardianRecoveryValidatorFilterer) ParseGuardianRemoved(log types.Log) (*GuardianRecoveryValidatorGuardianRemoved, error) {
	event := new(GuardianRecoveryValidatorGuardianRemoved)
	if err := _GuardianRecoveryValidator.contract.UnpackLog(event, "GuardianRemoved", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// GuardianRecoveryValidatorHashedOriginDomainDisabledForAccountIterator is returned from FilterHashedOriginDomainDisabledForAccount and is used to iterate over the raw logs and unpacked data for HashedOriginDomainDisabledForAccount events raised by the GuardianRecoveryValidator contract.
type GuardianRecoveryValidatorHashedOriginDomainDisabledForAccountIterator struct {
	Event *GuardianRecoveryValidatorHashedOriginDomainDisabledForAccount // Event containing the contract specifics and raw log

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
func (it *GuardianRecoveryValidatorHashedOriginDomainDisabledForAccountIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(GuardianRecoveryValidatorHashedOriginDomainDisabledForAccount)
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
		it.Event = new(GuardianRecoveryValidatorHashedOriginDomainDisabledForAccount)
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
func (it *GuardianRecoveryValidatorHashedOriginDomainDisabledForAccountIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *GuardianRecoveryValidatorHashedOriginDomainDisabledForAccountIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// GuardianRecoveryValidatorHashedOriginDomainDisabledForAccount represents a HashedOriginDomainDisabledForAccount event raised by the GuardianRecoveryValidator contract.
type GuardianRecoveryValidatorHashedOriginDomainDisabledForAccount struct {
	Account            common.Address
	HashedOriginDomain [32]byte
	Raw                types.Log // Blockchain specific contextual infos
}

// FilterHashedOriginDomainDisabledForAccount is a free log retrieval operation binding the contract event 0x1373f4151802c1c04091d06c3f13c1948e40afad22bd00382d4702808d0320eb.
//
// Solidity: event HashedOriginDomainDisabledForAccount(address indexed account, bytes32 indexed hashedOriginDomain)
func (_GuardianRecoveryValidator *GuardianRecoveryValidatorFilterer) FilterHashedOriginDomainDisabledForAccount(opts *bind.FilterOpts, account []common.Address, hashedOriginDomain [][32]byte) (*GuardianRecoveryValidatorHashedOriginDomainDisabledForAccountIterator, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}
	var hashedOriginDomainRule []interface{}
	for _, hashedOriginDomainItem := range hashedOriginDomain {
		hashedOriginDomainRule = append(hashedOriginDomainRule, hashedOriginDomainItem)
	}

	logs, sub, err := _GuardianRecoveryValidator.contract.FilterLogs(opts, "HashedOriginDomainDisabledForAccount", accountRule, hashedOriginDomainRule)
	if err != nil {
		return nil, err
	}
	return &GuardianRecoveryValidatorHashedOriginDomainDisabledForAccountIterator{contract: _GuardianRecoveryValidator.contract, event: "HashedOriginDomainDisabledForAccount", logs: logs, sub: sub}, nil
}

// WatchHashedOriginDomainDisabledForAccount is a free log subscription operation binding the contract event 0x1373f4151802c1c04091d06c3f13c1948e40afad22bd00382d4702808d0320eb.
//
// Solidity: event HashedOriginDomainDisabledForAccount(address indexed account, bytes32 indexed hashedOriginDomain)
func (_GuardianRecoveryValidator *GuardianRecoveryValidatorFilterer) WatchHashedOriginDomainDisabledForAccount(opts *bind.WatchOpts, sink chan<- *GuardianRecoveryValidatorHashedOriginDomainDisabledForAccount, account []common.Address, hashedOriginDomain [][32]byte) (event.Subscription, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}
	var hashedOriginDomainRule []interface{}
	for _, hashedOriginDomainItem := range hashedOriginDomain {
		hashedOriginDomainRule = append(hashedOriginDomainRule, hashedOriginDomainItem)
	}

	logs, sub, err := _GuardianRecoveryValidator.contract.WatchLogs(opts, "HashedOriginDomainDisabledForAccount", accountRule, hashedOriginDomainRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(GuardianRecoveryValidatorHashedOriginDomainDisabledForAccount)
				if err := _GuardianRecoveryValidator.contract.UnpackLog(event, "HashedOriginDomainDisabledForAccount", log); err != nil {
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

// ParseHashedOriginDomainDisabledForAccount is a log parse operation binding the contract event 0x1373f4151802c1c04091d06c3f13c1948e40afad22bd00382d4702808d0320eb.
//
// Solidity: event HashedOriginDomainDisabledForAccount(address indexed account, bytes32 indexed hashedOriginDomain)
func (_GuardianRecoveryValidator *GuardianRecoveryValidatorFilterer) ParseHashedOriginDomainDisabledForAccount(log types.Log) (*GuardianRecoveryValidatorHashedOriginDomainDisabledForAccount, error) {
	event := new(GuardianRecoveryValidatorHashedOriginDomainDisabledForAccount)
	if err := _GuardianRecoveryValidator.contract.UnpackLog(event, "HashedOriginDomainDisabledForAccount", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// GuardianRecoveryValidatorHashedOriginDomainEnabledForAccountIterator is returned from FilterHashedOriginDomainEnabledForAccount and is used to iterate over the raw logs and unpacked data for HashedOriginDomainEnabledForAccount events raised by the GuardianRecoveryValidator contract.
type GuardianRecoveryValidatorHashedOriginDomainEnabledForAccountIterator struct {
	Event *GuardianRecoveryValidatorHashedOriginDomainEnabledForAccount // Event containing the contract specifics and raw log

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
func (it *GuardianRecoveryValidatorHashedOriginDomainEnabledForAccountIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(GuardianRecoveryValidatorHashedOriginDomainEnabledForAccount)
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
		it.Event = new(GuardianRecoveryValidatorHashedOriginDomainEnabledForAccount)
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
func (it *GuardianRecoveryValidatorHashedOriginDomainEnabledForAccountIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *GuardianRecoveryValidatorHashedOriginDomainEnabledForAccountIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// GuardianRecoveryValidatorHashedOriginDomainEnabledForAccount represents a HashedOriginDomainEnabledForAccount event raised by the GuardianRecoveryValidator contract.
type GuardianRecoveryValidatorHashedOriginDomainEnabledForAccount struct {
	Account            common.Address
	HashedOriginDomain [32]byte
	Raw                types.Log // Blockchain specific contextual infos
}

// FilterHashedOriginDomainEnabledForAccount is a free log retrieval operation binding the contract event 0x49e308297bf7e3373130f54d11e9944d094c586954301b5c7da3b918195c7580.
//
// Solidity: event HashedOriginDomainEnabledForAccount(address indexed account, bytes32 indexed hashedOriginDomain)
func (_GuardianRecoveryValidator *GuardianRecoveryValidatorFilterer) FilterHashedOriginDomainEnabledForAccount(opts *bind.FilterOpts, account []common.Address, hashedOriginDomain [][32]byte) (*GuardianRecoveryValidatorHashedOriginDomainEnabledForAccountIterator, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}
	var hashedOriginDomainRule []interface{}
	for _, hashedOriginDomainItem := range hashedOriginDomain {
		hashedOriginDomainRule = append(hashedOriginDomainRule, hashedOriginDomainItem)
	}

	logs, sub, err := _GuardianRecoveryValidator.contract.FilterLogs(opts, "HashedOriginDomainEnabledForAccount", accountRule, hashedOriginDomainRule)
	if err != nil {
		return nil, err
	}
	return &GuardianRecoveryValidatorHashedOriginDomainEnabledForAccountIterator{contract: _GuardianRecoveryValidator.contract, event: "HashedOriginDomainEnabledForAccount", logs: logs, sub: sub}, nil
}

// WatchHashedOriginDomainEnabledForAccount is a free log subscription operation binding the contract event 0x49e308297bf7e3373130f54d11e9944d094c586954301b5c7da3b918195c7580.
//
// Solidity: event HashedOriginDomainEnabledForAccount(address indexed account, bytes32 indexed hashedOriginDomain)
func (_GuardianRecoveryValidator *GuardianRecoveryValidatorFilterer) WatchHashedOriginDomainEnabledForAccount(opts *bind.WatchOpts, sink chan<- *GuardianRecoveryValidatorHashedOriginDomainEnabledForAccount, account []common.Address, hashedOriginDomain [][32]byte) (event.Subscription, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}
	var hashedOriginDomainRule []interface{}
	for _, hashedOriginDomainItem := range hashedOriginDomain {
		hashedOriginDomainRule = append(hashedOriginDomainRule, hashedOriginDomainItem)
	}

	logs, sub, err := _GuardianRecoveryValidator.contract.WatchLogs(opts, "HashedOriginDomainEnabledForAccount", accountRule, hashedOriginDomainRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(GuardianRecoveryValidatorHashedOriginDomainEnabledForAccount)
				if err := _GuardianRecoveryValidator.contract.UnpackLog(event, "HashedOriginDomainEnabledForAccount", log); err != nil {
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

// ParseHashedOriginDomainEnabledForAccount is a log parse operation binding the contract event 0x49e308297bf7e3373130f54d11e9944d094c586954301b5c7da3b918195c7580.
//
// Solidity: event HashedOriginDomainEnabledForAccount(address indexed account, bytes32 indexed hashedOriginDomain)
func (_GuardianRecoveryValidator *GuardianRecoveryValidatorFilterer) ParseHashedOriginDomainEnabledForAccount(log types.Log) (*GuardianRecoveryValidatorHashedOriginDomainEnabledForAccount, error) {
	event := new(GuardianRecoveryValidatorHashedOriginDomainEnabledForAccount)
	if err := _GuardianRecoveryValidator.contract.UnpackLog(event, "HashedOriginDomainEnabledForAccount", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// GuardianRecoveryValidatorInitializedIterator is returned from FilterInitialized and is used to iterate over the raw logs and unpacked data for Initialized events raised by the GuardianRecoveryValidator contract.
type GuardianRecoveryValidatorInitializedIterator struct {
	Event *GuardianRecoveryValidatorInitialized // Event containing the contract specifics and raw log

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
func (it *GuardianRecoveryValidatorInitializedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(GuardianRecoveryValidatorInitialized)
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
		it.Event = new(GuardianRecoveryValidatorInitialized)
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
func (it *GuardianRecoveryValidatorInitializedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *GuardianRecoveryValidatorInitializedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// GuardianRecoveryValidatorInitialized represents a Initialized event raised by the GuardianRecoveryValidator contract.
type GuardianRecoveryValidatorInitialized struct {
	Version uint8
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterInitialized is a free log retrieval operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_GuardianRecoveryValidator *GuardianRecoveryValidatorFilterer) FilterInitialized(opts *bind.FilterOpts) (*GuardianRecoveryValidatorInitializedIterator, error) {

	logs, sub, err := _GuardianRecoveryValidator.contract.FilterLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return &GuardianRecoveryValidatorInitializedIterator{contract: _GuardianRecoveryValidator.contract, event: "Initialized", logs: logs, sub: sub}, nil
}

// WatchInitialized is a free log subscription operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_GuardianRecoveryValidator *GuardianRecoveryValidatorFilterer) WatchInitialized(opts *bind.WatchOpts, sink chan<- *GuardianRecoveryValidatorInitialized) (event.Subscription, error) {

	logs, sub, err := _GuardianRecoveryValidator.contract.WatchLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(GuardianRecoveryValidatorInitialized)
				if err := _GuardianRecoveryValidator.contract.UnpackLog(event, "Initialized", log); err != nil {
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

// ParseInitialized is a log parse operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_GuardianRecoveryValidator *GuardianRecoveryValidatorFilterer) ParseInitialized(log types.Log) (*GuardianRecoveryValidatorInitialized, error) {
	event := new(GuardianRecoveryValidatorInitialized)
	if err := _GuardianRecoveryValidator.contract.UnpackLog(event, "Initialized", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// GuardianRecoveryValidatorRecoveryDiscardedIterator is returned from FilterRecoveryDiscarded and is used to iterate over the raw logs and unpacked data for RecoveryDiscarded events raised by the GuardianRecoveryValidator contract.
type GuardianRecoveryValidatorRecoveryDiscardedIterator struct {
	Event *GuardianRecoveryValidatorRecoveryDiscarded // Event containing the contract specifics and raw log

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
func (it *GuardianRecoveryValidatorRecoveryDiscardedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(GuardianRecoveryValidatorRecoveryDiscarded)
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
		it.Event = new(GuardianRecoveryValidatorRecoveryDiscarded)
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
func (it *GuardianRecoveryValidatorRecoveryDiscardedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *GuardianRecoveryValidatorRecoveryDiscardedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// GuardianRecoveryValidatorRecoveryDiscarded represents a RecoveryDiscarded event raised by the GuardianRecoveryValidator contract.
type GuardianRecoveryValidatorRecoveryDiscarded struct {
	Account            common.Address
	HashedOriginDomain [32]byte
	HashedCredentialId [32]byte
	Raw                types.Log // Blockchain specific contextual infos
}

// FilterRecoveryDiscarded is a free log retrieval operation binding the contract event 0xa4e8705d9ef0c51657f9c64b028c7199ff3cd3ee5d057d6e1088e64580853d13.
//
// Solidity: event RecoveryDiscarded(address indexed account, bytes32 indexed hashedOriginDomain, bytes32 indexed hashedCredentialId)
func (_GuardianRecoveryValidator *GuardianRecoveryValidatorFilterer) FilterRecoveryDiscarded(opts *bind.FilterOpts, account []common.Address, hashedOriginDomain [][32]byte, hashedCredentialId [][32]byte) (*GuardianRecoveryValidatorRecoveryDiscardedIterator, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}
	var hashedOriginDomainRule []interface{}
	for _, hashedOriginDomainItem := range hashedOriginDomain {
		hashedOriginDomainRule = append(hashedOriginDomainRule, hashedOriginDomainItem)
	}
	var hashedCredentialIdRule []interface{}
	for _, hashedCredentialIdItem := range hashedCredentialId {
		hashedCredentialIdRule = append(hashedCredentialIdRule, hashedCredentialIdItem)
	}

	logs, sub, err := _GuardianRecoveryValidator.contract.FilterLogs(opts, "RecoveryDiscarded", accountRule, hashedOriginDomainRule, hashedCredentialIdRule)
	if err != nil {
		return nil, err
	}
	return &GuardianRecoveryValidatorRecoveryDiscardedIterator{contract: _GuardianRecoveryValidator.contract, event: "RecoveryDiscarded", logs: logs, sub: sub}, nil
}

// WatchRecoveryDiscarded is a free log subscription operation binding the contract event 0xa4e8705d9ef0c51657f9c64b028c7199ff3cd3ee5d057d6e1088e64580853d13.
//
// Solidity: event RecoveryDiscarded(address indexed account, bytes32 indexed hashedOriginDomain, bytes32 indexed hashedCredentialId)
func (_GuardianRecoveryValidator *GuardianRecoveryValidatorFilterer) WatchRecoveryDiscarded(opts *bind.WatchOpts, sink chan<- *GuardianRecoveryValidatorRecoveryDiscarded, account []common.Address, hashedOriginDomain [][32]byte, hashedCredentialId [][32]byte) (event.Subscription, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}
	var hashedOriginDomainRule []interface{}
	for _, hashedOriginDomainItem := range hashedOriginDomain {
		hashedOriginDomainRule = append(hashedOriginDomainRule, hashedOriginDomainItem)
	}
	var hashedCredentialIdRule []interface{}
	for _, hashedCredentialIdItem := range hashedCredentialId {
		hashedCredentialIdRule = append(hashedCredentialIdRule, hashedCredentialIdItem)
	}

	logs, sub, err := _GuardianRecoveryValidator.contract.WatchLogs(opts, "RecoveryDiscarded", accountRule, hashedOriginDomainRule, hashedCredentialIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(GuardianRecoveryValidatorRecoveryDiscarded)
				if err := _GuardianRecoveryValidator.contract.UnpackLog(event, "RecoveryDiscarded", log); err != nil {
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

// ParseRecoveryDiscarded is a log parse operation binding the contract event 0xa4e8705d9ef0c51657f9c64b028c7199ff3cd3ee5d057d6e1088e64580853d13.
//
// Solidity: event RecoveryDiscarded(address indexed account, bytes32 indexed hashedOriginDomain, bytes32 indexed hashedCredentialId)
func (_GuardianRecoveryValidator *GuardianRecoveryValidatorFilterer) ParseRecoveryDiscarded(log types.Log) (*GuardianRecoveryValidatorRecoveryDiscarded, error) {
	event := new(GuardianRecoveryValidatorRecoveryDiscarded)
	if err := _GuardianRecoveryValidator.contract.UnpackLog(event, "RecoveryDiscarded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// GuardianRecoveryValidatorRecoveryFinishedIterator is returned from FilterRecoveryFinished and is used to iterate over the raw logs and unpacked data for RecoveryFinished events raised by the GuardianRecoveryValidator contract.
type GuardianRecoveryValidatorRecoveryFinishedIterator struct {
	Event *GuardianRecoveryValidatorRecoveryFinished // Event containing the contract specifics and raw log

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
func (it *GuardianRecoveryValidatorRecoveryFinishedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(GuardianRecoveryValidatorRecoveryFinished)
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
		it.Event = new(GuardianRecoveryValidatorRecoveryFinished)
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
func (it *GuardianRecoveryValidatorRecoveryFinishedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *GuardianRecoveryValidatorRecoveryFinishedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// GuardianRecoveryValidatorRecoveryFinished represents a RecoveryFinished event raised by the GuardianRecoveryValidator contract.
type GuardianRecoveryValidatorRecoveryFinished struct {
	Account            common.Address
	HashedOriginDomain [32]byte
	HashedCredentialId [32]byte
	Raw                types.Log // Blockchain specific contextual infos
}

// FilterRecoveryFinished is a free log retrieval operation binding the contract event 0x122f699d3998d7d483d87a65abbdba44780a412ca03b283c5c6fc45b6d7d94d2.
//
// Solidity: event RecoveryFinished(address indexed account, bytes32 indexed hashedOriginDomain, bytes32 indexed hashedCredentialId)
func (_GuardianRecoveryValidator *GuardianRecoveryValidatorFilterer) FilterRecoveryFinished(opts *bind.FilterOpts, account []common.Address, hashedOriginDomain [][32]byte, hashedCredentialId [][32]byte) (*GuardianRecoveryValidatorRecoveryFinishedIterator, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}
	var hashedOriginDomainRule []interface{}
	for _, hashedOriginDomainItem := range hashedOriginDomain {
		hashedOriginDomainRule = append(hashedOriginDomainRule, hashedOriginDomainItem)
	}
	var hashedCredentialIdRule []interface{}
	for _, hashedCredentialIdItem := range hashedCredentialId {
		hashedCredentialIdRule = append(hashedCredentialIdRule, hashedCredentialIdItem)
	}

	logs, sub, err := _GuardianRecoveryValidator.contract.FilterLogs(opts, "RecoveryFinished", accountRule, hashedOriginDomainRule, hashedCredentialIdRule)
	if err != nil {
		return nil, err
	}
	return &GuardianRecoveryValidatorRecoveryFinishedIterator{contract: _GuardianRecoveryValidator.contract, event: "RecoveryFinished", logs: logs, sub: sub}, nil
}

// WatchRecoveryFinished is a free log subscription operation binding the contract event 0x122f699d3998d7d483d87a65abbdba44780a412ca03b283c5c6fc45b6d7d94d2.
//
// Solidity: event RecoveryFinished(address indexed account, bytes32 indexed hashedOriginDomain, bytes32 indexed hashedCredentialId)
func (_GuardianRecoveryValidator *GuardianRecoveryValidatorFilterer) WatchRecoveryFinished(opts *bind.WatchOpts, sink chan<- *GuardianRecoveryValidatorRecoveryFinished, account []common.Address, hashedOriginDomain [][32]byte, hashedCredentialId [][32]byte) (event.Subscription, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}
	var hashedOriginDomainRule []interface{}
	for _, hashedOriginDomainItem := range hashedOriginDomain {
		hashedOriginDomainRule = append(hashedOriginDomainRule, hashedOriginDomainItem)
	}
	var hashedCredentialIdRule []interface{}
	for _, hashedCredentialIdItem := range hashedCredentialId {
		hashedCredentialIdRule = append(hashedCredentialIdRule, hashedCredentialIdItem)
	}

	logs, sub, err := _GuardianRecoveryValidator.contract.WatchLogs(opts, "RecoveryFinished", accountRule, hashedOriginDomainRule, hashedCredentialIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(GuardianRecoveryValidatorRecoveryFinished)
				if err := _GuardianRecoveryValidator.contract.UnpackLog(event, "RecoveryFinished", log); err != nil {
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

// ParseRecoveryFinished is a log parse operation binding the contract event 0x122f699d3998d7d483d87a65abbdba44780a412ca03b283c5c6fc45b6d7d94d2.
//
// Solidity: event RecoveryFinished(address indexed account, bytes32 indexed hashedOriginDomain, bytes32 indexed hashedCredentialId)
func (_GuardianRecoveryValidator *GuardianRecoveryValidatorFilterer) ParseRecoveryFinished(log types.Log) (*GuardianRecoveryValidatorRecoveryFinished, error) {
	event := new(GuardianRecoveryValidatorRecoveryFinished)
	if err := _GuardianRecoveryValidator.contract.UnpackLog(event, "RecoveryFinished", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// GuardianRecoveryValidatorRecoveryInitiatedIterator is returned from FilterRecoveryInitiated and is used to iterate over the raw logs and unpacked data for RecoveryInitiated events raised by the GuardianRecoveryValidator contract.
type GuardianRecoveryValidatorRecoveryInitiatedIterator struct {
	Event *GuardianRecoveryValidatorRecoveryInitiated // Event containing the contract specifics and raw log

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
func (it *GuardianRecoveryValidatorRecoveryInitiatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(GuardianRecoveryValidatorRecoveryInitiated)
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
		it.Event = new(GuardianRecoveryValidatorRecoveryInitiated)
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
func (it *GuardianRecoveryValidatorRecoveryInitiatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *GuardianRecoveryValidatorRecoveryInitiatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// GuardianRecoveryValidatorRecoveryInitiated represents a RecoveryInitiated event raised by the GuardianRecoveryValidator contract.
type GuardianRecoveryValidatorRecoveryInitiated struct {
	Account            common.Address
	HashedOriginDomain [32]byte
	HashedCredentialId [32]byte
	Guardian           common.Address
	Raw                types.Log // Blockchain specific contextual infos
}

// FilterRecoveryInitiated is a free log retrieval operation binding the contract event 0x750203e1c6151c9136eca3be6c61d060ea236395244e97c72ee3735c131c3802.
//
// Solidity: event RecoveryInitiated(address indexed account, bytes32 indexed hashedOriginDomain, bytes32 indexed hashedCredentialId, address guardian)
func (_GuardianRecoveryValidator *GuardianRecoveryValidatorFilterer) FilterRecoveryInitiated(opts *bind.FilterOpts, account []common.Address, hashedOriginDomain [][32]byte, hashedCredentialId [][32]byte) (*GuardianRecoveryValidatorRecoveryInitiatedIterator, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}
	var hashedOriginDomainRule []interface{}
	for _, hashedOriginDomainItem := range hashedOriginDomain {
		hashedOriginDomainRule = append(hashedOriginDomainRule, hashedOriginDomainItem)
	}
	var hashedCredentialIdRule []interface{}
	for _, hashedCredentialIdItem := range hashedCredentialId {
		hashedCredentialIdRule = append(hashedCredentialIdRule, hashedCredentialIdItem)
	}

	logs, sub, err := _GuardianRecoveryValidator.contract.FilterLogs(opts, "RecoveryInitiated", accountRule, hashedOriginDomainRule, hashedCredentialIdRule)
	if err != nil {
		return nil, err
	}
	return &GuardianRecoveryValidatorRecoveryInitiatedIterator{contract: _GuardianRecoveryValidator.contract, event: "RecoveryInitiated", logs: logs, sub: sub}, nil
}

// WatchRecoveryInitiated is a free log subscription operation binding the contract event 0x750203e1c6151c9136eca3be6c61d060ea236395244e97c72ee3735c131c3802.
//
// Solidity: event RecoveryInitiated(address indexed account, bytes32 indexed hashedOriginDomain, bytes32 indexed hashedCredentialId, address guardian)
func (_GuardianRecoveryValidator *GuardianRecoveryValidatorFilterer) WatchRecoveryInitiated(opts *bind.WatchOpts, sink chan<- *GuardianRecoveryValidatorRecoveryInitiated, account []common.Address, hashedOriginDomain [][32]byte, hashedCredentialId [][32]byte) (event.Subscription, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}
	var hashedOriginDomainRule []interface{}
	for _, hashedOriginDomainItem := range hashedOriginDomain {
		hashedOriginDomainRule = append(hashedOriginDomainRule, hashedOriginDomainItem)
	}
	var hashedCredentialIdRule []interface{}
	for _, hashedCredentialIdItem := range hashedCredentialId {
		hashedCredentialIdRule = append(hashedCredentialIdRule, hashedCredentialIdItem)
	}

	logs, sub, err := _GuardianRecoveryValidator.contract.WatchLogs(opts, "RecoveryInitiated", accountRule, hashedOriginDomainRule, hashedCredentialIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(GuardianRecoveryValidatorRecoveryInitiated)
				if err := _GuardianRecoveryValidator.contract.UnpackLog(event, "RecoveryInitiated", log); err != nil {
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

// ParseRecoveryInitiated is a log parse operation binding the contract event 0x750203e1c6151c9136eca3be6c61d060ea236395244e97c72ee3735c131c3802.
//
// Solidity: event RecoveryInitiated(address indexed account, bytes32 indexed hashedOriginDomain, bytes32 indexed hashedCredentialId, address guardian)
func (_GuardianRecoveryValidator *GuardianRecoveryValidatorFilterer) ParseRecoveryInitiated(log types.Log) (*GuardianRecoveryValidatorRecoveryInitiated, error) {
	event := new(GuardianRecoveryValidatorRecoveryInitiated)
	if err := _GuardianRecoveryValidator.contract.UnpackLog(event, "RecoveryInitiated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
