package stream

import (
	"common/pkg/util/base"
	"context"
	"errors"
	"fmt"
	"sort"
	"sync"
)

type NoBlockStream[T any] struct {
	ctx       context.Context
	out       chan chan T
	cancel    context.CancelFunc // 用于取消所有操作
	err       error
	closeOnce sync.Once      // 确保只关闭流一次
	wg        sync.WaitGroup // 跟踪所有活动协程

	linkedOrConsumed   bool // 标记流是否已被链接（添加新操作）或消耗（执行终端操作）。用于确保流的一次性使用，防止重复操作。
	parallelGoroutines int  // 并行协程数，等于1时为顺序流，大于1时为并行流。
	hasOperations      bool // 标记是否已有操作，用于判断是否是第一次添加操作
}

func newNoBlockStream[T any](ctx context.Context) *NoBlockStream[T] {
	ctx, cancel := context.WithCancel(ctx)
	return &NoBlockStream[T]{
		ctx:                ctx,
		cancel:             cancel,
		out:                make(chan chan T, 1),
		linkedOrConsumed:   false,
		parallelGoroutines: 1,
	}
}

// -------------------- 源操作：非阻塞流 --------------------

// OfNoBlock 从指定元素创建一个流（有限流）
func OfNoBlock[T any](ctx context.Context, values ...T) Stream[T] {
	if len(values) == 0 {
		return EmptyNoBlock[T](ctx) // 创建一个空流
	}
	p := newNoBlockStream[T](ctx)
	out := make(chan T, len(values)) // 使用缓冲通道优化性能
	p.out <- out
	go func() {
		defer close(out)
		for _, v := range values {
			select {
			case <-p.ctx.Done():
				return
			case out <- v:
			}
		}
	}()
	return p
}

// OfChanNoBlock 从指定通道创建一个流（无限流）
func OfChanNoBlock[T any](ctx context.Context, ins ...chan T) Stream[T] {
	p := newNoBlockStream[T](ctx)
	out := make(chan T, 1)
	p.out <- out
	go func() {
		defer close(out)
		var wg sync.WaitGroup
		for _, in := range ins {
			wg.Add(1)
			go func() {
				defer wg.Done()
				for v := range in {
					select {
					case <-p.ctx.Done():
						return
					case out <- v:
					}
				}
			}()
		}
		wg.Wait()
	}()
	return p
}

// GenerateNoBlock 返回一个无限流，由 Supplier 提供的元素组成
func GenerateNoBlock[T any](ctx context.Context, s base.Supplier[T]) Stream[T] {
	p := newNoBlockStream[T](ctx)
	out := make(chan T, 1)
	p.out <- out
	go func() {
		defer close(out)
		for {
			select {
			case <-p.ctx.Done():
				return
			case out <- s():
			}
		}
	}()
	return p
}

// ConcatNoBlock 返回一个流，该流由给定的多个流中的所有元素组成。
func ConcatNoBlock[T any](ctx context.Context, streams ...Stream[T]) Stream[T] {
	if len(streams) == 0 {
		return EmptyNoBlock[T](ctx) // 返回空流
	}
	p := newNoBlockStream[T](ctx)
	out := make(chan T, 1)
	p.out <- out
	go func() {
		defer close(out)
		for _, s := range streams {
			for v := range s.Iterator() {
				select {
				case <-p.ctx.Done():
					return
				case out <- v:
				}
			}
		}
	}()
	return p
}

// EmptyNoBlock 创建一个空流
func EmptyNoBlock[T any](ctx context.Context) Stream[T] {
	stream := newNoBlockStream[T](ctx)
	stream.out <- stream.closeChan()
	return stream
}

// --------------------- 中间操作 ---------------------

func (s *NoBlockStream[T]) Map(action base.UnaryOperator[T]) Stream[T] {
	in, out := s.initOp()
	s.wg.Add(s.parallelGoroutines)
	var currentWg sync.WaitGroup
	currentWg.Add(s.parallelGoroutines)
	go func() {
		currentWg.Wait()
		close(out)
	}()
	for i := 0; i < s.parallelGoroutines; i++ {
		go func() {
			defer s.wg.Done()
			defer currentWg.Done()
			for {
				select {
				case <-s.ctx.Done():
					s.close(s.ctx.Err())
					return
				case v, ok := <-in:
					if !ok {
						return
					}
					out <- action(v)
				}
			}
		}()
	}
	return s
}

