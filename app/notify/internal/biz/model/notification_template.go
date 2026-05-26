package model

import (
	commonenum "common/pkg/enum"
	"fmt"
	"notify/internal/enum"
	"time"
)

type NotificationTemplate struct {
	ID        int64
	EventType commonenum.EventType
	Channel   enum.NotificationChannel
	Language  enum.Language
	Title     string
	Content   string
	Enable    bool
	CreatedAt *time.Time
	UpdatedAt *time.Time
}

func (n *NotificationTemplate) GetKey() string {
	return fmt.Sprintf("%s_%s", n.EventType, n.Channel)
}
