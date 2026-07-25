package repo

import (
	"context"

	"common/pkg/apperror"
	utilent "common/pkg/util/ent"
	cerrors "common/proto/gen/common/errors"
	"game_town/internal/biz/model"
	bizrepo "game_town/internal/biz/repo"
	"game_town/internal/data/gen"
	"game_town/internal/data/gen/world"
	"game_town/internal/enum"

	"github.com/samber/lo"
)

var _ bizrepo.WorldRepo = (*WorldRepo)(nil)

type WorldRepo struct {
	db *gen.Client
}

func NewWorldRepo(
	db *gen.Client,
) bizrepo.WorldRepo {
	return &WorldRepo{
		db: db,
	}
}

func (r *WorldRepo) getClient(ctx context.Context) *gen.Client {
	if tx, ok := utilent.ClientFromCtx[*gen.Client](ctx); ok {
		return tx
	}
	return r.db
}

func (r *WorldRepo) Save(ctx context.Context, row *model.World) (*model.World, error) {
	saved, err := r.getClient(ctx).World.Create().
		SetCode(row.Code).
		SetName(row.Name).
		SetDescription(row.Description).
		SetStatus(world.Status(row.Status)).
		SetCreatorPlayerID(row.CreatorPlayerID).
		SetNillableDefaultLocationID(row.DefaultLocationID).
		SetAgentConfigID(row.AgentConfigID).
		SetSeed(row.Seed).
		SetGenerationSummary(row.GenerationSummary).
		Save(ctx)
	if err != nil {
		return nil, err
	}
	return &model.World{
		ID:                saved.ID,
		Code:              saved.Code,
		Name:              saved.Name,
		Description:       saved.Description,
		Status:            enum.WorldStatus(saved.Status),
		CreatorPlayerID:   saved.CreatorPlayerID,
		DefaultLocationID: saved.DefaultLocationID,
		AgentConfigID:     saved.AgentConfigID,
		Seed:              saved.Seed,
		GenerationSummary: saved.GenerationSummary,
		CreatedAt:         saved.CreatedAt,
		UpdatedAt:         saved.UpdatedAt,
	}, nil
}

func worldQuery(q *gen.WorldQuery, req *bizrepo.WorldQuery) *gen.WorldQuery {
	q = q.Where(world.DeletedAtIsNil())
	if req == nil {
		return q
	}
	if req.ID != nil {
		q = q.Where(world.ID(*req.ID))
	}
	if len(req.IDs) > 0 {
		q = q.Where(world.IDIn(req.IDs...))
	}
	if req.Code != nil {
		q = q.Where(world.Code(*req.Code))
	}
	if req.CreatorPlayerID != nil {
		q = q.Where(world.CreatorPlayerID(*req.CreatorPlayerID))
	}
	if req.Status != nil {
		q = q.Where(world.StatusEQ(world.Status(*req.Status)))
	}
	return q
}

func (r *WorldRepo) Get(ctx context.Context, req *bizrepo.WorldQuery) (*model.World, error) {
	row, err := worldQuery(r.getClient(ctx).World.Query(), req).Only(ctx)
	if gen.IsNotFound(err) {
		return nil, apperror.New(cerrors.BusinessErrorCode_BUSINESS_ERROR_CODE_GAME_TOWN_WORLD_NOT_FOUND)
	}
	if err != nil {
		return nil, err
	}
	return &model.World{
		ID:                row.ID,
		Code:              row.Code,
		Name:              row.Name,
		Description:       row.Description,
		Status:            enum.WorldStatus(row.Status),
		CreatorPlayerID:   row.CreatorPlayerID,
		DefaultLocationID: row.DefaultLocationID,
		AgentConfigID:     row.AgentConfigID,
		Seed:              row.Seed,
		GenerationSummary: row.GenerationSummary,
		CreatedAt:         row.CreatedAt,
		UpdatedAt:         row.UpdatedAt,
	}, nil
}

