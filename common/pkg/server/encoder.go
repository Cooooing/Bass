package server

import (
	"bytes"
	"common/pkg/apperror"
	"common/pkg/constant"
	"common/proto/gen/common"
	cerrors "common/proto/gen/common/errors"
	"encoding/json"
	"encoding/json/jsontext"
	jsonv2 "encoding/json/v2"
	"net/http"
	"runtime/debug"

	kratoslog "github.com/go-kratos/kratos/v3/log"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/codes"
	oteltrace "go.opentelemetry.io/otel/trace"
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
		result := NewResult[json.RawMessage](int(cerrors.BusinessErrorCode_BUSINESS_ERROR_CODE_SUCCESS), "success", b)
		return encodeJSONResponse(w, http.StatusOK, result)
	}

	if w.Header().Get("Content-Type") != "" && w.Header().Get("Content-Type") != "application/json" {
		return nil
	}

	result := NewResult[any](int(cerrors.BusinessErrorCode_BUSINESS_ERROR_CODE_SUCCESS), "success", data)
	return encodeJSONResponse(w, http.StatusOK, result)
}

// encodeJSONResponse 使用 encoding/json/v2 编码响应体，并以 statusCode 写出。
//
// 采用 v2 原生语义：
//   - <, >, & 按原字符输出，不做 \u003c 形式转义
//   - 字符串必须为合法 UTF-8，非法字节返回错误
//
// 响应先编码到内存，成功后再写状态码，避免 v2 编码中途失败时留下残缺 JSON。
func encodeJSONResponse(w http.ResponseWriter, statusCode int, result any) error {
	var buf bytes.Buffer
	if err := jsonv2.MarshalEncode(jsontext.NewEncoder(&buf), result); err != nil {
		return err
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	_, err := w.Write(buf.Bytes())
	return err
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
		logArgs := []any{
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
		}
		if statusCode >= http.StatusInternalServerError {
			logArgs = append(logArgs, "stack", string(debug.Stack()))
		}
		kratoslog.ErrorContext(r.Context(), "http error", logArgs...)

		res := NewResult[any](int(businessCode), message, errorData)

		_ = encodeJSONResponse(w, statusCode, res)
	}
}
