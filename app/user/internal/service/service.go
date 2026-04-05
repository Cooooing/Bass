package service

import (
	"user/internal/conf"
	"user/internal/data/ent/gen"

	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-kratos/kratos/v2/transport/grpc"
	"github.com/google/wire"
)

// ServiceProviderSet is service providers.
var ServiceProviderSet = wire.NewSet(
	NewBaseService,
	NewVerifyService,
	NewSystemService,
	NewAuthenticationService,
	NewUserService,
	NewUserRelationService,
	NewTwoFactorAuthenticationService,
	ProvideServices,
)

type BaseService struct {
	Conf *conf.Bootstrap
	Log  *log.Helper
	Db   *gen.Client
}

func NewBaseService(conf *conf.Bootstrap, logger *log.Helper, db *gen.Client) *BaseService {
	return &BaseService{
		Conf: conf,
		Log:  logger,
		Db:   db,
	}
}

// Service 接口，每个 service 实现它
type Service interface {
	RegisterGrpc(gs *grpc.Server)
}

func ProvideServices(
	systemService *SystemService,
	authenticationService *AuthenticationService,
	userService *UserService,
	userRelationService *UserRelationService,
	twoFactorAuthenticationService *TwoFactorAuthenticationService,
) []Service {
	return []Service{
		systemService,
		authenticationService,
		userService,
		userRelationService,
		twoFactorAuthenticationService,
	}
}
