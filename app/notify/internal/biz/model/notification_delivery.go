package model

import (
	commonenum "common/pkg/enum"
	"notify/internal/enum"
	"time"
)

type NotificationDelivery struct {
	ID         int64
	EventID    string
	EventType  commonenum.EventType
	ReceiverID *int64
	Channel    enum.NotificationChannel
	Target     string
	Title      string
	Content    string
	Status     enum.NotificationDeliveryStatus
	RetryCount int32
	SentAt     *time.Time
	CreatedAt  *time.Time
	UpdatedAt  *time.Time
}
