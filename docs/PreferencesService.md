# PreferencesService

All URIs are relative to *http://localhost*

| Method | HTTP request | Description |
|------------- | ------------- | -------------|
| [**getCurrentPreferences**](PreferencesService.md#getcurrentpreferences) | **POST** /v1/user/preference/get-current |  |
| [**updateCurrentPreferences**](PreferencesService.md#updatecurrentpreferences) | **POST** /v1/user/preference/update-current |  |



## getCurrentPreferences

> GetCurrentPreferencesResp getCurrentPreferences(body)



获取当前账号的偏好设置。

### Example

```ts
import {
  Configuration,
  PreferencesService,
} from '@bass/bbs-sdk-fetch';
import type { GetCurrentPreferencesRequest } from '@bass/bbs-sdk-fetch';

async function example() {
  console.log("🚀 Testing @bass/bbs-sdk-fetch SDK...");
  const api = new PreferencesService();

  const body = {
    // object
    body: Object,
  } satisfies GetCurrentPreferencesRequest;

  try {
    const data = await api.getCurrentPreferences(body);
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

[**GetCurrentPreferencesResp**](GetCurrentPreferencesResp.md)

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


## updateCurrentPreferences

> UpdateCurrentPreferencesResp updateCurrentPreferences(updateCurrentPreferencesReq)



更新当前账号的偏好设置。

### Example

```ts
import {
  Configuration,
  PreferencesService,
} from '@bass/bbs-sdk-fetch';
import type { UpdateCurrentPreferencesRequest } from '@bass/bbs-sdk-fetch';

async function example() {
  console.log("🚀 Testing @bass/bbs-sdk-fetch SDK...");
  const api = new PreferencesService();

  const body = {
    // UpdateCurrentPreferencesReq
    updateCurrentPreferencesReq: ...,
  } satisfies UpdateCurrentPreferencesRequest;

  try {
    const data = await api.updateCurrentPreferences(body);
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
| **updateCurrentPreferencesReq** | [UpdateCurrentPreferencesReq](UpdateCurrentPreferencesReq.md) |  | |

### Return type

[**UpdateCurrentPreferencesResp**](UpdateCurrentPreferencesResp.md)

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

