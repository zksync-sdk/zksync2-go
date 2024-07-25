package utils

import (
	"encoding/binary"
	"fmt"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/zksync-sdk/zksync2-go/contracts/erc20"
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
