package service

import (
	"common/pkg/util/server"

	"github.com/google/wire"
)

// ServiceProviderSet 是 service 层依赖集合。
var ServiceProviderSet = wire.NewSet(
	NewSystemService,
	NewAuthService,
	NewAccountService,
	NewRelationService,
	NewPreferencesService,
	NewPrivacySettingService,
	NewLocationService,
	NewTfaService,
	NewContentArticleService,
	NewContentPostscriptService,
	NewContentCommentService,
	NewContentDomainService,
	NewContentTagService,
	NewNotificationService,
	ProvideServices,
)

func ProvideServices(
	systemService *SystemService,
	authService *AuthService,
	accountService *AccountService,
	relationService *RelationService,
	preferencesService *PreferencesService,
	privacySettingService *PrivacySettingService,
	locationService *LocationService,
	tfaService *TfaService,
	contentArticleService *ContentArticleService,
	contentPostscriptService *ContentPostscriptService,
	contentCommentService *ContentCommentService,
	contentDomainService *ContentDomainService,
	contentTagService *ContentTagService,
	notificationService *NotificationService,
) []server.HttpService {
	return []server.HttpService{
		systemService,
		authService,
		accountService,
		relationService,
		preferencesService,
		privacySettingService,
		locationService,
		tfaService,
		contentArticleService,
		contentPostscriptService,
		contentCommentService,
		contentDomainService,
		contentTagService,
		notificationService,
	}
}
