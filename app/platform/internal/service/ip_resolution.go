package service

import (
	v1 "common/proto/gen/platform/v1"
	"context"
	"platform/internal/biz/usecase"

	"github.com/go-kratos/kratos/v3/transport/grpc"
	"github.com/go-kratos/kratos/v3/transport/http"
)

// IpResolutionService 实现 platform IP 解析 gRPC 服务。
type IpResolutionService struct {
	v1.UnimplementedPlatformIpResolutionServiceServer
	ipUsecase *usecase.IpResolutionUsecase
}

func NewIpResolutionService(
	ipUsecase *usecase.IpResolutionUsecase,
) *IpResolutionService {
	return &IpResolutionService{
		ipUsecase: ipUsecase,
	}
}

func (s *IpResolutionService) RegisterGrpc(
	gs *grpc.Server,
) {
	v1.RegisterPlatformIpResolutionServiceServer(gs, s)
}

func (s *IpResolutionService) RegisterHttp(
	hs *http.Server,
) {
}

func (s *IpResolutionService) ResolveIp(
	ctx context.Context,
	req *v1.ResolveIp_Req,
) (*v1.ResolveIp_Resp, error) {
	info, err := s.ipUsecase.Get(ctx, req.GetIp())
	if err != nil {
		return nil, err
	}
	return &v1.ResolveIp_Resp{
		Country:     info.Country,
		Province:    info.Province,
		City:        info.City,
		Isp:         info.ISP,
		CountryCode: info.CountryCode,
	}, nil
}

func (s *IpResolutionService) UpdateIpDataFromSource(
	ctx context.Context,
	req *v1.UpdateIpDataFromSource_Req,
) (*v1.UpdateIpDataFromSource_Resp, error) {
	if err := s.ipUsecase.UpdateIpDataFromSource(ctx); err != nil {
		return nil, err
	}
	return &v1.UpdateIpDataFromSource_Resp{}, nil
}

func (s *IpResolutionService) UpdateIpDataFromOss(
	ctx context.Context,
	req *v1.UpdateIpDataFromOss_Req,
) (*v1.UpdateIpDataFromOss_Resp, error) {
	if err := s.ipUsecase.UpdateIpDataFromOss(ctx); err != nil {
		return nil, err
	}
	return &v1.UpdateIpDataFromOss_Resp{}, nil
}
