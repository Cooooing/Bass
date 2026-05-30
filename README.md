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
    // string | 用于生成头像的账号名。 (optional)
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
*AccountService* | [**updateProfile**](docs/AccountService.md#updateprofile) | **POST** /v1/user/account/update-profile | 
*ArticleService* | [**_delete**](docs/ArticleService.md#_delete) | **POST** /v1/content/article/delete | 
*ArticleService* | [**acceptAnswer**](docs/ArticleService.md#acceptanswer) | **POST** /v1/content/article/accept-answer | 
*ArticleService* | [**collect**](docs/ArticleService.md#collect) | **POST** /v1/content/article/collect | 
*ArticleService* | [**create**](docs/ArticleService.md#create) | **POST** /v1/content/article/create | 
*ArticleService* | [**get**](docs/ArticleService.md#get) | **POST** /v1/content/article/get | 
*ArticleService* | [**like**](docs/ArticleService.md#like) | **POST** /v1/content/article/like | 
*ArticleService* | [**list**](docs/ArticleService.md#list) | **POST** /v1/content/article/list | 
*ArticleService* | [**publish**](docs/ArticleService.md#publish) | **POST** /v1/content/article/publish | 
*ArticleService* | [**reward**](docs/ArticleService.md#reward) | **POST** /v1/content/article/reward | 
*ArticleService* | [**thank**](docs/ArticleService.md#thank) | **POST** /v1/content/article/thank | 
*ArticleService* | [**updateDraft**](docs/ArticleService.md#updatedraft) | **POST** /v1/content/article/update-draft | 
*ArticleService* | [**watch**](docs/ArticleService.md#watch) | **POST** /v1/content/article/watch | 
*AuthService* | [**loginByPassword**](docs/AuthService.md#loginbypasswordoperation) | **POST** /v1/user/auth/login-by-password | 
*AuthService* | [**logout**](docs/AuthService.md#logout) | **POST** /v1/user/auth/logout | 
*AuthService* | [**startEmailRegistration**](docs/AuthService.md#startemailregistrationoperation) | **POST** /v1/user/auth/start-email-registration | 
*AuthService* | [**startPhoneRegistration**](docs/AuthService.md#startphoneregistrationoperation) | **POST** /v1/user/auth/start-phone-registration | 
*AuthService* | [**verifyEmailRegistration**](docs/AuthService.md#verifyemailregistrationoperation) | **POST** /v1/user/auth/verify-email-registration | 
*AuthService* | [**verifyPhoneRegistration**](docs/AuthService.md#verifyphoneregistrationoperation) | **POST** /v1/user/auth/verify-phone-registration | 
*CommentService* | [**create**](docs/CommentService.md#create) | **POST** /v1/content/comment/create | 
*CommentService* | [**like**](docs/CommentService.md#like) | **POST** /v1/content/comment/like | 
*CommentService* | [**list**](docs/CommentService.md#list) | **POST** /v1/content/comment/list | 
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
*TagService* | [**list**](docs/TagService.md#list) | **POST** /v1/content/tag/list | 
*TotpService* | [**beginEnable**](docs/TotpService.md#beginenable) | **POST** /v1/user/totp/begin-enable | 
*TotpService* | [**confirmEnable**](docs/TotpService.md#confirmenable) | **POST** /v1/user/totp/confirm-enable | 
*TotpService* | [**disable**](docs/TotpService.md#disable) | **POST** /v1/user/totp/disable | 
*TotpService* | [**getCurrent**](docs/TotpService.md#getcurrent) | **POST** /v1/user/totp/get-current | 


### Models

- [AcceptAnswerArticleRequest](docs/AcceptAnswerArticleRequest.md)
- [Account](docs/Account.md)
- [AccountContact](docs/AccountContact.md)
- [AccountProfile](docs/AccountProfile.md)
- [AddPostscriptReply](docs/AddPostscriptReply.md)
- [AddPostscriptRequest](docs/AddPostscriptRequest.md)
- [Any](docs/Any.md)
- [Article](docs/Article.md)
- [ArticlePostscript](docs/ArticlePostscript.md)
- [ArticleQuery](docs/ArticleQuery.md)
- [ArticleSave](docs/ArticleSave.md)
- [BeginEnableTotpReply](docs/BeginEnableTotpReply.md)
- [BlockRelationRequest](docs/BlockRelationRequest.md)
- [CollectArticleRequest](docs/CollectArticleRequest.md)
- [Comment](docs/Comment.md)
- [CommentQuery](docs/CommentQuery.md)
- [ConfirmEnableTotpRequest](docs/ConfirmEnableTotpRequest.md)
- [CountUnreadNotificationsReply](docs/CountUnreadNotificationsReply.md)
- [CreateArticleReply](docs/CreateArticleReply.md)
- [CreateArticleRequest](docs/CreateArticleRequest.md)
- [CreateCommentReply](docs/CreateCommentReply.md)
- [CreateCommentRequest](docs/CreateCommentRequest.md)
- [DeleteArticleRequest](docs/DeleteArticleRequest.md)
- [DisableTotpRequest](docs/DisableTotpRequest.md)
- [Domain](docs/Domain.md)
- [DomainQuery](docs/DomainQuery.md)
- [ExternalDocs](docs/ExternalDocs.md)
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
- [GoogleProtobufAny](docs/GoogleProtobufAny.md)
- [ImageReply](docs/ImageReply.md)
- [LikeArticleRequest](docs/LikeArticleRequest.md)
- [LikeCommentRequest](docs/LikeCommentRequest.md)
- [ListArticlesReply](docs/ListArticlesReply.md)
- [ListArticlesRequest](docs/ListArticlesRequest.md)
- [ListBlockedRelationsReply](docs/ListBlockedRelationsReply.md)
- [ListBlockedRelationsRequest](docs/ListBlockedRelationsRequest.md)
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
- [NamedAny](docs/NamedAny.md)
- [Notification](docs/Notification.md)
- [PageReply](docs/PageReply.md)
- [PageRequest](docs/PageRequest.md)
- [Preference](docs/Preference.md)
- [PrivacySetting](docs/PrivacySetting.md)
- [PublishArticleRequest](docs/PublishArticleRequest.md)
- [Relation](docs/Relation.md)
- [RelationStatus](docs/RelationStatus.md)
- [RewardArticleRequest](docs/RewardArticleRequest.md)
- [StartEmailRegistrationReply](docs/StartEmailRegistrationReply.md)
- [StartEmailRegistrationRequest](docs/StartEmailRegistrationRequest.md)
- [StartPhoneRegistrationReply](docs/StartPhoneRegistrationReply.md)
- [StartPhoneRegistrationRequest](docs/StartPhoneRegistrationRequest.md)
- [Tag](docs/Tag.md)
- [TagQuery](docs/TagQuery.md)
- [TagSave](docs/TagSave.md)
- [ThankArticleRequest](docs/ThankArticleRequest.md)
- [ThankCommentRequest](docs/ThankCommentRequest.md)
- [Totp](docs/Totp.md)
- [UnblockRelationRequest](docs/UnblockRelationRequest.md)
- [UnfollowRelationRequest](docs/UnfollowRelationRequest.md)
- [UpdateCurrentPreferencesReply](docs/UpdateCurrentPreferencesReply.md)
- [UpdateCurrentPreferencesRequest](docs/UpdateCurrentPreferencesRequest.md)
- [UpdateCurrentPrivacySettingReply](docs/UpdateCurrentPrivacySettingReply.md)
- [UpdateCurrentPrivacySettingRequest](docs/UpdateCurrentPrivacySettingRequest.md)
- [UpdateDraftArticleReply](docs/UpdateDraftArticleReply.md)
- [UpdateDraftArticleRequest](docs/UpdateDraftArticleRequest.md)
- [UpdateProfileAccountReply](docs/UpdateProfileAccountReply.md)
- [UpdateProfileAccountRequest](docs/UpdateProfileAccountRequest.md)
- [UpsertCurrentLocationReply](docs/UpsertCurrentLocationReply.md)
- [UpsertCurrentLocationRequest](docs/UpsertCurrentLocationRequest.md)
- [VerifyEmailRegistrationRequest](docs/VerifyEmailRegistrationRequest.md)
- [VerifyPhoneRegistrationRequest](docs/VerifyPhoneRegistrationRequest.md)
- [WatchArticleRequest](docs/WatchArticleRequest.md)

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
