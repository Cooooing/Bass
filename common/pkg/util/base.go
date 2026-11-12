package util

import "reflect"

// IsNil 判断一个值是否为 nil
// 注意：只有引用类型（chan, func, interface, map, pointer, slice）才可能为 nil
func IsNil(v any) bool {
	if v == nil {
		return true
	}
	rv := reflect.ValueOf(v)
	switch rv.Kind() {
	case reflect.Chan, reflect.Func, reflect.Interface, reflect.Map, reflect.Pointer, reflect.Slice:
		return rv.IsNil()
	default:
	}
	return false
}

// OrDefault 如果 v 为 nil，返回默认值，否则返回 v。
func OrDefault[T any](v T, defaultValue T) T {
	if IsNil(v) {
		return defaultValue
	}
	return v
}

// DerefOrDefault 如果指针 v 为 nil，返回默认值，否则返回 *v。
func DerefOrDefault[T any](v *T, defaultValue T) T {
	if v == nil {
		return defaultValue
	}
	return *v
}
