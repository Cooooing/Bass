# PreferencesServiceApi

All URIs are relative to *http://localhost*

| Method | HTTP request | Description |
|------------- | ------------- | -------------|
| [**preferencesServiceGetCurrent**](PreferencesServiceApi.md#preferencesservicegetcurrent) | **POST** /v1/user/preference/get-current |  |
| [**preferencesServiceUpdateCurrent**](PreferencesServiceApi.md#preferencesserviceupdatecurrent) | **POST** /v1/user/preference/update-current |  |



## preferencesServiceGetCurrent

> GetCurrentPreferencesReply preferencesServiceGetCurrent(body)



获取当前登录账号的偏好设置

### Example

```ts
import {
  Configuration,
  PreferencesServiceApi,
} from '@bass/bbs-sdk';
import type { PreferencesServiceGetCurrentRequest } from '@bass/bbs-sdk';

async function example() {
  console.log("🚀 Testing @bass/bbs-sdk SDK...");
  const api = new PreferencesServiceApi();

  const body = {
    // object
    body: Object,
  } satisfies PreferencesServiceGetCurrentRequest;

  try {
    const data = await api.preferencesServiceGetCurrent(body);
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


## preferencesServiceUpdateCurrent

> UpdateCurrentPreferencesReply preferencesServiceUpdateCurrent(updateCurrentPreferencesRequest)



更新当前登录账号的偏好设置

### Example

```ts
import {
  Configuration,
  PreferencesServiceApi,
} from '@bass/bbs-sdk';
import type { PreferencesServiceUpdateCurrentRequest } from '@bass/bbs-sdk';

async function example() {
  console.log("🚀 Testing @bass/bbs-sdk SDK...");
  const api = new PreferencesServiceApi();

  const body = {
    // UpdateCurrentPreferencesRequest
    updateCurrentPreferencesRequest: ...,
  } satisfies PreferencesServiceUpdateCurrentRequest;

  try {
    const data = await api.preferencesServiceUpdateCurrent(body);
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

