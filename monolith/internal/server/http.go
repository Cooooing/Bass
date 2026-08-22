package server

import (
	commonserver "common/pkg/server"
	cerrors "common/proto/gen/common/errors"
	"encoding/json"
	"fmt"
	"log/slog"
	"monolith/internal/config"
	stdhttp "net/http"

	"github.com/go-kratos/kratos/contrib/middleware/validate/v3"
	"github.com/go-kratos/kratos/v3/middleware"
	"github.com/go-kratos/kratos/v3/middleware/recovery"
	kratoshttp "github.com/go-kratos/kratos/v3/transport/http"
)

func NewHTTPServer(
	c *config.Bootstrap,
	logger *slog.Logger,
	services []commonserver.Service,
	moduleMiddlewares []middleware.Middleware,
) *kratoshttp.Server {
	var middlewares []middleware.Middleware
	middlewares = append(middlewares,
		commonserver.RequestLogContextMiddleware(),
		recovery.Recovery(),
	)
	middlewares = append(middlewares, moduleMiddlewares...)
	middlewares = append(middlewares, validate.ProtoValidate())

	var opts = []kratoshttp.ServerOption{
		kratoshttp.Filter(commonserver.HTTPAccessLogMiddleware(logger)),
		kratoshttp.Middleware(middlewares...),
		kratoshttp.RequestDecoder(commonserver.ProtoJSONRequestDecoder),
		kratoshttp.ResponseEncoder(commonserver.HttpRespEncoder),
		kratoshttp.ErrorEncoder(commonserver.HttpErrorEncoder(func(_ *stdhttp.Request, _ cerrors.BusinessErrorCode, _ json.RawMessage) string {
			return ""
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
	for _, s := range services {
		s.RegisterHttp(srv)
	}
	return srv
}
