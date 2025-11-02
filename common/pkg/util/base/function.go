package base

type (
	Function[T any, R any]   func(T) R
	Consumer[T any]          func(T)
	Supplier[T any]          func() T
	Predicate[T any]         func(T) bool
	Equator[T any]           func(T, T) bool
	Comparator[T any]        func(T, T) int
	UnaryOperator[T any]     func(T) T
	BinaryOperator[T any]    func(T, T) T
	BiFunction[T any, R any] func(T, R) R
)
