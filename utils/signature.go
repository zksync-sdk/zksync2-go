package utils

import (
	"github.com/ethereum/go-ethereum/accounts"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/signer/core/apitypes"
)

// IsMessageSignatureCorrect checks whether the message ECDSA signature is correct.
func IsMessageSignatureCorrect(address common.Address, msg, sig []byte) (bool, error) {
	signature := common.CopyBytes(sig)
	if signature[64] > 4 {
		signature[64] = signature[64] - 27
	}
	publicKey, err := crypto.SigToPub(accounts.TextHash(msg), signature)
	if err != nil {
		return false, err
	}
	return address == crypto.PubkeyToAddress(*publicKey), nil
}

// IsTypedDataSignatureCorrect checks whether the typed data ECDSA signature is correct.
func IsTypedDataSignatureCorrect(address common.Address, typedData apitypes.TypedData, sig []byte) (bool, error) {
	signature := common.CopyBytes(sig)
	if signature[64] > 4 {
		signature[64] = signature[64] - 27
	}
	hash, _, err := apitypes.TypedDataAndHash(typedData)
	if err != nil {
		return false, err
	}
	publicKey, err := crypto.SigToPub(hash, signature)
	if err != nil {
		return false, err
	}
	return address == crypto.PubkeyToAddress(*publicKey), nil
}
