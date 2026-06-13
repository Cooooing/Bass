package model

import (
	commonenum "common/pkg/enum"
	"time"
)

type OutboxEvent struct {
	ID         int64
	EventID    string
	EventType  commonenum.EventType
	Subject    commonenum.EventSubject
	Payload    string
	Headers    map[string]string
	Status     commonenum.OutboxEventStatus
	RetryCount int32
	LastError  *string
	UpdatedAt  *time.Time
}
