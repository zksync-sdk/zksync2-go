package utils

import (
	"context"
	"encoding/binary"
	"fmt"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/zksync-sdk/zksync2-go/contracts/erc20"
	"github.com/zksync-sdk/zksync2-go/contracts/l1nativetokenvault"
	"github.com/zksync-sdk/zksync2-go/contracts/l2bridge"
	"github.com/zksync-sdk/zksync2-go/types"
	"math/big"
)

// CreateETH creates ETH token with appropriate Name, Symbol and Decimals values.
func CreateETH() *types.Token {
	return &types.Token{
		L1Address: common.Address{},
		L2Address: common.Address{},
		Name:      `ETH`,
		Symbol:    `ETH`,
		Decimals:  18,
	}
}

// Erc20DefaultBridgeData Returns the data needed for correct initialization of an L1 token counterpart on L2.
func Erc20DefaultBridgeData(l1TokenAddress common.Address, backend bind.ContractBackend) ([]byte, error) {
	var (
		name     = "Ether"
		symbol   = "ETH"
		decimals = uint8(18)
	)

	if l1TokenAddress != EthAddressInContracts {
		token, err := erc20.NewIERC20(l1TokenAddress, backend)
		if err != nil {
			return nil, fmt.Errorf("failed to load IERC20: %w", err)
		}
		name, err = token.Name(nil)
		if err != nil {
			return nil, err
		}
		symbol, err = token.Symbol(nil)
		if err != nil {
			return nil, err
		}
		decimals, err = token.Decimals(nil)
		if err != nil {
			return nil, err
		}
	}

	stringAbiType, err := abi.NewType("string", "", nil)
	if err != nil {
		return nil, err
	}
	uint256AbiType, err := abi.NewType("uint256", "", nil)
	if err != nil {
		return nil, err
	}
	bytesAbiType, err := abi.NewType("bytes", "", nil)
	if err != nil {
		return nil, err
	}

	nameEncoded, err := abi.Arguments{{Type: stringAbiType}}.Pack(name)
	if err != nil {
		return nil, err
	}
	symbolEncoded, err := abi.Arguments{{Type: stringAbiType}}.Pack(symbol)
	if err != nil {
		return nil, err
	}
	decimalsEncoded, err := abi.Arguments{{Type: uint256AbiType}}.Pack(big.NewInt(int64(decimals)))
	if err != nil {
		return nil, err
	}

	return abi.Arguments{
		{Type: bytesAbiType},
		{Type: bytesAbiType},
		{Type: bytesAbiType},
	}.Pack(nameEncoded, symbolEncoded, decimalsEncoded)
}

// Erc20BridgeCalldata returns the calldata that will be sent by an L1 ERC20 bridge to its L2 counterpart
// during bridging of a token.
func Erc20BridgeCalldata(l1TokenAddress, l1Sender, l2Receiver common.Address, amount *big.Int, bridgeData []byte) ([]byte, error) {
	l2BridgeAbi, err := l2bridge.IL2BridgeMetaData.GetAbi()
	if err != nil {
		return nil, fmt.Errorf("failed to load L2 bridge ABI: %w", err)
	}
	return l2BridgeAbi.Pack("finalizeDeposit", l1Sender, l2Receiver, l1TokenAddress, amount, bridgeData)
}

// HashedL2ToL1Msg returns a `keccak` encoded message with a given sender address and
// block number from the L1 messenger contract.
func HashedL2ToL1Msg(sender common.Address, msg []byte, txNumberInBlock uint16) common.Hash {
	txNumberInBlockBytes := make([]byte, 2)
	binary.BigEndian.PutUint16(txNumberInBlockBytes, txNumberInBlock)

	encodedMsg := append([]byte{0, 1}, txNumberInBlockBytes...)
	encodedMsg = append(encodedMsg, L1MessengerAddress.Bytes()...)
	encodedMsg = append(encodedMsg, common.LeftPadBytes(sender.Bytes(), 32)...)
	encodedMsg = append(encodedMsg, crypto.Keccak256(msg)...)

	return crypto.Keccak256Hash(encodedMsg)
}

