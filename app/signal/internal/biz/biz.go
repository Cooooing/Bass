package biz

import (
	"common/pkg/client"
	"common/pkg/constant"
	"common/pkg/util"
	"common/pkg/util/task"
	"signal/internal/biz/domain"
	"signal/internal/biz/domain/handler"

	"github.com/google/wire"
)

// BizProviderSet is biz providers.
var BizProviderSet = wire.NewSet(
	util.NewEventPool,
	task.NewAsynqCache,

	client.NewAsynqServer,
	client.NewAsynqClient,
	client.NewAsynqScheduler,
	client.NewProducer,
	ProvideTasks,
	handler.NewPingHandler,
	handler.NewPowHandler,
	handler.NewSessionHandler,

	domain.NewNodeDomain,
)

func ProvideTasks(
	ping *handler.PingHandler,
	pow *handler.PowHandler,
	session *handler.SessionHandler,
) map[constant.TaskName]client.MessageHandler {
	d := make(map[constant.TaskName]client.MessageHandler)
	d[ping.Name()] = ping
	d[pow.Name()] = pow
	d[session.Name()] = session
	return d
}
