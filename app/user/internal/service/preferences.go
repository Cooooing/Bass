package service

import (
	"common/pkg/apperror"
	commonenums "common/proto/gen/common/enums"
	cerrors "common/proto/gen/common/errors"
	v1 "common/proto/gen/user/v1"
	"context"
	"user/internal/biz/model"
	"user/internal/biz/usecase"
	"user/internal/enum"

	"github.com/go-kratos/kratos/v3/transport/grpc"
	"github.com/go-kratos/kratos/v3/transport/http"
)

type PreferencesService struct {
	v1.UnimplementedPreferencesServiceServer
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
	v1.RegisterPreferencesServiceServer(gs, s)
}

func (s *PreferencesService) RegisterHttp(hs *http.Server) {
}

func (s *PreferencesService) Get(ctx context.Context, req *v1.GetPreferences_Req) (*v1.GetPreferences_Resp, error) {
	res, err := s.preferencesUsecase.GetByUserID(ctx, req.GetUserId())
	if err != nil {
		return nil, err
	}
	reply := &v1.GetPreferences_Resp_Preferences{
		UserId: req.GetUserId(),
	}
	if res != nil {
		if res.Language != nil {
			reply.Language = enum.LanguageMap.MustToProto(*res.Language)
		}
		reply.Timezone = res.Timezone
		reply.Theme = res.Theme
		reply.MobileTheme = res.MobileTheme
	}
	return &v1.GetPreferences_Resp{
		Preferences: reply,
	}, nil
}

func (s *PreferencesService) Update(ctx context.Context, req *v1.UpdatePreferences_Req) (*v1.UpdatePreferences_Resp, error) {
	var language *enum.Language
	if req.Language != nil {
		if *req.Language != commonenums.Language_LANGUAGE_UNSPECIFIED {
			value, ok := enum.LanguageMap.ToEnum(*req.Language)
			if !ok {
				return nil, apperror.New(cerrors.BusinessErrorCode_BUSINESS_ERROR_CODE_USER_PREFERENCE_INVALID)
			}
			language = new(value)
		}
	}
	res, err := s.preferencesUsecase.UpsertByUserID(ctx, &model.Preferences{
		UserID:      req.GetUserId(),
		Language:    language,
		Timezone:    req.Timezone,
		Theme:       req.Theme,
		MobileTheme: req.MobileTheme,
	})
	if err != nil {
		return nil, err
	}
	preferences := res
	reply := &v1.UpdatePreferences_Resp_Preferences{
		UserId:      req.GetUserId(),
		Timezone:    preferences.Timezone,
		Theme:       preferences.Theme,
		MobileTheme: preferences.MobileTheme,
	}
	if preferences.Language != nil {
		reply.Language = enum.LanguageMap.MustToProto(*preferences.Language)
	}
	return &v1.UpdatePreferences_Resp{
		Preferences: reply,
	}, nil
}
