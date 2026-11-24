# PrivacySettingService

All URIs are relative to *http://localhost*

|Method | HTTP request | Description|
|------------- | ------------- | -------------|
|[**getCurrentPrivacySetting**](#getcurrentprivacysetting) | **POST** /v1/user/privacy-setting/get-current | |
|[**updateCurrentPrivacySetting**](#updatecurrentprivacysetting) | **POST** /v1/user/privacy-setting/update-current | |

# **getCurrentPrivacySetting**
> GetCurrentPrivacySettingResp getCurrentPrivacySetting(body)

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

const { status, data } = await apiInstance.getCurrentPrivacySetting(
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

# **updateCurrentPrivacySetting**
> UpdateCurrentPrivacySettingResp updateCurrentPrivacySetting(updateCurrentPrivacySettingReq)

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

const { status, data } = await apiInstance.updateCurrentPrivacySetting(
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

