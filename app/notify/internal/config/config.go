package config

import (
	serverconfig "common/pkg/server/config"
	"common/proto/gen/common"
)

func NewConfig(bootstrapPath string, path string) (*Bootstrap, *common.Bootstrap, func(), error) {
	c, bc, hot, cleanup, err := serverconfig.LoadConfig[*Bootstrap](bootstrapPath, path)
	if err != nil {
		return nil, nil, cleanup, err
	}

	if err := hot.BindProtoHotFields(
		&c.Business.NotificationRateLimit,
		&c.Alert.LarkWebhook,
		&c.Event.Inbox,
		&c.Event.DeadLetter,
	); err != nil {
		cleanup()
		return nil, nil, cleanup, err
	}

	return c, bc, cleanup, nil
}
