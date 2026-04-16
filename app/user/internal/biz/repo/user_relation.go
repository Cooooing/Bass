package repo

import (
	"common/api/gen/common"
	v1 "common/api/gen/user/v1"
	"context"
	"user/internal/biz/model"
	"user/internal/data/ent/gen"
)

type UserRelationRepo interface {
	Save(ctx context.Context, tx *gen.Client, u *model.UserRelation) (*model.UserRelation, error)

	Delete(ctx context.Context, tx *gen.Client, u *model.UserRelation) (int, error)

	Exist(ctx context.Context, tx *gen.Client, req *UserRelationGetReq) (bool, error)
	GetList(ctx context.Context, tx *gen.Client, req *UserRelationGetReq) ([]*model.UserRelation, error)
	GetPage(ctx context.Context, tx *gen.Client, page *common.PageRequest, req *UserRelationGetReq) ([]*model.UserRelation, *common.PageReply, error)
}

type UserRelationGetReq struct {
	ActorId         *int64
	TargetId        *int64
	ActorOrTargetId *int64
	Type            *v1.UserRelationType

	WithActor  bool
	WithTarget bool
}
