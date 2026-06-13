package usecase

import (
	"bbs/internal/biz/repo"
	bbsuserv1 "common/proto/gen/bbs/v1/user"
	"context"
)

type LocationUsecase struct {
	locationClient repo.LocationClient
}

func NewLocationUsecase(locationClient repo.LocationClient) *LocationUsecase {
	return &LocationUsecase{locationClient: locationClient}
}

func (u *LocationUsecase) GetCurrentLocation(ctx context.Context, req *bbsuserv1.GetCurrentLocation_Request) (*bbsuserv1.GetCurrentLocation_Reply, error) {
	return u.locationClient.GetCurrentLocation(ctx, req)
}

func (u *LocationUsecase) UpsertCurrentLocation(ctx context.Context, req *bbsuserv1.UpsertCurrentLocation_Request) (*bbsuserv1.UpsertCurrentLocation_Reply, error) {
	return u.locationClient.UpsertCurrentLocation(ctx, req)
}
