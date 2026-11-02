package repo

import (
	"badge/internal/biz/model"
	commonmodel "common/pkg/model"
	"context"
)

// BadgeRepo persists badges owned by the badge domain.
type BadgeRepo interface {
	Get(context.Context, *BadgeGetReq) (*model.Badge, error)
	List(context.Context, *BadgeGetReq) ([]*model.Badge, error)
	Page(context.Context, *BadgeGetReq) (*BadgePageResp, error)
	Create(context.Context, *model.Badge) (*model.Badge, error)
	Update(context.Context, *model.Badge) (*model.Badge, error)
	Delete(context.Context, int64) error
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
