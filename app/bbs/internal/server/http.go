package server

import (
	"bbs/internal/conf"
	commonClient "common/pkg/client"
	"common/pkg/constant"
	"common/pkg/server"
	bbsuserv1 "common/proto/gen/bbs/v1/user"
	cerrors "common/proto/gen/common/errors"
	userv1 "common/proto/gen/user/v1"
	"context"
	"encoding/json"
	"fmt"
	"log/slog"
	stdhttp "net/http"

	"github.com/go-kratos/kratos/contrib/middleware/validate/v3"
	"github.com/go-kratos/kratos/v3/middleware/recovery"
	"github.com/go-kratos/kratos/v3/middleware/selector"
	kratoshttp "github.com/go-kratos/kratos/v3/transport/http"
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

func NewHTTPServer(c *conf.Bootstrap, logger *slog.Logger, obs *commonClient.Observer, services []server.HttpService, authClient userv1.AuthServiceClient) *kratoshttp.Server {
	authRequiredMatch := func(_ context.Context, operation string) bool {
		_, ok := bbsHTTPAuthOperationWhitelist[operation]
		return !ok
	}

	var opts = []kratoshttp.ServerOption{
		kratoshttp.Middleware(
			server.RequestLogContextMiddleware(),
			selector.Server(server.UserAuthMiddleware(authClient)).Match(authRequiredMatch).Build(),
			obs.ServerMiddleware(),
			recovery.Recovery(),
			validate.ProtoValidate(),
		),
		kratoshttp.ResponseEncoder(server.HttpResponseEncoder),
		kratoshttp.ErrorEncoder(server.HttpErrorEncoder(func(r *stdhttp.Request, code cerrors.BusinessErrorCode, data json.RawMessage) string {
			return bbsErrorMessages.Resolve(r, code, data)
		})),
	}
	if c.Http.Network != "" {
		opts = append(opts, kratoshttp.Network(c.Http.Network))
	}
	if c.Http.Host != "" && c.Http.Port != 0 {
		opts = append(opts, kratoshttp.Address(fmt.Sprintf("%s:%d", c.Http.Host, c.Http.Port)))
	}
	if c.Http.Timeout != nil {
		opts = append(opts, kratoshttp.Timeout(c.Http.Timeout.AsDuration()))
	}
	srv := kratoshttp.NewServer(opts...)
	if obsConf := c.GetObservability(); obsConf != nil && obsConf.GetEnableMetrics() {
		srv.Handle("/metrics", promhttp.Handler())
		logger.Info("metrics endpoint registered", slog.String(constant.LogFieldPath, "/metrics"))
	}
	for _, s := range services {
		s.RegisterHttp(srv)
	}
	return srv
}
