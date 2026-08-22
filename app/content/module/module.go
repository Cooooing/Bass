package module

import (
	commonclient "common/pkg/client"
	"common/pkg/client/rpc"
	"common/pkg/server"
	"common/proto/gen/common"
	"content/internal/biz"
	"content/internal/config"
	"content/internal/data"
	dataclient "content/internal/data/client"
	"content/internal/data/repo"
	"content/internal/service"
	"fmt"

	"github.com/google/wire"
)

// ProviderSet 提供内容模块的依赖项。
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
	repo.NewArticleRepo,
	repo.NewCommentRepo,
	repo.NewCommentActionRecordRepo,
	repo.NewPostscriptRepo,
	repo.NewArticleActionRecordRepo,
	repo.NewArticleViewCacheRepo,
	repo.NewArticleViewRecordRepo,
	repo.NewOutboxEventRepo,
	repo.NewContentModerationRecordRepo,
	repo.NewDomainRepo,
	repo.NewTagRepo,
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
	Content *rpc.ContentClient
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
	return &Clients{Content: rpc.NewLocalContentClient(services)}
}
