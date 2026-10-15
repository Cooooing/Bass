package service

import (
	"common/pkg/server"
	platformv1 "common/proto/gen/bff_bbs/v1/platform"
	"context"

	"bff_bbs/internal/biz/usecase"

	"github.com/go-kratos/kratos/v3/transport/grpc"
	"github.com/go-kratos/kratos/v3/transport/http"
)

type LocationService struct {
	platformv1.UnimplementedLocationServiceServer
	locationUsecase *usecase.LocationUsecase
}

func NewLocationService(
	locationUsecase *usecase.LocationUsecase,
) *LocationService {
	return &LocationService{
		locationUsecase: locationUsecase,
	}
}

func (s *LocationService) RegisterGrpc(gs *grpc.Server) {
}

func (s *LocationService) RegisterHTTP(hs *http.Server) {
	platformv1.RegisterLocationServiceHTTPServer(hs, s)
}

func (s *LocationService) GetCurrent(ctx context.Context, req *platformv1.GetCurrentLocation_Req) (*platformv1.GetCurrentLocation_Resp, error) {
	location, err := s.locationUsecase.GetCurrent(ctx, server.ClientIP(ctx))
	if err != nil {
		return nil, err
	}
	return &platformv1.GetCurrentLocation_Resp{
		Location: location,
	}, nil
}
