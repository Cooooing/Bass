# PreferencesServiceApi

All URIs are relative to *http://localhost*

| Method | HTTP request | Description |
|------------- | ------------- | -------------|
| [**preferencesServiceGetCurrent**](PreferencesServiceApi.md#preferencesServiceGetCurrent) | **POST** /v1/user/preference/get-current |  |
| [**preferencesServiceGetCurrentWithHttpInfo**](PreferencesServiceApi.md#preferencesServiceGetCurrentWithHttpInfo) | **POST** /v1/user/preference/get-current |  |
| [**preferencesServiceUpdateCurrent**](PreferencesServiceApi.md#preferencesServiceUpdateCurrent) | **POST** /v1/user/preference/update-current |  |
| [**preferencesServiceUpdateCurrentWithHttpInfo**](PreferencesServiceApi.md#preferencesServiceUpdateCurrentWithHttpInfo) | **POST** /v1/user/preference/update-current |  |



## preferencesServiceGetCurrent

> GetCurrentPreferencesReply preferencesServiceGetCurrent(preferencesServiceGetCurrentRequest)



获取当前登录账号的偏好设置

### Example

```java
// Import classes:
import com.bass.bbs.ApiClient;
import com.bass.bbs.ApiException;
import com.bass.bbs.Configuration;
import com.bass.bbs.models.*;
import com.bass.bbs.api.PreferencesServiceApi;
import com.bass.bbs.api.PreferencesServiceApi.*;

public class Example {
    public static void main(String[] args) {
        ApiClient defaultClient = Configuration.getDefaultApiClient();
        defaultClient.setBasePath("http://localhost");

        PreferencesServiceApi apiInstance = new PreferencesServiceApi(defaultClient);
        Object body = null; // Object | 
        try {
            APIpreferencesServiceGetCurrentRequest request = APIpreferencesServiceGetCurrentRequest.newBuilder()
                .body(body)
                .build();
            GetCurrentPreferencesReply result = apiInstance.preferencesServiceGetCurrent(request);
            System.out.println(result);
        } catch (ApiException e) {
            System.err.println("Exception when calling PreferencesServiceApi#preferencesServiceGetCurrent");
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
| preferencesServiceGetCurrentRequest | [**APIpreferencesServiceGetCurrentRequest**](PreferencesServiceApi.md#APIpreferencesServiceGetCurrentRequest)|-|-|

### Return type

[**GetCurrentPreferencesReply**](GetCurrentPreferencesReply.md)


### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

### HTTP response details
| Status code | Description | Response headers |
|-------------|-------------|------------------|
| **200** | OK |  -  |

## preferencesServiceGetCurrentWithHttpInfo

> ApiResponse<GetCurrentPreferencesReply> preferencesServiceGetCurrentWithHttpInfo(preferencesServiceGetCurrentRequest)



获取当前登录账号的偏好设置

### Example

```java
// Import classes:
import com.bass.bbs.ApiClient;
import com.bass.bbs.ApiException;
import com.bass.bbs.ApiResponse;
import com.bass.bbs.Configuration;
import com.bass.bbs.models.*;
import com.bass.bbs.api.PreferencesServiceApi;
import com.bass.bbs.api.PreferencesServiceApi.*;

public class Example {
    public static void main(String[] args) {
        ApiClient defaultClient = Configuration.getDefaultApiClient();
        defaultClient.setBasePath("http://localhost");

        PreferencesServiceApi apiInstance = new PreferencesServiceApi(defaultClient);
        Object body = null; // Object | 
        try {
            APIpreferencesServiceGetCurrentRequest request = APIpreferencesServiceGetCurrentRequest.newBuilder()
                .body(body)
                .build();
            ApiResponse<GetCurrentPreferencesReply> response = apiInstance.preferencesServiceGetCurrentWithHttpInfo(request);
            System.out.println("Status code: " + response.getStatusCode());
            System.out.println("Response headers: " + response.getHeaders());
            System.out.println("Response body: " + response.getData());
        } catch (ApiException e) {
            System.err.println("Exception when calling PreferencesServiceApi#preferencesServiceGetCurrent");
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
| preferencesServiceGetCurrentRequest | [**APIpreferencesServiceGetCurrentRequest**](PreferencesServiceApi.md#APIpreferencesServiceGetCurrentRequest)|-|-|

### Return type

ApiResponse<[**GetCurrentPreferencesReply**](GetCurrentPreferencesReply.md)>


### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

### HTTP response details
| Status code | Description | Response headers |
|-------------|-------------|------------------|
| **200** | OK |  -  |


<a id="APIpreferencesServiceGetCurrentRequest"></a>
## APIpreferencesServiceGetCurrentRequest
### Properties

|     Name      |    Type       | Description   |     Notes    |
| ------------- | ------------- | ------------- | -------------|
| **body** | **Object** |  | |



## preferencesServiceUpdateCurrent

> UpdateCurrentPreferencesReply preferencesServiceUpdateCurrent(preferencesServiceUpdateCurrentRequest)



更新当前登录账号的偏好设置

### Example

```java
// Import classes:
import com.bass.bbs.ApiClient;
import com.bass.bbs.ApiException;
import com.bass.bbs.Configuration;
import com.bass.bbs.models.*;
import com.bass.bbs.api.PreferencesServiceApi;
import com.bass.bbs.api.PreferencesServiceApi.*;

public class Example {
    public static void main(String[] args) {
        ApiClient defaultClient = Configuration.getDefaultApiClient();
        defaultClient.setBasePath("http://localhost");

        PreferencesServiceApi apiInstance = new PreferencesServiceApi(defaultClient);
        UpdateCurrentPreferencesRequest updateCurrentPreferencesRequest = new UpdateCurrentPreferencesRequest(); // UpdateCurrentPreferencesRequest | 
        try {
            APIpreferencesServiceUpdateCurrentRequest request = APIpreferencesServiceUpdateCurrentRequest.newBuilder()
                .updateCurrentPreferencesRequest(updateCurrentPreferencesRequest)
                .build();
            UpdateCurrentPreferencesReply result = apiInstance.preferencesServiceUpdateCurrent(request);
            System.out.println(result);
        } catch (ApiException e) {
            System.err.println("Exception when calling PreferencesServiceApi#preferencesServiceUpdateCurrent");
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
| preferencesServiceUpdateCurrentRequest | [**APIpreferencesServiceUpdateCurrentRequest**](PreferencesServiceApi.md#APIpreferencesServiceUpdateCurrentRequest)|-|-|

### Return type

[**UpdateCurrentPreferencesReply**](UpdateCurrentPreferencesReply.md)


### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

### HTTP response details
| Status code | Description | Response headers |
|-------------|-------------|------------------|
| **200** | OK |  -  |

## preferencesServiceUpdateCurrentWithHttpInfo

> ApiResponse<UpdateCurrentPreferencesReply> preferencesServiceUpdateCurrentWithHttpInfo(preferencesServiceUpdateCurrentRequest)



更新当前登录账号的偏好设置

### Example

```java
// Import classes:
import com.bass.bbs.ApiClient;
import com.bass.bbs.ApiException;
import com.bass.bbs.ApiResponse;
import com.bass.bbs.Configuration;
import com.bass.bbs.models.*;
import com.bass.bbs.api.PreferencesServiceApi;
import com.bass.bbs.api.PreferencesServiceApi.*;

public class Example {
    public static void main(String[] args) {
        ApiClient defaultClient = Configuration.getDefaultApiClient();
        defaultClient.setBasePath("http://localhost");

        PreferencesServiceApi apiInstance = new PreferencesServiceApi(defaultClient);
        UpdateCurrentPreferencesRequest updateCurrentPreferencesRequest = new UpdateCurrentPreferencesRequest(); // UpdateCurrentPreferencesRequest | 
        try {
            APIpreferencesServiceUpdateCurrentRequest request = APIpreferencesServiceUpdateCurrentRequest.newBuilder()
                .updateCurrentPreferencesRequest(updateCurrentPreferencesRequest)
                .build();
            ApiResponse<UpdateCurrentPreferencesReply> response = apiInstance.preferencesServiceUpdateCurrentWithHttpInfo(request);
            System.out.println("Status code: " + response.getStatusCode());
            System.out.println("Response headers: " + response.getHeaders());
            System.out.println("Response body: " + response.getData());
        } catch (ApiException e) {
            System.err.println("Exception when calling PreferencesServiceApi#preferencesServiceUpdateCurrent");
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
| preferencesServiceUpdateCurrentRequest | [**APIpreferencesServiceUpdateCurrentRequest**](PreferencesServiceApi.md#APIpreferencesServiceUpdateCurrentRequest)|-|-|

### Return type

ApiResponse<[**UpdateCurrentPreferencesReply**](UpdateCurrentPreferencesReply.md)>


### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

### HTTP response details
| Status code | Description | Response headers |
|-------------|-------------|------------------|
| **200** | OK |  -  |


<a id="APIpreferencesServiceUpdateCurrentRequest"></a>
## APIpreferencesServiceUpdateCurrentRequest
### Properties

|     Name      |    Type       | Description   |     Notes    |
| ------------- | ------------- | ------------- | -------------|
| **updateCurrentPreferencesRequest** | [**UpdateCurrentPreferencesRequest**](UpdateCurrentPreferencesRequest.md) |  | |


