package config

import (
	serverconfig "common/pkg/server/config"
	"common/proto/gen/common"
)

// LoadConfig 加载 game_idle 配置，并注册可热更新配置。
func NewConfig(bootstrapPath string, path string) (*Bootstrap, *common.Bootstrap, func(), error) {
	c, bc, hot, cleanup, err := serverconfig.LoadConfig[*Bootstrap](bootstrapPath, path)
	if err != nil {
		return nil, nil, cleanup, err
	}
	if err := hot.BindProtoHotFields(); err != nil {
		cleanup()
		return nil, nil, cleanup, err
	}
	return c, bc, cleanup, nil
}
