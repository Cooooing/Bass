package data

import (
	"common/pkg/client/rpc"
	gameidlev1 "common/proto/gen/game_idle/v1"
	"context"
	"game_idle_bff/internal/biz/model"
	"game_idle_bff/internal/biz/repo"
)

var _ repo.ItemRepo = (*ItemRepo)(nil)

type ItemRepo struct {
	gameIdleClient *rpc.GameIdleClient
}

func NewItemRepo(gameIdleClient *rpc.GameIdleClient) repo.ItemRepo {
	return &ItemRepo{
		gameIdleClient: gameIdleClient,
	}
}

func (r *ItemRepo) List(ctx context.Context) ([]*model.ItemConfig, error) {
	reply, err := r.gameIdleClient.Item.List(ctx, &gameidlev1.ListItems_Request{})
	if err != nil {
		return nil, err
	}
	rows := make([]*model.ItemConfig, 0, len(reply.GetRows()))
	for _, row := range reply.GetRows() {
		rows = append(rows, &model.ItemConfig{
			ItemID:      row.GetItemId(),
			Name:        row.GetName(),
			ItemType:    row.GetItemType(),
			Category:    row.GetItemType(),
			Description: row.GetDescription(),
			Enabled:     row.GetEnabled(),
			Sort:        row.GetSort(),
		})
	}
	return rows, nil
}
