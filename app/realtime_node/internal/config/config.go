package config

import (
	serverconfig "common/pkg/server/config"
	"common/proto/gen/common"
)

func NewConfig(bootstrapPath, path string) (*Bootstrap, *common.Bootstrap, func(), error) {
	config, bootstrap, hot, cleanup, err := serverconfig.LoadConfig[*Bootstrap](bootstrapPath, path)
	if err != nil {
		return nil, nil, cleanup, err
	}
	if err := hot.BindProtoHotFields(); err != nil {
		cleanup()
		return nil, nil, cleanup, err
	}
	return config, bootstrap, cleanup, nil
}
