package module

import (
	"bff_bbs_admin/internal/biz"
	"bff_bbs_admin/internal/config"
	"bff_bbs_admin/internal/data"
	"bff_bbs_admin/internal/service"
	commonmodule "common/pkg/module"
	commonserver "common/pkg/server"

	"github.com/google/wire"
)

// ProviderSet 提供 bff_bbs_admin 模块的依赖项。
var ProviderSet = wire.NewSet(provideBootstrap, data.ModuleProviderSet, biz.BizProviderSet, service.ServiceProviderSet, newModule)

type Config = commonmodule.Config[*config.Bootstrap]

type Module struct {
	Name     string
	Services []commonserver.Service
}

func newModule(config *Config, services []commonserver.Service) *Module {
	return &Module{Name: config.Server().GetName(), Services: services}
}

func provideBootstrap(c *Config) *config.Bootstrap { return c.Bootstrap() }

// Build 构造管理端 BFF 模块并返回单体可收集的模块能力。
func Build(runtime *commonmodule.Runtime, name string) (commonmodule.Mounted, func(), error) {
	values, err := runtime.Values(name)
	if err != nil {
		return commonmodule.Mounted{}, func() {}, err
	}
	moduleConfig, err := commonmodule.NewConfig(runtime.Config, values, name, func() *config.Bootstrap { return &config.Bootstrap{} })
	if err != nil {
		return commonmodule.Mounted{}, func() {}, err
	}
	module, err := wireModule(moduleConfig)
	if err != nil {
		return commonmodule.Mounted{}, func() {}, err
	}
	return commonmodule.Mounted{Module: module, Services: module.Services}, func() {}, nil
}

func Descriptor() commonmodule.Descriptor {
	return commonmodule.NewDescriptor(Build, commonmodule.WithExternal())
}
