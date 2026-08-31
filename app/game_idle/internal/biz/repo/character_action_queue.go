package repo

import (
	"context"
	"game_idle/internal/biz/model"
)

// CharacterActionQueueRepo 管理玩家运行时行动队列；本地缓存承载热状态，数据库保存冷备快照。
type CharacterActionQueueRepo interface {
	ListCharacterIDs(ctx context.Context) ([]int64, error)
	Load(ctx context.Context, characterID int64) (*model.CharacterActionQueue, error)
	Save(ctx context.Context, queue *model.CharacterActionQueue) error
	Persist(ctx context.Context, characterID int64) error
}
