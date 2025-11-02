package stream

import (
	"common/pkg/util/base"
	"context"
	"reflect"
	"strconv"
	"testing"
	"time"
)

func TestOfBlock(t *testing.T) {
	t.Parallel()
	tests := []struct {
		name     string
		input    []int
		expected []int
	}{
		{name: "empty stream", input: []int{}, expected: []int{}},
		{name: "single element", input: []int{1}, expected: []int{1}},
		{name: "multiple elements", input: []int{1, 2, 3}, expected: []int{1, 2, 3}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctx := context.Background()
			stream := OfBlock(ctx, tt.input...)
			result, err := stream.ToArray()
			if err != nil {
				t.Errorf("OfBlock() error = %v, want nil", err)
			}
			if !reflect.DeepEqual(result, tt.expected) {
				t.Errorf("OfBlock() = %v, want %v", result, tt.expected)
			}
		})
	}
}

func TestOfChanBlock(t *testing.T) {
	t.Parallel()
	tests := []struct {
		name      string
		inputFunc func() chan int
		expected  []int
	}{
		{
			name: "empty stream",
			inputFunc: func() chan int {
				ch := make(chan int)
				close(ch)
				return ch
			},
			expected: []int{},
		},
		{
			name: "single element",
			inputFunc: func() chan int {
				ch := make(chan int)
				go func() {
					defer close(ch)
					ch <- 1
				}()
				return ch
			},
			expected: []int{1},
		},
		{
			name: "multiple elements",
			inputFunc: func() chan int {
				ch := make(chan int)
				go func() {
					defer close(ch)
					ch <- 1
					ch <- 2
					ch <- 3
				}()
				return ch
			},
			expected: []int{1, 2, 3},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctx := context.Background()
			stream := OfChanBlock(ctx, tt.inputFunc())
			result, err := stream.ToArray()
			if err != nil {
				t.Errorf("OfBlock() error = %v, want nil", err)
			}
			if !reflect.DeepEqual(result, tt.expected) {
				t.Errorf("OfBlock() = %v, want %v", result, tt.expected)
			}
		})
	}
}

func TestGenerateBlock(t *testing.T) {
	t.Parallel()
	tests := []struct {
		name     string
		count    int
		input    base.Supplier[int]
		expected []int
	}{
		{name: "generate numbers", count: 3, input: func() int { return 1 }, expected: []int{1, 1, 1}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctx := context.Background()
			stream := GenerateBlock(ctx, tt.input, tt.count).Limit(tt.count)
			result, err := stream.ToArray()
			if err != nil {
				t.Errorf("GenerateBlock() error = %v, want nil", err)
			}
			if !reflect.DeepEqual(result, tt.expected) {
				t.Errorf("GenerateBlock() = %v, want %v", result, tt.expected)
			}
		})
	}
}

func TestConcatBlock(t *testing.T) {
	t.Parallel()
	tests := []struct {
		name     string
		input    []Stream[int]
		expected []int
	}{
		{name: "concat empty stream", input: []Stream[int]{EmptyBlock[int](context.Background())}, expected: []int{}},
		{name: "concat one stream", input: []Stream[int]{OfBlock(context.Background(), 1, 2)}, expected: []int{1, 2}},
		{name: "concat two streams", input: []Stream[int]{OfBlock(context.Background(), 1, 2), OfBlock(context.Background(), 3, 4)}, expected: []int{1, 2, 3, 4}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctx := context.Background()
			stream := ConcatBlock(ctx, tt.input...)
			result, err := stream.ToArray()
			if err != nil {
				t.Errorf("ConcatBlock() error = %v, want nil", err)
			}
			if !reflect.DeepEqual(result, tt.expected) {
				t.Errorf("ConcatBlock() = %v, want %v", result, tt.expected)
			}
		})
	}
}

func TestEmptyBlock(t *testing.T) {
	t.Parallel()
	tests := []struct {
		name     string
		expected []int
	}{
		{name: "empty stream", expected: []int{}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctx := context.Background()
			stream := EmptyBlock[int](ctx)
			result, err := stream.ToArray()
			if err != nil {
				t.Errorf("EmptyBlock() error = %v, want nil", err)
			}
			if !reflect.DeepEqual(result, tt.expected) {
				t.Errorf("EmptyBlock() = %v, want %v", result, tt.expected)
			}
		})
	}
}

