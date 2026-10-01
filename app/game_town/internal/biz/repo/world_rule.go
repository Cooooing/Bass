package repo

import (
	commonmodel "common/pkg/model"
	"context"
	"game_town/internal/biz/model"
)

type WorldRuleRepo interface {
	Save(context.Context, *model.WorldRule) (*model.WorldRule, error)
	Get(context.Context, *WorldRuleQuery) (*model.WorldRule, error)
	List(context.Context, *WorldRuleQuery) ([]*model.WorldRule, error)
	Map(context.Context, *WorldRuleQuery) (map[int64]*model.WorldRule, error)
	Count(context.Context, *WorldRuleQuery) (int, error)
	Page(context.Context, *WorldRulePageReq) (*WorldRulePageResp, error)
}

type WorldRuleQuery struct{ ID, WorldID *int64 }
type WorldRulePageReq struct {
	Page  commonmodel.PageReq
	Query WorldRuleQuery
}
type WorldRulePageResp struct {
	Rows []*model.WorldRule
	Page commonmodel.PageResp
}
