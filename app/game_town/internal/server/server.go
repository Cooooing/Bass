package server

import (
	"game_town/internal/biz/usecase"

	"github.com/go-kratos/kratos/v3/transport"
	"github.com/go-kratos/kratos/v3/transport/grpc"
	"github.com/go-kratos/kratos/v3/transport/http"
	"github.com/google/wire"
)

var ServerProviderSet = wire.NewSet(ProvideServers, NewGRPCServer, NewHTTPServer)

func ProvideServers(
	grpcServer *grpc.Server,
	httpServer *http.Server,
	agentRunner *usecase.WorldAgentRunner,
) []transport.Server {
	return []transport.Server{grpcServer, httpServer, agentRunner}
}
