package usecase

import (
	"bbs/internal/biz/repo"
	bbscontentv1 "common/proto/gen/bbs/v1/content"
	"context"
)

type ContentPostscriptUsecase struct {
	contentPostscriptClient repo.ContentPostscriptClient
}

func NewContentPostscriptUsecase(contentPostscriptClient repo.ContentPostscriptClient) *ContentPostscriptUsecase {
	return &ContentPostscriptUsecase{contentPostscriptClient: contentPostscriptClient}
}

func (u *ContentPostscriptUsecase) AddPostscript(ctx context.Context, req *bbscontentv1.AddPostscript_Request) (*bbscontentv1.AddPostscript_Reply, error) {
	return u.contentPostscriptClient.AddPostscript(ctx, req)
}
