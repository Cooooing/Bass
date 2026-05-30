package usecase

import (
	"bbs/internal/biz/repo"
	bbscontentv1 "common/api/gen/bbs/v1/content"
	"context"
)

type ContentDomainUsecase struct {
	contentDomainRepo repo.ContentDomainRepo
}

func NewContentDomainUsecase(contentDomainRepo repo.ContentDomainRepo) *ContentDomainUsecase {
	return &ContentDomainUsecase{contentDomainRepo: contentDomainRepo}
}

func (u *ContentDomainUsecase) ListDomains(ctx context.Context, req *bbscontentv1.ListDomains_Request) (*bbscontentv1.ListDomains_Reply, error) {
	return u.contentDomainRepo.ListDomains(ctx, req)
}
