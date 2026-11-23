package server

import (
	commonclient "common/pkg/client"
	"common/pkg/constant"
	commonserver "common/pkg/server"
	servermiddleware "common/pkg/server/middleware"
	"fmt"
	"log/slog"

	"realtime_hub/internal/config"

	"github.com/go-kratos/kratos/contrib/middleware/validate/v3"
	"github.com/go-kratos/kratos/v3/middleware/recovery"
	"github.com/go-kratos/kratos/v3/transport/http"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

func NewHTTPServer(
	config *config.Bootstrap,
	logger *slog.Logger,
	observer *commonclient.Observer,
	services []commonserver.Service,
) *http.Server {
	options := []http.ServerOption{
		http.Middleware(
			servermiddleware.RequestLogContextMiddleware(),
			observer.ServerMiddleware(),
			recovery.Recovery(),
			validate.ProtoValidate(),
		),
	}
	if config.GetHttp().GetNetwork() != "" {
		options = append(options, http.Network(config.GetHttp().GetNetwork()))
	}
	if config.GetHttp().GetHost() != "" && config.GetHttp().GetPort() != 0 {
		options = append(options, http.Address(fmt.Sprintf("%s:%d", config.GetHttp().GetHost(), config.GetHttp().GetPort())))
	}
	if config.GetHttp().GetTimeout() != nil {
		options = append(options, http.Timeout(config.GetHttp().GetTimeout().AsDuration()))
	}
	server := http.NewServer(options...)
	if config.GetObservability().GetEnableMetrics() {
		server.Handle("/metrics", promhttp.Handler())
		logger.Info("metrics endpoint registered", slog.String(constant.LogFieldPath, "/metrics"))
	}
	for _, service := range services {
		service.RegisterHTTP(server)
	}
	return server
}
