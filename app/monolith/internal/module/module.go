package module

import (
	bbsmodule "bbs/module"
	"common/pkg/client/localrpc"
	"common/pkg/client/rpc"
	"common/pkg/constant"
	commonserver "common/pkg/server"
	contentmodule "content/module"
	economymodule "economy/module"
	"fmt"
	gametownmodule "game_town/module"
	immodule "im/module"
	"log/slog"
	"monolith/internal/config"
	notifymodule "notify/module"
	platformmodule "platform/module"
	pushhubmodule "push_hub/module"
	pushnodemodule "push_node/module"
	schedulermodule "scheduler/module"
	usermodule "user/module"

	"github.com/go-kratos/kratos/v3/middleware"
	"github.com/google/wire"
)

var ProviderSet = wire.NewSet(
	NewRegistry,
	ProvideUserClient,
	ProvideContentClient,
	ProvideEconomyClient,
	ProvideGameTownClient,
	ProvideIMClient,
	ProvideNotifyClient,
	ProvidePlatformClient,
	ProvidePushHubClient,
	ProvideSchedulerClient,
	ProvideUserConfig,
	ProvideUserModule,
	ProvideContentModule,
	ProvideEconomyModule,
	ProvideGameTownModule,
	ProvideIMModule,
	ProvideNotifyModule,
	ProvidePlatformModule,
	ProvidePushHubModule,
	ProvidePushNodeModule,
	ProvideSchedulerModule,
	ProvideBBSConfig,
	ProvideBBSModule,
	MountModules,
	ProvideServices,
	ProvideHTTPMiddlewares,
)

type Registry struct {
	modules         map[string]any
	conns           map[string]*localrpc.Conn
	clients         map[string]any
	services        []commonserver.Service
	httpMiddlewares []middleware.Middleware
}

type MountedRegistry struct {
	registry *Registry
}

type ModuleNode struct {
	Name            string
	Module          any
	Services        []commonserver.Service
	Mount           func(*localrpc.Conn, []commonserver.Service)
	External        bool
	HTTPMiddlewares []middleware.Middleware
}

func NewRegistry() *Registry {
	return &Registry{
		modules: make(map[string]any),
		conns:   make(map[string]*localrpc.Conn),
		clients: make(map[string]any),
	}
}

func MountModules(
	registry *Registry,
	userModule *usermodule.Module,
	contentModule *contentmodule.Module,
	economyModule *economymodule.Module,
	gameTownModule *gametownmodule.Module,
	imModule *immodule.Module,
	notifyModule *notifymodule.Module,
	platformModule *platformmodule.Module,
	pushHubModule *pushhubmodule.Module,
	pushNodeModule *pushnodemodule.Module,
	schedulerModule *schedulermodule.Module,
	bbsModule *bbsmodule.Module,
) *MountedRegistry {
	nodes := []ModuleNode{
		NewModuleNode(userModule.Name, userModule, userModule.Services, rpc.MountUserServices[commonserver.Service]),
		NewModuleNode(contentModule.Name, contentModule, contentModule.Services, rpc.MountContentServices[commonserver.Service]),
		NewModuleNode(economyModule.Name, economyModule, economyModule.Services, rpc.MountEconomyServices[commonserver.Service]),
		NewModuleNode(gameTownModule.Name, gameTownModule, gameTownModule.Services, rpc.MountGameTownServices[commonserver.Service]),
		NewModuleNode(imModule.Name, imModule, imModule.Services, rpc.MountIMServices[commonserver.Service]),
		NewModuleNode(notifyModule.Name, notifyModule, notifyModule.Services, rpc.MountNotifyServices[commonserver.Service]),
		NewModuleNode(platformModule.Name, platformModule, platformModule.Services, rpc.MountPlatformServices[commonserver.Service]),
		NewModuleNode(pushHubModule.Name, pushHubModule, pushHubModule.Services, rpc.MountPushHubServices[commonserver.Service]),
		NewModuleNode(pushNodeModule.Name, pushNodeModule, pushNodeModule.Services, nil),
		NewModuleNode(schedulerModule.Name, schedulerModule, schedulerModule.Services, rpc.MountSchedulerServices[commonserver.Service]),
		NewExternalModuleNode(bbsModule.Name, bbsModule, bbsModule.Services, bbsModule.HTTPMiddlewares()),
	}
	registry.Mount(nodes...)
	return &MountedRegistry{registry: registry}
}

