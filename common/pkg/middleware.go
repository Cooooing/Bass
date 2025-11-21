package pkg

import (
	"common/pkg/constant"
	"common/pkg/model"
	"common/pkg/util"
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"strings"

	"github.com/go-kratos/kratos/v2/errors"
	"github.com/go-kratos/kratos/v2/middleware"
	"github.com/go-kratos/kratos/v2/transport"
	"google.golang.org/grpc/metadata"
)

// HttpResponseEncoder 自定义响应编码器（统一处理正常与错误返回）
func HttpResponseEncoder(w http.ResponseWriter, r *http.Request, data any) error {
	w.Header().Set("Content-Type", "application/json")

	// --- 错误响应 ---
	if err, ok := data.(error); ok && err != nil {
		se := errors.FromError(err)

		code := int(se.Code)
		if code == 0 {
			code = http.StatusInternalServerError
		}

		// 判断是否是业务错误（proto 定义）
		if se.Reason != "" {
			// Reason 存在 → 业务错误
			w.WriteHeader(code)
			res := NewResult[any](code, se.Message, nil)
			return json.NewEncoder(w).Encode(res)
		}

		// 没有 reason → 系统错误
		w.WriteHeader(http.StatusInternalServerError)
		res := NewResult[any](500, "Internal Server Error", nil)
		return json.NewEncoder(w).Encode(res)
	}

	// --- 正常响应 ---
	w.WriteHeader(http.StatusOK)
	result := SuccessData(data)
	return json.NewEncoder(w).Encode(result)
}

// HttpErrorEncoder 兜底错误编码器（Kratos 框架异常时调用）
func HttpErrorEncoder(w http.ResponseWriter, r *http.Request, err error) {
	w.Header().Set("Content-Type", "application/json")

	se := errors.FromError(err)
	code := int(se.Code)
	if code == 0 {
		code = http.StatusInternalServerError
	}

	w.WriteHeader(code)
	res := NewResult[any](code, se.Message, nil)
	_ = json.NewEncoder(w).Encode(res)
}

// AuthMiddleware 返回一个 Kratos 中间件，用于认证
func AuthMiddleware(tokenRepo *util.TokenRepo) middleware.Middleware {
	return func(handler middleware.Handler) middleware.Handler {
		return func(ctx context.Context, req interface{}) (interface{}, error) {

			const bearerPrefix = "Bearer "
			var token string
			// gRPC
			if md, ok := metadata.FromIncomingContext(ctx); ok {
				if vals := md.Get(strings.ToLower(constant.Authentication)); len(vals) > 0 {
					token = strings.TrimPrefix(vals[0], bearerPrefix)
				}
			}

			// HTTP
			if tr, ok := transport.FromServerContext(ctx); ok {
				if h := tr.RequestHeader(); h != nil {
					if s := h.Get(constant.Authentication); s != "" {
						token = strings.TrimPrefix(s, bearerPrefix)
					}
				}
			}

			if tr, ok := transport.FromServerContext(ctx); ok {
				token = strings.TrimPrefix(tr.RequestHeader().Get(constant.Authentication), bearerPrefix)
			}

			userInfo, err := tokenRepo.GetToken(ctx, token)
			if err != nil {
				return nil, fmt.Errorf("invalid token: %w", err)
			}

			ctx = util.SetContextValue[string](ctx, constant.CtxToken, token)
			ctx = util.SetContextValue[*model.User](ctx, constant.CtxUserInfo, userInfo)

			return handler(ctx, req)
		}
	}
}
