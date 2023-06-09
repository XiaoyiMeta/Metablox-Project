// =================================================================================
// This is auto-generated by GoFrame CLI tool only once. Fill this file as you wish.
// =================================================================================

package dao

import (
	"go-metablox/app/foundation/internal/dao/internal"
)

// internalDidVerifiableCredentialDao is internal type for wrapping internal DAO implements.
type internalDidVerifiableCredentialDao = *internal.DidVerifiableCredentialDao

// didVerifiableCredentialDao is the data access object for table did_verifiable_credential.
// You can define custom methods on it to extend its functionality as you wish.
type didVerifiableCredentialDao struct {
	internalDidVerifiableCredentialDao
}

var (
	// DidVerifiableCredential is globally public accessible object for table did_verifiable_credential operations.
	DidVerifiableCredential = didVerifiableCredentialDao{
		internal.NewDidVerifiableCredentialDao(),
	}
)

// Fill with you ideas below.
