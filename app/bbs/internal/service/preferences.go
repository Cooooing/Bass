package service

import (
	"bbs/internal/biz/usecase"
	bbsuserv1 "common/proto/gen/bbs/v1/user"
	"context"

	"github.com/go-kratos/kratos/v3/transport/grpc"
	"github.com/go-kratos/kratos/v3/transport/http"
)

type PreferencesService struct {
	bbsuserv1.UnimplementedPreferencesServiceServer
	preferencesUsecase *usecase.PreferencesUsecase
}

func NewPreferencesService(
	preferencesUsecase *usecase.PreferencesUsecase,
) *PreferencesService {
	return &PreferencesService{
		preferencesUsecase: preferencesUsecase,
	}
}

func (s *PreferencesService) RegisterGrpc(gs *grpc.Server) {
}

func (s *PreferencesService) RegisterHttp(hs *http.Server) {
	bbsuserv1.RegisterPreferencesServiceHTTPServer(hs, s)
}

func (s *PreferencesService) GetCurrent(ctx context.Context, req *bbsuserv1.GetCurrentPreferences_Req) (*bbsuserv1.GetCurrentPreferences_Resp, error) {
	userID, err := currentUserID(ctx)
	if err != nil {
		return nil, err
	}
	preference, err := s.preferencesUsecase.GetCurrentPreferences(ctx, userID)
	if err != nil {
		return nil, err
	}
	return &bbsuserv1.GetCurrentPreferences_Resp{
		Preference: preference,
	}, nil
}

func (s *PreferencesService) UpdateCurrent(ctx context.Context, req *bbsuserv1.UpdateCurrentPreferences_Req) (*bbsuserv1.UpdateCurrentPreferences_Resp, error) {
	userID, err := currentUserID(ctx)
	if err != nil {
		return nil, err
	}
	preference, err := s.preferencesUsecase.UpdateCurrentPreferences(ctx, &usecase.UpdateCurrentPreferencesReq{
		UserID:      userID,
		Timezone:    req.Timezone,
		Theme:       req.Theme,
		MobileTheme: req.MobileTheme,
		Language:    req.Language,
	})
	if err != nil {
		return nil, err
	}
	return &bbsuserv1.UpdateCurrentPreferences_Resp{
		Preference: preference,
	}, nil
}
