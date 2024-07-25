package utils

import (
	"github.com/ethereum/go-ethereum/common"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestHashedL2ToL1Msg(t *testing.T) {
	sender := common.HexToAddress("0x36615Cf349d7F6344891B1e7CA7C72883F5dc049")
	withdrawETHMessage := common.FromHex("0x6c0960f936615cf349d7f6344891b1e7ca7c72883f5dc04900000000000000000000000000000000000000000000000000000001a13b8600")
	withdrawETHMessageHash := common.HexToHash("0x521bd25904766c83fe868d6a29cbffa011afd8a1754f6c9a52b053b693e42f18")
	expected := HashedL2ToL1Msg(sender, withdrawETHMessage, 0)
	assert.Equal(t, expected, withdrawETHMessageHash, "Hashes should be the same")
}
