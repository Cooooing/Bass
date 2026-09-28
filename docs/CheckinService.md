# CheckinService

All URIs are relative to *http://localhost*

|Method | HTTP request | Description|
|------------- | ------------- | -------------|
|[**checkIn**](#checkin) | **POST** /v1/user/checkin/check-in | |
|[**getOverview**](#getoverview) | **POST** /v1/user/checkin/get-overview | |

# **checkIn**
> CheckInResp checkIn(body)


### Example

```typescript
import {
    CheckinService,
    Configuration
} from '@bass/bbs-sdk-axios';

const configuration = new Configuration();
const apiInstance = new CheckinService(configuration);

let body: object; //

const { status, data } = await apiInstance.checkIn(
    body
);
```

### Parameters

|Name | Type | Description  | Notes|
|------------- | ------------- | ------------- | -------------|
| **body** | **object**|  | |


### Return type

**CheckInResp**

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

# **getOverview**
> GetCheckinOverviewResp getOverview(getCheckinOverviewReq)


### Example

```typescript
import {
    CheckinService,
    Configuration,
    GetCheckinOverviewReq
} from '@bass/bbs-sdk-axios';

const configuration = new Configuration();
const apiInstance = new CheckinService(configuration);

let getCheckinOverviewReq: GetCheckinOverviewReq; //

const { status, data } = await apiInstance.getOverview(
    getCheckinOverviewReq
);
```

### Parameters

|Name | Type | Description  | Notes|
|------------- | ------------- | ------------- | -------------|
| **getCheckinOverviewReq** | **GetCheckinOverviewReq**|  | |


### Return type

**GetCheckinOverviewResp**

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

