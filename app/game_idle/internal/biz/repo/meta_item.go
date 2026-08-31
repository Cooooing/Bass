package repo

import (
	"context"
	"game_idle/internal/biz/model"
)

// MetaItemRepo 管理物品配置缓存，构造时会从配置表全量初始化。
type MetaItemRepo interface {
	Refresh(ctx context.Context) error
	Get(ctx context.Context, itemID string) (*model.MetaItem, error)
	Map(ctx context.Context, itemIDs []string) (map[string]*model.MetaItem, error)
}
