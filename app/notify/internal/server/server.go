package server

import (
	"notify/internal/biz/usecase"
	"notify/internal/biz/usecase/consumer"

	"github.com/go-kratos/kratos/v3/transport"
	"github.com/go-kratos/kratos/v3/transport/grpc"
	"github.com/go-kratos/kratos/v3/transport/http"
	"github.com/google/wire"
)

// ServerProviderSet 是 server 层依赖集合。
var ServerProviderSet = wire.NewSet(
	ProvideServers,
	NewGRPCServer,
	NewHTTPServer,
)

func ProvideServers(grpcServer *grpc.Server, httpServer *http.Server, consumerServer *consumer.Consumer, inboxDeadLetterScanner *usecase.InboxDeadLetterScanner) []transport.Server {
	return []transport.Server{
		grpcServer,
		httpServer,
		consumerServer,
		inboxDeadLetterScanner,
	}
}
