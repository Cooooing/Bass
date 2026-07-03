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

func (s *LocationService) Get(ctx context.Context, req *v1.GetLocation_Request) (*v1.GetLocation_Reply, error) {
	location, err := s.locationUsecase.GetByUserID(ctx, req.GetUserId())
	if err != nil {
		return nil, err
	}
	reply := &v1.Location{UserId: req.GetUserId()}
	if location != nil {
		reply.Country = location.Country
		reply.Province = location.Province
		reply.City = location.City
	}
	return &v1.GetLocation_Reply{Location: reply}, nil
}

func (s *LocationService) Upsert(ctx context.Context, req *v1.UpsertLocation_Request) (*v1.UpsertLocation_Reply, error) {
	location, err := s.locationUsecase.UpsertByUserID(ctx, &model.Location{
		UserID:   req.GetUserId(),
		Country:  req.Country,
		Province: req.Province,
		City:     req.City,
	})
	if err != nil {
		return nil, err
	}
	return &v1.UpsertLocation_Reply{Location: &v1.Location{
		UserId:   req.GetUserId(),
		Country:  location.Country,
		Province: location.Province,
		City:     location.City,
	}}, nil
}
