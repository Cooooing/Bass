package service

import (
	v1 "common/proto/gen/user/v1"
	"context"
	"user/internal/biz/model"
	"user/internal/biz/usecase"

	"github.com/go-kratos/kratos/v3/transport/grpc"
	"github.com/go-kratos/kratos/v3/transport/http"
)

type LocationService struct {
	v1.UnimplementedLocationServiceServer
	locationUsecase *usecase.LocationUsecase
}

func NewLocationService(locationUsecase *usecase.LocationUsecase) *LocationService {
	return &LocationService{locationUsecase: locationUsecase}
}

func (s *LocationService) RegisterGrpc(gs *grpc.Server) {
	v1.RegisterLocationServiceServer(gs, s)
}

func (s *LocationService) RegisterHttp(hs *http.Server) {}

func (s *LocationService) Get(ctx context.Context, req *v1.GetLocation_Req) (*v1.GetLocation_Resp, error) {
	res, err := s.locationUsecase.GetByUserID(ctx, req.GetUserId())
	if err != nil {
		return nil, err
	}
	reply := &v1.GetLocation_Resp_Location{UserId: req.GetUserId()}
	if res != nil {
		reply.Country = res.Country
		reply.Province = res.Province
		reply.City = res.City
	}
	return &v1.GetLocation_Resp{Location: reply}, nil
}

func (s *LocationService) Upsert(ctx context.Context, req *v1.UpsertLocation_Req) (*v1.UpsertLocation_Resp, error) {
	res, err := s.locationUsecase.UpsertByUserID(ctx, &model.Location{
		UserID:   req.GetUserId(),
		Country:  req.Country,
		Province: req.Province,
		City:     req.City,
	})
	if err != nil {
		return nil, err
	}
	location := res
	return &v1.UpsertLocation_Resp{Location: &v1.UpsertLocation_Resp_Location{
		UserId:   req.GetUserId(),
		Country:  location.Country,
		Province: location.Province,
		City:     location.City,
	}}, nil
}
