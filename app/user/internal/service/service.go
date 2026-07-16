package service

import (
	"common/pkg/server"

	"github.com/google/wire"
)

var ServiceProviderSet = wire.NewSet(
	ProvideServices,
	NewCommonSystemService,
	NewAuthService,
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
	accountService *AccountService,
	preferencesService *PreferencesService,
	privacySettingService *PrivacySettingService,
	locationService *LocationService,
	relationService *RelationService,
	totpService *TotpService,
) []server.Service {
	return []server.Service{
		commonSystemService,
		authService,
		accountService,
		preferencesService,
		privacySettingService,
		locationService,
		relationService,
		totpService,
	}
}
