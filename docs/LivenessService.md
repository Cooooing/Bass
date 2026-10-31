# LivenessService

All URIs are relative to *http://localhost*

| Method | HTTP request | Description |
|------------- | ------------- | -------------|
| [**claimYesterdayReward**](LivenessService.md#claimYesterdayReward) | **POST** /v1/user/liveness/yesterday-reward/claim |  |
| [**claimYesterdayRewardWithHttpInfo**](LivenessService.md#claimYesterdayRewardWithHttpInfo) | **POST** /v1/user/liveness/yesterday-reward/claim |  |
| [**getLiveness**](LivenessService.md#getLiveness) | **GET** /v1/user/liveness |  |
| [**getLivenessWithHttpInfo**](LivenessService.md#getLivenessWithHttpInfo) | **GET** /v1/user/liveness |  |



## claimYesterdayReward

> ClaimYesterdayLivenessRewardResp claimYesterdayReward(claimYesterdayRewardRequest)



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
            APIclaimYesterdayRewardRequest request = APIclaimYesterdayRewardRequest.newBuilder()
                .body(body)
                .build();
            ClaimYesterdayLivenessRewardResp result = apiInstance.claimYesterdayReward(request);
            System.out.println(result);
        } catch (ApiException e) {
            System.err.println("Exception when calling LivenessService#claimYesterdayReward");
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
| claimYesterdayRewardRequest | [**APIclaimYesterdayRewardRequest**](LivenessService.md#APIclaimYesterdayRewardRequest)|-|-|

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

## claimYesterdayRewardWithHttpInfo

> ApiResponse<ClaimYesterdayLivenessRewardResp> claimYesterdayRewardWithHttpInfo(claimYesterdayRewardRequest)



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
            APIclaimYesterdayRewardRequest request = APIclaimYesterdayRewardRequest.newBuilder()
                .body(body)
                .build();
            ApiResponse<ClaimYesterdayLivenessRewardResp> response = apiInstance.claimYesterdayRewardWithHttpInfo(request);
            System.out.println("Status code: " + response.getStatusCode());
            System.out.println("Response headers: " + response.getHeaders());
            System.out.println("Response body: " + response.getData());
        } catch (ApiException e) {
            System.err.println("Exception when calling LivenessService#claimYesterdayReward");
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
| claimYesterdayRewardRequest | [**APIclaimYesterdayRewardRequest**](LivenessService.md#APIclaimYesterdayRewardRequest)|-|-|

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


<a id="APIclaimYesterdayRewardRequest"></a>
## APIclaimYesterdayRewardRequest
### Properties

|     Name      |    Type       | Description   |     Notes    |
| ------------- | ------------- | ------------- | -------------|
| **body** | **Object** |  | |



## getLiveness

> GetLivenessResp getLiveness()



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

