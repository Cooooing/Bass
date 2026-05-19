package service

import (
	"common/pkg/util/server"

	"github.com/google/wire"
)

// ServiceProviderSet is service providers.
var ServiceProviderSet = wire.NewSet(
	NewVerifyService,
	NewSystemService,
	NewAuthService,
	NewUserService,
	NewUserRelationService,
	NewTwoFactorAuthService,
	ProvideServices,
)

func ProvideServices(
	systemService *SystemService,
	authenticationService *AuthService,
	userService *UserService,
	userRelationService *UserRelationService,
	twoFactorAuthService *TwoFactorAuthService,
) []server.GrpcService {
	return []server.GrpcService{
		systemService,
		authenticationService,
		userService,
		userRelationService,
		twoFactorAuthService,
	}
}
