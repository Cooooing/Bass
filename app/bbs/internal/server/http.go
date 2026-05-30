package server

import (
	"bbs/internal/conf"
	bbsuserv1 "common/api/gen/bbs/v1/user"
	userv1 "common/api/gen/user/v1"
	"common/pkg/util/server"
	"context"
	"fmt"

	"github.com/go-kratos/kratos/contrib/middleware/validate/v2"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-kratos/kratos/v2/middleware/logging"
	"github.com/go-kratos/kratos/v2/middleware/metrics"
	"github.com/go-kratos/kratos/v2/middleware/recovery"
	"github.com/go-kratos/kratos/v2/middleware/selector"
	"github.com/go-kratos/kratos/v2/middleware/tracing"
	"github.com/go-kratos/kratos/v2/transport/http"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

var bbsHTTPAuthOperationWhitelist = map[string]struct{}{
	bbsuserv1.OperationAuthServiceStartEmailRegistration:  {},
	bbsuserv1.OperationAuthServiceVerifyEmailRegistration: {},
	bbsuserv1.OperationAuthServiceStartPhoneRegistration:  {},
	bbsuserv1.OperationAuthServiceVerifyPhoneRegistration: {},
	bbsuserv1.OperationAuthServiceLoginByPassword:         {},
	bbsuserv1.OperationAccountServiceAvatar:               {},
}

// NewHTTPServer 创建 HTTP 服务。
func NewHTTPServer(c *conf.Bootstrap, logger log.Logger, services []server.HttpService, authClient userv1.AuthServiceClient) *http.Server {
	authRequiredMatch := func(_ context.Context, operation string) bool {
		_, ok := bbsHTTPAuthOperationWhitelist[operation]
		return !ok
	}

	var opts = []http.ServerOption{
		http.Middleware(
			recovery.Recovery(),
			tracing.Server(),
			metrics.Server(
				metrics.WithSeconds(_metricSeconds),
				metrics.WithRequests(_metricRequests),
			),
			logging.Server(logger),
			selector.Server(server.UserAuthMiddleware(authClient)).Match(authRequiredMatch).Build(),
			validate.ProtoValidate(),
		),
		http.ResponseEncoder(server.HttpResponseEncoder),
		// http.ErrorEncoder(pkg.HttpErrorEncoder),
	}
	if c.Server.Http.Network != "" {
		opts = append(opts, http.Network(c.Server.Http.Network))
	}
	if c.Server.Http.Host != "" && c.Server.Http.Port != 0 {
		opts = append(opts, http.Address(fmt.Sprintf("%s:%d", c.Server.Http.Host, c.Server.Http.Port)))
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
