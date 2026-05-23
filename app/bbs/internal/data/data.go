package data

import (
	"bbs/internal/conf"
	"common/api/gen/common"
	commonClient "common/pkg/client"
	"common/pkg/client/rpc"

	"github.com/google/wire"
)

// DataProviderSet 是 data 层依赖集合。
var DataProviderSet = wire.NewSet(
	ProvideConsul,
	commonClient.NewConsulClient,
	rpc.ProvideUserClient,
)

func ProvideConsul(c *conf.Bootstrap) *common.Consul {
	return c.Data.Consul
}
