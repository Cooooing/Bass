package repo

import (
	"context"
	"game_town/internal/biz/model"
)

type WorldMetricDefinitionRepo interface {
	ListWorldMetrics(ctx context.Context, req *ListWorldMetricsReq) (*ListWorldMetricsResponse, error)
}

type ListWorldMetricsReq struct {
	WorldID int64
}

type ListWorldMetricsResponse struct {
	Rows []*model.WorldMetricDefinition
}
