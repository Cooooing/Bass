package set

import (
	"common/pkg/util/base"
	"encoding/json"
	"fmt"
	"strings"
)

// ComparableSet 适用于可比较的类型的Set集合，非线程安全
type ComparableSet[T comparable] map[T]struct{}

func New[T comparable](size int) Set[T] {
	return NewComparableSet[T](size)
}

func NewComparableSet[T comparable](size int, items ...T) Set[T] {
	s := make(ComparableSet[T], size)
	s.AddAll(items...)
	return &s
}

func (s *ComparableSet[T]) Add(item T) bool {
	prevLen := s.Len()
	(*s)[item] = struct{}{}
	return prevLen != s.Len()
}

func (s *ComparableSet[T]) AddAll(items ...T) int {
	prevLen := s.Len()
	for _, item := range items {
		(*s)[item] = struct{}{}
	}
	return prevLen - s.Len()
}

func (s *ComparableSet[T]) Remove(item T) {
	delete(*s, item)
}

func (s *ComparableSet[T]) RemoveAll(items ...T) {
	for _, item := range items {
		delete(*s, item)
	}
}

func (s *ComparableSet[T]) Pop() (T, bool) {
	for item := range *s {
		delete(*s, item)
		return item, true
	}
	var v T
	return v, false
}

func (s *ComparableSet[T]) PopN(n int) ([]T, int) {
	if n <= 0 || len(*s) == 0 {
		return make([]T, 0), 0
	}
	sn := s.Len()
	if n > sn {
		n = sn
	}

	count := 0
	items := make([]T, 0, sn)
	for item := range *s {
		if count >= n {
			break
		}
		delete(*s, item)
		items = append(items, item)
		count++
	}
	return items, count
}

func (s *ComparableSet[T]) ForEach(action base.Predicate[T]) {
	for item := range *s {
		if !action(item) {
			return
		}
	}
}

func (s *ComparableSet[T]) Contains(item T) bool {
	if _, ok := (*s)[item]; ok {
		return true
	}
	return false
}

func (s *ComparableSet[T]) ContainsAll(items ...T) bool {
	for _, item := range items {
		if _, ok := (*s)[item]; !ok {
			return false
		}
	}
	return true
}

func (s *ComparableSet[T]) ContainsAny(items ...T) bool {
	for _, item := range items {
		if _, ok := (*s)[item]; ok {
			return true
		}
	}
	return false
}

func (s *ComparableSet[T]) Clear() {
	for key := range *s {
		delete(*s, key)
	}
}

func (s *ComparableSet[T]) Reset() {
	*s = make(ComparableSet[T])
}

func (s *ComparableSet[T]) Len() int {
	return len(*s)
}

func (s *ComparableSet[T]) ToSlice() []T {
	keys := make([]T, 0, s.Len())
	for elem := range *s {
		keys = append(keys, elem)
	}
	return keys
}

func (s *ComparableSet[T]) Clone() Set[T] {
	clonedSet := NewComparableSet[T](s.Len())
	for elem := range *s {
		clonedSet.Add(elem)
	}
	return clonedSet
}

func (s *ComparableSet[T]) MarshalJSON() ([]byte, error) {
	items := make([]string, 0, s.Len())
	for elem := range *s {
		b, err := json.Marshal(elem)
		if err != nil {
			return nil, err
		}
		items = append(items, string(b))
	}
	return []byte(fmt.Sprintf("[%s]", strings.Join(items, ","))), nil
}

func (s *ComparableSet[T]) UnmarshalJSON(b []byte) error {
	var i []T
	err := json.Unmarshal(b, &i)
	if err != nil {
		return err
	}
	s.AddAll(i...)
	return nil
}

func (s *ComparableSet[T]) ContainsAnyElement(other Set[T]) bool {
	if s.Len() < other.Len() {
		for elem := range *s {
			if other.ContainsAll(elem) {
				return true
			}
		}
	} else {
		found := false
		other.ForEach(func(t T) bool {
			if s.ContainsAll(t) {
				found = true
				return false
			}
			return true
		})
		return found
	}
	return false
}

func (s *ComparableSet[T]) Union(other Set[T]) Set[T] {
	n := s.Len()
	if other.Len() > n {
		n = other.Len()
	}
	unionedSet := make(ComparableSet[T], n)

	for elem := range *s {
		unionedSet.Add(elem)
	}
	other.ForEach(func(t T) bool {
		unionedSet.Add(t)
		return true
	})
	return &unionedSet
}

func (s *ComparableSet[T]) Intersection(other Set[T]) Set[T] {
	n := s.Len()
	if other.Len() < n {
		n = other.Len()
	}
	intersectedSet := make(ComparableSet[T], n)
	if s.Len() < other.Len() {
		for elem := range *s {
			if other.ContainsAll(elem) {
				intersectedSet.Add(elem)
			}
		}
	} else {
		other.ForEach(func(t T) bool {
			if s.ContainsAll(t) {
				intersectedSet.Add(t)
			}
			return true
		})
	}
	return &intersectedSet
}

func (s *ComparableSet[T]) Difference(other Set[T]) Set[T] {
	diffSet := make(ComparableSet[T], s.Len())
	for elem := range *s {
		if !other.ContainsAll(elem) {
			diffSet.Add(elem)
		}
	}
	return &diffSet
}

func (s *ComparableSet[T]) SymmetricDifference(other Set[T]) Set[T] {
	sdSet := make(ComparableSet[T], s.Len()+other.Len())
	for elem := range *s {
		if !other.ContainsAll(elem) {
			sdSet.Add(elem)
		}
	}
	other.ForEach(func(t T) bool {
		if !s.ContainsAll(t) {
			sdSet.Add(t)
		}
		return true
	})
	return &sdSet
}

func (s *ComparableSet[T]) IsEmpty() bool {
	return s.Len() == 0
}

func (s *ComparableSet[T]) Equal(other Set[T]) bool {
	if s.Len() != other.Len() {
		return false
	}
	for elem := range *s {
		if !other.Contains(elem) {
			return false
		}
	}
	return true
}

func (s *ComparableSet[T]) IsSubset(other Set[T]) bool {
	if s.Len() > other.Len() {
		return false
	}
	for elem := range *s {
		if !other.ContainsAll(elem) {
			return false
		}
	}
	return true
}

func (s *ComparableSet[T]) IsProperSubset(other Set[T]) bool {
	return s.Len() < other.Len() && s.IsSubset(other)
}

func (s *ComparableSet[T]) IsSuperset(other Set[T]) bool {
	return other.IsSubset(s)
}

func (s *ComparableSet[T]) IsProperSuperset(other Set[T]) bool {
	return s.Len() > other.Len() && s.IsSuperset(other)
}

func (s *ComparableSet[T]) Lock() {}

func (s *ComparableSet[T]) Unlock() {}

func (s *ComparableSet[T]) RLock() {}

func (s *ComparableSet[T]) RUnlock() {}
