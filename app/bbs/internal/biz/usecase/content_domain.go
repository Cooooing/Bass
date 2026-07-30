package usecase

import (
	"bbs/internal/biz/repo"
	"common/pkg/apperror"
	bbscontentv1 "common/proto/gen/bbs/v1/content"
	bbscontentv1enum "common/proto/gen/bbs/v1/content/enum"
	"common/proto/gen/common"
	cerrors "common/proto/gen/common/errors"
	"context"
)

type ContentDomainUsecase struct {
	contentDomainClient repo.ContentDomainClient
}

func NewContentDomainUsecase(
	contentDomainClient repo.ContentDomainClient,
) *ContentDomainUsecase {
	return &ContentDomainUsecase{
		contentDomainClient: contentDomainClient,
	}
}

type ContentDomainSave struct {
	Code        string
	Name        string
	Description *string
	Status      *bbscontentv1enum.DomainStatus
	URL         *string
	Icon        *string
	IsNav       bool
	Sort        int32
}

type CreateDomainReq struct {
	UserID int64
	Domain *ContentDomainSave
}

func (u *ContentDomainUsecase) CreateDomain(ctx context.Context, req *CreateDomainReq) (*repo.Domain, error) {
	if req == nil || req.Domain == nil {
		return nil, apperror.New(cerrors.BusinessErrorCode_BUSINESS_ERROR_CODE_COMMON_INVALID_ARGUMENT)
	}
	domain := req.Domain
	var status *int32
	if domain.Status != nil {
		status = new(int32(*domain.Status))
	}
	resp, err := u.contentDomainClient.CreateDomain(ctx, &repo.CreateDomainReq{
		UserID: req.UserID,
		Domain: &repo.DomainSave{
			Code:        domain.Code,
			Name:        domain.Name,
			Description: domain.Description,
			Status:      status,
			URL:         domain.URL,
			Icon:        domain.Icon,
			IsNav:       domain.IsNav,
			Sort:        domain.Sort,
		},
	})
	if err != nil {
		return nil, err
	}
	return resp, nil
}

type UpdateDomainReq struct {
	UserID   int64
	DomainID int64
	Domain   *ContentDomainSave
}

func (u *ContentDomainUsecase) UpdateDomain(ctx context.Context, req *UpdateDomainReq) (*repo.Domain, error) {
	if req == nil || req.Domain == nil || req.DomainID <= 0 {
		return nil, apperror.New(cerrors.BusinessErrorCode_BUSINESS_ERROR_CODE_COMMON_INVALID_ARGUMENT)
	}
	domain := req.Domain
	var status *int32
	if domain.Status != nil {
		status = new(int32(*domain.Status))
	}
	resp, err := u.contentDomainClient.UpdateDomain(ctx, &repo.UpdateDomainReq{
		UserID:   req.UserID,
		DomainID: req.DomainID,
		Domain: &repo.DomainSave{
			Code:        domain.Code,
			Name:        domain.Name,
			Description: domain.Description,
			Status:      status,
			URL:         domain.URL,
			Icon:        domain.Icon,
			IsNav:       domain.IsNav,
			Sort:        domain.Sort,
		},
	})
	if err != nil {
		return nil, err
	}
	return resp, nil
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
		page = &repo.PageReq{
			Page: req.Page.GetPage(),
			Size: req.Page.GetSize(),
		}
	}
	query := &repo.DomainQuery{}
	if req.Query != nil {
		query.IDs = req.Query.GetIds()
		query.Code = req.Query.Code
		query.Name = req.Query.Name
		query.Description = req.Query.Description
		query.URL = req.Query.Url
		query.Icon = req.Query.Icon
		query.IsNav = req.Query.IsNav
		if req.Query.Status != nil {
			query.Status = new(int32(*req.Query.Status))
		}
	}
	if query.Status == nil {
		query.Status = new(int32(bbscontentv1enum.DomainStatus_DOMAIN_STATUS_ENABLED))
	}
	resp, err := u.contentDomainClient.ListDomains(ctx, &repo.ListDomainsReq{
		Page:  page,
		Query: query,
	})
	if err != nil {
		return nil, err
	}
	return &ListDomainsResp{
		Page: resp.Page,
		Rows: resp.Rows,
	}, nil
}
