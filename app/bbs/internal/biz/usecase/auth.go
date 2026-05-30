package usecase

import (
	"bbs/internal/biz/repo"
	bbsuserv1 "common/api/gen/bbs/v1/user"
	"context"
)

type AuthUsecase struct {
	authRepo repo.AuthRepo
}

func NewAuthUsecase(authRepo repo.AuthRepo) *AuthUsecase {
	return &AuthUsecase{authRepo: authRepo}
}

func (u *AuthUsecase) StartEmailRegistration(ctx context.Context, req *bbsuserv1.StartEmailRegistration_Request) (*bbsuserv1.StartEmailRegistration_Reply, error) {
	return u.authRepo.StartEmailRegistration(ctx, req)
}

func (u *AuthUsecase) VerifyEmailRegistration(ctx context.Context, req *bbsuserv1.VerifyEmailRegistration_Request) (*bbsuserv1.VerifyEmailRegistration_Reply, error) {
	return u.authRepo.VerifyEmailRegistration(ctx, req)
}

func (u *AuthUsecase) StartPhoneRegistration(ctx context.Context, req *bbsuserv1.StartPhoneRegistration_Request) (*bbsuserv1.StartPhoneRegistration_Reply, error) {
	return u.authRepo.StartPhoneRegistration(ctx, req)
}

func (u *AuthUsecase) VerifyPhoneRegistration(ctx context.Context, req *bbsuserv1.VerifyPhoneRegistration_Request) (*bbsuserv1.VerifyPhoneRegistration_Reply, error) {
	return u.authRepo.VerifyPhoneRegistration(ctx, req)
}

func (u *AuthUsecase) LoginByPassword(ctx context.Context, req *bbsuserv1.LoginByPassword_Request) (*bbsuserv1.LoginByPassword_Reply, error) {
	return u.authRepo.LoginByPassword(ctx, req)
}

func (u *AuthUsecase) Logout(ctx context.Context, req *bbsuserv1.Logout_Request) (*bbsuserv1.Logout_Reply, error) {
	return u.authRepo.Logout(ctx, req)
}
