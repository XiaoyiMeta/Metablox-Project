// ==========================================================================
// Code generated by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// DappContractDao is the data access object for table dapp_contract.
type DappContractDao struct {
	table   string              // table is the underlying table name of the DAO.
	group   string              // group is the database configuration group name of current DAO.
	columns DappContractColumns // columns contains all the column names of Table for convenient usage.
}

// DappContractColumns defines and stores column names for table dapp_contract.
type DappContractColumns struct {
	Id              string //
	BizType         string //
	ContractType    string //
	Description     string //
	ChainId         string //
	ContractAddress string //
	CreatorAddress  string //
	TxHash          string //
	BlockTimestamp  string //
	BlockNumber     string //
	Disabled        string //
	Version         string //
	Abi             string //
	CreatedAt       string //
	UpdatedAt       string //
	DeletedAt       string //
}

// dappContractColumns holds the columns for table dapp_contract.
var dappContractColumns = DappContractColumns{
	Id:              "id",
	BizType:         "biz_type",
	ContractType:    "contract_type",
	Description:     "description",
	ChainId:         "chain_id",
	ContractAddress: "contract_address",
	CreatorAddress:  "creator_address",
	TxHash:          "tx_hash",
	BlockTimestamp:  "block_timestamp",
	BlockNumber:     "block_number",
	Disabled:        "disabled",
	Version:         "version",
	Abi:             "abi",
	CreatedAt:       "created_at",
	UpdatedAt:       "updated_at",
	DeletedAt:       "deleted_at",
}

// NewDappContractDao creates and returns a new DAO object for table data access.
func NewDappContractDao() *DappContractDao {
	return &DappContractDao{
		group:   "default",
		table:   "dapp_contract",
		columns: dappContractColumns,
	}
}

// DB retrieves and returns the underlying raw database management object of current DAO.
func (dao *DappContractDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of current dao.
func (dao *DappContractDao) Table() string {
	return dao.table
}

// Columns returns all column names of current dao.
func (dao *DappContractDao) Columns() DappContractColumns {
	return dao.columns
}

// Group returns the configuration group name of database of current dao.
func (dao *DappContractDao) Group() string {
	return dao.group
}

// Ctx creates and returns the Model for current DAO, It automatically sets the context for current operation.
func (dao *DappContractDao) Ctx(ctx context.Context) *gdb.Model {
	return dao.DB().Model(dao.table).Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
// It rollbacks the transaction and returns the error from function f if it returns non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note that, you should not Commit or Rollback the transaction in function f
// as it is automatically handled by this function.
func (dao *DappContractDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
