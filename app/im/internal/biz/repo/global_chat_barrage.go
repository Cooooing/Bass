package repo

import (
	"context"

	"im/internal/biz/model"
)

type GlobalChatBarrageRepo interface {
	Get(ctx context.Context, barrageID int64) (*model.GlobalChatBarrage, error)
	NextID(ctx context.Context) (int64, error)
	CreatePending(ctx context.Context, barrage *model.GlobalChatBarrage) (*model.GlobalChatBarrage, error)
	Publish(ctx context.Context, barrageID int64) (*model.GlobalChatBarrage, error)
	Cancel(ctx context.Context, barrageID int64) error
	List(ctx context.Context, afterID *int64, limit, offset int) ([]*model.GlobalChatBarrage, int64, error)
}
