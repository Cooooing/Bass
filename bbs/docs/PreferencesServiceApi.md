# PreferencesServiceApi

All URIs are relative to *http://localhost*

| Method | HTTP request | Description |
|------------- | ------------- | -------------|
| [**preferencesServiceGetCurrent**](PreferencesServiceApi.md#preferencesservicegetcurrent) | **POST** /v1/user/preference/get-current |  |
| [**preferencesServiceUpdate**](PreferencesServiceApi.md#preferencesserviceupdate) | **POST** /v1/user/preference/update-current |  |



## preferencesServiceGetCurrent

> CommonApiAppBbsV1UserGetCurrentPreferencesReply preferencesServiceGetCurrent(body)



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

[**CommonApiAppBbsV1UserGetCurrentPreferencesReply**](CommonApiAppBbsV1UserGetCurrentPreferencesReply.md)

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


## preferencesServiceUpdate

> CommonApiAppBbsV1UserUpdatePreferencesReply preferencesServiceUpdate(commonApiAppBbsV1UserUpdatePreferencesRequest)



更新当前登录账号的偏好设置

### Example

```ts
import {
  Configuration,
  PreferencesServiceApi,
} from '@bass/bbs-sdk';
import type { PreferencesServiceUpdateRequest } from '@bass/bbs-sdk';

async function example() {
  console.log("🚀 Testing @bass/bbs-sdk SDK...");
  const api = new PreferencesServiceApi();

  const body = {
    // CommonApiAppBbsV1UserUpdatePreferencesRequest
    commonApiAppBbsV1UserUpdatePreferencesRequest: ...,
  } satisfies PreferencesServiceUpdateRequest;

  try {
    const data = await api.preferencesServiceUpdate(body);
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
| **commonApiAppBbsV1UserUpdatePreferencesRequest** | [CommonApiAppBbsV1UserUpdatePreferencesRequest](CommonApiAppBbsV1UserUpdatePreferencesRequest.md) |  | |

### Return type

[**CommonApiAppBbsV1UserUpdatePreferencesReply**](CommonApiAppBbsV1UserUpdatePreferencesReply.md)

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

