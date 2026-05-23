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
import type { AccountServiceBatchGetProfileRequest } from '@bass/bbs-sdk';

async function example() {
  console.log("🚀 Testing @bass/bbs-sdk SDK...");
  const api = new AccountServiceApi();

  const body = {
    // CommonApiAppBbsV1UserBatchGetProfileAccountRequest
    commonApiAppBbsV1UserBatchGetProfileAccountRequest: ...,
  } satisfies AccountServiceBatchGetProfileRequest;

  try {
    const data = await api.accountServiceBatchGetProfile(body);
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
*AccountServiceApi* | [**accountServiceBatchGetProfile**](docs/AccountServiceApi.md#accountservicebatchgetprofile) | **POST** /v1/user/account/batch-get-profile | 
*AccountServiceApi* | [**accountServiceGetCurrent**](docs/AccountServiceApi.md#accountservicegetcurrent) | **POST** /v1/user/account/get-current | 
*AccountServiceApi* | [**accountServiceGetProfile**](docs/AccountServiceApi.md#accountservicegetprofile) | **POST** /v1/user/account/get-profile | 
*AccountServiceApi* | [**accountServiceUpdateProfile**](docs/AccountServiceApi.md#accountserviceupdateprofile) | **POST** /v1/user/account/update-profile | 
*AuthServiceApi* | [**authServiceLoginPassword**](docs/AuthServiceApi.md#authserviceloginpassword) | **POST** /v1/user/auth/login-password | 
*AuthServiceApi* | [**authServiceLogout**](docs/AuthServiceApi.md#authservicelogout) | **POST** /v1/user/auth/logout | 
*AuthServiceApi* | [**authServiceRegisterEmail**](docs/AuthServiceApi.md#authserviceregisteremail) | **POST** /v1/user/auth/register-email | 
*AuthServiceApi* | [**authServiceRegisterPhone**](docs/AuthServiceApi.md#authserviceregisterphone) | **POST** /v1/user/auth/register-phone | 
*AuthServiceApi* | [**authServiceVerifyEmailRegister**](docs/AuthServiceApi.md#authserviceverifyemailregister) | **POST** /v1/user/auth/verify-email-register | 
*AuthServiceApi* | [**authServiceVerifyPhoneRegister**](docs/AuthServiceApi.md#authserviceverifyphoneregister) | **POST** /v1/user/auth/verify-phone-register | 
*LocationServiceApi* | [**locationServiceGetCurrent**](docs/LocationServiceApi.md#locationservicegetcurrent) | **POST** /v1/user/location/get-current | 
*LocationServiceApi* | [**locationServiceUpsert**](docs/LocationServiceApi.md#locationserviceupsert) | **POST** /v1/user/location/upsert-current | 
*PreferencesServiceApi* | [**preferencesServiceGetCurrent**](docs/PreferencesServiceApi.md#preferencesservicegetcurrent) | **POST** /v1/user/preference/get-current | 
*PreferencesServiceApi* | [**preferencesServiceUpdate**](docs/PreferencesServiceApi.md#preferencesserviceupdate) | **POST** /v1/user/preference/update-current | 
*PrivacySettingServiceApi* | [**privacySettingServiceGetCurrent**](docs/PrivacySettingServiceApi.md#privacysettingservicegetcurrent) | **POST** /v1/user/privacy-setting/get-current | 
*PrivacySettingServiceApi* | [**privacySettingServiceUpdate**](docs/PrivacySettingServiceApi.md#privacysettingserviceupdate) | **POST** /v1/user/privacy-setting/update-current | 
*RelationServiceApi* | [**relationServiceBatchGetStatus**](docs/RelationServiceApi.md#relationservicebatchgetstatus) | **POST** /v1/user/relation/batch-get-status | 
*RelationServiceApi* | [**relationServiceBlock**](docs/RelationServiceApi.md#relationserviceblock) | **POST** /v1/user/relation/block | 
*RelationServiceApi* | [**relationServiceFollow**](docs/RelationServiceApi.md#relationservicefollow) | **POST** /v1/user/relation/follow | 
*RelationServiceApi* | [**relationServicePageBlocked**](docs/RelationServiceApi.md#relationservicepageblocked) | **POST** /v1/user/relation/page-blocked | 
*RelationServiceApi* | [**relationServicePageFollowers**](docs/RelationServiceApi.md#relationservicepagefollowers) | **POST** /v1/user/relation/page-followers | 
*RelationServiceApi* | [**relationServicePageFollowing**](docs/RelationServiceApi.md#relationservicepagefollowing) | **POST** /v1/user/relation/page-following | 
*RelationServiceApi* | [**relationServiceUnblock**](docs/RelationServiceApi.md#relationserviceunblock) | **POST** /v1/user/relation/unblock | 
*RelationServiceApi* | [**relationServiceUnfollow**](docs/RelationServiceApi.md#relationserviceunfollow) | **POST** /v1/user/relation/unfollow | 
*TfaServiceApi* | [**tfaServiceBeginEnable**](docs/TfaServiceApi.md#tfaservicebeginenable) | **POST** /v1/user/tfa/begin-enable | 
*TfaServiceApi* | [**tfaServiceConfirmEnable**](docs/TfaServiceApi.md#tfaserviceconfirmenable) | **POST** /v1/user/tfa/confirm-enable | 
*TfaServiceApi* | [**tfaServiceDisable**](docs/TfaServiceApi.md#tfaservicedisable) | **POST** /v1/user/tfa/disable | 
*TfaServiceApi* | [**tfaServiceGetCurrent**](docs/TfaServiceApi.md#tfaservicegetcurrent) | **POST** /v1/user/tfa/get-current | 
*TfaServiceApi* | [**tfaServiceValidate**](docs/TfaServiceApi.md#tfaservicevalidate) | **POST** /v1/user/tfa/validate | 


### Models

- [CommonApiAppBbsV1UserAccount](docs/CommonApiAppBbsV1UserAccount.md)
- [CommonApiAppBbsV1UserAccountContact](docs/CommonApiAppBbsV1UserAccountContact.md)
- [CommonApiAppBbsV1UserAccountProfile](docs/CommonApiAppBbsV1UserAccountProfile.md)
- [CommonApiAppBbsV1UserBatchGetProfileAccountReply](docs/CommonApiAppBbsV1UserBatchGetProfileAccountReply.md)
- [CommonApiAppBbsV1UserBatchGetProfileAccountRequest](docs/CommonApiAppBbsV1UserBatchGetProfileAccountRequest.md)
- [CommonApiAppBbsV1UserBatchGetStatusRelationReply](docs/CommonApiAppBbsV1UserBatchGetStatusRelationReply.md)
- [CommonApiAppBbsV1UserBatchGetStatusRelationRequest](docs/CommonApiAppBbsV1UserBatchGetStatusRelationRequest.md)
- [CommonApiAppBbsV1UserBeginEnableTfaReply](docs/CommonApiAppBbsV1UserBeginEnableTfaReply.md)
- [CommonApiAppBbsV1UserBlockRelationRequest](docs/CommonApiAppBbsV1UserBlockRelationRequest.md)
- [CommonApiAppBbsV1UserConfirmEnableTfaRequest](docs/CommonApiAppBbsV1UserConfirmEnableTfaRequest.md)
- [CommonApiAppBbsV1UserDisableTfaRequest](docs/CommonApiAppBbsV1UserDisableTfaRequest.md)
- [CommonApiAppBbsV1UserFollowRelationRequest](docs/CommonApiAppBbsV1UserFollowRelationRequest.md)
- [CommonApiAppBbsV1UserGetCurrentAccountReply](docs/CommonApiAppBbsV1UserGetCurrentAccountReply.md)
- [CommonApiAppBbsV1UserGetCurrentLocationReply](docs/CommonApiAppBbsV1UserGetCurrentLocationReply.md)
- [CommonApiAppBbsV1UserGetCurrentPreferencesReply](docs/CommonApiAppBbsV1UserGetCurrentPreferencesReply.md)
- [CommonApiAppBbsV1UserGetCurrentPrivacySettingReply](docs/CommonApiAppBbsV1UserGetCurrentPrivacySettingReply.md)
- [CommonApiAppBbsV1UserGetCurrentTfaReply](docs/CommonApiAppBbsV1UserGetCurrentTfaReply.md)
- [CommonApiAppBbsV1UserGetProfileAccountReply](docs/CommonApiAppBbsV1UserGetProfileAccountReply.md)
- [CommonApiAppBbsV1UserGetProfileAccountRequest](docs/CommonApiAppBbsV1UserGetProfileAccountRequest.md)
- [CommonApiAppBbsV1UserLocation](docs/CommonApiAppBbsV1UserLocation.md)
- [CommonApiAppBbsV1UserLoginPasswordReply](docs/CommonApiAppBbsV1UserLoginPasswordReply.md)
- [CommonApiAppBbsV1UserLoginPasswordRequest](docs/CommonApiAppBbsV1UserLoginPasswordRequest.md)
- [CommonApiAppBbsV1UserPageBlockedRelationReply](docs/CommonApiAppBbsV1UserPageBlockedRelationReply.md)
- [CommonApiAppBbsV1UserPageBlockedRelationRequest](docs/CommonApiAppBbsV1UserPageBlockedRelationRequest.md)
- [CommonApiAppBbsV1UserPageFollowersRelationReply](docs/CommonApiAppBbsV1UserPageFollowersRelationReply.md)
- [CommonApiAppBbsV1UserPageFollowersRelationRequest](docs/CommonApiAppBbsV1UserPageFollowersRelationRequest.md)
- [CommonApiAppBbsV1UserPageFollowingRelationReply](docs/CommonApiAppBbsV1UserPageFollowingRelationReply.md)
- [CommonApiAppBbsV1UserPageFollowingRelationRequest](docs/CommonApiAppBbsV1UserPageFollowingRelationRequest.md)
- [CommonApiAppBbsV1UserPageReply](docs/CommonApiAppBbsV1UserPageReply.md)
- [CommonApiAppBbsV1UserPageRequest](docs/CommonApiAppBbsV1UserPageRequest.md)
- [CommonApiAppBbsV1UserPreference](docs/CommonApiAppBbsV1UserPreference.md)
- [CommonApiAppBbsV1UserPrivacySetting](docs/CommonApiAppBbsV1UserPrivacySetting.md)
- [CommonApiAppBbsV1UserRegisterEmailReply](docs/CommonApiAppBbsV1UserRegisterEmailReply.md)
- [CommonApiAppBbsV1UserRegisterEmailRequest](docs/CommonApiAppBbsV1UserRegisterEmailRequest.md)
- [CommonApiAppBbsV1UserRegisterPhoneReply](docs/CommonApiAppBbsV1UserRegisterPhoneReply.md)
- [CommonApiAppBbsV1UserRegisterPhoneRequest](docs/CommonApiAppBbsV1UserRegisterPhoneRequest.md)
- [CommonApiAppBbsV1UserRelation](docs/CommonApiAppBbsV1UserRelation.md)
- [CommonApiAppBbsV1UserRelationStatus](docs/CommonApiAppBbsV1UserRelationStatus.md)
- [CommonApiAppBbsV1UserTfa](docs/CommonApiAppBbsV1UserTfa.md)
- [CommonApiAppBbsV1UserUnblockRelationRequest](docs/CommonApiAppBbsV1UserUnblockRelationRequest.md)
- [CommonApiAppBbsV1UserUnfollowRelationRequest](docs/CommonApiAppBbsV1UserUnfollowRelationRequest.md)
- [CommonApiAppBbsV1UserUpdatePreferencesReply](docs/CommonApiAppBbsV1UserUpdatePreferencesReply.md)
- [CommonApiAppBbsV1UserUpdatePreferencesRequest](docs/CommonApiAppBbsV1UserUpdatePreferencesRequest.md)
- [CommonApiAppBbsV1UserUpdatePrivacySettingReply](docs/CommonApiAppBbsV1UserUpdatePrivacySettingReply.md)
- [CommonApiAppBbsV1UserUpdatePrivacySettingRequest](docs/CommonApiAppBbsV1UserUpdatePrivacySettingRequest.md)
- [CommonApiAppBbsV1UserUpdateProfileAccountReply](docs/CommonApiAppBbsV1UserUpdateProfileAccountReply.md)
- [CommonApiAppBbsV1UserUpdateProfileAccountRequest](docs/CommonApiAppBbsV1UserUpdateProfileAccountRequest.md)
- [CommonApiAppBbsV1UserUpsertLocationReply](docs/CommonApiAppBbsV1UserUpsertLocationReply.md)
- [CommonApiAppBbsV1UserUpsertLocationRequest](docs/CommonApiAppBbsV1UserUpsertLocationRequest.md)
- [CommonApiAppBbsV1UserValidateTfaReply](docs/CommonApiAppBbsV1UserValidateTfaReply.md)
- [CommonApiAppBbsV1UserValidateTfaRequest](docs/CommonApiAppBbsV1UserValidateTfaRequest.md)
- [CommonApiAppBbsV1UserVerifyEmailRegisterRequest](docs/CommonApiAppBbsV1UserVerifyEmailRegisterRequest.md)
- [CommonApiAppBbsV1UserVerifyPhoneRegisterRequest](docs/CommonApiAppBbsV1UserVerifyPhoneRegisterRequest.md)

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
