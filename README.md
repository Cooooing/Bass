# @bass/bbs-sdk-fetch@0.1.0

A TypeScript SDK client for the localhost API.

## Usage

First, install the SDK from npm.

```bash
npm install @bass/bbs-sdk-fetch --save
```

Next, try it out.


```ts
import {
  Configuration,
  AccountService,
} from '@bass/bbs-sdk-fetch';
import type { AvatarAccountRequest } from '@bass/bbs-sdk-fetch';

async function example() {
  console.log("🚀 Testing @bass/bbs-sdk-fetch SDK...");
  const api = new AccountService();

  const body = {
    // string | 账号名，必填 (optional)
    name: name_example,
  } satisfies AvatarAccountRequest;

  try {
    const data = await api.avatarAccount(body);
    console.log(data);
  } catch (error) {
    console.error(error);
  }
}

// Run the test
example().catch(console.error);
```


## Documentation

### API Endpoints

All URIs are relative to *http://localhost*

| Class | Method | HTTP request | Description
| ----- | ------ | ------------ | -------------
*AccountService* | [**avatarAccount**](docs/AccountService.md#avataraccount) | **GET** /v1/user/account/avatar | 
*AccountService* | [**completeProfileImageUploadAccount**](docs/AccountService.md#completeprofileimageuploadaccount) | **POST** /v1/user/account/complete-profile-image-upload | 
*AccountService* | [**getAccountEconomy**](docs/AccountService.md#getaccounteconomy) | **POST** /v1/user/account/economy/get | 
*AccountService* | [**getCurrentAccount**](docs/AccountService.md#getcurrentaccount) | **POST** /v1/user/account/get-current | 
*AccountService* | [**getProfile**](docs/AccountService.md#getprofile) | **POST** /v1/user/account/get-profile | 
*AccountService* | [**listFollowers**](docs/AccountService.md#listfollowers) | **POST** /v1/user/account/list-followers | 
*AccountService* | [**listFollowing**](docs/AccountService.md#listfollowing) | **POST** /v1/user/account/list-following | 
*AccountService* | [**listPointRecords**](docs/AccountService.md#listpointrecords) | **POST** /v1/user/account/list-point-records | 
*AccountService* | [**prepareProfileImageUploadAccount**](docs/AccountService.md#prepareprofileimageuploadaccount) | **POST** /v1/user/account/prepare-profile-image-upload | 
*AccountService* | [**searchMentionAccounts**](docs/AccountService.md#searchmentionaccounts) | **POST** /v1/user/account/search-mention | 
*AccountService* | [**transferAccountPoints**](docs/AccountService.md#transferaccountpoints) | **POST** /v1/user/account/economy/transfer | 
*AccountService* | [**updateEmailAccount**](docs/AccountService.md#updateemailaccount) | **POST** /v1/user/account/update-email | 
*AccountService* | [**updatePasswordAccount**](docs/AccountService.md#updatepasswordaccount) | **POST** /v1/user/account/update-password | 
*AccountService* | [**updatePhoneAccount**](docs/AccountService.md#updatephoneaccount) | **POST** /v1/user/account/update-phone | 
*AccountService* | [**updateProfileAccount**](docs/AccountService.md#updateprofileaccount) | **POST** /v1/user/account/update-profile | 
*ArticleSearchService* | [**searchArticles**](docs/ArticleSearchService.md#searcharticles) | **POST** /v1/content/article/search | 
*ArticleService* | [**acceptArticleAnswer**](docs/ArticleService.md#acceptarticleanswer) | **POST** /v1/content/article/answer/accept | 
*ArticleService* | [**archiveArticle**](docs/ArticleService.md#archivearticle) | **POST** /v1/content/article/archive | 
*ArticleService* | [**cancelPublishArticle**](docs/ArticleService.md#cancelpublisharticle) | **POST** /v1/content/article/publish/cancel | 
*ArticleService* | [**collectArticle**](docs/ArticleService.md#collectarticle) | **POST** /v1/content/article/collect | 
*ArticleService* | [**countCollectedArticles**](docs/ArticleService.md#countcollectedarticles) | **POST** /v1/content/article/collect/count | 
*ArticleService* | [**createDraftArticle**](docs/ArticleService.md#createdraftarticle) | **POST** /v1/content/article/draft/create | 
*ArticleService* | [**discardDraftArticle**](docs/ArticleService.md#discarddraftarticle) | **POST** /v1/content/article/draft/discard | 
*ArticleService* | [**enterArticleLottery**](docs/ArticleService.md#enterarticlelottery) | **POST** /v1/content/article/lottery/entry | 
*ArticleService* | [**getArticle**](docs/ArticleService.md#getarticle) | **POST** /v1/content/article/get | 
*ArticleService* | [**likeArticle**](docs/ArticleService.md#likearticle) | **POST** /v1/content/article/like | 
*ArticleService* | [**listArticles**](docs/ArticleService.md#listarticles) | **POST** /v1/content/article/list | 
*ArticleService* | [**listFollowingArticles**](docs/ArticleService.md#listfollowingarticles) | **POST** /v1/content/article/following/list | 
*ArticleService* | [**listNearbyArticles**](docs/ArticleService.md#listnearbyarticles) | **POST** /v1/content/article/nearby/list | 
*ArticleService* | [**pageArticleViewHistory**](docs/ArticleService.md#pagearticleviewhistory) | **POST** /v1/content/article/view-history/page | 
*ArticleService* | [**pageCollectedArticles**](docs/ArticleService.md#pagecollectedarticles) | **POST** /v1/content/article/collected/page | 
*ArticleService* | [**pageWatchedArticles**](docs/ArticleService.md#pagewatchedarticles) | **POST** /v1/content/article/watched/page | 
*ArticleService* | [**publishArticle**](docs/ArticleService.md#publisharticle) | **POST** /v1/content/article/publish | 
*ArticleService* | [**rewardArticle**](docs/ArticleService.md#rewardarticle) | **POST** /v1/content/article/reward | 
*ArticleService* | [**setPinnedArticleComment**](docs/ArticleService.md#setpinnedarticlecomment) | **POST** /v1/content/article/comment/pin | 
*ArticleService* | [**thankArticle**](docs/ArticleService.md#thankarticle) | **POST** /v1/content/article/thank | 
*ArticleService* | [**updateDraftArticle**](docs/ArticleService.md#updatedraftarticle) | **POST** /v1/content/article/draft/update | 
*ArticleService* | [**voteArticlePoll**](docs/ArticleService.md#votearticlepoll) | **POST** /v1/content/article/poll/vote | 
*ArticleService* | [**watchArticle**](docs/ArticleService.md#watcharticle) | **POST** /v1/content/article/watch | 
*AssetService* | [**completeBadgeFontUpload**](docs/AssetService.md#completebadgefontupload) | **POST** /v1/asset/badge-font/complete | 
*AssetService* | [**completeBadgeImageUpload**](docs/AssetService.md#completebadgeimageupload) | **POST** /v1/asset/badge-image/complete | 
*AssetService* | [**completeMarkdownImageUpload**](docs/AssetService.md#completemarkdownimageupload) | **POST** /v1/asset/markdown-image/complete | 
*AssetService* | [**prepareBadgeFontUpload**](docs/AssetService.md#preparebadgefontupload) | **POST** /v1/asset/badge-font/prepare | 
*AssetService* | [**prepareBadgeImageUpload**](docs/AssetService.md#preparebadgeimageupload) | **POST** /v1/asset/badge-image/prepare | 
*AssetService* | [**prepareMarkdownImageUpload**](docs/AssetService.md#preparemarkdownimageupload) | **POST** /v1/asset/markdown-image/prepare | 
*AuthService* | [**cancelAccount**](docs/AuthService.md#cancelaccount) | **POST** /v1/user/auth/cancel-account | 
*AuthService* | [**checkRegistrationAvailability**](docs/AuthService.md#checkregistrationavailability) | **GET** /v1/user/auth/register-availability | 
*AuthService* | [**login**](docs/AuthService.md#login) | **POST** /v1/user/auth/login | 
*AuthService* | [**logout**](docs/AuthService.md#logout) | **POST** /v1/user/auth/logout | 
*AuthService* | [**refreshToken**](docs/AuthService.md#refreshtoken) | **POST** /v1/user/auth/refresh-token | 
*AuthService* | [**register**](docs/AuthService.md#register) | **POST** /v1/user/auth/register | 
*BadgeService* | [**deleteBadgeDraft**](docs/BadgeService.md#deletebadgedraft) | **POST** /v1/badge/draft/delete | 
*BadgeService* | [**listBadgeDrafts**](docs/BadgeService.md#listbadgedrafts) | **POST** /v1/badge/draft/list | 
*BadgeService* | [**renderMakerBadge**](docs/BadgeService.md#rendermakerbadge) | **GET** /v1/badge/maker/render | 
*BadgeService* | [**renderPublishedBadge**](docs/BadgeService.md#renderpublishedbadge) | **GET** /v1/badge/render | 
*BadgeService* | [**upsertBadgeDraft**](docs/BadgeService.md#upsertbadgedraft) | **POST** /v1/badge/draft/upsert | 
*CommentService* | [**createComment**](docs/CommentService.md#createcomment) | **POST** /v1/content/comment/create | 
*CommentService* | [**likeComment**](docs/CommentService.md#likecomment) | **POST** /v1/content/comment/like | 
*CommentService* | [**listCommentReplies**](docs/CommentService.md#listcommentreplies) | **POST** /v1/content/comment/list-replies | 
*CommentService* | [**listCommentThreads**](docs/CommentService.md#listcommentthreads) | **POST** /v1/content/comment/list-threads | 
*CommentService* | [**listCommentTimeline**](docs/CommentService.md#listcommenttimeline) | **POST** /v1/content/comment/list-timeline | 
*CommentService* | [**listComments**](docs/CommentService.md#listcomments) | **POST** /v1/content/comment/list | 
*CommentService* | [**thankComment**](docs/CommentService.md#thankcomment) | **POST** /v1/content/comment/thank | 
*DomainService* | [**createDomain**](docs/DomainService.md#createdomain) | **POST** /v1/content/domain/create | 
*DomainService* | [**listDomains**](docs/DomainService.md#listdomains) | **POST** /v1/content/domain/list | 
*DomainService* | [**updateDomain**](docs/DomainService.md#updatedomain) | **POST** /v1/content/domain/update | 
*EmojiService* | [**addEmojiItem**](docs/EmojiService.md#addemojiitem) | **POST** /v1/im/emoji/items | 
*EmojiService* | [**createEmojiGroup**](docs/EmojiService.md#createemojigroup) | **POST** /v1/im/emoji/groups | 
*EmojiService* | [**createEmojiShare**](docs/EmojiService.md#createemojishare) | **POST** /v1/im/emoji/shares | 
*EmojiService* | [**deleteEmojiGroup**](docs/EmojiService.md#deleteemojigroup) | **POST** /v1/im/emoji/groups/delete | 
*EmojiService* | [**deleteEmojiItem**](docs/EmojiService.md#deleteemojiitem) | **POST** /v1/im/emoji/items/delete | 
*EmojiService* | [**importEmojiShare**](docs/EmojiService.md#importemojishare) | **POST** /v1/im/emoji/shares/import | 
*EmojiService* | [**listEmojiGroups**](docs/EmojiService.md#listemojigroups) | **POST** /v1/im/emoji/groups/list | 
*EmojiService* | [**moveEmojiItem**](docs/EmojiService.md#moveemojiitem) | **POST** /v1/im/emoji/items/move | 
*EmojiService* | [**renameEmojiGroup**](docs/EmojiService.md#renameemojigroup) | **POST** /v1/im/emoji/groups/rename | 
*EmojiService* | [**renameEmojiItem**](docs/EmojiService.md#renameemojiitem) | **POST** /v1/im/emoji/items/rename | 
*EmojiService* | [**reorderEmojiGroups**](docs/EmojiService.md#reorderemojigroups) | **POST** /v1/im/emoji/groups/reorder | 
*EmojiService* | [**reorderEmojiItems**](docs/EmojiService.md#reorderemojiitems) | **POST** /v1/im/emoji/items/reorder | 
*GlobalChatService* | [**getGlobalChatTopic**](docs/GlobalChatService.md#getglobalchattopic) | **POST** /v1/im/global-chat/topic/get | 
*GlobalChatService* | [**listGlobalChatBarrages**](docs/GlobalChatService.md#listglobalchatbarrages) | **POST** /v1/im/global-chat/barrages/list | 
*GlobalChatService* | [**listGlobalChatMessages**](docs/GlobalChatService.md#listglobalchatmessages) | **POST** /v1/im/global-chat/messages/list | 
*GlobalChatService* | [**listGlobalChatOnlineMembers**](docs/GlobalChatService.md#listglobalchatonlinemembers) | **POST** /v1/im/global-chat/online-members/list | 
*GlobalChatService* | [**revokeGlobalChatMessage**](docs/GlobalChatService.md#revokeglobalchatmessage) | **POST** /v1/im/global-chat/messages/revoke | 
*GlobalChatService* | [**sendGlobalChatBarrage**](docs/GlobalChatService.md#sendglobalchatbarrage) | **POST** /v1/im/global-chat/barrages | 
*GlobalChatService* | [**sendGlobalChatMessage**](docs/GlobalChatService.md#sendglobalchatmessage) | **POST** /v1/im/global-chat/messages | 
*GlobalChatService* | [**setGlobalChatTopic**](docs/GlobalChatService.md#setglobalchattopic) | **POST** /v1/im/global-chat/topic/set | 
*HolidayService* | [**getCurrentHoliday**](docs/HolidayService.md#getcurrentholiday) | **POST** /v1/platform/holiday/get-current | 
*LivenessService* | [**claimYesterdayLivenessReward**](docs/LivenessService.md#claimyesterdaylivenessreward) | **POST** /v1/user/liveness/yesterday-reward/claim | 
*LivenessService* | [**getLiveness**](docs/LivenessService.md#getliveness) | **GET** /v1/user/liveness | 
*LocationService* | [**getCurrentLocation**](docs/LocationService.md#getcurrentlocation) | **POST** /v1/platform/location/get-current | 
*MoonbreezeService* | [**createMoonbreeze**](docs/MoonbreezeService.md#createmoonbreeze) | **POST** /v1/content/moonbreeze/create | 
*MoonbreezeService* | [**pageMemberMoonbreezes**](docs/MoonbreezeService.md#pagemembermoonbreezes) | **POST** /v1/content/moonbreeze/page-member | 
*MoonbreezeService* | [**pagePublicMoonbreezes**](docs/MoonbreezeService.md#pagepublicmoonbreezes) | **POST** /v1/content/moonbreeze/page-public | 
*MoonbreezeService* | [**pageWatchingMoonbreezes**](docs/MoonbreezeService.md#pagewatchingmoonbreezes) | **POST** /v1/content/moonbreeze/page-watching | 
*NotificationService* | [**countUnreadNotifications**](docs/NotificationService.md#countunreadnotifications) | **POST** /v1/notify/notification/count-unread | 
*NotificationService* | [**listNotifications**](docs/NotificationService.md#listnotifications) | **POST** /v1/notify/notification/list | 
*NotificationService* | [**markReadNotification**](docs/NotificationService.md#markreadnotification) | **POST** /v1/notify/notification/mark-read | 
*OtpService* | [**beginEnableTotp**](docs/OtpService.md#beginenabletotp) | **POST** /v1/user/otp/totp/begin-enable | 
*OtpService* | [**confirmEnableTotp**](docs/OtpService.md#confirmenabletotp) | **POST** /v1/user/otp/totp/confirm-enable | 
*OtpService* | [**disableTotp**](docs/OtpService.md#disabletotp) | **POST** /v1/user/otp/totp/disable | 
*OtpService* | [**getCurrentTotp**](docs/OtpService.md#getcurrenttotp) | **POST** /v1/user/otp/totp/get-current | 
*OtpService* | [**sendEmailOtp**](docs/OtpService.md#sendemailotp) | **POST** /v1/user/otp/email/send | 
*OtpService* | [**sendPhoneOtp**](docs/OtpService.md#sendphoneotp) | **POST** /v1/user/otp/phone/send | 
*PostscriptService* | [**addPostscript**](docs/PostscriptService.md#addpostscript) | **POST** /v1/content/postscript/add | 
*PostscriptService* | [**listPostscripts**](docs/PostscriptService.md#listpostscripts) | **POST** /v1/content/postscript/list | 
*PreferencesService* | [**getCurrentPreferences**](docs/PreferencesService.md#getcurrentpreferences) | **POST** /v1/user/preference/get-current | 
*PreferencesService* | [**updateCurrentPreferences**](docs/PreferencesService.md#updatecurrentpreferences) | **POST** /v1/user/preference/update-current | 
*PrivacySettingService* | [**getCurrentPrivacySetting**](docs/PrivacySettingService.md#getcurrentprivacysetting) | **POST** /v1/user/privacy-setting/get-current | 
*PrivacySettingService* | [**updateCurrentPrivacySetting**](docs/PrivacySettingService.md#updatecurrentprivacysetting) | **POST** /v1/user/privacy-setting/update-current | 
*PrivateChatService* | [**countPrivateChatUnread**](docs/PrivateChatService.md#countprivatechatunread) | **POST** /v1/im/private-chat/unread-count | 
*PrivateChatService* | [**listPrivateChatMessages**](docs/PrivateChatService.md#listprivatechatmessages) | **POST** /v1/im/private-chat/messages/list | 
*PrivateChatService* | [**listPrivateChatSessions**](docs/PrivateChatService.md#listprivatechatsessions) | **GET** /v1/im/private-chat/sessions | 
*PrivateChatService* | [**markPrivateChatSessionRead**](docs/PrivateChatService.md#markprivatechatsessionread) | **POST** /v1/im/private-chat/read | 
*PrivateChatService* | [**revokePrivateChatMessage**](docs/PrivateChatService.md#revokeprivatechatmessage) | **POST** /v1/im/private-chat/messages/revoke | 
*PrivateChatService* | [**sendPrivateChatMessage**](docs/PrivateChatService.md#sendprivatechatmessage) | **POST** /v1/im/private-chat/messages | 
*PrivateChatService* | [**setPrivateChatSessionMuted**](docs/PrivateChatService.md#setprivatechatsessionmuted) | **POST** /v1/im/private-chat/muted | 
*PrivateChatService* | [**setPrivateChatSessionPinned**](docs/PrivateChatService.md#setprivatechatsessionpinned) | **POST** /v1/im/private-chat/pinned | 
*RealtimeService* | [**createRealtimeTicket**](docs/RealtimeService.md#createrealtimeticket) | **POST** /v1/realtime/ticket | 
*RelationService* | [**blockRelation**](docs/RelationService.md#blockrelation) | **POST** /v1/user/relation/block | 
*RelationService* | [**followRelation**](docs/RelationService.md#followrelation) | **POST** /v1/user/relation/follow | 
*RelationService* | [**getStatusRelation**](docs/RelationService.md#getstatusrelation) | **POST** /v1/user/relation/get-status | 
*RelationService* | [**listBlockedRelations**](docs/RelationService.md#listblockedrelations) | **POST** /v1/user/relation/list-blocked | 
*RelationService* | [**listFollowersRelations**](docs/RelationService.md#listfollowersrelations) | **POST** /v1/user/relation/list-followers | 
*RelationService* | [**listFollowingRelations**](docs/RelationService.md#listfollowingrelations) | **POST** /v1/user/relation/list-following | 
*RelationService* | [**unblockRelation**](docs/RelationService.md#unblockrelation) | **POST** /v1/user/relation/unblock | 
*RelationService* | [**unfollowRelation**](docs/RelationService.md#unfollowrelation) | **POST** /v1/user/relation/unfollow | 
*TagService* | [**bindArticleTags**](docs/TagService.md#bindarticletags) | **POST** /v1/content/tag/bind-article | 
*TagService* | [**createTag**](docs/TagService.md#createtag) | **POST** /v1/content/tag/create | 
*TagService* | [**listArticleTags**](docs/TagService.md#listarticletags) | **POST** /v1/content/tag/list-article-tags | 
*TagService* | [**listTags**](docs/TagService.md#listtags) | **POST** /v1/content/tag/list | 
*TagService* | [**unbindArticleTags**](docs/TagService.md#unbindarticletags) | **POST** /v1/content/tag/unbind-article | 
*TagService* | [**updateTag**](docs/TagService.md#updatetag) | **POST** /v1/content/tag/update | 


### Models

- [AcceptArticleAnswerReq](docs/AcceptArticleAnswerReq.md)
- [AccountContact](docs/AccountContact.md)
- [AccountProfile](docs/AccountProfile.md)
- [AccountProfileListItem](docs/AccountProfileListItem.md)
- [AddEmojiItemReq](docs/AddEmojiItemReq.md)
- [AddEmojiItemResp](docs/AddEmojiItemResp.md)
- [AddPostscriptReq](docs/AddPostscriptReq.md)
- [AddPostscriptResp](docs/AddPostscriptResp.md)
- [ArchiveArticleReq](docs/ArchiveArticleReq.md)
- [ArticleDetail](docs/ArticleDetail.md)
- [ArticleDomain](docs/ArticleDomain.md)
- [ArticleListItem](docs/ArticleListItem.md)
- [ArticleLottery](docs/ArticleLottery.md)
- [ArticleLotteryDraft](docs/ArticleLotteryDraft.md)
- [ArticleLotteryPrize](docs/ArticleLotteryPrize.md)
- [ArticleLotteryPrizeDraft](docs/ArticleLotteryPrizeDraft.md)
- [ArticlePoll](docs/ArticlePoll.md)
- [ArticlePollDraft](docs/ArticlePollDraft.md)
- [ArticlePollOption](docs/ArticlePollOption.md)
- [ArticlePostscript](docs/ArticlePostscript.md)
- [ArticleTag](docs/ArticleTag.md)
- [ArticleViewHistoryItem](docs/ArticleViewHistoryItem.md)
- [ArticleViewerActionState](docs/ArticleViewerActionState.md)
- [AvatarAccountResp](docs/AvatarAccountResp.md)
- [BadgeDraft](docs/BadgeDraft.md)
- [BeginEnableTotpResp](docs/BeginEnableTotpResp.md)
- [BindArticleTagsReq](docs/BindArticleTagsReq.md)
- [BlockRelationReq](docs/BlockRelationReq.md)
- [CancelAccountReq](docs/CancelAccountReq.md)
- [CancelPublishArticleReq](docs/CancelPublishArticleReq.md)
- [CheckRegistrationAvailabilityResp](docs/CheckRegistrationAvailabilityResp.md)
- [ClaimYesterdayLivenessRewardResp](docs/ClaimYesterdayLivenessRewardResp.md)
- [CollectArticleReq](docs/CollectArticleReq.md)
- [CollectArticleResp](docs/CollectArticleResp.md)
- [CommentListItem](docs/CommentListItem.md)
- [CommentViewerActionState](docs/CommentViewerActionState.md)
- [CompleteBadgeFontUploadReq](docs/CompleteBadgeFontUploadReq.md)
- [CompleteBadgeFontUploadResp](docs/CompleteBadgeFontUploadResp.md)
- [CompleteBadgeImageUploadReq](docs/CompleteBadgeImageUploadReq.md)
- [CompleteBadgeImageUploadResp](docs/CompleteBadgeImageUploadResp.md)
- [CompleteMarkdownImageUploadReq](docs/CompleteMarkdownImageUploadReq.md)
- [CompleteMarkdownImageUploadResp](docs/CompleteMarkdownImageUploadResp.md)
- [CompleteProfileImageUploadAccountReq](docs/CompleteProfileImageUploadAccountReq.md)
- [CompleteProfileImageUploadAccountResp](docs/CompleteProfileImageUploadAccountResp.md)
- [ConfirmEnableTotpReq](docs/ConfirmEnableTotpReq.md)
- [CountCollectedArticlesResp](docs/CountCollectedArticlesResp.md)
- [CountPrivateChatUnreadResp](docs/CountPrivateChatUnreadResp.md)
- [CountUnreadNotificationsResp](docs/CountUnreadNotificationsResp.md)
- [CreateCommentReq](docs/CreateCommentReq.md)
- [CreateCommentResp](docs/CreateCommentResp.md)
- [CreateDomainReq](docs/CreateDomainReq.md)
- [CreateDomainResp](docs/CreateDomainResp.md)
- [CreateDraftArticleReq](docs/CreateDraftArticleReq.md)
- [CreateDraftArticleResp](docs/CreateDraftArticleResp.md)
- [CreateEmojiGroupReq](docs/CreateEmojiGroupReq.md)
- [CreateEmojiGroupResp](docs/CreateEmojiGroupResp.md)
- [CreateEmojiShareReq](docs/CreateEmojiShareReq.md)
- [CreateEmojiShareResp](docs/CreateEmojiShareResp.md)
- [CreateMoonbreezeReq](docs/CreateMoonbreezeReq.md)
- [CreateMoonbreezeResp](docs/CreateMoonbreezeResp.md)
- [CreateRealtimeTicketResp](docs/CreateRealtimeTicketResp.md)
- [CreateTagReq](docs/CreateTagReq.md)
- [CreateTagResp](docs/CreateTagResp.md)
- [DeleteBadgeDraftReq](docs/DeleteBadgeDraftReq.md)
- [DeleteEmojiGroupReq](docs/DeleteEmojiGroupReq.md)
- [DeleteEmojiItemReq](docs/DeleteEmojiItemReq.md)
- [DisableTotpReq](docs/DisableTotpReq.md)
- [DiscardDraftArticleReq](docs/DiscardDraftArticleReq.md)
- [EmojiGroup](docs/EmojiGroup.md)
- [EmojiItem](docs/EmojiItem.md)
- [EnterArticleLotteryReq](docs/EnterArticleLotteryReq.md)
- [FollowRelationReq](docs/FollowRelationReq.md)
- [GetAccountEconomyResp](docs/GetAccountEconomyResp.md)
- [GetArticleReq](docs/GetArticleReq.md)
- [GetArticleResp](docs/GetArticleResp.md)
- [GetCurrentAccountResp](docs/GetCurrentAccountResp.md)
- [GetCurrentHolidayResp](docs/GetCurrentHolidayResp.md)
- [GetCurrentLocationResp](docs/GetCurrentLocationResp.md)
- [GetCurrentPreferencesResp](docs/GetCurrentPreferencesResp.md)
- [GetCurrentPrivacySettingResp](docs/GetCurrentPrivacySettingResp.md)
- [GetCurrentTotpResp](docs/GetCurrentTotpResp.md)
- [GetGlobalChatTopicResp](docs/GetGlobalChatTopicResp.md)
- [GetLivenessResp](docs/GetLivenessResp.md)
- [GetProfileReq](docs/GetProfileReq.md)
- [GetProfileResp](docs/GetProfileResp.md)
- [GetStatusRelationReq](docs/GetStatusRelationReq.md)
- [GetStatusRelationResp](docs/GetStatusRelationResp.md)
- [GlobalChatBarrage](docs/GlobalChatBarrage.md)
- [GlobalChatMember](docs/GlobalChatMember.md)
- [GlobalChatMessage](docs/GlobalChatMessage.md)
- [GlobalChatOnlineMember](docs/GlobalChatOnlineMember.md)
- [GlobalChatTopic](docs/GlobalChatTopic.md)
- [ImageResp](docs/ImageResp.md)
- [ImportEmojiShareReq](docs/ImportEmojiShareReq.md)
- [ImportEmojiShareResp](docs/ImportEmojiShareResp.md)
- [LikeArticleReq](docs/LikeArticleReq.md)
- [LikeArticleResp](docs/LikeArticleResp.md)
- [LikeCommentReq](docs/LikeCommentReq.md)
- [LikeCommentResp](docs/LikeCommentResp.md)
- [ListArticleTagsReq](docs/ListArticleTagsReq.md)
- [ListArticleTagsResp](docs/ListArticleTagsResp.md)
- [ListArticlesReq](docs/ListArticlesReq.md)
- [ListArticlesResp](docs/ListArticlesResp.md)
- [ListBadgeDraftsResp](docs/ListBadgeDraftsResp.md)
- [ListBlockedRelationsReq](docs/ListBlockedRelationsReq.md)
- [ListBlockedRelationsResp](docs/ListBlockedRelationsResp.md)
- [ListCommentRepliesReq](docs/ListCommentRepliesReq.md)
- [ListCommentRepliesResp](docs/ListCommentRepliesResp.md)
- [ListCommentThreadsReq](docs/ListCommentThreadsReq.md)
- [ListCommentThreadsResp](docs/ListCommentThreadsResp.md)
- [ListCommentTimelineReq](docs/ListCommentTimelineReq.md)
- [ListCommentTimelineResp](docs/ListCommentTimelineResp.md)
- [ListCommentsReq](docs/ListCommentsReq.md)
- [ListCommentsResp](docs/ListCommentsResp.md)
- [ListDomainsReq](docs/ListDomainsReq.md)
- [ListDomainsResp](docs/ListDomainsResp.md)
- [ListEmojiGroupsReq](docs/ListEmojiGroupsReq.md)
- [ListEmojiGroupsResp](docs/ListEmojiGroupsResp.md)
- [ListFollowersRelationsReq](docs/ListFollowersRelationsReq.md)
- [ListFollowersRelationsResp](docs/ListFollowersRelationsResp.md)
- [ListFollowersReq](docs/ListFollowersReq.md)
- [ListFollowersResp](docs/ListFollowersResp.md)
- [ListFollowingArticlesReq](docs/ListFollowingArticlesReq.md)
- [ListFollowingArticlesResp](docs/ListFollowingArticlesResp.md)
- [ListFollowingRelationsReq](docs/ListFollowingRelationsReq.md)
- [ListFollowingRelationsResp](docs/ListFollowingRelationsResp.md)
- [ListFollowingReq](docs/ListFollowingReq.md)
- [ListFollowingResp](docs/ListFollowingResp.md)
- [ListGlobalChatBarragesReq](docs/ListGlobalChatBarragesReq.md)
- [ListGlobalChatBarragesResp](docs/ListGlobalChatBarragesResp.md)
- [ListGlobalChatMessagesReq](docs/ListGlobalChatMessagesReq.md)
- [ListGlobalChatMessagesResp](docs/ListGlobalChatMessagesResp.md)
- [ListGlobalChatOnlineMembersResp](docs/ListGlobalChatOnlineMembersResp.md)
- [ListNearbyArticlesReq](docs/ListNearbyArticlesReq.md)
- [ListNearbyArticlesResp](docs/ListNearbyArticlesResp.md)
- [ListNotificationsReq](docs/ListNotificationsReq.md)
- [ListNotificationsResp](docs/ListNotificationsResp.md)
- [ListPointRecordsReq](docs/ListPointRecordsReq.md)
- [ListPointRecordsResp](docs/ListPointRecordsResp.md)
- [ListPostscriptsReq](docs/ListPostscriptsReq.md)
- [ListPostscriptsResp](docs/ListPostscriptsResp.md)
- [ListPrivateChatMessagesReq](docs/ListPrivateChatMessagesReq.md)
- [ListPrivateChatMessagesResp](docs/ListPrivateChatMessagesResp.md)
- [ListPrivateChatSessionsResp](docs/ListPrivateChatSessionsResp.md)
- [ListTagsReq](docs/ListTagsReq.md)
- [ListTagsResp](docs/ListTagsResp.md)
- [LoginReq](docs/LoginReq.md)
- [LoginResp](docs/LoginResp.md)
- [MarkPrivateChatSessionReadReq](docs/MarkPrivateChatSessionReadReq.md)
- [MarkReadNotificationReq](docs/MarkReadNotificationReq.md)
- [MarkReadNotificationResp](docs/MarkReadNotificationResp.md)
- [Moonbreeze](docs/Moonbreeze.md)
- [MoveEmojiItemReq](docs/MoveEmojiItemReq.md)
- [MoveEmojiItemResp](docs/MoveEmojiItemResp.md)
- [PageArticleViewHistoryReq](docs/PageArticleViewHistoryReq.md)
- [PageArticleViewHistoryResp](docs/PageArticleViewHistoryResp.md)
- [PageCollectedArticlesReq](docs/PageCollectedArticlesReq.md)
- [PageCollectedArticlesResp](docs/PageCollectedArticlesResp.md)
- [PageMemberMoonbreezesReq](docs/PageMemberMoonbreezesReq.md)
- [PageMemberMoonbreezesResp](docs/PageMemberMoonbreezesResp.md)
- [PagePublicMoonbreezesReq](docs/PagePublicMoonbreezesReq.md)
- [PagePublicMoonbreezesResp](docs/PagePublicMoonbreezesResp.md)
- [PageReq](docs/PageReq.md)
- [PageResp](docs/PageResp.md)
- [PageWatchedArticlesReq](docs/PageWatchedArticlesReq.md)
- [PageWatchedArticlesResp](docs/PageWatchedArticlesResp.md)
- [PageWatchingMoonbreezesReq](docs/PageWatchingMoonbreezesReq.md)
- [PageWatchingMoonbreezesResp](docs/PageWatchingMoonbreezesResp.md)
- [PrepareBadgeFontUploadReq](docs/PrepareBadgeFontUploadReq.md)
- [PrepareBadgeFontUploadResp](docs/PrepareBadgeFontUploadResp.md)
- [PrepareBadgeImageUploadReq](docs/PrepareBadgeImageUploadReq.md)
- [PrepareBadgeImageUploadResp](docs/PrepareBadgeImageUploadResp.md)
- [PrepareMarkdownImageUploadReq](docs/PrepareMarkdownImageUploadReq.md)
- [PrepareMarkdownImageUploadResp](docs/PrepareMarkdownImageUploadResp.md)
- [PrepareProfileImageUploadAccountReq](docs/PrepareProfileImageUploadAccountReq.md)
- [PrepareProfileImageUploadAccountResp](docs/PrepareProfileImageUploadAccountResp.md)
- [PrivateChatMember](docs/PrivateChatMember.md)
- [PrivateChatMessage](docs/PrivateChatMessage.md)
- [PrivateChatSession](docs/PrivateChatSession.md)
- [Profile](docs/Profile.md)
- [ProfileLocation](docs/ProfileLocation.md)
- [ProfileRelation](docs/ProfileRelation.md)
- [ProfileVisibility](docs/ProfileVisibility.md)
- [PublishArticleReq](docs/PublishArticleReq.md)
- [RefreshTokenReq](docs/RefreshTokenReq.md)
- [RefreshTokenResp](docs/RefreshTokenResp.md)
- [RegisterReq](docs/RegisterReq.md)
- [Relation](docs/Relation.md)
- [RenameEmojiGroupReq](docs/RenameEmojiGroupReq.md)
- [RenameEmojiGroupResp](docs/RenameEmojiGroupResp.md)
- [RenameEmojiItemReq](docs/RenameEmojiItemReq.md)
- [RenameEmojiItemResp](docs/RenameEmojiItemResp.md)
- [RenderMakerBadgeResp](docs/RenderMakerBadgeResp.md)
- [RenderPublishedBadgeResp](docs/RenderPublishedBadgeResp.md)
- [RenderSpec](docs/RenderSpec.md)
- [ReorderEmojiGroupsReq](docs/ReorderEmojiGroupsReq.md)
- [ReorderEmojiItemsReq](docs/ReorderEmojiItemsReq.md)
- [ReqArticle](docs/ReqArticle.md)
- [ReqArticleQuery](docs/ReqArticleQuery.md)
- [ReqCommentQuery](docs/ReqCommentQuery.md)
- [ReqDomain](docs/ReqDomain.md)
- [ReqDomainQuery](docs/ReqDomainQuery.md)
- [ReqEmailCredential](docs/ReqEmailCredential.md)
- [ReqPasswordCredential](docs/ReqPasswordCredential.md)
- [ReqPhoneCredential](docs/ReqPhoneCredential.md)
- [ReqTag](docs/ReqTag.md)
- [ReqTagQuery](docs/ReqTagQuery.md)
- [RespAccount](docs/RespAccount.md)
- [RespAccountBasic](docs/RespAccountBasic.md)
- [RespAccountContact](docs/RespAccountContact.md)
- [RespAccountProfile](docs/RespAccountProfile.md)
- [RespArticleBrief](docs/RespArticleBrief.md)
- [RespCommentDetail](docs/RespCommentDetail.md)
- [RespCommentListItem](docs/RespCommentListItem.md)
- [RespCommentThread](docs/RespCommentThread.md)
- [RespCommentViewerActionState](docs/RespCommentViewerActionState.md)
- [RespCurrentAccount](docs/RespCurrentAccount.md)
- [RespDomain](docs/RespDomain.md)
- [RespLocation](docs/RespLocation.md)
- [RespNotification](docs/RespNotification.md)
- [RespPinnedComment](docs/RespPinnedComment.md)
- [RespPointRecord](docs/RespPointRecord.md)
- [RespPreference](docs/RespPreference.md)
- [RespPrivacySetting](docs/RespPrivacySetting.md)
- [RespRelationStatus](docs/RespRelationStatus.md)
- [RespTag](docs/RespTag.md)
- [RespTotp](docs/RespTotp.md)
- [RevokeGlobalChatMessageReq](docs/RevokeGlobalChatMessageReq.md)
- [RevokeGlobalChatMessageResp](docs/RevokeGlobalChatMessageResp.md)
- [RevokePrivateChatMessageReq](docs/RevokePrivateChatMessageReq.md)
- [RevokePrivateChatMessageResp](docs/RevokePrivateChatMessageResp.md)
- [RewardArticleReq](docs/RewardArticleReq.md)
- [RewardArticleResp](docs/RewardArticleResp.md)
- [SearchArticlesReq](docs/SearchArticlesReq.md)
- [SearchArticlesResp](docs/SearchArticlesResp.md)
- [SearchMentionAccountsReq](docs/SearchMentionAccountsReq.md)
- [SearchMentionAccountsResp](docs/SearchMentionAccountsResp.md)
- [SendEmailOtpReq](docs/SendEmailOtpReq.md)
- [SendEmailOtpResp](docs/SendEmailOtpResp.md)
- [SendGlobalChatBarrageReq](docs/SendGlobalChatBarrageReq.md)
- [SendGlobalChatBarrageResp](docs/SendGlobalChatBarrageResp.md)
- [SendGlobalChatMessageReq](docs/SendGlobalChatMessageReq.md)
- [SendGlobalChatMessageResp](docs/SendGlobalChatMessageResp.md)
- [SendPhoneOtpReq](docs/SendPhoneOtpReq.md)
- [SendPhoneOtpResp](docs/SendPhoneOtpResp.md)
- [SendPrivateChatMessageReq](docs/SendPrivateChatMessageReq.md)
- [SendPrivateChatMessageResp](docs/SendPrivateChatMessageResp.md)
- [SetGlobalChatTopicReq](docs/SetGlobalChatTopicReq.md)
- [SetGlobalChatTopicResp](docs/SetGlobalChatTopicResp.md)
- [SetPinnedArticleCommentReq](docs/SetPinnedArticleCommentReq.md)
- [SetPrivateChatSessionMutedReq](docs/SetPrivateChatSessionMutedReq.md)
- [SetPrivateChatSessionPinnedReq](docs/SetPrivateChatSessionPinnedReq.md)
- [ThankArticleReq](docs/ThankArticleReq.md)
- [ThankArticleResp](docs/ThankArticleResp.md)
- [ThankCommentReq](docs/ThankCommentReq.md)
- [ThankCommentResp](docs/ThankCommentResp.md)
- [TransferAccountPointsReq](docs/TransferAccountPointsReq.md)
- [TransferAccountPointsResp](docs/TransferAccountPointsResp.md)
- [UnbindArticleTagsReq](docs/UnbindArticleTagsReq.md)
- [UnblockRelationReq](docs/UnblockRelationReq.md)
- [UnfollowRelationReq](docs/UnfollowRelationReq.md)
- [UpdateCurrentPreferencesReq](docs/UpdateCurrentPreferencesReq.md)
- [UpdateCurrentPreferencesResp](docs/UpdateCurrentPreferencesResp.md)
- [UpdateCurrentPrivacySettingReq](docs/UpdateCurrentPrivacySettingReq.md)
- [UpdateCurrentPrivacySettingResp](docs/UpdateCurrentPrivacySettingResp.md)
- [UpdateDomainReq](docs/UpdateDomainReq.md)
- [UpdateDomainResp](docs/UpdateDomainResp.md)
- [UpdateDraftArticleReq](docs/UpdateDraftArticleReq.md)
- [UpdateDraftArticleResp](docs/UpdateDraftArticleResp.md)
- [UpdateEmailAccountReq](docs/UpdateEmailAccountReq.md)
- [UpdatePasswordAccountReq](docs/UpdatePasswordAccountReq.md)
- [UpdatePhoneAccountReq](docs/UpdatePhoneAccountReq.md)
- [UpdateProfileAccountReq](docs/UpdateProfileAccountReq.md)
- [UpdateProfileAccountResp](docs/UpdateProfileAccountResp.md)
- [UpdateTagReq](docs/UpdateTagReq.md)
- [UpdateTagResp](docs/UpdateTagResp.md)
- [UpsertBadgeDraftReq](docs/UpsertBadgeDraftReq.md)
- [UpsertBadgeDraftResp](docs/UpsertBadgeDraftResp.md)
- [VoteArticlePollReq](docs/VoteArticlePollReq.md)
- [WatchArticleReq](docs/WatchArticleReq.md)
- [WatchArticleResp](docs/WatchArticleResp.md)

### Authorization

Endpoints do not require authorization.


## About

This TypeScript SDK client supports the [Fetch API](https://fetch.spec.whatwg.org/)
and is automatically generated by the
[OpenAPI Generator](https://openapi-generator.tech) project:

- API version: `1.0.0`
- Package version: `0.1.0`
- Generator version: `7.22.0`
- Build package: `org.openapitools.codegen.languages.TypeScriptFetchClientCodegen`

The generated npm module supports the following:

- Environments
  * Node.js
  * Webpack
  * Browserify
- Language levels
  * ES5 - you must have a Promises/A+ library installed
  * ES6
- Module systems
  * CommonJS
  * ES6 module system


## Development

### Building

To build the TypeScript source code, you need to have Node.js and npm installed.
After cloning the repository, navigate to the project directory and run:

```bash
npm install
npm run build
```

### Publishing

Once you've built the package, you can publish it to npm:

```bash
npm publish
```

## License

[]()
