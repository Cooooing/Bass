package server

import (
	"bff_game_idle/internal/config"
	"bff_game_idle/internal/errormessage"
	commonClient "common/pkg/client"
	"common/pkg/client/rpc"
	"common/pkg/constant"
	commonenum "common/pkg/enum"
	"common/pkg/server"
	v1 "common/proto/gen/bff_game_idle/v1"
	cerrors "common/proto/gen/common/errors"
	"context"
	"encoding/json"
	"fmt"
	"log/slog"
	stdhttp "net/http"

	"github.com/go-kratos/kratos/contrib/middleware/validate/v3"
	"github.com/go-kratos/kratos/v3/middleware"
	"github.com/go-kratos/kratos/v3/middleware/recovery"
	"github.com/go-kratos/kratos/v3/middleware/selector"
	kratoshttp "github.com/go-kratos/kratos/v3/transport/http"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

func NewHTTPAuthMiddlewares(authClient *rpc.UserClient) []middleware.Middleware {
	publicOperations := map[string]struct{}{
		v1.OperationAuthServiceRegister: {},
		v1.OperationAuthServiceLogin:    {},
	}
	authRequiredMatch := func(_ context.Context, operation string) bool {
		_, ok := publicOperations[operation]
		return !ok
	}

	return []middleware.Middleware{
		selector.Server(server.UserAuthMiddleware(authClient.Auth, commonenum.LoginRealmGameIdle)).Match(authRequiredMatch).Build(),
	}
}

func NewHTTPServer(
	c *config.Bootstrap,
	logger *slog.Logger,
	obs *commonClient.Observer,
	services []server.Service,
	userClient *rpc.UserClient,
) *kratoshttp.Server {
	_ = userClient
	middlewares := []middleware.Middleware{
		server.RequestLogContextMiddleware(),
		obs.ServerMiddleware(),
		recovery.Recovery(),
	}
	middlewares = append(middlewares, NewHTTPAuthMiddlewares(userClient)...)
	middlewares = append(middlewares, validate.ProtoValidate())

	opts := []kratoshttp.ServerOption{
		kratoshttp.Filter(server.HTTPCORSFilter(), server.HTTPTraceMiddleware(), server.HTTPAccessLogMiddleware(logger)),
		kratoshttp.Middleware(middlewares...),
		kratoshttp.RequestDecoder(server.ProtoJSONRequestDecoder),
		kratoshttp.ResponseEncoder(server.HttpRespEncoder),
		kratoshttp.ErrorEncoder(server.HttpErrorEncoder(func(r *stdhttp.Request, code cerrors.BusinessErrorCode, data json.RawMessage) string {
			return errormessage.ResolveHTTP(r, code, data)
		})),
	}
	if c.GetHttp().GetNetwork() != "" {
		opts = append(opts, kratoshttp.Network(c.GetHttp().GetNetwork()))
	}
	if c.GetHttp().GetHost() != "" && c.GetHttp().GetPort() != 0 {
		opts = append(opts, kratoshttp.Address(fmt.Sprintf("%s:%d", c.GetHttp().GetHost(), c.GetHttp().GetPort())))
	}
	if c.GetHttp().GetTimeout() != nil {
		opts = append(opts, kratoshttp.Timeout(c.GetHttp().GetTimeout().AsDuration()))
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
