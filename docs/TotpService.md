# TotpService

All URIs are relative to *http://localhost*

| Method | HTTP request | Description |
|------------- | ------------- | -------------|
| [**beginEnable**](TotpService.md#beginenable) | **POST** /v1/user/totp/begin-enable |  |
| [**confirmEnable**](TotpService.md#confirmenable) | **POST** /v1/user/totp/confirm-enable |  |
| [**disable**](TotpService.md#disable) | **POST** /v1/user/totp/disable |  |
| [**getCurrent**](TotpService.md#getcurrent) | **POST** /v1/user/totp/get-current |  |



## beginEnable

> BeginEnableTotpReply beginEnable(body)



开始启用 TOTP。

### Example

```ts
import {
  Configuration,
  TotpService,
} from '@bass/bbs-sdk-fetch';
import type { BeginEnableRequest } from '@bass/bbs-sdk-fetch';

async function example() {
  console.log("🚀 Testing @bass/bbs-sdk-fetch SDK...");
  const api = new TotpService();

  const body = {
    // object
    body: Object,
  } satisfies BeginEnableRequest;

  try {
    const data = await api.beginEnable(body);
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

[**BeginEnableTotpReply**](BeginEnableTotpReply.md)

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


## confirmEnable

> object confirmEnable(confirmEnableTotpRequest)



确认 TOTP 验证码并正式启用 TOTP。

### Example

```ts
import {
  Configuration,
  TotpService,
} from '@bass/bbs-sdk-fetch';
import type { ConfirmEnableRequest } from '@bass/bbs-sdk-fetch';

async function example() {
  console.log("🚀 Testing @bass/bbs-sdk-fetch SDK...");
  const api = new TotpService();

  const body = {
    // ConfirmEnableTotpRequest
    confirmEnableTotpRequest: ...,
  } satisfies ConfirmEnableRequest;

  try {
    const data = await api.confirmEnable(body);
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
| **confirmEnableTotpRequest** | [ConfirmEnableTotpRequest](ConfirmEnableTotpRequest.md) |  | |

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


## disable

> object disable(disableTotpRequest)



校验 TOTP 验证码并关闭 TOTP。

### Example

```ts
import {
  Configuration,
  TotpService,
} from '@bass/bbs-sdk-fetch';
import type { DisableRequest } from '@bass/bbs-sdk-fetch';

async function example() {
  console.log("🚀 Testing @bass/bbs-sdk-fetch SDK...");
  const api = new TotpService();

  const body = {
    // DisableTotpRequest
    disableTotpRequest: ...,
  } satisfies DisableRequest;

  try {
    const data = await api.disable(body);
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
| **disableTotpRequest** | [DisableTotpRequest](DisableTotpRequest.md) |  | |

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


## getCurrent

> GetCurrentTotpReply getCurrent(body)



获取当前账号的 TOTP 状态。

### Example

```ts
import {
  Configuration,
  TotpService,
} from '@bass/bbs-sdk-fetch';
import type { GetCurrentRequest } from '@bass/bbs-sdk-fetch';

async function example() {
  console.log("🚀 Testing @bass/bbs-sdk-fetch SDK...");
  const api = new TotpService();

  const body = {
    // object
    body: Object,
  } satisfies GetCurrentRequest;

  try {
    const data = await api.getCurrent(body);
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

[**GetCurrentTotpReply**](GetCurrentTotpReply.md)

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

