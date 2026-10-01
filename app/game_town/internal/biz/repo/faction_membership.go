package repo

import (
	commonmodel "common/pkg/model"
	"context"
	"game_town/internal/biz/model"
	"game_town/internal/enum"
)

type FactionMembershipRepo interface {
	Save(context.Context, *model.FactionMembership) (*model.FactionMembership, error)
	Get(context.Context, *FactionMembershipQuery) (*model.FactionMembership, error)
	List(context.Context, *FactionMembershipQuery) ([]*model.FactionMembership, error)
	Map(context.Context, *FactionMembershipQuery) (map[int64]*model.FactionMembership, error)
	Count(context.Context, *FactionMembershipQuery) (int, error)
	Page(context.Context, *FactionMembershipPageReq) (*FactionMembershipPageResp, error)
}

type FactionMembershipQuery struct {
	ID, WorldID, FactionID, MemberID *int64
	MemberType                       *enum.EntityType
	ActiveOnly                       bool
}
type FactionMembershipPageReq struct {
	Page  commonmodel.PageReq
	Query FactionMembershipQuery
}
type FactionMembershipPageResp struct {
	Rows []*model.FactionMembership
	Page commonmodel.PageResp
}
