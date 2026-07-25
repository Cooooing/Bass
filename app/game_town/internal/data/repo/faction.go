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
	"game_town/internal/data/gen/faction"
	"game_town/internal/enum"
	"github.com/samber/lo"
)

var _ bizrepo.FactionRepo = (*FactionRepo)(nil)

type FactionRepo struct {
	pageHelper
	db *gen.Client
}

func NewFactionRepo(
	db *gen.Client,
) bizrepo.FactionRepo {
	return &FactionRepo{
		db: db,
	}
}

func (r *FactionRepo) getClient(ctx context.Context) *gen.Client {
	if tx, ok := utilent.ClientFromCtx[*gen.Client](ctx); ok {
		return tx
	}
	return r.db
}

func (r *FactionRepo) Save(ctx context.Context, row *model.Faction) (*model.Faction, error) {
	saved, err := r.getClient(ctx).Faction.Create().
		SetWorldID(row.WorldID).
		SetCode(row.Code).
		SetName(row.Name).
		SetDescription(row.Description).
		SetPublicGoal(row.PublicGoal).
		SetStatus(faction.Status(row.Status)).
		SetAttributes(row.Attributes).
		SetVersion(row.Version).
		Save(ctx)
	if err != nil {
		return nil, err
	}
	return r.faction(saved), nil
}

func (r *FactionRepo) factionQuery(q *gen.FactionQuery, req *bizrepo.FactionQuery) *gen.FactionQuery {
	q = q.Where(faction.DeletedAtIsNil())
	if req == nil {
		return q
	}
	if req.ID != nil {
		q = q.Where(faction.ID(*req.ID))
	}
	if len(req.IDs) > 0 {
		q = q.Where(faction.IDIn(req.IDs...))
	}
	if req.WorldID != nil {
		q = q.Where(faction.WorldID(*req.WorldID))
	}
	if req.Code != nil {
		q = q.Where(faction.Code(*req.Code))
	}
	return q
}

func (r *FactionRepo) Get(ctx context.Context, req *bizrepo.FactionQuery) (*model.Faction, error) {
	row, err := r.factionQuery(r.getClient(ctx).Faction.Query(), req).Only(ctx)
	if gen.IsNotFound(err) {
		return nil, apperror.New(cerrors.BusinessErrorCode_BUSINESS_ERROR_CODE_COMMON_NOT_FOUND)
	}
	if err != nil {
		return nil, err
	}
	return r.faction(row), nil
}

func (r *FactionRepo) List(ctx context.Context, req *bizrepo.FactionQuery) ([]*model.Faction, error) {
	rows, err := r.factionQuery(r.getClient(ctx).Faction.Query(), req).
		Order(faction.ByID()).
		All(ctx)
	if err != nil {
		return nil, err
	}
	return lo.Map(rows, func(row *gen.Faction, _ int) *model.Faction {
		return r.faction(row)
	}), nil
}

func (r *FactionRepo) Map(ctx context.Context, req *bizrepo.FactionQuery) (map[int64]*model.Faction, error) {
	rows, err := r.List(ctx, req)
	if err != nil {
		return nil, err
	}
	out := make(map[int64]*model.Faction, len(rows))
	for _, row := range rows {
		out[row.ID] = row
	}
	return out, nil
}

func (r *FactionRepo) Count(ctx context.Context, req *bizrepo.FactionQuery) (int, error) {
	return r.factionQuery(r.getClient(ctx).Faction.Query(), req).Count(ctx)
}

func (r *FactionRepo) Page(ctx context.Context, req *bizrepo.FactionPageReq) (*bizrepo.FactionPageResp, error) {
	p := r.page(req.Page)
	q := r.factionQuery(r.getClient(ctx).Faction.Query(), &req.Query)
	total, err := q.Clone().Count(ctx)
	if err != nil {
		return nil, err
	}
	rows, err := q.Order(faction.ByID()).
		Offset(r.pageOffset(p)).
		Limit(r.pageLimit(p)).
		All(ctx)
	if err != nil {
		return nil, err
	}
	return &bizrepo.FactionPageResp{
		Rows: lo.Map(rows, func(row *gen.Faction, _ int) *model.Faction {
			return r.faction(row)
		}),
		Page: r.basePage(total, p),
	}, nil
}

func (r *FactionRepo) faction(row *gen.Faction) *model.Faction {
	return &model.Faction{
		ID:          row.ID,
		WorldID:     row.WorldID,
		Code:        row.Code,
		Name:        row.Name,
		Description: row.Description,
		PublicGoal:  row.PublicGoal,
		Status:      enum.FactionStatus(row.Status),
		Attributes:  row.Attributes,
		Version:     row.Version,
		CreatedAt:   row.CreatedAt,
		UpdatedAt:   row.UpdatedAt,
	}
}

func (r *FactionRepo) UpdateState(ctx context.Context, req *bizrepo.FactionStateUpdateReq) (*model.Faction, error) {
	update := r.getClient(ctx).Faction.Update().Where(faction.ID(req.FactionID), faction.Version(req.Version), faction.DeletedAtIsNil())
	if req.Status != nil {
		update.SetStatus(faction.Status(*req.Status))
	}
	if req.Description != "" {
		update.SetDescription(req.Description)
	}
	if req.PublicGoal != "" {
		update.SetPublicGoal(req.PublicGoal)
	}
	if req.Attributes != nil {
		update.SetAttributes(req.Attributes)
	}
	update.AddVersion(1)
	count, err := update.Save(ctx)
	if err != nil {
		return nil, err
	}
	if count == 0 {
		return nil, fmt.Errorf("faction version conflict")
	}
	return r.Get(ctx, &bizrepo.FactionQuery{
		ID: new(req.FactionID),
	})
}
