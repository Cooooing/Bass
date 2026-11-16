package model

import (
	"time"

	notifyenum "notify/internal/enum"
)

type NotificationLarkWebhookDelivery struct {
	ID               int64
	EventID          string
	NotificationType notifyenum.NotificationType
	WebhookID        string
	RequestBody      string
	Status           notifyenum.NotificationChannelStatus
	AttemptCount     int32
	LastAttemptAt    *time.Time
	HTTPStatus       *int
	RespBody         *string
	SentAt           *time.Time
	CreatedAt        *time.Time
	UpdatedAt        *time.Time
}
