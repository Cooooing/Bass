package module

import (
	"bbs/internal/biz"
	"bbs/internal/config"
	"bbs/internal/data"
	bbsserver "bbs/internal/server"
	"bbs/internal/service"
	"common/pkg/client"
	"common/pkg/client/rpc"
	"common/pkg/server"
	"common/proto/gen/common"
	userv1 "common/proto/gen/user/v1"
	"fmt"
	"log/slog"

	"github.com/go-kratos/kratos/v3/middleware"
	"github.com/go-kratos/kratos/v3/transport"
	"github.com/google/wire"
)

// ProviderSet 提供 BBS 模块的依赖项。
var ProviderSet = wire.NewSet(
	data.NewAuthClient,
	data.NewAccountClient,
	data.NewRelationClient,
	data.NewPreferencesClient,
	data.NewPrivacySettingClient,
	data.NewLocationClient,
	data.NewOtpClient,
	data.NewCheckinClient,
	data.NewAssetClient,
	data.NewEconomyClient,
	data.NewNotificationClient,
	data.NewContentArticleClient,
	data.NewContentPostscriptClient,
	data.NewContentCommentClient,
	data.NewContentDomainClient,
	data.NewContentTagClient,
	provideUserAuthClient,
	biz.BizProviderSet,
	service.ServiceProviderSet,
	newModule,
)

type Config struct{ bootstrap *config.Bootstrap }

type Module struct {
	Name           string
	Services       []server.Service
	UserAuthClient userv1.AuthServiceClient
	config         *Config
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

func newModule(config *Config, services []server.Service, userAuthClient userv1.AuthServiceClient) *Module {
	return &Module{Name: config.Server().GetName(), Services: services, UserAuthClient: userAuthClient, config: config}
}

func (m *Module) Servers(commonServer *common.Server, logger *slog.Logger) []transport.Server {
	return newServers(m, commonServer, logger)
}

func (m *Module) HTTPMiddlewares() []middleware.Middleware {
	return bbsserver.NewHTTPAuthMiddlewares(m.UserAuthClient)
}

func provideUserAuthClient(userClient *rpc.UserClient) userv1.AuthServiceClient {
	return userClient.Auth
}

func newServers(m *Module, commonServer *common.Server, logger *slog.Logger) []transport.Server {
	observer := client.NewObservability(logger, commonServer)
	grpcServer := bbsserver.NewGRPCServer(m.config.bootstrap, logger, observer, m.Services)
	httpServer := bbsserver.NewHTTPServer(m.config.bootstrap, logger, observer, m.Services, m.UserAuthClient)
	return bbsserver.ProvideServers(grpcServer, httpServer)
}
