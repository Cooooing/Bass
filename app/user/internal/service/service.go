package service

import (
	"common/pkg/server"

	"github.com/google/wire"
)

// ServiceProviderSet 是 service 层依赖集合。
var ServiceProviderSet = wire.NewSet(
	NewCommonSystemService,
	NewAuthService,
	NewAccountService,
	NewPreferencesService,
	NewPrivacySettingService,
	NewLocationService,
	NewRelationService,
	NewTotpService,
	ProvideGrpcServices,
	ProvideHttpServices,
)

func ProvideGrpcServices(
	commonSystemService *CommonSystemService,
	authService *AuthService,
	accountService *AccountService,
	preferencesService *PreferencesService,
	privacySettingService *PrivacySettingService,
	locationService *LocationService,
	relationService *RelationService,
	totpService *TotpService,
) []server.GrpcService {
	return []server.GrpcService{
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

func ProvideHttpServices(
	commonSystemService *CommonSystemService,
) []server.HttpService {
	return []server.HttpService{
		commonSystemService,
	}
}
