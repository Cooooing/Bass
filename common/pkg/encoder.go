package pkg

import (
	"encoding/json"
	nhttp "net/http"

	"github.com/go-kratos/kratos/v2/errors"
	"github.com/go-kratos/kratos/v2/transport/http"
)

// HttpResponseEncoder 自定义的响应编码器
func HttpResponseEncoder(w http.ResponseWriter, r *http.Request, data any) error {
	// 如果 data 是 error 类型，可以在这里做判断
	if err, ok := data.(error); ok && err != nil {
		// 使用 kratos 标准错误
		se := errors.FromError(err)
		w.WriteHeader(int(se.Code))
		result := NewResult[any](int(se.Code), se.Message, nil)
		w.Header().Set("Content-Type", "application/json")
		return json.NewEncoder(w).Encode(result)
	}

	// 正常返回时包装为 Result 结构
	result := SuccessData(data)
	w.Header().Set("Content-Type", "application/json")
	return json.NewEncoder(w).Encode(result)
}

// HttpErrorEncoder 自定义错误响应编码器
func HttpErrorEncoder(w http.ResponseWriter, r *http.Request, err error) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(nhttp.StatusInternalServerError)
	res := Error(err)
	_ = json.NewEncoder(w).Encode(res)
}
