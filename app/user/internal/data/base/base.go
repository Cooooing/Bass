package base

import (
	commonClient "common/pkg/client"
	"user/internal/conf"
	"user/internal/data/ent/gen"

	"github.com/go-kratos/kratos/v2/log"
)

type BaseData struct {
	Conf     *conf.Bootstrap
	Log      *log.Helper
	Db       *gen.Client
	consul   *commonClient.ConsulClient
	Redis    *commonClient.RedisClient
	Rabbitmq *commonClient.RabbitMQClient
}

func NewBaseData(conf *conf.Bootstrap, log *log.Helper, db *gen.Client, consul *commonClient.ConsulClient, redis *commonClient.RedisClient, rabbitmq *commonClient.RabbitMQClient) *BaseData {
	return &BaseData{
		Conf:     conf,
		Log:      log,
		Etcd:     etcd,
		Db:       db,
		Redis:    redis,
		Rabbitmq: rabbitmq,
	}
}
