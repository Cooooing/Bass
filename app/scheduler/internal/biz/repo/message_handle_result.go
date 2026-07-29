package repo

import (
	schedulerenum "scheduler/internal/enum"
	"time"
)

type MessageHandleResult struct {
	Action     schedulerenum.MessageHandleAction
	RetryAfter time.Duration
}
