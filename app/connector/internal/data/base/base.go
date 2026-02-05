package base

import (
	commonClient "common/pkg/client"
	"connector/internal/conf"

	"github.com/go-kratos/kratos/v2/log"
)

type BaseData struct {
	Conf  *conf.Bootstrap
	Log   *log.Helper
	Redis *commonClient.RedisClient
}

func NewBaseData(conf *conf.Bootstrap, log *log.Helper, redis *commonClient.RedisClient) *BaseData {
	return &BaseData{
		Conf:  conf,
		Log:   log,
		Redis: redis,
	}
}