func TestMapToAnotherStreamBlock(t *testing.T) {
	t.Parallel()
	tests := []struct {
		name     string
		input    []int
		mapper   func(int) string
		expected []string
	}{
		{name: "empty stream", input: []int{}, mapper: func(x int) string { return strconv.Itoa(x * 2) }, expected: []string{}},
		{name: "double elements", input: []int{1, 2, 3}, mapper: func(x int) string { return strconv.Itoa(x * 2) }, expected: []string{"2", "4", "6"}},
		{name: "add one", input: []int{1, 2, 3}, mapper: func(x int) string { return strconv.Itoa(x + 1) }, expected: []string{"2", "3", "4"}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctx := context.Background()
			stream := OfBlock(ctx, tt.input...)
			result, err := Map[int, string](stream, tt.mapper).ToArray()
			if err != nil {
				t.Errorf("MapToAnotherStream() error = %v, want nil", err)
			}
			if !reflect.DeepEqual(result, tt.expected) {
				t.Errorf("MapToAnotherStream() = %v, want %v", result, tt.expected)
			}
		})
	}
}

func TestFlatMapToAnotherStreamBlock(t *testing.T) {
	t.Parallel()
	tests := []struct {
		name     string
		input    []string
		mapper   func(string) Stream[int]
		expected []int
	}{
		{name: "empty stream", input: []string{},
			mapper: func(x string) Stream[int] {
				var result []int
				for _, c := range x {
					result = append(result, int(c))
				}
				return OfBlock[int](context.Background(), result...)
			},
			expected: []int{},
		},
		{name: "word to ascii value of char", input: []string{"hello", "golang"},
			mapper: func(x string) Stream[int] {
				var result []int
				for _, c := range x {
					result = append(result, int(c))
				}
				return OfBlock[int](context.Background(), result...)
			},
			expected: []int{104, 101, 108, 108, 111, 103, 111, 108, 97, 110, 103},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctx := context.Background()
			stream := OfBlock(ctx, tt.input...)
			result, err := FlatMap[string, int](stream, tt.mapper).ToArray()
			if err != nil {
				t.Errorf("MapToAnotherStream() error = %v, want nil", err)
			}
			if !reflect.DeepEqual(result, tt.expected) {
				t.Errorf("MapToAnotherStream() = %v, want %v", result, tt.expected)
			}
		})
	}
}

func TestReduceToAnotherTypeBlock(t *testing.T) {
	t.Parallel()
	tests := []struct {
		name     string
		input    []string
		identity int
		mapper   func(string, int) int
		combiner func(int, int) int
		expected int
	}{
		{name: "empty stream", input: []string{}, identity: 0,
			mapper:   func(s string, i int) int { return i + len(s) },
			combiner: func(a int, b int) int { return a + b },
			expected: 0,
		},
		{name: "calculate the sum of word lengths", input: []string{"hello", "golang"}, identity: 0,
			mapper:   func(s string, i int) int { return i + len(s) },
			combiner: func(a int, b int) int { return a + b },
			expected: 11,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctx := context.Background()
			stream := OfBlock(ctx, tt.input...)
			result, err := Reduce[string, int](stream, tt.identity, tt.mapper, tt.combiner)
			if err != nil {
				t.Errorf("ReduceToAnotherType() error = %v, want nil", err)
			}
			if !reflect.DeepEqual(result, tt.expected) {
				t.Errorf("ReduceToAnotherType() = %v, want %v", result, tt.expected)
			}
		})
	}
}

func TestGroupByBlock(t *testing.T) {
	t.Parallel()
	tests := []struct {
		name       string
		input      []string
		classifier func(string) string
		expected   map[string][]string
	}{
		{name: "empty stream", input: []string{}, classifier: func(s string) string { return s[0:1] }, expected: map[string][]string{}},
		{name: "group by first char", input: []string{"hello", "golang", "hi", "gopher"}, classifier: func(s string) string { return s[0:1] }, expected: map[string][]string{"h": {"hello", "hi"}, "g": {"golang", "gopher"}}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctx := context.Background()
			stream := OfBlock(ctx, tt.input...)
			result, err := GroupBy[string, string](stream, tt.classifier)
			if err != nil {
				t.Errorf("GroupBy() error = %v, want nil", err)
			}
			if !reflect.DeepEqual(result, tt.expected) {
				t.Errorf("GroupBy() = %v, want %v", result, tt.expected)
			}
		})
	}

}

// --------