func NewModuleNode(name string, module any, services []commonserver.Service, mount func(*localrpc.Conn, []commonserver.Service)) ModuleNode {
	return ModuleNode{Name: name, Module: module, Services: services, Mount: mount}
}

func NewExternalModuleNode(name string, module any, services []commonserver.Service, middlewares []middleware.Middleware) ModuleNode {
	return ModuleNode{Name: name, Module: module, Services: services, External: true, HTTPMiddlewares: middlewares}
}

func (r *Registry) Conn(name string) *localrpc.Conn {
	conn, ok := r.conns[name]
	if ok {
		return conn
	}
	conn = localrpc.NewConn()
	r.conns[name] = conn
	return conn
}

func (r *Registry) SetClient(name string, client any) {
	r.clients[name] = client
}

func (r *Registry) Module(name string) (any, bool) {
	module, ok := r.modules[name]
	return module, ok
}

func (r *Registry) Client(name string) (any, bool) {
	client, ok := r.clients[name]
	return client, ok
}

func (r *Registry) Mount(nodes ...ModuleNode) {
	for _, node := range nodes {
		r.modules[node.Name] = node.Module
		if node.Mount != nil {
			node.Mount(r.Conn(node.Name), node.Services)
		}
		if node.External {
			r.services = append(r.services, node.Services...)
			r.httpMiddlewares = append(r.httpMiddlewares, node.HTTPMiddlewares...)
		}
	}
}

func ProvideServices(registry *MountedRegistry) []commonserver.Service {
	return registry.registry.services
}

func ProvideHTTPMiddlewares(registry *MountedRegistry) []middleware.Middleware {
	return registry.registry.httpMiddlewares
}

func ProvideUserClient(registry *Registry) *rpc.UserClient {
	client := rpc.NewUserClient(registry.Conn("user"))
	registry.SetClient("user", client)
	return client
}

func ProvideContentClient(registry *Registry) *rpc.ContentClient {
	client := rpc.NewContentClient(registry.Conn("content"))
	registry.SetClient("content", client)
	return client
}

func ProvideEconomyClient(registry *Registry) *rpc.EconomyClient {
	client := rpc.NewEconomyClient(registry.Conn("economy"))
	registry.SetClient("economy", client)
	return client
}

func ProvideGameTownClient(registry *Registry) *rpc.GameTownClient {
	client := rpc.NewGameTownClient(registry.Conn("game_town"))
	registry.SetClient("game_town", client)
	return client
}

func ProvideIMClient(registry *Registry) *rpc.IMClient {
	client := rpc.NewIMClient(registry.Conn("im"))
	registry.SetClient("im", client)
	return client
}

func ProvideNotifyClient(registry *Registry) *rpc.NotifyClient {
	client := rpc.NewNotifyClient(registry.Conn("notify"))
	registry.SetClient("notify", client)
	return client
}

func ProvidePlatformClient(registry *Registry) *rpc.PlatformClient {
	client := rpc.NewPlatformClient(registry.Conn("platform"))
	registry.SetClient("platform", client)
	return client
}

func ProvidePushHubClient(registry *Registry) *rpc.PushHubClient {
	client := rpc.NewPushHubClient(registry.Conn("push_hub"))
	registry.SetClient("push_hub", client)
	return client
}

func ProvideSchedulerClient(registry *Registry) *rpc.SchedulerClient {
	client := rpc.NewSchedulerClient(registry.Conn("scheduler"))
	registry.SetClient("scheduler", client)
	return client
}

func ProvideUserModule(
	c *usermodule.Config,
	logger *slog.Logger,
	notifyClient *rpc.NotifyClient,
	platformClient *rpc.PlatformClient,
	schedulerClient *rpc.SchedulerClient,
) (*usermodule.Module, func(), error) {
	return usermodule.New(c, logger, notifyClient, platformClient, schedulerClient)
}

func ProvideContentModule(c *config.Bootstrap, logger *slog.Logger, schedulerClient *rpc.SchedulerClient) (*contentmodule.Module, func(), error) {
	moduleConfig := c.GetModules().GetContent()
	if !moduleEnabled(moduleConfig) {
		return nil, func() {}, fmt.Errorf("%s module is disabled", constant.ContentServiceName)
	}
	cfg, cleanup, err := contentmodule.LoadRequiredConfig(constant.ContentServiceName.String(), moduleConfig.GetBootstrap(), moduleConfig.GetConfig())
	if err != nil {
		return nil, cleanup, err
	}
	module, moduleCleanup, err := contentmodule.New(cfg, logger, schedulerClient)
	if err != nil {
		composeCleanup(moduleCleanup, cleanup)()
		return nil, func() {}, err
	}
	return module, composeCleanup(moduleCleanup, cleanup), nil
}

