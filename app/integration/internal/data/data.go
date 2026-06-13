package data

import (
	"common/proto/gen/common"
	commonClient "common/pkg/client"
	"integration/internal/conf"

	"github.com/google/wire"
)

// DataProviderSet 是 data 层依赖集合。
var DataProviderSet = wire.NewSet(
	ProvideConsul,
	commonClient.NewConsulClient,
)

func ProvideConsul(c *conf.Bootstrap) *common.Consul {
	return c.Data.Consul
}
