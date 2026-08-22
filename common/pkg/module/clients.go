package module

import (
	"fmt"
	"reflect"
)

// Clients 保存单体内所有模块提供的本地客户端。
type Clients struct {
	byType map[reflect.Type]any
}

// NewClients 创建空的本地客户端集合。
func NewClients() *Clients {
	return &Clients{
		byType: make(map[reflect.Type]any),
	}
}

// Add 注册模块提供的本地客户端。
func (c *Clients) Add(client any) {
	if c == nil || client == nil {
		return
	}
	c.byType[reflect.TypeOf(client)] = client
}

// ResolveClient 按类型查找本地客户端。
func (c *Clients) ResolveClient(target reflect.Type) (any, bool) {
	if c == nil || target == nil {
		return nil, false
	}
	if client, ok := c.byType[target]; ok {
		return client, true
	}
	for _, client := range c.byType {
		clientType := reflect.TypeOf(client)
		if clientType.AssignableTo(target) || target.Kind() == reflect.Interface && clientType.Implements(target) {
			return client, true
		}
	}
	return nil, false
}

// Client 按类型返回本地客户端。
func Client[T any](clients *Clients) (T, error) {
	var zero T
	if clients == nil {
		return zero, fmt.Errorf("module clients is required")
	}
	target := reflect.TypeFor[T]()
	if client, ok := clients.ResolveClient(target); ok {
		return client.(T), nil
	}
	return zero, fmt.Errorf("%s module client is required", target.String())
}
