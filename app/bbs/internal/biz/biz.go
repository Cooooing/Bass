package biz

import (
	"bbs/internal/biz/usecase"

	"github.com/google/wire"
)

// BizProviderSet 是 biz 层依赖集合。
var BizProviderSet = wire.NewSet(
	usecase.NewUserUsecase,
	usecase.NewContentUsecase,
	usecase.NewNotifyUsecase,
)
