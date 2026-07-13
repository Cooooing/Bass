package conf

import (
	commonserver "common/pkg/server"
	"common/proto/gen/common"
)

func LoadConfig(bootstrapPath string, path string) (*Bootstrap, *common.Bootstrap, func(), error) {
	c, bc, hot, cleanup, err := commonserver.LoadConfig[*Bootstrap](bootstrapPath, path)
	if err != nil {
		return nil, nil, cleanup, err
	}

	if err := hot.BindProtoHotFields(
		&c.Notify.NotificationRateLimit,
		&c.Alert.LarkWebhook,
		&c.Event.Inbox,
		&c.Event.DeadLetter,
	); err != nil {
		cleanup()
		return nil, nil, cleanup, err
	}

	return c, bc, cleanup, nil
}
