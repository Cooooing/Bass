# LocationService

All URIs are relative to *http://localhost*

| Method | HTTP request | Description |
|------------- | ------------- | -------------|
| [**getCurrentLocation**](LocationService.md#getCurrentLocation) | **POST** /v1/platform/location/get-current |  |
| [**getCurrentLocationWithHttpInfo**](LocationService.md#getCurrentLocationWithHttpInfo) | **POST** /v1/platform/location/get-current |  |



## getCurrentLocation

> GetCurrentLocationResp getCurrentLocation(getCurrentLocationRequest)



按当前请求来源 IP 获取地理数据。

### Example

```java
// Import classes:
import com.bass.bbs.ApiClient;
import com.bass.bbs.ApiException;
import com.bass.bbs.Configuration;
import com.bass.bbs.models.*;
import com.bass.bbs.api.LocationService;
import com.bass.bbs.api.LocationService.*;

public class Example {
    public static void main(String[] args) {
        ApiClient defaultClient = Configuration.getDefaultApiClient();
        defaultClient.setBasePath("http://localhost");

        LocationService apiInstance = new LocationService(defaultClient);
        Object body = null; // Object | 
        try {
            APIgetCurrentLocationRequest request = APIgetCurrentLocationRequest.newBuilder()
                .body(body)
                .build();
            GetCurrentLocationResp result = apiInstance.getCurrentLocation(request);
            System.out.println(result);
        } catch (ApiException e) {
            System.err.println("Exception when calling LocationService#getCurrentLocation");
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
| getCurrentLocationRequest | [**APIgetCurrentLocationRequest**](LocationService.md#APIgetCurrentLocationRequest)|-|-|

### Return type

[**GetCurrentLocationResp**](GetCurrentLocationResp.md)


### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

### HTTP response details
| Status code | Description | Response headers |
|-------------|-------------|------------------|
| **200** | OK |  -  |

## getCurrentLocationWithHttpInfo

> ApiResponse<GetCurrentLocationResp> getCurrentLocationWithHttpInfo(getCurrentLocationRequest)



按当前请求来源 IP 获取地理数据。

### Example

```java
// Import classes:
import com.bass.bbs.ApiClient;
import com.bass.bbs.ApiException;
import com.bass.bbs.ApiResponse;
import com.bass.bbs.Configuration;
import com.bass.bbs.models.*;
import com.bass.bbs.api.LocationService;
import com.bass.bbs.api.LocationService.*;

public class Example {
    public static void main(String[] args) {
        ApiClient defaultClient = Configuration.getDefaultApiClient();
        defaultClient.setBasePath("http://localhost");

        LocationService apiInstance = new LocationService(defaultClient);
        Object body = null; // Object | 
        try {
            APIgetCurrentLocationRequest request = APIgetCurrentLocationRequest.newBuilder()
                .body(body)
                .build();
            ApiResponse<GetCurrentLocationResp> response = apiInstance.getCurrentLocationWithHttpInfo(request);
            System.out.println("Status code: " + response.getStatusCode());
            System.out.println("Response headers: " + response.getHeaders());
            System.out.println("Response body: " + response.getData());
        } catch (ApiException e) {
            System.err.println("Exception when calling LocationService#getCurrentLocation");
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
| getCurrentLocationRequest | [**APIgetCurrentLocationRequest**](LocationService.md#APIgetCurrentLocationRequest)|-|-|

### Return type

ApiResponse<[**GetCurrentLocationResp**](GetCurrentLocationResp.md)>


### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

### HTTP response details
| Status code | Description | Response headers |
|-------------|-------------|------------------|
| **200** | OK |  -  |


<a id="APIgetCurrentLocationRequest"></a>
## APIgetCurrentLocationRequest
### Properties

|     Name      |    Type       | Description   |     Notes    |
| ------------- | ------------- | ------------- | -------------|
| **body** | **Object** |  | |


