package dict

import (
	"common/pkg/util/base"
	"common/pkg/util/stream"
	"context"
	"sync"
)

type ThreadSafeMap[K comparable, V any] struct {
	unsafeMap *ComparableMap[K, V]
	sync.RWMutex
}

func NewSafeMap[K comparable, V any](size int) Map[K, V] {
	unsafeMap := NewComparableMap[K, V](size)
	return &ThreadSafeMap[K, V]{
		unsafeMap: unsafeMap.(*ComparableMap[K, V]),
	}
}

func (m *ThreadSafeMap[K, V]) Set(key K, value V) {
	m.Lock()
	m.unsafeMap.Set(key, value)
	m.Unlock()
}

func (m *ThreadSafeMap[K, V]) Get(key K) (V, bool) {
	m.RLock()
	v, ok := m.unsafeMap.Get(key)
	m.RUnlock()
	return v, ok
}

func (m *ThreadSafeMap[K, V]) Remove(key K) {
	m.Lock()
	m.Remove(key)
	m.Unlock()
}

func (m *ThreadSafeMap[K, V]) RemoveAll(keys ...K) {
	m.Lock()
	m.RemoveAll(keys...)
	m.Unlock()
}

func (m *ThreadSafeMap[K, V]) Pop(key K) (V, bool) {
	m.Lock()
	v, ok := m.unsafeMap.Pop(key)
	m.Unlock()
	return v, ok
}

func (m *ThreadSafeMap[K, V]) Contains(key K) bool {
	m.RLock()
	ok := m.Contains(key)
	m.RUnlock()
	return ok
}

func (m *ThreadSafeMap[K, V]) ContainsAll(keys ...K) bool {
	m.RLock()
	ok := m.ContainsAll(keys...)
	m.RUnlock()
	return ok
}

func (m *ThreadSafeMap[K, V]) ContainsAny(keys ...K) bool {
	m.RLock()
	ok := m.ContainsAny(keys...)
	m.RUnlock()
	return ok
}

func (m *ThreadSafeMap[K, V]) Keys() []K {
	m.RLock()
	keys := m.unsafeMap.Keys()
	m.RUnlock()
	return keys
}

func (m *ThreadSafeMap[K, V]) Values() []V {
	m.Lock()
	values := m.unsafeMap.Values()
	m.Unlock()
	return values
}

func (m *ThreadSafeMap[K, V]) Entries() []*Entry[K, V] {
	m.RLock()
	entries := m.unsafeMap.Entries()
	m.RUnlock()
	return entries
}

func (m *ThreadSafeMap[K, V]) Len() int {
	m.RLock()
	l := m.unsafeMap.Len()
	m.Unlock()
	return l
}

func (m *ThreadSafeMap[K, V]) IsEmpty() bool {
	return m.Len() == 0
}

func (m *ThreadSafeMap[K, V]) Merge(other Map[K, V]) {
	m.Lock()
	other.RLock()
	other.Foreach(func(e *Entry[K, V]) bool {
		m.Set(e.Key, e.Value)
		return true
	})
	other.RUnlock()
	m.Unlock()
}

func (m *ThreadSafeMap[K, V]) Equal(other Map[K, V]) bool {
	m.RLock()
	other.RLock()
	equal := m.Equal(other)
	other.RUnlock()
	m.Unlock()
	return equal
}

func (m *ThreadSafeMap[K, V]) EqualFunc(other Map[K, V], fn base.Equator[V]) bool {
	m.RLock()
	other.RLock()
	equal := m.EqualFunc(other, fn)
	other.RUnlock()
	m.Unlock()
	return equal
}

func (m *ThreadSafeMap[K, V]) Clone() Map[K, V] {
	m.RLock()
	clone := m.unsafeMap.Clone().(*ComparableMap[K, V])
	m.RUnlock()
	return &ThreadSafeMap[K, V]{
		unsafeMap: clone,
	}

}

func (m *ThreadSafeMap[K, V]) Clear() {
	m.Lock()
	m.unsafeMap.Clear()
	m.Unlock()
}

func (m *ThreadSafeMap[K, V]) Reset() {
	m.Lock()
	m.unsafeMap.Reset()
	m.Unlock()
}

func (m *ThreadSafeMap[K, V]) Foreach(action base.Predicate[*Entry[K, V]]) {
	m.RLock()
	m.unsafeMap.Foreach(action)
	m.RUnlock()
}

func (m *ThreadSafeMap[K, V]) Stream(ctx context.Context) stream.Stream[*Entry[K, V]] {
	return stream.OfBlock[*Entry[K, V]](ctx, m.Entries()...)
}

func (m *ThreadSafeMap[K, V]) Lock() {
	m.Lock()
}

func (m *ThreadSafeMap[K, V]) Unlock() {
	m.Unlock()
}

func (m *ThreadSafeMap[K, V]) RLock() {
	m.RLock()
}

func (m *ThreadSafeMap[K, V]) RUnlock() {
	m.RUnlock()
}
