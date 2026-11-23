package middleware

import (
	cerrors "common/proto/gen/common/errors/v1"
	"context"
	"log/slog"
	"net"
	"net/http"
	"strings"
	"time"

	"common/pkg/apperror"
	"common/pkg/client"
	"common/pkg/constant"
	"common/pkg/server"

	"github.com/felixge/httpsnoop"
	kratoslog "github.com/go-kratos/kratos/v3/log"
	"github.com/go-kratos/kratos/v3/middleware"
	kratoshttp "github.com/go-kratos/kratos/v3/transport/http"
	"github.com/google/uuid"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/codes"
	"go.opentelemetry.io/otel/propagation"
	oteltrace "go.opentelemetry.io/otel/trace"
)

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
				for item := range strings.SplitSeq(r.Header.Get(header), ",") {
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

func HTTPCORSFilter() kratoshttp.FilterFunc {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			w.Header().Set("Access-Control-Allow-Origin", "*")
			w.Header().Set("Access-Control-Allow-Methods", "GET,POST,PUT,PATCH,DELETE,OPTIONS")
			w.Header().Set("Access-Control-Allow-Headers", "Authorization,Content-Type,X-Request-ID,X-Trace-ID,X-Timestamp,X-Nonce,X-Bass-App-Name,X-Bass-App-Version")
			w.Header().Set("Access-Control-Expose-Headers", "X-Request-ID,X-Trace-ID")
			w.Header().Set("Access-Control-Max-Age", "86400")
			if r.Method == http.MethodOptions {
				w.WriteHeader(http.StatusNoContent)
				return
			}
			next.ServeHTTP(w, r)
		})
	}
}

func NonceMiddleware(redisClient *client.RedisClient, mode string) middleware.Middleware {
	return func(handler middleware.Handler) middleware.Handler {
		return func(ctx context.Context, req any) (any, error) {
			if mode == constant.Dev {
				return handler(ctx, req)
			}
			nonce := server.GetHeader(ctx, constant.HeaderNonce)
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
		return func(ctx context.Context, req any) (any, error) {
			attrs := make([]slog.Attr, 0, 2)
			if requestID := server.GetHeader(ctx, constant.HeaderRequestID); requestID != "" {
				attrs = append(attrs, slog.String(constant.LogFieldRequestID, requestID))
			}
			if clientIP := server.ClientIP(ctx); clientIP != "" {
				attrs = append(attrs, slog.String(constant.LogFieldClientIP, clientIP))
			}
			if len(attrs) > 0 {
				ctx = kratoslog.ContextWithAttrs(ctx, attrs...)
			}
			return handler(ctx, req)
		}
	}
}
