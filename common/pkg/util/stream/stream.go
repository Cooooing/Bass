package stream

import (
	"common/pkg/util/base"
	"context"
	"sync"
)

type Stream[T any] interface {
	// 源操作（创建流）

	// 中间操作

	// Map 返回一个流，该流由该流中的每个元素通过指定的映射函数映射生成。转换后的元素类型不变。
	Map(mapper base.UnaryOperator[T]) Stream[T]
	// Peek 返回一个由该流的元素组成的流，并在从结果流中消耗元素时对每个元素执行提供的操作。
	Peek(action base.Consumer[T]) Stream[T]
	// Filter 返回一个流，该流由与给定条件函数匹配的元素组成。
	Filter(predicate base.Predicate[T]) Stream[T]
	// Skip 在丢弃流的前n个元素后，返回由该流的其余元素组成的流。如果此流包含的元素少于n个，则会阻塞。
	Skip(n int) Stream[T]
	// Limit 返回一个由该流的元素组成的流，该流被截断为长度不超过maxSize。如果此流包含的元素少于maxSize个，则会阻塞。
	Limit(maxSize int) Stream[T]
	// Distinct 返回一个去重后的流，元素顺序保持不变。
	Distinct() Stream[T]
	// Sorted 返回一个已排序的流，元素顺序保持不变。
	Sorted(comparator base.Comparator[T]) Stream[T]

	// 终止操作

	// ForEach 迭代流中的每个元素，按顺序执行提供的操作。
	ForEach(action base.Consumer[T]) error
	// ForEachOrdered 迭代流中的每个元素，按顺序执行提供的操作。
	ForEachOrdered(comparator base.Comparator[T], action base.Consumer[T]) error
	// AnyMatch 检查是否存在满足条件的元素
	AnyMatch(predicate base.Predicate[T]) (bool, error)
	// AllMatch 检查是否所有元素都满足条件
	AllMatch(predicate base.Predicate[T]) (bool, error)
	// NoneMatch 检查是否没有元素满足条件
	NoneMatch(predicate base.Predicate[T]) (bool, error)
	// ToArray 返回一个包含此流中所有元素的数组。
	ToArray() ([]T, error)
	// Count 返回此流中的元素数。
	Count() (int, error)
	// Min 返回此流中的最小元素。
	Min(comparator base.Comparator[T]) (T, error)
	// Max 返回此流中的最大元素。
	Max(comparator base.Comparator[T]) (T, error)
	// FindFirst 尝试返回此流中的第一个元素。
	FindFirst() (T, error)
	// FindAny 尝试返回此流中的任意元素。
	FindAny() (T, error)
	// Reduce 尝试将此流中元素归约为单个值。初始值为流中第一个元素。
	Reduce(accumulator base.BinaryOperator[T]) (T, error)
	// ReduceByDefault 尝试将此流中元素归约为单个值。给定初始值。
	ReduceByDefault(identity T, accumulator base.BinaryOperator[T]) (T, error)

	// Iterator 返回一个通道，用于迭代流中的元素。需要自行处理通道的生命周期。
	Iterator() chan T

	// 其他辅助函数

	// getCtx 返回上下文
	getCtx() context.Context
	// close 关闭流
	close(err error)
	// IsParallel 返回是否是并行流
	IsParallel() bool
	// GetParallelGoroutines 返回并行流并行线程数
	GetParallelGoroutines() int
	// Parallel 将流转为并行流，设置并发协程数。必须在创建流后紧接着调用。并行流不保证元素原始顺序。
	Parallel(n int) Stream[T]
}

// 流操作，返回一个流或结果，但流中数据或结果类型发生改变。（解决go中方法不能增加泛型的问题）

func Map[T any, R any](stream Stream[T], mapper base.Function[T, R]) Stream[R] {
	switch s := stream.(type) {
	case *BlockStream[T]:
		arr, _ := s.ToArray()
		newArr := make([]R, 0, len(arr))
		for _, v := range arr {
			newArr = append(newArr, mapper(v))
		}
		return newBlockStream(s.getCtx(), newArr)

	default: // 非阻塞流
		in := stream.Iterator()
		p := newNoBlockStream[R](stream.getCtx())
		out := make(chan R, 1)
		p.out <- out
		go func() {
			defer close(out)
			for v := range in {
				select {
				case <-p.ctx.Done():
					return
				case out <- mapper(v):
				}
			}
		}()
		return p
	}
}

func FlatMap[T any, R any](stream Stream[T], mapper base.Function[T, Stream[R]]) Stream[R] {
	switch s := stream.(type) {
	case *BlockStream[T]:
		arr, _ := s.ToArray()
		newArr := make([]R, 0)
		for _, v := range arr {
			mapped := mapper(v)
			subArr, _ := mapped.ToArray()
			newArr = append(newArr, subArr...)
		}
		return newBlockStream(s.getCtx(), newArr)

	default: // 非阻塞流
		in := stream.Iterator()
		p := newNoBlockStream[R](stream.getCtx())
		out := make(chan R, 1)
		p.out <- out
		go func() {
			defer close(out)
			for v := range in {
				mapped := mapper(v)
				for mv := range mapped.Iterator() {
					select {
					case <-p.ctx.Done():
						return
					case out <- mv:
					}
				}
			}
		}()
		return p
	}
}

func Reduce[T any, R any](stream Stream[T], identity R,
	accumulator base.BiFunction[T, R],
	combiner base.BinaryOperator[R]) (R, error) {

	switch s := stream.(type) {
	case *BlockStream[T]:
		arr, _ := s.ToArray()
		result := identity
		for _, v := range arr {
			result = accumulator(v, result)
		}
		return result, nil

	default: // 非阻塞流
		var err error
		result := identity
		if stream.IsParallel() {
			subResults := make(chan R, stream.GetParallelGoroutines())
			iterator := stream.Iterator()
			var wg sync.WaitGroup
			wg.Add(stream.GetParallelGoroutines())
			go func() {
				wg.Wait()
				stream.close(nil)
				close(subResults)
			}()
			for i := 0; i < stream.GetParallelGoroutines(); i++ {
				go func() {
					defer wg.Done()
					var localResult R
					for v := range iterator {
						localResult = accumulator(v, localResult)
					}
					subResults <- localResult
				}()
			}
			finalResult := identity
			for r := range subResults {
				finalResult = combiner(finalResult, r)
			}
			return finalResult, nil
		} else {
			err = stream.ForEach(func(item T) {
				result = accumulator(item, result)
			})
		}
		return result, err
	}
}

func GroupBy[T any, K comparable](stream Stream[T], classifier base.Function[T, K]) (map[K][]T, error) {
	switch s := stream.(type) {
	case *BlockStream[T]:
		arr, _ := s.ToArray()
		result := make(map[K][]T)
		for _, v := range arr {
			key := classifier(v)
			result[key] = append(result[key], v)
		}
		return result, nil

	default:
		result := make(map[K][]T)
		err := stream.ForEach(func(item T) {
			key := classifier(item)
			result[key] = append(result[key], item)
		})
		return result, err
	}
}
