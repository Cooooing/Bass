package usecase

import (
	commonenum "common/pkg/enum"
	"common/pkg/util/jwt"
	commonenums "common/proto/gen/common/enums"
	"time"
	"user/internal/biz/model"
	"user/internal/config"

	"google.golang.org/protobuf/types/known/durationpb"
)

const (
	tokenTypeAccess  = "access"
	tokenTypeRefresh = "refresh"
)

type TokenUsecase struct {
	conf     *config.Bootstrap
	TokenGen *jwt.TokenGenerator[model.Token]
}

func NewTokenUsecase(conf *config.Bootstrap) *TokenUsecase {
	return &TokenUsecase{
		conf:     conf,
		TokenGen: jwt.NewTokenGenerator[model.Token](sessionSecret(conf)),
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

func (u *TokenUsecase) GenerateAccess(req *GenerateAccessTokenReq) (string, time.Time, error) {
	ttl := u.AccessTokenTTL()
	expiresAt := time.Now().Add(ttl)
	token, err := u.TokenGen.Generate(model.Token{
		Type:      tokenTypeAccess,
		UserID:    req.UserID,
		SessionID: req.SessionID,
		Realm:     req.Realm,
		Name:      req.Name,
		Nickname:  req.Nickname,
		Language:  req.Language,
		Timezone:  req.Timezone,
	}, ttl)
	return token, expiresAt, err
}

type GenerateRefreshTokenReq struct {
	UserID    int64
	SessionID string
	Realm     commonenum.LoginRealm
	JTI       string
}

func (u *TokenUsecase) GenerateRefresh(req *GenerateRefreshTokenReq) (string, time.Time, error) {
	ttl := u.RefreshTokenTTL()
	expiresAt := time.Now().Add(ttl)
	token, err := u.TokenGen.Generate(model.Token{
		Type:      tokenTypeRefresh,
		UserID:    req.UserID,
		SessionID: req.SessionID,
		Realm:     req.Realm,
		JTI:       req.JTI,
	}, ttl)
	return token, expiresAt, err
}

func (u *TokenUsecase) Parse(token string) (model.Token, error) {
	return u.TokenGen.Parse(token)
}

func (u *TokenUsecase) AccessTokenTTL() time.Duration {
	return durationOrDefault(u.conf.GetBusiness().GetAuth().GetSession().GetAccessTokenTtl(), 15*time.Minute)
}

func (u *TokenUsecase) RefreshTokenTTL() time.Duration {
	return durationOrDefault(u.conf.GetBusiness().GetAuth().GetSession().GetRefreshTokenTtl(), 30*24*time.Hour)
}

func (u *TokenUsecase) SessionTTL() time.Duration {
	return durationOrDefault(u.conf.GetBusiness().GetAuth().GetSession().GetSessionTtl(), 180*24*time.Hour)
}

func sessionSecret(conf *config.Bootstrap) string {
	secret := conf.GetBusiness().GetAuth().GetSession().GetSecret()
	if secret == "" {
		return "dev-user-session-secret"
	}
	return secret
}

func durationOrDefault(duration *durationpb.Duration, fallback time.Duration) time.Duration {
	if duration == nil || duration.AsDuration() <= 0 {
		return fallback
	}
	return duration.AsDuration()
}
