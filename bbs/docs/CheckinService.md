# CheckinService

All URIs are relative to *http://localhost*

| Method | HTTP request | Description |
|------------- | ------------- | -------------|
| [**checkIn**](CheckinService.md#checkin) | **POST** /v1/user/checkin/check-in |  |
| [**getOverview**](CheckinService.md#getoverview) | **POST** /v1/user/checkin/get-overview |  |



## checkIn

> CheckInResp checkIn(body)



### Example

```ts
import {
  Configuration,
  CheckinService,
} from '@bass/bbs-sdk-fetch';
import type { CheckInRequest } from '@bass/bbs-sdk-fetch';

async function example() {
  console.log("🚀 Testing @bass/bbs-sdk-fetch SDK...");
  const api = new CheckinService();

  const body = {
    // object
    body: Object,
  } satisfies CheckInRequest;

  try {
    const data = await api.checkIn(body);
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

[**CheckInResp**](CheckInResp.md)

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


## getOverview

> GetCheckinOverviewResp getOverview(getCheckinOverviewReq)



### Example

```ts
import {
  Configuration,
  CheckinService,
} from '@bass/bbs-sdk-fetch';
import type { GetOverviewRequest } from '@bass/bbs-sdk-fetch';

async function example() {
  console.log("🚀 Testing @bass/bbs-sdk-fetch SDK...");
  const api = new CheckinService();

  const body = {
    // GetCheckinOverviewReq
    getCheckinOverviewReq: ...,
  } satisfies GetOverviewRequest;

  try {
    const data = await api.getOverview(body);
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
| **getCheckinOverviewReq** | [GetCheckinOverviewReq](GetCheckinOverviewReq.md) |  | |

### Return type

[**GetCheckinOverviewResp**](GetCheckinOverviewResp.md)

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

