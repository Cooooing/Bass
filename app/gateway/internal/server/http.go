package server

import (
	"common/pkg"
	"common/pkg/client"
	"common/pkg/constant"
	"common/pkg/util"
	"context"
	"errors"
	"gateway/internal/conf"
	"gateway/internal/service"
	"io"
	"net/http"
	"path"
	"strings"

	"github.com/go-kratos/kratos/contrib/middleware/validate/v2"
	errors2 "github.com/go-kratos/kratos/v2/errors"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-kratos/kratos/v2/middleware"
	"github.com/go-kratos/kratos/v2/middleware/logging"
	"github.com/go-kratos/kratos/v2/middleware/metrics"
	"github.com/go-kratos/kratos/v2/middleware/recovery"
	"github.com/go-kratos/kratos/v2/middleware/tracing"
	"github.com/go-kratos/kratos/v2/transport"
	transporthttp "github.com/go-kratos/kratos/v2/transport/http"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"go.opentelemetry.io/otel/propagation"
)

// NewHTTPServer new an HTTP server.
func NewHTTPServer(c *conf.Bootstrap, logger log.Logger, etcdClient *client.EtcdClient, services []service.Service, tokenRepo *util.TokenRepo) *transporthttp.Server {

	middlewares := []middleware.Middleware{
		recovery.Recovery(),
		tracing.Server(),
		metrics.Server(
			metrics.WithSeconds(_metricSeconds),
			metrics.WithRequests(_metricRequests),
		),
		logging.Server(logger),
		AuthMiddleware(tokenRepo),
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
	if c.Server.Http.Addr != "" {
		opts = append(opts, transporthttp.Address(c.Server.Http.Addr))
	}
	if c.Server.Http.Timeout != nil {
		opts = append(opts, transporthttp.Timeout(c.Server.Http.Timeout.AsDuration()))
	}
	srv := transporthttp.NewServer(opts...)
	srv.Handle("/metrics", promhttp.Handler())
	// 代理 handler
	srv.HandlePrefix("/user", NewProxyHandler(middlewares, etcdClient, constant.UserServiceName.String(), "/user", logger))
	srv.HandlePrefix("/content", NewProxyHandler(middlewares, etcdClient, constant.ContentServiceName.String(), "/content", logger))

	for _, s := range services {
		s.RegisterHttp(srv)
	}
	return srv
}

// NewProxyHandler 实现反向代理
func NewProxyHandler(middlewares []middleware.Middleware, etcdClient *client.EtcdClient, serviceName, prefix string, l log.Logger) http.Handler {
	logger := log.NewHelper(l)
	propagator := propagation.TraceContext{}
	handlerFunc := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		logger.Infof("proxy -> %s headers: %v", serviceName, r.Header)
		conn, err := etcdClient.NewHTTPConn(serviceName)
		r.URL.Path = strings.TrimPrefix(r.URL.Path, prefix)
		r.RequestURI = ""
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
		defer response.Body.Close()
		w.Header().Set("Content-Type", response.Header.Get("Content-Type"))
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
			// 用 middleware 提供的 ctx 注入 trace header 到下游请求头
			propagator.Inject(ctx, propagation.HeaderCarrier(r.Header))
			handlerFunc(w, r)
			return nil, nil
		})
		_, _ = h(r.Context(), nil)
	})
}

// AuthMiddleware 返回一个 Kratos 中间件，用于认证
func AuthMiddleware(tokenRepo *util.TokenRepo) middleware.Middleware {
	return func(handler middleware.Handler) middleware.Handler {
		return func(ctx context.Context, req interface{}) (interface{}, error) {
			tr, ok := transport.FromServerContext(ctx)
			if !ok {
				return nil, errors.New("transport not found")
			}

			// 是否需要鉴权 Todo 用户组权限规则后续持久化入库
			var allow bool
			for pattern := range NoAuthEndpoints {
				match, err := path.Match(pattern, tr.Operation())
				if err == nil && match {
					allow = true
				}
			}
			if allow {
				return handler(ctx, req)
			}

			// 获取 token
			token := strings.TrimPrefix(tr.RequestHeader().Get(constant.Authentication), "Bearer ")

			// 验证 token
			userInfo, err := tokenRepo.GetToken(ctx, token)
			if err != nil {
				return nil, err
			}

			// 权限范围 Todo 用户组权限规则后续持久化入库

			// 设置上下文
			ctx = context.WithValue(ctx, constant.UserInfo, userInfo)

			return handler(ctx, req)
		}
	}
}

var NoAuthEndpoints = map[string]struct{}{
	"/common.api.common.v1.System/Health": {},
}
