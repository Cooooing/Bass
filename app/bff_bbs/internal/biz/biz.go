package biz

import (
	"bff_bbs/internal/biz/usecase"

	"github.com/google/wire"
)

var BizProviderSet = wire.NewSet(
	usecase.NewAuthUsecase,
	usecase.NewAccountUsecase,
	usecase.NewPreferencesUsecase,
	usecase.NewPrivacySettingUsecase,
	usecase.NewLocationUsecase,
	usecase.NewRelationUsecase,
	usecase.NewOtpUsecase,
	usecase.NewLivenessUsecase,
	usecase.NewContentArticleUsecase,
	usecase.NewContentPostscriptUsecase,
	usecase.NewContentCommentUsecase,
	usecase.NewContentDomainUsecase,
	usecase.NewContentTagUsecase,
	usecase.NewNotificationUsecase,
	usecase.NewContentMoonbreezeUsecase,
	usecase.NewHolidayUsecase,
	usecase.NewRealtimeUsecase,
	usecase.NewIMUsecase,
	usecase.NewEmojiUsecase,
	usecase.NewAssetUsecase,
)
