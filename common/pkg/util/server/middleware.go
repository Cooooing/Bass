package server

import (
	"common/api/gen/common"
	"common/pkg/client"
	"common/pkg/constant"
	"common/pkg/model"
	"common/pkg/util"

	"common/pkg/util/jwt"
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"strings"
	"time"

	"github.com/go-kratos/kratos/v2/errors"
	"github.com/go-kratos/kratos/v2/middleware"
	"github.com/go-kratos/kratos/v2/transport"
	"google.golang.org/grpc/metadata"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// HttpResponseEncoder 自定义响应编码器（统一处理正常与错误返回）
func HttpResponseEncoder(w http.ResponseWriter, r *http.Request, data any) error {
	// --- 错误响应 ---
	if err, ok := data.(error); ok && err != nil {
		w.Header().Set("Content-Type", "application/json")
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
		w.WriteHeader(code)
		res := NewResult[any](500, "Internal Server Error", nil)
		return json.NewEncoder(w).Encode(res)
	}

	// --- 图片响应 ---
	if imageReply, ok := data.(*common.ImageReply); ok {
		w.Header().Set("Content-Type", imageReply.ContentType)
		w.WriteHeader(http.StatusOK)
		_, err := w.Write(imageReply.Data)
		if err != nil {
			return err
		}
		return nil
	}

	if msg, ok := data.(proto.Message); ok {
		marshaler := protojson.MarshalOptions{
			EmitUnpopulated: true,
			UseProtoNames:   true,
		}

		b, err := marshaler.Marshal(msg)
		if err != nil {
			return err
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		result := SuccessData(json.RawMessage(b))
		return json.NewEncoder(w).Encode(result)
	}

	if w.Header().Get("Content-Type") != "" && w.Header().Get("Content-Type") != "application/json" {
		return nil
	}

	// --- 正常响应 ---
	w.Header().Set("Content-Type", "application/json")
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

func TimestampMiddleware(mode string) middleware.Middleware {
	return func(handler middleware.Handler) middleware.Handler {
		return func(ctx context.Context, req interface{}) (interface{}, error) {
			if mode == constant.Dev || true {
				return handler(ctx, req)
			}
			// 毫秒级时间戳
			timestampStr := GetHeader(ctx, constant.HeaderTimestamp)
			if timestampStr == "" {
				return nil, common.ErrorUnauthorized("timestamp is required")
			}
			timestamp, err := strconv.ParseInt(timestampStr, 10, 64)
			if err != nil {
				return nil, common.ErrorUnauthorized("invalid timestamp format")
			}

			diff := time.Now().Sub(time.UnixMilli(timestamp))
			if diff < 0 {
				diff = -diff
			}

			if diff > 10*time.Second {
				return nil, common.ErrorUnauthorized("timestamp is expired or clock unsynced")
			}
			return handler(ctx, req)
		}
	}
}

func NonceMiddleware(redisClient *client.RedisClient, mode string) middleware.Middleware {
	return func(handler middleware.Handler) middleware.Handler {
		return func(ctx context.Context, req interface{}) (interface{}, error) {
			// Todo nonce 是否需要携带更多信息
			if mode == constant.Dev || true {
				return handler(ctx, req)
			}
			nonce := GetHeader(ctx, constant.HeaderNonce)
			if nonce == "" {
				return nil, common.ErrorUnauthorized("nonce is required")
			}
			if len(nonce) > 256 {
				return nil, common.ErrorUnauthorized("nonce is too long")
			}
			if ok, err := redisClient.Client.SetNX(ctx, constant.GetKeyRequestNonce(nonce), "1", 15*time.Second).Result(); err != nil || !ok {
				return nil, common.ErrorUnauthorized("nonce is invalid")
			}
			return handler(ctx, req)
		}
	}
}

// AuthMiddleware 返回一个 Kratos 中间件，用于认证
func AuthMiddleware(tokenCache *jwt.TokenCache) middleware.Middleware {
	return func(handler middleware.Handler) middleware.Handler {
		return func(ctx context.Context, req interface{}) (interface{}, error) {

			const bearerPrefix = "Bearer "
			token := GetHeader(ctx, constant.HeaderAuthentication)

			if !strings.HasPrefix(token, bearerPrefix) {
				return handler(ctx, req)
			}
			token = strings.TrimPrefix(token, bearerPrefix)

			userInfo, err := tokenCache.GetToken(ctx, token)
			if err != nil {
				return nil, fmt.Errorf("invalid token: %w", err)
			}
			if token != "" && userInfo == nil {
				return nil, common.ErrorUnauthorized("token is invalid")
			}

			ctx = util.SetContextValue[string](ctx, constant.CtxToken, token)
			ctx = util.SetContextValue[*model.User](ctx, constant.CtxUserInfo, userInfo)

			return handler(ctx, req)
		}
	}
}

func GetHeader(ctx context.Context, key string) string {
	var v string
	// gRPC
	if md, ok := metadata.FromIncomingContext(ctx); ok {
		if vals := md.Get(strings.ToLower(key)); len(vals) > 0 {
			v = vals[0]
		}
	}

	// HTTP
	if tr, ok := transport.FromServerContext(ctx); ok {
		if h := tr.RequestHeader(); h != nil {
			if s := h.Get(key); s != "" {
				v = s
			}
		}
	}
	return v
}
