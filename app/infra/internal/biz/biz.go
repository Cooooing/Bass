package biz

import (
	"common/pkg/util"
	doaminbase "infra/internal/biz/base"
	"infra/internal/biz/domain"

	"github.com/google/wire"
)

// BizProviderSet is biz providers.
var BizProviderSet = wire.NewSet(
	doaminbase.NewBaseDomain,
	util.NewEventPool,
	domain.NewObjectStorageDomain,
)
