package repo

import (
	cv1 "common/gen/common/v1"
	v1 "common/gen/user/v1"
	"context"
	"user/internal/biz/model"
	"user/internal/data/ent/gen"
)

type UserRelationRepo interface {
	Save(ctx context.Context, tx *gen.Client, u *model.UserRelation) (*model.UserRelation, error)

	Delete(ctx context.Context, tx *gen.Client, u *model.UserRelation) (int, error)

	Exist(ctx context.Context, tx *gen.Client, req *UserRelationGetReq) (bool, error)
	GetList(ctx context.Context, tx *gen.Client, req *UserRelationGetReq) ([]*model.UserRelation, error)
	GetPage(ctx context.Context, tx *gen.Client, page *cv1.PageRequest, req *UserRelationGetReq) ([]*model.UserRelation, *cv1.PageReply, error)
}

type UserRelationGetReq struct {
	ActorId         *int64
	TargetId        *int64
	ActorOrTargetId *int64
	Type            *v1.UserRelationType

	WithActor  bool
	WithTarget bool
}