// NativeTokenVaultAssetId returns the assetId for a token in the Native Token Vault with specific
// origin chainId and address.
func NativeTokenVaultAssetId(chainId *big.Int, address common.Address) (common.Hash, error) {
	uint256AbiType, err := abi.NewType("uint256", "", nil)
	if err != nil {
		return common.Hash{}, err
	}
	addressAbiType, err := abi.NewType("address", "", nil)
	if err != nil {
		return common.Hash{}, err
	}

	chainIdEncoded, err := abi.Arguments{{Type: uint256AbiType}}.Pack(chainId)
	if err != nil {
		return common.Hash{}, err
	}
	l2NativeTokenVaultAddressEncoded, err := abi.Arguments{{Type: addressAbiType}}.Pack(L2NativeTokenVaultAddress)
	if err != nil {
		return common.Hash{}, err
	}
	addressEncoded, err := abi.Arguments{{Type: addressAbiType}}.Pack(address)
	if err != nil {
		return common.Hash{}, err
	}

	encodedMsg := append(chainIdEncoded, l2NativeTokenVaultAddressEncoded...)
	encodedMsg = append(encodedMsg, addressEncoded...)
	return crypto.Keccak256Hash(encodedMsg), nil
}

// NativeTokenVaultTransferData encodes the data for a transfer of a token through the Native Token Vault.
func NativeTokenVaultTransferData(amount *big.Int, receiver, token common.Address) ([]byte, error) {
	uint256AbiType, err := abi.NewType("uint256", "", nil)
	if err != nil {
		return nil, err
	}
	addressAbiType, err := abi.NewType("address", "", nil)
	if err != nil {
		return nil, err
	}

	amountEncoded, err := abi.Arguments{{Type: uint256AbiType}}.Pack(amount)
	if err != nil {
		return nil, err
	}
	receiverEncoded, err := abi.Arguments{{Type: addressAbiType}}.Pack(receiver)
	if err != nil {
		return nil, err
	}
	tokenEncoded, err := abi.Arguments{{Type: addressAbiType}}.Pack(token)
	if err != nil {
		return nil, err
	}

	encodedMsg := append(amountEncoded, receiverEncoded...)
	encodedMsg = append(encodedMsg, tokenEncoded...)
	return encodedMsg, nil
}

// SecondBridgeDataV1 encodes asset transfer data for BridgeHub contract, using v1 encoding scheme (introduced in v26 upgrade).
// Can be utilized to encode deposit initiation data.
func SecondBridgeDataV1(assetId [32]byte, transferData []byte) ([]byte, error) {
	bytes32AbiType, err := abi.NewType("bytes32", "", nil)
	if err != nil {
		return nil, err
	}
	bytesAbiType, err := abi.NewType("bytes", "", nil)
	if err != nil {
		return nil, err
	}

	data, err := abi.Arguments{
		{Type: bytes32AbiType},
		{Type: bytesAbiType},
	}.Pack(assetId, transferData)
	if err != nil {
		return nil, err
	}
	encodedMsg := append(common.Hex2Bytes("01"), data...)
	return encodedMsg, nil
}

// ResolveAssetId resolves the assetId for a token.
func ResolveAssetId(ctx context.Context, vault *l1nativetokenvault.IL1NativeTokenVault, token common.Address, chainID *big.Int) ([32]byte, error) {
	if token == LegacyEthAddress {
		token = EthAddressInContracts
	}

	// In case only token is provided, we expect that it is a token inside Native Token Vault
	assetIdFromNTV, err := vault.AssetId(&bind.CallOpts{Context: ctx}, token)
	if err != nil {
		return [32]byte{}, err
	}

	if assetIdFromNTV != (common.Hash{}) {
		return assetIdFromNTV, nil
	}

	// Okay, the token have not been registered within the Native token vault.
	// There are two cases when it is possible:
	// - The token is native to L1 (it may or may not be bridged), but it has not been
	// registered within NTV after the Gateway upgrade. We assume that this is not the case
	// as the SDK is expected to work only after the full migration is done.
	// - The token is native to the current chain, and it has never been bridged.
	return NativeTokenVaultAssetId(chainID, token)
}
