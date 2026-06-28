package server

import (
	"common/pkg/server"
	"fmt"
	"github.com/go-kratos/kratos/contrib/middleware/validate/v3"
	"github.com/go-kratos/kratos/v3/middleware/logging"
	"github.com/go-kratos/kratos/v3/middleware/recovery"
	"github.com/go-kratos/kratos/v3/transport/grpc"
	ggrpc "google.golang.org/grpc"
	"google.golang.org/grpc/keepalive"
	"integration/internal/conf"
	"log/slog"
	"time"
)

// NewGRPCServer 创建 gRPC 服务。
func NewGRPCServer(c *conf.Bootstrap, logger *slog.Logger, services []server.GrpcService) *grpc.Server {
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
			server.MetricsMiddleware(c.Server.Name),
			server.TracingMiddleware(c.Server.Name),
			recovery.Recovery(),
			logging.Server(logger),
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
