# AssetService

All URIs are relative to *http://localhost*

| Method | HTTP request | Description |
|------------- | ------------- | -------------|
| [**completeMarkdownImageUpload**](AssetService.md#completeMarkdownImageUpload) | **POST** /v1/asset/markdown-image/complete |  |
| [**completeMarkdownImageUploadWithHttpInfo**](AssetService.md#completeMarkdownImageUploadWithHttpInfo) | **POST** /v1/asset/markdown-image/complete |  |
| [**prepareMarkdownImageUpload**](AssetService.md#prepareMarkdownImageUpload) | **POST** /v1/asset/markdown-image/prepare |  |
| [**prepareMarkdownImageUploadWithHttpInfo**](AssetService.md#prepareMarkdownImageUploadWithHttpInfo) | **POST** /v1/asset/markdown-image/prepare |  |



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