func (r *WorldRepo) List(ctx context.Context, req *bizrepo.WorldQuery) ([]*model.World, error) {
	rows, err := worldQuery(r.getClient(ctx).World.Query(), req).Order(world.ByID()).All(ctx)
	if err != nil {
		return nil, err
	}
	out := lo.Map(rows, func(row *gen.World, _ int) *model.World {
		return &model.World{
			ID:                row.ID,
			Code:              row.Code,
			Name:              row.Name,
			Description:       row.Description,
			Status:            enum.WorldStatus(row.Status),
			CreatorPlayerID:   row.CreatorPlayerID,
			DefaultLocationID: row.DefaultLocationID,
			AgentConfigID:     row.AgentConfigID,
			Seed:              row.Seed,
			GenerationSummary: row.GenerationSummary,
			CreatedAt:         row.CreatedAt,
			UpdatedAt:         row.UpdatedAt,
		}
	})
	return out, nil
}

func (r *WorldRepo) Map(ctx context.Context, req *bizrepo.WorldQuery) (map[int64]*model.World, error) {
	rows, err := r.List(ctx, req)
	if err != nil {
		return nil, err
	}
	out := make(map[int64]*model.World, len(rows))
	for _, row := range rows {
		out[row.ID] = row
	}
	return out, nil
}

func (r *WorldRepo) Count(ctx context.Context, req *bizrepo.WorldQuery) (int, error) {
	return worldQuery(r.getClient(ctx).World.Query(), req).Count(ctx)
}

func (r *WorldRepo) Page(ctx context.Context, req *bizrepo.WorldPageReq) (*bizrepo.WorldPageResp, error) {
	p := page(req.Page)
	q := worldQuery(r.getClient(ctx).World.Query(), &req.Query)
	total, err := q.Clone().Count(ctx)
	if err != nil {
		return nil, err
	}
	rows, err := q.Order(world.ByID()).Offset(pageOffset(p)).Limit(pageLimit(p)).All(ctx)
	if err != nil {
		return nil, err
	}
	out := lo.Map(rows, func(row *gen.World, _ int) *model.World {
		return &model.World{
			ID:                row.ID,
			Code:              row.Code,
			Name:              row.Name,
			Description:       row.Description,
			Status:            enum.WorldStatus(row.Status),
			CreatorPlayerID:   row.CreatorPlayerID,
			DefaultLocationID: row.DefaultLocationID,
			AgentConfigID:     row.AgentConfigID,
			Seed:              row.Seed,
			GenerationSummary: row.GenerationSummary,
			CreatedAt:         row.CreatedAt,
			UpdatedAt:         row.UpdatedAt,
		}
	})
	return &bizrepo.WorldPageResp{
		Rows: out,
		Page: basePage(total, p),
	}, nil
}

func (r *WorldRepo) Update(ctx context.Context, row *model.World) (*model.World, error) {
	saved, err := r.getClient(ctx).World.UpdateOneID(row.ID).
		SetName(row.Name).
		SetDescription(row.Description).
		SetStatus(world.Status(row.Status)).
		SetNillableDefaultLocationID(row.DefaultLocationID).
		SetGenerationSummary(row.GenerationSummary).
		Save(ctx)
	if err != nil {
		return nil, err
	}
	return &model.World{
		ID:                saved.ID,
		Code:              saved.Code,
		Name:              saved.Name,
		Description:       saved.Description,
		Status:            enum.WorldStatus(saved.Status),
		CreatorPlayerID:   saved.CreatorPlayerID,
		DefaultLocationID: saved.DefaultLocationID,
		AgentConfigID:     saved.AgentConfigID,
		Seed:              saved.Seed,
		GenerationSummary: saved.GenerationSummary,
		CreatedAt:         saved.CreatedAt,
		UpdatedAt:         saved.UpdatedAt,
	}, nil
}
