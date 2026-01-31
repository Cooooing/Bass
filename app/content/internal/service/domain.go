package service

import (
	v1 "common/api/content/v1"
	"common/pkg/cutil/base"
	commonModel "common/pkg/model"
	"content/internal/biz"
	"content/internal/biz/model"
	"content/internal/biz/repo"
	"content/internal/data/ent/gen"
	"context"

	"github.com/go-kratos/kratos/v2/transport/grpc"
	"github.com/go-kratos/kratos/v2/transport/http"
)

type DomainService struct {
	v1.UnimplementedContentDomainServiceServer
	*BaseService
	domainDomain *biz.DomainDomain
}

func (s *DomainService) RegisterGrpc(gs *grpc.Server) {
	v1.RegisterContentDomainServiceServer(gs, s)
}

func (s *DomainService) RegisterHttp(hs *http.Server) {
	v1.RegisterContentDomainServiceHTTPServer(hs, s)
}

func NewDomainService(baseService *BaseService, domainDomain *biz.DomainDomain) *DomainService {
	return &DomainService{
		BaseService:  baseService,
		domainDomain: domainDomain,
	}
}

func (s *DomainService) Adds(ctx context.Context, req *v1.AddDomainsRequest) (*v1.AddDomainsReply, error) {
	domains := make([]*model.Domain, len(req.Domains))
	for i, domain := range req.Domains {
		domains[i] = &model.Domain{Domain: &gen.Domain{
			Name:        domain.Name,
			Description: domain.Description,
			Status:      int32(base.DerefOrDefault(domain.Status, v1.DomainStatus_DOMAIN_STATUS_NORMAL)),
			URL:         domain.Url,
			Icon:        domain.Icon,
			IsNav:       domain.IsNav,
		}}
	}
	_, err := s.domainDomain.Adds(ctx, domains)
	if err != nil {
		return nil, err
	}
	return &v1.AddDomainsReply{}, nil
}

func (s *DomainService) Update(ctx context.Context, req *v1.UpdateDomainRequest) (*v1.UpdateDomainReply, error) {
	data, err := s.domainDomain.Update(ctx, &model.Domain{Domain: &gen.Domain{
		Name:        req.Domain.Name,
		Description: req.Domain.Description,
		Status:      int32(base.DerefOrDefault(req.Domain.Status, v1.DomainStatus_DOMAIN_STATUS_NORMAL)),
		URL:         req.Domain.Url,
		Icon:        req.Domain.Icon,
		IsNav:       req.Domain.IsNav,
	}})
	if err != nil {
		return nil, err
	}
	return &v1.UpdateDomainReply{
		Data: data.ConvertToRpc(),
	}, err
}

func (s *DomainService) Page(ctx context.Context, req *v1.PageDomainRequest) (*v1.PageDomainReply, error) {
	req.Query = base.OrDefault(req.Query, &v1.DomainQueryParams{})
	getReq := &repo.DomainGetReq{
		DomainIds:   req.Query.Ids,
		Name:        req.Query.Name,
		Description: req.Query.Description,
		Status:      base.Ptr(v1.DomainStatus_DOMAIN_STATUS_NORMAL),
		Url:         req.Query.Url,
		Icon:        req.Query.Icon,
		TagCount:    nil,
		IsNav:       req.Query.IsNav,
	}
	if req.Query.Status != nil {
		getReq.Status = base.Ptr(*req.Query.Status)
	}
	if req.Query.TagCount != nil {
		getReq.TagCount = &commonModel.Range[int32]{
			Start: req.Query.TagCount.Start,
			End:   req.Query.TagCount.End,
		}
	}
	data, page, err := s.domainDomain.Page(ctx, req.Page, getReq)
	reply := make([]*v1.Domain, 0, len(data))
	for _, datum := range data {
		reply = append(reply, datum.ConvertToRpc())
	}
	return &v1.PageDomainReply{
		Page: page,
		Rows: reply,
	}, err
}
