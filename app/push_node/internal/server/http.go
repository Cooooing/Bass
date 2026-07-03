package server

import (
	commonClient "common/pkg/client"
	"common/pkg/constant"
	"common/pkg/server"
	"fmt"
	"log/slog"
	"net/http"
	"push_node/internal/biz/usecase"
	"push_node/internal/conf"

	"github.com/go-kratos/kratos/v3/middleware/recovery"
	transporthttp "github.com/go-kratos/kratos/v3/transport/http"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

func NewHTTPServer(
	c *conf.Bootstrap,
	logger *slog.Logger,
	obs *commonClient.Observer,
	services []server.HttpService,
	sseUc *usecase.SSEUsecase,
) *transporthttp.Server {
	var opts = []transporthttp.ServerOption{
		transporthttp.Middleware(
			server.RequestLogContextMiddleware(),
			obs.ServerMiddleware(),
			recovery.Recovery(),
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
	srv.Handle("/push/v1/sse/connect", newSSEHandler(sseUc, logger))
	if obsConf := c.GetObservability(); obsConf != nil && obsConf.GetEnableMetrics() {
		srv.Handle("/metrics", promhttp.Handler())
		logger.Info("metrics endpoint registered", slog.String(constant.LogFieldPath, "/metrics"))
	}
	for _, s := range services {
		s.RegisterHttp(srv)
	}

	return srv
}

func newSSEHandler(sseUc *usecase.SSEUsecase, logger *slog.Logger) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		token := r.URL.Query().Get("token")
		if token == "" {
			http.Error(w, "missing token", http.StatusUnauthorized)
			return
		}

		if _, ok := w.(http.Flusher); !ok {
			http.Error(w, "streaming unsupported", http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "text/event-stream")
		w.Header().Set("Cache-Control", "no-cache")
		w.Header().Set("Connection", "keep-alive")
		w.Header().Set("X-Accel-Buffering", "no")
		w.Header().Set("Access-Control-Allow-Origin", "*")

		if f, ok := w.(http.Flusher); ok {
			f.Flush()
		}

		logger.Info("sse client connecting", constant.LogFieldAddress, r.RemoteAddr)
		sseUc.Connect(r.Context(), token, w)
	}
}
