package service

import (
	commonClient "common/pkg/client"
	"user/internal/conf"
	"user/internal/data/ent/gen"

	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-kratos/kratos/v2/transport/grpc"
	"github.com/go-kratos/kratos/v2/transport/http"
	"github.com/google/wire"
)

// ServiceProviderSet is service providers.
var ServiceProviderSet = wire.NewSet(
	NewBaseService,
	NewSystemService,
	NewAuthenticationService,
	NewUserService,
	ProvideServices,
)

type BaseService struct {
	conf *conf.Bootstrap
	log  *log.Helper
	etcd *commonClient.EtcdClient
	db   *gen.Client
}

func NewBaseService(conf *conf.Bootstrap, logger *log.Helper, etcd *commonClient.EtcdClient, db *gen.Client) *BaseService {
	return &BaseService{
		conf: conf,
		log:  logger,
		etcd: etcd,
		db:   db,
	}
}

// Service 接口，每个 service 实现它
type Service interface {
	RegisterGrpc(gs *grpc.Server)
	RegisterHttp(hs *http.Server)
}

func ProvideServices(
	systemService *SystemService,
	authenticationService *AuthenticationService,
	userService *UserService,
) []Service {
	return []Service{
		systemService,
		authenticationService,
		userService,
	}
}
