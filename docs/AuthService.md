# AuthService

All URIs are relative to *http://localhost*

| Method | HTTP request | Description |
|------------- | ------------- | -------------|
| [**cancelAccount**](AuthService.md#cancelaccount) | **POST** /v1/user/auth/cancel-account |  |
| [**checkRegistrationAvailability**](AuthService.md#checkregistrationavailability) | **GET** /v1/user/auth/register-availability |  |
| [**login**](AuthService.md#login) | **POST** /v1/user/auth/login |  |
| [**logout**](AuthService.md#logout) | **POST** /v1/user/auth/logout |  |
| [**refreshToken**](AuthService.md#refreshtoken) | **POST** /v1/user/auth/refresh-token |  |
| [**register**](AuthService.md#register) | **POST** /v1/user/auth/register |  |



## cancelAccount

> object cancelAccount(cancelAccountReq)



注销账号。

### Example

```ts
import {
  Configuration,
  AuthService,
} from '@bass/bbs-sdk-fetch';
import type { CancelAccountRequest } from '@bass/bbs-sdk-fetch';

async function example() {
  console.log("🚀 Testing @bass/bbs-sdk-fetch SDK...");
  const api = new AuthService();

  const body = {
    // CancelAccountReq
    cancelAccountReq: ...,
  } satisfies CancelAccountRequest;

  try {
    const data = await api.cancelAccount(body);
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
| **cancelAccountReq** | [CancelAccountReq](CancelAccountReq.md) |  | |

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


## checkRegistrationAvailability

> CheckRegistrationAvailabilityResp checkRegistrationAvailability(name, email, phone)



检查注册字段是否可用。

### Example

```ts
import {
  Configuration,
  AuthService,
} from '@bass/bbs-sdk-fetch';
import type { CheckRegistrationAvailabilityRequest } from '@bass/bbs-sdk-fetch';

async function example() {
  console.log("🚀 Testing @bass/bbs-sdk-fetch SDK...");
  const api = new AuthService();

  const body = {
    // string (optional)
    name: name_example,
    // string (optional)
    email: email_example,
    // string (optional)
    phone: phone_example,
  } satisfies CheckRegistrationAvailabilityRequest;

  try {
    const data = await api.checkRegistrationAvailability(body);
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
| **name** | `string` |  | [Optional] [Defaults to `undefined`] |
| **email** | `string` |  | [Optional] [Defaults to `undefined`] |
| **phone** | `string` |  | [Optional] [Defaults to `undefined`] |

### Return type

[**CheckRegistrationAvailabilityResp**](CheckRegistrationAvailabilityResp.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: `application/json`


### HTTP response details
| Status code | Description | Response headers |
|-------------|-------------|------------------|
| **200** | OK |  -  |

[[Back to top]](#) [[Back to API list]](../README.md#api-endpoints) [[Back to Model list]](../README.md#models) [[Back to README]](../README.md)


## login

> LoginResp login(loginReq)



登录账号。

### Example

```ts
import {
  Configuration,
  AuthService,
} from '@bass/bbs-sdk-fetch';
import type { LoginRequest } from '@bass/bbs-sdk-fetch';

async function example() {
  console.log("🚀 Testing @bass/bbs-sdk-fetch SDK...");
  const api = new AuthService();

  const body = {
    // LoginReq
    loginReq: ...,
  } satisfies LoginRequest;

  try {
    const data = await api.login(body);
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
| **loginReq** | [LoginReq](LoginReq.md) |  | |

### Return type

[**LoginResp**](LoginResp.md)

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



退出登录。

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


## refreshToken

> RefreshTokenResp refreshToken(refreshTokenReq)



刷新登录令牌。

### Example

```ts
import {
  Configuration,
  AuthService,
} from '@bass/bbs-sdk-fetch';
import type { RefreshTokenRequest } from '@bass/bbs-sdk-fetch';

async function example() {
  console.log("🚀 Testing @bass/bbs-sdk-fetch SDK...");
  const api = new AuthService();

  const body = {
    // RefreshTokenReq
    refreshTokenReq: ...,
  } satisfies RefreshTokenRequest;

  try {
    const data = await api.refreshToken(body);
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
| **refreshTokenReq** | [RefreshTokenReq](RefreshTokenReq.md) |  | |

### Return type

[**RefreshTokenResp**](RefreshTokenResp.md)

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


## register

> object register(registerReq)



注册账号。

### Example

```ts
import {
  Configuration,
  AuthService,
} from '@bass/bbs-sdk-fetch';
import type { RegisterRequest } from '@bass/bbs-sdk-fetch';

async function example() {
  console.log("🚀 Testing @bass/bbs-sdk-fetch SDK...");
  const api = new AuthService();

  const body = {
    // RegisterReq
    registerReq: ...,
  } satisfies RegisterRequest;

  try {
    const data = await api.register(body);
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
| **registerReq** | [RegisterReq](RegisterReq.md) |  | |

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

