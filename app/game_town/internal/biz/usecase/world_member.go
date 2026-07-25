package usecase

import (
	"context"
	"strings"
	"time"

	"common/pkg/apperror"
	"game_town/internal/biz/base"
	"game_town/internal/biz/model"
	"game_town/internal/biz/repo"
	"game_town/internal/enum"
)

type WorldMemberUsecase struct {
	tx              base.Tx
	playerRepo      repo.PlayerRepo
	worldRepo       repo.WorldRepo
	worldMemberRepo repo.WorldMemberRepo
	worldStateRepo  repo.WorldStateRepo
	eventUsecase    *EventUsecase
}

func NewWorldMemberUsecase(
	tx base.Tx,
	playerRepo repo.PlayerRepo,
	worldRepo repo.WorldRepo,
	worldMemberRepo repo.WorldMemberRepo,
	worldStateRepo repo.WorldStateRepo,
	eventUsecase *EventUsecase,
) *WorldMemberUsecase {
	return &WorldMemberUsecase{
		tx:              tx,
		playerRepo:      playerRepo,
		worldRepo:       worldRepo,
		worldMemberRepo: worldMemberRepo,
		worldStateRepo:  worldStateRepo,
		eventUsecase:    eventUsecase,
	}
}

type JoinWorldReq struct {
	PlayerID            int64
	WorldCode           string
	CharacterPreference string
}

type JoinWorldResp struct {
	WorldID    int64
	MemberID   int64
	LocationID int64
	EventID    int64
}

func (u *WorldMemberUsecase) Join(ctx context.Context, req *JoinWorldReq) (*JoinWorldResp, error) {
	worldCode := strings.TrimSpace(req.WorldCode)
	characterPreference := strings.TrimSpace(req.CharacterPreference)
	if req.PlayerID <= 0 || worldCode == "" {
		return nil, apperror.CommonInvalidArgument()
	}

	if _, err := u.playerRepo.Get(ctx, &repo.PlayerQuery{
		ID: new(req.PlayerID),
	}); err != nil {
		return nil, err
	}

	world, err := u.worldRepo.Get(ctx, &repo.WorldQuery{
		Code: new(worldCode),
	})
	if err != nil {
		return nil, err
	}
	if world.Status != enum.WorldStatusActive || world.DefaultLocationID == nil {
		return nil, apperror.GameTownWorldInvalid()
	}

	query := &repo.WorldMemberQuery{
		WorldID:  new(world.ID),
		PlayerID: new(req.PlayerID),
	}
	count, err := u.worldMemberRepo.Count(ctx, query)
	if err != nil {
		return nil, err
	}
	if count > 0 {
		member, err := u.worldMemberRepo.Get(ctx, query)
		if err != nil {
			return nil, err
		}
		return &JoinWorldResp{
			WorldID:    world.ID,
			MemberID:   member.ID,
			LocationID: member.CurrentLocationID,
		}, nil
	}

	var member *model.WorldMember
	var event *model.Event
	err = u.tx(ctx, func(ctx context.Context) error {
		now := time.Now()
		var err error
		member, err = u.worldMemberRepo.Save(ctx, &model.WorldMember{
			WorldID:             world.ID,
			PlayerID:            req.PlayerID,
			CurrentLocationID:   *world.DefaultLocationID,
			Role:                enum.WorldMemberRoleMember,
			CharacterPreference: characterPreference,
			CharacterTraits:     []string{},
			JoinedAt:            now,
			LastSeenAt:          now,
		})
		if err != nil {
			return err
		}

		event, err = u.eventUsecase.AppendInTx(ctx, &AppendEventReq{
			WorldID:       world.ID,
			Type:          enum.EventTypePlayerJoined,
			ActorPlayerID: new(req.PlayerID),
			LocationID:    world.DefaultLocationID,
			Summary:       "玩家加入世界，等待世界生成角色",
			Payload: map[string]any{
				"public":               true,
				"character_preference": characterPreference,
			},
		})
		return err
	})
	if err != nil {
		return nil, err
	}

	u.eventUsecase.Publish(event)
	return &JoinWorldResp{
		WorldID:    world.ID,
		MemberID:   member.ID,
		LocationID: member.CurrentLocationID,
		EventID:    event.ID,
	}, nil
}

type GetWorldMemberReq struct {
	WorldID  int64
	PlayerID int64
}

type GetWorldMemberResp struct {
	Member *model.WorldMember
	State  *model.WorldState
}

func (u *WorldMemberUsecase) Get(ctx context.Context, req *GetWorldMemberReq) (*GetWorldMemberResp, error) {
	member, err := u.worldMemberRepo.Get(ctx, &repo.WorldMemberQuery{
		WorldID:  new(req.WorldID),
		PlayerID: new(req.PlayerID),
	})
	if err != nil {
		return nil, err
	}

	state, err := u.worldStateRepo.Get(ctx, &repo.WorldStateQuery{
		WorldID: new(req.WorldID),
	})
	if err != nil {
		return nil, err
	}

	return &GetWorldMemberResp{
		Member: member,
		State:  state,
	}, nil
}

type SubmitActionTarget struct {
	Type enum.EntityType
	ID   int64
}

type SubmitActionReq struct {
	WorldID         int64
	PlayerID        int64
	Content         string
	Targets         []SubmitActionTarget
	ClientRequestID string
}

func (u *WorldMemberUsecase) SubmitAction(ctx context.Context, req *SubmitActionReq) (*model.Event, error) {
	content := strings.TrimSpace(req.Content)
	if req.WorldID <= 0 || req.PlayerID <= 0 || content == "" {
		return nil, apperror.CommonInvalidArgument()
	}

	member, err := u.worldMemberRepo.Get(ctx, &repo.WorldMemberQuery{
		WorldID:  new(req.WorldID),
		PlayerID: new(req.PlayerID),
	})
	if err != nil {
		return nil, err
	}

	targets, npcID := u.actionPayloadTargets(req.Targets)
	var event *model.Event
	err = u.tx(ctx, func(ctx context.Context) error {
		var err error
		event, err = u.eventUsecase.AppendInTx(ctx, &AppendEventReq{
			WorldID:       req.WorldID,
			Type:          enum.EventTypePlayerActionSubmitted,
			ActorPlayerID: new(req.PlayerID),
			NpcID:         npcID,
			LocationID:    new(member.CurrentLocationID),
			Summary:       "玩家提交行动",
			Content:       content,
			Payload: map[string]any{
				"targets":           targets,
				"client_request_id": strings.TrimSpace(req.ClientRequestID),
			},
		})
		return err
	})
	if err != nil {
		return nil, err
	}

	u.eventUsecase.Publish(event)
	return event, nil
}

func (u *WorldMemberUsecase) actionPayloadTargets(rows []SubmitActionTarget) ([]any, *int64) {
	targets := make([]any, 0, len(rows))
	var npcID *int64
	for _, target := range rows {
		if target.ID <= 0 {
			continue
		}
		targets = append(targets, map[string]any{
			"type": string(target.Type),
			"id":   target.ID,
		})
		if npcID == nil && target.Type == enum.EntityTypeNpc {
			npcID = new(target.ID)
		}
	}
	return targets, npcID
}
