# PreferencesService

All URIs are relative to *http://localhost*

| Method | HTTP request | Description |
|------------- | ------------- | -------------|
| [**getCurrent**](PreferencesService.md#getcurrent) | **POST** /v1/user/preference/get-current |  |
| [**updateCurrent**](PreferencesService.md#updatecurrent) | **POST** /v1/user/preference/update-current |  |



## getCurrent

> GetCurrentPreferencesReply getCurrent(body)



获取当前登录账号的偏好设置

### Example

```ts
import {
  Configuration,
  PreferencesService,
} from '@bass/bbs-sdk-fetch';
import type { GetCurrentRequest } from '@bass/bbs-sdk-fetch';

async function example() {
  console.log("🚀 Testing @bass/bbs-sdk-fetch SDK...");
  const api = new PreferencesService();

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

[**GetCurrentPreferencesReply**](GetCurrentPreferencesReply.md)

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

> UpdateCurrentPreferencesReply updateCurrent(updateCurrentPreferencesRequest)



更新当前登录账号的偏好设置

### Example

```ts
import {
  Configuration,
  PreferencesService,
} from '@bass/bbs-sdk-fetch';
import type { UpdateCurrentRequest } from '@bass/bbs-sdk-fetch';

async function example() {
  console.log("🚀 Testing @bass/bbs-sdk-fetch SDK...");
  const api = new PreferencesService();

  const body = {
    // UpdateCurrentPreferencesRequest
    updateCurrentPreferencesRequest: ...,
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
| **updateCurrentPreferencesRequest** | [UpdateCurrentPreferencesRequest](UpdateCurrentPreferencesRequest.md) |  | |

### Return type

[**UpdateCurrentPreferencesReply**](UpdateCurrentPreferencesReply.md)

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

