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

	NewChatSessionService,
	NewChatMessageService,
)

type BaseService struct {
	Conf   *conf.Bootstrap
	Log    *log.Helper
	Db     *gen.Client
	Consul *client.ConsulClient
	Redis  *client.RedisClient
}

func NewBaseService(conf *conf.Bootstrap, logger log.Logger, db *gen.Client, consul *client.ConsulClient, redis *client.RedisClient) *BaseService {
	return &BaseService{
		Conf:   conf,
		Log:    log.NewHelper(logger),
		Db:     db,
		Consul: consul,
		Redis:  redis,
	}
}

// Service 接口，每个 service 实现它
type Service interface {
	RegisterGrpc(gs *grpc.Server)
	RegisterHttp(gs *http.Server)
}

func ProvideServices(
	systemService *SystemService,
	chatSessionService *ChatSessionService,
	chatMessageService *ChatMessageService,
) []Service {
	return []Service{
		systemService,
		chatSessionService,
		chatMessageService,
	}
}
