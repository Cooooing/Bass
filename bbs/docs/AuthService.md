# AuthService

All URIs are relative to *http://localhost*

| Method | HTTP request | Description |
|------------- | ------------- | -------------|
| [**loginByPassword**](AuthService.md#loginByPassword) | **POST** /v1/user/auth/login-by-password |  |
| [**loginByPasswordWithHttpInfo**](AuthService.md#loginByPasswordWithHttpInfo) | **POST** /v1/user/auth/login-by-password |  |
| [**logout**](AuthService.md#logout) | **POST** /v1/user/auth/logout |  |
| [**logoutWithHttpInfo**](AuthService.md#logoutWithHttpInfo) | **POST** /v1/user/auth/logout |  |
| [**startEmailRegistration**](AuthService.md#startEmailRegistration) | **POST** /v1/user/auth/start-email-registration |  |
| [**startEmailRegistrationWithHttpInfo**](AuthService.md#startEmailRegistrationWithHttpInfo) | **POST** /v1/user/auth/start-email-registration |  |
| [**startPhoneRegistration**](AuthService.md#startPhoneRegistration) | **POST** /v1/user/auth/start-phone-registration |  |
| [**startPhoneRegistrationWithHttpInfo**](AuthService.md#startPhoneRegistrationWithHttpInfo) | **POST** /v1/user/auth/start-phone-registration |  |
| [**verifyEmailRegistration**](AuthService.md#verifyEmailRegistration) | **POST** /v1/user/auth/verify-email-registration |  |
| [**verifyEmailRegistrationWithHttpInfo**](AuthService.md#verifyEmailRegistrationWithHttpInfo) | **POST** /v1/user/auth/verify-email-registration |  |
| [**verifyPhoneRegistration**](AuthService.md#verifyPhoneRegistration) | **POST** /v1/user/auth/verify-phone-registration |  |
| [**verifyPhoneRegistrationWithHttpInfo**](AuthService.md#verifyPhoneRegistrationWithHttpInfo) | **POST** /v1/user/auth/verify-phone-registration |  |



## loginByPassword

> LoginByPasswordReply loginByPassword(loginByPasswordRequest)



使用密码登录账号。

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
        LoginByPasswordRequest loginByPasswordRequest = new LoginByPasswordRequest(); // LoginByPasswordRequest | 
        try {
            APIloginByPasswordRequest request = APIloginByPasswordRequest.newBuilder()
                .loginByPasswordRequest(loginByPasswordRequest)
                .build();
            LoginByPasswordReply result = apiInstance.loginByPassword(request);
            System.out.println(result);
        } catch (ApiException e) {
            System.err.println("Exception when calling AuthService#loginByPassword");
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
| loginByPasswordRequest | [**APIloginByPasswordRequest**](AuthService.md#APIloginByPasswordRequest)|-|-|

### Return type

[**LoginByPasswordReply**](LoginByPasswordReply.md)


### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

### HTTP response details
| Status code | Description | Response headers |
|-------------|-------------|------------------|
| **200** | OK |  -  |

## loginByPasswordWithHttpInfo

> ApiResponse<LoginByPasswordReply> loginByPasswordWithHttpInfo(loginByPasswordRequest)



使用密码登录账号。

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
        LoginByPasswordRequest loginByPasswordRequest = new LoginByPasswordRequest(); // LoginByPasswordRequest | 
        try {
            APIloginByPasswordRequest request = APIloginByPasswordRequest.newBuilder()
                .loginByPasswordRequest(loginByPasswordRequest)
                .build();
            ApiResponse<LoginByPasswordReply> response = apiInstance.loginByPasswordWithHttpInfo(request);
            System.out.println("Status code: " + response.getStatusCode());
            System.out.println("Response headers: " + response.getHeaders());
            System.out.println("Response body: " + response.getData());
        } catch (ApiException e) {
            System.err.println("Exception when calling AuthService#loginByPassword");
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
| loginByPasswordRequest | [**APIloginByPasswordRequest**](AuthService.md#APIloginByPasswordRequest)|-|-|

### Return type

ApiResponse<[**LoginByPasswordReply**](LoginByPasswordReply.md)>


### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

### HTTP response details
| Status code | Description | Response headers |
|-------------|-------------|------------------|
| **200** | OK |  -  |


<a id="APIloginByPasswordRequest"></a>
## APIloginByPasswordRequest
### Properties

|     Name      |    Type       | Description   |     Notes    |
| ------------- | ------------- | ------------- | -------------|
| **loginByPasswordRequest** | [**LoginByPasswordRequest**](LoginByPasswordRequest.md) |  | |



## logout

> Object logout(logoutRequest)



登出当前账号。

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



登出当前账号。

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



## startEmailRegistration

> StartEmailRegistrationReply startEmailRegistration(startEmailRegistrationRequest)



使用邮箱发起账号注册。

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
        StartEmailRegistrationRequest startEmailRegistrationRequest = new StartEmailRegistrationRequest(); // StartEmailRegistrationRequest | 
        try {
            APIstartEmailRegistrationRequest request = APIstartEmailRegistrationRequest.newBuilder()
                .startEmailRegistrationRequest(startEmailRegistrationRequest)
                .build();
            StartEmailRegistrationReply result = apiInstance.startEmailRegistration(request);
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

[**StartEmailRegistrationReply**](StartEmailRegistrationReply.md)


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

> ApiResponse<StartEmailRegistrationReply> startEmailRegistrationWithHttpInfo(startEmailRegistrationRequest)



使用邮箱发起账号注册。

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
        StartEmailRegistrationRequest startEmailRegistrationRequest = new StartEmailRegistrationRequest(); // StartEmailRegistrationRequest | 
        try {
            APIstartEmailRegistrationRequest request = APIstartEmailRegistrationRequest.newBuilder()
                .startEmailRegistrationRequest(startEmailRegistrationRequest)
                .build();
            ApiResponse<StartEmailRegistrationReply> response = apiInstance.startEmailRegistrationWithHttpInfo(request);
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

ApiResponse<[**StartEmailRegistrationReply**](StartEmailRegistrationReply.md)>


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
| **startEmailRegistrationRequest** | [**StartEmailRegistrationRequest**](StartEmailRegistrationRequest.md) |  | |



## startPhoneRegistration

> StartPhoneRegistrationReply startPhoneRegistration(startPhoneRegistrationRequest)



使用手机号发起账号注册。

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
        StartPhoneRegistrationRequest startPhoneRegistrationRequest = new StartPhoneRegistrationRequest(); // StartPhoneRegistrationRequest | 
        try {
            APIstartPhoneRegistrationRequest request = APIstartPhoneRegistrationRequest.newBuilder()
                .startPhoneRegistrationRequest(startPhoneRegistrationRequest)
                .build();
            StartPhoneRegistrationReply result = apiInstance.startPhoneRegistration(request);
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

[**StartPhoneRegistrationReply**](StartPhoneRegistrationReply.md)


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

> ApiResponse<StartPhoneRegistrationReply> startPhoneRegistrationWithHttpInfo(startPhoneRegistrationRequest)



使用手机号发起账号注册。

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
        StartPhoneRegistrationRequest startPhoneRegistrationRequest = new StartPhoneRegistrationRequest(); // StartPhoneRegistrationRequest | 
        try {
            APIstartPhoneRegistrationRequest request = APIstartPhoneRegistrationRequest.newBuilder()
                .startPhoneRegistrationRequest(startPhoneRegistrationRequest)
                .build();
            ApiResponse<StartPhoneRegistrationReply> response = apiInstance.startPhoneRegistrationWithHttpInfo(request);
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

ApiResponse<[**StartPhoneRegistrationReply**](StartPhoneRegistrationReply.md)>


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
| **startPhoneRegistrationRequest** | [**StartPhoneRegistrationRequest**](StartPhoneRegistrationRequest.md) |  | |



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
        VerifyEmailRegistrationRequest verifyEmailRegistrationRequest = new VerifyEmailRegistrationRequest(); // VerifyEmailRegistrationRequest | 
        try {
            APIverifyEmailRegistrationRequest request = APIverifyEmailRegistrationRequest.newBuilder()
                .verifyEmailRegistrationRequest(verifyEmailRegistrationRequest)
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
        VerifyEmailRegistrationRequest verifyEmailRegistrationRequest = new VerifyEmailRegistrationRequest(); // VerifyEmailRegistrationRequest | 
        try {
            APIverifyEmailRegistrationRequest request = APIverifyEmailRegistrationRequest.newBuilder()
                .verifyEmailRegistrationRequest(verifyEmailRegistrationRequest)
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
| **verifyEmailRegistrationRequest** | [**VerifyEmailRegistrationRequest**](VerifyEmailRegistrationRequest.md) |  | |



## verifyPhoneRegistration

> Object verifyPhoneRegistration(verifyPhoneRegistrationRequest)



校验手机号注册验证码。

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
        VerifyPhoneRegistrationRequest verifyPhoneRegistrationRequest = new VerifyPhoneRegistrationRequest(); // VerifyPhoneRegistrationRequest | 
        try {
            APIverifyPhoneRegistrationRequest request = APIverifyPhoneRegistrationRequest.newBuilder()
                .verifyPhoneRegistrationRequest(verifyPhoneRegistrationRequest)
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



校验手机号注册验证码。

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
        VerifyPhoneRegistrationRequest verifyPhoneRegistrationRequest = new VerifyPhoneRegistrationRequest(); // VerifyPhoneRegistrationRequest | 
        try {
            APIverifyPhoneRegistrationRequest request = APIverifyPhoneRegistrationRequest.newBuilder()
                .verifyPhoneRegistrationRequest(verifyPhoneRegistrationRequest)
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
| **verifyPhoneRegistrationRequest** | [**VerifyPhoneRegistrationRequest**](VerifyPhoneRegistrationRequest.md) |  | |


