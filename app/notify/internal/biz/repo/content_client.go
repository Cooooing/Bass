package repo

import (
	"context"
	"notify/internal/biz/model"
)

type ContentClient interface {
	GetArticle(ctx context.Context, req *ContentGetArticleReq) (*ContentGetArticleResponse, error)
	GetComment(ctx context.Context, req *ContentGetCommentReq) (*ContentGetCommentResponse, error)
}

type ContentGetArticleReq struct {
	ArticleID int64
}

type ContentGetArticleResponse struct {
	Article *model.ContentArticle
}

type ContentGetCommentReq struct {
	CommentID int64
}

type ContentGetCommentResponse struct {
	Comment *model.ContentComment
}
