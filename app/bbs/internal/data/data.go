package data

import (
	"bbs/internal/conf"
	"common/api/gen/common"
	commonClient "common/pkg/client"

	"github.com/google/wire"
)

// DataProviderSet is data providers.
var DataProviderSet = wire.NewSet(
	ProvideConsul,
	commonClient.NewConsulClient,
)

func ProvideConsul(c *conf.Bootstrap) *common.Consul {
	return c.Data.Consul
}
