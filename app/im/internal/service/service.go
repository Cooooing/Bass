package service

import (
	"common/pkg/client"
	"im/internal/conf"
	"im/internal/data/ent/gen"

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
)

type BaseService struct {
	conf     *conf.Bootstrap
	log      *log.Helper
	db       *gen.Client
	etcd     *client.EtcdClient
	redis    *client.RedisClient
	rabbitmq *client.RabbitMQClient
}

func NewBaseService(conf *conf.Bootstrap, logger *log.Helper, db *gen.Client, etcd *client.EtcdClient, redis *client.RedisClient, rabbitmq *client.RabbitMQClient) *BaseService {
	return &BaseService{
		conf:     conf,
		log:      logger,
		db:       db,
		etcd:     etcd,
		redis:    redis,
		rabbitmq: rabbitmq,
	}
}

// Service 接口，每个 service 实现它
type Service interface {
	RegisterGrpc(gs *grpc.Server)
	RegisterHttp(hs *http.Server)
}

func ProvideServices(
	systemService *SystemService,
) []Service {
	return []Service{
		systemService,
	}
}
