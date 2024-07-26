package utils

import (
	"github.com/ethereum/go-ethereum/accounts"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
)

// IsMessageSignatureCorrect checks whether the message ECDSA signature is correct.
func IsMessageSignatureCorrect(address common.Address, msg, sig []byte) (bool, error) {
	publicKey, err := crypto.SigToPub(accounts.TextHash(msg), sig)
	if err != nil {
		return false, err
	}
	return address == crypto.PubkeyToAddress(*publicKey), nil
}
