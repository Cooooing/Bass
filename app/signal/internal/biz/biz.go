package biz

import (
	"common/pkg/client"
	"common/pkg/constant"
	"common/pkg/cutil/collections/dict"
	"common/pkg/util"
	"signal/internal/biz/base"

	"github.com/google/wire"
)

// BizProviderSet is biz providers.
var BizProviderSet = wire.NewSet(
	base.NewBaseDomain,
	util.NewEventPool,
	util.NewAsynqCache,

	client.NewAsynqServer,
	client.NewAsynqClient,
	client.NewAsynqScheduler,
	client.NewProducer,
	ProvideTasks,
	NewNodePingTaskHandler,
	NewNodePowTaskHandler,
	NewNodeSessionTaskHandler,

	NewNodeDomain,
)

func ProvideTasks(
	ping *NodePingTaskHandler,
	pow *NodePowTaskHandler,
	session *NodeSessionTaskHandler,
) dict.Map[constant.TaskName, client.Handler] {
	d := dict.New[constant.TaskName, client.Handler](0)
	d.Set(ping.Name(), ping)
	d.Set(pow.Name(), pow)
	d.Set(session.Name(), session)
	return d
}
