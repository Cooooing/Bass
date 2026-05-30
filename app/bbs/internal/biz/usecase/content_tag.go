package usecase

import (
	"bbs/internal/biz/repo"
	bbscontentv1 "common/api/gen/bbs/v1/content"
	"context"
)

type ContentTagUsecase struct {
	contentTagRepo repo.ContentTagRepo
}

func NewContentTagUsecase(contentTagRepo repo.ContentTagRepo) *ContentTagUsecase {
	return &ContentTagUsecase{contentTagRepo: contentTagRepo}
}

func (u *ContentTagUsecase) ListTags(ctx context.Context, req *bbscontentv1.ListTags_Request) (*bbscontentv1.ListTags_Reply, error) {
	return u.contentTagRepo.ListTags(ctx, req)
}
