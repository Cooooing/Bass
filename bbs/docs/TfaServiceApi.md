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

> CommonApiAppBbsV1UserBeginEnableTfaReply tfaServiceBeginEnable(body)



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

[**CommonApiAppBbsV1UserBeginEnableTfaReply**](CommonApiAppBbsV1UserBeginEnableTfaReply.md)

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

> object tfaServiceConfirmEnable(commonApiAppBbsV1UserConfirmEnableTfaRequest)



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
    // CommonApiAppBbsV1UserConfirmEnableTfaRequest
    commonApiAppBbsV1UserConfirmEnableTfaRequest: ...,
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
| **commonApiAppBbsV1UserConfirmEnableTfaRequest** | [CommonApiAppBbsV1UserConfirmEnableTfaRequest](CommonApiAppBbsV1UserConfirmEnableTfaRequest.md) |  | |

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

> object tfaServiceDisable(commonApiAppBbsV1UserDisableTfaRequest)



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
    // CommonApiAppBbsV1UserDisableTfaRequest
    commonApiAppBbsV1UserDisableTfaRequest: ...,
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
| **commonApiAppBbsV1UserDisableTfaRequest** | [CommonApiAppBbsV1UserDisableTfaRequest](CommonApiAppBbsV1UserDisableTfaRequest.md) |  | |

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

> CommonApiAppBbsV1UserGetCurrentTfaReply tfaServiceGetCurrent(body)



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

[**CommonApiAppBbsV1UserGetCurrentTfaReply**](CommonApiAppBbsV1UserGetCurrentTfaReply.md)

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

> CommonApiAppBbsV1UserValidateTfaReply tfaServiceValidate(commonApiAppBbsV1UserValidateTfaRequest)



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
    // CommonApiAppBbsV1UserValidateTfaRequest
    commonApiAppBbsV1UserValidateTfaRequest: ...,
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
| **commonApiAppBbsV1UserValidateTfaRequest** | [CommonApiAppBbsV1UserValidateTfaRequest](CommonApiAppBbsV1UserValidateTfaRequest.md) |  | |

### Return type

[**CommonApiAppBbsV1UserValidateTfaReply**](CommonApiAppBbsV1UserValidateTfaReply.md)

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

