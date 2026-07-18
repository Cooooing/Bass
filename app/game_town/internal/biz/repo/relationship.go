package repo

import (
	"context"
	"game_town/internal/biz/model"
)

type RelationshipRepo interface {
	GetRelationship(ctx context.Context, req *GetRelationshipReq) (*model.Relationship, error)
	UpsertRelationship(ctx context.Context, row *model.Relationship) (*model.Relationship, error)
}

type GetRelationshipReq struct {
	WorldID  int64
	PlayerID int64
	NpcID    int64
}
