package server

import (
	"github.com/go-kratos/kratos/v3/transport"
	"github.com/go-kratos/kratos/v3/transport/grpc"
	"github.com/go-kratos/kratos/v3/transport/http"
	"github.com/google/wire"
)

var ProviderSet = wire.NewSet(
	NewGRPCServer,
	NewHTTPServer,
	ProvideServers,
)

func ProvideServers(grpcServer *grpc.Server, httpServer *http.Server) []transport.Server {
	return []transport.Server{
		grpcServer,
		httpServer,
	}
}
