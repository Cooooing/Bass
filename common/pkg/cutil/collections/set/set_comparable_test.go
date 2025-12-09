package set

import (
	"encoding/json"
	"testing"
)

func TestAdd(t *testing.T) {
	t.Parallel()
	tests := []struct {
		name     string
		set      Set[int]
		add      []int
		expected Set[int]
	}{
		{"empty add", NewComparableSet[int](0), []int{}, NewComparableSet[int](0)},
		{"add one", NewComparableSet[int](0), []int{1}, NewComparableSet[int](0, 1)},
		{"add duplicate", NewComparableSet[int](0, 1, 1), []int{1}, NewComparableSet[int](0, 1)},
		{"not contains", NewComparableSet[int](0, 2, 1, 2), []int{3}, NewComparableSet[int](0, 1, 2, 3)},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.set.AddAll(tt.add...)
			if !tt.set.Equal(tt.expected) {
				t.Errorf("Add() = %v, want %v", tt.set.ToSlice(), tt.expected)
			}
		})
	}
}

func TestRemove(t *testing.T) {
	t.Parallel()
	tests := []struct {
		name     string
		set      Set[int]
		remove   []int
		expected Set[int]
	}{
		{"remove existing", NewComparableSet[int](0, 1, 2, 3), []int{2}, NewComparableSet[int](0, 1, 3)},
		{"remove non-existing", NewComparableSet[int](0, 1, 2), []int{3}, NewComparableSet[int](0, 1, 2)},
		{"remove all", NewComparableSet[int](0, 1, 2), []int{1, 2}, NewComparableSet[int](0)},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.set.RemoveAll(tt.remove...)
			if !tt.set.Equal(tt.expected) {
				t.Errorf("Remove() = %v, want %v", tt.set.ToSlice(), tt.expected)
			}
		})
	}
}

func TestPopAndPopN(t *testing.T) {
	t.Parallel()
	tests := []struct {
		name     string
		set      Set[int]
		n        int
		expected int
	}{
		{"pop from empty", NewComparableSet[int](0), 1, 0},
		{"pop one", NewComparableSet[int](0, 1), 1, 0},
		{"pop multiple", NewComparableSet[int](0, 1, 2, 3), 2, 1},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if _, _ = tt.set.PopN(tt.n); tt.set.Len() != tt.expected {
				t.Errorf("Pop() = %v, want %v", tt.set.Len(), tt.expected)
			}
		})
	}
}

func TestForEach(t *testing.T) {
	t.Parallel()
	tests := []struct {
		name     string
		set      Set[int]
		expected int
	}{
		{"forEach", NewComparableSet[int](0, 1, 2, 3), 6},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			sum := 0
			tt.set.ForEach(func(item int) bool {
				sum += item
				return true
			})
			if sum != tt.expected {
				t.Errorf("ForEach() = %v, want %v", sum, tt.expected)
			}
		})
	}
}

func TestClearAndReset(t *testing.T) {
	t.Parallel()
	tests := []struct {
		name     string
		set      Set[int]
		fn       func(s Set[int])
		expected Set[int]
	}{
		{"clear", NewComparableSet[int](0, 1, 2, 3), func(s Set[int]) { s.Clear() }, NewComparableSet[int](0)},
		{"reset", NewComparableSet[int](0, 1, 2, 3), func(s Set[int]) { s.Reset() }, NewComparableSet[int](0)},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.fn(tt.set)
			if !tt.set.IsEmpty() {
				t.Errorf("%s failed, set not empty", tt.name)
			}
		})
	}
}

func TestClone(t *testing.T) {
	t.Parallel()
	tests := []struct {
		name string
		set  Set[int]
	}{
		{"toSlice", NewComparableSet[int](0, 1, 2, 3)},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			set := tt.set.Clone()
			if !tt.set.Equal(set) {
				t.Errorf("Clone() = %v, want %v", set.ToSlice(), tt.set.ToSlice())
			}
		})
	}
}

func TestMarshalUnmarshalJSON(t *testing.T) {
	t.Parallel()
	tests := []struct {
		name string
		set  Set[int]
	}{
		{"empty set", NewComparableSet[int](0)},
		{"single element", NewComparableSet[int](0, 1)},
		{"multiple elements", NewComparableSet[int](0, 1, 2, 3)},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			bytes, _ := json.Marshal(tt.set)
			s2 := NewComparableSet[int](0)
			_ = json.Unmarshal(bytes, &s2)
			if !tt.set.Equal(s2) {
				t.Errorf("Unmarshal not equal, got %v, want %v", s2.ToSlice(), tt.set.ToSlice())
			}
		})
	}
}

func TestContainsAnyElement(t *testing.T) {
	t.Parallel()
	tests := []struct {
		name     string
		set      Set[int]
		other    Set[int]
		expected bool
	}{
		{"no overlap", NewComparableSet[int](0, 1, 2), NewComparableSet[int](0, 3, 4), false},
		{"some overlap", NewComparableSet[int](0, 1, 2), NewComparableSet[int](0, 2, 3), true},
		{"all overlap", NewComparableSet[int](0, 1, 2), NewComparableSet[int](0, 1, 2), true},
		{"empty other", NewComparableSet[int](0, 1, 2), NewComparableSet[int](0), false},
		{"empty set", NewComparableSet[int](0), NewComparableSet[int](0, 1, 2), false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.set.ContainsAnyElement(tt.other)
			if got != tt.expected {
				t.Errorf("ContainsAnyElement() = %v, want %v", got, tt.expected)
			}
		})
	}
}

