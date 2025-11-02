package set

import (
	"common/pkg/util/base"
	"encoding/json"
	"fmt"
	"strings"
)

// KeySet 适用于自定义键值的Set集合，注意其元素需要实现collections.Keyer接口，非线程安全
type KeySet[T Keyer] map[string]T

func NewKeySet[T Keyer](size int, items ...T) Set[T] {
	s := make(KeySet[T], size)
	s.AddAll(items...)
	return &s
}

func (s *KeySet[T]) Add(item T) bool {
	prevLen := s.Len()
	(*s)[item.Key()] = item
	return prevLen != s.Len()
}

func (s *KeySet[T]) AddAll(items ...T) int {
	prevLen := s.Len()
	for _, item := range items {
		(*s)[item.Key()] = item
	}
	return prevLen - s.Len()
}

func (s *KeySet[T]) Remove(item T) {
	delete(*s, item.Key())
}

func (s *KeySet[T]) RemoveAll(items ...T) {
	for _, item := range items {
		delete(*s, item.Key())
	}
}

func (s *KeySet[T]) Pop() (T, bool) {
	for k, v := range *s {
		delete(*s, k)
		return v, true
	}
	var v T
	return v, false
}

func (s *KeySet[T]) PopN(n int) ([]T, int) {
	if n <= 0 || len(*s) == 0 {
		return make([]T, 0), 0
	}
	sn := s.Len()
	if n > sn {
		n = sn
	}

	count := 0
	items := make([]T, 0, sn)
	for k, v := range *s {
		if count >= n {
			break
		}
		delete(*s, k)
		items = append(items, v)
		count++
	}
	return items, count
}

func (s *KeySet[T]) ForEach(action base.Predicate[T]) {
	for _, v := range *s {
		if !action(v) {
			return
		}
	}
}

func (s *KeySet[T]) Contains(item T) bool {
	if _, ok := (*s)[item.Key()]; ok {
		return true
	}
	return false
}

func (s *KeySet[T]) ContainsAll(items ...T) bool {
	for _, item := range items {
		if _, ok := (*s)[item.Key()]; !ok {
			return false
		}
	}
	return true
}

func (s *KeySet[T]) ContainsAny(items ...T) bool {
	for _, item := range items {
		if _, ok := (*s)[item.Key()]; ok {
			return true
		}
	}
	return false
}

func (s *KeySet[T]) Clear() {
	for key := range *s {
		delete(*s, key)
	}
}

func (s *KeySet[T]) Reset() {
	*s = make(KeySet[T])
}

func (s *KeySet[T]) Len() int {
	return len(*s)
}

func (s *KeySet[T]) ToSlice() []T {
	keys := make([]T, 0, s.Len())
	for _, v := range *s {
		keys = append(keys, v)
	}
	return keys
}

func (s *KeySet[T]) Clone() Set[T] {
	clonedSet := NewKeySet[T](s.Len())
	for _, v := range *s {
		clonedSet.Add(v)
	}
	return clonedSet
}

func (s *KeySet[T]) MarshalJSON() ([]byte, error) {
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

func (s *KeySet[T]) UnmarshalJSON(b []byte) error {
	var i []T
	err := json.Unmarshal(b, &i)
	if err != nil {
		return err
	}
	s.AddAll(i...)
	return nil
}

func (s *KeySet[T]) ContainsAnyElement(other Set[T]) bool {
	if s.Len() < other.Len() {
		for _, v := range *s {
			if other.ContainsAll(v) {
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

func (s *KeySet[T]) Union(other Set[T]) Set[T] {
	n := s.Len()
	if other.Len() > n {
		n = other.Len()
	}
	unionedSet := make(KeySet[T], n)

	for _, v := range *s {
		unionedSet.Add(v)
	}
	other.ForEach(func(t T) bool {
		unionedSet.Add(t)
		return true
	})
	return &unionedSet
}

func (s *KeySet[T]) Intersection(other Set[T]) Set[T] {
	n := s.Len()
	if other.Len() < n {
		n = other.Len()
	}
	intersectedSet := make(KeySet[T], n)
	if s.Len() < other.Len() {
		for _, v := range *s {
			if other.ContainsAll(v) {
				intersectedSet.Add(v)
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

func (s *KeySet[T]) Difference(other Set[T]) Set[T] {
	diffSet := make(KeySet[T], s.Len())
	for _, v := range *s {
		if !other.ContainsAll(v) {
			diffSet.Add(v)
		}
	}
	return &diffSet
}

func (s *KeySet[T]) SymmetricDifference(other Set[T]) Set[T] {
	sdSet := make(KeySet[T], s.Len()+other.Len())
	for _, v := range *s {
		if !other.ContainsAll(v) {
			sdSet.Add(v)
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

func (s *KeySet[T]) IsEmpty() bool {
	return s.Len() == 0
}

func (s *KeySet[T]) Equal(other Set[T]) bool {
	if s.Len() != other.Len() {
		return false
	}
	for _, v := range *s {
		if !other.Contains(v) {
			return false
		}
	}
	return true
}

func (s *KeySet[T]) IsSubset(other Set[T]) bool {
	if s.Len() > other.Len() {
		return false
	}
	for _, v := range *s {
		if !other.ContainsAll(v) {
			return false
		}
	}
	return true
}

func (s *KeySet[T]) IsProperSubset(other Set[T]) bool {
	return s.Len() < other.Len() && s.IsSubset(other)
}

func (s *KeySet[T]) IsSuperset(other Set[T]) bool {
	return other.IsSubset(s)
}

func (s *KeySet[T]) IsProperSuperset(other Set[T]) bool {
	return s.Len() > other.Len() && s.IsSuperset(other)
}

func (s *KeySet[T]) Lock() {}

func (s *KeySet[T]) Unlock() {}

func (s *KeySet[T]) RLock() {}

func (s *KeySet[T]) RUnlock() {}
