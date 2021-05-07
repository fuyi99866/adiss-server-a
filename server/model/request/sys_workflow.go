package request

import "adiss-server-a/model"

type WorkflowProcessSearch struct {
	model.WorkflowProcess
	PageInfo
}
