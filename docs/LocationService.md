# LocationService

All URIs are relative to *http://localhost*

|Method | HTTP request | Description|
|------------- | ------------- | -------------|
|[**getCurrent**](#getcurrent) | **POST** /v1/user/location/get-current | |
|[**upsertCurrent**](#upsertcurrent) | **POST** /v1/user/location/upsert-current | |

# **getCurrent**
> GetCurrentLocationReply getCurrent(body)

获取当前登录账号的地理资料

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

**GetCurrentLocationReply**

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
> UpsertCurrentLocationReply upsertCurrent(upsertCurrentLocationRequest)

更新当前登录账号的地理资料

### Example

```typescript
import {
    LocationService,
    Configuration,
    UpsertCurrentLocationRequest
} from '@bass/bbs-sdk-axios';

const configuration = new Configuration();
const apiInstance = new LocationService(configuration);

let upsertCurrentLocationRequest: UpsertCurrentLocationRequest; //

const { status, data } = await apiInstance.upsertCurrent(
    upsertCurrentLocationRequest
);
```

### Parameters

|Name | Type | Description  | Notes|
|------------- | ------------- | ------------- | -------------|
| **upsertCurrentLocationRequest** | **UpsertCurrentLocationRequest**|  | |


### Return type

**UpsertCurrentLocationReply**

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

