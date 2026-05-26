package data

import (
	"common/api/gen/common"
	commonClient "common/pkg/client"
	"common/pkg/client/rpc"
	"common/pkg/util/jwt"
	"content/internal/conf"
	"content/internal/data/client"
	"content/internal/data/repo"

	"github.com/google/wire"
)

// DataProviderSet 是 data 层依赖集合。
var DataProviderSet = wire.NewSet(
	client.NewDataBaseClient,
	ProvideRedis,
	ProvideConsul,
	ProvideNats,
	commonClient.NewConsulClient,
	commonClient.NewRedisClient,
	commonClient.NewNatsClient,
	rpc.ProvideUserClient,

	client.ProvideTx,

	repo.NewArticleRepo,
	repo.NewCommentRepo,
	repo.NewCommentActionRecordRepo,
	repo.NewArticlePostscriptRepo,
	repo.NewArticleActionRecordRepo,
	repo.NewOutboxEventRepo,
	repo.NewDomainRepo,
	repo.NewTagRepo,
	repo.NewUserClient,
	NewOutboxPublisher,

	jwt.NewTokenCache,
)

func ProvideRedis(c *conf.Bootstrap) *common.Redis {
	return c.Data.Redis
}

func ProvideConsul(c *conf.Bootstrap) *common.Consul {
	return c.Data.Consul
}

func ProvideNats(c *conf.Bootstrap) *common.Nats {
	return c.Data.Nats
}
