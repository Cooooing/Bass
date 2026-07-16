package server

import (
	commonClient "common/pkg/client"
	"common/pkg/constant"
	"common/pkg/server"
	"fmt"
	"game_town/internal/config"
	"log/slog"

	"github.com/go-kratos/kratos/v3/middleware/recovery"
	"github.com/go-kratos/kratos/v3/transport/http"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

func NewHTTPServer(c *config.Bootstrap, logger *slog.Logger, obs *commonClient.Observer, services []server.Service) *http.Server {
	serverOpts := []http.ServerOption{http.Middleware(obs.ServerMiddleware(), recovery.Recovery())}
	if c.GetHttp().GetHost() != "" && c.GetHttp().GetPort() != 0 {
		serverOpts = append(serverOpts, http.Address(fmt.Sprintf("%s:%d", c.GetHttp().GetHost(), c.GetHttp().GetPort())))
	}
	if c.GetHttp().GetTimeout() != nil {
		serverOpts = append(serverOpts, http.Timeout(c.GetHttp().GetTimeout().AsDuration()))
	}
	srv := http.NewServer(serverOpts...)
	if c.GetObservability().GetEnableMetrics() {
		srv.Handle("/metrics", promhttp.Handler())
		logger.Info("metrics endpoint registered", slog.String(constant.LogFieldPath, "/metrics"))
	}
	for _, s := range services {
		s.RegisterHttp(srv)
	}
	return srv
}
