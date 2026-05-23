package service

import (
	"common/pkg/util/server"

	"github.com/google/wire"
)

// ServiceProviderSet 是 service 层依赖集合。
var ServiceProviderSet = wire.NewSet(
	NewVerifyService,
	NewSystemService,
	NewAuthService,
	NewAccountService,
	NewPreferencesService,
	NewPrivacySettingService,
	NewLocationService,
	NewRelationService,
	NewTfaService,
	ProvideServices,
)

func ProvideServices(
	systemService *SystemService,
	authService *AuthService,
	accountService *AccountService,
	preferencesService *PreferencesService,
	privacySettingService *PrivacySettingService,
	locationService *LocationService,
	relationService *RelationService,
	tfaService *TfaService,
) []server.GrpcService {
	return []server.GrpcService{
		systemService,
		authService,
		accountService,
		preferencesService,
		privacySettingService,
		locationService,
		relationService,
		tfaService,
	}
}
