# PreferencesService

All URIs are relative to *http://localhost*

| Method | HTTP request | Description |
|------------- | ------------- | -------------|
| [**getCurrent**](PreferencesService.md#getCurrent) | **POST** /v1/user/preference/get-current |  |
| [**getCurrentWithHttpInfo**](PreferencesService.md#getCurrentWithHttpInfo) | **POST** /v1/user/preference/get-current |  |
| [**updateCurrent**](PreferencesService.md#updateCurrent) | **POST** /v1/user/preference/update-current |  |
| [**updateCurrentWithHttpInfo**](PreferencesService.md#updateCurrentWithHttpInfo) | **POST** /v1/user/preference/update-current |  |



## getCurrent

> GetCurrentPreferencesReply getCurrent(getCurrentRequest)



获取当前账号的偏好设置。

### Example

```java
// Import classes:
import com.bass.bbs.ApiClient;
import com.bass.bbs.ApiException;
import com.bass.bbs.Configuration;
import com.bass.bbs.models.*;
import com.bass.bbs.api.PreferencesService;
import com.bass.bbs.api.PreferencesService.*;

public class Example {
    public static void main(String[] args) {
        ApiClient defaultClient = Configuration.getDefaultApiClient();
        defaultClient.setBasePath("http://localhost");

        PreferencesService apiInstance = new PreferencesService(defaultClient);
        Object body = null; // Object | 
        try {
            APIgetCurrentRequest request = APIgetCurrentRequest.newBuilder()
                .body(body)
                .build();
            GetCurrentPreferencesReply result = apiInstance.getCurrent(request);
            System.out.println(result);
        } catch (ApiException e) {
            System.err.println("Exception when calling PreferencesService#getCurrent");
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
| getCurrentRequest | [**APIgetCurrentRequest**](PreferencesService.md#APIgetCurrentRequest)|-|-|

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

## getCurrentWithHttpInfo

> ApiResponse<GetCurrentPreferencesReply> getCurrentWithHttpInfo(getCurrentRequest)



获取当前账号的偏好设置。

### Example

```java
// Import classes:
import com.bass.bbs.ApiClient;
import com.bass.bbs.ApiException;
import com.bass.bbs.ApiResponse;
import com.bass.bbs.Configuration;
import com.bass.bbs.models.*;
import com.bass.bbs.api.PreferencesService;
import com.bass.bbs.api.PreferencesService.*;

public class Example {
    public static void main(String[] args) {
        ApiClient defaultClient = Configuration.getDefaultApiClient();
        defaultClient.setBasePath("http://localhost");

        PreferencesService apiInstance = new PreferencesService(defaultClient);
        Object body = null; // Object | 
        try {
            APIgetCurrentRequest request = APIgetCurrentRequest.newBuilder()
                .body(body)
                .build();
            ApiResponse<GetCurrentPreferencesReply> response = apiInstance.getCurrentWithHttpInfo(request);
            System.out.println("Status code: " + response.getStatusCode());
            System.out.println("Response headers: " + response.getHeaders());
            System.out.println("Response body: " + response.getData());
        } catch (ApiException e) {
            System.err.println("Exception when calling PreferencesService#getCurrent");
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
| getCurrentRequest | [**APIgetCurrentRequest**](PreferencesService.md#APIgetCurrentRequest)|-|-|

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


<a id="APIgetCurrentRequest"></a>
## APIgetCurrentRequest
### Properties

|     Name      |    Type       | Description   |     Notes    |
| ------------- | ------------- | ------------- | -------------|
| **body** | **Object** |  | |



## updateCurrent

> UpdateCurrentPreferencesReply updateCurrent(updateCurrentRequest)



更新当前账号的偏好设置。

### Example

```java
// Import classes:
import com.bass.bbs.ApiClient;
import com.bass.bbs.ApiException;
import com.bass.bbs.Configuration;
import com.bass.bbs.models.*;
import com.bass.bbs.api.PreferencesService;
import com.bass.bbs.api.PreferencesService.*;

public class Example {
    public static void main(String[] args) {
        ApiClient defaultClient = Configuration.getDefaultApiClient();
        defaultClient.setBasePath("http://localhost");

        PreferencesService apiInstance = new PreferencesService(defaultClient);
        UpdateCurrentPreferencesRequest updateCurrentPreferencesRequest = new UpdateCurrentPreferencesRequest(); // UpdateCurrentPreferencesRequest | 
        try {
            APIupdateCurrentRequest request = APIupdateCurrentRequest.newBuilder()
                .updateCurrentPreferencesRequest(updateCurrentPreferencesRequest)
                .build();
            UpdateCurrentPreferencesReply result = apiInstance.updateCurrent(request);
            System.out.println(result);
        } catch (ApiException e) {
            System.err.println("Exception when calling PreferencesService#updateCurrent");
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
| updateCurrentRequest | [**APIupdateCurrentRequest**](PreferencesService.md#APIupdateCurrentRequest)|-|-|

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

## updateCurrentWithHttpInfo

> ApiResponse<UpdateCurrentPreferencesReply> updateCurrentWithHttpInfo(updateCurrentRequest)



更新当前账号的偏好设置。

### Example

```java
// Import classes:
import com.bass.bbs.ApiClient;
import com.bass.bbs.ApiException;
import com.bass.bbs.ApiResponse;
import com.bass.bbs.Configuration;
import com.bass.bbs.models.*;
import com.bass.bbs.api.PreferencesService;
import com.bass.bbs.api.PreferencesService.*;

public class Example {
    public static void main(String[] args) {
        ApiClient defaultClient = Configuration.getDefaultApiClient();
        defaultClient.setBasePath("http://localhost");

        PreferencesService apiInstance = new PreferencesService(defaultClient);
        UpdateCurrentPreferencesRequest updateCurrentPreferencesRequest = new UpdateCurrentPreferencesRequest(); // UpdateCurrentPreferencesRequest | 
        try {
            APIupdateCurrentRequest request = APIupdateCurrentRequest.newBuilder()
                .updateCurrentPreferencesRequest(updateCurrentPreferencesRequest)
                .build();
            ApiResponse<UpdateCurrentPreferencesReply> response = apiInstance.updateCurrentWithHttpInfo(request);
            System.out.println("Status code: " + response.getStatusCode());
            System.out.println("Response headers: " + response.getHeaders());
            System.out.println("Response body: " + response.getData());
        } catch (ApiException e) {
            System.err.println("Exception when calling PreferencesService#updateCurrent");
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
| updateCurrentRequest | [**APIupdateCurrentRequest**](PreferencesService.md#APIupdateCurrentRequest)|-|-|

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


<a id="APIupdateCurrentRequest"></a>
## APIupdateCurrentRequest
### Properties

|     Name      |    Type       | Description   |     Notes    |
| ------------- | ------------- | ------------- | -------------|
| **updateCurrentPreferencesRequest** | [**UpdateCurrentPreferencesRequest**](UpdateCurrentPreferencesRequest.md) |  | |


