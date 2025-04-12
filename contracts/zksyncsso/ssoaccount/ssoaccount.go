// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package ssoaccount

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

// Call is an auto generated low-level Go binding around an user-defined struct.
type Call struct {
	Target       common.Address
	AllowFailure bool
	Value        *big.Int
	CallData     []byte
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

// SsoAccountMetaData contains all meta data concerning the SsoAccount contract.
var SsoAccountMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"ADDRESS_CAST_OVERFLOW\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"actualValue\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"expectedValue\",\"type\":\"uint256\"}],\"name\":\"BATCH_MSG_VALUE_MISMATCH\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"FEE_PAYMENT_FAILED\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"hook\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"isValidation\",\"type\":\"bool\"}],\"name\":\"HOOK_ALREADY_EXISTS\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"hookAddress\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"isValidation\",\"type\":\"bool\"}],\"name\":\"HOOK_ERC165_FAIL\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"hook\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"isValidation\",\"type\":\"bool\"}],\"name\":\"HOOK_NOT_FOUND\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"required\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"available\",\"type\":\"uint256\"}],\"name\":\"INSUFFICIENT_FUNDS\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"INVALID_ACCOUNT_KEYS\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"METHOD_NOT_IMPLEMENTED\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"notBootloader\",\"type\":\"address\"}],\"name\":\"NOT_FROM_BOOTLOADER\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"notSelf\",\"type\":\"address\"}],\"name\":\"NOT_FROM_SELF\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"OWNER_ALREADY_EXISTS\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"OWNER_NOT_FOUND\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"validator\",\"type\":\"address\"}],\"name\":\"VALIDATOR_ALREADY_EXISTS\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"validator\",\"type\":\"address\"}],\"name\":\"VALIDATOR_ERC165_FAIL\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"validator\",\"type\":\"address\"}],\"name\":\"VALIDATOR_NOT_FOUND\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"revertData\",\"type\":\"bytes\"}],\"name\":\"BatchCallFailure\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"hook\",\"type\":\"address\"}],\"name\":\"HookAdded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"hook\",\"type\":\"address\"}],\"name\":\"HookRemoved\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint8\",\"name\":\"version\",\"type\":\"uint8\"}],\"name\":\"Initialized\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"addr\",\"type\":\"address\"}],\"name\":\"K1OwnerAdded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"addr\",\"type\":\"address\"}],\"name\":\"K1OwnerRemoved\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"validator\",\"type\":\"address\"}],\"name\":\"ValidatorAdded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"validator\",\"type\":\"address\"}],\"name\":\"ValidatorRemoved\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"hook\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"isValidation\",\"type\":\"bool\"},{\"internalType\":\"bytes\",\"name\":\"initData\",\"type\":\"bytes\"}],\"name\":\"addHook\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"addr\",\"type\":\"address\"}],\"name\":\"addK1Owner\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"validator\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"initData\",\"type\":\"bytes\"}],\"name\":\"addModuleValidator\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"target\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"allowFailure\",\"type\":\"bool\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"callData\",\"type\":\"bytes\"}],\"internalType\":\"structCall[]\",\"name\":\"_calls\",\"type\":\"tuple[]\"}],\"name\":\"batchCall\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"txType\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"from\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"to\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"gasLimit\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"gasPerPubdataByteLimit\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"maxFeePerGas\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"maxPriorityFeePerGas\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"paymaster\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"nonce\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"},{\"internalType\":\"uint256[4]\",\"name\":\"reserved\",\"type\":\"uint256[4]\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"signature\",\"type\":\"bytes\"},{\"internalType\":\"bytes32[]\",\"name\":\"factoryDeps\",\"type\":\"bytes32[]\"},{\"internalType\":\"bytes\",\"name\":\"paymasterInput\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"reservedDynamic\",\"type\":\"bytes\"}],\"internalType\":\"structTransaction\",\"name\":\"_transaction\",\"type\":\"tuple\"}],\"name\":\"executeTransaction\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"txType\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"from\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"to\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"gasLimit\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"gasPerPubdataByteLimit\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"maxFeePerGas\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"maxPriorityFeePerGas\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"paymaster\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"nonce\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"},{\"internalType\":\"uint256[4]\",\"name\":\"reserved\",\"type\":\"uint256[4]\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"signature\",\"type\":\"bytes\"},{\"internalType\":\"bytes32[]\",\"name\":\"factoryDeps\",\"type\":\"bytes32[]\"},{\"internalType\":\"bytes\",\"name\":\"paymasterInput\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"reservedDynamic\",\"type\":\"bytes\"}],\"internalType\":\"structTransaction\",\"name\":\"\",\"type\":\"tuple\"}],\"name\":\"executeTransactionFromOutside\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes[]\",\"name\":\"initialValidators\",\"type\":\"bytes[]\"},{\"internalType\":\"address[]\",\"name\":\"initialK1Owners\",\"type\":\"address[]\"}],\"name\":\"initialize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"addr\",\"type\":\"address\"}],\"name\":\"isHook\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"addr\",\"type\":\"address\"}],\"name\":\"isK1Owner\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"validator\",\"type\":\"address\"}],\"name\":\"isModuleValidator\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"hash\",\"type\":\"bytes32\"},{\"internalType\":\"bytes\",\"name\":\"signature\",\"type\":\"bytes\"}],\"name\":\"isValidSignature\",\"outputs\":[{\"internalType\":\"bytes4\",\"name\":\"magicValue\",\"type\":\"bytes4\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bool\",\"name\":\"isValidation\",\"type\":\"bool\"}],\"name\":\"listHooks\",\"outputs\":[{\"internalType\":\"address[]\",\"name\":\"hookList\",\"type\":\"address[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"listK1Owners\",\"outputs\":[{\"internalType\":\"address[]\",\"name\":\"k1OwnerList\",\"type\":\"address[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"listModuleValidators\",\"outputs\":[{\"internalType\":\"address[]\",\"name\":\"validatorList\",\"type\":\"address[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"uint256[]\",\"name\":\"\",\"type\":\"uint256[]\"},{\"internalType\":\"uint256[]\",\"name\":\"\",\"type\":\"uint256[]\"},{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"name\":\"onERC1155BatchReceived\",\"outputs\":[{\"internalType\":\"bytes4\",\"name\":\"\",\"type\":\"bytes4\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"name\":\"onERC1155Received\",\"outputs\":[{\"internalType\":\"bytes4\",\"name\":\"\",\"type\":\"bytes4\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"name\":\"onERC721Received\",\"outputs\":[{\"internalType\":\"bytes4\",\"name\":\"\",\"type\":\"bytes4\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"txType\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"from\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"to\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"gasLimit\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"gasPerPubdataByteLimit\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"maxFeePerGas\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"maxPriorityFeePerGas\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"paymaster\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"nonce\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"},{\"internalType\":\"uint256[4]\",\"name\":\"reserved\",\"type\":\"uint256[4]\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"signature\",\"type\":\"bytes\"},{\"internalType\":\"bytes32[]\",\"name\":\"factoryDeps\",\"type\":\"bytes32[]\"},{\"internalType\":\"bytes\",\"name\":\"paymasterInput\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"reservedDynamic\",\"type\":\"bytes\"}],\"internalType\":\"structTransaction\",\"name\":\"_transaction\",\"type\":\"tuple\"}],\"name\":\"payForTransaction\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"txType\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"from\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"to\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"gasLimit\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"gasPerPubdataByteLimit\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"maxFeePerGas\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"maxPriorityFeePerGas\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"paymaster\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"nonce\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"},{\"internalType\":\"uint256[4]\",\"name\":\"reserved\",\"type\":\"uint256[4]\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"signature\",\"type\":\"bytes\"},{\"internalType\":\"bytes32[]\",\"name\":\"factoryDeps\",\"type\":\"bytes32[]\"},{\"internalType\":\"bytes\",\"name\":\"paymasterInput\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"reservedDynamic\",\"type\":\"bytes\"}],\"internalType\":\"structTransaction\",\"name\":\"_transaction\",\"type\":\"tuple\"}],\"name\":\"prepareForPaymaster\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"hook\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"isValidation\",\"type\":\"bool\"},{\"internalType\":\"bytes\",\"name\":\"deinitData\",\"type\":\"bytes\"}],\"name\":\"removeHook\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"addr\",\"type\":\"address\"}],\"name\":\"removeK1Owner\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"validator\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"deinitData\",\"type\":\"bytes\"}],\"name\":\"removeModuleValidator\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes4\",\"name\":\"interfaceId\",\"type\":\"bytes4\"}],\"name\":\"supportsInterface\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"hook\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"isValidation\",\"type\":\"bool\"},{\"internalType\":\"bytes\",\"name\":\"deinitData\",\"type\":\"bytes\"}],\"name\":\"unlinkHook\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"validator\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"deinitData\",\"type\":\"bytes\"}],\"name\":\"unlinkModuleValidator\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"_suggestedSignedHash\",\"type\":\"bytes32\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"txType\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"from\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"to\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"gasLimit\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"gasPerPubdataByteLimit\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"maxFeePerGas\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"maxPriorityFeePerGas\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"paymaster\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"nonce\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"},{\"internalType\":\"uint256[4]\",\"name\":\"reserved\",\"type\":\"uint256[4]\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"signature\",\"type\":\"bytes\"},{\"internalType\":\"bytes32[]\",\"name\":\"factoryDeps\",\"type\":\"bytes32[]\"},{\"internalType\":\"bytes\",\"name\":\"paymasterInput\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"reservedDynamic\",\"type\":\"bytes\"}],\"internalType\":\"structTransaction\",\"name\":\"_transaction\",\"type\":\"tuple\"}],\"name\":\"validateTransaction\",\"outputs\":[{\"internalType\":\"bytes4\",\"name\":\"magic\",\"type\":\"bytes4\"}],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"stateMutability\":\"payable\",\"type\":\"receive\"}]",
}

