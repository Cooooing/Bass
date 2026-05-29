# TfaService

All URIs are relative to *http://localhost*

| Method | HTTP request | Description |
|------------- | ------------- | -------------|
| [**beginEnable**](TfaService.md#beginEnable) | **POST** /v1/user/tfa/begin-enable |  |
| [**beginEnableWithHttpInfo**](TfaService.md#beginEnableWithHttpInfo) | **POST** /v1/user/tfa/begin-enable |  |
| [**confirmEnable**](TfaService.md#confirmEnable) | **POST** /v1/user/tfa/confirm-enable |  |
| [**confirmEnableWithHttpInfo**](TfaService.md#confirmEnableWithHttpInfo) | **POST** /v1/user/tfa/confirm-enable |  |
| [**disable**](TfaService.md#disable) | **POST** /v1/user/tfa/disable |  |
| [**disableWithHttpInfo**](TfaService.md#disableWithHttpInfo) | **POST** /v1/user/tfa/disable |  |
| [**getCurrent**](TfaService.md#getCurrent) | **POST** /v1/user/tfa/get-current |  |
| [**getCurrentWithHttpInfo**](TfaService.md#getCurrentWithHttpInfo) | **POST** /v1/user/tfa/get-current |  |
| [**validate**](TfaService.md#validate) | **POST** /v1/user/tfa/validate |  |
| [**validateWithHttpInfo**](TfaService.md#validateWithHttpInfo) | **POST** /v1/user/tfa/validate |  |



## beginEnable

> BeginEnableTfaReply beginEnable(beginEnableRequest)



开始启用二步验证

### Example

```java
// Import classes:
import com.bass.bbs.ApiClient;
import com.bass.bbs.ApiException;
import com.bass.bbs.Configuration;
import com.bass.bbs.models.*;
import com.bass.bbs.api.TfaService;
import com.bass.bbs.api.TfaService.*;

public class Example {
    public static void main(String[] args) {
        ApiClient defaultClient = Configuration.getDefaultApiClient();
        defaultClient.setBasePath("http://localhost");

        TfaService apiInstance = new TfaService(defaultClient);
        Object body = null; // Object | 
        try {
            APIbeginEnableRequest request = APIbeginEnableRequest.newBuilder()
                .body(body)
                .build();
            BeginEnableTfaReply result = apiInstance.beginEnable(request);
            System.out.println(result);
        } catch (ApiException e) {
            System.err.println("Exception when calling TfaService#beginEnable");
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
| beginEnableRequest | [**APIbeginEnableRequest**](TfaService.md#APIbeginEnableRequest)|-|-|

### Return type

[**BeginEnableTfaReply**](BeginEnableTfaReply.md)


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

> ApiResponse<BeginEnableTfaReply> beginEnableWithHttpInfo(beginEnableRequest)



开始启用二步验证

### Example

```java
// Import classes:
import com.bass.bbs.ApiClient;
import com.bass.bbs.ApiException;
import com.bass.bbs.ApiResponse;
import com.bass.bbs.Configuration;
import com.bass.bbs.models.*;
import com.bass.bbs.api.TfaService;
import com.bass.bbs.api.TfaService.*;

public class Example {
    public static void main(String[] args) {
        ApiClient defaultClient = Configuration.getDefaultApiClient();
        defaultClient.setBasePath("http://localhost");

        TfaService apiInstance = new TfaService(defaultClient);
        Object body = null; // Object | 
        try {
            APIbeginEnableRequest request = APIbeginEnableRequest.newBuilder()
                .body(body)
                .build();
            ApiResponse<BeginEnableTfaReply> response = apiInstance.beginEnableWithHttpInfo(request);
            System.out.println("Status code: " + response.getStatusCode());
            System.out.println("Response headers: " + response.getHeaders());
            System.out.println("Response body: " + response.getData());
        } catch (ApiException e) {
            System.err.println("Exception when calling TfaService#beginEnable");
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
| beginEnableRequest | [**APIbeginEnableRequest**](TfaService.md#APIbeginEnableRequest)|-|-|

### Return type

ApiResponse<[**BeginEnableTfaReply**](BeginEnableTfaReply.md)>


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



确认二步验证码并正式启用二步验证

### Example

```java
// Import classes:
import com.bass.bbs.ApiClient;
import com.bass.bbs.ApiException;
import com.bass.bbs.Configuration;
import com.bass.bbs.models.*;
import com.bass.bbs.api.TfaService;
import com.bass.bbs.api.TfaService.*;

public class Example {
    public static void main(String[] args) {
        ApiClient defaultClient = Configuration.getDefaultApiClient();
        defaultClient.setBasePath("http://localhost");

        TfaService apiInstance = new TfaService(defaultClient);
        ConfirmEnableTfaRequest confirmEnableTfaRequest = new ConfirmEnableTfaRequest(); // ConfirmEnableTfaRequest | 
        try {
            APIconfirmEnableRequest request = APIconfirmEnableRequest.newBuilder()
                .confirmEnableTfaRequest(confirmEnableTfaRequest)
                .build();
            Object result = apiInstance.confirmEnable(request);
            System.out.println(result);
        } catch (ApiException e) {
            System.err.println("Exception when calling TfaService#confirmEnable");
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
| confirmEnableRequest | [**APIconfirmEnableRequest**](TfaService.md#APIconfirmEnableRequest)|-|-|

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



确认二步验证码并正式启用二步验证

### Example

```java
// Import classes:
import com.bass.bbs.ApiClient;
import com.bass.bbs.ApiException;
import com.bass.bbs.ApiResponse;
import com.bass.bbs.Configuration;
import com.bass.bbs.models.*;
import com.bass.bbs.api.TfaService;
import com.bass.bbs.api.TfaService.*;

public class Example {
    public static void main(String[] args) {
        ApiClient defaultClient = Configuration.getDefaultApiClient();
        defaultClient.setBasePath("http://localhost");

        TfaService apiInstance = new TfaService(defaultClient);
        ConfirmEnableTfaRequest confirmEnableTfaRequest = new ConfirmEnableTfaRequest(); // ConfirmEnableTfaRequest | 
        try {
            APIconfirmEnableRequest request = APIconfirmEnableRequest.newBuilder()
                .confirmEnableTfaRequest(confirmEnableTfaRequest)
                .build();
            ApiResponse<Object> response = apiInstance.confirmEnableWithHttpInfo(request);
            System.out.println("Status code: " + response.getStatusCode());
            System.out.println("Response headers: " + response.getHeaders());
            System.out.println("Response body: " + response.getData());
        } catch (ApiException e) {
            System.err.println("Exception when calling TfaService#confirmEnable");
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
| confirmEnableRequest | [**APIconfirmEnableRequest**](TfaService.md#APIconfirmEnableRequest)|-|-|

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
| **confirmEnableTfaRequest** | [**ConfirmEnableTfaRequest**](ConfirmEnableTfaRequest.md) |  | |



## disable

> Object disable(disableRequest)



校验二步验证码并关闭二步验证

### Example

```java
// Import classes:
import com.bass.bbs.ApiClient;
import com.bass.bbs.ApiException;
import com.bass.bbs.Configuration;
import com.bass.bbs.models.*;
import com.bass.bbs.api.TfaService;
import com.bass.bbs.api.TfaService.*;

public class Example {
    public static void main(String[] args) {
        ApiClient defaultClient = Configuration.getDefaultApiClient();
        defaultClient.setBasePath("http://localhost");

        TfaService apiInstance = new TfaService(defaultClient);
        DisableTfaRequest disableTfaRequest = new DisableTfaRequest(); // DisableTfaRequest | 
        try {
            APIdisableRequest request = APIdisableRequest.newBuilder()
                .disableTfaRequest(disableTfaRequest)
                .build();
            Object result = apiInstance.disable(request);
            System.out.println(result);
        } catch (ApiException e) {
            System.err.println("Exception when calling TfaService#disable");
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
| disableRequest | [**APIdisableRequest**](TfaService.md#APIdisableRequest)|-|-|

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



校验二步验证码并关闭二步验证

### Example

```java
// Import classes:
import com.bass.bbs.ApiClient;
import com.bass.bbs.ApiException;
import com.bass.bbs.ApiResponse;
import com.bass.bbs.Configuration;
import com.bass.bbs.models.*;
import com.bass.bbs.api.TfaService;
import com.bass.bbs.api.TfaService.*;

public class Example {
    public static void main(String[] args) {
        ApiClient defaultClient = Configuration.getDefaultApiClient();
        defaultClient.setBasePath("http://localhost");

        TfaService apiInstance = new TfaService(defaultClient);
        DisableTfaRequest disableTfaRequest = new DisableTfaRequest(); // DisableTfaRequest | 
        try {
            APIdisableRequest request = APIdisableRequest.newBuilder()
                .disableTfaRequest(disableTfaRequest)
                .build();
            ApiResponse<Object> response = apiInstance.disableWithHttpInfo(request);
            System.out.println("Status code: " + response.getStatusCode());
            System.out.println("Response headers: " + response.getHeaders());
            System.out.println("Response body: " + response.getData());
        } catch (ApiException e) {
            System.err.println("Exception when calling TfaService#disable");
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
| disableRequest | [**APIdisableRequest**](TfaService.md#APIdisableRequest)|-|-|

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
| **disableTfaRequest** | [**DisableTfaRequest**](DisableTfaRequest.md) |  | |



## getCurrent

> GetCurrentTfaReply getCurrent(getCurrentRequest)



获取当前登录账号的二步验证状态

### Example

```java
// Import classes:
import com.bass.bbs.ApiClient;
import com.bass.bbs.ApiException;
import com.bass.bbs.Configuration;
import com.bass.bbs.models.*;
import com.bass.bbs.api.TfaService;
import com.bass.bbs.api.TfaService.*;

public class Example {
    public static void main(String[] args) {
        ApiClient defaultClient = Configuration.getDefaultApiClient();
        defaultClient.setBasePath("http://localhost");

        TfaService apiInstance = new TfaService(defaultClient);
        Object body = null; // Object | 
        try {
            APIgetCurrentRequest request = APIgetCurrentRequest.newBuilder()
                .body(body)
                .build();
            GetCurrentTfaReply result = apiInstance.getCurrent(request);
            System.out.println(result);
        } catch (ApiException e) {
            System.err.println("Exception when calling TfaService#getCurrent");
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
| getCurrentRequest | [**APIgetCurrentRequest**](TfaService.md#APIgetCurrentRequest)|-|-|

### Return type

[**GetCurrentTfaReply**](GetCurrentTfaReply.md)


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

> ApiResponse<GetCurrentTfaReply> getCurrentWithHttpInfo(getCurrentRequest)



获取当前登录账号的二步验证状态

### Example

```java
// Import classes:
import com.bass.bbs.ApiClient;
import com.bass.bbs.ApiException;
import com.bass.bbs.ApiResponse;
import com.bass.bbs.Configuration;
import com.bass.bbs.models.*;
import com.bass.bbs.api.TfaService;
import com.bass.bbs.api.TfaService.*;

public class Example {
    public static void main(String[] args) {
        ApiClient defaultClient = Configuration.getDefaultApiClient();
        defaultClient.setBasePath("http://localhost");

        TfaService apiInstance = new TfaService(defaultClient);
        Object body = null; // Object | 
        try {
            APIgetCurrentRequest request = APIgetCurrentRequest.newBuilder()
                .body(body)
                .build();
            ApiResponse<GetCurrentTfaReply> response = apiInstance.getCurrentWithHttpInfo(request);
            System.out.println("Status code: " + response.getStatusCode());
            System.out.println("Response headers: " + response.getHeaders());
            System.out.println("Response body: " + response.getData());
        } catch (ApiException e) {
            System.err.println("Exception when calling TfaService#getCurrent");
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
| getCurrentRequest | [**APIgetCurrentRequest**](TfaService.md#APIgetCurrentRequest)|-|-|

### Return type

ApiResponse<[**GetCurrentTfaReply**](GetCurrentTfaReply.md)>


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



## validate

> ValidateTfaReply validate(validateRequest)



校验当前登录账号的二步验证码

### Example

```java
// Import classes:
import com.bass.bbs.ApiClient;
import com.bass.bbs.ApiException;
import com.bass.bbs.Configuration;
import com.bass.bbs.models.*;
import com.bass.bbs.api.TfaService;
import com.bass.bbs.api.TfaService.*;

public class Example {
    public static void main(String[] args) {
        ApiClient defaultClient = Configuration.getDefaultApiClient();
        defaultClient.setBasePath("http://localhost");

        TfaService apiInstance = new TfaService(defaultClient);
        ValidateTfaRequest validateTfaRequest = new ValidateTfaRequest(); // ValidateTfaRequest | 
        try {
            APIvalidateRequest request = APIvalidateRequest.newBuilder()
                .validateTfaRequest(validateTfaRequest)
                .build();
            ValidateTfaReply result = apiInstance.validate(request);
            System.out.println(result);
        } catch (ApiException e) {
            System.err.println("Exception when calling TfaService#validate");
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
| validateRequest | [**APIvalidateRequest**](TfaService.md#APIvalidateRequest)|-|-|

### Return type

[**ValidateTfaReply**](ValidateTfaReply.md)


### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

### HTTP response details
| Status code | Description | Response headers |
|-------------|-------------|------------------|
| **200** | OK |  -  |

## validateWithHttpInfo

> ApiResponse<ValidateTfaReply> validateWithHttpInfo(validateRequest)



校验当前登录账号的二步验证码

### Example

```java
// Import classes:
import com.bass.bbs.ApiClient;
import com.bass.bbs.ApiException;
import com.bass.bbs.ApiResponse;
import com.bass.bbs.Configuration;
import com.bass.bbs.models.*;
import com.bass.bbs.api.TfaService;
import com.bass.bbs.api.TfaService.*;

public class Example {
    public static void main(String[] args) {
        ApiClient defaultClient = Configuration.getDefaultApiClient();
        defaultClient.setBasePath("http://localhost");

        TfaService apiInstance = new TfaService(defaultClient);
        ValidateTfaRequest validateTfaRequest = new ValidateTfaRequest(); // ValidateTfaRequest | 
        try {
            APIvalidateRequest request = APIvalidateRequest.newBuilder()
                .validateTfaRequest(validateTfaRequest)
                .build();
            ApiResponse<ValidateTfaReply> response = apiInstance.validateWithHttpInfo(request);
            System.out.println("Status code: " + response.getStatusCode());
            System.out.println("Response headers: " + response.getHeaders());
            System.out.println("Response body: " + response.getData());
        } catch (ApiException e) {
            System.err.println("Exception when calling TfaService#validate");
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
| validateRequest | [**APIvalidateRequest**](TfaService.md#APIvalidateRequest)|-|-|

### Return type

ApiResponse<[**ValidateTfaReply**](ValidateTfaReply.md)>


### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

### HTTP response details
| Status code | Description | Response headers |
|-------------|-------------|------------------|
| **200** | OK |  -  |


<a id="APIvalidateRequest"></a>
## APIvalidateRequest
### Properties

|     Name      |    Type       | Description   |     Notes    |
| ------------- | ------------- | ------------- | -------------|
| **validateTfaRequest** | [**ValidateTfaRequest**](ValidateTfaRequest.md) |  | |


