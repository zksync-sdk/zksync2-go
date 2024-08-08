package utils

import (
	"crypto/sha256"
	"errors"
	"fmt"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/zksync-sdk/zksync2-go/contracts/contractdeployer"
	"github.com/zksync-sdk/zksync2-go/types"
	"math/big"
	"strings"
)

var (
	// EthAddress The address of the L1 ETH token.
	EthAddress = common.HexToAddress("0x0000000000000000000000000000000000000000")
	// LegacyEthAddress The address of the L1 ETH token.
	LegacyEthAddress = common.HexToAddress("0x0000000000000000000000000000000000000000")
	// EthAddressInContracts In the contracts the zero address can not be used, use one instead.
	EthAddressInContracts = common.HexToAddress("0x0000000000000000000000000000000000000001")
	// BootloaderFormalAddress The formal address for the Bootloader.
	BootloaderFormalAddress = common.HexToAddress("0x0000000000000000000000000000000000008001")
	// ContractDeployerAddress The address of the Contract deployer.
	ContractDeployerAddress = common.HexToAddress("0x0000000000000000000000000000000000008006")
	// L1MessengerAddress The address of the L1 messenger.
	L1MessengerAddress = common.HexToAddress("0x0000000000000000000000000000000000008008")
	// L2BaseTokenAddress The address of the base token.
	L2BaseTokenAddress = common.HexToAddress("0x000000000000000000000000000000000000800a")
	// NonceHolderAddress The address of the Nonce holder.
	NonceHolderAddress = common.HexToAddress("0x0000000000000000000000000000000000008003")

	// L1ToL2AliasOffset Used for applying and undoing aliases on contract addresses during bridging from L1 to L2.
	L1ToL2AliasOffset = common.HexToAddress("0x1111000000000000000000000000000000001111")
	AddressModulo     = new(big.Int).Exp(big.NewInt(2), big.NewInt(160), nil)

	Eip1271MagicValue = [4]byte(common.FromHex("0x1626ba7e"))
)

var contractDeployerABI *abi.ABI

// ApplyL1ToL2Alias converts the address of smart contract that submitted a transaction to the inbox on L1 to the
// `msg.sender` viewed on L2.
func ApplyL1ToL2Alias(address common.Address) common.Address {
	result := new(big.Int).Add(new(big.Int).SetBytes(address.Bytes()), new(big.Int).SetBytes(L1ToL2AliasOffset.Bytes()))
	result.Mod(result, AddressModulo)
	return common.BytesToAddress(result.Bytes())
}

// UndoL1ToL2Alias converts the address of smart contract that submitted a transaction to the inbox on L2 to the
// `msg.sender` viewed on L1.
func UndoL1ToL2Alias(address common.Address) common.Address {
	result := new(big.Int).Sub(new(big.Int).SetBytes(address.Bytes()), new(big.Int).SetBytes(L1ToL2AliasOffset.Bytes()))
	if result.Sign() < 0 {
		result = result.Add(result, AddressModulo)
	}
	return common.BytesToAddress(result.Bytes())
}

func getContractDeployerABI() (*abi.ABI, error) {
	if contractDeployerABI == nil {
		cda, err := abi.JSON(strings.NewReader(contractdeployer.IContractDeployerMetaData.ABI))
		if err != nil {
			return nil, fmt.Errorf("failed to load Deployer ABI: %w", err)
		}
		contractDeployerABI = &cda
	}
	return contractDeployerABI, nil
}

// Create2Address generates a future-proof contract address using salt plus bytecode which allows determination
// of an address before deployment.
func Create2Address(sender common.Address, bytecode, constructor, salt []byte) (common.Address, error) {
	if len(salt) == 0 {
		salt = make([]byte, 32)
	} else if len(salt) != 32 {
		return common.Address{}, errors.New("salt must be 32 bytes")
	}
	if constructor == nil {
		constructor = []byte{}
	}
	bytecodeHash, err := HashBytecode(bytecode)
	if err != nil {
		return common.Address{}, fmt.Errorf("failed to get hash of bytecode: %w", err)
	}
	bytes := make([]byte, 5*32)                                  // concatenate five 32-bytes slices
	copy(bytes[0:32], crypto.Keccak256([]byte("zksyncCreate2"))) // 1 - prefix
	copy(bytes[44:64], sender[:])                                // 2 - sender (20 bytes right padded to 32)
	copy(bytes[64:96], salt)                                     // 3 - salt
	copy(bytes[96:128], bytecodeHash)                            // 4 - bytecode hash
	copy(bytes[128:160], crypto.Keccak256(constructor))          // 5 - constructor hash
	result := crypto.Keccak256(bytes)
	return common.BytesToAddress(result[12:]), nil
}

