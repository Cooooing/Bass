package service

import (
	commonv1 "common/proto/gen/common/v1"
	"context"

	"github.com/go-kratos/kratos/v3/transport/grpc"
	"github.com/go-kratos/kratos/v3/transport/http"
)

// CommonSystemService 仅提供进程健康检查；不承载管理端业务接口。
type CommonSystemService struct {
	commonv1.UnimplementedCommonSystemServiceServer
}

func NewCommonSystemService() *CommonSystemService { return &CommonSystemService{} }

func (s *CommonSystemService) RegisterGrpc(gs *grpc.Server) {
	commonv1.RegisterCommonSystemServiceServer(gs, s)
}

func (s *CommonSystemService) RegisterHttp(hs *http.Server) {
	commonv1.RegisterCommonSystemServiceHTTPServer(hs, s)
}

func (s *CommonSystemService) Health(context.Context, *commonv1.HealthSystem_Req) (*commonv1.HealthSystem_Resp, error) {
	return &commonv1.HealthSystem_Resp{Message: "ok"}, nil
}
