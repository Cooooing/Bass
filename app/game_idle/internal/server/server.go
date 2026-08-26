package server

import (
	"game_idle/internal/service"

	"github.com/go-kratos/kratos/v3/transport"
	"github.com/go-kratos/kratos/v3/transport/grpc"
	"github.com/go-kratos/kratos/v3/transport/http"
	"github.com/google/wire"
)

var ServerProviderSet = wire.NewSet(
	ProvideServers,
	NewGRPCServer,
	NewHTTPServer,
	NewTimeWheelServer,
	NewActionQueueServer,
)

func ProvideServers(
	grpcServer *grpc.Server,
	httpServer *http.Server,
	timeWheelServer *TimeWheelServer,
	actionQueueServer *ActionQueueServer,
	webSocketService *service.WebSocketService,
) []transport.Server {
	return []transport.Server{
		grpcServer,
		httpServer,
		timeWheelServer,
		actionQueueServer,
		webSocketService,
	}
}
