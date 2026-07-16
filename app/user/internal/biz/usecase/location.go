package usecase

import (
	"context"
	"user/internal/biz/model"
	"user/internal/biz/repo"
)

type LocationUsecase struct {
	locationRepo repo.LocationRepo
}

func NewLocationUsecase(locationRepo repo.LocationRepo) *LocationUsecase {
	return &LocationUsecase{locationRepo: locationRepo}
}

type GetLocationByUserIDReq struct {
	UserID int64
}

type GetLocationByUserIDResponse struct {
	Location *model.Location
}

func (s *LocationUsecase) GetByUserID(ctx context.Context, req *GetLocationByUserIDReq) (*GetLocationByUserIDResponse, error) {
	locationResp, err := s.locationRepo.Get(ctx, &repo.LocationGetReq{UserID: &req.UserID})
	if err != nil {
		return nil, err
	}
	return &GetLocationByUserIDResponse{Location: locationResp.Location}, nil
}

type UpsertLocationByUserIDReq struct {
	Location *model.Location
}

type UpsertLocationByUserIDResponse struct {
	Location *model.Location
}

func (s *LocationUsecase) UpsertByUserID(ctx context.Context, req *UpsertLocationByUserIDReq) (*UpsertLocationByUserIDResponse, error) {
	locationResp, err := s.locationRepo.UpsertByUserID(ctx, &repo.LocationUpsertByUserIDReq{Location: req.Location})
	if err != nil {
		return nil, err
	}
	return &UpsertLocationByUserIDResponse{Location: locationResp.Location}, nil
}
