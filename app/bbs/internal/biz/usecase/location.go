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

type GetCurrentLocationReq struct {
	UserID int64
}

type GetCurrentLocationResponse struct {
	Location *bbsuserv1.GetCurrentLocation_Response_Location
}

func (u *LocationUsecase) GetCurrentLocation(ctx context.Context, req *GetCurrentLocationReq) (*GetCurrentLocationResponse, error) {
	reply, err := u.locationClient.GetCurrentLocation(ctx, &repo.GetCurrentLocationReq{UserID: req.UserID})
	if err != nil {
		return nil, err
	}
	var location *bbsuserv1.GetCurrentLocation_Response_Location
	if row := reply.Location; row != nil {
		location = &bbsuserv1.GetCurrentLocation_Response_Location{
			UserId:   row.UserID,
			Country:  row.Country,
			Province: row.Province,
			City:     row.City,
		}
	}
	return &GetCurrentLocationResponse{Location: location}, nil
}

type UpsertCurrentLocationReq struct {
	UserID   int64
	Country  *string
	Province *string
	City     *string
}

type UpsertCurrentLocationResponse struct {
	Location *bbsuserv1.UpsertCurrentLocation_Response_Location
}

func (u *LocationUsecase) UpsertCurrentLocation(ctx context.Context, req *UpsertCurrentLocationReq) (*UpsertCurrentLocationResponse, error) {
	reply, err := u.locationClient.UpsertCurrentLocation(ctx, &repo.UpsertCurrentLocationReq{
		UserID:   req.UserID,
		Country:  req.Country,
		Province: req.Province,
		City:     req.City,
	})
	if err != nil {
		return nil, err
	}
	var location *bbsuserv1.UpsertCurrentLocation_Response_Location
	if row := reply.Location; row != nil {
		location = &bbsuserv1.UpsertCurrentLocation_Response_Location{
			UserId:   row.UserID,
			Country:  row.Country,
			Province: row.Province,
			City:     row.City,
		}
	}
	return &UpsertCurrentLocationResponse{Location: location}, nil
}
