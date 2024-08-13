package test

import (
	"context"
	"github.com/stretchr/testify/assert"
	"github.com/zksync-sdk/zksync2-go/accounts"
	"github.com/zksync-sdk/zksync2-go/clients"
	"github.com/zksync-sdk/zksync2-go/types"
	"math/big"
	"testing"
)

func TestIntegration_PopulateTransactionECDSA(t *testing.T) {
	client, err := clients.Dial(L2ChainURL)
	defer client.Close()
	assert.NoError(t, err, "clients.DialBase should not return an error")

	tx := types.Transaction{
		To:      &Address2,
		Value:   big.NewInt(7_000_000_000),
		ChainID: big.NewInt(270),
		From:    &Address1,
	}

	err = accounts.PopulateTransactionECDSA(context.Background(), &tx, PrivateKey1, client)
	assert.NoError(t, err, "PopulateTransactionECDSA should not return an error")

	assert.Equal(t, &Address2, tx.To, "To addresses must be the same")
	assert.Equal(t, &Address1, tx.From, "From addresses must be the same")
	assert.Equal(t, big.NewInt(7_000_000_000), tx.Value, "Values must be the same")
	assert.NotNil(t, tx.ChainID, "Chain IDs must not be nil")
	assert.NotNil(t, tx.Data, "Data must not be nil")
	assert.NotNil(t, tx.Gas, "Gas must not be nil")
	assert.NotNil(t, tx.GasFeeCap, "GasFeeCap must not be nil")
	assert.NotNil(t, tx.GasTipCap, "GasTipCap must not be nil")
	assert.NotNil(t, tx.GasPerPubdata, "GasPerPubdata must not be nil")
}

func TestIntegration_PopulateTransactionECDSA_ErrorNoClientProvided(t *testing.T) {
	tx := types.Transaction{
		To:      &Address2,
		Value:   big.NewInt(7_000_000_000),
		ChainID: big.NewInt(270),
		From:    &Address1,
	}

	err := accounts.PopulateTransactionECDSA(context.Background(), &tx, PrivateKey1, nil)
	assert.Error(t, err, "PopulateTransactionECDSA should return an error when client is not provided")
}

func TestIntegration_PopulateTransactionMultipleECDSA_ErrorNoMultipleKeysProvided(t *testing.T) {
	client, err := clients.Dial(L2ChainURL)
	defer client.Close()
	assert.NoError(t, err, "clients.DialBase should not return an error")

	err = accounts.PopulateTransactionMultipleECDSA(context.Background(), &types.Transaction{}, [1]string{PrivateKey1}, client)
	assert.Error(t, err, "PopulateTransactionMultipleECDSA should return an error when only one private key is provided")
}
