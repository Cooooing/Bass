package repo

import (
	"context"
	"user/internal/biz/model"
	"user/internal/enum"
)

type RelationRepo interface {
	Create(ctx context.Context, req *RelationCreateReq) (*RelationCreateResponse, error)

	Delete(ctx context.Context, req *RelationDeleteReq) (*RelationDeleteResponse, error)

	Exists(ctx context.Context, req *RelationGetReq) (*RelationExistsResponse, error)
	Get(ctx context.Context, req *RelationGetReq) (*RelationGetResponse, error)
	List(ctx context.Context, req *RelationGetReq) (*RelationListResponse, error)
	Map(ctx context.Context, req *RelationGetReq) (*RelationMapResponse, error)
	Count(ctx context.Context, req *RelationGetReq) (*RelationCountResponse, error)
	Page(ctx context.Context, req *RelationPageReq) (*RelationPageResponse, error)
}

type RelationCreateReq struct {
	Relation *model.Relation
}

type RelationCreateResponse struct {
	Relation *model.Relation
}

type RelationDeleteReq struct {
	ActorID  int64
	TargetID int64
	Type     enum.RelationType
}

type RelationDeleteResponse struct {
	Deleted int
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

type RelationExistsResponse struct {
	Exists bool
}

type RelationGetResponse struct {
	Relation *model.Relation
}

type RelationListResponse struct {
	Rows []*model.Relation
}

type RelationMapResponse struct {
	Rows map[int64]*model.Relation
}

type RelationCountResponse struct {
	Count int
}

type RelationPageReq struct {
	Page  PageReq
	Query RelationGetReq
}

type RelationPageResponse struct {
	Rows []*model.Relation
	Page PageResponse
}