func (s *NoBlockStream[T]) Peek(action base.Consumer[T]) Stream[T] {
	in, out := s.initOp()
	s.wg.Add(s.parallelGoroutines)
	var currentWg sync.WaitGroup
	currentWg.Add(s.parallelGoroutines)
	go func() {
		currentWg.Wait()
		close(out)
	}()
	for i := 0; i < s.parallelGoroutines; i++ {
		go func() {
			defer s.wg.Done()
			defer currentWg.Done()
			for v := range in {
				action(v)
				select {
				case <-s.ctx.Done():
					s.close(s.ctx.Err())
					return
				case out <- v:
				}
			}
		}()
	}
	return s
}

func (s *NoBlockStream[T]) Filter(predicate base.Predicate[T]) Stream[T] {
	in, out := s.initOp()
	s.wg.Add(s.parallelGoroutines)
	var currentWg sync.WaitGroup
	currentWg.Add(s.parallelGoroutines)
	go func() {
		currentWg.Wait()
		close(out)
	}()
	for i := 0; i < s.parallelGoroutines; i++ {
		go func() {
			defer s.wg.Done()
			defer currentWg.Done()
			for v := range in {
				if predicate(v) {
					select {
					case <-s.ctx.Done():
						s.close(s.ctx.Err())
						return
					case out <- v:
					}
				}
			}
		}()
	}
	return s
}

func (s *NoBlockStream[T]) Skip(n int) Stream[T] {
	in, out := s.initOp()
	s.wg.Add(1)
	go func() {
		defer s.wg.Done()
		defer close(out)
		count := 0
		for v := range in {
			count++
			if count <= n {
				continue
			}
			select {
			case <-s.ctx.Done():
				s.close(s.ctx.Err())
				return
			case out <- v:
			}
		}
	}()
	return s
}
func (s *NoBlockStream[T]) Limit(maxSize int) Stream[T] {
	if maxSize <= 0 {
		return EmptyBlock[T](s.ctx)
	}
	in, out := s.initOp()
	s.wg.Add(1)
	go func() {
		defer s.wg.Done()
		defer close(out)
		count := 0
		for v := range in {
			count++
			if count > maxSize {
				break
			}
			select {
			case <-s.ctx.Done():
				s.close(s.ctx.Err())
				return
			case out <- v:
			}
		}
	}()
	return s
}

func (s *NoBlockStream[T]) Distinct() Stream[T] {
	in, out := s.initOp()
	s.wg.Add(1)
	go func() {
		defer s.wg.Done()
		defer close(out)
		seen := make(map[any]struct{}) // 使用 map 去重
		for v := range in {
			if _, exists := seen[v]; !exists {
				seen[v] = struct{}{}
				select {
				case <-s.ctx.Done():
					s.close(s.ctx.Err())
					return
				case out <- v:
				}
			}
		}
	}()
	return s
}

// Sorted 排序操作
func (s *NoBlockStream[T]) Sorted(comparator base.Comparator[T]) Stream[T] {
	in, out := s.initOp()
	s.wg.Add(1)
	go func() {
		defer s.wg.Done()
		defer close(out)
		elements := make([]T, 0)
		over := false
		for {
			select {
			case <-s.ctx.Done():
				s.close(s.ctx.Err())
				return
			case v, ok := <-in:
				if !ok {
					over = true
					break
				}
				elements = append(elements, v)
			}
			if over {
				break
			}
		}
		// 排序
		sort.Slice(elements, func(i, j int) bool {
			return comparator(elements[i], elements[j]) < 0
		})
		for _, v := range elements {
			select {
			case <-s.ctx.Done():
				s.close(s.ctx.Err())
				return
			case out <- v:
			}
		}
	}()
	return s
}

// --------------------- 终止操作 ---------------------

