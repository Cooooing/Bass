# PostscriptService

All URIs are relative to *http://localhost*

| Method | HTTP request | Description |
|------------- | ------------- | -------------|
| [**addPostscript**](PostscriptService.md#addPostscript) | **POST** /v1/content/postscript/add |  |
| [**addPostscriptWithHttpInfo**](PostscriptService.md#addPostscriptWithHttpInfo) | **POST** /v1/content/postscript/add |  |
| [**listPostscripts**](PostscriptService.md#listPostscripts) | **POST** /v1/content/postscript/list |  |
| [**listPostscriptsWithHttpInfo**](PostscriptService.md#listPostscriptsWithHttpInfo) | **POST** /v1/content/postscript/list |  |



## addPostscript

> AddPostscriptResp addPostscript(addPostscriptRequest)



添加文章附言。

### Example

```java
// Import classes:
import com.bass.bbs.ApiClient;
import com.bass.bbs.ApiException;
import com.bass.bbs.Configuration;
import com.bass.bbs.models.*;
import com.bass.bbs.api.PostscriptService;
import com.bass.bbs.api.PostscriptService.*;

public class Example {
    public static void main(String[] args) {
        ApiClient defaultClient = Configuration.getDefaultApiClient();
        defaultClient.setBasePath("http://localhost");

        PostscriptService apiInstance = new PostscriptService(defaultClient);
        AddPostscriptReq addPostscriptReq = new AddPostscriptReq(); // AddPostscriptReq | 
        try {
            APIaddPostscriptRequest request = APIaddPostscriptRequest.newBuilder()
                .addPostscriptReq(addPostscriptReq)
                .build();
            AddPostscriptResp result = apiInstance.addPostscript(request);
            System.out.println(result);
        } catch (ApiException e) {
            System.err.println("Exception when calling PostscriptService#addPostscript");
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
| addPostscriptRequest | [**APIaddPostscriptRequest**](PostscriptService.md#APIaddPostscriptRequest)|-|-|

### Return type

[**AddPostscriptResp**](AddPostscriptResp.md)


### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

### HTTP response details
| Status code | Description | Response headers |
|-------------|-------------|------------------|
| **200** | OK |  -  |

## addPostscriptWithHttpInfo

> ApiResponse<AddPostscriptResp> addPostscriptWithHttpInfo(addPostscriptRequest)



添加文章附言。

### Example

```java
// Import classes:
import com.bass.bbs.ApiClient;
import com.bass.bbs.ApiException;
import com.bass.bbs.ApiResponse;
import com.bass.bbs.Configuration;
import com.bass.bbs.models.*;
import com.bass.bbs.api.PostscriptService;
import com.bass.bbs.api.PostscriptService.*;

public class Example {
    public static void main(String[] args) {
        ApiClient defaultClient = Configuration.getDefaultApiClient();
        defaultClient.setBasePath("http://localhost");

        PostscriptService apiInstance = new PostscriptService(defaultClient);
        AddPostscriptReq addPostscriptReq = new AddPostscriptReq(); // AddPostscriptReq | 
        try {
            APIaddPostscriptRequest request = APIaddPostscriptRequest.newBuilder()
                .addPostscriptReq(addPostscriptReq)
                .build();
            ApiResponse<AddPostscriptResp> response = apiInstance.addPostscriptWithHttpInfo(request);
            System.out.println("Status code: " + response.getStatusCode());
            System.out.println("Response headers: " + response.getHeaders());
            System.out.println("Response body: " + response.getData());
        } catch (ApiException e) {
            System.err.println("Exception when calling PostscriptService#addPostscript");
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
| addPostscriptRequest | [**APIaddPostscriptRequest**](PostscriptService.md#APIaddPostscriptRequest)|-|-|

### Return type

ApiResponse<[**AddPostscriptResp**](AddPostscriptResp.md)>


### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

### HTTP response details
| Status code | Description | Response headers |
|-------------|-------------|------------------|
| **200** | OK |  -  |


<a id="APIaddPostscriptRequest"></a>
## APIaddPostscriptRequest
### Properties

|     Name      |    Type       | Description   |     Notes    |
| ------------- | ------------- | ------------- | -------------|
| **addPostscriptReq** | [**AddPostscriptReq**](AddPostscriptReq.md) |  | |



## listPostscripts

> ListPostscriptsResp listPostscripts(listPostscriptsRequest)



查询文章附言列表。

### Example

```java
// Import classes:
import com.bass.bbs.ApiClient;
import com.bass.bbs.ApiException;
import com.bass.bbs.Configuration;
import com.bass.bbs.models.*;
import com.bass.bbs.api.PostscriptService;
import com.bass.bbs.api.PostscriptService.*;

public class Example {
    public static void main(String[] args) {
        ApiClient defaultClient = Configuration.getDefaultApiClient();
        defaultClient.setBasePath("http://localhost");

        PostscriptService apiInstance = new PostscriptService(defaultClient);
        ListPostscriptsReq listPostscriptsReq = new ListPostscriptsReq(); // ListPostscriptsReq | 
        try {
            APIlistPostscriptsRequest request = APIlistPostscriptsRequest.newBuilder()
                .listPostscriptsReq(listPostscriptsReq)
                .build();
            ListPostscriptsResp result = apiInstance.listPostscripts(request);
            System.out.println(result);
        } catch (ApiException e) {
            System.err.println("Exception when calling PostscriptService#listPostscripts");
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
| listPostscriptsRequest | [**APIlistPostscriptsRequest**](PostscriptService.md#APIlistPostscriptsRequest)|-|-|

### Return type

[**ListPostscriptsResp**](ListPostscriptsResp.md)


### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

### HTTP response details
| Status code | Description | Response headers |
|-------------|-------------|------------------|
| **200** | OK |  -  |

## listPostscriptsWithHttpInfo

> ApiResponse<ListPostscriptsResp> listPostscriptsWithHttpInfo(listPostscriptsRequest)



查询文章附言列表。

### Example

```java
// Import classes:
import com.bass.bbs.ApiClient;
import com.bass.bbs.ApiException;
import com.bass.bbs.ApiResponse;
import com.bass.bbs.Configuration;
import com.bass.bbs.models.*;
import com.bass.bbs.api.PostscriptService;
import com.bass.bbs.api.PostscriptService.*;

public class Example {
    public static void main(String[] args) {
        ApiClient defaultClient = Configuration.getDefaultApiClient();
        defaultClient.setBasePath("http://localhost");

        PostscriptService apiInstance = new PostscriptService(defaultClient);
        ListPostscriptsReq listPostscriptsReq = new ListPostscriptsReq(); // ListPostscriptsReq | 
        try {
            APIlistPostscriptsRequest request = APIlistPostscriptsRequest.newBuilder()
                .listPostscriptsReq(listPostscriptsReq)
                .build();
            ApiResponse<ListPostscriptsResp> response = apiInstance.listPostscriptsWithHttpInfo(request);
            System.out.println("Status code: " + response.getStatusCode());
            System.out.println("Response headers: " + response.getHeaders());
            System.out.println("Response body: " + response.getData());
        } catch (ApiException e) {
            System.err.println("Exception when calling PostscriptService#listPostscripts");
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
| listPostscriptsRequest | [**APIlistPostscriptsRequest**](PostscriptService.md#APIlistPostscriptsRequest)|-|-|

### Return type

ApiResponse<[**ListPostscriptsResp**](ListPostscriptsResp.md)>


### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

### HTTP response details
| Status code | Description | Response headers |
|-------------|-------------|------------------|
| **200** | OK |  -  |


<a id="APIlistPostscriptsRequest"></a>
## APIlistPostscriptsRequest
### Properties

|     Name      |    Type       | Description   |     Notes    |
| ------------- | ------------- | ------------- | -------------|
| **listPostscriptsReq** | [**ListPostscriptsReq**](ListPostscriptsReq.md) |  | |


