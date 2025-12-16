package data

import (
	cv1 "common/api/common/v1"
	"common/pkg/constant"
	"context"
	"user/internal/biz/model"
	"user/internal/biz/repo"
	"user/internal/data/ent/gen"
	"user/internal/data/ent/gen/userrelation"
)

type UserRelationRepo struct {
	*BaseRepo
}

func NewUserRelationRepo(repo *BaseRepo) repo.UserRelationRepo {
	return &UserRelationRepo{
		BaseRepo: repo,
	}
}

func (r *UserRelationRepo) Save(ctx context.Context, tx *gen.Client, u *model.UserRelation) (*model.UserRelation, error) {
	userRelationCreate, err := tx.UserRelation.Create().
		SetActorID(u.ActorID).
		SetTargetID(u.TargetID).
		SetType(u.Type).
		Save(ctx)
	return &model.UserRelation{UserRelation: userRelationCreate}, err
}

func (r *UserRelationRepo) Delete(ctx context.Context, tx *gen.Client, u *model.UserRelation) (int, error) {
	if u == nil {
		return 0, nil
	}
	if u.ID != 0 {
		return 1, tx.UserRelation.DeleteOneID(u.ID).Exec(ctx)
	}
	return tx.UserRelation.Delete().
		Where(userrelation.ActorIDEQ(u.ActorID)).
		Where(userrelation.TargetIDEQ(u.TargetID)).
		Where(userrelation.TypeEQ(u.Type)).
		Exec(ctx)
}

func (r *UserRelationRepo) Exist(ctx context.Context, tx *gen.Client, req *repo.UserRelationGetReq) (bool, error) {
	query := tx.UserRelation.Query()
	query = r.getQuery(query, req)
	return query.Exist(ctx)
}

func (r *UserRelationRepo) GetList(ctx context.Context, tx *gen.Client, req *repo.UserRelationGetReq) ([]*model.UserRelation, error) {
	var (
		records []*model.UserRelation
		err     error
	)
	query := tx.UserRelation.Query()
	query = r.getQuery(query, req)
	list, err := query.All(ctx)
	if err != nil {
		return nil, err
	}
	for i := range list {
		records = append(records, &model.UserRelation{UserRelation: list[i]})
	}
	return records, nil
}

func (r *UserRelationRepo) GetPage(ctx context.Context, tx *gen.Client, page *cv1.PageRequest, req *repo.UserRelationGetReq) ([]*model.UserRelation, *cv1.PageReply, error) {
	var (
		notificationRecords []*model.UserRelation
		err                 error
	)
	page = constant.PageValid(page)
	query := tx.UserRelation.Query()
	query = r.getQuery(query, req)
	countQuery := query.Clone()
	count, err := countQuery.Count(ctx)
	if err != nil {
		return nil, nil, err
	}
	list, err := query.Limit(int(page.Size)).Offset(int((page.Page - 1) * page.Size)).All(ctx)
	if err != nil {
		return nil, nil, err
	}

	for _, item := range list {
		notificationRecords = append(notificationRecords, &model.UserRelation{UserRelation: item})
	}
	return notificationRecords, &cv1.PageReply{
		Total: uint32(count),
		Size:  page.Size,
		Page:  page.Page,
	}, nil
}

func (r *UserRelationRepo) getQuery(query *gen.UserRelationQuery, req *repo.UserRelationGetReq) *gen.UserRelationQuery {
	if req.ActorId != nil {
		query = query.Where(userrelation.ActorIDEQ(*req.ActorId))
	}
	if req.TargetId != nil {
		query = query.Where(userrelation.TargetIDEQ(*req.TargetId))
	}
	if req.ActorOrTargetId != nil {
		query = query.Where(userrelation.Or(
			userrelation.ActorIDEQ(*req.ActorOrTargetId),
			userrelation.TargetIDEQ(*req.ActorOrTargetId),
		))
	}
	if req.Type != nil {
		query = query.Where(userrelation.TypeEQ(int32(*req.Type)))
	}
	if req.WithActor {
		query = query.WithActor()
	}
	if req.WithTarget {
		query = query.WithTarget()
	}
	return query
}
