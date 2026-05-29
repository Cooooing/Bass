# TfaServiceApi

All URIs are relative to *http://localhost*

| Method | HTTP request | Description |
|------------- | ------------- | -------------|
| [**tfaServiceBeginEnable**](TfaServiceApi.md#tfaservicebeginenable) | **POST** /v1/user/tfa/begin-enable |  |
| [**tfaServiceConfirmEnable**](TfaServiceApi.md#tfaserviceconfirmenable) | **POST** /v1/user/tfa/confirm-enable |  |
| [**tfaServiceDisable**](TfaServiceApi.md#tfaservicedisable) | **POST** /v1/user/tfa/disable |  |
| [**tfaServiceGetCurrent**](TfaServiceApi.md#tfaservicegetcurrent) | **POST** /v1/user/tfa/get-current |  |
| [**tfaServiceValidate**](TfaServiceApi.md#tfaservicevalidate) | **POST** /v1/user/tfa/validate |  |



## tfaServiceBeginEnable

> BeginEnableTfaReply tfaServiceBeginEnable(body)



开始启用二步验证

### Example

```ts
import {
  Configuration,
  TfaServiceApi,
} from '@bass/bbs-sdk';
import type { TfaServiceBeginEnableRequest } from '@bass/bbs-sdk';

async function example() {
  console.log("🚀 Testing @bass/bbs-sdk SDK...");
  const api = new TfaServiceApi();

  const body = {
    // object
    body: Object,
  } satisfies TfaServiceBeginEnableRequest;

  try {
    const data = await api.tfaServiceBeginEnable(body);
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


## tfaServiceConfirmEnable

> object tfaServiceConfirmEnable(confirmEnableTfaRequest)



确认二步验证码并正式启用二步验证

### Example

```ts
import {
  Configuration,
  TfaServiceApi,
} from '@bass/bbs-sdk';
import type { TfaServiceConfirmEnableRequest } from '@bass/bbs-sdk';

async function example() {
  console.log("🚀 Testing @bass/bbs-sdk SDK...");
  const api = new TfaServiceApi();

  const body = {
    // ConfirmEnableTfaRequest
    confirmEnableTfaRequest: ...,
  } satisfies TfaServiceConfirmEnableRequest;

  try {
    const data = await api.tfaServiceConfirmEnable(body);
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


## tfaServiceDisable

> object tfaServiceDisable(disableTfaRequest)



校验二步验证码并关闭二步验证

### Example

```ts
import {
  Configuration,
  TfaServiceApi,
} from '@bass/bbs-sdk';
import type { TfaServiceDisableRequest } from '@bass/bbs-sdk';

async function example() {
  console.log("🚀 Testing @bass/bbs-sdk SDK...");
  const api = new TfaServiceApi();

  const body = {
    // DisableTfaRequest
    disableTfaRequest: ...,
  } satisfies TfaServiceDisableRequest;

  try {
    const data = await api.tfaServiceDisable(body);
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


## tfaServiceGetCurrent

> GetCurrentTfaReply tfaServiceGetCurrent(body)



获取当前登录账号的二步验证状态

### Example

```ts
import {
  Configuration,
  TfaServiceApi,
} from '@bass/bbs-sdk';
import type { TfaServiceGetCurrentRequest } from '@bass/bbs-sdk';

async function example() {
  console.log("🚀 Testing @bass/bbs-sdk SDK...");
  const api = new TfaServiceApi();

  const body = {
    // object
    body: Object,
  } satisfies TfaServiceGetCurrentRequest;

  try {
    const data = await api.tfaServiceGetCurrent(body);
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


## tfaServiceValidate

> ValidateTfaReply tfaServiceValidate(validateTfaRequest)



校验当前登录账号的二步验证码

### Example

```ts
import {
  Configuration,
  TfaServiceApi,
} from '@bass/bbs-sdk';
import type { TfaServiceValidateRequest } from '@bass/bbs-sdk';

async function example() {
  console.log("🚀 Testing @bass/bbs-sdk SDK...");
  const api = new TfaServiceApi();

  const body = {
    // ValidateTfaRequest
    validateTfaRequest: ...,
  } satisfies TfaServiceValidateRequest;

  try {
    const data = await api.tfaServiceValidate(body);
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

