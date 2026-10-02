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
	NewCheckinService,
	NewContentArticleService,
	NewContentArticleSearchService,
	NewContentPostscriptService,
	NewContentCommentService,
	NewContentDomainService,
	NewContentTagService,
	NewNotificationService,
	NewContentMoonbreezeService,
)

func ProvideServices(
	commonSystemService *CommonSystemService,
	authService *AuthService,
	accountService *AccountService,
	relationService *RelationService,
	preferencesService *PreferencesService,
	privacySettingService *PrivacySettingService,
	locationService *LocationService,
	otpService *OtpService,
	checkinService *CheckinService,
	contentArticleService *ContentArticleService,
	contentArticleSearchService *ContentArticleSearchService,
	contentPostscriptService *ContentPostscriptService,
	contentCommentService *ContentCommentService,
	contentDomainService *ContentDomainService,
	contentTagService *ContentTagService,
	notificationService *NotificationService,
	contentMoonbreezeService *ContentMoonbreezeService,
) []server.Service {
	return []server.Service{
		commonSystemService,
		authService,
		accountService,
		relationService,
		preferencesService,
		privacySettingService,
		locationService,
		otpService,
		checkinService,
		contentArticleService,
		contentArticleSearchService,
		contentPostscriptService,
		contentCommentService,
		contentDomainService,
		contentTagService,
		notificationService,
		contentMoonbreezeService,
	}
}
