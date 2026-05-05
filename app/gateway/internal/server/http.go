package server

import (
	"common/pkg/client"
	"common/pkg/constant"
	commonModel "common/pkg/model"
	"common/pkg/util"
	"common/pkg/util/jwt"
	"common/pkg/util/server"

	"context"
	"errors"
	"fmt"
	"gateway/internal/biz/domain"
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
	"github.com/go-kratos/kratos/v2/middleware/metadata"
	"github.com/go-kratos/kratos/v2/middleware/metrics"
	"github.com/go-kratos/kratos/v2/middleware/recovery"
	"github.com/go-kratos/kratos/v2/middleware/selector"
	"github.com/go-kratos/kratos/v2/middleware/tracing"
	transporthttp "github.com/go-kratos/kratos/v2/transport/http"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"go.opentelemetry.io/otel/propagation"
)

// NewHTTPServer new an HTTP server.
func NewHTTPServer(c *conf.Bootstrap, logger log.Logger, consulClient *client.ConsulClient, services []service.Service, tokenCache *jwt.TokenCache, ipDomain *domain.IpDomain) *transporthttp.Server {

	middlewares := []middleware.Middleware{
		metadata.Server(),
		recovery.Recovery(),
		tracing.Server(),
		metrics.Server(
			metrics.WithSeconds(_metricSeconds),
			metrics.WithRequests(_metricRequests),
		),
		logging.Server(logger),

		IpMiddleware(ipDomain),
		// 七牛回调验签
		selector.Server(QiniuCallbackSignMiddleware(c)).Match(QiniuCallbackMatch()).Build(),
		// 认证鉴权
		selector.Server(AuthMiddleware(tokenCache), PermissionMiddleware()).Match(UserAPIMatch()).Build(),

		validate.ProtoValidate(),
	}
	var opts = []transporthttp.ServerOption{
		transporthttp.Middleware(middlewares...),
		transporthttp.ResponseEncoder(server.HttpResponseEncoder),
		transporthttp.ErrorEncoder(server.HttpErrorEncoder),
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
	srv.HandlePrefix("/api/notify", NewProxyHandler(middlewares, consulClient, constant.MsgCenterServiceName.String(), "/api/notify", logger))
	srv.HandlePrefix("/api/im", NewProxyHandler(middlewares, consulClient, constant.IMServiceName.String(), "/api/im", logger))
	srv.HandlePrefix("/api/signal", NewProxyHandler(middlewares, consulClient, constant.SignalServiceName.String(), "/api/signal", logger))

	for _, s := range services {
		s.RegisterHttp(srv)
	}
	return srv
}

// NewProxyHandler 实现反向代理
func NewProxyHandler(middlewares []middleware.Middleware, consulClient *client.ConsulClient, serviceName, prefix string, l log.Logger) http.Handler {
	logger := log.NewHelper(l)
	propagator := propagation.TraceContext{}
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		h := middleware.Chain(middlewares...)(func(ctx context.Context, req interface{}) (interface{}, error) {
			http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				conn, err := consulClient.GetHTTPClient(serviceName)
				if err != nil {
					logger.Errorf("new http conn error: %v", err)
					server.HttpErrorEncoder(w, r, errors2.New(500, "Internal Server Error", "Internal Server Error"))
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

				ip, ok := util.GetContextValue[*commonModel.IpInfo](ctx, constant.CtxIpInfo)
				if !ok {
					logger.Errorf("proxy error: %v", "ip info not found")
					server.HttpErrorEncoder(w, r, errors2.New(500, "Internal Server Error", "Internal Server Error"))
					return
				}
				logger.Infof(
					"proxy ip:%s [%s-%s-%s] -> service=%s method=%s original_path=%s headers=%v",
					ip.Ip,
					ip.Country,
					ip.Province,
					ip.City,
					serviceName,
					r.Method,
					originalPath,
					r.Header,
				)

				response, err := conn.Do(r)

				if err != nil {
					if e, ok2 := errors.AsType[*errors2.Error](err); ok2 {
						server.HttpErrorEncoder(w, r, errors2.New(int(e.Code), e.Reason, e.Message))
						return
					}
					logger.Errorf("proxy error: %v", err)
					server.HttpErrorEncoder(w, r, errors2.New(500, "Internal Server Error", "Internal Server Error"))
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
					server.HttpErrorEncoder(w, r, errors2.New(500, "Internal Server Error", "Internal Server Error"))
					return
				}
			})(w, r)
			return nil, nil
		})
		ctx := r.Context()
		propagator.Inject(ctx, propagation.HeaderCarrier(r.Header))
		_, err := h(ctx, nil)
		if err != nil {
			// 捕获中间件错误，例如 AuthMiddleware 返回的 401
			if e, ok := errors.AsType[*errors2.Error](err); ok {
				server.HttpErrorEncoder(w, r, e)
			} else {
				server.HttpErrorEncoder(w, r, errors2.New(500, "Internal Server Error", err.Error()))
			}
		}
	})
}
