package service

import (
	"common/pkg/util/server"

	"github.com/google/wire"
)

// ServiceProviderSet is service providers.
var ServiceProviderSet = wire.NewSet(
	NewVerifyService,
	NewSystemService,
	NewAuthenticationService,
	NewUserService,
	NewUserRelationService,
	NewTwoFactorAuthenticationService,
	ProvideServices,
)

func ProvideServices(
	systemService *SystemService,
	authenticationService *AuthenticationService,
	userService *UserService,
	userRelationService *UserRelationService,
	twoFactorAuthenticationService *TwoFactorAuthenticationService,
) []server.GrpcService {
	return []server.GrpcService{
		systemService,
		authenticationService,
		userService,
		userRelationService,
		twoFactorAuthenticationService,
	}
}
