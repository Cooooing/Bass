package repo

import (
	"context"

	"common/pkg/apperror"
	utilent "common/pkg/util/ent"
	cerrors "common/proto/gen/common/errors"
	"game_town/internal/biz/model"
	bizrepo "game_town/internal/biz/repo"
	"game_town/internal/data/gen"
	"game_town/internal/data/gen/player"
	"game_town/internal/enum"

	"github.com/samber/lo"
)

var _ bizrepo.PlayerRepo = (*PlayerRepo)(nil)

type PlayerRepo struct {
	pageHelper
	db *gen.Client
}

func NewPlayerRepo(
	db *gen.Client,
) bizrepo.PlayerRepo {
	return &PlayerRepo{
		db: db,
	}
}

func (r *PlayerRepo) getClient(ctx context.Context) *gen.Client {
	if tx, ok := utilent.ClientFromCtx[*gen.Client](ctx); ok {
		return tx
	}
	return r.db
}

func (r *PlayerRepo) Save(ctx context.Context, row *model.Player) (*model.Player, error) {
	saved, err := r.getClient(ctx).Player.Create().
		SetName(row.Name).
		SetDisplayName(row.DisplayName).
		SetStatus(player.Status(row.Status)).
		Save(ctx)
	if err != nil {
		return nil, err
	}
	return &model.Player{
		ID:          saved.ID,
		Name:        saved.Name,
		DisplayName: saved.DisplayName,
		Status:      enum.PlayerStatus(saved.Status),
		CreatedAt:   saved.CreatedAt,
		UpdatedAt:   saved.UpdatedAt,
	}, nil
}

func (r *PlayerRepo) playerQuery(q *gen.PlayerQuery, req *bizrepo.PlayerQuery) *gen.PlayerQuery {
	q = q.Where(player.DeletedAtIsNil())
	if req == nil {
		return q
	}
	if req.ID != nil {
		q = q.Where(player.ID(*req.ID))
	}
	if len(req.IDs) > 0 {
		q = q.Where(player.IDIn(req.IDs...))
	}
	if req.Name != nil {
		q = q.Where(player.Name(*req.Name))
	}
	if req.Status != nil {
		q = q.Where(player.StatusEQ(player.Status(*req.Status)))
	}
	return q
}

func (r *PlayerRepo) Get(ctx context.Context, req *bizrepo.PlayerQuery) (*model.Player, error) {
	row, err := r.playerQuery(r.getClient(ctx).Player.Query(), req).Only(ctx)
	if gen.IsNotFound(err) {
		return nil, apperror.New(cerrors.BusinessErrorCode_BUSINESS_ERROR_CODE_GAME_TOWN_PLAYER_NOT_FOUND)
	}
	if err != nil {
		return nil, err
	}
	return &model.Player{
		ID:          row.ID,
		Name:        row.Name,
		DisplayName: row.DisplayName,
		Status:      enum.PlayerStatus(row.Status),
		CreatedAt:   row.CreatedAt,
		UpdatedAt:   row.UpdatedAt,
	}, nil
}

func (r *PlayerRepo) List(ctx context.Context, req *bizrepo.PlayerQuery) ([]*model.Player, error) {
	rows, err := r.playerQuery(r.getClient(ctx).Player.Query(), req).Order(player.ByID()).All(ctx)
	if err != nil {
		return nil, err
	}
	out := lo.Map(rows, func(row *gen.Player, _ int) *model.Player {
		return &model.Player{
			ID:          row.ID,
			Name:        row.Name,
			DisplayName: row.DisplayName,
			Status:      enum.PlayerStatus(row.Status),
			CreatedAt:   row.CreatedAt,
			UpdatedAt:   row.UpdatedAt,
		}
	})
	return out, nil
}

func (r *PlayerRepo) Map(ctx context.Context, req *bizrepo.PlayerQuery) (map[int64]*model.Player, error) {
	rows, err := r.List(ctx, req)
	if err != nil {
		return nil, err
	}
	out := make(map[int64]*model.Player, len(rows))
	for _, row := range rows {
		out[row.ID] = row
	}
	return out, nil
}

func (r *PlayerRepo) Count(ctx context.Context, req *bizrepo.PlayerQuery) (int, error) {
	return r.playerQuery(r.getClient(ctx).Player.Query(), req).Count(ctx)
}

func (r *PlayerRepo) Page(ctx context.Context, req *bizrepo.PlayerPageReq) (*bizrepo.PlayerPageResp, error) {
	p := r.page(req.Page)
	q := r.playerQuery(r.getClient(ctx).Player.Query(), &req.Query)
	total, err := q.Clone().Count(ctx)
	if err != nil {
		return nil, err
	}
	rows, err := q.Order(player.ByID()).Offset(r.pageOffset(p)).Limit(r.pageLimit(p)).All(ctx)
	if err != nil {
		return nil, err
	}
	out := lo.Map(rows, func(row *gen.Player, _ int) *model.Player {
		return &model.Player{
			ID:          row.ID,
			Name:        row.Name,
			DisplayName: row.DisplayName,
			Status:      enum.PlayerStatus(row.Status),
			CreatedAt:   row.CreatedAt,
			UpdatedAt:   row.UpdatedAt,
		}
	})
	return &bizrepo.PlayerPageResp{
		Rows: out,
		Page: r.basePage(total, p),
	}, nil
}
