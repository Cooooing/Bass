package model

import (
	commonenum "common/pkg/enum"
	notifyenum "notify/internal/enum"
	"time"
)

type NotificationLarkWebhookDelivery struct {
	ID            int64
	EventID       string
	EventType     commonenum.EventType
	WebhookID     string
	RequestBody   string
	Status        notifyenum.NotificationChannelStatus
	AttemptCount  int32
	LastAttemptAt *time.Time
	HTTPStatus    *int
	RespBody      *string
	SentAt        *time.Time
	CreatedAt     *time.Time
	UpdatedAt     *time.Time
}
