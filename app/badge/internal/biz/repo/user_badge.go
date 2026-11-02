package repo

import (
	"badge/internal/biz/model"
	commonmodel "common/pkg/model"
	"context"
)

// UserBadgeRepo persists a user's display state and awarded badge facts.
type UserBadgeRepo interface {
	Get(context.Context, *UserBadgeGetReq) (*model.UserBadge, error)
	List(context.Context, *UserBadgeGetReq) ([]*model.UserBadge, error)
	Page(context.Context, *UserBadgeGetReq) (*UserBadgePageResp, error)
	Grant(context.Context, []*model.UserBadge) ([]*model.UserBadge, error)
	Revoke(context.Context, []*model.UserBadge) (int, error)
	UpdatePresentation(context.Context, int64, int64, bool, int32) (*model.UserBadge, error)
}

// UserBadgeGetReq describes reusable user badge query conditions.
type UserBadgeGetReq struct {
	Page       *commonmodel.PageReq
	UserID     *int64
	BadgeID    *int64
	Displayed  *bool
	ActiveOnly bool
}

// UserBadgePageResp combines a user badge page with paging metadata.
type UserBadgePageResp struct {
	Rows []*model.UserBadge
	Page *commonmodel.PageResp
}
