package repo

import (
	"common/api/gen/common"
	"common/pkg/constant"
	"context"
	"user/internal/biz/model"
	"user/internal/biz/repo"
	"user/internal/conf"
	"user/internal/data/gen"
	"user/internal/data/gen/userrelation"
	"user/internal/enum"

	commonClient "common/pkg/client"
	utilent "common/pkg/util/ent"

	"github.com/go-kratos/kratos/v2/log"
)

var _ repo.UserRelationRepo = (*UserRelationRepo)(nil)

type UserRelationRepo struct {
	conf   *conf.Bootstrap
	log    *log.Helper
	db     *gen.Client
	consul *commonClient.ConsulClient
	redis  *commonClient.RedisClient
	nats   *commonClient.NatsClient
}

func NewUserRelationRepo(
	conf *conf.Bootstrap,
	logger log.Logger,
	db *gen.Client,
	consul *commonClient.ConsulClient,
	redis *commonClient.RedisClient,
	nats *commonClient.NatsClient,
) repo.UserRelationRepo {
	return &UserRelationRepo{
		conf:   conf,
		log:    log.NewHelper(logger),
		db:     db,
		consul: consul,
		redis:  redis,
		nats:   nats,
	}
}

func (r *UserRelationRepo) getClient(ctx context.Context) *gen.Client {
	if c, ok := utilent.ClientFromCtx[*gen.Client](ctx); ok {
		return c
	}
	return r.db
}

func toRelationDomain(rel *gen.UserRelation) *model.UserRelation {
	return &model.UserRelation{
		ID:        rel.ID,
		Type:      enum.UserRelationType(rel.Type),
		ActorID:   rel.ActorID,
		TargetID:  rel.TargetID,
		CreatedAt: rel.CreatedAt,
		UpdatedAt: rel.UpdatedAt,
	}
}

func (r *UserRelationRepo) Save(ctx context.Context, u *model.UserRelation) (*model.UserRelation, error) {
	tx := r.getClient(ctx)
	created, err := tx.UserRelation.Create().
		SetActorID(u.ActorID).
		SetTargetID(u.TargetID).
		SetType(userrelation.Type(u.Type)).
		Save(ctx)
	if err != nil {
		return nil, err
	}
	return toRelationDomain(created), nil
}

func (r *UserRelationRepo) Delete(ctx context.Context, u *model.UserRelation) (int, error) {
	if u == nil {
		return 0, nil
	}
	tx := r.getClient(ctx)
	if u.ID != 0 {
		return 1, tx.UserRelation.DeleteOneID(u.ID).Exec(ctx)
	}
	return tx.UserRelation.Delete().
		Where(userrelation.ActorIDEQ(u.ActorID)).
		Where(userrelation.TargetIDEQ(u.TargetID)).
		Where(userrelation.TypeEQ(userrelation.Type(u.Type))).
		Exec(ctx)
}

func (r *UserRelationRepo) Exist(ctx context.Context, req *repo.UserRelationGetReq) (bool, error) {
	tx := r.getClient(ctx)
	query := tx.UserRelation.Query()
	query = r.getQuery(query, req)
	return query.Exist(ctx)
}

func (r *UserRelationRepo) GetList(ctx context.Context, req *repo.UserRelationGetReq) ([]*model.UserRelation, error) {
	tx := r.getClient(ctx)
	query := tx.UserRelation.Query()
	query = r.getQuery(query, req)
	list, err := query.All(ctx)
	if err != nil {
		return nil, err
	}
	result := make([]*model.UserRelation, 0, len(list))
	for _, rel := range list {
		result = append(result, toRelationDomain(rel))
	}
	return result, nil
}

func (r *UserRelationRepo) GetPage(ctx context.Context, page *common.PageRequest, req *repo.UserRelationGetReq) ([]*model.UserRelation, *common.PageReply, error) {
	tx := r.getClient(ctx)
	page = constant.PageValid(page)
	query := tx.UserRelation.Query()
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

	result := make([]*model.UserRelation, 0, len(list))
	for _, rel := range list {
		result = append(result, toRelationDomain(rel))
	}
	return result, &common.PageReply{
		Total: uint32(total),
		Page:  page.Page,
		Size:  page.Size,
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
		dbVal, _ := enum.UserRelationTypeMap.ToEnum(*req.Type)
		query = query.Where(userrelation.TypeEQ(userrelation.Type(dbVal)))
	}
	if req.WithActor {
		query = query.WithActor()
	}
	if req.WithTarget {
		query = query.WithTarget()
	}
	return query
}
