# AuthService

All URIs are relative to *http://localhost*

|Method | HTTP request | Description|
|------------- | ------------- | -------------|
|[**loginByPassword**](#loginbypassword) | **POST** /v1/user/auth/login-by-password | |
|[**logout**](#logout) | **POST** /v1/user/auth/logout | |
|[**startEmailRegistration**](#startemailregistration) | **POST** /v1/user/auth/start-email-registration | |
|[**startPhoneRegistration**](#startphoneregistration) | **POST** /v1/user/auth/start-phone-registration | |
|[**verifyEmailRegistration**](#verifyemailregistration) | **POST** /v1/user/auth/verify-email-registration | |
|[**verifyPhoneRegistration**](#verifyphoneregistration) | **POST** /v1/user/auth/verify-phone-registration | |

# **loginByPassword**
> LoginByPasswordReply loginByPassword(loginByPasswordRequest)

使用密码登录账号。

### Example

```typescript
import {
    AuthService,
    Configuration,
    LoginByPasswordRequest
} from '@bass/bbs-sdk-axios';

const configuration = new Configuration();
const apiInstance = new AuthService(configuration);

let loginByPasswordRequest: LoginByPasswordRequest; //

const { status, data } = await apiInstance.loginByPassword(
    loginByPasswordRequest
);
```

### Parameters

|Name | Type | Description  | Notes|
|------------- | ------------- | ------------- | -------------|
| **loginByPasswordRequest** | **LoginByPasswordRequest**|  | |


### Return type

**LoginByPasswordReply**

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json


### HTTP response details
| Status code | Description | Response headers |
|-------------|-------------|------------------|
|**200** | OK |  -  |

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **logout**
> object logout(body)

登出当前账号。

### Example

```typescript
import {
    AuthService,
    Configuration
} from '@bass/bbs-sdk-axios';

const configuration = new Configuration();
const apiInstance = new AuthService(configuration);

let body: object; //

const { status, data } = await apiInstance.logout(
    body
);
```

### Parameters

|Name | Type | Description  | Notes|
|------------- | ------------- | ------------- | -------------|
| **body** | **object**|  | |


### Return type

**object**

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json


### HTTP response details
| Status code | Description | Response headers |
|-------------|-------------|------------------|
|**200** | OK |  -  |

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **startEmailRegistration**
> StartEmailRegistrationReply startEmailRegistration(startEmailRegistrationRequest)

使用邮箱发起账号注册。

### Example

```typescript
import {
    AuthService,
    Configuration,
    StartEmailRegistrationRequest
} from '@bass/bbs-sdk-axios';

const configuration = new Configuration();
const apiInstance = new AuthService(configuration);

let startEmailRegistrationRequest: StartEmailRegistrationRequest; //

const { status, data } = await apiInstance.startEmailRegistration(
    startEmailRegistrationRequest
);
```

### Parameters

|Name | Type | Description  | Notes|
|------------- | ------------- | ------------- | -------------|
| **startEmailRegistrationRequest** | **StartEmailRegistrationRequest**|  | |


### Return type

**StartEmailRegistrationReply**

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json


### HTTP response details
| Status code | Description | Response headers |
|-------------|-------------|------------------|
|**200** | OK |  -  |

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **startPhoneRegistration**
> StartPhoneRegistrationReply startPhoneRegistration(startPhoneRegistrationRequest)

使用手机号发起账号注册。

### Example

```typescript
import {
    AuthService,
    Configuration,
    StartPhoneRegistrationRequest
} from '@bass/bbs-sdk-axios';

const configuration = new Configuration();
const apiInstance = new AuthService(configuration);

let startPhoneRegistrationRequest: StartPhoneRegistrationRequest; //

const { status, data } = await apiInstance.startPhoneRegistration(
    startPhoneRegistrationRequest
);
```

### Parameters

|Name | Type | Description  | Notes|
|------------- | ------------- | ------------- | -------------|
| **startPhoneRegistrationRequest** | **StartPhoneRegistrationRequest**|  | |


### Return type

**StartPhoneRegistrationReply**

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json


### HTTP response details
| Status code | Description | Response headers |
|-------------|-------------|------------------|
|**200** | OK |  -  |

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **verifyEmailRegistration**
> object verifyEmailRegistration(verifyEmailRegistrationRequest)

校验邮箱注册验证码。

### Example

```typescript
import {
    AuthService,
    Configuration,
    VerifyEmailRegistrationRequest
} from '@bass/bbs-sdk-axios';

const configuration = new Configuration();
const apiInstance = new AuthService(configuration);

let verifyEmailRegistrationRequest: VerifyEmailRegistrationRequest; //

const { status, data } = await apiInstance.verifyEmailRegistration(
    verifyEmailRegistrationRequest
);
```

### Parameters

|Name | Type | Description  | Notes|
|------------- | ------------- | ------------- | -------------|
| **verifyEmailRegistrationRequest** | **VerifyEmailRegistrationRequest**|  | |


### Return type

**object**

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json


### HTTP response details
| Status code | Description | Response headers |
|-------------|-------------|------------------|
|**200** | OK |  -  |

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **verifyPhoneRegistration**
> object verifyPhoneRegistration(verifyPhoneRegistrationRequest)

校验手机号注册验证码。

### Example

```typescript
import {
    AuthService,
    Configuration,
    VerifyPhoneRegistrationRequest
} from '@bass/bbs-sdk-axios';

const configuration = new Configuration();
const apiInstance = new AuthService(configuration);

let verifyPhoneRegistrationRequest: VerifyPhoneRegistrationRequest; //

const { status, data } = await apiInstance.verifyPhoneRegistration(
    verifyPhoneRegistrationRequest
);
```

### Parameters

|Name | Type | Description  | Notes|
|------------- | ------------- | ------------- | -------------|
| **verifyPhoneRegistrationRequest** | **VerifyPhoneRegistrationRequest**|  | |


### Return type

**object**

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json


### HTTP response details
| Status code | Description | Response headers |
|-------------|-------------|------------------|
|**200** | OK |  -  |

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

