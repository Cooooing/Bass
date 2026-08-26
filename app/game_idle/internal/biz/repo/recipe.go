package repo

import (
	"context"
	"game_idle/internal/biz/model"
)

// RecipeRepo 管理配方配置缓存；刷新时需要同时载入输入和输出明细。
type RecipeRepo interface {
	Refresh(ctx context.Context) ([]*model.Recipe, error)
	Get(ctx context.Context, recipeID string) (*model.Recipe, error)
	Map(ctx context.Context, recipeIDs []string) (map[string]*model.Recipe, error)
}
