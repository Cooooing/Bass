package model

import (
	commonenum "common/pkg/enum"
	notifyenum "notify/internal/enum"
	"time"
)

type NotificationEmailDelivery struct {
	ID                int64
	EventID           string
	EventType         commonenum.EventType
	ReceiverID        *int64
	ToEmail           string
	Subject           string
	Body              string
	ContentType       string
	Status            notifyenum.NotificationChannelStatus
	AttemptCount      int32
	LastAttemptAt     *time.Time
	ProviderMessageID *string
	ProviderResp      *string
	SentAt            *time.Time
	CreatedAt         *time.Time
	UpdatedAt         *time.Time
}
