package service

import (
	v1 "common/proto/gen/platform/v1"
	"context"

	"platform/internal/biz/usecase"

	"github.com/go-kratos/kratos/v3/transport/grpc"
	"github.com/go-kratos/kratos/v3/transport/http"
)

// IPResolutionService 实现 platform IP 解析 gRPC 服务。
type IPResolutionService struct {
	v1.UnimplementedPlatformIPResolutionServiceServer
	ipUsecase *usecase.IPResolutionUsecase
}

func NewIPResolutionService(
	ipUsecase *usecase.IPResolutionUsecase,
) *IPResolutionService {
	return &IPResolutionService{
		ipUsecase: ipUsecase,
	}
}

func (s *IPResolutionService) RegisterGrpc(gs *grpc.Server) {
	v1.RegisterPlatformIPResolutionServiceServer(gs, s)
}

func (s *IPResolutionService) RegisterHTTP(hs *http.Server) {
}

func (s *IPResolutionService) ResolveIP(ctx context.Context, req *v1.ResolveIP_Req) (*v1.ResolveIP_Resp, error) {
	info, err := s.ipUsecase.Get(ctx, req.GetIp())
	if err != nil {
		return nil, err
	}
	return &v1.ResolveIP_Resp{
		Country:     info.Country,
		Province:    info.Province,
		City:        info.City,
		Isp:         info.ISP,
		CountryCode: info.CountryCode,
	}, nil
}

func (s *IPResolutionService) UpdateIPDataFromSource(ctx context.Context, req *v1.UpdateIPDataFromSource_Req) (*v1.UpdateIPDataFromSource_Resp, error) {
	if err := s.ipUsecase.UpdateIPDataFromSource(ctx); err != nil {
		return nil, err
	}
	return &v1.UpdateIPDataFromSource_Resp{}, nil
}

func (s *IPResolutionService) UpdateIPDataFromOSS(ctx context.Context, req *v1.UpdateIPDataFromOSS_Req) (*v1.UpdateIPDataFromOSS_Resp, error) {
	if err := s.ipUsecase.UpdateIPDataFromOss(ctx); err != nil {
		return nil, err
	}
	return &v1.UpdateIPDataFromOSS_Resp{}, nil
}
