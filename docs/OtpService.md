# OtpService

All URIs are relative to *http://localhost*

| Method | HTTP request | Description |
|------------- | ------------- | -------------|
| [**beginEnableTotp**](OtpService.md#beginEnableTotp) | **POST** /v1/user/otp/totp/begin-enable |  |
| [**beginEnableTotpWithHttpInfo**](OtpService.md#beginEnableTotpWithHttpInfo) | **POST** /v1/user/otp/totp/begin-enable |  |
| [**confirmEnableTotp**](OtpService.md#confirmEnableTotp) | **POST** /v1/user/otp/totp/confirm-enable |  |
| [**confirmEnableTotpWithHttpInfo**](OtpService.md#confirmEnableTotpWithHttpInfo) | **POST** /v1/user/otp/totp/confirm-enable |  |
| [**disableTotp**](OtpService.md#disableTotp) | **POST** /v1/user/otp/totp/disable |  |
| [**disableTotpWithHttpInfo**](OtpService.md#disableTotpWithHttpInfo) | **POST** /v1/user/otp/totp/disable |  |
| [**getCurrentTotp**](OtpService.md#getCurrentTotp) | **POST** /v1/user/otp/totp/get-current |  |
| [**getCurrentTotpWithHttpInfo**](OtpService.md#getCurrentTotpWithHttpInfo) | **POST** /v1/user/otp/totp/get-current |  |
| [**sendEmailOtp**](OtpService.md#sendEmailOtp) | **POST** /v1/user/otp/email/send |  |
| [**sendEmailOtpWithHttpInfo**](OtpService.md#sendEmailOtpWithHttpInfo) | **POST** /v1/user/otp/email/send |  |
| [**sendPhoneOtp**](OtpService.md#sendPhoneOtp) | **POST** /v1/user/otp/phone/send |  |
| [**sendPhoneOtpWithHttpInfo**](OtpService.md#sendPhoneOtpWithHttpInfo) | **POST** /v1/user/otp/phone/send |  |



## beginEnableTotp

> BeginEnableTotpResp beginEnableTotp(beginEnableTotpRequest)



开始启用 TOTP。

### Example

```java
// Import classes:
import com.bass.bbs.ApiClient;
import com.bass.bbs.ApiException;
import com.bass.bbs.Configuration;
import com.bass.bbs.models.*;
import com.bass.bbs.api.OtpService;
import com.bass.bbs.api.OtpService.*;

public class Example {
    public static void main(String[] args) {
        ApiClient defaultClient = Configuration.getDefaultApiClient();
        defaultClient.setBasePath("http://localhost");

        OtpService apiInstance = new OtpService(defaultClient);
        Object body = null; // Object | 
        try {
            APIbeginEnableTotpRequest request = APIbeginEnableTotpRequest.newBuilder()
                .body(body)
                .build();
            BeginEnableTotpResp result = apiInstance.beginEnableTotp(request);
            System.out.println(result);
        } catch (ApiException e) {
            System.err.println("Exception when calling OtpService#beginEnableTotp");
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
| beginEnableTotpRequest | [**APIbeginEnableTotpRequest**](OtpService.md#APIbeginEnableTotpRequest)|-|-|

### Return type

[**BeginEnableTotpResp**](BeginEnableTotpResp.md)


### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

### HTTP response details
| Status code | Description | Response headers |
|-------------|-------------|------------------|
| **200** | OK |  -  |

## beginEnableTotpWithHttpInfo

> ApiResponse<BeginEnableTotpResp> beginEnableTotpWithHttpInfo(beginEnableTotpRequest)



开始启用 TOTP。

### Example

```java
// Import classes:
import com.bass.bbs.ApiClient;
import com.bass.bbs.ApiException;
import com.bass.bbs.ApiResponse;
import com.bass.bbs.Configuration;
import com.bass.bbs.models.*;
import com.bass.bbs.api.OtpService;
import com.bass.bbs.api.OtpService.*;

public class Example {
    public static void main(String[] args) {
        ApiClient defaultClient = Configuration.getDefaultApiClient();
        defaultClient.setBasePath("http://localhost");

        OtpService apiInstance = new OtpService(defaultClient);
        Object body = null; // Object | 
        try {
            APIbeginEnableTotpRequest request = APIbeginEnableTotpRequest.newBuilder()
                .body(body)
                .build();
            ApiResponse<BeginEnableTotpResp> response = apiInstance.beginEnableTotpWithHttpInfo(request);
            System.out.println("Status code: " + response.getStatusCode());
            System.out.println("Response headers: " + response.getHeaders());
            System.out.println("Response body: " + response.getData());
        } catch (ApiException e) {
            System.err.println("Exception when calling OtpService#beginEnableTotp");
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
| beginEnableTotpRequest | [**APIbeginEnableTotpRequest**](OtpService.md#APIbeginEnableTotpRequest)|-|-|

### Return type

ApiResponse<[**BeginEnableTotpResp**](BeginEnableTotpResp.md)>


### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

### HTTP response details
| Status code | Description | Response headers |
|-------------|-------------|------------------|
| **200** | OK |  -  |


<a id="APIbeginEnableTotpRequest"></a>
## APIbeginEnableTotpRequest
### Properties

|     Name      |    Type       | Description   |     Notes    |
| ------------- | ------------- | ------------- | -------------|
| **body** | **Object** |  | |



## confirmEnableTotp

> Object confirmEnableTotp(confirmEnableTotpRequest)



确认启用 TOTP。

### Example

```java
// Import classes:
import com.bass.bbs.ApiClient;
import com.bass.bbs.ApiException;
import com.bass.bbs.Configuration;
import com.bass.bbs.models.*;
import com.bass.bbs.api.OtpService;
import com.bass.bbs.api.OtpService.*;

public class Example {
    public static void main(String[] args) {
        ApiClient defaultClient = Configuration.getDefaultApiClient();
        defaultClient.setBasePath("http://localhost");

        OtpService apiInstance = new OtpService(defaultClient);
        ConfirmEnableTotpReq confirmEnableTotpReq = new ConfirmEnableTotpReq(); // ConfirmEnableTotpReq | 
        try {
            APIconfirmEnableTotpRequest request = APIconfirmEnableTotpRequest.newBuilder()
                .confirmEnableTotpReq(confirmEnableTotpReq)
                .build();
            Object result = apiInstance.confirmEnableTotp(request);
            System.out.println(result);
        } catch (ApiException e) {
            System.err.println("Exception when calling OtpService#confirmEnableTotp");
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
| confirmEnableTotpRequest | [**APIconfirmEnableTotpRequest**](OtpService.md#APIconfirmEnableTotpRequest)|-|-|

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

## confirmEnableTotpWithHttpInfo

> ApiResponse<Object> confirmEnableTotpWithHttpInfo(confirmEnableTotpRequest)



确认启用 TOTP。

### Example

```java
// Import classes:
import com.bass.bbs.ApiClient;
import com.bass.bbs.ApiException;
import com.bass.bbs.ApiResponse;
import com.bass.bbs.Configuration;
import com.bass.bbs.models.*;
import com.bass.bbs.api.OtpService;
import com.bass.bbs.api.OtpService.*;

public class Example {
    public static void main(String[] args) {
        ApiClient defaultClient = Configuration.getDefaultApiClient();
        defaultClient.setBasePath("http://localhost");

        OtpService apiInstance = new OtpService(defaultClient);
        ConfirmEnableTotpReq confirmEnableTotpReq = new ConfirmEnableTotpReq(); // ConfirmEnableTotpReq | 
        try {
            APIconfirmEnableTotpRequest request = APIconfirmEnableTotpRequest.newBuilder()
                .confirmEnableTotpReq(confirmEnableTotpReq)
                .build();
            ApiResponse<Object> response = apiInstance.confirmEnableTotpWithHttpInfo(request);
            System.out.println("Status code: " + response.getStatusCode());
            System.out.println("Response headers: " + response.getHeaders());
            System.out.println("Response body: " + response.getData());
        } catch (ApiException e) {
            System.err.println("Exception when calling OtpService#confirmEnableTotp");
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
| confirmEnableTotpRequest | [**APIconfirmEnableTotpRequest**](OtpService.md#APIconfirmEnableTotpRequest)|-|-|

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


<a id="APIconfirmEnableTotpRequest"></a>
## APIconfirmEnableTotpRequest
### Properties

|     Name      |    Type       | Description   |     Notes    |
| ------------- | ------------- | ------------- | -------------|
| **confirmEnableTotpReq** | [**ConfirmEnableTotpReq**](ConfirmEnableTotpReq.md) |  | |



## disableTotp

> Object disableTotp(disableTotpRequest)



关闭 TOTP。

### Example

```java
// Import classes:
import com.bass.bbs.ApiClient;
import com.bass.bbs.ApiException;
import com.bass.bbs.Configuration;
import com.bass.bbs.models.*;
import com.bass.bbs.api.OtpService;
import com.bass.bbs.api.OtpService.*;

public class Example {
    public static void main(String[] args) {
        ApiClient defaultClient = Configuration.getDefaultApiClient();
        defaultClient.setBasePath("http://localhost");

        OtpService apiInstance = new OtpService(defaultClient);
        DisableTotpReq disableTotpReq = new DisableTotpReq(); // DisableTotpReq | 
        try {
            APIdisableTotpRequest request = APIdisableTotpRequest.newBuilder()
                .disableTotpReq(disableTotpReq)
                .build();
            Object result = apiInstance.disableTotp(request);
            System.out.println(result);
        } catch (ApiException e) {
            System.err.println("Exception when calling OtpService#disableTotp");
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
| disableTotpRequest | [**APIdisableTotpRequest**](OtpService.md#APIdisableTotpRequest)|-|-|

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

## disableTotpWithHttpInfo

> ApiResponse<Object> disableTotpWithHttpInfo(disableTotpRequest)



关闭 TOTP。

### Example

```java
// Import classes:
import com.bass.bbs.ApiClient;
import com.bass.bbs.ApiException;
import com.bass.bbs.ApiResponse;
import com.bass.bbs.Configuration;
import com.bass.bbs.models.*;
import com.bass.bbs.api.OtpService;
import com.bass.bbs.api.OtpService.*;

public class Example {
    public static void main(String[] args) {
        ApiClient defaultClient = Configuration.getDefaultApiClient();
        defaultClient.setBasePath("http://localhost");

        OtpService apiInstance = new OtpService(defaultClient);
        DisableTotpReq disableTotpReq = new DisableTotpReq(); // DisableTotpReq | 
        try {
            APIdisableTotpRequest request = APIdisableTotpRequest.newBuilder()
                .disableTotpReq(disableTotpReq)
                .build();
            ApiResponse<Object> response = apiInstance.disableTotpWithHttpInfo(request);
            System.out.println("Status code: " + response.getStatusCode());
            System.out.println("Response headers: " + response.getHeaders());
            System.out.println("Response body: " + response.getData());
        } catch (ApiException e) {
            System.err.println("Exception when calling OtpService#disableTotp");
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
| disableTotpRequest | [**APIdisableTotpRequest**](OtpService.md#APIdisableTotpRequest)|-|-|

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


<a id="APIdisableTotpRequest"></a>
## APIdisableTotpRequest
### Properties

|     Name      |    Type       | Description   |     Notes    |
| ------------- | ------------- | ------------- | -------------|
| **disableTotpReq** | [**DisableTotpReq**](DisableTotpReq.md) |  | |



## getCurrentTotp

> GetCurrentTotpResp getCurrentTotp(getCurrentTotpRequest)



获取当前账号 TOTP 状态。

### Example

```java
// Import classes:
import com.bass.bbs.ApiClient;
import com.bass.bbs.ApiException;
import com.bass.bbs.Configuration;
import com.bass.bbs.models.*;
import com.bass.bbs.api.OtpService;
import com.bass.bbs.api.OtpService.*;

public class Example {
    public static void main(String[] args) {
        ApiClient defaultClient = Configuration.getDefaultApiClient();
        defaultClient.setBasePath("http://localhost");

        OtpService apiInstance = new OtpService(defaultClient);
        Object body = null; // Object | 
        try {
            APIgetCurrentTotpRequest request = APIgetCurrentTotpRequest.newBuilder()
                .body(body)
                .build();
            GetCurrentTotpResp result = apiInstance.getCurrentTotp(request);
            System.out.println(result);
        } catch (ApiException e) {
            System.err.println("Exception when calling OtpService#getCurrentTotp");
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
| getCurrentTotpRequest | [**APIgetCurrentTotpRequest**](OtpService.md#APIgetCurrentTotpRequest)|-|-|

### Return type

[**GetCurrentTotpResp**](GetCurrentTotpResp.md)


### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

### HTTP response details
| Status code | Description | Response headers |
|-------------|-------------|------------------|
| **200** | OK |  -  |

## getCurrentTotpWithHttpInfo

> ApiResponse<GetCurrentTotpResp> getCurrentTotpWithHttpInfo(getCurrentTotpRequest)



获取当前账号 TOTP 状态。

### Example

```java
// Import classes:
import com.bass.bbs.ApiClient;
import com.bass.bbs.ApiException;
import com.bass.bbs.ApiResponse;
import com.bass.bbs.Configuration;
import com.bass.bbs.models.*;
import com.bass.bbs.api.OtpService;
import com.bass.bbs.api.OtpService.*;

public class Example {
    public static void main(String[] args) {
        ApiClient defaultClient = Configuration.getDefaultApiClient();
        defaultClient.setBasePath("http://localhost");

        OtpService apiInstance = new OtpService(defaultClient);
        Object body = null; // Object | 
        try {
            APIgetCurrentTotpRequest request = APIgetCurrentTotpRequest.newBuilder()
                .body(body)
                .build();
            ApiResponse<GetCurrentTotpResp> response = apiInstance.getCurrentTotpWithHttpInfo(request);
            System.out.println("Status code: " + response.getStatusCode());
            System.out.println("Response headers: " + response.getHeaders());
            System.out.println("Response body: " + response.getData());
        } catch (ApiException e) {
            System.err.println("Exception when calling OtpService#getCurrentTotp");
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
| getCurrentTotpRequest | [**APIgetCurrentTotpRequest**](OtpService.md#APIgetCurrentTotpRequest)|-|-|

### Return type

ApiResponse<[**GetCurrentTotpResp**](GetCurrentTotpResp.md)>


### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

### HTTP response details
| Status code | Description | Response headers |
|-------------|-------------|------------------|
| **200** | OK |  -  |


<a id="APIgetCurrentTotpRequest"></a>
## APIgetCurrentTotpRequest
### Properties

|     Name      |    Type       | Description   |     Notes    |
| ------------- | ------------- | ------------- | -------------|
| **body** | **Object** |  | |



## sendEmailOtp

> SendEmailOtpResp sendEmailOtp(sendEmailOtpRequest)



发送邮箱 OTP。

### Example

```java
// Import classes:
import com.bass.bbs.ApiClient;
import com.bass.bbs.ApiException;
import com.bass.bbs.Configuration;
import com.bass.bbs.models.*;
import com.bass.bbs.api.OtpService;
import com.bass.bbs.api.OtpService.*;

public class Example {
    public static void main(String[] args) {
        ApiClient defaultClient = Configuration.getDefaultApiClient();
        defaultClient.setBasePath("http://localhost");

        OtpService apiInstance = new OtpService(defaultClient);
        SendEmailOtpReq sendEmailOtpReq = new SendEmailOtpReq(); // SendEmailOtpReq | 
        try {
            APIsendEmailOtpRequest request = APIsendEmailOtpRequest.newBuilder()
                .sendEmailOtpReq(sendEmailOtpReq)
                .build();
            SendEmailOtpResp result = apiInstance.sendEmailOtp(request);
            System.out.println(result);
        } catch (ApiException e) {
            System.err.println("Exception when calling OtpService#sendEmailOtp");
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
| sendEmailOtpRequest | [**APIsendEmailOtpRequest**](OtpService.md#APIsendEmailOtpRequest)|-|-|

### Return type

[**SendEmailOtpResp**](SendEmailOtpResp.md)


### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

### HTTP response details
| Status code | Description | Response headers |
|-------------|-------------|------------------|
| **200** | OK |  -  |

## sendEmailOtpWithHttpInfo

> ApiResponse<SendEmailOtpResp> sendEmailOtpWithHttpInfo(sendEmailOtpRequest)



发送邮箱 OTP。

### Example

```java
// Import classes:
import com.bass.bbs.ApiClient;
import com.bass.bbs.ApiException;
import com.bass.bbs.ApiResponse;
import com.bass.bbs.Configuration;
import com.bass.bbs.models.*;
import com.bass.bbs.api.OtpService;
import com.bass.bbs.api.OtpService.*;

public class Example {
    public static void main(String[] args) {
        ApiClient defaultClient = Configuration.getDefaultApiClient();
        defaultClient.setBasePath("http://localhost");

        OtpService apiInstance = new OtpService(defaultClient);
        SendEmailOtpReq sendEmailOtpReq = new SendEmailOtpReq(); // SendEmailOtpReq | 
        try {
            APIsendEmailOtpRequest request = APIsendEmailOtpRequest.newBuilder()
                .sendEmailOtpReq(sendEmailOtpReq)
                .build();
            ApiResponse<SendEmailOtpResp> response = apiInstance.sendEmailOtpWithHttpInfo(request);
            System.out.println("Status code: " + response.getStatusCode());
            System.out.println("Response headers: " + response.getHeaders());
            System.out.println("Response body: " + response.getData());
        } catch (ApiException e) {
            System.err.println("Exception when calling OtpService#sendEmailOtp");
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
| sendEmailOtpRequest | [**APIsendEmailOtpRequest**](OtpService.md#APIsendEmailOtpRequest)|-|-|

### Return type

ApiResponse<[**SendEmailOtpResp**](SendEmailOtpResp.md)>


### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

### HTTP response details
| Status code | Description | Response headers |
|-------------|-------------|------------------|
| **200** | OK |  -  |


<a id="APIsendEmailOtpRequest"></a>
## APIsendEmailOtpRequest
### Properties

|     Name      |    Type       | Description   |     Notes    |
| ------------- | ------------- | ------------- | -------------|
| **sendEmailOtpReq** | [**SendEmailOtpReq**](SendEmailOtpReq.md) |  | |



## sendPhoneOtp

> SendPhoneOtpResp sendPhoneOtp(sendPhoneOtpRequest)



发送手机 OTP。

### Example

```java
// Import classes:
import com.bass.bbs.ApiClient;
import com.bass.bbs.ApiException;
import com.bass.bbs.Configuration;
import com.bass.bbs.models.*;
import com.bass.bbs.api.OtpService;
import com.bass.bbs.api.OtpService.*;

public class Example {
    public static void main(String[] args) {
        ApiClient defaultClient = Configuration.getDefaultApiClient();
        defaultClient.setBasePath("http://localhost");

        OtpService apiInstance = new OtpService(defaultClient);
        SendPhoneOtpReq sendPhoneOtpReq = new SendPhoneOtpReq(); // SendPhoneOtpReq | 
        try {
            APIsendPhoneOtpRequest request = APIsendPhoneOtpRequest.newBuilder()
                .sendPhoneOtpReq(sendPhoneOtpReq)
                .build();
            SendPhoneOtpResp result = apiInstance.sendPhoneOtp(request);
            System.out.println(result);
        } catch (ApiException e) {
            System.err.println("Exception when calling OtpService#sendPhoneOtp");
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
| sendPhoneOtpRequest | [**APIsendPhoneOtpRequest**](OtpService.md#APIsendPhoneOtpRequest)|-|-|

### Return type

[**SendPhoneOtpResp**](SendPhoneOtpResp.md)


### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

### HTTP response details
| Status code | Description | Response headers |
|-------------|-------------|------------------|
| **200** | OK |  -  |

## sendPhoneOtpWithHttpInfo

> ApiResponse<SendPhoneOtpResp> sendPhoneOtpWithHttpInfo(sendPhoneOtpRequest)



发送手机 OTP。

### Example

```java
// Import classes:
import com.bass.bbs.ApiClient;
import com.bass.bbs.ApiException;
import com.bass.bbs.ApiResponse;
import com.bass.bbs.Configuration;
import com.bass.bbs.models.*;
import com.bass.bbs.api.OtpService;
import com.bass.bbs.api.OtpService.*;

public class Example {
    public static void main(String[] args) {
        ApiClient defaultClient = Configuration.getDefaultApiClient();
        defaultClient.setBasePath("http://localhost");

        OtpService apiInstance = new OtpService(defaultClient);
        SendPhoneOtpReq sendPhoneOtpReq = new SendPhoneOtpReq(); // SendPhoneOtpReq | 
        try {
            APIsendPhoneOtpRequest request = APIsendPhoneOtpRequest.newBuilder()
                .sendPhoneOtpReq(sendPhoneOtpReq)
                .build();
            ApiResponse<SendPhoneOtpResp> response = apiInstance.sendPhoneOtpWithHttpInfo(request);
            System.out.println("Status code: " + response.getStatusCode());
            System.out.println("Response headers: " + response.getHeaders());
            System.out.println("Response body: " + response.getData());
        } catch (ApiException e) {
            System.err.println("Exception when calling OtpService#sendPhoneOtp");
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
| sendPhoneOtpRequest | [**APIsendPhoneOtpRequest**](OtpService.md#APIsendPhoneOtpRequest)|-|-|

### Return type

ApiResponse<[**SendPhoneOtpResp**](SendPhoneOtpResp.md)>


### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

### HTTP response details
| Status code | Description | Response headers |
|-------------|-------------|------------------|
| **200** | OK |  -  |


<a id="APIsendPhoneOtpRequest"></a>
## APIsendPhoneOtpRequest
### Properties

|     Name      |    Type       | Description   |     Notes    |
| ------------- | ------------- | ------------- | -------------|
| **sendPhoneOtpReq** | [**SendPhoneOtpReq**](SendPhoneOtpReq.md) |  | |


