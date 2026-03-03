package server

import (
	"common/pkg/constant"
	"common/pkg/util/jwt"
	"common/pkg/util/server"
	"fmt"
	"infra/internal/conf"
	"infra/internal/service"
	http2 "net/http"
	"os"

	"github.com/go-kratos/kratos/contrib/middleware/validate/v2"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-kratos/kratos/v2/middleware/logging"
	"github.com/go-kratos/kratos/v2/middleware/metrics"
	"github.com/go-kratos/kratos/v2/middleware/recovery"
	"github.com/go-kratos/kratos/v2/middleware/tracing"
	"github.com/go-kratos/kratos/v2/transport/http"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

// NewHTTPServer new an HTTP server.
func NewHTTPServer(c *conf.Bootstrap, logger log.Logger, services []service.Service, tokenCache *jwt.TokenCache) *http.Server {
	var opts = []http.ServerOption{
		http.Middleware(
			recovery.Recovery(),
			tracing.Server(),
			metrics.Server(
				metrics.WithSeconds(_metricSeconds),
				metrics.WithRequests(_metricRequests),
			),
			logging.Server(logger),
			server.AuthMiddleware(tokenCache),
			validate.ProtoValidate(),
		),
		http.ResponseEncoder(server.HttpResponseEncoder),
		// http.ErrorEncoder(pkg.HttpErrorEncoder),
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
	srv.Handle("/metrics", promhttp.Handler())
	// 接口文档
	if c.Server.Mode != constant.Prod {
		srv.Handle("/openapi.yaml", DocHandler())
	}
	for _, s := range services {
		s.RegisterHttp(srv)
	}
	return srv
}

func DocHandler() http2.Handler {
	return http2.HandlerFunc(func(w http2.ResponseWriter, r *http2.Request) {
		// 只允许 GET
		if r.Method != http2.MethodGet {
			http2.Error(w, "method not allowed", http2.StatusMethodNotAllowed)
			return
		}

		data, err := os.ReadFile("../../common/api/openapi.yaml")
		if err != nil {
			http2.Error(w, "swagger file not found", http2.StatusNotFound)
			return
		}

		w.Header().Set("Content-Type", "text/plain; charset=utf-8")
		w.WriteHeader(http2.StatusOK)
		_, _ = w.Write(data)
	})
}
