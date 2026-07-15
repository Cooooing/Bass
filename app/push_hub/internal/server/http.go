package server

import (
	commonClient "common/pkg/client"
	"common/pkg/constant"
	"common/pkg/server"
	"fmt"
	"log/slog"
	"push_hub/internal/config"

	"github.com/go-kratos/kratos/contrib/middleware/validate/v3"
	"github.com/go-kratos/kratos/v3/middleware/recovery"
	"github.com/go-kratos/kratos/v3/transport/http"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

func NewHTTPServer(c *config.Bootstrap, logger *slog.Logger, obs *commonClient.Observer, services []server.HttpService) *http.Server {
	var opts = []http.ServerOption{
		http.Middleware(
			server.RequestLogContextMiddleware(),
			obs.ServerMiddleware(),
			recovery.Recovery(),
			validate.ProtoValidate(),
		),
	}
	if c.Http.Network != "" {
		opts = append(opts, http.Network(c.Http.Network))
	}
	if c.Http.Host != "" && c.Http.Port != 0 {
		opts = append(opts, http.Address(fmt.Sprintf("%s:%d", c.Http.Host, c.Http.Port)))
	}
	if c.Http.Timeout != nil {
		opts = append(opts, http.Timeout(c.Http.Timeout.AsDuration()))
	}
	srv := http.NewServer(opts...)
	if obsConf := c.GetObservability(); obsConf != nil && obsConf.GetEnableMetrics() {
		srv.Handle("/metrics", promhttp.Handler())
		logger.Info("metrics endpoint registered", slog.String(constant.LogFieldPath, "/metrics"))
	}
	for _, s := range services {
		s.RegisterHttp(srv)
	}
	return srv
}
