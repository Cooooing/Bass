package data

import (
	"bbs/internal/conf"
	"common/api/gen/common"
	commonClient "common/pkg/client"
	"common/pkg/client/rpc"
	"common/pkg/util/jwt"

	"github.com/google/wire"
)

// DataProviderSet 是 data 层依赖集合。
var DataProviderSet = wire.NewSet(
	ProvideConsul,
	ProvideRedis,
	commonClient.NewConsulClient,
	commonClient.NewRedisClient,
	jwt.NewTokenCache,
	rpc.ProvideUserClient,
	rpc.ProvideContentClient,
	rpc.ProvideNotifyClient,
	NewUserRepo,
	NewContentRepo,
	NewNotifyRepo,
)

func ProvideRedis(c *conf.Bootstrap) *common.Redis {
	return c.Data.Redis
}

func ProvideConsul(c *conf.Bootstrap) *common.Consul {
	return c.Data.Consul
}
