package server

import (
	"scheduler/internal/biz/usecase"

	"github.com/go-kratos/kratos/v3/transport"
	"github.com/go-kratos/kratos/v3/transport/grpc"
	"github.com/go-kratos/kratos/v3/transport/http"
	"github.com/google/wire"
)

var ServerProviderSet = wire.NewSet(
	ProvideServers,
	NewGRPCServer,
	NewHTTPServer,
)

func ProvideServers(
	grpcServer *grpc.Server,
	httpServer *http.Server,
	schedulerRunner *usecase.SchedulerRunner,
	delayedTaskRunner *usecase.DelayedTaskRunner,
) []transport.Server {
	return []transport.Server{grpcServer, httpServer, schedulerRunner, delayedTaskRunner}
}
