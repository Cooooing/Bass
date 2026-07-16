package repo

import "context"

type LocationClient interface {
	GetCurrentLocation(ctx context.Context, req *GetCurrentLocationReq) (*GetCurrentLocationResponse, error)
	UpsertCurrentLocation(ctx context.Context, req *UpsertCurrentLocationReq) (*UpsertCurrentLocationResponse, error)
}

type Location struct {
	UserID   int64
	Country  *string
	Province *string
	City     *string
}

type GetCurrentLocationReq struct {
	UserID int64
}

type GetCurrentLocationResponse struct {
	Location *Location
}

type UpsertCurrentLocationReq struct {
	UserID   int64
	Country  *string
	Province *string
	City     *string
}

type UpsertCurrentLocationResponse struct {
	Location *Location
}
