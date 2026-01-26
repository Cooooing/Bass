package service

import (
	"common/api/common/v1"
	"context"
	"fmt"

	"github.com/go-kratos/kratos/v2/transport/grpc"
	"github.com/go-kratos/kratos/v2/transport/http"
)

type SystemService struct {
	v1.UnimplementedSystemServer
	*BaseService
}

func NewSystemService(baseService *BaseService) *SystemService {
	return &SystemService{
		BaseService: baseService,
	}
}

func (s *SystemService) RegisterGrpc(gs *grpc.Server) {
	v1.RegisterSystemServer(gs, s)
}

func (s *SystemService) RegisterHttp(hs *http.Server) {
	v1.RegisterSystemHTTPServer(hs, s)
}

func (s *SystemService) Health(ctx context.Context, req *v1.HealthRequest) (*v1.HealthReply, error) {
	return &v1.HealthReply{Message: fmt.Sprintf("%s %s is ok", s.conf.Server.Name, s.conf.Server.Version)}, nil
}
