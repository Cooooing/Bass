package data

import (
	"common/pkg/auth"
	commonClient "common/pkg/client"
	"common/pkg/client/rpc"
	"common/proto/gen/common"
	"content/internal/conf"
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
	auth.NewTokenCache,
)

func ProvideRedis(c *conf.Bootstrap) *common.Redis   { return c.Redis }
func ProvideConsul(c *conf.Bootstrap) *common.Consul { return c.Consul }
func ProvideNats(c *conf.Bootstrap) *common.Nats     { return c.Nats }
