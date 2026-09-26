package server

import (
	"bff_bbs_admin/internal/config"
	commonClient "common/pkg/client"
	commonserver "common/pkg/server"
	"fmt"
	"log/slog"
	"time"

	"github.com/go-kratos/kratos/contrib/middleware/validate/v3"
	"github.com/go-kratos/kratos/v3/middleware/recovery"
	"github.com/go-kratos/kratos/v3/transport/grpc"
	ggrpc "google.golang.org/grpc"
	"google.golang.org/grpc/keepalive"
)

func NewGRPCServer(c *config.Bootstrap, _ *slog.Logger, obs *commonClient.Observer, services []commonserver.Service) *grpc.Server {
	ka := []ggrpc.ServerOption{
		ggrpc.KeepaliveEnforcementPolicy(keepalive.EnforcementPolicy{MinTime: 10 * time.Second}),
		ggrpc.KeepaliveParams(keepalive.ServerParameters{
			MaxConnectionIdle:     300 * time.Second,
			MaxConnectionAge:      600 * time.Second,
			MaxConnectionAgeGrace: 30 * time.Second,
			Time:                  60 * time.Second,
			Timeout:               20 * time.Second,
		}),
	}
	opts := []grpc.ServerOption{
		grpc.Middleware(
			commonserver.RequestLogContextMiddleware(),
			obs.ServerMiddleware(),
			recovery.Recovery(),
			validate.ProtoValidate(),
		),
		grpc.Options(ka...),
	}
	if c.GetGrpc().GetHost() != "" && c.GetGrpc().GetPort() != 0 {
		opts = append(opts, grpc.Address(fmt.Sprintf("%s:%d", c.GetGrpc().GetHost(), c.GetGrpc().GetPort())))
	}
	if c.GetGrpc().GetTimeout() != nil {
		opts = append(opts, grpc.Timeout(c.GetGrpc().GetTimeout().AsDuration()))
	}
	srv := grpc.NewServer(opts...)
	for _, s := range services {
		s.RegisterGrpc(srv)
	}
	return srv
}
