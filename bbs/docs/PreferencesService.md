# PreferencesService

All URIs are relative to *http://localhost*

|Method | HTTP request | Description|
|------------- | ------------- | -------------|
|[**getCurrent**](#getcurrent) | **POST** /v1/user/preference/get-current | |
|[**updateCurrent**](#updatecurrent) | **POST** /v1/user/preference/update-current | |

# **getCurrent**
> GetCurrentPreferencesReply getCurrent(body)

获取当前账号的偏好设置。

### Example

```typescript
import {
    PreferencesService,
    Configuration
} from '@bass/bbs-sdk-axios';

const configuration = new Configuration();
const apiInstance = new PreferencesService(configuration);

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

**GetCurrentPreferencesReply**

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

# **updateCurrent**
> UpdateCurrentPreferencesReply updateCurrent(updateCurrentPreferencesRequest)

更新当前账号的偏好设置。

### Example

```typescript
import {
    PreferencesService,
    Configuration,
    UpdateCurrentPreferencesRequest
} from '@bass/bbs-sdk-axios';

const configuration = new Configuration();
const apiInstance = new PreferencesService(configuration);

let updateCurrentPreferencesRequest: UpdateCurrentPreferencesRequest; //

const { status, data } = await apiInstance.updateCurrent(
    updateCurrentPreferencesRequest
);
```

### Parameters

|Name | Type | Description  | Notes|
|------------- | ------------- | ------------- | -------------|
| **updateCurrentPreferencesRequest** | **UpdateCurrentPreferencesRequest**|  | |


### Return type

**UpdateCurrentPreferencesReply**

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

