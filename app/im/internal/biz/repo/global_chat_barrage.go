package repo

import (
	"context"
	"im/internal/biz/model"
)

type GlobalChatBarrageRepo interface {
	GetByClientRequestID(context.Context, string) (*model.GlobalChatBarrage, error)
	CreatePending(context.Context, *model.GlobalChatBarrage) (*model.GlobalChatBarrage, error)
	Publish(context.Context, string) (*model.GlobalChatBarrage, error)
	Cancel(context.Context, string) error
	List(context.Context, *int64, int, int) ([]*model.GlobalChatBarrage, int64, error)
}
