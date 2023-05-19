package ethutil

import (
	"context"
	"fmt"
	"github.com/ethereum/go-ethereum/accounts"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/accounts/keystore"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/signer/core/apitypes"
	"math/big"
	"regexp"
)

var _ ISigner = &Signer{}

// ISigner defines the external API through which signing requests are made.
type ISigner interface {
	// SignTextData - request to sign the given data (plus prefix)
	// EIP-191
	SignTextData(address common.MixedcaseAddress, data string) (hexutil.Bytes, error)
	// SignTypedData - request to sign the given structured data (plus prefix)
	// EIP-712
	SignTypedData(address common.MixedcaseAddress, data apitypes.TypedData) (hexutil.Bytes, error)
	//// EcRecover - recover public key from given message and signature
	//EcRecover(data hexutil.Bytes, sig hexutil.Bytes) (common.Address, error)
}

// Signer defines the actual implementation of ISigner
type Signer struct {
	chainID *big.Int
	store   *keystore.KeyStore
}

func NewSigner(keystore *keystore.KeyStore, chainID int64) *Signer {
	signer := &Signer{big.NewInt(chainID), keystore}
	return signer
}

// SignTextData signs EIP191 conformance text data
func (api *Signer) SignTextData(address common.MixedcaseAddress, textData string) (hexutil.Bytes, error) {
	textHash, rawData, err := TextDataAndHash(textData)
	fmt.Println("Sign text data:", rawData)
	if err != nil {
		return nil, err
	}
	account, err := api.store.Find(accounts.Account{Address: address.Address()})
	if err != nil {
		return nil, err
	}
	signature, err := api.store.SignHash(account, textHash)
	return signature, err
}

// SignTypedData signs EIP-712 conformance typed data
// hash = keccak256("\x19${byteVersion}${domainSeparator}${hashStruct(message)}")
// It returns the signature and/or any error
func (api *Signer) SignTypedData(address common.MixedcaseAddress, typedData apitypes.TypedData) (hexutil.Bytes, error) {
	hash, rawData, err := apitypes.TypedDataAndHash(typedData)
	fmt.Println("Sign typed data:", rawData)
	if err != nil {
		return nil, err
	}
	account, err := api.store.Find(accounts.Account{Address: address.Address()})
	if err != nil {
		return nil, err
	}
	signature, err := api.store.SignHash(account, hash)
	return signature, err
}

// EcRecover recovers the address associated with the given sig.
// Only compatible with `text/plain`
func EcRecover(data hexutil.Bytes, sig hexutil.Bytes) (common.Address, error) {
	// Returns the address for the Account that was used to create the signature.
	//
	// Note, this function is compatible with eth_sign and personal_sign. As such it recovers
	// the address of:
	// hash = keccak256("\x19Ethereum Signed Message:\n${message length}${message}")
	// addr = ecrecover(hash, signature)
	//
	// Note, the signature must conform to the secp256k1 curve R, S and V values, where
	// the V value must be 27 or 28 for legacy reasons.
	//
	// https://github.com/ethereum/go-ethereum/wiki/Management-APIs#personal_ecRecover
	if len(sig) != 65 {
		return common.Address{}, fmt.Errorf("signature must be 65 bytes long")
	}
	if sig[64] != 27 && sig[64] != 28 {
		return common.Address{}, fmt.Errorf("invalid Ethereum signature (V is not 27 or 28)")
	}
	sig[64] -= 27 // Transform yellow paper V from 27/28 to 0/1
	hash := accounts.TextHash(data)
	rpk, err := crypto.SigToPub(hash, sig)
	if err != nil {
		return common.Address{}, err
	}
	return crypto.PubkeyToAddress(*rpk), nil
}

// TypedDataAndHash is a helper function that calculates a hash for typed data conforming to EIP-712.
// This hash can then be safely used to calculate a signature.
// This gives context to the signed typed data and prevents signing of transactions.
func TypedDataAndHash(typedData apitypes.TypedData) ([]byte, string, error) {
	return apitypes.TypedDataAndHash(typedData)
}

// TextDataAndHash EIP-191
func TextDataAndHash(textData string) ([]byte, string, error) {
	decode, err := hexutil.Decode(textData)
	if err != nil {
		return nil, "", err
	}
	sighash, rawData := accounts.TextAndHash(decode)
	return sighash, rawData, nil
}

func IsHexTxHash(input string) bool {
	txHashPattern := "^0x([A-Fa-f0-9]{64})$"
	match, _ := regexp.MatchString(txHashPattern, input)
	return match
}

func IsHexAddress(s string) bool {
	return common.IsHexAddress(s)
}

func IsContract(ctx context.Context, backend bind.ContractBackend, address common.Address) (bool, error) {
	code, err := backend.CodeAt(ctx, address, nil)
	if err != nil {
		return false, err
	}
	return len(code) > 0, nil
}
