package set

import (
	"common/pkg/util/base"
	"sync"
)

// ThreadSafeKeySet 适用于可比较的类型
type ThreadSafeKeySet[T Keyer] struct {
	unsafeSet *KeySet[T]
	sync.RWMutex
}

func NewThreadSafeKeySet[T Keyer](size int, items ...T) Set[T] {
	return newThreadSafeKeySet(size, items...)
}

func newThreadSafeKeySet[T Keyer](size int, items ...T) *ThreadSafeKeySet[T] {
	s := make(KeySet[T], size)
	s.AddAll(items...)
	return &ThreadSafeKeySet[T]{
		unsafeSet: &s,
	}
}

func (s *ThreadSafeKeySet[T]) Add(item T) bool {
	s.Lock()
	added := s.unsafeSet.Add(item)
	s.Unlock()
	return added
}

func (s *ThreadSafeKeySet[T]) AddAll(items ...T) int {
	s.Lock()
	added := s.unsafeSet.AddAll(items...)
	s.Unlock()
	return added
}

func (s *ThreadSafeKeySet[T]) Remove(item T) {
	s.Lock()
	s.unsafeSet.Remove(item)
	s.Unlock()
}

func (s *ThreadSafeKeySet[T]) RemoveAll(items ...T) {
	s.Lock()
	s.unsafeSet.RemoveAll(items...)
	s.Unlock()
}

func (s *ThreadSafeKeySet[T]) Pop() (T, bool) {
	s.Lock()
	pop, b := s.unsafeSet.Pop()
	s.Unlock()
	return pop, b
}

func (s *ThreadSafeKeySet[T]) PopN(n int) ([]T, int) {
	s.Lock()
	pop, n := s.unsafeSet.PopN(n)
	s.Unlock()
	return pop, n
}

func (s *ThreadSafeKeySet[T]) ForEach(action base.Predicate[T]) {
	s.RLock()
	s.unsafeSet.ForEach(action)
	s.RUnlock()
}

func (s *ThreadSafeKeySet[T]) Contains(item T) bool {
	s.RLock()
	contains := s.unsafeSet.Contains(item)
	s.RUnlock()
	return contains
}

func (s *ThreadSafeKeySet[T]) ContainsAll(items ...T) bool {
	s.RLock()
	one := s.unsafeSet.ContainsAll(items...)
	s.RUnlock()
	return one
}

func (s *ThreadSafeKeySet[T]) ContainsAny(items ...T) bool {
	s.RLock()
	contains := s.unsafeSet.ContainsAny(items...)
	s.RUnlock()
	return contains
}

func (s *ThreadSafeKeySet[T]) Clear() {
	s.Lock()
	s.unsafeSet.Clear()
	s.Unlock()
}

func (s *ThreadSafeKeySet[T]) Reset() {
	s.Lock()
	s.unsafeSet.Reset()
	s.Unlock()
}

func (s *ThreadSafeKeySet[T]) Len() int {
	s.RLock()
	i := s.unsafeSet.Len()
	s.RUnlock()
	return i
}

func (s *ThreadSafeKeySet[T]) ToSlice() []T {
	s.RLock()
	slice := s.unsafeSet.ToSlice()
	s.RUnlock()
	return slice
}

func (s *ThreadSafeKeySet[T]) Clone() Set[T] {
	s.RLock()
	clone := s.unsafeSet.Clone().(*KeySet[T])
	s.RUnlock()
	return &ThreadSafeKeySet[T]{
		unsafeSet: clone,
	}
}

func (s *ThreadSafeKeySet[T]) MarshalJSON() ([]byte, error) {
	s.RLock()
	bytes, err := s.unsafeSet.MarshalJSON()
	s.RUnlock()
	return bytes, err
}

func (s *ThreadSafeKeySet[T]) UnmarshalJSON(b []byte) error {
	s.Lock()
	err := s.unsafeSet.UnmarshalJSON(b)
	s.Unlock()
	return err
}

func (s *ThreadSafeKeySet[T]) ContainsAnyElement(other Set[T]) bool {
	s.RLock()
	other.RLock()
	contains := s.unsafeSet.ContainsAnyElement(other)
	other.RUnlock()
	s.RUnlock()
	return contains
}

func (s *ThreadSafeKeySet[T]) Union(other Set[T]) Set[T] {
	s.RLock()
	other.RLock()
	unionedSet := s.unsafeSet.Union(other).(*KeySet[T])
	other.RUnlock()
	s.RUnlock()
	return &ThreadSafeKeySet[T]{
		unsafeSet: unionedSet,
	}
}

func (s *ThreadSafeKeySet[T]) Intersection(other Set[T]) Set[T] {
	s.RLock()
	other.RLock()
	intersectedSet := s.unsafeSet.Intersection(other).(*KeySet[T])
	other.RUnlock()
	s.RUnlock()
	return &ThreadSafeKeySet[T]{
		unsafeSet: intersectedSet,
	}
}

func (s *ThreadSafeKeySet[T]) Difference(other Set[T]) Set[T] {
	s.RLock()
	other.RLock()
	differenceSet := s.unsafeSet.Difference(other).(*KeySet[T])
	other.RUnlock()
	s.RUnlock()
	return &ThreadSafeKeySet[T]{
		unsafeSet: differenceSet,
	}
}

func (s *ThreadSafeKeySet[T]) SymmetricDifference(other Set[T]) Set[T] {
	s.RLock()
	other.RLock()
	symmetricDifferenceSet := s.unsafeSet.SymmetricDifference(other).(*KeySet[T])
	other.RUnlock()
	s.RUnlock()
	return &ThreadSafeKeySet[T]{
		unsafeSet: symmetricDifferenceSet,
	}
}

func (s *ThreadSafeKeySet[T]) IsEmpty() bool {
	return s.Len() == 0
}

func (s *ThreadSafeKeySet[T]) Equal(other Set[T]) bool {
	s.RLock()
	other.RLock()
	equal := s.unsafeSet.Equal(other)
	other.RUnlock()
	s.RUnlock()
	return equal
}

func (s *ThreadSafeKeySet[T]) IsSubset(other Set[T]) bool {
	s.RLock()
	other.RLock()
	isSubset := s.unsafeSet.IsSubset(other)
	other.RUnlock()
	s.RUnlock()
	return isSubset
}

func (s *ThreadSafeKeySet[T]) IsProperSubset(other Set[T]) bool {
	return s.Len() < other.Len() && s.IsSubset(other)
}

func (s *ThreadSafeKeySet[T]) IsSuperset(other Set[T]) bool {
	return other.IsSubset(s)
}

func (s *ThreadSafeKeySet[T]) IsProperSuperset(other Set[T]) bool {
	return s.Len() > other.Len() && s.IsSuperset(other)
}

func (s *ThreadSafeKeySet[T]) Lock() {
	s.Lock()
}

func (s *ThreadSafeKeySet[T]) Unlock() {
	s.Unlock()
}

func (s *ThreadSafeKeySet[T]) RLock() {
	s.RLock()
}

func (s *ThreadSafeKeySet[T]) RUnlock() {
	s.RUnlock()
}
