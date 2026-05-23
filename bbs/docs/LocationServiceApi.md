# LocationServiceApi

All URIs are relative to *http://localhost*

| Method | HTTP request | Description |
|------------- | ------------- | -------------|
| [**locationServiceGetCurrent**](LocationServiceApi.md#locationservicegetcurrent) | **POST** /v1/user/location/get-current |  |
| [**locationServiceUpsert**](LocationServiceApi.md#locationserviceupsert) | **POST** /v1/user/location/upsert-current |  |



## locationServiceGetCurrent

> CommonApiAppBbsV1UserGetCurrentLocationReply locationServiceGetCurrent(body)



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

[**CommonApiAppBbsV1UserGetCurrentLocationReply**](CommonApiAppBbsV1UserGetCurrentLocationReply.md)

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


## locationServiceUpsert

> CommonApiAppBbsV1UserUpsertLocationReply locationServiceUpsert(commonApiAppBbsV1UserUpsertLocationRequest)



更新当前登录账号的地理资料

### Example

```ts
import {
  Configuration,
  LocationServiceApi,
} from '@bass/bbs-sdk';
import type { LocationServiceUpsertRequest } from '@bass/bbs-sdk';

async function example() {
  console.log("🚀 Testing @bass/bbs-sdk SDK...");
  const api = new LocationServiceApi();

  const body = {
    // CommonApiAppBbsV1UserUpsertLocationRequest
    commonApiAppBbsV1UserUpsertLocationRequest: ...,
  } satisfies LocationServiceUpsertRequest;

  try {
    const data = await api.locationServiceUpsert(body);
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
| **commonApiAppBbsV1UserUpsertLocationRequest** | [CommonApiAppBbsV1UserUpsertLocationRequest](CommonApiAppBbsV1UserUpsertLocationRequest.md) |  | |

### Return type

[**CommonApiAppBbsV1UserUpsertLocationReply**](CommonApiAppBbsV1UserUpsertLocationReply.md)

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

