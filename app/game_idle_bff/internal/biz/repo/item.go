package repo

import (
	"context"
	"game_idle_bff/internal/biz/model"
)

// ItemRepo 查询物品配置。
type ItemRepo interface {
	List(ctx context.Context) ([]*model.ItemConfig, error)
}
