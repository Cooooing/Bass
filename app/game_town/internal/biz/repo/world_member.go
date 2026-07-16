package repo

import (
	"context"
	"game_town/internal/biz/model"
)

type WorldMemberRepo interface {
	JoinWorld(ctx context.Context, req *JoinWorldReq) (*JoinWorldResponse, error)
	GetMember(ctx context.Context, req *GetMemberReq) (*GetMemberResponse, error)
	MoveMember(ctx context.Context, req *MoveMemberReq) (*MoveMemberResponse, error)
}

type JoinWorldReq struct {
	PlayerID  int64
	WorldCode string
}

type JoinWorldResponse struct {
	World    *model.World
	Member   *model.WorldMember
	Location *model.Location
}

type GetMemberReq struct {
	WorldID  int64
	PlayerID int64
}

type GetMemberResponse struct {
	Row *model.WorldMember
}

type MoveMemberReq struct {
	WorldID    int64
	PlayerID   int64
	LocationID int64
}

type MoveMemberResponse struct {
	Row *model.WorldMember
}
