package repo

import (
	"context"

	"im/internal/biz/model"
)

type RealtimeEventPublisher interface {
	Ensure(ctx context.Context) error
	PublishPrivateChatMessageCreated(ctx context.Context, message *model.PrivateChatMessage) error
	PublishGlobalChatMessageCreated(ctx context.Context, message *model.GlobalChatMessage) error
	PublishGlobalChatBarrageCreated(ctx context.Context, barrage *model.GlobalChatBarrage) error
	PublishGlobalChatTopicUpdated(ctx context.Context, topic *model.GlobalChatTopic) error
}
