package module

import (
	commonclient "common/pkg/client"
	"common/pkg/client/rpc"
	"common/pkg/server"
	"common/proto/gen/common"
	"fmt"
	"notify/internal/biz/usecase"
	"notify/internal/config"
	"notify/internal/data"
	dataclient "notify/internal/data/client"
	"notify/internal/data/repo"
	"notify/internal/service"

	"github.com/google/wire"
)

// ProviderSet 提供通知模块的依赖项。
var ProviderSet = wire.NewSet(
	provideBootstrap,
	dataclient.NewDataBaseClient,
	dataclient.ProvideTx,
	dataProviders,
	service.ServiceProviderSet,
	newClients,
	newModule,
)

var dataProviders = wire.NewSet(
	commonclient.NewRedisClient,
	data.ProvideRedis,
	repo.NewNotificationStationMessageRepo,
	repo.NewNotificationRuleRepo,
	repo.NewNotificationStationTemplateRepo,
	repo.NewNotificationEmailTemplateRepo,
	repo.NewNotificationTencentSMSTemplateRepo,
	repo.NewNotificationLarkWebhookTemplateRepo,
	repo.NewNotificationRateLimitCache,
	usecase.NewStationMessageUsecase,
	usecase.NewRateLimitUsecase,
	usecase.NewTemplateUsecase,
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
	Notify *rpc.NotifyClient
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

func newClients(services []server.Service) *Clients {
	return &Clients{Notify: rpc.NewLocalNotifyClient(services)}
}
