package usecase

import (
	"bbs/internal/biz/model"
	"bbs/internal/biz/repo"
	"bbs/internal/enum"
	"context"
	"time"
)

type AuthUsecase struct{ authRepo repo.AuthRepo }

func NewAuthUsecase(
	authRepo repo.AuthRepo,
) *AuthUsecase {
	return &AuthUsecase{
		authRepo: authRepo,
	}
}

type RegisterReq struct {
	Type     enum.RegisterType
	Name     string
	Password string
	Nickname *string
	Email    string
	Phone    string
	Code     string
}

func (u *AuthUsecase) Register(ctx context.Context, req *RegisterReq) error {
	return u.authRepo.Register(ctx, &repo.RegisterReq{
		Type:     req.Type,
		Name:     req.Name,
		Password: req.Password,
		Nickname: req.Nickname,
		Email:    req.Email,
		Phone:    req.Phone,
		Code:     req.Code,
	})
}

type LoginReq struct {
	Type     enum.LoginType
	Account  string
	Password string
	Email    string
	Phone    string
	Code     string
}
type LoginResp struct {
	AccessToken           string
	RefreshToken          string
	AccessTokenExpiresAt  *time.Time
	RefreshTokenExpiresAt *time.Time
	SessionExpiresAt      *time.Time
	Account               *model.Account
}

func (u *AuthUsecase) Login(ctx context.Context, req *LoginReq) (*LoginResp, error) {
	reply, err := u.authRepo.Login(ctx, &repo.LoginReq{
		Type:     req.Type,
		Account:  req.Account,
		Password: req.Password,
		Email:    req.Email,
		Phone:    req.Phone,
		Code:     req.Code,
	})
	if err != nil {
		return nil, err
	}
	var account *model.Account
	if reply.Account != nil {
		account = &model.Account{}
		if profile := reply.Account.Profile; profile != nil {
			account.Profile = &model.AccountProfile{
				ID:            profile.ID,
				Name:          profile.Name,
				Nickname:      profile.Nickname,
				URL:           profile.URL,
				AvatarURL:     profile.AvatarURL,
				Introduction:  profile.Introduction,
				Status:        profile.Status,
				MBTI:          profile.MBTI,
				FollowCount:   profile.FollowCount,
				FollowerCount: profile.FollowerCount,
				CreatedAt:     profile.CreatedAt,
				UpdatedAt:     profile.UpdatedAt,
			}
		}
		if contact := reply.Account.Contact; contact != nil {
			account.Contact = &model.AccountContact{
				UserID: contact.UserID,
				Email:  contact.Email,
				Phone:  contact.Phone,
			}
		}
	}
	return &LoginResp{
		AccessToken:           reply.Token.AccessToken,
		RefreshToken:          reply.Token.RefreshToken,
		AccessTokenExpiresAt:  reply.Token.AccessTokenExpiresAt,
		RefreshTokenExpiresAt: reply.Token.RefreshTokenExpiresAt,
		SessionExpiresAt:      reply.Token.SessionExpiresAt,
		Account:               account,
	}, nil
}

type RefreshTokenResp struct {
	AccessToken           string
	RefreshToken          string
	AccessTokenExpiresAt  *time.Time
	RefreshTokenExpiresAt *time.Time
	SessionExpiresAt      *time.Time
}

func (u *AuthUsecase) RefreshToken(ctx context.Context, refreshToken string) (*RefreshTokenResp, error) {
	reply, err := u.authRepo.RefreshToken(ctx, refreshToken)
	if err != nil {
		return nil, err
	}
	return &RefreshTokenResp{
		AccessToken:           reply.AccessToken,
		RefreshToken:          reply.RefreshToken,
		AccessTokenExpiresAt:  reply.AccessTokenExpiresAt,
		RefreshTokenExpiresAt: reply.RefreshTokenExpiresAt,
		SessionExpiresAt:      reply.SessionExpiresAt,
	}, nil
}

func (u *AuthUsecase) Logout(ctx context.Context, accessToken string) error {
	return u.authRepo.Logout(ctx, accessToken)
}

func (u *AuthUsecase) CancelAccount(ctx context.Context, userID int64, password string, code string) error {
	return u.authRepo.CancelAccount(ctx, &repo.CancelAccountReq{
		UserID:   userID,
		Password: password,
		Code:     code,
	})
}
