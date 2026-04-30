package util

import (
	"common/pkg/constant"
	"context"
)

func SetContextValue[T any](ctx context.Context, key constant.CtxKey, value T) context.Context {
	return context.WithValue(ctx, key, value)
}

func GetContextValue[T any](ctx context.Context, key constant.CtxKey) (T, bool) {
	var zero T
	value := ctx.Value(key)
	if IsNil(value) {
		return zero, false
	}
	t, ok := value.(T)
	if !ok {
		return zero, false
	}
	return t, true
}
