# LivenessService

All URIs are relative to *http://localhost*

|Method | HTTP request | Description|
|------------- | ------------- | -------------|
|[**claimYesterdayLivenessReward**](#claimyesterdaylivenessreward) | **POST** /v1/user/liveness/yesterday-reward/claim | |
|[**getLiveness**](#getliveness) | **GET** /v1/user/liveness | |

# **claimYesterdayLivenessReward**
> ClaimYesterdayLivenessRewardResp claimYesterdayLivenessReward(body)

服务方法。

### Example

```typescript
import {
    LivenessService,
    Configuration
} from '@bass/bbs-sdk-axios';

const configuration = new Configuration();
const apiInstance = new LivenessService(configuration);

let body: object; //

const { status, data } = await apiInstance.claimYesterdayLivenessReward(
    body
);
```

### Parameters

|Name | Type | Description  | Notes|
|------------- | ------------- | ------------- | -------------|
| **body** | **object**|  | |


### Return type

**ClaimYesterdayLivenessRewardResp**

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

# **getLiveness**
> GetLivenessResp getLiveness()

服务方法。

### Example

```typescript
import {
    LivenessService,
    Configuration
} from '@bass/bbs-sdk-axios';

const configuration = new Configuration();
const apiInstance = new LivenessService(configuration);

const { status, data } = await apiInstance.getLiveness();
```

### Parameters
This endpoint does not have any parameters.


### Return type

**GetLivenessResp**

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json


### HTTP response details
| Status code | Description | Response headers |
|-------------|-------------|------------------|
|**200** | OK |  -  |

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

