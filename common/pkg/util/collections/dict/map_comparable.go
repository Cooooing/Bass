package dict

import (
	"common/pkg/util/base"
	"common/pkg/util/stream"
	"context"
	"reflect"
)

type ComparableMap[K comparable, V any] map[K]V

func New[K comparable, V any](size int) Map[K, V] {
	return NewComparableMap[K, V](size)
}

func NewComparableMap[K comparable, V any](size int) Map[K, V] {
	m := make(ComparableMap[K, V], size)
	return &m
}

func (m *ComparableMap[K, V]) Set(key K, value V) {
	(*m)[key] = value
}

func (m *ComparableMap[K, V]) Get(key K) (V, bool) {
	v, ok := (*m)[key]
	return v, ok
}

func (m *ComparableMap[K, V]) Remove(key K) {
	delete(*m, key)
}

func (m *ComparableMap[K, V]) RemoveAll(keys ...K) {
	for _, key := range keys {
		delete(*m, key)
	}
}

func (m *ComparableMap[K, V]) Pop(key K) (V, bool) {
	v, ok := (*m)[key]
	if ok {
		delete(*m, key)
	}
	return v, ok
}

func (m *ComparableMap[K, V]) Contains(key K) bool {
	_, ok := (*m)[key]
	return ok
}

func (m *ComparableMap[K, V]) ContainsAll(keys ...K) bool {
	for _, key := range keys {
		if _, ok := (*m)[key]; !ok {
			return false
		}
	}
	return true
}

func (m *ComparableMap[K, V]) ContainsAny(keys ...K) bool {
	for _, key := range keys {
		if _, ok := (*m)[key]; ok {
			return true
		}
	}
	return false
}

func (m *ComparableMap[K, V]) Keys() []K {
	keys := make([]K, 0, m.Len())
	for k, _ := range *m {
		keys = append(keys, k)
	}
	return keys
}

func (m *ComparableMap[K, V]) Values() []V {
	values := make([]V, 0, m.Len())
	for _, v := range *m {
		values = append(values, v)
	}
	return values
}

func (m *ComparableMap[K, V]) Entries() []*Entry[K, V] {
	entries := make([]*Entry[K, V], 0, m.Len())
	for k, v := range *m {
		entries = append(entries, &Entry[K, V]{
			Key:   k,
			Value: v,
		})
	}
	return entries
}

func (m *ComparableMap[K, V]) Len() int {
	return len(*m)
}

func (m *ComparableMap[K, V]) IsEmpty() bool {
	return m.Len() == 0
}

func (m *ComparableMap[K, V]) Merge(other Map[K, V]) {
	other.Foreach(func(e *Entry[K, V]) bool {
		m.Set(e.Key, e.Value)
		return true
	})
}

func (m *ComparableMap[K, V]) Equal(other Map[K, V]) bool {
	if m.Len() != other.Len() {
		return false
	}
	for k, v := range *m {
		if ov, ok := other.Get(k); !ok || !reflect.DeepEqual(ov, v) {
			return false
		}
	}
	return true
}

func (m *ComparableMap[K, V]) EqualFunc(other Map[K, V], fn base.Equator[V]) bool {
	if m.Len() != other.Len() {
		return false
	}
	for k, v := range *m {
		if ov, ok := other.Get(k); !ok || !fn(ov, v) {
			return false
		}
	}
	return true
}

func (m *ComparableMap[K, V]) Clone() Map[K, V] {
	clone := NewComparableMap[K, V](m.Len())
	for k, v := range *m {
		clone.Set(k, v)
	}
	return clone
}

func (m *ComparableMap[K, V]) Clear() {
	for k, _ := range *m {
		delete(*m, k)
	}
}

func (m *ComparableMap[K, V]) Reset() {
	*m = make(ComparableMap[K, V])
}

func (m *ComparableMap[K, V]) Foreach(action base.Predicate[*Entry[K, V]]) {
	for item := range *m {
		action(&Entry[K, V]{
			Key:   item,
			Value: (*m)[item],
		})
	}
}

func (m *ComparableMap[K, V]) Stream(ctx context.Context) stream.Stream[*Entry[K, V]] {
	return stream.OfBlock[*Entry[K, V]](ctx, m.Entries()...)
}

func (m *ComparableMap[K, V]) Lock() {}

func (m *ComparableMap[K, V]) Unlock() {}

func (m *ComparableMap[K, V]) RLock() {}

func (m *ComparableMap[K, V]) RUnlock() {}
