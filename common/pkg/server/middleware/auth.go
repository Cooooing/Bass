package middleware

import (
	cerrors "common/proto/gen/common/errors"
	userv1 "common/proto/gen/user/v1"
	"context"
	"log/slog"
	"strings"

	"common/pkg/apperror"
	"common/pkg/constant"
	commonenum "common/pkg/enum"
	"common/pkg/model"
	"common/pkg/server"
	"common/pkg/util"

	kratoslog "github.com/go-kratos/kratos/v3/log"
	"github.com/go-kratos/kratos/v3/middleware"
)

func UserAuthMiddleware(authClient userv1.AuthServiceClient, realm commonenum.LoginRealm) middleware.Middleware {
	return func(handler middleware.Handler) middleware.Handler {
		return func(ctx context.Context, req any) (any, error) {
			const bearerPrefix = "Bearer "
			token := server.GetHeader(ctx, constant.HeaderAuthentication)

			if !strings.HasPrefix(token, bearerPrefix) {
				return nil, apperror.New(cerrors.BusinessErrorCode_BUSINESS_ERROR_CODE_USER_TOKEN_REQUIRED)
			}
			token = strings.TrimSpace(strings.TrimPrefix(token, bearerPrefix))
			if token == "" {
				return nil, apperror.New(cerrors.BusinessErrorCode_BUSINESS_ERROR_CODE_USER_TOKEN_REQUIRED)
			}

			reply, err := authClient.ParseToken(ctx, &userv1.ParseToken_Req{
				AccessToken: token,
				Realm:       commonenum.LoginRealmMap.MustToProto(realm),
			})
			if err != nil {
				return nil, apperror.New(cerrors.BusinessErrorCode_BUSINESS_ERROR_CODE_USER_TOKEN_INVALID).WithCause(err)
			}
			tokenUser := reply.GetUser()
			if tokenUser == nil {
				return nil, apperror.New(cerrors.BusinessErrorCode_BUSINESS_ERROR_CODE_USER_TOKEN_INVALID)
			}
			userInfo := &model.User{
				ID:       tokenUser.GetId(),
				Name:     tokenUser.GetName(),
				Nickname: tokenUser.GetNickname(),
				Language: tokenUser.GetLanguage(),
				Timezone: tokenUser.GetTimezone(),
			}

			ctx = util.SetContextValue[string](ctx, constant.CtxToken, token)
			ctx = util.SetContextValue[*model.User](ctx, constant.CtxUserInfo, userInfo)
			ctx = kratoslog.ContextWithAttrs(ctx, slog.Int64(constant.LogFieldUserID, userInfo.ID))

			return handler(ctx, req)
		}
	}
}

func OptionalUserAuthMiddleware(authClient userv1.AuthServiceClient, realm commonenum.LoginRealm) middleware.Middleware {
	return func(handler middleware.Handler) middleware.Handler {
		return func(ctx context.Context, req any) (any, error) {
			bearerPrefix := "Bearer "
			token := server.GetHeader(ctx, constant.HeaderAuthentication)
			if token == "" {
				return handler(ctx, req)
			}
			if !strings.HasPrefix(token, bearerPrefix) {
				return nil, apperror.New(cerrors.BusinessErrorCode_BUSINESS_ERROR_CODE_USER_TOKEN_INVALID)
			}
			token = strings.TrimSpace(strings.TrimPrefix(token, bearerPrefix))
			if token == "" {
				return nil, apperror.New(cerrors.BusinessErrorCode_BUSINESS_ERROR_CODE_USER_TOKEN_INVALID)
			}

			reply, err := authClient.ParseToken(ctx, &userv1.ParseToken_Req{
				AccessToken: token,
				Realm:       commonenum.LoginRealmMap.MustToProto(realm),
			})
			if err != nil {
				return nil, apperror.New(cerrors.BusinessErrorCode_BUSINESS_ERROR_CODE_USER_TOKEN_INVALID).WithCause(err)
			}
			tokenUser := reply.GetUser()
			if tokenUser == nil {
				return nil, apperror.New(cerrors.BusinessErrorCode_BUSINESS_ERROR_CODE_USER_TOKEN_INVALID)
			}
			userInfo := &model.User{
				ID:       tokenUser.GetId(),
				Name:     tokenUser.GetName(),
				Nickname: tokenUser.GetNickname(),
				Language: tokenUser.GetLanguage(),
				Timezone: tokenUser.GetTimezone(),
			}

			ctx = util.SetContextValue[string](ctx, constant.CtxToken, token)
			ctx = util.SetContextValue[*model.User](ctx, constant.CtxUserInfo, userInfo)
			ctx = kratoslog.ContextWithAttrs(ctx, slog.Int64(constant.LogFieldUserID, userInfo.ID))

			return handler(ctx, req)
		}
	}
}
