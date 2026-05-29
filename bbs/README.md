# @bass/bbs-sdk@0.1.0

A TypeScript SDK client for the localhost API.

## Usage

First, install the SDK from npm.

```bash
npm install @bass/bbs-sdk --save
```

Next, try it out.


```ts
import {
  Configuration,
  AccountServiceApi,
} from '@bass/bbs-sdk';
import type { AccountServiceGetCurrentRequest } from '@bass/bbs-sdk';

async function example() {
  console.log("🚀 Testing @bass/bbs-sdk SDK...");
  const api = new AccountServiceApi();

  const body = {
    // object
    body: Object,
  } satisfies AccountServiceGetCurrentRequest;

  try {
    const data = await api.accountServiceGetCurrent(body);
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
*AccountServiceApi* | [**accountServiceGetCurrent**](docs/AccountServiceApi.md#accountservicegetcurrent) | **POST** /v1/user/account/get-current | 
*AccountServiceApi* | [**accountServiceGetProfile**](docs/AccountServiceApi.md#accountservicegetprofile) | **POST** /v1/user/account/get-profile | 
*AccountServiceApi* | [**accountServiceUpdateProfile**](docs/AccountServiceApi.md#accountserviceupdateprofile) | **POST** /v1/user/account/update-profile | 
*ArticleServiceApi* | [**articleServiceAcceptAnswer**](docs/ArticleServiceApi.md#articleserviceacceptanswer) | **POST** /v1/content/article/accept-answer | 
*ArticleServiceApi* | [**articleServiceCollect**](docs/ArticleServiceApi.md#articleservicecollect) | **POST** /v1/content/article/collect | 
*ArticleServiceApi* | [**articleServiceCreate**](docs/ArticleServiceApi.md#articleservicecreate) | **POST** /v1/content/article/create | 
*ArticleServiceApi* | [**articleServiceDelete**](docs/ArticleServiceApi.md#articleservicedelete) | **POST** /v1/content/article/delete | 
*ArticleServiceApi* | [**articleServiceGet**](docs/ArticleServiceApi.md#articleserviceget) | **POST** /v1/content/article/get | 
*ArticleServiceApi* | [**articleServiceLike**](docs/ArticleServiceApi.md#articleservicelike) | **POST** /v1/content/article/like | 
*ArticleServiceApi* | [**articleServiceList**](docs/ArticleServiceApi.md#articleservicelist) | **POST** /v1/content/article/list | 
*ArticleServiceApi* | [**articleServicePublish**](docs/ArticleServiceApi.md#articleservicepublish) | **POST** /v1/content/article/publish | 
*ArticleServiceApi* | [**articleServiceReward**](docs/ArticleServiceApi.md#articleservicereward) | **POST** /v1/content/article/reward | 
*ArticleServiceApi* | [**articleServiceThank**](docs/ArticleServiceApi.md#articleservicethank) | **POST** /v1/content/article/thank | 
*ArticleServiceApi* | [**articleServiceUpdateDraft**](docs/ArticleServiceApi.md#articleserviceupdatedraft) | **POST** /v1/content/article/update-draft | 
*ArticleServiceApi* | [**articleServiceWatch**](docs/ArticleServiceApi.md#articleservicewatch) | **POST** /v1/content/article/watch | 
*AuthServiceApi* | [**authServiceLoginByPassword**](docs/AuthServiceApi.md#authserviceloginbypassword) | **POST** /v1/user/auth/login-by-password | 
*AuthServiceApi* | [**authServiceLogout**](docs/AuthServiceApi.md#authservicelogout) | **POST** /v1/user/auth/logout | 
*AuthServiceApi* | [**authServiceStartEmailRegistration**](docs/AuthServiceApi.md#authservicestartemailregistration) | **POST** /v1/user/auth/start-email-registration | 
*AuthServiceApi* | [**authServiceStartPhoneRegistration**](docs/AuthServiceApi.md#authservicestartphoneregistration) | **POST** /v1/user/auth/start-phone-registration | 
*AuthServiceApi* | [**authServiceVerifyEmailRegistration**](docs/AuthServiceApi.md#authserviceverifyemailregistration) | **POST** /v1/user/auth/verify-email-registration | 
*AuthServiceApi* | [**authServiceVerifyPhoneRegistration**](docs/AuthServiceApi.md#authserviceverifyphoneregistration) | **POST** /v1/user/auth/verify-phone-registration | 
*CommentServiceApi* | [**commentServiceCreate**](docs/CommentServiceApi.md#commentservicecreate) | **POST** /v1/content/comment/create | 
*CommentServiceApi* | [**commentServiceLike**](docs/CommentServiceApi.md#commentservicelike) | **POST** /v1/content/comment/like | 
*CommentServiceApi* | [**commentServiceList**](docs/CommentServiceApi.md#commentservicelist) | **POST** /v1/content/comment/list | 
*CommentServiceApi* | [**commentServiceThank**](docs/CommentServiceApi.md#commentservicethank) | **POST** /v1/content/comment/thank | 
*DomainServiceApi* | [**domainServiceList**](docs/DomainServiceApi.md#domainservicelist) | **POST** /v1/content/domain/list | 
*LocationServiceApi* | [**locationServiceGetCurrent**](docs/LocationServiceApi.md#locationservicegetcurrent) | **POST** /v1/user/location/get-current | 
*LocationServiceApi* | [**locationServiceUpsertCurrent**](docs/LocationServiceApi.md#locationserviceupsertcurrent) | **POST** /v1/user/location/upsert-current | 
*NotificationServiceApi* | [**notificationServiceCountUnread**](docs/NotificationServiceApi.md#notificationservicecountunread) | **POST** /v1/notify/notification/count-unread | 
*NotificationServiceApi* | [**notificationServiceList**](docs/NotificationServiceApi.md#notificationservicelist) | **POST** /v1/notify/notification/list | 
*NotificationServiceApi* | [**notificationServiceMarkRead**](docs/NotificationServiceApi.md#notificationservicemarkread) | **POST** /v1/notify/notification/mark-read | 
*PostscriptServiceApi* | [**postscriptServiceAdd**](docs/PostscriptServiceApi.md#postscriptserviceadd) | **POST** /v1/content/postscript/add | 
*PreferencesServiceApi* | [**preferencesServiceGetCurrent**](docs/PreferencesServiceApi.md#preferencesservicegetcurrent) | **POST** /v1/user/preference/get-current | 
*PreferencesServiceApi* | [**preferencesServiceUpdateCurrent**](docs/PreferencesServiceApi.md#preferencesserviceupdatecurrent) | **POST** /v1/user/preference/update-current | 
*PrivacySettingServiceApi* | [**privacySettingServiceGetCurrent**](docs/PrivacySettingServiceApi.md#privacysettingservicegetcurrent) | **POST** /v1/user/privacy-setting/get-current | 
*PrivacySettingServiceApi* | [**privacySettingServiceUpdateCurrent**](docs/PrivacySettingServiceApi.md#privacysettingserviceupdatecurrent) | **POST** /v1/user/privacy-setting/update-current | 
*RelationServiceApi* | [**relationServiceBlock**](docs/RelationServiceApi.md#relationserviceblock) | **POST** /v1/user/relation/block | 
*RelationServiceApi* | [**relationServiceFollow**](docs/RelationServiceApi.md#relationservicefollow) | **POST** /v1/user/relation/follow | 
*RelationServiceApi* | [**relationServiceGetStatus**](docs/RelationServiceApi.md#relationservicegetstatus) | **POST** /v1/user/relation/get-status | 
*RelationServiceApi* | [**relationServiceListBlocked**](docs/RelationServiceApi.md#relationservicelistblocked) | **POST** /v1/user/relation/list-blocked | 
*RelationServiceApi* | [**relationServiceListFollowers**](docs/RelationServiceApi.md#relationservicelistfollowers) | **POST** /v1/user/relation/list-followers | 
*RelationServiceApi* | [**relationServiceListFollowing**](docs/RelationServiceApi.md#relationservicelistfollowing) | **POST** /v1/user/relation/list-following | 
*RelationServiceApi* | [**relationServiceUnblock**](docs/RelationServiceApi.md#relationserviceunblock) | **POST** /v1/user/relation/unblock | 
*RelationServiceApi* | [**relationServiceUnfollow**](docs/RelationServiceApi.md#relationserviceunfollow) | **POST** /v1/user/relation/unfollow | 
*TagServiceApi* | [**tagServiceList**](docs/TagServiceApi.md#tagservicelist) | **POST** /v1/content/tag/list | 
*TfaServiceApi* | [**tfaServiceBeginEnable**](docs/TfaServiceApi.md#tfaservicebeginenable) | **POST** /v1/user/tfa/begin-enable | 
*TfaServiceApi* | [**tfaServiceConfirmEnable**](docs/TfaServiceApi.md#tfaserviceconfirmenable) | **POST** /v1/user/tfa/confirm-enable | 
*TfaServiceApi* | [**tfaServiceDisable**](docs/TfaServiceApi.md#tfaservicedisable) | **POST** /v1/user/tfa/disable | 
*TfaServiceApi* | [**tfaServiceGetCurrent**](docs/TfaServiceApi.md#tfaservicegetcurrent) | **POST** /v1/user/tfa/get-current | 
*TfaServiceApi* | [**tfaServiceValidate**](docs/TfaServiceApi.md#tfaservicevalidate) | **POST** /v1/user/tfa/validate | 


### Models

- [AcceptAnswerArticleRequest](docs/AcceptAnswerArticleRequest.md)
- [Account](docs/Account.md)
- [AccountContact](docs/AccountContact.md)
- [AccountProfile](docs/AccountProfile.md)
- [AddPostscriptReply](docs/AddPostscriptReply.md)
- [AddPostscriptRequest](docs/AddPostscriptRequest.md)
- [Article](docs/Article.md)
- [ArticlePostscript](docs/ArticlePostscript.md)
- [ArticleQuery](docs/ArticleQuery.md)
- [ArticleSave](docs/ArticleSave.md)
- [BeginEnableTfaReply](docs/BeginEnableTfaReply.md)
- [BlockRelationRequest](docs/BlockRelationRequest.md)
- [CollectArticleRequest](docs/CollectArticleRequest.md)
- [Comment](docs/Comment.md)
- [CommentQuery](docs/CommentQuery.md)
- [ConfirmEnableTfaRequest](docs/ConfirmEnableTfaRequest.md)
- [CountUnreadNotificationsReply](docs/CountUnreadNotificationsReply.md)
- [CreateArticleReply](docs/CreateArticleReply.md)
- [CreateArticleRequest](docs/CreateArticleRequest.md)
- [CreateCommentReply](docs/CreateCommentReply.md)
- [CreateCommentRequest](docs/CreateCommentRequest.md)
- [DeleteArticleRequest](docs/DeleteArticleRequest.md)
- [DisableTfaRequest](docs/DisableTfaRequest.md)
- [Domain](docs/Domain.md)
- [DomainQuery](docs/DomainQuery.md)
- [FollowRelationRequest](docs/FollowRelationRequest.md)
- [GetArticleReply](docs/GetArticleReply.md)
- [GetArticleRequest](docs/GetArticleRequest.md)
- [GetCurrentAccountReply](docs/GetCurrentAccountReply.md)
- [GetCurrentLocationReply](docs/GetCurrentLocationReply.md)
- [GetCurrentPreferencesReply](docs/GetCurrentPreferencesReply.md)
- [GetCurrentPrivacySettingReply](docs/GetCurrentPrivacySettingReply.md)
- [GetCurrentTfaReply](docs/GetCurrentTfaReply.md)
- [GetProfileAccountReply](docs/GetProfileAccountReply.md)
- [GetProfileAccountRequest](docs/GetProfileAccountRequest.md)
- [GetStatusRelationReply](docs/GetStatusRelationReply.md)
- [GetStatusRelationRequest](docs/GetStatusRelationRequest.md)
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
- [Tfa](docs/Tfa.md)
- [ThankArticleRequest](docs/ThankArticleRequest.md)
- [ThankCommentRequest](docs/ThankCommentRequest.md)
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
- [ValidateTfaReply](docs/ValidateTfaReply.md)
- [ValidateTfaRequest](docs/ValidateTfaRequest.md)
- [VerifyEmailRegistrationRequest](docs/VerifyEmailRegistrationRequest.md)
- [VerifyPhoneRegistrationRequest](docs/VerifyPhoneRegistrationRequest.md)
- [WatchArticleRequest](docs/WatchArticleRequest.md)

### Authorization

Endpoints do not require authorization.


## About

This TypeScript SDK client supports the [Fetch API](https://fetch.spec.whatwg.org/)
and is automatically generated by the
[OpenAPI Generator](https://openapi-generator.tech) project:

- API version: `0.0.1`
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
