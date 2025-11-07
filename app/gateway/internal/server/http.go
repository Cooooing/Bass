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
	"net/http"
	"path"
	"strings"

	"github.com/go-kratos/kratos/contrib/middleware/validate/v2"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-kratos/kratos/v2/middleware"
	"github.com/go-kratos/kratos/v2/middleware/logging"
	"github.com/go-kratos/kratos/v2/middleware/metrics"
	"github.com/go-kratos/kratos/v2/middleware/recovery"
	"github.com/go-kratos/kratos/v2/middleware/tracing"
	"github.com/go-kratos/kratos/v2/transport"
	transporthttp "github.com/go-kratos/kratos/v2/transport/http"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

// NewHTTPServer new an HTTP server.
func NewHTTPServer(c *conf.Bootstrap, logger log.Logger, etcdClient *client.EtcdClient, services []service.Service, tokenRepo *util.TokenRepo) *transporthttp.Server {

	var opts = []transporthttp.ServerOption{
		transporthttp.Middleware(
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
	// routes := map[string]string{
	// 	"/api/user":  "127.0.0.1:8001",
	// 	"/api/order": "127.0.0.1:8002",
	// }

	// 代理 handler
	srv.HandlePrefix("/api/user", NewProxyHandler(etcdClient, constant.UserServiceName.String(), "/api/user", logger))
	srv.HandlePrefix("/api/content", NewProxyHandler(etcdClient, constant.ContentServiceName.String(), "/api/content", logger))

	// srv.HandlePrefix("/api", NewProxyHandler(routes, logger))

	for _, s := range services {
		s.RegisterHttp(srv)
	}
	return srv
}

// NewProxyHandler 实现反向代理
func NewProxyHandler(etcdClient *client.EtcdClient, serviceName, prefix string, l log.Logger) http.Handler {
	logger := log.NewHelper(l)

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		conn, err := etcdClient.NewHTTPConn(serviceName)
		r.URL.Path = strings.TrimPrefix(r.URL.Path, prefix)
		response, err := conn.Do(r)

		if err != nil {
			logger.Errorf("proxy error: %v", err)
			http.Error(w, "service unavailable", http.StatusInternalServerError)
			return
		}
		defer response.Body.Close()
		err = response.Write(w)
		if err != nil {
			logger.Errorf("proxy error: %v", err)
			http.Error(w, "service unavailable", http.StatusInternalServerError)
			return
		}

		// u, err := url.Parse(node.Address())
		// if err != nil {
		// 	logger.Errorf("invalid node address: %v", err)
		// 	http.Error(w, "service unavailable", http.StatusInternalServerError)
		// 	return
		// }
		//
		// targetURL := &url.URL{Scheme: u.Scheme, Host: u.Host}
		// proxy := httputil.NewSingleHostReverseProxy(targetURL)
		//
		// originalDirector := proxy.Director
		// proxy.Director = func(req *http.Request) {
		// 	originalDirector(req)
		// 	// 去掉前缀
		// 	req.URL.Path = strings.TrimPrefix(req.URL.Path, prefix)
		// 	if req.URL.Path == "" {
		// 		req.URL.Path = "/"
		// 	}
		// 	// 保留原始 query
		// 	req.URL.RawPath = req.URL.EscapedPath()
		// 	req.Host = u.Host
		// 	log.Infof("proxy request: %s %s %s", req.Method, req.URL.Path)
		// }
		//
		// proxy.ErrorHandler = func(w http.ResponseWriter, req *http.Request, err error) {
		// 	log.Errorf("proxy error: %v", err)
		// 	http.Error(w, "proxy error", http.StatusBadGateway)
		// }
		//
		// proxy.ServeHTTP(w, r)
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
