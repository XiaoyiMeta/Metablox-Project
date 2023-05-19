package didutil

import (
	"crypto/ecdsa"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/mr-tron/base58"
)

func ParseDID(privateKey *ecdsa.PrivateKey) string {
	pubData := crypto.FromECDSAPub(&privateKey.PublicKey)
	hash := crypto.Keccak256(pubData)
	didString := base58.Encode(hash)
	returnString := "did:metablox:" + didString
	return returnString
}