func TestMapBlock(t *testing.T) {
	t.Parallel()
	tests := []struct {
		name     string
		input    []int
		mapper   base.UnaryOperator[int]
		expected []int
	}{
		{name: "empty stream", input: []int{}, mapper: func(x int) int { return x * 2 }, expected: []int{}},
		{name: "double elements", input: []int{1, 2, 3}, mapper: func(x int) int { return x * 2 }, expected: []int{2, 4, 6}},
		{name: "add one", input: []int{1, 2, 3}, mapper: func(x int) int { return x + 1 }, expected: []int{2, 3, 4}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctx := context.Background()
			stream := OfBlock(ctx, tt.input...).Map(tt.mapper)
			result, err := stream.ToArray()
			if err != nil {
				t.Errorf("Map() error = %v, want nil", err)
			}
			if !reflect.DeepEqual(result, tt.expected) {
				t.Errorf("Map() = %v, want %v", result, tt.expected)
			}
		})
	}
}

func TestFilterBlock(t *testing.T) {
	t.Parallel()
	tests := []struct {
		name      string
		input     []int
		predicate base.Predicate[int]
		expected  []int
	}{
		{name: "empty stream", input: []int{}, predicate: func(x int) bool { return x%2 == 0 }, expected: []int{}},
		{name: "even numbers", input: []int{1, 2, 3, 4}, predicate: func(x int) bool { return x%2 == 0 }, expected: []int{2, 4}},
		{name: "greater than 2", input: []int{1, 2, 3, 4}, predicate: func(x int) bool { return x > 2 }, expected: []int{3, 4}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctx := context.Background()
			stream := OfBlock(ctx, tt.input...).Filter(tt.predicate)
			result, err := stream.ToArray()
			if err != nil {
				t.Errorf("Filter() error = %v, want nil", err)
			}
			if !reflect.DeepEqual(result, tt.expected) {
				t.Errorf("Filter() = %v, want %v", result, tt.expected)
			}
		})
	}
}

// TestSkip 测试 Skip 方法
func TestSkipBlock(t *testing.T) {
	t.Parallel()
	tests := []struct {
		name     string
		input    []int
		n        int
		expected []int
	}{
		{name: "empty stream", input: []int{}, expected: []int{}},
		{name: "skip 2", input: []int{1, 2, 3, 4}, n: 2, expected: []int{3, 4}},
		{name: "skip all", input: []int{1, 2}, n: 3, expected: []int{}},
		{name: "skip zero", input: []int{1, 2, 3}, n: 0, expected: []int{1, 2, 3}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctx := context.Background()
			stream := OfBlock(ctx, tt.input...).Skip(tt.n)
			result, err := stream.ToArray()
			if err != nil {
				t.Errorf("Skip() error = %v, want nil", err)
			}
			if !reflect.DeepEqual(result, tt.expected) {
				t.Errorf("Skip() = %v, want %v", result, tt.expected)
			}
		})
	}
}

// TestLimit 测试 Limit 方法
func TestLimitBlock(t *testing.T) {
	t.Parallel()
	tests := []struct {
		name     string
		input    []int
		maxSize  int
		expected []int
	}{
		{name: "empty stream", input: []int{}, expected: []int{}},
		{name: "limit 2", input: []int{1, 2, 3, 4}, maxSize: 2, expected: []int{1, 2}},
		{name: "limit exceed", input: []int{1, 2}, maxSize: 3, expected: []int{1, 2}},
		{name: "limit zero", input: []int{1, 2, 3}, maxSize: 0, expected: []int{}},
		{name: "limit negative", input: []int{1, 2, 3}, maxSize: -1, expected: []int{}},
		{name: "limit all", input: []int{1, 2, 3, 4}, maxSize: 4, expected: []int{1, 2, 3, 4}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctx := context.Background()
			stream := OfBlock(ctx, tt.input...).Limit(tt.maxSize)
			result, err := stream.ToArray()
			if err != nil {
				t.Errorf("Limit() error = %v, want nil", err)
			}
			if !reflect.DeepEqual(result, tt.expected) {
				t.Errorf("Limit() = %v, want %v", result, tt.expected)
			}
		})
	}
}

// TestDistinct 测试 Distinct 方法
func TestDistinctBlock(t *testing.T) {
	t.Parallel()
	tests := []struct {
		name     string
		input    []int
		expected []int
	}{
		{name: "empty stream", input: []int{}, expected: []int{}},
		{name: "remove duplicates", input: []int{1, 2, 2, 3, 1}, expected: []int{1, 2, 3}},
		{name: "no duplicates", input: []int{1, 2, 3}, expected: []int{1, 2, 3}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctx := context.Background()
			stream := OfBlock(ctx, tt.input...).Distinct()
			result, err := stream.ToArray()
			if err != nil {
				t.Errorf("Distinct() error = %v, want nil", err)
			}
			if !reflect.DeepEqual(result, tt.expected) {
				t.Errorf("Distinct() = %v, want %v", result, tt.expected)
			}
		})
	}
}

