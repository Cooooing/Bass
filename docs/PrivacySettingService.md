# PrivacySettingService

All URIs are relative to *http://localhost*

|Method | HTTP request | Description|
|------------- | ------------- | -------------|
|[**getCurrent**](#getcurrent) | **POST** /v1/user/privacy-setting/get-current | |
|[**updateCurrent**](#updatecurrent) | **POST** /v1/user/privacy-setting/update-current | |

# **getCurrent**
> GetCurrentPrivacySettingResp getCurrent(body)

获取当前账号的隐私设置。

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

**GetCurrentPrivacySettingResp**

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
> UpdateCurrentPrivacySettingResp updateCurrent(updateCurrentPrivacySettingReq)

更新当前账号的隐私设置。

### Example

```typescript
import {
    PrivacySettingService,
    Configuration,
    UpdateCurrentPrivacySettingReq
} from '@bass/bbs-sdk-axios';

const configuration = new Configuration();
const apiInstance = new PrivacySettingService(configuration);

let updateCurrentPrivacySettingReq: UpdateCurrentPrivacySettingReq; //

const { status, data } = await apiInstance.updateCurrent(
    updateCurrentPrivacySettingReq
);
```

### Parameters

|Name | Type | Description  | Notes|
|------------- | ------------- | ------------- | -------------|
| **updateCurrentPrivacySettingReq** | **UpdateCurrentPrivacySettingReq**|  | |


### Return type

**UpdateCurrentPrivacySettingResp**

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

