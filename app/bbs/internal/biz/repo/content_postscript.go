package repo

import "context"

type AddPostscriptReq struct {
	UserID    int64
	ArticleID int64
	Content   string
}

type ListPostscriptsReq struct {
	ArticleID int64
}

type ContentPostscriptClient interface {
	AddPostscript(ctx context.Context, req *AddPostscriptReq) (*ArticlePostscript, error)
	ListPostscripts(ctx context.Context, req *ListPostscriptsReq) ([]*ArticlePostscript, error)
}
