package data

import (
	commonClient "common/pkg/client"
	commonModel "common/pkg/model"
	"user/internal/conf"
	"user/internal/data/base"
	"user/internal/data/client"

	"github.com/go-kratos/kratos/v2/log"
	"github.com/google/wire"
	"github.com/jinzhu/copier"
)

// DataProviderSet is data providers.
var DataProviderSet = wire.NewSet(
	base.NewBaseData,

	client.NewDataBaseClient,
	NewEtcdClient,
	NewRedisClient,
	NewRabbitMQClient,

	NewUserRepo,
	NewUserRelationRepo,
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
