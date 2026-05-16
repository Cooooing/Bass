package service

import (
	"common/pkg/client"
	"common/pkg/util/jwt"
	"common/pkg/util/server"
	"signal/internal/conf"
	"signal/internal/data/ent/gen"

	"github.com/go-kratos/kratos/v2/log"
	"github.com/google/wire"
)

// ServiceProviderSet is service providers.
var ServiceProviderSet = wire.NewSet(
	NewBaseService,
	NewSystemService,
	ProvideServices,

	NewNodeService,
)

type BaseService struct {
	Conf       *conf.Bootstrap
	Log        *log.Helper
	Db         *gen.Client
	Consul     *client.ConsulClient
	Redis      *client.RedisClient
	TokenCache *jwt.TokenCache
}

func NewBaseService(conf *conf.Bootstrap, logger log.Logger, db *gen.Client, consul *client.ConsulClient, redis *client.RedisClient, tokenCache *jwt.TokenCache) *BaseService {
	return &BaseService{
		Conf:       conf,
		Log:        log.NewHelper(logger),
		Db:         db,
		Consul:     consul,
		Redis:      redis,
		TokenCache: tokenCache,
	}
}

func ProvideServices(
	systemService *SystemService,
	nodeService *NodeService,
) []server.Service {
	return []server.Service{
		systemService,
		nodeService,
	}
}
