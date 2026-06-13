package server

import "time"

type Result[T any] struct {
	Code    int        `json:"code"`
	Message string     `json:"message"`
	Data    T          `json:"data"`
	Time    *time.Time `json:"time"`
}

func NewResult[T any](code int, message string, data T) *Result[T] {
	return &Result[T]{
		Code:    code,
		Message: message,
		Data:    data,
		Time:    new(time.Now().UTC()),
	}
}
