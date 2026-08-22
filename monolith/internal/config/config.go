package config

import (
	commonmodule "common/pkg/module"
	commonserver "common/pkg/server"
	"common/proto/gen/common"
	"fmt"

	"google.golang.org/protobuf/types/known/structpb"
)

func LoadConfig(bootstrapPath string, path string) (*Bootstrap, *common.Bootstrap, func(), error) {
	c, bc, hot, cleanup, err := commonserver.LoadConfig[*Bootstrap](bootstrapPath, path)
	if err != nil {
		return nil, nil, cleanup, err
	}

	if err := hot.BindProtoHotFields(); err != nil {
		cleanup()
		return nil, nil, cleanup, err
	}

	return c, bc, cleanup, nil
}

// RuntimeConfig 返回模块化单体共享的运行时配置。
func (c *Bootstrap) RuntimeConfig() *commonmodule.RuntimeConfig {
	return &commonmodule.RuntimeConfig{
		Server:        c.GetServer(),
		Http:          c.GetHttp(),
		Grpc:          c.GetGrpc(),
		Consul:        c.GetConsul(),
		Database:      c.GetDatabase(),
		Redis:         c.GetRedis(),
		Nats:          c.GetNats(),
		Trace:         c.GetTrace(),
		Observability: c.GetObservability(),
		Dtm:           c.GetDtm(),
		Alert:         c.GetAlert(),
		Event:         c.GetEvent(),
	}
}

// ModuleConfig 返回指定模块的私有配置。
func (c *Bootstrap) ModuleConfig(name string) (*structpb.Struct, error) {
	modules := c.GetModules()
	if len(modules) == 0 {
		return nil, fmt.Errorf("modules config is required")
	}
	value, ok := modules[name]
	if !ok || value == nil {
		return nil, fmt.Errorf("%s module config is required", name)
	}
	return value, nil
}
