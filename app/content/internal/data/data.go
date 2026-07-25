package data

import (
	commonClient "common/pkg/client"
	"common/pkg/client/rpc"
	"common/proto/gen/common"
	"content/internal/config"
	"content/internal/data/client"
	"content/internal/data/repo"

	"github.com/google/wire"
)

var DataProviderSet = wire.NewSet(
	client.NewDataBaseClient,
	ProvideRedis,
	ProvideConsul,
	ProvideNats,
	commonClient.NewObservability,
	commonClient.NewConsulClient,
	commonClient.NewRedisClient,
	commonClient.NewRedisLock,
	commonClient.NewNatsClient,
	commonClient.NewLarkWebhookClient,
	commonClient.NewDeadLetterAlertClient,
	rpc.ProvideUserClient,
	client.ProvideTx,
	repo.NewArticleRepo,
	repo.NewCommentRepo,
	repo.NewCommentActionRecordRepo,
	repo.NewArticlePostscriptRepo,
	repo.NewArticleActionRecordRepo,
	repo.NewOutboxEventRepo,
	repo.NewContentModerationRecordRepo,
	repo.NewDomainRepo,
	repo.NewTagRepo,
	repo.NewUserClient,
	repo.NewNatsEventClient,
)

func ProvideRedis(c *config.Bootstrap) *common.Redis {
	return c.Redis
}

func ProvideConsul(c *config.Bootstrap) *common.Consul {
	return c.Consul
}

func ProvideNats(c *config.Bootstrap) *common.Nats {
	return c.Nats
}
