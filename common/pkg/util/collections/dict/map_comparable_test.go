package dict

import "testing"

func TestMapBasicOps(t *testing.T) {
	t.Parallel()
	tests := []struct {
		name    string
		prepare func() Map[string, int]
		verify  func(t *testing.T, m Map[string, int])
	}{
		{
			name: "set and get",
			prepare: func() Map[string, int] {
				m := NewComparableMap[string, int](0)
				m.Set("a", 1)
				return m
			},
			verify: func(t *testing.T, m Map[string, int]) {
				v, ok := m.Get("a")
				if !ok || v != 1 {
					t.Errorf("Get() = %v,%v; want 1,true", v, ok)
				}
			},
		},
		{
			name: "remove key",
			prepare: func() Map[string, int] {
				m := NewComparableMap[string, int](0)
				m.Set("a", 1)
				m.Remove("a")
				return m
			},
			verify: func(t *testing.T, m Map[string, int]) {
				if m.Contains("a") {
					t.Errorf("Remove() failed, key still exists")
				}
			},
		},
		{
			name: "pop key",
			prepare: func() Map[string, int] {
				m := NewComparableMap[string, int](0)
				m.Set("x", 42)
				return m
			},
			verify: func(t *testing.T, m Map[string, int]) {
				v, ok := m.Pop("x")
				if !ok || v != 42 {
					t.Errorf("Pop() = %v,%v; want 42,true", v, ok)
				}
				if m.Contains("x") {
					t.Errorf("Pop() did not remove key")
				}
			},
		},
		{
			name: "merge maps",
			prepare: func() Map[string, int] {
				m1 := NewComparableMap[string, int](0)
				m1.Set("a", 1)
				m2 := NewComparableMap[string, int](0)
				m2.Set("b", 2)
				m1.Merge(m2)
				return m1
			},
			verify: func(t *testing.T, m Map[string, int]) {
				if !m.ContainsAll("a", "b") {
					t.Errorf("Merge() failed, got keys %v", m.Keys())
				}
			},
		},
		{
			name: "clone map",
			prepare: func() Map[string, int] {
				m := NewComparableMap[string, int](0)
				m.Set("a", 100)
				return m.Clone()
			},
			verify: func(t *testing.T, m Map[string, int]) {
				if v, ok := m.Get("a"); !ok || v != 100 {
					t.Errorf("Clone() failed, got %v,%v", v, ok)
				}
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := tt.prepare()
			tt.verify(t, m)
		})
	}
}
