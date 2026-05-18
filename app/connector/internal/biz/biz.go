package biz

import (
	"common/pkg/client"
	"common/pkg/client/rpc"
	"common/pkg/constant"
	"common/pkg/util"
	"common/pkg/util/jwt"
	"common/pkg/util/task"
	"connector/internal/biz/domain"
	"connector/internal/biz/domain/handler"

	"github.com/google/wire"
)

// BizProviderSet is biz providers.
var BizProviderSet = wire.NewSet(
	util.NewEventPool,
	task.NewAsynqCache,

	jwt.NewTokenCache,

	client.NewAsynqServer,
	client.NewAsynqClient,
	client.NewAsynqScheduler,
	client.NewProducer,
	ProvideTasks,
	handler.NewRegisterHandler,

	rpc.ProvideSignalNodeClient,

	domain.NewSessionDomain,
	domain.NewServerDomain,
)

func ProvideTasks(
	register *handler.RegisterHandler,
) map[constant.TaskName]client.Handler {
	d := make(map[constant.TaskName]client.Handler)
	d[constant.TaskConnectorRegister] = register
	return d
}
