package usecase

import (
	"context"
	"strings"

	"common/pkg/apperror"
	"game_town/internal/biz/base"
	"game_town/internal/biz/model"
	"game_town/internal/biz/repo"
	"game_town/internal/enum"
)

type NpcUsecase struct {
	tx              base.Tx
	npcRepo         repo.NpcRepo
	worldMemberRepo repo.WorldMemberRepo
	eventUsecase    *EventUsecase
}

func NewNpcUsecase(
	tx base.Tx,
	npcRepo repo.NpcRepo,
	worldMemberRepo repo.WorldMemberRepo,
	eventUsecase *EventUsecase,
) *NpcUsecase {
	return &NpcUsecase{
		tx:              tx,
		npcRepo:         npcRepo,
		worldMemberRepo: worldMemberRepo,
		eventUsecase:    eventUsecase,
	}
}

func (u *NpcUsecase) Get(
	ctx context.Context,
	npcID int64,
) (*model.Npc, error) {
	return u.npcRepo.Get(ctx, &repo.NpcQuery{
		ID: new(npcID),
	})
}

type ListNpcsReq struct {
	WorldID    int64
	LocationID *int64
}

func (u *NpcUsecase) List(
	ctx context.Context,
	req *ListNpcsReq,
) ([]*model.Npc, error) {
	return u.npcRepo.List(ctx, &repo.NpcQuery{
		WorldID:    new(req.WorldID),
		LocationID: req.LocationID,
	})
}

type TalkNpcReq struct {
	WorldID  int64
	PlayerID int64
	NpcID    int64
	Content  string
}

func (u *NpcUsecase) Talk(
	ctx context.Context,
	req *TalkNpcReq,
) (*model.Event, error) {
	content := strings.TrimSpace(req.Content)
	if content == "" {
		return nil, apperror.CommonInvalidArgument()
	}
	member, err := u.worldMemberRepo.Get(ctx, &repo.WorldMemberQuery{
		WorldID:  new(req.WorldID),
		PlayerID: new(req.PlayerID),
	})
	if err != nil {
		return nil, err
	}
	npc, err := u.npcRepo.Get(ctx, &repo.NpcQuery{
		ID:      new(req.NpcID),
		WorldID: new(req.WorldID),
	})
	if err != nil {
		return nil, err
	}
	if npc.CurrentLocationID != member.CurrentLocationID {
		return nil, apperror.GameTownWorldInvalidMessage("NPC 不在当前场景，请先使用 /look 查看其位置并移动")
	}
	var event *model.Event
	err = u.tx(ctx, func(ctx context.Context) error {
		var err error
		event, err = u.eventUsecase.AppendInTx(ctx, &AppendEventReq{
			WorldID:       req.WorldID,
			Type:          enum.EventTypePlayerTalked,
			ActorPlayerID: new(req.PlayerID),
			NpcID:         new(req.NpcID),
			LocationID:    new(member.CurrentLocationID),
			Summary:       "玩家与 NPC 对话",
			Content:       content,
		})
		return err
	})
	if err != nil {
		return nil, err
	}
	u.eventUsecase.Publish(event)
	return event, nil
}
