package model

import (
	"math"
	"time"

	schedulerenum "scheduler/internal/enum"
)

type ScheduledTask struct {
	ID            int64
	HandlerName   schedulerenum.TaskHandlerName
	Title         string
	Description   string
	Enabled       bool
	CronSpec      string
	Payload       string
	Timeout       time.Duration
	StaleAfter    *time.Duration
	MaxAttempts   int32
	MisfirePolicy schedulerenum.TaskMisfirePolicy
	AllowOverlap  bool
	Version       int64
	CreatedAt     *time.Time
	UpdatedAt     *time.Time
}

func (t *ScheduledTask) TimeoutSeconds() int32 {
	seconds := t.Timeout / time.Second
	if seconds > time.Duration(math.MaxInt32) {
		return math.MaxInt32
	}
	return int32(seconds) // #nosec G115 -- bounded by MaxInt32 above.
}
func (t *ScheduledTask) StaleAfterSeconds() *int32 {
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
