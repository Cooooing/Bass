package model

import "time"

type ChatSession struct {
	ID                int64
	ReceiverID        *int64
	GroupID           *int64
	IsMuted           bool
	IsPinned          bool
	LastReadMessageID *int64
	ReadCount         uint32
	MessageCount      uint32
	LastMessageID     *int64
	CreatedAt         *time.Time
	UpdatedAt         *time.Time
	CreatedBy         *int64
	UpdatedBy         *int64

	Group       *ChatGroup
	UnreadCount uint32
}
