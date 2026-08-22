package module

import (
	commonclient "common/pkg/client"
	"common/pkg/client/rpc"
	"common/pkg/server"
	"common/proto/gen/common"
	"fmt"
	"scheduler/internal/biz"
	"scheduler/internal/config"
	"scheduler/internal/data"
	dataclient "scheduler/internal/data/client"
	"scheduler/internal/data/repo"
	"scheduler/internal/service"

	"github.com/google/wire"
)

// ProviderSet 提供调度模块的依赖项。
var ProviderSet = wire.NewSet(
	provideBootstrap,
	provideCommonServer,
	commonclient.NewObservability,
	data.ProvideRedis,
	data.ProvideNats,
	commonclient.NewRedisClient,
	commonclient.NewNatsClient,
	dataclient.NewDataBaseClient,
	dataclient.ProvideTx,
	repo.NewScheduledTaskCacheRepo,
	repo.NewDelayedTaskCacheRepo,
	repo.NewScheduledTaskRepo,
	repo.NewScheduledTaskVersionRepo,
	repo.NewScheduledTaskExecutionRecordRepo,
	repo.NewScheduledTaskScheduleNatsRepo,
	repo.NewDelayedTaskScheduleNatsRepo,
	repo.NewDelayedTaskRepo,
	repo.NewDelayedTaskVersionRepo,
	repo.NewDelayedTaskExecutionRecordRepo,
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
	Scheduler *rpc.SchedulerClient
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
	return &Clients{Scheduler: rpc.NewLocalSchedulerClient(services)}
}
