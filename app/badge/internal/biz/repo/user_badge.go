package repo

import (
	commonmodel "common/pkg/model"
	"context"

	"badge/internal/biz/model"
)

// UserBadgeRepo persists a user's display state and awarded badge facts.
type UserBadgeRepo interface {
	Get(ctx context.Context, req *UserBadgeGetReq) (*model.UserBadge, error)
	List(ctx context.Context, req *UserBadgeGetReq) ([]*model.UserBadge, error)
	Page(ctx context.Context, req *UserBadgeGetReq) (*UserBadgePageResp, error)
	Grant(ctx context.Context, userBadges []*model.UserBadge) ([]*model.UserBadge, error)
	Revoke(ctx context.Context, userBadges []*model.UserBadge) (int, error)
	UpdatePresentation(ctx context.Context, userID, badgeID int64, displayed bool, position int32) (*model.UserBadge, error)
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
