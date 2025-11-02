package stream

import (
	"common/pkg/util/base"
	"context"
	"errors"
	"fmt"
	"sort"
)

// BlockStream 阻塞流实现，内部直接存储在 slice 中。
type BlockStream[T any] struct {
	ctx      context.Context
	elements []T
	parallel bool
	workers  int
}

func newBlockStream[T any](ctx context.Context, elems []T) *BlockStream[T] {
	return &BlockStream[T]{ctx: ctx, elements: elems, parallel: false, workers: 1}
}

// -------------------- 源操作：阻塞流 --------------------

func OfBlock[T any](ctx context.Context, values ...T) Stream[T] {
	return newBlockStream(ctx, values)
}

func OfChanBlock[T any](ctx context.Context, ins ...chan T) Stream[T] {
	all := make([]T, 0)
	for _, ch := range ins {
		for v := range ch {
			all = append(all, v)
		}
	}
	return newBlockStream(ctx, all)
}

// GenerateBlock 生成阻塞流，需要给定生成数量
func GenerateBlock[T any](ctx context.Context, s base.Supplier[T], count int) Stream[T] {
	all := make([]T, 0, count)
	for i := 0; i < count; i++ {
		all = append(all, s())
	}
	return newBlockStream(ctx, all)
}

func ConcatBlock[T any](ctx context.Context, streams ...Stream[T]) Stream[T] {
	all := make([]T, 0)
	for _, s := range streams {
		arr, _ := s.ToArray()
		all = append(all, arr...)
	}
	return newBlockStream(ctx, all)
}

func EmptyBlock[T any](ctx context.Context) Stream[T] {
	return newBlockStream(ctx, []T{})
}

// --------------------- 中间操作 ---------------------

func (s *BlockStream[T]) Map(mapper base.UnaryOperator[T]) Stream[T] {
	newElems := make([]T, 0, len(s.elements))
	for _, v := range s.elements {
		newElems = append(newElems, mapper(v))
	}
	return newBlockStream(s.ctx, newElems)
}

func (s *BlockStream[T]) Peek(action base.Consumer[T]) Stream[T] {
	for _, v := range s.elements {
		action(v)
	}
	return newBlockStream(s.ctx, s.elements)
}

func (s *BlockStream[T]) Filter(predicate base.Predicate[T]) Stream[T] {
	newElems := make([]T, 0)
	for _, v := range s.elements {
		if predicate(v) {
			newElems = append(newElems, v)
		}
	}
	return newBlockStream(s.ctx, newElems)
}

func (s *BlockStream[T]) Skip(n int) Stream[T] {
	if n >= len(s.elements) {
		return newBlockStream(s.ctx, []T{})
	}
	return newBlockStream(s.ctx, s.elements[n:])
}

func (s *BlockStream[T]) Limit(maxSize int) Stream[T] {
	if maxSize <= 0 {
		return EmptyBlock[T](s.ctx)
	}
	if maxSize >= len(s.elements) {
		return newBlockStream(s.ctx, s.elements)
	}
	return newBlockStream(s.ctx, s.elements[:maxSize])
}

func (s *BlockStream[T]) Distinct() Stream[T] {
	seen := make(map[any]struct{})
	newElems := make([]T, 0, len(s.elements))
	for _, v := range s.elements {
		if _, ok := seen[v]; !ok {
			seen[v] = struct{}{}
			newElems = append(newElems, v)
		}
	}
	return newBlockStream(s.ctx, newElems)
}

func (s *BlockStream[T]) Sorted(comparator base.Comparator[T]) Stream[T] {
	sort.Slice(s.elements, func(i, j int) bool {
		return comparator(s.elements[i], s.elements[j]) < 0
	})

	return s
}

// --------------------- 终止操作 ---------------------

func (s *BlockStream[T]) ForEach(action base.Consumer[T]) error {
	for _, v := range s.elements {
		select {
		case <-s.ctx.Done():
			return s.ctx.Err()
		default:
			action(v)
		}
	}
	return nil
}

