# HolidayService

All URIs are relative to *http://localhost*

|Method | HTTP request | Description|
|------------- | ------------- | -------------|
|[**getCurrentHoliday**](#getcurrentholiday) | **POST** /v1/platform/holiday/get-current | |

# **getCurrentHoliday**
> GetCurrentHolidayResp getCurrentHoliday(body)

获取当前节假日状态。

### Example

```typescript
import {
    HolidayService,
    Configuration
} from '@bass/bbs-sdk-axios';

const configuration = new Configuration();
const apiInstance = new HolidayService(configuration);

let body: object; //

const { status, data } = await apiInstance.getCurrentHoliday(
    body
);
```

### Parameters

|Name | Type | Description  | Notes|
|------------- | ------------- | ------------- | -------------|
| **body** | **object**|  | |


### Return type

**GetCurrentHolidayResp**

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

