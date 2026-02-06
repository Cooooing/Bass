package server

import (
	"common/pkg"
	"common/pkg/util"
	"content/internal/conf"
	"content/internal/service"
	"fmt"
	"net/url"

	"github.com/go-kratos/kratos/contrib/middleware/validate/v2"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-kratos/kratos/v2/middleware/logging"
	"github.com/go-kratos/kratos/v2/middleware/metrics"
	"github.com/go-kratos/kratos/v2/middleware/recovery"
	"github.com/go-kratos/kratos/v2/middleware/tracing"
	"github.com/go-kratos/kratos/v2/transport/grpc"
)

// NewGRPCServer new a gRPC server.
func NewGRPCServer(c *conf.Bootstrap, logger log.Logger, services []service.Service, tokenCache *util.TokenCache) *grpc.Server {
	var opts = []grpc.ServerOption{
		grpc.Middleware(
			recovery.Recovery(),
			tracing.Server(),
			metrics.Server(
				metrics.WithSeconds(_metricSeconds),
				metrics.WithRequests(_metricRequests),
			),
			logging.Server(logger),
			pkg.AuthMiddleware(tokenCache),
			validate.ProtoValidate(),
		),
	}
	if c.Server.RegisterAddress != "" {
		opts = append(opts, grpc.Endpoint(&url.URL{
			Scheme: "grpc",
			Host:   fmt.Sprintf("%s:%d", c.Server.RegisterAddress, c.Server.Grpc.Port),
		}))
	}
	if c.Server.Grpc.Host != "" && c.Server.Http.Port != 0 {
		opts = append(opts, grpc.Address(fmt.Sprintf("%s:%d", c.Server.Grpc.Host, c.Server.Grpc.Port)))
	}
	if c.Server.Grpc.Timeout != nil {
		opts = append(opts, grpc.Timeout(c.Server.Grpc.Timeout.AsDuration()))
	}
	srv := grpc.NewServer(opts...)
	for _, s := range services {
		s.RegisterGrpc(srv)
	}
	return srv
}
