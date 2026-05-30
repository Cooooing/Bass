package usecase

import (
	"bbs/internal/biz/repo"
	bbscontentv1 "common/api/gen/bbs/v1/content"
	"context"
)

type ContentPostscriptUsecase struct {
	contentPostscriptRepo repo.ContentPostscriptRepo
}

func NewContentPostscriptUsecase(contentPostscriptRepo repo.ContentPostscriptRepo) *ContentPostscriptUsecase {
	return &ContentPostscriptUsecase{contentPostscriptRepo: contentPostscriptRepo}
}

func (u *ContentPostscriptUsecase) AddPostscript(ctx context.Context, req *bbscontentv1.AddPostscript_Request) (*bbscontentv1.AddPostscript_Reply, error) {
	return u.contentPostscriptRepo.AddPostscript(ctx, req)
}
