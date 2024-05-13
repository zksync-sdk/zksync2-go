package test

import (
	"context"
	"github.com/stretchr/testify/assert"
	"github.com/zksync-sdk/zksync2-go/accounts"
	"github.com/zksync-sdk/zksync2-go/clients"
	zkTypes "github.com/zksync-sdk/zksync2-go/types"
	"math/big"
	"testing"
)

func TestIntegration_PopulateTransactionECDSA(t *testing.T) {
	client, err := clients.DialBase(ZkSyncEraProvider)
	defer client.Close()
	assert.NoError(t, err, "clients.DialBase should not return an error")

	tx := zkTypes.Transaction712{
		To:      &Address2,
		Value:   big.NewInt(7_000_000_000),
		ChainID: big.NewInt(270),
		From:    &Address1,
	}

	err = accounts.PopulateTransactionECDSA(context.Background(), &tx, PrivateKey1, client)
	assert.NoError(t, err, "PopulateTransactionECDSA should not return an error")

	assert.Equal(t, tx.To, &Address2, "To addresses must be the same")
	assert.Equal(t, tx.From, &Address1, "From addresses must be the same")
	assert.Equal(t, tx.ChainID, big.NewInt(270), "Chain IDs must be the same")
	assert.Equal(t, tx.Value, big.NewInt(7_000_000_000), "Values must be the same")
	assert.NotNil(t, tx.Data, "Data must not be nil")
	assert.NotNil(t, tx.Gas, "Gas must not be nil")
	assert.NotNil(t, tx.GasFeeCap, "GasFeeCap must not be nil")
	assert.NotNil(t, tx.GasTipCap, "GasTipCap must not be nil")
	assert.NotNil(t, tx.Meta, "Meta must not be nil")
}

func TestIntegration_PopulateTransactionECDSA_ErrorNoClientProvided(t *testing.T) {
	tx := zkTypes.Transaction712{
		To:      &Address2,
		Value:   big.NewInt(7_000_000_000),
		ChainID: big.NewInt(270),
		From:    &Address1,
	}

	err := accounts.PopulateTransactionECDSA(context.Background(), &tx, PrivateKey1, nil)
	assert.Error(t, err, "PopulateTransactionECDSA should return an error when client is not provided")
}

func TestIntegration_PopulateTransactionMultipleECDSA_ErrorNoMultipleKeysProvided(t *testing.T) {
	client, err := clients.DialBase(ZkSyncEraProvider)
	defer client.Close()
	assert.NoError(t, err, "clients.DialBase should not return an error")

	err = accounts.PopulateTransactionMultipleECDSA(context.Background(), &zkTypes.Transaction712{}, [1]string{PrivateKey1}, client)
	assert.Error(t, err, "PopulateTransactionMultipleECDSA should return an error when only one private key is provided")
}
