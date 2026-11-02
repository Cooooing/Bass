package data

import (
	"badge/internal/config"
	"badge/internal/data/client"
	"badge/internal/data/gen"
	"badge/internal/data/repo"
	commonclient "common/pkg/client"
	"common/pkg/client/rpc"
	"common/proto/gen/common"

	"github.com/google/wire"
)

var DataProviderSet = wire.NewSet(ModuleProviderSet, commonclient.NewObservability, ProvideConsul, commonclient.NewConsulClient, rpc.ProvidePlatformClient)
var ModuleProviderSet = wire.NewSet(commonclient.NewSQLClient, ProvideDatabase, client.NewEntClient, ProvideGenClient, repo.NewBadgeRepo, repo.NewUserBadgeRepo, repo.NewBadgeResourceRepo)

func ProvideDatabase(c *config.Bootstrap) *common.Database { return c.Database }
func ProvideConsul(c *config.Bootstrap) *common.Consul     { return c.Consul }
func ProvideGenClient(c *client.EntClient) *gen.Client     { return c.Client }
