package repo

import (
	"context"

	"im/internal/biz/model"
)

// GlobalChatTopicRepo 管理全局聊天室话题设置事实。
type GlobalChatTopicRepo interface {
	GetCurrent(context.Context) (*model.GlobalChatTopic, error)
	Get(context.Context, int64) (*model.GlobalChatTopic, error)
	NextID(context.Context) (int64, error)
	CreatePending(context.Context, *model.GlobalChatTopic) (*model.GlobalChatTopic, error)
	Publish(context.Context, int64) (*model.GlobalChatTopic, error)
	Cancel(context.Context, int64) error
}
