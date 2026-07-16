package repo

import (
	"common/pkg/apperror"
	"common/pkg/server"
	"common/proto/gen/common"
	cerrors "common/proto/gen/common/errors"
	"context"
	"game_town/internal/biz/model"
	bizrepo "game_town/internal/biz/repo"
	"game_town/internal/data/gen"
	"game_town/internal/data/gen/world"
	"time"
)

type WorldRepo struct{ *baseRepo }

func NewWorldRepo(db *gen.Client) bizrepo.WorldRepo {
	return &WorldRepo{baseRepo: &baseRepo{db: db}}
}

func (r *WorldRepo) CreateWorld(ctx context.Context, req *bizrepo.CreateWorldReq) (*bizrepo.CreateWorldResponse, error) {
	tx, err := r.db.Tx(ctx)
	if err != nil {
		return nil, err
	}
	defer func() { _ = tx.Rollback() }()
	client := tx.Client()
	now := time.Now()
	code := "w" + time.Now().Format("20060102150405")
	worldRow, err := client.World.Create().SetCode(code).SetName(req.Generated.WorldName).SetDescription(req.Description).SetScale(req.Scale).SetStatus("generating").SetCreatorPlayerID(req.CreatorPlayerID).SetSeed(req.Seed).SetGenerationParams(map[string]any{"npc_count": float64(req.NpcCount), "location_count": float64(req.LocationCount), "style_tags": req.StyleTags, "agent_config_id": optionalNumber(req.AgentConfigID)}).SetGenerationSummary(req.Generated.WorldSummary).SetNillableAgentConfigID(req.AgentConfigID).Save(ctx)
	if err != nil {
		return nil, err
	}
	locations := make(map[string]*gen.Location, len(req.Generated.Locations))
	var defaultLocation *gen.Location
	for i, item := range req.Generated.Locations {
		row, err := client.Location.Create().SetWorldID(worldRow.ID).SetCode(item.Code).SetName(item.Name).SetDescription(item.Description).SetTags(item.Tags).SetSort(int32(i)).SetEnabled(true).Save(ctx)
		if err != nil {
			return nil, err
		}
		locations[item.Code] = row
		if defaultLocation == nil {
			defaultLocation = row
		}
	}
	if defaultLocation == nil {
		return nil, apperror.New(cerrors.BusinessErrorCode_BUSINESS_ERROR_CODE_GAME_TOWN_WORLD_INVALID)
	}
	npcs := make([]*model.Npc, 0, len(req.Generated.Npcs))
	for _, item := range req.Generated.Npcs {
		loc := locations[item.LocationCode]
		if loc == nil {
			loc = defaultLocation
		}
		row, err := client.Npc.Create().SetWorldID(worldRow.ID).SetCode(item.Code).SetName(item.Name).SetRole(item.Role).SetPersonality(item.Personality).SetGoal(item.Goal).SetBackground(item.Background).SetCurrentLocationID(loc.ID).SetState("idle").SetSystemPrompt(item.SystemPrompt).SetGeneratedProfile(item.Profile).SetEnabled(true).Save(ctx)
		if err != nil {
			return nil, err
		}
		npcs = append(npcs, r.npc(row))
	}
	for _, item := range req.Generated.Metrics {
		if _, err := client.WorldMetricDefinition.Create().SetWorldID(worldRow.ID).SetKey(item.Key).SetName(item.Name).SetDescription(item.Description).SetMinValue(item.MinValue).SetMaxValue(item.MaxValue).SetInitialValue(item.InitialValue).Save(ctx); err != nil {
			return nil, err
		}
	}
	stateRow, err := client.WorldStateSnapshot.Create().SetWorldID(worldRow.ID).SetTickCount(0).SetCurrentArc(req.Generated.CurrentArc).SetMetrics(req.Generated.InitialMetrics).SetSummary(req.Generated.WorldSummary).SetCreatedAt(now).Save(ctx)
	if err != nil {
		return nil, err
	}
	events := make([]*model.Event, 0, len(req.Generated.OpeningEvents)+1)
	summary := "世界已创建"
	if len(req.Generated.OpeningEvents) > 0 {
		summary = req.Generated.OpeningEvents[0]
	}
	eventRow, err := client.Event.Create().SetWorldID(worldRow.ID).SetType("world_created").SetActorPlayerID(req.CreatorPlayerID).SetSummary(summary).SetContent(req.Generated.WorldSummary).SetEffects(map[string]any{}).SetMetadata(map[string]any{"seed": float64(req.Seed)}).SetOccurredAt(now).SetCreatedAt(now).Save(ctx)
	if err != nil {
		return nil, err
	}
	events = append(events, r.event(eventRow))
	memberRow, err := client.WorldMember.Create().SetWorldID(worldRow.ID).SetPlayerID(req.CreatorPlayerID).SetCurrentLocationID(defaultLocation.ID).SetRole("owner").SetJoinedAt(now).SetLastSeenAt(now).Save(ctx)
	if err != nil {
		return nil, err
	}
	_ = memberRow
	worldRow, err = client.World.UpdateOneID(worldRow.ID).SetStatus("active").SetDefaultLocationID(defaultLocation.ID).Save(ctx)
	if err != nil {
		return nil, err
	}
	if err := tx.Commit(); err != nil {
		return nil, err
	}
	return &bizrepo.CreateWorldResponse{World: r.world(worldRow), DefaultLocation: r.location(defaultLocation), Npcs: npcs, State: r.state(stateRow), Events: events}, nil
}

func (r *WorldRepo) Get(ctx context.Context, req *bizrepo.WorldGetReq) (*bizrepo.WorldGetResponse, error) {
	row, err := r.db.World.Query().Where(world.ID(req.ID), world.DeletedAtIsNil()).Only(ctx)
	if gen.IsNotFound(err) {
		return nil, apperror.New(cerrors.BusinessErrorCode_BUSINESS_ERROR_CODE_GAME_TOWN_WORLD_NOT_FOUND)
	}
	if err != nil {
		return nil, err
	}
	return &bizrepo.WorldGetResponse{Row: r.world(row)}, nil
}

func (r *WorldRepo) Page(ctx context.Context, req *bizrepo.WorldPageReq) (*bizrepo.WorldPageResponse, error) {
	pageReq := server.PageValid(req.Page)
	queryReq := req.Query
	query := r.db.World.Query().Where(world.DeletedAtIsNil())
	if queryReq.CreatorPlayerID != nil {
		query = query.Where(world.CreatorPlayerID(*queryReq.CreatorPlayerID))
	}
	if queryReq.Status != nil {
		query = query.Where(world.Status(*queryReq.Status))
	}
	total, err := query.Clone().Count(ctx)
	if err != nil {
		return nil, err
	}
	rows, err := query.Order(world.ByID()).Limit(int(pageReq.Size)).Offset(int((pageReq.Page - 1) * pageReq.Size)).All(ctx)
	if err != nil {
		return nil, err
	}
	result := make([]*model.World, 0, len(rows))
	for _, row := range rows {
		result = append(result, r.world(row))
	}
	return &bizrepo.WorldPageResponse{Rows: result, Page: &common.PageResponse{Total: uint32(total), Page: pageReq.Page, Size: pageReq.Size}}, nil
}