// SsoAccountABI is the input ABI used to generate the binding from.
// Deprecated: Use SsoAccountMetaData.ABI instead.
var SsoAccountABI = SsoAccountMetaData.ABI

// SsoAccount is an auto generated Go binding around an Ethereum contract.
type SsoAccount struct {
	SsoAccountCaller     // Read-only binding to the contract
	SsoAccountTransactor // Write-only binding to the contract
	SsoAccountFilterer   // Log filterer for contract events
}

// SsoAccountCaller is an auto generated read-only Go binding around an Ethereum contract.
type SsoAccountCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SsoAccountTransactor is an auto generated write-only Go binding around an Ethereum contract.
type SsoAccountTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SsoAccountFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type SsoAccountFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SsoAccountSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type SsoAccountSession struct {
	Contract     *SsoAccount       // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// SsoAccountCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type SsoAccountCallerSession struct {
	Contract *SsoAccountCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts     // Call options to use throughout this session
}

// SsoAccountTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type SsoAccountTransactorSession struct {
	Contract     *SsoAccountTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts     // Transaction auth options to use throughout this session
}

// SsoAccountRaw is an auto generated low-level Go binding around an Ethereum contract.
type SsoAccountRaw struct {
	Contract *SsoAccount // Generic contract binding to access the raw methods on
}

// SsoAccountCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type SsoAccountCallerRaw struct {
	Contract *SsoAccountCaller // Generic read-only contract binding to access the raw methods on
}

// SsoAccountTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type SsoAccountTransactorRaw struct {
	Contract *SsoAccountTransactor // Generic write-only contract binding to access the raw methods on
}