func (s *NoBlockStream[T]) ForEach(action base.Consumer[T]) error {
	in := s.initTerminalOp()
	if s.err != nil {
		return s.err
	}
	for {
		select {
		case <-s.ctx.Done():
			s.close(s.ctx.Err())
			return s.ctx.Err()
		case v, ok := <-in:
			if !ok {
				s.close(s.err)
				return s.err
			}
			action(v)
		}
	}
}

func (s *NoBlockStream[T]) ForEachOrdered(comparator base.Comparator[T], action base.Consumer[T]) error {
	s.Sorted(comparator)
	return s.ForEach(action)
}

func (s *NoBlockStream[T]) AnyMatch(predicate base.Predicate[T]) (bool, error) {
	in := s.initTerminalOp()
	if s.err != nil {
		return false, s.err
	}
	for {
		select {
		case <-s.ctx.Done():
			s.close(s.ctx.Err())
			return false, s.ctx.Err()
		case v, ok := <-in:
			if !ok {
				s.close(s.err)
				return false, s.err
			}
			return predicate(v), s.err
		}
	}
}

func (s *NoBlockStream[T]) AllMatch(predicate base.Predicate[T]) (bool, error) {
	in := s.initTerminalOp()
	if s.err != nil {
		return false, s.err
	}
	for {
		select {
		case <-s.ctx.Done():
			s.close(s.ctx.Err())
			return true, s.ctx.Err()
		case v, ok := <-in:
			if !ok {
				s.close(s.err)
				return true, s.err
			}
			if !predicate(v) {
				s.close(s.err)
				return false, s.err
			}
		}
	}
}

func (s *NoBlockStream[T]) NoneMatch(predicate base.Predicate[T]) (bool, error) {
	in := s.initTerminalOp()
	if s.err != nil {
		return false, s.err
	}
	for {
		select {
		case <-s.ctx.Done():
			s.close(s.ctx.Err())
			return true, s.ctx.Err()
		case v, ok := <-in:
			if !ok {
				s.close(s.err)
				return true, s.err
			}
			if predicate(v) {
				s.close(s.err)
				return false, s.err
			}
		}
	}
}

func (s *NoBlockStream[T]) ToArray() ([]T, error) {
	in := s.initTerminalOp()
	if s.err != nil {
		return nil, s.err
	}
	array := make([]T, 0)
	for {
		select {
		case <-s.ctx.Done():
			s.close(s.ctx.Err())
			return array, s.ctx.Err()
		case v, ok := <-in:
			if !ok {
				s.close(s.err)
				return array, s.err
			}
			array = append(array, v)
		}
	}
}

func (s *NoBlockStream[T]) Count() (int, error) {
	in := s.initTerminalOp()
	if s.err != nil {
		return 0, s.err
	}
	count := 0
	for {
		select {
		case <-s.ctx.Done():
			s.close(s.ctx.Err())
			return count, s.ctx.Err()
		case _, ok := <-in:
			if !ok {
				s.close(s.err)
				return count, s.err
			}
			count++
		}
	}
}

func (s *NoBlockStream[T]) Min(comparator base.Comparator[T]) (T, error) {
	in := s.initTerminalOp()
	var zero T
	if s.err != nil {
		return zero, s.err
	}
	var m T
	for {
		select {
		case <-s.ctx.Done():
			s.close(s.ctx.Err())
			return zero, s.ctx.Err()
		case v, ok := <-in:
			if !ok {
				s.close(s.err)
				return m, s.err
			}
			if comparator(v, m) < 0 {
				m = v
			}
		}
	}
}

func (s *NoBlockStream[T]) Max(comparator base.Comparator[T]) (T, error) {
	in := s.initTerminalOp()
	var zero T
	if s.err != nil {
		return zero, s.err
	}
	var m T
	for {
		select {
		case <-s.ctx.Done():
			s.close(s.ctx.Err())
			return zero, s.ctx.Err()
		case v, ok := <-in:
			if !ok {
				s.close(s.err)
				return m, s.err
			}
			if comparator(v, m) > 0 {
				m = v
			}
		}
	}
}

