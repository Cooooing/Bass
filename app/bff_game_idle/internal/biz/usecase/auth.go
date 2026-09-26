package usecase

import (
	"bff_game_idle/internal/biz/model"
	"bff_game_idle/internal/biz/repo"
	"context"
)

type AuthUsecase struct {
	authRepo repo.AuthRepo
}

func NewAuthUsecase(
	authRepo repo.AuthRepo,
) *AuthUsecase {
	return &AuthUsecase{
		authRepo: authRepo,
	}
}

type RegisterReq struct {
	Password string
	Email    string
}

func (u *AuthUsecase) Register(ctx context.Context, req *RegisterReq) error {
	return u.authRepo.Register(ctx, &repo.RegisterReq{
		Password: req.Password,
		Email:    req.Email,
	})
}

type LoginReq struct {
	Email    string
	Password string
}

func (u *AuthUsecase) Login(ctx context.Context, req *LoginReq) (*model.LoginToken, error) {
	return u.authRepo.Login(ctx, &repo.LoginReq{
		Email:    req.Email,
		Password: req.Password,
	})
}
