# LocationServiceApi

All URIs are relative to *http://localhost*

| Method | HTTP request | Description |
|------------- | ------------- | -------------|
| [**locationServiceGetCurrent**](LocationServiceApi.md#locationservicegetcurrent) | **POST** /v1/user/location/get-current |  |
| [**locationServiceUpsertCurrent**](LocationServiceApi.md#locationserviceupsertcurrent) | **POST** /v1/user/location/upsert-current |  |



## locationServiceGetCurrent

> GetCurrentLocationReply locationServiceGetCurrent(body)



获取当前登录账号的地理资料

### Example

```ts
import {
  Configuration,
  LocationServiceApi,
} from '@bass/bbs-sdk';
import type { LocationServiceGetCurrentRequest } from '@bass/bbs-sdk';

async function example() {
  console.log("🚀 Testing @bass/bbs-sdk SDK...");
  const api = new LocationServiceApi();

  const body = {
    // object
    body: Object,
  } satisfies LocationServiceGetCurrentRequest;

  try {
    const data = await api.locationServiceGetCurrent(body);
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


## locationServiceUpsertCurrent

> UpsertCurrentLocationReply locationServiceUpsertCurrent(upsertCurrentLocationRequest)



更新当前登录账号的地理资料

### Example

```ts
import {
  Configuration,
  LocationServiceApi,
} from '@bass/bbs-sdk';
import type { LocationServiceUpsertCurrentRequest } from '@bass/bbs-sdk';

async function example() {
  console.log("🚀 Testing @bass/bbs-sdk SDK...");
  const api = new LocationServiceApi();

  const body = {
    // UpsertCurrentLocationRequest
    upsertCurrentLocationRequest: ...,
  } satisfies LocationServiceUpsertCurrentRequest;

  try {
    const data = await api.locationServiceUpsertCurrent(body);
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

