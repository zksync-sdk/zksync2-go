package utils

import (
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/common/math"
	"github.com/ethereum/go-ethereum/signer/core/apitypes"
	"github.com/stretchr/testify/assert"
	"testing"
)

var Address1 = common.HexToAddress("0x36615Cf349d7F6344891B1e7CA7C72883F5dc049")

func TestIsMessageSignatureCorrect(t *testing.T) {
	const message = "Hello, world!"
	const signature = "0xb04f825363596418c630425916f73447d636094a75e47b45e2eb59d8a6c7d5035355f03b903b84700376f0efa23f3b095815776c5c6daf2b371a0a61b5f703451b"
	address := common.HexToAddress("0x36615Cf349d7F6344891B1e7CA7C72883F5dc049")
	actual, err := IsMessageSignatureCorrect(address, []byte(message), common.FromHex(signature))
	assert.NoError(t, err, "IsMessageSignatureCorrect should not return error for valid message signature")
	assert.True(t, actual, "Message signature should be valid")
}

func TestIsMessageSignatureCorrectInvalidSignature(t *testing.T) {
	const message = "Hello world"
	const invalidSignature = "0xb04f825363596418c630425916f73447d636094a75e47b45e2eb59d8a6c7d5035355f03b903b84700376f0efa23f3b095815776c5c6daf2b371a0a61b5f703451b"
	address := common.HexToAddress("0x36615Cf349d7F6344891B1e7CA7C72883F5dc049")
	actual, err := IsMessageSignatureCorrect(address, []byte(message), common.FromHex(invalidSignature))
	assert.NoError(t, err, "IsMessageSignatureCorrect should not return error for invalid message signature")
	assert.False(t, actual, "Message signature should be invalid")
}

func TestIsTypedDataSignatureCorrect(t *testing.T) {
	typedData := apitypes.TypedData{
		Domain: apitypes.TypedDataDomain{
			Name:    "Example",
			Version: "1",
			ChainId: math.NewHexOrDecimal256(270),
		},
		Types: apitypes.Types{
			"Person": []apitypes.Type{
				{Name: "name", Type: "string"},
				{Name: "age", Type: "uint8"},
			},
			"EIP712Domain": []apitypes.Type{
				{Name: "name", Type: "string"},
				{Name: "version", Type: "string"},
				{Name: "chainId", Type: "uint256"},
			},
		},
		PrimaryType: "Person",
		Message: apitypes.TypedDataMessage{
			"name": "John",
			"age":  hexutil.EncodeUint64(30),
		},
	}

	const invalidSignature = "0xbcaf0673c0c2b0e120165d207d42281d0c6e85f0a7f6b8044b0578a91cf5bda66b4aeb62aca4ae17012a38d71c9943e27285792fa7d788d848f849e3ea2e614b1b"
	actual, err := IsTypedDataSignatureCorrect(Address1, typedData, common.FromHex(invalidSignature))
	assert.NoError(t, err, "IsTypedDataSignatureCorrect should not return an error")
	assert.True(t, actual, "Typed data signature should be valid")
}

func TestIsTypedDataSignatureCorrectInvalidSignature(t *testing.T) {
	typedData := apitypes.TypedData{
		Domain: apitypes.TypedDataDomain{
			Name:    "Example",
			Version: "1",
			ChainId: math.NewHexOrDecimal256(270),
		},
		Types: apitypes.Types{
			"Person": []apitypes.Type{
				{Name: "name", Type: "string"},
				{Name: "age", Type: "uint8"},
			},
			"EIP712Domain": []apitypes.Type{
				{Name: "name", Type: "string"},
				{Name: "version", Type: "string"},
				{Name: "chainId", Type: "uint256"},
			},
		},
		PrimaryType: "Person",
		Message: apitypes.TypedDataMessage{
			"name": "Bob", // instead of John
			"age":  hexutil.EncodeUint64(30),
		},
	}

	const signature = "0xbcaf0673c0c2b0e120165d207d42281d0c6e85f0a7f6b8044b0578a91cf5bda66b4aeb62aca4ae17012a38d71c9943e27285792fa7d788d848f849e3ea2e614b1b"
	actual, err := IsTypedDataSignatureCorrect(Address1, typedData, common.FromHex(signature))
	assert.NoError(t, err, "IsTypedDataSignatureCorrect should not return an error")
	assert.False(t, actual, "Typed data signature should be valid")
}
