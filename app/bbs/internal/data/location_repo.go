package data

import (
	"bbs/internal/biz/repo"
	bbsuserv1 "common/api/gen/bbs/v1/user"
	userv1 "common/api/gen/user/v1"
	"common/pkg/client/rpc"
	"context"
)

var _ repo.LocationRepo = (*LocationRepo)(nil)

type LocationRepo struct {
	userClient *rpc.UserClient
}

func NewLocationRepo(userClient *rpc.UserClient) repo.LocationRepo {
	return &LocationRepo{userClient: userClient}
}

func (r *LocationRepo) GetCurrentLocation(ctx context.Context, req *bbsuserv1.GetCurrentLocation_Request) (*bbsuserv1.GetCurrentLocation_Reply, error) {
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

func (r *LocationRepo) UpsertCurrentLocation(ctx context.Context, req *bbsuserv1.UpsertCurrentLocation_Request) (*bbsuserv1.UpsertCurrentLocation_Reply, error) {
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
