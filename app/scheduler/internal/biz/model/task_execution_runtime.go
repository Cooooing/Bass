package model

import schedulerenum "scheduler/internal/enum"

type TaskExecutionRuntime struct {
	ExecutionRecordID int64
	TaskID            int64
	State             schedulerenum.TaskExecutionRuntimeState
}
