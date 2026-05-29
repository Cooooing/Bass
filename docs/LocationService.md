# LocationService

All URIs are relative to *http://localhost*

| Method | HTTP request | Description |
|------------- | ------------- | -------------|
| [**getCurrent**](LocationService.md#getcurrent) | **POST** /v1/user/location/get-current |  |
| [**upsertCurrent**](LocationService.md#upsertcurrent) | **POST** /v1/user/location/upsert-current |  |



## getCurrent

> GetCurrentLocationReply getCurrent(body)



获取当前登录账号的地理资料

### Example

```ts
import {
  Configuration,
  LocationService,
} from '@bass/bbs-sdk-fetch';
import type { GetCurrentRequest } from '@bass/bbs-sdk-fetch';

async function example() {
  console.log("🚀 Testing @bass/bbs-sdk-fetch SDK...");
  const api = new LocationService();

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

[**GetCurrentLocationReply**](GetCurrentLocationReply.md)

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


## upsertCurrent

> UpsertCurrentLocationReply upsertCurrent(upsertCurrentLocationRequest)



更新当前登录账号的地理资料

### Example

```ts
import {
  Configuration,
  LocationService,
} from '@bass/bbs-sdk-fetch';
import type { UpsertCurrentRequest } from '@bass/bbs-sdk-fetch';

async function example() {
  console.log("🚀 Testing @bass/bbs-sdk-fetch SDK...");
  const api = new LocationService();

  const body = {
    // UpsertCurrentLocationRequest
    upsertCurrentLocationRequest: ...,
  } satisfies UpsertCurrentRequest;

  try {
    const data = await api.upsertCurrent(body);
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
| **upsertCurrentLocationRequest** | [UpsertCurrentLocationRequest](UpsertCurrentLocationRequest.md) |  | |

### Return type

[**UpsertCurrentLocationReply**](UpsertCurrentLocationReply.md)

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

