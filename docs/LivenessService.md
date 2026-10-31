# LivenessService

All URIs are relative to *http://localhost*

| Method | HTTP request | Description |
|------------- | ------------- | -------------|
| [**claimYesterdayReward**](LivenessService.md#claimyesterdayreward) | **POST** /v1/user/liveness/yesterday-reward/claim |  |
| [**getLiveness**](LivenessService.md#getliveness) | **GET** /v1/user/liveness |  |



## claimYesterdayReward

> ClaimYesterdayLivenessRewardResp claimYesterdayReward(body)



### Example

```ts
import {
  Configuration,
  LivenessService,
} from '@bass/bbs-sdk-fetch';
import type { ClaimYesterdayRewardRequest } from '@bass/bbs-sdk-fetch';

async function example() {
  console.log("🚀 Testing @bass/bbs-sdk-fetch SDK...");
  const api = new LivenessService();

  const body = {
    // object
    body: Object,
  } satisfies ClaimYesterdayRewardRequest;

  try {
    const data = await api.claimYesterdayReward(body);
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

[**ClaimYesterdayLivenessRewardResp**](ClaimYesterdayLivenessRewardResp.md)

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


## getLiveness

> GetLivenessResp getLiveness()



### Example

```ts
import {
  Configuration,
  LivenessService,
} from '@bass/bbs-sdk-fetch';
import type { GetLivenessRequest } from '@bass/bbs-sdk-fetch';

async function example() {
  console.log("🚀 Testing @bass/bbs-sdk-fetch SDK...");
  const api = new LivenessService();

  try {
    const data = await api.getLiveness();
    console.log(data);
  } catch (error) {
    console.error(error);
  }
}

// Run the test
example().catch(console.error);
```

### Parameters

This endpoint does not need any parameter.

### Return type

[**GetLivenessResp**](GetLivenessResp.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: `application/json`


### HTTP response details
| Status code | Description | Response headers |
|-------------|-------------|------------------|
| **200** | OK |  -  |

[[Back to top]](#) [[Back to API list]](../README.md#api-endpoints) [[Back to Model list]](../README.md#models) [[Back to README]](../README.md)

