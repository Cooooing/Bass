package service

import (
	"common/pkg/util/server"
	"gateway/internal/conf"

	"github.com/go-kratos/kratos/v2/log"
	"github.com/google/wire"
)

// ServiceProviderSet is service providers.
var ServiceProviderSet = wire.NewSet(
	NewBaseService,
	NewSystemService,
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

func ProvideServices(
	systemService *SystemService,
) []server.Service {
	return []server.Service{
		systemService,
	}
}
