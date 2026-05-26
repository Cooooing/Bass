package service

import (
	"context"

	cerrors "common/api/gen/common/errors"
	v1 "common/api/gen/user/v1"
	"common/pkg/constant"
	commonModel "common/pkg/model"
	"common/pkg/util"
	"user/internal/biz/model"
	"user/internal/biz/repo"

	"github.com/go-kratos/kratos/v2/transport/grpc"
	"github.com/go-kratos/kratos/v2/transport/http"
)

type PrivacySettingService struct {
	v1.UnimplementedPrivacySettingServiceServer
	privacySettingRepo repo.PrivacySettingRepo
}

func NewPrivacySettingService(privacySettingRepo repo.PrivacySettingRepo) *PrivacySettingService {
	return &PrivacySettingService{privacySettingRepo: privacySettingRepo}
}

func (s *PrivacySettingService) RegisterGrpc(gs *grpc.Server) {
	v1.RegisterPrivacySettingServiceServer(gs, s)
}

func (s *PrivacySettingService) RegisterHttp(hs *http.Server) {}

func (s *PrivacySettingService) GetCurrent(ctx context.Context, req *v1.GetCurrentPrivacySetting_Request) (*v1.GetCurrentPrivacySetting_Reply, error) {
	current, ok := util.GetContextValue[*commonModel.User](ctx, constant.CtxUserInfo)
	if !ok {
		return nil, cerrors.ErrorUnauthorized("user not login")
	}
	privacySetting, err := s.privacySettingRepo.FindByUserID(ctx, current.ID)
	if err != nil {
		return nil, err
	}
	reply := &v1.PrivacySetting{UserId: current.ID}
	if privacySetting != nil {
		reply.PublicPoints = privacySetting.PublicPoints
		reply.PublicFollowers = privacySetting.PublicFollowers
		reply.PublicArticles = privacySetting.PublicArticles
		reply.PublicComments = privacySetting.PublicComments
		reply.PublicOnlineStatus = privacySetting.PublicOnlineStatus
		reply.PublicLocation = privacySetting.PublicLocation
	}
	return &v1.GetCurrentPrivacySetting_Reply{PrivacySetting: reply}, nil
}

func (s *PrivacySettingService) UpdateCurrent(ctx context.Context, req *v1.UpdateCurrentPrivacySetting_Request) (*v1.UpdateCurrentPrivacySetting_Reply, error) {
	current, ok := util.GetContextValue[*commonModel.User](ctx, constant.CtxUserInfo)
	if !ok {
		return nil, cerrors.ErrorUnauthorized("user not login")
	}
	privacySetting, err := s.privacySettingRepo.UpsertByUserID(ctx, &model.PrivacySetting{
		UserID:             current.ID,
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
	return &v1.UpdateCurrentPrivacySetting_Reply{PrivacySetting: &v1.PrivacySetting{
		UserId:             current.ID,
		PublicPoints:       privacySetting.PublicPoints,
		PublicFollowers:    privacySetting.PublicFollowers,
		PublicArticles:     privacySetting.PublicArticles,
		PublicComments:     privacySetting.PublicComments,
		PublicOnlineStatus: privacySetting.PublicOnlineStatus,
		PublicLocation:     privacySetting.PublicLocation,
	}}, nil
}
