package repo

import (
	"content/internal/biz/base"
	"content/internal/biz/model"
	"content/internal/enum"
	"context"
)

type PostscriptRepo interface {
	Save(ctx context.Context, postscript *model.Postscript) (*model.Postscript, error)
	Get(ctx context.Context, req *PostscriptGetReq) (*model.Postscript, error)
	List(ctx context.Context, req *PostscriptGetReq) ([]*model.Postscript, error)
	Map(ctx context.Context, req *PostscriptGetReq) (map[int64]*model.Postscript, error)
	Count(ctx context.Context, req *PostscriptGetReq) (int, error)
	Page(ctx context.Context, req *PostscriptGetReq) (*PostscriptPageResp, error)
}

type PostscriptPageResp struct {
	Rows []*model.Postscript
	Page *base.PageResp
}

type PostscriptGetReq struct {
	Page         *base.PageRequest
	ID           *int64
	IDs          []int64
	ArticleID    *int64
	ArticleIDs   []int64
	CreatedBy    *int64
	Restriction  *enum.ContentRestriction
	Restrictions []enum.ContentRestriction
}
