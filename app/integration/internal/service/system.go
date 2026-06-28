package service

import (
	commonv1 "common/proto/gen/common/v1"
	"context"
	"fmt"
	"integration/internal/conf"

	"github.com/go-kratos/kratos/v3/transport/grpc"
	"github.com/go-kratos/kratos/v3/transport/http"
)

type SystemService struct {
	commonv1.UnimplementedCommonSystemServiceServer
	conf *conf.Bootstrap
}

func NewSystemService(conf *conf.Bootstrap) *SystemService {
	return &SystemService{conf: conf}
}

func (s *SystemService) RegisterGrpc(gs *grpc.Server) {
	commonv1.RegisterCommonSystemServiceServer(gs, s)
}

func (s *SystemService) RegisterHttp(hs *http.Server) {
}

func (s *SystemService) Health(ctx context.Context, req *commonv1.HealthSystem_Request) (*commonv1.HealthSystem_Reply, error) {
	return &commonv1.HealthSystem_Reply{Message: fmt.Sprintf("%s %s is ok", s.conf.Server.Name, s.conf.Server.Version)}, nil
}
