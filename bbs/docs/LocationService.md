# LocationService

All URIs are relative to *http://localhost*

|Method | HTTP request | Description|
|------------- | ------------- | -------------|
|[**getCurrent**](#getcurrent) | **POST** /v1/user/location/get-current | |
|[**upsertCurrent**](#upsertcurrent) | **POST** /v1/user/location/upsert-current | |

# **getCurrent**
> GetCurrentLocationResp getCurrent(body)

获取当前账号的地理资料。

### Example

```typescript
import {
    LocationService,
    Configuration
} from '@bass/bbs-sdk-axios';

const configuration = new Configuration();
const apiInstance = new LocationService(configuration);

let body: object; //

const { status, data } = await apiInstance.getCurrent(
    body
);
```

### Parameters

|Name | Type | Description  | Notes|
|------------- | ------------- | ------------- | -------------|
| **body** | **object**|  | |


### Return type

**GetCurrentLocationResp**

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

# **upsertCurrent**
> UpsertCurrentLocationResp upsertCurrent(upsertCurrentLocationReq)

更新当前账号的地理资料。

### Example

```typescript
import {
    LocationService,
    Configuration,
    UpsertCurrentLocationReq
} from '@bass/bbs-sdk-axios';

const configuration = new Configuration();
const apiInstance = new LocationService(configuration);

let upsertCurrentLocationReq: UpsertCurrentLocationReq; //

const { status, data } = await apiInstance.upsertCurrent(
    upsertCurrentLocationReq
);
```

### Parameters

|Name | Type | Description  | Notes|
|------------- | ------------- | ------------- | -------------|
| **upsertCurrentLocationReq** | **UpsertCurrentLocationReq**|  | |


### Return type

**UpsertCurrentLocationResp**

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

