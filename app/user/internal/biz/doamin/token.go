package doamin

import (
	"common/pkg/cutil/jwt"
	"user/internal/biz/model"
	"user/internal/conf"
)

type TokenService struct {
	conf                      *conf.Bootstrap
	VerityCodeAccountTokenGen *jwt.TokenGenerator[model.TokenVerityCodeAccount]
	TokenGen                  *jwt.TokenGenerator[model.Token]
}

func NewTokenService(conf *conf.Bootstrap) *TokenService {
	verityCodeAccountTokenGen := jwt.NewTokenGenerator[model.TokenVerityCodeAccount](conf.Jwt.Secret)
	tokenGen := jwt.NewTokenGenerator[model.Token](conf.Jwt.Secret)
	return &TokenService{
		conf:                      conf,
		VerityCodeAccountTokenGen: verityCodeAccountTokenGen,
		TokenGen:                  tokenGen,
	}
}
