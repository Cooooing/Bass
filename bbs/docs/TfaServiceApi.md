# TfaServiceApi

All URIs are relative to *http://localhost*

| Method | HTTP request | Description |
|------------- | ------------- | -------------|
| [**tfaServiceBeginEnable**](TfaServiceApi.md#tfaServiceBeginEnable) | **POST** /v1/user/tfa/begin-enable |  |
| [**tfaServiceBeginEnableWithHttpInfo**](TfaServiceApi.md#tfaServiceBeginEnableWithHttpInfo) | **POST** /v1/user/tfa/begin-enable |  |
| [**tfaServiceConfirmEnable**](TfaServiceApi.md#tfaServiceConfirmEnable) | **POST** /v1/user/tfa/confirm-enable |  |
| [**tfaServiceConfirmEnableWithHttpInfo**](TfaServiceApi.md#tfaServiceConfirmEnableWithHttpInfo) | **POST** /v1/user/tfa/confirm-enable |  |
| [**tfaServiceDisable**](TfaServiceApi.md#tfaServiceDisable) | **POST** /v1/user/tfa/disable |  |
| [**tfaServiceDisableWithHttpInfo**](TfaServiceApi.md#tfaServiceDisableWithHttpInfo) | **POST** /v1/user/tfa/disable |  |
| [**tfaServiceGetCurrent**](TfaServiceApi.md#tfaServiceGetCurrent) | **POST** /v1/user/tfa/get-current |  |
| [**tfaServiceGetCurrentWithHttpInfo**](TfaServiceApi.md#tfaServiceGetCurrentWithHttpInfo) | **POST** /v1/user/tfa/get-current |  |
| [**tfaServiceValidate**](TfaServiceApi.md#tfaServiceValidate) | **POST** /v1/user/tfa/validate |  |
| [**tfaServiceValidateWithHttpInfo**](TfaServiceApi.md#tfaServiceValidateWithHttpInfo) | **POST** /v1/user/tfa/validate |  |



## tfaServiceBeginEnable

> BeginEnableTfaReply tfaServiceBeginEnable(tfaServiceBeginEnableRequest)



开始启用二步验证

### Example

```java
// Import classes:
import com.bass.bbs.ApiClient;
import com.bass.bbs.ApiException;
import com.bass.bbs.Configuration;
import com.bass.bbs.models.*;
import com.bass.bbs.api.TfaServiceApi;
import com.bass.bbs.api.TfaServiceApi.*;

public class Example {
    public static void main(String[] args) {
        ApiClient defaultClient = Configuration.getDefaultApiClient();
        defaultClient.setBasePath("http://localhost");

        TfaServiceApi apiInstance = new TfaServiceApi(defaultClient);
        Object body = null; // Object | 
        try {
            APItfaServiceBeginEnableRequest request = APItfaServiceBeginEnableRequest.newBuilder()
                .body(body)
                .build();
            BeginEnableTfaReply result = apiInstance.tfaServiceBeginEnable(request);
            System.out.println(result);
        } catch (ApiException e) {
            System.err.println("Exception when calling TfaServiceApi#tfaServiceBeginEnable");
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
| tfaServiceBeginEnableRequest | [**APItfaServiceBeginEnableRequest**](TfaServiceApi.md#APItfaServiceBeginEnableRequest)|-|-|

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

## tfaServiceBeginEnableWithHttpInfo

> ApiResponse<BeginEnableTfaReply> tfaServiceBeginEnableWithHttpInfo(tfaServiceBeginEnableRequest)



开始启用二步验证

### Example

```java
// Import classes:
import com.bass.bbs.ApiClient;
import com.bass.bbs.ApiException;
import com.bass.bbs.ApiResponse;
import com.bass.bbs.Configuration;
import com.bass.bbs.models.*;
import com.bass.bbs.api.TfaServiceApi;
import com.bass.bbs.api.TfaServiceApi.*;

public class Example {
    public static void main(String[] args) {
        ApiClient defaultClient = Configuration.getDefaultApiClient();
        defaultClient.setBasePath("http://localhost");

        TfaServiceApi apiInstance = new TfaServiceApi(defaultClient);
        Object body = null; // Object | 
        try {
            APItfaServiceBeginEnableRequest request = APItfaServiceBeginEnableRequest.newBuilder()
                .body(body)
                .build();
            ApiResponse<BeginEnableTfaReply> response = apiInstance.tfaServiceBeginEnableWithHttpInfo(request);
            System.out.println("Status code: " + response.getStatusCode());
            System.out.println("Response headers: " + response.getHeaders());
            System.out.println("Response body: " + response.getData());
        } catch (ApiException e) {
            System.err.println("Exception when calling TfaServiceApi#tfaServiceBeginEnable");
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
| tfaServiceBeginEnableRequest | [**APItfaServiceBeginEnableRequest**](TfaServiceApi.md#APItfaServiceBeginEnableRequest)|-|-|

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


<a id="APItfaServiceBeginEnableRequest"></a>
## APItfaServiceBeginEnableRequest
### Properties

|     Name      |    Type       | Description   |     Notes    |
| ------------- | ------------- | ------------- | -------------|
| **body** | **Object** |  | |



## tfaServiceConfirmEnable

> Object tfaServiceConfirmEnable(tfaServiceConfirmEnableRequest)



确认二步验证码并正式启用二步验证

### Example

```java
// Import classes:
import com.bass.bbs.ApiClient;
import com.bass.bbs.ApiException;
import com.bass.bbs.Configuration;
import com.bass.bbs.models.*;
import com.bass.bbs.api.TfaServiceApi;
import com.bass.bbs.api.TfaServiceApi.*;

public class Example {
    public static void main(String[] args) {
        ApiClient defaultClient = Configuration.getDefaultApiClient();
        defaultClient.setBasePath("http://localhost");

        TfaServiceApi apiInstance = new TfaServiceApi(defaultClient);
        ConfirmEnableTfaRequest confirmEnableTfaRequest = new ConfirmEnableTfaRequest(); // ConfirmEnableTfaRequest | 
        try {
            APItfaServiceConfirmEnableRequest request = APItfaServiceConfirmEnableRequest.newBuilder()
                .confirmEnableTfaRequest(confirmEnableTfaRequest)
                .build();
            Object result = apiInstance.tfaServiceConfirmEnable(request);
            System.out.println(result);
        } catch (ApiException e) {
            System.err.println("Exception when calling TfaServiceApi#tfaServiceConfirmEnable");
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
| tfaServiceConfirmEnableRequest | [**APItfaServiceConfirmEnableRequest**](TfaServiceApi.md#APItfaServiceConfirmEnableRequest)|-|-|

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

## tfaServiceConfirmEnableWithHttpInfo

> ApiResponse<Object> tfaServiceConfirmEnableWithHttpInfo(tfaServiceConfirmEnableRequest)



确认二步验证码并正式启用二步验证

### Example

```java
// Import classes:
import com.bass.bbs.ApiClient;
import com.bass.bbs.ApiException;
import com.bass.bbs.ApiResponse;
import com.bass.bbs.Configuration;
import com.bass.bbs.models.*;
import com.bass.bbs.api.TfaServiceApi;
import com.bass.bbs.api.TfaServiceApi.*;

public class Example {
    public static void main(String[] args) {
        ApiClient defaultClient = Configuration.getDefaultApiClient();
        defaultClient.setBasePath("http://localhost");

        TfaServiceApi apiInstance = new TfaServiceApi(defaultClient);
        ConfirmEnableTfaRequest confirmEnableTfaRequest = new ConfirmEnableTfaRequest(); // ConfirmEnableTfaRequest | 
        try {
            APItfaServiceConfirmEnableRequest request = APItfaServiceConfirmEnableRequest.newBuilder()
                .confirmEnableTfaRequest(confirmEnableTfaRequest)
                .build();
            ApiResponse<Object> response = apiInstance.tfaServiceConfirmEnableWithHttpInfo(request);
            System.out.println("Status code: " + response.getStatusCode());
            System.out.println("Response headers: " + response.getHeaders());
            System.out.println("Response body: " + response.getData());
        } catch (ApiException e) {
            System.err.println("Exception when calling TfaServiceApi#tfaServiceConfirmEnable");
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
| tfaServiceConfirmEnableRequest | [**APItfaServiceConfirmEnableRequest**](TfaServiceApi.md#APItfaServiceConfirmEnableRequest)|-|-|

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


<a id="APItfaServiceConfirmEnableRequest"></a>
## APItfaServiceConfirmEnableRequest
### Properties

|     Name      |    Type       | Description   |     Notes    |
| ------------- | ------------- | ------------- | -------------|
| **confirmEnableTfaRequest** | [**ConfirmEnableTfaRequest**](ConfirmEnableTfaRequest.md) |  | |



## tfaServiceDisable

> Object tfaServiceDisable(tfaServiceDisableRequest)



校验二步验证码并关闭二步验证

### Example

```java
// Import classes:
import com.bass.bbs.ApiClient;
import com.bass.bbs.ApiException;
import com.bass.bbs.Configuration;
import com.bass.bbs.models.*;
import com.bass.bbs.api.TfaServiceApi;
import com.bass.bbs.api.TfaServiceApi.*;

public class Example {
    public static void main(String[] args) {
        ApiClient defaultClient = Configuration.getDefaultApiClient();
        defaultClient.setBasePath("http://localhost");

        TfaServiceApi apiInstance = new TfaServiceApi(defaultClient);
        DisableTfaRequest disableTfaRequest = new DisableTfaRequest(); // DisableTfaRequest | 
        try {
            APItfaServiceDisableRequest request = APItfaServiceDisableRequest.newBuilder()
                .disableTfaRequest(disableTfaRequest)
                .build();
            Object result = apiInstance.tfaServiceDisable(request);
            System.out.println(result);
        } catch (ApiException e) {
            System.err.println("Exception when calling TfaServiceApi#tfaServiceDisable");
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
| tfaServiceDisableRequest | [**APItfaServiceDisableRequest**](TfaServiceApi.md#APItfaServiceDisableRequest)|-|-|

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

## tfaServiceDisableWithHttpInfo

> ApiResponse<Object> tfaServiceDisableWithHttpInfo(tfaServiceDisableRequest)



校验二步验证码并关闭二步验证

### Example

```java
// Import classes:
import com.bass.bbs.ApiClient;
import com.bass.bbs.ApiException;
import com.bass.bbs.ApiResponse;
import com.bass.bbs.Configuration;
import com.bass.bbs.models.*;
import com.bass.bbs.api.TfaServiceApi;
import com.bass.bbs.api.TfaServiceApi.*;

public class Example {
    public static void main(String[] args) {
        ApiClient defaultClient = Configuration.getDefaultApiClient();
        defaultClient.setBasePath("http://localhost");

        TfaServiceApi apiInstance = new TfaServiceApi(defaultClient);
        DisableTfaRequest disableTfaRequest = new DisableTfaRequest(); // DisableTfaRequest | 
        try {
            APItfaServiceDisableRequest request = APItfaServiceDisableRequest.newBuilder()
                .disableTfaRequest(disableTfaRequest)
                .build();
            ApiResponse<Object> response = apiInstance.tfaServiceDisableWithHttpInfo(request);
            System.out.println("Status code: " + response.getStatusCode());
            System.out.println("Response headers: " + response.getHeaders());
            System.out.println("Response body: " + response.getData());
        } catch (ApiException e) {
            System.err.println("Exception when calling TfaServiceApi#tfaServiceDisable");
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
| tfaServiceDisableRequest | [**APItfaServiceDisableRequest**](TfaServiceApi.md#APItfaServiceDisableRequest)|-|-|

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


<a id="APItfaServiceDisableRequest"></a>
## APItfaServiceDisableRequest
### Properties

|     Name      |    Type       | Description   |     Notes    |
| ------------- | ------------- | ------------- | -------------|
| **disableTfaRequest** | [**DisableTfaRequest**](DisableTfaRequest.md) |  | |



## tfaServiceGetCurrent

> GetCurrentTfaReply tfaServiceGetCurrent(tfaServiceGetCurrentRequest)



获取当前登录账号的二步验证状态

### Example

```java
// Import classes:
import com.bass.bbs.ApiClient;
import com.bass.bbs.ApiException;
import com.bass.bbs.Configuration;
import com.bass.bbs.models.*;
import com.bass.bbs.api.TfaServiceApi;
import com.bass.bbs.api.TfaServiceApi.*;

public class Example {
    public static void main(String[] args) {
        ApiClient defaultClient = Configuration.getDefaultApiClient();
        defaultClient.setBasePath("http://localhost");

        TfaServiceApi apiInstance = new TfaServiceApi(defaultClient);
        Object body = null; // Object | 
        try {
            APItfaServiceGetCurrentRequest request = APItfaServiceGetCurrentRequest.newBuilder()
                .body(body)
                .build();
            GetCurrentTfaReply result = apiInstance.tfaServiceGetCurrent(request);
            System.out.println(result);
        } catch (ApiException e) {
            System.err.println("Exception when calling TfaServiceApi#tfaServiceGetCurrent");
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
| tfaServiceGetCurrentRequest | [**APItfaServiceGetCurrentRequest**](TfaServiceApi.md#APItfaServiceGetCurrentRequest)|-|-|

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

## tfaServiceGetCurrentWithHttpInfo

> ApiResponse<GetCurrentTfaReply> tfaServiceGetCurrentWithHttpInfo(tfaServiceGetCurrentRequest)



获取当前登录账号的二步验证状态

### Example

```java
// Import classes:
import com.bass.bbs.ApiClient;
import com.bass.bbs.ApiException;
import com.bass.bbs.ApiResponse;
import com.bass.bbs.Configuration;
import com.bass.bbs.models.*;
import com.bass.bbs.api.TfaServiceApi;
import com.bass.bbs.api.TfaServiceApi.*;

public class Example {
    public static void main(String[] args) {
        ApiClient defaultClient = Configuration.getDefaultApiClient();
        defaultClient.setBasePath("http://localhost");

        TfaServiceApi apiInstance = new TfaServiceApi(defaultClient);
        Object body = null; // Object | 
        try {
            APItfaServiceGetCurrentRequest request = APItfaServiceGetCurrentRequest.newBuilder()
                .body(body)
                .build();
            ApiResponse<GetCurrentTfaReply> response = apiInstance.tfaServiceGetCurrentWithHttpInfo(request);
            System.out.println("Status code: " + response.getStatusCode());
            System.out.println("Response headers: " + response.getHeaders());
            System.out.println("Response body: " + response.getData());
        } catch (ApiException e) {
            System.err.println("Exception when calling TfaServiceApi#tfaServiceGetCurrent");
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
| tfaServiceGetCurrentRequest | [**APItfaServiceGetCurrentRequest**](TfaServiceApi.md#APItfaServiceGetCurrentRequest)|-|-|

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


<a id="APItfaServiceGetCurrentRequest"></a>
## APItfaServiceGetCurrentRequest
### Properties

|     Name      |    Type       | Description   |     Notes    |
| ------------- | ------------- | ------------- | -------------|
| **body** | **Object** |  | |



## tfaServiceValidate

> ValidateTfaReply tfaServiceValidate(tfaServiceValidateRequest)



校验当前登录账号的二步验证码

### Example

```java
// Import classes:
import com.bass.bbs.ApiClient;
import com.bass.bbs.ApiException;
import com.bass.bbs.Configuration;
import com.bass.bbs.models.*;
import com.bass.bbs.api.TfaServiceApi;
import com.bass.bbs.api.TfaServiceApi.*;

public class Example {
    public static void main(String[] args) {
        ApiClient defaultClient = Configuration.getDefaultApiClient();
        defaultClient.setBasePath("http://localhost");

        TfaServiceApi apiInstance = new TfaServiceApi(defaultClient);
        ValidateTfaRequest validateTfaRequest = new ValidateTfaRequest(); // ValidateTfaRequest | 
        try {
            APItfaServiceValidateRequest request = APItfaServiceValidateRequest.newBuilder()
                .validateTfaRequest(validateTfaRequest)
                .build();
            ValidateTfaReply result = apiInstance.tfaServiceValidate(request);
            System.out.println(result);
        } catch (ApiException e) {
            System.err.println("Exception when calling TfaServiceApi#tfaServiceValidate");
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
| tfaServiceValidateRequest | [**APItfaServiceValidateRequest**](TfaServiceApi.md#APItfaServiceValidateRequest)|-|-|

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

## tfaServiceValidateWithHttpInfo

> ApiResponse<ValidateTfaReply> tfaServiceValidateWithHttpInfo(tfaServiceValidateRequest)



校验当前登录账号的二步验证码

### Example

```java
// Import classes:
import com.bass.bbs.ApiClient;
import com.bass.bbs.ApiException;
import com.bass.bbs.ApiResponse;
import com.bass.bbs.Configuration;
import com.bass.bbs.models.*;
import com.bass.bbs.api.TfaServiceApi;
import com.bass.bbs.api.TfaServiceApi.*;

public class Example {
    public static void main(String[] args) {
        ApiClient defaultClient = Configuration.getDefaultApiClient();
        defaultClient.setBasePath("http://localhost");

        TfaServiceApi apiInstance = new TfaServiceApi(defaultClient);
        ValidateTfaRequest validateTfaRequest = new ValidateTfaRequest(); // ValidateTfaRequest | 
        try {
            APItfaServiceValidateRequest request = APItfaServiceValidateRequest.newBuilder()
                .validateTfaRequest(validateTfaRequest)
                .build();
            ApiResponse<ValidateTfaReply> response = apiInstance.tfaServiceValidateWithHttpInfo(request);
            System.out.println("Status code: " + response.getStatusCode());
            System.out.println("Response headers: " + response.getHeaders());
            System.out.println("Response body: " + response.getData());
        } catch (ApiException e) {
            System.err.println("Exception when calling TfaServiceApi#tfaServiceValidate");
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
| tfaServiceValidateRequest | [**APItfaServiceValidateRequest**](TfaServiceApi.md#APItfaServiceValidateRequest)|-|-|

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


<a id="APItfaServiceValidateRequest"></a>
## APItfaServiceValidateRequest
### Properties

|     Name      |    Type       | Description   |     Notes    |
| ------------- | ------------- | ------------- | -------------|
| **validateTfaRequest** | [**ValidateTfaRequest**](ValidateTfaRequest.md) |  | |


