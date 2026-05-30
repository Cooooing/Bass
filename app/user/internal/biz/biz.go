package biz

import (
	"common/pkg/util/jwt"
	"user/internal/biz/usecase"

	"github.com/google/wire"
)

// BizProviderSet 是 biz 层依赖集合。
var BizProviderSet = wire.NewSet(
	jwt.NewTokenCache,
	usecase.NewTokenUsecase,

	usecase.NewAccountUsecase,
	usecase.NewAuthUsecase,
	usecase.NewPreferencesUsecase,
	usecase.NewPrivacySettingUsecase,
	usecase.NewLocationUsecase,
	usecase.NewRelationUsecase,
	usecase.NewTotpUsecase,
	usecase.NewOutboxPublisher,
)
