package base

import (
	commonClient "common/pkg/client"
	"content/internal/conf"
	"content/internal/data/ent/gen"

	"github.com/go-kratos/kratos/v2/log"
)

type BaseData struct {
	Conf   *conf.Bootstrap
	Log    *log.Helper
	Db     *gen.Client
	Consul *commonClient.ConsulClient
	Redis  *commonClient.RedisClient
}

func NewBaseData(conf *conf.Bootstrap, logger log.Logger, db *gen.Client, consul *commonClient.ConsulClient, redis *commonClient.RedisClient) *BaseData {
	return &BaseData{
		Conf:   conf,
		Log:    log.NewHelper(logger),
		Consul: consul,
		Db:     db,
		Redis:  redis,
	}
}
