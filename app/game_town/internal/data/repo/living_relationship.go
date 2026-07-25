package repo

import (
	"context"

	"common/pkg/apperror"
	utilent "common/pkg/util/ent"
	cerrors "common/proto/gen/common/errors"
	"game_town/internal/biz/model"
	bizrepo "game_town/internal/biz/repo"
	"game_town/internal/data/gen"
	"game_town/internal/data/gen/livingrelationship"
	"game_town/internal/enum"
	"github.com/samber/lo"
)

var _ bizrepo.RelationshipRepo = (*RelationshipRepo)(nil)

type RelationshipRepo struct {
	pageHelper
	db *gen.Client
}

func NewRelationshipRepo(
	db *gen.Client,
) bizrepo.RelationshipRepo {
	return &RelationshipRepo{
		db: db,
	}
}

func (r *RelationshipRepo) getClient(ctx context.Context) *gen.Client {
	if tx, ok := utilent.ClientFromCtx[*gen.Client](ctx); ok {
		return tx
	}
	return r.db
}

func (r *RelationshipRepo) Save(ctx context.Context, row *model.Relationship) (*model.Relationship, error) {
	saved, err := r.getClient(ctx).LivingRelationship.Create().
		SetWorldID(row.WorldID).
		SetSourceType(livingrelationship.SourceType(row.SourceType)).
		SetSourceID(row.SourceID).
		SetTargetType(livingrelationship.TargetType(row.TargetType)).
		SetTargetID(row.TargetID).
		SetMetrics(row.Metrics).
		SetTags(row.Tags).
		SetVersion(row.Version).
		Save(ctx)
	if err != nil {
		return nil, err
	}
	return r.relationship(saved), nil
}

func (r *RelationshipRepo) relationshipQuery(q *gen.LivingRelationshipQuery, req *bizrepo.RelationshipQuery) *gen.LivingRelationshipQuery {
	if req == nil {
		return q
	}
	if req.ID != nil {
		q = q.Where(livingrelationship.ID(*req.ID))
	}
	if req.WorldID != nil {
		q = q.Where(livingrelationship.WorldID(*req.WorldID))
	}
	if req.SourceType != nil {
		q = q.Where(livingrelationship.SourceTypeEQ(livingrelationship.SourceType(*req.SourceType)))
	}
	if req.SourceID != nil {
		q = q.Where(livingrelationship.SourceID(*req.SourceID))
	}
	if req.TargetType != nil {
		q = q.Where(livingrelationship.TargetTypeEQ(livingrelationship.TargetType(*req.TargetType)))
	}
	if req.TargetID != nil {
		q = q.Where(livingrelationship.TargetID(*req.TargetID))
	}
	return q
}

func (r *RelationshipRepo) Get(ctx context.Context, req *bizrepo.RelationshipQuery) (*model.Relationship, error) {
	row, err := r.relationshipQuery(r.getClient(ctx).LivingRelationship.Query(), req).Only(ctx)
	if gen.IsNotFound(err) {
		return nil, apperror.New(cerrors.BusinessErrorCode_BUSINESS_ERROR_CODE_COMMON_NOT_FOUND)
	}
	if err != nil {
		return nil, err
	}
	return r.relationship(row), nil
}

func (r *RelationshipRepo) List(ctx context.Context, req *bizrepo.RelationshipQuery) ([]*model.Relationship, error) {
	rows, err := r.relationshipQuery(r.getClient(ctx).LivingRelationship.Query(), req).
		Order(livingrelationship.ByID()).
		All(ctx)
	if err != nil {
		return nil, err
	}
	return lo.Map(rows, func(row *gen.LivingRelationship, _ int) *model.Relationship {
		return r.relationship(row)
	}), nil
}

func (r *RelationshipRepo) Map(ctx context.Context, req *bizrepo.RelationshipQuery) (map[int64]*model.Relationship, error) {
	rows, err := r.List(ctx, req)
	if err != nil {
		return nil, err
	}
	out := make(map[int64]*model.Relationship, len(rows))
	for _, row := range rows {
		out[row.ID] = row
	}
	return out, nil
}

func (r *RelationshipRepo) Count(ctx context.Context, req *bizrepo.RelationshipQuery) (int, error) {
	return r.relationshipQuery(r.getClient(ctx).LivingRelationship.Query(), req).Count(ctx)
}

func (r *RelationshipRepo) Page(ctx context.Context, req *bizrepo.RelationshipPageReq) (*bizrepo.RelationshipPageResp, error) {
	p := r.page(req.Page)
	q := r.relationshipQuery(r.getClient(ctx).LivingRelationship.Query(), &req.Query)
	total, err := q.Clone().Count(ctx)
	if err != nil {
		return nil, err
	}
	rows, err := q.Order(livingrelationship.ByID()).
		Offset(r.pageOffset(p)).
		Limit(r.pageLimit(p)).
		All(ctx)
	if err != nil {
		return nil, err
	}
	return &bizrepo.RelationshipPageResp{
		Rows: lo.Map(rows, func(row *gen.LivingRelationship, _ int) *model.Relationship {
			return r.relationship(row)
		}),
		Page: r.basePage(total, p),
	}, nil
}

func (r *RelationshipRepo) relationship(row *gen.LivingRelationship) *model.Relationship {
	return &model.Relationship{
		ID:         row.ID,
		WorldID:    row.WorldID,
		SourceType: enum.EntityType(row.SourceType),
		SourceID:   row.SourceID,
		TargetType: enum.EntityType(row.TargetType),
		TargetID:   row.TargetID,
		Metrics:    row.Metrics,
		Tags:       row.Tags,
		Version:    row.Version,
		CreatedAt:  row.CreatedAt,
		UpdatedAt:  row.UpdatedAt,
	}
}

func (r *RelationshipRepo) Upsert(ctx context.Context, req *bizrepo.RelationshipUpsertReq) (*model.Relationship, error) {
	query := &bizrepo.RelationshipQuery{
		WorldID:    new(req.WorldID),
		SourceType: new(req.SourceType),
		SourceID:   new(req.SourceID),
		TargetType: new(req.TargetType),
		TargetID:   new(req.TargetID),
	}
	row, err := r.relationshipQuery(r.getClient(ctx).LivingRelationship.Query(), query).Only(ctx)
	if gen.IsNotFound(err) {
		return r.Save(ctx, &model.Relationship{
			WorldID:    req.WorldID,
			SourceType: req.SourceType,
			SourceID:   req.SourceID,
			TargetType: req.TargetType,
			TargetID:   req.TargetID,
			Metrics:    req.Metrics,
			Tags:       req.Tags,
		})
	}
	if err != nil {
		return nil, err
	}
	updated, err := r.getClient(ctx).LivingRelationship.UpdateOneID(row.ID).
		SetMetrics(req.Metrics).
		SetTags(req.Tags).
		AddVersion(1).
		Save(ctx)
	if err != nil {
		return nil, err
	}
	return r.relationship(updated), nil
}
