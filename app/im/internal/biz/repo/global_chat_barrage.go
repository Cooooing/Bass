package repo

import (
	"context"
	"im/internal/biz/model"
)

type GlobalChatBarrageRepo interface {
	Get(context.Context, int64) (*model.GlobalChatBarrage, error)
	NextID(context.Context) (int64, error)
	CreatePending(context.Context, *model.GlobalChatBarrage) (*model.GlobalChatBarrage, error)
	Publish(context.Context, int64) (*model.GlobalChatBarrage, error)
	Cancel(context.Context, int64) error
	List(context.Context, *int64, int, int) ([]*model.GlobalChatBarrage, int64, error)
}