func (s *NoBlockStream[T]) FindFirst() (T, error) {
	in := s.initTerminalOp()
	var m T
	if s.err != nil {
		return m, s.err
	}
	for {
		select {
		case <-s.ctx.Done():
			s.close(s.ctx.Err())
			return m, s.ctx.Err()
		case v, ok := <-in:
			if !ok {
				s.close(s.err)
				return m, s.err
			}
			return v, s.err
		}
	}
}
func (s *NoBlockStream[T]) FindAny() (T, error) {
	in := s.initTerminalOp()
	var m T
	if s.err != nil {
		return m, s.err
	}
	for {
		select {
		case <-s.ctx.Done():
			s.close(s.ctx.Err())
			return m, s.ctx.Err()
		case v, ok := <-in:
			if !ok {
				s.close(s.err)
				return m, s.err
			}
			return v, s.err
		}
	}
}

func (s *NoBlockStream[T]) Reduce(accumulator base.BinaryOperator[T]) (T, error) {
	in := s.initTerminalOp()
	var result T
	for {
		select {
		case <-s.ctx.Done():
			var zero T
			return zero, s.ctx.Err()
		case v, ok := <-in:
			if !ok {
				s.close(s.err)
				return result, s.err
			}
			result = accumulator(result, v)
		}
	}
}

func (s *NoBlockStream[T]) ReduceByDefault(identity T, accumulator base.BinaryOperator[T]) (T, error) {
	in := s.initTerminalOp()
	result := identity
	for {
		select {
		case <-s.ctx.Done():
			var zero T
			return zero, s.ctx.Err()
		case v, ok := <-in:
			if !ok {
				s.close(s.err)
				return result, s.err
			}
			result = accumulator(result, v)
		}
	}
}

func (s *NoBlockStream[T]) Iterator() chan T {
	s.linkedOrConsumed = true
	return <-s.out
}

func (s *NoBlockStream[T]) IsParallel() bool {
	return s.parallelGoroutines != 1
}

func (s *NoBlockStream[T]) GetParallelGoroutines() int {
	return s.parallelGoroutines
}

// Parallel 并行流 n = 0 时，退化为阻塞流
func (s *NoBlockStream[T]) Parallel(n int) Stream[T] {
	if n < 0 {
		s.close(fmt.Errorf("parallelism must be non-negative number,but now is %d", n))
	}
	if s.hasOperations {
		s.close(fmt.Errorf("parallel operation must be the first operation"))
	}
	if n == 0 {
		array, err := s.ToArray()
		if err != nil {
			s.close(err)
		}
		return OfBlock[T](s.ctx, array...)
	}
	s.hasOperations = true
	s.parallelGoroutines = n
	return s
}

// --------------------- 辅助函数 ---------------------

func (s *NoBlockStream[T]) getCtx() context.Context {
	return s.ctx
}

// initOp 初始化中间操作
func (s *NoBlockStream[T]) initOp() (chan T, chan T) {
	if s.linkedOrConsumed {
		s.close(errors.New("stream already operated upon or closed"))
	}
	s.hasOperations = true
	in, ok := <-s.out
	if !ok {
		return s.closeChan(), s.closeChan()
	}
	out := make(chan T, s.parallelGoroutines)
	s.out <- out
	return in, out
}

// initTerminalOp 初始化终端操作
func (s *NoBlockStream[T]) initTerminalOp() chan T {
	if s.linkedOrConsumed {
		s.close(errors.New("stream already operated upon or closed"))
	}
	s.hasOperations = true
	in, ok := <-s.out
	if !ok {
		return s.closeChan()
	}
	return in
}

func (s *NoBlockStream[T]) closeChan() chan T {
	ch := make(chan T)
	close(ch)
	return ch
}

func (s *NoBlockStream[T]) close(err error) {
	go func() {
		s.closeOnce.Do(func() {
			if err != nil {
				s.err = err
				s.linkedOrConsumed = true
				s.cancel()  // 取消所有操作
				s.wg.Wait() // 等待所有协程结束
				close(s.out)
			}
		})
	}()
}
