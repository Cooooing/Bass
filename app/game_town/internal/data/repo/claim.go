package repo

import (
	"context"

	"common/pkg/apperror"
	utilent "common/pkg/util/ent"
	cerrors "common/proto/gen/common/errors"
	"game_town/internal/biz/model"
	bizrepo "game_town/internal/biz/repo"
	"game_town/internal/data/gen"
	"game_town/internal/data/gen/claim"
	"game_town/internal/enum"
	"github.com/samber/lo"
)

var _ bizrepo.ClaimRepo = (*ClaimRepo)(nil)

type ClaimRepo struct {
	db *gen.Client
}

func NewClaimRepo(
	db *gen.Client,
) bizrepo.ClaimRepo {
	return &ClaimRepo{
		db: db,
	}
}

func (r *ClaimRepo) getClient(
	ctx context.Context,
) *gen.Client {
	if tx, ok := utilent.ClientFromCtx[*gen.Client](ctx); ok {
		return tx
	}
	return r.db
}

func (r *ClaimRepo) Save(
	ctx context.Context,
	row *model.Claim,
) (*model.Claim, error) {
	saved, err := r.getClient(ctx).Claim.Create().
		SetWorldID(row.WorldID).
		SetNillableOriginEventID(row.OriginEventID).
		SetSubjectType(claim.SubjectType(row.SubjectType)).
		SetSubjectID(row.SubjectID).
		SetPredicate(row.Predicate).
		SetObject(row.Object).
		SetTruth(claim.Truth(row.Truth)).
		Save(ctx)
	if err != nil {
		return nil, err
	}
	return claimModel(saved), nil
}

func claimQuery(
	q *gen.ClaimQuery,
	req *bizrepo.ClaimQuery,
) *gen.ClaimQuery {
	if req == nil {
		return q
	}
	if req.ID != nil {
		q = q.Where(claim.ID(*req.ID))
	}
	if req.WorldID != nil {
		q = q.Where(claim.WorldID(*req.WorldID))
	}
	if req.OriginEventID != nil {
		q = q.Where(claim.OriginEventID(*req.OriginEventID))
	}
	if req.SubjectType != nil {
		q = q.Where(claim.SubjectTypeEQ(claim.SubjectType(*req.SubjectType)))
	}
	if req.SubjectID != nil {
		q = q.Where(claim.SubjectID(*req.SubjectID))
	}
	if req.Predicate != nil {
		q = q.Where(claim.Predicate(*req.Predicate))
	}
	return q
}

func (r *ClaimRepo) Get(
	ctx context.Context,
	req *bizrepo.ClaimQuery,
) (*model.Claim, error) {
	row, err := claimQuery(r.getClient(ctx).Claim.Query(), req).Only(ctx)
	if gen.IsNotFound(err) {
		return nil, apperror.New(cerrors.BusinessErrorCode_BUSINESS_ERROR_CODE_COMMON_NOT_FOUND)
	}
	if err != nil {
		return nil, err
	}
	return claimModel(row), nil
}

func (r *ClaimRepo) List(
	ctx context.Context,
	req *bizrepo.ClaimQuery,
) ([]*model.Claim, error) {
	rows, err := claimQuery(r.getClient(ctx).Claim.Query(), req).
		Order(claim.ByID()).
		All(ctx)
	if err != nil {
		return nil, err
	}
	return lo.Map(rows, func(row *gen.Claim, _ int) *model.Claim {
		return claimModel(row)
	}), nil
}

func (r *ClaimRepo) Map(
	ctx context.Context,
	req *bizrepo.ClaimQuery,
) (map[int64]*model.Claim, error) {
	rows, err := r.List(ctx, req)
	if err != nil {
		return nil, err
	}
	out := make(map[int64]*model.Claim, len(rows))
	for _, row := range rows {
		out[row.ID] = row
	}
	return out, nil
}

func (r *ClaimRepo) Count(
	ctx context.Context,
	req *bizrepo.ClaimQuery,
) (int, error) {
	return claimQuery(r.getClient(ctx).Claim.Query(), req).Count(ctx)
}

func (r *ClaimRepo) Page(
	ctx context.Context,
	req *bizrepo.ClaimPageReq,
) (*bizrepo.ClaimPageResp, error) {
	p := page(req.Page)
	q := claimQuery(r.getClient(ctx).Claim.Query(), &req.Query)
	total, err := q.Clone().Count(ctx)
	if err != nil {
		return nil, err
	}
	rows, err := q.Order(claim.ByID()).
		Offset(pageOffset(p)).
		Limit(pageLimit(p)).
		All(ctx)
	if err != nil {
		return nil, err
	}
	return &bizrepo.ClaimPageResp{
		Rows: lo.Map(rows, func(row *gen.Claim, _ int) *model.Claim {
			return claimModel(row)
		}),
		Page: basePage(total, p),
	}, nil
}

func claimModel(
	row *gen.Claim,
) *model.Claim {
	return &model.Claim{
		ID:            row.ID,
		WorldID:       row.WorldID,
		OriginEventID: row.OriginEventID,
		SubjectType:   enum.EntityType(row.SubjectType),
		SubjectID:     row.SubjectID,
		Predicate:     row.Predicate,
		Object:        row.Object,
		Truth:         enum.ClaimTruth(row.Truth),
		CreatedAt:     row.CreatedAt,
		UpdatedAt:     row.UpdatedAt,
	}
}
