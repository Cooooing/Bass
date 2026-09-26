package repo

import (
	"bff_game_idle/internal/biz/model"
	"context"
)

// ActionRepo 查询行动配置。
type ActionRepo interface {
	List(ctx context.Context) ([]*model.ActionConfig, error)
	GetDetail(ctx context.Context, actionID string) (*model.ActionDetailConfig, error)
}
