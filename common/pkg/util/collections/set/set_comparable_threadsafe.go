package set

import (
	"common/pkg/util/base"
	"sync"
)

// ThreadSafeComparableSet 适用于可比较的类型
type ThreadSafeComparableSet[T comparable] struct {
	unsafeSet *ComparableSet[T]
	sync.RWMutex
}

func NewThreadSafeComparableSet[T comparable](size int, items ...T) Set[T] {
	return newThreadSafeComparableSet(size, items...)
}

func newThreadSafeComparableSet[T comparable](size int, items ...T) *ThreadSafeComparableSet[T] {
	s := make(ComparableSet[T], size)
	s.AddAll(items...)
	return &ThreadSafeComparableSet[T]{
		unsafeSet: &s,
	}
}

func (s *ThreadSafeComparableSet[T]) Add(item T) bool {
	s.Lock()
	added := s.unsafeSet.Add(item)
	s.Unlock()
	return added
}

func (s *ThreadSafeComparableSet[T]) AddAll(items ...T) int {
	s.Lock()
	added := s.unsafeSet.AddAll(items...)
	s.Unlock()
	return added
}

func (s *ThreadSafeComparableSet[T]) Remove(item T) {
	s.Lock()
	s.unsafeSet.Remove(item)
	s.Unlock()
}

func (s *ThreadSafeComparableSet[T]) RemoveAll(items ...T) {
	s.Lock()
	s.unsafeSet.RemoveAll(items...)
	s.Unlock()
}

func (s *ThreadSafeComparableSet[T]) Pop() (T, bool) {
	s.Lock()
	pop, b := s.unsafeSet.Pop()
	s.Unlock()
	return pop, b
}

func (s *ThreadSafeComparableSet[T]) PopN(n int) ([]T, int) {
	s.Lock()
	pop, n := s.unsafeSet.PopN(n)
	s.Unlock()
	return pop, n
}

func (s *ThreadSafeComparableSet[T]) ForEach(action base.Predicate[T]) {
	s.RLock()
	s.unsafeSet.ForEach(action)
	s.RUnlock()
}

func (s *ThreadSafeComparableSet[T]) Contains(item T) bool {
	s.RLock()
	contains := s.unsafeSet.Contains(item)
	s.RUnlock()
	return contains
}

func (s *ThreadSafeComparableSet[T]) ContainsAll(items ...T) bool {
	s.RLock()
	one := s.unsafeSet.ContainsAll(items...)
	s.RUnlock()
	return one
}

func (s *ThreadSafeComparableSet[T]) ContainsAny(items ...T) bool {
	s.RLock()
	contains := s.unsafeSet.ContainsAny(items...)
	s.RUnlock()
	return contains
}

func (s *ThreadSafeComparableSet[T]) Clear() {
	s.Lock()
	s.unsafeSet.Clear()
	s.Unlock()
}

func (s *ThreadSafeComparableSet[T]) Reset() {
	s.Lock()
	s.unsafeSet.Reset()
	s.Unlock()
}

func (s *ThreadSafeComparableSet[T]) Len() int {
	s.RLock()
	i := s.unsafeSet.Len()
	s.RUnlock()
	return i
}

func (s *ThreadSafeComparableSet[T]) ToSlice() []T {
	s.RLock()
	slice := s.unsafeSet.ToSlice()
	s.RUnlock()
	return slice
}

func (s *ThreadSafeComparableSet[T]) Clone() Set[T] {
	s.RLock()
	clone := s.unsafeSet.Clone().(*ComparableSet[T])
	s.RUnlock()
	return &ThreadSafeComparableSet[T]{
		unsafeSet: clone,
	}
}

func (s *ThreadSafeComparableSet[T]) MarshalJSON() ([]byte, error) {
	s.RLock()
	bytes, err := s.unsafeSet.MarshalJSON()
	s.RUnlock()
	return bytes, err
}

func (s *ThreadSafeComparableSet[T]) UnmarshalJSON(b []byte) error {
	s.Lock()
	err := s.unsafeSet.UnmarshalJSON(b)
	s.Unlock()
	return err
}

func (s *ThreadSafeComparableSet[T]) ContainsAnyElement(other Set[T]) bool {
	s.RLock()
	other.RLock()
	contains := s.unsafeSet.ContainsAnyElement(other)
	other.RUnlock()
	s.RUnlock()
	return contains
}

func (s *ThreadSafeComparableSet[T]) Union(other Set[T]) Set[T] {
	s.RLock()
	other.RLock()
	unionedSet := s.unsafeSet.Union(other).(*ComparableSet[T])
	other.RUnlock()
	s.RUnlock()
	return &ThreadSafeComparableSet[T]{
		unsafeSet: unionedSet,
	}
}

func (s *ThreadSafeComparableSet[T]) Intersection(other Set[T]) Set[T] {
	s.RLock()
	other.RLock()
	intersectedSet := s.unsafeSet.Intersection(other).(*ComparableSet[T])
	other.RUnlock()
	s.RUnlock()
	return &ThreadSafeComparableSet[T]{
		unsafeSet: intersectedSet,
	}
}

func (s *ThreadSafeComparableSet[T]) Difference(other Set[T]) Set[T] {
	s.RLock()
	other.RLock()
	differenceSet := s.unsafeSet.Difference(other).(*ComparableSet[T])
	other.RUnlock()
	s.RUnlock()
	return &ThreadSafeComparableSet[T]{
		unsafeSet: differenceSet,
	}
}

func (s *ThreadSafeComparableSet[T]) SymmetricDifference(other Set[T]) Set[T] {
	s.RLock()
	other.RLock()
	symmetricDifferenceSet := s.unsafeSet.SymmetricDifference(other).(*ComparableSet[T])
	other.RUnlock()
	s.RUnlock()
	return &ThreadSafeComparableSet[T]{
		unsafeSet: symmetricDifferenceSet,
	}
}

func (s *ThreadSafeComparableSet[T]) IsEmpty() bool {
	return s.Len() == 0
}

func (s *ThreadSafeComparableSet[T]) Equal(other Set[T]) bool {
	s.RLock()
	other.RLock()
	equal := s.unsafeSet.Equal(other)
	other.RUnlock()
	s.RUnlock()
	return equal
}

func (s *ThreadSafeComparableSet[T]) IsSubset(other Set[T]) bool {
	s.RLock()
	other.RLock()
	isSubset := s.unsafeSet.IsSubset(other)
	other.RUnlock()
	s.RUnlock()
	return isSubset
}

func (s *ThreadSafeComparableSet[T]) IsProperSubset(other Set[T]) bool {
	return s.Len() < other.Len() && s.IsSubset(other)
}

func (s *ThreadSafeComparableSet[T]) IsSuperset(other Set[T]) bool {
	return other.IsSubset(s)
}

func (s *ThreadSafeComparableSet[T]) IsProperSuperset(other Set[T]) bool {
	return s.Len() > other.Len() && s.IsSuperset(other)
}

func (s *ThreadSafeComparableSet[T]) Lock() {
	s.Lock()
}

func (s *ThreadSafeComparableSet[T]) Unlock() {
	s.Unlock()
}

func (s *ThreadSafeComparableSet[T]) RLock() {
	s.RLock()
}

func (s *ThreadSafeComparableSet[T]) RUnlock() {
	s.RUnlock()
}
