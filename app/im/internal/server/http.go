package server

import (
	"common/pkg/server"
	"fmt"
	"github.com/go-kratos/kratos/contrib/middleware/validate/v3"
	"github.com/go-kratos/kratos/v3/middleware/logging"
	"github.com/go-kratos/kratos/v3/middleware/recovery"
	"github.com/go-kratos/kratos/v3/transport/http"
	"im/internal/conf"
	"log/slog"
)

// NewHTTPServer 创建 HTTP 服务。
func NewHTTPServer(c *conf.Bootstrap, logger *slog.Logger, services []server.HttpService) *http.Server {
	var opts = []http.ServerOption{
		http.Middleware(
			server.MetricsMiddleware(c.Server.Name),
			server.TracingMiddleware(c.Server.Name),
			recovery.Recovery(),
			logging.Server(logger),
			validate.ProtoValidate(),
		),
	}
	if c.Server.Http.Network != "" {
		opts = append(opts, http.Network(c.Server.Http.Network))
	}
	if c.Server.Http.Host != "" && c.Server.Http.Port != 0 {
		opts = append(opts, http.Address(fmt.Sprintf("%s:%d", c.Server.Http.Host, c.Server.Http.Port)))
	}
	if c.Server.Http.Timeout != nil {
		opts = append(opts, http.Timeout(c.Server.Http.Timeout.AsDuration()))
	}
	srv := http.NewServer(opts...)
	for _, s := range services {
		s.RegisterHttp(srv)
	}
	return srv
}
