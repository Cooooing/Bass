package service

import (
	"common/proto/gen/common/v1"
	"context"
	"fmt"
	"platform/internal/conf"

	"github.com/go-kratos/kratos/v2/transport/grpc"
	"github.com/go-kratos/kratos/v2/transport/http"
)

type SystemService struct {
	v1.UnimplementedCommonSystemServiceServer
	conf *conf.Bootstrap
}

func NewSystemService(conf *conf.Bootstrap) *SystemService {
	return &SystemService{conf: conf}
}

func (s *SystemService) RegisterGrpc(gs *grpc.Server) {
	v1.RegisterCommonSystemServiceServer(gs, s)
}

func (s *SystemService) RegisterHttp(hs *http.Server) {
}

func (s *SystemService) Health(ctx context.Context, req *v1.HealthSystem_Request) (*v1.HealthSystem_Reply, error) {
	return &v1.HealthSystem_Reply{Message: fmt.Sprintf("%s %s is ok", s.conf.Server.Name, s.conf.Server.Version)}, nil
}
