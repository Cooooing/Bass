package server

import (
	"github.com/go-kratos/kratos/v3/transport"
	"github.com/go-kratos/kratos/v3/transport/grpc"
	"github.com/go-kratos/kratos/v3/transport/http"
	"github.com/google/wire"
)

var ServerProviderSet = wire.NewSet(ProvideServers, NewGRPCServer, NewHTTPServer, NewRealtimeServer)

func ProvideServers(realtime *RealtimeServer, grpcServer *grpc.Server, httpServer *http.Server) []transport.Server {
	return []transport.Server{realtime, grpcServer, httpServer}
}
