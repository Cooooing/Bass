# AuthService

All URIs are relative to *http://localhost*

| Method | HTTP request | Description |
|------------- | ------------- | -------------|
| [**login**](AuthService.md#login) | **POST** /v1/game-idle/auth/login |  |
| [**register**](AuthService.md#register) | **POST** /v1/game-idle/auth/register |  |



## login

> LoginAccountResp login(loginAccountReq)



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
    // LoginAccountReq
    loginAccountReq: ...,
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
| **loginAccountReq** | [LoginAccountReq](LoginAccountReq.md) |  | |

### Return type

[**LoginAccountResp**](LoginAccountResp.md)

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

> object register(registerAccountReq)



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
    // RegisterAccountReq
    registerAccountReq: ...,
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
| **registerAccountReq** | [RegisterAccountReq](RegisterAccountReq.md) |  | |

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

