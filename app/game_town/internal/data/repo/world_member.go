package repo

import (
	"context"
	"time"

	"common/pkg/apperror"
	utilent "common/pkg/util/ent"
	cerrors "common/proto/gen/common/errors"
	"game_town/internal/biz/model"
	bizrepo "game_town/internal/biz/repo"
	"game_town/internal/data/gen"
	"game_town/internal/data/gen/worldmember"
	"game_town/internal/enum"

	"github.com/samber/lo"
)

var _ bizrepo.WorldMemberRepo = (*WorldMemberRepo)(nil)

type WorldMemberRepo struct {
	db *gen.Client
}

func NewWorldMemberRepo(
	db *gen.Client,
) bizrepo.WorldMemberRepo {
	return &WorldMemberRepo{
		db: db,
	}
}

func (r *WorldMemberRepo) getClient(ctx context.Context) *gen.Client {
	if tx, ok := utilent.ClientFromCtx[*gen.Client](ctx); ok {
		return tx
	}
	return r.db
}

func (r *WorldMemberRepo) Save(ctx context.Context, row *model.WorldMember) (*model.WorldMember, error) {
	saved, err := r.getClient(ctx).WorldMember.Create().
		SetWorldID(row.WorldID).
		SetPlayerID(row.PlayerID).
		SetCurrentLocationID(row.CurrentLocationID).
		SetRole(worldmember.Role(row.Role)).
		SetCharacterPreference(row.CharacterPreference).
		SetCharacterName(row.CharacterName).
		SetCharacterBackground(row.CharacterBackground).
		SetCharacterGoal(row.CharacterGoal).
		SetCharacterTraits(row.CharacterTraits).
		SetCharacterReady(row.CharacterReady).
		SetJoinedAt(row.JoinedAt).
		SetLastSeenAt(row.LastSeenAt).
		Save(ctx)
	if err != nil {
		return nil, err
	}
	return memberModel(saved), nil
}

func memberQuery(q *gen.WorldMemberQuery, req *bizrepo.WorldMemberQuery) *gen.WorldMemberQuery {
	if req == nil {
		return q
	}
	if req.ID != nil {
		q = q.Where(worldmember.ID(*req.ID))
	}
	if len(req.IDs) > 0 {
		q = q.Where(worldmember.IDIn(req.IDs...))
	}
	if req.WorldID != nil {
		q = q.Where(worldmember.WorldID(*req.WorldID))
	}
	if req.PlayerID != nil {
		q = q.Where(worldmember.PlayerID(*req.PlayerID))
	}
	if req.LocationID != nil {
		q = q.Where(worldmember.CurrentLocationID(*req.LocationID))
	}
	return q
}

func (r *WorldMemberRepo) Get(ctx context.Context, req *bizrepo.WorldMemberQuery) (*model.WorldMember, error) {
	row, err := memberQuery(r.getClient(ctx).WorldMember.Query(), req).Only(ctx)
	if gen.IsNotFound(err) {
		return nil, apperror.New(cerrors.BusinessErrorCode_BUSINESS_ERROR_CODE_COMMON_NOT_FOUND)
	}
	if err != nil {
		return nil, err
	}
	return memberModel(row), nil
}

func (r *WorldMemberRepo) List(ctx context.Context, req *bizrepo.WorldMemberQuery) ([]*model.WorldMember, error) {
	rows, err := memberQuery(r.getClient(ctx).WorldMember.Query(), req).Order(worldmember.ByID()).All(ctx)
	if err != nil {
		return nil, err
	}
	return lo.Map(rows, func(row *gen.WorldMember, _ int) *model.WorldMember {
		return memberModel(row)
	}), nil
}

func (r *WorldMemberRepo) Map(ctx context.Context, req *bizrepo.WorldMemberQuery) (map[int64]*model.WorldMember, error) {
	rows, err := r.List(ctx, req)
	if err != nil {
		return nil, err
	}
	out := make(map[int64]*model.WorldMember, len(rows))
	for _, row := range rows {
		out[row.ID] = row
	}
	return out, nil
}

func (r *WorldMemberRepo) Count(ctx context.Context, req *bizrepo.WorldMemberQuery) (int, error) {
	return memberQuery(r.getClient(ctx).WorldMember.Query(), req).Count(ctx)
}

func (r *WorldMemberRepo) Page(ctx context.Context, req *bizrepo.WorldMemberPageReq) (*bizrepo.WorldMemberPageResp, error) {
	p := page(req.Page)
	q := memberQuery(r.getClient(ctx).WorldMember.Query(), &req.Query)
	total, err := q.Clone().Count(ctx)
	if err != nil {
		return nil, err
	}
	rows, err := q.Order(worldmember.ByID()).Offset(pageOffset(p)).Limit(pageLimit(p)).All(ctx)
	if err != nil {
		return nil, err
	}
	return &bizrepo.WorldMemberPageResp{
		Rows: lo.Map(rows, func(row *gen.WorldMember, _ int) *model.WorldMember {
			return memberModel(row)
		}),
		Page: basePage(total, p),
	}, nil
}

func (r *WorldMemberRepo) Move(ctx context.Context, id int64, locationID int64) (*model.WorldMember, error) {
	row, err := r.getClient(ctx).WorldMember.UpdateOneID(id).
		SetCurrentLocationID(locationID).
		SetLastSeenAt(time.Now()).
		Save(ctx)
	if err != nil {
		return nil, err
	}
	return memberModel(row), nil
}

func (r *WorldMemberRepo) UpdateCharacter(ctx context.Context, req *bizrepo.WorldMemberCharacterReq) (*model.WorldMember, error) {
	row, err := r.getClient(ctx).WorldMember.UpdateOneID(req.MemberID).
		SetCharacterName(req.Name).
		SetCharacterBackground(req.Background).
		SetCharacterGoal(req.Goal).
		SetCharacterTraits(req.Traits).
		SetCharacterReady(req.Ready).
		SetLastSeenAt(time.Now()).
		Save(ctx)
	if err != nil {
		return nil, err
	}
	return memberModel(row), nil
}

func memberModel(row *gen.WorldMember) *model.WorldMember {
	return &model.WorldMember{
		ID:                  row.ID,
		WorldID:             row.WorldID,
		PlayerID:            row.PlayerID,
		CurrentLocationID:   row.CurrentLocationID,
		Role:                enum.WorldMemberRole(row.Role),
		CharacterPreference: row.CharacterPreference,
		CharacterName:       row.CharacterName,
		CharacterBackground: row.CharacterBackground,
		CharacterGoal:       row.CharacterGoal,
		CharacterTraits:     row.CharacterTraits,
		CharacterReady:      row.CharacterReady,
		JoinedAt:            row.JoinedAt,
		LastSeenAt:          row.LastSeenAt,
		CreatedAt:           row.CreatedAt,
		UpdatedAt:           row.UpdatedAt,
	}
}
