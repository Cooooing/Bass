package server

import (
	"common/pkg/server"
	"fmt"
	"github.com/go-kratos/kratos/contrib/middleware/validate/v3"
	"github.com/go-kratos/kratos/v3/middleware/logging"
	"github.com/go-kratos/kratos/v3/middleware/recovery"
	transporthttp "github.com/go-kratos/kratos/v3/transport/http"
	"log/slog"
	"platform/internal/conf"
)

// NewHTTPServer 创建 HTTP 服务，用于接收 OSS 回调。
func NewHTTPServer(c *conf.Bootstrap, logger *slog.Logger, services []server.HttpService) *transporthttp.Server {
	var opts = []transporthttp.ServerOption{
		transporthttp.Middleware(
			server.MetricsMiddleware(c.Server.Name),
			server.TracingMiddleware(c.Server.Name),
			recovery.Recovery(),
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
