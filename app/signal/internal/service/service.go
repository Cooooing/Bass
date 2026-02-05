package service

import (
	"common/pkg/client"
	"common/pkg/util"
	"signal/internal/conf"
	"signal/internal/data/ent/gen"

	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-kratos/kratos/v2/transport/grpc"
	"github.com/go-kratos/kratos/v2/transport/http"
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
	conf       *conf.Bootstrap
	log        *log.Helper
	db         *gen.Client
	consul     *client.ConsulClient
	redis      *client.RedisClient
	rabbitmq   *client.RabbitMQClient
	tokenCache *util.TokenCache
}

func NewBaseService(conf *conf.Bootstrap, logger *log.Helper, db *gen.Client, consul *client.ConsulClient, redis *client.RedisClient, rabbitmq *client.RabbitMQClient, tokenCache *util.TokenCache) *BaseService {
	return &BaseService{
		conf:       conf,
		log:        logger,
		db:         db,
		consul:     consul,
		redis:      redis,
		rabbitmq:   rabbitmq,
		tokenCache: tokenCache,
	}
}

// Service 接口，每个 service 实现它
type Service interface {
	RegisterGrpc(gs *grpc.Server)
	RegisterHttp(hs *http.Server)
}

func ProvideServices(
	systemService *SystemService,
	nodeService *NodeService,
) []Service {
	return []Service{
		systemService,
		nodeService,
	}
}
