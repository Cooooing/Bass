package data

import (
	"bbs/internal/biz/repo"
	"common/pkg/client/rpc"
	"common/proto/gen/common"
	contentv1 "common/proto/gen/content/v1"
	contentv1enum "common/proto/gen/content/v1/enum"
	"context"
)

var _ repo.ContentDomainClient = (*ContentDomainClient)(nil)

type ContentDomainClient struct {
	contentClient *rpc.ContentClient
}

func NewContentDomainClient(
	contentClient *rpc.ContentClient,
) repo.ContentDomainClient {
	return &ContentDomainClient{
		contentClient: contentClient,
	}
}

func (r *ContentDomainClient) CreateDomain(ctx context.Context, req *repo.CreateDomainReq) (*repo.Domain, error) {
	domain := req.Domain
	var status *contentv1enum.DomainStatus
	if domain.Status != nil {
		status = new(contentv1enum.DomainStatus(*domain.Status))
	}
	reply, err := r.contentClient.Domain.BatchCreate(ctx, &contentv1.BatchCreateDomains_Req{
		UserId: req.UserID,
		Domains: []*contentv1.BatchCreateDomains_Req_Domain{
			{
				Code:        domain.Code,
				Name:        domain.Name,
				Description: domain.Description,
				Status:      status,
				Url:         domain.URL,
				Icon:        domain.Icon,
				IsNav:       domain.IsNav,
				Sort:        domain.Sort,
			},
		},
	})
	if err != nil {
		return nil, err
	}
	rows := reply.GetRows()
	if len(rows) == 0 {
		return nil, nil
	}
	item := rows[0]
	row := &repo.Domain{
		ID:          item.GetId(),
		Code:        item.GetCode(),
		Name:        item.GetName(),
		Description: item.Description,
		Status:      int32(item.GetStatus()),
		URL:         item.Url,
		Icon:        item.Icon,
		IsNav:       item.GetIsNav(),
		Sort:        item.GetSort(),
		CreatedBy:   item.CreatedBy,
		UpdatedBy:   item.UpdatedBy,
		CreatedAt:   new(item.GetCreatedAt().AsTime()),
		UpdatedAt:   new(item.GetUpdatedAt().AsTime()),
	}
	return row, nil
}

func (r *ContentDomainClient) UpdateDomain(ctx context.Context, req *repo.UpdateDomainReq) (*repo.Domain, error) {
	domain := req.Domain
	var status *contentv1enum.DomainStatus
	if domain.Status != nil {
		status = new(contentv1enum.DomainStatus(*domain.Status))
	}
	reply, err := r.contentClient.Domain.Update(ctx, &contentv1.UpdateDomain_Req{
		UserId: req.UserID,
		Domain: &contentv1.UpdateDomain_Req_Domain{
			Id:          req.DomainID,
			Code:        domain.Code,
			Name:        domain.Name,
			Description: domain.Description,
			Status:      status,
			Url:         domain.URL,
			Icon:        domain.Icon,
			IsNav:       domain.IsNav,
			Sort:        domain.Sort,
		},
	})
	if err != nil {
		return nil, err
	}
	item := reply.GetDomain()
	row := &repo.Domain{
		ID:          item.GetId(),
		Code:        item.GetCode(),
		Name:        item.GetName(),
		Description: item.Description,
		Status:      int32(item.GetStatus()),
		URL:         item.Url,
		Icon:        item.Icon,
		IsNav:       item.GetIsNav(),
		Sort:        item.GetSort(),
		CreatedBy:   item.CreatedBy,
		UpdatedBy:   item.UpdatedBy,
		CreatedAt:   new(item.GetCreatedAt().AsTime()),
		UpdatedAt:   new(item.GetUpdatedAt().AsTime()),
	}
	return row, nil
}

func (r *ContentDomainClient) ListDomains(ctx context.Context, req *repo.ListDomainsReq) (*repo.ListDomainsResp, error) {
	query := req.Query
	if query == nil {
		query = &repo.DomainQuery{}
	}
	contentQuery := &contentv1.PageDomains_Req_Query{
		Ids:         query.IDs,
		Code:        query.Code,
		Name:        query.Name,
		Description: query.Description,
		Url:         query.URL,
		Icon:        query.Icon,
		IsNav:       query.IsNav,
	}
	if query.Status != nil {
		contentQuery.Status = new(contentv1enum.DomainStatus(*query.Status))
	}
	var pageReq *common.PageReq
	if req.Page != nil {
		pageReq = &common.PageReq{
			Page: req.Page.Page,
			Size: req.Page.Size,
		}
	}
	reply, err := r.contentClient.Domain.Page(ctx, &contentv1.PageDomains_Req{
		Page:  pageReq,
		Query: contentQuery,
	})
	if err != nil {
		return nil, err
	}
	rows := make([]*repo.Domain, 0, len(reply.GetRows()))
	for _, item := range reply.GetRows() {
		row := &repo.Domain{
			ID:          item.GetId(),
			Code:        item.GetCode(),
			Name:        item.GetName(),
			Description: item.Description,
			Status:      int32(item.GetStatus()),
			URL:         item.Url,
			Icon:        item.Icon,
			IsNav:       item.GetIsNav(),
			Sort:        item.GetSort(),
			CreatedBy:   item.CreatedBy,
			UpdatedBy:   item.UpdatedBy,
			CreatedAt:   new(item.GetCreatedAt().AsTime()),
			UpdatedAt:   new(item.GetUpdatedAt().AsTime()),
		}
		rows = append(rows, row)
	}
	var page *repo.PageResp
	if reply.GetPage() != nil {
		page = &repo.PageResp{
			Page:  reply.GetPage().GetPage(),
			Size:  reply.GetPage().GetSize(),
			Total: reply.GetPage().GetTotal(),
		}
	}
	return &repo.ListDomainsResp{
		Page: page,
		Rows: rows,
	}, nil
}
