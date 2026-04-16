package server

import (
	"common/api/gen/common"
	"common/pkg/constant"
	"common/pkg/model"
	"common/pkg/util"

	"common/pkg/util/jwt"
	"context"
	"encoding/json"
	"errors"
	"gateway/internal/biz/domain"
	"gateway/internal/conf"
	"net"
	"regexp"
	"strings"

	"github.com/go-kratos/kratos/v2/metadata"
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
			errNoAuth := common.ErrorForbidden("verify callback failed")
			mac := auth.New(c.Server.Oss.Qiniu.AccessKey, c.Server.Oss.Qiniu.SecretKey)
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
func AuthMiddleware(tokenCache *jwt.TokenCache) middleware.Middleware {
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
					return nil, common.ErrorUnauthorized("token is not provided")
				}

				// 验证 token
				userInfo, err := tokenCache.GetToken(ctx, token)
				if err != nil {
					return nil, err
				}

				// 权限范围 Todo 用户组权限规则后续持久化入库

				// 设置上下文
				ctx = util.SetContextValue[*model.User](ctx, constant.CtxUserInfo, userInfo)
				ctx = util.SetContextValue[string](ctx, constant.CtxToken, token)

				return handler(ctx, req)
			}

			return nil, errors.New("transport not http")
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

// IpMiddleware Ip 中间件
func IpMiddleware(ipDomain *domain.IpDomain) middleware.Middleware {
	return func(handler middleware.Handler) middleware.Handler {
		return func(ctx context.Context, req interface{}) (interface{}, error) {
			tr, ok := transport.FromServerContext(ctx)
			if !ok {
				return nil, errors.New("transport not found")
			}
			var ip string
			if _, ok := tr.(*transporthttp.Transport); ok {

				if xff := tr.RequestHeader().Get("X-Forwarded-For"); xff != "" {
					ip = strings.TrimSpace(strings.Split(xff, ",")[0])
				}
				if ip == "" {
					ip = tr.RequestHeader().Get("X-Real-IP")
				}
				if ip == "" {
					ip, _, _ = net.SplitHostPort(tr.(*transporthttp.Transport).Request().RemoteAddr)
				}

				ipInfo, err := ipDomain.GetInfo(ctx, ip)
				if err != nil {
					return nil, err
				}
				ctx = util.SetContextValue[*model.IpInfo](ctx, constant.CtxIpInfo, ipInfo)
				ipInfoBytes, err := json.Marshal(ipInfo)
				if err != nil {
					return nil, err
				}
				ctx = metadata.AppendToClientContext(ctx, constant.CtxIpInfo.String(), string(ipInfoBytes))
				return handler(ctx, req)
			}

			return nil, errors.New("transport not http")
		}
	}
}
