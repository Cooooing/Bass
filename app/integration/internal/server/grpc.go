package server

import (
	commonClient "common/pkg/client"
	"common/pkg/server"
	"fmt"
	"integration/internal/conf"
	"log/slog"
	"time"

	"github.com/go-kratos/kratos/contrib/middleware/validate/v3"
	"github.com/go-kratos/kratos/v3/middleware/recovery"
	"github.com/go-kratos/kratos/v3/transport/grpc"
	ggrpc "google.golang.org/grpc"
	"google.golang.org/grpc/keepalive"
)

func NewGRPCServer(c *conf.Bootstrap, logger *slog.Logger, obs *commonClient.Observer, services []server.GrpcService) *grpc.Server {
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
	var serverOpts = []grpc.ServerOption{
		grpc.Middleware(
			server.RequestLogContextMiddleware(),
			obs.ServerMiddleware(),
			recovery.Recovery(),
			validate.ProtoValidate(),
		),
		grpc.Options(ka...),
	}
	if c.Server.Grpc.Host != "" && c.Server.Grpc.Port != 0 {
		serverOpts = append(serverOpts, grpc.Address(fmt.Sprintf("%s:%d", c.Server.Grpc.Host, c.Server.Grpc.Port)))
	}
	if c.Server.Grpc.Timeout != nil {
		serverOpts = append(serverOpts, grpc.Timeout(c.Server.Grpc.Timeout.AsDuration()))
	}
	srv := grpc.NewServer(serverOpts...)
	for _, s := range services {
		s.RegisterGrpc(srv)
	}
	return srv
}
