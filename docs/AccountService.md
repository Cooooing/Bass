# AccountService

All URIs are relative to *http://localhost*

| Method | HTTP request | Description |
|------------- | ------------- | -------------|
| [**avatar**](AccountService.md#avatar) | **GET** /v1/user/account/avatar |  |
| [**avatarWithHttpInfo**](AccountService.md#avatarWithHttpInfo) | **GET** /v1/user/account/avatar |  |
| [**completeProfileImageUpload**](AccountService.md#completeProfileImageUpload) | **POST** /v1/user/account/complete-profile-image-upload |  |
| [**completeProfileImageUploadWithHttpInfo**](AccountService.md#completeProfileImageUploadWithHttpInfo) | **POST** /v1/user/account/complete-profile-image-upload |  |
| [**getCurrent**](AccountService.md#getCurrent) | **POST** /v1/user/account/get-current |  |
| [**getCurrentWithHttpInfo**](AccountService.md#getCurrentWithHttpInfo) | **POST** /v1/user/account/get-current |  |
| [**getProfile**](AccountService.md#getProfile) | **POST** /v1/user/account/get-profile |  |
| [**getProfileWithHttpInfo**](AccountService.md#getProfileWithHttpInfo) | **POST** /v1/user/account/get-profile |  |
| [**listFollowers**](AccountService.md#listFollowers) | **POST** /v1/user/account/list-followers |  |
| [**listFollowersWithHttpInfo**](AccountService.md#listFollowersWithHttpInfo) | **POST** /v1/user/account/list-followers |  |
| [**listFollowing**](AccountService.md#listFollowing) | **POST** /v1/user/account/list-following |  |
| [**listFollowingWithHttpInfo**](AccountService.md#listFollowingWithHttpInfo) | **POST** /v1/user/account/list-following |  |
| [**prepareProfileImageUpload**](AccountService.md#prepareProfileImageUpload) | **POST** /v1/user/account/prepare-profile-image-upload |  |
| [**prepareProfileImageUploadWithHttpInfo**](AccountService.md#prepareProfileImageUploadWithHttpInfo) | **POST** /v1/user/account/prepare-profile-image-upload |  |
| [**updateEmail**](AccountService.md#updateEmail) | **POST** /v1/user/account/update-email |  |
| [**updateEmailWithHttpInfo**](AccountService.md#updateEmailWithHttpInfo) | **POST** /v1/user/account/update-email |  |
| [**updatePassword**](AccountService.md#updatePassword) | **POST** /v1/user/account/update-password |  |
| [**updatePasswordWithHttpInfo**](AccountService.md#updatePasswordWithHttpInfo) | **POST** /v1/user/account/update-password |  |
| [**updatePhone**](AccountService.md#updatePhone) | **POST** /v1/user/account/update-phone |  |
| [**updatePhoneWithHttpInfo**](AccountService.md#updatePhoneWithHttpInfo) | **POST** /v1/user/account/update-phone |  |
| [**updateProfile**](AccountService.md#updateProfile) | **POST** /v1/user/account/update-profile |  |
| [**updateProfileWithHttpInfo**](AccountService.md#updateProfileWithHttpInfo) | **POST** /v1/user/account/update-profile |  |



## avatar

> ImageResp avatar(avatarRequest)



生成默认账号头像

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
        String name = "name_example"; // String | 
        try {
            APIavatarRequest request = APIavatarRequest.newBuilder()
                .name(name)
                .build();
            ImageResp result = apiInstance.avatar(request);
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

[**ImageResp**](ImageResp.md)


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

> ApiResponse<ImageResp> avatarWithHttpInfo(avatarRequest)



生成默认账号头像

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
        String name = "name_example"; // String | 
        try {
            APIavatarRequest request = APIavatarRequest.newBuilder()
                .name(name)
                .build();
            ApiResponse<ImageResp> response = apiInstance.avatarWithHttpInfo(request);
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

ApiResponse<[**ImageResp**](ImageResp.md)>


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
| **name** | **String** |  | [optional] |



## completeProfileImageUpload

> CompleteProfileImageUploadAccountResp completeProfileImageUpload(completeProfileImageUploadRequest)



将已由 MinIO 回调确认的资源绑定到当前账号资料。

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
        CompleteProfileImageUploadAccountReq completeProfileImageUploadAccountReq = new CompleteProfileImageUploadAccountReq(); // CompleteProfileImageUploadAccountReq | 
        try {
            APIcompleteProfileImageUploadRequest request = APIcompleteProfileImageUploadRequest.newBuilder()
                .completeProfileImageUploadAccountReq(completeProfileImageUploadAccountReq)
                .build();
            CompleteProfileImageUploadAccountResp result = apiInstance.completeProfileImageUpload(request);
            System.out.println(result);
        } catch (ApiException e) {
            System.err.println("Exception when calling AccountService#completeProfileImageUpload");
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
| completeProfileImageUploadRequest | [**APIcompleteProfileImageUploadRequest**](AccountService.md#APIcompleteProfileImageUploadRequest)|-|-|

### Return type

[**CompleteProfileImageUploadAccountResp**](CompleteProfileImageUploadAccountResp.md)


### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

### HTTP response details
| Status code | Description | Response headers |
|-------------|-------------|------------------|
| **200** | OK |  -  |

## completeProfileImageUploadWithHttpInfo

> ApiResponse<CompleteProfileImageUploadAccountResp> completeProfileImageUploadWithHttpInfo(completeProfileImageUploadRequest)



将已由 MinIO 回调确认的资源绑定到当前账号资料。

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
        CompleteProfileImageUploadAccountReq completeProfileImageUploadAccountReq = new CompleteProfileImageUploadAccountReq(); // CompleteProfileImageUploadAccountReq | 
        try {
            APIcompleteProfileImageUploadRequest request = APIcompleteProfileImageUploadRequest.newBuilder()
                .completeProfileImageUploadAccountReq(completeProfileImageUploadAccountReq)
                .build();
            ApiResponse<CompleteProfileImageUploadAccountResp> response = apiInstance.completeProfileImageUploadWithHttpInfo(request);
            System.out.println("Status code: " + response.getStatusCode());
            System.out.println("Response headers: " + response.getHeaders());
            System.out.println("Response body: " + response.getData());
        } catch (ApiException e) {
            System.err.println("Exception when calling AccountService#completeProfileImageUpload");
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
| completeProfileImageUploadRequest | [**APIcompleteProfileImageUploadRequest**](AccountService.md#APIcompleteProfileImageUploadRequest)|-|-|

### Return type

ApiResponse<[**CompleteProfileImageUploadAccountResp**](CompleteProfileImageUploadAccountResp.md)>


### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

### HTTP response details
| Status code | Description | Response headers |
|-------------|-------------|------------------|
| **200** | OK |  -  |


<a id="APIcompleteProfileImageUploadRequest"></a>
## APIcompleteProfileImageUploadRequest
### Properties

|     Name      |    Type       | Description   |     Notes    |
| ------------- | ------------- | ------------- | -------------|
| **completeProfileImageUploadAccountReq** | [**CompleteProfileImageUploadAccountReq**](CompleteProfileImageUploadAccountReq.md) |  | |



## getCurrent

> GetCurrentAccountResp getCurrent(getCurrentRequest)



获取当前账号完整资料

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
            GetCurrentAccountResp result = apiInstance.getCurrent(request);
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

[**GetCurrentAccountResp**](GetCurrentAccountResp.md)


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

> ApiResponse<GetCurrentAccountResp> getCurrentWithHttpInfo(getCurrentRequest)



获取当前账号完整资料

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
            ApiResponse<GetCurrentAccountResp> response = apiInstance.getCurrentWithHttpInfo(request);
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

ApiResponse<[**GetCurrentAccountResp**](GetCurrentAccountResp.md)>


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

> GetProfileResp getProfile(getProfileRequest)



按账号名获取个人主页资料。

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
        GetProfileReq getProfileReq = new GetProfileReq(); // GetProfileReq | 
        try {
            APIgetProfileRequest request = APIgetProfileRequest.newBuilder()
                .getProfileReq(getProfileReq)
                .build();
            GetProfileResp result = apiInstance.getProfile(request);
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

[**GetProfileResp**](GetProfileResp.md)


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

> ApiResponse<GetProfileResp> getProfileWithHttpInfo(getProfileRequest)



按账号名获取个人主页资料。

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
        GetProfileReq getProfileReq = new GetProfileReq(); // GetProfileReq | 
        try {
            APIgetProfileRequest request = APIgetProfileRequest.newBuilder()
                .getProfileReq(getProfileReq)
                .build();
            ApiResponse<GetProfileResp> response = apiInstance.getProfileWithHttpInfo(request);
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

ApiResponse<[**GetProfileResp**](GetProfileResp.md)>


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
| **getProfileReq** | [**GetProfileReq**](GetProfileReq.md) |  | |



## listFollowers

> ListFollowersResp listFollowers(listFollowersRequest)



查询账号公开的粉丝列表。

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
        ListFollowersReq listFollowersReq = new ListFollowersReq(); // ListFollowersReq | 
        try {
            APIlistFollowersRequest request = APIlistFollowersRequest.newBuilder()
                .listFollowersReq(listFollowersReq)
                .build();
            ListFollowersResp result = apiInstance.listFollowers(request);
            System.out.println(result);
        } catch (ApiException e) {
            System.err.println("Exception when calling AccountService#listFollowers");
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
| listFollowersRequest | [**APIlistFollowersRequest**](AccountService.md#APIlistFollowersRequest)|-|-|

### Return type

[**ListFollowersResp**](ListFollowersResp.md)


### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

### HTTP response details
| Status code | Description | Response headers |
|-------------|-------------|------------------|
| **200** | OK |  -  |

## listFollowersWithHttpInfo

> ApiResponse<ListFollowersResp> listFollowersWithHttpInfo(listFollowersRequest)



查询账号公开的粉丝列表。

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
        ListFollowersReq listFollowersReq = new ListFollowersReq(); // ListFollowersReq | 
        try {
            APIlistFollowersRequest request = APIlistFollowersRequest.newBuilder()
                .listFollowersReq(listFollowersReq)
                .build();
            ApiResponse<ListFollowersResp> response = apiInstance.listFollowersWithHttpInfo(request);
            System.out.println("Status code: " + response.getStatusCode());
            System.out.println("Response headers: " + response.getHeaders());
            System.out.println("Response body: " + response.getData());
        } catch (ApiException e) {
            System.err.println("Exception when calling AccountService#listFollowers");
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
| listFollowersRequest | [**APIlistFollowersRequest**](AccountService.md#APIlistFollowersRequest)|-|-|

### Return type

ApiResponse<[**ListFollowersResp**](ListFollowersResp.md)>


### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

### HTTP response details
| Status code | Description | Response headers |
|-------------|-------------|------------------|
| **200** | OK |  -  |


<a id="APIlistFollowersRequest"></a>
## APIlistFollowersRequest
### Properties

|     Name      |    Type       | Description   |     Notes    |
| ------------- | ------------- | ------------- | -------------|
| **listFollowersReq** | [**ListFollowersReq**](ListFollowersReq.md) |  | |



## listFollowing

> ListFollowingResp listFollowing(listFollowingRequest)



查询账号公开的关注列表。

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
        ListFollowingReq listFollowingReq = new ListFollowingReq(); // ListFollowingReq | 
        try {
            APIlistFollowingRequest request = APIlistFollowingRequest.newBuilder()
                .listFollowingReq(listFollowingReq)
                .build();
            ListFollowingResp result = apiInstance.listFollowing(request);
            System.out.println(result);
        } catch (ApiException e) {
            System.err.println("Exception when calling AccountService#listFollowing");
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
| listFollowingRequest | [**APIlistFollowingRequest**](AccountService.md#APIlistFollowingRequest)|-|-|

### Return type

[**ListFollowingResp**](ListFollowingResp.md)


### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

### HTTP response details
| Status code | Description | Response headers |
|-------------|-------------|------------------|
| **200** | OK |  -  |

## listFollowingWithHttpInfo

> ApiResponse<ListFollowingResp> listFollowingWithHttpInfo(listFollowingRequest)



查询账号公开的关注列表。

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
        ListFollowingReq listFollowingReq = new ListFollowingReq(); // ListFollowingReq | 
        try {
            APIlistFollowingRequest request = APIlistFollowingRequest.newBuilder()
                .listFollowingReq(listFollowingReq)
                .build();
            ApiResponse<ListFollowingResp> response = apiInstance.listFollowingWithHttpInfo(request);
            System.out.println("Status code: " + response.getStatusCode());
            System.out.println("Response headers: " + response.getHeaders());
            System.out.println("Response body: " + response.getData());
        } catch (ApiException e) {
            System.err.println("Exception when calling AccountService#listFollowing");
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
| listFollowingRequest | [**APIlistFollowingRequest**](AccountService.md#APIlistFollowingRequest)|-|-|

### Return type

ApiResponse<[**ListFollowingResp**](ListFollowingResp.md)>


### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

### HTTP response details
| Status code | Description | Response headers |
|-------------|-------------|------------------|
| **200** | OK |  -  |


<a id="APIlistFollowingRequest"></a>
## APIlistFollowingRequest
### Properties

|     Name      |    Type       | Description   |     Notes    |
| ------------- | ------------- | ------------- | -------------|
| **listFollowingReq** | [**ListFollowingReq**](ListFollowingReq.md) |  | |



## prepareProfileImageUpload

> PrepareProfileImageUploadAccountResp prepareProfileImageUpload(prepareProfileImageUploadRequest)



申请当前账号资料图片的内容寻址直传能力。

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
        PrepareProfileImageUploadAccountReq prepareProfileImageUploadAccountReq = new PrepareProfileImageUploadAccountReq(); // PrepareProfileImageUploadAccountReq | 
        try {
            APIprepareProfileImageUploadRequest request = APIprepareProfileImageUploadRequest.newBuilder()
                .prepareProfileImageUploadAccountReq(prepareProfileImageUploadAccountReq)
                .build();
            PrepareProfileImageUploadAccountResp result = apiInstance.prepareProfileImageUpload(request);
            System.out.println(result);
        } catch (ApiException e) {
            System.err.println("Exception when calling AccountService#prepareProfileImageUpload");
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
| prepareProfileImageUploadRequest | [**APIprepareProfileImageUploadRequest**](AccountService.md#APIprepareProfileImageUploadRequest)|-|-|

### Return type

[**PrepareProfileImageUploadAccountResp**](PrepareProfileImageUploadAccountResp.md)


### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

### HTTP response details
| Status code | Description | Response headers |
|-------------|-------------|------------------|
| **200** | OK |  -  |

## prepareProfileImageUploadWithHttpInfo

> ApiResponse<PrepareProfileImageUploadAccountResp> prepareProfileImageUploadWithHttpInfo(prepareProfileImageUploadRequest)



申请当前账号资料图片的内容寻址直传能力。

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
        PrepareProfileImageUploadAccountReq prepareProfileImageUploadAccountReq = new PrepareProfileImageUploadAccountReq(); // PrepareProfileImageUploadAccountReq | 
        try {
            APIprepareProfileImageUploadRequest request = APIprepareProfileImageUploadRequest.newBuilder()
                .prepareProfileImageUploadAccountReq(prepareProfileImageUploadAccountReq)
                .build();
            ApiResponse<PrepareProfileImageUploadAccountResp> response = apiInstance.prepareProfileImageUploadWithHttpInfo(request);
            System.out.println("Status code: " + response.getStatusCode());
            System.out.println("Response headers: " + response.getHeaders());
            System.out.println("Response body: " + response.getData());
        } catch (ApiException e) {
            System.err.println("Exception when calling AccountService#prepareProfileImageUpload");
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
| prepareProfileImageUploadRequest | [**APIprepareProfileImageUploadRequest**](AccountService.md#APIprepareProfileImageUploadRequest)|-|-|

### Return type

ApiResponse<[**PrepareProfileImageUploadAccountResp**](PrepareProfileImageUploadAccountResp.md)>


### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

### HTTP response details
| Status code | Description | Response headers |
|-------------|-------------|------------------|
| **200** | OK |  -  |


<a id="APIprepareProfileImageUploadRequest"></a>
## APIprepareProfileImageUploadRequest
### Properties

|     Name      |    Type       | Description   |     Notes    |
| ------------- | ------------- | ------------- | -------------|
| **prepareProfileImageUploadAccountReq** | [**PrepareProfileImageUploadAccountReq**](PrepareProfileImageUploadAccountReq.md) |  | |



## updateEmail

> Object updateEmail(updateEmailRequest)



更新当前账号邮箱

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
        UpdateEmailAccountReq updateEmailAccountReq = new UpdateEmailAccountReq(); // UpdateEmailAccountReq | 
        try {
            APIupdateEmailRequest request = APIupdateEmailRequest.newBuilder()
                .updateEmailAccountReq(updateEmailAccountReq)
                .build();
            Object result = apiInstance.updateEmail(request);
            System.out.println(result);
        } catch (ApiException e) {
            System.err.println("Exception when calling AccountService#updateEmail");
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
| updateEmailRequest | [**APIupdateEmailRequest**](AccountService.md#APIupdateEmailRequest)|-|-|

### Return type

**Object**


### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

### HTTP response details
| Status code | Description | Response headers |
|-------------|-------------|------------------|
| **200** | OK |  -  |

## updateEmailWithHttpInfo

> ApiResponse<Object> updateEmailWithHttpInfo(updateEmailRequest)



更新当前账号邮箱

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
        UpdateEmailAccountReq updateEmailAccountReq = new UpdateEmailAccountReq(); // UpdateEmailAccountReq | 
        try {
            APIupdateEmailRequest request = APIupdateEmailRequest.newBuilder()
                .updateEmailAccountReq(updateEmailAccountReq)
                .build();
            ApiResponse<Object> response = apiInstance.updateEmailWithHttpInfo(request);
            System.out.println("Status code: " + response.getStatusCode());
            System.out.println("Response headers: " + response.getHeaders());
            System.out.println("Response body: " + response.getData());
        } catch (ApiException e) {
            System.err.println("Exception when calling AccountService#updateEmail");
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
| updateEmailRequest | [**APIupdateEmailRequest**](AccountService.md#APIupdateEmailRequest)|-|-|

### Return type

ApiResponse<**Object**>


### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

### HTTP response details
| Status code | Description | Response headers |
|-------------|-------------|------------------|
| **200** | OK |  -  |


<a id="APIupdateEmailRequest"></a>
## APIupdateEmailRequest
### Properties

|     Name      |    Type       | Description   |     Notes    |
| ------------- | ------------- | ------------- | -------------|
| **updateEmailAccountReq** | [**UpdateEmailAccountReq**](UpdateEmailAccountReq.md) |  | |



## updatePassword

> Object updatePassword(updatePasswordRequest)



更新当前账号密码

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
        UpdatePasswordAccountReq updatePasswordAccountReq = new UpdatePasswordAccountReq(); // UpdatePasswordAccountReq | 
        try {
            APIupdatePasswordRequest request = APIupdatePasswordRequest.newBuilder()
                .updatePasswordAccountReq(updatePasswordAccountReq)
                .build();
            Object result = apiInstance.updatePassword(request);
            System.out.println(result);
        } catch (ApiException e) {
            System.err.println("Exception when calling AccountService#updatePassword");
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
| updatePasswordRequest | [**APIupdatePasswordRequest**](AccountService.md#APIupdatePasswordRequest)|-|-|

### Return type

**Object**


### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

### HTTP response details
| Status code | Description | Response headers |
|-------------|-------------|------------------|
| **200** | OK |  -  |

## updatePasswordWithHttpInfo

> ApiResponse<Object> updatePasswordWithHttpInfo(updatePasswordRequest)



更新当前账号密码

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
        UpdatePasswordAccountReq updatePasswordAccountReq = new UpdatePasswordAccountReq(); // UpdatePasswordAccountReq | 
        try {
            APIupdatePasswordRequest request = APIupdatePasswordRequest.newBuilder()
                .updatePasswordAccountReq(updatePasswordAccountReq)
                .build();
            ApiResponse<Object> response = apiInstance.updatePasswordWithHttpInfo(request);
            System.out.println("Status code: " + response.getStatusCode());
            System.out.println("Response headers: " + response.getHeaders());
            System.out.println("Response body: " + response.getData());
        } catch (ApiException e) {
            System.err.println("Exception when calling AccountService#updatePassword");
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
| updatePasswordRequest | [**APIupdatePasswordRequest**](AccountService.md#APIupdatePasswordRequest)|-|-|

### Return type

ApiResponse<**Object**>


### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

### HTTP response details
| Status code | Description | Response headers |
|-------------|-------------|------------------|
| **200** | OK |  -  |


<a id="APIupdatePasswordRequest"></a>
## APIupdatePasswordRequest
### Properties

|     Name      |    Type       | Description   |     Notes    |
| ------------- | ------------- | ------------- | -------------|
| **updatePasswordAccountReq** | [**UpdatePasswordAccountReq**](UpdatePasswordAccountReq.md) |  | |



## updatePhone

> Object updatePhone(updatePhoneRequest)



更新当前账号手机号

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
        UpdatePhoneAccountReq updatePhoneAccountReq = new UpdatePhoneAccountReq(); // UpdatePhoneAccountReq | 
        try {
            APIupdatePhoneRequest request = APIupdatePhoneRequest.newBuilder()
                .updatePhoneAccountReq(updatePhoneAccountReq)
                .build();
            Object result = apiInstance.updatePhone(request);
            System.out.println(result);
        } catch (ApiException e) {
            System.err.println("Exception when calling AccountService#updatePhone");
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
| updatePhoneRequest | [**APIupdatePhoneRequest**](AccountService.md#APIupdatePhoneRequest)|-|-|

### Return type

**Object**


### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

### HTTP response details
| Status code | Description | Response headers |
|-------------|-------------|------------------|
| **200** | OK |  -  |

## updatePhoneWithHttpInfo

> ApiResponse<Object> updatePhoneWithHttpInfo(updatePhoneRequest)



更新当前账号手机号

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
        UpdatePhoneAccountReq updatePhoneAccountReq = new UpdatePhoneAccountReq(); // UpdatePhoneAccountReq | 
        try {
            APIupdatePhoneRequest request = APIupdatePhoneRequest.newBuilder()
                .updatePhoneAccountReq(updatePhoneAccountReq)
                .build();
            ApiResponse<Object> response = apiInstance.updatePhoneWithHttpInfo(request);
            System.out.println("Status code: " + response.getStatusCode());
            System.out.println("Response headers: " + response.getHeaders());
            System.out.println("Response body: " + response.getData());
        } catch (ApiException e) {
            System.err.println("Exception when calling AccountService#updatePhone");
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
| updatePhoneRequest | [**APIupdatePhoneRequest**](AccountService.md#APIupdatePhoneRequest)|-|-|

### Return type

ApiResponse<**Object**>


### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

### HTTP response details
| Status code | Description | Response headers |
|-------------|-------------|------------------|
| **200** | OK |  -  |


<a id="APIupdatePhoneRequest"></a>
## APIupdatePhoneRequest
### Properties

|     Name      |    Type       | Description   |     Notes    |
| ------------- | ------------- | ------------- | -------------|
| **updatePhoneAccountReq** | [**UpdatePhoneAccountReq**](UpdatePhoneAccountReq.md) |  | |



## updateProfile

> UpdateProfileAccountResp updateProfile(updateProfileRequest)



更新当前账号展示资料

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
        UpdateProfileAccountReq updateProfileAccountReq = new UpdateProfileAccountReq(); // UpdateProfileAccountReq | 
        try {
            APIupdateProfileRequest request = APIupdateProfileRequest.newBuilder()
                .updateProfileAccountReq(updateProfileAccountReq)
                .build();
            UpdateProfileAccountResp result = apiInstance.updateProfile(request);
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

[**UpdateProfileAccountResp**](UpdateProfileAccountResp.md)


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

> ApiResponse<UpdateProfileAccountResp> updateProfileWithHttpInfo(updateProfileRequest)



更新当前账号展示资料

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
        UpdateProfileAccountReq updateProfileAccountReq = new UpdateProfileAccountReq(); // UpdateProfileAccountReq | 
        try {
            APIupdateProfileRequest request = APIupdateProfileRequest.newBuilder()
                .updateProfileAccountReq(updateProfileAccountReq)
                .build();
            ApiResponse<UpdateProfileAccountResp> response = apiInstance.updateProfileWithHttpInfo(request);
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

ApiResponse<[**UpdateProfileAccountResp**](UpdateProfileAccountResp.md)>


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
| **updateProfileAccountReq** | [**UpdateProfileAccountReq**](UpdateProfileAccountReq.md) |  | |


