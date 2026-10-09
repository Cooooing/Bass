package repo

import (
	"context"
	"time"

	"content/internal/biz/model"
)

type MoonbreezeRepo interface {
	Create(ctx context.Context, moonbreeze *model.Moonbreeze) (*model.Moonbreeze, error)
	Page(ctx context.Context, req *MoonbreezePageReq) (*MoonbreezePageResp, error)
}

type MoonbreezePageReq struct {
	AuthorIDs         []int64
	ExcludedAuthorIDs []int64
	BeforeAt          *time.Time
	BeforeID          *int64
	Size              int
}

type MoonbreezePageResp struct {
	Rows []*model.Moonbreeze
}
