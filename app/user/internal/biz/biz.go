package biz

import (
	"user/internal/biz/usecase"

	"github.com/google/wire"
)

var BizProviderSet = wire.NewSet(
	usecase.NewTokenUsecase,
	wire.Struct(new(usecase.AuthUsecaseDeps), "*"),
	usecase.NewAccountUsecase,
	usecase.NewAuthUsecase,
	usecase.NewRbacUsecase,
	usecase.NewPreferencesUsecase,
	usecase.NewPrivacySettingUsecase,
	usecase.NewLocationUsecase,
	usecase.NewRelationUsecase,
	usecase.NewTotpUsecase,
	usecase.NewEmailOtpUsecase,
	usecase.NewSmsOtpUsecase,
	usecase.NewOutboxPublisher,
	usecase.NewOutboxDeadLetterScanner,
)
