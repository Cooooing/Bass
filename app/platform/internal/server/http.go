package server

import (
	"common/pkg/server"
	"fmt"
	"platform/internal/conf"

	"github.com/go-kratos/kratos/contrib/middleware/validate/v2"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-kratos/kratos/v2/middleware/logging"
	"github.com/go-kratos/kratos/v2/middleware/metrics"
	"github.com/go-kratos/kratos/v2/middleware/recovery"
	"github.com/go-kratos/kratos/v2/middleware/tracing"
	transporthttp "github.com/go-kratos/kratos/v2/transport/http"
)

// NewHTTPServer 创建 HTTP 服务，用于接收 OSS 回调。
func NewHTTPServer(c *conf.Bootstrap, logger log.Logger, services []server.HttpService) *transporthttp.Server {
	var opts = []transporthttp.ServerOption{
		transporthttp.Middleware(
			recovery.Recovery(),
			tracing.Server(),
			metrics.Server(
				metrics.WithSeconds(_metricSeconds),
				metrics.WithRequests(_metricRequests),
			),
			logging.Server(logger),
			validate.ProtoValidate(),
		),
		transporthttp.ResponseEncoder(server.HttpResponseEncoder),
		transporthttp.ErrorEncoder(server.HttpErrorEncoder(nil)),
	}
	if c.Server.Http.Network != "" {
		opts = append(opts, transporthttp.Network(c.Server.Http.Network))
	}
	if c.Server.Http.Host != "" && c.Server.Http.Port != 0 {
		opts = append(opts, transporthttp.Address(fmt.Sprintf("%s:%d", c.Server.Http.Host, c.Server.Http.Port)))
	}
	if c.Server.Http.Timeout != nil {
		opts = append(opts, transporthttp.Timeout(c.Server.Http.Timeout.AsDuration()))
	}
	srv := transporthttp.NewServer(opts...)
	for _, s := range services {
		s.RegisterHttp(srv)
	}
	return srv
}
