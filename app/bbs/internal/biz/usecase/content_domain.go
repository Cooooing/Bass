package usecase

import (
	"bbs/internal/biz/repo"
	bbscontentv1 "common/proto/gen/bbs/v1/content"
	"context"
)

type ContentDomainUsecase struct {
	contentDomainClient repo.ContentDomainClient
}

func NewContentDomainUsecase(contentDomainClient repo.ContentDomainClient) *ContentDomainUsecase {
	return &ContentDomainUsecase{contentDomainClient: contentDomainClient}
}

func (u *ContentDomainUsecase) ListDomains(ctx context.Context, req *bbscontentv1.ListDomains_Request) (*bbscontentv1.ListDomains_Reply, error) {
	if req == nil {
		req = &bbscontentv1.ListDomains_Request{}
	}
	if req.Query == nil {
		req.Query = &bbscontentv1.DomainQuery{}
	}
	if req.Query.Status == nil {
		req.Query.Status = new(bbscontentv1.DomainStatus_DOMAIN_STATUS_ENABLED)
	}
	return u.contentDomainClient.ListDomains(ctx, req)
}
