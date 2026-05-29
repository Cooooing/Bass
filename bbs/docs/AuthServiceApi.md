# AuthServiceApi

All URIs are relative to *http://localhost*

| Method | HTTP request | Description |
|------------- | ------------- | -------------|
| [**authServiceLoginByPassword**](AuthServiceApi.md#authServiceLoginByPassword) | **POST** /v1/user/auth/login-by-password |  |
| [**authServiceLoginByPasswordWithHttpInfo**](AuthServiceApi.md#authServiceLoginByPasswordWithHttpInfo) | **POST** /v1/user/auth/login-by-password |  |
| [**authServiceLogout**](AuthServiceApi.md#authServiceLogout) | **POST** /v1/user/auth/logout |  |
| [**authServiceLogoutWithHttpInfo**](AuthServiceApi.md#authServiceLogoutWithHttpInfo) | **POST** /v1/user/auth/logout |  |
| [**authServiceStartEmailRegistration**](AuthServiceApi.md#authServiceStartEmailRegistration) | **POST** /v1/user/auth/start-email-registration |  |
| [**authServiceStartEmailRegistrationWithHttpInfo**](AuthServiceApi.md#authServiceStartEmailRegistrationWithHttpInfo) | **POST** /v1/user/auth/start-email-registration |  |
| [**authServiceStartPhoneRegistration**](AuthServiceApi.md#authServiceStartPhoneRegistration) | **POST** /v1/user/auth/start-phone-registration |  |
| [**authServiceStartPhoneRegistrationWithHttpInfo**](AuthServiceApi.md#authServiceStartPhoneRegistrationWithHttpInfo) | **POST** /v1/user/auth/start-phone-registration |  |
| [**authServiceVerifyEmailRegistration**](AuthServiceApi.md#authServiceVerifyEmailRegistration) | **POST** /v1/user/auth/verify-email-registration |  |
| [**authServiceVerifyEmailRegistrationWithHttpInfo**](AuthServiceApi.md#authServiceVerifyEmailRegistrationWithHttpInfo) | **POST** /v1/user/auth/verify-email-registration |  |
| [**authServiceVerifyPhoneRegistration**](AuthServiceApi.md#authServiceVerifyPhoneRegistration) | **POST** /v1/user/auth/verify-phone-registration |  |
| [**authServiceVerifyPhoneRegistrationWithHttpInfo**](AuthServiceApi.md#authServiceVerifyPhoneRegistrationWithHttpInfo) | **POST** /v1/user/auth/verify-phone-registration |  |



## authServiceLoginByPassword

> LoginByPasswordReply authServiceLoginByPassword(authServiceLoginByPasswordRequest)



使用密码登录账号

### Example

```java
// Import classes:
import com.bass.bbs.ApiClient;
import com.bass.bbs.ApiException;
import com.bass.bbs.Configuration;
import com.bass.bbs.models.*;
import com.bass.bbs.api.AuthServiceApi;
import com.bass.bbs.api.AuthServiceApi.*;

public class Example {
    public static void main(String[] args) {
        ApiClient defaultClient = Configuration.getDefaultApiClient();
        defaultClient.setBasePath("http://localhost");

        AuthServiceApi apiInstance = new AuthServiceApi(defaultClient);
        LoginByPasswordRequest loginByPasswordRequest = new LoginByPasswordRequest(); // LoginByPasswordRequest | 
        try {
            APIauthServiceLoginByPasswordRequest request = APIauthServiceLoginByPasswordRequest.newBuilder()
                .loginByPasswordRequest(loginByPasswordRequest)
                .build();
            LoginByPasswordReply result = apiInstance.authServiceLoginByPassword(request);
            System.out.println(result);
        } catch (ApiException e) {
            System.err.println("Exception when calling AuthServiceApi#authServiceLoginByPassword");
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
| authServiceLoginByPasswordRequest | [**APIauthServiceLoginByPasswordRequest**](AuthServiceApi.md#APIauthServiceLoginByPasswordRequest)|-|-|

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

## authServiceLoginByPasswordWithHttpInfo

> ApiResponse<LoginByPasswordReply> authServiceLoginByPasswordWithHttpInfo(authServiceLoginByPasswordRequest)



使用密码登录账号

### Example

```java
// Import classes:
import com.bass.bbs.ApiClient;
import com.bass.bbs.ApiException;
import com.bass.bbs.ApiResponse;
import com.bass.bbs.Configuration;
import com.bass.bbs.models.*;
import com.bass.bbs.api.AuthServiceApi;
import com.bass.bbs.api.AuthServiceApi.*;

public class Example {
    public static void main(String[] args) {
        ApiClient defaultClient = Configuration.getDefaultApiClient();
        defaultClient.setBasePath("http://localhost");

        AuthServiceApi apiInstance = new AuthServiceApi(defaultClient);
        LoginByPasswordRequest loginByPasswordRequest = new LoginByPasswordRequest(); // LoginByPasswordRequest | 
        try {
            APIauthServiceLoginByPasswordRequest request = APIauthServiceLoginByPasswordRequest.newBuilder()
                .loginByPasswordRequest(loginByPasswordRequest)
                .build();
            ApiResponse<LoginByPasswordReply> response = apiInstance.authServiceLoginByPasswordWithHttpInfo(request);
            System.out.println("Status code: " + response.getStatusCode());
            System.out.println("Response headers: " + response.getHeaders());
            System.out.println("Response body: " + response.getData());
        } catch (ApiException e) {
            System.err.println("Exception when calling AuthServiceApi#authServiceLoginByPassword");
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
| authServiceLoginByPasswordRequest | [**APIauthServiceLoginByPasswordRequest**](AuthServiceApi.md#APIauthServiceLoginByPasswordRequest)|-|-|

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


<a id="APIauthServiceLoginByPasswordRequest"></a>
## APIauthServiceLoginByPasswordRequest
### Properties

|     Name      |    Type       | Description   |     Notes    |
| ------------- | ------------- | ------------- | -------------|
| **loginByPasswordRequest** | [**LoginByPasswordRequest**](LoginByPasswordRequest.md) |  | |



## authServiceLogout

> Object authServiceLogout(authServiceLogoutRequest)



登出当前登录账号

### Example

```java
// Import classes:
import com.bass.bbs.ApiClient;
import com.bass.bbs.ApiException;
import com.bass.bbs.Configuration;
import com.bass.bbs.models.*;
import com.bass.bbs.api.AuthServiceApi;
import com.bass.bbs.api.AuthServiceApi.*;

public class Example {
    public static void main(String[] args) {
        ApiClient defaultClient = Configuration.getDefaultApiClient();
        defaultClient.setBasePath("http://localhost");

        AuthServiceApi apiInstance = new AuthServiceApi(defaultClient);
        Object body = null; // Object | 
        try {
            APIauthServiceLogoutRequest request = APIauthServiceLogoutRequest.newBuilder()
                .body(body)
                .build();
            Object result = apiInstance.authServiceLogout(request);
            System.out.println(result);
        } catch (ApiException e) {
            System.err.println("Exception when calling AuthServiceApi#authServiceLogout");
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
| authServiceLogoutRequest | [**APIauthServiceLogoutRequest**](AuthServiceApi.md#APIauthServiceLogoutRequest)|-|-|

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

## authServiceLogoutWithHttpInfo

> ApiResponse<Object> authServiceLogoutWithHttpInfo(authServiceLogoutRequest)



登出当前登录账号

### Example

```java
// Import classes:
import com.bass.bbs.ApiClient;
import com.bass.bbs.ApiException;
import com.bass.bbs.ApiResponse;
import com.bass.bbs.Configuration;
import com.bass.bbs.models.*;
import com.bass.bbs.api.AuthServiceApi;
import com.bass.bbs.api.AuthServiceApi.*;

public class Example {
    public static void main(String[] args) {
        ApiClient defaultClient = Configuration.getDefaultApiClient();
        defaultClient.setBasePath("http://localhost");

        AuthServiceApi apiInstance = new AuthServiceApi(defaultClient);
        Object body = null; // Object | 
        try {
            APIauthServiceLogoutRequest request = APIauthServiceLogoutRequest.newBuilder()
                .body(body)
                .build();
            ApiResponse<Object> response = apiInstance.authServiceLogoutWithHttpInfo(request);
            System.out.println("Status code: " + response.getStatusCode());
            System.out.println("Response headers: " + response.getHeaders());
            System.out.println("Response body: " + response.getData());
        } catch (ApiException e) {
            System.err.println("Exception when calling AuthServiceApi#authServiceLogout");
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
| authServiceLogoutRequest | [**APIauthServiceLogoutRequest**](AuthServiceApi.md#APIauthServiceLogoutRequest)|-|-|

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


<a id="APIauthServiceLogoutRequest"></a>
## APIauthServiceLogoutRequest
### Properties

|     Name      |    Type       | Description   |     Notes    |
| ------------- | ------------- | ------------- | -------------|
| **body** | **Object** |  | |



## authServiceStartEmailRegistration

> StartEmailRegistrationReply authServiceStartEmailRegistration(authServiceStartEmailRegistrationRequest)



使用邮箱发起账号注册

### Example

```java
// Import classes:
import com.bass.bbs.ApiClient;
import com.bass.bbs.ApiException;
import com.bass.bbs.Configuration;
import com.bass.bbs.models.*;
import com.bass.bbs.api.AuthServiceApi;
import com.bass.bbs.api.AuthServiceApi.*;

public class Example {
    public static void main(String[] args) {
        ApiClient defaultClient = Configuration.getDefaultApiClient();
        defaultClient.setBasePath("http://localhost");

        AuthServiceApi apiInstance = new AuthServiceApi(defaultClient);
        StartEmailRegistrationRequest startEmailRegistrationRequest = new StartEmailRegistrationRequest(); // StartEmailRegistrationRequest | 
        try {
            APIauthServiceStartEmailRegistrationRequest request = APIauthServiceStartEmailRegistrationRequest.newBuilder()
                .startEmailRegistrationRequest(startEmailRegistrationRequest)
                .build();
            StartEmailRegistrationReply result = apiInstance.authServiceStartEmailRegistration(request);
            System.out.println(result);
        } catch (ApiException e) {
            System.err.println("Exception when calling AuthServiceApi#authServiceStartEmailRegistration");
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
| authServiceStartEmailRegistrationRequest | [**APIauthServiceStartEmailRegistrationRequest**](AuthServiceApi.md#APIauthServiceStartEmailRegistrationRequest)|-|-|

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

## authServiceStartEmailRegistrationWithHttpInfo

> ApiResponse<StartEmailRegistrationReply> authServiceStartEmailRegistrationWithHttpInfo(authServiceStartEmailRegistrationRequest)



使用邮箱发起账号注册

### Example

```java
// Import classes:
import com.bass.bbs.ApiClient;
import com.bass.bbs.ApiException;
import com.bass.bbs.ApiResponse;
import com.bass.bbs.Configuration;
import com.bass.bbs.models.*;
import com.bass.bbs.api.AuthServiceApi;
import com.bass.bbs.api.AuthServiceApi.*;

public class Example {
    public static void main(String[] args) {
        ApiClient defaultClient = Configuration.getDefaultApiClient();
        defaultClient.setBasePath("http://localhost");

        AuthServiceApi apiInstance = new AuthServiceApi(defaultClient);
        StartEmailRegistrationRequest startEmailRegistrationRequest = new StartEmailRegistrationRequest(); // StartEmailRegistrationRequest | 
        try {
            APIauthServiceStartEmailRegistrationRequest request = APIauthServiceStartEmailRegistrationRequest.newBuilder()
                .startEmailRegistrationRequest(startEmailRegistrationRequest)
                .build();
            ApiResponse<StartEmailRegistrationReply> response = apiInstance.authServiceStartEmailRegistrationWithHttpInfo(request);
            System.out.println("Status code: " + response.getStatusCode());
            System.out.println("Response headers: " + response.getHeaders());
            System.out.println("Response body: " + response.getData());
        } catch (ApiException e) {
            System.err.println("Exception when calling AuthServiceApi#authServiceStartEmailRegistration");
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
| authServiceStartEmailRegistrationRequest | [**APIauthServiceStartEmailRegistrationRequest**](AuthServiceApi.md#APIauthServiceStartEmailRegistrationRequest)|-|-|

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


<a id="APIauthServiceStartEmailRegistrationRequest"></a>
## APIauthServiceStartEmailRegistrationRequest
### Properties

|     Name      |    Type       | Description   |     Notes    |
| ------------- | ------------- | ------------- | -------------|
| **startEmailRegistrationRequest** | [**StartEmailRegistrationRequest**](StartEmailRegistrationRequest.md) |  | |



## authServiceStartPhoneRegistration

> StartPhoneRegistrationReply authServiceStartPhoneRegistration(authServiceStartPhoneRegistrationRequest)



使用手机号发起账号注册

### Example

```java
// Import classes:
import com.bass.bbs.ApiClient;
import com.bass.bbs.ApiException;
import com.bass.bbs.Configuration;
import com.bass.bbs.models.*;
import com.bass.bbs.api.AuthServiceApi;
import com.bass.bbs.api.AuthServiceApi.*;

public class Example {
    public static void main(String[] args) {
        ApiClient defaultClient = Configuration.getDefaultApiClient();
        defaultClient.setBasePath("http://localhost");

        AuthServiceApi apiInstance = new AuthServiceApi(defaultClient);
        StartPhoneRegistrationRequest startPhoneRegistrationRequest = new StartPhoneRegistrationRequest(); // StartPhoneRegistrationRequest | 
        try {
            APIauthServiceStartPhoneRegistrationRequest request = APIauthServiceStartPhoneRegistrationRequest.newBuilder()
                .startPhoneRegistrationRequest(startPhoneRegistrationRequest)
                .build();
            StartPhoneRegistrationReply result = apiInstance.authServiceStartPhoneRegistration(request);
            System.out.println(result);
        } catch (ApiException e) {
            System.err.println("Exception when calling AuthServiceApi#authServiceStartPhoneRegistration");
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
| authServiceStartPhoneRegistrationRequest | [**APIauthServiceStartPhoneRegistrationRequest**](AuthServiceApi.md#APIauthServiceStartPhoneRegistrationRequest)|-|-|

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

## authServiceStartPhoneRegistrationWithHttpInfo

> ApiResponse<StartPhoneRegistrationReply> authServiceStartPhoneRegistrationWithHttpInfo(authServiceStartPhoneRegistrationRequest)



使用手机号发起账号注册

### Example

```java
// Import classes:
import com.bass.bbs.ApiClient;
import com.bass.bbs.ApiException;
import com.bass.bbs.ApiResponse;
import com.bass.bbs.Configuration;
import com.bass.bbs.models.*;
import com.bass.bbs.api.AuthServiceApi;
import com.bass.bbs.api.AuthServiceApi.*;

public class Example {
    public static void main(String[] args) {
        ApiClient defaultClient = Configuration.getDefaultApiClient();
        defaultClient.setBasePath("http://localhost");

        AuthServiceApi apiInstance = new AuthServiceApi(defaultClient);
        StartPhoneRegistrationRequest startPhoneRegistrationRequest = new StartPhoneRegistrationRequest(); // StartPhoneRegistrationRequest | 
        try {
            APIauthServiceStartPhoneRegistrationRequest request = APIauthServiceStartPhoneRegistrationRequest.newBuilder()
                .startPhoneRegistrationRequest(startPhoneRegistrationRequest)
                .build();
            ApiResponse<StartPhoneRegistrationReply> response = apiInstance.authServiceStartPhoneRegistrationWithHttpInfo(request);
            System.out.println("Status code: " + response.getStatusCode());
            System.out.println("Response headers: " + response.getHeaders());
            System.out.println("Response body: " + response.getData());
        } catch (ApiException e) {
            System.err.println("Exception when calling AuthServiceApi#authServiceStartPhoneRegistration");
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
| authServiceStartPhoneRegistrationRequest | [**APIauthServiceStartPhoneRegistrationRequest**](AuthServiceApi.md#APIauthServiceStartPhoneRegistrationRequest)|-|-|

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


<a id="APIauthServiceStartPhoneRegistrationRequest"></a>
## APIauthServiceStartPhoneRegistrationRequest
### Properties

|     Name      |    Type       | Description   |     Notes    |
| ------------- | ------------- | ------------- | -------------|
| **startPhoneRegistrationRequest** | [**StartPhoneRegistrationRequest**](StartPhoneRegistrationRequest.md) |  | |



## authServiceVerifyEmailRegistration

> Object authServiceVerifyEmailRegistration(authServiceVerifyEmailRegistrationRequest)



校验邮箱注册验证码

### Example

```java
// Import classes:
import com.bass.bbs.ApiClient;
import com.bass.bbs.ApiException;
import com.bass.bbs.Configuration;
import com.bass.bbs.models.*;
import com.bass.bbs.api.AuthServiceApi;
import com.bass.bbs.api.AuthServiceApi.*;

public class Example {
    public static void main(String[] args) {
        ApiClient defaultClient = Configuration.getDefaultApiClient();
        defaultClient.setBasePath("http://localhost");

        AuthServiceApi apiInstance = new AuthServiceApi(defaultClient);
        VerifyEmailRegistrationRequest verifyEmailRegistrationRequest = new VerifyEmailRegistrationRequest(); // VerifyEmailRegistrationRequest | 
        try {
            APIauthServiceVerifyEmailRegistrationRequest request = APIauthServiceVerifyEmailRegistrationRequest.newBuilder()
                .verifyEmailRegistrationRequest(verifyEmailRegistrationRequest)
                .build();
            Object result = apiInstance.authServiceVerifyEmailRegistration(request);
            System.out.println(result);
        } catch (ApiException e) {
            System.err.println("Exception when calling AuthServiceApi#authServiceVerifyEmailRegistration");
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
| authServiceVerifyEmailRegistrationRequest | [**APIauthServiceVerifyEmailRegistrationRequest**](AuthServiceApi.md#APIauthServiceVerifyEmailRegistrationRequest)|-|-|

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

## authServiceVerifyEmailRegistrationWithHttpInfo

> ApiResponse<Object> authServiceVerifyEmailRegistrationWithHttpInfo(authServiceVerifyEmailRegistrationRequest)



校验邮箱注册验证码

### Example

```java
// Import classes:
import com.bass.bbs.ApiClient;
import com.bass.bbs.ApiException;
import com.bass.bbs.ApiResponse;
import com.bass.bbs.Configuration;
import com.bass.bbs.models.*;
import com.bass.bbs.api.AuthServiceApi;
import com.bass.bbs.api.AuthServiceApi.*;

public class Example {
    public static void main(String[] args) {
        ApiClient defaultClient = Configuration.getDefaultApiClient();
        defaultClient.setBasePath("http://localhost");

        AuthServiceApi apiInstance = new AuthServiceApi(defaultClient);
        VerifyEmailRegistrationRequest verifyEmailRegistrationRequest = new VerifyEmailRegistrationRequest(); // VerifyEmailRegistrationRequest | 
        try {
            APIauthServiceVerifyEmailRegistrationRequest request = APIauthServiceVerifyEmailRegistrationRequest.newBuilder()
                .verifyEmailRegistrationRequest(verifyEmailRegistrationRequest)
                .build();
            ApiResponse<Object> response = apiInstance.authServiceVerifyEmailRegistrationWithHttpInfo(request);
            System.out.println("Status code: " + response.getStatusCode());
            System.out.println("Response headers: " + response.getHeaders());
            System.out.println("Response body: " + response.getData());
        } catch (ApiException e) {
            System.err.println("Exception when calling AuthServiceApi#authServiceVerifyEmailRegistration");
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
| authServiceVerifyEmailRegistrationRequest | [**APIauthServiceVerifyEmailRegistrationRequest**](AuthServiceApi.md#APIauthServiceVerifyEmailRegistrationRequest)|-|-|

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


<a id="APIauthServiceVerifyEmailRegistrationRequest"></a>
## APIauthServiceVerifyEmailRegistrationRequest
### Properties

|     Name      |    Type       | Description   |     Notes    |
| ------------- | ------------- | ------------- | -------------|
| **verifyEmailRegistrationRequest** | [**VerifyEmailRegistrationRequest**](VerifyEmailRegistrationRequest.md) |  | |



## authServiceVerifyPhoneRegistration

> Object authServiceVerifyPhoneRegistration(authServiceVerifyPhoneRegistrationRequest)



校验手机号注册验证码

### Example

```java
// Import classes:
import com.bass.bbs.ApiClient;
import com.bass.bbs.ApiException;
import com.bass.bbs.Configuration;
import com.bass.bbs.models.*;
import com.bass.bbs.api.AuthServiceApi;
import com.bass.bbs.api.AuthServiceApi.*;

public class Example {
    public static void main(String[] args) {
        ApiClient defaultClient = Configuration.getDefaultApiClient();
        defaultClient.setBasePath("http://localhost");

        AuthServiceApi apiInstance = new AuthServiceApi(defaultClient);
        VerifyPhoneRegistrationRequest verifyPhoneRegistrationRequest = new VerifyPhoneRegistrationRequest(); // VerifyPhoneRegistrationRequest | 
        try {
            APIauthServiceVerifyPhoneRegistrationRequest request = APIauthServiceVerifyPhoneRegistrationRequest.newBuilder()
                .verifyPhoneRegistrationRequest(verifyPhoneRegistrationRequest)
                .build();
            Object result = apiInstance.authServiceVerifyPhoneRegistration(request);
            System.out.println(result);
        } catch (ApiException e) {
            System.err.println("Exception when calling AuthServiceApi#authServiceVerifyPhoneRegistration");
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
| authServiceVerifyPhoneRegistrationRequest | [**APIauthServiceVerifyPhoneRegistrationRequest**](AuthServiceApi.md#APIauthServiceVerifyPhoneRegistrationRequest)|-|-|

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

## authServiceVerifyPhoneRegistrationWithHttpInfo

> ApiResponse<Object> authServiceVerifyPhoneRegistrationWithHttpInfo(authServiceVerifyPhoneRegistrationRequest)



校验手机号注册验证码

### Example

```java
// Import classes:
import com.bass.bbs.ApiClient;
import com.bass.bbs.ApiException;
import com.bass.bbs.ApiResponse;
import com.bass.bbs.Configuration;
import com.bass.bbs.models.*;
import com.bass.bbs.api.AuthServiceApi;
import com.bass.bbs.api.AuthServiceApi.*;

public class Example {
    public static void main(String[] args) {
        ApiClient defaultClient = Configuration.getDefaultApiClient();
        defaultClient.setBasePath("http://localhost");

        AuthServiceApi apiInstance = new AuthServiceApi(defaultClient);
        VerifyPhoneRegistrationRequest verifyPhoneRegistrationRequest = new VerifyPhoneRegistrationRequest(); // VerifyPhoneRegistrationRequest | 
        try {
            APIauthServiceVerifyPhoneRegistrationRequest request = APIauthServiceVerifyPhoneRegistrationRequest.newBuilder()
                .verifyPhoneRegistrationRequest(verifyPhoneRegistrationRequest)
                .build();
            ApiResponse<Object> response = apiInstance.authServiceVerifyPhoneRegistrationWithHttpInfo(request);
            System.out.println("Status code: " + response.getStatusCode());
            System.out.println("Response headers: " + response.getHeaders());
            System.out.println("Response body: " + response.getData());
        } catch (ApiException e) {
            System.err.println("Exception when calling AuthServiceApi#authServiceVerifyPhoneRegistration");
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
| authServiceVerifyPhoneRegistrationRequest | [**APIauthServiceVerifyPhoneRegistrationRequest**](AuthServiceApi.md#APIauthServiceVerifyPhoneRegistrationRequest)|-|-|

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


<a id="APIauthServiceVerifyPhoneRegistrationRequest"></a>
## APIauthServiceVerifyPhoneRegistrationRequest
### Properties

|     Name      |    Type       | Description   |     Notes    |
| ------------- | ------------- | ------------- | -------------|
| **verifyPhoneRegistrationRequest** | [**VerifyPhoneRegistrationRequest**](VerifyPhoneRegistrationRequest.md) |  | |


