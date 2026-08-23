package module

import (
	"common/pkg/client/localrpc"
	commonmodule "common/pkg/module"
	commonserver "common/pkg/server"
	"fmt"
	"log/slog"
	"monolith/internal/catalog"
	"monolith/internal/config"

	"github.com/go-kratos/kratos/v3/middleware"
	"github.com/go-kratos/kratos/v3/transport"
	"github.com/google/wire"
	"google.golang.org/protobuf/types/known/structpb"
)

// ProviderSet 提供模块化单体的通用装配依赖。
var ProviderSet = wire.NewSet(
	ProvideRegistry,
	ProvideServices,
	ProvideHTTPMiddlewares,
)

// Registry 保存单体中的模块实例、本地连接、本地客户端和对外服务。
type Registry struct {
	modules         map[string]any
	conns           map[string]*localrpc.Conn
	clients         *commonmodule.Clients
	services        []commonserver.Service
	servers         []transport.Server
	httpMiddlewares []middleware.Middleware
}

// NewRegistry 创建空模块注册表。
func NewRegistry() *Registry {
	return &Registry{
		modules: make(map[string]any),
		conns:   make(map[string]*localrpc.Conn),
		clients: commonmodule.NewClients(),
	}
}

// ProvideRegistry 按模块 catalog 构造单体注册表。
func ProvideRegistry(c *config.Bootstrap, logger *slog.Logger) (*Registry, func(), error) {
	return Assemble(catalog.Descriptors, c.RuntimeConfig(), c.GetModules(), logger)
}

// ProvideServices 提供统一 HTTP/gRPC 端口需要注册的服务。
func ProvideServices(registry *Registry) []commonserver.Service {
	return registry.Services()
}

// ProvideHTTPMiddlewares 提供模块对外 HTTP 中间件。
func ProvideHTTPMiddlewares(registry *Registry) []middleware.Middleware {
	return registry.HTTPMiddlewares()
}

// Assemble 按模块描述装配模块化单体。
func Assemble(
	descriptors []commonmodule.Descriptor,
	config *commonmodule.RuntimeConfig,
	modules map[string]*structpb.Struct,
	logger *slog.Logger,
) (*Registry, func(), error) {
	registry := NewRegistry()
	if err := registry.Prepare(descriptors); err != nil {
		return nil, func() {}, err
	}

	var cleanups []func()
	cleanup := func() {
		for i := len(cleanups) - 1; i >= 0; i-- {
			cleanups[i]()
		}
	}
	infrastructure, infrastructureCleanup, err := commonmodule.NewInfrastructure(logger, config)
	if err != nil {
		return nil, func() {}, err
	}
	cleanups = append(cleanups, infrastructureCleanup)

	runtime := &commonmodule.Runtime{
		Config:         config,
		Logger:         logger,
		Modules:        modules,
		Clients:        registry.clients,
		Infrastructure: infrastructure,
	}
	type assembledModule struct {
		descriptor commonmodule.Descriptor
		mounted    commonmodule.Mounted
	}
	var assembled []assembledModule

	for _, descriptor := range descriptors {
		if descriptor.Build == nil {
			cleanup()
			return nil, func() {}, fmt.Errorf("%s module builder is required", descriptor.Name)
		}
		mounted, moduleCleanup, err := descriptor.Build(runtime, descriptor.Name)
		if err != nil {
			cleanup()
			return nil, func() {}, fmt.Errorf("build %s module: %w", descriptor.Name, err)
		}
		if descriptor.External {
			mounted.External = true
		}
		if moduleCleanup != nil {
			cleanups = append(cleanups, moduleCleanup)
		}
		assembled = append(assembled, assembledModule{descriptor: descriptor, mounted: mounted})
	}

	for _, item := range assembled {
		registry.Mount(item.descriptor, item.mounted)
	}

	return registry, cleanup, nil
}

// Prepare 先创建所有本地连接和客户端，允许模块之间存在启动期引用。
func (r *Registry) Prepare(descriptors []commonmodule.Descriptor) error {
	for _, descriptor := range descriptors {
		if descriptor.Name == "" {
			return fmt.Errorf("module name is required")
		}
		if _, ok := r.conns[descriptor.Name]; ok {
			return fmt.Errorf("%s module is duplicated", descriptor.Name)
		}
		r.conns[descriptor.Name] = localrpc.NewConn()
	}

	for _, descriptor := range descriptors {
		if descriptor.NewClient == nil {
			continue
		}
		client := descriptor.NewClient(r.conns[descriptor.Name])
		if client == nil {
			continue
		}
		r.clients.Add(client)
	}
	return nil
}

// Mount 挂载单个模块并收集对外能力。
func (r *Registry) Mount(descriptor commonmodule.Descriptor, mounted commonmodule.Mounted) {
	r.modules[descriptor.Name] = mounted.Module
	if descriptor.Mount != nil {
		descriptor.Mount(r.conns[descriptor.Name], mounted.Services)
	}
	r.servers = append(r.servers, mounted.Servers...)
	if mounted.External {
		r.services = append(r.services, mounted.Services...)
		r.httpMiddlewares = append(r.httpMiddlewares, mounted.HTTPMiddlewares...)
	}
}

// Services 返回单体统一端口需要注册的对外服务。
func (r *Registry) Services() []commonserver.Service {
	return r.services
}

// Servers 返回模块内部需要跟随单体启停的后台服务。
func (r *Registry) Servers() []transport.Server {
	return r.servers
}

// HTTPMiddlewares 返回对外 HTTP 服务需要追加的模块中间件。
func (r *Registry) HTTPMiddlewares() []middleware.Middleware {
	return r.httpMiddlewares
}

// Module 返回指定模块实例。
func (r *Registry) Module(name string) (any, bool) {
	module, ok := r.modules[name]
	return module, ok
}

// Conn 返回指定模块的本地连接。
func (r *Registry) Conn(name string) (*localrpc.Conn, bool) {
	conn, ok := r.conns[name]
	return conn, ok
}
