package data

import (
	"bff_game_idle/internal/biz/model"
	"bff_game_idle/internal/biz/repo"
	"common/pkg/client/rpc"
	gameidlev1 "common/proto/gen/game_idle/v1"
	"context"
)

var _ repo.BackpackRepo = (*BackpackRepo)(nil)

type BackpackRepo struct {
	gameIdleClient *rpc.GameIdleClient
}

func NewBackpackRepo(
	gameIdleClient *rpc.GameIdleClient,
) repo.BackpackRepo {
	return &BackpackRepo{
		gameIdleClient: gameIdleClient,
	}
}

func (r *BackpackRepo) Map(ctx context.Context, req *repo.BackpackMapReq) (map[string]*model.CharacterItem, error) {
	reply, err := r.gameIdleClient.Backpack.Get(ctx, &gameidlev1.GetBackpack_Request{
		CharacterId: req.CharacterID,
		ItemIds:     req.ItemIDs,
	})
	if err != nil {
		return nil, err
	}
	rows := make(map[string]*model.CharacterItem, len(reply.GetRows()))
	for _, row := range reply.GetRows() {
		rows[row.GetItemId()] = &model.CharacterItem{
			ItemID:   row.GetItemId(),
			Quantity: row.GetQuantity(),
		}
	}
	return rows, nil
}
