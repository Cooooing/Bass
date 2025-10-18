package pkg

import (
	"net/http"
	"time"
)

type Result[T any] struct {
	Code int        `json:"code"`
	Msg  string     `json:"msg"`
	Data T          `json:"data"`
	Time *time.Time `json:"time"`
}

func NewResult[T any](code int, msg string, data T) *Result[T] {
	now := time.Now()
	return &Result[T]{
		Code: code,
		Msg:  msg,
		Data: data,
		Time: &now,
	}
}

func Success() *Result[any] {
	return NewResult[any](http.StatusOK, "success", nil)
}

func SuccessData[T any](data T) *Result[T] {
	return NewResult(http.StatusOK, "success", data)
}

func BadRequest() *Result[any] {
	return NewResult[any](http.StatusBadRequest, http.StatusText(http.StatusBadRequest), nil)
}

func Forbidden() *Result[any] {
	return NewResult[any](http.StatusForbidden, http.StatusText(http.StatusForbidden), nil)
}

func TooManyRequests() *Result[any] {
	return NewResult[any](http.StatusTooManyRequests, http.StatusText(http.StatusTooManyRequests), nil)
}

func Error(err error) *Result[any] {
	return NewResult[any](http.StatusInternalServerError, err.Error(), nil)
}
