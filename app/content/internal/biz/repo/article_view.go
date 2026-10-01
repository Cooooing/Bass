package repo

import (
	"content/internal/biz/base"
	"content/internal/biz/model"
	"context"
)

type ArticleViewCacheRepo interface {
	Record(ctx context.Context, req *ArticleViewCacheRecordReq) (bool, error)
	PopCounts(ctx context.Context, limit int32) (map[int64]int32, error)
}

type ArticleViewCacheRecordReq struct {
	ArticleID          int64
	ViewerUserID       *int64
	IP                 *string
	UserAgent          *string
	BrowserFingerprint *string
}

type ArticleViewRecordRepo interface {
	Save(ctx context.Context, record *model.ArticleViewRecord) error
	Page(ctx context.Context, req *ArticleViewRecordPageReq) (*ArticleViewRecordPageResp, error)
}

type ArticleViewRecordPageReq struct {
	UserID int64
	Page   *base.PageRequest
}

type ArticleViewRecordPageResp struct {
	Rows []*model.ArticleViewRecord
	Page *base.PageResp
}
