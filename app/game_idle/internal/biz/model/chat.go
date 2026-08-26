package model

import (
	"game_idle/internal/enum"
	"time"
)

// ChatMessage 是挂机游戏内部聊天消息。
type ChatMessage struct {
	ID                  int64
	ChannelType         enum.ChatChannelType
	ChannelID           string
	SenderCharacterID   int64
	ReceiverCharacterID *int64
	Content             string
	Status              enum.ChatMessageStatus
	CreatedAt           *time.Time
	UpdatedAt           *time.Time
	DeletedAt           *time.Time
}
