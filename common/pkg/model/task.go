package model

import (
	"encoding/json"
	"time"
)

type Task struct {
	TaskName  string
	TaskId    string
	Version   int64         // 版本号
	Delay     bool          // 是否需要延迟执行（false = 立即执行）
	Interval  time.Duration // 间隔执行
	Cronspec  string        // cron 表达式
	MaxRetry  int           // 最大重试次数
	Retention time.Duration // 任务保留时间
	Data      json.RawMessage
}
