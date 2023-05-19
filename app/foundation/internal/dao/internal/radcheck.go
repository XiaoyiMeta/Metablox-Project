// ==========================================================================
// Code generated by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// RadcheckDao is the data access object for table radcheck.
type RadcheckDao struct {
	table   string          // table is the underlying table name of the DAO.
	group   string          // group is the database configuration group name of current DAO.
	columns RadcheckColumns // columns contains all the column names of Table for convenient usage.
}

// RadcheckColumns defines and stores column names for table radcheck.
type RadcheckColumns struct {
	Id        string //
	Username  string //
	Attribute string //
	Op        string //
	Value     string //
}

// radcheckColumns holds the columns for table radcheck.
var radcheckColumns = RadcheckColumns{
	Id:        "id",
	Username:  "username",
	Attribute: "attribute",
	Op:        "op",
	Value:     "value",
}

// NewRadcheckDao creates and returns a new DAO object for table data access.
func NewRadcheckDao() *RadcheckDao {
	return &RadcheckDao{
		group:   "default",
		table:   "radcheck",
		columns: radcheckColumns,
	}
}

// DB retrieves and returns the underlying raw database management object of current DAO.
func (dao *RadcheckDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of current dao.
func (dao *RadcheckDao) Table() string {
	return dao.table
}

// Columns returns all column names of current dao.
func (dao *RadcheckDao) Columns() RadcheckColumns {
	return dao.columns
}

// Group returns the configuration group name of database of current dao.
func (dao *RadcheckDao) Group() string {
	return dao.group
}

// Ctx creates and returns the Model for current DAO, It automatically sets the context for current operation.
func (dao *RadcheckDao) Ctx(ctx context.Context) *gdb.Model {
	return dao.DB().Model(dao.table).Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
// It rollbacks the transaction and returns the error from function f if it returns non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note that, you should not Commit or Rollback the transaction in function f
// as it is automatically handled by this function.
func (dao *RadcheckDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}