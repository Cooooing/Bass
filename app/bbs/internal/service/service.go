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
	NewRelationService,
	NewPreferencesService,
	NewPrivacySettingService,
	NewLocationService,
	NewTotpService,
	NewContentArticleService,
	NewContentPostscriptService,
	NewContentCommentService,
	NewContentDomainService,
	NewContentTagService,
	NewNotificationService,
	ProvideGrpcServices,
	ProvideHttpServices,
)

func ProvideGrpcServices(
	commonSystemService *CommonSystemService,
) []server.GrpcService {
	return []server.GrpcService{
		commonSystemService,
	}
}

func ProvideHttpServices(
	commonSystemService *CommonSystemService,
	authService *AuthService,
	accountService *AccountService,
	relationService *RelationService,
	preferencesService *PreferencesService,
	privacySettingService *PrivacySettingService,
	locationService *LocationService,
	totpService *TotpService,
	contentArticleService *ContentArticleService,
	contentPostscriptService *ContentPostscriptService,
	contentCommentService *ContentCommentService,
	contentDomainService *ContentDomainService,
	contentTagService *ContentTagService,
	notificationService *NotificationService,
) []server.HttpService {
	return []server.HttpService{
		commonSystemService,
		authService,
		accountService,
		relationService,
		preferencesService,
		privacySettingService,
		locationService,
		totpService,
		contentArticleService,
		contentPostscriptService,
		contentCommentService,
		contentDomainService,
		contentTagService,
		notificationService,
	}
}
