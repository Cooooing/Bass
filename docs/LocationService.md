# LocationService

All URIs are relative to *http://localhost*

| Method | HTTP request | Description |
|------------- | ------------- | -------------|
| [**detectCurrent**](LocationService.md#detectCurrent) | **POST** /v1/user/location/detect-current |  |
| [**detectCurrentWithHttpInfo**](LocationService.md#detectCurrentWithHttpInfo) | **POST** /v1/user/location/detect-current |  |
| [**getCurrent**](LocationService.md#getCurrent) | **POST** /v1/user/location/get-current |  |
| [**getCurrentWithHttpInfo**](LocationService.md#getCurrentWithHttpInfo) | **POST** /v1/user/location/get-current |  |
| [**upsertCurrent**](LocationService.md#upsertCurrent) | **POST** /v1/user/location/upsert-current |  |
| [**upsertCurrentWithHttpInfo**](LocationService.md#upsertCurrentWithHttpInfo) | **POST** /v1/user/location/upsert-current |  |



## detectCurrent

> DetectCurrentLocationResp detectCurrent(detectCurrentRequest)



按当前请求 IP 解析并更新当前账号的地理资料。

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
            APIdetectCurrentRequest request = APIdetectCurrentRequest.newBuilder()
                .body(body)
                .build();
            DetectCurrentLocationResp result = apiInstance.detectCurrent(request);
            System.out.println(result);
        } catch (ApiException e) {
            System.err.println("Exception when calling LocationService#detectCurrent");
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
| detectCurrentRequest | [**APIdetectCurrentRequest**](LocationService.md#APIdetectCurrentRequest)|-|-|

### Return type

[**DetectCurrentLocationResp**](DetectCurrentLocationResp.md)


### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

### HTTP response details
| Status code | Description | Response headers |
|-------------|-------------|------------------|
| **200** | OK |  -  |

## detectCurrentWithHttpInfo

> ApiResponse<DetectCurrentLocationResp> detectCurrentWithHttpInfo(detectCurrentRequest)



按当前请求 IP 解析并更新当前账号的地理资料。

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
            APIdetectCurrentRequest request = APIdetectCurrentRequest.newBuilder()
                .body(body)
                .build();
            ApiResponse<DetectCurrentLocationResp> response = apiInstance.detectCurrentWithHttpInfo(request);
            System.out.println("Status code: " + response.getStatusCode());
            System.out.println("Response headers: " + response.getHeaders());
            System.out.println("Response body: " + response.getData());
        } catch (ApiException e) {
            System.err.println("Exception when calling LocationService#detectCurrent");
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
| detectCurrentRequest | [**APIdetectCurrentRequest**](LocationService.md#APIdetectCurrentRequest)|-|-|

### Return type

ApiResponse<[**DetectCurrentLocationResp**](DetectCurrentLocationResp.md)>


### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

### HTTP response details
| Status code | Description | Response headers |
|-------------|-------------|------------------|
| **200** | OK |  -  |


<a id="APIdetectCurrentRequest"></a>
## APIdetectCurrentRequest
### Properties

|     Name      |    Type       | Description   |     Notes    |
| ------------- | ------------- | ------------- | -------------|
| **body** | **Object** |  | |



## getCurrent

> GetCurrentLocationResp getCurrent(getCurrentRequest)



获取当前账号的地理资料。

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
            APIgetCurrentRequest request = APIgetCurrentRequest.newBuilder()
                .body(body)
                .build();
            GetCurrentLocationResp result = apiInstance.getCurrent(request);
            System.out.println(result);
        } catch (ApiException e) {
            System.err.println("Exception when calling LocationService#getCurrent");
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
| getCurrentRequest | [**APIgetCurrentRequest**](LocationService.md#APIgetCurrentRequest)|-|-|

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

## getCurrentWithHttpInfo

> ApiResponse<GetCurrentLocationResp> getCurrentWithHttpInfo(getCurrentRequest)



获取当前账号的地理资料。

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
            APIgetCurrentRequest request = APIgetCurrentRequest.newBuilder()
                .body(body)
                .build();
            ApiResponse<GetCurrentLocationResp> response = apiInstance.getCurrentWithHttpInfo(request);
            System.out.println("Status code: " + response.getStatusCode());
            System.out.println("Response headers: " + response.getHeaders());
            System.out.println("Response body: " + response.getData());
        } catch (ApiException e) {
            System.err.println("Exception when calling LocationService#getCurrent");
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
| getCurrentRequest | [**APIgetCurrentRequest**](LocationService.md#APIgetCurrentRequest)|-|-|

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


<a id="APIgetCurrentRequest"></a>
## APIgetCurrentRequest
### Properties

|     Name      |    Type       | Description   |     Notes    |
| ------------- | ------------- | ------------- | -------------|
| **body** | **Object** |  | |



## upsertCurrent

> UpsertCurrentLocationResp upsertCurrent(upsertCurrentRequest)



更新当前账号的地理资料。

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
        UpsertCurrentLocationReq upsertCurrentLocationReq = new UpsertCurrentLocationReq(); // UpsertCurrentLocationReq | 
        try {
            APIupsertCurrentRequest request = APIupsertCurrentRequest.newBuilder()
                .upsertCurrentLocationReq(upsertCurrentLocationReq)
                .build();
            UpsertCurrentLocationResp result = apiInstance.upsertCurrent(request);
            System.out.println(result);
        } catch (ApiException e) {
            System.err.println("Exception when calling LocationService#upsertCurrent");
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
| upsertCurrentRequest | [**APIupsertCurrentRequest**](LocationService.md#APIupsertCurrentRequest)|-|-|

### Return type

[**UpsertCurrentLocationResp**](UpsertCurrentLocationResp.md)


### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

### HTTP response details
| Status code | Description | Response headers |
|-------------|-------------|------------------|
| **200** | OK |  -  |

## upsertCurrentWithHttpInfo

> ApiResponse<UpsertCurrentLocationResp> upsertCurrentWithHttpInfo(upsertCurrentRequest)



更新当前账号的地理资料。

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
        UpsertCurrentLocationReq upsertCurrentLocationReq = new UpsertCurrentLocationReq(); // UpsertCurrentLocationReq | 
        try {
            APIupsertCurrentRequest request = APIupsertCurrentRequest.newBuilder()
                .upsertCurrentLocationReq(upsertCurrentLocationReq)
                .build();
            ApiResponse<UpsertCurrentLocationResp> response = apiInstance.upsertCurrentWithHttpInfo(request);
            System.out.println("Status code: " + response.getStatusCode());
            System.out.println("Response headers: " + response.getHeaders());
            System.out.println("Response body: " + response.getData());
        } catch (ApiException e) {
            System.err.println("Exception when calling LocationService#upsertCurrent");
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
| upsertCurrentRequest | [**APIupsertCurrentRequest**](LocationService.md#APIupsertCurrentRequest)|-|-|

### Return type

ApiResponse<[**UpsertCurrentLocationResp**](UpsertCurrentLocationResp.md)>


### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

### HTTP response details
| Status code | Description | Response headers |
|-------------|-------------|------------------|
| **200** | OK |  -  |


<a id="APIupsertCurrentRequest"></a>
## APIupsertCurrentRequest
### Properties

|     Name      |    Type       | Description   |     Notes    |
| ------------- | ------------- | ------------- | -------------|
| **upsertCurrentLocationReq** | [**UpsertCurrentLocationReq**](UpsertCurrentLocationReq.md) |  | |


