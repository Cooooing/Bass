package service

import (
	"common/pkg/server"

	"github.com/google/wire"
)

var ServiceProviderSet = wire.NewSet(
	ProvideServices,
	NewCommonSystemService,
	NewAuthService,
	NewRbacService,
	NewAccountService,
	NewPreferencesService,
	NewPrivacySettingService,
	NewLocationService,
	NewRelationService,
	NewTotpService,
)

func ProvideServices(
	commonSystemService *CommonSystemService,
	authService *AuthService,
	rbacService *RbacService,
	accountService *AccountService,
	preferencesService *PreferencesService,
	privacySettingService *PrivacySettingService,
	locationService *LocationService,
	relationService *RelationService,
	totpService *TotpService,
) []server.Service {
	return []server.Service{commonSystemService, authService, rbacService, accountService, preferencesService, privacySettingService, locationService, relationService, totpService}
}
