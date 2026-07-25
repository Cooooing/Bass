package repo

import (
	"context"
	"fmt"

	"common/pkg/apperror"
	utilent "common/pkg/util/ent"
	cerrors "common/proto/gen/common/errors"
	"game_town/internal/biz/model"
	bizrepo "game_town/internal/biz/repo"
	"game_town/internal/data/gen"
	"game_town/internal/data/gen/npc"
	"game_town/internal/enum"

	"github.com/samber/lo"
)

var _ bizrepo.NpcRepo = (*NpcRepo)(nil)

type NpcRepo struct {
	db *gen.Client
}

func NewNpcRepo(
	db *gen.Client,
) bizrepo.NpcRepo {
	return &NpcRepo{
		db: db,
	}
}

func (r *NpcRepo) getClient(
	ctx context.Context,
) *gen.Client {
	if tx, ok := utilent.ClientFromCtx[*gen.Client](ctx); ok {
		return tx
	}
	return r.db
}

func (r *NpcRepo) Save(
	ctx context.Context,
	row *model.Npc,
) (*model.Npc, error) {
	saved, err := r.getClient(ctx).Npc.Create().
		SetWorldID(row.WorldID).
		SetCode(row.Code).
		SetName(row.Name).
		SetRole(row.Role).
		SetSpecies(row.Species).
		SetPersonality(row.Personality).
		SetGoal(row.Goal).
		SetBackground(row.Background).
		SetCurrentLocationID(row.CurrentLocationID).
		SetSystemPrompt(row.SystemPrompt).
		SetContextSummary(row.ContextSummary).
		SetLifeStatus(npc.LifeStatus(row.LifeStatus)).
		SetStateTags(row.StateTags).
		SetAttributes(row.Attributes).
		SetNillableBirthWorldTime(row.BirthWorldTime).
		SetNillableDeathWorldTime(row.DeathWorldTime).
		SetNillableNextDecisionAt(row.NextDecisionAt).
		SetNillableLastPlannedAt(row.LastPlannedAt).
		SetVersion(row.Version).
		Save(ctx)
	if err != nil {
		return nil, err
	}
	return &model.Npc{
		ID:                saved.ID,
		WorldID:           saved.WorldID,
		Code:              saved.Code,
		Name:              saved.Name,
		Role:              saved.Role,
		Species:           saved.Species,
		Personality:       saved.Personality,
		Goal:              saved.Goal,
		Background:        saved.Background,
		CurrentLocationID: saved.CurrentLocationID,
		SystemPrompt:      saved.SystemPrompt,
		ContextSummary:    saved.ContextSummary,
		LifeStatus:        enum.NpcLifeStatus(saved.LifeStatus),
		StateTags:         saved.StateTags,
		Attributes:        saved.Attributes,
		BirthWorldTime:    saved.BirthWorldTime,
		DeathWorldTime:    saved.DeathWorldTime,
		NextDecisionAt:    saved.NextDecisionAt,
		LastPlannedAt:     saved.LastPlannedAt,
		Version:           saved.Version,
		CreatedAt:         saved.CreatedAt,
		UpdatedAt:         saved.UpdatedAt,
	}, nil
}

func npcQuery(
	q *gen.NpcQuery,
	req *bizrepo.NpcQuery,
) *gen.NpcQuery {
	q = q.Where(npc.DeletedAtIsNil())
	if req == nil {
		return q
	}
	if req.ID != nil {
		q = q.Where(npc.ID(*req.ID))
	}
	if len(req.IDs) > 0 {
		q = q.Where(npc.IDIn(req.IDs...))
	}
	if req.WorldID != nil {
		q = q.Where(npc.WorldID(*req.WorldID))
	}
	if req.LocationID != nil {
		q = q.Where(npc.CurrentLocationID(*req.LocationID))
	}
	if req.Code != nil {
		q = q.Where(npc.Code(*req.Code))
	}
	if req.NextDecisionBefore != nil {
		q = q.Where(npc.NextDecisionAtNotNil(), npc.NextDecisionAtLTE(*req.NextDecisionBefore), npc.LifeStatusEQ(npc.LifeStatus(enum.NpcLifeStatusAlive)))
	}
	return q
}

func (r *NpcRepo) Get(
	ctx context.Context,
	req *bizrepo.NpcQuery,
) (*model.Npc, error) {
	row, err := npcQuery(r.getClient(ctx).Npc.Query(), req).Only(ctx)
	if gen.IsNotFound(err) {
		return nil, apperror.New(cerrors.BusinessErrorCode_BUSINESS_ERROR_CODE_GAME_TOWN_NPC_NOT_FOUND)
	}
	if err != nil {
		return nil, err
	}
	return &model.Npc{
		ID:                row.ID,
		WorldID:           row.WorldID,
		Code:              row.Code,
		Name:              row.Name,
		Role:              row.Role,
		Personality:       row.Personality,
		Species:           row.Species,
		Goal:              row.Goal,
		Background:        row.Background,
		CurrentLocationID: row.CurrentLocationID,
		SystemPrompt:      row.SystemPrompt,
		ContextSummary:    row.ContextSummary,
		Version:           row.Version,
		LifeStatus:        enum.NpcLifeStatus(row.LifeStatus),
		StateTags:         row.StateTags,
		Attributes:        row.Attributes,
		BirthWorldTime:    row.BirthWorldTime,
		DeathWorldTime:    row.DeathWorldTime,
		NextDecisionAt:    row.NextDecisionAt,
		LastPlannedAt:     row.LastPlannedAt,
		CreatedAt:         row.CreatedAt,
		UpdatedAt:         row.UpdatedAt,
	}, nil
}

