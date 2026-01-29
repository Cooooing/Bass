package model

import (
	"encoding/json"
	"time"
)

type Task struct {
	TaskName string
	TaskId   string
	Delay    bool          // 是否需要延迟执行（false = 立即执行）
	Interval time.Duration // 执行间隔
	MaxRetry int           // 最大重试次数
	Data     json.RawMessage
}
