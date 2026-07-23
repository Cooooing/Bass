package repo

import (
	"context"

	"game_town/internal/biz/base"
	"game_town/internal/biz/model"
	"game_town/internal/enum"
)

type ClaimRepo interface {
	Save(context.Context, *model.Claim) (*model.Claim, error)
	Get(context.Context, *ClaimQuery) (*model.Claim, error)
	List(context.Context, *ClaimQuery) ([]*model.Claim, error)
	Map(context.Context, *ClaimQuery) (map[int64]*model.Claim, error)
	Count(context.Context, *ClaimQuery) (int, error)
	Page(context.Context, *ClaimPageReq) (*ClaimPageResp, error)
}

type ClaimQuery struct {
	ID, WorldID, OriginEventID, SubjectID *int64
	SubjectType                           *enum.EntityType
	Predicate                             *string
}
type ClaimPageReq struct {
	Page  base.PageRequest
	Query ClaimQuery
}
type ClaimPageResp struct {
	Rows []*model.Claim
	Page base.PageResp
}
