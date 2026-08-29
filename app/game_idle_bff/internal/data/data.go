package data

import (
	commonClient "common/pkg/client"
	"common/pkg/client/rpc"
	"common/proto/gen/common"
	userv1 "common/proto/gen/user/v1"
	"game_idle_bff/internal/config"

	"github.com/google/wire"
)

// DataProviderSet 提供微服务模式的数据层依赖。
var DataProviderSet = wire.NewSet(
	ModuleProviderSet,
	ProvideConsul,
	ProvideNats,
	ProvideRedis,
	commonClient.NewConsulClient,
	commonClient.NewRedisClient,
	commonClient.NewNatsClient,
	rpc.ProvideUserClient,
	rpc.ProvideGameIdleClient,
)

// ModuleProviderSet 提供不依赖服务发现的模块数据层依赖。
var ModuleProviderSet = wire.NewSet(
	ProvideUserAuthClient,
	NewAuthRepo,
	NewCharacterRepo,
	NewCharacterAbilityRepo,
	NewBackpackRepo,
	NewActionQueueRepo,
	NewChatRepo,
	NewRegionRepo,
	NewActionRepo,
	NewItemRepo,
	NewWebSocketEventRepo,
	NewWebSocketTicketRepo,
	NewConfigVersionRepo,
)

func ProvideNats(c *config.Bootstrap) *common.Nats {
	return c.Nats
}

func ProvideConsul(c *config.Bootstrap) *common.Consul {
	return c.Consul
}

func ProvideRedis(c *config.Bootstrap) *common.Redis {
	return c.Redis
}

func ProvideUserAuthClient(userClient *rpc.UserClient) userv1.AuthServiceClient {
	return userClient.Auth
}
