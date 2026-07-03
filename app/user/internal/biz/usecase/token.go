package usecase

import (
	"common/pkg/util/jwt"
	"user/internal/biz/model"
	"user/internal/conf"
)

type TokenUsecase struct {
	conf                      *conf.Bootstrap
	VerityCodeAccountTokenGen *jwt.TokenGenerator[model.TokenVerityCodeAccount]
	TokenGen                  *jwt.TokenGenerator[model.Token]
}

func NewTokenUsecase(conf *conf.Bootstrap) *TokenUsecase {
	verityCodeAccountTokenGen := jwt.NewTokenGenerator[model.TokenVerityCodeAccount](conf.Business.Jwt.Secret)
	tokenGen := jwt.NewTokenGenerator[model.Token](conf.Business.Jwt.Secret)
	return &TokenUsecase{
		conf:                      conf,
		VerityCodeAccountTokenGen: verityCodeAccountTokenGen,
		TokenGen:                  tokenGen,
	}
}
