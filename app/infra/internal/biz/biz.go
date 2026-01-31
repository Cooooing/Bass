package biz

import (
	"common/pkg/client"
	"common/pkg/util"
	"infra/internal/conf"
	"infra/internal/data/ent/gen"

	"github.com/go-kratos/kratos/v2/log"
	"github.com/google/wire"
)

// BizProviderSet is biz providers.
var BizProviderSet = wire.NewSet(
	NewBaseDomain,
	util.NewEventPool,
)

type BaseDomain struct {
	conf      *conf.Bootstrap
	log       *log.Helper
	db        *gen.Client
	etcd      *client.EtcdClient
	redis     *client.RedisClient
	rabbitmq  *client.RabbitMQClient
	eventPool *util.EventPool
}

func NewBaseDomain(conf *conf.Bootstrap, log *log.Helper, db *gen.Client, etcd *client.EtcdClient, redis *client.RedisClient, rabbitmq *client.RabbitMQClient, eventPool *util.EventPool) *BaseDomain {
	return &BaseDomain{
		conf:      conf,
		log:       log,
		db:        db,
		etcd:      etcd,
		redis:     redis,
		rabbitmq:  rabbitmq,
		eventPool: eventPool,
	}
}
