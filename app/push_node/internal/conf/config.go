package conf

import (
	"common/proto/gen/common"
	commonserver "common/pkg/server"
)

// LoadConfig 加载本模块配置，并注册可热更新配置。
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
