package model

import (
	commonenum "common/pkg/enum"
	"notify/internal/enum"
	"time"
)

type NotificationSetting struct {
	ID        int64
	UserID    int64
	EventType commonenum.EventType
	Channel   enum.NotificationChannel
	Enable    bool
	CreatedAt *time.Time
	UpdatedAt *time.Time
}
