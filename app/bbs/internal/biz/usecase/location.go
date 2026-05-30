package usecase

import (
	"bbs/internal/biz/repo"
	bbsuserv1 "common/api/gen/bbs/v1/user"
	"context"
)

type LocationUsecase struct {
	locationRepo repo.LocationRepo
}

func NewLocationUsecase(locationRepo repo.LocationRepo) *LocationUsecase {
	return &LocationUsecase{locationRepo: locationRepo}
}

func (u *LocationUsecase) GetCurrentLocation(ctx context.Context, req *bbsuserv1.GetCurrentLocation_Request) (*bbsuserv1.GetCurrentLocation_Reply, error) {
	return u.locationRepo.GetCurrentLocation(ctx, req)
}

func (u *LocationUsecase) UpsertCurrentLocation(ctx context.Context, req *bbsuserv1.UpsertCurrentLocation_Request) (*bbsuserv1.UpsertCurrentLocation_Reply, error) {
	return u.locationRepo.UpsertCurrentLocation(ctx, req)
}
