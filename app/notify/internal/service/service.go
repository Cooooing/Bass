package service

import (
	commonClient "common/pkg/client"
	"notify/internal/conf"
	"notify/internal/data/ent/gen"

	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-kratos/kratos/v2/transport/grpc"
	"github.com/google/wire"
)

// ServiceProviderSet is service providers.
var ServiceProviderSet = wire.NewSet(
	NewBaseService,
	NewSystemService,
	NewNotificationMetaService,
	NewNotificationRecordService,
	ProvideServices,
)

type BaseService struct {
	Conf   *conf.Bootstrap
	Log    *log.Helper
	Consul *commonClient.ConsulClient
	Db     *gen.Client
}

func NewBaseService(conf *conf.Bootstrap, logger *log.Helper, consul *commonClient.ConsulClient, db *gen.Client) *BaseService {
	return &BaseService{
		Conf:   conf,
		Log:    logger,
		Consul: consul,
		Db:     db,
	}
}

// Service 接口，每个 service 实现它
type Service interface {
	RegisterGrpc(gs *grpc.Server)
}

func ProvideServices(
	systemService *SystemService,
	notificationMetaService *NotificationMetaService,
	notificationRecordService *NotificationRecordService,
) []Service {
	return []Service{
		systemService,
		notificationMetaService,
		notificationRecordService,
	}
}
