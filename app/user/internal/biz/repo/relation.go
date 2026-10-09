package repo

import (
	commonmodel "common/pkg/model"
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
	ID                *int64
	IDs               []int64
	ActorID           *int64
	TargetID          *int64
	ActorIDs          []int64
	TargetIDs         []int64
	ActorOrTargetID   *int64
	Type              *enum.RelationType
	ExcludedActorIDs  []int64
	ExcludedTargetIDs []int64

	WithActor  bool
	WithTarget bool
}

type RelationPageReq struct {
	Page  commonmodel.PageReq
	Query RelationGetReq
}

type RelationPageResp struct {
	Rows []*model.Relation
	Page commonmodel.PageResp
}
