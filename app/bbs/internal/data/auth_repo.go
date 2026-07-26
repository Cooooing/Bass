package data

import (
	"bbs/internal/biz/repo"
	"bbs/internal/enum"
	"common/pkg/client/rpc"
	"common/pkg/constant"
	commonenum "common/pkg/enum"
	"common/pkg/server"

	userv1 "common/proto/gen/user/v1"
	userenum "common/proto/gen/user/v1/enum"
	"context"

	"github.com/mileusna/useragent"
)

var _ repo.AuthRepo = (*AuthRepo)(nil)

type AuthRepo struct {
	userClient *rpc.UserClient
}

func NewAuthRepo(
	userClient *rpc.UserClient,
) repo.AuthRepo {
	return &AuthRepo{
		userClient: userClient,
	}
}

func NewAuthClient(
	userClient *rpc.UserClient,
) repo.AuthRepo {
	return NewAuthRepo(userClient)
}

func (r *AuthRepo) StartEmailRegistration(ctx context.Context, req *repo.StartEmailRegistrationReq) (*repo.StartEmailRegistrationResp, error) {
	reply, err := r.userClient.Auth.StartEmailRegistration(ctx, &userv1.StartEmailRegistration_Req{
		Email:    req.Email,
		Password: req.Password,
		Name:     req.Name,
		Nickname: req.Nickname,
	})
	if err != nil {
		return nil, err
	}
	return &repo.StartEmailRegistrationResp{
		Code: reply.GetCode(),
	}, nil
}

func (r *AuthRepo) VerifyEmailRegistration(ctx context.Context, req *repo.VerifyEmailRegistrationReq) error {
	_, err := r.userClient.Auth.VerifyEmailRegistration(ctx, &userv1.VerifyEmailRegistration_Req{
		Email: req.Email,
		Code:  req.Code,
	})
	return err
}

func (r *AuthRepo) StartPhoneRegistration(ctx context.Context, req *repo.StartPhoneRegistrationReq) (*repo.StartPhoneRegistrationResp, error) {
	reply, err := r.userClient.Auth.StartPhoneRegistration(ctx, &userv1.StartPhoneRegistration_Req{
		Phone:    req.Phone,
		Password: req.Password,
		Name:     req.Name,
		Nickname: req.Nickname,
	})
	if err != nil {
		return nil, err
	}
	return &repo.StartPhoneRegistrationResp{
		Code: reply.GetCode(),
	}, nil
}

func (r *AuthRepo) VerifyPhoneRegistration(ctx context.Context, req *repo.VerifyPhoneRegistrationReq) error {
	_, err := r.userClient.Auth.VerifyPhoneRegistration(ctx, &userv1.VerifyPhoneRegistration_Req{
		Phone: req.Phone,
		Code:  req.Code,
	})
	return err
}

