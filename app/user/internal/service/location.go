package service

import (
	cerrors "common/api/gen/common/errors"
	v1 "common/api/gen/user/v1"
	"common/pkg/constant"
	commonModel "common/pkg/model"
	"common/pkg/util"
	"context"
	"user/internal/biz/model"
	"user/internal/biz/repo"

	"github.com/go-kratos/kratos/v2/transport/grpc"
	"github.com/go-kratos/kratos/v2/transport/http"
)

type LocationService struct {
	v1.UnimplementedLocationServiceServer
	locationRepo repo.LocationRepo
}

func NewLocationService(locationRepo repo.LocationRepo) *LocationService {
	return &LocationService{locationRepo: locationRepo}
}

func (s *LocationService) RegisterGrpc(gs *grpc.Server) {
	v1.RegisterLocationServiceServer(gs, s)
}

func (s *LocationService) RegisterHttp(hs *http.Server) {}

func (s *LocationService) GetCurrent(ctx context.Context, req *v1.GetCurrentLocation_Request) (*v1.GetCurrentLocation_Reply, error) {
	current, ok := util.GetContextValue[*commonModel.User](ctx, constant.CtxUserInfo)
	if !ok {
		return nil, cerrors.ErrorUnauthorized("user not login")
	}
	location, err := s.locationRepo.FindByUserID(ctx, current.ID)
	if err != nil {
		return nil, err
	}
	reply := &v1.Location{UserId: current.ID}
	if location != nil {
		reply.Country = location.Country
		reply.Province = location.Province
		reply.City = location.City
	}
	return &v1.GetCurrentLocation_Reply{Location: reply}, nil
}

func (s *LocationService) UpsertCurrent(ctx context.Context, req *v1.UpsertCurrentLocation_Request) (*v1.UpsertCurrentLocation_Reply, error) {
	current, ok := util.GetContextValue[*commonModel.User](ctx, constant.CtxUserInfo)
	if !ok {
		return nil, cerrors.ErrorUnauthorized("user not login")
	}
	location, err := s.locationRepo.UpsertByUserID(ctx, &model.Location{
		UserID:   current.ID,
		Country:  req.Country,
		Province: req.Province,
		City:     req.City,
	})
	if err != nil {
		return nil, err
	}
	return &v1.UpsertCurrentLocation_Reply{Location: &v1.Location{
		UserId:   current.ID,
		Country:  location.Country,
		Province: location.Province,
		City:     location.City,
	}}, nil
}
