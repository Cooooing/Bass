package server

import (
	"github.com/go-kratos/kratos/v3/transport"
	"github.com/go-kratos/kratos/v3/transport/grpc"
	"github.com/go-kratos/kratos/v3/transport/http"
	"github.com/google/wire"
)

// ServerProviderSet 是 server 层依赖集合。
var ServerProviderSet = wire.NewSet(
	ProvideServers,
	NewEventStreamServer,
	NewGRPCServer,
	NewHTTPServer,
)

func ProvideServers(grpcServer *grpc.Server, httpServer *http.Server, eventStreamServer *EventStreamServer) []transport.Server {
	return []transport.Server{
		eventStreamServer,
		grpcServer,
		httpServer,
	}
}
