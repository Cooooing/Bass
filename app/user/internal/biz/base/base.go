package base

import (
	"common/pkg/client"
	"common/pkg/util"
	"user/internal/conf"
	"user/internal/data/ent/gen"

	"github.com/go-kratos/kratos/v2/log"
)

type BaseDomain struct {
	Conf      *conf.Bootstrap
	Log       *log.Helper
	Db        *gen.Client
	Consul    *client.ConsulClient
	Redis     *client.RedisClient
	Rabbitmq  *client.RabbitMQClient
	EventPool *util.EventPool
}

func NewBaseDomain(conf *conf.Bootstrap, logger log.Logger, db *gen.Client, consul *client.ConsulClient, redis *client.RedisClient, rabbitmq *client.RabbitMQClient, eventPool *util.EventPool) *BaseDomain {
	return &BaseDomain{
		Conf:      conf,
		Log:       log.NewHelper(logger),
		Db:        db,
		Consul:    consul,
		Redis:     redis,
		Rabbitmq:  rabbitmq,
		EventPool: eventPool,
	}
}
