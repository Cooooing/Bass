package model

import (
	"im/internal/enum"
	"time"
)

type ChatMessage struct {
	ID         int64
	SenderID   int64
	ReceiverID *int64
	GroupID    *int64
	SessionID  *int64
	Type       enum.MessageType
	Content    string
	Status     enum.MessageStatus
	CreatedAt  *time.Time
	UpdatedAt  *time.Time
	CreatedBy  *int64
	UpdatedBy  *int64

	SenderUser   *AccountBasic
	ReceiverUser *AccountBasic
}
