# HolidayService

All URIs are relative to *http://localhost*

| Method | HTTP request | Description |
|------------- | ------------- | -------------|
| [**getCurrent**](HolidayService.md#getcurrent) | **POST** /v1/platform/holiday/get-current |  |



## getCurrent

> GetCurrentHolidayResp getCurrent(body)



获取当前节假日状态。

### Example

```ts
import {
  Configuration,
  HolidayService,
} from '@bass/bbs-sdk-fetch';
import type { GetCurrentRequest } from '@bass/bbs-sdk-fetch';

async function example() {
  console.log("🚀 Testing @bass/bbs-sdk-fetch SDK...");
  const api = new HolidayService();

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

[**GetCurrentHolidayResp**](GetCurrentHolidayResp.md)

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

