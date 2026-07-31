package server

import (
	"common/pkg/apperror"
	"common/pkg/client"
	"common/pkg/constant"
	commonenum "common/pkg/enum"
	"common/pkg/model"
	"common/pkg/util"
	"common/proto/gen/common"
	cerrors "common/proto/gen/common/errors"
	userv1 "common/proto/gen/user/v1"
	"context"
	"encoding/json"
	"log/slog"
	"net"
	"net/http"
	"strconv"
	"strings"
	"time"

	"github.com/felixge/httpsnoop"
	kratoslog "github.com/go-kratos/kratos/v3/log"
	"github.com/go-kratos/kratos/v3/middleware"
	"github.com/go-kratos/kratos/v3/transport"
	kratoshttp "github.com/go-kratos/kratos/v3/transport/http"
	"github.com/google/uuid"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/codes"
	"go.opentelemetry.io/otel/propagation"
	oteltrace "go.opentelemetry.io/otel/trace"
	"google.golang.org/grpc/metadata"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

func HttpRespEncoder(w http.ResponseWriter, r *http.Request, data any) error {
	if ImageResp, ok := data.(*common.ImageResp); ok {
		w.Header().Set("Content-Type", ImageResp.ContentType)
		w.WriteHeader(http.StatusOK)
		_, err := w.Write(ImageResp.Data)
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

		span := oteltrace.SpanFromContext(r.Context())
		if span.SpanContext().IsValid() {
			span.RecordError(err)
			span.SetStatus(codes.Error, err.Error())
			span.SetAttributes(attribute.Int("http.response.status_code", statusCode), attribute.Int("bass.business_code", int(businessCode)))
		}
		kratoslog.ErrorContext(r.Context(), "http error",
			constant.LogFieldErr, err,
			"method", r.Method,
			constant.LogFieldPath, r.URL.Path,
			"query", r.URL.RawQuery,
			"remote_addr", r.RemoteAddr,
			"host", r.Host,
			"request_uri", r.RequestURI,
			constant.LogFieldStatusCode, statusCode,
			"business_code", int(businessCode),
			"business_reason", businessCode.String(),
			"message", message,
			"error_data", string(errorData),
		)

		w.WriteHeader(statusCode)
		res := NewResult[any](int(businessCode), message, errorData)
		_ = json.NewEncoder(w).Encode(res)
	}
}

func HTTPTraceMiddleware() kratoshttp.FilterFunc {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			ctx := otel.GetTextMapPropagator().Extract(r.Context(), propagation.HeaderCarrier(r.Header))
			requestID := strings.TrimSpace(r.Header.Get(constant.HeaderRequestID))
			if requestID == "" {
				requestID = uuid.NewString()
			}
			w.Header().Set(constant.HeaderRequestID, requestID)
			r.Header.Set(constant.HeaderRequestID, requestID)
			clientIP := ""
			for _, header := range [...]string{constant.HeaderForwardedFor, constant.HeaderRealIP, constant.HeaderClientIP} {
				for _, item := range strings.Split(r.Header.Get(header), ",") {
					if ip := strings.TrimSpace(item); ip != "" {
						clientIP = ip
						break
					}
				}
				if clientIP != "" {
					break
				}
			}
			if clientIP == "" {
				if host, _, err := net.SplitHostPort(r.RemoteAddr); err == nil {
					clientIP = host
				} else {
					clientIP = r.RemoteAddr
				}
			}
			ctx = kratoslog.ContextWithAttrs(ctx, slog.String(constant.LogFieldRequestID, requestID), slog.String(constant.LogFieldClientIP, clientIP))
			spanName := r.Method
			if r.URL != nil && r.URL.Path != "" {
				spanName += " " + r.URL.Path
			}
			ctx, span := otel.Tracer("common.http.server").Start(ctx, spanName, oteltrace.WithSpanKind(oteltrace.SpanKindServer), oteltrace.WithAttributes(
				attribute.String("http.request.method", r.Method),
				attribute.String("url.path", r.URL.Path),
				attribute.String("url.query", r.URL.RawQuery),
				attribute.String("server.address", r.Host),
				attribute.String("client.address", clientIP),
				attribute.String(constant.LogFieldRequestID, requestID),
			))
			defer span.End()

			metrics := httpsnoop.CaptureMetrics(next, w, r.WithContext(ctx))
			statusCode := metrics.Code
			if statusCode == 0 {
				statusCode = http.StatusOK
			}
			span.SetAttributes(attribute.Int("http.response.status_code", statusCode), attribute.Int64("http.response.body.size", metrics.Written))
			if statusCode >= http.StatusInternalServerError {
				span.SetStatus(codes.Error, http.StatusText(statusCode))
			}
		})
	}
}

