package server

import (
	commonClient "common/pkg/client"
	"common/pkg/client/rpc"
	"common/pkg/constant"
	commonenum "common/pkg/enum"
	"common/pkg/server"
	cerrors "common/proto/gen/common/errors"
	v1 "common/proto/gen/game_idle_bff/v1"
	"context"
	"encoding/json"
	"fmt"
	"game_idle_bff/internal/config"
	"log/slog"
	stdhttp "net/http"

	"github.com/go-kratos/kratos/contrib/middleware/validate/v3"
	"github.com/go-kratos/kratos/v3/middleware"
	"github.com/go-kratos/kratos/v3/middleware/recovery"
	"github.com/go-kratos/kratos/v3/middleware/selector"
	kratoshttp "github.com/go-kratos/kratos/v3/transport/http"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

var gameIdleBffHTTPPublicOperations = map[string]struct{}{
	v1.OperationGameIdleBffAuthServiceRegister: {},
	v1.OperationGameIdleBffAuthServiceLogin:    {},
}

func NewHTTPAuthMiddlewares(authClient *rpc.UserClient) []middleware.Middleware {
	authRequiredMatch := func(_ context.Context, operation string) bool {
		_, ok := gameIdleBffHTTPPublicOperations[operation]
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
		kratoshttp.Filter(server.HTTPTraceMiddleware(), server.HTTPAccessLogMiddleware(logger)),
		kratoshttp.Middleware(middlewares...),
		kratoshttp.RequestDecoder(server.ProtoJSONRequestDecoder),
		kratoshttp.ResponseEncoder(server.HttpRespEncoder),
		kratoshttp.ErrorEncoder(server.HttpErrorEncoder(func(r *stdhttp.Request, code cerrors.BusinessErrorCode, data json.RawMessage) string {
			return code.String()
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
