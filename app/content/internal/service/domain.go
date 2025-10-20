package service

import (
	v1 "common/api/content/v1"

	"github.com/go-kratos/kratos/v2/transport/grpc"
	"github.com/go-kratos/kratos/v2/transport/http"
)

type DomainService struct {
	v1.UnimplementedDomainServiceServer
	*BaseService
}

func (s *DomainService) RegisterGrpc(gs *grpc.Server) {
	v1.RegisterDomainServiceServer(gs, s)
}

func (s *DomainService) RegisterHttp(hs *http.Server) {
	v1.RegisterDomainServiceHTTPServer(hs, s)
}

func NewDomainService(baseService *BaseService) *DomainService {
	return &DomainService{
		BaseService: baseService,
	}
}
