package service

import (
	"bbs/internal/biz/usecase"
	bbsuserv1 "common/proto/gen/bbs/v1/user"
	"context"

	"github.com/go-kratos/kratos/v3/transport/grpc"
	"github.com/go-kratos/kratos/v3/transport/http"
)

type PrivacySettingService struct {
	bbsuserv1.UnimplementedPrivacySettingServiceServer
	privacySettingUsecase *usecase.PrivacySettingUsecase
}

func NewPrivacySettingService(privacySettingUsecase *usecase.PrivacySettingUsecase) *PrivacySettingService {
	return &PrivacySettingService{privacySettingUsecase: privacySettingUsecase}
}

func (s *PrivacySettingService) RegisterGrpc(gs *grpc.Server) {}

func (s *PrivacySettingService) RegisterHttp(hs *http.Server) {
	bbsuserv1.RegisterPrivacySettingServiceHTTPServer(hs, s)
}

func (s *PrivacySettingService) GetCurrent(ctx context.Context, req *bbsuserv1.GetCurrentPrivacySetting_Req) (*bbsuserv1.GetCurrentPrivacySetting_Resp, error) {
	userID, err := currentUserID(ctx)
	if err != nil {
		return nil, err
	}
	setting, err := s.privacySettingUsecase.GetCurrentPrivacySetting(ctx, userID)
	if err != nil {
		return nil, err
	}
	return &bbsuserv1.GetCurrentPrivacySetting_Resp{PrivacySetting: setting}, nil
}

func (s *PrivacySettingService) UpdateCurrent(ctx context.Context, req *bbsuserv1.UpdateCurrentPrivacySetting_Req) (*bbsuserv1.UpdateCurrentPrivacySetting_Resp, error) {
	userID, err := currentUserID(ctx)
	if err != nil {
		return nil, err
	}
	setting, err := s.privacySettingUsecase.UpdateCurrentPrivacySetting(ctx, &usecase.UpdateCurrentPrivacySettingReq{UserID: userID, PublicPoints: req.PublicPoints, PublicFollowers: req.PublicFollowers, PublicArticles: req.PublicArticles, PublicComments: req.PublicComments, PublicOnlineStatus: req.PublicOnlineStatus, PublicLocation: req.PublicLocation})
	if err != nil {
		return nil, err
	}
	return &bbsuserv1.UpdateCurrentPrivacySetting_Resp{PrivacySetting: setting}, nil
}
