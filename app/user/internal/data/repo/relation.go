package repo

import (
	"common/api/gen/common"
	"common/pkg/constant"
	"context"
	"user/internal/biz/model"
	"user/internal/biz/repo"
	"user/internal/data/gen"
	"user/internal/data/gen/relation"
	"user/internal/enum"

	utilent "common/pkg/util/ent"
)

var _ repo.RelationRepo = (*RelationRepo)(nil)

type RelationRepo struct {
	db *gen.Client
}

func NewRelationRepo(db *gen.Client) repo.RelationRepo {
	return &RelationRepo{
		db: db,
	}
}

func (r *RelationRepo) getClient(ctx context.Context) *gen.Client {
	if c, ok := utilent.ClientFromCtx[*gen.Client](ctx); ok {
		return c
	}
	return r.db
}

func (r *RelationRepo) Create(ctx context.Context, u *model.Relation) (*model.Relation, error) {
	tx := r.getClient(ctx)
	created, err := tx.Relation.Create().
		SetActorID(u.ActorID).
		SetTargetID(u.TargetID).
		SetType(relation.Type(u.Type)).
		Save(ctx)
	if err != nil {
		return nil, err
	}
	return &model.Relation{
		ID:        created.ID,
		Type:      enum.RelationType(created.Type),
		ActorID:   created.ActorID,
		TargetID:  created.TargetID,
		CreatedAt: created.CreatedAt,
		UpdatedAt: created.UpdatedAt,
	}, nil
}

func (r *RelationRepo) Delete(ctx context.Context, req *repo.RelationDeleteReq) (int, error) {
	tx := r.getClient(ctx)
	return tx.Relation.Delete().
		Where(relation.ActorIDEQ(req.ActorID)).
		Where(relation.TargetIDEQ(req.TargetID)).
		Where(relation.TypeEQ(relation.Type(req.Type))).
		Exec(ctx)
}

func (r *RelationRepo) Exists(ctx context.Context, req *repo.RelationGetReq) (bool, error) {
	tx := r.getClient(ctx)
	query := tx.Relation.Query()
	query = r.getQuery(query, req)
	return query.Exist(ctx)
}

func (r *RelationRepo) List(ctx context.Context, req *repo.RelationGetReq) ([]*model.Relation, error) {
	tx := r.getClient(ctx)
	query := tx.Relation.Query()
	query = r.getQuery(query, req)
	list, err := query.All(ctx)
	if err != nil {
		return nil, err
	}
	result := make([]*model.Relation, 0, len(list))
	for _, rel := range list {
		result = append(result, &model.Relation{
			ID:        rel.ID,
			Type:      enum.RelationType(rel.Type),
			ActorID:   rel.ActorID,
			TargetID:  rel.TargetID,
			CreatedAt: rel.CreatedAt,
			UpdatedAt: rel.UpdatedAt,
		})
	}
	return result, nil
}

func (r *RelationRepo) Page(ctx context.Context, page *common.PageRequest, req *repo.RelationGetReq) ([]*model.Relation, *common.PageReply, error) {
	tx := r.getClient(ctx)
	page = constant.PageValid(page)
	query := tx.Relation.Query()
	query = r.getQuery(query, req)

	total, err := query.Clone().Count(ctx)
	if err != nil {
		return nil, nil, err
	}

	list, err := query.
		Limit(int(page.Size)).
		Offset(int((page.Page - 1) * page.Size)).
		All(ctx)
	if err != nil {
		return nil, nil, err
	}

	result := make([]*model.Relation, 0, len(list))
	for _, rel := range list {
		result = append(result, &model.Relation{
			ID:        rel.ID,
			Type:      enum.RelationType(rel.Type),
			ActorID:   rel.ActorID,
			TargetID:  rel.TargetID,
			CreatedAt: rel.CreatedAt,
			UpdatedAt: rel.UpdatedAt,
		})
	}
	return result, &common.PageReply{
		Total: uint32(total),
		Page:  page.Page,
		Size:  page.Size,
	}, nil
}

func (r *RelationRepo) getQuery(query *gen.RelationQuery, req *repo.RelationGetReq) *gen.RelationQuery {
	if req.ActorId != nil {
		query = query.Where(relation.ActorIDEQ(*req.ActorId))
	}
	if req.TargetId != nil {
		query = query.Where(relation.TargetIDEQ(*req.TargetId))
	}
	if req.ActorOrTargetId != nil {
		query = query.Where(relation.Or(
			relation.ActorIDEQ(*req.ActorOrTargetId),
			relation.TargetIDEQ(*req.ActorOrTargetId),
		))
	}
	if req.Type != nil {
		dbVal, _ := enum.RelationTypeMap.ToEnum(*req.Type)
		query = query.Where(relation.TypeEQ(relation.Type(dbVal)))
	}
	if req.WithActor {
		query = query.WithActor()
	}
	if req.WithTarget {
		query = query.WithTarget()
	}
	return query
}
