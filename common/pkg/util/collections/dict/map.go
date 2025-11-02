package dict

import (
	"common/pkg/util/base"
	"common/pkg/util/collections"
	"common/pkg/util/stream"
	"context"
)

// Map 接口，定义字典应具备的基本操作
type Map[K comparable, V any] interface {

	// 基础操作

	// Set 设置键值对
	Set(key K, value V)
	// Get 获取键对应的值
	Get(key K) (V, bool)
	// Remove 删除键
	Remove(key K)
	// RemoveAll 删除多个键
	RemoveAll(keys ...K)
	// Pop 删除键并返回键对应的值
	Pop(key K) (V, bool)
	// Contains 判断键是否存在
	Contains(key K) bool
	// ContainsAll 判断所有键是否存在
	ContainsAll(keys ...K) bool
	// ContainsAny 判断任意键是否存在
	ContainsAny(keys ...K) bool

	// 集合操作

	// Keys 获取所有键集合
	Keys() []K
	// Values 获取所有值集合
	Values() []V
	// Entries 获取所有键值对集合
	Entries() []*Entry[K, V]
	// Len 获取字典长度
	Len() int
	// IsEmpty 判断字典是否为空
	IsEmpty() bool
	// Merge 合并字典
	Merge(other Map[K, V])
	// Equal 判断字典是否相等
	Equal(other Map[K, V]) bool
	// EqualFunc 判断字典是否相等，使用自定义比较函数
	EqualFunc(other Map[K, V], fn base.Equator[V]) bool
	// Clone 创建字典的副本
	Clone() Map[K, V]
	// Clear 清空字典
	Clear()
	// Reset 重置字典
	Reset()
	// Foreach 遍历字典
	Foreach(action base.Predicate[*Entry[K, V]])
	// Stream 获取字典的流
	Stream(ctx context.Context) stream.Stream[*Entry[K, V]]

	collections.Lockable
}

type Entry[T any, R any] struct {
	Key   T
	Value R
}
