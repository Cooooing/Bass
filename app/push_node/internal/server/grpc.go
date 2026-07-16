package server

import (
	commonClient "common/pkg/client"
	"common/pkg/server"
	"fmt"
	"log/slog"
	"push_node/internal/config"
	"time"

	"github.com/go-kratos/kratos/v3/middleware/recovery"
	"github.com/go-kratos/kratos/v3/transport/grpc"
	ggrpc "google.golang.org/grpc"
	"google.golang.org/grpc/keepalive"
)

func NewGRPCServer(c *config.Bootstrap, logger *slog.Logger, obs *commonClient.Observer, services []server.Service) *grpc.Server {
	ka := []ggrpc.ServerOption{
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
	serverOpts := []grpc.ServerOption{
		grpc.Middleware(
			server.RequestLogContextMiddleware(),
			obs.ServerMiddleware(),
			recovery.Recovery(),
		),
		grpc.Options(ka...),
	}
	if c.GetGrpc().GetHost() != "" && c.GetGrpc().GetPort() != 0 {
		serverOpts = append(serverOpts, grpc.Address(fmt.Sprintf("%s:%d", c.GetGrpc().GetHost(), c.GetGrpc().GetPort())))
	}
	if c.GetGrpc().GetTimeout() != nil {
		serverOpts = append(serverOpts, grpc.Timeout(c.GetGrpc().GetTimeout().AsDuration()))
	}
	srv := grpc.NewServer(serverOpts...)
	for _, s := range services {
		s.RegisterGrpc(srv)
	}
	return srv
}
