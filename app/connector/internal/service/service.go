package service

import (
	"common/pkg/util/server"
	"connector/internal/conf"

	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-kratos/kratos/v2/transport/grpc"
	"github.com/go-kratos/kratos/v2/transport/http"
	"github.com/google/wire"
)

// ServiceProviderSet is service providers.
var ServiceProviderSet = wire.NewSet(
	NewBaseService,
	NewSystemService,
	NewCallbackService,
	ProvideServices,
)

type BaseService struct {
	Conf *conf.Bootstrap
	Log  *log.Helper
}

func NewBaseService(conf *conf.Bootstrap, logger log.Logger) *BaseService {
	return &BaseService{
		Conf: conf,
		Log:  log.NewHelper(logger),
	}
}

// Service 接口，每个 service 实现它
type Service interface {
	RegisterGrpc(gs *grpc.Server)
	RegisterHttp(gs *http.Server)
}

func ProvideServices(
	systemService *SystemService,
	callbackService *CallbackService,
) []server.Service {
	return []server.Service{
		systemService,
		callbackService,
	}
}
