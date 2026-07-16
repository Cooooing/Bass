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

func (s *PreferencesService) GetCurrent(ctx context.Context, req *bbsuserv1.GetCurrentPreferences_Request) (*bbsuserv1.GetCurrentPreferences_Response, error) {
	userID, err := currentUserID(ctx)
	if err != nil {
		return nil, err
	}
	response, err := s.preferencesUsecase.GetCurrentPreferences(ctx, &usecase.GetCurrentPreferencesReq{UserID: userID})
	if err != nil {
		return nil, err
	}
	return &bbsuserv1.GetCurrentPreferences_Response{Preference: response.Preference}, nil
}

func (s *PreferencesService) UpdateCurrent(ctx context.Context, req *bbsuserv1.UpdateCurrentPreferences_Request) (*bbsuserv1.UpdateCurrentPreferences_Response, error) {
	userID, err := currentUserID(ctx)
	if err != nil {
		return nil, err
	}
	response, err := s.preferencesUsecase.UpdateCurrentPreferences(ctx, &usecase.UpdateCurrentPreferencesReq{UserID: userID, Timezone: req.Timezone, Theme: req.Theme, MobileTheme: req.MobileTheme, Language: req.Language})
	if err != nil {
		return nil, err
	}
	return &bbsuserv1.UpdateCurrentPreferences_Response{Preference: response.Preference}, nil
}
