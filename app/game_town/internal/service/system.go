package service

import (
	commonv1 "common/proto/gen/common/v1"
	"context"

	"github.com/go-kratos/kratos/v3/transport/grpc"
	"github.com/go-kratos/kratos/v3/transport/http"
)

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
func (s *CommonSystemService) Health(ctx context.Context, req *commonv1.HealthSystem_Request) (*commonv1.HealthSystem_Response, error) {
	return &commonv1.HealthSystem_Response{Message: "ok"}, nil
}
