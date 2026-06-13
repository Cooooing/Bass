package service

import (
	"common/proto/gen/integration/v1"
	"context"
	"platform/internal/biz/usecase"

	"github.com/go-kratos/kratos/v2/transport/grpc"
	"github.com/go-kratos/kratos/v2/transport/http"
)

// IpResolutionService 实现 IP 解析 gRPC 服务。
type IpResolutionService struct {
	v1.UnimplementedIntegrationIpResolutionServiceServer
	ipUsecase *usecase.IpResolutionUsecase
}

func NewIpResolutionService(ipUsecase *usecase.IpResolutionUsecase) *IpResolutionService {
	return &IpResolutionService{ipUsecase: ipUsecase}
}

func (s *IpResolutionService) RegisterGrpc(gs *grpc.Server) {
	v1.RegisterIntegrationIpResolutionServiceServer(gs, s)
}

func (s *IpResolutionService) RegisterHttp(hs *http.Server) {
}

// ResolveIp 解析 IP 地理信息。
func (s *IpResolutionService) ResolveIp(ctx context.Context, req *v1.ResolveIp_Request) (*v1.ResolveIp_Reply, error) {
	info, err := s.ipUsecase.Get(ctx, req.GetIp())
	if err != nil {
		return nil, err
	}
	return &v1.ResolveIp_Reply{
		Country:     info.Country,
		Province:    info.Province,
		City:        info.City,
		Isp:         info.ISP,
		CountryCode: info.CountryCode,
	}, nil
}
