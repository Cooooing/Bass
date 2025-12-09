package biz

import (
	"common/pkg/cutil/jwt"
	"user/internal/biz/model"
	"user/internal/conf"
)

type TokenService struct {
	conf          *conf.Bootstrap
	EmailTokenGen *jwt.TokenGenerator[model.TokenEmail]
	TokenGen      *jwt.TokenGenerator[model.Token]
}

func NewTokenService(conf *conf.Bootstrap) *TokenService {
	emailTokenGen := jwt.NewTokenGenerator[model.TokenEmail](conf.Jwt.Secret, conf.Jwt.EmailExpire.AsDuration())
	tokenGen := jwt.NewTokenGenerator[model.Token](conf.Jwt.Secret, conf.Jwt.Expires.AsDuration())
	return &TokenService{
		conf:          conf,
		EmailTokenGen: emailTokenGen,
		TokenGen:      tokenGen,
	}
}
