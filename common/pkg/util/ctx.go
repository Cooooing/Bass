package util

import (
	"common/pkg/constant"
	"common/pkg/util/base"
	"context"
)

func SetContextValue[T any](ctx context.Context, key constant.CtxKey, value T) context.Context {
	return context.WithValue(ctx, key, value)
}

func GetContextValue[T any](ctx context.Context, key constant.CtxKey) (T, bool) {
	value := ctx.Value(key)
	if base.IsNil(value) {
		var t T
		return t, false
	}
	return value.(T), true
}

func MustGetContextValue[T any](ctx context.Context, key constant.CtxKey) T {
	if value, ok := ctx.Value(key).(T); ok {
		return value
	}
	panic("context value not found")
}