func (s *BlockStream[T]) ForEachOrdered(comparator base.Comparator[T], action base.Consumer[T]) error {
	newElems := append([]T(nil), s.elements...)
	sort.Slice(newElems, func(i, j int) bool {
		return comparator(newElems[i], newElems[j]) < 0
	})
	for _, v := range newElems {
		select {
		case <-s.ctx.Done():
			return s.ctx.Err()
		default:
			action(v)
		}
	}
	return nil
}

func (s *BlockStream[T]) AnyMatch(predicate base.Predicate[T]) (bool, error) {
	for _, v := range s.elements {
		if predicate(v) {
			return true, nil
		}
	}
	return false, nil
}

func (s *BlockStream[T]) AllMatch(predicate base.Predicate[T]) (bool, error) {
	for _, v := range s.elements {
		if !predicate(v) {
			return false, nil
		}
	}
	return true, nil
}

func (s *BlockStream[T]) NoneMatch(predicate base.Predicate[T]) (bool, error) {
	for _, v := range s.elements {
		if predicate(v) {
			return false, nil
		}
	}
	return true, nil
}

func (s *BlockStream[T]) ToArray() ([]T, error) {
	return s.elements, nil
}

func (s *BlockStream[T]) Count() (int, error) {
	return len(s.elements), nil
}

func (s *BlockStream[T]) Min(comparator base.Comparator[T]) (T, error) {
	if len(s.elements) == 0 {
		var zero T
		return zero, errors.New("stream is empty")
	}
	min := s.elements[0]
	for _, v := range s.elements[1:] {
		if comparator(v, min) < 0 {
			min = v
		}
	}
	return min, nil
}

func (s *BlockStream[T]) Max(comparator base.Comparator[T]) (T, error) {
	if len(s.elements) == 0 {
		var zero T
		return zero, errors.New("stream is empty")
	}
	max := s.elements[0]
	for _, v := range s.elements[1:] {
		if comparator(v, max) > 0 {
			max = v
		}
	}
	return max, nil
}

func (s *BlockStream[T]) FindFirst() (T, error) {
	if len(s.elements) == 0 {
		var zero T
		return zero, errors.New("stream is empty")
	}
	return s.elements[0], nil
}

func (s *BlockStream[T]) FindAny() (T, error) {
	if len(s.elements) == 0 {
		var zero T
		return zero, errors.New("stream is empty")
	}
	// 阻塞流里 FindAny 就返回第一个
	return s.elements[0], nil
}

func (s *BlockStream[T]) Reduce(accumulator base.BinaryOperator[T]) (T, error) {
	if len(s.elements) == 0 {
		var zero T
		return zero, errors.New("stream is empty")
	}
	result := s.elements[0]
	for _, v := range s.elements[1:] {
		result = accumulator(result, v)
	}
	return result, nil
}

func (s *BlockStream[T]) ReduceByDefault(identity T, accumulator base.BinaryOperator[T]) (T, error) {
	result := identity
	for _, v := range s.elements {
		result = accumulator(result, v)
	}
	return result, nil
}

// --------------------- 辅助函数 ---------------------

func (s *BlockStream[T]) Iterator() chan T {
	ch := make(chan T)
	go func() {
		defer close(ch)
		for _, v := range s.elements {
			ch <- v
		}
	}()
	return ch
}

func (s *BlockStream[T]) getCtx() context.Context { return s.ctx }
func (s *BlockStream[T]) close(err error)         {}
func (s *BlockStream[T]) IsParallel() bool        { return s.parallel }
func (s *BlockStream[T]) GetParallelGoroutines() int {
	return s.workers
}

// Parallel 阻塞流设置并行，返回一个新的非阻塞流
func (s *BlockStream[T]) Parallel(n int) Stream[T] {
	if n < 0 {
		s.close(fmt.Errorf("parallelism must be non-negative number,but now is %d", n))
	}
	// 新建一个非阻塞流返回
	return OfNoBlock[T](s.ctx, s.elements...).Parallel(n)
}
