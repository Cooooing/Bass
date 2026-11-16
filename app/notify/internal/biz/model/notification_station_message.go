package model

import (
	"time"

	notifyenum "notify/internal/enum"
)

type NotificationStationMessage struct {
	ID               int64
	EventID          string
	NotificationType notifyenum.NotificationType
	ReceiverID       int64
	BusinessKey      *string
	Title            string
	Content          string
	Status           notifyenum.NotificationChannelStatus
	ReadAt           *time.Time
	CreatedAt        *time.Time
	UpdatedAt        *time.Time
}
