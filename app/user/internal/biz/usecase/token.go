package usecase

import (
	commonenum "common/pkg/enum"
	"common/pkg/util/jwt"
	commonenums "common/proto/gen/common/enums"
	"time"
	"user/internal/biz/model"
	"user/internal/config"
	"user/internal/enum"
)

type TokenUsecase struct {
	conf     *config.Bootstrap
	TokenGen *jwt.TokenGenerator[model.Token]
}

func NewTokenUsecase(
	conf *config.Bootstrap,
) *TokenUsecase {
	secret := conf.GetBusiness().GetAuth().GetSession().GetSecret()
	if secret == "" {
		secret = "dev-user-session-secret"
	}
	return &TokenUsecase{
		conf:     conf,
		TokenGen: jwt.NewTokenGenerator[model.Token](secret),
	}
}

type GenerateAccessTokenReq struct {
	UserID    int64
	SessionID string
	Realm     commonenum.LoginRealm
	Name      string
	Nickname  string
	Language  commonenums.Language
	Timezone  string
}

func (u *TokenUsecase) GenerateAccess(req *GenerateAccessTokenReq) (string, *time.Time, error) {
	ttl := u.AccessTokenTTL()
	expiresAt := time.Now().Add(ttl)
	token, err := u.TokenGen.Generate(model.Token{
		Type:      enum.TokenTypeAccess,
		UserID:    req.UserID,
		SessionID: req.SessionID,
		Realm:     req.Realm,
		Name:      req.Name,
		Nickname:  req.Nickname,
		Language:  req.Language,
		Timezone:  req.Timezone,
	}, ttl)
	return token, new(expiresAt), err
}

type GenerateRefreshTokenReq struct {
	UserID    int64
	SessionID string
	Realm     commonenum.LoginRealm
	JTI       string
}

func (u *TokenUsecase) GenerateRefresh(req *GenerateRefreshTokenReq) (string, *time.Time, error) {
	ttl := u.RefreshTokenTTL()
	expiresAt := time.Now().Add(ttl)
	token, err := u.TokenGen.Generate(model.Token{
		Type:      enum.TokenTypeRefresh,
		UserID:    req.UserID,
		SessionID: req.SessionID,
		Realm:     req.Realm,
		JTI:       req.JTI,
	}, ttl)
	return token, new(expiresAt), err
}

func (u *TokenUsecase) Parse(token string) (model.Token, error) {
	return u.TokenGen.Parse(token)
}

func (u *TokenUsecase) AccessTokenTTL() time.Duration {
	ttl := u.conf.GetBusiness().GetAuth().GetSession().GetAccessTokenTtl()
	if ttl == nil || ttl.AsDuration() <= 0 {
		return 15 * time.Minute
	}
	return ttl.AsDuration()
}

func (u *TokenUsecase) RefreshTokenTTL() time.Duration {
	ttl := u.conf.GetBusiness().GetAuth().GetSession().GetRefreshTokenTtl()
	if ttl == nil || ttl.AsDuration() <= 0 {
		return 30 * 24 * time.Hour
	}
	return ttl.AsDuration()
}

func (u *TokenUsecase) SessionTTL() time.Duration {
	ttl := u.conf.GetBusiness().GetAuth().GetSession().GetSessionTtl()
	if ttl == nil || ttl.AsDuration() <= 0 {
		return 180 * 24 * time.Hour
	}
	return ttl.AsDuration()
}
