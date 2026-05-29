# AuthServiceApi

All URIs are relative to *http://localhost*

| Method | HTTP request | Description |
|------------- | ------------- | -------------|
| [**authServiceLoginByPassword**](AuthServiceApi.md#authserviceloginbypassword) | **POST** /v1/user/auth/login-by-password |  |
| [**authServiceLogout**](AuthServiceApi.md#authservicelogout) | **POST** /v1/user/auth/logout |  |
| [**authServiceStartEmailRegistration**](AuthServiceApi.md#authservicestartemailregistration) | **POST** /v1/user/auth/start-email-registration |  |
| [**authServiceStartPhoneRegistration**](AuthServiceApi.md#authservicestartphoneregistration) | **POST** /v1/user/auth/start-phone-registration |  |
| [**authServiceVerifyEmailRegistration**](AuthServiceApi.md#authserviceverifyemailregistration) | **POST** /v1/user/auth/verify-email-registration |  |
| [**authServiceVerifyPhoneRegistration**](AuthServiceApi.md#authserviceverifyphoneregistration) | **POST** /v1/user/auth/verify-phone-registration |  |



## authServiceLoginByPassword

> LoginByPasswordReply authServiceLoginByPassword(loginByPasswordRequest)



使用密码登录账号

### Example

```ts
import {
  Configuration,
  AuthServiceApi,
} from '@bass/bbs-sdk';
import type { AuthServiceLoginByPasswordRequest } from '@bass/bbs-sdk';

async function example() {
  console.log("🚀 Testing @bass/bbs-sdk SDK...");
  const api = new AuthServiceApi();

  const body = {
    // LoginByPasswordRequest
    loginByPasswordRequest: ...,
  } satisfies AuthServiceLoginByPasswordRequest;

  try {
    const data = await api.authServiceLoginByPassword(body);
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


## authServiceLogout

> object authServiceLogout(body)



登出当前登录账号

### Example

```ts
import {
  Configuration,
  AuthServiceApi,
} from '@bass/bbs-sdk';
import type { AuthServiceLogoutRequest } from '@bass/bbs-sdk';

async function example() {
  console.log("🚀 Testing @bass/bbs-sdk SDK...");
  const api = new AuthServiceApi();

  const body = {
    // object
    body: Object,
  } satisfies AuthServiceLogoutRequest;

  try {
    const data = await api.authServiceLogout(body);
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


## authServiceStartEmailRegistration

> StartEmailRegistrationReply authServiceStartEmailRegistration(startEmailRegistrationRequest)



使用邮箱发起账号注册

### Example

```ts
import {
  Configuration,
  AuthServiceApi,
} from '@bass/bbs-sdk';
import type { AuthServiceStartEmailRegistrationRequest } from '@bass/bbs-sdk';

async function example() {
  console.log("🚀 Testing @bass/bbs-sdk SDK...");
  const api = new AuthServiceApi();

  const body = {
    // StartEmailRegistrationRequest
    startEmailRegistrationRequest: ...,
  } satisfies AuthServiceStartEmailRegistrationRequest;

  try {
    const data = await api.authServiceStartEmailRegistration(body);
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


## authServiceStartPhoneRegistration

> StartPhoneRegistrationReply authServiceStartPhoneRegistration(startPhoneRegistrationRequest)



使用手机号发起账号注册

### Example

```ts
import {
  Configuration,
  AuthServiceApi,
} from '@bass/bbs-sdk';
import type { AuthServiceStartPhoneRegistrationRequest } from '@bass/bbs-sdk';

async function example() {
  console.log("🚀 Testing @bass/bbs-sdk SDK...");
  const api = new AuthServiceApi();

  const body = {
    // StartPhoneRegistrationRequest
    startPhoneRegistrationRequest: ...,
  } satisfies AuthServiceStartPhoneRegistrationRequest;

  try {
    const data = await api.authServiceStartPhoneRegistration(body);
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


## authServiceVerifyEmailRegistration

> object authServiceVerifyEmailRegistration(verifyEmailRegistrationRequest)



校验邮箱注册验证码

### Example

```ts
import {
  Configuration,
  AuthServiceApi,
} from '@bass/bbs-sdk';
import type { AuthServiceVerifyEmailRegistrationRequest } from '@bass/bbs-sdk';

async function example() {
  console.log("🚀 Testing @bass/bbs-sdk SDK...");
  const api = new AuthServiceApi();

  const body = {
    // VerifyEmailRegistrationRequest
    verifyEmailRegistrationRequest: ...,
  } satisfies AuthServiceVerifyEmailRegistrationRequest;

  try {
    const data = await api.authServiceVerifyEmailRegistration(body);
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


## authServiceVerifyPhoneRegistration

> object authServiceVerifyPhoneRegistration(verifyPhoneRegistrationRequest)



校验手机号注册验证码

### Example

```ts
import {
  Configuration,
  AuthServiceApi,
} from '@bass/bbs-sdk';
import type { AuthServiceVerifyPhoneRegistrationRequest } from '@bass/bbs-sdk';

async function example() {
  console.log("🚀 Testing @bass/bbs-sdk SDK...");
  const api = new AuthServiceApi();

  const body = {
    // VerifyPhoneRegistrationRequest
    verifyPhoneRegistrationRequest: ...,
  } satisfies AuthServiceVerifyPhoneRegistrationRequest;

  try {
    const data = await api.authServiceVerifyPhoneRegistration(body);
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

