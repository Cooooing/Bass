package service

import (
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
	NewWebsocketService, // websocket 接口，手动注册
	ProvideServices,
)

type BaseService struct {
	Conf *conf.Bootstrap
	Log  *log.Helper
}

func NewBaseService(conf *conf.Bootstrap, logger *log.Helper) *BaseService {
	return &BaseService{
		Conf: conf,
		Log:  logger,
	}
}

// Service 接口，每个 service 实现它
type Service interface {
	RegisterGrpc(gs *grpc.Server)
	RegisterHttp(hs *http.Server)
}

func ProvideServices(
	systemService *SystemService,
	callbackService *CallbackService,
) []Service {
	return []Service{
		systemService,
		callbackService,
	}
}
