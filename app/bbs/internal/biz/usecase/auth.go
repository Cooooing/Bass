package usecase

import (
	"bbs/internal/biz/repo"
	bbsuserv1 "common/proto/gen/bbs/v1/user"
	"time"

	"context"

	"google.golang.org/protobuf/types/known/timestamppb"
)

func bbsTime(value string) *timestamppb.Timestamp {
	if value == "" {
		return nil
	}
	parsed, err := time.Parse(time.RFC3339Nano, value)
	if err != nil {
		return nil
	}
	return timestamppb.New(parsed)
}

type AuthUsecase struct {
	authRepo repo.AuthRepo
}

func NewAuthUsecase(authRepo repo.AuthRepo) *AuthUsecase {
	return &AuthUsecase{authRepo: authRepo}
}

type StartEmailRegistrationReq struct {
	Email    string
	Password string
	Name     string
	Nickname *string
}

type StartEmailRegistrationResponse struct {
	CodeToken string
	Code      string
}

func (u *AuthUsecase) StartEmailRegistration(ctx context.Context, req *StartEmailRegistrationReq) (*StartEmailRegistrationResponse, error) {
	reply, err := u.authRepo.StartEmailRegistration(ctx, &repo.StartEmailRegistrationReq{
		Email:    req.Email,
		Password: req.Password,
		Name:     req.Name,
		Nickname: req.Nickname,
	})
	if err != nil {
		return nil, err
	}
	return &StartEmailRegistrationResponse{CodeToken: reply.CodeToken, Code: reply.Code}, nil
}

type VerifyEmailRegistrationReq struct {
	Code      string
	CodeToken string
}

func (u *AuthUsecase) VerifyEmailRegistration(ctx context.Context, req *VerifyEmailRegistrationReq) error {
	_, err := u.authRepo.VerifyEmailRegistration(ctx, &repo.VerifyEmailRegistrationReq{Code: req.Code, CodeToken: req.CodeToken})
	return err
}

type StartPhoneRegistrationReq struct {
	Phone    string
	Password string
	Name     string
	Nickname *string
}

type StartPhoneRegistrationResponse struct {
	CodeToken string
	Code      string
}

func (u *AuthUsecase) StartPhoneRegistration(ctx context.Context, req *StartPhoneRegistrationReq) (*StartPhoneRegistrationResponse, error) {
	reply, err := u.authRepo.StartPhoneRegistration(ctx, &repo.StartPhoneRegistrationReq{
		Phone:    req.Phone,
		Password: req.Password,
		Name:     req.Name,
		Nickname: req.Nickname,
	})
	if err != nil {
		return nil, err
	}
	return &StartPhoneRegistrationResponse{CodeToken: reply.CodeToken, Code: reply.Code}, nil
}

type VerifyPhoneRegistrationReq struct {
	Code      string
	CodeToken string
}

func (u *AuthUsecase) VerifyPhoneRegistration(ctx context.Context, req *VerifyPhoneRegistrationReq) error {
	_, err := u.authRepo.VerifyPhoneRegistration(ctx, &repo.VerifyPhoneRegistrationReq{Code: req.Code, CodeToken: req.CodeToken})
	return err
}

type LoginByPasswordReq struct {
	Account  string
	Password string
}

type LoginByPasswordResponse struct {
	Token   string
	Account *bbsuserv1.LoginByPassword_Response_Account
}

func (u *AuthUsecase) LoginByPassword(ctx context.Context, req *LoginByPasswordReq) (*LoginByPasswordResponse, error) {
	reply, err := u.authRepo.LoginByPassword(ctx, &repo.LoginByPasswordReq{Account: req.Account, Password: req.Password})
	if err != nil {
		return nil, err
	}
	var account *bbsuserv1.LoginByPassword_Response_Account
	if reply.Account != nil {
		account = &bbsuserv1.LoginByPassword_Response_Account{}
		if profile := reply.Account.Profile; profile != nil {
			account.Basic = &bbsuserv1.LoginByPassword_Response_AccountBasic{
				Id:            profile.ID,
				Name:          profile.Name,
				Nickname:      profile.Nickname,
				Url:           profile.URL,
				AvatarUrl:     profile.AvatarURL,
				Introduction:  profile.Introduction,
				Status:        bbsuserv1.AccountStatus(profile.Status),
				Mbti:          bbsuserv1.MBTI(profile.MBTI),
				FollowCount:   profile.FollowCount,
				FollowerCount: profile.FollowerCount,
				CreatedAt:     bbsTime(profile.CreatedAt),
				UpdatedAt:     bbsTime(profile.UpdatedAt),
			}
		}
		if contact := reply.Account.Contact; contact != nil {
			account.Contact = &bbsuserv1.LoginByPassword_Response_AccountContact{
				UserId: contact.UserID,
				Email:  contact.Email,
				Phone:  contact.Phone,
			}
		}
	}
	return &LoginByPasswordResponse{Token: reply.Token, Account: account}, nil
}

type LogoutReq struct {
	Token string
}

func (u *AuthUsecase) Logout(ctx context.Context, req *LogoutReq) error {
	_, err := u.authRepo.Logout(ctx, &repo.LogoutReq{Token: req.Token})
	return err
}
