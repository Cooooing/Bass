package handlerchain

import (
	"context"
)

// Handler 定义责任链接口，T 是数据类型
type Handler[T any] interface {
	Handle(ctx context.Context, data T) (T, error)
	SetNext(next Handler[T])
	Name() string
}

// BaseHandler 提供基础链实现，可嵌入具体 Handler
type BaseHandler[T any] struct {
	next Handler[T]
	name string
}

// SetNext 设置下一个 Handler
func (b *BaseHandler[T]) SetNext(next Handler[T]) {
	b.next = next
}

// Next 调用下一个 Handler
func (b *BaseHandler[T]) Next(ctx context.Context, data T) (T, error) {
	if b.next != nil {
		return b.next.Handle(ctx, data)
	}
	return data, nil
}

// Name 返回 Handler 名称
func (b *BaseHandler[T]) Name() string {
	return b.name
}

// NewBaseHandler 创建一个带名字的 BaseHandler
func NewBaseHandler[T any](name string) BaseHandler[T] {
	return BaseHandler[T]{name: name}
}
