package biz

import (
	"common/pkg/client/rpc"
	"common/pkg/util"
	"content/internal/biz/usecase"

	"github.com/google/wire"
)

// BizProviderSet is biz providers.
var BizProviderSet = wire.NewSet(
	util.NewEventPool,

	rpc.ProvideUserClient,

	usecase.NewArticleUsecase,
	usecase.NewCommentUsecase,
	usecase.NewContentUsecase,
	usecase.NewTagUsecase,
)
