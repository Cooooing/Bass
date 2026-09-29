# AuthService

All URIs are relative to *http://localhost*

| Method | HTTP request | Description |
|------------- | ------------- | -------------|
| [**login**](AuthService.md#login) | **POST** /v1/game-idle/auth/login |  |
| [**loginWithHttpInfo**](AuthService.md#loginWithHttpInfo) | **POST** /v1/game-idle/auth/login |  |
| [**register**](AuthService.md#register) | **POST** /v1/game-idle/auth/register |  |
| [**registerWithHttpInfo**](AuthService.md#registerWithHttpInfo) | **POST** /v1/game-idle/auth/register |  |



## login

> LoginAccountResp login(loginRequest)



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
            APIloginRequest request = APIloginRequest.newBuilder()
                .loginAccountReq(loginAccountReq)
                .build();
            LoginAccountResp result = apiInstance.login(request);
            System.out.println(result);
        } catch (ApiException e) {
            System.err.println("Exception when calling AuthService#login");
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
| loginRequest | [**APIloginRequest**](AuthService.md#APIloginRequest)|-|-|

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

## loginWithHttpInfo

> ApiResponse<LoginAccountResp> loginWithHttpInfo(loginRequest)



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
            APIloginRequest request = APIloginRequest.newBuilder()
                .loginAccountReq(loginAccountReq)
                .build();
            ApiResponse<LoginAccountResp> response = apiInstance.loginWithHttpInfo(request);
            System.out.println("Status code: " + response.getStatusCode());
            System.out.println("Response headers: " + response.getHeaders());
            System.out.println("Response body: " + response.getData());
        } catch (ApiException e) {
            System.err.println("Exception when calling AuthService#login");
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
| loginRequest | [**APIloginRequest**](AuthService.md#APIloginRequest)|-|-|

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


<a id="APIloginRequest"></a>
## APIloginRequest
### Properties

|     Name      |    Type       | Description   |     Notes    |
| ------------- | ------------- | ------------- | -------------|
| **loginAccountReq** | [**LoginAccountReq**](LoginAccountReq.md) |  | |



## register

> Object register(registerRequest)



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
            APIregisterRequest request = APIregisterRequest.newBuilder()
                .registerAccountReq(registerAccountReq)
                .build();
            Object result = apiInstance.register(request);
            System.out.println(result);
        } catch (ApiException e) {
            System.err.println("Exception when calling AuthService#register");
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
| registerRequest | [**APIregisterRequest**](AuthService.md#APIregisterRequest)|-|-|

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

## registerWithHttpInfo

> ApiResponse<Object> registerWithHttpInfo(registerRequest)



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
            APIregisterRequest request = APIregisterRequest.newBuilder()
                .registerAccountReq(registerAccountReq)
                .build();
            ApiResponse<Object> response = apiInstance.registerWithHttpInfo(request);
            System.out.println("Status code: " + response.getStatusCode());
            System.out.println("Response headers: " + response.getHeaders());
            System.out.println("Response body: " + response.getData());
        } catch (ApiException e) {
            System.err.println("Exception when calling AuthService#register");
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
| registerRequest | [**APIregisterRequest**](AuthService.md#APIregisterRequest)|-|-|

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


<a id="APIregisterRequest"></a>
## APIregisterRequest
### Properties

|     Name      |    Type       | Description   |     Notes    |
| ------------- | ------------- | ------------- | -------------|
| **registerAccountReq** | [**RegisterAccountReq**](RegisterAccountReq.md) |  | |


