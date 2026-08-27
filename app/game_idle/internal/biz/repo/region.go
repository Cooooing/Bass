package repo

import (
	"context"
	"game_idle/internal/biz/model"
)

// RegionRepo 管理前端展示区域配置缓存。
type RegionRepo interface {
	Refresh(ctx context.Context) ([]*model.Region, error)
	Map(ctx context.Context, regionIDs []string) (map[string]*model.Region, error)
}
