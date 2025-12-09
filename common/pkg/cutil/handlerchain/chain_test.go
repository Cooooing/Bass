package handlerchain

import (
	"context"
	"fmt"
	"testing"
)

type ValidateHandler struct {
	BaseHandler[string]
}

func (h *ValidateHandler) Handle(ctx context.Context, data string) (string, error) {
	fmt.Println("ValidateHandler executed")
	return h.Next(ctx, data)
}

type TransformHandler struct {
	BaseHandler[string]
}

func (h *TransformHandler) Handle(ctx context.Context, data string) (string, error) {
	fmt.Println("TransformHandler executed")
	return h.Next(ctx, data)
}

func TestChain(t *testing.T) {
	// 创建独立工厂实例
	factory := NewHandlerFactory[string]()

	// 注册 Handler
	factory.Register("validate", func() Handler[string] {
		return &ValidateHandler{NewBaseHandler[string]("validate")}
	})
	factory.Register("transform", func() Handler[string] {
		return &TransformHandler{NewBaseHandler[string]("transform")}
	})

	// 从名称数组构建链条
	names := []string{"validate", "transform", "validate"}
	chain, err := factory.BuildChainByNames(names)
	if err != nil {
		panic(err)
	}

	// 执行链条
	result, err := chain.Handle(context.Background(), "input data")
	if err != nil {
		panic(err)
	}

	fmt.Println("Final Result:", result)
}
