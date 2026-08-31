package server

import (
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
	NewCharacterActionQueueServer,
)

func ProvideServers(
	grpcServer *grpc.Server,
	httpServer *http.Server,
	timeWheelServer *TimeWheelServer,
	actionQueueServer *CharacterActionQueueServer,
) []transport.Server {
	return []transport.Server{
		grpcServer,
		httpServer,
		timeWheelServer,
		actionQueueServer,
	}
}
