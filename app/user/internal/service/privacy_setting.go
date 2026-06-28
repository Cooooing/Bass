package service

import (
	"context"

	v1 "common/proto/gen/user/v1"
	"user/internal/biz/model"
	"user/internal/biz/usecase"

	"github.com/go-kratos/kratos/v3/transport/grpc"
	"github.com/go-kratos/kratos/v3/transport/http"
)

type PrivacySettingService struct {
	v1.UnimplementedPrivacySettingServiceServer
	privacySettingUsecase *usecase.PrivacySettingUsecase
}

func NewPrivacySettingService(privacySettingUsecase *usecase.PrivacySettingUsecase) *PrivacySettingService {
	return &PrivacySettingService{privacySettingUsecase: privacySettingUsecase}
}

func (s *PrivacySettingService) RegisterGrpc(gs *grpc.Server) {
	v1.RegisterPrivacySettingServiceServer(gs, s)
}

func (s *PrivacySettingService) RegisterHttp(hs *http.Server) {}

func (s *PrivacySettingService) Get(ctx context.Context, req *v1.GetPrivacySetting_Request) (*v1.GetPrivacySetting_Reply, error) {
	privacySetting, err := s.privacySettingUsecase.GetByUserID(ctx, req.GetUserId())
	if err != nil {
		return nil, err
	}
	reply := &v1.PrivacySetting{UserId: req.GetUserId()}
	if privacySetting != nil {
		reply.PublicPoints = privacySetting.PublicPoints
		reply.PublicFollowers = privacySetting.PublicFollowers
		reply.PublicArticles = privacySetting.PublicArticles
		reply.PublicComments = privacySetting.PublicComments
		reply.PublicOnlineStatus = privacySetting.PublicOnlineStatus
		reply.PublicLocation = privacySetting.PublicLocation
	}
	return &v1.GetPrivacySetting_Reply{PrivacySetting: reply}, nil
}

func (s *PrivacySettingService) Update(ctx context.Context, req *v1.UpdatePrivacySetting_Request) (*v1.UpdatePrivacySetting_Reply, error) {
	privacySetting, err := s.privacySettingUsecase.UpsertByUserID(ctx, &model.PrivacySetting{
		UserID:             req.GetUserId(),
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
	return &v1.UpdatePrivacySetting_Reply{PrivacySetting: &v1.PrivacySetting{
		UserId:             req.GetUserId(),
		PublicPoints:       privacySetting.PublicPoints,
		PublicFollowers:    privacySetting.PublicFollowers,
		PublicArticles:     privacySetting.PublicArticles,
		PublicComments:     privacySetting.PublicComments,
		PublicOnlineStatus: privacySetting.PublicOnlineStatus,
		PublicLocation:     privacySetting.PublicLocation,
	}}, nil
}
