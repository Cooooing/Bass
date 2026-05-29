# AccountServiceApi

All URIs are relative to *http://localhost*

| Method | HTTP request | Description |
|------------- | ------------- | -------------|
| [**accountServiceGetCurrent**](AccountServiceApi.md#accountServiceGetCurrent) | **POST** /v1/user/account/get-current |  |
| [**accountServiceGetCurrentWithHttpInfo**](AccountServiceApi.md#accountServiceGetCurrentWithHttpInfo) | **POST** /v1/user/account/get-current |  |
| [**accountServiceGetProfile**](AccountServiceApi.md#accountServiceGetProfile) | **POST** /v1/user/account/get-profile |  |
| [**accountServiceGetProfileWithHttpInfo**](AccountServiceApi.md#accountServiceGetProfileWithHttpInfo) | **POST** /v1/user/account/get-profile |  |
| [**accountServiceUpdateProfile**](AccountServiceApi.md#accountServiceUpdateProfile) | **POST** /v1/user/account/update-profile |  |
| [**accountServiceUpdateProfileWithHttpInfo**](AccountServiceApi.md#accountServiceUpdateProfileWithHttpInfo) | **POST** /v1/user/account/update-profile |  |



## accountServiceGetCurrent

> GetCurrentAccountReply accountServiceGetCurrent(accountServiceGetCurrentRequest)



获取当前登录账号的完整资料

### Example

```java
// Import classes:
import com.bass.bbs.ApiClient;
import com.bass.bbs.ApiException;
import com.bass.bbs.Configuration;
import com.bass.bbs.models.*;
import com.bass.bbs.api.AccountServiceApi;
import com.bass.bbs.api.AccountServiceApi.*;

public class Example {
    public static void main(String[] args) {
        ApiClient defaultClient = Configuration.getDefaultApiClient();
        defaultClient.setBasePath("http://localhost");

        AccountServiceApi apiInstance = new AccountServiceApi(defaultClient);
        Object body = null; // Object | 
        try {
            APIaccountServiceGetCurrentRequest request = APIaccountServiceGetCurrentRequest.newBuilder()
                .body(body)
                .build();
            GetCurrentAccountReply result = apiInstance.accountServiceGetCurrent(request);
            System.out.println(result);
        } catch (ApiException e) {
            System.err.println("Exception when calling AccountServiceApi#accountServiceGetCurrent");
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
| accountServiceGetCurrentRequest | [**APIaccountServiceGetCurrentRequest**](AccountServiceApi.md#APIaccountServiceGetCurrentRequest)|-|-|

### Return type

[**GetCurrentAccountReply**](GetCurrentAccountReply.md)


### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

### HTTP response details
| Status code | Description | Response headers |
|-------------|-------------|------------------|
| **200** | OK |  -  |

## accountServiceGetCurrentWithHttpInfo

> ApiResponse<GetCurrentAccountReply> accountServiceGetCurrentWithHttpInfo(accountServiceGetCurrentRequest)



获取当前登录账号的完整资料

### Example

```java
// Import classes:
import com.bass.bbs.ApiClient;
import com.bass.bbs.ApiException;
import com.bass.bbs.ApiResponse;
import com.bass.bbs.Configuration;
import com.bass.bbs.models.*;
import com.bass.bbs.api.AccountServiceApi;
import com.bass.bbs.api.AccountServiceApi.*;

public class Example {
    public static void main(String[] args) {
        ApiClient defaultClient = Configuration.getDefaultApiClient();
        defaultClient.setBasePath("http://localhost");

        AccountServiceApi apiInstance = new AccountServiceApi(defaultClient);
        Object body = null; // Object | 
        try {
            APIaccountServiceGetCurrentRequest request = APIaccountServiceGetCurrentRequest.newBuilder()
                .body(body)
                .build();
            ApiResponse<GetCurrentAccountReply> response = apiInstance.accountServiceGetCurrentWithHttpInfo(request);
            System.out.println("Status code: " + response.getStatusCode());
            System.out.println("Response headers: " + response.getHeaders());
            System.out.println("Response body: " + response.getData());
        } catch (ApiException e) {
            System.err.println("Exception when calling AccountServiceApi#accountServiceGetCurrent");
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
| accountServiceGetCurrentRequest | [**APIaccountServiceGetCurrentRequest**](AccountServiceApi.md#APIaccountServiceGetCurrentRequest)|-|-|

### Return type

ApiResponse<[**GetCurrentAccountReply**](GetCurrentAccountReply.md)>


### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

### HTTP response details
| Status code | Description | Response headers |
|-------------|-------------|------------------|
| **200** | OK |  -  |


<a id="APIaccountServiceGetCurrentRequest"></a>
## APIaccountServiceGetCurrentRequest
### Properties

|     Name      |    Type       | Description   |     Notes    |
| ------------- | ------------- | ------------- | -------------|
| **body** | **Object** |  | |



## accountServiceGetProfile

> GetProfileAccountReply accountServiceGetProfile(accountServiceGetProfileRequest)



按账号 ID 获取账号展示资料

### Example

```java
// Import classes:
import com.bass.bbs.ApiClient;
import com.bass.bbs.ApiException;
import com.bass.bbs.Configuration;
import com.bass.bbs.models.*;
import com.bass.bbs.api.AccountServiceApi;
import com.bass.bbs.api.AccountServiceApi.*;

public class Example {
    public static void main(String[] args) {
        ApiClient defaultClient = Configuration.getDefaultApiClient();
        defaultClient.setBasePath("http://localhost");

        AccountServiceApi apiInstance = new AccountServiceApi(defaultClient);
        GetProfileAccountRequest getProfileAccountRequest = new GetProfileAccountRequest(); // GetProfileAccountRequest | 
        try {
            APIaccountServiceGetProfileRequest request = APIaccountServiceGetProfileRequest.newBuilder()
                .getProfileAccountRequest(getProfileAccountRequest)
                .build();
            GetProfileAccountReply result = apiInstance.accountServiceGetProfile(request);
            System.out.println(result);
        } catch (ApiException e) {
            System.err.println("Exception when calling AccountServiceApi#accountServiceGetProfile");
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
| accountServiceGetProfileRequest | [**APIaccountServiceGetProfileRequest**](AccountServiceApi.md#APIaccountServiceGetProfileRequest)|-|-|

### Return type

[**GetProfileAccountReply**](GetProfileAccountReply.md)


### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

### HTTP response details
| Status code | Description | Response headers |
|-------------|-------------|------------------|
| **200** | OK |  -  |

## accountServiceGetProfileWithHttpInfo

> ApiResponse<GetProfileAccountReply> accountServiceGetProfileWithHttpInfo(accountServiceGetProfileRequest)



按账号 ID 获取账号展示资料

### Example

```java
// Import classes:
import com.bass.bbs.ApiClient;
import com.bass.bbs.ApiException;
import com.bass.bbs.ApiResponse;
import com.bass.bbs.Configuration;
import com.bass.bbs.models.*;
import com.bass.bbs.api.AccountServiceApi;
import com.bass.bbs.api.AccountServiceApi.*;

public class Example {
    public static void main(String[] args) {
        ApiClient defaultClient = Configuration.getDefaultApiClient();
        defaultClient.setBasePath("http://localhost");

        AccountServiceApi apiInstance = new AccountServiceApi(defaultClient);
        GetProfileAccountRequest getProfileAccountRequest = new GetProfileAccountRequest(); // GetProfileAccountRequest | 
        try {
            APIaccountServiceGetProfileRequest request = APIaccountServiceGetProfileRequest.newBuilder()
                .getProfileAccountRequest(getProfileAccountRequest)
                .build();
            ApiResponse<GetProfileAccountReply> response = apiInstance.accountServiceGetProfileWithHttpInfo(request);
            System.out.println("Status code: " + response.getStatusCode());
            System.out.println("Response headers: " + response.getHeaders());
            System.out.println("Response body: " + response.getData());
        } catch (ApiException e) {
            System.err.println("Exception when calling AccountServiceApi#accountServiceGetProfile");
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
| accountServiceGetProfileRequest | [**APIaccountServiceGetProfileRequest**](AccountServiceApi.md#APIaccountServiceGetProfileRequest)|-|-|

### Return type

ApiResponse<[**GetProfileAccountReply**](GetProfileAccountReply.md)>


### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

### HTTP response details
| Status code | Description | Response headers |
|-------------|-------------|------------------|
| **200** | OK |  -  |


<a id="APIaccountServiceGetProfileRequest"></a>
## APIaccountServiceGetProfileRequest
### Properties

|     Name      |    Type       | Description   |     Notes    |
| ------------- | ------------- | ------------- | -------------|
| **getProfileAccountRequest** | [**GetProfileAccountRequest**](GetProfileAccountRequest.md) |  | |



## accountServiceUpdateProfile

> UpdateProfileAccountReply accountServiceUpdateProfile(accountServiceUpdateProfileRequest)



更新当前登录账号的展示资料

### Example

```java
// Import classes:
import com.bass.bbs.ApiClient;
import com.bass.bbs.ApiException;
import com.bass.bbs.Configuration;
import com.bass.bbs.models.*;
import com.bass.bbs.api.AccountServiceApi;
import com.bass.bbs.api.AccountServiceApi.*;

public class Example {
    public static void main(String[] args) {
        ApiClient defaultClient = Configuration.getDefaultApiClient();
        defaultClient.setBasePath("http://localhost");

        AccountServiceApi apiInstance = new AccountServiceApi(defaultClient);
        UpdateProfileAccountRequest updateProfileAccountRequest = new UpdateProfileAccountRequest(); // UpdateProfileAccountRequest | 
        try {
            APIaccountServiceUpdateProfileRequest request = APIaccountServiceUpdateProfileRequest.newBuilder()
                .updateProfileAccountRequest(updateProfileAccountRequest)
                .build();
            UpdateProfileAccountReply result = apiInstance.accountServiceUpdateProfile(request);
            System.out.println(result);
        } catch (ApiException e) {
            System.err.println("Exception when calling AccountServiceApi#accountServiceUpdateProfile");
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
| accountServiceUpdateProfileRequest | [**APIaccountServiceUpdateProfileRequest**](AccountServiceApi.md#APIaccountServiceUpdateProfileRequest)|-|-|

### Return type

[**UpdateProfileAccountReply**](UpdateProfileAccountReply.md)


### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

### HTTP response details
| Status code | Description | Response headers |
|-------------|-------------|------------------|
| **200** | OK |  -  |

## accountServiceUpdateProfileWithHttpInfo

> ApiResponse<UpdateProfileAccountReply> accountServiceUpdateProfileWithHttpInfo(accountServiceUpdateProfileRequest)



更新当前登录账号的展示资料

### Example

```java
// Import classes:
import com.bass.bbs.ApiClient;
import com.bass.bbs.ApiException;
import com.bass.bbs.ApiResponse;
import com.bass.bbs.Configuration;
import com.bass.bbs.models.*;
import com.bass.bbs.api.AccountServiceApi;
import com.bass.bbs.api.AccountServiceApi.*;

public class Example {
    public static void main(String[] args) {
        ApiClient defaultClient = Configuration.getDefaultApiClient();
        defaultClient.setBasePath("http://localhost");

        AccountServiceApi apiInstance = new AccountServiceApi(defaultClient);
        UpdateProfileAccountRequest updateProfileAccountRequest = new UpdateProfileAccountRequest(); // UpdateProfileAccountRequest | 
        try {
            APIaccountServiceUpdateProfileRequest request = APIaccountServiceUpdateProfileRequest.newBuilder()
                .updateProfileAccountRequest(updateProfileAccountRequest)
                .build();
            ApiResponse<UpdateProfileAccountReply> response = apiInstance.accountServiceUpdateProfileWithHttpInfo(request);
            System.out.println("Status code: " + response.getStatusCode());
            System.out.println("Response headers: " + response.getHeaders());
            System.out.println("Response body: " + response.getData());
        } catch (ApiException e) {
            System.err.println("Exception when calling AccountServiceApi#accountServiceUpdateProfile");
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
| accountServiceUpdateProfileRequest | [**APIaccountServiceUpdateProfileRequest**](AccountServiceApi.md#APIaccountServiceUpdateProfileRequest)|-|-|

### Return type

ApiResponse<[**UpdateProfileAccountReply**](UpdateProfileAccountReply.md)>


### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

### HTTP response details
| Status code | Description | Response headers |
|-------------|-------------|------------------|
| **200** | OK |  -  |


<a id="APIaccountServiceUpdateProfileRequest"></a>
## APIaccountServiceUpdateProfileRequest
### Properties

|     Name      |    Type       | Description   |     Notes    |
| ------------- | ------------- | ------------- | -------------|
| **updateProfileAccountRequest** | [**UpdateProfileAccountRequest**](UpdateProfileAccountRequest.md) |  | |


