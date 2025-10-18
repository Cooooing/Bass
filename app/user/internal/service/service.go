package service

import (
	"user/internal/conf"
	"user/internal/data/client"

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
	ProvideServices,
)

type BaseService struct {
	conf *conf.Bootstrap
	log  *log.Helper
	etcd *client.EtcdClient
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
	authenticationService *AuthenticationService,
	systemService *SystemService,
) []Service {
	return []Service{
		authenticationService,
		systemService,
	}
}
