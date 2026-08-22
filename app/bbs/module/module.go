package module

import (
	"bbs/internal/biz"
	"bbs/internal/config"
	"bbs/internal/data"
	bbsserver "bbs/internal/server"
	"bbs/internal/service"
	"common/pkg/client/rpc"
	commonmodule "common/pkg/module"
	"common/pkg/server"
	userv1 "common/proto/gen/user/v1"

	"github.com/go-kratos/kratos/v3/middleware"
	"github.com/google/wire"
)

// ProviderSet 提供 BBS 模块的依赖项。
var ProviderSet = wire.NewSet(
	rpc.LocalProviderSet,
	data.ModuleProviderSet,
	biz.BizProviderSet,
	service.ServiceProviderSet,
	newModule,
)

type Config = commonmodule.Config[*config.Bootstrap]

type Module struct {
	Name           string
	Services       []server.Service
	UserAuthClient userv1.AuthServiceClient
}

func newModule(config *Config, services []server.Service, userAuthClient userv1.AuthServiceClient) *Module {
	return &Module{Name: config.Server().GetName(), Services: services, UserAuthClient: userAuthClient}
}

func (m *Module) HTTPMiddlewares() []middleware.Middleware {
	return bbsserver.NewHTTPAuthMiddlewares(m.UserAuthClient)
}

// Build 构造 BBS 模块并返回单体可收集的模块能力。
func Build(runtime *commonmodule.Runtime, name string) (commonmodule.Mounted, func(), error) {
	values, err := runtime.Values(name)
	if err != nil {
		return commonmodule.Mounted{}, func() {}, err
	}
	moduleConfig, err := commonmodule.NewConfig(runtime.Config, values, name, func() *config.Bootstrap { return &config.Bootstrap{} })
	if err != nil {
		return commonmodule.Mounted{}, func() {}, err
	}
	module, err := wireModule(moduleConfig, runtime.Clients)
	if err != nil {
		return commonmodule.Mounted{}, func() {}, err
	}
	return commonmodule.Mounted{
		Module:          module,
		Services:        module.Services,
		HTTPMiddlewares: module.HTTPMiddlewares(),
	}, func() {}, nil
}

func Descriptor() commonmodule.Descriptor {
	return commonmodule.NewDescriptor(
		Build,
		commonmodule.WithExternal(),
	)
}
