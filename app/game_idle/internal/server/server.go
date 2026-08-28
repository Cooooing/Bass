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
	NewMetadataCacheServer,
	NewActionQueueServer,
)

func ProvideServers(
	grpcServer *grpc.Server,
	httpServer *http.Server,
	timeWheelServer *TimeWheelServer,
	metadataCacheServer *MetadataCacheServer,
	actionQueueServer *ActionQueueServer,
) []transport.Server {
	return []transport.Server{
		grpcServer,
		httpServer,
		timeWheelServer,
		metadataCacheServer,
		actionQueueServer,
	}
}
