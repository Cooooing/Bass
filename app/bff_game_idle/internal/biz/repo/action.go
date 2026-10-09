package repo

import (
	"context"

	"bff_game_idle/internal/biz/model"
)

// ActionRepo 查询行动配置。
type ActionRepo interface {
	List(ctx context.Context) ([]*model.ActionConfig, error)
	GetDetail(ctx context.Context, actionID string) (*model.ActionDetailConfig, error)
}
