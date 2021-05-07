package request

import "adiss-server-a/model"

type SysOperationRecordSearch struct {
	model.SysOperationRecord
	PageInfo
}