func HTTPAccessLogMiddleware(logger *slog.Logger) kratoshttp.FilterFunc {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			start := time.Now()
			metrics := httpsnoop.CaptureMetrics(next, w, r)
			statusCode := metrics.Code
			if statusCode == 0 {
				statusCode = http.StatusOK
			}
			level := slog.LevelInfo
			if statusCode >= http.StatusInternalServerError {
				level = slog.LevelError
			} else if statusCode >= http.StatusBadRequest {
				level = slog.LevelWarn
			}
			requestHeaders := make(map[string][]string, len(r.Header))
			for key, values := range r.Header {
				if strings.EqualFold(key, constant.LogFieldAuthorization) || strings.EqualFold(key, constant.LogFieldCookie) || strings.EqualFold(key, "Set-Cookie") {
					requestHeaders[key] = []string{"***"}
					continue
				}
				copied := make([]string, len(values))
				copy(copied, values)
				requestHeaders[key] = copied
			}
			responseHeaders := make(map[string][]string, len(w.Header()))
			for key, values := range w.Header() {
				if strings.EqualFold(key, constant.LogFieldAuthorization) || strings.EqualFold(key, constant.LogFieldCookie) || strings.EqualFold(key, "Set-Cookie") {
					responseHeaders[key] = []string{"***"}
					continue
				}
				copied := make([]string, len(values))
				copy(copied, values)
				responseHeaders[key] = copied
			}
			logger.LogAttrs(r.Context(), level, "http access",
				slog.String(constant.LogFieldKind, constant.LogKindServer),
				slog.String("method", r.Method),
				slog.String(constant.LogFieldPath, r.URL.Path),
				slog.String("query", r.URL.RawQuery),
				slog.String("remote_addr", r.RemoteAddr),
				slog.String("host", r.Host),
				slog.String("request_uri", r.RequestURI),
				slog.Int(constant.LogFieldStatusCode, statusCode),
				slog.Int(constant.LogFieldLatencyMS, int(time.Since(start).Milliseconds())),
				slog.Int64("response_bytes", metrics.Written),
				slog.Any("request_headers", requestHeaders),
				slog.Any("response_headers", responseHeaders),
			)
		})
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

func RequestLogContextMiddleware() middleware.Middleware {
	return func(handler middleware.Handler) middleware.Handler {
		return func(ctx context.Context, req interface{}) (interface{}, error) {
			attrs := make([]slog.Attr, 0, 2)
			if requestID := GetHeader(ctx, constant.HeaderRequestID); requestID != "" {
				attrs = append(attrs, slog.String(constant.LogFieldRequestID, requestID))
			}
			if clientIP := ClientIP(ctx); clientIP != "" {
				attrs = append(attrs, slog.String(constant.LogFieldClientIP, clientIP))
			}
			if len(attrs) > 0 {
				ctx = kratoslog.ContextWithAttrs(ctx, attrs...)
			}
			return handler(ctx, req)
		}
	}
}

func UserAuthMiddleware(authClient userv1.AuthServiceClient, realm commonenum.LoginRealm) middleware.Middleware {
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
		return func(ctx context.Context, req interface{}) (interface{}, error) {
			bearerPrefix := "Bearer "
			token := GetHeader(ctx, constant.HeaderAuthentication)
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
