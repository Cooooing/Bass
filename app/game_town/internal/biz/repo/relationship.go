package repo

import (
	"context"
	"game_town/internal/biz/model"
)

type RelationshipRepo interface {
	GetRelationship(ctx context.Context, req *GetRelationshipReq) (*GetRelationshipResponse, error)
	UpsertRelationship(ctx context.Context, req *UpsertRelationshipReq) (*UpsertRelationshipResponse, error)
}

type GetRelationshipReq struct {
	WorldID  int64
	PlayerID int64
	NpcID    int64
}

type GetRelationshipResponse struct {
	Row *model.Relationship
}

type UpsertRelationshipReq struct {
	Row *model.Relationship
}

type UpsertRelationshipResponse struct {
	Row *model.Relationship
}
