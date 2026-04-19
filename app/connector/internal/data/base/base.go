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

func NewBaseData(conf *conf.Bootstrap, logger log.Logger, redis *commonClient.RedisClient) *BaseData {
	return &BaseData{
		Conf:  conf,
		Log:   log.NewHelper(logger),
		Redis: redis,
	}
}
