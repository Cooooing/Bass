package data

import (
	"bbs/internal/biz/repo"
	"common/pkg/client/rpc"
	bbsuserv1 "common/proto/gen/bbs/v1/user"
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

func (r *LocationClient) GetCurrentLocation(ctx context.Context, req *bbsuserv1.GetCurrentLocation_Request) (*bbsuserv1.GetCurrentLocation_Reply, error) {
	userID, err := currentUserID(ctx)
	if err != nil {
		return nil, err
	}
	reply, err := r.userClient.Location.Get(ctx, &userv1.GetLocation_Request{UserId: userID})
	if err != nil {
		return nil, err
	}
	location := reply.GetLocation()
	var out *bbsuserv1.Location
	if location != nil {
		out = &bbsuserv1.Location{
			UserId:   location.GetUserId(),
			Country:  location.Country,
			Province: location.Province,
			City:     location.City,
		}
	}
	return &bbsuserv1.GetCurrentLocation_Reply{Location: out}, nil
}

func (r *LocationClient) UpsertCurrentLocation(ctx context.Context, req *bbsuserv1.UpsertCurrentLocation_Request) (*bbsuserv1.UpsertCurrentLocation_Reply, error) {
	userID, err := currentUserID(ctx)
	if err != nil {
		return nil, err
	}
	reply, err := r.userClient.Location.Upsert(ctx, &userv1.UpsertLocation_Request{
		UserId:   userID,
		Country:  req.Country,
		Province: req.Province,
		City:     req.City,
	})
	if err != nil {
		return nil, err
	}
	location := reply.GetLocation()
	var out *bbsuserv1.Location
	if location != nil {
		out = &bbsuserv1.Location{
			UserId:   location.GetUserId(),
			Country:  location.Country,
			Province: location.Province,
			City:     location.City,
		}
	}
	return &bbsuserv1.UpsertCurrentLocation_Reply{Location: out}, nil
}
