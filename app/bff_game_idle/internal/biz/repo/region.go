package repo

import (
	"bff_game_idle/internal/biz/model"
	"context"
)

// RegionRepo 查询区域配置。
type RegionRepo interface {
	List(ctx context.Context) ([]*model.RegionConfig, error)
}
