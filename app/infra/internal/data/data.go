package data

import (
	commonClient "common/pkg/client"
	commonModel "common/pkg/model"
	"common/pkg/util"
	"infra/internal/conf"
	"infra/internal/data/base"
	"infra/internal/data/client"
	"infra/internal/data/oss"

	"github.com/go-kratos/kratos/v2/log"
	"github.com/google/wire"
	"github.com/jinzhu/copier"
)

// DataProviderSet is data providers.
var DataProviderSet = wire.NewSet(
	base.NewBaseData,
	oss.ProviderSet,
	client.NewDataBaseClient,
	NewEtcdClient,
	NewRedisClient,
	NewRabbitMQClient,
	NewObjectStorageRepo,

	util.NewTokenCache,
)

func NewEtcdClient(log *log.Helper, conf *conf.Bootstrap) (*commonClient.EtcdClient, func(), error) {
	c := &commonModel.EtcdConf{}
	err := copier.Copy(c, conf.Registry.Etcd)
	if err != nil {
		return nil, nil, err
	}
	return commonClient.NewEtcdClient(log, c)
}

func NewRedisClient(log *log.Helper, conf *conf.Bootstrap) (*commonClient.RedisClient, func(), error) {
	c := &commonModel.RedisConf{}
	err := copier.Copy(c, conf.Data.Redis)
	if err != nil {
		return nil, nil, err
	}
	return commonClient.NewRedisClient(log, c)
}

func NewRabbitMQClient(log *log.Helper, conf *conf.Bootstrap) (*commonClient.RabbitMQClient, func(), error) {
	c := &commonModel.RabbitmqConf{}
	err := copier.Copy(c, conf.Data.Rabbitmq)
	if err != nil {
		return nil, nil, err
	}
	return commonClient.NewRabbitMQClient(log, c)
}
