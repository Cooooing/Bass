package server

import (
	"fmt"
	"log/slog"
	"time"

	commonclient "common/pkg/client"
	"common/pkg/server"
	"game_town/internal/config"

	"github.com/go-kratos/kratos/contrib/middleware/validate/v3"
	"github.com/go-kratos/kratos/v3/middleware/recovery"
	"github.com/go-kratos/kratos/v3/transport/grpc"
	ggrpc "google.golang.org/grpc"
	"google.golang.org/grpc/keepalive"
)

func NewGRPCServer(
	conf *config.Bootstrap,
	log *slog.Logger,
	observer *commonclient.Observer,
	services []server.Service,
) *grpc.Server {
	grpcOptions := []ggrpc.ServerOption{
		ggrpc.KeepaliveEnforcementPolicy(keepalive.EnforcementPolicy{
			MinTime:             10 * time.Second,
			PermitWithoutStream: false,
		}),
		ggrpc.KeepaliveParams(keepalive.ServerParameters{
			MaxConnectionIdle:     300 * time.Second,
			MaxConnectionAge:      600 * time.Second,
			MaxConnectionAgeGrace: 30 * time.Second,
			Time:                  60 * time.Second,
			Timeout:               20 * time.Second,
		}),
	}
	serverOptions := []grpc.ServerOption{
		grpc.Middleware(
			server.RequestLogContextMiddleware(),
			observer.ServerMiddleware(),
			recovery.Recovery(),
			validate.ProtoValidate(),
		),
		grpc.Options(grpcOptions...),
	}
	if conf.GetGrpc().GetHost() != "" && conf.GetGrpc().GetPort() != 0 {
		serverOptions = append(serverOptions, grpc.Address(fmt.Sprintf(
			"%s:%d",
			conf.GetGrpc().GetHost(),
			conf.GetGrpc().GetPort(),
		)))
	}
	if conf.GetGrpc().GetTimeout() != nil {
		serverOptions = append(
			serverOptions,
			grpc.Timeout(conf.GetGrpc().GetTimeout().AsDuration()),
		)
	}
	grpcServer := grpc.NewServer(serverOptions...)
	for _, service := range services {
		service.RegisterGrpc(grpcServer)
	}
	return grpcServer
}
