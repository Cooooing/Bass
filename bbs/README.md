## @bass/bbs-sdk-axios@0.1.0

This generator creates TypeScript/JavaScript client that utilizes [axios](https://github.com/axios/axios). The generated Node module can be used in the following environments:

Environment
* Node.js
* Webpack
* Browserify

Language level
* ES5 - you must have a Promises/A+ library installed
* ES6

Module system
* CommonJS
* ES6 module system

It can be used in both TypeScript and JavaScript. In TypeScript, the definition will be automatically resolved via `package.json`. ([Reference](https://www.typescriptlang.org/docs/handbook/declaration-files/consumption.html))

### Building

To build and compile the typescript sources to javascript use:
```
npm install
npm run build
```

### Publishing

First build the package then run `npm publish`

### Consuming

navigate to the folder of your consuming project and run one of the following commands.

_published:_

```
npm install @bass/bbs-sdk-axios@0.1.0 --save
```

_unPublished (not recommended):_

```
npm install PATH_TO_GENERATED_PACKAGE --save
```

### Documentation for API Endpoints

All URIs are relative to *http://localhost*

Class | Method | HTTP request | Description
------------ | ------------- | ------------- | -------------
*AccountService* | [**avatar**](docs/AccountService.md#avatar) | **GET** /v1/user/account/avatar | 
*AccountService* | [**getCurrent**](docs/AccountService.md#getcurrent) | **POST** /v1/user/account/get-current | 
*AccountService* | [**getProfile**](docs/AccountService.md#getprofile) | **POST** /v1/user/account/get-profile | 
*AccountService* | [**updateProfile**](docs/AccountService.md#updateprofile) | **POST** /v1/user/account/update-profile | 
*ArticleService* | [**acceptAnswer**](docs/ArticleService.md#acceptanswer) | **POST** /v1/content/article/accept-answer | 
*ArticleService* | [**collect**](docs/ArticleService.md#collect) | **POST** /v1/content/article/collect | 
*ArticleService* | [**create**](docs/ArticleService.md#create) | **POST** /v1/content/article/create | 
*ArticleService* | [**discardDraft**](docs/ArticleService.md#discarddraft) | **POST** /v1/content/article/discard-draft | 
*ArticleService* | [**get**](docs/ArticleService.md#get) | **POST** /v1/content/article/get | 
*ArticleService* | [**like**](docs/ArticleService.md#like) | **POST** /v1/content/article/like | 
*ArticleService* | [**list**](docs/ArticleService.md#list) | **POST** /v1/content/article/list | 
*ArticleService* | [**publish**](docs/ArticleService.md#publish) | **POST** /v1/content/article/publish | 
*ArticleService* | [**reward**](docs/ArticleService.md#reward) | **POST** /v1/content/article/reward | 
*ArticleService* | [**thank**](docs/ArticleService.md#thank) | **POST** /v1/content/article/thank | 
*ArticleService* | [**update**](docs/ArticleService.md#update) | **POST** /v1/content/article/update | 
*ArticleService* | [**updateDraft**](docs/ArticleService.md#updatedraft) | **POST** /v1/content/article/update-draft | 
*ArticleService* | [**watch**](docs/ArticleService.md#watch) | **POST** /v1/content/article/watch | 
*AuthService* | [**loginByPassword**](docs/AuthService.md#loginbypassword) | **POST** /v1/user/auth/login-by-password | 
*AuthService* | [**logout**](docs/AuthService.md#logout) | **POST** /v1/user/auth/logout | 
*AuthService* | [**startEmailRegistration**](docs/AuthService.md#startemailregistration) | **POST** /v1/user/auth/start-email-registration | 
*AuthService* | [**startPhoneRegistration**](docs/AuthService.md#startphoneregistration) | **POST** /v1/user/auth/start-phone-registration | 
*AuthService* | [**verifyEmailRegistration**](docs/AuthService.md#verifyemailregistration) | **POST** /v1/user/auth/verify-email-registration | 
*AuthService* | [**verifyPhoneRegistration**](docs/AuthService.md#verifyphoneregistration) | **POST** /v1/user/auth/verify-phone-registration | 
*CommentService* | [**create**](docs/CommentService.md#create) | **POST** /v1/content/comment/create | 
*CommentService* | [**like**](docs/CommentService.md#like) | **POST** /v1/content/comment/like | 
*CommentService* | [**list**](docs/CommentService.md#list) | **POST** /v1/content/comment/list | 
*CommentService* | [**listReplies**](docs/CommentService.md#listreplies) | **POST** /v1/content/comment/list-replies | 
*CommentService* | [**listThreads**](docs/CommentService.md#listthreads) | **POST** /v1/content/comment/list-threads | 
*CommentService* | [**listTimeline**](docs/CommentService.md#listtimeline) | **POST** /v1/content/comment/list-timeline | 
*CommentService* | [**thank**](docs/CommentService.md#thank) | **POST** /v1/content/comment/thank | 
*DomainService* | [**list**](docs/DomainService.md#list) | **POST** /v1/content/domain/list | 
*LocationService* | [**getCurrent**](docs/LocationService.md#getcurrent) | **POST** /v1/user/location/get-current | 
*LocationService* | [**upsertCurrent**](docs/LocationService.md#upsertcurrent) | **POST** /v1/user/location/upsert-current | 
*NotificationService* | [**countUnread**](docs/NotificationService.md#countunread) | **POST** /v1/notify/notification/count-unread | 
*NotificationService* | [**list**](docs/NotificationService.md#list) | **POST** /v1/notify/notification/list | 
*NotificationService* | [**markRead**](docs/NotificationService.md#markread) | **POST** /v1/notify/notification/mark-read | 
*PostscriptService* | [**add**](docs/PostscriptService.md#add) | **POST** /v1/content/postscript/add | 
*PreferencesService* | [**getCurrent**](docs/PreferencesService.md#getcurrent) | **POST** /v1/user/preference/get-current | 
*PreferencesService* | [**updateCurrent**](docs/PreferencesService.md#updatecurrent) | **POST** /v1/user/preference/update-current | 
*PrivacySettingService* | [**getCurrent**](docs/PrivacySettingService.md#getcurrent) | **POST** /v1/user/privacy-setting/get-current | 
*PrivacySettingService* | [**updateCurrent**](docs/PrivacySettingService.md#updatecurrent) | **POST** /v1/user/privacy-setting/update-current | 
*RelationService* | [**block**](docs/RelationService.md#block) | **POST** /v1/user/relation/block | 
*RelationService* | [**follow**](docs/RelationService.md#follow) | **POST** /v1/user/relation/follow | 
*RelationService* | [**getStatus**](docs/RelationService.md#getstatus) | **POST** /v1/user/relation/get-status | 
*RelationService* | [**listBlocked**](docs/RelationService.md#listblocked) | **POST** /v1/user/relation/list-blocked | 
*RelationService* | [**listFollowers**](docs/RelationService.md#listfollowers) | **POST** /v1/user/relation/list-followers | 
*RelationService* | [**listFollowing**](docs/RelationService.md#listfollowing) | **POST** /v1/user/relation/list-following | 
*RelationService* | [**unblock**](docs/RelationService.md#unblock) | **POST** /v1/user/relation/unblock | 
*RelationService* | [**unfollow**](docs/RelationService.md#unfollow) | **POST** /v1/user/relation/unfollow | 
*TagService* | [**create**](docs/TagService.md#create) | **POST** /v1/content/tag/create | 
*TagService* | [**list**](docs/TagService.md#list) | **POST** /v1/content/tag/list | 
*TagService* | [**update**](docs/TagService.md#update) | **POST** /v1/content/tag/update | 
*TotpService* | [**beginEnable**](docs/TotpService.md#beginenable) | **POST** /v1/user/totp/begin-enable | 
*TotpService* | [**confirmEnable**](docs/TotpService.md#confirmenable) | **POST** /v1/user/totp/confirm-enable | 
*TotpService* | [**disable**](docs/TotpService.md#disable) | **POST** /v1/user/totp/disable | 
*TotpService* | [**getCurrent**](docs/TotpService.md#getcurrent) | **POST** /v1/user/totp/get-current | 


### Documentation For Models

 - [AcceptAnswerArticleRequest](docs/AcceptAnswerArticleRequest.md)
 - [Account](docs/Account.md)
 - [AccountContact](docs/AccountContact.md)
 - [AccountProfile](docs/AccountProfile.md)
 - [AddPostscriptReply](docs/AddPostscriptReply.md)
 - [AddPostscriptRequest](docs/AddPostscriptRequest.md)
 - [ArticleBrief](docs/ArticleBrief.md)
 - [ArticleDetail](docs/ArticleDetail.md)
 - [ArticleListItem](docs/ArticleListItem.md)
 - [ArticlePostscript](docs/ArticlePostscript.md)
 - [ArticleQuery](docs/ArticleQuery.md)
 - [ArticleViewerActionState](docs/ArticleViewerActionState.md)
 - [BeginEnableTotpReply](docs/BeginEnableTotpReply.md)
 - [BlockRelationRequest](docs/BlockRelationRequest.md)
 - [CollectArticleReply](docs/CollectArticleReply.md)
 - [CollectArticleRequest](docs/CollectArticleRequest.md)
 - [CommentDetail](docs/CommentDetail.md)
 - [CommentListItem](docs/CommentListItem.md)
 - [CommentQuery](docs/CommentQuery.md)
 - [CommentThread](docs/CommentThread.md)
 - [CommentViewerActionState](docs/CommentViewerActionState.md)
 - [ConfirmEnableTotpRequest](docs/ConfirmEnableTotpRequest.md)
 - [CountUnreadNotificationsReply](docs/CountUnreadNotificationsReply.md)
 - [CreateArticleReply](docs/CreateArticleReply.md)
 - [CreateArticleRequest](docs/CreateArticleRequest.md)
 - [CreateCommentReply](docs/CreateCommentReply.md)
 - [CreateCommentRequest](docs/CreateCommentRequest.md)
 - [CreateTagReply](docs/CreateTagReply.md)
 - [CreateTagRequest](docs/CreateTagRequest.md)
 - [DisableTotpRequest](docs/DisableTotpRequest.md)
 - [DiscardDraftArticleRequest](docs/DiscardDraftArticleRequest.md)
 - [Domain](docs/Domain.md)
 - [DomainQuery](docs/DomainQuery.md)
 - [FollowRelationRequest](docs/FollowRelationRequest.md)
 - [GetArticleReply](docs/GetArticleReply.md)
 - [GetArticleRequest](docs/GetArticleRequest.md)
 - [GetCurrentAccountReply](docs/GetCurrentAccountReply.md)
 - [GetCurrentLocationReply](docs/GetCurrentLocationReply.md)
 - [GetCurrentPreferencesReply](docs/GetCurrentPreferencesReply.md)
 - [GetCurrentPrivacySettingReply](docs/GetCurrentPrivacySettingReply.md)
 - [GetCurrentTotpReply](docs/GetCurrentTotpReply.md)
 - [GetProfileAccountReply](docs/GetProfileAccountReply.md)
 - [GetProfileAccountRequest](docs/GetProfileAccountRequest.md)
 - [GetStatusRelationReply](docs/GetStatusRelationReply.md)
 - [GetStatusRelationRequest](docs/GetStatusRelationRequest.md)
 - [ImageReply](docs/ImageReply.md)
 - [LikeArticleReply](docs/LikeArticleReply.md)
 - [LikeArticleRequest](docs/LikeArticleRequest.md)
 - [LikeCommentReply](docs/LikeCommentReply.md)
 - [LikeCommentRequest](docs/LikeCommentRequest.md)
 - [ListArticlesReply](docs/ListArticlesReply.md)
 - [ListArticlesRequest](docs/ListArticlesRequest.md)
 - [ListBlockedRelationsReply](docs/ListBlockedRelationsReply.md)
 - [ListBlockedRelationsRequest](docs/ListBlockedRelationsRequest.md)
 - [ListCommentRepliesReply](docs/ListCommentRepliesReply.md)
 - [ListCommentRepliesRequest](docs/ListCommentRepliesRequest.md)
 - [ListCommentThreadsReply](docs/ListCommentThreadsReply.md)
 - [ListCommentThreadsRequest](docs/ListCommentThreadsRequest.md)
 - [ListCommentTimelineReply](docs/ListCommentTimelineReply.md)
 - [ListCommentTimelineRequest](docs/ListCommentTimelineRequest.md)
 - [ListCommentsReply](docs/ListCommentsReply.md)
 - [ListCommentsRequest](docs/ListCommentsRequest.md)
 - [ListDomainsReply](docs/ListDomainsReply.md)
 - [ListDomainsRequest](docs/ListDomainsRequest.md)
 - [ListFollowersRelationsReply](docs/ListFollowersRelationsReply.md)
 - [ListFollowersRelationsRequest](docs/ListFollowersRelationsRequest.md)
 - [ListFollowingRelationsReply](docs/ListFollowingRelationsReply.md)
 - [ListFollowingRelationsRequest](docs/ListFollowingRelationsRequest.md)
 - [ListNotificationsReply](docs/ListNotificationsReply.md)
 - [ListNotificationsRequest](docs/ListNotificationsRequest.md)
 - [ListTagsReply](docs/ListTagsReply.md)
 - [ListTagsRequest](docs/ListTagsRequest.md)
 - [Location](docs/Location.md)
 - [LoginByPasswordReply](docs/LoginByPasswordReply.md)
 - [LoginByPasswordRequest](docs/LoginByPasswordRequest.md)
 - [MarkReadNotificationReply](docs/MarkReadNotificationReply.md)
 - [MarkReadNotificationRequest](docs/MarkReadNotificationRequest.md)
 - [Notification](docs/Notification.md)
 - [PageReply](docs/PageReply.md)
 - [PageRequest](docs/PageRequest.md)
 - [Preference](docs/Preference.md)
 - [PrivacySetting](docs/PrivacySetting.md)
 - [PublishArticleRequest](docs/PublishArticleRequest.md)
 - [Relation](docs/Relation.md)
 - [RelationStatus](docs/RelationStatus.md)
 - [RequestArticle](docs/RequestArticle.md)
 - [RequestTag](docs/RequestTag.md)
 - [RewardArticleRequest](docs/RewardArticleRequest.md)
 - [StartEmailRegistrationReply](docs/StartEmailRegistrationReply.md)
 - [StartEmailRegistrationRequest](docs/StartEmailRegistrationRequest.md)
 - [StartPhoneRegistrationReply](docs/StartPhoneRegistrationReply.md)
 - [StartPhoneRegistrationRequest](docs/StartPhoneRegistrationRequest.md)
 - [Tag](docs/Tag.md)
 - [TagQuery](docs/TagQuery.md)
 - [ThankArticleReply](docs/ThankArticleReply.md)
 - [ThankArticleRequest](docs/ThankArticleRequest.md)
 - [ThankCommentReply](docs/ThankCommentReply.md)
 - [ThankCommentRequest](docs/ThankCommentRequest.md)
 - [Totp](docs/Totp.md)
 - [UnblockRelationRequest](docs/UnblockRelationRequest.md)
 - [UnfollowRelationRequest](docs/UnfollowRelationRequest.md)
 - [UpdateArticleReply](docs/UpdateArticleReply.md)
 - [UpdateArticleRequest](docs/UpdateArticleRequest.md)
 - [UpdateCurrentPreferencesReply](docs/UpdateCurrentPreferencesReply.md)
 - [UpdateCurrentPreferencesRequest](docs/UpdateCurrentPreferencesRequest.md)
 - [UpdateCurrentPrivacySettingReply](docs/UpdateCurrentPrivacySettingReply.md)
 - [UpdateCurrentPrivacySettingRequest](docs/UpdateCurrentPrivacySettingRequest.md)
 - [UpdateDraftArticleReply](docs/UpdateDraftArticleReply.md)
 - [UpdateDraftArticleRequest](docs/UpdateDraftArticleRequest.md)
 - [UpdateProfileAccountReply](docs/UpdateProfileAccountReply.md)
 - [UpdateProfileAccountRequest](docs/UpdateProfileAccountRequest.md)
 - [UpdateTagReply](docs/UpdateTagReply.md)
 - [UpdateTagRequest](docs/UpdateTagRequest.md)
 - [UpsertCurrentLocationReply](docs/UpsertCurrentLocationReply.md)
 - [UpsertCurrentLocationRequest](docs/UpsertCurrentLocationRequest.md)
 - [VerifyEmailRegistrationRequest](docs/VerifyEmailRegistrationRequest.md)
 - [VerifyPhoneRegistrationRequest](docs/VerifyPhoneRegistrationRequest.md)
 - [WatchArticleReply](docs/WatchArticleReply.md)
 - [WatchArticleRequest](docs/WatchArticleRequest.md)


<a id="documentation-for-authorization"></a>
## Documentation For Authorization

Endpoints do not require authorization.

