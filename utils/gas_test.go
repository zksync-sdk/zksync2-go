package utils

import (
	"github.com/stretchr/testify/assert"
	"math/big"
	"testing"
)

func TestCheckBaseCost(t *testing.T) {
	err := CheckBaseCost(big.NewInt(100), big.NewInt(99))
	assert.Error(t, err, "CheckBaseCost Should return error when base cost is greater than value")
	assert.ErrorContains(t, err, "The base cost of performing the priority operation is higher")
}

func TestScaleGasLimit(t *testing.T) {
	actual := ScaleGasLimit(big.NewInt(100))
	assert.True(t, big.NewInt(120).Cmp(actual) == 0, "Amounts should be the same")
}
