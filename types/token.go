package types

import (
	"bytes"
	"github.com/ethereum/go-ethereum/common"
)

// Token represents a token with addresses on both L1 and L2 chains.
type Token struct {
	L1Address common.Address `json:"l1Address"` // Token address on L1.
	L2Address common.Address `json:"l2Address"` // Token address on L2.
	Name      string         `json:"name"`      // Token name.
	Symbol    string         `json:"symbol"`    // Token symbol.
	Decimals  uint8          `json:"decimals"`  // Number of decimals for the token.
}

func (t *Token) IsETH() bool {
	return bytes.Equal(t.L2Address.Bytes(), common.Address{}.Bytes()) && t.Symbol == "ETH"
}
