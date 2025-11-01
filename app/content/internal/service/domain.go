package service

import (
	v1 "common/api/content/v1"
	"content/internal/biz"
	"content/internal/biz/model"
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

func (s *DomainService) Add(ctx context.Context, req *v1.AddDomainRequest) (*v1.AddDomainReply, error) {
	_, err := s.domainDomain.AddDomain(ctx, &model.Domain{
		Name:        req.Name,
		Description: req.Description,
		URL:         &req.Url,
		Icon:        &req.Icon,
		IsNav:       req.IsNav,
	})
	if err != nil {
		return nil, err
	}
	return &v1.AddDomainReply{}, nil
}

func (s *DomainService) Get(ctx context.Context, req *v1.GetDomainRequest) (*v1.GetDomainReply, error) {

	// TODO implement me
	panic("implement me")
}

func (s *DomainService) Update(ctx context.Context, req *v1.UpdateDomainRequest) (*v1.UpdateDomainReply, error) {
	// TODO implement me
	panic("implement me")
}
