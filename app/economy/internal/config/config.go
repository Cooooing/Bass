package config

import (
	commonserver "common/pkg/server"
	"common/proto/gen/common"
)

// LoadConfig 加载 economy 配置
func LoadConfig(bootstrapPath string, path string) (*Bootstrap, *common.Bootstrap, func(), error) {
	c, bc, _, cleanup, err := commonserver.LoadConfig[*Bootstrap](bootstrapPath, path)
	if err != nil {
		return nil, nil, cleanup, err
	}
	return c, bc, cleanup, nil
}
