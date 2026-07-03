package data

import (
	commonClient "common/pkg/client"
	"common/proto/gen/common"
	"integration/internal/conf"

	"github.com/google/wire"
)

var DataProviderSet = wire.NewSet(
	ProvideConsul,
	commonClient.NewObservability,
	commonClient.NewConsulClient,
)

func ProvideConsul(c *conf.Bootstrap) *common.Consul { return c.Data.Consul }
