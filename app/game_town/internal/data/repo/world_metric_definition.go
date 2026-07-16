package repo

import (
	"context"
	"game_town/internal/biz/model"
	bizrepo "game_town/internal/biz/repo"
	"game_town/internal/data/gen"
	"game_town/internal/data/gen/worldmetricdefinition"
)

type WorldMetricDefinitionRepo struct{ *baseRepo }

func NewWorldMetricDefinitionRepo(db *gen.Client) bizrepo.WorldMetricDefinitionRepo {
	return &WorldMetricDefinitionRepo{baseRepo: &baseRepo{db: db}}
}

func (r *WorldMetricDefinitionRepo) ListWorldMetrics(ctx context.Context, req *bizrepo.ListWorldMetricsReq) (*bizrepo.ListWorldMetricsResponse, error) {
	rows, err := r.db.WorldMetricDefinition.Query().Where(worldmetricdefinition.WorldID(req.WorldID)).Order(worldmetricdefinition.ByID()).All(ctx)
	if err != nil {
		return nil, err
	}
	result := make([]*model.WorldMetricDefinition, 0, len(rows))
	for _, row := range rows {
		result = append(result, r.metric(row))
	}
	return &bizrepo.ListWorldMetricsResponse{Rows: result}, nil
}
