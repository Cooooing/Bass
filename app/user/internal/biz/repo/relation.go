package repo

import (
	"common/proto/gen/common"
	v1 "common/proto/gen/user/v1"
	"context"
	"user/internal/biz/model"
	"user/internal/enum"
)

type RelationRepo interface {
	Create(ctx context.Context, u *model.Relation) (*model.Relation, error)

	Delete(ctx context.Context, req *RelationDeleteReq) (int, error)

	Exists(ctx context.Context, req *RelationGetReq) (bool, error)
	Get(ctx context.Context, req *RelationGetReq) (*model.Relation, error)
	List(ctx context.Context, req *RelationGetReq) ([]*model.Relation, error)
	Map(ctx context.Context, req *RelationGetReq) (map[int64]*model.Relation, error)
	Count(ctx context.Context, req *RelationGetReq) (int, error)
	Page(ctx context.Context, page *common.PageRequest, req *RelationGetReq) ([]*model.Relation, *common.PageReply, error)
}

type RelationDeleteReq struct {
	ActorID  int64
	TargetID int64
	Type     enum.RelationType
}

type RelationGetReq struct {
	ID              *int64
	IDs             []int64
	ActorId         *int64
	TargetId        *int64
	ActorOrTargetId *int64
	Type            *v1.RelationType

	WithActor  bool
	WithTarget bool
}
