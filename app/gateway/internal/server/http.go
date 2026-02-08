package server

import (
	"common/pkg"
	"common/pkg/client"
	"common/pkg/constant"
	"common/pkg/util"
	"context"
	"errors"
	"fmt"
	"gateway/internal/conf"
	"gateway/internal/service"
	"io"
	"net/http"
	"strings"

	"github.com/go-kratos/kratos/contrib/middleware/validate/v2"
	errors2 "github.com/go-kratos/kratos/v2/errors"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-kratos/kratos/v2/middleware"
	"github.com/go-kratos/kratos/v2/middleware/logging"
	"github.com/go-kratos/kratos/v2/middleware/metrics"
	"github.com/go-kratos/kratos/v2/middleware/recovery"
	"github.com/go-kratos/kratos/v2/middleware/selector"
	"github.com/go-kratos/kratos/v2/middleware/tracing"
	transporthttp "github.com/go-kratos/kratos/v2/transport/http"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"go.opentelemetry.io/otel/propagation"
)

// NewHTTPServer new an HTTP server.
func NewHTTPServer(c *conf.Bootstrap, logger log.Logger, consulClient *client.ConsulClient, services []service.Service, tokenCache *util.TokenCache) *transporthttp.Server {

	middlewares := []middleware.Middleware{
		recovery.Recovery(),
		tracing.Server(),
		metrics.Server(
			metrics.WithSeconds(_metricSeconds),
			metrics.WithRequests(_metricRequests),
		),
		logging.Server(logger),

		// 七牛回调验签
		selector.Server(QiniuCallbackSignMiddleware(c)).Match(QiniuCallbackMatch()).Build(),
		// 认证鉴权
		selector.Server(AuthMiddleware(tokenCache), PermissionMiddleware()).Match(UserAPIMatch()).Build(),

		validate.ProtoValidate(),
	}
	var opts = []transporthttp.ServerOption{
		transporthttp.Middleware(middlewares...),
		transporthttp.ResponseEncoder(pkg.HttpResponseEncoder),
		transporthttp.ErrorEncoder(pkg.HttpErrorEncoder),
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
	srv.Handle("/metrics", promhttp.Handler())
	// 代理 handler
	srv.HandlePrefix("/api/user", NewProxyHandler(middlewares, consulClient, constant.UserServiceName.String(), "/api/user", logger))
	srv.HandlePrefix("/api/content", NewProxyHandler(middlewares, consulClient, constant.ContentServiceName.String(), "/api/content", logger))
	srv.HandlePrefix("/api/notify", NewProxyHandler(middlewares, consulClient, constant.NotifyServiceName.String(), "/api/notify", logger))
	srv.HandlePrefix("/api/im", NewProxyHandler(middlewares, consulClient, constant.IMServiceName.String(), "/api/im", logger))
	srv.HandlePrefix("/api/signal", NewProxyHandler(middlewares, consulClient, constant.SignalServiceName.String(), "/api/signal", logger))
	srv.HandlePrefix("/api/infra", NewProxyHandler(middlewares, consulClient, constant.InfraServiceName.String(), "/api/infra", logger))

	for _, s := range services {
		s.RegisterHttp(srv)
	}
	return srv
}

// NewProxyHandler 实现反向代理
func NewProxyHandler(middlewares []middleware.Middleware, consulClient *client.ConsulClient, serviceName, prefix string, l log.Logger) http.Handler {
	logger := log.NewHelper(l)
	propagator := propagation.TraceContext{}
	handlerFunc := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		conn, err := consulClient.GetHTTPClient(serviceName)
		if err != nil {
			logger.Errorf("new http conn error: %v", err)
			pkg.HttpErrorEncoder(w, r, errors2.New(500, "Internal Server Error", "Internal Server Error"))
			return
		}
		originalPath := r.URL.Path
		r.URL.Path = strings.TrimPrefix(r.URL.Path, prefix)
		r.RequestURI = ""
		r.Header.Set("Host", r.Host)

		proxiedPath := r.URL.Path
		if r.URL.RawQuery != "" {
			proxiedPath += "?" + r.URL.RawQuery
		}

		logger.Infof(
			"proxy -> service=%s method=%s original_path=%s headers=%v",
			serviceName,
			r.Method,
			originalPath,
			r.Header,
		)

		response, err := conn.Do(r)

		if err != nil {
			var e *errors2.Error
			if errors.As(err, &e) {
				pkg.HttpErrorEncoder(w, r, errors2.New(int(e.Code), e.Reason, e.Message))
				return
			}
			logger.Errorf("proxy error: %v", err)
			pkg.HttpErrorEncoder(w, r, errors2.New(500, "Internal Server Error", "Internal Server Error"))
			return
		}
		defer func(Body io.ReadCloser) {
			_ = Body.Close()
		}(response.Body)
		// 复制后端 Header（过滤 hop-by-hop headers）
		for k, vv := range response.Header {
			if strings.EqualFold(k, "Connection") ||
				strings.EqualFold(k, "Proxy-Connection") ||
				strings.EqualFold(k, "Keep-Alive") ||
				strings.EqualFold(k, "Transfer-Encoding") ||
				strings.EqualFold(k, "TE") ||
				strings.EqualFold(k, "Trailer") ||
				strings.EqualFold(k, "Upgrade") {
				continue
			}
			for _, v := range vv {
				w.Header().Add(k, v)
			}
		}
		w.Header().Add("Access-Control-Allow-Origin", "*") // 允许跨域
		w.WriteHeader(response.StatusCode)
		_, err = io.Copy(w, response.Body)
		if err != nil {
			logger.Errorf("proxy error: %v", err)
			pkg.HttpErrorEncoder(w, r, errors2.New(500, "Internal Server Error", "Internal Server Error"))
			return
		}
	})
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		h := middleware.Chain(middlewares...)(func(ctx context.Context, req interface{}) (interface{}, error) {
			handlerFunc(w, r)
			return nil, nil
		})
		ctx := r.Context()
		propagator.Inject(ctx, propagation.HeaderCarrier(r.Header))
		_, err := h(ctx, nil)
		if err != nil {
			// 捕获中间件错误，例如 AuthMiddleware 返回的 401
			var e *errors2.Error
			if errors.As(err, &e) {
				pkg.HttpErrorEncoder(w, r, e)
			} else {
				pkg.HttpErrorEncoder(w, r, errors2.New(500, "Internal Server Error", err.Error()))
			}
		}
	})
}
