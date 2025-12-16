package biz

import (
	commonClient "common/pkg/client"
	"common/pkg/util"
	"user/internal/conf"
	"user/internal/data/ent/gen"

	"github.com/go-kratos/kratos/v2/log"
	"github.com/google/wire"
)

// BizProviderSet is biz providers.
var BizProviderSet = wire.NewSet(
	NewBaseDomain,

	util.NewTokenRepo,
	NewTokenService,
	util.NewEventPool,

	NewAuthenticationDomain,
	NewUserDomain,
	NewUserRelationDomain,
)

type BaseDomain struct {
	conf      *conf.Bootstrap
	log       *log.Helper
	db        *gen.Client
	etcd      *commonClient.EtcdClient
	redis     *commonClient.RedisClient
	rabbitmq  *commonClient.RabbitMQClient
	eventPool *util.EventPool
}

func NewBaseDomain(conf *conf.Bootstrap, log *log.Helper, db *gen.Client, etcd *commonClient.EtcdClient, redis *commonClient.RedisClient, rabbitmq *commonClient.RabbitMQClient, eventPool *util.EventPool) *BaseDomain {
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
