package repo

import (
	"context"

	"bff_bbs/internal/biz/model"
)

type AddPostscriptReq struct {
	UserID    int64
	ArticleID int64
	Content   string
}

type ListPostscriptsReq struct {
	ArticleID int64
	UserID    int64
}

type ContentPostscriptClient interface {
	AddPostscript(ctx context.Context, req *AddPostscriptReq) (*model.ArticlePostscript, error)
	ListPostscripts(ctx context.Context, req *ListPostscriptsReq) ([]*model.ArticlePostscript, error)
}
