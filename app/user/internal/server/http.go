package server

import (
	"common/pkg"
	"common/pkg/constant"
	"common/pkg/util"
	"context"
	"errors"
	"strings"
	"user/internal/conf"
	"user/internal/service"

	"github.com/go-kratos/kratos/contrib/middleware/validate/v2"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-kratos/kratos/v2/middleware"
	"github.com/go-kratos/kratos/v2/middleware/logging"
	"github.com/go-kratos/kratos/v2/middleware/metrics"
	"github.com/go-kratos/kratos/v2/middleware/recovery"
	"github.com/go-kratos/kratos/v2/middleware/tracing"
	"github.com/go-kratos/kratos/v2/transport"
	"github.com/go-kratos/kratos/v2/transport/http"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

// NewHTTPServer new an HTTP server.
func NewHTTPServer(c *conf.Bootstrap, logger log.Logger, services []service.Service, tokenRepo *util.TokenRepo) *http.Server {
	var opts = []http.ServerOption{
		http.Middleware(
			recovery.Recovery(),
			tracing.Server(),
			metrics.Server(
				metrics.WithSeconds(_metricSeconds),
				metrics.WithRequests(_metricRequests),
			),
			logging.Server(logger),
			AuthMiddleware(tokenRepo),
			validate.ProtoValidate(),
		),
		http.ResponseEncoder(pkg.HttpResponseEncoder),
		// http.ErrorEncoder(pkg.HttpErrorEncoder),
	}
	if c.Server.Http.Network != "" {
		opts = append(opts, http.Network(c.Server.Http.Network))
	}
	if c.Server.Http.Addr != "" {
		opts = append(opts, http.Address(c.Server.Http.Addr))
	}
	if c.Server.Http.Timeout != nil {
		opts = append(opts, http.Timeout(c.Server.Http.Timeout.AsDuration()))
	}
	srv := http.NewServer(opts...)
	srv.Handle("/metrics", promhttp.Handler())
	for _, s := range services {
		s.RegisterHttp(srv)
	}
	return srv
}

// AuthMiddleware 返回一个 Kratos 中间件，用于认证
func AuthMiddleware(tokenRepo *util.TokenRepo) middleware.Middleware {
	return func(handler middleware.Handler) middleware.Handler {
		return func(ctx context.Context, req interface{}) (interface{}, error) {
			tr, ok := transport.FromServerContext(ctx)
			if !ok {
				return nil, errors.New("transport not found")
			}

			// 获取 token
			token := strings.TrimPrefix(tr.RequestHeader().Get(constant.Authentication), "Bearer ")
			// 验证 token
			userInfo, err := tokenRepo.GetToken(ctx, token)
			if err != nil {
				return nil, err
			}
			// 设置上下文
			ctx = context.WithValue(ctx, constant.UserInfo, userInfo)
			return handler(ctx, req)
		}
	}
}
