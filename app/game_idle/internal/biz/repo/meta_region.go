package repo

import (
	"context"
	"game_idle/internal/biz/model"
)

// MetaRegionRepo 管理前端展示区域配置缓存。
type MetaRegionRepo interface {
	Refresh(ctx context.Context) error
	Map(ctx context.Context, regionIDs []string) (map[string]*model.MetaRegion, error)
}
