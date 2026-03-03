package service

import (
	"common/pkg/client"
	"common/pkg/util/jwt"
	"infra/internal/conf"
	"infra/internal/data/ent/gen"

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
	NewOssService,
	NewEmailService,
	NewSmsService,
)

type BaseService struct {
	Conf       *conf.Bootstrap
	Log        *log.Helper
	Db         *gen.Client
	Consul     *client.ConsulClient
	Redis      *client.RedisClient
	Rabbitmq   *client.RabbitMQClient
	TokenCache *jwt.TokenCache
}

func NewBaseService(conf *conf.Bootstrap, logger *log.Helper, db *gen.Client, consul *client.ConsulClient, redis *client.RedisClient, rabbitmq *client.RabbitMQClient, tokenCache *jwt.TokenCache) *BaseService {
	return &BaseService{
		Conf:       conf,
		Log:        logger,
		Db:         db,
		Consul:     consul,
		Redis:      redis,
		Rabbitmq:   rabbitmq,
		TokenCache: tokenCache,
	}
}

// Service 接口，每个 service 实现它
type Service interface {
	RegisterGrpc(gs *grpc.Server)
	RegisterHttp(hs *http.Server)
}

func ProvideServices(
	systemService *SystemService,
	ossService *OssService,
) []Service {
	return []Service{
		systemService,
		ossService,
	}
}
