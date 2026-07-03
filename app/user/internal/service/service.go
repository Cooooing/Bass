package service

import (
	"common/pkg/server"

	"github.com/google/wire"
)

// ServiceProviderSet 是 service 层依赖集合。
var ServiceProviderSet = wire.NewSet(
	NewSystemService,
	NewAuthService,
	NewAccountService,
	NewPreferencesService,
	NewPrivacySettingService,
	NewLocationService,
	NewRelationService,
	NewTotpService,
	ProvideGrpcServices,
)

func ProvideGrpcServices(
	systemService *SystemService,
	authService *AuthService,
	accountService *AccountService,
	preferencesService *PreferencesService,
	privacySettingService *PrivacySettingService,
	locationService *LocationService,
	relationService *RelationService,
	totpService *TotpService,
) []server.GrpcService {
	return []server.GrpcService{
		systemService,
		authService,
		accountService,
		preferencesService,
		privacySettingService,
		locationService,
		relationService,
		totpService,
	}
}