// CreateAddress generates a contract address from deployer's account and nonce.
func CreateAddress(sender common.Address, nonce *big.Int) (common.Address, error) {
	nonceBytes := nonce.Bytes()
	bytes := make([]byte, 3*32)                                 // concatenate three 32-bytes slices
	copy(bytes[0:32], crypto.Keccak256([]byte("zksyncCreate"))) // 1 - prefix
	copy(bytes[44:64], sender[:])                               // 2 - sender (20 bytes right padded to 32)
	copy(bytes[64+(32-len(nonceBytes)):96], nonceBytes)         // 3 - nonce (some bytes right padded to 32)
	result := crypto.Keccak256(bytes)
	return common.BytesToAddress(result[12:]), nil
}

// EncodeCreate2 returns the encoded constructor data for CREATE2 method used for smart contract deployment.
func EncodeCreate2(bytecode, calldata, salt []byte) ([]byte, error) {
	cdABI, err := getContractDeployerABI()
	if err != nil {
		return nil, err
	}
	// prepare
	if len(salt) == 0 {
		salt = make([]byte, 32)
	} else if len(salt) != 32 {
		return nil, errors.New("salt must be 32 bytes")
	}
	hash, err := HashBytecode(bytecode)
	if err != nil {
		return nil, fmt.Errorf("failed to get hash of bytecode: %w", err)
	}
	salt32 := [32]byte{}
	copy(salt32[:], salt)
	hash32 := [32]byte{}
	copy(hash32[:], hash)

	res, err := cdABI.Pack("create2", salt32, hash32, calldata)
	if err != nil {
		return nil, fmt.Errorf("failed to pack create2 function: %w", err)
	}
	return res, nil
}

// EncodeCreate encodes the constructor data for CREATE method used for smart contract deployment.
func EncodeCreate(bytecode, calldata []byte) ([]byte, error) {
	cdABI, err := getContractDeployerABI()
	if err != nil {
		return nil, err
	}
	hash, err := HashBytecode(bytecode)
	if err != nil {
		return nil, fmt.Errorf("failed to get hash of bytecode: %w", err)
	}
	salt32 := [32]byte{} // will be empty
	hash32 := [32]byte{}
	copy(hash32[:], hash)

	res, err := cdABI.Pack("create", salt32, hash32, calldata)
	if err != nil {
		return nil, fmt.Errorf("failed to pack create function: %w", err)
	}
	return res, nil
}

// EncodeCreate2Account encodes the constructor data for CREATE2 method used for smart account deployment.
func EncodeCreate2Account(bytecode, calldata, salt []byte, version types.AccountAbstractionVersion) ([]byte, error) {
	cdABI, err := getContractDeployerABI()
	if err != nil {
		return nil, err
	}
	// prepare
	if len(salt) == 0 {
		salt = make([]byte, 32)
	} else if len(salt) != 32 {
		return nil, errors.New("salt must be 32 bytes")
	}
	hash, err := HashBytecode(bytecode)
	if err != nil {
		return nil, fmt.Errorf("failed to get hash of bytecode: %w", err)
	}
	salt32 := [32]byte{}
	copy(salt32[:], salt)
	hash32 := [32]byte{}
	copy(hash32[:], hash)

	res, err := cdABI.Pack("create2Account", salt32, hash32, calldata, version)
	if err != nil {
		return nil, fmt.Errorf("failed to pack create2Account function: %w", err)
	}
	return res, nil
}

// EncodeCreateAccount encodes the constructor data for CREATE method used for smart account deployment.
func EncodeCreateAccount(bytecode, calldata []byte, version types.AccountAbstractionVersion) ([]byte, error) {
	cdABI, err := getContractDeployerABI()
	if err != nil {
		return nil, err
	}
	hash, err := HashBytecode(bytecode)
	if err != nil {
		return nil, fmt.Errorf("failed to get hash of bytecode: %w", err)
	}
	salt32 := [32]byte{} // will be empty
	hash32 := [32]byte{}
	copy(hash32[:], hash)

	res, err := cdABI.Pack("createAccount", salt32, hash32, calldata, version)
	if err != nil {
		return nil, fmt.Errorf("failed to pack createAccount function: %w", err)
	}
	return res, nil
}

// HashBytecode returns the hash of given bytecode.
func HashBytecode(bytecode []byte) ([]byte, error) {
	if len(bytecode)%32 != 0 {
		return nil, errors.New("bytecode length in bytes must be divisible by 32")
	}
	bytecodeHash := sha256.Sum256(bytecode)
	// get real length of bytecode, which is presented as 32-byte words
	length := big.NewInt(int64(len(bytecode) / 32))
	if length.BitLen() > 16 {
		return nil, errors.New("bytecode length must be less than 2^16 bytes")
	}
	// replace first 2 bytes of hash with version
	version := []byte{1, 0}
	copy(bytecodeHash[0:2], version)
	// replace second 2 bytes of hash with bytecode length
	length2b := make([]byte, 2)
	length2b = length.FillBytes(length2b) // 0-padded in 2 bytes
	copy(bytecodeHash[2:4], length2b)
	return bytecodeHash[:], nil
}
