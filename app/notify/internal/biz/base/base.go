package base

import (
	"common/pkg/client"
	"common/pkg/client/rpc"
	"notify/internal/conf"
	"notify/internal/data/ent/gen"

	"github.com/go-kratos/kratos/v2/log"
)

type BaseDomain struct {
	Conf       *conf.Bootstrap
	Log        *log.Helper
	Db         *gen.Client
	Consul     *client.ConsulClient
	UserClient *rpc.UserClient
	Infra      *rpc.InfraClient
	Redis      *client.RedisClient
	Rabbitmq   *client.RabbitMQClient
}

func NewBaseDomain(conf *conf.Bootstrap, log *log.Helper, db *gen.Client, consul *client.ConsulClient, userClient *rpc.UserClient, infra *rpc.InfraClient, redis *client.RedisClient, rabbitmq *client.RabbitMQClient) *BaseDomain {
	return &BaseDomain{
		Conf:       conf,
		Log:        log,
		Db:         db,
		Consul:     consul,
		UserClient: userClient,
		Infra:      infra,
		Redis:      redis,
		Rabbitmq:   rabbitmq,
	}
}
