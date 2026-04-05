package service

import (
	"common/pkg/client"
	"common/pkg/util/jwt"
	"content/internal/conf"
	"content/internal/data/ent/gen"

	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-kratos/kratos/v2/transport/grpc"
	"github.com/google/wire"
)

// ServiceProviderSet is service providers.
var ServiceProviderSet = wire.NewSet(
	NewBaseService,

	NewSystemService,
	NewArticleService,
	NewCommentService,
	NewDomainService,
	NewTagService,

	ProvideServices,
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
}

func ProvideServices(
	systemService *SystemService,
	articleService *ArticleService,
	domainService *DomainService,
	commentService *CommentService,
	tagService *TagService,
) []Service {
	return []Service{
		systemService,
		articleService,
		domainService,
		commentService,
		tagService,
	}
}
