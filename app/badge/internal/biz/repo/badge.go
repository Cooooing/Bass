package repo

import (
	commonmodel "common/pkg/model"
	"context"

	"badge/internal/biz/model"
)

// BadgeRepo persists badges owned by the badge domain.
type BadgeRepo interface {
	Get(ctx context.Context, req *BadgeGetReq) (*model.Badge, error)
	List(ctx context.Context, req *BadgeGetReq) ([]*model.Badge, error)
	Page(ctx context.Context, req *BadgeGetReq) (*BadgePageResp, error)
	Create(ctx context.Context, badge *model.Badge) (*model.Badge, error)
	Update(ctx context.Context, badge *model.Badge) (*model.Badge, error)
	Delete(ctx context.Context, badgeID int64) error
}

// BadgeGetReq describes reusable badge query conditions.
type BadgeGetReq struct {
	Page           *commonmodel.PageReq
	BadgeID        *int64
	Code           *string
	DesignerUserID *int64
	Status         *string
}

// BadgePageResp combines a badge page with paging metadata.
type BadgePageResp struct {
	Rows []*model.Badge
	Page *commonmodel.PageResp
}
