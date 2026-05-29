# TfaService

All URIs are relative to *http://localhost*

|Method | HTTP request | Description|
|------------- | ------------- | -------------|
|[**beginEnable**](#beginenable) | **POST** /v1/user/tfa/begin-enable | |
|[**confirmEnable**](#confirmenable) | **POST** /v1/user/tfa/confirm-enable | |
|[**disable**](#disable) | **POST** /v1/user/tfa/disable | |
|[**getCurrent**](#getcurrent) | **POST** /v1/user/tfa/get-current | |
|[**validate**](#validate) | **POST** /v1/user/tfa/validate | |

# **beginEnable**
> BeginEnableTfaReply beginEnable(body)

开始启用二步验证

### Example

```typescript
import {
    TfaService,
    Configuration
} from '@bass/bbs-sdk-axios';

const configuration = new Configuration();
const apiInstance = new TfaService(configuration);

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

**BeginEnableTfaReply**

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
> object confirmEnable(confirmEnableTfaRequest)

确认二步验证码并正式启用二步验证

### Example

```typescript
import {
    TfaService,
    Configuration,
    ConfirmEnableTfaRequest
} from '@bass/bbs-sdk-axios';

const configuration = new Configuration();
const apiInstance = new TfaService(configuration);

let confirmEnableTfaRequest: ConfirmEnableTfaRequest; //

const { status, data } = await apiInstance.confirmEnable(
    confirmEnableTfaRequest
);
```

### Parameters

|Name | Type | Description  | Notes|
|------------- | ------------- | ------------- | -------------|
| **confirmEnableTfaRequest** | **ConfirmEnableTfaRequest**|  | |


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
> object disable(disableTfaRequest)

校验二步验证码并关闭二步验证

### Example

```typescript
import {
    TfaService,
    Configuration,
    DisableTfaRequest
} from '@bass/bbs-sdk-axios';

const configuration = new Configuration();
const apiInstance = new TfaService(configuration);

let disableTfaRequest: DisableTfaRequest; //

const { status, data } = await apiInstance.disable(
    disableTfaRequest
);
```

### Parameters

|Name | Type | Description  | Notes|
|------------- | ------------- | ------------- | -------------|
| **disableTfaRequest** | **DisableTfaRequest**|  | |


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
> GetCurrentTfaReply getCurrent(body)

获取当前登录账号的二步验证状态

### Example

```typescript
import {
    TfaService,
    Configuration
} from '@bass/bbs-sdk-axios';

const configuration = new Configuration();
const apiInstance = new TfaService(configuration);

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

**GetCurrentTfaReply**

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

# **validate**
> ValidateTfaReply validate(validateTfaRequest)

校验当前登录账号的二步验证码

### Example

```typescript
import {
    TfaService,
    Configuration,
    ValidateTfaRequest
} from '@bass/bbs-sdk-axios';

const configuration = new Configuration();
const apiInstance = new TfaService(configuration);

let validateTfaRequest: ValidateTfaRequest; //

const { status, data } = await apiInstance.validate(
    validateTfaRequest
);
```

### Parameters

|Name | Type | Description  | Notes|
|------------- | ------------- | ------------- | -------------|
| **validateTfaRequest** | **ValidateTfaRequest**|  | |


### Return type

**ValidateTfaReply**

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

