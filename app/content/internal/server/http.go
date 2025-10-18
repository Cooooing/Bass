package server

import (
	"common/pkg"
	"content/internal/conf"
	"content/internal/service"

	"github.com/go-kratos/kratos/contrib/middleware/validate/v2"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-kratos/kratos/v2/middleware/logging"
	"github.com/go-kratos/kratos/v2/middleware/metrics"
	"github.com/go-kratos/kratos/v2/middleware/recovery"
	"github.com/go-kratos/kratos/v2/transport/http"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

// NewHTTPServer new an HTTP server.
func NewHTTPServer(c *conf.Bootstrap, logger log.Logger, services []service.Service) *http.Server {
	var opts = []http.ServerOption{
		http.Middleware(
			recovery.Recovery(),
			logging.Server(logger),
			metrics.Server(
				metrics.WithSeconds(_metricSeconds),
				metrics.WithRequests(_metricRequests),
			),
			validate.ProtoValidate(),
		),
		http.ResponseEncoder(pkg.HttpResponseEncoder),
		http.ErrorEncoder(pkg.HttpErrorEncoder),
	}
	if c.Server.Http.Network != "" {
		opts = append(opts, http.Network(c.Server.Http.Network))
	}
	if c.Server.Http.Addr != "" {
		opts = append(opts, http.Address(c.Server.Http.Addr))
	}
	if c.Server.Http.Timeout != nil {
		opts = append(opts, http.Timeout(c.Server.Http.Timeout.AsDuration()))
	}
	srv := http.NewServer(opts...)
	srv.Handle("/metrics", promhttp.Handler())
	for _, s := range services {
		s.RegisterHttp(srv)
	}
	return srv
}
