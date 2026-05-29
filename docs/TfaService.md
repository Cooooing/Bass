# TfaService

All URIs are relative to *http://localhost*

| Method | HTTP request | Description |
|------------- | ------------- | -------------|
| [**beginEnable**](TfaService.md#beginenable) | **POST** /v1/user/tfa/begin-enable |  |
| [**confirmEnable**](TfaService.md#confirmenable) | **POST** /v1/user/tfa/confirm-enable |  |
| [**disable**](TfaService.md#disable) | **POST** /v1/user/tfa/disable |  |
| [**getCurrent**](TfaService.md#getcurrent) | **POST** /v1/user/tfa/get-current |  |
| [**validate**](TfaService.md#validate) | **POST** /v1/user/tfa/validate |  |



## beginEnable

> BeginEnableTfaReply beginEnable(body)



开始启用二步验证

### Example

```ts
import {
  Configuration,
  TfaService,
} from '@bass/bbs-sdk-fetch';
import type { BeginEnableRequest } from '@bass/bbs-sdk-fetch';

async function example() {
  console.log("🚀 Testing @bass/bbs-sdk-fetch SDK...");
  const api = new TfaService();

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

[**BeginEnableTfaReply**](BeginEnableTfaReply.md)

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

> object confirmEnable(confirmEnableTfaRequest)



确认二步验证码并正式启用二步验证

### Example

```ts
import {
  Configuration,
  TfaService,
} from '@bass/bbs-sdk-fetch';
import type { ConfirmEnableRequest } from '@bass/bbs-sdk-fetch';

async function example() {
  console.log("🚀 Testing @bass/bbs-sdk-fetch SDK...");
  const api = new TfaService();

  const body = {
    // ConfirmEnableTfaRequest
    confirmEnableTfaRequest: ...,
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
| **confirmEnableTfaRequest** | [ConfirmEnableTfaRequest](ConfirmEnableTfaRequest.md) |  | |

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

> object disable(disableTfaRequest)



校验二步验证码并关闭二步验证

### Example

```ts
import {
  Configuration,
  TfaService,
} from '@bass/bbs-sdk-fetch';
import type { DisableRequest } from '@bass/bbs-sdk-fetch';

async function example() {
  console.log("🚀 Testing @bass/bbs-sdk-fetch SDK...");
  const api = new TfaService();

  const body = {
    // DisableTfaRequest
    disableTfaRequest: ...,
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
| **disableTfaRequest** | [DisableTfaRequest](DisableTfaRequest.md) |  | |

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

> GetCurrentTfaReply getCurrent(body)



获取当前登录账号的二步验证状态

### Example

```ts
import {
  Configuration,
  TfaService,
} from '@bass/bbs-sdk-fetch';
import type { GetCurrentRequest } from '@bass/bbs-sdk-fetch';

async function example() {
  console.log("🚀 Testing @bass/bbs-sdk-fetch SDK...");
  const api = new TfaService();

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

[**GetCurrentTfaReply**](GetCurrentTfaReply.md)

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


## validate

> ValidateTfaReply validate(validateTfaRequest)



校验当前登录账号的二步验证码

### Example

```ts
import {
  Configuration,
  TfaService,
} from '@bass/bbs-sdk-fetch';
import type { ValidateRequest } from '@bass/bbs-sdk-fetch';

async function example() {
  console.log("🚀 Testing @bass/bbs-sdk-fetch SDK...");
  const api = new TfaService();

  const body = {
    // ValidateTfaRequest
    validateTfaRequest: ...,
  } satisfies ValidateRequest;

  try {
    const data = await api.validate(body);
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
| **validateTfaRequest** | [ValidateTfaRequest](ValidateTfaRequest.md) |  | |

### Return type

[**ValidateTfaReply**](ValidateTfaReply.md)

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

