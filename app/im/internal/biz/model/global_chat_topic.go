package model

import (
	"time"

	"im/internal/enum"
)

// GlobalChatTopic 是一次全局聊天室话题设置的业务事实。
type GlobalChatTopic struct {
	ID              int64
	UserID          int64
	ClientRequestID string
	Content         string
	PointsCost      int64
	Status          enum.GlobalChatTopicStatus
	CreatedAt       *time.Time
}
