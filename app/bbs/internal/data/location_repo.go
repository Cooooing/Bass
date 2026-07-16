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

func NewLocationClient(userClient *rpc.UserClient) repo.LocationClient {
	return &LocationClient{userClient: userClient}
}

func (r *LocationClient) GetCurrentLocation(ctx context.Context, req *repo.GetCurrentLocationReq) (*repo.GetCurrentLocationResponse, error) {
	reply, err := r.userClient.Location.Get(ctx, &userv1.GetLocation_Request{UserId: req.UserID})
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
	return &repo.GetCurrentLocationResponse{Location: out}, nil
}

func (r *LocationClient) UpsertCurrentLocation(ctx context.Context, req *repo.UpsertCurrentLocationReq) (*repo.UpsertCurrentLocationResponse, error) {
	reply, err := r.userClient.Location.Upsert(ctx, &userv1.UpsertLocation_Request{
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
	return &repo.UpsertCurrentLocationResponse{Location: out}, nil
}
