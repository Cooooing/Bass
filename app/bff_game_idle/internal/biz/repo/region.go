package repo

import (
	"context"

	"bff_game_idle/internal/biz/model"
)

// RegionRepo 查询区域配置。
type RegionRepo interface {
	List(ctx context.Context) ([]*model.RegionConfig, error)
}
