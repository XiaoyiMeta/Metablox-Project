// ==========================================================================
// Code generated by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// DidDocumentDao is the data access object for table did_document.
type DidDocumentDao struct {
	table   string             // table is the underlying table name of the DAO.
	group   string             // group is the database configuration group name of current DAO.
	columns DidDocumentColumns // columns contains all the column names of Table for convenient usage.
}

// DidDocumentColumns defines and stores column names for table did_document.
type DidDocumentColumns struct {
	Id        string //
	ChainId   string //
	Did       string //
	RawText   string //
	CreatedAt string //
	UpdatedAt string //
	DeletedAt string //
}

// didDocumentColumns holds the columns for table did_document.
var didDocumentColumns = DidDocumentColumns{
	Id:        "id",
	ChainId:   "chain_id",
	Did:       "did",
	RawText:   "raw_text",
	CreatedAt: "created_at",
	UpdatedAt: "updated_at",
	DeletedAt: "deleted_at",
}

// NewDidDocumentDao creates and returns a new DAO object for table data access.
func NewDidDocumentDao() *DidDocumentDao {
	return &DidDocumentDao{
		group:   "default",
		table:   "did_document",
		columns: didDocumentColumns,
	}
}

// DB retrieves and returns the underlying raw database management object of current DAO.
func (dao *DidDocumentDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of current dao.
func (dao *DidDocumentDao) Table() string {
	return dao.table
}

// Columns returns all column names of current dao.
func (dao *DidDocumentDao) Columns() DidDocumentColumns {
	return dao.columns
}

// Group returns the configuration group name of database of current dao.
func (dao *DidDocumentDao) Group() string {
	return dao.group
}

// Ctx creates and returns the Model for current DAO, It automatically sets the context for current operation.
func (dao *DidDocumentDao) Ctx(ctx context.Context) *gdb.Model {
	return dao.DB().Model(dao.table).Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
// It rollbacks the transaction and returns the error from function f if it returns non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note that, you should not Commit or Rollback the transaction in function f
// as it is automatically handled by this function.
func (dao *DidDocumentDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
