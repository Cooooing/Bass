package service

import (
	commonenums "common/api/gen/common/enums"
	cerrors "common/api/gen/common/errors"
	v1 "common/api/gen/user/v1"
	"context"
	"user/internal/biz/model"
	"user/internal/biz/usecase"
	"user/internal/enum"

	"github.com/go-kratos/kratos/v2/transport/grpc"
	"github.com/go-kratos/kratos/v2/transport/http"
)

type PreferencesService struct {
	v1.UnimplementedPreferencesServiceServer
	preferencesUsecase *usecase.PreferencesUsecase
}

func NewPreferencesService(preferencesUsecase *usecase.PreferencesUsecase) *PreferencesService {
	return &PreferencesService{preferencesUsecase: preferencesUsecase}
}

func (s *PreferencesService) RegisterGrpc(gs *grpc.Server) {
	v1.RegisterPreferencesServiceServer(gs, s)
}

func (s *PreferencesService) RegisterHttp(hs *http.Server) {}

func (s *PreferencesService) Get(ctx context.Context, req *v1.GetPreferences_Request) (*v1.GetPreferences_Reply, error) {
	preferences, err := s.preferencesUsecase.GetByUserID(ctx, req.GetUserId())
	if err != nil {
		return nil, err
	}
	reply := &v1.Preferences{UserId: req.GetUserId()}
	if preferences != nil {
		if preferences.Language != nil {
			reply.Language = enum.LanguageMap.MustToProto(*preferences.Language)
		}
		reply.Timezone = preferences.Timezone
		reply.Theme = preferences.Theme
		reply.MobileTheme = preferences.MobileTheme
	}
	return &v1.GetPreferences_Reply{Preferences: reply}, nil
}

func (s *PreferencesService) Update(ctx context.Context, req *v1.UpdatePreferences_Request) (*v1.UpdatePreferences_Reply, error) {
	var language *enum.Language
	if req.Language != nil {
		if *req.Language != commonenums.Language_LANGUAGE_UNSPECIFIED {
			value, ok := enum.LanguageMap.ToEnum(*req.Language)
			if !ok {
				return nil, cerrors.ErrorBadRequest("language is invalid")
			}
			language = new(value)
		}
	}
	preferences, err := s.preferencesUsecase.UpsertByUserID(ctx, &model.Preferences{
		UserID:      req.GetUserId(),
		Language:    language,
		Timezone:    req.Timezone,
		Theme:       req.Theme,
		MobileTheme: req.MobileTheme,
	})
	if err != nil {
		return nil, err
	}
	reply := &v1.Preferences{
		UserId:      req.GetUserId(),
		Timezone:    preferences.Timezone,
		Theme:       preferences.Theme,
		MobileTheme: preferences.MobileTheme,
	}
	if preferences.Language != nil {
		reply.Language = enum.LanguageMap.MustToProto(*preferences.Language)
	}
	return &v1.UpdatePreferences_Reply{Preferences: reply}, nil
}
