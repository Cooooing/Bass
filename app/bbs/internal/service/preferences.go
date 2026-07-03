package service

import (
	"bbs/internal/biz/usecase"
	bbsuserv1 "common/proto/gen/bbs/v1/user"
	"context"

	"github.com/go-kratos/kratos/v3/transport/http"
)

type PreferencesService struct {
	bbsuserv1.UnimplementedPreferencesServiceServer
	preferencesUsecase *usecase.PreferencesUsecase
}

func NewPreferencesService(preferencesUsecase *usecase.PreferencesUsecase) *PreferencesService {
	return &PreferencesService{preferencesUsecase: preferencesUsecase}
}

func (s *PreferencesService) RegisterHttp(hs *http.Server) {
	bbsuserv1.RegisterPreferencesServiceHTTPServer(hs, s)
}

func (s *PreferencesService) GetCurrent(ctx context.Context, req *bbsuserv1.GetCurrentPreferences_Request) (*bbsuserv1.GetCurrentPreferences_Reply, error) {
	return s.preferencesUsecase.GetCurrentPreferences(ctx, req)
}

func (s *PreferencesService) UpdateCurrent(ctx context.Context, req *bbsuserv1.UpdateCurrentPreferences_Request) (*bbsuserv1.UpdateCurrentPreferences_Reply, error) {
	return s.preferencesUsecase.UpdateCurrentPreferences(ctx, req)
}
