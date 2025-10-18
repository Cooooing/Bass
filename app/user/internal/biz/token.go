package biz

import (
	"common/pkg/util"
	"user/internal/biz/model"
	"user/internal/conf"
)

type TokenService struct {
	conf          *conf.Bootstrap
	EmailTokenGen *util.TokenGenerator[model.TokenEmail]
	TokenGen      *util.TokenGenerator[model.Token]
}

func NewTokenService(conf *conf.Bootstrap) *TokenService {
	emailTokenGen := util.NewTokenGenerator[model.TokenEmail](conf.Jwt.Secret, conf.Jwt.EmailExpire.AsDuration())
	tokenGen := util.NewTokenGenerator[model.Token](conf.Jwt.Secret, conf.Jwt.EmailExpire.AsDuration())
	return &TokenService{
		conf:          conf,
		EmailTokenGen: emailTokenGen,
		TokenGen:      tokenGen,
	}
}