// NewSsoAccount creates a new instance of SsoAccount, bound to a specific deployed contract.
func NewSsoAccount(address common.Address, backend bind.ContractBackend) (*SsoAccount, error) {
	contract, err := bindSsoAccount(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &SsoAccount{SsoAccountCaller: SsoAccountCaller{contract: contract}, SsoAccountTransactor: SsoAccountTransactor{contract: contract}, SsoAccountFilterer: SsoAccountFilterer{contract: contract}}, nil
}

// NewSsoAccountCaller creates a new read-only instance of SsoAccount, bound to a specific deployed contract.
func NewSsoAccountCaller(address common.Address, caller bind.ContractCaller) (*SsoAccountCaller, error) {
	contract, err := bindSsoAccount(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &SsoAccountCaller{contract: contract}, nil
}

// NewSsoAccountTransactor creates a new write-only instance of SsoAccount, bound to a specific deployed contract.
func NewSsoAccountTransactor(address common.Address, transactor bind.ContractTransactor) (*SsoAccountTransactor, error) {
	contract, err := bindSsoAccount(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &SsoAccountTransactor{contract: contract}, nil
}

// NewSsoAccountFilterer creates a new log filterer instance of SsoAccount, bound to a specific deployed contract.
func NewSsoAccountFilterer(address common.Address, filterer bind.ContractFilterer) (*SsoAccountFilterer, error) {
	contract, err := bindSsoAccount(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &SsoAccountFilterer{contract: contract}, nil
}

// bindSsoAccount binds a generic wrapper to an already deployed contract.
func bindSsoAccount(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := SsoAccountMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_SsoAccount *SsoAccountRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _SsoAccount.Contract.SsoAccountCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_SsoAccount *SsoAccountRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _SsoAccount.Contract.SsoAccountTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_SsoAccount *SsoAccountRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _SsoAccount.Contract.SsoAccountTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_SsoAccount *SsoAccountCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _SsoAccount.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_SsoAccount *SsoAccountTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _SsoAccount.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_SsoAccount *SsoAccountTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _SsoAccount.Contract.contract.Transact(opts, method, params...)
}

// IsHook is a free data retrieval call binding the contract method 0xd2676529.
//
// Solidity: function isHook(address addr) view returns(bool)
func (_SsoAccount *SsoAccountCaller) IsHook(opts *bind.CallOpts, addr common.Address) (bool, error) {
	var out []interface{}
	err := _SsoAccount.contract.Call(opts, &out, "isHook", addr)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsHook is a free data retrieval call binding the contract method 0xd2676529.
//
// Solidity: function isHook(address addr) view returns(bool)
func (_SsoAccount *SsoAccountSession) IsHook(addr common.Address) (bool, error) {
	return _SsoAccount.Contract.IsHook(&_SsoAccount.CallOpts, addr)
}

// IsHook is a free data retrieval call binding the contract method 0xd2676529.
//
// Solidity: function isHook(address addr) view returns(bool)
func (_SsoAccount *SsoAccountCallerSession) IsHook(addr common.Address) (bool, error) {
	return _SsoAccount.Contract.IsHook(&_SsoAccount.CallOpts, addr)
}

// IsK1Owner is a free data retrieval call binding the contract method 0x6e354cb8.
//
// Solidity: function isK1Owner(address addr) view returns(bool)
func (_SsoAccount *SsoAccountCaller) IsK1Owner(opts *bind.CallOpts, addr common.Address) (bool, error) {
	var out []interface{}
	err := _SsoAccount.contract.Call(opts, &out, "isK1Owner", addr)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsK1Owner is a free data retrieval call binding the contract method 0x6e354cb8.
//
// Solidity: function isK1Owner(address addr) view returns(bool)
func (_SsoAccount *SsoAccountSession) IsK1Owner(addr common.Address) (bool, error) {
	return _SsoAccount.Contract.IsK1Owner(&_SsoAccount.CallOpts, addr)
}

// IsK1Owner is a free data retrieval call binding the contract method 0x6e354cb8.
//
// Solidity: function isK1Owner(address addr) view returns(bool)
func (_SsoAccount *SsoAccountCallerSession) IsK1Owner(addr common.Address) (bool, error) {
	return _SsoAccount.Contract.IsK1Owner(&_SsoAccount.CallOpts, addr)
}

// IsModuleValidator is a free data retrieval call binding the contract method 0x9b7be156.
//
// Solidity: function isModuleValidator(address validator) view returns(bool)
func (_SsoAccount *SsoAccountCaller) IsModuleValidator(opts *bind.CallOpts, validator common.Address) (bool, error) {
	var out []interface{}
	err := _SsoAccount.contract.Call(opts, &out, "isModuleValidator", validator)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsModuleValidator is a free data retrieval call binding the contract method 0x9b7be156.
//
// Solidity: function isModuleValidator(address validator) view returns(bool)
func (_SsoAccount *SsoAccountSession) IsModuleValidator(validator common.Address) (bool, error) {
	return _SsoAccount.Contract.IsModuleValidator(&_SsoAccount.CallOpts, validator)
}

// IsModuleValidator is a free data retrieval call binding the contract method 0x9b7be156.
//
// Solidity: function isModuleValidator(address validator) view returns(bool)
func (_SsoAccount *SsoAccountCallerSession) IsModuleValidator(validator common.Address) (bool, error) {
	return _SsoAccount.Contract.IsModuleValidator(&_SsoAccount.CallOpts, validator)
}

// IsValidSignature is a free data retrieval call binding the contract method 0x1626ba7e.
//
// Solidity: function isValidSignature(bytes32 hash, bytes signature) view returns(bytes4 magicValue)
func (_SsoAccount *SsoAccountCaller) IsValidSignature(opts *bind.CallOpts, hash [32]byte, signature []byte) ([4]byte, error) {
	var out []interface{}
	err := _SsoAccount.contract.Call(opts, &out, "isValidSignature", hash, signature)

	if err != nil {
		return *new([4]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([4]byte)).(*[4]byte)

	return out0, err

}

// IsValidSignature is a free data retrieval call binding the contract method 0x1626ba7e.
//
// Solidity: function isValidSignature(bytes32 hash, bytes signature) view returns(bytes4 magicValue)
func (_SsoAccount *SsoAccountSession) IsValidSignature(hash [32]byte, signature []byte) ([4]byte, error) {
	return _SsoAccount.Contract.IsValidSignature(&_SsoAccount.CallOpts, hash, signature)
}

// IsValidSignature is a free data retrieval call binding the contract method 0x1626ba7e.
//
// Solidity: function isValidSignature(bytes32 hash, bytes signature) view returns(bytes4 magicValue)
func (_SsoAccount *SsoAccountCallerSession) IsValidSignature(hash [32]byte, signature []byte) ([4]byte, error) {
	return _SsoAccount.Contract.IsValidSignature(&_SsoAccount.CallOpts, hash, signature)
}

// ListHooks is a free data retrieval call binding the contract method 0xdb8a323f.
//
// Solidity: function listHooks(bool isValidation) view returns(address[] hookList)
func (_SsoAccount *SsoAccountCaller) ListHooks(opts *bind.CallOpts, isValidation bool) ([]common.Address, error) {
	var out []interface{}
	err := _SsoAccount.contract.Call(opts, &out, "listHooks", isValidation)

	if err != nil {
		return *new([]common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)

	return out0, err

}

// ListHooks is a free data retrieval call binding the contract method 0xdb8a323f.
//
// Solidity: function listHooks(bool isValidation) view returns(address[] hookList)
func (_SsoAccount *SsoAccountSession) ListHooks(isValidation bool) ([]common.Address, error) {
	return _SsoAccount.Contract.ListHooks(&_SsoAccount.CallOpts, isValidation)
}

// ListHooks is a free data retrieval call binding the contract method 0xdb8a323f.
//
// Solidity: function listHooks(bool isValidation) view returns(address[] hookList)
func (_SsoAccount *SsoAccountCallerSession) ListHooks(isValidation bool) ([]common.Address, error) {
	return _SsoAccount.Contract.ListHooks(&_SsoAccount.CallOpts, isValidation)
}

// ListK1Owners is a free data retrieval call binding the contract method 0x4f16eec8.
//
// Solidity: function listK1Owners() view returns(address[] k1OwnerList)
func (_SsoAccount *SsoAccountCaller) ListK1Owners(opts *bind.CallOpts) ([]common.Address, error) {
	var out []interface{}
	err := _SsoAccount.contract.Call(opts, &out, "listK1Owners")

	if err != nil {
		return *new([]common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)

	return out0, err

}

// ListK1Owners is a free data retrieval call binding the contract method 0x4f16eec8.
//
// Solidity: function listK1Owners() view returns(address[] k1OwnerList)
func (_SsoAccount *SsoAccountSession) ListK1Owners() ([]common.Address, error) {
	return _SsoAccount.Contract.ListK1Owners(&_SsoAccount.CallOpts)
}

// ListK1Owners is a free data retrieval call binding the contract method 0x4f16eec8.
//
// Solidity: function listK1Owners() view returns(address[] k1OwnerList)
func (_SsoAccount *SsoAccountCallerSession) ListK1Owners() ([]common.Address, error) {
	return _SsoAccount.Contract.ListK1Owners(&_SsoAccount.CallOpts)
}

// ListModuleValidators is a free data retrieval call binding the contract method 0x3d3a69bf.
//
// Solidity: function listModuleValidators() view returns(address[] validatorList)
func (_SsoAccount *SsoAccountCaller) ListModuleValidators(opts *bind.CallOpts) ([]common.Address, error) {
	var out []interface{}
	err := _SsoAccount.contract.Call(opts, &out, "listModuleValidators")

	if err != nil {
		return *new([]common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)

	return out0, err

}

// ListModuleValidators is a free data retrieval call binding the contract method 0x3d3a69bf.
//
// Solidity: function listModuleValidators() view returns(address[] validatorList)
func (_SsoAccount *SsoAccountSession) ListModuleValidators() ([]common.Address, error) {
	return _SsoAccount.Contract.ListModuleValidators(&_SsoAccount.CallOpts)
}

// ListModuleValidators is a free data retrieval call binding the contract method 0x3d3a69bf.
//
// Solidity: function listModuleValidators() view returns(address[] validatorList)
func (_SsoAccount *SsoAccountCallerSession) ListModuleValidators() ([]common.Address, error) {
	return _SsoAccount.Contract.ListModuleValidators(&_SsoAccount.CallOpts)
}

// OnERC1155BatchReceived is a free data retrieval call binding the contract method 0xbc197c81.
//
// Solidity: function onERC1155BatchReceived(address , address , uint256[] , uint256[] , bytes ) pure returns(bytes4)
func (_SsoAccount *SsoAccountCaller) OnERC1155BatchReceived(opts *bind.CallOpts, arg0 common.Address, arg1 common.Address, arg2 []*big.Int, arg3 []*big.Int, arg4 []byte) ([4]byte, error) {
	var out []interface{}
	err := _SsoAccount.contract.Call(opts, &out, "onERC1155BatchReceived", arg0, arg1, arg2, arg3, arg4)

	if err != nil {
		return *new([4]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([4]byte)).(*[4]byte)

	return out0, err

}

// OnERC1155BatchReceived is a free data retrieval call binding the contract method 0xbc197c81.
//
// Solidity: function onERC1155BatchReceived(address , address , uint256[] , uint256[] , bytes ) pure returns(bytes4)
func (_SsoAccount *SsoAccountSession) OnERC1155BatchReceived(arg0 common.Address, arg1 common.Address, arg2 []*big.Int, arg3 []*big.Int, arg4 []byte) ([4]byte, error) {
	return _SsoAccount.Contract.OnERC1155BatchReceived(&_SsoAccount.CallOpts, arg0, arg1, arg2, arg3, arg4)
}

// OnERC1155BatchReceived is a free data retrieval call binding the contract method 0xbc197c81.
//
// Solidity: function onERC1155BatchReceived(address , address , uint256[] , uint256[] , bytes ) pure returns(bytes4)
func (_SsoAccount *SsoAccountCallerSession) OnERC1155BatchReceived(arg0 common.Address, arg1 common.Address, arg2 []*big.Int, arg3 []*big.Int, arg4 []byte) ([4]byte, error) {
	return _SsoAccount.Contract.OnERC1155BatchReceived(&_SsoAccount.CallOpts, arg0, arg1, arg2, arg3, arg4)
}

// OnERC1155Received is a free data retrieval call binding the contract method 0xf23a6e61.
//
// Solidity: function onERC1155Received(address , address , uint256 , uint256 , bytes ) pure returns(bytes4)
func (_SsoAccount *SsoAccountCaller) OnERC1155Received(opts *bind.CallOpts, arg0 common.Address, arg1 common.Address, arg2 *big.Int, arg3 *big.Int, arg4 []byte) ([4]byte, error) {
	var out []interface{}
	err := _SsoAccount.contract.Call(opts, &out, "onERC1155Received", arg0, arg1, arg2, arg3, arg4)

	if err != nil {
		return *new([4]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([4]byte)).(*[4]byte)

	return out0, err

}

// OnERC1155Received is a free data retrieval call binding the contract method 0xf23a6e61.
//
// Solidity: function onERC1155Received(address , address , uint256 , uint256 , bytes ) pure returns(bytes4)
func (_SsoAccount *SsoAccountSession) OnERC1155Received(arg0 common.Address, arg1 common.Address, arg2 *big.Int, arg3 *big.Int, arg4 []byte) ([4]byte, error) {
	return _SsoAccount.Contract.OnERC1155Received(&_SsoAccount.CallOpts, arg0, arg1, arg2, arg3, arg4)
}

// OnERC1155Received is a free data retrieval call binding the contract method 0xf23a6e61.
//
// Solidity: function onERC1155Received(address , address , uint256 , uint256 , bytes ) pure returns(bytes4)
func (_SsoAccount *SsoAccountCallerSession) OnERC1155Received(arg0 common.Address, arg1 common.Address, arg2 *big.Int, arg3 *big.Int, arg4 []byte) ([4]byte, error) {
	return _SsoAccount.Contract.OnERC1155Received(&_SsoAccount.CallOpts, arg0, arg1, arg2, arg3, arg4)
}

// OnERC721Received is a free data retrieval call binding the contract method 0x150b7a02.
//
// Solidity: function onERC721Received(address , address , uint256 , bytes ) pure returns(bytes4)
func (_SsoAccount *SsoAccountCaller) OnERC721Received(opts *bind.CallOpts, arg0 common.Address, arg1 common.Address, arg2 *big.Int, arg3 []byte) ([4]byte, error) {
	var out []interface{}
	err := _SsoAccount.contract.Call(opts, &out, "onERC721Received", arg0, arg1, arg2, arg3)

	if err != nil {
		return *new([4]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([4]byte)).(*[4]byte)

	return out0, err

}

// OnERC721Received is a free data retrieval call binding the contract method 0x150b7a02.
//
// Solidity: function onERC721Received(address , address , uint256 , bytes ) pure returns(bytes4)
func (_SsoAccount *SsoAccountSession) OnERC721Received(arg0 common.Address, arg1 common.Address, arg2 *big.Int, arg3 []byte) ([4]byte, error) {
	return _SsoAccount.Contract.OnERC721Received(&_SsoAccount.CallOpts, arg0, arg1, arg2, arg3)
}

// OnERC721Received is a free data retrieval call binding the contract method 0x150b7a02.
//
// Solidity: function onERC721Received(address , address , uint256 , bytes ) pure returns(bytes4)
func (_SsoAccount *SsoAccountCallerSession) OnERC721Received(arg0 common.Address, arg1 common.Address, arg2 *big.Int, arg3 []byte) ([4]byte, error) {
	return _SsoAccount.Contract.OnERC721Received(&_SsoAccount.CallOpts, arg0, arg1, arg2, arg3)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_SsoAccount *SsoAccountCaller) SupportsInterface(opts *bind.CallOpts, interfaceId [4]byte) (bool, error) {
	var out []interface{}
	err := _SsoAccount.contract.Call(opts, &out, "supportsInterface", interfaceId)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_SsoAccount *SsoAccountSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _SsoAccount.Contract.SupportsInterface(&_SsoAccount.CallOpts, interfaceId)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_SsoAccount *SsoAccountCallerSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _SsoAccount.Contract.SupportsInterface(&_SsoAccount.CallOpts, interfaceId)
}

// AddHook is a paid mutator transaction binding the contract method 0x0775a94e.
//
// Solidity: function addHook(address hook, bool isValidation, bytes initData) returns()
func (_SsoAccount *SsoAccountTransactor) AddHook(opts *bind.TransactOpts, hook common.Address, isValidation bool, initData []byte) (*types.Transaction, error) {
	return _SsoAccount.contract.Transact(opts, "addHook", hook, isValidation, initData)
}

// AddHook is a paid mutator transaction binding the contract method 0x0775a94e.
//
// Solidity: function addHook(address hook, bool isValidation, bytes initData) returns()
func (_SsoAccount *SsoAccountSession) AddHook(hook common.Address, isValidation bool, initData []byte) (*types.Transaction, error) {
	return _SsoAccount.Contract.AddHook(&_SsoAccount.TransactOpts, hook, isValidation, initData)
}

// AddHook is a paid mutator transaction binding the contract method 0x0775a94e.
//
// Solidity: function addHook(address hook, bool isValidation, bytes initData) returns()
func (_SsoAccount *SsoAccountTransactorSession) AddHook(hook common.Address, isValidation bool, initData []byte) (*types.Transaction, error) {
	return _SsoAccount.Contract.AddHook(&_SsoAccount.TransactOpts, hook, isValidation, initData)
}

// AddK1Owner is a paid mutator transaction binding the contract method 0x23cef13e.
//
// Solidity: function addK1Owner(address addr) returns()
func (_SsoAccount *SsoAccountTransactor) AddK1Owner(opts *bind.TransactOpts, addr common.Address) (*types.Transaction, error) {
	return _SsoAccount.contract.Transact(opts, "addK1Owner", addr)
}

// AddK1Owner is a paid mutator transaction binding the contract method 0x23cef13e.
//
// Solidity: function addK1Owner(address addr) returns()
func (_SsoAccount *SsoAccountSession) AddK1Owner(addr common.Address) (*types.Transaction, error) {
	return _SsoAccount.Contract.AddK1Owner(&_SsoAccount.TransactOpts, addr)
}

// AddK1Owner is a paid mutator transaction binding the contract method 0x23cef13e.
//
// Solidity: function addK1Owner(address addr) returns()
func (_SsoAccount *SsoAccountTransactorSession) AddK1Owner(addr common.Address) (*types.Transaction, error) {
	return _SsoAccount.Contract.AddK1Owner(&_SsoAccount.TransactOpts, addr)
}

// AddModuleValidator is a paid mutator transaction binding the contract method 0x2e0b437d.
//
// Solidity: function addModuleValidator(address validator, bytes initData) returns()
func (_SsoAccount *SsoAccountTransactor) AddModuleValidator(opts *bind.TransactOpts, validator common.Address, initData []byte) (*types.Transaction, error) {
	return _SsoAccount.contract.Transact(opts, "addModuleValidator", validator, initData)
}

// AddModuleValidator is a paid mutator transaction binding the contract method 0x2e0b437d.
//
// Solidity: function addModuleValidator(address validator, bytes initData) returns()
func (_SsoAccount *SsoAccountSession) AddModuleValidator(validator common.Address, initData []byte) (*types.Transaction, error) {
	return _SsoAccount.Contract.AddModuleValidator(&_SsoAccount.TransactOpts, validator, initData)
}

// AddModuleValidator is a paid mutator transaction binding the contract method 0x2e0b437d.
//
// Solidity: function addModuleValidator(address validator, bytes initData) returns()
func (_SsoAccount *SsoAccountTransactorSession) AddModuleValidator(validator common.Address, initData []byte) (*types.Transaction, error) {
	return _SsoAccount.Contract.AddModuleValidator(&_SsoAccount.TransactOpts, validator, initData)
}

// BatchCall is a paid mutator transaction binding the contract method 0x8f0273a9.
//
// Solidity: function batchCall((address,bool,uint256,bytes)[] _calls) payable returns()
func (_SsoAccount *SsoAccountTransactor) BatchCall(opts *bind.TransactOpts, _calls []Call) (*types.Transaction, error) {
	return _SsoAccount.contract.Transact(opts, "batchCall", _calls)
}

// BatchCall is a paid mutator transaction binding the contract method 0x8f0273a9.
//
// Solidity: function batchCall((address,bool,uint256,bytes)[] _calls) payable returns()
func (_SsoAccount *SsoAccountSession) BatchCall(_calls []Call) (*types.Transaction, error) {
	return _SsoAccount.Contract.BatchCall(&_SsoAccount.TransactOpts, _calls)
}

// BatchCall is a paid mutator transaction binding the contract method 0x8f0273a9.
//
// Solidity: function batchCall((address,bool,uint256,bytes)[] _calls) payable returns()
func (_SsoAccount *SsoAccountTransactorSession) BatchCall(_calls []Call) (*types.Transaction, error) {
	return _SsoAccount.Contract.BatchCall(&_SsoAccount.TransactOpts, _calls)
}

// ExecuteTransaction is a paid mutator transaction binding the contract method 0xdf9c1589.
//
// Solidity: function executeTransaction(bytes32 , bytes32 , (uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256[4],bytes,bytes,bytes32[],bytes,bytes) _transaction) payable returns()
func (_SsoAccount *SsoAccountTransactor) ExecuteTransaction(opts *bind.TransactOpts, arg0 [32]byte, arg1 [32]byte, _transaction Transaction) (*types.Transaction, error) {
	return _SsoAccount.contract.Transact(opts, "executeTransaction", arg0, arg1, _transaction)
}

// ExecuteTransaction is a paid mutator transaction binding the contract method 0xdf9c1589.
//
// Solidity: function executeTransaction(bytes32 , bytes32 , (uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256[4],bytes,bytes,bytes32[],bytes,bytes) _transaction) payable returns()
func (_SsoAccount *SsoAccountSession) ExecuteTransaction(arg0 [32]byte, arg1 [32]byte, _transaction Transaction) (*types.Transaction, error) {
	return _SsoAccount.Contract.ExecuteTransaction(&_SsoAccount.TransactOpts, arg0, arg1, _transaction)
}

// ExecuteTransaction is a paid mutator transaction binding the contract method 0xdf9c1589.
//
// Solidity: function executeTransaction(bytes32 , bytes32 , (uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256[4],bytes,bytes,bytes32[],bytes,bytes) _transaction) payable returns()
func (_SsoAccount *SsoAccountTransactorSession) ExecuteTransaction(arg0 [32]byte, arg1 [32]byte, _transaction Transaction) (*types.Transaction, error) {
	return _SsoAccount.Contract.ExecuteTransaction(&_SsoAccount.TransactOpts, arg0, arg1, _transaction)
}

// ExecuteTransactionFromOutside is a paid mutator transaction binding the contract method 0xeeb8cb09.
//
// Solidity: function executeTransactionFromOutside((uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256[4],bytes,bytes,bytes32[],bytes,bytes) ) payable returns()
func (_SsoAccount *SsoAccountTransactor) ExecuteTransactionFromOutside(opts *bind.TransactOpts, arg0 Transaction) (*types.Transaction, error) {
	return _SsoAccount.contract.Transact(opts, "executeTransactionFromOutside", arg0)
}

// ExecuteTransactionFromOutside is a paid mutator transaction binding the contract method 0xeeb8cb09.
//
// Solidity: function executeTransactionFromOutside((uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256[4],bytes,bytes,bytes32[],bytes,bytes) ) payable returns()
func (_SsoAccount *SsoAccountSession) ExecuteTransactionFromOutside(arg0 Transaction) (*types.Transaction, error) {
	return _SsoAccount.Contract.ExecuteTransactionFromOutside(&_SsoAccount.TransactOpts, arg0)
}

// ExecuteTransactionFromOutside is a paid mutator transaction binding the contract method 0xeeb8cb09.
//
// Solidity: function executeTransactionFromOutside((uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256[4],bytes,bytes,bytes32[],bytes,bytes) ) payable returns()
func (_SsoAccount *SsoAccountTransactorSession) ExecuteTransactionFromOutside(arg0 Transaction) (*types.Transaction, error) {
	return _SsoAccount.Contract.ExecuteTransactionFromOutside(&_SsoAccount.TransactOpts, arg0)
}

// Initialize is a paid mutator transaction binding the contract method 0xb9094997.
//
// Solidity: function initialize(bytes[] initialValidators, address[] initialK1Owners) returns()
func (_SsoAccount *SsoAccountTransactor) Initialize(opts *bind.TransactOpts, initialValidators [][]byte, initialK1Owners []common.Address) (*types.Transaction, error) {
	return _SsoAccount.contract.Transact(opts, "initialize", initialValidators, initialK1Owners)
}

// Initialize is a paid mutator transaction binding the contract method 0xb9094997.
//
// Solidity: function initialize(bytes[] initialValidators, address[] initialK1Owners) returns()
func (_SsoAccount *SsoAccountSession) Initialize(initialValidators [][]byte, initialK1Owners []common.Address) (*types.Transaction, error) {
	return _SsoAccount.Contract.Initialize(&_SsoAccount.TransactOpts, initialValidators, initialK1Owners)
}

// Initialize is a paid mutator transaction binding the contract method 0xb9094997.
//
// Solidity: function initialize(bytes[] initialValidators, address[] initialK1Owners) returns()
func (_SsoAccount *SsoAccountTransactorSession) Initialize(initialValidators [][]byte, initialK1Owners []common.Address) (*types.Transaction, error) {
	return _SsoAccount.Contract.Initialize(&_SsoAccount.TransactOpts, initialValidators, initialK1Owners)
}

// PayForTransaction is a paid mutator transaction binding the contract method 0xe2f318e3.
//
// Solidity: function payForTransaction(bytes32 , bytes32 , (uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256[4],bytes,bytes,bytes32[],bytes,bytes) _transaction) payable returns()
func (_SsoAccount *SsoAccountTransactor) PayForTransaction(opts *bind.TransactOpts, arg0 [32]byte, arg1 [32]byte, _transaction Transaction) (*types.Transaction, error) {
	return _SsoAccount.contract.Transact(opts, "payForTransaction", arg0, arg1, _transaction)
}

// PayForTransaction is a paid mutator transaction binding the contract method 0xe2f318e3.
//
// Solidity: function payForTransaction(bytes32 , bytes32 , (uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256[4],bytes,bytes,bytes32[],bytes,bytes) _transaction) payable returns()
func (_SsoAccount *SsoAccountSession) PayForTransaction(arg0 [32]byte, arg1 [32]byte, _transaction Transaction) (*types.Transaction, error) {
	return _SsoAccount.Contract.PayForTransaction(&_SsoAccount.TransactOpts, arg0, arg1, _transaction)
}

// PayForTransaction is a paid mutator transaction binding the contract method 0xe2f318e3.
//
// Solidity: function payForTransaction(bytes32 , bytes32 , (uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256[4],bytes,bytes,bytes32[],bytes,bytes) _transaction) payable returns()
func (_SsoAccount *SsoAccountTransactorSession) PayForTransaction(arg0 [32]byte, arg1 [32]byte, _transaction Transaction) (*types.Transaction, error) {
	return _SsoAccount.Contract.PayForTransaction(&_SsoAccount.TransactOpts, arg0, arg1, _transaction)
}

// PrepareForPaymaster is a paid mutator transaction binding the contract method 0xa28c1aee.
//
// Solidity: function prepareForPaymaster(bytes32 , bytes32 , (uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256[4],bytes,bytes,bytes32[],bytes,bytes) _transaction) payable returns()
func (_SsoAccount *SsoAccountTransactor) PrepareForPaymaster(opts *bind.TransactOpts, arg0 [32]byte, arg1 [32]byte, _transaction Transaction) (*types.Transaction, error) {
	return _SsoAccount.contract.Transact(opts, "prepareForPaymaster", arg0, arg1, _transaction)
}

// PrepareForPaymaster is a paid mutator transaction binding the contract method 0xa28c1aee.
//
// Solidity: function prepareForPaymaster(bytes32 , bytes32 , (uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256[4],bytes,bytes,bytes32[],bytes,bytes) _transaction) payable returns()
func (_SsoAccount *SsoAccountSession) PrepareForPaymaster(arg0 [32]byte, arg1 [32]byte, _transaction Transaction) (*types.Transaction, error) {
	return _SsoAccount.Contract.PrepareForPaymaster(&_SsoAccount.TransactOpts, arg0, arg1, _transaction)
}

// PrepareForPaymaster is a paid mutator transaction binding the contract method 0xa28c1aee.
//
// Solidity: function prepareForPaymaster(bytes32 , bytes32 , (uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256[4],bytes,bytes,bytes32[],bytes,bytes) _transaction) payable returns()
func (_SsoAccount *SsoAccountTransactorSession) PrepareForPaymaster(arg0 [32]byte, arg1 [32]byte, _transaction Transaction) (*types.Transaction, error) {
	return _SsoAccount.Contract.PrepareForPaymaster(&_SsoAccount.TransactOpts, arg0, arg1, _transaction)
}

// RemoveHook is a paid mutator transaction binding the contract method 0xe1551286.
//
// Solidity: function removeHook(address hook, bool isValidation, bytes deinitData) returns()
func (_SsoAccount *SsoAccountTransactor) RemoveHook(opts *bind.TransactOpts, hook common.Address, isValidation bool, deinitData []byte) (*types.Transaction, error) {
	return _SsoAccount.contract.Transact(opts, "removeHook", hook, isValidation, deinitData)
}

// RemoveHook is a paid mutator transaction binding the contract method 0xe1551286.
//
// Solidity: function removeHook(address hook, bool isValidation, bytes deinitData) returns()
func (_SsoAccount *SsoAccountSession) RemoveHook(hook common.Address, isValidation bool, deinitData []byte) (*types.Transaction, error) {
	return _SsoAccount.Contract.RemoveHook(&_SsoAccount.TransactOpts, hook, isValidation, deinitData)
}

// RemoveHook is a paid mutator transaction binding the contract method 0xe1551286.
//
// Solidity: function removeHook(address hook, bool isValidation, bytes deinitData) returns()
func (_SsoAccount *SsoAccountTransactorSession) RemoveHook(hook common.Address, isValidation bool, deinitData []byte) (*types.Transaction, error) {
	return _SsoAccount.Contract.RemoveHook(&_SsoAccount.TransactOpts, hook, isValidation, deinitData)
}

// RemoveK1Owner is a paid mutator transaction binding the contract method 0x714da018.
//
// Solidity: function removeK1Owner(address addr) returns()
func (_SsoAccount *SsoAccountTransactor) RemoveK1Owner(opts *bind.TransactOpts, addr common.Address) (*types.Transaction, error) {
	return _SsoAccount.contract.Transact(opts, "removeK1Owner", addr)
}

// RemoveK1Owner is a paid mutator transaction binding the contract method 0x714da018.
//
// Solidity: function removeK1Owner(address addr) returns()
func (_SsoAccount *SsoAccountSession) RemoveK1Owner(addr common.Address) (*types.Transaction, error) {
	return _SsoAccount.Contract.RemoveK1Owner(&_SsoAccount.TransactOpts, addr)
}

// RemoveK1Owner is a paid mutator transaction binding the contract method 0x714da018.
//
// Solidity: function removeK1Owner(address addr) returns()
func (_SsoAccount *SsoAccountTransactorSession) RemoveK1Owner(addr common.Address) (*types.Transaction, error) {
	return _SsoAccount.Contract.RemoveK1Owner(&_SsoAccount.TransactOpts, addr)
}

// RemoveModuleValidator is a paid mutator transaction binding the contract method 0xb027e50a.
//
// Solidity: function removeModuleValidator(address validator, bytes deinitData) returns()
func (_SsoAccount *SsoAccountTransactor) RemoveModuleValidator(opts *bind.TransactOpts, validator common.Address, deinitData []byte) (*types.Transaction, error) {
	return _SsoAccount.contract.Transact(opts, "removeModuleValidator", validator, deinitData)
}

// RemoveModuleValidator is a paid mutator transaction binding the contract method 0xb027e50a.
//
// Solidity: function removeModuleValidator(address validator, bytes deinitData) returns()
func (_SsoAccount *SsoAccountSession) RemoveModuleValidator(validator common.Address, deinitData []byte) (*types.Transaction, error) {
	return _SsoAccount.Contract.RemoveModuleValidator(&_SsoAccount.TransactOpts, validator, deinitData)
}

// RemoveModuleValidator is a paid mutator transaction binding the contract method 0xb027e50a.
//
// Solidity: function removeModuleValidator(address validator, bytes deinitData) returns()
func (_SsoAccount *SsoAccountTransactorSession) RemoveModuleValidator(validator common.Address, deinitData []byte) (*types.Transaction, error) {
	return _SsoAccount.Contract.RemoveModuleValidator(&_SsoAccount.TransactOpts, validator, deinitData)
}

// UnlinkHook is a paid mutator transaction binding the contract method 0x126c46c1.
//
// Solidity: function unlinkHook(address hook, bool isValidation, bytes deinitData) returns()
func (_SsoAccount *SsoAccountTransactor) UnlinkHook(opts *bind.TransactOpts, hook common.Address, isValidation bool, deinitData []byte) (*types.Transaction, error) {
	return _SsoAccount.contract.Transact(opts, "unlinkHook", hook, isValidation, deinitData)
}

// UnlinkHook is a paid mutator transaction binding the contract method 0x126c46c1.
//
// Solidity: function unlinkHook(address hook, bool isValidation, bytes deinitData) returns()
func (_SsoAccount *SsoAccountSession) UnlinkHook(hook common.Address, isValidation bool, deinitData []byte) (*types.Transaction, error) {
	return _SsoAccount.Contract.UnlinkHook(&_SsoAccount.TransactOpts, hook, isValidation, deinitData)
}

// UnlinkHook is a paid mutator transaction binding the contract method 0x126c46c1.
//
// Solidity: function unlinkHook(address hook, bool isValidation, bytes deinitData) returns()
func (_SsoAccount *SsoAccountTransactorSession) UnlinkHook(hook common.Address, isValidation bool, deinitData []byte) (*types.Transaction, error) {
	return _SsoAccount.Contract.UnlinkHook(&_SsoAccount.TransactOpts, hook, isValidation, deinitData)
}

// UnlinkModuleValidator is a paid mutator transaction binding the contract method 0xa7d525e8.
//
// Solidity: function unlinkModuleValidator(address validator, bytes deinitData) returns()
func (_SsoAccount *SsoAccountTransactor) UnlinkModuleValidator(opts *bind.TransactOpts, validator common.Address, deinitData []byte) (*types.Transaction, error) {
	return _SsoAccount.contract.Transact(opts, "unlinkModuleValidator", validator, deinitData)
}

// UnlinkModuleValidator is a paid mutator transaction binding the contract method 0xa7d525e8.
//
// Solidity: function unlinkModuleValidator(address validator, bytes deinitData) returns()
func (_SsoAccount *SsoAccountSession) UnlinkModuleValidator(validator common.Address, deinitData []byte) (*types.Transaction, error) {
	return _SsoAccount.Contract.UnlinkModuleValidator(&_SsoAccount.TransactOpts, validator, deinitData)
}

// UnlinkModuleValidator is a paid mutator transaction binding the contract method 0xa7d525e8.
//
// Solidity: function unlinkModuleValidator(address validator, bytes deinitData) returns()
func (_SsoAccount *SsoAccountTransactorSession) UnlinkModuleValidator(validator common.Address, deinitData []byte) (*types.Transaction, error) {
	return _SsoAccount.Contract.UnlinkModuleValidator(&_SsoAccount.TransactOpts, validator, deinitData)
}

// ValidateTransaction is a paid mutator transaction binding the contract method 0x202bcce7.
//
// Solidity: function validateTransaction(bytes32 , bytes32 _suggestedSignedHash, (uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256[4],bytes,bytes,bytes32[],bytes,bytes) _transaction) payable returns(bytes4 magic)
func (_SsoAccount *SsoAccountTransactor) ValidateTransaction(opts *bind.TransactOpts, arg0 [32]byte, _suggestedSignedHash [32]byte, _transaction Transaction) (*types.Transaction, error) {
	return _SsoAccount.contract.Transact(opts, "validateTransaction", arg0, _suggestedSignedHash, _transaction)
}

// ValidateTransaction is a paid mutator transaction binding the contract method 0x202bcce7.
//
// Solidity: function validateTransaction(bytes32 , bytes32 _suggestedSignedHash, (uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256[4],bytes,bytes,bytes32[],bytes,bytes) _transaction) payable returns(bytes4 magic)
func (_SsoAccount *SsoAccountSession) ValidateTransaction(arg0 [32]byte, _suggestedSignedHash [32]byte, _transaction Transaction) (*types.Transaction, error) {
	return _SsoAccount.Contract.ValidateTransaction(&_SsoAccount.TransactOpts, arg0, _suggestedSignedHash, _transaction)
}

// ValidateTransaction is a paid mutator transaction binding the contract method 0x202bcce7.
//
// Solidity: function validateTransaction(bytes32 , bytes32 _suggestedSignedHash, (uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256[4],bytes,bytes,bytes32[],bytes,bytes) _transaction) payable returns(bytes4 magic)
func (_SsoAccount *SsoAccountTransactorSession) ValidateTransaction(arg0 [32]byte, _suggestedSignedHash [32]byte, _transaction Transaction) (*types.Transaction, error) {
	return _SsoAccount.Contract.ValidateTransaction(&_SsoAccount.TransactOpts, arg0, _suggestedSignedHash, _transaction)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_SsoAccount *SsoAccountTransactor) Receive(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _SsoAccount.contract.RawTransact(opts, nil) // calldata is disallowed for receive function
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_SsoAccount *SsoAccountSession) Receive() (*types.Transaction, error) {
	return _SsoAccount.Contract.Receive(&_SsoAccount.TransactOpts)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_SsoAccount *SsoAccountTransactorSession) Receive() (*types.Transaction, error) {
	return _SsoAccount.Contract.Receive(&_SsoAccount.TransactOpts)
}

// SsoAccountBatchCallFailureIterator is returned from FilterBatchCallFailure and is used to iterate over the raw logs and unpacked data for BatchCallFailure events raised by the SsoAccount contract.
type SsoAccountBatchCallFailureIterator struct {
	Event *SsoAccountBatchCallFailure // Event containing the contract specifics and raw log

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
func (it *SsoAccountBatchCallFailureIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SsoAccountBatchCallFailure)
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
		it.Event = new(SsoAccountBatchCallFailure)
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
func (it *SsoAccountBatchCallFailureIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SsoAccountBatchCallFailureIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SsoAccountBatchCallFailure represents a BatchCallFailure event raised by the SsoAccount contract.
type SsoAccountBatchCallFailure struct {
	Index      *big.Int
	RevertData []byte
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterBatchCallFailure is a free log retrieval operation binding the contract event 0x860d60a3c3ed451a144846eba67081eb134233fed0aff20997e5110b942b3d0e.
//
// Solidity: event BatchCallFailure(uint256 indexed index, bytes revertData)
func (_SsoAccount *SsoAccountFilterer) FilterBatchCallFailure(opts *bind.FilterOpts, index []*big.Int) (*SsoAccountBatchCallFailureIterator, error) {

	var indexRule []interface{}
	for _, indexItem := range index {
		indexRule = append(indexRule, indexItem)
	}

	logs, sub, err := _SsoAccount.contract.FilterLogs(opts, "BatchCallFailure", indexRule)
	if err != nil {
		return nil, err
	}
	return &SsoAccountBatchCallFailureIterator{contract: _SsoAccount.contract, event: "BatchCallFailure", logs: logs, sub: sub}, nil
}

// WatchBatchCallFailure is a free log subscription operation binding the contract event 0x860d60a3c3ed451a144846eba67081eb134233fed0aff20997e5110b942b3d0e.
//
// Solidity: event BatchCallFailure(uint256 indexed index, bytes revertData)
func (_SsoAccount *SsoAccountFilterer) WatchBatchCallFailure(opts *bind.WatchOpts, sink chan<- *SsoAccountBatchCallFailure, index []*big.Int) (event.Subscription, error) {

	var indexRule []interface{}
	for _, indexItem := range index {
		indexRule = append(indexRule, indexItem)
	}

	logs, sub, err := _SsoAccount.contract.WatchLogs(opts, "BatchCallFailure", indexRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SsoAccountBatchCallFailure)
				if err := _SsoAccount.contract.UnpackLog(event, "BatchCallFailure", log); err != nil {
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

// ParseBatchCallFailure is a log parse operation binding the contract event 0x860d60a3c3ed451a144846eba67081eb134233fed0aff20997e5110b942b3d0e.
//
// Solidity: event BatchCallFailure(uint256 indexed index, bytes revertData)
func (_SsoAccount *SsoAccountFilterer) ParseBatchCallFailure(log types.Log) (*SsoAccountBatchCallFailure, error) {
	event := new(SsoAccountBatchCallFailure)
	if err := _SsoAccount.contract.UnpackLog(event, "BatchCallFailure", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SsoAccountHookAddedIterator is returned from FilterHookAdded and is used to iterate over the raw logs and unpacked data for HookAdded events raised by the SsoAccount contract.
type SsoAccountHookAddedIterator struct {
	Event *SsoAccountHookAdded // Event containing the contract specifics and raw log

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
func (it *SsoAccountHookAddedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SsoAccountHookAdded)
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
		it.Event = new(SsoAccountHookAdded)
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
func (it *SsoAccountHookAddedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SsoAccountHookAddedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SsoAccountHookAdded represents a HookAdded event raised by the SsoAccount contract.
type SsoAccountHookAdded struct {
	Hook common.Address
	Raw  types.Log // Blockchain specific contextual infos
}

// FilterHookAdded is a free log retrieval operation binding the contract event 0x28e00134722f84e69c391c81e4fe022ee3e61048222a8ea2f98c9f235f797508.
//
// Solidity: event HookAdded(address indexed hook)
func (_SsoAccount *SsoAccountFilterer) FilterHookAdded(opts *bind.FilterOpts, hook []common.Address) (*SsoAccountHookAddedIterator, error) {

	var hookRule []interface{}
	for _, hookItem := range hook {
		hookRule = append(hookRule, hookItem)
	}

	logs, sub, err := _SsoAccount.contract.FilterLogs(opts, "HookAdded", hookRule)
	if err != nil {
		return nil, err
	}
	return &SsoAccountHookAddedIterator{contract: _SsoAccount.contract, event: "HookAdded", logs: logs, sub: sub}, nil
}

// WatchHookAdded is a free log subscription operation binding the contract event 0x28e00134722f84e69c391c81e4fe022ee3e61048222a8ea2f98c9f235f797508.
//
// Solidity: event HookAdded(address indexed hook)
func (_SsoAccount *SsoAccountFilterer) WatchHookAdded(opts *bind.WatchOpts, sink chan<- *SsoAccountHookAdded, hook []common.Address) (event.Subscription, error) {

	var hookRule []interface{}
	for _, hookItem := range hook {
		hookRule = append(hookRule, hookItem)
	}

	logs, sub, err := _SsoAccount.contract.WatchLogs(opts, "HookAdded", hookRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SsoAccountHookAdded)
				if err := _SsoAccount.contract.UnpackLog(event, "HookAdded", log); err != nil {
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

// ParseHookAdded is a log parse operation binding the contract event 0x28e00134722f84e69c391c81e4fe022ee3e61048222a8ea2f98c9f235f797508.
//
// Solidity: event HookAdded(address indexed hook)
func (_SsoAccount *SsoAccountFilterer) ParseHookAdded(log types.Log) (*SsoAccountHookAdded, error) {
	event := new(SsoAccountHookAdded)
	if err := _SsoAccount.contract.UnpackLog(event, "HookAdded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SsoAccountHookRemovedIterator is returned from FilterHookRemoved and is used to iterate over the raw logs and unpacked data for HookRemoved events raised by the SsoAccount contract.
type SsoAccountHookRemovedIterator struct {
	Event *SsoAccountHookRemoved // Event containing the contract specifics and raw log

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
func (it *SsoAccountHookRemovedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SsoAccountHookRemoved)
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
		it.Event = new(SsoAccountHookRemoved)
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
func (it *SsoAccountHookRemovedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SsoAccountHookRemovedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SsoAccountHookRemoved represents a HookRemoved event raised by the SsoAccount contract.
type SsoAccountHookRemoved struct {
	Hook common.Address
	Raw  types.Log // Blockchain specific contextual infos
}

// FilterHookRemoved is a free log retrieval operation binding the contract event 0x47d0871e905ac6550f54ba266e0d90d2dc8ed67a957c064ca3438eddf4e3fd89.
//
// Solidity: event HookRemoved(address indexed hook)
func (_SsoAccount *SsoAccountFilterer) FilterHookRemoved(opts *bind.FilterOpts, hook []common.Address) (*SsoAccountHookRemovedIterator, error) {

	var hookRule []interface{}
	for _, hookItem := range hook {
		hookRule = append(hookRule, hookItem)
	}

	logs, sub, err := _SsoAccount.contract.FilterLogs(opts, "HookRemoved", hookRule)
	if err != nil {
		return nil, err
	}
	return &SsoAccountHookRemovedIterator{contract: _SsoAccount.contract, event: "HookRemoved", logs: logs, sub: sub}, nil
}

// WatchHookRemoved is a free log subscription operation binding the contract event 0x47d0871e905ac6550f54ba266e0d90d2dc8ed67a957c064ca3438eddf4e3fd89.
//
// Solidity: event HookRemoved(address indexed hook)
func (_SsoAccount *SsoAccountFilterer) WatchHookRemoved(opts *bind.WatchOpts, sink chan<- *SsoAccountHookRemoved, hook []common.Address) (event.Subscription, error) {

	var hookRule []interface{}
	for _, hookItem := range hook {
		hookRule = append(hookRule, hookItem)
	}

	logs, sub, err := _SsoAccount.contract.WatchLogs(opts, "HookRemoved", hookRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SsoAccountHookRemoved)
				if err := _SsoAccount.contract.UnpackLog(event, "HookRemoved", log); err != nil {
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

// ParseHookRemoved is a log parse operation binding the contract event 0x47d0871e905ac6550f54ba266e0d90d2dc8ed67a957c064ca3438eddf4e3fd89.
//
// Solidity: event HookRemoved(address indexed hook)
func (_SsoAccount *SsoAccountFilterer) ParseHookRemoved(log types.Log) (*SsoAccountHookRemoved, error) {
	event := new(SsoAccountHookRemoved)
	if err := _SsoAccount.contract.UnpackLog(event, "HookRemoved", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SsoAccountInitializedIterator is returned from FilterInitialized and is used to iterate over the raw logs and unpacked data for Initialized events raised by the SsoAccount contract.
type SsoAccountInitializedIterator struct {
	Event *SsoAccountInitialized // Event containing the contract specifics and raw log

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
func (it *SsoAccountInitializedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SsoAccountInitialized)
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
		it.Event = new(SsoAccountInitialized)
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
func (it *SsoAccountInitializedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SsoAccountInitializedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SsoAccountInitialized represents a Initialized event raised by the SsoAccount contract.
type SsoAccountInitialized struct {
	Version uint8
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterInitialized is a free log retrieval operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_SsoAccount *SsoAccountFilterer) FilterInitialized(opts *bind.FilterOpts) (*SsoAccountInitializedIterator, error) {

	logs, sub, err := _SsoAccount.contract.FilterLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return &SsoAccountInitializedIterator{contract: _SsoAccount.contract, event: "Initialized", logs: logs, sub: sub}, nil
}

// WatchInitialized is a free log subscription operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_SsoAccount *SsoAccountFilterer) WatchInitialized(opts *bind.WatchOpts, sink chan<- *SsoAccountInitialized) (event.Subscription, error) {

	logs, sub, err := _SsoAccount.contract.WatchLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SsoAccountInitialized)
				if err := _SsoAccount.contract.UnpackLog(event, "Initialized", log); err != nil {
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
func (_SsoAccount *SsoAccountFilterer) ParseInitialized(log types.Log) (*SsoAccountInitialized, error) {
	event := new(SsoAccountInitialized)
	if err := _SsoAccount.contract.UnpackLog(event, "Initialized", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SsoAccountK1OwnerAddedIterator is returned from FilterK1OwnerAdded and is used to iterate over the raw logs and unpacked data for K1OwnerAdded events raised by the SsoAccount contract.
type SsoAccountK1OwnerAddedIterator struct {
	Event *SsoAccountK1OwnerAdded // Event containing the contract specifics and raw log

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
func (it *SsoAccountK1OwnerAddedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SsoAccountK1OwnerAdded)
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
		it.Event = new(SsoAccountK1OwnerAdded)
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
func (it *SsoAccountK1OwnerAddedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SsoAccountK1OwnerAddedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SsoAccountK1OwnerAdded represents a K1OwnerAdded event raised by the SsoAccount contract.
type SsoAccountK1OwnerAdded struct {
	Addr common.Address
	Raw  types.Log // Blockchain specific contextual infos
}

// FilterK1OwnerAdded is a free log retrieval operation binding the contract event 0x40097f116787d282c09163d089084b2d950545433dad84146baf8da81064e84c.
//
// Solidity: event K1OwnerAdded(address indexed addr)
func (_SsoAccount *SsoAccountFilterer) FilterK1OwnerAdded(opts *bind.FilterOpts, addr []common.Address) (*SsoAccountK1OwnerAddedIterator, error) {

	var addrRule []interface{}
	for _, addrItem := range addr {
		addrRule = append(addrRule, addrItem)
	}

	logs, sub, err := _SsoAccount.contract.FilterLogs(opts, "K1OwnerAdded", addrRule)
	if err != nil {
		return nil, err
	}
	return &SsoAccountK1OwnerAddedIterator{contract: _SsoAccount.contract, event: "K1OwnerAdded", logs: logs, sub: sub}, nil
}

// WatchK1OwnerAdded is a free log subscription operation binding the contract event 0x40097f116787d282c09163d089084b2d950545433dad84146baf8da81064e84c.
//
// Solidity: event K1OwnerAdded(address indexed addr)
func (_SsoAccount *SsoAccountFilterer) WatchK1OwnerAdded(opts *bind.WatchOpts, sink chan<- *SsoAccountK1OwnerAdded, addr []common.Address) (event.Subscription, error) {

	var addrRule []interface{}
	for _, addrItem := range addr {
		addrRule = append(addrRule, addrItem)
	}

	logs, sub, err := _SsoAccount.contract.WatchLogs(opts, "K1OwnerAdded", addrRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SsoAccountK1OwnerAdded)
				if err := _SsoAccount.contract.UnpackLog(event, "K1OwnerAdded", log); err != nil {
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

// ParseK1OwnerAdded is a log parse operation binding the contract event 0x40097f116787d282c09163d089084b2d950545433dad84146baf8da81064e84c.
//
// Solidity: event K1OwnerAdded(address indexed addr)
func (_SsoAccount *SsoAccountFilterer) ParseK1OwnerAdded(log types.Log) (*SsoAccountK1OwnerAdded, error) {
	event := new(SsoAccountK1OwnerAdded)
	if err := _SsoAccount.contract.UnpackLog(event, "K1OwnerAdded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SsoAccountK1OwnerRemovedIterator is returned from FilterK1OwnerRemoved and is used to iterate over the raw logs and unpacked data for K1OwnerRemoved events raised by the SsoAccount contract.
type SsoAccountK1OwnerRemovedIterator struct {
	Event *SsoAccountK1OwnerRemoved // Event containing the contract specifics and raw log

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
func (it *SsoAccountK1OwnerRemovedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SsoAccountK1OwnerRemoved)
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
		it.Event = new(SsoAccountK1OwnerRemoved)
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
func (it *SsoAccountK1OwnerRemovedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SsoAccountK1OwnerRemovedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SsoAccountK1OwnerRemoved represents a K1OwnerRemoved event raised by the SsoAccount contract.
type SsoAccountK1OwnerRemoved struct {
	Addr common.Address
	Raw  types.Log // Blockchain specific contextual infos
}

// FilterK1OwnerRemoved is a free log retrieval operation binding the contract event 0xe09f0b842a072d50a0b373f3255e2a4216aeb61908e35a93a499f0ff2aa4c8fa.
//
// Solidity: event K1OwnerRemoved(address indexed addr)
func (_SsoAccount *SsoAccountFilterer) FilterK1OwnerRemoved(opts *bind.FilterOpts, addr []common.Address) (*SsoAccountK1OwnerRemovedIterator, error) {

	var addrRule []interface{}
	for _, addrItem := range addr {
		addrRule = append(addrRule, addrItem)
	}

	logs, sub, err := _SsoAccount.contract.FilterLogs(opts, "K1OwnerRemoved", addrRule)
	if err != nil {
		return nil, err
	}
	return &SsoAccountK1OwnerRemovedIterator{contract: _SsoAccount.contract, event: "K1OwnerRemoved", logs: logs, sub: sub}, nil
}

// WatchK1OwnerRemoved is a free log subscription operation binding the contract event 0xe09f0b842a072d50a0b373f3255e2a4216aeb61908e35a93a499f0ff2aa4c8fa.
//
// Solidity: event K1OwnerRemoved(address indexed addr)
func (_SsoAccount *SsoAccountFilterer) WatchK1OwnerRemoved(opts *bind.WatchOpts, sink chan<- *SsoAccountK1OwnerRemoved, addr []common.Address) (event.Subscription, error) {

	var addrRule []interface{}
	for _, addrItem := range addr {
		addrRule = append(addrRule, addrItem)
	}

	logs, sub, err := _SsoAccount.contract.WatchLogs(opts, "K1OwnerRemoved", addrRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SsoAccountK1OwnerRemoved)
				if err := _SsoAccount.contract.UnpackLog(event, "K1OwnerRemoved", log); err != nil {
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

// ParseK1OwnerRemoved is a log parse operation binding the contract event 0xe09f0b842a072d50a0b373f3255e2a4216aeb61908e35a93a499f0ff2aa4c8fa.
//
// Solidity: event K1OwnerRemoved(address indexed addr)
func (_SsoAccount *SsoAccountFilterer) ParseK1OwnerRemoved(log types.Log) (*SsoAccountK1OwnerRemoved, error) {
	event := new(SsoAccountK1OwnerRemoved)
	if err := _SsoAccount.contract.UnpackLog(event, "K1OwnerRemoved", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SsoAccountValidatorAddedIterator is returned from FilterValidatorAdded and is used to iterate over the raw logs and unpacked data for ValidatorAdded events raised by the SsoAccount contract.
type SsoAccountValidatorAddedIterator struct {
	Event *SsoAccountValidatorAdded // Event containing the contract specifics and raw log

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
func (it *SsoAccountValidatorAddedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SsoAccountValidatorAdded)
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
		it.Event = new(SsoAccountValidatorAdded)
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
func (it *SsoAccountValidatorAddedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SsoAccountValidatorAddedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SsoAccountValidatorAdded represents a ValidatorAdded event raised by the SsoAccount contract.
type SsoAccountValidatorAdded struct {
	Validator common.Address
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterValidatorAdded is a free log retrieval operation binding the contract event 0xe366c1c0452ed8eec96861e9e54141ebff23c9ec89fe27b996b45f5ec3884987.
//
// Solidity: event ValidatorAdded(address indexed validator)
func (_SsoAccount *SsoAccountFilterer) FilterValidatorAdded(opts *bind.FilterOpts, validator []common.Address) (*SsoAccountValidatorAddedIterator, error) {

	var validatorRule []interface{}
	for _, validatorItem := range validator {
		validatorRule = append(validatorRule, validatorItem)
	}

	logs, sub, err := _SsoAccount.contract.FilterLogs(opts, "ValidatorAdded", validatorRule)
	if err != nil {
		return nil, err
	}
	return &SsoAccountValidatorAddedIterator{contract: _SsoAccount.contract, event: "ValidatorAdded", logs: logs, sub: sub}, nil
}

// WatchValidatorAdded is a free log subscription operation binding the contract event 0xe366c1c0452ed8eec96861e9e54141ebff23c9ec89fe27b996b45f5ec3884987.
//
// Solidity: event ValidatorAdded(address indexed validator)
func (_SsoAccount *SsoAccountFilterer) WatchValidatorAdded(opts *bind.WatchOpts, sink chan<- *SsoAccountValidatorAdded, validator []common.Address) (event.Subscription, error) {

	var validatorRule []interface{}
	for _, validatorItem := range validator {
		validatorRule = append(validatorRule, validatorItem)
	}

	logs, sub, err := _SsoAccount.contract.WatchLogs(opts, "ValidatorAdded", validatorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SsoAccountValidatorAdded)
				if err := _SsoAccount.contract.UnpackLog(event, "ValidatorAdded", log); err != nil {
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

// ParseValidatorAdded is a log parse operation binding the contract event 0xe366c1c0452ed8eec96861e9e54141ebff23c9ec89fe27b996b45f5ec3884987.
//
// Solidity: event ValidatorAdded(address indexed validator)
func (_SsoAccount *SsoAccountFilterer) ParseValidatorAdded(log types.Log) (*SsoAccountValidatorAdded, error) {
	event := new(SsoAccountValidatorAdded)
	if err := _SsoAccount.contract.UnpackLog(event, "ValidatorAdded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SsoAccountValidatorRemovedIterator is returned from FilterValidatorRemoved and is used to iterate over the raw logs and unpacked data for ValidatorRemoved events raised by the SsoAccount contract.
type SsoAccountValidatorRemovedIterator struct {
	Event *SsoAccountValidatorRemoved // Event containing the contract specifics and raw log

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
func (it *SsoAccountValidatorRemovedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SsoAccountValidatorRemoved)
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
		it.Event = new(SsoAccountValidatorRemoved)
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
func (it *SsoAccountValidatorRemovedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SsoAccountValidatorRemovedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SsoAccountValidatorRemoved represents a ValidatorRemoved event raised by the SsoAccount contract.
type SsoAccountValidatorRemoved struct {
	Validator common.Address
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterValidatorRemoved is a free log retrieval operation binding the contract event 0xe1434e25d6611e0db941968fdc97811c982ac1602e951637d206f5fdda9dd8f1.
//
// Solidity: event ValidatorRemoved(address indexed validator)
func (_SsoAccount *SsoAccountFilterer) FilterValidatorRemoved(opts *bind.FilterOpts, validator []common.Address) (*SsoAccountValidatorRemovedIterator, error) {

	var validatorRule []interface{}
	for _, validatorItem := range validator {
		validatorRule = append(validatorRule, validatorItem)
	}

	logs, sub, err := _SsoAccount.contract.FilterLogs(opts, "ValidatorRemoved", validatorRule)
	if err != nil {
		return nil, err
	}
	return &SsoAccountValidatorRemovedIterator{contract: _SsoAccount.contract, event: "ValidatorRemoved", logs: logs, sub: sub}, nil
}

// WatchValidatorRemoved is a free log subscription operation binding the contract event 0xe1434e25d6611e0db941968fdc97811c982ac1602e951637d206f5fdda9dd8f1.
//
// Solidity: event ValidatorRemoved(address indexed validator)
func (_SsoAccount *SsoAccountFilterer) WatchValidatorRemoved(opts *bind.WatchOpts, sink chan<- *SsoAccountValidatorRemoved, validator []common.Address) (event.Subscription, error) {

	var validatorRule []interface{}
	for _, validatorItem := range validator {
		validatorRule = append(validatorRule, validatorItem)
	}

	logs, sub, err := _SsoAccount.contract.WatchLogs(opts, "ValidatorRemoved", validatorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SsoAccountValidatorRemoved)
				if err := _SsoAccount.contract.UnpackLog(event, "ValidatorRemoved", log); err != nil {
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

// ParseValidatorRemoved is a log parse operation binding the contract event 0xe1434e25d6611e0db941968fdc97811c982ac1602e951637d206f5fdda9dd8f1.
//
// Solidity: event ValidatorRemoved(address indexed validator)
func (_SsoAccount *SsoAccountFilterer) ParseValidatorRemoved(log types.Log) (*SsoAccountValidatorRemoved, error) {
	event := new(SsoAccountValidatorRemoved)
	if err := _SsoAccount.contract.UnpackLog(event, "ValidatorRemoved", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
