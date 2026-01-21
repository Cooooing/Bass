package biz

import (
	"common/pkg/cutil/collections/dict"
	"common/pkg/util"
	"signal/internal/biz/base"
	"signal/internal/biz/task"

	"github.com/google/wire"
)

// BizProviderSet is biz providers.
var BizProviderSet = wire.NewSet(
	base.NewBaseDomain,
	util.NewEventPool,

	task.NewAsynqServer,
	task.NewAsynqClient,
	ProvideTasks,
	NewProducer,
	NewNodePingTaskHandler,

	NewNodeDomain,
)

func ProvideTasks(
	ping *NodePingTaskHandler,
) dict.Map[task.TaskName, task.Handler] {
	d := dict.New[task.TaskName, task.Handler](0)
	d.Set(task.TaskNodePing, ping)
	return d
}
