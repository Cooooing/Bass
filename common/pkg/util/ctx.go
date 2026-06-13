package util

import "context"

func SetContextValue[T any](ctx context.Context, key any, value T) context.Context {
	return context.WithValue(ctx, key, value)
}

func GetContextValue[T any](ctx context.Context, key any) (T, bool) {
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
