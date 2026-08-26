package data

import (
	"common/pkg/client/rpc"
	gameidlev1 "common/proto/gen/game_idle/v1"
	"context"
	"game_idle_bff/internal/biz/model"
	"game_idle_bff/internal/biz/repo"
)

var _ repo.CharacterRepo = (*CharacterRepo)(nil)

type CharacterRepo struct {
	gameIdleClient *rpc.GameIdleClient
}

func NewCharacterRepo(
	gameIdleClient *rpc.GameIdleClient,
) repo.CharacterRepo {
	return &CharacterRepo{
		gameIdleClient: gameIdleClient,
	}
}

func (r *CharacterRepo) Create(ctx context.Context, req *repo.CreateCharacterReq) (*model.Character, error) {
	reply, err := r.gameIdleClient.Character.Create(ctx, &gameidlev1.CreateGameIdleCharacter_Request{
		UserId: req.UserID,
		Name:   req.Name,
	})
	if err != nil {
		return nil, err
	}
	row := reply.GetRow()
	out := &model.Character{
		ID:                  row.GetId(),
		Name:                row.GetName(),
		Status:              row.GetStatus(),
		Slot:                row.GetSlot(),
		ActionQueueCapacity: row.GetActionQueueCapacity(),
		MaxOfflineSeconds:   row.GetMaxOfflineSeconds(),
	}
	if row.CreatedAt != nil {
		out.CreatedAt = new(row.CreatedAt.AsTime())
	}
	if row.UpdatedAt != nil {
		out.UpdatedAt = new(row.UpdatedAt.AsTime())
	}
	return out, nil
}

func (r *CharacterRepo) List(ctx context.Context, userID int64) ([]*model.Character, error) {
	reply, err := r.gameIdleClient.Character.Get(ctx, &gameidlev1.GetGameIdleCharacter_Request{
		UserId: userID,
	})
	if err != nil {
		return nil, err
	}
	rows := make([]*model.Character, 0, len(reply.GetRows()))
	for _, row := range reply.GetRows() {
		item := &model.Character{
			ID:                  row.GetId(),
			Name:                row.GetName(),
			Status:              row.GetStatus(),
			Slot:                row.GetSlot(),
			ActionQueueCapacity: row.GetActionQueueCapacity(),
			MaxOfflineSeconds:   row.GetMaxOfflineSeconds(),
		}
		if row.CreatedAt != nil {
			item.CreatedAt = new(row.CreatedAt.AsTime())
		}
		if row.UpdatedAt != nil {
			item.UpdatedAt = new(row.UpdatedAt.AsTime())
		}
		rows = append(rows, item)
	}
	return rows, nil
}