// TestSorted 测试 Sorted 方法
func TestSortedBlock(t *testing.T) {
	t.Parallel()
	tests := []struct {
		name       string
		input      []int
		comparator base.Comparator[int]
		expected   []int
	}{
		{name: "empty stream", input: []int{}, comparator: func(a, b int) int { return a - b }, expected: []int{}},
		{name: "ascending order", input: []int{3, 1, 2}, comparator: func(a, b int) int { return a - b }, expected: []int{1, 2, 3}},
		{name: "descending order", input: []int{3, 1, 2}, comparator: func(a, b int) int { return b - a }, expected: []int{3, 2, 1}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctx := context.Background()
			stream := OfBlock(ctx, tt.input...).Sorted(tt.comparator)
			result, err := stream.ToArray()
			if err != nil {
				t.Errorf("Sorted() error = %v, want nil", err)
			}
			if !reflect.DeepEqual(result, tt.expected) {
				t.Errorf("Sorted() = %v, want %v", result, tt.expected)
			}
		})
	}
}

// TestForEach 测试 ForEach 方法
func TestForEachBlock(t *testing.T) {
	t.Parallel()
	tests := []struct {
		name     string
		input    []int
		consumer base.Consumer[int]
		expected []int
	}{
		{name: "empty stream", input: []int{}, consumer: func(x int) {}, expected: []int{}},
		{name: "collect elements", input: []int{1, 2, 3}, consumer: func(x int) {}, expected: []int{1, 2, 3}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctx := context.Background()
			collected := make([]int, 0)
			consumer := func(x int) { collected = append(collected, x) }
			stream := OfBlock(ctx, tt.input...)
			err := stream.ForEach(consumer)
			if err != nil {
				t.Errorf("ForEach() error = %v, want nil", err)
			}
			if !reflect.DeepEqual(collected, tt.expected) {
				t.Errorf("ForEach() collected = %v, want %v", collected, tt.expected)
			}
		})
	}
}

// TestCount 测试 Count 方法
func TestCountBlock(t *testing.T) {
	t.Parallel()
	tests := []struct {
		name     string
		input    []int
		expected int
	}{
		{name: "empty stream", input: []int{}, expected: 0},
		{name: "three elements", input: []int{1, 2, 3}, expected: 3},
		{name: "single element", input: []int{1}, expected: 1},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctx := context.Background()
			stream := OfBlock(ctx, tt.input...)
			count, err := stream.Count()
			if err != nil {
				t.Errorf("Count() error = %v, want nil", err)
			}
			if count != tt.expected {
				t.Errorf("Count() = %d, want %d", count, tt.expected)
			}
		})
	}
}

// TestParallel 测试并行流
func TestParallelBlock(t *testing.T) {
	t.Parallel()
	ctx := context.Background()
	input := []int{1, 2, 3, 4, 5}
	stream := OfBlock(ctx, input...).
		Parallel(2).
		Map(func(x int) int {
			time.Sleep(100 * time.Millisecond) // 模拟耗时操作
			return x * 2
		}).
		Sorted(func(a, b int) int {
			return a - b
		})
	result, err := stream.ToArray()
	if err != nil {
		t.Errorf("Parallel() error = %v, want nil", err)
	}
	expected := []int{2, 4, 6, 8, 10}
	if !reflect.DeepEqual(result, expected) {
		t.Errorf("Parallel() = %v, want %v", result, expected)
	}
}

// TestContextCancel 测试上下文取消
// func TestContextCancelBlock(t *testing.T) {
// 	t.Parallel()
// 	ctx, cancel := context.WithCancel(context.Background())
// 	defer cancel()
//
// 	stream := OfBlock(ctx, 1, 2, 3, 4, 5).Map(func(x int) int {
// 		time.Sleep(100 * time.Millisecond) // 模拟耗时操作
// 		return x
// 	})
//
// 	cancel() // 立即取消上下文
// 	_, err := stream.ToArray()
// 	if err == nil {
// 		t.Error("ToArray() expected context canceled error, got nil")
// 	}
// 	if !errors.Is(err, context.Canceled) {
// 		t.Errorf("ToArray() error = %v, want context.Canceled", err)
// 	}
// }
