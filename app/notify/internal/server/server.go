package server

import (
	"github.com/go-kratos/kratos/v3/transport"
	"github.com/go-kratos/kratos/v3/transport/grpc"
	"github.com/go-kratos/kratos/v3/transport/http"
	"github.com/google/wire"
)

// ServerProviderSet is the dependency provider set for server.
var ServerProviderSet = wire.NewSet(
	ProvideServers,
	NewConsumer,
	NewGRPCServer,
	NewHTTPServer,
	NewTemplateInitializationServer,
)

func ProvideServers(
	grpcServer *grpc.Server,
	httpServer *http.Server,
	templateInitializationServer *TemplateInitializationServer,
	consumerServer *Consumer,
) []transport.Server {
	return []transport.Server{
		templateInitializationServer,
		grpcServer,
		httpServer,
		consumerServer,
	}
}
