package repo

import (
	"context"
	"game_idle/internal/biz/model"
	"game_idle/internal/enum"
)

// ChatMessageRepo 管理游戏内部聊天消息。
type ChatMessageRepo interface {
	Create(ctx context.Context, message *model.ChatMessage) (*model.ChatMessage, error)
	List(ctx context.Context, req *ChatMessageListReq) ([]*model.ChatMessage, error)
}

// ChatMessageEventRepo 发布聊天消息事件。
type ChatMessageEventRepo interface {
	PublishWorldMessage(ctx context.Context, message *model.ChatMessage) error
}

type ChatMessageListReq struct {
	ChannelType enum.ChatChannelType
	ChannelID   string
	BeforeID    int64
	Size        int
}
