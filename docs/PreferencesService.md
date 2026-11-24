# PreferencesService

All URIs are relative to *http://localhost*

|Method | HTTP request | Description|
|------------- | ------------- | -------------|
|[**getCurrentPreferences**](#getcurrentpreferences) | **POST** /v1/user/preference/get-current | |
|[**updateCurrentPreferences**](#updatecurrentpreferences) | **POST** /v1/user/preference/update-current | |

# **getCurrentPreferences**
> GetCurrentPreferencesResp getCurrentPreferences(body)

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

const { status, data } = await apiInstance.getCurrentPreferences(
    body
);
```

### Parameters

|Name | Type | Description  | Notes|
|------------- | ------------- | ------------- | -------------|
| **body** | **object**|  | |


### Return type

**GetCurrentPreferencesResp**

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

# **updateCurrentPreferences**
> UpdateCurrentPreferencesResp updateCurrentPreferences(updateCurrentPreferencesReq)

更新当前账号的偏好设置。

### Example

```typescript
import {
    PreferencesService,
    Configuration,
    UpdateCurrentPreferencesReq
} from '@bass/bbs-sdk-axios';

const configuration = new Configuration();
const apiInstance = new PreferencesService(configuration);

let updateCurrentPreferencesReq: UpdateCurrentPreferencesReq; //

const { status, data } = await apiInstance.updateCurrentPreferences(
    updateCurrentPreferencesReq
);
```

### Parameters

|Name | Type | Description  | Notes|
|------------- | ------------- | ------------- | -------------|
| **updateCurrentPreferencesReq** | **UpdateCurrentPreferencesReq**|  | |


### Return type

**UpdateCurrentPreferencesResp**

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