func (r *AuthRepo) Login(ctx context.Context, req *repo.LoginReq) (*repo.LoginResp, error) {
	loginType := userenum.LoginType_LOGIN_TYPE_PASSWORD
	switch req.Type {
	case enum.LoginTypeEmail:
		loginType = userenum.LoginType_LOGIN_TYPE_EMAIL
	case enum.LoginTypePhone:
		loginType = userenum.LoginType_LOGIN_TYPE_PHONE
	}

	uaRaw := server.GetHeader(ctx, constant.HeaderUserAgent)
	ua := useragent.Parse(uaRaw)
	deviceType := userenum.DeviceType_DEVICE_TYPE_DESKTOP
	if ua.Bot {
		deviceType = userenum.DeviceType_DEVICE_TYPE_BOT
	} else if ua.Tablet {
		deviceType = userenum.DeviceType_DEVICE_TYPE_TABLET
	} else if ua.Mobile {
		deviceType = userenum.DeviceType_DEVICE_TYPE_MOBILE
	}

	loginReq := &userv1.Login_Req{
		Type:  loginType,
		Realm: commonenum.LoginRealmMap.MustToProto(commonenum.LoginRealmBBS),
		Client: &userv1.Login_Req_ClientInfo{
			Ip:             server.ClientIP(ctx),
			UserAgent:      uaRaw,
			ClientType:     userenum.ClientType_CLIENT_TYPE_WEB,
			DeviceType:     deviceType,
			OsName:         ua.OS,
			OsVersion:      ua.OSVersion,
			BrowserName:    ua.Name,
			BrowserVersion: ua.Version,
			AppName:        server.GetHeader(ctx, constant.HeaderBassAppName),
			AppVersion:     server.GetHeader(ctx, constant.HeaderBassAppVersion),
		},
	}
	switch req.Type {
	case enum.LoginTypePassword:
		loginReq.Credential = &userv1.Login_Req_PasswordCredential_{
			PasswordCredential: &userv1.Login_Req_PasswordCredential{
				Account:  req.Account,
				Password: req.Password,
				Code:     req.Code,
			},
		}
	case enum.LoginTypeEmail:
		loginReq.Credential = &userv1.Login_Req_EmailCredential_{
			EmailCredential: &userv1.Login_Req_EmailCredential{
				Email: req.Email,
				Code:  req.Code,
			},
		}
	case enum.LoginTypePhone:
		loginReq.Credential = &userv1.Login_Req_PhoneCredential_{
			PhoneCredential: &userv1.Login_Req_PhoneCredential{
				Phone: req.Phone,
				Code:  req.Code,
			},
		}
	}

	reply, err := r.userClient.Auth.Login(ctx, loginReq)
	if err != nil {
		return nil, err
	}

	var account *repo.Account
	if reply.GetAccount() != nil {
		account = &repo.Account{}
		if basic := reply.GetAccount().GetBasic(); basic != nil {
			account.Profile = &repo.AccountProfile{
				ID:            basic.GetId(),
				Name:          basic.GetName(),
				Nickname:      basic.Nickname,
				URL:           basic.Url,
				AvatarURL:     basic.AvatarUrl,
				Introduction:  basic.Introduction,
				Status:        int32(basic.GetStatus()),
				MBTI:          int32(basic.GetMbti()),
				FollowCount:   basic.FollowCount,
				FollowerCount: basic.FollowerCount,
			}
			if basic.CreatedAt != nil {
				account.Profile.CreatedAt = new(basic.CreatedAt.AsTime())
			}
			if basic.UpdatedAt != nil {
				account.Profile.UpdatedAt = new(basic.UpdatedAt.AsTime())
			}
		}
		if contact := reply.GetAccount().GetContact(); contact != nil {
			account.Contact = &repo.AccountContact{
				UserID: contact.GetUserId(),
				Email:  contact.Email,
				Phone:  contact.Phone,
			}
		}
	}

	token := repo.TokenResp{
		AccessToken:  reply.GetAccessToken(),
		RefreshToken: reply.GetRefreshToken(),
	}
	if reply.AccessTokenExpiresAt != nil {
		token.AccessTokenExpiresAt = new(reply.AccessTokenExpiresAt.AsTime())
	}
	if reply.RefreshTokenExpiresAt != nil {
		token.RefreshTokenExpiresAt = new(reply.RefreshTokenExpiresAt.AsTime())
	}
	if reply.SessionExpiresAt != nil {
		token.SessionExpiresAt = new(reply.SessionExpiresAt.AsTime())
	}
	return &repo.LoginResp{
		Token:   token,
		Account: account,
	}, nil
}

func (r *AuthRepo) RefreshToken(ctx context.Context, refreshToken string) (*repo.TokenResp, error) {
	reply, err := r.userClient.Auth.RefreshToken(ctx, &userv1.RefreshToken_Req{
		RefreshToken: refreshToken,
		Realm:        commonenum.LoginRealmMap.MustToProto(commonenum.LoginRealmBBS),
	})
	if err != nil {
		return nil, err
	}
	token := &repo.TokenResp{
		AccessToken:  reply.GetAccessToken(),
		RefreshToken: reply.GetRefreshToken(),
	}
	if reply.AccessTokenExpiresAt != nil {
		token.AccessTokenExpiresAt = new(reply.AccessTokenExpiresAt.AsTime())
	}
	if reply.RefreshTokenExpiresAt != nil {
		token.RefreshTokenExpiresAt = new(reply.RefreshTokenExpiresAt.AsTime())
	}
	if reply.SessionExpiresAt != nil {
		token.SessionExpiresAt = new(reply.SessionExpiresAt.AsTime())
	}
	return token, nil
}

func (r *AuthRepo) Logout(ctx context.Context, accessToken string) error {
	_, err := r.userClient.Auth.Logout(ctx, &userv1.Logout_Req{
		AccessToken: accessToken,
		Realm:       commonenum.LoginRealmMap.MustToProto(commonenum.LoginRealmBBS),
	})
	return err
}

func (r *AuthRepo) CancelAccount(ctx context.Context, req *repo.CancelAccountReq) error {
	_, err := r.userClient.Auth.CancelAccount(ctx, &userv1.CancelAccount_Req{
		UserId:   req.UserID,
		Password: req.Password,
		Code:     req.Code,
	})
	return err
}
