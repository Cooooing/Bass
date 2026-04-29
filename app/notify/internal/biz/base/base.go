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
}

func NewBaseDomain(conf *conf.Bootstrap, logger log.Logger, db *gen.Client, consul *client.ConsulClient, userClient *rpc.UserClient, infra *rpc.InfraClient, redis *client.RedisClient) *BaseDomain {
	return &BaseDomain{
		Conf:       conf,
		Log:        log.NewHelper(logger),
		Db:         db,
		Consul:     consul,
		UserClient: userClient,
		Infra:      infra,
		Redis:      redis,
	}
}
