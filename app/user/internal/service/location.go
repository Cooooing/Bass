package service

import (
	v1 "common/proto/gen/user/v1"
	"context"
	"user/internal/biz/usecase"

	"github.com/go-kratos/kratos/v3/transport/grpc"
	"github.com/go-kratos/kratos/v3/transport/http"
)

type LocationService struct {
	v1.UnimplementedLocationServiceServer
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
	v1.RegisterLocationServiceServer(gs, s)
}

func (s *LocationService) RegisterHttp(hs *http.Server) {
}

func (s *LocationService) Get(ctx context.Context, req *v1.GetLocation_Req) (*v1.GetLocation_Resp, error) {
	res, err := s.locationUsecase.GetByUserID(ctx, req.GetUserId())
	if err != nil {
		return nil, err
	}
	if res == nil {
		return &v1.GetLocation_Resp{}, nil
	}
	return &v1.GetLocation_Resp{
		Location: &v1.GetLocation_Resp_Location{
			UserId:   res.UserID,
			Country:  res.Country,
			Province: res.Province,
			City:     res.City,
		},
	}, nil
}
