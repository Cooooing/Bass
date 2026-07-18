package repo

import (
	"context"
	"game_town/internal/biz/model"
)

type WorldMetricDefinitionRepo interface {
	ListWorldMetrics(ctx context.Context, worldID int64) ([]*model.WorldMetricDefinition, error)
}
