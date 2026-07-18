package repo

import (
	"context"
	"user/internal/biz/model"
	"user/internal/enum"
)

type RelationRepo interface {
	Create(ctx context.Context, relation *model.Relation) (*model.Relation, error)

	Delete(ctx context.Context, req *RelationDeleteReq) (int, error)

	Exists(ctx context.Context, req *RelationGetReq) (bool, error)
	Get(ctx context.Context, req *RelationGetReq) (*model.Relation, error)
	List(ctx context.Context, req *RelationGetReq) ([]*model.Relation, error)
	Map(ctx context.Context, req *RelationGetReq) (map[int64]*model.Relation, error)
	Count(ctx context.Context, req *RelationGetReq) (int, error)
	Page(ctx context.Context, req *RelationPageReq) (*RelationPageResp, error)
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
	Type            *enum.RelationType

	WithActor  bool
	WithTarget bool
}

type RelationPageReq struct {
	Page  PageReq
	Query RelationGetReq
}

type RelationPageResp struct {
	Rows []*model.Relation
	Page PageResp
}
