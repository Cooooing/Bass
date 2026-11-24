# AccountService

All URIs are relative to *http://localhost*

| Method | HTTP request | Description |
|------------- | ------------- | -------------|
| [**avatarAccount**](AccountService.md#avatarAccount) | **GET** /v1/user/account/avatar |  |
| [**avatarAccountWithHttpInfo**](AccountService.md#avatarAccountWithHttpInfo) | **GET** /v1/user/account/avatar |  |
| [**completeProfileImageUploadAccount**](AccountService.md#completeProfileImageUploadAccount) | **POST** /v1/user/account/complete-profile-image-upload |  |
| [**completeProfileImageUploadAccountWithHttpInfo**](AccountService.md#completeProfileImageUploadAccountWithHttpInfo) | **POST** /v1/user/account/complete-profile-image-upload |  |
| [**getAccountEconomy**](AccountService.md#getAccountEconomy) | **POST** /v1/user/account/economy/get |  |
| [**getAccountEconomyWithHttpInfo**](AccountService.md#getAccountEconomyWithHttpInfo) | **POST** /v1/user/account/economy/get |  |
| [**getCurrentAccount**](AccountService.md#getCurrentAccount) | **POST** /v1/user/account/get-current |  |
| [**getCurrentAccountWithHttpInfo**](AccountService.md#getCurrentAccountWithHttpInfo) | **POST** /v1/user/account/get-current |  |
| [**getProfile**](AccountService.md#getProfile) | **POST** /v1/user/account/get-profile |  |
| [**getProfileWithHttpInfo**](AccountService.md#getProfileWithHttpInfo) | **POST** /v1/user/account/get-profile |  |
| [**listFollowers**](AccountService.md#listFollowers) | **POST** /v1/user/account/list-followers |  |
| [**listFollowersWithHttpInfo**](AccountService.md#listFollowersWithHttpInfo) | **POST** /v1/user/account/list-followers |  |
| [**listFollowing**](AccountService.md#listFollowing) | **POST** /v1/user/account/list-following |  |
| [**listFollowingWithHttpInfo**](AccountService.md#listFollowingWithHttpInfo) | **POST** /v1/user/account/list-following |  |
| [**listPointRecords**](AccountService.md#listPointRecords) | **POST** /v1/user/account/list-point-records |  |
| [**listPointRecordsWithHttpInfo**](AccountService.md#listPointRecordsWithHttpInfo) | **POST** /v1/user/account/list-point-records |  |
| [**prepareProfileImageUploadAccount**](AccountService.md#prepareProfileImageUploadAccount) | **POST** /v1/user/account/prepare-profile-image-upload |  |
| [**prepareProfileImageUploadAccountWithHttpInfo**](AccountService.md#prepareProfileImageUploadAccountWithHttpInfo) | **POST** /v1/user/account/prepare-profile-image-upload |  |
| [**searchMentionAccounts**](AccountService.md#searchMentionAccounts) | **POST** /v1/user/account/search-mention |  |
| [**searchMentionAccountsWithHttpInfo**](AccountService.md#searchMentionAccountsWithHttpInfo) | **POST** /v1/user/account/search-mention |  |
| [**transferAccountPoints**](AccountService.md#transferAccountPoints) | **POST** /v1/user/account/economy/transfer |  |
| [**transferAccountPointsWithHttpInfo**](AccountService.md#transferAccountPointsWithHttpInfo) | **POST** /v1/user/account/economy/transfer |  |
| [**updateEmailAccount**](AccountService.md#updateEmailAccount) | **POST** /v1/user/account/update-email |  |
| [**updateEmailAccountWithHttpInfo**](AccountService.md#updateEmailAccountWithHttpInfo) | **POST** /v1/user/account/update-email |  |
| [**updatePasswordAccount**](AccountService.md#updatePasswordAccount) | **POST** /v1/user/account/update-password |  |
| [**updatePasswordAccountWithHttpInfo**](AccountService.md#updatePasswordAccountWithHttpInfo) | **POST** /v1/user/account/update-password |  |
| [**updatePhoneAccount**](AccountService.md#updatePhoneAccount) | **POST** /v1/user/account/update-phone |  |
| [**updatePhoneAccountWithHttpInfo**](AccountService.md#updatePhoneAccountWithHttpInfo) | **POST** /v1/user/account/update-phone |  |
| [**updateProfileAccount**](AccountService.md#updateProfileAccount) | **POST** /v1/user/account/update-profile |  |
| [**updateProfileAccountWithHttpInfo**](AccountService.md#updateProfileAccountWithHttpInfo) | **POST** /v1/user/account/update-profile |  |



## avatarAccount

> AvatarAccountResp avatarAccount(avatarAccountRequest)



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
        String name = "name_example"; // String | 账号名，必填
        try {
            APIavatarAccountRequest request = APIavatarAccountRequest.newBuilder()
                .name(name)
                .build();
            AvatarAccountResp result = apiInstance.avatarAccount(request);
            System.out.println(result);
        } catch (ApiException e) {
            System.err.println("Exception when calling AccountService#avatarAccount");
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
| avatarAccountRequest | [**APIavatarAccountRequest**](AccountService.md#APIavatarAccountRequest)|-|-|

### Return type

[**AvatarAccountResp**](AvatarAccountResp.md)


### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

### HTTP response details
| Status code | Description | Response headers |
|-------------|-------------|------------------|
| **200** | OK |  -  |

## avatarAccountWithHttpInfo

> ApiResponse<AvatarAccountResp> avatarAccountWithHttpInfo(avatarAccountRequest)



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
        String name = "name_example"; // String | 账号名，必填
        try {
            APIavatarAccountRequest request = APIavatarAccountRequest.newBuilder()
                .name(name)
                .build();
            ApiResponse<AvatarAccountResp> response = apiInstance.avatarAccountWithHttpInfo(request);
            System.out.println("Status code: " + response.getStatusCode());
            System.out.println("Response headers: " + response.getHeaders());
            System.out.println("Response body: " + response.getData());
        } catch (ApiException e) {
            System.err.println("Exception when calling AccountService#avatarAccount");
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
| avatarAccountRequest | [**APIavatarAccountRequest**](AccountService.md#APIavatarAccountRequest)|-|-|

### Return type

ApiResponse<[**AvatarAccountResp**](AvatarAccountResp.md)>


### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

### HTTP response details
| Status code | Description | Response headers |
|-------------|-------------|------------------|
| **200** | OK |  -  |


<a id="APIavatarAccountRequest"></a>
## APIavatarAccountRequest
### Properties

|     Name      |    Type       | Description   |     Notes    |
| ------------- | ------------- | ------------- | -------------|
| **name** | **String** | 账号名，必填 | [optional] |



## completeProfileImageUploadAccount

> CompleteProfileImageUploadAccountResp completeProfileImageUploadAccount(completeProfileImageUploadAccountRequest)



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
            APIcompleteProfileImageUploadAccountRequest request = APIcompleteProfileImageUploadAccountRequest.newBuilder()
                .completeProfileImageUploadAccountReq(completeProfileImageUploadAccountReq)
                .build();
            CompleteProfileImageUploadAccountResp result = apiInstance.completeProfileImageUploadAccount(request);
            System.out.println(result);
        } catch (ApiException e) {
            System.err.println("Exception when calling AccountService#completeProfileImageUploadAccount");
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
| completeProfileImageUploadAccountRequest | [**APIcompleteProfileImageUploadAccountRequest**](AccountService.md#APIcompleteProfileImageUploadAccountRequest)|-|-|

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

## completeProfileImageUploadAccountWithHttpInfo

> ApiResponse<CompleteProfileImageUploadAccountResp> completeProfileImageUploadAccountWithHttpInfo(completeProfileImageUploadAccountRequest)



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
            APIcompleteProfileImageUploadAccountRequest request = APIcompleteProfileImageUploadAccountRequest.newBuilder()
                .completeProfileImageUploadAccountReq(completeProfileImageUploadAccountReq)
                .build();
            ApiResponse<CompleteProfileImageUploadAccountResp> response = apiInstance.completeProfileImageUploadAccountWithHttpInfo(request);
            System.out.println("Status code: " + response.getStatusCode());
            System.out.println("Response headers: " + response.getHeaders());
            System.out.println("Response body: " + response.getData());
        } catch (ApiException e) {
            System.err.println("Exception when calling AccountService#completeProfileImageUploadAccount");
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
| completeProfileImageUploadAccountRequest | [**APIcompleteProfileImageUploadAccountRequest**](AccountService.md#APIcompleteProfileImageUploadAccountRequest)|-|-|

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


<a id="APIcompleteProfileImageUploadAccountRequest"></a>
## APIcompleteProfileImageUploadAccountRequest
### Properties

|     Name      |    Type       | Description   |     Notes    |
| ------------- | ------------- | ------------- | -------------|
| **completeProfileImageUploadAccountReq** | [**CompleteProfileImageUploadAccountReq**](CompleteProfileImageUploadAccountReq.md) |  | |



## getAccountEconomy

> GetAccountEconomyResp getAccountEconomy(getAccountEconomyRequest)



查询当前账号的积分余额和累计收支。

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
            APIgetAccountEconomyRequest request = APIgetAccountEconomyRequest.newBuilder()
                .body(body)
                .build();
            GetAccountEconomyResp result = apiInstance.getAccountEconomy(request);
            System.out.println(result);
        } catch (ApiException e) {
            System.err.println("Exception when calling AccountService#getAccountEconomy");
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
| getAccountEconomyRequest | [**APIgetAccountEconomyRequest**](AccountService.md#APIgetAccountEconomyRequest)|-|-|

### Return type

[**GetAccountEconomyResp**](GetAccountEconomyResp.md)


### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

### HTTP response details
| Status code | Description | Response headers |
|-------------|-------------|------------------|
| **200** | OK |  -  |

## getAccountEconomyWithHttpInfo

> ApiResponse<GetAccountEconomyResp> getAccountEconomyWithHttpInfo(getAccountEconomyRequest)



查询当前账号的积分余额和累计收支。

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
            APIgetAccountEconomyRequest request = APIgetAccountEconomyRequest.newBuilder()
                .body(body)
                .build();
            ApiResponse<GetAccountEconomyResp> response = apiInstance.getAccountEconomyWithHttpInfo(request);
            System.out.println("Status code: " + response.getStatusCode());
            System.out.println("Response headers: " + response.getHeaders());
            System.out.println("Response body: " + response.getData());
        } catch (ApiException e) {
            System.err.println("Exception when calling AccountService#getAccountEconomy");
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
| getAccountEconomyRequest | [**APIgetAccountEconomyRequest**](AccountService.md#APIgetAccountEconomyRequest)|-|-|

### Return type

ApiResponse<[**GetAccountEconomyResp**](GetAccountEconomyResp.md)>


### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

### HTTP response details
| Status code | Description | Response headers |
|-------------|-------------|------------------|
| **200** | OK |  -  |


<a id="APIgetAccountEconomyRequest"></a>
## APIgetAccountEconomyRequest
### Properties

|     Name      |    Type       | Description   |     Notes    |
| ------------- | ------------- | ------------- | -------------|
| **body** | **Object** |  | |



## getCurrentAccount

> GetCurrentAccountResp getCurrentAccount(getCurrentAccountRequest)



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
            APIgetCurrentAccountRequest request = APIgetCurrentAccountRequest.newBuilder()
                .body(body)
                .build();
            GetCurrentAccountResp result = apiInstance.getCurrentAccount(request);
            System.out.println(result);
        } catch (ApiException e) {
            System.err.println("Exception when calling AccountService#getCurrentAccount");
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
| getCurrentAccountRequest | [**APIgetCurrentAccountRequest**](AccountService.md#APIgetCurrentAccountRequest)|-|-|

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

## getCurrentAccountWithHttpInfo

> ApiResponse<GetCurrentAccountResp> getCurrentAccountWithHttpInfo(getCurrentAccountRequest)



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
            APIgetCurrentAccountRequest request = APIgetCurrentAccountRequest.newBuilder()
                .body(body)
                .build();
            ApiResponse<GetCurrentAccountResp> response = apiInstance.getCurrentAccountWithHttpInfo(request);
            System.out.println("Status code: " + response.getStatusCode());
            System.out.println("Response headers: " + response.getHeaders());
            System.out.println("Response body: " + response.getData());
        } catch (ApiException e) {
            System.err.println("Exception when calling AccountService#getCurrentAccount");
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
| getCurrentAccountRequest | [**APIgetCurrentAccountRequest**](AccountService.md#APIgetCurrentAccountRequest)|-|-|

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


<a id="APIgetCurrentAccountRequest"></a>
## APIgetCurrentAccountRequest
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



## listPointRecords

> ListPointRecordsResp listPointRecords(listPointRecordsRequest)



查询账号公开的积分流水。

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
        ListPointRecordsReq listPointRecordsReq = new ListPointRecordsReq(); // ListPointRecordsReq | 
        try {
            APIlistPointRecordsRequest request = APIlistPointRecordsRequest.newBuilder()
                .listPointRecordsReq(listPointRecordsReq)
                .build();
            ListPointRecordsResp result = apiInstance.listPointRecords(request);
            System.out.println(result);
        } catch (ApiException e) {
            System.err.println("Exception when calling AccountService#listPointRecords");
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
| listPointRecordsRequest | [**APIlistPointRecordsRequest**](AccountService.md#APIlistPointRecordsRequest)|-|-|

### Return type

[**ListPointRecordsResp**](ListPointRecordsResp.md)


### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

### HTTP response details
| Status code | Description | Response headers |
|-------------|-------------|------------------|
| **200** | OK |  -  |

## listPointRecordsWithHttpInfo

> ApiResponse<ListPointRecordsResp> listPointRecordsWithHttpInfo(listPointRecordsRequest)



查询账号公开的积分流水。

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
        ListPointRecordsReq listPointRecordsReq = new ListPointRecordsReq(); // ListPointRecordsReq | 
        try {
            APIlistPointRecordsRequest request = APIlistPointRecordsRequest.newBuilder()
                .listPointRecordsReq(listPointRecordsReq)
                .build();
            ApiResponse<ListPointRecordsResp> response = apiInstance.listPointRecordsWithHttpInfo(request);
            System.out.println("Status code: " + response.getStatusCode());
            System.out.println("Response headers: " + response.getHeaders());
            System.out.println("Response body: " + response.getData());
        } catch (ApiException e) {
            System.err.println("Exception when calling AccountService#listPointRecords");
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
| listPointRecordsRequest | [**APIlistPointRecordsRequest**](AccountService.md#APIlistPointRecordsRequest)|-|-|

### Return type

ApiResponse<[**ListPointRecordsResp**](ListPointRecordsResp.md)>


### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

### HTTP response details
| Status code | Description | Response headers |
|-------------|-------------|------------------|
| **200** | OK |  -  |


<a id="APIlistPointRecordsRequest"></a>
## APIlistPointRecordsRequest
### Properties

|     Name      |    Type       | Description   |     Notes    |
| ------------- | ------------- | ------------- | -------------|
| **listPointRecordsReq** | [**ListPointRecordsReq**](ListPointRecordsReq.md) |  | |



## prepareProfileImageUploadAccount

> PrepareProfileImageUploadAccountResp prepareProfileImageUploadAccount(prepareProfileImageUploadAccountRequest)



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
            APIprepareProfileImageUploadAccountRequest request = APIprepareProfileImageUploadAccountRequest.newBuilder()
                .prepareProfileImageUploadAccountReq(prepareProfileImageUploadAccountReq)
                .build();
            PrepareProfileImageUploadAccountResp result = apiInstance.prepareProfileImageUploadAccount(request);
            System.out.println(result);
        } catch (ApiException e) {
            System.err.println("Exception when calling AccountService#prepareProfileImageUploadAccount");
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
| prepareProfileImageUploadAccountRequest | [**APIprepareProfileImageUploadAccountRequest**](AccountService.md#APIprepareProfileImageUploadAccountRequest)|-|-|

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

## prepareProfileImageUploadAccountWithHttpInfo

> ApiResponse<PrepareProfileImageUploadAccountResp> prepareProfileImageUploadAccountWithHttpInfo(prepareProfileImageUploadAccountRequest)



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
            APIprepareProfileImageUploadAccountRequest request = APIprepareProfileImageUploadAccountRequest.newBuilder()
                .prepareProfileImageUploadAccountReq(prepareProfileImageUploadAccountReq)
                .build();
            ApiResponse<PrepareProfileImageUploadAccountResp> response = apiInstance.prepareProfileImageUploadAccountWithHttpInfo(request);
            System.out.println("Status code: " + response.getStatusCode());
            System.out.println("Response headers: " + response.getHeaders());
            System.out.println("Response body: " + response.getData());
        } catch (ApiException e) {
            System.err.println("Exception when calling AccountService#prepareProfileImageUploadAccount");
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
| prepareProfileImageUploadAccountRequest | [**APIprepareProfileImageUploadAccountRequest**](AccountService.md#APIprepareProfileImageUploadAccountRequest)|-|-|

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


<a id="APIprepareProfileImageUploadAccountRequest"></a>
## APIprepareProfileImageUploadAccountRequest
### Properties

|     Name      |    Type       | Description   |     Notes    |
| ------------- | ------------- | ------------- | -------------|
| **prepareProfileImageUploadAccountReq** | [**PrepareProfileImageUploadAccountReq**](PrepareProfileImageUploadAccountReq.md) |  | |



## searchMentionAccounts

> SearchMentionAccountsResp searchMentionAccounts(searchMentionAccountsRequest)



检索可提及账号的候选项。

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
        SearchMentionAccountsReq searchMentionAccountsReq = new SearchMentionAccountsReq(); // SearchMentionAccountsReq | 
        try {
            APIsearchMentionAccountsRequest request = APIsearchMentionAccountsRequest.newBuilder()
                .searchMentionAccountsReq(searchMentionAccountsReq)
                .build();
            SearchMentionAccountsResp result = apiInstance.searchMentionAccounts(request);
            System.out.println(result);
        } catch (ApiException e) {
            System.err.println("Exception when calling AccountService#searchMentionAccounts");
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
| searchMentionAccountsRequest | [**APIsearchMentionAccountsRequest**](AccountService.md#APIsearchMentionAccountsRequest)|-|-|

### Return type

[**SearchMentionAccountsResp**](SearchMentionAccountsResp.md)


### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

### HTTP response details
| Status code | Description | Response headers |
|-------------|-------------|------------------|
| **200** | OK |  -  |

## searchMentionAccountsWithHttpInfo

> ApiResponse<SearchMentionAccountsResp> searchMentionAccountsWithHttpInfo(searchMentionAccountsRequest)



检索可提及账号的候选项。

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
        SearchMentionAccountsReq searchMentionAccountsReq = new SearchMentionAccountsReq(); // SearchMentionAccountsReq | 
        try {
            APIsearchMentionAccountsRequest request = APIsearchMentionAccountsRequest.newBuilder()
                .searchMentionAccountsReq(searchMentionAccountsReq)
                .build();
            ApiResponse<SearchMentionAccountsResp> response = apiInstance.searchMentionAccountsWithHttpInfo(request);
            System.out.println("Status code: " + response.getStatusCode());
            System.out.println("Response headers: " + response.getHeaders());
            System.out.println("Response body: " + response.getData());
        } catch (ApiException e) {
            System.err.println("Exception when calling AccountService#searchMentionAccounts");
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
| searchMentionAccountsRequest | [**APIsearchMentionAccountsRequest**](AccountService.md#APIsearchMentionAccountsRequest)|-|-|

### Return type

ApiResponse<[**SearchMentionAccountsResp**](SearchMentionAccountsResp.md)>


### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

### HTTP response details
| Status code | Description | Response headers |
|-------------|-------------|------------------|
| **200** | OK |  -  |


<a id="APIsearchMentionAccountsRequest"></a>
## APIsearchMentionAccountsRequest
### Properties

|     Name      |    Type       | Description   |     Notes    |
| ------------- | ------------- | ------------- | -------------|
| **searchMentionAccountsReq** | [**SearchMentionAccountsReq**](SearchMentionAccountsReq.md) |  | |



## transferAccountPoints

> TransferAccountPointsResp transferAccountPoints(transferAccountPointsRequest)



向指定账号转移积分。

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
        TransferAccountPointsReq transferAccountPointsReq = new TransferAccountPointsReq(); // TransferAccountPointsReq | 
        try {
            APItransferAccountPointsRequest request = APItransferAccountPointsRequest.newBuilder()
                .transferAccountPointsReq(transferAccountPointsReq)
                .build();
            TransferAccountPointsResp result = apiInstance.transferAccountPoints(request);
            System.out.println(result);
        } catch (ApiException e) {
            System.err.println("Exception when calling AccountService#transferAccountPoints");
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
| transferAccountPointsRequest | [**APItransferAccountPointsRequest**](AccountService.md#APItransferAccountPointsRequest)|-|-|

### Return type

[**TransferAccountPointsResp**](TransferAccountPointsResp.md)


### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

### HTTP response details
| Status code | Description | Response headers |
|-------------|-------------|------------------|
| **200** | OK |  -  |

## transferAccountPointsWithHttpInfo

> ApiResponse<TransferAccountPointsResp> transferAccountPointsWithHttpInfo(transferAccountPointsRequest)



向指定账号转移积分。

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
        TransferAccountPointsReq transferAccountPointsReq = new TransferAccountPointsReq(); // TransferAccountPointsReq | 
        try {
            APItransferAccountPointsRequest request = APItransferAccountPointsRequest.newBuilder()
                .transferAccountPointsReq(transferAccountPointsReq)
                .build();
            ApiResponse<TransferAccountPointsResp> response = apiInstance.transferAccountPointsWithHttpInfo(request);
            System.out.println("Status code: " + response.getStatusCode());
            System.out.println("Response headers: " + response.getHeaders());
            System.out.println("Response body: " + response.getData());
        } catch (ApiException e) {
            System.err.println("Exception when calling AccountService#transferAccountPoints");
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
| transferAccountPointsRequest | [**APItransferAccountPointsRequest**](AccountService.md#APItransferAccountPointsRequest)|-|-|

### Return type

ApiResponse<[**TransferAccountPointsResp**](TransferAccountPointsResp.md)>


### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

### HTTP response details
| Status code | Description | Response headers |
|-------------|-------------|------------------|
| **200** | OK |  -  |


<a id="APItransferAccountPointsRequest"></a>
## APItransferAccountPointsRequest
### Properties

|     Name      |    Type       | Description   |     Notes    |
| ------------- | ------------- | ------------- | -------------|
| **transferAccountPointsReq** | [**TransferAccountPointsReq**](TransferAccountPointsReq.md) |  | |



## updateEmailAccount

> Object updateEmailAccount(updateEmailAccountRequest)



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
            APIupdateEmailAccountRequest request = APIupdateEmailAccountRequest.newBuilder()
                .updateEmailAccountReq(updateEmailAccountReq)
                .build();
            Object result = apiInstance.updateEmailAccount(request);
            System.out.println(result);
        } catch (ApiException e) {
            System.err.println("Exception when calling AccountService#updateEmailAccount");
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
| updateEmailAccountRequest | [**APIupdateEmailAccountRequest**](AccountService.md#APIupdateEmailAccountRequest)|-|-|

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

## updateEmailAccountWithHttpInfo

> ApiResponse<Object> updateEmailAccountWithHttpInfo(updateEmailAccountRequest)



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
            APIupdateEmailAccountRequest request = APIupdateEmailAccountRequest.newBuilder()
                .updateEmailAccountReq(updateEmailAccountReq)
                .build();
            ApiResponse<Object> response = apiInstance.updateEmailAccountWithHttpInfo(request);
            System.out.println("Status code: " + response.getStatusCode());
            System.out.println("Response headers: " + response.getHeaders());
            System.out.println("Response body: " + response.getData());
        } catch (ApiException e) {
            System.err.println("Exception when calling AccountService#updateEmailAccount");
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
| updateEmailAccountRequest | [**APIupdateEmailAccountRequest**](AccountService.md#APIupdateEmailAccountRequest)|-|-|

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


<a id="APIupdateEmailAccountRequest"></a>
## APIupdateEmailAccountRequest
### Properties

|     Name      |    Type       | Description   |     Notes    |
| ------------- | ------------- | ------------- | -------------|
| **updateEmailAccountReq** | [**UpdateEmailAccountReq**](UpdateEmailAccountReq.md) |  | |



## updatePasswordAccount

> Object updatePasswordAccount(updatePasswordAccountRequest)



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
            APIupdatePasswordAccountRequest request = APIupdatePasswordAccountRequest.newBuilder()
                .updatePasswordAccountReq(updatePasswordAccountReq)
                .build();
            Object result = apiInstance.updatePasswordAccount(request);
            System.out.println(result);
        } catch (ApiException e) {
            System.err.println("Exception when calling AccountService#updatePasswordAccount");
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
| updatePasswordAccountRequest | [**APIupdatePasswordAccountRequest**](AccountService.md#APIupdatePasswordAccountRequest)|-|-|

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

## updatePasswordAccountWithHttpInfo

> ApiResponse<Object> updatePasswordAccountWithHttpInfo(updatePasswordAccountRequest)



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
            APIupdatePasswordAccountRequest request = APIupdatePasswordAccountRequest.newBuilder()
                .updatePasswordAccountReq(updatePasswordAccountReq)
                .build();
            ApiResponse<Object> response = apiInstance.updatePasswordAccountWithHttpInfo(request);
            System.out.println("Status code: " + response.getStatusCode());
            System.out.println("Response headers: " + response.getHeaders());
            System.out.println("Response body: " + response.getData());
        } catch (ApiException e) {
            System.err.println("Exception when calling AccountService#updatePasswordAccount");
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
| updatePasswordAccountRequest | [**APIupdatePasswordAccountRequest**](AccountService.md#APIupdatePasswordAccountRequest)|-|-|

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


<a id="APIupdatePasswordAccountRequest"></a>
## APIupdatePasswordAccountRequest
### Properties

|     Name      |    Type       | Description   |     Notes    |
| ------------- | ------------- | ------------- | -------------|
| **updatePasswordAccountReq** | [**UpdatePasswordAccountReq**](UpdatePasswordAccountReq.md) |  | |



## updatePhoneAccount

> Object updatePhoneAccount(updatePhoneAccountRequest)



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
            APIupdatePhoneAccountRequest request = APIupdatePhoneAccountRequest.newBuilder()
                .updatePhoneAccountReq(updatePhoneAccountReq)
                .build();
            Object result = apiInstance.updatePhoneAccount(request);
            System.out.println(result);
        } catch (ApiException e) {
            System.err.println("Exception when calling AccountService#updatePhoneAccount");
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
| updatePhoneAccountRequest | [**APIupdatePhoneAccountRequest**](AccountService.md#APIupdatePhoneAccountRequest)|-|-|

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

## updatePhoneAccountWithHttpInfo

> ApiResponse<Object> updatePhoneAccountWithHttpInfo(updatePhoneAccountRequest)



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
            APIupdatePhoneAccountRequest request = APIupdatePhoneAccountRequest.newBuilder()
                .updatePhoneAccountReq(updatePhoneAccountReq)
                .build();
            ApiResponse<Object> response = apiInstance.updatePhoneAccountWithHttpInfo(request);
            System.out.println("Status code: " + response.getStatusCode());
            System.out.println("Response headers: " + response.getHeaders());
            System.out.println("Response body: " + response.getData());
        } catch (ApiException e) {
            System.err.println("Exception when calling AccountService#updatePhoneAccount");
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
| updatePhoneAccountRequest | [**APIupdatePhoneAccountRequest**](AccountService.md#APIupdatePhoneAccountRequest)|-|-|

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


<a id="APIupdatePhoneAccountRequest"></a>
## APIupdatePhoneAccountRequest
### Properties

|     Name      |    Type       | Description   |     Notes    |
| ------------- | ------------- | ------------- | -------------|
| **updatePhoneAccountReq** | [**UpdatePhoneAccountReq**](UpdatePhoneAccountReq.md) |  | |



## updateProfileAccount

> UpdateProfileAccountResp updateProfileAccount(updateProfileAccountRequest)



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
            APIupdateProfileAccountRequest request = APIupdateProfileAccountRequest.newBuilder()
                .updateProfileAccountReq(updateProfileAccountReq)
                .build();
            UpdateProfileAccountResp result = apiInstance.updateProfileAccount(request);
            System.out.println(result);
        } catch (ApiException e) {
            System.err.println("Exception when calling AccountService#updateProfileAccount");
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
| updateProfileAccountRequest | [**APIupdateProfileAccountRequest**](AccountService.md#APIupdateProfileAccountRequest)|-|-|

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

## updateProfileAccountWithHttpInfo

> ApiResponse<UpdateProfileAccountResp> updateProfileAccountWithHttpInfo(updateProfileAccountRequest)



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
            APIupdateProfileAccountRequest request = APIupdateProfileAccountRequest.newBuilder()
                .updateProfileAccountReq(updateProfileAccountReq)
                .build();
            ApiResponse<UpdateProfileAccountResp> response = apiInstance.updateProfileAccountWithHttpInfo(request);
            System.out.println("Status code: " + response.getStatusCode());
            System.out.println("Response headers: " + response.getHeaders());
            System.out.println("Response body: " + response.getData());
        } catch (ApiException e) {
            System.err.println("Exception when calling AccountService#updateProfileAccount");
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
| updateProfileAccountRequest | [**APIupdateProfileAccountRequest**](AccountService.md#APIupdateProfileAccountRequest)|-|-|

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


<a id="APIupdateProfileAccountRequest"></a>
## APIupdateProfileAccountRequest
### Properties

|     Name      |    Type       | Description   |     Notes    |
| ------------- | ------------- | ------------- | -------------|
| **updateProfileAccountReq** | [**UpdateProfileAccountReq**](UpdateProfileAccountReq.md) |  | |


