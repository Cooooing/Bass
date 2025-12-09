package handlerchain

import (
	"common/pkg/cutil/base"
	"fmt"
)

type HandlerFactory[T any] struct {
	registry map[string]base.Supplier[Handler[T]]
}

// NewHandlerFactory 创建新的工厂实例
func NewHandlerFactory[T any]() *HandlerFactory[T] {
	return &HandlerFactory[T]{registry: make(map[string]base.Supplier[Handler[T]])}
}

func NewHandlerFactoryWithHandlers[T any](handlers ...Handler[T]) *HandlerFactory[T] {
	factory := &HandlerFactory[T]{registry: make(map[string]base.Supplier[Handler[T]])}
	for _, handler := range handlers {
		factory.Register(handler.Name(), func() Handler[T] {
			return handler
		})
	}
	return factory
}

// Register 注册 Handler 构造函数
func (f *HandlerFactory[T]) Register(name string, constructor base.Supplier[Handler[T]]) {
	f.registry[name] = constructor
}

// BuildChainByNames 根据 Handler 名称数组构建责任链
func (f *HandlerFactory[T]) BuildChainByNames(names []string) (Handler[T], error) {
	if len(names) == 0 {
		return nil, fmt.Errorf("handler names list is empty")
	}

	var handlers []Handler[T]
	for _, name := range names {
		constructor, ok := f.registry[name]
		if !ok {
			return nil, fmt.Errorf("handler %s not registered", name)
		}
		handlers = append(handlers, constructor())
	}

	// 构建链条
	for i := 0; i < len(handlers)-1; i++ {
		handlers[i].SetNext(handlers[i+1])
	}
	return handlers[0], nil
}
