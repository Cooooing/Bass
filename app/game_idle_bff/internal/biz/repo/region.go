package repo

import (
	"context"
	"game_idle_bff/internal/biz/model"
)

// RegionRepo 查询区域配置。
type RegionRepo interface {
	List(ctx context.Context) ([]*model.RegionConfig, error)
}
