package pkg

import (
	"encoding/json"
	"net/http"

	"github.com/go-kratos/kratos/v2/errors"
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
