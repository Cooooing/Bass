package biz

import (
	"common/pkg/client"
	"common/pkg/constant"
	"common/pkg/cutil/collections/dict"
	"common/pkg/util"
	"connector/internal/biz/base"
	"connector/internal/biz/domain"
	"connector/internal/biz/domain/handler"

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
	handler.NewRegisterHandler,

	domain.NewSessionDomain,
	domain.NewServerDomain,
)

func ProvideTasks(
	register *handler.RegisterHandler,
) dict.Map[constant.TaskName, client.Handler] {
	d := dict.New[constant.TaskName, client.Handler](0)
	d.Set(constant.TaskConnectorRegister, register)
	return d
}
