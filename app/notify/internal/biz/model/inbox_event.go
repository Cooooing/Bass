package model

import (
	commonenum "common/pkg/enum"
	"time"
)

type InboxEvent struct {
	ID          int64
	EventID     string
	EventType   commonenum.EventType
	Subject     commonenum.EventSubject
	Payload     []byte
	Status      commonenum.InboxEventStatus
	RetryCount  int32
	ReceivedAt  time.Time
	ProcessedAt *time.Time
	CreatedAt   *time.Time
	UpdatedAt   *time.Time
}
