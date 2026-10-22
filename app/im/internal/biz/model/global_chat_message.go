package model

import (
	"time"

	"im/internal/enum"
)

type GlobalChatMessage struct {
	ID              int64
	SenderID        int64
	ClientMessageID string
	Content         string
	Type            enum.MessageType
	Revoked         bool
	RevokedAt       *time.Time
	RevokedBy       *int64
	CreatedAt       *time.Time
}
