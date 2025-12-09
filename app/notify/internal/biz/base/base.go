package base

import (
	"common/pkg/client"
	"notify/internal/conf"
	"notify/internal/data/ent/gen"

	"github.com/go-kratos/kratos/v2/log"
)

type BaseDomain struct {
	Conf     *conf.Bootstrap
	Log      *log.Helper
	Db       *gen.Client
	Etcd     *client.EtcdClient
	Redis    *client.RedisClient
	Rabbitmq *client.RabbitMQClient
}

func NewBaseDomain(conf *conf.Bootstrap, log *log.Helper, db *gen.Client, etcd *client.EtcdClient, redis *client.RedisClient, rabbitmq *client.RabbitMQClient) *BaseDomain {
	return &BaseDomain{
		Conf:     conf,
		Log:      log,
		Db:       db,
		Etcd:     etcd,
		Redis:    redis,
		Rabbitmq: rabbitmq,
	}
}
