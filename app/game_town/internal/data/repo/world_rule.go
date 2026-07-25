package repo

import (
	"context"

	"common/pkg/apperror"
	utilent "common/pkg/util/ent"
	cerrors "common/proto/gen/common/errors"
	"game_town/internal/biz/model"
	bizrepo "game_town/internal/biz/repo"
	"game_town/internal/data/gen"
	"game_town/internal/data/gen/worldrule"

	"github.com/samber/lo"
)

var _ bizrepo.WorldRuleRepo = (*WorldRuleRepo)(nil)

type WorldRuleRepo struct {
	pageHelper
	db *gen.Client
}

func NewWorldRuleRepo(
	db *gen.Client,
) bizrepo.WorldRuleRepo {
	return &WorldRuleRepo{
		db: db,
	}
}

func (r *WorldRuleRepo) getClient(ctx context.Context) *gen.Client {
	if tx, ok := utilent.ClientFromCtx[*gen.Client](ctx); ok {
		return tx
	}
	return r.db
}

func (r *WorldRuleRepo) Save(ctx context.Context, row *model.WorldRule) (*model.WorldRule, error) {
	saved, err := r.getClient(ctx).WorldRule.Create().
		SetWorldID(row.WorldID).
		SetVersion(row.Version).
		SetRules(row.Rules).
		Save(ctx)
	if err != nil {
		return nil, err
	}
	return &model.WorldRule{
		ID:        saved.ID,
		WorldID:   saved.WorldID,
		Version:   saved.Version,
		Rules:     saved.Rules,
		CreatedAt: saved.CreatedAt,
		UpdatedAt: saved.UpdatedAt,
	}, nil
}

func (r *WorldRuleRepo) worldRuleQuery(q *gen.WorldRuleQuery, req *bizrepo.WorldRuleQuery) *gen.WorldRuleQuery {
	if req == nil {
		return q
	}
	if req.ID != nil {
		q = q.Where(worldrule.ID(*req.ID))
	}
	if req.WorldID != nil {
		q = q.Where(worldrule.WorldID(*req.WorldID))
	}
	return q
}

func (r *WorldRuleRepo) Get(ctx context.Context, req *bizrepo.WorldRuleQuery) (*model.WorldRule, error) {
	row, err := r.worldRuleQuery(r.getClient(ctx).WorldRule.Query(), req).Only(ctx)
	if gen.IsNotFound(err) {
		return nil, apperror.New(cerrors.BusinessErrorCode_BUSINESS_ERROR_CODE_COMMON_NOT_FOUND)
	}
	if err != nil {
		return nil, err
	}
	return &model.WorldRule{
		ID:        row.ID,
		WorldID:   row.WorldID,
		Version:   row.Version,
		Rules:     row.Rules,
		CreatedAt: row.CreatedAt,
		UpdatedAt: row.UpdatedAt,
	}, nil
}

func (r *WorldRuleRepo) List(ctx context.Context, req *bizrepo.WorldRuleQuery) ([]*model.WorldRule, error) {
	rows, err := r.worldRuleQuery(r.getClient(ctx).WorldRule.Query(), req).Order(worldrule.ByID()).All(ctx)
	if err != nil {
		return nil, err
	}
	return lo.Map(rows, func(row *gen.WorldRule, _ int) *model.WorldRule {
		return &model.WorldRule{
			ID:        row.ID,
			WorldID:   row.WorldID,
			Version:   row.Version,
			Rules:     row.Rules,
			CreatedAt: row.CreatedAt,
			UpdatedAt: row.UpdatedAt,
		}
	}), nil
}

func (r *WorldRuleRepo) Map(ctx context.Context, req *bizrepo.WorldRuleQuery) (map[int64]*model.WorldRule, error) {
	rows, err := r.List(ctx, req)
	if err != nil {
		return nil, err
	}
	out := make(map[int64]*model.WorldRule, len(rows))
	for _, row := range rows {
		out[row.ID] = row
	}
	return out, nil
}

func (r *WorldRuleRepo) Count(ctx context.Context, req *bizrepo.WorldRuleQuery) (int, error) {
	return r.worldRuleQuery(r.getClient(ctx).WorldRule.Query(), req).Count(ctx)
}

func (r *WorldRuleRepo) Page(ctx context.Context, req *bizrepo.WorldRulePageReq) (*bizrepo.WorldRulePageResp, error) {
	p := r.page(req.Page)
	q := r.worldRuleQuery(r.getClient(ctx).WorldRule.Query(), &req.Query)
	total, err := q.Clone().Count(ctx)
	if err != nil {
		return nil, err
	}
	rows, err := q.Order(worldrule.ByID()).Offset(r.pageOffset(p)).Limit(r.pageLimit(p)).All(ctx)
	if err != nil {
		return nil, err
	}
	out := lo.Map(rows, func(row *gen.WorldRule, _ int) *model.WorldRule {
		return &model.WorldRule{
			ID:        row.ID,
			WorldID:   row.WorldID,
			Version:   row.Version,
			Rules:     row.Rules,
			CreatedAt: row.CreatedAt,
			UpdatedAt: row.UpdatedAt,
		}
	})
	return &bizrepo.WorldRulePageResp{
		Rows: out,
		Page: r.basePage(total, p),
	}, nil
}
