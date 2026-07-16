package repo

import "context"

type AddPostscriptReq struct {
	UserID    int64
	ArticleID int64
	Content   string
}

type AddPostscriptResponse struct {
	Postscript *ArticlePostscript
}

type ContentPostscriptClient interface {
	AddPostscript(ctx context.Context, req *AddPostscriptReq) (*AddPostscriptResponse, error)
}
