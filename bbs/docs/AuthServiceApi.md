# AuthServiceApi

All URIs are relative to *http://localhost*

| Method | HTTP request | Description |
|------------- | ------------- | -------------|
| [**authServiceLoginPassword**](AuthServiceApi.md#authserviceloginpassword) | **POST** /v1/user/auth/login-password |  |
| [**authServiceLogout**](AuthServiceApi.md#authservicelogout) | **POST** /v1/user/auth/logout |  |
| [**authServiceRegisterEmail**](AuthServiceApi.md#authserviceregisteremail) | **POST** /v1/user/auth/register-email |  |
| [**authServiceRegisterPhone**](AuthServiceApi.md#authserviceregisterphone) | **POST** /v1/user/auth/register-phone |  |
| [**authServiceVerifyEmailRegister**](AuthServiceApi.md#authserviceverifyemailregister) | **POST** /v1/user/auth/verify-email-register |  |
| [**authServiceVerifyPhoneRegister**](AuthServiceApi.md#authserviceverifyphoneregister) | **POST** /v1/user/auth/verify-phone-register |  |



## authServiceLoginPassword

> CommonApiAppBbsV1UserLoginPasswordReply authServiceLoginPassword(commonApiAppBbsV1UserLoginPasswordRequest)



使用密码登录账号

### Example

```ts
import {
  Configuration,
  AuthServiceApi,
} from '@bass/bbs-sdk';
import type { AuthServiceLoginPasswordRequest } from '@bass/bbs-sdk';

async function example() {
  console.log("🚀 Testing @bass/bbs-sdk SDK...");
  const api = new AuthServiceApi();

  const body = {
    // CommonApiAppBbsV1UserLoginPasswordRequest
    commonApiAppBbsV1UserLoginPasswordRequest: ...,
  } satisfies AuthServiceLoginPasswordRequest;

  try {
    const data = await api.authServiceLoginPassword(body);
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
| **commonApiAppBbsV1UserLoginPasswordRequest** | [CommonApiAppBbsV1UserLoginPasswordRequest](CommonApiAppBbsV1UserLoginPasswordRequest.md) |  | |

### Return type

[**CommonApiAppBbsV1UserLoginPasswordReply**](CommonApiAppBbsV1UserLoginPasswordReply.md)

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


## authServiceRegisterEmail

> CommonApiAppBbsV1UserRegisterEmailReply authServiceRegisterEmail(commonApiAppBbsV1UserRegisterEmailRequest)



使用邮箱发起账号注册

### Example

```ts
import {
  Configuration,
  AuthServiceApi,
} from '@bass/bbs-sdk';
import type { AuthServiceRegisterEmailRequest } from '@bass/bbs-sdk';

async function example() {
  console.log("🚀 Testing @bass/bbs-sdk SDK...");
  const api = new AuthServiceApi();

  const body = {
    // CommonApiAppBbsV1UserRegisterEmailRequest
    commonApiAppBbsV1UserRegisterEmailRequest: ...,
  } satisfies AuthServiceRegisterEmailRequest;

  try {
    const data = await api.authServiceRegisterEmail(body);
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
| **commonApiAppBbsV1UserRegisterEmailRequest** | [CommonApiAppBbsV1UserRegisterEmailRequest](CommonApiAppBbsV1UserRegisterEmailRequest.md) |  | |

### Return type

[**CommonApiAppBbsV1UserRegisterEmailReply**](CommonApiAppBbsV1UserRegisterEmailReply.md)

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


## authServiceRegisterPhone

> CommonApiAppBbsV1UserRegisterPhoneReply authServiceRegisterPhone(commonApiAppBbsV1UserRegisterPhoneRequest)



使用手机号发起账号注册

### Example

```ts
import {
  Configuration,
  AuthServiceApi,
} from '@bass/bbs-sdk';
import type { AuthServiceRegisterPhoneRequest } from '@bass/bbs-sdk';

async function example() {
  console.log("🚀 Testing @bass/bbs-sdk SDK...");
  const api = new AuthServiceApi();

  const body = {
    // CommonApiAppBbsV1UserRegisterPhoneRequest
    commonApiAppBbsV1UserRegisterPhoneRequest: ...,
  } satisfies AuthServiceRegisterPhoneRequest;

  try {
    const data = await api.authServiceRegisterPhone(body);
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
| **commonApiAppBbsV1UserRegisterPhoneRequest** | [CommonApiAppBbsV1UserRegisterPhoneRequest](CommonApiAppBbsV1UserRegisterPhoneRequest.md) |  | |

### Return type

[**CommonApiAppBbsV1UserRegisterPhoneReply**](CommonApiAppBbsV1UserRegisterPhoneReply.md)

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


## authServiceVerifyEmailRegister

> object authServiceVerifyEmailRegister(commonApiAppBbsV1UserVerifyEmailRegisterRequest)



校验邮箱注册验证码

### Example

```ts
import {
  Configuration,
  AuthServiceApi,
} from '@bass/bbs-sdk';
import type { AuthServiceVerifyEmailRegisterRequest } from '@bass/bbs-sdk';

async function example() {
  console.log("🚀 Testing @bass/bbs-sdk SDK...");
  const api = new AuthServiceApi();

  const body = {
    // CommonApiAppBbsV1UserVerifyEmailRegisterRequest
    commonApiAppBbsV1UserVerifyEmailRegisterRequest: ...,
  } satisfies AuthServiceVerifyEmailRegisterRequest;

  try {
    const data = await api.authServiceVerifyEmailRegister(body);
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
| **commonApiAppBbsV1UserVerifyEmailRegisterRequest** | [CommonApiAppBbsV1UserVerifyEmailRegisterRequest](CommonApiAppBbsV1UserVerifyEmailRegisterRequest.md) |  | |

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


## authServiceVerifyPhoneRegister

> object authServiceVerifyPhoneRegister(commonApiAppBbsV1UserVerifyPhoneRegisterRequest)



校验手机号注册验证码

### Example

```ts
import {
  Configuration,
  AuthServiceApi,
} from '@bass/bbs-sdk';
import type { AuthServiceVerifyPhoneRegisterRequest } from '@bass/bbs-sdk';

async function example() {
  console.log("🚀 Testing @bass/bbs-sdk SDK...");
  const api = new AuthServiceApi();

  const body = {
    // CommonApiAppBbsV1UserVerifyPhoneRegisterRequest
    commonApiAppBbsV1UserVerifyPhoneRegisterRequest: ...,
  } satisfies AuthServiceVerifyPhoneRegisterRequest;

  try {
    const data = await api.authServiceVerifyPhoneRegister(body);
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
| **commonApiAppBbsV1UserVerifyPhoneRegisterRequest** | [CommonApiAppBbsV1UserVerifyPhoneRegisterRequest](CommonApiAppBbsV1UserVerifyPhoneRegisterRequest.md) |  | |

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

