package repo

import (
	"context"

	"common/pkg/apperror"
	utilent "common/pkg/util/ent"
	cerrors "common/proto/gen/common/errors"
	"game_town/internal/biz/model"
	bizrepo "game_town/internal/biz/repo"
	"game_town/internal/data/gen"
	"game_town/internal/data/gen/npcbelief"
	"game_town/internal/enum"
	"github.com/samber/lo"
)

var _ bizrepo.NpcBeliefRepo = (*NpcBeliefRepo)(nil)

type NpcBeliefRepo struct {
	pageHelper
	db *gen.Client
}

func NewNpcBeliefRepo(
	db *gen.Client,
) bizrepo.NpcBeliefRepo {
	return &NpcBeliefRepo{
		db: db,
	}
}

func (r *NpcBeliefRepo) getClient(ctx context.Context) *gen.Client {
	if tx, ok := utilent.ClientFromCtx[*gen.Client](ctx); ok {
		return tx
	}
	return r.db
}

func (r *NpcBeliefRepo) Save(ctx context.Context, row *model.NpcBelief) (*model.NpcBelief, error) {
	saved, err := r.getClient(ctx).NpcBelief.Create().
		SetWorldID(row.WorldID).
		SetNpcID(row.NpcID).
		SetClaimID(row.ClaimID).
		SetNillableSourceNpcID(row.SourceNpcID).
		SetNillableSourcePlayerID(row.SourcePlayerID).
		SetNillableSourceEventID(row.SourceEventID).
		SetStance(npcbelief.Stance(row.Stance)).
		SetConfidence(row.Confidence).
		SetLearnedAt(row.LearnedAt).
		SetUpdatedAt(row.LearnedAt).
		Save(ctx)
	if err != nil {
		return nil, err
	}
	return r.npcBelief(saved), nil
}

func (r *NpcBeliefRepo) npcBeliefQuery(q *gen.NpcBeliefQuery, req *bizrepo.NpcBeliefQuery) *gen.NpcBeliefQuery {
	if req == nil {
		return q
	}
	if req.ID != nil {
		q = q.Where(npcbelief.ID(*req.ID))
	}
	if req.WorldID != nil {
		q = q.Where(npcbelief.WorldID(*req.WorldID))
	}
	if req.NpcID != nil {
		q = q.Where(npcbelief.NpcID(*req.NpcID))
	}
	if req.ClaimID != nil {
		q = q.Where(npcbelief.ClaimID(*req.ClaimID))
	}
	if req.MinConfidence != nil {
		q = q.Where(npcbelief.ConfidenceGTE(*req.MinConfidence))
	}
	return q
}

func (r *NpcBeliefRepo) Get(ctx context.Context, req *bizrepo.NpcBeliefQuery) (*model.NpcBelief, error) {
	row, err := r.npcBeliefQuery(r.getClient(ctx).NpcBelief.Query(), req).Only(ctx)
	if gen.IsNotFound(err) {
		return nil, apperror.New(cerrors.BusinessErrorCode_BUSINESS_ERROR_CODE_COMMON_NOT_FOUND)
	}
	if err != nil {
		return nil, err
	}
	return r.npcBelief(row), nil
}

func (r *NpcBeliefRepo) List(ctx context.Context, req *bizrepo.NpcBeliefQuery) ([]*model.NpcBelief, error) {
	rows, err := r.npcBeliefQuery(r.getClient(ctx).NpcBelief.Query(), req).
		Order(npcbelief.ByConfidence()).
		All(ctx)
	if err != nil {
		return nil, err
	}
	return lo.Map(rows, func(row *gen.NpcBelief, _ int) *model.NpcBelief {
		return r.npcBelief(row)
	}), nil
}

func (r *NpcBeliefRepo) Map(ctx context.Context, req *bizrepo.NpcBeliefQuery) (map[int64]*model.NpcBelief, error) {
	rows, err := r.List(ctx, req)
	if err != nil {
		return nil, err
	}
	out := make(map[int64]*model.NpcBelief, len(rows))
	for _, row := range rows {
		out[row.ID] = row
	}
	return out, nil
}

func (r *NpcBeliefRepo) Count(ctx context.Context, req *bizrepo.NpcBeliefQuery) (int, error) {
	return r.npcBeliefQuery(r.getClient(ctx).NpcBelief.Query(), req).Count(ctx)
}

func (r *NpcBeliefRepo) Page(ctx context.Context, req *bizrepo.NpcBeliefPageReq) (*bizrepo.NpcBeliefPageResp, error) {
	p := r.page(req.Page)
	q := r.npcBeliefQuery(r.getClient(ctx).NpcBelief.Query(), &req.Query)
	total, err := q.Clone().Count(ctx)
	if err != nil {
		return nil, err
	}
	rows, err := q.Order(npcbelief.ByConfidence()).
		Offset(r.pageOffset(p)).
		Limit(r.pageLimit(p)).
		All(ctx)
	if err != nil {
		return nil, err
	}
	return &bizrepo.NpcBeliefPageResp{
		Rows: lo.Map(rows, func(row *gen.NpcBelief, _ int) *model.NpcBelief {
			return r.npcBelief(row)
		}),
		Page: r.basePage(total, p),
	}, nil
}

func (r *NpcBeliefRepo) npcBelief(row *gen.NpcBelief) *model.NpcBelief {
	return &model.NpcBelief{
		ID:             row.ID,
		WorldID:        row.WorldID,
		NpcID:          row.NpcID,
		ClaimID:        row.ClaimID,
		SourceNpcID:    row.SourceNpcID,
		SourcePlayerID: row.SourcePlayerID,
		SourceEventID:  row.SourceEventID,
		Stance:         enum.BeliefStance(row.Stance),
		Confidence:     row.Confidence,
		LearnedAt:      row.LearnedAt,
		UpdatedAt:      row.UpdatedAt,
	}
}
