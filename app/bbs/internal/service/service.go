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
	NewRelationService,
	NewPreferencesService,
	NewPrivacySettingService,
	NewLocationService,
	NewOtpService,
	NewContentArticleService,
	NewContentPostscriptService,
	NewContentCommentService,
	NewContentDomainService,
	NewContentTagService,
	NewNotificationService,
)

func ProvideServices(commonSystemService *CommonSystemService, authService *AuthService, accountService *AccountService, relationService *RelationService, preferencesService *PreferencesService, privacySettingService *PrivacySettingService, locationService *LocationService, otpService *OtpService, contentArticleService *ContentArticleService, contentPostscriptService *ContentPostscriptService, contentCommentService *ContentCommentService, contentDomainService *ContentDomainService, contentTagService *ContentTagService, notificationService *NotificationService) []server.Service {
	return []server.Service{
		commonSystemService,
		authService,
		accountService,
		relationService,
		preferencesService,
		privacySettingService,
		locationService,
		otpService,
		contentArticleService,
		contentPostscriptService,
		contentCommentService,
		contentDomainService,
		contentTagService,
		notificationService,
	}
}
