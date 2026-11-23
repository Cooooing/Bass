package data

import (
	"common/pkg/client/rpc"
	gameidlev1 "common/proto/gen/game_idle/v1"
	"context"
	"time"

	"bff_game_idle/internal/biz/model"
	"bff_game_idle/internal/biz/repo"
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
	reply, err := r.gameIdleClient.Character.CreateCharacter(ctx, &gameidlev1.CreateCharacter_Req{
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
	if row.GetCreatedAt() != nil {
		out.CreatedAt = new(row.GetCreatedAt().AsTime())
	}
	if row.GetUpdatedAt() != nil {
		out.UpdatedAt = new(row.GetUpdatedAt().AsTime())
	}
	if row.GetLastOfflineAt() != nil {
		out.LastOfflineAt = new(row.GetLastOfflineAt().AsTime())
	}
	return out, nil
}

func (r *CharacterRepo) List(ctx context.Context, req *repo.ListCharacterReq) ([]*model.Character, error) {
	reply, err := r.gameIdleClient.Character.GetCharacter(ctx, &gameidlev1.GetCharacter_Req{
		UserId:      req.UserID,
		CharacterId: req.CharacterID,
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
		if row.GetCreatedAt() != nil {
			item.CreatedAt = new(row.GetCreatedAt().AsTime())
		}
		if row.GetUpdatedAt() != nil {
			item.UpdatedAt = new(row.GetUpdatedAt().AsTime())
		}
		if row.GetLastOfflineAt() != nil {
			item.LastOfflineAt = new(row.GetLastOfflineAt().AsTime())
		}
		rows = append(rows, item)
	}
	return rows, nil
}

func (r *CharacterRepo) Online(ctx context.Context, req *repo.OnlineCharacterReq) (*model.CharacterOnlineSession, error) {
	reply, err := r.gameIdleClient.Character.OnlineCharacter(ctx, &gameidlev1.OnlineCharacter_Req{
		CharacterId: req.CharacterID,
	})
	if err != nil {
		return nil, err
	}
	return &model.CharacterOnlineSession{
		CharacterID:       req.CharacterID,
		SessionID:         reply.GetSessionId(),
		RemainingDuration: time.Duration(reply.GetExpiresInSeconds()) * time.Second,
	}, nil
}

func (r *CharacterRepo) Ping(ctx context.Context, req *repo.PingCharacterReq) (*model.CharacterOnlineSession, error) {
	reply, err := r.gameIdleClient.Character.PingCharacter(ctx, &gameidlev1.PingCharacter_Req{
		CharacterId: req.CharacterID,
		SessionId:   req.SessionID,
	})
	if err != nil {
		return nil, err
	}
	return &model.CharacterOnlineSession{
		CharacterID:       req.CharacterID,
		SessionID:         req.SessionID,
		RemainingDuration: time.Duration(reply.GetExpiresInSeconds()) * time.Second,
	}, nil
}

func (r *CharacterRepo) Offline(ctx context.Context, req *repo.OfflineCharacterReq) error {
	_, err := r.gameIdleClient.Character.OfflineCharacter(ctx, &gameidlev1.OfflineCharacter_Req{
		CharacterId: req.CharacterID,
		SessionId:   req.SessionID,
		Timeout:     req.Timeout,
	})
	return err
}
