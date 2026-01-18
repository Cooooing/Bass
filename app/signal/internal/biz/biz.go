package biz

import (
	"common/pkg/util"
	"signal/internal/biz/base"
	"signal/internal/biz/task"

	"github.com/google/wire"
)

// BizProviderSet is biz providers.
var BizProviderSet = wire.NewSet(
	base.NewBaseDomain,
	util.NewEventPool,

	task.NewNodeTaskProducer,

	NewNodeDomain,
)
