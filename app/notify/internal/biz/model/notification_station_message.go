package model

import (
	commonenum "common/pkg/enum"
	notifyenum "notify/internal/enum"
	"time"
)

type NotificationStationMessage struct {
	ID         int64
	EventID    string
	EventType  commonenum.EventType
	ReceiverID int64
	Title      string
	Content    string
	Status     notifyenum.NotificationChannelStatus
	ReadAt     *time.Time
	CreatedAt  *time.Time
	UpdatedAt  *time.Time
}
