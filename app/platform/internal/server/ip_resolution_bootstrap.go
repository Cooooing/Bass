package server

import (
	"context"

	"platform/internal/biz/usecase"
)

// IPResolutionBootstrapServer 在应用启动时加载 IP 地址数据库。
type IPResolutionBootstrapServer struct {
	ipResolutionUsecase *usecase.IPResolutionUsecase
}

func NewIPResolutionBootstrapServer(ipResolutionUsecase *usecase.IPResolutionUsecase) *IPResolutionBootstrapServer {
	return &IPResolutionBootstrapServer{ipResolutionUsecase: ipResolutionUsecase}
}

func (s *IPResolutionBootstrapServer) Start(ctx context.Context) error {
	return s.ipResolutionUsecase.Start(ctx)
}

func (*IPResolutionBootstrapServer) Stop(context.Context) error { return nil }
