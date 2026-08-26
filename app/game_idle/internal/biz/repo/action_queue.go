package repo

import (
	"context"
	"game_idle/internal/biz/model"
)

// ActionQueueRepo 管理玩家运行时行动队列；数据层使用 Redis 承载缓存和持久化，不写业务数据库。
type ActionQueueRepo interface {
	ListCharacterIDs(ctx context.Context) ([]int64, error)
	Load(ctx context.Context, characterID int64) (*model.ActionQueue, error)
	Save(ctx context.Context, queue *model.ActionQueue) error
}
