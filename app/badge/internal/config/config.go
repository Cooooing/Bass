package config

import (
	serverconfig "common/pkg/server/config"
	"common/proto/gen/common"
)

// LoadConfig 加载 badge 配置
func NewConfig(bootstrapPath string, path string) (*Bootstrap, *common.Bootstrap, func(), error) {
	c, bc, _, cleanup, err := serverconfig.LoadConfig[*Bootstrap](bootstrapPath, path)
	if err != nil {
		return nil, nil, cleanup, err
	}
	return c, bc, cleanup, nil
}
