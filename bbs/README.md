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
import type { AvatarRequest } from '@bass/bbs-sdk-fetch';

async function example() {
  console.log("🚀 Testing @bass/bbs-sdk-fetch SDK...");
  const api = new AccountService();

  const body = {
    // string (optional)
    name: name_example,
  } satisfies AvatarRequest;

  try {
    const data = await api.avatar(body);
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
*AccountService* | [**avatar**](docs/AccountService.md#avatar) | **GET** /v1/user/account/avatar | 
*AccountService* | [**getCurrent**](docs/AccountService.md#getcurrent) | **POST** /v1/user/account/get-current | 
*AccountService* | [**getProfile**](docs/AccountService.md#getprofile) | **POST** /v1/user/account/get-profile | 
*AccountService* | [**updateEmail**](docs/AccountService.md#updateemail) | **POST** /v1/user/account/update-email | 
*AccountService* | [**updatePassword**](docs/AccountService.md#updatepassword) | **POST** /v1/user/account/update-password | 
*AccountService* | [**updatePhone**](docs/AccountService.md#updatephone) | **POST** /v1/user/account/update-phone | 
*AccountService* | [**updateProfile**](docs/AccountService.md#updateprofile) | **POST** /v1/user/account/update-profile | 
*ArticleService* | [**archive**](docs/ArticleService.md#archive) | **POST** /v1/content/article/archive | 
*ArticleService* | [**cancelPublish**](docs/ArticleService.md#cancelpublish) | **POST** /v1/content/article/publish/cancel | 
*ArticleService* | [**collect**](docs/ArticleService.md#collect) | **POST** /v1/content/article/collect | 
*ArticleService* | [**createDraft**](docs/ArticleService.md#createdraft) | **POST** /v1/content/article/draft/create | 
*ArticleService* | [**discardDraft**](docs/ArticleService.md#discarddraft) | **POST** /v1/content/article/draft/discard | 
*ArticleService* | [**get**](docs/ArticleService.md#get) | **POST** /v1/content/article/get | 
*ArticleService* | [**like**](docs/ArticleService.md#like) | **POST** /v1/content/article/like | 
*ArticleService* | [**list**](docs/ArticleService.md#list) | **POST** /v1/content/article/list | 
*ArticleService* | [**publish**](docs/ArticleService.md#publish) | **POST** /v1/content/article/publish | 
*ArticleService* | [**reward**](docs/ArticleService.md#reward) | **POST** /v1/content/article/reward | 
*ArticleService* | [**schedulePublish**](docs/ArticleService.md#schedulepublish) | **POST** /v1/content/article/publish/schedule | 
*ArticleService* | [**thank**](docs/ArticleService.md#thank) | **POST** /v1/content/article/thank | 
*ArticleService* | [**updateDraft**](docs/ArticleService.md#updatedraft) | **POST** /v1/content/article/draft/update | 
*AuthService* | [**cancelAccount**](docs/AuthService.md#cancelaccount) | **POST** /v1/user/auth/cancel-account | 
*AuthService* | [**login**](docs/AuthService.md#login) | **POST** /v1/user/auth/login | 
*AuthService* | [**logout**](docs/AuthService.md#logout) | **POST** /v1/user/auth/logout | 
*AuthService* | [**refreshToken**](docs/AuthService.md#refreshtoken) | **POST** /v1/user/auth/refresh-token | 
*AuthService* | [**register**](docs/AuthService.md#register) | **POST** /v1/user/auth/register | 
*CommentService* | [**create**](docs/CommentService.md#create) | **POST** /v1/content/comment/create | 
*CommentService* | [**like**](docs/CommentService.md#like) | **POST** /v1/content/comment/like | 
*CommentService* | [**list**](docs/CommentService.md#list) | **POST** /v1/content/comment/list | 
*CommentService* | [**listReplies**](docs/CommentService.md#listreplies) | **POST** /v1/content/comment/list-replies | 
*CommentService* | [**listThreads**](docs/CommentService.md#listthreads) | **POST** /v1/content/comment/list-threads | 
*CommentService* | [**listTimeline**](docs/CommentService.md#listtimeline) | **POST** /v1/content/comment/list-timeline | 
*CommentService* | [**thank**](docs/CommentService.md#thank) | **POST** /v1/content/comment/thank | 
*DomainService* | [**create**](docs/DomainService.md#create) | **POST** /v1/content/domain/create | 
*DomainService* | [**list**](docs/DomainService.md#list) | **POST** /v1/content/domain/list | 
*DomainService* | [**update**](docs/DomainService.md#update) | **POST** /v1/content/domain/update | 
*LocationService* | [**getCurrent**](docs/LocationService.md#getcurrent) | **POST** /v1/user/location/get-current | 
*LocationService* | [**upsertCurrent**](docs/LocationService.md#upsertcurrent) | **POST** /v1/user/location/upsert-current | 
*NotificationService* | [**countUnread**](docs/NotificationService.md#countunread) | **POST** /v1/notify/notification/count-unread | 
*NotificationService* | [**list**](docs/NotificationService.md#list) | **POST** /v1/notify/notification/list | 
*NotificationService* | [**markRead**](docs/NotificationService.md#markread) | **POST** /v1/notify/notification/mark-read | 
*OtpService* | [**beginEnableTotp**](docs/OtpService.md#beginenabletotp) | **POST** /v1/user/otp/totp/begin-enable | 
*OtpService* | [**confirmEnableTotp**](docs/OtpService.md#confirmenabletotp) | **POST** /v1/user/otp/totp/confirm-enable | 
*OtpService* | [**disableTotp**](docs/OtpService.md#disabletotp) | **POST** /v1/user/otp/totp/disable | 
*OtpService* | [**getCurrentTotp**](docs/OtpService.md#getcurrenttotp) | **POST** /v1/user/otp/totp/get-current | 
*OtpService* | [**sendEmailOtp**](docs/OtpService.md#sendemailotp) | **POST** /v1/user/otp/email/send | 
*OtpService* | [**sendPhoneOtp**](docs/OtpService.md#sendphoneotp) | **POST** /v1/user/otp/phone/send | 
*PostscriptService* | [**add**](docs/PostscriptService.md#add) | **POST** /v1/content/postscript/add | 
*PostscriptService* | [**list**](docs/PostscriptService.md#list) | **POST** /v1/content/postscript/list | 
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
*TagService* | [**bindArticle**](docs/TagService.md#bindarticle) | **POST** /v1/content/tag/bind-article | 
*TagService* | [**create**](docs/TagService.md#create) | **POST** /v1/content/tag/create | 
*TagService* | [**list**](docs/TagService.md#list) | **POST** /v1/content/tag/list | 
*TagService* | [**listArticleTags**](docs/TagService.md#listarticletags) | **POST** /v1/content/tag/list-article-tags | 
*TagService* | [**unbindArticle**](docs/TagService.md#unbindarticle) | **POST** /v1/content/tag/unbind-article | 
*TagService* | [**update**](docs/TagService.md#update) | **POST** /v1/content/tag/update | 


### Models

- [AccountProfile](docs/AccountProfile.md)
- [AddPostscriptReq](docs/AddPostscriptReq.md)
- [AddPostscriptResp](docs/AddPostscriptResp.md)
- [ArchiveArticleReq](docs/ArchiveArticleReq.md)
- [ArticleDetail](docs/ArticleDetail.md)
- [ArticleListItem](docs/ArticleListItem.md)
- [ArticlePostscript](docs/ArticlePostscript.md)
- [ArticleViewerActionState](docs/ArticleViewerActionState.md)
- [BeginEnableTotpResp](docs/BeginEnableTotpResp.md)
- [BindArticleTagsReq](docs/BindArticleTagsReq.md)
- [BlockRelationReq](docs/BlockRelationReq.md)
- [CancelAccountReq](docs/CancelAccountReq.md)
- [CancelPublishArticleReq](docs/CancelPublishArticleReq.md)
- [CollectArticleReq](docs/CollectArticleReq.md)
- [CollectArticleResp](docs/CollectArticleResp.md)
- [ConfirmEnableTotpReq](docs/ConfirmEnableTotpReq.md)
- [CountUnreadNotificationsResp](docs/CountUnreadNotificationsResp.md)
- [CreateCommentReq](docs/CreateCommentReq.md)
- [CreateCommentResp](docs/CreateCommentResp.md)
- [CreateDomainReq](docs/CreateDomainReq.md)
- [CreateDomainResp](docs/CreateDomainResp.md)
- [CreateDraftArticleReq](docs/CreateDraftArticleReq.md)
- [CreateDraftArticleResp](docs/CreateDraftArticleResp.md)
- [CreateTagReq](docs/CreateTagReq.md)
- [CreateTagResp](docs/CreateTagResp.md)
- [DisableTotpReq](docs/DisableTotpReq.md)
- [DiscardDraftArticleReq](docs/DiscardDraftArticleReq.md)
- [FollowRelationReq](docs/FollowRelationReq.md)
- [GetArticleReq](docs/GetArticleReq.md)
- [GetArticleResp](docs/GetArticleResp.md)
- [GetCurrentAccountResp](docs/GetCurrentAccountResp.md)
- [GetCurrentLocationResp](docs/GetCurrentLocationResp.md)
- [GetCurrentPreferencesResp](docs/GetCurrentPreferencesResp.md)
- [GetCurrentPrivacySettingResp](docs/GetCurrentPrivacySettingResp.md)
- [GetCurrentTotpResp](docs/GetCurrentTotpResp.md)
- [GetProfileAccountReq](docs/GetProfileAccountReq.md)
- [GetProfileAccountResp](docs/GetProfileAccountResp.md)
- [GetStatusRelationReq](docs/GetStatusRelationReq.md)
- [GetStatusRelationResp](docs/GetStatusRelationResp.md)
- [ImageResp](docs/ImageResp.md)
- [LikeArticleReq](docs/LikeArticleReq.md)
- [LikeArticleResp](docs/LikeArticleResp.md)
- [LikeCommentReq](docs/LikeCommentReq.md)
- [LikeCommentResp](docs/LikeCommentResp.md)
- [ListArticleTagsReq](docs/ListArticleTagsReq.md)
- [ListArticleTagsResp](docs/ListArticleTagsResp.md)
- [ListArticlesReq](docs/ListArticlesReq.md)
- [ListArticlesResp](docs/ListArticlesResp.md)
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
- [ListFollowersRelationsReq](docs/ListFollowersRelationsReq.md)
- [ListFollowersRelationsResp](docs/ListFollowersRelationsResp.md)
- [ListFollowingRelationsReq](docs/ListFollowingRelationsReq.md)
- [ListFollowingRelationsResp](docs/ListFollowingRelationsResp.md)
- [ListNotificationsReq](docs/ListNotificationsReq.md)
- [ListNotificationsResp](docs/ListNotificationsResp.md)
- [ListPostscriptsReq](docs/ListPostscriptsReq.md)
- [ListPostscriptsResp](docs/ListPostscriptsResp.md)
- [ListTagsReq](docs/ListTagsReq.md)
- [ListTagsResp](docs/ListTagsResp.md)
- [LoginReq](docs/LoginReq.md)
- [LoginResp](docs/LoginResp.md)
- [MarkReadNotificationReq](docs/MarkReadNotificationReq.md)
- [MarkReadNotificationResp](docs/MarkReadNotificationResp.md)
- [PageReq](docs/PageReq.md)
- [PageResp](docs/PageResp.md)
- [PublishArticleReq](docs/PublishArticleReq.md)
- [RefreshTokenReq](docs/RefreshTokenReq.md)
- [RefreshTokenResp](docs/RefreshTokenResp.md)
- [RegisterReq](docs/RegisterReq.md)
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
- [RespDomain](docs/RespDomain.md)
- [RespLocation](docs/RespLocation.md)
- [RespNotification](docs/RespNotification.md)
- [RespPreference](docs/RespPreference.md)
- [RespPrivacySetting](docs/RespPrivacySetting.md)
- [RespRelation](docs/RespRelation.md)
- [RespRelationStatus](docs/RespRelationStatus.md)
- [RespTag](docs/RespTag.md)
- [RespTotp](docs/RespTotp.md)
- [RewardArticleReq](docs/RewardArticleReq.md)
- [SchedulePublishArticleReq](docs/SchedulePublishArticleReq.md)
- [SendEmailOtpReq](docs/SendEmailOtpReq.md)
- [SendEmailOtpResp](docs/SendEmailOtpResp.md)
- [SendPhoneOtpReq](docs/SendPhoneOtpReq.md)
- [SendPhoneOtpResp](docs/SendPhoneOtpResp.md)
- [ThankArticleReq](docs/ThankArticleReq.md)
- [ThankArticleResp](docs/ThankArticleResp.md)
- [ThankCommentReq](docs/ThankCommentReq.md)
- [ThankCommentResp](docs/ThankCommentResp.md)
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
- [UpsertCurrentLocationReq](docs/UpsertCurrentLocationReq.md)
- [UpsertCurrentLocationResp](docs/UpsertCurrentLocationResp.md)

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
