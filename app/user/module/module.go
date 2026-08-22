package module

import (
	commonclient "common/pkg/client"
	"common/pkg/client/rpc"
	"common/pkg/server"
	"common/proto/gen/common"
	"fmt"
	"user/internal/biz"
	"user/internal/config"
	"user/internal/data"
	userclient "user/internal/data/client"
	"user/internal/data/repo"
	"user/internal/service"

	"github.com/google/wire"
)

// ProviderSet 提供用户模块的依赖项。
var ProviderSet = wire.NewSet(
	provideBootstrap,
	provideCommonServer,
	data.ProvideRedis,
	data.ProvideNats,
	commonclient.NewObservability,
	commonclient.NewRedisClient,
	commonclient.NewNatsClient,
	userclient.NewDataBaseClient,
	userclient.ProvideTx,
	repo.NewAccountRepo,
	repo.NewRelationRepo,
	repo.NewPreferencesRepo,
	repo.NewPrivacySettingRepo,
	repo.NewLocationRepo,
	repo.NewTotpRepo,
	repo.NewCheckinRecordRepo,
	repo.NewCheckinStatRepo,
	repo.NewLoginLogRepo,
	repo.NewBanRecordRepo,
	repo.NewRbacRepo,
	repo.NewOutboxEventRepo,
	repo.NewTotpSecretCache,
	repo.NewAuthCacheRepo,
	repo.NewNotificationRateLimitClient,
	repo.NewIPClient,
	repo.NewDelayedTaskClient,
	repo.NewNatsEventClient,
	biz.BizProviderSet,
	service.ServiceProviderSet,
	newClients,
	newModule,
)

type Config struct {
	bootstrap *config.Bootstrap
}

type Module struct {
	Name     string
	Services []server.Service
	Clients  *Clients
}

type Clients struct {
	User *rpc.UserClient
}

func newModule(config *Config, services []server.Service, clients *Clients) *Module {
	return &Module{Name: config.Server().GetName(), Services: services, Clients: clients}
}

func LoadConfig(bootstrapPath string, path string) (*Config, *common.Bootstrap, func(), error) {
	c, bc, cleanup, err := config.LoadConfig(bootstrapPath, path)
	if err != nil {
		return nil, nil, cleanup, err
	}
	return &Config{bootstrap: c}, bc, cleanup, nil
}

func LoadRequiredConfig(name string, bootstrapPath string, path string) (*Config, func(), error) {
	if bootstrapPath == "" || path == "" {
		return nil, func() {}, fmt.Errorf("%s module config path is required", name)
	}
	c, _, cleanup, err := LoadConfig(bootstrapPath, path)
	return c, cleanup, err
}

func (c *Config) Server() *common.Server { return c.bootstrap.GetServer() }
func (c *Config) Trace() *common.Trace   { return c.bootstrap.GetTrace() }

func provideBootstrap(c *Config) *config.Bootstrap { return c.bootstrap }
func provideCommonServer(c *Config) *common.Server { return c.Server() }

func newClients(services []server.Service) *Clients {
	return &Clients{User: rpc.NewLocalUserClient(services)}
}
