package model

import (
	"math"
	"time"

	schedulerenum "scheduler/internal/enum"
)

type DelayedTaskVersion struct {
	ID            int64
	DelayedTaskID int64
	Version       int64
	TaskKey       string
	HandlerName   schedulerenum.TaskHandlerName
	Title         string
	Description   string
	Enabled       bool
	Timeout       time.Duration
	StaleAfter    *time.Duration
	MaxAttempts   int32
	MisfirePolicy schedulerenum.TaskMisfirePolicy
	CreatedAt     *time.Time
	UpdatedAt     *time.Time
}

func (t *DelayedTaskVersion) TimeoutSeconds() int32 {
	seconds := t.Timeout / time.Second
	if seconds > time.Duration(math.MaxInt32) {
		return math.MaxInt32
	}
	return int32(seconds) // #nosec G115 -- bounded by MaxInt32 above.
}
func (t *DelayedTaskVersion) StaleAfterSeconds() *int32 {
	if t.StaleAfter == nil {
		return nil
	}
	seconds := *t.StaleAfter / time.Second
	if seconds > time.Duration(math.MaxInt32) {
		value := int32(math.MaxInt32)
		return &value
	}
	value := int32(seconds) // #nosec G115 -- bounded by MaxInt32 above.
	return &value
}
