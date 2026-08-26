package data

import (
	"common/pkg/client/rpc"
	commonenum "common/pkg/enum"
	"common/pkg/server"
	userv1 "common/proto/gen/user/v1"
	userenum "common/proto/gen/user/v1/enum"
	"context"
	"crypto/sha1"
	"encoding/hex"
	"game_idle_bff/internal/biz/model"
	"game_idle_bff/internal/biz/repo"

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

func (r *AuthRepo) Register(ctx context.Context, req *repo.RegisterReq) error {
	sum := sha1.Sum([]byte(req.Email))
	_, err := r.userClient.Auth.Register(ctx, &userv1.Register_Req{
		Type:     userenum.RegisterType_REGISTER_TYPE_EMAIL,
		Name:     "idle_" + hex.EncodeToString(sum[:])[:20],
		Password: req.Password,
		Credential: &userv1.Register_Req_EmailCredential_{
			EmailCredential: &userv1.Register_Req_EmailCredential{
				Email: req.Email,
			},
		},
	})
	return err
}

func (r *AuthRepo) Login(ctx context.Context, req *repo.LoginReq) (*model.LoginToken, error) {
	uaRaw := server.GetHeader(ctx, "User-Agent")
	ua := useragent.Parse(uaRaw)
	deviceType := userenum.DeviceType_DEVICE_TYPE_DESKTOP
	if ua.Bot {
		deviceType = userenum.DeviceType_DEVICE_TYPE_BOT
	} else if ua.Tablet {
		deviceType = userenum.DeviceType_DEVICE_TYPE_TABLET
	} else if ua.Mobile {
		deviceType = userenum.DeviceType_DEVICE_TYPE_MOBILE
	}
	reply, err := r.userClient.Auth.Login(ctx, &userv1.Login_Req{
		Type:  userenum.LoginType_LOGIN_TYPE_PASSWORD,
		Realm: commonenum.LoginRealmMap.MustToProto(commonenum.LoginRealmGameIdle),
		Client: &userv1.Login_Req_ClientInfo{
			Ip:             server.ClientIP(ctx),
			UserAgent:      uaRaw,
			ClientType:     userenum.ClientType_CLIENT_TYPE_WEB,
			DeviceType:     deviceType,
			OsName:         ua.OS,
			OsVersion:      ua.OSVersion,
			BrowserName:    ua.Name,
			BrowserVersion: ua.Version,
		},
		Credential: &userv1.Login_Req_PasswordCredential_{
			PasswordCredential: &userv1.Login_Req_PasswordCredential{
				Account:  req.Email,
				Password: req.Password,
			},
		},
	})
	if err != nil {
		return nil, err
	}
	token := &model.LoginToken{
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
	if account := reply.GetAccount(); account != nil && account.GetBasic() != nil {
		basic := account.GetBasic()
		token.UserID = basic.GetId()
		token.Name = basic.GetName()
		token.Nickname = basic.GetNickname()
	}
	return token, nil
}
