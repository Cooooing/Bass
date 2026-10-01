package repo

import (
	"content/internal/biz/model"
	"context"
	"time"
)

type MoonbreezeRepo interface {
	Create(ctx context.Context, moonbreeze *model.Moonbreeze) (*model.Moonbreeze, error)
	Page(ctx context.Context, req *MoonbreezePageReq) (*MoonbreezePageResp, error)
}

type MoonbreezePageReq struct {
	AuthorIDs []int64
	BeforeAt  *time.Time
	BeforeID  *int64
	Size      int
}

type MoonbreezePageResp struct {
	Rows []*model.Moonbreeze
}
