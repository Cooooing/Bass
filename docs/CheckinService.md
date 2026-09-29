# CheckinService

All URIs are relative to *http://localhost*

| Method | HTTP request | Description |
|------------- | ------------- | -------------|
| [**checkIn**](CheckinService.md#checkIn) | **POST** /v1/user/checkin/check-in |  |
| [**checkInWithHttpInfo**](CheckinService.md#checkInWithHttpInfo) | **POST** /v1/user/checkin/check-in |  |
| [**getOverview**](CheckinService.md#getOverview) | **POST** /v1/user/checkin/get-overview |  |
| [**getOverviewWithHttpInfo**](CheckinService.md#getOverviewWithHttpInfo) | **POST** /v1/user/checkin/get-overview |  |



## checkIn

> CheckInResp checkIn(checkInRequest)



### Example

```java
// Import classes:
import com.bass.bbs.ApiClient;
import com.bass.bbs.ApiException;
import com.bass.bbs.Configuration;
import com.bass.bbs.models.*;
import com.bass.bbs.api.CheckinService;
import com.bass.bbs.api.CheckinService.*;

public class Example {
    public static void main(String[] args) {
        ApiClient defaultClient = Configuration.getDefaultApiClient();
        defaultClient.setBasePath("http://localhost");

        CheckinService apiInstance = new CheckinService(defaultClient);
        Object body = null; // Object | 
        try {
            APIcheckInRequest request = APIcheckInRequest.newBuilder()
                .body(body)
                .build();
            CheckInResp result = apiInstance.checkIn(request);
            System.out.println(result);
        } catch (ApiException e) {
            System.err.println("Exception when calling CheckinService#checkIn");
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
| checkInRequest | [**APIcheckInRequest**](CheckinService.md#APIcheckInRequest)|-|-|

### Return type

[**CheckInResp**](CheckInResp.md)


### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

### HTTP response details
| Status code | Description | Response headers |
|-------------|-------------|------------------|
| **200** | OK |  -  |

## checkInWithHttpInfo

> ApiResponse<CheckInResp> checkInWithHttpInfo(checkInRequest)



### Example

```java
// Import classes:
import com.bass.bbs.ApiClient;
import com.bass.bbs.ApiException;
import com.bass.bbs.ApiResponse;
import com.bass.bbs.Configuration;
import com.bass.bbs.models.*;
import com.bass.bbs.api.CheckinService;
import com.bass.bbs.api.CheckinService.*;

public class Example {
    public static void main(String[] args) {
        ApiClient defaultClient = Configuration.getDefaultApiClient();
        defaultClient.setBasePath("http://localhost");

        CheckinService apiInstance = new CheckinService(defaultClient);
        Object body = null; // Object | 
        try {
            APIcheckInRequest request = APIcheckInRequest.newBuilder()
                .body(body)
                .build();
            ApiResponse<CheckInResp> response = apiInstance.checkInWithHttpInfo(request);
            System.out.println("Status code: " + response.getStatusCode());
            System.out.println("Response headers: " + response.getHeaders());
            System.out.println("Response body: " + response.getData());
        } catch (ApiException e) {
            System.err.println("Exception when calling CheckinService#checkIn");
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
| checkInRequest | [**APIcheckInRequest**](CheckinService.md#APIcheckInRequest)|-|-|

### Return type

ApiResponse<[**CheckInResp**](CheckInResp.md)>


### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

### HTTP response details
| Status code | Description | Response headers |
|-------------|-------------|------------------|
| **200** | OK |  -  |


<a id="APIcheckInRequest"></a>
## APIcheckInRequest
### Properties

|     Name      |    Type       | Description   |     Notes    |
| ------------- | ------------- | ------------- | -------------|
| **body** | **Object** |  | |



## getOverview

> GetCheckinOverviewResp getOverview(getOverviewRequest)



### Example

```java
// Import classes:
import com.bass.bbs.ApiClient;
import com.bass.bbs.ApiException;
import com.bass.bbs.Configuration;
import com.bass.bbs.models.*;
import com.bass.bbs.api.CheckinService;
import com.bass.bbs.api.CheckinService.*;

public class Example {
    public static void main(String[] args) {
        ApiClient defaultClient = Configuration.getDefaultApiClient();
        defaultClient.setBasePath("http://localhost");

        CheckinService apiInstance = new CheckinService(defaultClient);
        GetCheckinOverviewReq getCheckinOverviewReq = new GetCheckinOverviewReq(); // GetCheckinOverviewReq | 
        try {
            APIgetOverviewRequest request = APIgetOverviewRequest.newBuilder()
                .getCheckinOverviewReq(getCheckinOverviewReq)
                .build();
            GetCheckinOverviewResp result = apiInstance.getOverview(request);
            System.out.println(result);
        } catch (ApiException e) {
            System.err.println("Exception when calling CheckinService#getOverview");
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
| getOverviewRequest | [**APIgetOverviewRequest**](CheckinService.md#APIgetOverviewRequest)|-|-|

### Return type

[**GetCheckinOverviewResp**](GetCheckinOverviewResp.md)


### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

### HTTP response details
| Status code | Description | Response headers |
|-------------|-------------|------------------|
| **200** | OK |  -  |

## getOverviewWithHttpInfo

> ApiResponse<GetCheckinOverviewResp> getOverviewWithHttpInfo(getOverviewRequest)



### Example

```java
// Import classes:
import com.bass.bbs.ApiClient;
import com.bass.bbs.ApiException;
import com.bass.bbs.ApiResponse;
import com.bass.bbs.Configuration;
import com.bass.bbs.models.*;
import com.bass.bbs.api.CheckinService;
import com.bass.bbs.api.CheckinService.*;

public class Example {
    public static void main(String[] args) {
        ApiClient defaultClient = Configuration.getDefaultApiClient();
        defaultClient.setBasePath("http://localhost");

        CheckinService apiInstance = new CheckinService(defaultClient);
        GetCheckinOverviewReq getCheckinOverviewReq = new GetCheckinOverviewReq(); // GetCheckinOverviewReq | 
        try {
            APIgetOverviewRequest request = APIgetOverviewRequest.newBuilder()
                .getCheckinOverviewReq(getCheckinOverviewReq)
                .build();
            ApiResponse<GetCheckinOverviewResp> response = apiInstance.getOverviewWithHttpInfo(request);
            System.out.println("Status code: " + response.getStatusCode());
            System.out.println("Response headers: " + response.getHeaders());
            System.out.println("Response body: " + response.getData());
        } catch (ApiException e) {
            System.err.println("Exception when calling CheckinService#getOverview");
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
| getOverviewRequest | [**APIgetOverviewRequest**](CheckinService.md#APIgetOverviewRequest)|-|-|

### Return type

ApiResponse<[**GetCheckinOverviewResp**](GetCheckinOverviewResp.md)>


### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

### HTTP response details
| Status code | Description | Response headers |
|-------------|-------------|------------------|
| **200** | OK |  -  |


<a id="APIgetOverviewRequest"></a>
## APIgetOverviewRequest
### Properties

|     Name      |    Type       | Description   |     Notes    |
| ------------- | ------------- | ------------- | -------------|
| **getCheckinOverviewReq** | [**GetCheckinOverviewReq**](GetCheckinOverviewReq.md) |  | |


