# TotpService

All URIs are relative to *http://localhost*

| Method | HTTP request | Description |
|------------- | ------------- | -------------|
| [**beginEnable**](TotpService.md#beginEnable) | **POST** /v1/user/totp/begin-enable |  |
| [**beginEnableWithHttpInfo**](TotpService.md#beginEnableWithHttpInfo) | **POST** /v1/user/totp/begin-enable |  |
| [**confirmEnable**](TotpService.md#confirmEnable) | **POST** /v1/user/totp/confirm-enable |  |
| [**confirmEnableWithHttpInfo**](TotpService.md#confirmEnableWithHttpInfo) | **POST** /v1/user/totp/confirm-enable |  |
| [**disable**](TotpService.md#disable) | **POST** /v1/user/totp/disable |  |
| [**disableWithHttpInfo**](TotpService.md#disableWithHttpInfo) | **POST** /v1/user/totp/disable |  |
| [**getCurrent**](TotpService.md#getCurrent) | **POST** /v1/user/totp/get-current |  |
| [**getCurrentWithHttpInfo**](TotpService.md#getCurrentWithHttpInfo) | **POST** /v1/user/totp/get-current |  |



## beginEnable

> BeginEnableTotpReply beginEnable(beginEnableRequest)



开始启用 TOTP。

### Example

```java
// Import classes:
import com.bass.bbs.ApiClient;
import com.bass.bbs.ApiException;
import com.bass.bbs.Configuration;
import com.bass.bbs.models.*;
import com.bass.bbs.api.TotpService;
import com.bass.bbs.api.TotpService.*;

public class Example {
    public static void main(String[] args) {
        ApiClient defaultClient = Configuration.getDefaultApiClient();
        defaultClient.setBasePath("http://localhost");

        TotpService apiInstance = new TotpService(defaultClient);
        Object body = null; // Object | 
        try {
            APIbeginEnableRequest request = APIbeginEnableRequest.newBuilder()
                .body(body)
                .build();
            BeginEnableTotpReply result = apiInstance.beginEnable(request);
            System.out.println(result);
        } catch (ApiException e) {
            System.err.println("Exception when calling TotpService#beginEnable");
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
| beginEnableRequest | [**APIbeginEnableRequest**](TotpService.md#APIbeginEnableRequest)|-|-|

### Return type

[**BeginEnableTotpReply**](BeginEnableTotpReply.md)


### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

### HTTP response details
| Status code | Description | Response headers |
|-------------|-------------|------------------|
| **200** | OK |  -  |

## beginEnableWithHttpInfo

> ApiResponse<BeginEnableTotpReply> beginEnableWithHttpInfo(beginEnableRequest)



开始启用 TOTP。

### Example

```java
// Import classes:
import com.bass.bbs.ApiClient;
import com.bass.bbs.ApiException;
import com.bass.bbs.ApiResponse;
import com.bass.bbs.Configuration;
import com.bass.bbs.models.*;
import com.bass.bbs.api.TotpService;
import com.bass.bbs.api.TotpService.*;

public class Example {
    public static void main(String[] args) {
        ApiClient defaultClient = Configuration.getDefaultApiClient();
        defaultClient.setBasePath("http://localhost");

        TotpService apiInstance = new TotpService(defaultClient);
        Object body = null; // Object | 
        try {
            APIbeginEnableRequest request = APIbeginEnableRequest.newBuilder()
                .body(body)
                .build();
            ApiResponse<BeginEnableTotpReply> response = apiInstance.beginEnableWithHttpInfo(request);
            System.out.println("Status code: " + response.getStatusCode());
            System.out.println("Response headers: " + response.getHeaders());
            System.out.println("Response body: " + response.getData());
        } catch (ApiException e) {
            System.err.println("Exception when calling TotpService#beginEnable");
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
| beginEnableRequest | [**APIbeginEnableRequest**](TotpService.md#APIbeginEnableRequest)|-|-|

### Return type

ApiResponse<[**BeginEnableTotpReply**](BeginEnableTotpReply.md)>


### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

### HTTP response details
| Status code | Description | Response headers |
|-------------|-------------|------------------|
| **200** | OK |  -  |


<a id="APIbeginEnableRequest"></a>
## APIbeginEnableRequest
### Properties

|     Name      |    Type       | Description   |     Notes    |
| ------------- | ------------- | ------------- | -------------|
| **body** | **Object** |  | |



## confirmEnable

> Object confirmEnable(confirmEnableRequest)



确认 TOTP 验证码并正式启用 TOTP。

### Example

```java
// Import classes:
import com.bass.bbs.ApiClient;
import com.bass.bbs.ApiException;
import com.bass.bbs.Configuration;
import com.bass.bbs.models.*;
import com.bass.bbs.api.TotpService;
import com.bass.bbs.api.TotpService.*;

public class Example {
    public static void main(String[] args) {
        ApiClient defaultClient = Configuration.getDefaultApiClient();
        defaultClient.setBasePath("http://localhost");

        TotpService apiInstance = new TotpService(defaultClient);
        ConfirmEnableTotpRequest confirmEnableTotpRequest = new ConfirmEnableTotpRequest(); // ConfirmEnableTotpRequest | 
        try {
            APIconfirmEnableRequest request = APIconfirmEnableRequest.newBuilder()
                .confirmEnableTotpRequest(confirmEnableTotpRequest)
                .build();
            Object result = apiInstance.confirmEnable(request);
            System.out.println(result);
        } catch (ApiException e) {
            System.err.println("Exception when calling TotpService#confirmEnable");
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
| confirmEnableRequest | [**APIconfirmEnableRequest**](TotpService.md#APIconfirmEnableRequest)|-|-|

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

## confirmEnableWithHttpInfo

> ApiResponse<Object> confirmEnableWithHttpInfo(confirmEnableRequest)



确认 TOTP 验证码并正式启用 TOTP。

### Example

```java
// Import classes:
import com.bass.bbs.ApiClient;
import com.bass.bbs.ApiException;
import com.bass.bbs.ApiResponse;
import com.bass.bbs.Configuration;
import com.bass.bbs.models.*;
import com.bass.bbs.api.TotpService;
import com.bass.bbs.api.TotpService.*;

public class Example {
    public static void main(String[] args) {
        ApiClient defaultClient = Configuration.getDefaultApiClient();
        defaultClient.setBasePath("http://localhost");

        TotpService apiInstance = new TotpService(defaultClient);
        ConfirmEnableTotpRequest confirmEnableTotpRequest = new ConfirmEnableTotpRequest(); // ConfirmEnableTotpRequest | 
        try {
            APIconfirmEnableRequest request = APIconfirmEnableRequest.newBuilder()
                .confirmEnableTotpRequest(confirmEnableTotpRequest)
                .build();
            ApiResponse<Object> response = apiInstance.confirmEnableWithHttpInfo(request);
            System.out.println("Status code: " + response.getStatusCode());
            System.out.println("Response headers: " + response.getHeaders());
            System.out.println("Response body: " + response.getData());
        } catch (ApiException e) {
            System.err.println("Exception when calling TotpService#confirmEnable");
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
| confirmEnableRequest | [**APIconfirmEnableRequest**](TotpService.md#APIconfirmEnableRequest)|-|-|

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


<a id="APIconfirmEnableRequest"></a>
## APIconfirmEnableRequest
### Properties

|     Name      |    Type       | Description   |     Notes    |
| ------------- | ------------- | ------------- | -------------|
| **confirmEnableTotpRequest** | [**ConfirmEnableTotpRequest**](ConfirmEnableTotpRequest.md) |  | |



## disable

> Object disable(disableRequest)



校验 TOTP 验证码并关闭 TOTP。

### Example

```java
// Import classes:
import com.bass.bbs.ApiClient;
import com.bass.bbs.ApiException;
import com.bass.bbs.Configuration;
import com.bass.bbs.models.*;
import com.bass.bbs.api.TotpService;
import com.bass.bbs.api.TotpService.*;

public class Example {
    public static void main(String[] args) {
        ApiClient defaultClient = Configuration.getDefaultApiClient();
        defaultClient.setBasePath("http://localhost");

        TotpService apiInstance = new TotpService(defaultClient);
        DisableTotpRequest disableTotpRequest = new DisableTotpRequest(); // DisableTotpRequest | 
        try {
            APIdisableRequest request = APIdisableRequest.newBuilder()
                .disableTotpRequest(disableTotpRequest)
                .build();
            Object result = apiInstance.disable(request);
            System.out.println(result);
        } catch (ApiException e) {
            System.err.println("Exception when calling TotpService#disable");
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
| disableRequest | [**APIdisableRequest**](TotpService.md#APIdisableRequest)|-|-|

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

## disableWithHttpInfo

> ApiResponse<Object> disableWithHttpInfo(disableRequest)



校验 TOTP 验证码并关闭 TOTP。

### Example

```java
// Import classes:
import com.bass.bbs.ApiClient;
import com.bass.bbs.ApiException;
import com.bass.bbs.ApiResponse;
import com.bass.bbs.Configuration;
import com.bass.bbs.models.*;
import com.bass.bbs.api.TotpService;
import com.bass.bbs.api.TotpService.*;

public class Example {
    public static void main(String[] args) {
        ApiClient defaultClient = Configuration.getDefaultApiClient();
        defaultClient.setBasePath("http://localhost");

        TotpService apiInstance = new TotpService(defaultClient);
        DisableTotpRequest disableTotpRequest = new DisableTotpRequest(); // DisableTotpRequest | 
        try {
            APIdisableRequest request = APIdisableRequest.newBuilder()
                .disableTotpRequest(disableTotpRequest)
                .build();
            ApiResponse<Object> response = apiInstance.disableWithHttpInfo(request);
            System.out.println("Status code: " + response.getStatusCode());
            System.out.println("Response headers: " + response.getHeaders());
            System.out.println("Response body: " + response.getData());
        } catch (ApiException e) {
            System.err.println("Exception when calling TotpService#disable");
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
| disableRequest | [**APIdisableRequest**](TotpService.md#APIdisableRequest)|-|-|

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


<a id="APIdisableRequest"></a>
## APIdisableRequest
### Properties

|     Name      |    Type       | Description   |     Notes    |
| ------------- | ------------- | ------------- | -------------|
| **disableTotpRequest** | [**DisableTotpRequest**](DisableTotpRequest.md) |  | |



## getCurrent

> GetCurrentTotpReply getCurrent(getCurrentRequest)



获取当前账号的 TOTP 状态。

### Example

```java
// Import classes:
import com.bass.bbs.ApiClient;
import com.bass.bbs.ApiException;
import com.bass.bbs.Configuration;
import com.bass.bbs.models.*;
import com.bass.bbs.api.TotpService;
import com.bass.bbs.api.TotpService.*;

public class Example {
    public static void main(String[] args) {
        ApiClient defaultClient = Configuration.getDefaultApiClient();
        defaultClient.setBasePath("http://localhost");

        TotpService apiInstance = new TotpService(defaultClient);
        Object body = null; // Object | 
        try {
            APIgetCurrentRequest request = APIgetCurrentRequest.newBuilder()
                .body(body)
                .build();
            GetCurrentTotpReply result = apiInstance.getCurrent(request);
            System.out.println(result);
        } catch (ApiException e) {
            System.err.println("Exception when calling TotpService#getCurrent");
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
| getCurrentRequest | [**APIgetCurrentRequest**](TotpService.md#APIgetCurrentRequest)|-|-|

### Return type

[**GetCurrentTotpReply**](GetCurrentTotpReply.md)


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

> ApiResponse<GetCurrentTotpReply> getCurrentWithHttpInfo(getCurrentRequest)



获取当前账号的 TOTP 状态。

### Example

```java
// Import classes:
import com.bass.bbs.ApiClient;
import com.bass.bbs.ApiException;
import com.bass.bbs.ApiResponse;
import com.bass.bbs.Configuration;
import com.bass.bbs.models.*;
import com.bass.bbs.api.TotpService;
import com.bass.bbs.api.TotpService.*;

public class Example {
    public static void main(String[] args) {
        ApiClient defaultClient = Configuration.getDefaultApiClient();
        defaultClient.setBasePath("http://localhost");

        TotpService apiInstance = new TotpService(defaultClient);
        Object body = null; // Object | 
        try {
            APIgetCurrentRequest request = APIgetCurrentRequest.newBuilder()
                .body(body)
                .build();
            ApiResponse<GetCurrentTotpReply> response = apiInstance.getCurrentWithHttpInfo(request);
            System.out.println("Status code: " + response.getStatusCode());
            System.out.println("Response headers: " + response.getHeaders());
            System.out.println("Response body: " + response.getData());
        } catch (ApiException e) {
            System.err.println("Exception when calling TotpService#getCurrent");
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
| getCurrentRequest | [**APIgetCurrentRequest**](TotpService.md#APIgetCurrentRequest)|-|-|

### Return type

ApiResponse<[**GetCurrentTotpReply**](GetCurrentTotpReply.md)>


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


