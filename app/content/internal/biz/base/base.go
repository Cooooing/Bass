package base

import (
	"common/pkg/client"
	"common/pkg/client/rpc"
	"common/pkg/util"
	"content/internal/conf"
	"content/internal/data/ent/gen"

	"github.com/go-kratos/kratos/v2/log"
)

type BaseDomain struct {
	Conf       *conf.Bootstrap
	Log        *log.Helper
	Db         *gen.Client
	Consul     *client.ConsulClient
	UserClient *rpc.UserClient
	Redis      *client.RedisClient
	Rabbitmq   *client.RabbitMQClient
	EventPool  *util.EventPool
}

func NewBaseDomain(conf *conf.Bootstrap, logger log.Logger, db *gen.Client, consul *client.ConsulClient, userClient *rpc.UserClient, redis *client.RedisClient, rabbitmq *client.RabbitMQClient, eventPool *util.EventPool) *BaseDomain {
	return &BaseDomain{
		Conf:       conf,
		Log:        log.NewHelper(logger),
		Db:         db,
		Consul:     consul,
		UserClient: userClient,
		Redis:      redis,
		Rabbitmq:   rabbitmq,
		EventPool:  eventPool,
	}
}
