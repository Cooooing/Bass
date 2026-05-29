# AuthService

All URIs are relative to *http://localhost*

| Method | HTTP request | Description |
|------------- | ------------- | -------------|
| [**loginByPassword**](AuthService.md#loginbypasswordoperation) | **POST** /v1/user/auth/login-by-password |  |
| [**logout**](AuthService.md#logout) | **POST** /v1/user/auth/logout |  |
| [**startEmailRegistration**](AuthService.md#startemailregistrationoperation) | **POST** /v1/user/auth/start-email-registration |  |
| [**startPhoneRegistration**](AuthService.md#startphoneregistrationoperation) | **POST** /v1/user/auth/start-phone-registration |  |
| [**verifyEmailRegistration**](AuthService.md#verifyemailregistrationoperation) | **POST** /v1/user/auth/verify-email-registration |  |
| [**verifyPhoneRegistration**](AuthService.md#verifyphoneregistrationoperation) | **POST** /v1/user/auth/verify-phone-registration |  |



## loginByPassword

> LoginByPasswordReply loginByPassword(loginByPasswordRequest)



使用密码登录账号

### Example

```ts
import {
  Configuration,
  AuthService,
} from '@bass/bbs-sdk-fetch';
import type { LoginByPasswordOperationRequest } from '@bass/bbs-sdk-fetch';

async function example() {
  console.log("🚀 Testing @bass/bbs-sdk-fetch SDK...");
  const api = new AuthService();

  const body = {
    // LoginByPasswordRequest
    loginByPasswordRequest: ...,
  } satisfies LoginByPasswordOperationRequest;

  try {
    const data = await api.loginByPassword(body);
    console.log(data);
  } catch (error) {
    console.error(error);
  }
}

// Run the test
example().catch(console.error);
```

### Parameters


| Name | Type | Description  | Notes |
|------------- | ------------- | ------------- | -------------|
| **loginByPasswordRequest** | [LoginByPasswordRequest](LoginByPasswordRequest.md) |  | |

### Return type

[**LoginByPasswordReply**](LoginByPasswordReply.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: `application/json`
- **Accept**: `application/json`


### HTTP response details
| Status code | Description | Response headers |
|-------------|-------------|------------------|
| **200** | OK |  -  |

[[Back to top]](#) [[Back to API list]](../README.md#api-endpoints) [[Back to Model list]](../README.md#models) [[Back to README]](../README.md)


## logout

> object logout(body)



登出当前登录账号

### Example

```ts
import {
  Configuration,
  AuthService,
} from '@bass/bbs-sdk-fetch';
import type { LogoutRequest } from '@bass/bbs-sdk-fetch';

async function example() {
  console.log("🚀 Testing @bass/bbs-sdk-fetch SDK...");
  const api = new AuthService();

  const body = {
    // object
    body: Object,
  } satisfies LogoutRequest;

  try {
    const data = await api.logout(body);
    console.log(data);
  } catch (error) {
    console.error(error);
  }
}

// Run the test
example().catch(console.error);
```

### Parameters


| Name | Type | Description  | Notes |
|------------- | ------------- | ------------- | -------------|
| **body** | `object` |  | |

### Return type

**object**

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: `application/json`
- **Accept**: `application/json`


### HTTP response details
| Status code | Description | Response headers |
|-------------|-------------|------------------|
| **200** | OK |  -  |

[[Back to top]](#) [[Back to API list]](../README.md#api-endpoints) [[Back to Model list]](../README.md#models) [[Back to README]](../README.md)


## startEmailRegistration

> StartEmailRegistrationReply startEmailRegistration(startEmailRegistrationRequest)



使用邮箱发起账号注册

### Example

```ts
import {
  Configuration,
  AuthService,
} from '@bass/bbs-sdk-fetch';
import type { StartEmailRegistrationOperationRequest } from '@bass/bbs-sdk-fetch';

async function example() {
  console.log("🚀 Testing @bass/bbs-sdk-fetch SDK...");
  const api = new AuthService();

  const body = {
    // StartEmailRegistrationRequest
    startEmailRegistrationRequest: ...,
  } satisfies StartEmailRegistrationOperationRequest;

  try {
    const data = await api.startEmailRegistration(body);
    console.log(data);
  } catch (error) {
    console.error(error);
  }
}

// Run the test
example().catch(console.error);
```

### Parameters


| Name | Type | Description  | Notes |
|------------- | ------------- | ------------- | -------------|
| **startEmailRegistrationRequest** | [StartEmailRegistrationRequest](StartEmailRegistrationRequest.md) |  | |

### Return type

[**StartEmailRegistrationReply**](StartEmailRegistrationReply.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: `application/json`
- **Accept**: `application/json`


### HTTP response details
| Status code | Description | Response headers |
|-------------|-------------|------------------|
| **200** | OK |  -  |

[[Back to top]](#) [[Back to API list]](../README.md#api-endpoints) [[Back to Model list]](../README.md#models) [[Back to README]](../README.md)


## startPhoneRegistration

> StartPhoneRegistrationReply startPhoneRegistration(startPhoneRegistrationRequest)



使用手机号发起账号注册

### Example

```ts
import {
  Configuration,
  AuthService,
} from '@bass/bbs-sdk-fetch';
import type { StartPhoneRegistrationOperationRequest } from '@bass/bbs-sdk-fetch';

async function example() {
  console.log("🚀 Testing @bass/bbs-sdk-fetch SDK...");
  const api = new AuthService();

  const body = {
    // StartPhoneRegistrationRequest
    startPhoneRegistrationRequest: ...,
  } satisfies StartPhoneRegistrationOperationRequest;

  try {
    const data = await api.startPhoneRegistration(body);
    console.log(data);
  } catch (error) {
    console.error(error);
  }
}

// Run the test
example().catch(console.error);
```

### Parameters


| Name | Type | Description  | Notes |
|------------- | ------------- | ------------- | -------------|
| **startPhoneRegistrationRequest** | [StartPhoneRegistrationRequest](StartPhoneRegistrationRequest.md) |  | |

### Return type

[**StartPhoneRegistrationReply**](StartPhoneRegistrationReply.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: `application/json`
- **Accept**: `application/json`


### HTTP response details
| Status code | Description | Response headers |
|-------------|-------------|------------------|
| **200** | OK |  -  |

[[Back to top]](#) [[Back to API list]](../README.md#api-endpoints) [[Back to Model list]](../README.md#models) [[Back to README]](../README.md)


## verifyEmailRegistration

> object verifyEmailRegistration(verifyEmailRegistrationRequest)



校验邮箱注册验证码

### Example

```ts
import {
  Configuration,
  AuthService,
} from '@bass/bbs-sdk-fetch';
import type { VerifyEmailRegistrationOperationRequest } from '@bass/bbs-sdk-fetch';

async function example() {
  console.log("🚀 Testing @bass/bbs-sdk-fetch SDK...");
  const api = new AuthService();

  const body = {
    // VerifyEmailRegistrationRequest
    verifyEmailRegistrationRequest: ...,
  } satisfies VerifyEmailRegistrationOperationRequest;

  try {
    const data = await api.verifyEmailRegistration(body);
    console.log(data);
  } catch (error) {
    console.error(error);
  }
}

// Run the test
example().catch(console.error);
```

### Parameters


| Name | Type | Description  | Notes |
|------------- | ------------- | ------------- | -------------|
| **verifyEmailRegistrationRequest** | [VerifyEmailRegistrationRequest](VerifyEmailRegistrationRequest.md) |  | |

### Return type

**object**

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: `application/json`
- **Accept**: `application/json`


### HTTP response details
| Status code | Description | Response headers |
|-------------|-------------|------------------|
| **200** | OK |  -  |

[[Back to top]](#) [[Back to API list]](../README.md#api-endpoints) [[Back to Model list]](../README.md#models) [[Back to README]](../README.md)


## verifyPhoneRegistration

> object verifyPhoneRegistration(verifyPhoneRegistrationRequest)



校验手机号注册验证码

### Example

```ts
import {
  Configuration,
  AuthService,
} from '@bass/bbs-sdk-fetch';
import type { VerifyPhoneRegistrationOperationRequest } from '@bass/bbs-sdk-fetch';

async function example() {
  console.log("🚀 Testing @bass/bbs-sdk-fetch SDK...");
  const api = new AuthService();

  const body = {
    // VerifyPhoneRegistrationRequest
    verifyPhoneRegistrationRequest: ...,
  } satisfies VerifyPhoneRegistrationOperationRequest;

  try {
    const data = await api.verifyPhoneRegistration(body);
    console.log(data);
  } catch (error) {
    console.error(error);
  }
}

// Run the test
example().catch(console.error);
```

### Parameters


| Name | Type | Description  | Notes |
|------------- | ------------- | ------------- | -------------|
| **verifyPhoneRegistrationRequest** | [VerifyPhoneRegistrationRequest](VerifyPhoneRegistrationRequest.md) |  | |

### Return type

**object**

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: `application/json`
- **Accept**: `application/json`


### HTTP response details
| Status code | Description | Response headers |
|-------------|-------------|------------------|
| **200** | OK |  -  |

[[Back to top]](#) [[Back to API list]](../README.md#api-endpoints) [[Back to Model list]](../README.md#models) [[Back to README]](../README.md)

