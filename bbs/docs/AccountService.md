# AccountService

All URIs are relative to *http://localhost*

| Method | HTTP request | Description |
|------------- | ------------- | -------------|
| [**avatar**](AccountService.md#avatar) | **GET** /v1/user/account/avatar |  |
| [**avatarWithHttpInfo**](AccountService.md#avatarWithHttpInfo) | **GET** /v1/user/account/avatar |  |
| [**getCurrent**](AccountService.md#getCurrent) | **POST** /v1/user/account/get-current |  |
| [**getCurrentWithHttpInfo**](AccountService.md#getCurrentWithHttpInfo) | **POST** /v1/user/account/get-current |  |
| [**getProfile**](AccountService.md#getProfile) | **POST** /v1/user/account/get-profile |  |
| [**getProfileWithHttpInfo**](AccountService.md#getProfileWithHttpInfo) | **POST** /v1/user/account/get-profile |  |
| [**updateProfile**](AccountService.md#updateProfile) | **POST** /v1/user/account/update-profile |  |
| [**updateProfileWithHttpInfo**](AccountService.md#updateProfileWithHttpInfo) | **POST** /v1/user/account/update-profile |  |



## avatar

> ImageReply avatar(avatarRequest)



生成默认账号头像。

### Example

```java
// Import classes:
import com.bass.bbs.ApiClient;
import com.bass.bbs.ApiException;
import com.bass.bbs.Configuration;
import com.bass.bbs.models.*;
import com.bass.bbs.api.AccountService;
import com.bass.bbs.api.AccountService.*;

public class Example {
    public static void main(String[] args) {
        ApiClient defaultClient = Configuration.getDefaultApiClient();
        defaultClient.setBasePath("http://localhost");

        AccountService apiInstance = new AccountService(defaultClient);
        String name = "name_example"; // String | 用于生成头像的账号名。
        try {
            APIavatarRequest request = APIavatarRequest.newBuilder()
                .name(name)
                .build();
            ImageReply result = apiInstance.avatar(request);
            System.out.println(result);
        } catch (ApiException e) {
            System.err.println("Exception when calling AccountService#avatar");
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
| avatarRequest | [**APIavatarRequest**](AccountService.md#APIavatarRequest)|-|-|

### Return type

[**ImageReply**](ImageReply.md)


### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

### HTTP response details
| Status code | Description | Response headers |
|-------------|-------------|------------------|
| **200** | OK |  -  |

## avatarWithHttpInfo

> ApiResponse<ImageReply> avatarWithHttpInfo(avatarRequest)



生成默认账号头像。

### Example

```java
// Import classes:
import com.bass.bbs.ApiClient;
import com.bass.bbs.ApiException;
import com.bass.bbs.ApiResponse;
import com.bass.bbs.Configuration;
import com.bass.bbs.models.*;
import com.bass.bbs.api.AccountService;
import com.bass.bbs.api.AccountService.*;

public class Example {
    public static void main(String[] args) {
        ApiClient defaultClient = Configuration.getDefaultApiClient();
        defaultClient.setBasePath("http://localhost");

        AccountService apiInstance = new AccountService(defaultClient);
        String name = "name_example"; // String | 用于生成头像的账号名。
        try {
            APIavatarRequest request = APIavatarRequest.newBuilder()
                .name(name)
                .build();
            ApiResponse<ImageReply> response = apiInstance.avatarWithHttpInfo(request);
            System.out.println("Status code: " + response.getStatusCode());
            System.out.println("Response headers: " + response.getHeaders());
            System.out.println("Response body: " + response.getData());
        } catch (ApiException e) {
            System.err.println("Exception when calling AccountService#avatar");
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
| avatarRequest | [**APIavatarRequest**](AccountService.md#APIavatarRequest)|-|-|

### Return type

ApiResponse<[**ImageReply**](ImageReply.md)>


### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

### HTTP response details
| Status code | Description | Response headers |
|-------------|-------------|------------------|
| **200** | OK |  -  |


<a id="APIavatarRequest"></a>
## APIavatarRequest
### Properties

|     Name      |    Type       | Description   |     Notes    |
| ------------- | ------------- | ------------- | -------------|
| **name** | **String** | 用于生成头像的账号名。 | [optional] |



## getCurrent

> GetCurrentAccountReply getCurrent(getCurrentRequest)



获取当前账号的完整资料。

### Example

```java
// Import classes:
import com.bass.bbs.ApiClient;
import com.bass.bbs.ApiException;
import com.bass.bbs.Configuration;
import com.bass.bbs.models.*;
import com.bass.bbs.api.AccountService;
import com.bass.bbs.api.AccountService.*;

public class Example {
    public static void main(String[] args) {
        ApiClient defaultClient = Configuration.getDefaultApiClient();
        defaultClient.setBasePath("http://localhost");

        AccountService apiInstance = new AccountService(defaultClient);
        Object body = null; // Object | 
        try {
            APIgetCurrentRequest request = APIgetCurrentRequest.newBuilder()
                .body(body)
                .build();
            GetCurrentAccountReply result = apiInstance.getCurrent(request);
            System.out.println(result);
        } catch (ApiException e) {
            System.err.println("Exception when calling AccountService#getCurrent");
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
| getCurrentRequest | [**APIgetCurrentRequest**](AccountService.md#APIgetCurrentRequest)|-|-|

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

## getCurrentWithHttpInfo

> ApiResponse<GetCurrentAccountReply> getCurrentWithHttpInfo(getCurrentRequest)



获取当前账号的完整资料。

### Example

```java
// Import classes:
import com.bass.bbs.ApiClient;
import com.bass.bbs.ApiException;
import com.bass.bbs.ApiResponse;
import com.bass.bbs.Configuration;
import com.bass.bbs.models.*;
import com.bass.bbs.api.AccountService;
import com.bass.bbs.api.AccountService.*;

public class Example {
    public static void main(String[] args) {
        ApiClient defaultClient = Configuration.getDefaultApiClient();
        defaultClient.setBasePath("http://localhost");

        AccountService apiInstance = new AccountService(defaultClient);
        Object body = null; // Object | 
        try {
            APIgetCurrentRequest request = APIgetCurrentRequest.newBuilder()
                .body(body)
                .build();
            ApiResponse<GetCurrentAccountReply> response = apiInstance.getCurrentWithHttpInfo(request);
            System.out.println("Status code: " + response.getStatusCode());
            System.out.println("Response headers: " + response.getHeaders());
            System.out.println("Response body: " + response.getData());
        } catch (ApiException e) {
            System.err.println("Exception when calling AccountService#getCurrent");
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
| getCurrentRequest | [**APIgetCurrentRequest**](AccountService.md#APIgetCurrentRequest)|-|-|

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


<a id="APIgetCurrentRequest"></a>
## APIgetCurrentRequest
### Properties

|     Name      |    Type       | Description   |     Notes    |
| ------------- | ------------- | ------------- | -------------|
| **body** | **Object** |  | |



## getProfile

> GetProfileAccountReply getProfile(getProfileRequest)



按账号 ID 获取账号展示资料。

### Example

```java
// Import classes:
import com.bass.bbs.ApiClient;
import com.bass.bbs.ApiException;
import com.bass.bbs.Configuration;
import com.bass.bbs.models.*;
import com.bass.bbs.api.AccountService;
import com.bass.bbs.api.AccountService.*;

public class Example {
    public static void main(String[] args) {
        ApiClient defaultClient = Configuration.getDefaultApiClient();
        defaultClient.setBasePath("http://localhost");

        AccountService apiInstance = new AccountService(defaultClient);
        GetProfileAccountRequest getProfileAccountRequest = new GetProfileAccountRequest(); // GetProfileAccountRequest | 
        try {
            APIgetProfileRequest request = APIgetProfileRequest.newBuilder()
                .getProfileAccountRequest(getProfileAccountRequest)
                .build();
            GetProfileAccountReply result = apiInstance.getProfile(request);
            System.out.println(result);
        } catch (ApiException e) {
            System.err.println("Exception when calling AccountService#getProfile");
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
| getProfileRequest | [**APIgetProfileRequest**](AccountService.md#APIgetProfileRequest)|-|-|

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

## getProfileWithHttpInfo

> ApiResponse<GetProfileAccountReply> getProfileWithHttpInfo(getProfileRequest)



按账号 ID 获取账号展示资料。

### Example

```java
// Import classes:
import com.bass.bbs.ApiClient;
import com.bass.bbs.ApiException;
import com.bass.bbs.ApiResponse;
import com.bass.bbs.Configuration;
import com.bass.bbs.models.*;
import com.bass.bbs.api.AccountService;
import com.bass.bbs.api.AccountService.*;

public class Example {
    public static void main(String[] args) {
        ApiClient defaultClient = Configuration.getDefaultApiClient();
        defaultClient.setBasePath("http://localhost");

        AccountService apiInstance = new AccountService(defaultClient);
        GetProfileAccountRequest getProfileAccountRequest = new GetProfileAccountRequest(); // GetProfileAccountRequest | 
        try {
            APIgetProfileRequest request = APIgetProfileRequest.newBuilder()
                .getProfileAccountRequest(getProfileAccountRequest)
                .build();
            ApiResponse<GetProfileAccountReply> response = apiInstance.getProfileWithHttpInfo(request);
            System.out.println("Status code: " + response.getStatusCode());
            System.out.println("Response headers: " + response.getHeaders());
            System.out.println("Response body: " + response.getData());
        } catch (ApiException e) {
            System.err.println("Exception when calling AccountService#getProfile");
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
| getProfileRequest | [**APIgetProfileRequest**](AccountService.md#APIgetProfileRequest)|-|-|

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


<a id="APIgetProfileRequest"></a>
## APIgetProfileRequest
### Properties

|     Name      |    Type       | Description   |     Notes    |
| ------------- | ------------- | ------------- | -------------|
| **getProfileAccountRequest** | [**GetProfileAccountRequest**](GetProfileAccountRequest.md) |  | |



## updateProfile

> UpdateProfileAccountReply updateProfile(updateProfileRequest)



更新当前账号的展示资料。

### Example

```java
// Import classes:
import com.bass.bbs.ApiClient;
import com.bass.bbs.ApiException;
import com.bass.bbs.Configuration;
import com.bass.bbs.models.*;
import com.bass.bbs.api.AccountService;
import com.bass.bbs.api.AccountService.*;

public class Example {
    public static void main(String[] args) {
        ApiClient defaultClient = Configuration.getDefaultApiClient();
        defaultClient.setBasePath("http://localhost");

        AccountService apiInstance = new AccountService(defaultClient);
        UpdateProfileAccountRequest updateProfileAccountRequest = new UpdateProfileAccountRequest(); // UpdateProfileAccountRequest | 
        try {
            APIupdateProfileRequest request = APIupdateProfileRequest.newBuilder()
                .updateProfileAccountRequest(updateProfileAccountRequest)
                .build();
            UpdateProfileAccountReply result = apiInstance.updateProfile(request);
            System.out.println(result);
        } catch (ApiException e) {
            System.err.println("Exception when calling AccountService#updateProfile");
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
| updateProfileRequest | [**APIupdateProfileRequest**](AccountService.md#APIupdateProfileRequest)|-|-|

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

## updateProfileWithHttpInfo

> ApiResponse<UpdateProfileAccountReply> updateProfileWithHttpInfo(updateProfileRequest)



更新当前账号的展示资料。

### Example

```java
// Import classes:
import com.bass.bbs.ApiClient;
import com.bass.bbs.ApiException;
import com.bass.bbs.ApiResponse;
import com.bass.bbs.Configuration;
import com.bass.bbs.models.*;
import com.bass.bbs.api.AccountService;
import com.bass.bbs.api.AccountService.*;

public class Example {
    public static void main(String[] args) {
        ApiClient defaultClient = Configuration.getDefaultApiClient();
        defaultClient.setBasePath("http://localhost");

        AccountService apiInstance = new AccountService(defaultClient);
        UpdateProfileAccountRequest updateProfileAccountRequest = new UpdateProfileAccountRequest(); // UpdateProfileAccountRequest | 
        try {
            APIupdateProfileRequest request = APIupdateProfileRequest.newBuilder()
                .updateProfileAccountRequest(updateProfileAccountRequest)
                .build();
            ApiResponse<UpdateProfileAccountReply> response = apiInstance.updateProfileWithHttpInfo(request);
            System.out.println("Status code: " + response.getStatusCode());
            System.out.println("Response headers: " + response.getHeaders());
            System.out.println("Response body: " + response.getData());
        } catch (ApiException e) {
            System.err.println("Exception when calling AccountService#updateProfile");
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
| updateProfileRequest | [**APIupdateProfileRequest**](AccountService.md#APIupdateProfileRequest)|-|-|

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


<a id="APIupdateProfileRequest"></a>
## APIupdateProfileRequest
### Properties

|     Name      |    Type       | Description   |     Notes    |
| ------------- | ------------- | ------------- | -------------|
| **updateProfileAccountRequest** | [**UpdateProfileAccountRequest**](UpdateProfileAccountRequest.md) |  | |


