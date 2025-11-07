package service

import (
	commonClient "common/pkg/client"
	"gateway/internal/conf"

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
	conf *conf.Bootstrap
	log  *log.Helper
	etcd *commonClient.EtcdClient
}

func NewBaseService(conf *conf.Bootstrap, logger *log.Helper, etcd *commonClient.EtcdClient) *BaseService {
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
) []Service {
	return []Service{
		systemService,
	}
}
