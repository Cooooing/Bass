package data

import (
	commonclient "common/pkg/client"
	"common/pkg/client/rpc"
	"common/proto/gen/common/v1"

	"badge/internal/config"
	"badge/internal/data/client"
	"badge/internal/data/gen"
	"badge/internal/data/repo"

	"github.com/google/wire"
)

var DataProviderSet = wire.NewSet(ModuleProviderSet, commonclient.NewObservability, ProvideConsul, commonclient.NewConsulClient, rpc.ProvidePlatformClient)
var ModuleProviderSet = wire.NewSet(commonclient.NewSQLClient, ProvideDatabase, client.NewEntClient, ProvideGenClient, repo.NewBadgeRepo, repo.NewUserBadgeRepo, repo.NewBadgeResourceRepo)

func ProvideDatabase(c *config.Bootstrap) *common.Database { return c.GetDatabase() }
func ProvideConsul(c *config.Bootstrap) *common.Consul     { return c.GetConsul() }
func ProvideGenClient(c *client.EntClient) *gen.Client     { return c.Client }
