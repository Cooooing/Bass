package service

import (
	cerrors "common/api/gen/common/errors"
	v1 "common/api/gen/user/v1"
	"common/pkg/constant"
	commonModel "common/pkg/model"
	"common/pkg/util"
	"context"
	"user/internal/biz/model"
	"user/internal/biz/repo"

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
	preferences, _ := s.preferencesRepo.FindByUserID(ctx, current.ID)
	reply := &v1.Preferences{UserId: current.ID}
	if preferences != nil {
		reply.Language = preferences.Language
		reply.Timezone = preferences.Timezone
		reply.Theme = preferences.Theme
		reply.MobileTheme = preferences.MobileTheme
		reply.EnableWebNotify = preferences.EnableWebNotify
		reply.EnableEmailSubscribe = preferences.EnableEmailSubscribe
	}
	return &v1.GetCurrentPreferences_Reply{Preferences: reply}, nil
}

func (s *PreferencesService) Update(ctx context.Context, req *v1.UpdatePreferences_Request) (*v1.UpdatePreferences_Reply, error) {
	current, ok := util.GetContextValue[*commonModel.User](ctx, constant.CtxUserInfo)
	if !ok {
		return nil, cerrors.ErrorUnauthorized("user not login")
	}
	preferences, err := s.preferencesRepo.UpsertByUserID(ctx, &model.Preferences{
		UserID:               current.ID,
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
	return &v1.UpdatePreferences_Reply{Preferences: &v1.Preferences{
		UserId:               current.ID,
		Language:             preferences.Language,
		Timezone:             preferences.Timezone,
		Theme:                preferences.Theme,
		MobileTheme:          preferences.MobileTheme,
		EnableWebNotify:      preferences.EnableWebNotify,
		EnableEmailSubscribe: preferences.EnableEmailSubscribe,
	}}, nil
}
