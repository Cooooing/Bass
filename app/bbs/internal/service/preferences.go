package service

import (
	"bbs/internal/biz/usecase"
	bbsuserv1 "common/api/gen/bbs/v1/user"
	"context"

	"github.com/go-kratos/kratos/v2/transport/grpc"
	"github.com/go-kratos/kratos/v2/transport/http"
)

type PreferencesService struct {
	bbsuserv1.UnimplementedPreferencesServiceServer
	userUsecase *usecase.UserUsecase
}

func NewPreferencesService(userUsecase *usecase.UserUsecase) *PreferencesService {
	return &PreferencesService{userUsecase: userUsecase}
}

func (s *PreferencesService) RegisterGrpc(gs *grpc.Server) {
	bbsuserv1.RegisterPreferencesServiceServer(gs, s)
}

func (s *PreferencesService) RegisterHttp(hs *http.Server) {
	bbsuserv1.RegisterPreferencesServiceHTTPServer(hs, s)
}

func (s *PreferencesService) GetCurrent(ctx context.Context, req *bbsuserv1.GetCurrentPreferences_Request) (*bbsuserv1.GetCurrentPreferences_Reply, error) {
	return s.userUsecase.GetCurrentPreferences(ctx, req)
}

func (s *PreferencesService) UpdateCurrent(ctx context.Context, req *bbsuserv1.UpdateCurrentPreferences_Request) (*bbsuserv1.UpdateCurrentPreferences_Reply, error) {
	return s.userUsecase.UpdateCurrentPreferences(ctx, req)
}
