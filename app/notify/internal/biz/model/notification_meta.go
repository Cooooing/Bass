package model

import (
	"notify/internal/enum"
	"time"
)

type NotificationMeta struct {
	ID        int64
	Title     string
	Content   string
	Status    enum.NotificationStatus
	CreatedAt *time.Time
	UpdatedAt *time.Time
}
