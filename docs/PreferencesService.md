# PreferencesService

All URIs are relative to *http://localhost*

| Method | HTTP request | Description |
|------------- | ------------- | -------------|
| [**getCurrentPreferences**](PreferencesService.md#getCurrentPreferences) | **POST** /v1/user/preference/get-current |  |
| [**getCurrentPreferencesWithHttpInfo**](PreferencesService.md#getCurrentPreferencesWithHttpInfo) | **POST** /v1/user/preference/get-current |  |
| [**updateCurrentPreferences**](PreferencesService.md#updateCurrentPreferences) | **POST** /v1/user/preference/update-current |  |
| [**updateCurrentPreferencesWithHttpInfo**](PreferencesService.md#updateCurrentPreferencesWithHttpInfo) | **POST** /v1/user/preference/update-current |  |



## getCurrentPreferences

> GetCurrentPreferencesResp getCurrentPreferences(getCurrentPreferencesRequest)



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
            APIgetCurrentPreferencesRequest request = APIgetCurrentPreferencesRequest.newBuilder()
                .body(body)
                .build();
            GetCurrentPreferencesResp result = apiInstance.getCurrentPreferences(request);
            System.out.println(result);
        } catch (ApiException e) {
            System.err.println("Exception when calling PreferencesService#getCurrentPreferences");
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
| getCurrentPreferencesRequest | [**APIgetCurrentPreferencesRequest**](PreferencesService.md#APIgetCurrentPreferencesRequest)|-|-|

### Return type

[**GetCurrentPreferencesResp**](GetCurrentPreferencesResp.md)


### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

### HTTP response details
| Status code | Description | Response headers |
|-------------|-------------|------------------|
| **200** | OK |  -  |

## getCurrentPreferencesWithHttpInfo

> ApiResponse<GetCurrentPreferencesResp> getCurrentPreferencesWithHttpInfo(getCurrentPreferencesRequest)



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
            APIgetCurrentPreferencesRequest request = APIgetCurrentPreferencesRequest.newBuilder()
                .body(body)
                .build();
            ApiResponse<GetCurrentPreferencesResp> response = apiInstance.getCurrentPreferencesWithHttpInfo(request);
            System.out.println("Status code: " + response.getStatusCode());
            System.out.println("Response headers: " + response.getHeaders());
            System.out.println("Response body: " + response.getData());
        } catch (ApiException e) {
            System.err.println("Exception when calling PreferencesService#getCurrentPreferences");
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
| getCurrentPreferencesRequest | [**APIgetCurrentPreferencesRequest**](PreferencesService.md#APIgetCurrentPreferencesRequest)|-|-|

### Return type

ApiResponse<[**GetCurrentPreferencesResp**](GetCurrentPreferencesResp.md)>


### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

### HTTP response details
| Status code | Description | Response headers |
|-------------|-------------|------------------|
| **200** | OK |  -  |


<a id="APIgetCurrentPreferencesRequest"></a>
## APIgetCurrentPreferencesRequest
### Properties

|     Name      |    Type       | Description   |     Notes    |
| ------------- | ------------- | ------------- | -------------|
| **body** | **Object** |  | |



## updateCurrentPreferences

> UpdateCurrentPreferencesResp updateCurrentPreferences(updateCurrentPreferencesRequest)



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
        UpdateCurrentPreferencesReq updateCurrentPreferencesReq = new UpdateCurrentPreferencesReq(); // UpdateCurrentPreferencesReq | 
        try {
            APIupdateCurrentPreferencesRequest request = APIupdateCurrentPreferencesRequest.newBuilder()
                .updateCurrentPreferencesReq(updateCurrentPreferencesReq)
                .build();
            UpdateCurrentPreferencesResp result = apiInstance.updateCurrentPreferences(request);
            System.out.println(result);
        } catch (ApiException e) {
            System.err.println("Exception when calling PreferencesService#updateCurrentPreferences");
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
| updateCurrentPreferencesRequest | [**APIupdateCurrentPreferencesRequest**](PreferencesService.md#APIupdateCurrentPreferencesRequest)|-|-|

### Return type

[**UpdateCurrentPreferencesResp**](UpdateCurrentPreferencesResp.md)


### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

### HTTP response details
| Status code | Description | Response headers |
|-------------|-------------|------------------|
| **200** | OK |  -  |

## updateCurrentPreferencesWithHttpInfo

> ApiResponse<UpdateCurrentPreferencesResp> updateCurrentPreferencesWithHttpInfo(updateCurrentPreferencesRequest)



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
        UpdateCurrentPreferencesReq updateCurrentPreferencesReq = new UpdateCurrentPreferencesReq(); // UpdateCurrentPreferencesReq | 
        try {
            APIupdateCurrentPreferencesRequest request = APIupdateCurrentPreferencesRequest.newBuilder()
                .updateCurrentPreferencesReq(updateCurrentPreferencesReq)
                .build();
            ApiResponse<UpdateCurrentPreferencesResp> response = apiInstance.updateCurrentPreferencesWithHttpInfo(request);
            System.out.println("Status code: " + response.getStatusCode());
            System.out.println("Response headers: " + response.getHeaders());
            System.out.println("Response body: " + response.getData());
        } catch (ApiException e) {
            System.err.println("Exception when calling PreferencesService#updateCurrentPreferences");
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
| updateCurrentPreferencesRequest | [**APIupdateCurrentPreferencesRequest**](PreferencesService.md#APIupdateCurrentPreferencesRequest)|-|-|

### Return type

ApiResponse<[**UpdateCurrentPreferencesResp**](UpdateCurrentPreferencesResp.md)>


### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

### HTTP response details
| Status code | Description | Response headers |
|-------------|-------------|------------------|
| **200** | OK |  -  |


<a id="APIupdateCurrentPreferencesRequest"></a>
## APIupdateCurrentPreferencesRequest
### Properties

|     Name      |    Type       | Description   |     Notes    |
| ------------- | ------------- | ------------- | -------------|
| **updateCurrentPreferencesReq** | [**UpdateCurrentPreferencesReq**](UpdateCurrentPreferencesReq.md) |  | |


