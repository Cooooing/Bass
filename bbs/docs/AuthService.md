# AuthService

All URIs are relative to *http://localhost*

| Method | HTTP request | Description |
|------------- | ------------- | -------------|
| [**cancelAccount**](AuthService.md#cancelaccount) | **POST** /v1/user/auth/cancel-account |  |
| [**login**](AuthService.md#login) | **POST** /v1/user/auth/login |  |
| [**logout**](AuthService.md#logout) | **POST** /v1/user/auth/logout |  |
| [**refreshToken**](AuthService.md#refreshtoken) | **POST** /v1/user/auth/refresh-token |  |
| [**startEmailRegistration**](AuthService.md#startemailregistration) | **POST** /v1/user/auth/start-email-registration |  |
| [**startPhoneRegistration**](AuthService.md#startphoneregistration) | **POST** /v1/user/auth/start-phone-registration |  |
| [**verifyEmailRegistration**](AuthService.md#verifyemailregistration) | **POST** /v1/user/auth/verify-email-registration |  |
| [**verifyPhoneRegistration**](AuthService.md#verifyphoneregistration) | **POST** /v1/user/auth/verify-phone-registration |  |



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


## startEmailRegistration

> StartEmailRegistrationResp startEmailRegistration(startEmailRegistrationReq)



开始邮箱注册。

### Example

```ts
import {
  Configuration,
  AuthService,
} from '@bass/bbs-sdk-fetch';
import type { StartEmailRegistrationRequest } from '@bass/bbs-sdk-fetch';

async function example() {
  console.log("🚀 Testing @bass/bbs-sdk-fetch SDK...");
  const api = new AuthService();

  const body = {
    // StartEmailRegistrationReq
    startEmailRegistrationReq: ...,
  } satisfies StartEmailRegistrationRequest;

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
| **startEmailRegistrationReq** | [StartEmailRegistrationReq](StartEmailRegistrationReq.md) |  | |

### Return type

[**StartEmailRegistrationResp**](StartEmailRegistrationResp.md)

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

> StartPhoneRegistrationResp startPhoneRegistration(startPhoneRegistrationReq)



开始手机注册。

### Example

```ts
import {
  Configuration,
  AuthService,
} from '@bass/bbs-sdk-fetch';
import type { StartPhoneRegistrationRequest } from '@bass/bbs-sdk-fetch';

async function example() {
  console.log("🚀 Testing @bass/bbs-sdk-fetch SDK...");
  const api = new AuthService();

  const body = {
    // StartPhoneRegistrationReq
    startPhoneRegistrationReq: ...,
  } satisfies StartPhoneRegistrationRequest;

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
| **startPhoneRegistrationReq** | [StartPhoneRegistrationReq](StartPhoneRegistrationReq.md) |  | |

### Return type

[**StartPhoneRegistrationResp**](StartPhoneRegistrationResp.md)

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

> object verifyEmailRegistration(verifyEmailRegistrationReq)



校验邮箱注册验证码。

### Example

```ts
import {
  Configuration,
  AuthService,
} from '@bass/bbs-sdk-fetch';
import type { VerifyEmailRegistrationRequest } from '@bass/bbs-sdk-fetch';

async function example() {
  console.log("🚀 Testing @bass/bbs-sdk-fetch SDK...");
  const api = new AuthService();

  const body = {
    // VerifyEmailRegistrationReq
    verifyEmailRegistrationReq: ...,
  } satisfies VerifyEmailRegistrationRequest;

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
| **verifyEmailRegistrationReq** | [VerifyEmailRegistrationReq](VerifyEmailRegistrationReq.md) |  | |

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

> object verifyPhoneRegistration(verifyPhoneRegistrationReq)



校验手机注册验证码。

### Example

```ts
import {
  Configuration,
  AuthService,
} from '@bass/bbs-sdk-fetch';
import type { VerifyPhoneRegistrationRequest } from '@bass/bbs-sdk-fetch';

async function example() {
  console.log("🚀 Testing @bass/bbs-sdk-fetch SDK...");
  const api = new AuthService();

  const body = {
    // VerifyPhoneRegistrationReq
    verifyPhoneRegistrationReq: ...,
  } satisfies VerifyPhoneRegistrationRequest;

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
| **verifyPhoneRegistrationReq** | [VerifyPhoneRegistrationReq](VerifyPhoneRegistrationReq.md) |  | |

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

