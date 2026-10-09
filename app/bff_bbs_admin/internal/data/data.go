package data

import (
	commonClient "common/pkg/client"
	"common/proto/gen/common"

	"bff_bbs_admin/internal/config"

	"github.com/google/wire"
)

// DataProviderSet 提供独立服务运行所需的基础设施依赖。
var DataProviderSet = wire.NewSet(
	ModuleProviderSet,
	commonClient.NewObservability,
	ProvideConsul,
	commonClient.NewConsulClient,
)

// ModuleProviderSet 预留管理端模块的数据层依赖。
var ModuleProviderSet = wire.NewSet()

func ProvideConsul(c *config.Bootstrap) *common.Consul {
	return c.Consul
}
