package set

import (
	"common/pkg/util/base"
	"common/pkg/util/collections"
)

// Set 接口，定义集合应具备的基本操作
type Set[T any] interface {

	// 基础操作

	// Add 添加一个元素
	Add(item T) bool
	// AddAll 添加多个元素
	AddAll(items ...T) int
	// Remove 删除一个元素
	Remove(item T)
	// RemoveAll 删除多个元素
	RemoveAll(items ...T)
	// Pop 删除并返回一个元素
	Pop() (T, bool)
	// PopN 删除并返回n个元素
	PopN(n int) ([]T, int)
	// ForEach 迭代集合
	ForEach(consumer base.Predicate[T])

	// Contains 判断是否存在元素
	Contains(item T) bool
	// ContainsAll 判断是否存在所有元素
	ContainsAll(items ...T) bool
	// ContainsAny 判断是否包含任意元素
	ContainsAny(items ...T) bool

	// Clear 清空集合
	Clear()
	// Reset 重置集合
	Reset()
	// Len 返回元素数量
	Len() int
	// ToSlice 转换为切片
	ToSlice() []T
	// Clone 创建一个副本
	Clone() Set[T]
	// MarshalJSON 序列化为JSON数据
	MarshalJSON() ([]byte, error)
	// UnmarshalJSON 解析JSON数据
	UnmarshalJSON(b []byte) error

	// 集合运算

	// ContainsAnyElement 判断是否包含传入Set的任意元素
	ContainsAnyElement(other Set[T]) bool
	// Union 返回并集
	Union(other Set[T]) Set[T]
	// Intersection 返回交集
	Intersection(other Set[T]) Set[T]
	// Difference 返回差集
	Difference(other Set[T]) Set[T]
	// SymmetricDifference 返回对称差集 (只在一个集合里出现的元素)
	SymmetricDifference(other Set[T]) Set[T]

	// 关系判断

	// IsEmpty 集合是否为空
	IsEmpty() bool
	// Equal 判断两个集合是否相等
	Equal(other Set[T]) bool
	// IsSubset 判断是否为子集（可以相等）
	IsSubset(other Set[T]) bool
	// IsProperSubset 判断是否为真子集（严格小于）
	IsProperSubset(other Set[T]) bool
	// IsSuperset 判断是否为超集（可以相等）
	IsSuperset(other Set[T]) bool
	// IsProperSuperset 判断是否为真超集（严格大于）
	IsProperSuperset(other Set[T]) bool

	collections.Lockable
}

// Keyer 接口，用于自定义集合元素的键值
type Keyer interface {
	// Key 获取键值
	Key() string
}
