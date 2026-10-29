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
	NewLivenessService,
	NewContentArticleService,
	NewContentArticleSearchService,
	NewContentPostscriptService,
	NewContentCommentService,
	NewContentDomainService,
	NewContentTagService,
	NewNotificationService,
	NewContentMoonbreezeService,
	NewHolidayService,
	NewRealtimeService,
	NewIMService,
	NewEmojiService,
	NewAssetService,
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
	livenessService *LivenessService,
	contentArticleService *ContentArticleService,
	contentArticleSearchService *ContentArticleSearchService,
	contentPostscriptService *ContentPostscriptService,
	contentCommentService *ContentCommentService,
	contentDomainService *ContentDomainService,
	contentTagService *ContentTagService,
	notificationService *NotificationService,
	contentMoonbreezeService *ContentMoonbreezeService,
	holidayService *HolidayService,
	realtimeService *RealtimeService,
	imService *IMService,
	emojiService *EmojiService,
	assetService *AssetService,
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
		livenessService,
		contentArticleService,
		contentArticleSearchService,
		contentPostscriptService,
		contentCommentService,
		contentDomainService,
		contentTagService,
		notificationService,
		contentMoonbreezeService,
		holidayService,
		realtimeService,
		imService,
		emojiService,
		assetService,
	}
}
