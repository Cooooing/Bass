package repo

import (
	"time"

	schedulerenum "scheduler/internal/enum"
)

type MessageHandleResult struct {
	Action     schedulerenum.MessageHandleAction
	RetryAfter time.Duration
}
