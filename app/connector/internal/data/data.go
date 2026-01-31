package data

import (
	commonClient "common/pkg/client"
	commonModel "common/pkg/model"
	"connector/internal/conf"

	"github.com/go-kratos/kratos/v2/log"
	"github.com/google/wire"
	"github.com/jinzhu/copier"
)

// DataProviderSet is data providers.
var DataProviderSet = wire.NewSet(
	NewBaseData,
	NewRedisClient,

	NewSessionCache,
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

func NewRedisClient(log *log.Helper, conf *conf.Bootstrap) (*commonClient.RedisClient, func(), error) {
	c := &commonModel.RedisConf{}
	err := copier.Copy(c, conf.Data.Redis)
	if err != nil {
		return nil, nil, err
	}
	return commonClient.NewRedisClient(log, c)
}
