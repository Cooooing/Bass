package config

import (
	commonserver "common/pkg/server"
	"common/proto/gen/common"
)

// LoadConfig 加载本模块配置，并注册可热更新配置。
func LoadConfig(
	bootstrapPath string,
	path string,
) (*Bootstrap, *common.Bootstrap, func(), error) {
	c, bc, hot, cleanup, err := commonserver.LoadConfig[*Bootstrap](bootstrapPath, path)
	if err != nil {
		return nil, nil, cleanup, err
	}

	if err := hot.BindProtoHotFields(
		&c.Business.App,
		&c.Business.Avatar,
		&c.Alert.LarkWebhook,
		&c.Event.Outbox,
		&c.Event.DeadLetter,
	); err != nil {
		cleanup()
		return nil, nil, cleanup, err
	}

	return c, bc, cleanup, nil
}
