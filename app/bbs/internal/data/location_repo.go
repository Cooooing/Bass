package data

import (
	"bbs/internal/biz/repo"
	"common/pkg/client/rpc"
	userv1 "common/proto/gen/user/v1"
	"context"
)

var _ repo.LocationClient = (*LocationClient)(nil)

type LocationClient struct {
	userClient *rpc.UserClient
}

func NewLocationClient(
	userClient *rpc.UserClient,
) repo.LocationClient {
	return &LocationClient{
		userClient: userClient,
	}
}

func (r *LocationClient) GetCurrentLocation(ctx context.Context, userID int64) (*repo.Location, error) {
	reply, err := r.userClient.Location.Get(ctx, &userv1.GetLocation_Req{
		UserId: userID,
	})
	if err != nil {
		return nil, err
	}
	location := reply.GetLocation()
	var out *repo.Location
	if location != nil {
		out = &repo.Location{
			UserID:   location.GetUserId(),
			Country:  location.Country,
			Province: location.Province,
			City:     location.City,
		}
	}
	return out, nil
}

func (r *LocationClient) UpsertCurrentLocation(ctx context.Context, req *repo.UpsertCurrentLocationReq) (*repo.Location, error) {
	reply, err := r.userClient.Location.Upsert(ctx, &userv1.UpsertLocation_Req{
		UserId:   req.UserID,
		Country:  req.Country,
		Province: req.Province,
		City:     req.City,
	})
	if err != nil {
		return nil, err
	}
	location := reply.GetLocation()
	var out *repo.Location
	if location != nil {
		out = &repo.Location{
			UserID:   location.GetUserId(),
			Country:  location.Country,
			Province: location.Province,
			City:     location.City,
		}
	}
	return out, nil
}
