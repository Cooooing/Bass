package repo

import (
	"context"
	"game_town/internal/biz/model"
)

type WorldMemberRepo interface {
	JoinWorld(ctx context.Context, req *JoinWorldReq) (*JoinWorldResp, error)
	GetMember(ctx context.Context, req *GetMemberReq) (*model.WorldMember, error)
	MoveMember(ctx context.Context, req *MoveMemberReq) (*model.WorldMember, error)
}

type JoinWorldReq struct {
	PlayerID  int64
	WorldCode string
}

type JoinWorldResp struct {
	World    *model.World
	Member   *model.WorldMember
	Location *model.Location
}

type GetMemberReq struct {
	WorldID  int64
	PlayerID int64
}

type MoveMemberReq struct {
	WorldID    int64
	PlayerID   int64
	LocationID int64
}
