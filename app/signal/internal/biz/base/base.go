package base

import (
	"common/pkg/client"
	"common/pkg/util"
	"signal/internal/conf"
	"signal/internal/data/ent/gen"

	"github.com/go-kratos/kratos/v2/log"
)

type BaseDomain struct {
	Conf      *conf.Bootstrap
	Log       *log.Helper
	Db        *gen.Client
	Etcd      *client.EtcdClient
	Redis     *client.RedisClient
	Rabbitmq  *client.RabbitMQClient
	EventPool *util.EventPool
}

func NewBaseDomain(conf *conf.Bootstrap, log *log.Helper, db *gen.Client, etcd *client.EtcdClient, redis *client.RedisClient, rabbitmq *client.RabbitMQClient, eventPool *util.EventPool) *BaseDomain {
	return &BaseDomain{
		Conf:      conf,
		Log:       log,
		Db:        db,
		Etcd:      etcd,
		Redis:     redis,
		Rabbitmq:  rabbitmq,
		EventPool: eventPool,
	}
}
