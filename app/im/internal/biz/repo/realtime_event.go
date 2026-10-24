package repo

import (
	"context"

	"im/internal/biz/model"
)

type RealtimeEventPublisher interface {
	PublishPrivateChatMessageCreated(context.Context, *model.PrivateChatMessage) error
	PublishGlobalChatMessageCreated(context.Context, *model.GlobalChatMessage) error
	PublishGlobalChatBarrageCreated(context.Context, *model.GlobalChatBarrage) error
}
