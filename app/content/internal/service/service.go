package service

import (
	"common/pkg/client"
	"content/internal/conf"
	"content/internal/data/ent/gen"

	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-kratos/kratos/v2/transport/grpc"
	"github.com/go-kratos/kratos/v2/transport/http"
	"github.com/google/wire"
)

// ServiceProviderSet is service providers.
var ServiceProviderSet = wire.NewSet(
	NewBaseService,

	NewSystemService,
	NewArticleService,
	NewDomainService,

	ProvideServices,
)

type BaseService struct {
	conf *conf.Bootstrap
	log  *log.Helper
	etcd *client.EtcdClient
	db   *gen.Client
}

func NewBaseService(conf *conf.Bootstrap, logger *log.Helper, etcd *client.EtcdClient) *BaseService {
	return &BaseService{
		conf: conf,
		log:  logger,
		etcd: etcd,
	}
}

// Service 接口，每个 service 实现它
type Service interface {
	RegisterGrpc(gs *grpc.Server)
	RegisterHttp(hs *http.Server)
}

func ProvideServices(
	systemService *SystemService,
	articleService *ArticleService,
	domainService *DomainService,
) []Service {
	return []Service{
		systemService,
		articleService,
		domainService,
	}
}
