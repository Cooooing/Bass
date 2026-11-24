# LivenessService

All URIs are relative to *http://localhost*

| Method | HTTP request | Description |
|------------- | ------------- | -------------|
| [**claimYesterdayLivenessReward**](LivenessService.md#claimYesterdayLivenessReward) | **POST** /v1/user/liveness/yesterday-reward/claim |  |
| [**claimYesterdayLivenessRewardWithHttpInfo**](LivenessService.md#claimYesterdayLivenessRewardWithHttpInfo) | **POST** /v1/user/liveness/yesterday-reward/claim |  |
| [**getLiveness**](LivenessService.md#getLiveness) | **GET** /v1/user/liveness |  |
| [**getLivenessWithHttpInfo**](LivenessService.md#getLivenessWithHttpInfo) | **GET** /v1/user/liveness |  |



## claimYesterdayLivenessReward

> ClaimYesterdayLivenessRewardResp claimYesterdayLivenessReward(claimYesterdayLivenessRewardRequest)



服务方法。

### Example

```java
// Import classes:
import com.bass.bbs.ApiClient;
import com.bass.bbs.ApiException;
import com.bass.bbs.Configuration;
import com.bass.bbs.models.*;
import com.bass.bbs.api.LivenessService;
import com.bass.bbs.api.LivenessService.*;

public class Example {
    public static void main(String[] args) {
        ApiClient defaultClient = Configuration.getDefaultApiClient();
        defaultClient.setBasePath("http://localhost");

        LivenessService apiInstance = new LivenessService(defaultClient);
        Object body = null; // Object | 
        try {
            APIclaimYesterdayLivenessRewardRequest request = APIclaimYesterdayLivenessRewardRequest.newBuilder()
                .body(body)
                .build();
            ClaimYesterdayLivenessRewardResp result = apiInstance.claimYesterdayLivenessReward(request);
            System.out.println(result);
        } catch (ApiException e) {
            System.err.println("Exception when calling LivenessService#claimYesterdayLivenessReward");
            System.err.println("Status code: " + e.getCode());
            System.err.println("Reason: " + e.getResponseBody());
            System.err.println("Response headers: " + e.getResponseHeaders());
            e.printStackTrace();
        }
    }
}
```

### Parameters

|    Name      |    Type       | Description   |     Notes    |
|------------- | ------------- | ------------- | -------------|
| claimYesterdayLivenessRewardRequest | [**APIclaimYesterdayLivenessRewardRequest**](LivenessService.md#APIclaimYesterdayLivenessRewardRequest)|-|-|

### Return type

[**ClaimYesterdayLivenessRewardResp**](ClaimYesterdayLivenessRewardResp.md)


### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

### HTTP response details
| Status code | Description | Response headers |
|-------------|-------------|------------------|
| **200** | OK |  -  |

## claimYesterdayLivenessRewardWithHttpInfo

> ApiResponse<ClaimYesterdayLivenessRewardResp> claimYesterdayLivenessRewardWithHttpInfo(claimYesterdayLivenessRewardRequest)



服务方法。

### Example

```java
// Import classes:
import com.bass.bbs.ApiClient;
import com.bass.bbs.ApiException;
import com.bass.bbs.ApiResponse;
import com.bass.bbs.Configuration;
import com.bass.bbs.models.*;
import com.bass.bbs.api.LivenessService;
import com.bass.bbs.api.LivenessService.*;

public class Example {
    public static void main(String[] args) {
        ApiClient defaultClient = Configuration.getDefaultApiClient();
        defaultClient.setBasePath("http://localhost");

        LivenessService apiInstance = new LivenessService(defaultClient);
        Object body = null; // Object | 
        try {
            APIclaimYesterdayLivenessRewardRequest request = APIclaimYesterdayLivenessRewardRequest.newBuilder()
                .body(body)
                .build();
            ApiResponse<ClaimYesterdayLivenessRewardResp> response = apiInstance.claimYesterdayLivenessRewardWithHttpInfo(request);
            System.out.println("Status code: " + response.getStatusCode());
            System.out.println("Response headers: " + response.getHeaders());
            System.out.println("Response body: " + response.getData());
        } catch (ApiException e) {
            System.err.println("Exception when calling LivenessService#claimYesterdayLivenessReward");
            System.err.println("Status code: " + e.getCode());
            System.err.println("Response headers: " + e.getResponseHeaders());
            System.err.println("Reason: " + e.getResponseBody());
            e.printStackTrace();
        }
    }
}
```

### Parameters

|    Name      |    Type       | Description   |     Notes    |
|------------- | ------------- | ------------- | -------------|
| claimYesterdayLivenessRewardRequest | [**APIclaimYesterdayLivenessRewardRequest**](LivenessService.md#APIclaimYesterdayLivenessRewardRequest)|-|-|

### Return type

ApiResponse<[**ClaimYesterdayLivenessRewardResp**](ClaimYesterdayLivenessRewardResp.md)>


### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

### HTTP response details
| Status code | Description | Response headers |
|-------------|-------------|------------------|
| **200** | OK |  -  |


<a id="APIclaimYesterdayLivenessRewardRequest"></a>
## APIclaimYesterdayLivenessRewardRequest
### Properties

|     Name      |    Type       | Description   |     Notes    |
| ------------- | ------------- | ------------- | -------------|
| **body** | **Object** |  | |



## getLiveness

> GetLivenessResp getLiveness()



服务方法。

### Example

```java
// Import classes:
import com.bass.bbs.ApiClient;
import com.bass.bbs.ApiException;
import com.bass.bbs.Configuration;
import com.bass.bbs.models.*;
import com.bass.bbs.api.LivenessService;
import com.bass.bbs.api.LivenessService.*;

public class Example {
    public static void main(String[] args) {
        ApiClient defaultClient = Configuration.getDefaultApiClient();
        defaultClient.setBasePath("http://localhost");

        LivenessService apiInstance = new LivenessService(defaultClient);
        try {
            GetLivenessResp result = apiInstance.getLiveness();
            System.out.println(result);
        } catch (ApiException e) {
            System.err.println("Exception when calling LivenessService#getLiveness");
            System.err.println("Status code: " + e.getCode());
            System.err.println("Reason: " + e.getResponseBody());
            System.err.println("Response headers: " + e.getResponseHeaders());
            e.printStackTrace();
        }
    }
}
```

### Parameters

This endpoint does not need any parameter.

### Return type

[**GetLivenessResp**](GetLivenessResp.md)


### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

### HTTP response details
| Status code | Description | Response headers |
|-------------|-------------|------------------|
| **200** | OK |  -  |

## getLivenessWithHttpInfo

> ApiResponse<GetLivenessResp> getLivenessWithHttpInfo()



服务方法。

### Example

```java
// Import classes:
import com.bass.bbs.ApiClient;
import com.bass.bbs.ApiException;
import com.bass.bbs.ApiResponse;
import com.bass.bbs.Configuration;
import com.bass.bbs.models.*;
import com.bass.bbs.api.LivenessService;
import com.bass.bbs.api.LivenessService.*;

public class Example {
    public static void main(String[] args) {
        ApiClient defaultClient = Configuration.getDefaultApiClient();
        defaultClient.setBasePath("http://localhost");

        LivenessService apiInstance = new LivenessService(defaultClient);
        try {
            ApiResponse<GetLivenessResp> response = apiInstance.getLivenessWithHttpInfo();
            System.out.println("Status code: " + response.getStatusCode());
            System.out.println("Response headers: " + response.getHeaders());
            System.out.println("Response body: " + response.getData());
        } catch (ApiException e) {
            System.err.println("Exception when calling LivenessService#getLiveness");
            System.err.println("Status code: " + e.getCode());
            System.err.println("Response headers: " + e.getResponseHeaders());
            System.err.println("Reason: " + e.getResponseBody());
            e.printStackTrace();
        }
    }
}
```

### Parameters

This endpoint does not need any parameter.

### Return type

ApiResponse<[**GetLivenessResp**](GetLivenessResp.md)>


### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

### HTTP response details
| Status code | Description | Response headers |
|-------------|-------------|------------------|
| **200** | OK |  -  |

