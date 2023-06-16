package types

import (
	"bytes"
	"github.com/ethereum/go-ethereum/common"
)

type Token struct {
	L1Address common.Address `json:"l1Address"`
	L2Address common.Address `json:"l2Address"`
	Name      string         `json:"name"`
	Symbol    string         `json:"symbol"`
	Decimals  uint           `json:"decimals"`
}

func (t *Token) IsETH() bool {
	return bytes.Equal(t.L2Address.Bytes(), common.Address{}.Bytes()) && t.Symbol == "ETH"
}
