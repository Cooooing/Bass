package data

import (
	"common/pkg/client/rpc"
	gameidlev1 "common/proto/gen/game_idle/v1"
	"context"

	"bff_game_idle/internal/biz/model"
	"bff_game_idle/internal/biz/repo"
)

var _ repo.ActionQueueRepo = (*ActionQueueRepo)(nil)

type ActionQueueRepo struct {
	gameIdleClient *rpc.GameIdleClient
}

func NewActionQueueRepo(
	gameIdleClient *rpc.GameIdleClient,
) repo.ActionQueueRepo {
	return &ActionQueueRepo{
		gameIdleClient: gameIdleClient,
	}
}

func (r *ActionQueueRepo) List(ctx context.Context, characterID int64) (*model.ActionQueue, error) {
	reply, err := r.gameIdleClient.ActionQueue.ListActionQueue(ctx, &gameidlev1.ListActionQueue_Req{
		CharacterId: characterID,
	})
	if err != nil {
		return nil, err
	}
	queue := reply.GetQueue()
	out := &model.ActionQueue{
		CharacterID: queue.GetCharacterId(),
		Items:       make([]*model.ActionQueueItem, 0, len(queue.GetItems())),
	}
	for _, item := range queue.GetItems() {
		out.Items = append(out.Items, &model.ActionQueueItem{
			ActionID:  item.GetActionId(),
			Times:     item.GetTimes(),
			CreatedAt: item.GetCreatedAt().AsTime(),
		})
	}
	return out, nil
}

func (r *ActionQueueRepo) Add(ctx context.Context, req *repo.AddActionReq) error {
	_, err := r.gameIdleClient.ActionQueue.AddAction(ctx, &gameidlev1.AddAction_Req{
		CharacterId: req.CharacterID,
		ActionId:    req.ActionID,
		Times:       req.Times,
		Position:    req.Position,
	})
	return err
}

func (r *ActionQueueRepo) Move(ctx context.Context, req *repo.MoveActionReq) error {
	_, err := r.gameIdleClient.ActionQueue.MoveAction(ctx, &gameidlev1.MoveAction_Req{
		CharacterId:     req.CharacterID,
		CurrentPosition: req.CurrentPosition,
		TargetPosition:  req.TargetPosition,
	})
	return err
}

func (r *ActionQueueRepo) Remove(ctx context.Context, req *repo.RemoveActionReq) error {
	_, err := r.gameIdleClient.ActionQueue.RemoveAction(ctx, &gameidlev1.RemoveAction_Req{
		CharacterId: req.CharacterID,
		Position:    req.Position,
	})
	return err
}

func (r *ActionQueueRepo) Clear(ctx context.Context, characterID int64) error {
	_, err := r.gameIdleClient.ActionQueue.ClearActionQueue(ctx, &gameidlev1.ClearActionQueue_Req{
		CharacterId: characterID,
	})
	return err
}
