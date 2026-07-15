package usecase

import (
	"common/pkg/util/jwt"
	"user/internal/biz/model"
	"user/internal/config"
)

type TokenUsecase struct {
	conf                      *config.Bootstrap
	VerityCodeAccountTokenGen *jwt.TokenGenerator[model.TokenVerityCodeAccount]
	TokenGen                  *jwt.TokenGenerator[model.Token]
}

func NewTokenUsecase(conf *config.Bootstrap) *TokenUsecase {
	verityCodeAccountTokenGen := jwt.NewTokenGenerator[model.TokenVerityCodeAccount](conf.Business.Jwt.Secret)
	tokenGen := jwt.NewTokenGenerator[model.Token](conf.Business.Jwt.Secret)
	return &TokenUsecase{
		conf:                      conf,
		VerityCodeAccountTokenGen: verityCodeAccountTokenGen,
		TokenGen:                  tokenGen,
	}
}
