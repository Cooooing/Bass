package repo

import (
	"context"
	"game_idle/internal/biz/model"
)

// ItemRepo 管理物品配置缓存，构造时会从配置表全量初始化。
type ItemRepo interface {
	Refresh(ctx context.Context) ([]*model.Item, error)
	Get(ctx context.Context, itemID string) (*model.Item, error)
	Map(ctx context.Context, itemIDs []string) (map[string]*model.Item, error)
}
