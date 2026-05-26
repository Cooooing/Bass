package service

import (
	"bbs/internal/biz/usecase"
	bbsuserv1 "common/api/gen/bbs/v1/user"
	"context"

	"github.com/go-kratos/kratos/v2/transport/grpc"
	"github.com/go-kratos/kratos/v2/transport/http"
)

type LocationService struct {
	bbsuserv1.UnimplementedLocationServiceServer
	userUsecase *usecase.UserUsecase
}

func NewLocationService(userUsecase *usecase.UserUsecase) *LocationService {
	return &LocationService{userUsecase: userUsecase}
}

func (s *LocationService) RegisterGrpc(gs *grpc.Server) {
	bbsuserv1.RegisterLocationServiceServer(gs, s)
}

func (s *LocationService) RegisterHttp(hs *http.Server) {
	bbsuserv1.RegisterLocationServiceHTTPServer(hs, s)
}

func (s *LocationService) GetCurrent(ctx context.Context, req *bbsuserv1.GetCurrentLocation_Request) (*bbsuserv1.GetCurrentLocation_Reply, error) {
	return s.userUsecase.GetCurrentLocation(ctx, req)
}

func (s *LocationService) UpsertCurrent(ctx context.Context, req *bbsuserv1.UpsertCurrentLocation_Request) (*bbsuserv1.UpsertCurrentLocation_Reply, error) {
	return s.userUsecase.UpsertCurrentLocation(ctx, req)
}
