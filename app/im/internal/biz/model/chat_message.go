package model

import (
	"time"

	userv1 "common/proto/gen/user/v1"
	"im/internal/enum"
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

	SenderUser   *userv1.AccountBasic
	ReceiverUser *userv1.AccountBasic
}
