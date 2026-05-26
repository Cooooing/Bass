package service

import (
	commonenums "common/api/gen/common/enums"
	cerrors "common/api/gen/common/errors"
	v1 "common/api/gen/user/v1"
	"common/pkg/constant"
	commonModel "common/pkg/model"
	"common/pkg/util"
	"context"
	"user/internal/biz/model"
	"user/internal/biz/repo"
	"user/internal/enum"

	"github.com/go-kratos/kratos/v2/transport/grpc"
	"github.com/go-kratos/kratos/v2/transport/http"
)

type PreferencesService struct {
	v1.UnimplementedPreferencesServiceServer
	preferencesRepo repo.PreferencesRepo
}

func NewPreferencesService(preferencesRepo repo.PreferencesRepo) *PreferencesService {
	return &PreferencesService{preferencesRepo: preferencesRepo}
}

func (s *PreferencesService) RegisterGrpc(gs *grpc.Server) {
	v1.RegisterPreferencesServiceServer(gs, s)
}

func (s *PreferencesService) RegisterHttp(hs *http.Server) {}

func (s *PreferencesService) GetCurrent(ctx context.Context, req *v1.GetCurrentPreferences_Request) (*v1.GetCurrentPreferences_Reply, error) {
	current, ok := util.GetContextValue[*commonModel.User](ctx, constant.CtxUserInfo)
	if !ok {
		return nil, cerrors.ErrorUnauthorized("user not login")
	}
	preferences, err := s.preferencesRepo.FindByUserID(ctx, current.ID)
	if err != nil {
		return nil, err
	}
	reply := &v1.Preferences{UserId: current.ID}
	if preferences != nil {
		if preferences.Language != nil {
			reply.Language = enum.LanguageMap.MustToProto(*preferences.Language)
		}
		reply.Timezone = preferences.Timezone
		reply.Theme = preferences.Theme
		reply.MobileTheme = preferences.MobileTheme
	}
	return &v1.GetCurrentPreferences_Reply{Preferences: reply}, nil
}

func (s *PreferencesService) UpdateCurrent(ctx context.Context, req *v1.UpdateCurrentPreferences_Request) (*v1.UpdateCurrentPreferences_Reply, error) {
	current, ok := util.GetContextValue[*commonModel.User](ctx, constant.CtxUserInfo)
	if !ok {
		return nil, cerrors.ErrorUnauthorized("user not login")
	}
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
	preferences, err := s.preferencesRepo.UpsertByUserID(ctx, &model.Preferences{
		UserID:      current.ID,
		Language:    language,
		Timezone:    req.Timezone,
		Theme:       req.Theme,
		MobileTheme: req.MobileTheme,
	})
	if err != nil {
		return nil, err
	}
	reply := &v1.Preferences{
		UserId:      current.ID,
		Timezone:    preferences.Timezone,
		Theme:       preferences.Theme,
		MobileTheme: preferences.MobileTheme,
	}
	if preferences.Language != nil {
		reply.Language = enum.LanguageMap.MustToProto(*preferences.Language)
	}
	return &v1.UpdateCurrentPreferences_Reply{Preferences: reply}, nil
}
