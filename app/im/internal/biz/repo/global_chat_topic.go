package repo

import (
	"context"

	"im/internal/biz/model"
)

// GlobalChatTopicRepo 管理全局聊天室话题设置事实。
type GlobalChatTopicRepo interface {
	GetCurrent(context.Context) (*model.GlobalChatTopic, error)
	GetByClientRequestID(context.Context, string) (*model.GlobalChatTopic, error)
	CreatePending(context.Context, *model.GlobalChatTopic) (*model.GlobalChatTopic, error)
	Publish(context.Context, string) (*model.GlobalChatTopic, error)
	Cancel(context.Context, string) error
}
