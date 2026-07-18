package usecase

import (
	"bbs/internal/biz/repo"
	bbsuserv1 "common/proto/gen/bbs/v1/user"
	"time"

	"context"

	"google.golang.org/protobuf/types/known/timestamppb"
)

func (u *AuthUsecase) bbsTime(value string) *timestamppb.Timestamp {
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

type StartEmailRegistrationResp struct {
	CodeToken string
	Code      string
}

func (u *AuthUsecase) StartEmailRegistration(ctx context.Context, req *StartEmailRegistrationReq) (*StartEmailRegistrationResp, error) {
	reply, err := u.authRepo.StartEmailRegistration(ctx, &repo.StartEmailRegistrationReq{
		Email:    req.Email,
		Password: req.Password,
		Name:     req.Name,
		Nickname: req.Nickname,
	})
	if err != nil {
		return nil, err
	}
	return &StartEmailRegistrationResp{CodeToken: reply.CodeToken, Code: reply.Code}, nil
}

type VerifyEmailRegistrationReq struct {
	Code      string
	CodeToken string
}

func (u *AuthUsecase) VerifyEmailRegistration(ctx context.Context, req *VerifyEmailRegistrationReq) error {
	err := u.authRepo.VerifyEmailRegistration(ctx, &repo.VerifyEmailRegistrationReq{Code: req.Code, CodeToken: req.CodeToken})
	return err
}

type StartPhoneRegistrationReq struct {
	Phone    string
	Password string
	Name     string
	Nickname *string
}

type StartPhoneRegistrationResp struct {
	CodeToken string
	Code      string
}

func (u *AuthUsecase) StartPhoneRegistration(ctx context.Context, req *StartPhoneRegistrationReq) (*StartPhoneRegistrationResp, error) {
	reply, err := u.authRepo.StartPhoneRegistration(ctx, &repo.StartPhoneRegistrationReq{
		Phone:    req.Phone,
		Password: req.Password,
		Name:     req.Name,
		Nickname: req.Nickname,
	})
	if err != nil {
		return nil, err
	}
	return &StartPhoneRegistrationResp{CodeToken: reply.CodeToken, Code: reply.Code}, nil
}

type VerifyPhoneRegistrationReq struct {
	Code      string
	CodeToken string
}

func (u *AuthUsecase) VerifyPhoneRegistration(ctx context.Context, req *VerifyPhoneRegistrationReq) error {
	err := u.authRepo.VerifyPhoneRegistration(ctx, &repo.VerifyPhoneRegistrationReq{Code: req.Code, CodeToken: req.CodeToken})
	return err
}

type LoginByPasswordReq struct {
	Account  string
	Password string
}

type LoginByPasswordResp struct {
	Token   string
	Account *bbsuserv1.LoginByPassword_Resp_Account
}

func (u *AuthUsecase) LoginByPassword(ctx context.Context, req *LoginByPasswordReq) (*LoginByPasswordResp, error) {
	reply, err := u.authRepo.LoginByPassword(ctx, &repo.LoginByPasswordReq{Account: req.Account, Password: req.Password})
	if err != nil {
		return nil, err
	}
	var account *bbsuserv1.LoginByPassword_Resp_Account
	if reply.Account != nil {
		account = &bbsuserv1.LoginByPassword_Resp_Account{}
		if profile := reply.Account.Profile; profile != nil {
			account.Basic = &bbsuserv1.LoginByPassword_Resp_AccountBasic{
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
				CreatedAt:     u.bbsTime(profile.CreatedAt),
				UpdatedAt:     u.bbsTime(profile.UpdatedAt),
			}
		}
		if contact := reply.Account.Contact; contact != nil {
			account.Contact = &bbsuserv1.LoginByPassword_Resp_AccountContact{
				UserId: contact.UserID,
				Email:  contact.Email,
				Phone:  contact.Phone,
			}
		}
	}
	return &LoginByPasswordResp{Token: reply.Token, Account: account}, nil
}

func (u *AuthUsecase) Logout(ctx context.Context, token string) error {
	return u.authRepo.Logout(ctx, token)
}
