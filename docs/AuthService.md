# AuthService

All URIs are relative to *http://localhost*

| Method | HTTP request | Description |
|------------- | ------------- | -------------|
| [**loginAccount**](AuthService.md#loginAccount) | **POST** /v1/game-idle/auth/login |  |
| [**loginAccountWithHttpInfo**](AuthService.md#loginAccountWithHttpInfo) | **POST** /v1/game-idle/auth/login |  |
| [**registerAccount**](AuthService.md#registerAccount) | **POST** /v1/game-idle/auth/register |  |
| [**registerAccountWithHttpInfo**](AuthService.md#registerAccountWithHttpInfo) | **POST** /v1/game-idle/auth/register |  |



## loginAccount

> LoginAccountResp loginAccount(loginAccountRequest)



登录账号。

### Example

```java
// Import classes:
import com.bass.bbs.ApiClient;
import com.bass.bbs.ApiException;
import com.bass.bbs.Configuration;
import com.bass.bbs.models.*;
import com.bass.bbs.api.AuthService;
import com.bass.bbs.api.AuthService.*;

public class Example {
    public static void main(String[] args) {
        ApiClient defaultClient = Configuration.getDefaultApiClient();
        defaultClient.setBasePath("http://localhost");

        AuthService apiInstance = new AuthService(defaultClient);
        LoginAccountReq loginAccountReq = new LoginAccountReq(); // LoginAccountReq | 
        try {
            APIloginAccountRequest request = APIloginAccountRequest.newBuilder()
                .loginAccountReq(loginAccountReq)
                .build();
            LoginAccountResp result = apiInstance.loginAccount(request);
            System.out.println(result);
        } catch (ApiException e) {
            System.err.println("Exception when calling AuthService#loginAccount");
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
| loginAccountRequest | [**APIloginAccountRequest**](AuthService.md#APIloginAccountRequest)|-|-|

### Return type

[**LoginAccountResp**](LoginAccountResp.md)


### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

### HTTP response details
| Status code | Description | Response headers |
|-------------|-------------|------------------|
| **200** | OK |  -  |

## loginAccountWithHttpInfo

> ApiResponse<LoginAccountResp> loginAccountWithHttpInfo(loginAccountRequest)



登录账号。

### Example

```java
// Import classes:
import com.bass.bbs.ApiClient;
import com.bass.bbs.ApiException;
import com.bass.bbs.ApiResponse;
import com.bass.bbs.Configuration;
import com.bass.bbs.models.*;
import com.bass.bbs.api.AuthService;
import com.bass.bbs.api.AuthService.*;

public class Example {
    public static void main(String[] args) {
        ApiClient defaultClient = Configuration.getDefaultApiClient();
        defaultClient.setBasePath("http://localhost");

        AuthService apiInstance = new AuthService(defaultClient);
        LoginAccountReq loginAccountReq = new LoginAccountReq(); // LoginAccountReq | 
        try {
            APIloginAccountRequest request = APIloginAccountRequest.newBuilder()
                .loginAccountReq(loginAccountReq)
                .build();
            ApiResponse<LoginAccountResp> response = apiInstance.loginAccountWithHttpInfo(request);
            System.out.println("Status code: " + response.getStatusCode());
            System.out.println("Response headers: " + response.getHeaders());
            System.out.println("Response body: " + response.getData());
        } catch (ApiException e) {
            System.err.println("Exception when calling AuthService#loginAccount");
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
| loginAccountRequest | [**APIloginAccountRequest**](AuthService.md#APIloginAccountRequest)|-|-|

### Return type

ApiResponse<[**LoginAccountResp**](LoginAccountResp.md)>


### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

### HTTP response details
| Status code | Description | Response headers |
|-------------|-------------|------------------|
| **200** | OK |  -  |


<a id="APIloginAccountRequest"></a>
## APIloginAccountRequest
### Properties

|     Name      |    Type       | Description   |     Notes    |
| ------------- | ------------- | ------------- | -------------|
| **loginAccountReq** | [**LoginAccountReq**](LoginAccountReq.md) |  | |



## registerAccount

> Object registerAccount(registerAccountRequest)



注册账号。

### Example

```java
// Import classes:
import com.bass.bbs.ApiClient;
import com.bass.bbs.ApiException;
import com.bass.bbs.Configuration;
import com.bass.bbs.models.*;
import com.bass.bbs.api.AuthService;
import com.bass.bbs.api.AuthService.*;

public class Example {
    public static void main(String[] args) {
        ApiClient defaultClient = Configuration.getDefaultApiClient();
        defaultClient.setBasePath("http://localhost");

        AuthService apiInstance = new AuthService(defaultClient);
        RegisterAccountReq registerAccountReq = new RegisterAccountReq(); // RegisterAccountReq | 
        try {
            APIregisterAccountRequest request = APIregisterAccountRequest.newBuilder()
                .registerAccountReq(registerAccountReq)
                .build();
            Object result = apiInstance.registerAccount(request);
            System.out.println(result);
        } catch (ApiException e) {
            System.err.println("Exception when calling AuthService#registerAccount");
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
| registerAccountRequest | [**APIregisterAccountRequest**](AuthService.md#APIregisterAccountRequest)|-|-|

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

## registerAccountWithHttpInfo

> ApiResponse<Object> registerAccountWithHttpInfo(registerAccountRequest)



注册账号。

### Example

```java
// Import classes:
import com.bass.bbs.ApiClient;
import com.bass.bbs.ApiException;
import com.bass.bbs.ApiResponse;
import com.bass.bbs.Configuration;
import com.bass.bbs.models.*;
import com.bass.bbs.api.AuthService;
import com.bass.bbs.api.AuthService.*;

public class Example {
    public static void main(String[] args) {
        ApiClient defaultClient = Configuration.getDefaultApiClient();
        defaultClient.setBasePath("http://localhost");

        AuthService apiInstance = new AuthService(defaultClient);
        RegisterAccountReq registerAccountReq = new RegisterAccountReq(); // RegisterAccountReq | 
        try {
            APIregisterAccountRequest request = APIregisterAccountRequest.newBuilder()
                .registerAccountReq(registerAccountReq)
                .build();
            ApiResponse<Object> response = apiInstance.registerAccountWithHttpInfo(request);
            System.out.println("Status code: " + response.getStatusCode());
            System.out.println("Response headers: " + response.getHeaders());
            System.out.println("Response body: " + response.getData());
        } catch (ApiException e) {
            System.err.println("Exception when calling AuthService#registerAccount");
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
| registerAccountRequest | [**APIregisterAccountRequest**](AuthService.md#APIregisterAccountRequest)|-|-|

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


<a id="APIregisterAccountRequest"></a>
## APIregisterAccountRequest
### Properties

|     Name      |    Type       | Description   |     Notes    |
| ------------- | ------------- | ------------- | -------------|
| **registerAccountReq** | [**RegisterAccountReq**](RegisterAccountReq.md) |  | |


