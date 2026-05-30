package biz

import (
	"bbs/internal/biz/usecase"

	"github.com/google/wire"
)

// BizProviderSet 是 biz 层依赖集合。
var BizProviderSet = wire.NewSet(
	usecase.NewAuthUsecase,
	usecase.NewAccountUsecase,
	usecase.NewPreferencesUsecase,
	usecase.NewPrivacySettingUsecase,
	usecase.NewLocationUsecase,
	usecase.NewRelationUsecase,
	usecase.NewTotpUsecase,
	usecase.NewContentArticleUsecase,
	usecase.NewContentPostscriptUsecase,
	usecase.NewContentCommentUsecase,
	usecase.NewContentDomainUsecase,
	usecase.NewContentTagUsecase,
	usecase.NewNotificationUsecase,
)
