package service

import (
	v1 "common/proto/gen/user/v1"
	"context"
	"user/internal/biz/model"
	"user/internal/biz/usecase"

	"github.com/go-kratos/kratos/v3/transport/grpc"
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

func (s *LocationService) Get(ctx context.Context, req *v1.GetLocation_Request) (*v1.GetLocation_Response, error) {
	res, err := s.locationUsecase.GetByUserID(ctx, &usecase.GetLocationByUserIDReq{UserID: req.GetUserId()})
	if err != nil {
		return nil, err
	}
	reply := &v1.GetLocation_Response_Location{UserId: req.GetUserId()}
	if res.Location != nil {
		reply.Country = res.Location.Country
		reply.Province = res.Location.Province
		reply.City = res.Location.City
	}
	return &v1.GetLocation_Response{Location: reply}, nil
}

func (s *LocationService) Upsert(ctx context.Context, req *v1.UpsertLocation_Request) (*v1.UpsertLocation_Response, error) {
	res, err := s.locationUsecase.UpsertByUserID(ctx, &usecase.UpsertLocationByUserIDReq{Location: &model.Location{
		UserID:   req.GetUserId(),
		Country:  req.Country,
		Province: req.Province,
		City:     req.City,
	}})
	if err != nil {
		return nil, err
	}
	location := res.Location
	return &v1.UpsertLocation_Response{Location: &v1.UpsertLocation_Response_Location{
		UserId:   req.GetUserId(),
		Country:  location.Country,
		Province: location.Province,
		City:     location.City,
	}}, nil
}
