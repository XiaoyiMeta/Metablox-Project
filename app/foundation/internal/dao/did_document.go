// =================================================================================
// This is auto-generated by GoFrame CLI tool only once. Fill this file as you wish.
// =================================================================================

package dao

import (
	"go-metablox/app/foundation/internal/dao/internal"
)

// internalDidDocumentDao is internal type for wrapping internal DAO implements.
type internalDidDocumentDao = *internal.DidDocumentDao

// didDocumentDao is the data access object for table did_document.
// You can define custom methods on it to extend its functionality as you wish.
type didDocumentDao struct {
	internalDidDocumentDao
}

var (
	// DidDocument is globally public accessible object for table did_document operations.
	DidDocument = didDocumentDao{
		internal.NewDidDocumentDao(),
	}
)

// Fill with you ideas below.