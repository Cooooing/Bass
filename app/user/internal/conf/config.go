package conf

import (
	"common/api/gen/common"
	commonserver "common/pkg/util/server"
)

// LoadConfig 加载本模块配置，并注册可热更新配置。
func LoadConfig(bootstrapPath string, path string) (*Bootstrap, *common.Bootstrap, func(), error) {
	c, bc, hot, cleanup, err := commonserver.LoadConfig[*Bootstrap](bootstrapPath, path)
	if err != nil {
		return nil, nil, cleanup, err
	}

	if err := hot.BindProtoHotFields(
		&c.Server.App,
		&c.Server.Avatar,
		&c.Server.Jwt,
	); err != nil {
		cleanup()
		return nil, nil, cleanup, err
	}

	return c, bc, cleanup, nil
}
