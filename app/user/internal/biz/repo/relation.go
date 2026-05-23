package repo

import (
	"common/api/gen/common"
	v1 "common/api/gen/user/v1"
	"context"
	"user/internal/biz/model"
	"user/internal/enum"
)

type RelationRepo interface {
	Create(ctx context.Context, u *model.Relation) (*model.Relation, error)

	Delete(ctx context.Context, req *RelationDeleteReq) (int, error)

	Exists(ctx context.Context, req *RelationGetReq) (bool, error)
	List(ctx context.Context, req *RelationGetReq) ([]*model.Relation, error)
	Page(ctx context.Context, page *common.PageRequest, req *RelationGetReq) ([]*model.Relation, *common.PageReply, error)
}

type RelationDeleteReq struct {
	ActorID  int64
	TargetID int64
	Type     enum.RelationType
}

type RelationGetReq struct {
	ActorId         *int64
	TargetId        *int64
	ActorOrTargetId *int64
	Type            *v1.RelationType

	WithActor  bool
	WithTarget bool
}
