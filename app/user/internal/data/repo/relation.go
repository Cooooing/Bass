package repo

import (
	"common/proto/gen/common"
	"context"
	"user/internal/biz/model"
	"user/internal/biz/repo"
	"user/internal/data/gen"
	"user/internal/data/gen/relation"
	"user/internal/enum"

	"common/pkg/server"
	utilent "common/pkg/util/ent"
)

var _ repo.RelationRepo = (*RelationRepo)(nil)

type RelationRepo struct {
	db *gen.Client
}

func NewRelationRepo(
	db *gen.Client,
) repo.RelationRepo {
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

func (r *RelationRepo) Create(ctx context.Context, relation *model.Relation) (*model.Relation, error) {
	relation, err := r.create(ctx, relation)
	if err != nil {
		return nil, err
	}
	return relation, nil
}

func (r *RelationRepo) Delete(ctx context.Context, req *repo.RelationDeleteReq) (int, error) {
	deleted, err := r.delete(ctx, req)
	if err != nil {
		return 0, err
	}
	return deleted, nil
}

func (r *RelationRepo) Exists(ctx context.Context, req *repo.RelationGetReq) (bool, error) {
	exists, err := r.exists(ctx, req)
	if err != nil {
		return false, err
	}
	return exists, nil
}

func (r *RelationRepo) Get(ctx context.Context, req *repo.RelationGetReq) (*model.Relation, error) {
	relation, err := r.get(ctx, req)
	if err != nil {
		return nil, err
	}
	return relation, nil
}

func (r *RelationRepo) List(ctx context.Context, req *repo.RelationGetReq) ([]*model.Relation, error) {
	rows, err := r.list(ctx, req)
	if err != nil {
		return nil, err
	}
	return rows, nil
}

func (r *RelationRepo) Map(ctx context.Context, req *repo.RelationGetReq) (map[int64]*model.Relation, error) {
	rows, err := r.mapRows(ctx, req)
	if err != nil {
		return nil, err
	}
	return rows, nil
}

func (r *RelationRepo) Count(ctx context.Context, req *repo.RelationGetReq) (int, error) {
	count, err := r.count(ctx, req)
	if err != nil {
		return 0, err
	}
	return count, nil
}

func (r *RelationRepo) Page(ctx context.Context, req *repo.RelationPageReq) (*repo.RelationPageResp, error) {
	rows, page, err := r.page(ctx, &common.PageReq{
		Page: req.Page.Page,
		Size: req.Page.Size,
	}, &req.Query)
	if err != nil {
		return nil, err
	}
	resp := repo.PageResp{}
	if page != nil {
		resp = repo.PageResp{
			Total: page.GetTotal(),
			Page:  page.GetPage(),
			Size:  page.GetSize(),
		}
	}
	return &repo.RelationPageResp{
		Rows: rows,
		Page: resp,
	}, nil
}

func (r *RelationRepo) create(ctx context.Context, u *model.Relation) (*model.Relation, error) {
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

func (r *RelationRepo) delete(ctx context.Context, req *repo.RelationDeleteReq) (int, error) {
	tx := r.getClient(ctx)
	return tx.Relation.Delete().
		Where(relation.ActorIDEQ(req.ActorID)).
		Where(relation.TargetIDEQ(req.TargetID)).
		Where(relation.TypeEQ(relation.Type(req.Type))).
		Exec(ctx)
}

func (r *RelationRepo) exists(ctx context.Context, req *repo.RelationGetReq) (bool, error) {
	tx := r.getClient(ctx)
	query := tx.Relation.Query()
	query = r.getQuery(query, req)
	return query.Exist(ctx)
}

func (r *RelationRepo) get(ctx context.Context, req *repo.RelationGetReq) (*model.Relation, error) {
	tx := r.getClient(ctx)
	query := tx.Relation.Query()
	query = r.getQuery(query, req)
	rel, err := query.First(ctx)
	if gen.IsNotFound(err) {
		return nil, nil
	}
	if err != nil {
		return nil, err
	}
	return &model.Relation{
		ID:        rel.ID,
		Type:      enum.RelationType(rel.Type),
		ActorID:   rel.ActorID,
		TargetID:  rel.TargetID,
		CreatedAt: rel.CreatedAt,
		UpdatedAt: rel.UpdatedAt,
	}, nil
}

func (r *RelationRepo) list(ctx context.Context, req *repo.RelationGetReq) ([]*model.Relation, error) {
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

func (r *RelationRepo) mapRows(ctx context.Context, req *repo.RelationGetReq) (map[int64]*model.Relation, error) {
	list, err := r.list(ctx, req)
	if err != nil {
		return nil, err
	}
	result := make(map[int64]*model.Relation, len(list))
	for _, item := range list {
		result[item.ID] = item
	}
	return result, nil
}

func (r *RelationRepo) count(ctx context.Context, req *repo.RelationGetReq) (int, error) {
	tx := r.getClient(ctx)
	query := tx.Relation.Query()
	query = r.getQuery(query, req)
	return query.Count(ctx)
}

func (r *RelationRepo) page(ctx context.Context, page *common.PageReq, req *repo.RelationGetReq) ([]*model.Relation, *common.PageResp, error) {
	tx := r.getClient(ctx)
	page = server.PageValid(page)
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
	return result, &common.PageResp{
		Total: uint32(total),
		Page:  page.Page,
		Size:  page.Size,
	}, nil
}

func (r *RelationRepo) getQuery(query *gen.RelationQuery, req *repo.RelationGetReq) *gen.RelationQuery {
	if req == nil {
		return query
	}
	if req.ID != nil {
		query = query.Where(relation.ID(*req.ID))
	}
	if len(req.IDs) > 0 {
		query = query.Where(relation.IDIn(req.IDs...))
	}
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
		query = query.Where(relation.TypeEQ(relation.Type(*req.Type)))
	}
	if req.WithActor {
		query = query.WithActor()
	}
	if req.WithTarget {
		query = query.WithTarget()
	}
	return query
}
