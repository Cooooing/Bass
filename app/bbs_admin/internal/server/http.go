package server

import (
	"bbs_admin/internal/config"
	commonClient "common/pkg/client"
	"common/pkg/constant"
	commonserver "common/pkg/server"
	"fmt"
	"log/slog"

	"github.com/go-kratos/kratos/contrib/middleware/validate/v3"
	"github.com/go-kratos/kratos/v3/middleware/recovery"
	transporthttp "github.com/go-kratos/kratos/v3/transport/http"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

func NewHTTPServer(c *config.Bootstrap, logger *slog.Logger, obs *commonClient.Observer, services []commonserver.Service) *transporthttp.Server {
	opts := []transporthttp.ServerOption{
		transporthttp.Middleware(
			commonserver.RequestLogContextMiddleware(),
			obs.ServerMiddleware(),
			recovery.Recovery(),
			validate.ProtoValidate(),
		),
		transporthttp.ResponseEncoder(commonserver.HttpRespEncoder),
		transporthttp.ErrorEncoder(commonserver.HttpErrorEncoder(nil)),
	}
	if c.GetHttp().GetNetwork() != "" {
		opts = append(opts, transporthttp.Network(c.GetHttp().GetNetwork()))
	}
	if c.GetHttp().GetHost() != "" && c.GetHttp().GetPort() != 0 {
		opts = append(opts, transporthttp.Address(fmt.Sprintf("%s:%d", c.GetHttp().GetHost(), c.GetHttp().GetPort())))
	}
	if c.GetHttp().GetTimeout() != nil {
		opts = append(opts, transporthttp.Timeout(c.GetHttp().GetTimeout().AsDuration()))
	}
	srv := transporthttp.NewServer(opts...)
	if obsConf := c.GetObservability(); obsConf != nil && obsConf.GetEnableMetrics() {
		srv.Handle("/metrics", promhttp.Handler())
		logger.Info("metrics endpoint registered", slog.String(constant.LogFieldPath, "/metrics"))
	}
	for _, s := range services {
		s.RegisterHttp(srv)
	}
	return srv
}
