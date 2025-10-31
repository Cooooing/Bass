package server

import (
	"common/pkg"
	"common/pkg/constant"
	"context"
	"errors"
	"path"
	"strings"
	"user/internal/biz"
	"user/internal/conf"
	"user/internal/service"

	"github.com/go-kratos/kratos/contrib/middleware/validate/v2"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-kratos/kratos/v2/middleware"
	"github.com/go-kratos/kratos/v2/middleware/logging"
	"github.com/go-kratos/kratos/v2/middleware/metrics"
	"github.com/go-kratos/kratos/v2/middleware/recovery"
	"github.com/go-kratos/kratos/v2/transport"
	"github.com/go-kratos/kratos/v2/transport/http"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

// NewHTTPServer new an HTTP server.
func NewHTTPServer(c *conf.Bootstrap, logger log.Logger, services []service.Service, tokenService *biz.TokenService) *http.Server {
	var opts = []http.ServerOption{
		http.Middleware(
			recovery.Recovery(),
			logging.Server(logger),
			metrics.Server(
				metrics.WithSeconds(_metricSeconds),
				metrics.WithRequests(_metricRequests),
			),
			AuthMiddleware(tokenService),
			validate.ProtoValidate(),
		),
		http.ResponseEncoder(pkg.HttpResponseEncoder),
		http.ErrorEncoder(pkg.HttpErrorEncoder),
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
func AuthMiddleware(tokenService *biz.TokenService) middleware.Middleware {
	return func(handler middleware.Handler) middleware.Handler {
		return func(ctx context.Context, req interface{}) (interface{}, error) {
			tr, ok := transport.FromServerContext(ctx)
			if !ok {
				return nil, errors.New("transport not found")
			}

			log.Infof(tr.Operation())
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
			userInfo, err := tokenService.TokenGen.Parse(token)
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
	"/user.v1.Authentication/*": {},
}
