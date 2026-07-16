package usecase

import (
	"bbs/internal/biz/repo"
	bbscontentv1 "common/proto/gen/bbs/v1/content"
	"common/proto/gen/common"
	"context"
)

type ContentDomainUsecase struct {
	contentDomainClient repo.ContentDomainClient
}

func NewContentDomainUsecase(contentDomainClient repo.ContentDomainClient) *ContentDomainUsecase {
	return &ContentDomainUsecase{contentDomainClient: contentDomainClient}
}

type ListDomainsReq struct {
	Page  *common.PageRequest
	Query *bbscontentv1.ListDomains_Request_DomainQuery
}

type ListDomainsResponse struct {
	Page *repo.PageResponse
	Rows []*repo.Domain
}

func (u *ContentDomainUsecase) ListDomains(ctx context.Context, req *ListDomainsReq) (*ListDomainsResponse, error) {
	if req == nil {
		req = &ListDomainsReq{}
	}
	var page *repo.PageReq
	if req.Page != nil {
		page = &repo.PageReq{Page: req.Page.GetPage(), Size: req.Page.GetSize()}
	}
	query := &repo.DomainQuery{}
	if req.Query != nil {
		query.IDs = req.Query.GetIds()
		query.Name = req.Query.Name
		query.Description = req.Query.Description
		query.URL = req.Query.Url
		query.Icon = req.Query.Icon
		query.IsNav = req.Query.IsNav
		if req.Query.Status != nil {
			value := int32(*req.Query.Status)
			query.Status = &value
		}
	}
	if query.Status == nil {
		value := int32(bbscontentv1.DomainStatus_DOMAIN_STATUS_ENABLED)
		query.Status = &value
	}
	response, err := u.contentDomainClient.ListDomains(ctx, &repo.ListDomainsReq{Page: page, Query: query})
	if err != nil {
		return nil, err
	}
	return &ListDomainsResponse{Page: response.Page, Rows: response.Rows}, nil
}
