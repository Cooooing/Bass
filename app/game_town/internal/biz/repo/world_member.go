package repo

import (
	"context"

	"game_town/internal/biz/base"
	"game_town/internal/biz/model"
)

type WorldMemberRepo interface {
	Save(ctx context.Context, member *model.WorldMember) (*model.WorldMember, error)
	Move(ctx context.Context, memberID int64, locationID int64) (*model.WorldMember, error)
	UpdateCharacter(ctx context.Context, req *WorldMemberCharacterReq) (*model.WorldMember, error)
	Get(ctx context.Context, req *WorldMemberQuery) (*model.WorldMember, error)
	List(ctx context.Context, req *WorldMemberQuery) ([]*model.WorldMember, error)
	Map(ctx context.Context, req *WorldMemberQuery) (map[int64]*model.WorldMember, error)
	Count(ctx context.Context, req *WorldMemberQuery) (int, error)
	Page(ctx context.Context, req *WorldMemberPageReq) (*WorldMemberPageResp, error)
}

type WorldMemberCharacterReq struct {
	MemberID   int64
	Name       string
	Background string
	Goal       string
	Traits     []string
	Ready      bool
}

type WorldMemberQuery struct {
	ID         *int64
	IDs        []int64
	WorldID    *int64
	PlayerID   *int64
	LocationID *int64
}

type WorldMemberPageReq struct {
	Page  base.PageRequest
	Query WorldMemberQuery
}

type WorldMemberPageResp struct {
	Rows []*model.WorldMember
	Page base.PageResp
}
