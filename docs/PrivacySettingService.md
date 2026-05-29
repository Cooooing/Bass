# PrivacySettingService

All URIs are relative to *http://localhost*

|Method | HTTP request | Description|
|------------- | ------------- | -------------|
|[**getCurrent**](#getcurrent) | **POST** /v1/user/privacy-setting/get-current | |
|[**updateCurrent**](#updatecurrent) | **POST** /v1/user/privacy-setting/update-current | |

# **getCurrent**
> GetCurrentPrivacySettingReply getCurrent(body)

获取当前登录账号的隐私设置

### Example

```typescript
import {
    PrivacySettingService,
    Configuration
} from '@bass/bbs-sdk-axios';

const configuration = new Configuration();
const apiInstance = new PrivacySettingService(configuration);

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

**GetCurrentPrivacySettingReply**

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
> UpdateCurrentPrivacySettingReply updateCurrent(updateCurrentPrivacySettingRequest)

更新当前登录账号的隐私设置

### Example

```typescript
import {
    PrivacySettingService,
    Configuration,
    UpdateCurrentPrivacySettingRequest
} from '@bass/bbs-sdk-axios';

const configuration = new Configuration();
const apiInstance = new PrivacySettingService(configuration);

let updateCurrentPrivacySettingRequest: UpdateCurrentPrivacySettingRequest; //

const { status, data } = await apiInstance.updateCurrent(
    updateCurrentPrivacySettingRequest
);
```

### Parameters

|Name | Type | Description  | Notes|
|------------- | ------------- | ------------- | -------------|
| **updateCurrentPrivacySettingRequest** | **UpdateCurrentPrivacySettingRequest**|  | |


### Return type

**UpdateCurrentPrivacySettingReply**

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

