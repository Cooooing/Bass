package server

import (
	"common/pkg/apperror"
	"common/pkg/client"
	"common/pkg/constant"
	"common/pkg/model"
	"common/pkg/util"
	"common/proto/gen/common"
	cerrors "common/proto/gen/common/errors"
	userv1 "common/proto/gen/user/v1"
	"context"
	"encoding/json"
	"net/http"
	"strconv"
	"strings"
	"time"

	"github.com/go-kratos/kratos/v2/middleware"
	"github.com/go-kratos/kratos/v2/transport"
	"google.golang.org/grpc/metadata"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

func HttpResponseEncoder(w http.ResponseWriter, r *http.Request, data any) error {
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
		result := NewResult[json.RawMessage](int(cerrors.BusinessErrorCode_BUSINESS_ERROR_CODE_SUCCESS), "success", b)
		return json.NewEncoder(w).Encode(result)
	}

	if w.Header().Get("Content-Type") != "" && w.Header().Get("Content-Type") != "application/json" {
		return nil
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	result := NewResult[any](int(cerrors.BusinessErrorCode_BUSINESS_ERROR_CODE_SUCCESS), "success", data)
	return json.NewEncoder(w).Encode(result)
}

func HttpErrorEncoder(resolve func(r *http.Request, code cerrors.BusinessErrorCode, data json.RawMessage) string) func(w http.ResponseWriter, r *http.Request, err error) {
	return func(w http.ResponseWriter, r *http.Request, err error) {
		w.Header().Set("Content-Type", "application/json")

		businessCode, ok := apperror.BusinessCode(err)
		if !ok || businessCode == cerrors.BusinessErrorCode_BUSINESS_ERROR_CODE_SUCCESS {
			businessCode = cerrors.BusinessErrorCode_BUSINESS_ERROR_CODE_COMMON_UNKNOWN
		}
		statusCode := apperror.StatusCode(businessCode)

		errorData := apperror.Data(err)
		message := businessCode.String()
		if resolve != nil {
			if resolved := resolve(r, businessCode, errorData); resolved != "" {
				message = resolved
			}
		}

		w.WriteHeader(statusCode)
		res := NewResult[any](int(businessCode), message, errorData)
		_ = json.NewEncoder(w).Encode(res)
	}
}

func TimestampMiddleware(mode string) middleware.Middleware {
	return func(handler middleware.Handler) middleware.Handler {
		return func(ctx context.Context, req interface{}) (interface{}, error) {
			if mode == constant.Dev {
				return handler(ctx, req)
			}
			timestampStr := GetHeader(ctx, constant.HeaderTimestamp)
			if timestampStr == "" {
				return nil, apperror.New(cerrors.BusinessErrorCode_BUSINESS_ERROR_CODE_COMMON_UNAUTHORIZED)
			}
			timestamp, err := strconv.ParseInt(timestampStr, 10, 64)
			if err != nil {
				return nil, apperror.New(cerrors.BusinessErrorCode_BUSINESS_ERROR_CODE_COMMON_UNAUTHORIZED)
			}

			diff := time.Now().Sub(time.UnixMilli(timestamp))
			if diff < 0 {
				diff = -diff
			}

			if diff > 10*time.Second {
				return nil, apperror.New(cerrors.BusinessErrorCode_BUSINESS_ERROR_CODE_COMMON_UNAUTHORIZED)
			}
			return handler(ctx, req)
		}
	}
}

func NonceMiddleware(redisClient *client.RedisClient, mode string) middleware.Middleware {
	return func(handler middleware.Handler) middleware.Handler {
		return func(ctx context.Context, req interface{}) (interface{}, error) {
			if mode == constant.Dev {
				return handler(ctx, req)
			}
			nonce := GetHeader(ctx, constant.HeaderNonce)
			if nonce == "" {
				return nil, apperror.New(cerrors.BusinessErrorCode_BUSINESS_ERROR_CODE_COMMON_UNAUTHORIZED)
			}
			if len(nonce) > 256 {
				return nil, apperror.New(cerrors.BusinessErrorCode_BUSINESS_ERROR_CODE_COMMON_UNAUTHORIZED)
			}
			if ok, err := redisClient.Client.SetNX(ctx, constant.GetKeyRequestNonce(nonce), "1", 15*time.Second).Result(); err != nil || !ok {
				return nil, apperror.New(cerrors.BusinessErrorCode_BUSINESS_ERROR_CODE_COMMON_UNAUTHORIZED)
			}
			return handler(ctx, req)
		}
	}
}

// UserAuthMiddleware 通过 user 服务强制解析登录令牌。
func UserAuthMiddleware(authClient userv1.AuthServiceClient) middleware.Middleware {
	return func(handler middleware.Handler) middleware.Handler {
		return func(ctx context.Context, req interface{}) (interface{}, error) {
			const bearerPrefix = "Bearer "
			token := GetHeader(ctx, constant.HeaderAuthentication)

			if !strings.HasPrefix(token, bearerPrefix) {
				return nil, apperror.New(cerrors.BusinessErrorCode_BUSINESS_ERROR_CODE_USER_TOKEN_REQUIRED)
			}
			token = strings.TrimSpace(strings.TrimPrefix(token, bearerPrefix))
			if token == "" {
				return nil, apperror.New(cerrors.BusinessErrorCode_BUSINESS_ERROR_CODE_USER_TOKEN_REQUIRED)
			}

			reply, err := authClient.ParseToken(ctx, &userv1.ParseToken_Request{Token: token})
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

			return handler(ctx, req)
		}
	}
}

func GetHeader(ctx context.Context, key string) string {
	var v string
	if md, ok := metadata.FromIncomingContext(ctx); ok {
		if vals := md.Get(strings.ToLower(key)); len(vals) > 0 {
			v = vals[0]
		}
	}

	if tr, ok := transport.FromServerContext(ctx); ok {
		if h := tr.RequestHeader(); h != nil {
			if s := h.Get(key); s != "" {
				v = s
			}
		}
	}
	return v
}
