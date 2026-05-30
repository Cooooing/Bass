package service

import (
	"bbs/internal/biz/usecase"
	bbsuserv1 "common/api/gen/bbs/v1/user"
	"context"

	"github.com/go-kratos/kratos/v2/transport/grpc"
	"github.com/go-kratos/kratos/v2/transport/http"
)

type PrivacySettingService struct {
	bbsuserv1.UnimplementedPrivacySettingServiceServer
	privacySettingUsecase *usecase.PrivacySettingUsecase
}

func NewPrivacySettingService(privacySettingUsecase *usecase.PrivacySettingUsecase) *PrivacySettingService {
	return &PrivacySettingService{privacySettingUsecase: privacySettingUsecase}
}

func (s *PrivacySettingService) RegisterGrpc(gs *grpc.Server) {
	bbsuserv1.RegisterPrivacySettingServiceServer(gs, s)
}

func (s *PrivacySettingService) RegisterHttp(hs *http.Server) {
	bbsuserv1.RegisterPrivacySettingServiceHTTPServer(hs, s)
}

func (s *PrivacySettingService) GetCurrent(ctx context.Context, req *bbsuserv1.GetCurrentPrivacySetting_Request) (*bbsuserv1.GetCurrentPrivacySetting_Reply, error) {
	return s.privacySettingUsecase.GetCurrentPrivacySetting(ctx, req)
}

func (s *PrivacySettingService) UpdateCurrent(ctx context.Context, req *bbsuserv1.UpdateCurrentPrivacySetting_Request) (*bbsuserv1.UpdateCurrentPrivacySetting_Reply, error) {
	return s.privacySettingUsecase.UpdateCurrentPrivacySetting(ctx, req)
}
