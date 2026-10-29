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
	NewOtpService,
	NewOutboxService,
	NewLivenessService,
	NewTCCCheckinService,
	NewTCCLivenessRewardService,
)

func ProvideServices(commonSystemService *CommonSystemService, authService *AuthService, rbacService *RbacService, accountService *AccountService, preferencesService *PreferencesService, privacySettingService *PrivacySettingService, locationService *LocationService, relationService *RelationService, otpService *OtpService, outboxService *OutboxService, livenessService *LivenessService, tccCheckinService *TCCCheckinService, tccLivenessRewardService *TCCLivenessRewardService) []server.Service {
	return []server.Service{
		commonSystemService,
		authService,
		rbacService,
		accountService,
		preferencesService,
		privacySettingService,
		locationService,
		relationService,
		otpService,
		outboxService,
		livenessService,
		tccCheckinService,
		tccLivenessRewardService,
	}
}
