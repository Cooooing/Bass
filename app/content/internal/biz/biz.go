package biz

import (
	"content/internal/biz/usecase"

	"github.com/google/wire"
)

// BizProviderSet 提供 biz 层依赖集合
var BizProviderSet = wire.NewSet(
	usecase.NewArticleAccessUsecase,
	usecase.NewCommentAccessUsecase,
	usecase.NewArticleUsecase,
	usecase.NewCommentUsecase,
	usecase.NewPostscriptUsecase,
	usecase.NewContentUsecase,
	usecase.NewTagUsecase,
	usecase.NewOutboxUsecase,
)
