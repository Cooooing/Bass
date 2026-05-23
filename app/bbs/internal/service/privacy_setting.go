package service

import (
	bbsuserv1 "common/api/gen/bbs/v1/user"
	userv1 "common/api/gen/user/v1"
	"common/pkg/client/rpc"
	"context"

	"github.com/go-kratos/kratos/v2/transport/grpc"
	"github.com/go-kratos/kratos/v2/transport/http"
)

type PrivacySettingService struct {
	bbsuserv1.UnimplementedPrivacySettingServiceServer
	userClient *rpc.UserClient
}

func NewPrivacySettingService(userClient *rpc.UserClient) *PrivacySettingService {
	return &PrivacySettingService{userClient: userClient}
}

func (s *PrivacySettingService) RegisterGrpc(gs *grpc.Server) {
	bbsuserv1.RegisterPrivacySettingServiceServer(gs, s)
}

func (s *PrivacySettingService) RegisterHttp(hs *http.Server) {
	bbsuserv1.RegisterPrivacySettingServiceHTTPServer(hs, s)
}

func (s *PrivacySettingService) GetCurrent(ctx context.Context, req *bbsuserv1.GetCurrentPrivacySetting_Request) (*bbsuserv1.GetCurrentPrivacySetting_Reply, error) {
	reply, err := s.userClient.PrivacySetting.GetCurrent(forwardAuth(ctx), &userv1.GetCurrentPrivacySetting_Request{})
	if err != nil {
		return nil, err
	}
	return &bbsuserv1.GetCurrentPrivacySetting_Reply{PrivacySetting: toBFFPrivacySetting(reply.GetPrivacySetting())}, nil
}

func (s *PrivacySettingService) Update(ctx context.Context, req *bbsuserv1.UpdatePrivacySetting_Request) (*bbsuserv1.UpdatePrivacySetting_Reply, error) {
	reply, err := s.userClient.PrivacySetting.Update(forwardAuth(ctx), &userv1.UpdatePrivacySetting_Request{
		PublicPoints:       req.PublicPoints,
		PublicFollowers:    req.PublicFollowers,
		PublicArticles:     req.PublicArticles,
		PublicComments:     req.PublicComments,
		PublicOnlineStatus: req.PublicOnlineStatus,
		PublicLocation:     req.PublicLocation,
	})
	if err != nil {
		return nil, err
	}
	return &bbsuserv1.UpdatePrivacySetting_Reply{PrivacySetting: toBFFPrivacySetting(reply.GetPrivacySetting())}, nil
}

func toBFFPrivacySetting(in *userv1.PrivacySetting) *bbsuserv1.PrivacySetting {
	if in == nil {
		return nil
	}
	return &bbsuserv1.PrivacySetting{
		UserId:             in.GetUserId(),
		PublicPoints:       in.PublicPoints,
		PublicFollowers:    in.PublicFollowers,
		PublicArticles:     in.PublicArticles,
		PublicComments:     in.PublicComments,
		PublicOnlineStatus: in.PublicOnlineStatus,
		PublicLocation:     in.PublicLocation,
	}
}
