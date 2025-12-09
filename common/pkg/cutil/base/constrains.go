package base

// Signed 有符号类型
type Signed interface {
	~int | ~int8 | ~int16 | ~int32 | ~int64
}

// Unsigned 无符号类型
type Unsigned interface {
	~uint | ~uint8 | ~uint16 | ~uint32 | ~uint64 | ~uintptr
}

// Integer 整型
type Integer interface {
	Signed | Unsigned
}

// Float 浮点型
type Float interface {
	~float32 | ~float64
}

// Complex 复数
type Complex interface {
	~complex64 | ~complex128
}
