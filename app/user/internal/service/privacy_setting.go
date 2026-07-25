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

func NewPrivacySettingService(
	privacySettingUsecase *usecase.PrivacySettingUsecase,
) *PrivacySettingService {
	return &PrivacySettingService{
		privacySettingUsecase: privacySettingUsecase,
	}
}

func (s *PrivacySettingService) RegisterGrpc(gs *grpc.Server) {
	v1.RegisterPrivacySettingServiceServer(gs, s)
}

func (s *PrivacySettingService) RegisterHttp(hs *http.Server) {
}

func (s *PrivacySettingService) Get(ctx context.Context, req *v1.GetPrivacySetting_Req) (*v1.GetPrivacySetting_Resp, error) {
	res, err := s.privacySettingUsecase.GetByUserID(ctx, req.GetUserId())
	if err != nil {
		return nil, err
	}
	reply := &v1.GetPrivacySetting_Resp_PrivacySetting{
		UserId: req.GetUserId(),
	}
	if res != nil {
		reply.PublicPoints = res.PublicPoints
		reply.PublicFollowers = res.PublicFollowers
		reply.PublicArticles = res.PublicArticles
		reply.PublicComments = res.PublicComments
		reply.PublicOnlineStatus = res.PublicOnlineStatus
		reply.PublicLocation = res.PublicLocation
	}
	return &v1.GetPrivacySetting_Resp{
		PrivacySetting: reply,
	}, nil
}

func (s *PrivacySettingService) Update(ctx context.Context, req *v1.UpdatePrivacySetting_Req) (*v1.UpdatePrivacySetting_Resp, error) {
	res, err := s.privacySettingUsecase.UpsertByUserID(ctx, &model.PrivacySetting{
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
	privacySetting := res
	return &v1.UpdatePrivacySetting_Resp{
		PrivacySetting: &v1.UpdatePrivacySetting_Resp_PrivacySetting{
			UserId:             req.GetUserId(),
			PublicPoints:       privacySetting.PublicPoints,
			PublicFollowers:    privacySetting.PublicFollowers,
			PublicArticles:     privacySetting.PublicArticles,
			PublicComments:     privacySetting.PublicComments,
			PublicOnlineStatus: privacySetting.PublicOnlineStatus,
			PublicLocation:     privacySetting.PublicLocation,
		},
	}, nil
}
