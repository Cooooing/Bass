# PrivacySettingService

All URIs are relative to *http://localhost*

| Method | HTTP request | Description |
|------------- | ------------- | -------------|
| [**getCurrent**](PrivacySettingService.md#getCurrent) | **POST** /v1/user/privacy-setting/get-current |  |
| [**getCurrentWithHttpInfo**](PrivacySettingService.md#getCurrentWithHttpInfo) | **POST** /v1/user/privacy-setting/get-current |  |
| [**updateCurrent**](PrivacySettingService.md#updateCurrent) | **POST** /v1/user/privacy-setting/update-current |  |
| [**updateCurrentWithHttpInfo**](PrivacySettingService.md#updateCurrentWithHttpInfo) | **POST** /v1/user/privacy-setting/update-current |  |



## getCurrent

> GetCurrentPrivacySettingResp getCurrent(getCurrentRequest)



获取当前账号的隐私设置。

### Example

```java
// Import classes:
import com.bass.bbs.ApiClient;
import com.bass.bbs.ApiException;
import com.bass.bbs.Configuration;
import com.bass.bbs.models.*;
import com.bass.bbs.api.PrivacySettingService;
import com.bass.bbs.api.PrivacySettingService.*;

public class Example {
    public static void main(String[] args) {
        ApiClient defaultClient = Configuration.getDefaultApiClient();
        defaultClient.setBasePath("http://localhost");

        PrivacySettingService apiInstance = new PrivacySettingService(defaultClient);
        Object body = null; // Object | 
        try {
            APIgetCurrentRequest request = APIgetCurrentRequest.newBuilder()
                .body(body)
                .build();
            GetCurrentPrivacySettingResp result = apiInstance.getCurrent(request);
            System.out.println(result);
        } catch (ApiException e) {
            System.err.println("Exception when calling PrivacySettingService#getCurrent");
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
| getCurrentRequest | [**APIgetCurrentRequest**](PrivacySettingService.md#APIgetCurrentRequest)|-|-|

### Return type

[**GetCurrentPrivacySettingResp**](GetCurrentPrivacySettingResp.md)


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

> ApiResponse<GetCurrentPrivacySettingResp> getCurrentWithHttpInfo(getCurrentRequest)



获取当前账号的隐私设置。

### Example

```java
// Import classes:
import com.bass.bbs.ApiClient;
import com.bass.bbs.ApiException;
import com.bass.bbs.ApiResponse;
import com.bass.bbs.Configuration;
import com.bass.bbs.models.*;
import com.bass.bbs.api.PrivacySettingService;
import com.bass.bbs.api.PrivacySettingService.*;

public class Example {
    public static void main(String[] args) {
        ApiClient defaultClient = Configuration.getDefaultApiClient();
        defaultClient.setBasePath("http://localhost");

        PrivacySettingService apiInstance = new PrivacySettingService(defaultClient);
        Object body = null; // Object | 
        try {
            APIgetCurrentRequest request = APIgetCurrentRequest.newBuilder()
                .body(body)
                .build();
            ApiResponse<GetCurrentPrivacySettingResp> response = apiInstance.getCurrentWithHttpInfo(request);
            System.out.println("Status code: " + response.getStatusCode());
            System.out.println("Response headers: " + response.getHeaders());
            System.out.println("Response body: " + response.getData());
        } catch (ApiException e) {
            System.err.println("Exception when calling PrivacySettingService#getCurrent");
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
| getCurrentRequest | [**APIgetCurrentRequest**](PrivacySettingService.md#APIgetCurrentRequest)|-|-|

### Return type

ApiResponse<[**GetCurrentPrivacySettingResp**](GetCurrentPrivacySettingResp.md)>


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

> UpdateCurrentPrivacySettingResp updateCurrent(updateCurrentRequest)



更新当前账号的隐私设置。

### Example

```java
// Import classes:
import com.bass.bbs.ApiClient;
import com.bass.bbs.ApiException;
import com.bass.bbs.Configuration;
import com.bass.bbs.models.*;
import com.bass.bbs.api.PrivacySettingService;
import com.bass.bbs.api.PrivacySettingService.*;

public class Example {
    public static void main(String[] args) {
        ApiClient defaultClient = Configuration.getDefaultApiClient();
        defaultClient.setBasePath("http://localhost");

        PrivacySettingService apiInstance = new PrivacySettingService(defaultClient);
        UpdateCurrentPrivacySettingReq updateCurrentPrivacySettingReq = new UpdateCurrentPrivacySettingReq(); // UpdateCurrentPrivacySettingReq | 
        try {
            APIupdateCurrentRequest request = APIupdateCurrentRequest.newBuilder()
                .updateCurrentPrivacySettingReq(updateCurrentPrivacySettingReq)
                .build();
            UpdateCurrentPrivacySettingResp result = apiInstance.updateCurrent(request);
            System.out.println(result);
        } catch (ApiException e) {
            System.err.println("Exception when calling PrivacySettingService#updateCurrent");
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
| updateCurrentRequest | [**APIupdateCurrentRequest**](PrivacySettingService.md#APIupdateCurrentRequest)|-|-|

### Return type

[**UpdateCurrentPrivacySettingResp**](UpdateCurrentPrivacySettingResp.md)


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

> ApiResponse<UpdateCurrentPrivacySettingResp> updateCurrentWithHttpInfo(updateCurrentRequest)



更新当前账号的隐私设置。

### Example

```java
// Import classes:
import com.bass.bbs.ApiClient;
import com.bass.bbs.ApiException;
import com.bass.bbs.ApiResponse;
import com.bass.bbs.Configuration;
import com.bass.bbs.models.*;
import com.bass.bbs.api.PrivacySettingService;
import com.bass.bbs.api.PrivacySettingService.*;

public class Example {
    public static void main(String[] args) {
        ApiClient defaultClient = Configuration.getDefaultApiClient();
        defaultClient.setBasePath("http://localhost");

        PrivacySettingService apiInstance = new PrivacySettingService(defaultClient);
        UpdateCurrentPrivacySettingReq updateCurrentPrivacySettingReq = new UpdateCurrentPrivacySettingReq(); // UpdateCurrentPrivacySettingReq | 
        try {
            APIupdateCurrentRequest request = APIupdateCurrentRequest.newBuilder()
                .updateCurrentPrivacySettingReq(updateCurrentPrivacySettingReq)
                .build();
            ApiResponse<UpdateCurrentPrivacySettingResp> response = apiInstance.updateCurrentWithHttpInfo(request);
            System.out.println("Status code: " + response.getStatusCode());
            System.out.println("Response headers: " + response.getHeaders());
            System.out.println("Response body: " + response.getData());
        } catch (ApiException e) {
            System.err.println("Exception when calling PrivacySettingService#updateCurrent");
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
| updateCurrentRequest | [**APIupdateCurrentRequest**](PrivacySettingService.md#APIupdateCurrentRequest)|-|-|

### Return type

ApiResponse<[**UpdateCurrentPrivacySettingResp**](UpdateCurrentPrivacySettingResp.md)>


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
| **updateCurrentPrivacySettingReq** | [**UpdateCurrentPrivacySettingReq**](UpdateCurrentPrivacySettingReq.md) |  | |


