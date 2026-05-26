package model

import "time"

type NotificationRecord struct {
	ID               int64
	NotificationID   int64
	ReceiverID       int64
	ReadTime         *time.Time
	CreatedAt        *time.Time
	UpdatedAt        *time.Time
	NotificationMeta *NotificationMeta
}