func ProvideEconomyModule(c *config.Bootstrap, logger *slog.Logger) (*economymodule.Module, func(), error) {
	moduleConfig := c.GetModules().GetEconomy()
	if !moduleEnabled(moduleConfig) {
		return nil, func() {}, fmt.Errorf("%s module is disabled", constant.EconomyServiceName)
	}
	cfg, cleanup, err := economymodule.LoadRequiredConfig(constant.EconomyServiceName.String(), moduleConfig.GetBootstrap(), moduleConfig.GetConfig())
	if err != nil {
		return nil, cleanup, err
	}
	module, moduleCleanup, err := economymodule.New(cfg, logger)
	if err != nil {
		composeCleanup(moduleCleanup, cleanup)()
		return nil, func() {}, err
	}
	return module, composeCleanup(moduleCleanup, cleanup), nil
}

func ProvideGameTownModule(c *config.Bootstrap, logger *slog.Logger) (*gametownmodule.Module, func(), error) {
	moduleConfig := c.GetModules().GetGameTown()
	if !moduleEnabled(moduleConfig) {
		return nil, func() {}, fmt.Errorf("%s module is disabled", constant.GameTownServiceName)
	}
	cfg, cleanup, err := gametownmodule.LoadRequiredConfig(constant.GameTownServiceName.String(), moduleConfig.GetBootstrap(), moduleConfig.GetConfig())
	if err != nil {
		return nil, cleanup, err
	}
	module, moduleCleanup, err := gametownmodule.New(cfg, logger)
	if err != nil {
		composeCleanup(moduleCleanup, cleanup)()
		return nil, func() {}, err
	}
	return module, composeCleanup(moduleCleanup, cleanup), nil
}

func ProvideIMModule(c *config.Bootstrap, logger *slog.Logger) (*immodule.Module, func(), error) {
	moduleConfig := c.GetModules().GetIm()
	if !moduleEnabled(moduleConfig) {
		return nil, func() {}, fmt.Errorf("%s module is disabled", constant.IMServiceName)
	}
	cfg, cleanup, err := immodule.LoadRequiredConfig(constant.IMServiceName.String(), moduleConfig.GetBootstrap(), moduleConfig.GetConfig())
	if err != nil {
		return nil, cleanup, err
	}
	module, moduleCleanup, err := immodule.New(cfg, logger)
	if err != nil {
		composeCleanup(moduleCleanup, cleanup)()
		return nil, func() {}, err
	}
	return module, composeCleanup(moduleCleanup, cleanup), nil
}

func ProvideNotifyModule(c *config.Bootstrap, logger *slog.Logger, userClient *rpc.UserClient, contentClient *rpc.ContentClient) (*notifymodule.Module, func(), error) {
	moduleConfig := c.GetModules().GetNotify()
	if !moduleEnabled(moduleConfig) {
		return nil, func() {}, fmt.Errorf("%s module is disabled", constant.NotifyServiceName)
	}
	cfg, cleanup, err := notifymodule.LoadRequiredConfig(constant.NotifyServiceName.String(), moduleConfig.GetBootstrap(), moduleConfig.GetConfig())
	if err != nil {
		return nil, cleanup, err
	}
	module, moduleCleanup, err := notifymodule.New(cfg, logger, userClient, contentClient)
	if err != nil {
		composeCleanup(moduleCleanup, cleanup)()
		return nil, func() {}, err
	}
	return module, composeCleanup(moduleCleanup, cleanup), nil
}

func ProvidePlatformModule(c *config.Bootstrap, logger *slog.Logger) (*platformmodule.Module, func(), error) {
	moduleConfig := c.GetModules().GetPlatform()
	if !moduleEnabled(moduleConfig) {
		return nil, func() {}, fmt.Errorf("%s module is disabled", constant.PlatformServiceName)
	}
	cfg, cleanup, err := platformmodule.LoadRequiredConfig(constant.PlatformServiceName.String(), moduleConfig.GetBootstrap(), moduleConfig.GetConfig())
	if err != nil {
		return nil, cleanup, err
	}
	module, moduleCleanup, err := platformmodule.New(cfg, logger)
	if err != nil {
		composeCleanup(moduleCleanup, cleanup)()
		return nil, func() {}, err
	}
	return module, composeCleanup(moduleCleanup, cleanup), nil
}

