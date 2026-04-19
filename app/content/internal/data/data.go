package data

import (
	"common/api/gen/common"
	commonClient "common/pkg/client"
	"common/pkg/util/jwt"
	"content/internal/conf"
	"content/internal/data/base"
	"content/internal/data/client"
	"content/internal/data/repo"

	"github.com/google/wire"
)

// DataProviderSet is data providers.
var DataProviderSet = wire.NewSet(
	base.NewBaseData,

	client.NewDataBaseClient,
	ProvideRedis,
	ProvideConsul,
	ProvideRabbitMQ,
	commonClient.NewConsulClient,
	commonClient.NewRedisClient,
	commonClient.NewRabbitMQClient,

	repo.NewArticleRepo,
	repo.NewCommentRepo,
	repo.NewCommentActionRecordRepo,
	repo.NewArticlePostscriptRepo,
	repo.NewArticleActionRecordRepo,
	repo.NewDomainRepo,
	repo.NewTagRepo,

	jwt.NewTokenCache,
)

func ProvideRedis(c *conf.Bootstrap) *common.Redis {
	return c.Data.Redis
}

func ProvideConsul(c *conf.Bootstrap) *common.Consul {
	return c.Data.Consul
}

func ProvideRabbitMQ(c *conf.Bootstrap) *common.RabbitMQ {
	return c.Data.Rabbitmq
}
