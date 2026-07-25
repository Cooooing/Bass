package usecase

import (
	"context"
	"strconv"
	"strings"
	"time"

	"common/pkg/apperror"
	"game_town/internal/biz/base"
	"game_town/internal/biz/model"
	"game_town/internal/biz/repo"
	"game_town/internal/config"
	"game_town/internal/enum"
)

type WorldUsecase struct {
	conf            *config.Bootstrap
	tx              base.Tx
	playerRepo      repo.PlayerRepo
	agentConfigRepo repo.AgentConfigRepo
	worldRepo       repo.WorldRepo
	worldStateRepo  repo.WorldStateRepo
	eventUsecase    *EventUsecase
}

func NewWorldUsecase(
	conf *config.Bootstrap,
	tx base.Tx,
	playerRepo repo.PlayerRepo,
	agentConfigRepo repo.AgentConfigRepo,
	worldRepo repo.WorldRepo,
	worldStateRepo repo.WorldStateRepo,
	eventUsecase *EventUsecase,
) *WorldUsecase {
	return &WorldUsecase{
		conf:            conf,
		tx:              tx,
		playerRepo:      playerRepo,
		agentConfigRepo: agentConfigRepo,
		worldRepo:       worldRepo,
		worldStateRepo:  worldStateRepo,
		eventUsecase:    eventUsecase,
	}
}

type CreateWorldReq struct {
	CreatorPlayerID int64
	Description     string
	NpcCount        uint32
	LocationCount   uint32
	Seed            *int64
	AgentConfigID   int64
}

type CreateWorldResp struct {
	World *model.World
	Event *model.Event
}

func (u *WorldUsecase) Create(ctx context.Context, req *CreateWorldReq) (*CreateWorldResp, error) {
	description := strings.TrimSpace(req.Description)
	if req.CreatorPlayerID <= 0 || req.AgentConfigID <= 0 || description == "" {
		return nil, apperror.GameTownWorldInvalid()
	}
	npcCount := req.NpcCount
	if npcCount == 0 {
		npcCount = 4
	}
	locationCount := req.LocationCount
	if locationCount == 0 {
		locationCount = 4
	}
	if npcCount > u.conf.GetGameTown().GetWorld().GetMaxNpcCount() ||
		locationCount > u.conf.GetGameTown().GetWorld().GetMaxLocationCount() {
		return nil, apperror.GameTownWorldInvalid()
	}
	if _, err := u.playerRepo.Get(ctx, &repo.PlayerQuery{
		ID: new(req.CreatorPlayerID),
	}); err != nil {
		return nil, err
	}
	if _, err := u.agentConfigRepo.Get(ctx, &repo.AgentConfigQuery{
		ID: new(req.AgentConfigID),
	}); err != nil {
		return nil, err
	}
	seed := time.Now().UnixNano()
	if req.Seed != nil {
		seed = *req.Seed
	}
	var world *model.World
	var event *model.Event
	err := u.tx(ctx, func(ctx context.Context) error {
		var err error
		world, err = u.worldRepo.Save(ctx, &model.World{
			Code:            "w" + strconv.FormatInt(time.Now().UnixNano(), 36),
			Description:     description,
			Status:          enum.WorldStatusGenerating,
			CreatorPlayerID: req.CreatorPlayerID,
			AgentConfigID:   req.AgentConfigID,
			Seed:            seed,
		})
		if err != nil {
			return err
		}
		nextTickAt := time.Now().Add(u.conf.GetGameTown().GetAgent().GetTickInterval().AsDuration())
		if _, err = u.worldStateRepo.Save(ctx, &model.WorldState{
			WorldID:    world.ID,
			Summary:    description,
			CurrentArc: "世界正在生成",
			NextTickAt: new(nextTickAt),
		}); err != nil {
			return err
		}
		event, err = u.eventUsecase.AppendInTx(ctx, &AppendEventReq{
			WorldID:       world.ID,
			Type:          enum.EventTypeWorldGenerationRequested,
			ActorPlayerID: new(req.CreatorPlayerID),
			Summary:       "请求生成世界",
			Content:       description,
			Payload: map[string]any{
				"npc_count":      float64(npcCount),
				"location_count": float64(locationCount),
			},
		})
		return err
	})
	if err != nil {
		return nil, err
	}
	u.eventUsecase.Publish(event)
	return &CreateWorldResp{
		World: world,
		Event: event,
	}, nil
}

func (u *WorldUsecase) Get(ctx context.Context, worldID int64) (*model.World, error) {
	return u.worldRepo.Get(ctx, &repo.WorldQuery{
		ID: new(worldID),
	})
}

type PageWorldsReq struct {
	Page            base.PageRequest
	CreatorPlayerID *int64
	Status          *enum.WorldStatus
}

type PageWorldsResp struct {
	Rows []*model.World
	Page base.PageResp
}

func (u *WorldUsecase) Page(ctx context.Context, req *PageWorldsReq) (*PageWorldsResp, error) {
	resp, err := u.worldRepo.Page(ctx, &repo.WorldPageReq{
		Page: req.Page,
		Query: repo.WorldQuery{
			CreatorPlayerID: req.CreatorPlayerID,
			Status:          req.Status,
		},
	})
	if err != nil {
		return nil, err
	}
	return &PageWorldsResp{
		Rows: resp.Rows,
		Page: resp.Page,
	}, nil
}
