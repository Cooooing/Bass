package repo

import (
	"context"

	"im/internal/biz/model"
)

type GlobalChatMessageRepo interface {
	Get(context.Context, int64) (*model.GlobalChatMessage, error)
	GetByClientMessageID(context.Context, int64, string) (*model.GlobalChatMessage, error)
	Save(context.Context, *model.GlobalChatMessage) (*model.GlobalChatMessage, error)
	List(context.Context, int64, int, int) ([]*model.GlobalChatMessage, int64, error)
	Revoke(context.Context, int64, int64) (*model.GlobalChatMessage, error)
}
