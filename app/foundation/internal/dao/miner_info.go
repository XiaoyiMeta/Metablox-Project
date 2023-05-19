// =================================================================================
// This is auto-generated by GoFrame CLI tool only once. Fill this file as you wish.
// =================================================================================

package dao

import (
	"go-metablox/app/foundation/internal/dao/internal"
)

// internalMinerInfoDao is internal type for wrapping internal DAO implements.
type internalMinerInfoDao = *internal.MinerInfoDao

// minerInfoDao is the data access object for table miner_info.
// You can define custom methods on it to extend its functionality as you wish.
type minerInfoDao struct {
	internalMinerInfoDao
}

var (
	// MinerInfo is globally public accessible object for table miner_info operations.
	MinerInfo = minerInfoDao{
		internal.NewMinerInfoDao(),
	}
)

// Fill with you ideas below.