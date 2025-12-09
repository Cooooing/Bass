package base

import (
	"testing"
)

func TestPtr(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name string
		in   any
	}{
		{"int", 42},
		{"string", "hello"},
		{"float", 3.14},
		{"bool", true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			switch v := tt.in.(type) {
			case int:
				ptr := Ptr(v)
				if ptr == nil {
					t.Errorf("Ptr(%v) returned nil", v)
				}
				if *ptr != v {
					t.Errorf("Ptr(%v) = %v, want %v", v, *ptr, v)
				}
			case string:
				ptr := Ptr(v)
				if ptr == nil {
					t.Errorf("Ptr(%v) returned nil", v)
				}
				if *ptr != v {
					t.Errorf("Ptr(%v) = %v, want %v", v, *ptr, v)
				}
			case float64:
				ptr := Ptr(v)
				if ptr == nil {
					t.Errorf("Ptr(%v) returned nil", v)
				}
				if *ptr != v {
					t.Errorf("Ptr(%v) = %v, want %v", v, *ptr, v)
				}
			case bool:
				ptr := Ptr(v)
				if ptr == nil {
					t.Errorf("Ptr(%v) returned nil", v)
				}
				if *ptr != v {
					t.Errorf("Ptr(%v) = %v, want %v", v, *ptr, v)
				}
			default:
				t.Fatalf("unsupported type: %T", v)
			}
		})
	}
}

func TestIf(t *testing.T) {
	t.Parallel()
	tests := []struct {
		condition  bool
		trueValue  any
		falseValue any
		expected   any
	}{
		{true, 1, 2, 1},
		{false, 1, 2, 2},
	}
	for _, tt := range tests {
		actual := If(tt.condition, tt.trueValue, tt.falseValue)
		if actual != tt.expected {
			t.Errorf("If(%v, %v, %v) = %v, want %v", tt.condition, tt.trueValue, tt.falseValue, actual, tt.expected)
		}
	}
}

func TestIsNil(t *testing.T) {
	t.Parallel()

	type TestStruct struct {
		Ptr *int
		Map map[string]int
	}

	var (
		nilIntPtr    *int
		nonNilInt    = 42
		nonNilIntPtr = &nonNilInt // 先生成 *int
		nilSlice     []int
		emptySlice   = []int{}
		nilMap       map[string]int
		emptyMap     = map[string]int{}
		ch           = make(chan int)
		f            = func() {}
		ts           = TestStruct{}
		nilInterface any
		intInterface any = 0
		ptrInterface any = nilIntPtr
		multiPtr     **int
	)

	// multiPtr 需要指向一个 *int
	multiPtr = &nonNilIntPtr

	tests := []struct {
		value any
		isNil bool
	}{
		{nil, true},           // nil
		{nilIntPtr, true},     // nil pointer
		{&nonNilInt, false},   // non-nil pointer
		{nilSlice, true},      // nil slice
		{emptySlice, false},   // empty slice (not nil)
		{nilMap, true},        // nil map
		{emptyMap, false},     // empty map (not nil)
		{ch, false},           // non-nil channel
		{f, false},            // non-nil function
		{&ts, false},          // struct pointer
		{&[0]int{}, false},    // pointer to empty array
		{&[1]int{}, false},    // pointer to non-empty array
		{nilInterface, true},  // nil interface
		{intInterface, false}, // interface holding value
		{ptrInterface, true},  // interface holding nil pointer
		{multiPtr, false},     // multi-level pointer
	}

	for _, tt := range tests {
		isNil := IsNil(tt.value)
		if isNil != tt.isNil {
			t.Errorf("IsNil(%#v) = %v, want %v", tt.value, isNil, tt.isNil)
		}
	}
}
