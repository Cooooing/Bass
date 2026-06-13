package biz

import (
	"common/pkg/auth"
	"user/internal/biz/usecase"

	"github.com/google/wire"
)

// BizProviderSet 是 biz 层依赖集合。
var BizProviderSet = wire.NewSet(
	auth.NewTokenCache,
	usecase.NewTokenUsecase,
	wire.Struct(new(usecase.AuthUsecaseDeps), "*"),

	usecase.NewAccountUsecase,
	usecase.NewAuthUsecase,
	usecase.NewPreferencesUsecase,
	usecase.NewPrivacySettingUsecase,
	usecase.NewLocationUsecase,
	usecase.NewRelationUsecase,
	usecase.NewTotpUsecase,
	usecase.NewOutboxPublisher,
	usecase.NewOutboxDeadLetterScanner,
)
