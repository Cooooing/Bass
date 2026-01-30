package model

import (
	"encoding/json"
)

type Task struct {
	TaskName string
	TaskId   string
	Cronspec string // cron 表达式
	MaxRetry int    // 最大重试次数
	Data     json.RawMessage
}
