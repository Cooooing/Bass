package biz

import (
	"common/pkg/client"
	"common/pkg/constant"
	"common/pkg/cutil/collections/dict"
	"common/pkg/util"
	"connector/internal/biz/base"

	"github.com/google/wire"
)

// BizProviderSet is biz providers.
var BizProviderSet = wire.NewSet(
	base.NewBaseDomain,
	util.NewEventPool,

	client.NewAsynqServer,
	client.NewAsynqClient,
	client.NewProducer,
	ProvideTasks,
	NewNodeRegisterTaskHandler,

	NewSessionDomain,
	NewServerDomain,
)

func ProvideTasks(
	register *NodeRegisterTaskHandler,
) dict.Map[constant.TaskName, client.Handler] {
	d := dict.New[constant.TaskName, client.Handler](0)
	d.Set(constant.TaskConnectorRegister, register)
	return d
}
