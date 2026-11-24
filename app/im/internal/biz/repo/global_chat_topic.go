package repo

import (
	"context"

	"im/internal/biz/model"
)

// GlobalChatTopicRepo 管理全局聊天室话题设置事实。
type GlobalChatTopicRepo interface {
	GetCurrent(ctx context.Context) (*model.GlobalChatTopic, error)
	Get(ctx context.Context, topicID int64) (*model.GlobalChatTopic, error)
	NextID(ctx context.Context) (int64, error)
	CreatePending(ctx context.Context, topic *model.GlobalChatTopic) (*model.GlobalChatTopic, error)
	Publish(ctx context.Context, topicID int64) (*model.GlobalChatTopic, error)
	Cancel(ctx context.Context, topicID int64) error
}
