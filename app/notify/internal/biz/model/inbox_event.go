package model

import (
	commonenum "common/pkg/enum"
	"time"
)

type InboxEvent struct {
	ID                  int64
	EventID             string
	EventType           commonenum.EventType
	Subject             commonenum.EventSubject
	Payload             string
	Status              commonenum.InboxEventStatus
	AttemptCount        int32
	LastError           *string
	ProcessingStartedAt time.Time
	ProcessedAt         *time.Time
	CreatedAt           *time.Time
	UpdatedAt           *time.Time
}
