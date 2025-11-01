package service

import (
	"common/pkg/client"
	"common/pkg/util"
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
	NewCommentService,
	NewDomainService,

	ProvideServices,
)

type BaseService struct {
	conf      *conf.Bootstrap
	log       *log.Helper
	etcd      *client.EtcdClient
	db        *gen.Client
	tokenRepo *util.TokenRepo
}

func NewBaseService(conf *conf.Bootstrap, logger *log.Helper, etcd *client.EtcdClient, db *gen.Client, tokenRepo *util.TokenRepo) *BaseService {
	return &BaseService{
		conf:      conf,
		log:       logger,
		etcd:      etcd,
		db:        db,
		tokenRepo: tokenRepo,
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
	commentService *CommentService,
) []Service {
	return []Service{
		systemService,
		articleService,
		domainService,
		commentService,
	}
}
