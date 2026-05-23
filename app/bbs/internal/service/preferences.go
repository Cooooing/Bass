package service

import (
	bbsuserv1 "common/api/gen/bbs/v1/user"
	userv1 "common/api/gen/user/v1"
	"common/pkg/client/rpc"
	"context"

	"github.com/go-kratos/kratos/v2/transport/grpc"
	"github.com/go-kratos/kratos/v2/transport/http"
)

type PreferencesService struct {
	bbsuserv1.UnimplementedPreferencesServiceServer
	userClient *rpc.UserClient
}

func NewPreferencesService(userClient *rpc.UserClient) *PreferencesService {
	return &PreferencesService{userClient: userClient}
}

func (s *PreferencesService) RegisterGrpc(gs *grpc.Server) {
	bbsuserv1.RegisterPreferencesServiceServer(gs, s)
}

func (s *PreferencesService) RegisterHttp(hs *http.Server) {
	bbsuserv1.RegisterPreferencesServiceHTTPServer(hs, s)
}

func (s *PreferencesService) GetCurrent(ctx context.Context, req *bbsuserv1.GetCurrentPreferences_Request) (*bbsuserv1.GetCurrentPreferences_Reply, error) {
	reply, err := s.userClient.Preferences.GetCurrent(forwardAuth(ctx), &userv1.GetCurrentPreferences_Request{})
	if err != nil {
		return nil, err
	}
	return &bbsuserv1.GetCurrentPreferences_Reply{Preference: toBFFPreference(reply.GetPreferences())}, nil
}

func (s *PreferencesService) Update(ctx context.Context, req *bbsuserv1.UpdatePreferences_Request) (*bbsuserv1.UpdatePreferences_Reply, error) {
	reply, err := s.userClient.Preferences.Update(forwardAuth(ctx), &userv1.UpdatePreferences_Request{
		Language:             req.Language,
		Timezone:             req.Timezone,
		Theme:                req.Theme,
		MobileTheme:          req.MobileTheme,
		EnableWebNotify:      req.EnableWebNotify,
		EnableEmailSubscribe: req.EnableEmailSubscribe,
	})
	if err != nil {
		return nil, err
	}
	return &bbsuserv1.UpdatePreferences_Reply{Preference: toBFFPreference(reply.GetPreferences())}, nil
}

func toBFFPreference(in *userv1.Preferences) *bbsuserv1.Preference {
	if in == nil {
		return nil
	}
	return &bbsuserv1.Preference{
		UserId:               in.GetUserId(),
		Language:             in.Language,
		Timezone:             in.Timezone,
		Theme:                in.Theme,
		MobileTheme:          in.MobileTheme,
		EnableWebNotify:      in.EnableWebNotify,
		EnableEmailSubscribe: in.EnableEmailSubscribe,
	}
}
