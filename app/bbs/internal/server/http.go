package server

import (
	"bbs/internal/conf"
	"common/pkg/server"
	bbsuserv1 "common/proto/gen/bbs/v1/user"
	cerrors "common/proto/gen/common/errors"
	userv1 "common/proto/gen/user/v1"
	"context"
	"encoding/json"
	"fmt"
	stdhttp "net/http"

	"github.com/go-kratos/kratos/contrib/middleware/validate/v2"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-kratos/kratos/v2/middleware/logging"
	"github.com/go-kratos/kratos/v2/middleware/metrics"
	"github.com/go-kratos/kratos/v2/middleware/recovery"
	"github.com/go-kratos/kratos/v2/middleware/selector"
	"github.com/go-kratos/kratos/v2/middleware/tracing"
	kratoshttp "github.com/go-kratos/kratos/v2/transport/http"
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
func NewHTTPServer(c *conf.Bootstrap, logger log.Logger, services []server.HttpService, authClient userv1.AuthServiceClient) *kratoshttp.Server {
	authRequiredMatch := func(_ context.Context, operation string) bool {
		_, ok := bbsHTTPAuthOperationWhitelist[operation]
		return !ok
	}

	var opts = []kratoshttp.ServerOption{
		kratoshttp.Middleware(
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
		kratoshttp.ResponseEncoder(server.HttpResponseEncoder),
		kratoshttp.ErrorEncoder(server.HttpErrorEncoder(func(r *stdhttp.Request, code cerrors.BusinessErrorCode, data json.RawMessage) string {
			return bbsErrorMessages.Resolve(r, code, data)
		})),
	}
	if c.Server.Http.Network != "" {
		opts = append(opts, kratoshttp.Network(c.Server.Http.Network))
	}
	if c.Server.Http.Host != "" && c.Server.Http.Port != 0 {
		opts = append(opts, kratoshttp.Address(fmt.Sprintf("%s:%d", c.Server.Http.Host, c.Server.Http.Port)))
	}
	if c.Server.Http.Timeout != nil {
		opts = append(opts, kratoshttp.Timeout(c.Server.Http.Timeout.AsDuration()))
	}
	srv := kratoshttp.NewServer(opts...)
	srv.Handle("/metrics", promhttp.Handler())
	for _, s := range services {
		s.RegisterHttp(srv)
	}
	return srv
}
