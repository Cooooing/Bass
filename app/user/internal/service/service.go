package service

import (
	"common/pkg/util/server"
	"user/internal/conf"

	"github.com/go-kratos/kratos/v2/log"
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
}

func NewBaseService(conf *conf.Bootstrap, logger log.Logger) *BaseService {
	return &BaseService{
		Conf: conf,
		Log:  log.NewHelper(logger),
	}
}

func ProvideServices(
	systemService *SystemService,
	authenticationService *AuthenticationService,
	userService *UserService,
	userRelationService *UserRelationService,
	twoFactorAuthenticationService *TwoFactorAuthenticationService,
) []server.Service {
	return []server.Service{
		systemService,
		authenticationService,
		userService,
		userRelationService,
		twoFactorAuthenticationService,
	}
}
