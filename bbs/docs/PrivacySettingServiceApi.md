# PrivacySettingServiceApi

All URIs are relative to *http://localhost*

| Method | HTTP request | Description |
|------------- | ------------- | -------------|
| [**privacySettingServiceGetCurrent**](PrivacySettingServiceApi.md#privacySettingServiceGetCurrent) | **POST** /v1/user/privacy-setting/get-current |  |
| [**privacySettingServiceGetCurrentWithHttpInfo**](PrivacySettingServiceApi.md#privacySettingServiceGetCurrentWithHttpInfo) | **POST** /v1/user/privacy-setting/get-current |  |
| [**privacySettingServiceUpdateCurrent**](PrivacySettingServiceApi.md#privacySettingServiceUpdateCurrent) | **POST** /v1/user/privacy-setting/update-current |  |
| [**privacySettingServiceUpdateCurrentWithHttpInfo**](PrivacySettingServiceApi.md#privacySettingServiceUpdateCurrentWithHttpInfo) | **POST** /v1/user/privacy-setting/update-current |  |



## privacySettingServiceGetCurrent

> GetCurrentPrivacySettingReply privacySettingServiceGetCurrent(privacySettingServiceGetCurrentRequest)



获取当前登录账号的隐私设置

### Example

```java
// Import classes:
import com.bass.bbs.ApiClient;
import com.bass.bbs.ApiException;
import com.bass.bbs.Configuration;
import com.bass.bbs.models.*;
import com.bass.bbs.api.PrivacySettingServiceApi;
import com.bass.bbs.api.PrivacySettingServiceApi.*;

public class Example {
    public static void main(String[] args) {
        ApiClient defaultClient = Configuration.getDefaultApiClient();
        defaultClient.setBasePath("http://localhost");

        PrivacySettingServiceApi apiInstance = new PrivacySettingServiceApi(defaultClient);
        Object body = null; // Object | 
        try {
            APIprivacySettingServiceGetCurrentRequest request = APIprivacySettingServiceGetCurrentRequest.newBuilder()
                .body(body)
                .build();
            GetCurrentPrivacySettingReply result = apiInstance.privacySettingServiceGetCurrent(request);
            System.out.println(result);
        } catch (ApiException e) {
            System.err.println("Exception when calling PrivacySettingServiceApi#privacySettingServiceGetCurrent");
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
| privacySettingServiceGetCurrentRequest | [**APIprivacySettingServiceGetCurrentRequest**](PrivacySettingServiceApi.md#APIprivacySettingServiceGetCurrentRequest)|-|-|

### Return type

[**GetCurrentPrivacySettingReply**](GetCurrentPrivacySettingReply.md)


### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

### HTTP response details
| Status code | Description | Response headers |
|-------------|-------------|------------------|
| **200** | OK |  -  |

## privacySettingServiceGetCurrentWithHttpInfo

> ApiResponse<GetCurrentPrivacySettingReply> privacySettingServiceGetCurrentWithHttpInfo(privacySettingServiceGetCurrentRequest)



获取当前登录账号的隐私设置

### Example

```java
// Import classes:
import com.bass.bbs.ApiClient;
import com.bass.bbs.ApiException;
import com.bass.bbs.ApiResponse;
import com.bass.bbs.Configuration;
import com.bass.bbs.models.*;
import com.bass.bbs.api.PrivacySettingServiceApi;
import com.bass.bbs.api.PrivacySettingServiceApi.*;

public class Example {
    public static void main(String[] args) {
        ApiClient defaultClient = Configuration.getDefaultApiClient();
        defaultClient.setBasePath("http://localhost");

        PrivacySettingServiceApi apiInstance = new PrivacySettingServiceApi(defaultClient);
        Object body = null; // Object | 
        try {
            APIprivacySettingServiceGetCurrentRequest request = APIprivacySettingServiceGetCurrentRequest.newBuilder()
                .body(body)
                .build();
            ApiResponse<GetCurrentPrivacySettingReply> response = apiInstance.privacySettingServiceGetCurrentWithHttpInfo(request);
            System.out.println("Status code: " + response.getStatusCode());
            System.out.println("Response headers: " + response.getHeaders());
            System.out.println("Response body: " + response.getData());
        } catch (ApiException e) {
            System.err.println("Exception when calling PrivacySettingServiceApi#privacySettingServiceGetCurrent");
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
| privacySettingServiceGetCurrentRequest | [**APIprivacySettingServiceGetCurrentRequest**](PrivacySettingServiceApi.md#APIprivacySettingServiceGetCurrentRequest)|-|-|

### Return type

ApiResponse<[**GetCurrentPrivacySettingReply**](GetCurrentPrivacySettingReply.md)>


### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

### HTTP response details
| Status code | Description | Response headers |
|-------------|-------------|------------------|
| **200** | OK |  -  |


<a id="APIprivacySettingServiceGetCurrentRequest"></a>
## APIprivacySettingServiceGetCurrentRequest
### Properties

|     Name      |    Type       | Description   |     Notes    |
| ------------- | ------------- | ------------- | -------------|
| **body** | **Object** |  | |



## privacySettingServiceUpdateCurrent

> UpdateCurrentPrivacySettingReply privacySettingServiceUpdateCurrent(privacySettingServiceUpdateCurrentRequest)



更新当前登录账号的隐私设置

### Example

```java
// Import classes:
import com.bass.bbs.ApiClient;
import com.bass.bbs.ApiException;
import com.bass.bbs.Configuration;
import com.bass.bbs.models.*;
import com.bass.bbs.api.PrivacySettingServiceApi;
import com.bass.bbs.api.PrivacySettingServiceApi.*;

public class Example {
    public static void main(String[] args) {
        ApiClient defaultClient = Configuration.getDefaultApiClient();
        defaultClient.setBasePath("http://localhost");

        PrivacySettingServiceApi apiInstance = new PrivacySettingServiceApi(defaultClient);
        UpdateCurrentPrivacySettingRequest updateCurrentPrivacySettingRequest = new UpdateCurrentPrivacySettingRequest(); // UpdateCurrentPrivacySettingRequest | 
        try {
            APIprivacySettingServiceUpdateCurrentRequest request = APIprivacySettingServiceUpdateCurrentRequest.newBuilder()
                .updateCurrentPrivacySettingRequest(updateCurrentPrivacySettingRequest)
                .build();
            UpdateCurrentPrivacySettingReply result = apiInstance.privacySettingServiceUpdateCurrent(request);
            System.out.println(result);
        } catch (ApiException e) {
            System.err.println("Exception when calling PrivacySettingServiceApi#privacySettingServiceUpdateCurrent");
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
| privacySettingServiceUpdateCurrentRequest | [**APIprivacySettingServiceUpdateCurrentRequest**](PrivacySettingServiceApi.md#APIprivacySettingServiceUpdateCurrentRequest)|-|-|

### Return type

[**UpdateCurrentPrivacySettingReply**](UpdateCurrentPrivacySettingReply.md)


### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

### HTTP response details
| Status code | Description | Response headers |
|-------------|-------------|------------------|
| **200** | OK |  -  |

## privacySettingServiceUpdateCurrentWithHttpInfo

> ApiResponse<UpdateCurrentPrivacySettingReply> privacySettingServiceUpdateCurrentWithHttpInfo(privacySettingServiceUpdateCurrentRequest)



更新当前登录账号的隐私设置

### Example

```java
// Import classes:
import com.bass.bbs.ApiClient;
import com.bass.bbs.ApiException;
import com.bass.bbs.ApiResponse;
import com.bass.bbs.Configuration;
import com.bass.bbs.models.*;
import com.bass.bbs.api.PrivacySettingServiceApi;
import com.bass.bbs.api.PrivacySettingServiceApi.*;

public class Example {
    public static void main(String[] args) {
        ApiClient defaultClient = Configuration.getDefaultApiClient();
        defaultClient.setBasePath("http://localhost");

        PrivacySettingServiceApi apiInstance = new PrivacySettingServiceApi(defaultClient);
        UpdateCurrentPrivacySettingRequest updateCurrentPrivacySettingRequest = new UpdateCurrentPrivacySettingRequest(); // UpdateCurrentPrivacySettingRequest | 
        try {
            APIprivacySettingServiceUpdateCurrentRequest request = APIprivacySettingServiceUpdateCurrentRequest.newBuilder()
                .updateCurrentPrivacySettingRequest(updateCurrentPrivacySettingRequest)
                .build();
            ApiResponse<UpdateCurrentPrivacySettingReply> response = apiInstance.privacySettingServiceUpdateCurrentWithHttpInfo(request);
            System.out.println("Status code: " + response.getStatusCode());
            System.out.println("Response headers: " + response.getHeaders());
            System.out.println("Response body: " + response.getData());
        } catch (ApiException e) {
            System.err.println("Exception when calling PrivacySettingServiceApi#privacySettingServiceUpdateCurrent");
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
| privacySettingServiceUpdateCurrentRequest | [**APIprivacySettingServiceUpdateCurrentRequest**](PrivacySettingServiceApi.md#APIprivacySettingServiceUpdateCurrentRequest)|-|-|

### Return type

ApiResponse<[**UpdateCurrentPrivacySettingReply**](UpdateCurrentPrivacySettingReply.md)>


### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

### HTTP response details
| Status code | Description | Response headers |
|-------------|-------------|------------------|
| **200** | OK |  -  |


<a id="APIprivacySettingServiceUpdateCurrentRequest"></a>
## APIprivacySettingServiceUpdateCurrentRequest
### Properties

|     Name      |    Type       | Description   |     Notes    |
| ------------- | ------------- | ------------- | -------------|
| **updateCurrentPrivacySettingRequest** | [**UpdateCurrentPrivacySettingRequest**](UpdateCurrentPrivacySettingRequest.md) |  | |