func ProvidePushHubModule(c *config.Bootstrap, logger *slog.Logger) (*pushhubmodule.Module, func(), error) {
	moduleConfig := c.GetModules().GetPushHub()
	if !moduleEnabled(moduleConfig) {
		return nil, func() {}, fmt.Errorf("%s module is disabled", constant.PushHubServiceName)
	}
	cfg, cleanup, err := pushhubmodule.LoadRequiredConfig(constant.PushHubServiceName.String(), moduleConfig.GetBootstrap(), moduleConfig.GetConfig())
	if err != nil {
		return nil, cleanup, err
	}
	module, moduleCleanup, err := pushhubmodule.New(cfg, logger)
	if err != nil {
		composeCleanup(moduleCleanup, cleanup)()
		return nil, func() {}, err
	}
	return module, composeCleanup(moduleCleanup, cleanup), nil
}

func ProvidePushNodeModule(c *config.Bootstrap, logger *slog.Logger) (*pushnodemodule.Module, func(), error) {
	moduleConfig := c.GetModules().GetPushNode()
	if !moduleEnabled(moduleConfig) {
		return nil, func() {}, fmt.Errorf("%s module is disabled", constant.PushNodeServiceName)
	}
	cfg, cleanup, err := pushnodemodule.LoadRequiredConfig(constant.PushNodeServiceName.String(), moduleConfig.GetBootstrap(), moduleConfig.GetConfig())
	if err != nil {
		return nil, cleanup, err
	}
	module, moduleCleanup, err := pushnodemodule.New(cfg, logger)
	if err != nil {
		composeCleanup(moduleCleanup, cleanup)()
		return nil, func() {}, err
	}
	return module, composeCleanup(moduleCleanup, cleanup), nil
}

func ProvideSchedulerModule(c *config.Bootstrap, logger *slog.Logger, userClient *rpc.UserClient, contentClient *rpc.ContentClient) (*schedulermodule.Module, func(), error) {
	moduleConfig := c.GetModules().GetScheduler()
	if !moduleEnabled(moduleConfig) {
		return nil, func() {}, fmt.Errorf("%s module is disabled", constant.SchedulerServiceName)
	}
	cfg, cleanup, err := schedulermodule.LoadRequiredConfig(constant.SchedulerServiceName.String(), moduleConfig.GetBootstrap(), moduleConfig.GetConfig())
	if err != nil {
		return nil, cleanup, err
	}
	module, moduleCleanup, err := schedulermodule.New(cfg, logger, userClient, contentClient)
	if err != nil {
		composeCleanup(moduleCleanup, cleanup)()
		return nil, func() {}, err
	}
	return module, composeCleanup(moduleCleanup, cleanup), nil
}

func ProvideBBSModule(
	c *bbsmodule.Config,
	userClient *rpc.UserClient,
	contentClient *rpc.ContentClient,
	economyClient *rpc.EconomyClient,
	notifyClient *rpc.NotifyClient,
	platformClient *rpc.PlatformClient,
) (*bbsmodule.Module, error) {
	return bbsmodule.New(c, userClient, contentClient, economyClient, notifyClient, platformClient)
}

func ProvideUserConfig(c *config.Bootstrap) (*usermodule.Config, func(), error) {
	moduleConfig := c.GetModules().GetUser()
	if !moduleEnabled(moduleConfig) {
		return nil, func() {}, fmt.Errorf("%s module is disabled", constant.UserServiceName)
	}
	return usermodule.LoadRequiredConfig(c.GetServer().GetName(), moduleConfig.GetBootstrap(), moduleConfig.GetConfig())
}

func ProvideBBSConfig(c *config.Bootstrap) (*bbsmodule.Config, func(), error) {
	moduleConfig := c.GetModules().GetBbs()
	if !moduleEnabled(moduleConfig) {
		return nil, func() {}, fmt.Errorf("bbs module is disabled")
	}
	return bbsmodule.LoadRequiredConfig(c.GetServer().GetName(), moduleConfig.GetBootstrap(), moduleConfig.GetConfig())
}

func composeCleanup(cleanups ...func()) func() {
	return func() {
		for _, cleanup := range cleanups {
			if cleanup != nil {
				cleanup()
			}
		}
	}
}

func moduleEnabled(module *config.Module) bool {
	return module != nil && module.GetEnabled()
}
