package server

import (
	"common/pkg/server"
	"common/pkg/util"
	"fmt"
	"net/http"

	"github.com/go-kratos/kratos/v3/middleware/logging"
	"github.com/go-kratos/kratos/v3/middleware/recovery"
	transporthttp "github.com/go-kratos/kratos/v3/transport/http"
	"log/slog"
	"push_node/internal/biz/usecase"
	"push_node/internal/conf"
)

// NewHTTPServer 创建 HTTP 服务，注册 SSE 和健康检查端点。
func NewHTTPServer(
	c *conf.Bootstrap,
	logger *slog.Logger,
	services []server.HttpService,
	sseUc *usecase.SSEUsecase,
) *transporthttp.Server {
	var opts = []transporthttp.ServerOption{
		transporthttp.Middleware(
			server.MetricsMiddleware(c.Server.Name),
			server.TracingMiddleware(c.Server.Name),
			recovery.Recovery(),
			logging.Server(logger),
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

	// 注册 SSE 端点（原始 HTTP handler，不走 proto）
	srv.Handle("/push/v1/sse/connect", newSSEHandler(sseUc, logger))

	// 注册 HTTP 服务（健康检查等）
	for _, s := range services {
		s.RegisterHttp(srv)
	}

	return srv
}

// newSSEHandler 创建 SSE 连接的原始 HTTP handler。
func newSSEHandler(sseUc *usecase.SSEUsecase, logger *slog.Logger) http.HandlerFunc {
	helper := util.NewLogHelper(logger)
	return func(w http.ResponseWriter, r *http.Request) {
		// 获取 query 中的 token
		token := r.URL.Query().Get("token")
		if token == "" {
			http.Error(w, "缺少 token 参数", http.StatusUnauthorized)
			return
		}

		// 检查是否支持 Flusher
		if _, ok := w.(http.Flusher); !ok {
			http.Error(w, "服务端不支持流式响应", http.StatusInternalServerError)
			return
		}

		// 设置 SSE 响应头
		w.Header().Set("Content-Type", "text/event-stream")
		w.Header().Set("Cache-Control", "no-cache")
		w.Header().Set("Connection", "keep-alive")
		w.Header().Set("X-Accel-Buffering", "no")
		w.Header().Set("Access-Control-Allow-Origin", "*")

		// 刷新头部，让客户端立即收到
		if f, ok := w.(http.Flusher); ok {
			f.Flush()
		}

		helper.Infof("SSE 客户端正在连接: remote=%s", r.RemoteAddr)

		// 委托给 usecase 处理连接生命周期
		sseUc.Connect(r.Context(), token, w)
	}
}
