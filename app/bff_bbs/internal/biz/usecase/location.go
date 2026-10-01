package usecase

import (
	"bff_bbs/internal/biz/repo"
	bbsuserv1 "common/proto/gen/bff_bbs/v1/user"
	"context"
)

type LocationUsecase struct {
	locationClient repo.LocationClient
}

func NewLocationUsecase(
	locationClient repo.LocationClient,
) *LocationUsecase {
	return &LocationUsecase{
		locationClient: locationClient,
	}
}

func (u *LocationUsecase) GetCurrentLocation(ctx context.Context, userID int64) (*bbsuserv1.GetCurrentLocation_Resp_Location, error) {
	reply, err := u.locationClient.GetCurrentLocation(ctx, userID)
	if err != nil {
		return nil, err
	}
	var location *bbsuserv1.GetCurrentLocation_Resp_Location
	if row := reply; row != nil {
		location = &bbsuserv1.GetCurrentLocation_Resp_Location{
			UserId:   row.UserID,
			Country:  row.Country,
			Province: row.Province,
			City:     row.City,
		}
	}
	return location, nil
}
