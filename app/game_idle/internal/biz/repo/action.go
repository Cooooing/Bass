package repo

import (
	"context"
	"game_idle/internal/biz/model"
)

// ActionRepo 管理可执行行动配置缓存。
type ActionRepo interface {
	Refresh(ctx context.Context) ([]*model.Action, error)
	Get(ctx context.Context, actionID string) (*model.Action, error)
	Map(ctx context.Context, actionIDs []string) (map[string]*model.Action, error)
}