func TestUnion(t *testing.T) {
	t.Parallel()
	tests := []struct {
		name     string
		set      Set[int]
		other    Set[int]
		expected Set[int]
	}{
		{"disjoint", NewComparableSet[int](0, 1, 2), NewComparableSet[int](0, 3, 4), NewComparableSet[int](0, 1, 2, 3, 4)},
		{"overlap", NewComparableSet[int](0, 1, 2), NewComparableSet[int](0, 2, 3), NewComparableSet[int](0, 1, 2, 3)},
		{"empty other", NewComparableSet[int](0, 1, 2), NewComparableSet[int](0), NewComparableSet[int](0, 1, 2)},
		{"empty set", NewComparableSet[int](0), NewComparableSet[int](0, 1, 2), NewComparableSet[int](0, 1, 2)},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.set.Union(tt.other)
			if !got.Equal(tt.expected) {
				t.Errorf("Union() = %v, want %v", got.ToSlice(), tt.expected)
			}
		})
	}
}

func TestIntersection(t *testing.T) {
	t.Parallel()
	tests := []struct {
		name     string
		set      Set[int]
		other    Set[int]
		expected Set[int]
	}{
		{"no overlap", NewComparableSet[int](0, 1, 2), NewComparableSet[int](0, 3, 4), NewComparableSet[int](0)},
		{"some overlap", NewComparableSet[int](0), NewComparableSet[int](0, 1, 2), NewComparableSet[int](0)},
		{"all overlap", NewComparableSet[int](0, 1, 2), NewComparableSet[int](0, 1, 2), NewComparableSet[int](0, 1, 2)},
		{"empty set", NewComparableSet[int](0), NewComparableSet[int](0, 1, 2), NewComparableSet[int](0)},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.set.Intersection(tt.other)
			if !got.Equal(tt.expected) {
				t.Errorf("Intersection() = %v, want %v", got.ToSlice(), tt.expected)
			}
		})
	}
}

func TestDifference(t *testing.T) {
	t.Parallel()
	tests := []struct {
		name     string
		set      Set[int]
		other    Set[int]
		expected Set[int]
	}{
		{"disjoint", NewComparableSet[int](0, 1, 2), NewComparableSet[int](0, 3, 4), NewComparableSet[int](0, 1, 2)},
		{"some overlap", NewComparableSet[int](0, 1, 2), NewComparableSet[int](0, 2, 3), NewComparableSet[int](0, 1)},
		{"all overlap", NewComparableSet[int](0, 1, 2), NewComparableSet[int](0, 1, 2), NewComparableSet[int](0)},
		{"empty other", NewComparableSet[int](0, 1, 2), NewComparableSet[int](0), NewComparableSet[int](0, 1, 2)},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.set.Difference(tt.other)
			if !got.Equal(tt.expected) {
				t.Errorf("Difference() = %v, want %v", got.ToSlice(), tt.expected)
			}
		})
	}
}

func TestSymmetricDifference(t *testing.T) {
	t.Parallel()
	tests := []struct {
		name     string
		set      Set[int]
		other    Set[int]
		expected Set[int]
	}{
		{"disjoint", NewComparableSet[int](0, 1, 2), NewComparableSet[int](0, 3, 4), NewComparableSet[int](0, 1, 2, 3, 4)},
		{"some overlap", NewComparableSet[int](0, 1, 2), NewComparableSet[int](0, 2, 3), NewComparableSet[int](0, 1, 3)},
		{"all overlap", NewComparableSet[int](0, 1, 2), NewComparableSet[int](0, 1, 2), NewComparableSet[int](0)},
		{"empty set", NewComparableSet[int](0), NewComparableSet[int](0, 1, 2), NewComparableSet[int](0, 1, 2)},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.set.SymmetricDifference(tt.other)
			if !got.Equal(tt.expected) {
				t.Errorf("SymmetricDifference() = %v, want %v", got.ToSlice(), tt.expected)
			}
		})
	}
}

func TestRelations(t *testing.T) {
	t.Parallel()
	tests := []struct {
		name                                                           string
		set                                                            Set[int]
		other                                                          Set[int]
		isEmpty, equal, subset, properSubset, superset, properSuperset bool
	}{
		{"empty sets", NewComparableSet[int](0), NewComparableSet[int](0), true, true, true, false, true, false},
		{"subset", NewComparableSet[int](0), NewComparableSet[int](0, 1, 2), true, false, true, true, false, false},
		{"superset", NewComparableSet[int](0), NewComparableSet[int](0, 1, 2), true, false, true, true, false, false},
		{"equal sets", NewComparableSet[int](0, 1, 2), NewComparableSet[int](0, 1, 2), false, true, true, false, true, false},
		{"disjoint", NewComparableSet[int](0, 1, 2), NewComparableSet[int](0, 3, 4), false, false, false, false, false, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.set.IsEmpty(); got != tt.isEmpty {
				t.Errorf("IsEmpty() = %v, want %v", got, tt.isEmpty)
			}
			if got := tt.set.Equal(tt.other); got != tt.equal {
				t.Errorf("Equal() = %v, want %v", got, tt.equal)
			}
			if got := tt.set.IsSubset(tt.other); got != tt.subset {
				t.Errorf("IsSubset() = %v, want %v", got, tt.subset)
			}
			if got := tt.set.IsProperSubset(tt.other); got != tt.properSubset {
				t.Errorf("IsProperSubset() = %v, want %v", got, tt.properSubset)
			}
			if got := tt.set.IsSuperset(tt.other); got != tt.superset {
				t.Errorf("IsSuperset() = %v, want %v", got, tt.superset)
			}
			if got := tt.set.IsProperSuperset(tt.other); got != tt.properSuperset {
				t.Errorf("IsProperSuperset() = %v, want %v", got, tt.properSuperset)
			}
		})
	}
}
