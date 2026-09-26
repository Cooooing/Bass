package repo

import (
	"bff_game_idle/internal/biz/model"
	"context"
)

// ItemRepo 查询物品配置。
type ItemRepo interface {
	List(ctx context.Context) ([]*model.ItemConfig, error)
}
