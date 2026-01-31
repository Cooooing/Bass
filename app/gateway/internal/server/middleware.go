package server

import (
	cv1 "common/api/common/v1"
	"common/pkg/constant"
	"common/pkg/util"
	"context"
	"errors"
	"gateway/internal/conf"
	"regexp"
	"strings"

	"github.com/go-kratos/kratos/v2/middleware"
	"github.com/go-kratos/kratos/v2/middleware/selector"
	"github.com/go-kratos/kratos/v2/transport"
	transporthttp "github.com/go-kratos/kratos/v2/transport/http"
	"github.com/qiniu/go-sdk/v7/auth"
	"github.com/qiniu/go-sdk/v7/auth/qbox"
)

var NoAuthEndpoints = map[string]struct{}{
	"^.*/v1/system/health$": {},

	"^/user/v1/authentication/.*$":                {},
	"^/user/v1/oss/qiniu/uploadCallback$":         {},
	"^/user/v1/oss/qiniu/incrementAuditCallback$": {},

	"^/signal/v1/node/register":   {},
	"^/signal/v1/node/unregister": {},
	"^/signal/v1/node/online":     {},
	"^/signal/v1/node/offline":    {},
	"^/signal/v1/node/list":       {},
}

var QiniuCallbackEndpoints = map[string]struct{}{
	"/user/v1/oss/qiniu/uploadCallback":         {},
	"/user/v1/oss/qiniu/incrementAuditCallback": {},
}

// QiniuCallbackMatch 七牛回调匹配
func QiniuCallbackMatch() selector.MatchFunc {
	return func(ctx context.Context, operation string) bool {
		if _, exist := QiniuCallbackEndpoints[operation]; exist {
			return true
		}
		return false
	}
}

// QiniuCallbackSignMiddleware 七牛回调验签
func QiniuCallbackSignMiddleware(c *conf.Bootstrap) middleware.Middleware {
	return func(handler middleware.Handler) middleware.Handler {
		return func(ctx context.Context, req interface{}) (interface{}, error) {
			errNoAuth := cv1.ErrorForbidden("verify callback failed")
			mac := auth.New(c.Oss.Qiniu.AccessKey, c.Oss.Qiniu.SecretKey)
			if r, ok := transporthttp.RequestFromServerContext(ctx); ok {
				verify, err := qbox.VerifyCallback(mac, r)
				if err != nil {
					return nil, err
				}
				if verify {
					return handler(ctx, req)
				}
			}
			return nil, errNoAuth
		}
	}
}

func UserAPIMatch() selector.MatchFunc {
	return func(ctx context.Context, operation string) bool {
		allow := false
		for pattern := range NoAuthEndpoints {
			matched, err := regexp.MatchString(pattern, operation)
			if err == nil && matched {
				allow = true
				break
			}
		}
		return allow
	}
}

// AuthMiddleware 认证中间件
func AuthMiddleware(tokenCache *util.TokenCache) middleware.Middleware {
	return func(handler middleware.Handler) middleware.Handler {
		return func(ctx context.Context, req interface{}) (interface{}, error) {
			tr, ok := transport.FromServerContext(ctx)
			if !ok {
				return nil, errors.New("transport not found")
			}
			// 使用请求路径重写 operation
			if _, ok := tr.(*transporthttp.Transport); ok {
				// 获取 token
				token := strings.TrimPrefix(tr.RequestHeader().Get(constant.HeaderAuthentication), "Bearer ")

				if token == "" {
					return nil, cv1.ErrorUnauthorized("token is not provided")
				}

				// 验证 token
				userInfo, err := tokenCache.GetToken(ctx, token)
				if err != nil {
					return nil, err
				}

				// 权限范围 Todo 用户组权限规则后续持久化入库

				// 设置上下文
				ctx = context.WithValue(ctx, constant.CtxUserInfo, userInfo)
				ctx = context.WithValue(ctx, constant.Token, token)

				return handler(ctx, req)
			} else {
				return nil, errors.New("transport not http")
			}
		}
	}
}

// PermissionMiddleware 鉴权中间件
func PermissionMiddleware() middleware.Middleware {
	return func(handler middleware.Handler) middleware.Handler {
		return func(ctx context.Context, req interface{}) (interface{}, error) {
			return handler(ctx, req)
		}
	}
}
