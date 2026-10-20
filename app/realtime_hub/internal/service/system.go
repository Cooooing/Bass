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
func (s *CommonSystemService) RegisterGrpc(server *grpc.Server) {
	commonv1.RegisterCommonSystemServiceServer(server, s)
}
func (s *CommonSystemService) RegisterHTTP(server *http.Server) {
	commonv1.RegisterCommonSystemServiceHTTPServer(server, s)
}
func (*CommonSystemService) Health(context.Context, *commonv1.HealthSystem_Req) (*commonv1.HealthSystem_Resp, error) {
	return &commonv1.HealthSystem_Resp{Message: "ok"}, nil
}
