# PrivacySettingService

All URIs are relative to *http://localhost*

| Method | HTTP request | Description |
|------------- | ------------- | -------------|
| [**getCurrentPrivacySetting**](PrivacySettingService.md#getcurrentprivacysetting) | **POST** /v1/user/privacy-setting/get-current |  |
| [**updateCurrentPrivacySetting**](PrivacySettingService.md#updatecurrentprivacysetting) | **POST** /v1/user/privacy-setting/update-current |  |



## getCurrentPrivacySetting

> GetCurrentPrivacySettingResp getCurrentPrivacySetting(body)



获取当前账号的隐私设置。

### Example

```ts
import {
  Configuration,
  PrivacySettingService,
} from '@bass/bbs-sdk-fetch';
import type { GetCurrentPrivacySettingRequest } from '@bass/bbs-sdk-fetch';

async function example() {
  console.log("🚀 Testing @bass/bbs-sdk-fetch SDK...");
  const api = new PrivacySettingService();

  const body = {
    // object
    body: Object,
  } satisfies GetCurrentPrivacySettingRequest;

  try {
    const data = await api.getCurrentPrivacySetting(body);
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

[**GetCurrentPrivacySettingResp**](GetCurrentPrivacySettingResp.md)

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


## updateCurrentPrivacySetting

> UpdateCurrentPrivacySettingResp updateCurrentPrivacySetting(updateCurrentPrivacySettingReq)



更新当前账号的隐私设置。

### Example

```ts
import {
  Configuration,
  PrivacySettingService,
} from '@bass/bbs-sdk-fetch';
import type { UpdateCurrentPrivacySettingRequest } from '@bass/bbs-sdk-fetch';

async function example() {
  console.log("🚀 Testing @bass/bbs-sdk-fetch SDK...");
  const api = new PrivacySettingService();

  const body = {
    // UpdateCurrentPrivacySettingReq
    updateCurrentPrivacySettingReq: ...,
  } satisfies UpdateCurrentPrivacySettingRequest;

  try {
    const data = await api.updateCurrentPrivacySetting(body);
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
| **updateCurrentPrivacySettingReq** | [UpdateCurrentPrivacySettingReq](UpdateCurrentPrivacySettingReq.md) |  | |

### Return type

[**UpdateCurrentPrivacySettingResp**](UpdateCurrentPrivacySettingResp.md)

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

