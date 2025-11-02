package base

import "reflect"

// Ptr 返回一个指向 v 的指针
func Ptr[T any](v T) *T {
	return &v
}

// If 模拟三元运算符：condition ? trueValue : falseValue
func If[T any](condition bool, trueValue, falseValue T) T {
	if condition {
		return trueValue
	}
	return falseValue
}

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
	}
	return false
}

func IsNotNil(v any) bool {
	return !IsNil(v)
}

// IfNilDefault 如果 v 是 nil，返回 defaultValue，否则返回 v
func IfNilDefault[T any](v T, defaultValue T) T {
	if IsNil(v) {
		return defaultValue
	}
	return v
}
