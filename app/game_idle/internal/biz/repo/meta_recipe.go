package repo

import (
	"context"
	"game_idle/internal/biz/model"
)

// MetaRecipeRepo 管理配方配置缓存；刷新时需要同时载入输入和输出明细。
type MetaRecipeRepo interface {
	Refresh(ctx context.Context) error
	Get(ctx context.Context, recipeID string) (*model.MetaRecipe, error)
	Map(ctx context.Context, recipeIDs []string) (map[string]*model.MetaRecipe, error)
}
