package repo

import (
	"content/internal/biz/model"
	"context"
	"time"
)

type BreezemoonRepo interface {
	Create(ctx context.Context, breezemoon *model.Breezemoon) (*model.Breezemoon, error)
	Page(ctx context.Context, req *BreezemoonPageReq) (*BreezemoonPageResp, error)
}

type BreezemoonPageReq struct {
	AuthorIDs []int64
	BeforeAt  *time.Time
	BeforeID  *int64
	Size      int
}

type BreezemoonPageResp struct {
	Rows []*model.Breezemoon
}
