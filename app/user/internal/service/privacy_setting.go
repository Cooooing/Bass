package service

import (
	"context"

	v1 "common/proto/gen/user/v1"
	"user/internal/biz/model"
	"user/internal/biz/usecase"

	"github.com/go-kratos/kratos/v3/transport/grpc"
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

func (s *PrivacySettingService) Get(ctx context.Context, req *v1.GetPrivacySetting_Request) (*v1.GetPrivacySetting_Response, error) {
	res, err := s.privacySettingUsecase.GetByUserID(ctx, &usecase.GetPrivacySettingByUserIDReq{UserID: req.GetUserId()})
	if err != nil {
		return nil, err
	}
	reply := &v1.GetPrivacySetting_Response_PrivacySetting{UserId: req.GetUserId()}
	if res.PrivacySetting != nil {
		reply.PublicPoints = res.PrivacySetting.PublicPoints
		reply.PublicFollowers = res.PrivacySetting.PublicFollowers
		reply.PublicArticles = res.PrivacySetting.PublicArticles
		reply.PublicComments = res.PrivacySetting.PublicComments
		reply.PublicOnlineStatus = res.PrivacySetting.PublicOnlineStatus
		reply.PublicLocation = res.PrivacySetting.PublicLocation
	}
	return &v1.GetPrivacySetting_Response{PrivacySetting: reply}, nil
}

func (s *PrivacySettingService) Update(ctx context.Context, req *v1.UpdatePrivacySetting_Request) (*v1.UpdatePrivacySetting_Response, error) {
	res, err := s.privacySettingUsecase.UpsertByUserID(ctx, &usecase.UpsertPrivacySettingByUserIDReq{PrivacySetting: &model.PrivacySetting{
		UserID:             req.GetUserId(),
		PublicPoints:       req.PublicPoints,
		PublicFollowers:    req.PublicFollowers,
		PublicArticles:     req.PublicArticles,
		PublicComments:     req.PublicComments,
		PublicOnlineStatus: req.PublicOnlineStatus,
		PublicLocation:     req.PublicLocation,
	}})
	if err != nil {
		return nil, err
	}
	privacySetting := res.PrivacySetting
	return &v1.UpdatePrivacySetting_Response{PrivacySetting: &v1.UpdatePrivacySetting_Response_PrivacySetting{
		UserId:             req.GetUserId(),
		PublicPoints:       privacySetting.PublicPoints,
		PublicFollowers:    privacySetting.PublicFollowers,
		PublicArticles:     privacySetting.PublicArticles,
		PublicComments:     privacySetting.PublicComments,
		PublicOnlineStatus: privacySetting.PublicOnlineStatus,
		PublicLocation:     privacySetting.PublicLocation,
	}}, nil
}
