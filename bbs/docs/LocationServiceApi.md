# LocationServiceApi

All URIs are relative to *http://localhost*

| Method | HTTP request | Description |
|------------- | ------------- | -------------|
| [**locationServiceGetCurrent**](LocationServiceApi.md#locationServiceGetCurrent) | **POST** /v1/user/location/get-current |  |
| [**locationServiceGetCurrentWithHttpInfo**](LocationServiceApi.md#locationServiceGetCurrentWithHttpInfo) | **POST** /v1/user/location/get-current |  |
| [**locationServiceUpsertCurrent**](LocationServiceApi.md#locationServiceUpsertCurrent) | **POST** /v1/user/location/upsert-current |  |
| [**locationServiceUpsertCurrentWithHttpInfo**](LocationServiceApi.md#locationServiceUpsertCurrentWithHttpInfo) | **POST** /v1/user/location/upsert-current |  |



## locationServiceGetCurrent

> GetCurrentLocationReply locationServiceGetCurrent(locationServiceGetCurrentRequest)



获取当前登录账号的地理资料

### Example

```java
// Import classes:
import com.bass.bbs.ApiClient;
import com.bass.bbs.ApiException;
import com.bass.bbs.Configuration;
import com.bass.bbs.models.*;
import com.bass.bbs.api.LocationServiceApi;
import com.bass.bbs.api.LocationServiceApi.*;

public class Example {
    public static void main(String[] args) {
        ApiClient defaultClient = Configuration.getDefaultApiClient();
        defaultClient.setBasePath("http://localhost");

        LocationServiceApi apiInstance = new LocationServiceApi(defaultClient);
        Object body = null; // Object | 
        try {
            APIlocationServiceGetCurrentRequest request = APIlocationServiceGetCurrentRequest.newBuilder()
                .body(body)
                .build();
            GetCurrentLocationReply result = apiInstance.locationServiceGetCurrent(request);
            System.out.println(result);
        } catch (ApiException e) {
            System.err.println("Exception when calling LocationServiceApi#locationServiceGetCurrent");
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
| locationServiceGetCurrentRequest | [**APIlocationServiceGetCurrentRequest**](LocationServiceApi.md#APIlocationServiceGetCurrentRequest)|-|-|

### Return type

[**GetCurrentLocationReply**](GetCurrentLocationReply.md)


### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

### HTTP response details
| Status code | Description | Response headers |
|-------------|-------------|------------------|
| **200** | OK |  -  |

## locationServiceGetCurrentWithHttpInfo

> ApiResponse<GetCurrentLocationReply> locationServiceGetCurrentWithHttpInfo(locationServiceGetCurrentRequest)



获取当前登录账号的地理资料

### Example

```java
// Import classes:
import com.bass.bbs.ApiClient;
import com.bass.bbs.ApiException;
import com.bass.bbs.ApiResponse;
import com.bass.bbs.Configuration;
import com.bass.bbs.models.*;
import com.bass.bbs.api.LocationServiceApi;
import com.bass.bbs.api.LocationServiceApi.*;

public class Example {
    public static void main(String[] args) {
        ApiClient defaultClient = Configuration.getDefaultApiClient();
        defaultClient.setBasePath("http://localhost");

        LocationServiceApi apiInstance = new LocationServiceApi(defaultClient);
        Object body = null; // Object | 
        try {
            APIlocationServiceGetCurrentRequest request = APIlocationServiceGetCurrentRequest.newBuilder()
                .body(body)
                .build();
            ApiResponse<GetCurrentLocationReply> response = apiInstance.locationServiceGetCurrentWithHttpInfo(request);
            System.out.println("Status code: " + response.getStatusCode());
            System.out.println("Response headers: " + response.getHeaders());
            System.out.println("Response body: " + response.getData());
        } catch (ApiException e) {
            System.err.println("Exception when calling LocationServiceApi#locationServiceGetCurrent");
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
| locationServiceGetCurrentRequest | [**APIlocationServiceGetCurrentRequest**](LocationServiceApi.md#APIlocationServiceGetCurrentRequest)|-|-|

### Return type

ApiResponse<[**GetCurrentLocationReply**](GetCurrentLocationReply.md)>


### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

### HTTP response details
| Status code | Description | Response headers |
|-------------|-------------|------------------|
| **200** | OK |  -  |


<a id="APIlocationServiceGetCurrentRequest"></a>
## APIlocationServiceGetCurrentRequest
### Properties

|     Name      |    Type       | Description   |     Notes    |
| ------------- | ------------- | ------------- | -------------|
| **body** | **Object** |  | |



## locationServiceUpsertCurrent

> UpsertCurrentLocationReply locationServiceUpsertCurrent(locationServiceUpsertCurrentRequest)



更新当前登录账号的地理资料

### Example

```java
// Import classes:
import com.bass.bbs.ApiClient;
import com.bass.bbs.ApiException;
import com.bass.bbs.Configuration;
import com.bass.bbs.models.*;
import com.bass.bbs.api.LocationServiceApi;
import com.bass.bbs.api.LocationServiceApi.*;

public class Example {
    public static void main(String[] args) {
        ApiClient defaultClient = Configuration.getDefaultApiClient();
        defaultClient.setBasePath("http://localhost");

        LocationServiceApi apiInstance = new LocationServiceApi(defaultClient);
        UpsertCurrentLocationRequest upsertCurrentLocationRequest = new UpsertCurrentLocationRequest(); // UpsertCurrentLocationRequest | 
        try {
            APIlocationServiceUpsertCurrentRequest request = APIlocationServiceUpsertCurrentRequest.newBuilder()
                .upsertCurrentLocationRequest(upsertCurrentLocationRequest)
                .build();
            UpsertCurrentLocationReply result = apiInstance.locationServiceUpsertCurrent(request);
            System.out.println(result);
        } catch (ApiException e) {
            System.err.println("Exception when calling LocationServiceApi#locationServiceUpsertCurrent");
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
| locationServiceUpsertCurrentRequest | [**APIlocationServiceUpsertCurrentRequest**](LocationServiceApi.md#APIlocationServiceUpsertCurrentRequest)|-|-|

### Return type

[**UpsertCurrentLocationReply**](UpsertCurrentLocationReply.md)


### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

### HTTP response details
| Status code | Description | Response headers |
|-------------|-------------|------------------|
| **200** | OK |  -  |

## locationServiceUpsertCurrentWithHttpInfo

> ApiResponse<UpsertCurrentLocationReply> locationServiceUpsertCurrentWithHttpInfo(locationServiceUpsertCurrentRequest)



更新当前登录账号的地理资料

### Example

```java
// Import classes:
import com.bass.bbs.ApiClient;
import com.bass.bbs.ApiException;
import com.bass.bbs.ApiResponse;
import com.bass.bbs.Configuration;
import com.bass.bbs.models.*;
import com.bass.bbs.api.LocationServiceApi;
import com.bass.bbs.api.LocationServiceApi.*;

public class Example {
    public static void main(String[] args) {
        ApiClient defaultClient = Configuration.getDefaultApiClient();
        defaultClient.setBasePath("http://localhost");

        LocationServiceApi apiInstance = new LocationServiceApi(defaultClient);
        UpsertCurrentLocationRequest upsertCurrentLocationRequest = new UpsertCurrentLocationRequest(); // UpsertCurrentLocationRequest | 
        try {
            APIlocationServiceUpsertCurrentRequest request = APIlocationServiceUpsertCurrentRequest.newBuilder()
                .upsertCurrentLocationRequest(upsertCurrentLocationRequest)
                .build();
            ApiResponse<UpsertCurrentLocationReply> response = apiInstance.locationServiceUpsertCurrentWithHttpInfo(request);
            System.out.println("Status code: " + response.getStatusCode());
            System.out.println("Response headers: " + response.getHeaders());
            System.out.println("Response body: " + response.getData());
        } catch (ApiException e) {
            System.err.println("Exception when calling LocationServiceApi#locationServiceUpsertCurrent");
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
| locationServiceUpsertCurrentRequest | [**APIlocationServiceUpsertCurrentRequest**](LocationServiceApi.md#APIlocationServiceUpsertCurrentRequest)|-|-|

### Return type

ApiResponse<[**UpsertCurrentLocationReply**](UpsertCurrentLocationReply.md)>


### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

### HTTP response details
| Status code | Description | Response headers |
|-------------|-------------|------------------|
| **200** | OK |  -  |


<a id="APIlocationServiceUpsertCurrentRequest"></a>
## APIlocationServiceUpsertCurrentRequest
### Properties

|     Name      |    Type       | Description   |     Notes    |
| ------------- | ------------- | ------------- | -------------|
| **upsertCurrentLocationRequest** | [**UpsertCurrentLocationRequest**](UpsertCurrentLocationRequest.md) |  | |


