package server

import (
	commonClient "common/pkg/client"
	"common/pkg/constant"
	"common/pkg/server"
	servermiddleware "common/pkg/server/middleware"
	"fmt"
	"log/slog"

	"game_idle/internal/config"

	"github.com/go-kratos/kratos/contrib/middleware/validate/v3"
	"github.com/go-kratos/kratos/v3/middleware/recovery"
	"github.com/go-kratos/kratos/v3/transport/http"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

func NewHTTPServer(
	c *config.Bootstrap,
	logger *slog.Logger,
	obs *commonClient.Observer,
	services []server.Service,
) *http.Server {
	var opts = []http.ServerOption{
		http.Middleware(
			servermiddleware.RequestLogContextMiddleware(),
			obs.ServerMiddleware(),
			recovery.Recovery(),
			validate.ProtoValidate(),
		),
	}
	if c.GetHttp().GetNetwork() != "" {
		opts = append(opts, http.Network(c.GetHttp().GetNetwork()))
	}
	if c.GetHttp().GetHost() != "" && c.GetHttp().GetPort() != 0 {
		opts = append(opts, http.Address(fmt.Sprintf("%s:%d", c.GetHttp().GetHost(), c.GetHttp().GetPort())))
	}
	if c.GetHttp().GetTimeout() != nil {
		opts = append(opts, http.Timeout(c.GetHttp().GetTimeout().AsDuration()))
	}
	srv := http.NewServer(opts...)
	if obsConf := c.GetObservability(); obsConf != nil && obsConf.GetEnableMetrics() {
		srv.Handle("/metrics", promhttp.Handler())
		logger.Info("metrics endpoint registered", slog.String(constant.LogFieldPath, "/metrics"))
	}
	for _, s := range services {
		s.RegisterHTTP(srv)
	}
	return srv
}
