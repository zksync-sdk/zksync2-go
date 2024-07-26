package utils

import (
	"github.com/ethereum/go-ethereum/common"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestIsMessageSignatureCorrect(t *testing.T) {
	const message = "Hello, world!"
	const signature = "0xb04f825363596418c630425916f73447d636094a75e47b45e2eb59d8a6c7d5035355f03b903b84700376f0efa23f3b095815776c5c6daf2b371a0a61b5f7034500"
	address := common.HexToAddress("0x36615Cf349d7F6344891B1e7CA7C72883F5dc049")
	actual, err := IsMessageSignatureCorrect(address, []byte(message), common.FromHex(signature))
	assert.NoError(t, err, "Should throw error for valid message signature")
	assert.True(t, actual, "Message signature should be valid")
}

func TestIsMessageSignatureCorrectInvalidSignature(t *testing.T) {
	const message = "Hello world"
	const signature = "0xb04f825363596418c630425916f73447d636094a75e47b45e2eb59d8a6c7d5035355f03b903b84700376f0efa23f3b095815776c5c6daf2b371a0a61b5f7034500"
	address := common.HexToAddress("0x36615Cf349d7F6344891B1e7CA7C72883F5dc049")
	actual, err := IsMessageSignatureCorrect(address, []byte(message), common.FromHex(signature))
	assert.NoError(t, err, "Should throw error for valid message signature")
	assert.False(t, actual, "Message signature should be invalid")
}
