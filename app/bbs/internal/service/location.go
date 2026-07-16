package service

import (
	"bbs/internal/biz/usecase"
	bbsuserv1 "common/proto/gen/bbs/v1/user"
	"context"

	"github.com/go-kratos/kratos/v3/transport/grpc"
	"github.com/go-kratos/kratos/v3/transport/http"
)

type LocationService struct {
	bbsuserv1.UnimplementedLocationServiceServer
	locationUsecase *usecase.LocationUsecase
}

func NewLocationService(locationUsecase *usecase.LocationUsecase) *LocationService {
	return &LocationService{locationUsecase: locationUsecase}
}

func (s *LocationService) RegisterGrpc(gs *grpc.Server) {}

func (s *LocationService) RegisterHttp(hs *http.Server) {
	bbsuserv1.RegisterLocationServiceHTTPServer(hs, s)
}

func (s *LocationService) GetCurrent(ctx context.Context, req *bbsuserv1.GetCurrentLocation_Request) (*bbsuserv1.GetCurrentLocation_Response, error) {
	userID, err := currentUserID(ctx)
	if err != nil {
		return nil, err
	}
	response, err := s.locationUsecase.GetCurrentLocation(ctx, &usecase.GetCurrentLocationReq{UserID: userID})
	if err != nil {
		return nil, err
	}
	return &bbsuserv1.GetCurrentLocation_Response{Location: response.Location}, nil
}

func (s *LocationService) UpsertCurrent(ctx context.Context, req *bbsuserv1.UpsertCurrentLocation_Request) (*bbsuserv1.UpsertCurrentLocation_Response, error) {
	userID, err := currentUserID(ctx)
	if err != nil {
		return nil, err
	}
	response, err := s.locationUsecase.UpsertCurrentLocation(ctx, &usecase.UpsertCurrentLocationReq{UserID: userID, Country: req.Country, Province: req.Province, City: req.City})
	if err != nil {
		return nil, err
	}
	return &bbsuserv1.UpsertCurrentLocation_Response{Location: response.Location}, nil
}
