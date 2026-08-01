# AuthService

All URIs are relative to *http://localhost*

| Method | HTTP request | Description |
|------------- | ------------- | -------------|
| [**cancelAccount**](AuthService.md#cancelAccount) | **POST** /v1/user/auth/cancel-account |  |
| [**cancelAccountWithHttpInfo**](AuthService.md#cancelAccountWithHttpInfo) | **POST** /v1/user/auth/cancel-account |  |
| [**login**](AuthService.md#login) | **POST** /v1/user/auth/login |  |
| [**loginWithHttpInfo**](AuthService.md#loginWithHttpInfo) | **POST** /v1/user/auth/login |  |
| [**logout**](AuthService.md#logout) | **POST** /v1/user/auth/logout |  |
| [**logoutWithHttpInfo**](AuthService.md#logoutWithHttpInfo) | **POST** /v1/user/auth/logout |  |
| [**refreshToken**](AuthService.md#refreshToken) | **POST** /v1/user/auth/refresh-token |  |
| [**refreshTokenWithHttpInfo**](AuthService.md#refreshTokenWithHttpInfo) | **POST** /v1/user/auth/refresh-token |  |
| [**startEmailRegistration**](AuthService.md#startEmailRegistration) | **POST** /v1/user/auth/start-email-registration |  |
| [**startEmailRegistrationWithHttpInfo**](AuthService.md#startEmailRegistrationWithHttpInfo) | **POST** /v1/user/auth/start-email-registration |  |
| [**startPhoneRegistration**](AuthService.md#startPhoneRegistration) | **POST** /v1/user/auth/start-phone-registration |  |
| [**startPhoneRegistrationWithHttpInfo**](AuthService.md#startPhoneRegistrationWithHttpInfo) | **POST** /v1/user/auth/start-phone-registration |  |
| [**verifyEmailRegistration**](AuthService.md#verifyEmailRegistration) | **POST** /v1/user/auth/verify-email-registration |  |
| [**verifyEmailRegistrationWithHttpInfo**](AuthService.md#verifyEmailRegistrationWithHttpInfo) | **POST** /v1/user/auth/verify-email-registration |  |
| [**verifyPhoneRegistration**](AuthService.md#verifyPhoneRegistration) | **POST** /v1/user/auth/verify-phone-registration |  |
| [**verifyPhoneRegistrationWithHttpInfo**](AuthService.md#verifyPhoneRegistrationWithHttpInfo) | **POST** /v1/user/auth/verify-phone-registration |  |



## cancelAccount

> Object cancelAccount(cancelAccountRequest)



注销账号。

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
        CancelAccountReq cancelAccountReq = new CancelAccountReq(); // CancelAccountReq | 
        try {
            APIcancelAccountRequest request = APIcancelAccountRequest.newBuilder()
                .cancelAccountReq(cancelAccountReq)
                .build();
            Object result = apiInstance.cancelAccount(request);
            System.out.println(result);
        } catch (ApiException e) {
            System.err.println("Exception when calling AuthService#cancelAccount");
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
| cancelAccountRequest | [**APIcancelAccountRequest**](AuthService.md#APIcancelAccountRequest)|-|-|

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

## cancelAccountWithHttpInfo

> ApiResponse<Object> cancelAccountWithHttpInfo(cancelAccountRequest)



注销账号。

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
        CancelAccountReq cancelAccountReq = new CancelAccountReq(); // CancelAccountReq | 
        try {
            APIcancelAccountRequest request = APIcancelAccountRequest.newBuilder()
                .cancelAccountReq(cancelAccountReq)
                .build();
            ApiResponse<Object> response = apiInstance.cancelAccountWithHttpInfo(request);
            System.out.println("Status code: " + response.getStatusCode());
            System.out.println("Response headers: " + response.getHeaders());
            System.out.println("Response body: " + response.getData());
        } catch (ApiException e) {
            System.err.println("Exception when calling AuthService#cancelAccount");
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
| cancelAccountRequest | [**APIcancelAccountRequest**](AuthService.md#APIcancelAccountRequest)|-|-|

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


<a id="APIcancelAccountRequest"></a>
## APIcancelAccountRequest
### Properties

|     Name      |    Type       | Description   |     Notes    |
| ------------- | ------------- | ------------- | -------------|
| **cancelAccountReq** | [**CancelAccountReq**](CancelAccountReq.md) |  | |



## login

> LoginResp login(loginRequest)



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
        LoginReq loginReq = new LoginReq(); // LoginReq | 
        try {
            APIloginRequest request = APIloginRequest.newBuilder()
                .loginReq(loginReq)
                .build();
            LoginResp result = apiInstance.login(request);
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

[**LoginResp**](LoginResp.md)


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

> ApiResponse<LoginResp> loginWithHttpInfo(loginRequest)



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
        LoginReq loginReq = new LoginReq(); // LoginReq | 
        try {
            APIloginRequest request = APIloginRequest.newBuilder()
                .loginReq(loginReq)
                .build();
            ApiResponse<LoginResp> response = apiInstance.loginWithHttpInfo(request);
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

ApiResponse<[**LoginResp**](LoginResp.md)>


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
| **loginReq** | [**LoginReq**](LoginReq.md) |  | |



## logout

> Object logout(logoutRequest)



退出登录。

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
        Object body = null; // Object | 
        try {
            APIlogoutRequest request = APIlogoutRequest.newBuilder()
                .body(body)
                .build();
            Object result = apiInstance.logout(request);
            System.out.println(result);
        } catch (ApiException e) {
            System.err.println("Exception when calling AuthService#logout");
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
| logoutRequest | [**APIlogoutRequest**](AuthService.md#APIlogoutRequest)|-|-|

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

## logoutWithHttpInfo

> ApiResponse<Object> logoutWithHttpInfo(logoutRequest)



退出登录。

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
        Object body = null; // Object | 
        try {
            APIlogoutRequest request = APIlogoutRequest.newBuilder()
                .body(body)
                .build();
            ApiResponse<Object> response = apiInstance.logoutWithHttpInfo(request);
            System.out.println("Status code: " + response.getStatusCode());
            System.out.println("Response headers: " + response.getHeaders());
            System.out.println("Response body: " + response.getData());
        } catch (ApiException e) {
            System.err.println("Exception when calling AuthService#logout");
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
| logoutRequest | [**APIlogoutRequest**](AuthService.md#APIlogoutRequest)|-|-|

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


<a id="APIlogoutRequest"></a>
## APIlogoutRequest
### Properties

|     Name      |    Type       | Description   |     Notes    |
| ------------- | ------------- | ------------- | -------------|
| **body** | **Object** |  | |



## refreshToken

> RefreshTokenResp refreshToken(refreshTokenRequest)



刷新登录令牌。

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
        RefreshTokenReq refreshTokenReq = new RefreshTokenReq(); // RefreshTokenReq | 
        try {
            APIrefreshTokenRequest request = APIrefreshTokenRequest.newBuilder()
                .refreshTokenReq(refreshTokenReq)
                .build();
            RefreshTokenResp result = apiInstance.refreshToken(request);
            System.out.println(result);
        } catch (ApiException e) {
            System.err.println("Exception when calling AuthService#refreshToken");
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
| refreshTokenRequest | [**APIrefreshTokenRequest**](AuthService.md#APIrefreshTokenRequest)|-|-|

### Return type

[**RefreshTokenResp**](RefreshTokenResp.md)


### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

### HTTP response details
| Status code | Description | Response headers |
|-------------|-------------|------------------|
| **200** | OK |  -  |

## refreshTokenWithHttpInfo

> ApiResponse<RefreshTokenResp> refreshTokenWithHttpInfo(refreshTokenRequest)



刷新登录令牌。

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
        RefreshTokenReq refreshTokenReq = new RefreshTokenReq(); // RefreshTokenReq | 
        try {
            APIrefreshTokenRequest request = APIrefreshTokenRequest.newBuilder()
                .refreshTokenReq(refreshTokenReq)
                .build();
            ApiResponse<RefreshTokenResp> response = apiInstance.refreshTokenWithHttpInfo(request);
            System.out.println("Status code: " + response.getStatusCode());
            System.out.println("Response headers: " + response.getHeaders());
            System.out.println("Response body: " + response.getData());
        } catch (ApiException e) {
            System.err.println("Exception when calling AuthService#refreshToken");
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
| refreshTokenRequest | [**APIrefreshTokenRequest**](AuthService.md#APIrefreshTokenRequest)|-|-|

### Return type

ApiResponse<[**RefreshTokenResp**](RefreshTokenResp.md)>


### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

### HTTP response details
| Status code | Description | Response headers |
|-------------|-------------|------------------|
| **200** | OK |  -  |


<a id="APIrefreshTokenRequest"></a>
## APIrefreshTokenRequest
### Properties

|     Name      |    Type       | Description   |     Notes    |
| ------------- | ------------- | ------------- | -------------|
| **refreshTokenReq** | [**RefreshTokenReq**](RefreshTokenReq.md) |  | |



## startEmailRegistration

> StartEmailRegistrationResp startEmailRegistration(startEmailRegistrationRequest)



开始邮箱注册。

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
        StartEmailRegistrationReq startEmailRegistrationReq = new StartEmailRegistrationReq(); // StartEmailRegistrationReq | 
        try {
            APIstartEmailRegistrationRequest request = APIstartEmailRegistrationRequest.newBuilder()
                .startEmailRegistrationReq(startEmailRegistrationReq)
                .build();
            StartEmailRegistrationResp result = apiInstance.startEmailRegistration(request);
            System.out.println(result);
        } catch (ApiException e) {
            System.err.println("Exception when calling AuthService#startEmailRegistration");
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
| startEmailRegistrationRequest | [**APIstartEmailRegistrationRequest**](AuthService.md#APIstartEmailRegistrationRequest)|-|-|

### Return type

[**StartEmailRegistrationResp**](StartEmailRegistrationResp.md)


### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

### HTTP response details
| Status code | Description | Response headers |
|-------------|-------------|------------------|
| **200** | OK |  -  |

## startEmailRegistrationWithHttpInfo

> ApiResponse<StartEmailRegistrationResp> startEmailRegistrationWithHttpInfo(startEmailRegistrationRequest)



开始邮箱注册。

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
        StartEmailRegistrationReq startEmailRegistrationReq = new StartEmailRegistrationReq(); // StartEmailRegistrationReq | 
        try {
            APIstartEmailRegistrationRequest request = APIstartEmailRegistrationRequest.newBuilder()
                .startEmailRegistrationReq(startEmailRegistrationReq)
                .build();
            ApiResponse<StartEmailRegistrationResp> response = apiInstance.startEmailRegistrationWithHttpInfo(request);
            System.out.println("Status code: " + response.getStatusCode());
            System.out.println("Response headers: " + response.getHeaders());
            System.out.println("Response body: " + response.getData());
        } catch (ApiException e) {
            System.err.println("Exception when calling AuthService#startEmailRegistration");
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
| startEmailRegistrationRequest | [**APIstartEmailRegistrationRequest**](AuthService.md#APIstartEmailRegistrationRequest)|-|-|

### Return type

ApiResponse<[**StartEmailRegistrationResp**](StartEmailRegistrationResp.md)>


### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

### HTTP response details
| Status code | Description | Response headers |
|-------------|-------------|------------------|
| **200** | OK |  -  |


<a id="APIstartEmailRegistrationRequest"></a>
## APIstartEmailRegistrationRequest
### Properties

|     Name      |    Type       | Description   |     Notes    |
| ------------- | ------------- | ------------- | -------------|
| **startEmailRegistrationReq** | [**StartEmailRegistrationReq**](StartEmailRegistrationReq.md) |  | |



## startPhoneRegistration

> StartPhoneRegistrationResp startPhoneRegistration(startPhoneRegistrationRequest)



开始手机注册。

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
        StartPhoneRegistrationReq startPhoneRegistrationReq = new StartPhoneRegistrationReq(); // StartPhoneRegistrationReq | 
        try {
            APIstartPhoneRegistrationRequest request = APIstartPhoneRegistrationRequest.newBuilder()
                .startPhoneRegistrationReq(startPhoneRegistrationReq)
                .build();
            StartPhoneRegistrationResp result = apiInstance.startPhoneRegistration(request);
            System.out.println(result);
        } catch (ApiException e) {
            System.err.println("Exception when calling AuthService#startPhoneRegistration");
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
| startPhoneRegistrationRequest | [**APIstartPhoneRegistrationRequest**](AuthService.md#APIstartPhoneRegistrationRequest)|-|-|

### Return type

[**StartPhoneRegistrationResp**](StartPhoneRegistrationResp.md)


### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

### HTTP response details
| Status code | Description | Response headers |
|-------------|-------------|------------------|
| **200** | OK |  -  |

## startPhoneRegistrationWithHttpInfo

> ApiResponse<StartPhoneRegistrationResp> startPhoneRegistrationWithHttpInfo(startPhoneRegistrationRequest)



开始手机注册。

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
        StartPhoneRegistrationReq startPhoneRegistrationReq = new StartPhoneRegistrationReq(); // StartPhoneRegistrationReq | 
        try {
            APIstartPhoneRegistrationRequest request = APIstartPhoneRegistrationRequest.newBuilder()
                .startPhoneRegistrationReq(startPhoneRegistrationReq)
                .build();
            ApiResponse<StartPhoneRegistrationResp> response = apiInstance.startPhoneRegistrationWithHttpInfo(request);
            System.out.println("Status code: " + response.getStatusCode());
            System.out.println("Response headers: " + response.getHeaders());
            System.out.println("Response body: " + response.getData());
        } catch (ApiException e) {
            System.err.println("Exception when calling AuthService#startPhoneRegistration");
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
| startPhoneRegistrationRequest | [**APIstartPhoneRegistrationRequest**](AuthService.md#APIstartPhoneRegistrationRequest)|-|-|

### Return type

ApiResponse<[**StartPhoneRegistrationResp**](StartPhoneRegistrationResp.md)>


### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

### HTTP response details
| Status code | Description | Response headers |
|-------------|-------------|------------------|
| **200** | OK |  -  |


<a id="APIstartPhoneRegistrationRequest"></a>
## APIstartPhoneRegistrationRequest
### Properties

|     Name      |    Type       | Description   |     Notes    |
| ------------- | ------------- | ------------- | -------------|
| **startPhoneRegistrationReq** | [**StartPhoneRegistrationReq**](StartPhoneRegistrationReq.md) |  | |



## verifyEmailRegistration

> Object verifyEmailRegistration(verifyEmailRegistrationRequest)



校验邮箱注册验证码。

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
        VerifyEmailRegistrationReq verifyEmailRegistrationReq = new VerifyEmailRegistrationReq(); // VerifyEmailRegistrationReq | 
        try {
            APIverifyEmailRegistrationRequest request = APIverifyEmailRegistrationRequest.newBuilder()
                .verifyEmailRegistrationReq(verifyEmailRegistrationReq)
                .build();
            Object result = apiInstance.verifyEmailRegistration(request);
            System.out.println(result);
        } catch (ApiException e) {
            System.err.println("Exception when calling AuthService#verifyEmailRegistration");
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
| verifyEmailRegistrationRequest | [**APIverifyEmailRegistrationRequest**](AuthService.md#APIverifyEmailRegistrationRequest)|-|-|

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

## verifyEmailRegistrationWithHttpInfo

> ApiResponse<Object> verifyEmailRegistrationWithHttpInfo(verifyEmailRegistrationRequest)



校验邮箱注册验证码。

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
        VerifyEmailRegistrationReq verifyEmailRegistrationReq = new VerifyEmailRegistrationReq(); // VerifyEmailRegistrationReq | 
        try {
            APIverifyEmailRegistrationRequest request = APIverifyEmailRegistrationRequest.newBuilder()
                .verifyEmailRegistrationReq(verifyEmailRegistrationReq)
                .build();
            ApiResponse<Object> response = apiInstance.verifyEmailRegistrationWithHttpInfo(request);
            System.out.println("Status code: " + response.getStatusCode());
            System.out.println("Response headers: " + response.getHeaders());
            System.out.println("Response body: " + response.getData());
        } catch (ApiException e) {
            System.err.println("Exception when calling AuthService#verifyEmailRegistration");
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
| verifyEmailRegistrationRequest | [**APIverifyEmailRegistrationRequest**](AuthService.md#APIverifyEmailRegistrationRequest)|-|-|

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


<a id="APIverifyEmailRegistrationRequest"></a>
## APIverifyEmailRegistrationRequest
### Properties

|     Name      |    Type       | Description   |     Notes    |
| ------------- | ------------- | ------------- | -------------|
| **verifyEmailRegistrationReq** | [**VerifyEmailRegistrationReq**](VerifyEmailRegistrationReq.md) |  | |



## verifyPhoneRegistration

> Object verifyPhoneRegistration(verifyPhoneRegistrationRequest)



校验手机注册验证码。

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
        VerifyPhoneRegistrationReq verifyPhoneRegistrationReq = new VerifyPhoneRegistrationReq(); // VerifyPhoneRegistrationReq | 
        try {
            APIverifyPhoneRegistrationRequest request = APIverifyPhoneRegistrationRequest.newBuilder()
                .verifyPhoneRegistrationReq(verifyPhoneRegistrationReq)
                .build();
            Object result = apiInstance.verifyPhoneRegistration(request);
            System.out.println(result);
        } catch (ApiException e) {
            System.err.println("Exception when calling AuthService#verifyPhoneRegistration");
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
| verifyPhoneRegistrationRequest | [**APIverifyPhoneRegistrationRequest**](AuthService.md#APIverifyPhoneRegistrationRequest)|-|-|

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

## verifyPhoneRegistrationWithHttpInfo

> ApiResponse<Object> verifyPhoneRegistrationWithHttpInfo(verifyPhoneRegistrationRequest)



校验手机注册验证码。

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
        VerifyPhoneRegistrationReq verifyPhoneRegistrationReq = new VerifyPhoneRegistrationReq(); // VerifyPhoneRegistrationReq | 
        try {
            APIverifyPhoneRegistrationRequest request = APIverifyPhoneRegistrationRequest.newBuilder()
                .verifyPhoneRegistrationReq(verifyPhoneRegistrationReq)
                .build();
            ApiResponse<Object> response = apiInstance.verifyPhoneRegistrationWithHttpInfo(request);
            System.out.println("Status code: " + response.getStatusCode());
            System.out.println("Response headers: " + response.getHeaders());
            System.out.println("Response body: " + response.getData());
        } catch (ApiException e) {
            System.err.println("Exception when calling AuthService#verifyPhoneRegistration");
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
| verifyPhoneRegistrationRequest | [**APIverifyPhoneRegistrationRequest**](AuthService.md#APIverifyPhoneRegistrationRequest)|-|-|

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


<a id="APIverifyPhoneRegistrationRequest"></a>
## APIverifyPhoneRegistrationRequest
### Properties

|     Name      |    Type       | Description   |     Notes    |
| ------------- | ------------- | ------------- | -------------|
| **verifyPhoneRegistrationReq** | [**VerifyPhoneRegistrationReq**](VerifyPhoneRegistrationReq.md) |  | |


