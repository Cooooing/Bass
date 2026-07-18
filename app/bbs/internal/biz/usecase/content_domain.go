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
	Page  *common.PageReq
	Query *bbscontentv1.ListDomains_Req_DomainQuery
}

type ListDomainsResp struct {
	Page *repo.PageResp
	Rows []*repo.Domain
}

func (u *ContentDomainUsecase) ListDomains(ctx context.Context, req *ListDomainsReq) (*ListDomainsResp, error) {
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
	resp, err := u.contentDomainClient.ListDomains(ctx, &repo.ListDomainsReq{Page: page, Query: query})
	if err != nil {
		return nil, err
	}
	return &ListDomainsResp{Page: resp.Page, Rows: resp.Rows}, nil
}
