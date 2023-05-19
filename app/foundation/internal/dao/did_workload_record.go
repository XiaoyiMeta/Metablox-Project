// =================================================================================
// This is auto-generated by GoFrame CLI tool only once. Fill this file as you wish.
// =================================================================================

package dao

import (
	"go-metablox/app/foundation/internal/dao/internal"
)

// internalDidWorkloadRecordDao is internal type for wrapping internal DAO implements.
type internalDidWorkloadRecordDao = *internal.DidWorkloadRecordDao

// didWorkloadRecordDao is the data access object for table did_workload_record.
// You can define custom methods on it to extend its functionality as you wish.
type didWorkloadRecordDao struct {
	internalDidWorkloadRecordDao
}

var (
	// DidWorkloadRecord is globally public accessible object for table did_workload_record operations.
	DidWorkloadRecord = didWorkloadRecordDao{
		internal.NewDidWorkloadRecordDao(),
	}
)

// Fill with you ideas below.
