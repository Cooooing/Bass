package module

import (
	"common/pkg/client/localrpc"
	"common/pkg/server"
	"fmt"
	"log/slog"

	"github.com/go-kratos/kratos/v3/middleware"
	"github.com/go-kratos/kratos/v3/transport"
	"google.golang.org/grpc"
	"google.golang.org/protobuf/types/known/structpb"
)

// Builder 构造业务模块并返回单体可收集的模块能力。
type Builder func(*Runtime, string) (Mounted, func(), error)

// Descriptor 描述业务模块在单体中的装配方式。
type Descriptor struct {
	Name      string
	NewClient func(*localrpc.Conn) any
	Build     Builder
	Mount     func(*localrpc.Conn, []server.Service)
	External  bool
}

// DescriptorOption 调整业务模块装配描述。
type DescriptorOption func(*Descriptor)

// NewDescriptor 创建通用业务模块装配描述。
func NewDescriptor(builder Builder, options ...DescriptorOption) Descriptor {
	descriptor := Descriptor{Build: builder}
	for _, option := range options {
		option(&descriptor)
	}
	return descriptor
}

// WithLocalClient 声明模块向单体内提供的本地客户端。
func WithLocalClient[T any](newClient func(grpc.ClientConnInterface) T) DescriptorOption {
	return func(descriptor *Descriptor) {
		descriptor.NewClient = func(conn *localrpc.Conn) any {
			return newClient(conn)
		}
	}
}

// WithMount 声明模块服务挂载到本地连接的方式。
func WithMount(mount func(*localrpc.Conn, []server.Service)) DescriptorOption {
	return func(descriptor *Descriptor) {
		descriptor.Mount = mount
	}
}

// WithExternal 声明模块服务需要暴露到单体统一 HTTP/gRPC 端口。
func WithExternal() DescriptorOption {
	return func(descriptor *Descriptor) {
		descriptor.External = true
	}
}

// Runtime 为模块化单体提供共享配置、本地客户端和日志。
type Runtime struct {
	Config         *RuntimeConfig
	Logger         *slog.Logger
	Modules        map[string]*structpb.Struct
	Clients        *Clients
	Infrastructure *Infrastructure
}

// Values 返回指定模块的私有配置。
func (r *Runtime) Values(name string) (*structpb.Struct, error) {
	if r == nil {
		return nil, fmt.Errorf("module runtime is required")
	}
	values, ok := r.Modules[name]
	if !ok || values == nil {
		return nil, fmt.Errorf("%s module config is required", name)
	}
	return values, nil
}

// Mounted 表示已完成构造的业务模块。
type Mounted struct {
	Module          any
	Services        []server.Service
	Servers         []transport.Server
	External        bool
	HTTPMiddlewares []middleware.Middleware
}
