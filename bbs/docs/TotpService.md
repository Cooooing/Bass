# TotpService

All URIs are relative to *http://localhost*

|Method | HTTP request | Description|
|------------- | ------------- | -------------|
|[**beginEnable**](#beginenable) | **POST** /v1/user/totp/begin-enable | |
|[**confirmEnable**](#confirmenable) | **POST** /v1/user/totp/confirm-enable | |
|[**disable**](#disable) | **POST** /v1/user/totp/disable | |
|[**getCurrent**](#getcurrent) | **POST** /v1/user/totp/get-current | |

# **beginEnable**
> BeginEnableTotpReply beginEnable(body)

开始启用 TOTP。

### Example

```typescript
import {
    TotpService,
    Configuration
} from '@bass/bbs-sdk-axios';

const configuration = new Configuration();
const apiInstance = new TotpService(configuration);

let body: object; //

const { status, data } = await apiInstance.beginEnable(
    body
);
```

### Parameters

|Name | Type | Description  | Notes|
|------------- | ------------- | ------------- | -------------|
| **body** | **object**|  | |


### Return type

**BeginEnableTotpReply**

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json


### HTTP response details
| Status code | Description | Response headers |
|-------------|-------------|------------------|
|**200** | OK |  -  |

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **confirmEnable**
> object confirmEnable(confirmEnableTotpRequest)

确认 TOTP 验证码并正式启用 TOTP。

### Example

```typescript
import {
    TotpService,
    Configuration,
    ConfirmEnableTotpRequest
} from '@bass/bbs-sdk-axios';

const configuration = new Configuration();
const apiInstance = new TotpService(configuration);

let confirmEnableTotpRequest: ConfirmEnableTotpRequest; //

const { status, data } = await apiInstance.confirmEnable(
    confirmEnableTotpRequest
);
```

### Parameters

|Name | Type | Description  | Notes|
|------------- | ------------- | ------------- | -------------|
| **confirmEnableTotpRequest** | **ConfirmEnableTotpRequest**|  | |


### Return type

**object**

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json


### HTTP response details
| Status code | Description | Response headers |
|-------------|-------------|------------------|
|**200** | OK |  -  |

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **disable**
> object disable(disableTotpRequest)

校验 TOTP 验证码并关闭 TOTP。

### Example

```typescript
import {
    TotpService,
    Configuration,
    DisableTotpRequest
} from '@bass/bbs-sdk-axios';

const configuration = new Configuration();
const apiInstance = new TotpService(configuration);

let disableTotpRequest: DisableTotpRequest; //

const { status, data } = await apiInstance.disable(
    disableTotpRequest
);
```

### Parameters

|Name | Type | Description  | Notes|
|------------- | ------------- | ------------- | -------------|
| **disableTotpRequest** | **DisableTotpRequest**|  | |


### Return type

**object**

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json


### HTTP response details
| Status code | Description | Response headers |
|-------------|-------------|------------------|
|**200** | OK |  -  |

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **getCurrent**
> GetCurrentTotpReply getCurrent(body)

获取当前账号的 TOTP 状态。

### Example

```typescript
import {
    TotpService,
    Configuration
} from '@bass/bbs-sdk-axios';

const configuration = new Configuration();
const apiInstance = new TotpService(configuration);

let body: object; //

const { status, data } = await apiInstance.getCurrent(
    body
);
```

### Parameters

|Name | Type | Description  | Notes|
|------------- | ------------- | ------------- | -------------|
| **body** | **object**|  | |


### Return type

**GetCurrentTotpReply**

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json


### HTTP response details
| Status code | Description | Response headers |
|-------------|-------------|------------------|
|**200** | OK |  -  |

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

