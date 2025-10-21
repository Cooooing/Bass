package service

import (
	v1 "common/api/content/v1"
	"content/internal/biz"
	"content/internal/biz/model"
	"content/internal/data/ent/gen"
	"context"

	"github.com/go-kratos/kratos/v2/transport/grpc"
	"github.com/go-kratos/kratos/v2/transport/http"
)

type DomainService struct {
	v1.UnimplementedDomainServiceServer
	*BaseService
	domainDomain *biz.DomainDomain
}

func (s *DomainService) RegisterGrpc(gs *grpc.Server) {
	v1.RegisterDomainServiceServer(gs, s)
}

func (s *DomainService) RegisterHttp(hs *http.Server) {
	v1.RegisterDomainServiceHTTPServer(hs, s)
}

func NewDomainService(baseService *BaseService, domainDomain *biz.DomainDomain) *DomainService {
	return &DomainService{
		BaseService:  baseService,
		domainDomain: domainDomain,
	}
}

func (s *DomainService) Add(ctx context.Context, req *v1.AddDomainRequest) (*v1.AddDomainReply, error) {
	s.domainDomain.AddDomain(ctx, &model.Domain{
		Name:        req.Name,
		Description: "",
		URL:         nil,
		Icon:        nil,
		IsNav:       false,
	})
}

func (s *DomainService) Get(ctx context.Context, req *v1.GetDomainRequest) (*v1.GetDomainReply, error) {

	// TODO implement me
	panic("implement me")
}

func (s *DomainService) Update(ctx context.Context, req *v1.UpdateDomainRequest) (*v1.UpdateDomainReply, error) {
	// TODO implement me
	panic("implement me")
}
