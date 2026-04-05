package server

import (
	"common/pkg/util/jwt"
	"common/pkg/util/server"
	"fmt"
	"signal/internal/biz/domain"
	"signal/internal/conf"
	"signal/internal/service"
	"time"

	"github.com/go-kratos/kratos/contrib/middleware/validate/v2"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-kratos/kratos/v2/middleware/logging"
	"github.com/go-kratos/kratos/v2/middleware/metrics"
	"github.com/go-kratos/kratos/v2/middleware/recovery"
	"github.com/go-kratos/kratos/v2/middleware/selector"
	"github.com/go-kratos/kratos/v2/middleware/tracing"
	"github.com/go-kratos/kratos/v2/transport/grpc"
	ggrpc "google.golang.org/grpc"
	"google.golang.org/grpc/keepalive"
)

// NewGRPCServer new a gRPC server.
func NewGRPCServer(c *conf.Bootstrap, logger log.Logger, services []service.Service, tokenCache *jwt.TokenCache, nodeDomain *domain.NodeDomain) *grpc.Server {
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
			recovery.Recovery(),
			tracing.Server(),
			metrics.Server(
				metrics.WithSeconds(_metricSeconds),
				metrics.WithRequests(_metricRequests),
			),
			logging.Server(logger),
			server.AuthMiddleware(tokenCache),
			selector.Server(SignalAuthMiddleware(nodeDomain)).Match(NodeEndpointsMatch()).Build(),
			validate.ProtoValidate(),
		),
		grpc.Options(ka...),
	}
	if c.Server.Grpc.Host != "" && c.Server.Http.Port != 0 {
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
