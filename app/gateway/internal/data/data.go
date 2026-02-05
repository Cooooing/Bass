package data

import (
	commonClient "common/pkg/client"
	commonModel "common/pkg/model"
	"gateway/internal/conf"

	"github.com/go-kratos/kratos/v2/log"
	"github.com/google/wire"
	"github.com/jinzhu/copier"
)

// DataProviderSet is data providers.
var DataProviderSet = wire.NewSet(
	NewBaseData,

	NewConsulClient,
	NewRedisClient,
	NewRabbitMQClient,
)

type BaseData struct {
	conf     *conf.Bootstrap
	log      *log.Helper
	consul   *commonClient.ConsulClient
	redis    *commonClient.RedisClient
	rabbitmq *commonClient.RabbitMQClient
}

func NewBaseData(conf *conf.Bootstrap, log *log.Helper, consul *commonClient.ConsulClient, redis *commonClient.RedisClient, rabbitmq *commonClient.RabbitMQClient) *BaseData {
	return &BaseData{
		conf:     conf,
		log:      log,
		consul:   consul,
		redis:    redis,
		rabbitmq: rabbitmq,
	}
}

func NewConsulClient(log *log.Helper, conf *conf.Bootstrap) (*commonClient.ConsulClient, func(), error) {
	c := &commonModel.ConsulConf{}
	err := copier.Copy(c, conf.Registry.Consul)
	if err != nil {
		return nil, nil, err
	}
	return commonClient.NewConsulClient(log, c)
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
