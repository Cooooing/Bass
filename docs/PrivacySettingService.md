# PrivacySettingService

All URIs are relative to *http://localhost*

| Method | HTTP request | Description |
|------------- | ------------- | -------------|
| [**getCurrent**](PrivacySettingService.md#getcurrent) | **POST** /v1/user/privacy-setting/get-current |  |
| [**updateCurrent**](PrivacySettingService.md#updatecurrent) | **POST** /v1/user/privacy-setting/update-current |  |



## getCurrent

> GetCurrentPrivacySettingReply getCurrent(body)



获取当前登录账号的隐私设置

### Example

```ts
import {
  Configuration,
  PrivacySettingService,
} from '@bass/bbs-sdk-fetch';
import type { GetCurrentRequest } from '@bass/bbs-sdk-fetch';

async function example() {
  console.log("🚀 Testing @bass/bbs-sdk-fetch SDK...");
  const api = new PrivacySettingService();

  const body = {
    // object
    body: Object,
  } satisfies GetCurrentRequest;

  try {
    const data = await api.getCurrent(body);
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

[**GetCurrentPrivacySettingReply**](GetCurrentPrivacySettingReply.md)

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


## updateCurrent

> UpdateCurrentPrivacySettingReply updateCurrent(updateCurrentPrivacySettingRequest)



更新当前登录账号的隐私设置

### Example

```ts
import {
  Configuration,
  PrivacySettingService,
} from '@bass/bbs-sdk-fetch';
import type { UpdateCurrentRequest } from '@bass/bbs-sdk-fetch';

async function example() {
  console.log("🚀 Testing @bass/bbs-sdk-fetch SDK...");
  const api = new PrivacySettingService();

  const body = {
    // UpdateCurrentPrivacySettingRequest
    updateCurrentPrivacySettingRequest: ...,
  } satisfies UpdateCurrentRequest;

  try {
    const data = await api.updateCurrent(body);
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
| **updateCurrentPrivacySettingRequest** | [UpdateCurrentPrivacySettingRequest](UpdateCurrentPrivacySettingRequest.md) |  | |

### Return type

[**UpdateCurrentPrivacySettingReply**](UpdateCurrentPrivacySettingReply.md)

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

