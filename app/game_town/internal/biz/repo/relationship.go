package repo

import (
	"context"

	"game_town/internal/biz/base"
	"game_town/internal/biz/model"
	"game_town/internal/enum"
)

type RelationshipRepo interface {
	Save(context.Context, *model.Relationship) (*model.Relationship, error)
	Upsert(context.Context, *RelationshipUpsertReq) (*model.Relationship, error)
	Get(context.Context, *RelationshipQuery) (*model.Relationship, error)
	List(context.Context, *RelationshipQuery) ([]*model.Relationship, error)
	Map(context.Context, *RelationshipQuery) (map[int64]*model.Relationship, error)
	Count(context.Context, *RelationshipQuery) (int, error)
	Page(context.Context, *RelationshipPageReq) (*RelationshipPageResp, error)
}

type RelationshipUpsertReq struct {
	WorldID    int64
	SourceType enum.EntityType
	SourceID   int64
	TargetType enum.EntityType
	TargetID   int64
	Metrics    map[string]float64
	Tags       []string
}

type RelationshipQuery struct {
	ID         *int64
	WorldID    *int64
	SourceID   *int64
	TargetID   *int64
	SourceType *enum.EntityType
	TargetType *enum.EntityType
}

type RelationshipPageReq struct {
	Page  base.PageRequest
	Query RelationshipQuery
}

type RelationshipPageResp struct {
	Rows []*model.Relationship
	Page base.PageResp
}
