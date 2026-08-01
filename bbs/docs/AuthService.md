# AuthService

All URIs are relative to *http://localhost*

|Method | HTTP request | Description|
|------------- | ------------- | -------------|
|[**cancelAccount**](#cancelaccount) | **POST** /v1/user/auth/cancel-account | |
|[**login**](#login) | **POST** /v1/user/auth/login | |
|[**logout**](#logout) | **POST** /v1/user/auth/logout | |
|[**refreshToken**](#refreshtoken) | **POST** /v1/user/auth/refresh-token | |
|[**startEmailRegistration**](#startemailregistration) | **POST** /v1/user/auth/start-email-registration | |
|[**startPhoneRegistration**](#startphoneregistration) | **POST** /v1/user/auth/start-phone-registration | |
|[**verifyEmailRegistration**](#verifyemailregistration) | **POST** /v1/user/auth/verify-email-registration | |
|[**verifyPhoneRegistration**](#verifyphoneregistration) | **POST** /v1/user/auth/verify-phone-registration | |

# **cancelAccount**
> object cancelAccount(cancelAccountReq)

注销账号。

### Example

```typescript
import {
    AuthService,
    Configuration,
    CancelAccountReq
} from '@bass/bbs-sdk-axios';

const configuration = new Configuration();
const apiInstance = new AuthService(configuration);

let cancelAccountReq: CancelAccountReq; //

const { status, data } = await apiInstance.cancelAccount(
    cancelAccountReq
);
```

### Parameters

|Name | Type | Description  | Notes|
|------------- | ------------- | ------------- | -------------|
| **cancelAccountReq** | **CancelAccountReq**|  | |


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

# **login**
> LoginResp login(loginReq)

登录账号。

### Example

```typescript
import {
    AuthService,
    Configuration,
    LoginReq
} from '@bass/bbs-sdk-axios';

const configuration = new Configuration();
const apiInstance = new AuthService(configuration);

let loginReq: LoginReq; //

const { status, data } = await apiInstance.login(
    loginReq
);
```

### Parameters

|Name | Type | Description  | Notes|
|------------- | ------------- | ------------- | -------------|
| **loginReq** | **LoginReq**|  | |


### Return type

**LoginResp**

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

退出登录。

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

# **refreshToken**
> RefreshTokenResp refreshToken(refreshTokenReq)

刷新登录令牌。

### Example

```typescript
import {
    AuthService,
    Configuration,
    RefreshTokenReq
} from '@bass/bbs-sdk-axios';

const configuration = new Configuration();
const apiInstance = new AuthService(configuration);

let refreshTokenReq: RefreshTokenReq; //

const { status, data } = await apiInstance.refreshToken(
    refreshTokenReq
);
```

### Parameters

|Name | Type | Description  | Notes|
|------------- | ------------- | ------------- | -------------|
| **refreshTokenReq** | **RefreshTokenReq**|  | |


### Return type

**RefreshTokenResp**

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
> StartEmailRegistrationResp startEmailRegistration(startEmailRegistrationReq)

开始邮箱注册。

### Example

```typescript
import {
    AuthService,
    Configuration,
    StartEmailRegistrationReq
} from '@bass/bbs-sdk-axios';

const configuration = new Configuration();
const apiInstance = new AuthService(configuration);

let startEmailRegistrationReq: StartEmailRegistrationReq; //

const { status, data } = await apiInstance.startEmailRegistration(
    startEmailRegistrationReq
);
```

### Parameters

|Name | Type | Description  | Notes|
|------------- | ------------- | ------------- | -------------|
| **startEmailRegistrationReq** | **StartEmailRegistrationReq**|  | |


### Return type

**StartEmailRegistrationResp**

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
> StartPhoneRegistrationResp startPhoneRegistration(startPhoneRegistrationReq)

开始手机注册。

### Example

```typescript
import {
    AuthService,
    Configuration,
    StartPhoneRegistrationReq
} from '@bass/bbs-sdk-axios';

const configuration = new Configuration();
const apiInstance = new AuthService(configuration);

let startPhoneRegistrationReq: StartPhoneRegistrationReq; //

const { status, data } = await apiInstance.startPhoneRegistration(
    startPhoneRegistrationReq
);
```

### Parameters

|Name | Type | Description  | Notes|
|------------- | ------------- | ------------- | -------------|
| **startPhoneRegistrationReq** | **StartPhoneRegistrationReq**|  | |


### Return type

**StartPhoneRegistrationResp**

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
> object verifyEmailRegistration(verifyEmailRegistrationReq)

校验邮箱注册验证码。

### Example

```typescript
import {
    AuthService,
    Configuration,
    VerifyEmailRegistrationReq
} from '@bass/bbs-sdk-axios';

const configuration = new Configuration();
const apiInstance = new AuthService(configuration);

let verifyEmailRegistrationReq: VerifyEmailRegistrationReq; //

const { status, data } = await apiInstance.verifyEmailRegistration(
    verifyEmailRegistrationReq
);
```

### Parameters

|Name | Type | Description  | Notes|
|------------- | ------------- | ------------- | -------------|
| **verifyEmailRegistrationReq** | **VerifyEmailRegistrationReq**|  | |


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
> object verifyPhoneRegistration(verifyPhoneRegistrationReq)

校验手机注册验证码。

### Example

```typescript
import {
    AuthService,
    Configuration,
    VerifyPhoneRegistrationReq
} from '@bass/bbs-sdk-axios';

const configuration = new Configuration();
const apiInstance = new AuthService(configuration);

let verifyPhoneRegistrationReq: VerifyPhoneRegistrationReq; //

const { status, data } = await apiInstance.verifyPhoneRegistration(
    verifyPhoneRegistrationReq
);
```

### Parameters

|Name | Type | Description  | Notes|
|------------- | ------------- | ------------- | -------------|
| **verifyPhoneRegistrationReq** | **VerifyPhoneRegistrationReq**|  | |


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

