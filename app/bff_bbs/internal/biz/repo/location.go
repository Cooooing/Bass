package repo

import "context"

type LocationClient interface {
	GetCurrentLocation(ctx context.Context, userID int64) (*Location, error)
}

type Location struct {
	UserID   int64
	Country  *string
	Province *string
	City     *string
}
