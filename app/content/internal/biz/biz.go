package biz

import (
	"content/internal/biz/usecase"

	"github.com/google/wire"
)

// BizProviderSet 提供 biz 层依赖集合
var BizProviderSet = wire.NewSet(
	usecase.NewArticleAccessUsecase,
	usecase.NewArticleSearchUsecase,
	usecase.NewCommentAccessUsecase,
	usecase.NewArticleThankUsecase,
	usecase.NewTCCArticleThankUsecase,
	usecase.NewArticleRewardUsecase,
	usecase.NewTCCArticleRewardUsecase,
	usecase.NewTCCArticlePublishUsecase,
	usecase.NewCommentThankUsecase,
	usecase.NewTCCCommentThankUsecase,
	usecase.NewArticleUsecase,
	usecase.NewStatisticsUsecase,
	usecase.NewCommentUsecase,
	usecase.NewPostscriptUsecase,
	usecase.NewContentUsecase,
	usecase.NewTagUsecase,
	usecase.NewOutboxUsecase,
	usecase.NewMoonbreezeUsecase,
)
