package service

import (
	"bbs/internal/biz/usecase"
	bbsuserv1 "common/api/gen/bbs/v1/user"
	"context"

	"github.com/go-kratos/kratos/v2/transport/grpc"
	"github.com/go-kratos/kratos/v2/transport/http"
)

type AccountService struct {
	bbsuserv1.UnimplementedAccountServiceServer
	userUsecase *usecase.UserUsecase
}

func NewAccountService(userUsecase *usecase.UserUsecase) *AccountService {
	return &AccountService{userUsecase: userUsecase}
}

func (s *AccountService) RegisterGrpc(gs *grpc.Server) {
	bbsuserv1.RegisterAccountServiceServer(gs, s)
}

func (s *AccountService) RegisterHttp(hs *http.Server) {
	bbsuserv1.RegisterAccountServiceHTTPServer(hs, s)
}

func (s *AccountService) GetCurrent(ctx context.Context, req *bbsuserv1.GetCurrentAccount_Request) (*bbsuserv1.GetCurrentAccount_Reply, error) {
	return s.userUsecase.GetCurrentAccount(ctx, req)
}

func (s *AccountService) GetProfile(ctx context.Context, req *bbsuserv1.GetProfileAccount_Request) (*bbsuserv1.GetProfileAccount_Reply, error) {
	return s.userUsecase.GetProfileAccount(ctx, req)
}

func (s *AccountService) UpdateProfile(ctx context.Context, req *bbsuserv1.UpdateProfileAccount_Request) (*bbsuserv1.UpdateProfileAccount_Reply, error) {
	return s.userUsecase.UpdateProfileAccount(ctx, req)
}
