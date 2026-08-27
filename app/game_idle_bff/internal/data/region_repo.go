package data

import (
	"common/pkg/client/rpc"
	gameidlev1 "common/proto/gen/game_idle/v1"
	"context"
	"game_idle_bff/internal/biz/model"
	"game_idle_bff/internal/biz/repo"
)

var _ repo.RegionRepo = (*RegionRepo)(nil)

type RegionRepo struct {
	gameIdleClient *rpc.GameIdleClient
}

func NewRegionRepo(gameIdleClient *rpc.GameIdleClient) repo.RegionRepo {
	return &RegionRepo{
		gameIdleClient: gameIdleClient,
	}
}

func (r *RegionRepo) List(ctx context.Context) ([]*model.RegionConfig, error) {
	reply, err := r.gameIdleClient.Region.List(ctx, &gameidlev1.ListRegions_Request{})
	if err != nil {
		return nil, err
	}
	rows := make([]*model.RegionConfig, 0, len(reply.GetRows()))
	for _, row := range reply.GetRows() {
		rows = append(rows, &model.RegionConfig{
			RegionID:    row.GetRegionId(),
			Name:        row.GetName(),
			Description: row.GetDescription(),
			ActionKind:  row.GetActionKind(),
			Enabled:     row.GetEnabled(),
			Sort:        row.GetSort(),
		})
	}
	return rows, nil
}
