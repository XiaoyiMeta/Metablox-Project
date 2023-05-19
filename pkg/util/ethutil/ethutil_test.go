package ethutil_test

import (
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/stretchr/testify/assert"
	"go-metablox/pkg/util/ethutil"
	"testing"
)

func TestEcRecover(t *testing.T) {
	message := "Welcome to OpenSea!\n\nClick to sign in and accept the OpenSea Terms of Service: https://opensea.io/tos\n\nThis request will not trigger a blockchain transaction or cost any gas fees.\n\nYour authentication status will reset after 24 hours.\n\nWallet address:\n0x15c9254856b730f78a58b9b0b346e3c45f6af27f\n\nNonce:\n10526382-57f0-4cdb-94ab-bba2927b1e2a"
	signature := "0x499877d48df4a96a2b5fc9e9d720e4de0c5ec84aeecf25899a3f437f57e2c4e81a3565d603917192f0f2c244f75eebf1cc1040687a8a822b0b270aa81d6edc8a1c"
	account := "0x15c9254856b730f78a58b9b0b346e3c45f6af27f"
	ecRecover, err := ethutil.EcRecover([]byte(message), hexutil.MustDecode(signature))
	if err != nil {
		return
	}
	assert.Equal(t, common.HexToAddress(account), ecRecover)
}
