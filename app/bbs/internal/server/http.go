package server

import (
	"bbs/internal/config"
	commonClient "common/pkg/client"
	"common/pkg/constant"
	commonenum "common/pkg/enum"
	"common/pkg/server"
	bbscontentv1 "common/proto/gen/bbs/v1/content"
	bbsuserv1 "common/proto/gen/bbs/v1/user"
	cerrors "common/proto/gen/common/errors"
	userv1 "common/proto/gen/user/v1"
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

var bbsHTTPPublicOperations = map[string]struct{}{
	bbsuserv1.OperationAuthServiceRegister:      {},
	bbsuserv1.OperationAuthServiceLogin:         {},
	bbsuserv1.OperationAuthServiceRefreshToken:  {},
	bbsuserv1.OperationAccountServiceAvatar:     {},
	bbsuserv1.OperationAccountServiceGetProfile: {},
}

var bbsHTTPOptionalAuthOperations = map[string]struct{}{
	bbscontentv1.OperationArticleServiceList:         {},
	bbscontentv1.OperationArticleServiceGet:          {},
	bbscontentv1.OperationCommentServiceList:         {},
	bbscontentv1.OperationCommentServiceListThreads:  {},
	bbscontentv1.OperationCommentServiceListReplies:  {},
	bbscontentv1.OperationCommentServiceListTimeline: {},
	bbscontentv1.OperationDomainServiceList:          {},
	bbscontentv1.OperationTagServiceList:             {},
	bbscontentv1.OperationTagServiceListArticleTags:  {},
	bbscontentv1.OperationPostscriptServiceList:      {},
	bbsuserv1.OperationOtpServiceSendEmailOtp:        {},
	bbsuserv1.OperationOtpServiceSendPhoneOtp:        {},
	bbsuserv1.OperationRelationServiceGetStatus:      {},
}

func NewHTTPAuthMiddlewares(authClient userv1.AuthServiceClient) []middleware.Middleware {
	operationAuthGroups := map[string]string{}
	for operation := range bbsHTTPPublicOperations {
		operationAuthGroups[operation] = "public"
	}
	for operation := range bbsHTTPOptionalAuthOperations {
		if group, ok := operationAuthGroups[operation]; ok {
			panic(fmt.Sprintf("bbs http operation auth group conflict: %s in %s and optional", operation, group))
		}
		operationAuthGroups[operation] = "optional"
	}
	authRequiredMatch := func(_ context.Context, operation string) bool {
		if _, ok := bbsHTTPPublicOperations[operation]; ok {
			return false
		}
		if _, ok := bbsHTTPOptionalAuthOperations[operation]; ok {
			return false
		}
		return true
	}
	optionalAuthMatch := func(_ context.Context, operation string) bool {
		_, ok := bbsHTTPOptionalAuthOperations[operation]
		return ok
	}

	return []middleware.Middleware{
		selector.Server(server.OptionalUserAuthMiddleware(authClient, commonenum.LoginRealmBBS)).Match(optionalAuthMatch).Build(),
		selector.Server(server.UserAuthMiddleware(authClient, commonenum.LoginRealmBBS)).Match(authRequiredMatch).Build(),
	}
}

func NewHTTPServer(
	c *config.Bootstrap,
	logger *slog.Logger,
	obs *commonClient.Observer,
	services []server.Service,
	authClient userv1.AuthServiceClient,
) *kratoshttp.Server {
	middlewares := []middleware.Middleware{
		server.RequestLogContextMiddleware(),
		obs.ServerMiddleware(),
		recovery.Recovery(),
	}
	middlewares = append(middlewares, NewHTTPAuthMiddlewares(authClient)...)
	middlewares = append(middlewares, validate.ProtoValidate())

	var opts = []kratoshttp.ServerOption{
		kratoshttp.Filter(server.HTTPTraceMiddleware(), server.HTTPAccessLogMiddleware(logger)),
		kratoshttp.Middleware(middlewares...),
		kratoshttp.RequestDecoder(server.ProtoJSONRequestDecoder),
		kratoshttp.ResponseEncoder(server.HttpRespEncoder),
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
