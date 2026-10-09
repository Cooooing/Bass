package repo

import (
	"context"

	"bff_game_idle/internal/biz/model"
)

// ItemRepo 查询物品配置。
type ItemRepo interface {
	List(ctx context.Context) ([]*model.ItemConfig, error)
}
