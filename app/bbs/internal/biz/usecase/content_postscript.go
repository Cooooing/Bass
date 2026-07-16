package usecase

import (
	"bbs/internal/biz/repo"
	"context"
)

type ContentPostscriptUsecase struct {
	contentPostscriptClient repo.ContentPostscriptClient
}

func NewContentPostscriptUsecase(contentPostscriptClient repo.ContentPostscriptClient) *ContentPostscriptUsecase {
	return &ContentPostscriptUsecase{contentPostscriptClient: contentPostscriptClient}
}

type AddPostscriptReq struct {
	UserID    int64
	ArticleID int64
	Content   string
}

type AddPostscriptResponse struct {
	Postscript *repo.ArticlePostscript
}

func (u *ContentPostscriptUsecase) AddPostscript(ctx context.Context, req *AddPostscriptReq) (*AddPostscriptResponse, error) {
	response, err := u.contentPostscriptClient.AddPostscript(ctx, &repo.AddPostscriptReq{UserID: req.UserID, ArticleID: req.ArticleID, Content: req.Content})
	if err != nil {
		return nil, err
	}
	return &AddPostscriptResponse{Postscript: response.Postscript}, nil
}
