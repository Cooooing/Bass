package repo

import (
	"context"
	"game_idle/internal/biz/model"
)

// MetaActionRepo 管理可执行行动配置缓存。
type MetaActionRepo interface {
	Refresh(ctx context.Context) error
	Get(ctx context.Context, actionID string) (*model.MetaAction, error)
	Map(ctx context.Context, actionIDs []string) (map[string]*model.MetaAction, error)
}
