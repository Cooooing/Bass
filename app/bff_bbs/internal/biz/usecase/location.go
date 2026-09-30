package usecase

import (
	"bff_bbs/internal/biz/repo"
	bbsuserv1 "common/proto/gen/bff_bbs/v1/user"
	"context"
	"strings"
)

type LocationUsecase struct {
	locationClient     repo.LocationClient
	ipResolutionClient repo.IPResolutionClient
}

func NewLocationUsecase(
	locationClient repo.LocationClient,
	ipResolutionClient repo.IPResolutionClient,
) *LocationUsecase {
	return &LocationUsecase{
		locationClient:     locationClient,
		ipResolutionClient: ipResolutionClient,
	}
}

// DetectCurrentLocation resolves the current request IP, then persists only a usable result.
// An empty IP or an unresolved private address leaves the existing location untouched.
func (u *LocationUsecase) DetectCurrentLocation(ctx context.Context, userID int64, ip string) (*bbsuserv1.DetectCurrentLocation_Resp_Location, error) {
	ip = strings.TrimSpace(ip)
	if ip == "" || u.ipResolutionClient == nil {
		return u.currentDetectedLocation(ctx, userID)
	}
	resolved, err := u.ipResolutionClient.Resolve(ctx, ip)
	if err != nil {
		return nil, err
	}
	if resolved == nil || (resolved.Country == "" && resolved.Province == "" && resolved.City == "") {
		return u.currentDetectedLocation(ctx, userID)
	}
	row, err := u.locationClient.UpsertCurrentLocation(ctx, &repo.UpsertCurrentLocationReq{
		UserID:   userID,
		Country:  stringPointer(resolved.Country),
		Province: stringPointer(resolved.Province),
		City:     stringPointer(resolved.City),
	})
	if err != nil {
		return nil, err
	}
	return detectedLocation(row), nil
}

func (u *LocationUsecase) currentDetectedLocation(ctx context.Context, userID int64) (*bbsuserv1.DetectCurrentLocation_Resp_Location, error) {
	row, err := u.locationClient.GetCurrentLocation(ctx, userID)
	if err != nil {
		return nil, err
	}
	return detectedLocation(row), nil
}

func detectedLocation(row *repo.Location) *bbsuserv1.DetectCurrentLocation_Resp_Location {
	if row == nil {
		return nil
	}
	return &bbsuserv1.DetectCurrentLocation_Resp_Location{UserId: row.UserID, Country: row.Country, Province: row.Province, City: row.City}
}

func stringPointer(value string) *string {
	if value == "" {
		return nil
	}
	return &value
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

type UpsertCurrentLocationReq struct {
	UserID   int64
	Country  *string
	Province *string
	City     *string
}

func (u *LocationUsecase) UpsertCurrentLocation(ctx context.Context, req *UpsertCurrentLocationReq) (*bbsuserv1.UpsertCurrentLocation_Resp_Location, error) {
	reply, err := u.locationClient.UpsertCurrentLocation(ctx, &repo.UpsertCurrentLocationReq{
		UserID:   req.UserID,
		Country:  req.Country,
		Province: req.Province,
		City:     req.City,
	})
	if err != nil {
		return nil, err
	}
	var location *bbsuserv1.UpsertCurrentLocation_Resp_Location
	if row := reply; row != nil {
		location = &bbsuserv1.UpsertCurrentLocation_Resp_Location{
			UserId:   row.UserID,
			Country:  row.Country,
			Province: row.Province,
			City:     row.City,
		}
	}
	return location, nil
}
