package service

import (
	bbsuserv1 "common/api/gen/bbs/v1/user"
	userv1 "common/api/gen/user/v1"
	"common/pkg/client/rpc"
	"context"

	"github.com/go-kratos/kratos/v2/transport/grpc"
	"github.com/go-kratos/kratos/v2/transport/http"
)

type LocationService struct {
	bbsuserv1.UnimplementedLocationServiceServer
	userClient *rpc.UserClient
}

func NewLocationService(userClient *rpc.UserClient) *LocationService {
	return &LocationService{userClient: userClient}
}

func (s *LocationService) RegisterGrpc(gs *grpc.Server) {
	bbsuserv1.RegisterLocationServiceServer(gs, s)
}

func (s *LocationService) RegisterHttp(hs *http.Server) {
	bbsuserv1.RegisterLocationServiceHTTPServer(hs, s)
}

func (s *LocationService) GetCurrent(ctx context.Context, req *bbsuserv1.GetCurrentLocation_Request) (*bbsuserv1.GetCurrentLocation_Reply, error) {
	reply, err := s.userClient.Location.GetCurrent(forwardAuth(ctx), &userv1.GetCurrentLocation_Request{})
	if err != nil {
		return nil, err
	}
	return &bbsuserv1.GetCurrentLocation_Reply{Location: toBFFLocation(reply.GetLocation())}, nil
}

func (s *LocationService) Upsert(ctx context.Context, req *bbsuserv1.UpsertLocation_Request) (*bbsuserv1.UpsertLocation_Reply, error) {
	reply, err := s.userClient.Location.Upsert(forwardAuth(ctx), &userv1.UpsertLocation_Request{
		Country:  req.Country,
		Province: req.Province,
		City:     req.City,
	})
	if err != nil {
		return nil, err
	}
	return &bbsuserv1.UpsertLocation_Reply{Location: toBFFLocation(reply.GetLocation())}, nil
}

func toBFFLocation(in *userv1.Location) *bbsuserv1.Location {
	if in == nil {
		return nil
	}
	return &bbsuserv1.Location{
		UserId:   in.GetUserId(),
		Country:  in.Country,
		Province: in.Province,
		City:     in.City,
	}
}
