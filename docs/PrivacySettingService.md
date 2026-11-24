# PrivacySettingService

All URIs are relative to *http://localhost*

| Method | HTTP request | Description |
|------------- | ------------- | -------------|
| [**getCurrentPrivacySetting**](PrivacySettingService.md#getCurrentPrivacySetting) | **POST** /v1/user/privacy-setting/get-current |  |
| [**getCurrentPrivacySettingWithHttpInfo**](PrivacySettingService.md#getCurrentPrivacySettingWithHttpInfo) | **POST** /v1/user/privacy-setting/get-current |  |
| [**updateCurrentPrivacySetting**](PrivacySettingService.md#updateCurrentPrivacySetting) | **POST** /v1/user/privacy-setting/update-current |  |
| [**updateCurrentPrivacySettingWithHttpInfo**](PrivacySettingService.md#updateCurrentPrivacySettingWithHttpInfo) | **POST** /v1/user/privacy-setting/update-current |  |



## getCurrentPrivacySetting

> GetCurrentPrivacySettingResp getCurrentPrivacySetting(getCurrentPrivacySettingRequest)



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
            APIgetCurrentPrivacySettingRequest request = APIgetCurrentPrivacySettingRequest.newBuilder()
                .body(body)
                .build();
            GetCurrentPrivacySettingResp result = apiInstance.getCurrentPrivacySetting(request);
            System.out.println(result);
        } catch (ApiException e) {
            System.err.println("Exception when calling PrivacySettingService#getCurrentPrivacySetting");
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
| getCurrentPrivacySettingRequest | [**APIgetCurrentPrivacySettingRequest**](PrivacySettingService.md#APIgetCurrentPrivacySettingRequest)|-|-|

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

## getCurrentPrivacySettingWithHttpInfo

> ApiResponse<GetCurrentPrivacySettingResp> getCurrentPrivacySettingWithHttpInfo(getCurrentPrivacySettingRequest)



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
            APIgetCurrentPrivacySettingRequest request = APIgetCurrentPrivacySettingRequest.newBuilder()
                .body(body)
                .build();
            ApiResponse<GetCurrentPrivacySettingResp> response = apiInstance.getCurrentPrivacySettingWithHttpInfo(request);
            System.out.println("Status code: " + response.getStatusCode());
            System.out.println("Response headers: " + response.getHeaders());
            System.out.println("Response body: " + response.getData());
        } catch (ApiException e) {
            System.err.println("Exception when calling PrivacySettingService#getCurrentPrivacySetting");
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
| getCurrentPrivacySettingRequest | [**APIgetCurrentPrivacySettingRequest**](PrivacySettingService.md#APIgetCurrentPrivacySettingRequest)|-|-|

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


<a id="APIgetCurrentPrivacySettingRequest"></a>
## APIgetCurrentPrivacySettingRequest
### Properties

|     Name      |    Type       | Description   |     Notes    |
| ------------- | ------------- | ------------- | -------------|
| **body** | **Object** |  | |



## updateCurrentPrivacySetting

> UpdateCurrentPrivacySettingResp updateCurrentPrivacySetting(updateCurrentPrivacySettingRequest)



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
            APIupdateCurrentPrivacySettingRequest request = APIupdateCurrentPrivacySettingRequest.newBuilder()
                .updateCurrentPrivacySettingReq(updateCurrentPrivacySettingReq)
                .build();
            UpdateCurrentPrivacySettingResp result = apiInstance.updateCurrentPrivacySetting(request);
            System.out.println(result);
        } catch (ApiException e) {
            System.err.println("Exception when calling PrivacySettingService#updateCurrentPrivacySetting");
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
| updateCurrentPrivacySettingRequest | [**APIupdateCurrentPrivacySettingRequest**](PrivacySettingService.md#APIupdateCurrentPrivacySettingRequest)|-|-|

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

## updateCurrentPrivacySettingWithHttpInfo

> ApiResponse<UpdateCurrentPrivacySettingResp> updateCurrentPrivacySettingWithHttpInfo(updateCurrentPrivacySettingRequest)



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
            APIupdateCurrentPrivacySettingRequest request = APIupdateCurrentPrivacySettingRequest.newBuilder()
                .updateCurrentPrivacySettingReq(updateCurrentPrivacySettingReq)
                .build();
            ApiResponse<UpdateCurrentPrivacySettingResp> response = apiInstance.updateCurrentPrivacySettingWithHttpInfo(request);
            System.out.println("Status code: " + response.getStatusCode());
            System.out.println("Response headers: " + response.getHeaders());
            System.out.println("Response body: " + response.getData());
        } catch (ApiException e) {
            System.err.println("Exception when calling PrivacySettingService#updateCurrentPrivacySetting");
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
| updateCurrentPrivacySettingRequest | [**APIupdateCurrentPrivacySettingRequest**](PrivacySettingService.md#APIupdateCurrentPrivacySettingRequest)|-|-|

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


<a id="APIupdateCurrentPrivacySettingRequest"></a>
## APIupdateCurrentPrivacySettingRequest
### Properties

|     Name      |    Type       | Description   |     Notes    |
| ------------- | ------------- | ------------- | -------------|
| **updateCurrentPrivacySettingReq** | [**UpdateCurrentPrivacySettingReq**](UpdateCurrentPrivacySettingReq.md) |  | |


