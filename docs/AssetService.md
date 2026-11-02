# AssetService

All URIs are relative to *http://localhost*

| Method | HTTP request | Description |
|------------- | ------------- | -------------|
| [**completeBadgeFontUpload**](AssetService.md#completeBadgeFontUpload) | **POST** /v1/asset/badge-font/complete |  |
| [**completeBadgeFontUploadWithHttpInfo**](AssetService.md#completeBadgeFontUploadWithHttpInfo) | **POST** /v1/asset/badge-font/complete |  |
| [**completeBadgeImageUpload**](AssetService.md#completeBadgeImageUpload) | **POST** /v1/asset/badge-image/complete |  |
| [**completeBadgeImageUploadWithHttpInfo**](AssetService.md#completeBadgeImageUploadWithHttpInfo) | **POST** /v1/asset/badge-image/complete |  |
| [**completeMarkdownImageUpload**](AssetService.md#completeMarkdownImageUpload) | **POST** /v1/asset/markdown-image/complete |  |
| [**completeMarkdownImageUploadWithHttpInfo**](AssetService.md#completeMarkdownImageUploadWithHttpInfo) | **POST** /v1/asset/markdown-image/complete |  |
| [**prepareBadgeFontUpload**](AssetService.md#prepareBadgeFontUpload) | **POST** /v1/asset/badge-font/prepare |  |
| [**prepareBadgeFontUploadWithHttpInfo**](AssetService.md#prepareBadgeFontUploadWithHttpInfo) | **POST** /v1/asset/badge-font/prepare |  |
| [**prepareBadgeImageUpload**](AssetService.md#prepareBadgeImageUpload) | **POST** /v1/asset/badge-image/prepare |  |
| [**prepareBadgeImageUploadWithHttpInfo**](AssetService.md#prepareBadgeImageUploadWithHttpInfo) | **POST** /v1/asset/badge-image/prepare |  |
| [**prepareMarkdownImageUpload**](AssetService.md#prepareMarkdownImageUpload) | **POST** /v1/asset/markdown-image/prepare |  |
| [**prepareMarkdownImageUploadWithHttpInfo**](AssetService.md#prepareMarkdownImageUploadWithHttpInfo) | **POST** /v1/asset/markdown-image/prepare |  |



## completeBadgeFontUpload

> CompleteBadgeFontUploadResp completeBadgeFontUpload(completeBadgeFontUploadRequest)



查询已由对象存储回调确认的徽章 Maker 自定义字体。

### Example

```java
// Import classes:
import com.bass.bbs.ApiClient;
import com.bass.bbs.ApiException;
import com.bass.bbs.Configuration;
import com.bass.bbs.models.*;
import com.bass.bbs.api.AssetService;
import com.bass.bbs.api.AssetService.*;

public class Example {
    public static void main(String[] args) {
        ApiClient defaultClient = Configuration.getDefaultApiClient();
        defaultClient.setBasePath("http://localhost");

        AssetService apiInstance = new AssetService(defaultClient);
        CompleteBadgeFontUploadReq completeBadgeFontUploadReq = new CompleteBadgeFontUploadReq(); // CompleteBadgeFontUploadReq | 
        try {
            APIcompleteBadgeFontUploadRequest request = APIcompleteBadgeFontUploadRequest.newBuilder()
                .completeBadgeFontUploadReq(completeBadgeFontUploadReq)
                .build();
            CompleteBadgeFontUploadResp result = apiInstance.completeBadgeFontUpload(request);
            System.out.println(result);
        } catch (ApiException e) {
            System.err.println("Exception when calling AssetService#completeBadgeFontUpload");
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
| completeBadgeFontUploadRequest | [**APIcompleteBadgeFontUploadRequest**](AssetService.md#APIcompleteBadgeFontUploadRequest)|-|-|

### Return type

[**CompleteBadgeFontUploadResp**](CompleteBadgeFontUploadResp.md)


### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

### HTTP response details
| Status code | Description | Response headers |
|-------------|-------------|------------------|
| **200** | OK |  -  |

## completeBadgeFontUploadWithHttpInfo

> ApiResponse<CompleteBadgeFontUploadResp> completeBadgeFontUploadWithHttpInfo(completeBadgeFontUploadRequest)



查询已由对象存储回调确认的徽章 Maker 自定义字体。

### Example

```java
// Import classes:
import com.bass.bbs.ApiClient;
import com.bass.bbs.ApiException;
import com.bass.bbs.ApiResponse;
import com.bass.bbs.Configuration;
import com.bass.bbs.models.*;
import com.bass.bbs.api.AssetService;
import com.bass.bbs.api.AssetService.*;

public class Example {
    public static void main(String[] args) {
        ApiClient defaultClient = Configuration.getDefaultApiClient();
        defaultClient.setBasePath("http://localhost");

        AssetService apiInstance = new AssetService(defaultClient);
        CompleteBadgeFontUploadReq completeBadgeFontUploadReq = new CompleteBadgeFontUploadReq(); // CompleteBadgeFontUploadReq | 
        try {
            APIcompleteBadgeFontUploadRequest request = APIcompleteBadgeFontUploadRequest.newBuilder()
                .completeBadgeFontUploadReq(completeBadgeFontUploadReq)
                .build();
            ApiResponse<CompleteBadgeFontUploadResp> response = apiInstance.completeBadgeFontUploadWithHttpInfo(request);
            System.out.println("Status code: " + response.getStatusCode());
            System.out.println("Response headers: " + response.getHeaders());
            System.out.println("Response body: " + response.getData());
        } catch (ApiException e) {
            System.err.println("Exception when calling AssetService#completeBadgeFontUpload");
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
| completeBadgeFontUploadRequest | [**APIcompleteBadgeFontUploadRequest**](AssetService.md#APIcompleteBadgeFontUploadRequest)|-|-|

### Return type

ApiResponse<[**CompleteBadgeFontUploadResp**](CompleteBadgeFontUploadResp.md)>


### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

### HTTP response details
| Status code | Description | Response headers |
|-------------|-------------|------------------|
| **200** | OK |  -  |


<a id="APIcompleteBadgeFontUploadRequest"></a>
## APIcompleteBadgeFontUploadRequest
### Properties

|     Name      |    Type       | Description   |     Notes    |
| ------------- | ------------- | ------------- | -------------|
| **completeBadgeFontUploadReq** | [**CompleteBadgeFontUploadReq**](CompleteBadgeFontUploadReq.md) |  | |



## completeBadgeImageUpload

> CompleteBadgeImageUploadResp completeBadgeImageUpload(completeBadgeImageUploadRequest)



查询已由对象存储回调确认的徽章 Maker 图片。

### Example

```java
// Import classes:
import com.bass.bbs.ApiClient;
import com.bass.bbs.ApiException;
import com.bass.bbs.Configuration;
import com.bass.bbs.models.*;
import com.bass.bbs.api.AssetService;
import com.bass.bbs.api.AssetService.*;

public class Example {
    public static void main(String[] args) {
        ApiClient defaultClient = Configuration.getDefaultApiClient();
        defaultClient.setBasePath("http://localhost");

        AssetService apiInstance = new AssetService(defaultClient);
        CompleteBadgeImageUploadReq completeBadgeImageUploadReq = new CompleteBadgeImageUploadReq(); // CompleteBadgeImageUploadReq | 
        try {
            APIcompleteBadgeImageUploadRequest request = APIcompleteBadgeImageUploadRequest.newBuilder()
                .completeBadgeImageUploadReq(completeBadgeImageUploadReq)
                .build();
            CompleteBadgeImageUploadResp result = apiInstance.completeBadgeImageUpload(request);
            System.out.println(result);
        } catch (ApiException e) {
            System.err.println("Exception when calling AssetService#completeBadgeImageUpload");
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
| completeBadgeImageUploadRequest | [**APIcompleteBadgeImageUploadRequest**](AssetService.md#APIcompleteBadgeImageUploadRequest)|-|-|

### Return type

[**CompleteBadgeImageUploadResp**](CompleteBadgeImageUploadResp.md)


### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

### HTTP response details
| Status code | Description | Response headers |
|-------------|-------------|------------------|
| **200** | OK |  -  |

## completeBadgeImageUploadWithHttpInfo

> ApiResponse<CompleteBadgeImageUploadResp> completeBadgeImageUploadWithHttpInfo(completeBadgeImageUploadRequest)



查询已由对象存储回调确认的徽章 Maker 图片。

### Example

```java
// Import classes:
import com.bass.bbs.ApiClient;
import com.bass.bbs.ApiException;
import com.bass.bbs.ApiResponse;
import com.bass.bbs.Configuration;
import com.bass.bbs.models.*;
import com.bass.bbs.api.AssetService;
import com.bass.bbs.api.AssetService.*;

public class Example {
    public static void main(String[] args) {
        ApiClient defaultClient = Configuration.getDefaultApiClient();
        defaultClient.setBasePath("http://localhost");

        AssetService apiInstance = new AssetService(defaultClient);
        CompleteBadgeImageUploadReq completeBadgeImageUploadReq = new CompleteBadgeImageUploadReq(); // CompleteBadgeImageUploadReq | 
        try {
            APIcompleteBadgeImageUploadRequest request = APIcompleteBadgeImageUploadRequest.newBuilder()
                .completeBadgeImageUploadReq(completeBadgeImageUploadReq)
                .build();
            ApiResponse<CompleteBadgeImageUploadResp> response = apiInstance.completeBadgeImageUploadWithHttpInfo(request);
            System.out.println("Status code: " + response.getStatusCode());
            System.out.println("Response headers: " + response.getHeaders());
            System.out.println("Response body: " + response.getData());
        } catch (ApiException e) {
            System.err.println("Exception when calling AssetService#completeBadgeImageUpload");
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
| completeBadgeImageUploadRequest | [**APIcompleteBadgeImageUploadRequest**](AssetService.md#APIcompleteBadgeImageUploadRequest)|-|-|

### Return type

ApiResponse<[**CompleteBadgeImageUploadResp**](CompleteBadgeImageUploadResp.md)>


### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

### HTTP response details
| Status code | Description | Response headers |
|-------------|-------------|------------------|
| **200** | OK |  -  |


<a id="APIcompleteBadgeImageUploadRequest"></a>
## APIcompleteBadgeImageUploadRequest
### Properties

|     Name      |    Type       | Description   |     Notes    |
| ------------- | ------------- | ------------- | -------------|
| **completeBadgeImageUploadReq** | [**CompleteBadgeImageUploadReq**](CompleteBadgeImageUploadReq.md) |  | |



## completeMarkdownImageUpload

> CompleteMarkdownImageUploadResp completeMarkdownImageUpload(completeMarkdownImageUploadRequest)



查询已由对象存储回调确认的 Markdown 图片。

### Example

```java
// Import classes:
import com.bass.bbs.ApiClient;
import com.bass.bbs.ApiException;
import com.bass.bbs.Configuration;
import com.bass.bbs.models.*;
import com.bass.bbs.api.AssetService;
import com.bass.bbs.api.AssetService.*;

public class Example {
    public static void main(String[] args) {
        ApiClient defaultClient = Configuration.getDefaultApiClient();
        defaultClient.setBasePath("http://localhost");

        AssetService apiInstance = new AssetService(defaultClient);
        CompleteMarkdownImageUploadReq completeMarkdownImageUploadReq = new CompleteMarkdownImageUploadReq(); // CompleteMarkdownImageUploadReq | 
        try {
            APIcompleteMarkdownImageUploadRequest request = APIcompleteMarkdownImageUploadRequest.newBuilder()
                .completeMarkdownImageUploadReq(completeMarkdownImageUploadReq)
                .build();
            CompleteMarkdownImageUploadResp result = apiInstance.completeMarkdownImageUpload(request);
            System.out.println(result);
        } catch (ApiException e) {
            System.err.println("Exception when calling AssetService#completeMarkdownImageUpload");
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
| completeMarkdownImageUploadRequest | [**APIcompleteMarkdownImageUploadRequest**](AssetService.md#APIcompleteMarkdownImageUploadRequest)|-|-|

### Return type

[**CompleteMarkdownImageUploadResp**](CompleteMarkdownImageUploadResp.md)


### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

### HTTP response details
| Status code | Description | Response headers |
|-------------|-------------|------------------|
| **200** | OK |  -  |

## completeMarkdownImageUploadWithHttpInfo

> ApiResponse<CompleteMarkdownImageUploadResp> completeMarkdownImageUploadWithHttpInfo(completeMarkdownImageUploadRequest)



查询已由对象存储回调确认的 Markdown 图片。

### Example

```java
// Import classes:
import com.bass.bbs.ApiClient;
import com.bass.bbs.ApiException;
import com.bass.bbs.ApiResponse;
import com.bass.bbs.Configuration;
import com.bass.bbs.models.*;
import com.bass.bbs.api.AssetService;
import com.bass.bbs.api.AssetService.*;

public class Example {
    public static void main(String[] args) {
        ApiClient defaultClient = Configuration.getDefaultApiClient();
        defaultClient.setBasePath("http://localhost");

        AssetService apiInstance = new AssetService(defaultClient);
        CompleteMarkdownImageUploadReq completeMarkdownImageUploadReq = new CompleteMarkdownImageUploadReq(); // CompleteMarkdownImageUploadReq | 
        try {
            APIcompleteMarkdownImageUploadRequest request = APIcompleteMarkdownImageUploadRequest.newBuilder()
                .completeMarkdownImageUploadReq(completeMarkdownImageUploadReq)
                .build();
            ApiResponse<CompleteMarkdownImageUploadResp> response = apiInstance.completeMarkdownImageUploadWithHttpInfo(request);
            System.out.println("Status code: " + response.getStatusCode());
            System.out.println("Response headers: " + response.getHeaders());
            System.out.println("Response body: " + response.getData());
        } catch (ApiException e) {
            System.err.println("Exception when calling AssetService#completeMarkdownImageUpload");
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
| completeMarkdownImageUploadRequest | [**APIcompleteMarkdownImageUploadRequest**](AssetService.md#APIcompleteMarkdownImageUploadRequest)|-|-|

### Return type

ApiResponse<[**CompleteMarkdownImageUploadResp**](CompleteMarkdownImageUploadResp.md)>


### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

### HTTP response details
| Status code | Description | Response headers |
|-------------|-------------|------------------|
| **200** | OK |  -  |


<a id="APIcompleteMarkdownImageUploadRequest"></a>
## APIcompleteMarkdownImageUploadRequest
### Properties

|     Name      |    Type       | Description   |     Notes    |
| ------------- | ------------- | ------------- | -------------|
| **completeMarkdownImageUploadReq** | [**CompleteMarkdownImageUploadReq**](CompleteMarkdownImageUploadReq.md) |  | |



## prepareBadgeFontUpload

> PrepareBadgeFontUploadResp prepareBadgeFontUpload(prepareBadgeFontUploadRequest)



申请徽章 Maker 自定义字体的内容寻址直传凭证。

### Example

```java
// Import classes:
import com.bass.bbs.ApiClient;
import com.bass.bbs.ApiException;
import com.bass.bbs.Configuration;
import com.bass.bbs.models.*;
import com.bass.bbs.api.AssetService;
import com.bass.bbs.api.AssetService.*;

public class Example {
    public static void main(String[] args) {
        ApiClient defaultClient = Configuration.getDefaultApiClient();
        defaultClient.setBasePath("http://localhost");

        AssetService apiInstance = new AssetService(defaultClient);
        PrepareBadgeFontUploadReq prepareBadgeFontUploadReq = new PrepareBadgeFontUploadReq(); // PrepareBadgeFontUploadReq | 
        try {
            APIprepareBadgeFontUploadRequest request = APIprepareBadgeFontUploadRequest.newBuilder()
                .prepareBadgeFontUploadReq(prepareBadgeFontUploadReq)
                .build();
            PrepareBadgeFontUploadResp result = apiInstance.prepareBadgeFontUpload(request);
            System.out.println(result);
        } catch (ApiException e) {
            System.err.println("Exception when calling AssetService#prepareBadgeFontUpload");
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
| prepareBadgeFontUploadRequest | [**APIprepareBadgeFontUploadRequest**](AssetService.md#APIprepareBadgeFontUploadRequest)|-|-|

### Return type

[**PrepareBadgeFontUploadResp**](PrepareBadgeFontUploadResp.md)


### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

### HTTP response details
| Status code | Description | Response headers |
|-------------|-------------|------------------|
| **200** | OK |  -  |

## prepareBadgeFontUploadWithHttpInfo

> ApiResponse<PrepareBadgeFontUploadResp> prepareBadgeFontUploadWithHttpInfo(prepareBadgeFontUploadRequest)



申请徽章 Maker 自定义字体的内容寻址直传凭证。

### Example

```java
// Import classes:
import com.bass.bbs.ApiClient;
import com.bass.bbs.ApiException;
import com.bass.bbs.ApiResponse;
import com.bass.bbs.Configuration;
import com.bass.bbs.models.*;
import com.bass.bbs.api.AssetService;
import com.bass.bbs.api.AssetService.*;

public class Example {
    public static void main(String[] args) {
        ApiClient defaultClient = Configuration.getDefaultApiClient();
        defaultClient.setBasePath("http://localhost");

        AssetService apiInstance = new AssetService(defaultClient);
        PrepareBadgeFontUploadReq prepareBadgeFontUploadReq = new PrepareBadgeFontUploadReq(); // PrepareBadgeFontUploadReq | 
        try {
            APIprepareBadgeFontUploadRequest request = APIprepareBadgeFontUploadRequest.newBuilder()
                .prepareBadgeFontUploadReq(prepareBadgeFontUploadReq)
                .build();
            ApiResponse<PrepareBadgeFontUploadResp> response = apiInstance.prepareBadgeFontUploadWithHttpInfo(request);
            System.out.println("Status code: " + response.getStatusCode());
            System.out.println("Response headers: " + response.getHeaders());
            System.out.println("Response body: " + response.getData());
        } catch (ApiException e) {
            System.err.println("Exception when calling AssetService#prepareBadgeFontUpload");
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
| prepareBadgeFontUploadRequest | [**APIprepareBadgeFontUploadRequest**](AssetService.md#APIprepareBadgeFontUploadRequest)|-|-|

### Return type

ApiResponse<[**PrepareBadgeFontUploadResp**](PrepareBadgeFontUploadResp.md)>


### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

### HTTP response details
| Status code | Description | Response headers |
|-------------|-------------|------------------|
| **200** | OK |  -  |


<a id="APIprepareBadgeFontUploadRequest"></a>
## APIprepareBadgeFontUploadRequest
### Properties

|     Name      |    Type       | Description   |     Notes    |
| ------------- | ------------- | ------------- | -------------|
| **prepareBadgeFontUploadReq** | [**PrepareBadgeFontUploadReq**](PrepareBadgeFontUploadReq.md) |  | |



## prepareBadgeImageUpload

> PrepareBadgeImageUploadResp prepareBadgeImageUpload(prepareBadgeImageUploadRequest)



申请徽章 Maker 图片的内容寻址直传凭证。

### Example

```java
// Import classes:
import com.bass.bbs.ApiClient;
import com.bass.bbs.ApiException;
import com.bass.bbs.Configuration;
import com.bass.bbs.models.*;
import com.bass.bbs.api.AssetService;
import com.bass.bbs.api.AssetService.*;

public class Example {
    public static void main(String[] args) {
        ApiClient defaultClient = Configuration.getDefaultApiClient();
        defaultClient.setBasePath("http://localhost");

        AssetService apiInstance = new AssetService(defaultClient);
        PrepareBadgeImageUploadReq prepareBadgeImageUploadReq = new PrepareBadgeImageUploadReq(); // PrepareBadgeImageUploadReq | 
        try {
            APIprepareBadgeImageUploadRequest request = APIprepareBadgeImageUploadRequest.newBuilder()
                .prepareBadgeImageUploadReq(prepareBadgeImageUploadReq)
                .build();
            PrepareBadgeImageUploadResp result = apiInstance.prepareBadgeImageUpload(request);
            System.out.println(result);
        } catch (ApiException e) {
            System.err.println("Exception when calling AssetService#prepareBadgeImageUpload");
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
| prepareBadgeImageUploadRequest | [**APIprepareBadgeImageUploadRequest**](AssetService.md#APIprepareBadgeImageUploadRequest)|-|-|

### Return type

[**PrepareBadgeImageUploadResp**](PrepareBadgeImageUploadResp.md)


### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

### HTTP response details
| Status code | Description | Response headers |
|-------------|-------------|------------------|
| **200** | OK |  -  |

## prepareBadgeImageUploadWithHttpInfo

> ApiResponse<PrepareBadgeImageUploadResp> prepareBadgeImageUploadWithHttpInfo(prepareBadgeImageUploadRequest)



申请徽章 Maker 图片的内容寻址直传凭证。

### Example

```java
// Import classes:
import com.bass.bbs.ApiClient;
import com.bass.bbs.ApiException;
import com.bass.bbs.ApiResponse;
import com.bass.bbs.Configuration;
import com.bass.bbs.models.*;
import com.bass.bbs.api.AssetService;
import com.bass.bbs.api.AssetService.*;

public class Example {
    public static void main(String[] args) {
        ApiClient defaultClient = Configuration.getDefaultApiClient();
        defaultClient.setBasePath("http://localhost");

        AssetService apiInstance = new AssetService(defaultClient);
        PrepareBadgeImageUploadReq prepareBadgeImageUploadReq = new PrepareBadgeImageUploadReq(); // PrepareBadgeImageUploadReq | 
        try {
            APIprepareBadgeImageUploadRequest request = APIprepareBadgeImageUploadRequest.newBuilder()
                .prepareBadgeImageUploadReq(prepareBadgeImageUploadReq)
                .build();
            ApiResponse<PrepareBadgeImageUploadResp> response = apiInstance.prepareBadgeImageUploadWithHttpInfo(request);
            System.out.println("Status code: " + response.getStatusCode());
            System.out.println("Response headers: " + response.getHeaders());
            System.out.println("Response body: " + response.getData());
        } catch (ApiException e) {
            System.err.println("Exception when calling AssetService#prepareBadgeImageUpload");
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
| prepareBadgeImageUploadRequest | [**APIprepareBadgeImageUploadRequest**](AssetService.md#APIprepareBadgeImageUploadRequest)|-|-|

### Return type

ApiResponse<[**PrepareBadgeImageUploadResp**](PrepareBadgeImageUploadResp.md)>


### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

### HTTP response details
| Status code | Description | Response headers |
|-------------|-------------|------------------|
| **200** | OK |  -  |


<a id="APIprepareBadgeImageUploadRequest"></a>
## APIprepareBadgeImageUploadRequest
### Properties

|     Name      |    Type       | Description   |     Notes    |
| ------------- | ------------- | ------------- | -------------|
| **prepareBadgeImageUploadReq** | [**PrepareBadgeImageUploadReq**](PrepareBadgeImageUploadReq.md) |  | |



## prepareMarkdownImageUpload

> PrepareMarkdownImageUploadResp prepareMarkdownImageUpload(prepareMarkdownImageUploadRequest)



申请 Markdown 图片的内容寻址直传凭证。

### Example

```java
// Import classes:
import com.bass.bbs.ApiClient;
import com.bass.bbs.ApiException;
import com.bass.bbs.Configuration;
import com.bass.bbs.models.*;
import com.bass.bbs.api.AssetService;
import com.bass.bbs.api.AssetService.*;

public class Example {
    public static void main(String[] args) {
        ApiClient defaultClient = Configuration.getDefaultApiClient();
        defaultClient.setBasePath("http://localhost");

        AssetService apiInstance = new AssetService(defaultClient);
        PrepareMarkdownImageUploadReq prepareMarkdownImageUploadReq = new PrepareMarkdownImageUploadReq(); // PrepareMarkdownImageUploadReq | 
        try {
            APIprepareMarkdownImageUploadRequest request = APIprepareMarkdownImageUploadRequest.newBuilder()
                .prepareMarkdownImageUploadReq(prepareMarkdownImageUploadReq)
                .build();
            PrepareMarkdownImageUploadResp result = apiInstance.prepareMarkdownImageUpload(request);
            System.out.println(result);
        } catch (ApiException e) {
            System.err.println("Exception when calling AssetService#prepareMarkdownImageUpload");
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
| prepareMarkdownImageUploadRequest | [**APIprepareMarkdownImageUploadRequest**](AssetService.md#APIprepareMarkdownImageUploadRequest)|-|-|

### Return type

[**PrepareMarkdownImageUploadResp**](PrepareMarkdownImageUploadResp.md)


### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

### HTTP response details
| Status code | Description | Response headers |
|-------------|-------------|------------------|
| **200** | OK |  -  |

## prepareMarkdownImageUploadWithHttpInfo

> ApiResponse<PrepareMarkdownImageUploadResp> prepareMarkdownImageUploadWithHttpInfo(prepareMarkdownImageUploadRequest)



申请 Markdown 图片的内容寻址直传凭证。

### Example

```java
// Import classes:
import com.bass.bbs.ApiClient;
import com.bass.bbs.ApiException;
import com.bass.bbs.ApiResponse;
import com.bass.bbs.Configuration;
import com.bass.bbs.models.*;
import com.bass.bbs.api.AssetService;
import com.bass.bbs.api.AssetService.*;

public class Example {
    public static void main(String[] args) {
        ApiClient defaultClient = Configuration.getDefaultApiClient();
        defaultClient.setBasePath("http://localhost");

        AssetService apiInstance = new AssetService(defaultClient);
        PrepareMarkdownImageUploadReq prepareMarkdownImageUploadReq = new PrepareMarkdownImageUploadReq(); // PrepareMarkdownImageUploadReq | 
        try {
            APIprepareMarkdownImageUploadRequest request = APIprepareMarkdownImageUploadRequest.newBuilder()
                .prepareMarkdownImageUploadReq(prepareMarkdownImageUploadReq)
                .build();
            ApiResponse<PrepareMarkdownImageUploadResp> response = apiInstance.prepareMarkdownImageUploadWithHttpInfo(request);
            System.out.println("Status code: " + response.getStatusCode());
            System.out.println("Response headers: " + response.getHeaders());
            System.out.println("Response body: " + response.getData());
        } catch (ApiException e) {
            System.err.println("Exception when calling AssetService#prepareMarkdownImageUpload");
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
| prepareMarkdownImageUploadRequest | [**APIprepareMarkdownImageUploadRequest**](AssetService.md#APIprepareMarkdownImageUploadRequest)|-|-|

### Return type

ApiResponse<[**PrepareMarkdownImageUploadResp**](PrepareMarkdownImageUploadResp.md)>


### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

### HTTP response details
| Status code | Description | Response headers |
|-------------|-------------|------------------|
| **200** | OK |  -  |


<a id="APIprepareMarkdownImageUploadRequest"></a>
## APIprepareMarkdownImageUploadRequest
### Properties

|     Name      |    Type       | Description   |     Notes    |
| ------------- | ------------- | ------------- | -------------|
| **prepareMarkdownImageUploadReq** | [**PrepareMarkdownImageUploadReq**](PrepareMarkdownImageUploadReq.md) |  | |