func (r *NpcRepo) List(
	ctx context.Context,
	req *bizrepo.NpcQuery,
) ([]*model.Npc, error) {
	rows, err := npcQuery(r.getClient(ctx).Npc.Query(), req).Order(npc.ByID()).All(ctx)
	if err != nil {
		return nil, err
	}
	out := lo.Map(rows, func(row *gen.Npc, _ int) *model.Npc {
		return &model.Npc{
			ID:                row.ID,
			WorldID:           row.WorldID,
			Code:              row.Code,
			Name:              row.Name,
			Role:              row.Role,
			Personality:       row.Personality,
			Goal:              row.Goal,
			Species:           row.Species,
			Background:        row.Background,
			CurrentLocationID: row.CurrentLocationID,
			SystemPrompt:      row.SystemPrompt,
			ContextSummary:    row.ContextSummary,
			Version:           row.Version,
			CreatedAt:         row.CreatedAt,
			LifeStatus:        enum.NpcLifeStatus(row.LifeStatus),
			StateTags:         row.StateTags,
			Attributes:        row.Attributes,
			BirthWorldTime:    row.BirthWorldTime,
			DeathWorldTime:    row.DeathWorldTime,
			NextDecisionAt:    row.NextDecisionAt,
			LastPlannedAt:     row.LastPlannedAt,
			UpdatedAt:         row.UpdatedAt,
		}
	})
	return out, nil
}

func (r *NpcRepo) Map(
	ctx context.Context,
	req *bizrepo.NpcQuery,
) (map[int64]*model.Npc, error) {
	rows, err := r.List(ctx, req)
	if err != nil {
		return nil, err
	}
	out := make(map[int64]*model.Npc, len(rows))
	for _, row := range rows {
		out[row.ID] = row
	}
	return out, nil
}

func (r *NpcRepo) Count(
	ctx context.Context,
	req *bizrepo.NpcQuery,
) (int, error) {
	return npcQuery(r.getClient(ctx).Npc.Query(), req).Count(ctx)
}

func (r *NpcRepo) Page(
	ctx context.Context,
	req *bizrepo.NpcPageReq,
) (*bizrepo.NpcPageResp, error) {
	p := page(req.Page)
	q := npcQuery(r.getClient(ctx).Npc.Query(), &req.Query)
	total, err := q.Clone().Count(ctx)
	if err != nil {
		return nil, err
	}
	rows, err := q.Order(npc.ByID()).Offset(pageOffset(p)).Limit(pageLimit(p)).All(ctx)
	if err != nil {
		return nil, err
	}
	out := lo.Map(rows, func(row *gen.Npc, _ int) *model.Npc {
		return &model.Npc{
			ID:                row.ID,
			WorldID:           row.WorldID,
			Code:              row.Code,
			Name:              row.Name,
			Role:              row.Role,
			Personality:       row.Personality,
			Goal:              row.Goal,
			Background:        row.Background,
			Species:           row.Species,
			CurrentLocationID: row.CurrentLocationID,
			SystemPrompt:      row.SystemPrompt,
			ContextSummary:    row.ContextSummary,
			Version:           row.Version,
			CreatedAt:         row.CreatedAt,
			UpdatedAt:         row.UpdatedAt,
			LifeStatus:        enum.NpcLifeStatus(row.LifeStatus),
			StateTags:         row.StateTags,
			Attributes:        row.Attributes,
			BirthWorldTime:    row.BirthWorldTime,
			DeathWorldTime:    row.DeathWorldTime,
			NextDecisionAt:    row.NextDecisionAt,
			LastPlannedAt:     row.LastPlannedAt,
		}
	})
	return &bizrepo.NpcPageResp{
		Rows: out,
		Page: basePage(total, p),
	}, nil
}

func (r *NpcRepo) UpdateContext(
	ctx context.Context,
	id int64,
	version int64,
	summary string,
) (*model.Npc, error) {
	count, err := r.getClient(ctx).Npc.Update().Where(npc.ID(id), npc.Version(version), npc.DeletedAtIsNil()).SetContextSummary(summary).AddVersion(1).Save(ctx)
	if err != nil {
		return nil, err
	}
	if count == 0 {
		return nil, fmt.Errorf("npc version conflict")
	}
	return r.Get(ctx, &bizrepo.NpcQuery{
		ID: new(id),
	})
}
