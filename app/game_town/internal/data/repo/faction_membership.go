package repo

import (
	"context"

	"common/pkg/apperror"
	utilent "common/pkg/util/ent"
	cerrors "common/proto/gen/common/errors"
	"game_town/internal/biz/model"
	bizrepo "game_town/internal/biz/repo"
	"game_town/internal/data/gen"
	"game_town/internal/data/gen/factionmembership"
	"game_town/internal/enum"
	"github.com/samber/lo"
)

var _ bizrepo.FactionMembershipRepo = (*FactionMembershipRepo)(nil)

type FactionMembershipRepo struct {
	db *gen.Client
}

func NewFactionMembershipRepo(
	db *gen.Client,
) bizrepo.FactionMembershipRepo {
	return &FactionMembershipRepo{
		db: db,
	}
}

func (r *FactionMembershipRepo) getClient(ctx context.Context) *gen.Client {
	if tx, ok := utilent.ClientFromCtx[*gen.Client](ctx); ok {
		return tx
	}
	return r.db
}

func (r *FactionMembershipRepo) Save(ctx context.Context, row *model.FactionMembership) (*model.FactionMembership, error) {
	saved, err := r.getClient(ctx).FactionMembership.Create().
		SetWorldID(row.WorldID).
		SetFactionID(row.FactionID).
		SetMemberType(factionmembership.MemberType(row.MemberType)).
		SetMemberID(row.MemberID).
		SetRole(row.Role).
		SetReputation(row.Reputation).
		SetTags(row.Tags).
		SetJoinedAt(row.JoinedAt).
		SetNillableLeftAt(row.LeftAt).
		Save(ctx)
	if err != nil {
		return nil, err
	}
	return factionMembershipModel(saved), nil
}

func factionMembershipQuery(q *gen.FactionMembershipQuery, req *bizrepo.FactionMembershipQuery) *gen.FactionMembershipQuery {
	if req == nil {
		return q
	}
	if req.ID != nil {
		q = q.Where(factionmembership.ID(*req.ID))
	}
	if req.WorldID != nil {
		q = q.Where(factionmembership.WorldID(*req.WorldID))
	}
	if req.FactionID != nil {
		q = q.Where(factionmembership.FactionID(*req.FactionID))
	}
	if req.MemberType != nil {
		q = q.Where(factionmembership.MemberTypeEQ(factionmembership.MemberType(*req.MemberType)))
	}
	if req.MemberID != nil {
		q = q.Where(factionmembership.MemberID(*req.MemberID))
	}
	if req.ActiveOnly {
		q = q.Where(factionmembership.LeftAtIsNil())
	}
	return q
}

func (r *FactionMembershipRepo) Get(ctx context.Context, req *bizrepo.FactionMembershipQuery) (*model.FactionMembership, error) {
	row, err := factionMembershipQuery(r.getClient(ctx).FactionMembership.Query(), req).Only(ctx)
	if gen.IsNotFound(err) {
		return nil, apperror.New(cerrors.BusinessErrorCode_BUSINESS_ERROR_CODE_COMMON_NOT_FOUND)
	}
	if err != nil {
		return nil, err
	}
	return factionMembershipModel(row), nil
}

func (r *FactionMembershipRepo) List(ctx context.Context, req *bizrepo.FactionMembershipQuery) ([]*model.FactionMembership, error) {
	rows, err := factionMembershipQuery(r.getClient(ctx).FactionMembership.Query(), req).
		Order(factionmembership.ByID()).
		All(ctx)
	if err != nil {
		return nil, err
	}
	return lo.Map(rows, func(row *gen.FactionMembership, _ int) *model.FactionMembership {
		return factionMembershipModel(row)
	}), nil
}

func (r *FactionMembershipRepo) Map(ctx context.Context, req *bizrepo.FactionMembershipQuery) (map[int64]*model.FactionMembership, error) {
	rows, err := r.List(ctx, req)
	if err != nil {
		return nil, err
	}
	out := make(map[int64]*model.FactionMembership, len(rows))
	for _, row := range rows {
		out[row.ID] = row
	}
	return out, nil
}

func (r *FactionMembershipRepo) Count(ctx context.Context, req *bizrepo.FactionMembershipQuery) (int, error) {
	return factionMembershipQuery(r.getClient(ctx).FactionMembership.Query(), req).Count(ctx)
}

func (r *FactionMembershipRepo) Page(ctx context.Context, req *bizrepo.FactionMembershipPageReq) (*bizrepo.FactionMembershipPageResp, error) {
	p := page(req.Page)
	q := factionMembershipQuery(r.getClient(ctx).FactionMembership.Query(), &req.Query)
	total, err := q.Clone().Count(ctx)
	if err != nil {
		return nil, err
	}
	rows, err := q.Order(factionmembership.ByID()).
		Offset(pageOffset(p)).
		Limit(pageLimit(p)).
		All(ctx)
	if err != nil {
		return nil, err
	}
	return &bizrepo.FactionMembershipPageResp{
		Rows: lo.Map(rows, func(row *gen.FactionMembership, _ int) *model.FactionMembership {
			return factionMembershipModel(row)
		}),
		Page: basePage(total, p),
	}, nil
}

func factionMembershipModel(row *gen.FactionMembership) *model.FactionMembership {
	return &model.FactionMembership{
		ID:         row.ID,
		WorldID:    row.WorldID,
		FactionID:  row.FactionID,
		MemberType: enum.EntityType(row.MemberType),
		MemberID:   row.MemberID,
		Role:       row.Role,
		Reputation: row.Reputation,
		Tags:       row.Tags,
		JoinedAt:   row.JoinedAt,
		LeftAt:     row.LeftAt,
	}
}
